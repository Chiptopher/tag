
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

import "fmt"

func NewFileNotFoundError(filePath string) (*FileNotFoundError) {
	return &FileNotFoundError{filePath}
}

// FindNotFoundError reported when Find fails to find a data source folder at
// the given path.
type FindNotFoundError struct {
	filePath string
}

func (err FindNotFoundError) Error() string {
	return fmt.Sprintf("Data source not found at %s", err.filePath)
}

// FileNotFoundError reported when a desired file path in the data source doesn't
// exist.
type FileNotFoundError struct {
	filePath string
}

func (err FileNotFoundError) FilePath() string {
	return err.filePath
}

func (err FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found at '%s'", err.filePath)
}

// UnexpectedError reported when an unprepared-for error occurs.
type UnexpectedError struct {
	err error
}

func (err UnexpectedError) Error() (string) {
	return fmt.Sprintf("Unexpected error encountered: '%s'", err.err)
}

// TagNotFoundError reported when a desired tag in the data source doesn't
// exist.
type TagNotFoundError struct {
	filePath string
}

func (err TagNotFoundError) Error() (string) {
	return fmt.Sprintf("Tagnot found at '%s'", err.filePath)
}

// NoFileExistsError reported when tagging a file that doesn't exist in the
// file system.
type NoFileExistsError struct {
	filePath string
}

func (err NoFileExistsError) Error() (string) {
	return fmt.Sprintf("No file exists at: '%s'", err.filePath)
}

func (err NoFileExistsError) FilePath() string {
	return err.filePath
}


// PreexistingError reported when attempting to create a data source folder
// when one already exists.
type PreexistingError struct {
	FilePath string
}

func (f PreexistingError) Error() string {
	return fmt.Sprintf("file: Cannot create data folder at %s becuase" +
		" one already exists.", f.FilePath)
}