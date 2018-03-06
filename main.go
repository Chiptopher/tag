
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


// Package main provides the entrance to the application.
package main

import (
	"os"
	"log"
	"github.com/Chiptopher/tag/app"
	"fmt"
	"os/user"
	"path/filepath"
	"runtime/debug"
)

func main() {

	usr, _ := user.Current()
	f, _ := os.OpenFile(filepath.Join(usr.HomeDir, ".tag_errors"), os.O_WRONLY | os.O_CREATE, 0666)
	defer f.Close()
	log.SetOutput(f)

	// When a panic occurs, write it to the log.
	defer func() {
		if recover() != nil {
			log.Println(fmt.Sprintf("\n --------------\n %s \n -------------- \n", debug.Stack()))
			fmt.Printf("Error encountered. Info dumped to ~/.tag_errors.\n")
			fmt.Printf("\n")
		}
	}()

	cwd, _ := os.Getwd()
	result := app.Run(os.Args[1:], cwd)
	fmt.Println(result)

}