
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

import "fmt"

// NewTag creates a pointer to a Tag.
func NewTag(id int, name string) (*Tag) {
	state := entityStateClean
	if id == 0 {
		state = entityStateNew
	}
	return &Tag{id, name, state}
}

// CreateTag creates a pointer to a Tag that hasn't existed in the system
// before.
func CreateTag(name string) (*Tag) {
	return NewTag(0, name)
}

// Tag is the structure that contains information about a tag.
type Tag struct {
	id          int
	name        string
	entityState entityState
}

// Id reports the id of this Tag.
func (tag *Tag) Id() int {
	return tag.id
}

// Name reports the name of this Tag.
func (tag *Tag) Name() string {
	return tag.name
}

// SetName sets the name of this Tag.
func (tag *Tag) SetName(name string) {
	tag.setEntityState(entityStateDirty)
	tag.name = name
}

// ToString formats this Tag to a string.
func (tag Tag) ToString() string {
	return fmt.Sprintf("Tag: %d, %s", tag.id, tag.name)
}

// Copy copies the contents of this Tag to a new Tag.
func (tag Tag) Copy() (*Tag) {
	return NewTag(tag.id, tag.name)
}

func (tag *Tag) setEntityState(entityState entityState) {
	tag.entityState = entityState
}

func (tag Tag) getEntityState() (entityState) {
	return tag.entityState
}