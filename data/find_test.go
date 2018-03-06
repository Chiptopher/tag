
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


package data

import (
	"testing"
	"github.com/Chiptopher/tag/test"
	"path"
	"os"
	"reflect"
)

func TestFind(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)

	actual, err := find(test.GoPathTestFolder)
	var expected = path.Join(test.GoPathTestFolder, DataFolderName)

	if err != nil {
		t.Fatalf("find encountered error: %s", err)
	}

	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestFind_FromSubFolder(t *testing.T) {

	subFolderPath := test.CreatePathInTestFolder("sub")
	os.MkdirAll(subFolderPath, 0700)
	CreateDataFolder(test.GoPathTestFolder)

	var expected = path.Join(test.GoPathTestFolder, DataFolderName)

	actual, err := find(subFolderPath)

	if err != nil {
		t.Fatalf("find encountered error %s", err)
	}

	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestFind_NoDatFileReturnsNotFoundError(t *testing.T) {
	_, err := find(test.GoPath)

	if _, ok := err.(*FindNotFoundError); !ok {
		t.Fatalf("find returned error '%s' but should return error type '%s'.", reflect.TypeOf(err), reflect.TypeOf(FileNotFoundError{}))
	}
}