
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
	"testing"
	"fmt"
)

func TestNewFile(t *testing.T) {
	tag := &Tag{}
	file := NewFile(1, "name", []*Tag{tag})

	if file.id != 1 {
		t.Fatalf("Expected id %d but got %d", 1, file.id)
	}

	if file.fileName != "name" {
		t.Fatalf("Expected %s but got %s", "name", file.fileName)
	}

	if len(file.tags) != 1 || file.tags[0] != tag {
		t.Fatalf("Incorrectly loaded tags: %v", file.tags)
	}

	if len(file.originalTags) != 1 || file.originalTags[0] != tag {
		t.Fatalf("Incorrectly loaded tags: %v", file.tags)
	}
}

func TestFile_AddTag(t *testing.T) {
	file := NewFile(0, "fileName", []*Tag{})
	tag := NewTag(0, "TagName")
	file.AddTag(tag)

	if len(file.tags) != 1 || file.tags[0] != tag {
		t.Fatalf("Tag not added to file: %p", file)
	}
}


func TestFile_HasTag(t *testing.T) {
	tag := NewTag(1, "name")
	file := NewFile(1, "name", []*Tag{tag})

	if !file.HasTag(tag) {
		t.Fatalf("Excpected '%t', but got '%t'", true, false)
	}

	if file.HasTag(&Tag{}) {
		t.Fatalf("Excpected '%t', but got '%t'", false, true)
	}
}

func TestFileToString(t *testing.T) {

	tag := NewTag(1, "name")
	actual := NewFile(1, "name", []*Tag{tag}).ToString()
	expected := fmt.Sprintf("File: 1, name, [ %s ]", tag.ToString())

	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestFileToStringMultipleTags(t *testing.T) {

	tag1 := NewTag(1, "name")
	tag2 := NewTag(1, "name")

	actual := NewFile(1, "name", []*Tag{tag1, tag2}).ToString()
	expected := fmt.Sprintf("File: 1, name, [ %s, %s ]", tag1.ToString(), tag2.ToString())

	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestFile_AddedTags(t *testing.T) {
	file := NewFile(1, "name", []*Tag{})
	tag := NewTag(1, "name")

	file.AddTag(tag)

	addedTags := file.addedTags()

	if len(addedTags) != 1 || addedTags[0] != tag {
		t.Fatalf("Incorrectly loaded list: %v", addedTags)
	}
}

func TestFile_RemovedTags(t *testing.T) {

	tag := NewTag(1, "name")
	file := NewFile(1, "name", []*Tag{tag})

	file.RemoveTag(tag)

	removedTags := file.removedTags()

	if len(removedTags) != 1  {
		t.Fatalf("Incorrectly loaded list: %v", removedTags)
	}

	if removedTags[0] != tag {
		t.Fatalf("Incorrectly loaded tag: %p", removedTags[0])
	}
}

func TestFile_RemoveTag(t *testing.T) {
	file := NewFile(1, "name", []*Tag{})
	tag := &Tag{}

	file.AddTag(tag)
	file.RemoveTag(tag)

	if len(file.tags) != 0 {
		t.Fatalf("Tag not removed: %v", file.tags)
	}
}

func TestFile_RemoveTag_NonExistentValue(t *testing.T) {
	file := NewFile(1, "name", []*Tag{})
	tag := &Tag{}

	file.RemoveTag(tag)

	if len(file.tags) != 0 {
		t.Fatalf("Tag not removed: %v", file.tags)
	}
}

func TestFile_RemoveTag_MultipleValues(t *testing.T) {
	file := NewFile(1, "name", []*Tag{})
	tag1 := &Tag{}
	tag2 := &Tag{}

	file.AddTag(tag1)
	file.AddTag(tag2)
	file.RemoveTag(tag1)

	if len(file.tags) != 1 {
		t.Fatalf("Tag not removed: %v", file.tags)
	}
}

func TestFile_AddTag_CannotAddDuplicates(t *testing.T) {

	file := NewFile(0, "fileName", []*Tag{})
	tag := NewTag(0, "TagName")
	file.AddTag(tag)
	file.AddTag(tag)

	if len(file.tags) != 1 {
		t.Fatalf("Tag added to file that shouldn't have been")
	}
}

func TestFile_Copy(t *testing.T) {
	tag := NewTag(1, "name")
	file := NewFile(1, "fileName", []*Tag{tag})

	result := file.Copy()

	if result == file {
		t.Fatalf("Should not reference same thing.")
	}

	if file.tags[0] == result.tags[0] {
		t.Fatalf("Should copy tags too.")
	}

	if result.id != file.id {
		t.Fatalf("Expected %d but got %d", file.id, result.id)
	}

	if result.fileName != file.fileName {
		t.Fatalf("Expected '%s' but got '%s'", file.fileName, result.fileName)
	}
}

func TestFile_SetFileName_ChangesEntityState(t *testing.T) {
	file := NewFile(1, "name", []*Tag{})

	file.SetFileName("Name2")

	if file.getEntityState() != entityStateDirty {
		t.Fatalf("entity state not changed.")
	}
}

func TestFile_AddTag_ChangesEntityState(t *testing.T) {
	file := NewFile(1, "name", []*Tag{})
	file.AddTag(NewTag(1, "name"))
	if file.getEntityState() != entityStateDirty {
		t.Fatalf("Adding a tag doesn't change the entity state.")
	}
}

func TestFile_AddTag_StateStaysNewWhenChanged(t *testing.T) {

	file := CreateFile("name", []*Tag{})
	file.AddTag(NewTag(1, "name"))
	if file.getEntityState() != entityStateNew {
		t.Fatalf("Adding a tag shouldn't change the entity state when it's clean.")
	}
}

func TestFile_RemoveTag_SetsFileAsDirty(t *testing.T) {

	tag := NewTag(1, "name")
	file := NewFile(1, "name", []*Tag{tag})

	file.RemoveTag(tag)

	if file.getEntityState() != entityStateDirty {
		t.Fatalf("entity state not changed.")
	}
}