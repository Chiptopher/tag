
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
	"reflect"
)

func TestAddRemoveAndReadTagListEntity(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)
	gateway := TagListEntityGateway{}

	addResultId, err := gateway.AddTagListEntity(1, 1, test.GoPathTestFolder)
	expectedId := 1

	if err != nil {
		t.Fatalf("AddTagListEntity encountered an error: %s", err)
	}

	if addResultId != expectedId {
		t.Fatalf("Expected '%d' but got '%d'", expectedId, addResultId)
	}
}

func TestTagListEntityGateway_ReadTags(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)
	gateway := TagListEntityGateway{}
	id1, err := gateway.AddTagListEntity(3, 3, test.GoPathTestFolder)
	id2, err := gateway.AddTagListEntity(3, 4, test.GoPathTestFolder)
	_, err = gateway.AddTagListEntity(4, 3, test.GoPathTestFolder)

	result , err := gateway.ReadTags(3, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("AddTagListEntity encountered an error: %s", err)
	}


	expectedResult := []*TagListEntityDTO{ {id1, 3, 3}, {id2, 3, 4},}

	for i, element := range expectedResult {

		if element.Id != result[i].Id {
			t.Fatalf("Exprected %d but got %d", element.Id, result[i].Id)
		}

		if element.FileId != result[i].FileId {
			t.Fatalf("Exprected %d but got %d", element.FileId, result[i].FileId)
		}

		if element.TagId != result[i].TagId {
			t.Fatalf("Exprected %d but got %d", element.TagId, result[i].TagId)
		}
	}

}

func TestTagListEntityGateway_ReadFiles(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)
	gateway := TagListEntityGateway{}
	id1, err := gateway.AddTagListEntity(5, 5, test.GoPathTestFolder)
	_, err = gateway.AddTagListEntity(5, 6, test.GoPathTestFolder)
	id2, err := gateway.AddTagListEntity(6, 5, test.GoPathTestFolder)

	result, err := gateway.ReadFiles(5, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("AddTagListEntity encountered an error: %s", err)
	}

	expectedResult := []TagListEntityDTO{ {id1, 5, 5}, {id2, 6, 5},}

	for i, element := range expectedResult {
		if element != result[i] {
			t.Fatalf("Expected %p but got %p", expectedResult, result)
		}
	}
}

func TestTagListEntityGateway_RemoveTagListEntity(t *testing.T) {

	_, err := CreateDataFolder(test.GoPathTestFolder)

	gateway := TagListEntityGateway{}
	gateway.AddTagListEntity(7, 7, test.GoPathTestFolder)
	id, err := gateway.AddTagListEntity(7, 8, test.GoPathTestFolder)
	gateway.AddTagListEntity(8, 7, test.GoPathTestFolder)

	err = gateway.RemoveTagListEntity(7, 7, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("RemoveTagListEntity encountered an error: %s", err)
	}

	result , err := gateway.ReadTags(7, test.GoPathTestFolder)

	if err != nil {
		t.Fatalf("AddTagListEntity encountered an error: %s", err)
	}

	// since one thing was removed, it should be one item smaller.
	expectedResult := []*TagListEntityDTO{{id, 7, 8}}

	for i, element := range expectedResult {

		if !reflect.DeepEqual(element, result[i]) {
			t.Fatalf("Expected %p but got %p", element, result[i])
		}
	}
}