
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

package app

import (
	"testing"
	"github.com/Chiptopher/tag/data"
	"github.com/Chiptopher/tag/test"
	"path/filepath"
	"fmt"
	"os"
)

func TestAppRun_TagFile_NewFile_NewTag_Success(t *testing.T) {

	fileName := "TestAppRun_TagFile_NewFile_NewTag_Success_A.txt"
	os.Create(filepath.Join(e2eFolder, fileName))
	expectedPath := filepath.Join(e2eFolder, fileName)
	Run([]string{"init", e2eFolder}, "")
	result := Run([]string{"file", fileName, "name"}, e2eFolder)
	expectedResult := ""

	if result != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'", expectedResult, result)
	}

	fileId, err := data.NewFileGateway().GetFileIdByName(expectedPath, e2eFolder)
	file, err := data.NewFileGateway().ReadFile(fileId, e2eFolder)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if file.Name != expectedPath {
		t.Fatalf("Expected '%s', but recieved '%s'", expectedPath, file.Name)
	}

	tagId, err := data.NewTagGateway().GetTagIdByName("name", e2eFolder)
	tag, err := data.NewTagGateway().ReadTag(tagId, test.LinkToDataFileInPath(e2eFolder))

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if tag == nil {
		t.Fatalf("Tag not saved correctly.")
	}

	if tag.Name != "name" {
		t.Fatalf("Expected '%s' but got '%s'", "name", tag.Name)
	}

	tagEntityData, err := data.NewTagListEntityGateway().ReadTags(fileId, test.LinkToDataFileInPath(e2eFolder))

	if len(tagEntityData) != 1 {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}

	if tagEntityData[0].FileId != fileId {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}

	if  tagEntityData[0].TagId != tagId {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}
}

func TestAppRun_TagFile_ExistingFile_NewTag_Success(t *testing.T) {

	fileName := "TestAppRun_TagFile_ExistingFile_NewTag_Success_A.txt"
	os.Create(filepath.Join(e2eFolder, fileName))
	expectedFileName := filepath.Join(e2eFolder, fileName)
	Run([]string{"init", e2eFolder}, "")
	Run([]string{"file", fileName, "name"}, e2eFolder)
	result := Run([]string{"file", fileName, "Name2"}, e2eFolder)
	expectedResult := ""

	if result != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'", expectedResult, result)
	}

	tagId, err := data.NewTagGateway().GetTagIdByName("Name2", e2eFolder)
	tag, err := data.NewTagGateway().ReadTag(tagId, test.LinkToDataFileInPath(e2eFolder))

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if tag == nil {
		t.Fatalf("Tag not saved correctly.")
	}

	fileId, _ := data.NewFileGateway().GetFileIdByName(expectedFileName, e2eFolder)
	if tag.Name != "Name2" {
		t.Fatalf("Expected '%s' but got '%s'", "Name2", tag.Name)
	}

	tagEntityData, err := data.NewTagListEntityGateway().ReadTags(fileId, test.LinkToDataFileInPath(e2eFolder))

	if len(tagEntityData) != 2 {
		t.Fatalf("Expected %d but got %d", 2, len(tagEntityData))
	}

	if tagEntityData[1].FileId != fileId {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}

	if  tagEntityData[1].TagId != tagId {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}
}

func TestAppRun_TagFile_TagFileFromBelowDataFolder_Success(t *testing.T) {

	fileName := "TestAppRun_TagFile_TagFileFromBelowDataFolder_Success_B.txt"
	finalFileName := filepath.Join(e2eSubFolder, fileName)
	os.Create(finalFileName)

	Run([]string{"init", e2eFolder}, "")
	result := Run([]string{"file", fileName, "name"}, e2eSubFolder)
	expectedResult := ""

	if result != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'", expectedResult, result)
	}

	fileId, _ := data.NewFileGateway().GetFileIdByName(finalFileName, e2eSubFolder)
	file, err := data.NewFileGateway().ReadFile(fileId, e2eFolder)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if file.Name != finalFileName {
		t.Fatalf("Expected '%s', but recieved '%s'", finalFileName, file.Name)
	}

	tagId, _ := data.NewTagGateway().GetTagIdByName("name", e2eSubFolder)
	tag, err := data.NewTagGateway().ReadTag(tagId, e2eSubFolder)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if tag == nil {
		t.Fatalf("Tag not saved correctly.")
	}

	if tag.Name != "name" {
		t.Fatalf("Expected '%s' but got '%s'", "name", tag.Name)
	}

	tagEntityData, err := data.NewTagListEntityGateway().ReadTags(fileId, e2eSubFolder)

	if len(tagEntityData) != 1 {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}

	if tagEntityData[0].FileId != fileId {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}

	if  tagEntityData[0].TagId != tagId {
		t.Fatalf("TagEntityData not saved correctly: '%p'", tagEntityData)
	}
}

func TestAppRun_TagFile_CannotTagFilesThatDoNotExist(t *testing.T) {
	fileName := "TestAppRun_TagFile_CannotTagFilesThatDoNotExist"
	fullFileName := filepath.Join(e2eFolder, fileName)
	Run([]string{"init", e2eFolder}, "")
	result := Run([]string{"file", fileName, "Tag"}, e2eFolder)
	expectedResult := fmt.Sprintf("File %s doesn't exist.", fullFileName)

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}