
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

// TagIdentityMapI is the interface for accessing Tag objects currently loaded.
type TagIdentityMapI interface {
	GetTagById(id int) (*Tag, error)
	GetTagByName(name string) (*Tag, error)
	GetTagMapper() (TagMapperI)
}

// NewTagIdentityMap creates a pointer to a TagIdentityMap
func NewTagIdentityMap(tagMapper TagMapperI) (*TagIdentityMap) {
	return &TagIdentityMap{tagMapper, make(map[int]*Tag)}
}

// TagIdentityMap is the implementation of TagIdentitMapI that provides access
// to the Tag domain objects.
type TagIdentityMap struct {
	tagMapper TagMapperI
	tags map[int]*Tag
}

// GetTagById returns a Tag from the map, and loads it if it's not already been
// loaded. Any errors encountered are also returned.
func (t TagIdentityMap) GetTagById(id int) (*Tag, error) {

	tag := t.tags[id]
	var err error

	if tag == nil {
		tag, err = t.tagMapper.getTag(id)
		t.tags[id] = tag
	}

	return tag, err
}

// GetTagByName returns a Tag from the map, and loads it if it's not already been
// loaded. Any errors encountered are also returned.
func (t TagIdentityMap) GetTagByName(name string) (*Tag, error) {
	id := -1
	for _, tag := range t.tags {
		if tag.name == name {
			id = tag.id
		}
	}

	if id == -1 {
		tag, err := t.GetTagMapper().getTagByName(name)
		if err != nil {
			return nil, err
		}
		t.tags[tag.id] = tag
		return tag, err
	} else {
		return t.tags[id], nil
	}
}

// GetTagMapper gets the TagMapper associated with this TagIdentityMap.
func (t TagIdentityMap) GetTagMapper() (TagMapperI) {
	return t.tagMapper
}