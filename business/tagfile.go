
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
	"path/filepath"
	"fmt"
	"github.com/Chiptopher/tag/logic"
	"github.com/Chiptopher/tag/data"
)

// NewTagFileCommand creates a pointer to a TagFileCommand.
func NewTagFileCommand(args []string, transaction TransactionI) (Command){
	return &TagFileCommand{args, transaction}
}

// TagFileCommand tags a File with a Tag.
type TagFileCommand struct {
	args        []string
	transaction TransactionI
}

func (c TagFileCommand) Execute() (SubResult) {

	if len(c.args) > 2 {
		return TagFileResult{false, c.args, "", ""}
	} else if len(c.args) < 2 {
		return TagFileResult{false, c.args, "", ""}
	} else {

		path := filepath.Join(c.transaction.GetCWD(), c.args[0])

		file, err := c.transaction.GetFileByName(path)

		if err != nil {
			if _, ok := err.(*data.FileNotFoundError); ok {
				file = logic.CreateFile(path, []*logic.Tag{})
				c.transaction.RegisterFile(file)
			} else {
				return &TagFileResult{false, nil, "", ""}
			}
		}

		tag, err := c.transaction.GetTagByName(c.args[1])

		if err != nil {
			tag = logic.CreateTag(c.args[1])
		}

		file.AddTag(tag)

		return &TagFileResult{true, nil, "", ""}
	}
}

type TagFileResult struct {
	Success	 bool
	Args	 []string
	CWD      string
	FilePath string
}

func (r TagFileResult) Fmt() (string) {
	if !r.Success {
		if r.Args != nil {
			if len(r.Args) > 2 {
				return fmt.Sprintf("Too many arguments entered: %s", r.Args)
			} else {
				return fmt.Sprintf("Not enough arguments.")
			}
		} else if r.FilePath != "" && r.CWD != "" {
			return fmt.Sprintf("File %s doesn't exist at %s", r.FilePath, r.CWD)
		}
	}
	return ""
}

func (r TagFileResult) GetSuccess() (bool) {
	return r.Success
}