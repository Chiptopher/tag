
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


package app

import (
	"github.com/Chiptopher/tag/business"
	"github.com/Chiptopher/tag/data"
	"fmt"
)

// Run executes a command based on the given arguments and the current
// working directory, and formats the result.
func Run(args []string, cwd string) (string) {

	transaction := business.NewTransaction(cwd)

	initializer := business.NewInitializer(cwd, transaction)
	command := initializer.Parse(args)
	result := command.Execute()

	err := transaction.Commit()

	if err == nil {
		return result.Fmt()
	} else {
		return process(err)
	}
}

func process(err error) string {
	if v, ok := err.(*data.NoFileExistsError); ok {
		return fmt.Sprintf("File %s doesn't exist.", v.FilePath())
	}
	panic(err)
}