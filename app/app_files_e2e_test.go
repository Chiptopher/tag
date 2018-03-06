
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
	"os"
	"path/filepath"
	"fmt"
)

func TestAppRun_Files_Success(t *testing.T) {

	fileName := "TestAppRun_Files_Success_A.txt"
	file2Name := "TestAppRun_Files_Success_B.txt"
	os.Create(filepath.Join(e2eFolder, fileName))
	os.Create(filepath.Join(e2eFolder, file2Name))
	tagName := "TestAppRun_Files_Success_Name"

	Run([]string{"init", e2eFolder}, "")
	Run([]string{"file", fileName, tagName}, e2eFolder)
	Run([]string{"file", file2Name, tagName}, e2eFolder)
	result := Run([]string{"files", tagName}, e2eFolder)
	expectedResult := fmt.Sprintf("Files with tag '%s'\n  %s\n  %s\n", tagName, filepath.Join(e2eFolder, fileName), filepath.Join(e2eFolder, file2Name))

	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}

}

func TestAppRun_Files_TooFewArguments(t *testing.T) {

	result := Run([]string{"files"}, e2eFolder)
	expectedResult := "Too few arguments."
	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}

func TestAppRun_Files_TooManyArguments(t *testing.T) {

	result := Run([]string{"files", "file", "extra"}, e2eFolder)
	expectedResult := "Too many arguments."
	if result != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result)
	}
}