
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

package business

import (
	"testing"
	"github.com/Chiptopher/tag/logic"
)

func TestNewTransaction(t *testing.T) {

	transaction := NewTransaction("path")

	fileMapper := transaction.GetFileMapper()
	if fileMapper == nil {
		t.Fatalf("FileMapper not loaded.")
	}

	tagMapper := transaction.GetTagMapper()
	if tagMapper == nil {
		t.Fatalf("TagMapper not created.")
	}

	tagMap := transaction.GetTagIdentityMap()
	if tagMap == nil {
		t.Fatalf("TagMap not loaded.")
	}
	if tagMap.GetTagMapper() != tagMapper {
		t.Fatalf("TagMapper should be loaded into Tag identity map")
	}

	fileMap := transaction.GetFileIdentityMap()
	if fileMap == nil {
		t.Fatalf("FileMap not loaded.")
	}
	if fileMap.GetFileMapper() != fileMapper {
		t.Fatalf("FileMapper should be loaded into file identity map")
	}
}

func TestTransaction_GetFileById(t *testing.T) {

	file := logic.NewFile(1, "name", []*logic.Tag{})

	fileIMap := logic.MockNewFileIdentityMap()
	fileIMap.On("GetFileById", 1).Return(file, nil)

	transaction := NewTransaction("path")
	transaction.fileMap = fileIMap

	fileOut1, _ := transaction.GetFileById(1)
	fileOut2, _ := transaction.GetFileById(1)

	if fileOut1 != fileOut2 {
		t.Fatalf("Should load the same file.")
	}

}

func TestTransaction_GetFileById_GettingTheFileAgainKeepsOriginalCopyOfMap(t *testing.T) {

	file := logic.NewFile(1, "name", []*logic.Tag{})
	fileIMap := logic.MockNewFileIdentityMap()
	fileIMap.On("GetFileById", 1).Return(file, nil)

	transaction := NewTransaction("path")
	transaction.fileMap = fileIMap

	fileOut1, _ := transaction.GetFileById(1)
	fileOut1.SetFileName("Name2")
	transaction.GetFileById(1)

	if transaction.fileOriginal[1].FileName() != "name" {
		t.Fatalf("File name changed when reading the file a second time.")
	}
}

func TestTransaction_GetFileByName(t *testing.T) {
	file := logic.NewFile(1, "name", []*logic.Tag{})
	fileIMap := logic.MockNewFileIdentityMap()
	fileIMap.On("GetFileByName", "name").Return(file, nil)

	transaction := NewTransaction("path")
	transaction.fileMap = fileIMap

	fileOut1, _ := transaction.GetFileByName("name")
	fileOut2, _ := transaction.GetFileByName("name")

	if fileOut1 != fileOut2 {
		t.Fatalf("Doesn't load the same instance of the file.")
	}
}

func TestTransaction_GetTagById(t *testing.T) {

	tag := logic.NewTag(1, "name")

	tagIdentityMap := logic.MockNewTagIdentityMap()
	tagIdentityMap.On("GetTagById", 1).Return(tag, nil)

	transaction := NewTransaction("cwd")
	transaction.tagMap = tagIdentityMap

	transaction.GetTagById(1)

	if transaction.tagOriginal[1] == nil {
		t.Fatalf("Doesn't load an original version of the file into the map.")
	}
}

func TestTransaction_GetTagByName(t *testing.T) {

	tag := logic.NewTag(1, "name")

	tagIdentityMap := logic.MockNewTagIdentityMap()
	tagIdentityMap.On("GetTagByName", "name").Return(tag, nil)

	transaction := NewTransaction("cwd")
	transaction.tagMap = tagIdentityMap

	transaction.GetTagByName("name")

	if transaction.tagOriginal[1] == nil {
		t.Fatalf("Doesn't load an original version of the file into the map.")
	}
}

func TestTransaction_RegisterFile(t *testing.T) {
	file := logic.NewFile(1, "name", []*logic.Tag{})

	transaction := NewTransaction("path")
	transaction.RegisterFile(file)

	if transaction.addedFiles[0] != file {
		t.Fatalf("File not registered with the transaction.")
	}
}

func TestTransaction_Commit_RegisteredFiles(t *testing.T) {

	file := logic.NewFile(0, "name", []*logic.Tag{})

	filePath := "path"
	fileMapper := logic.MockNewFileMapper()
	fileMapper.On("SaveFile", file).Return(1, nil)

	transaction := NewTransaction(filePath)
	transaction.fileMapper = fileMapper

	transaction.RegisterFile(file)

	err := transaction.Commit()

	fileMapper.AssertCalled(t, "SaveFile", file)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}
}

func TestTransaction_Commit_ChangedFiles(t *testing.T) {
	filePath := "path"
	file := logic.NewFile(1, "name", []*logic.Tag{})
	mockMap := make(map[int]*logic.File)
	mockMap[1] = file

	fileIdentityMap := logic.MockNewFileIdentityMap()
	fileIdentityMap.On("GetMap").Return(mockMap)

	fileMapper := logic.MockNewFileMapper()
	fileMapper.On("SaveFile", file).Return(1, nil)

	transaction := NewTransaction(filePath)
	transaction.fileMap = fileIdentityMap
	transaction.fileMapper = fileMapper

	err := transaction.Commit()

	fileMapper.AssertCalled(t, "SaveFile", file)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}
}
