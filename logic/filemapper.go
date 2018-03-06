
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
)

// FileMapperI is the interface for mapping file data into a File.
type FileMapperI interface {
	getFile(id int) (*File, error)
	SaveFile(file *File) (int, error)
	RemoveFile(id int) (error)
	getFileByName(name string) (*File, error)
	getCWD() (string)
}

// NewFileMapper creates a pointer to a FileMapper
func NewFileMapper(tagIdentityMap TagIdentityMapI, filePath string) (*FileMapper) {
	return &FileMapper{data.NewFileGateway(),
	data.NewTagListEntityGateway(),
	tagIdentityMap.GetTagMapper(),
	tagIdentityMap,
	filePath,
	}
}

// FileMapper is the implementation of FileMapper that provides functionality
// for mapping changes to files.
type FileMapper struct {
	fileGateway 	data.FileGatewayI
	tagListGateway 	data.TagListEntityGatewayI
	tagMapper 		TagMapperI
	tagIdentityMap	TagIdentityMapI

	filePath 		string
}

// getFile maps a File, and returns the file and errors encountered.
func (m FileMapper) getFile(id int) (*File, error) {
	result, err := m.fileGateway.ReadFile(id, m.filePath)

	if err != nil {
		return nil, err
	}

	tags, err := m.tagListGateway.ReadTags(id, m.filePath)

	if err != nil {
		return nil, err
	}

	tagList := []*Tag{}

	for _, element := range tags {
		tag, err := m.tagIdentityMap.GetTagById(element.TagId)
		if err != nil {
			return nil, nil
		}
		tagList = append(tagList, tag)
	}

	if err != nil {
		return nil, err
	}
	return NewFile(result.Id, result.Name, tagList), nil
}

// SaveFile saves changes to the given File. SaveFile returns the id of the
// saved File and any errors encountered.
func (m FileMapper) SaveFile(file *File) (int, error) {

	if file.getEntityState() == entityStateDirty {
		return m.saveExistingFile(file)
	} else if file.getEntityState() == entityStateNew {
		return m.saveNewFile(file)
	}

	return file.id, nil
}

// RemoveFile removes the File with the given id, and returns any
// errors encountered.
func (m FileMapper) RemoveFile(id int) (error) {

	err := m.fileGateway.RemoveFile(id, m.filePath)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// getFileByName returns a File by it's name rather, and returns any errors
// encountered.
func (m FileMapper) getFileByName(name string) (*File, error) {

	id, err := m.fileGateway.GetFileIdByName(name, m.filePath)

	if err != nil {
		return nil, err
	}

	return m.getFile(id)
}

// getCWD gets the directory path associated with this FileMapper.
func (m FileMapper) getCWD() (string) {
	return m.filePath
}

// saveExistingFile saves an existing File.
func (m FileMapper) saveExistingFile(file *File) (int, error) {
	err := m.fileGateway.SaveFile(file.id, file.fileName, m.filePath)

	if err != nil {
		return 0, err
	}

	for _, tag := range file.tags {
		_, err = m.tagMapper.SaveTag(tag)
		if err != nil {
			return 0, err
		}
	}

	for _, tag := range file.addedTags() {
		_, err = m.tagListGateway.AddTagListEntity(file.id, tag.id, m.filePath)
		if err != nil {
			return 0, err
		}
	}

	for _, tag := range file.removedTags() {
		err = m.tagListGateway.RemoveTagListEntity(file.id, tag.id, m.filePath)
		if err != nil {
			return 0, err
		}
	}

	return file.id, nil
}

// saveNewFiles saves a new File.
func (m FileMapper) saveNewFile(file *File) (int, error) {

	result, err := m.fileGateway.AddFile(file.fileName, m.filePath)

	if err != nil {
		return 0, err
	}

	file.id = result

	for _, tag := range file.tags {
		_, err := m.tagMapper.SaveTag(tag)
		if err != nil {
			return 0, nil
		}
		_, err = m.tagListGateway.AddTagListEntity(file.id, tag.Id(), m.filePath)
		if err != nil {
			return 0, nil
		}
	}

	return result, nil
}