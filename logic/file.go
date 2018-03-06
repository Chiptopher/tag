
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
	"fmt"
)

// NewFile creates a pointer to a File.
func NewFile(id int, fileName string, tags []*Tag) (*File) {
	entityState := entityStateClean
	if id == 0 {
		entityState = entityStateNew
	}

	return  &File{id, fileName, tags, tags, entityState}
}

// CreateFile a pointer to a File that hasn't existed in the system before.
func CreateFile(fileName string, tags []*Tag) (*File) {
	return NewFile(0, fileName, tags)
}

// File is the structure that contains information about a file.
type File struct {
	id           int
	fileName     string
	tags         []*Tag
	originalTags []*Tag
	entityState  entityState
}

// Id reports the id of this File.
func (file *File) Id() int {
	return file.id
}

// FileName reports the name of this File.
func (file *File) FileName() string {
	return file.fileName
}

// SetFileName sets the name of this File.
func (file *File) SetFileName(name string) {
	file.setEntityState(entityStateDirty)
	file.fileName = name
}

// Tags gets the Tags associated with this Files.
func (file *File) Tags() []*Tag {
	return file.tags
}

// AddFile adds the given Tag to this File.
func (file *File) AddTag(tag *Tag) {

	if !inTagSlice(tag, file.tags) {
		file.tags = append(file.tags, tag)
		if file.getEntityState() != entityStateNew {
			file.setEntityState(entityStateDirty)
		}
	}

}

// RemoveTag removes the given Tag from this File.
func (file *File) RemoveTag(tag *Tag) {
	pos := -1
	for i, t := range file.tags {
		if tag == t {
			pos = i
		}
	}

	if pos != -1 {
		file.tags = append(file.tags[:pos], file.tags[pos+1:]...)
		file.setEntityState(entityStateDirty)
	}
}

// HasTag reports whether the given Tag is in this File.
func (file File) HasTag(tag *Tag) bool {
	for _, fTag := range file.tags {
		if tag == fTag {
			return true
		}
	}
	return false
}

// Copy copies the content of this File to a new File.
func (file File) Copy() *File {

	newTags := make([]*Tag, len(file.tags), cap(file.tags))
	for i, tag := range file.tags {
		newTags[i] = tag.Copy()
	}
	return NewFile(file.id, file.fileName, newTags)
}

// ToString writes this File to a string.
func (file File) ToString() string {
	tags := ""
	for i:= 0; i < len(file.tags); i++ {
		tags = tags + file.tags[i].ToString()
		if i < len(file.tags) - 1 {
			tags = tags + ", "
		}
	}

	return fmt.Sprintf("File: %d, %s, [ %s ]", file.id, file.fileName, tags)
}

func inTagSlice(tag *Tag, tags []*Tag) (bool) {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (file File) getEntityState() (entityState) {
	return file.entityState
}

func (file *File) setEntityState(entityState entityState) {
	file.entityState = entityState
}

// addedTags gets the Tags added to this File since it's original
// creation.
func (file File) addedTags() ([]*Tag) {
	result := []*Tag{}

	for _, tag := range file.tags {
		if !inTagSlice(tag, file.originalTags) {
			result = append(result, tag)
		}
	}

	return result
}

// removedTags gets the Tags added to this File since it's original
// creation.
func (file File) removedTags() ([]*Tag) {
	result := []*Tag{}

	for _, tag := range file.originalTags {
		if !inTagSlice(tag, file.tags) {
			result = append(result, tag)
		}
	}

	return result
}
