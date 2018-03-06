
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

// FileIdentityMapI is the interface for accessing File objects currently loaded into.
type FileIdentityMapI interface {
	GetFileById(id int) (*File, error)
	GetFileByName(name string) (*File, error)
	GetFileMapper() (FileMapperI)
	GetMap() (map[int]*File)
	GetAllFiles() ([]*File, error)
}

// NewFileIdentityMap creates a pointer to a FileIdentityMap.
func NewFileIdentityMap(fileMapper FileMapperI) (*FileIdentityMap) {
	return &FileIdentityMap{make(map[int]*File), fileMapper, fileMapper.getCWD(), data.NewFileGateway()}
}

// FileIdentityMap is the implementation of FileIdentityMapI that provides access to the
// File domain objects.
type FileIdentityMap struct {
	files       map[int]*File
	fileMapper  FileMapperI
	cwd         string
	fileGateway data.FileGatewayI
}

// GetFileById returns a File from the map, and loads it if it's not already been
// loaded. Any errors encountered are also returned.
func (f FileIdentityMap) GetFileById(id int) (*File, error) {
	file := f.files[id]
	var err error

	if file == nil {
		file, err = f.fileMapper.getFile(id)
		if err != nil {
			return nil, err
		}
		f.files[id] = file
	}

	return file, err
}

// GetFileByName returns a File from the map, and loads it if it's not already been
// loaded. Any errors encountered are returned.
func (f *FileIdentityMap) GetFileByName(name string) (*File, error) {
	id := -1
	for _, file := range f.files {
		if file.fileName == name {
			id = file.id
			break
		}
	}

	if id < 0 {
		file, err := f.fileMapper.getFileByName(name)
		if err != nil {
			return nil, err
		}
		f.files[file.id] = file
		return file, err
	} else {
		return f.GetFileById(id)
	}
}

// GetAllFiles reports all the Files, and any errors encountered while loading
// the Files.
func (f *FileIdentityMap) GetAllFiles() ([]*File, error) {

	ids, err := f.fileGateway.GetAllFileIds(f.cwd)
	if err != nil {
		return nil, err
	}

	files := make([]*File, len(ids))

	for i, id := range ids {
		file, err := f.GetFileById(id)
		if err != nil {
			return nil, err
		}
		files[i] = file
	}

	return files, nil
}

// GetFileMapper reports the FileMapperI associated with this FileIdentityMap.
func (f FileIdentityMap) GetFileMapper() (FileMapperI) {
	return f.fileMapper
}

// GetMap gets the map of all Files associated with this FileIdentityMap.
func (f FileIdentityMap) GetMap() (map[int]*File) {
	return f.files
}