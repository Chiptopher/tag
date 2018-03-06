
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
	"github.com/Chiptopher/tag/data"
)

func TestCreateDataFolderCommand_Execute_TooManyArguments(t *testing.T) {

	command := NewCreateDataFolderCommand([]string{"path", "extra"})
	result := command.Execute()

	expectedArgs := []string{"extra"}

	if result.(*createDataFolderResult).Args == nil {
		t.Fatalf("Args expected.")
	}

	for i, arg := range result.(*createDataFolderResult).Args {
		if arg != expectedArgs[i] {
			t.Fatalf("Expected args and actual args out of sync.")
		}
	}
}

func TestCreateDataFolderCommand_Execute_TooFewArguments(t *testing.T) {


	command := NewCreateDataFolderCommand([]string{})
	result := command.Execute()

	if result.(*createDataFolderResult).Args == nil {
		t.Fatalf("Args expected.")
	}

	if len(result.(*createDataFolderResult).Args) > 0 {
		t.Fatalf("Args should be empty.")
	}
}

func TestCreateDataFolderCommand_Execute_CorrectNumberOfArgs(t *testing.T) {

	mockCreate := data.MockCreateDataFolderData{true, nil}

	command := NewCreateDataFolderCommand([]string{"path"})
	command.createStrategy = mockCreate.MockCreateDataFolder

	result := command.Execute()

	if result.(*createDataFolderResult).Args != nil {
		t.Fatalf("Args not expected.")
	}

	err := result.(*createDataFolderResult).Err

	if result.(*createDataFolderResult).Err != nil {
		t.Fatalf("Error not expected: %s", err)
	}
}

var createDataFolderFmtTest = []struct {
	result         createDataFolderResult
	expectedOutput string
} {
	{createDataFolderResult{nil, nil, true}, ""},
}

func TestCreateDataFolderResult_Fmt(t *testing.T) {

	for _, ft := range createDataFolderFmtTest {
		output := ft.result.Fmt()
		if output != ft.expectedOutput {
			t.Fatalf("Expected '%s', but got '%s'", ft.expectedOutput, output)
		}
	}
}

