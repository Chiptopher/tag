
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
	"path/filepath"
	"os"
	"github.com/Chiptopher/tag/data"
)

func TestAppRun_RemoveTag_Success(t *testing.T) {

	fileName := "TestAppRun_RemoveTag_Success"
	fullFileName := filepath.Join(e2eFolder, fileName)
	tagName := "TestAppRun_RemoveTag_Success_Tag"

	_, err := os.Create(fullFileName)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	Run([]string{"init", e2eFolder}, "")
	Run([]string{"file", fileName, tagName}, e2eFolder)
	result := Run([]string{"remove", fileName, tagName}, e2eFolder)
	expectedResult := ""

	if result != expectedResult {
		t.Fatalf("Revieved unexpected result: %s", result)
	}

	tagId, _ := data.NewTagGateway().GetTagIdByName(tagName, e2eFolder)
	files, _ := data.NewTagListEntityGateway().ReadFiles(tagId, e2eFolder)

	if len(files) != 0 {
		t.Fatalf("Files longer than it should be: %d", len(files))
	}
}

func TestAppRun_RemoveTag_TooFewArguemnts(t *testing.T) {
	result := Run([]string{"remove", "fileName"}, e2eFolder)
	expectedResult := "Too few arguments."

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}

func TestAppRun_RemoveTag_TooManyrguemnts(t *testing.T) {
	result := Run([]string{"remove", "fileName", "tag", "extra"}, e2eFolder)
	expectedResult := "Too many arguments."

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}