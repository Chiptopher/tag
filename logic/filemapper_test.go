
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

func TestFileMapper_getFile(t *testing.T) {

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("ReadFile", 1, "path").Return(&data.FileDataDTO{1, "name"}, nil)
	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("ReadTags", 1, "path").	Return([]*data.TagListEntityDTO{}, nil)

	mapper := NewFileMapper(NewTagIdentityMap(nil), "path")
	mapper.fileGateway = fileGateway
	mapper.tagListGateway = tagListGateway

	result, err := mapper.getFile(1)

	var expectedName = "name"
	var expectedId = 1

	if err != nil {
		t.Fatalf("getFile encountered error: %s", err)
	}

	if result.id != expectedId {
		t.Fatalf("Expected %d but got %d", expectedId, result.id)
	}

	if result.fileName != expectedName {
		t.Fatalf("Expected '%s' but got '%s'", expectedName, result.fileName)
	}

}

func TestFileMapper_getFile_MapsTags(t *testing.T) {

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("ReadTags", 1, "path").Return([]*data.TagListEntityDTO{{1, 1, 1}}, nil)
	tagMapper := NewMockTagMapper()
	tagMapper.On("getTag", 1).Return(NewTag(1, "name"), nil)

	tagIdentityMap := NewTagIdentityMap(tagMapper)

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("ReadFile", 1, "path").Return(&data.FileDataDTO{1, "name"}, nil)

	mapper := NewFileMapper(tagIdentityMap, "path")
	mapper.fileGateway = fileGateway
	mapper.tagListGateway = tagListGateway

	result, err := mapper.getFile(1)

	if err != nil {
		t.Fatalf("getFile encountered an error: %s", err)
	}

	if len(result.tags) != 1 {
		t.Fatalf("Expected %d but got %d", 1, len(result.tags))
	}

	tag := result.tags[0]

	if tag.id != 1 {
		t.Fatalf("Expected %d but got %d", 1, tag.id)
	}

	if tag.name != "name" {
		t.Fatalf("Expected '%s' but got '%s'", "name", tag.name)
	}
}

func TestFileMapper_saveFile_NewFile(t *testing.T) {

	fileName := "name"
	filePath := "Path"
	tagName := "name"

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("AddFile", fileName, filePath).Return(1, nil)

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("AddTag", tagName, filePath).Return(1, nil)

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagIdentityMap := NewTagIdentityMap(tagMapper)

	mapper := NewFileMapper(tagIdentityMap, filePath)
	mapper.tagMapper = tagMapper
	mapper.fileGateway = fileGateway

	file := CreateFile(fileName, []*Tag{})
	file.fileName = fileName

	result, err := mapper.SaveFile(file)

	if err != nil {
		t.Fatalf("saveFile encountered an error: %s", err)
	}

	if result != 1 {
		t.Fatalf("Expected %d but got %d", 1, result)
	}

	if file.id != 1 {
		t.Fatalf("Expected %d but got %d", 1, file.id)
	}
}

func TestFileMapper_saveFile_MapsTags(t *testing.T) {

	filePath := "filePath"
	fileName := "name"

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("AddFile", fileName, filePath).Return(1, nil)

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("AddTag", "name", filePath).Return(1, nil)

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagIdentityMap := NewTagIdentityMap(tagMapper)

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("AddTagListEntity", 1, 1, filePath).Return(1, nil)

	file := NewFile(0, "fileName", []*Tag{})
	file.fileName = fileName

	tag := NewTag(0, "name")

	file.tags = append(file.tags, tag)

	mapper := NewFileMapper(tagIdentityMap, filePath)
	mapper.fileGateway = fileGateway
	mapper.tagListGateway = tagListGateway

	mapper.SaveFile(file)

	if tag.id != 1 {
		t.Fatalf("Expected %d but got %d.", 1, tag.id)
	}
}

func TestFileMapper_saveFile_EditsExistingFileButNotTags(t *testing.T) {

	filePath := "filePath"

	name := "name"
	id := 1

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("SaveFile", id, name, filePath).Return(nil)

	file := NewFile(id, name, []*Tag{})

	fileMapper := NewFileMapper(NewTagIdentityMap(nil), filePath)
	fileMapper.fileGateway = fileGateway

	resultId, err := fileMapper.SaveFile(file)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if resultId != id {
		t.Fatalf("Expected %d, but got %d", id, resultId)
	}
}

func TestFileMapper_saveFile_AddedNewTag(t *testing.T) {

	filePath := "filePath"
	fileName := "name"
	tagName := "name"
	fileId := 1
	tagId := 0

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("SaveFile", fileId, fileName, filePath).Return(nil)

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("AddTag", tagName, filePath).Return(1, nil)

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagIdentityMap := NewTagIdentityMap(tagMapper)

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("AddTagListEntity", fileId, 1, filePath).Return(1, nil)

	file := NewFile(fileId, fileName, []*Tag{})

	tag := NewTag(tagId, tagName)
	file.AddTag(tag)

	mapper := NewFileMapper(tagIdentityMap, filePath)
	mapper.fileGateway = fileGateway
	mapper.tagListGateway = tagListGateway

	resultId, err := mapper.SaveFile(file)

	if err != nil {
		t.Fatalf("Error encountered: '%s'", err)
	}

	if resultId != fileId {
		t.Fatalf("Expected '%d', but got '%d'", fileId, resultId)
	}
}

func TestFileMapper_saveFile_AddExistingTag(t *testing.T) {

	filePath := "filePath"
	fileName := "name"
	fileId := 1

	tagName := "name"
	tagId := 1

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("SaveFile", fileId, fileName, filePath).Return(nil)

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("SaveTag", tagId, tagName, filePath).Return(nil)

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagIdentityMap := NewTagIdentityMap(tagMapper)

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("AddTagListEntity", fileId, tagId, filePath).Return(1, nil)

	file := NewFile(fileId, fileName, []*Tag{})

	tag := NewTag(tagId, tagName)
	file.AddTag(tag)

	mapper := NewFileMapper(tagIdentityMap, filePath)
	mapper.fileGateway = fileGateway
	mapper.tagListGateway = tagListGateway


	resultId, err := mapper.SaveFile(file)

	if err != nil {
		t.Fatalf("Error encountered: '%s'", err)
	}

	if resultId != fileId {
		t.Fatalf("Expected '%d', but got '%d'", fileId, resultId)
	}
}

func TestFileMapper_saveFile_RemoveTag(t *testing.T) {

	filePath := "filePath"
	fileName := "name"
	tagName := "name"
	fileId := 1
	tagId := 1

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("SaveFile", fileId, fileName, filePath).Return(nil)

	tagGateway := data.MockNewTagGateway()

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagIdentityMap := NewTagIdentityMap(tagMapper)

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("RemoveTagListEntity", fileId, tagId, filePath).Return(nil)

	tag := NewTag(tagId, tagName)
	file := NewFile(fileId, fileName, []*Tag{tag})

	file.RemoveTag(tag)

	mapper := NewFileMapper(tagIdentityMap, filePath)
	mapper.fileGateway = fileGateway
	mapper.tagListGateway = tagListGateway

	resultId, err := mapper.SaveFile(file)

	if err != nil {
		t.Fatalf("Error encountered: '%s'", err)
	}

	if resultId != fileId {
		t.Fatalf("Expected '%d', but got '%d'", fileId, resultId)
	}
}

func TestFileMapper_deleteFile(t *testing.T) {

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("RemoveFile", 1, "path").Return(nil)

	mapper := FileMapper{}
	mapper.fileGateway = fileGateway
	mapper.filePath = "path"

	err := mapper.RemoveFile(1)

	if err != nil {
		t.Fatalf("removeFile encountered an error: %s", err)
	}
}

func TestFileMapper_getFileByName(t *testing.T) {

	fileId := 1
	fileName := "name"
	filePath := "path"

	fileGateway := data.MockNewFileGateway()
	fileGateway.On("GetFileIdByName", fileName, filePath).Return(fileId, nil)
	fileGateway.On("ReadFile", fileId, filePath).Return(&data.FileDataDTO{fileId, fileName}, nil)

	tagListGateway := data.MockNewTagListEntityGateway()
	tagListGateway.On("ReadTags", fileId, filePath).Return([]*data.TagListEntityDTO{}, nil)

	fileMapper := NewFileMapper(NewTagIdentityMap(nil), filePath)
	fileMapper.fileGateway = fileGateway
	fileMapper.tagListGateway = tagListGateway

	file, err := fileMapper.getFileByName(fileName)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if file.id != fileId {
		t.Fatalf("File incorrectly loaded.")
	}
}

func TestFileMapper_saveFile_OnlySavesExistingFileWhenTheFileHasChanged(t *testing.T) {

	fileId := 1
	fileName := "name"
	filePath := "path"
	newFileName := "Deep blue see"

	file := NewFile(fileId, fileName, []*Tag{})
	fileGateway := data.MockNewFileGateway()
	fileGateway.On("SaveFile", fileId, newFileName, filePath).Return(nil)
	fileMapper := NewFileMapper(NewTagIdentityMap(NewTagMapper(filePath)), filePath)
	fileMapper.fileGateway = fileGateway
	file.SetFileName(newFileName)
	fileMapper.SaveFile(file)

	fileGateway.AssertCalled(t, "SaveFile", fileId, newFileName, filePath)

}