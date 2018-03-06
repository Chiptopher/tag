
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

// InitializerI is an interface for initializing commands.
type InitializerI interface {
	Parse(args []string) (ExecutorI)
}

// NewInitializer creates a pointer to an Initializer.
func NewInitializer(filePath string, transaction TransactionI) (*Initializer) {
	return &Initializer{filePath, transaction}
}

// Initializer is the implementation of InitializerI.
type Initializer struct {
	filePath string
	transaction TransactionI
}

// Parse generates an ExecutorI with the appropriate Command.
func (a Initializer) Parse(args []string) (ExecutorI) {

	if len(args) == 0 || args[0] == "help" {
		return NewExecutor(NewNoArgsCommand())
	}

	if args[0] == "init" {
		return NewExecutor(NewCreateDataFolderCommand(args[1:]))
	} else if args[0] == "file" {
		return NewExecutor(NewTagFileCommand(args[1:], a.transaction))
	} else if args[0] == "list" {
		return NewExecutor(NewListTagsCommand(args[1:], a.transaction))
	} else if args[0] == "files" {
		return NewExecutor(NewFilesListCommand(args[1:], a.transaction))
	} else if args[0] == "remove" {
		return NewExecutor(NewRemoveTagCommand(args[1:], a.transaction))
	} else {
		return NewExecutor(NewNullCommand(args))
	}
}