
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
	"path/filepath"
	"github.com/stretchr/testify/mock"
	"github.com/Chiptopher/tag/data"
)

func TestListTagsCommand_Execute(t *testing.T) {

	fileName := "A.txt"

	filePath := "path"
	expectedName := filepath.Join(filePath, fileName)

	tag := logic.NewTag(1, "name")
	file := logic.NewFile(1, fileName, []*logic.Tag{tag})

	transaction := MockNewTransaction()
	transaction.On("GetFileByName", expectedName).Return(file, nil)
	transaction.On("GetCWD").Return(filePath)

	args := []string{"A.txt"}

	command := NewListTagsCommand(args, transaction)
	result := command.Execute()

	expectedOutput := "tags for A.txt\n  name\n"

	if result.Fmt() != expectedOutput {
		t.Fatalf("Expected '%s' but got '%s'.", expectedOutput, result.Fmt())
	}
}

func TestListTagsCommand_Execute_TooFewArguments(t *testing.T) {

	transaction := MockNewTransaction()
	args := []string{}

	command := NewListTagsCommand(args, transaction)
	result := command.Execute()

	expectedOutput := "Not enough arguments."

	if result.Fmt() != expectedOutput {
		t.Fatalf("Expected '%s' but got '%s'", expectedOutput, result.Fmt())
	}
}

func TestListTagsCommand_Execute_TooManyArguments(t *testing.T) {

	transaction := MockNewTransaction()
	args := []string{"a.txt", "extra"}

	command := NewListTagsCommand(args, transaction)
	result := command.Execute()

	expectedOutput := "Too many arguments."

	if result.Fmt() != expectedOutput {
		t.Fatalf("Expected '%s' but got '%s'", expectedOutput, result.Fmt())
	}
}

func TestListTagsCommand_Execute_FileDoesNotExit(t *testing.T) {

	transaction := MockNewTransaction()
	transaction.On("GetCWD").Return("")
	transaction.On("GetFileByName", mock.Anything).Return(nil, data.NewFileNotFoundError("FilePath"))
	args := []string{"filenonexistant"}

	command := NewListTagsCommand(args, transaction)
	result := command.Execute()

	expectedOutput := "File doesn't exist at: FilePath."

	if result.Fmt() != expectedOutput {
		t.Fatalf("Expected '%s' but got '%s'", expectedOutput, result.Fmt())
	}
}