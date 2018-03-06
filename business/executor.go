
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

// ExecutorI is an interface for executing commands.
type ExecutorI interface {
	Execute() (*Result)
	Command() (Command)
}

// NewExecutor creates a pointer to an Executor.
func NewExecutor(executable Command) (*Executor){
	return &Executor{executable}
}

// Executor is the implementation of ExecutorI.
type Executor struct {
	command Command
}

// Execute executes a command and packages it's result into a Result.
func (c Executor) Execute() (*Result) {
	subResult := c.command.Execute()
	return &Result{subResult}
}

// Command gets the command of this executor.
func (c Executor) Command() (Command) {
	return c.command
}