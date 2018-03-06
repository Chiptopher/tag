
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
	"github.com/Chiptopher/tag/data"
)


func TestInitializer_Parse_Init_TooFewArguments_e2e(t *testing.T) {

	result := Run([]string{"init"}, "")
	expectedResult := "Not enough arguments entered. Path to init required."

	if result !=  expectedResult {
		t.Fatalf("Expected '%s', but got '%s'.", expectedResult, result)
	}
}

func TestInitializer_Parse_TooManyArguments_e2e(t *testing.T) {

	result := Run([]string{"init", "path", "extra"}, "")
	expectedResult := "Too many arguments entered: [extra]"

	if result!=  expectedResult {
		t.Fatalf("Expected '%s', but got '%s'.", expectedResult, result)
	}
}

func TestInitializer_Parse_Success_e2e(t *testing.T) {

	result := Run([]string{"init", e2eFolder}, "")
	expectedResult := ""

	if result != expectedResult {
		t.Fatalf("Incorrect result returned. Expected '%s', but got '%s'.", "", result)
	}

	if _, err := os.Stat(filepath.Join(e2eFolder, data.DataFolderName)); os.IsNotExist(err) {
		t.Fatalf("Folder not created: %s", err)
	}

	if _,err := os.Stat(filepath.Join(e2eFolder, filepath.Join(data.DataFolderName, data.DatabaseName))); os.IsNotExist(err) {
		t.Fatalf("Did not create the data file.")
	}
}