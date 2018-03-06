
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
	"github.com/Chiptopher/tag/logic"
)

func TestRemoveTagCommand_Execute_Success(t *testing.T) {

	fileName := "fileName"
	tagName := "Tag"
	cwd := "CWD"

	tag := logic.NewTag(1, tagName)
	file := logic.NewFile(1, fileName, []*logic.Tag{tag})

	transaction := MockNewTransaction()
	transaction.On("GetFileByName", filepath.Join(cwd, fileName)).Return(file, nil)
	transaction.On("GetCWD").Return(cwd)
	transaction.On("GetTagByName", tagName).Return(tag, nil)

	args := []string{fileName, tagName}

	command := NewRemoveTagCommand(args, transaction)
	result := command.Execute()
	expectedResult := ""

	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result.Fmt())
	}

}

func TestRemoveTagCommand_Execute_TooFewArguments(t *testing.T) {
	command := NewRemoveTagCommand([]string{"file"}, nil)
	result := command.Execute()
	expectedResult := "Too few arguments."

	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result.Fmt())
	}

	if result.GetSuccess() {
		t.Fatalf("Shouldn't be success")
	}
}

func TestRemoveTagCommand_Execute_TooManyArguments(t *testing.T) {
	command := NewRemoveTagCommand([]string{"file", "tag", "extra"}, nil)
	result := command.Execute()
	expectedResult := "Too many arguments."

	if result.Fmt() != expectedResult {
		t.Fatalf("Expected '%s' but got '%s'", expectedResult, result.Fmt())
	}

	if result.GetSuccess() {
		t.Fatalf("Shouldn't be success")
	}
}