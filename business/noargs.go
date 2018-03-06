
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

const output = "Tag is a command line tool that makes keeping track of files and their\n" +
	"purposes easier. It has the follow commands available:\n\n" +
	"  init <PATH>                    Path to the root of tags.\n" +
	"  file <FILE NAME> <TAG NAME>    Add TAG NAME to the FILE NAME.\n"+
	"  list <FILE NAME>               List all tags associated with FILE NAME.\n"+
	"  files <TAG NAME>               List files with the TAG NAME associated with it.\n"+
	"  remove <FILE NAME> <TAG NAME>  Remove the TAG NAME from FILE NAME.\n"

// NewNoArgsCommand creates a pointer to a NoArgsCommand.
func NewNoArgsCommand() (*NoArgsCommand) {
	return &NoArgsCommand{}
}

// NoArgsCommand lists the help description of this program.
type NoArgsCommand struct {}

func (c NoArgsCommand) Execute() (SubResult) {
	return &noArgResult{}
}

type noArgResult struct {}

func (r noArgResult) Fmt() (string) {
	return output
}

func (r noArgResult) GetSuccess() (bool) {
	return false
}