
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
	"fmt"
	"path/filepath"
	"os"
)

func TestAppRun_ListTags_Success(t *testing.T) {

	Run([]string{"init", e2eFolder}, "")
	Run([]string{"file", "A.txt", "name"}, e2eFolder)
	result := Run([]string{"list", "A.txt"}, e2eFolder)
	expectedResult := "tags for A.txt\n  name\n"

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'.", expectedResult, result)
	}
}

func TestAppRun_ListTags_NotEnoughArgs(t *testing.T) {

	result := Run([]string{"list"}, e2eFolder)
	expectedResult := "Not enough arguments."

	if result != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'.", expectedResult, result)
	}
}

func TestAppRun_ListTags_TooManyArgs(t *testing.T) {

	result := Run([]string{"list", "A.txt", "Extra"}, e2eFolder)
	expectedResult := "Too many arguments."

	if result != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'", expectedResult, result)
	}

}

func TestAppRun_ListTags_FileThatDoesNotExit(t *testing.T) {

	Run([]string{"init", e2eFolder}, "")
	fileName := "TestAppRun_ListTags_FileThatDoesNotExit"
	result := Run([]string{"list", fileName}, e2eFolder)
	expectedResult := fmt.Sprintf("File doesn't exist at: %s.", filepath.Join(e2eFolder, fileName))

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}

func TestAppRun_ListTags_ReferenceFileWithRelativePathThatIsInUpperDirectory(t *testing.T) {
	Run([]string{"init", e2eFolder}, "")

	fileName := "TestAppRun_ListTags_ReferenceFileWithRelativePathThatIsInUpperDirectory"
	tag := "Tag"

	os.Create(filepath.Join(e2eFolder, fileName))

	Run([]string{"file", fileName, tag}, e2eFolder)
	result := Run([]string{"list", filepath.Join("..", fileName)}, e2eSubFolder)

	expectedResult := fmt.Sprintf("tags for %s\n  %s\n", filepath.Join("..", fileName), tag)

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}