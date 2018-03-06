
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
	"github.com/Chiptopher/tag/data"
)


func TestFileIdentityMap_GetFileById(t *testing.T) {

	filePath := "path"

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("ReadTags", 1, filePath).Return([]*data.TagListEntityDTO{}, nil)

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("ReadFile", 1, filePath).Return(&data.FileDataDTO{1, "name"}, nil)

	fileMapper := NewFileMapper(NewTagIdentityMap(nil), filePath)
	fileMapper.tagListGateway = tagListGateway
	fileMapper.fileGateway = fileGateway

	identityMap := NewFileIdentityMap(fileMapper)
	result1, err := identityMap.GetFileById(1)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	result2, err := identityMap.GetFileById(1)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if result1 != result2 {
		t.Fatalf("Doesn't get a single instance of the same file.")
	}
}

func TestFileIdentityMap_GetFileByName(t *testing.T) {
	
	filePath := "path"

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("ReadTags", 1, filePath).Return([]*data.TagListEntityDTO{}, nil)

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("ReadFile", 1, filePath).Return(&data.FileDataDTO{1, "name"}, nil)
	fileGateway.On("GetFileIdByName", "name", filePath).Return(1, nil)

	fileMapper := NewFileMapper(NewTagIdentityMap(nil), filePath)
	fileMapper.tagListGateway = tagListGateway
	fileMapper.fileGateway = fileGateway

	identityMap := NewFileIdentityMap(fileMapper)
	result1, err := identityMap.GetFileById(1)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	result2, err := identityMap.GetFileByName("name")

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if result1 != result2 {
		t.Fatalf("Doesn't get a single instance of the same file.")
	}
}

func TestFileIdentityMap_GetAllFiles(t *testing.T) {

	filePath := "path"

	file1 := NewFile(1, "Name1", []*Tag{})
	file2 := NewFile(2, "Name1", []*Tag{})

	fileMapper := MockNewFileMapper()
	fileMapper.On("getFile", 1).Return(file1, nil)
	fileMapper.On("getFile", 2).Return(file2, nil)
	fileMapper.On("getCWD").Return(filePath)

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("GetAllFileIds", filePath).Return([]int{1, 2}, nil)

	fileIdentityMap := NewFileIdentityMap(fileMapper)
	fileIdentityMap.fileGateway = fileGateway

	files, err := fileIdentityMap.GetAllFiles()
	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if files[0] != file1 {
		t.Fatalf("Loaded the wrong file.")
	}

	if files[1] != file2 {
		t.Fatalf("Loaded the wrong file.")
	}
}