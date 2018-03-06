
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

import "path/filepath"

// NewRemoveTagCommand creates a pointer to a RemoveTagCommand.
func NewRemoveTagCommand(args []string, transaction TransactionI) *RemoveTagCommand {
	return &RemoveTagCommand{args, transaction}
}

// RemoveTagCommand removes a Tag from a File.
type RemoveTagCommand struct {
	args        []string
	transaction TransactionI
}

func (c RemoveTagCommand) Execute() SubResult {

	if len(c.args) != 2 {
		return &removeTagResult{false, c.args}
	}

	fileName := c.args[0]
	fullFileName := filepath.Join(c.transaction.GetCWD(), fileName)

	file, err := c.transaction.GetFileByName(fullFileName)

	if err != nil {
		// TODO this
	}

	tagName := c.args[1]
	tag, err := c.transaction.GetTagByName(tagName)

	if err != nil {
		// TODO this
	}

	file.RemoveTag(tag)

	return removeTagResult{true, nil}
}

type removeTagResult struct {
	success bool
	args    []string
}

func (r removeTagResult) Fmt() string {
	if r.args != nil && len(r.args) < 2 {
		return "Too few arguments."
	} else if r.args != nil && len(r.args) > 2 {
		return "Too many arguments."
	}
	return ""
}

func (r removeTagResult) GetSuccess() bool {
	return r.success
}