
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
	"os"
	"github.com/Chiptopher/tag/test"
)

func TestCreateDataFolder(t *testing.T) {

	result, err := CreateDataFolder(test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Create data folder failed: %s", err)
	}

	if !result {
		t.Fatalf("Create data folder failed to create the folder but did not error.")
	}

	if _, err := os.Stat(test.CreatePathInTestFolder(DataFolderName)); os.IsNotExist(err) {
		t.Fatalf("Folder not created: %s", err)
	}

	if _,err = os.Stat(test.CreatePathInTestFolder(DataFolderName)+"/"+ DatabaseName); os.IsNotExist(err) {
		t.Fatalf("Did not create the data file.")
	}

	os.RemoveAll(test.CreatePathInTestFolder(DataFolderName))
}

func TestCreateDataFolder_ThrowsPreexistingErrorWhenFolderAlreadyExists(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)

	result, err := CreateDataFolder(test.GoPathTestFolder)

	if result {
		t.Fatalf("Create data folder should fail to create a folder.")
	}

	if _, ok := err.(*PreexistingError); !ok {
		t.Fatalf("Create data folder should return a preexisting error.")
	}

	os.RemoveAll(test.CreatePathInTestFolder(DataFolderName))
}
