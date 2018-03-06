
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
	"bytes"
	"fmt"
	"path/filepath"
	"github.com/Chiptopher/tag/data"
)

// NewListTagsCommand creates a pointer to a ListTagsCommand.
func NewListTagsCommand(args []string, transaction TransactionI) (*ListTagsCommand) {
	return &ListTagsCommand{args, transaction}
}

// ListTagsCommand gets the Tags associated with a File.
type ListTagsCommand struct {
	args        []string
	transaction TransactionI
}

func (c *ListTagsCommand) Execute() (SubResult) {

	if len(c.args) != 1 {
		return &listTagsResult{c.args, nil, nil}
	}

	actualPath := filepath.Join(c.transaction.GetCWD(), c.args[0])

	file, err := c.transaction.GetFileByName(actualPath)

	if err != nil {
		if _, ok := err.(*data.FileNotFoundError); ok {
			return listTagsResult{nil, nil, err}
		}

		// TODO fix this
	}

	var tagNames []string

	tagNames = make([]string, len(file.Tags()))
	for i, tag := range file.Tags() {
		tagNames[i] = tag.Name()
	}

	return listTagsResult{c.args, tagNames, err}
}

type listTagsResult struct {
	args         []string
	tagNames     []string
	fileNotFound error
}

func (r listTagsResult) Fmt() (string) {

	if r.fileNotFound != nil {
		return fmt.Sprintf("File doesn't exist at: %s.", r.fileNotFound.(*data.FileNotFoundError).FilePath())
	} else if len(r.args) < 1 {
		return "Not enough arguments."
	} else if len(r.args) > 1 {
		return "Too many arguments."
	}

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("tags for %s\n", r.args[0]))
	for _, tagName := range r.tagNames {
		buffer.WriteString(fmt.Sprintf("  %s\n", tagName))
	}

	return buffer.String()
}

func (r listTagsResult) GetSuccess() (bool) {
	return false
}