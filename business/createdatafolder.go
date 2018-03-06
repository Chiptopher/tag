
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
	fmt "fmt"
	"github.com/Chiptopher/tag/data"
)

// NewCreateDataFolderCommand creates a pointer to a CreateDataFolderCommand.
func NewCreateDataFolderCommand(args []string) (*CreateDataFolderCommand) {
	return &CreateDataFolderCommand{args, data.CreateDataFolder}
}

// CreateDataFolderCommand is the command for creating a data source folder.
type CreateDataFolderCommand struct {
	args []string
	createStrategy func(string) (bool, error)
}

func (c CreateDataFolderCommand) Execute() (SubResult) {

	if len(c.args) > 1 {
		return &createDataFolderResult{nil, c.args[1:], false}
	} else if len(c.args) < 1 {
		return &createDataFolderResult{nil, []string{}, false}
	} else {
		_, err := c.createStrategy(c.args[0])
		return &createDataFolderResult{err, nil, true}
	}
}

type createDataFolderResult struct {
	Err     error
	Args    []string
	Success bool
}

func (r createDataFolderResult) Fmt() (string) {
	if len(r.Args) > 0 {
		return fmt.Sprintf("Too many arguments entered: %s", r.Args)
	} else if len(r.Args) == 0 && r.Args != nil {
		return "Not enough arguments entered. Path to init required."
	} else {
		return ""
	}
}

func (r createDataFolderResult) GetSuccess() (bool) {
	return r.Success
}