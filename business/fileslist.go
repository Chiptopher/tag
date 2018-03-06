
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
	"github.com/Chiptopher/tag/logic"
	"fmt"
	"bytes"
)

// NewFilesListCommand creates a pointer to a FilesListCommand.
func NewFilesListCommand(args []string, transaction TransactionI) (*FilesListCommand) {
	return &FilesListCommand{args, transaction}
}

// FilesListCommand gets the list of all Files associated with a certain Tag.
type FilesListCommand struct {
	args        []string
	transaction TransactionI
}

func (c *FilesListCommand) Execute() (SubResult) {

	if len(c.args) != 1 {
		return &filesListResult{c.args, nil}
	}

	files, err := c.transaction.GetAllFiles()
	if err != nil {
		// TODO something with this
	}

	tagArg, err := c.transaction.GetTagByName(c.args[0])
	outFiles := []*logic.File{}

	if err != nil {
		// TODO something with this
	}

	for _, file := range files {
		for _, tag := range file.Tags() {
			if tag == tagArg {
				outFiles = append(outFiles, file)
			}
		}
	}

	return &filesListResult{c.args, outFiles}
}

type filesListResult struct {
	args     []string
	outFiles []*logic.File
}

func (r filesListResult) Fmt() (string) {

	if len(r.args) < 1 {
		return "Too few arguments."
	} else if len(r.args) > 1 {
		return "Too many arguments."
	} else {
		var buffer bytes.Buffer

		buffer.WriteString(fmt.Sprintf("Files with tag '%s'\n", r.args[0]))
		for _, file := range r.outFiles {
			buffer.WriteString(fmt.Sprintf("  %s\n", file.FileName()))
		}

		return buffer.String()
	}
}

func (r filesListResult) GetSuccess() (bool) {
	return false
}