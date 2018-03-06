
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
)

// NewNullCommand creates a pointer to a NullCommand.
func NewNullCommand(command []string) *NullCommand {
	return &NullCommand{command}
}

// NullCommand reports the invalid command that was issued.
type NullCommand struct {
	args []string
}

func (c NullCommand) Execute() SubResult {

	return nullResult{c.args[0], c.args[1:]}
}

type nullResult struct {
	command string
	args []string
}

func (r nullResult) Fmt() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Invalid command '%s' with args: [", r.command))
	for i, arg := range r.args {
		if i != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf(" %s", arg))
	}
	buffer.WriteString(" ].")
	return buffer.String()
}

func (r nullResult) GetSuccess() bool {
	return false
}