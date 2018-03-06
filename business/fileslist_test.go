
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

func TestFileListCommand_Execute_Success(t *testing.T) {

	tag := logic.NewTag(1, "name")
	file1 := logic.NewFile(1, "File1", []*logic.Tag{tag})
	file2 := logic.NewFile(2, "File2", []*logic.Tag{})

	args := []string{"name"}

	transaction := MockNewTransaction()
	transaction.On("GetAllFiles").Return([]*logic.File{file1, file2}, nil)
	transaction.On("GetTagByName", tag.Name()).Return(tag, nil)

	command := NewFilesListCommand(args, transaction)
	result := command.Execute()

	expectedResult := "Files with tag 'name'\n  File1\n"
	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result.Fmt())
	}
}

func TestFileListCommand_Execute_TooFewArguments(t *testing.T) {

	args := []string{}

	transaction := MockNewTransaction()

	command := NewFilesListCommand(args, transaction)
	result := command.Execute()

	expectedResult := "Too few arguments."

	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result.Fmt())
	}
}

func TestFileListCommand_Execute_TooManyArguments(t *testing.T) {
	args := []string{"name", "Extra"}
	transaction := MockNewTransaction()

	command := NewFilesListCommand(args, transaction)
	result := command.Execute()

	expectedResult := "Too many arguments."

	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result.Fmt())
	}
}