
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

package logic

import (
	"github.com/Chiptopher/tag/data"
	"testing"
)

func TestTagMapper_saveTag(t *testing.T) {

	expectedId := 1
	tagName := "name"
	tag := NewTag(0, tagName)
	testObj := data.MockNewTagGateway()
	testObj.On("AddTag", tagName, "path").Return(expectedId, nil)

	mapper := NewTagMapper("path")
	mapper.tagGateway = testObj

	result, err := mapper.SaveTag(tag)

	if err != nil {
		t.Fatalf("saveTag errored: %s", err)
	}

	if result != expectedId {
		t.Fatalf("Expected %d but got %d", expectedId, result)
	}

	if tag.id != result {
		t.Fatalf("Expected %d but got %d", result, tag.id)
	}
}

func TestTagMapper_saveTag_UpdateTag(t *testing.T) {

	filePath := "Path"

	id := 1
	tagName := "name"

	tag := NewTag(id, tagName)

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("SaveTag", id, tagName, filePath).Return(nil)

	mapper := NewTagMapper(filePath)
	mapper.tagGateway = tagGateway

	result, err := mapper.SaveTag(tag)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if result != id {
		t.Fatalf("Expected %d but got %d", id, result)
	}
}

func TestTagMapper_readTag(t *testing.T) {

	expectedId := 1
	testObj := data.MockNewTagGateway()
	testObj.On("ReadTag", expectedId, "path").Return(&data.TagDataDTO{expectedId, "name"})

	mapper := NewTagMapper("path")
	mapper.tagGateway = testObj
	result, err := mapper.getTag(expectedId)

	if err != nil {
		t.Fatalf("saveTag errored: %s", err)
	}

	if result.id != expectedId {
		t.Fatalf("Expected %d but got %d", expectedId, result.id)
	}

	if result.name != "name" {
		t.Fatalf("Expected '%s' but got '%s'", "name", result.name)
	}
}

func TestTagMapper_removeTag(t *testing.T) {

	expectedId := 1
	path := "Path"
	testObj := data.MockNewTagGateway()
	testObj.On("RemoveTag", expectedId, path).Return(nil)

	mapper := NewTagMapper(path)
	mapper.tagGateway = testObj
	err := mapper.RemoveTag(expectedId)

	if err != nil {
		t.Fatalf("removeTag encountered an error '%s'", err)
	}
}

func TestTagMapper_readTagByName(t *testing.T) {

	expectedId := 1
	tagName := "name"
	testObj := data.MockNewTagGateway()
	testObj.On("ReadTag", expectedId, "path").Return(&data.TagDataDTO{expectedId, tagName})
	testObj.On("GetTagIdByName", tagName, "path").Return(expectedId, nil)

	mapper := NewTagMapper("path")
	mapper.tagGateway = testObj
	result, err := mapper.getTagByName(tagName)

	if err != nil {
		t.Fatalf("saveTag errored: %s", err)
	}

	if result.id != expectedId {
		t.Fatalf("Expected %d but got %d", expectedId, result.id)
	}

	if result.name != "name" {
		t.Fatalf("Expected '%s' but got '%s'", "name", result.name)
	}
}

func TestTagMapper_saveTag_OnlySavesExistingTagsWhenTheyHaveChanged(t *testing.T) {

	tagName := "name"
	newTagName := "NewName"
	filePath := "path"

	mapper := NewTagMapper(filePath)

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("SaveTag", 1, newTagName, filePath).Return(nil)

	mapper.tagGateway = tagGateway

	tag := NewTag(1, tagName)
	tag.SetName(newTagName)

	_, err := mapper.SaveTag(tag)

	tagGateway.AssertCalled(t, "SaveTag", 1, newTagName, filePath)

	if err != nil {
		t.Fatalf("Error encountered; %s", err)
	}
}

func TestTagMapper_saveTag_DoesNotSaveWhenTagHasNoChanges(t *testing.T) {

	tagMapper := NewTagMapper("filePath")
	tagGateway := data.MockNewTagGateway()
	tag := NewTag(1, "name")
	tagMapper.tagGateway = tagGateway
	tagMapper.SaveTag(tag)

	tagGateway.AssertNotCalled(t, "SaveTag")
}