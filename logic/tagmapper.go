
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

import "github.com/Chiptopher/tag/data"

// TagMapperI is the interface for mapping tag data into a Tag.
type TagMapperI interface {
	SaveTag(tag *Tag) (int, error)
	getTag(id int) (*Tag, error)
	RemoveTag(id int) (error)
	getTagByName(name string) (*Tag, error)
}

// NewTagMapper creates a pointer to a TagMapper
func NewTagMapper(filePath string) (*TagMapper) {
	return &TagMapper{data.NewTagGateway(), filePath}
}

// TagMapper is the implementation of TagMapperI that provides functionality
// for mapping changes to Tags.
type TagMapper struct {
	tagGateway data.TagGatewayI
	filePath string
}

// getTag maps a Tag, and returns it and errors encountered.
func (m TagMapper) getTag(id int) (*Tag, error) {

	result, err := m.tagGateway.ReadTag(id, m.filePath)

	if err != nil {
		return nil, err
	}


	return NewTag(result.Id, result.Name), nil
}

// SaveTag saves changes to the given Tag. SaveTag returns the id of the
// saved Tag and any errors encountered.
func (m TagMapper) SaveTag(tag *Tag) (int, error) {

	if tag.getEntityState() == entityStateDirty {
		err := m.tagGateway.SaveTag(tag.id, tag.name, m.filePath)
		if err != nil {
			return 0, err
		}
	} else if tag.getEntityState() == entityStateNew {
		result, err := m.tagGateway.AddTag(tag.name, m.filePath)

		if err != nil {
			return 0, err
		}

		tag.id = result

		return result, nil
	}

	// Don't save tags where nothing has changed.
	return tag.id, nil
}

// RemoveTag removes the Tag with the given id, and returns any errors
// encountered.
func (m TagMapper) RemoveTag(id int) (error) {

	err := m.tagGateway.RemoveTag(id, m.filePath)

	return err
}

// getTagByName maps a Tag, and returns it and errors encountered.
func (m TagMapper) getTagByName(name string) (*Tag, error) {

	id, err := m.tagGateway.GetTagIdByName(name, m.filePath)

	if err != nil {
		return nil, err
	}

	return m.getTag(id)
}