
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
	"io/ioutil"
	"path"
	"path/filepath"
)

// Find finds the next closest data folder up from the given file path.
func Find(filePath string) (string, error) {
	return find(filePath)
}

func find(filePath string) (string, error) {

	var (
		previousDir = ""
	)

	dir, err := filepath.Abs(filePath)

	if err != nil {
		return "", err
	}

	// if the previous and current data folders match, then top level
	// has been hit.
	for dir != previousDir {

		files, err := ioutil.ReadDir(dir)

		if err != nil {
			return "", err
		}

		for _, f := range files {
			if f.Name() == DataFolderName {
				return path.Join(dir, f.Name()), nil
			}
		}

		previousDir = dir
		dir, err = filepath.Abs(filepath.Dir(dir))
	}

	return "", &FindNotFoundError{filePath}
}