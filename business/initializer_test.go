
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
	"testing"
	"reflect"
)

var emptyCreateDataFolderCommand = NewCreateDataFolderCommand([]string{})
var emptyTagFileCommand = NewTagFileCommand(nil, NewTransaction("cwd"))
var emptyListTagCommand = NewListTagsCommand(nil, NewTransaction(""))
var emptyFilesListCommand = NewFilesListCommand(nil, NewTransaction(""))
var emptyRemoveTagCommand = NewRemoveTagCommand(nil, NewTransaction(""))
var emptyNullCommand = NewNullCommand(nil)
var emptyNoArgsCommand = NewNoArgsCommand()

var parsetests = []struct {
	in  []string
	out Command
}{
	{[]string{"init", "path"}, emptyCreateDataFolderCommand},
	{[]string{"init"}, emptyCreateDataFolderCommand},
	{[]string{"init", "path", "extra"}, emptyCreateDataFolderCommand},
	{[]string{"file", "A.txt", "tag"}, emptyTagFileCommand},
	{[]string{"file", "A.txt"}, emptyTagFileCommand},
	{[]string{"file", "A.txt", "tag", "extra"}, emptyTagFileCommand},
	{[]string{"list", "A.txt"}, emptyListTagCommand},
	{[]string{"list"}, emptyListTagCommand},
	{[]string{"list", "a.txt", "extra"}, emptyListTagCommand},
	{[]string{"files", "name"}, emptyFilesListCommand},
	{[]string{"files"}, emptyFilesListCommand},
	{[]string{"files", "fileName", "extra"}, emptyFilesListCommand},
	{[]string{"remove", "fileName", "tag"}, emptyRemoveTagCommand},
	{[]string{"remove", "fileName", "tag", "extra"}, emptyRemoveTagCommand},
	{[]string{"remove", "fileName"}, emptyRemoveTagCommand},
	{[]string{"invalid_command"}, emptyNullCommand},
	{[]string{}, emptyNoArgsCommand},
	{[]string{"help"}, emptyNoArgsCommand},
}

func TestInitializer_Parse(t *testing.T) {
	for _, tt := range parsetests {
		command := NewInitializer("", nil).Parse(tt.in)
		if reflect.TypeOf(command.Command()) != reflect.TypeOf(tt.out) {
			t.Fatalf("Executor has incorrect subcommand. Expected %T but got %T",
				emptyCreateDataFolderCommand, command.Command())
		}
	}
}