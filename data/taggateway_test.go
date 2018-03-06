
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
	_ "github.com/mattn/go-sqlite3"
)


func TestTagGateway(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)
	gateway := TagGateway{}

	//
	// Test that a tag can be added to the data source
	//

	actualId, err := gateway.AddTag("name", test.GoPathTestFolder)
	var expectedId = 1

	if err != nil {
		t.Fatalf("AddTag encountered an error: %s", err)
	}

	if actualId != expectedId {
		t.Fatalf("Expected %d but got %d", expectedId, actualId)
	}

	//
	// Test that a tag can be read from the data source
	//

	tag, err := gateway.ReadTag(actualId, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("RemoveTag encountered an error: %s", err)
	}

	var expectedName = "name"

	if tag.Id != actualId {
		t.Fatalf("Expected %d but got %d", actualId, tag.Id)
	}

	if tag.Name != expectedName {
		t.Fatalf("Expected %s but got %s", expectedName, tag.Name)
	}

	//
	// test that a tag can be read by name
	//

	id, err := gateway.GetTagIdByName(tag.Name, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("ReadTagByName encountered an error: %s", err)
	}

	if id != tag.Id {
		t.Fatalf("Wrong tag loaded: %d", id)
	}

	//
	// Test that a tag can be edited
	//

	tag.Name = "Name2"

	err = gateway.SaveTag(tag.Id, tag.Name, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if tag.Name != "Name2" {
		t.Fatalf("Expected %s but got %s", "Name2", tag.Name)
	}

	//
	// Test that a tag can be deleted from the data source
	//

	err = gateway.RemoveTag(actualId, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("RemoveTag encountered an error: %s", err)
	}

	//
	// Test that trying to find a tag that doesn't exist by it's name returns the appropriate error
	//

	_, err = gateway.GetTagIdByName("DoesNotExist", test.GoPathTestFolder)

	if _, ok := err.(*TagNotFoundError); !ok {
		t.Fatalf("GetTagIdByName should encounter an error: '%s'", err)
	}
}