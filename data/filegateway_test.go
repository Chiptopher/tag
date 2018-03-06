
// MIT License
//
// Copyright (c) 2018 Christopher M. Boyer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package data

import (
	"testing"
	"github.com/Chiptopher/tag/test"
	"os"
	"path/filepath"
)

// TestFileGateway tests all of the functionality of the file gateway. Currently,
// they are all being ran in one tests because they cannot be made independent of
// each other.
func TestFileGateway(t *testing.T) {

	_, err := CreateDataFolder(test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Encountered error: %s", err)
	}

	_, err = os.Create(filepath.Join(test.GoPathTestFolder, "name"))

	if err != nil {
		t.Fatalf("Encountrered error: %s", err)
	}

	gateway := FileGateway{}

	//
	// test adding a file.
	//

	actualFileName := filepath.Join(test.GoPathTestFolder, "name")
	actualId, err := gateway.AddFile(actualFileName, test.GoPathTestFolder)

	var expectedId = 1

	if err != nil {
		t.Fatalf("AddFile errored: '%s'", err)
	}

	if actualId != expectedId {
		t.Fatalf("Expected '%d' but got '%d'", expectedId, actualId)
	}

	//
	// test reading a file by it's id
	//

	file, err := gateway.ReadFile(actualId, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("ReadFile encountered an error")
	}

	if file.Id != actualId {
		t.Fatalf("expected '%d' but got '%d'", actualId, file.Id)
	}

	if file.Id != actualId {
		t.Fatalf("expected '%s' but got '%s'", "name", file.Name)
	}

	//
	// test reading a file by it's name
	//

	id, err := gateway.GetFileIdByName(file.Name, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Error ecnountered: %s", err)
	}

	if id != file.Id {
		t.Fatalf("Incorrect file loaded: %d", id)
	}

	//
	// test edit file
	//

	file.Name = "Name2"

	err = gateway.SaveFile(file.Id, file.Name, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Encountrered error: %s", err)
	}

	file2, err := gateway.ReadFile(actualId, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Encountrered error: %s", err)
	}

	if file2.Name != "Name2" {
		t.Fatalf("Expected file name of '%s', but got '%s'", "Name2", file2.Name)
	}

	//
	// test removing file by id
	//

	err = gateway.RemoveFile(actualId, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("RemoveTag encountered an error: %s", err)
	}

	file, err = gateway.ReadFile(actualId, test.GoPathTestFolder)

	if err == nil {
		t.Fatalf("Error thrown when searching for file that doesn't exist.")
	}

	//
	// test reading a file that doesn't exist by it's name causes an error
	//

	_, err = gateway.GetFileIdByName("DoesNotExist", test.GoPathTestFolder)

	if _, ok := err.(*FileNotFoundError); !ok {
		t.Fatalf("Should encounter a FileNotFoundError: '%s'", err)
	}

	//
	// test get all the ids of files in the system
	//

	os.Create(filepath.Join(test.GoPathTestFolder, "Name1"))
	os.Create(filepath.Join(test.GoPathTestFolder, "Name2"))

	fileName1 := filepath.Join(test.GoPathTestFolder, "Name1")
	fileName2 := filepath.Join(test.GoPathTestFolder, "Name2")

	gateway.AddFile(fileName1, test.GoPathTestFolder)
	gateway.AddFile(fileName2, test.GoPathTestFolder)

	ids, err := gateway.GetAllFileIds(test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if len(ids) != 2 {
		t.Fatalf("Incorrect number of ids: %d", len(ids))
	}

	if ids[0] != 1 {
		t.Fatalf("id should be '%d'", 2)
	}

	if ids[1] != 2 {
		t.Fatalf("id should be '%d'", 3)
	}

	//
	// Trying to add a file to the gateway that doesn't exist deletes it
	//

	_, err = gateway.AddFile("FileThatDoesNotExist", test.GoPathTestFolder)
	if _, ok := err.(*NoFileExistsError); ok {

	}
}