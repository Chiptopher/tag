
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
	"path/filepath"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/Chiptopher/tag/logic"
)

func TestTagFileCommand_Execute_Success(t *testing.T) {

	filePath := "Path"
	fileName := "name"
	tagName := "Tag"

	expectedFileName := filepath.Join(filePath, fileName)

	transaction := MockNewTransaction()
	transaction.On("GetFileByName", expectedFileName).Return(logic.NewFile(1, expectedFileName, []*logic.Tag{}), nil)
	transaction.On("GetCWD").Return(filePath)
	transaction.On("GetTagByName", tagName).Return(logic.NewTag(1, tagName), nil)

	command := TagFileCommand{[]string{fileName, tagName}, transaction}
	result := command.Execute()

	if !result.GetSuccess() {
		t.Fatalf("Result should be marked as successful.")
	}

	if result.Fmt() != "" {
		t.Fatalf("Expected '%s' but got '%s'", "", result.Fmt())
	}
}

func TestTagFileCommand_Execute_TooManyArguments(t *testing.T) {

	command := TagFileCommand{[]string{"A.txt", "Tag", "Extra"}, nil}
	result := command.Execute()

	if result.GetSuccess() {
		t.Fatalf("Result should be markes as unsuccessful.")
	}

	expectedResultFmt := "Too many arguments entered: [A.txt Tag Extra]"

	if result.Fmt() != expectedResultFmt {
		t.Fatalf("Expected '%s' but got '%s'", expectedResultFmt, result.Fmt())
	}
}

func TestTagFileCommand_Execute_TooFewArguments(t *testing.T) {
	command := TagFileCommand{[]string{"A.txt"}, nil}
	result := command.Execute()

	if result.GetSuccess() {
		t.Fatalf("Result should be marked as unsuccessful.")
	}

	expectedResultFmt := "Not enough arguments."

	if result.Fmt() != expectedResultFmt {
		t.Fatalf("Expected '%s' but got '%s'", expectedResultFmt, result.Fmt())
	}
}

func TestTagFileCommand_Execute_TryingToTagANewFileSavesIt(t *testing.T) {

	filePath := "Path"
	fileName := "name"
	tagName := "Tag"

	expectedPath := filepath.Join(filePath, fileName)

	transaction := MockNewTransaction()
	transaction.On("GetFileByName", expectedPath).Return(nil, errors.New("error"))
	transaction.On("GetCWD").Return(filePath)
	transaction.On("GetTagByName", tagName).Return(logic.NewTag(1, tagName), nil)
	transaction.On("RegisterFile", mock.Anything)

	command := TagFileCommand{[]string{fileName, tagName}, transaction}
	result := command.Execute()

	expectedResult := ""

	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'.", expectedResult, result.Fmt())
	}
}

var tagFileFmtTest = []struct {
	result         TagFileResult
	expectedOutput string
} {
	{TagFileResult{true, nil, "", ""}, ""},
	{TagFileResult{false, []string{"A.txt", "Tag", "Extra"}, "", ""}, "Too many arguments entered: [A.txt Tag Extra]"},
	{TagFileResult{false, []string{"A.txt"}, "", ""}, "Not enough arguments."},
}

func TestTagFileResult_Fmt(t *testing.T) {
	for _, ft := range tagFileFmtTest {
		output := ft.result.Fmt()
		if output != ft.expectedOutput {
			t.Fatalf("Expected '%s', but got '%s'.", ft.expectedOutput, ft.result.Fmt())
		}
	}
}