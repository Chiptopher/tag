
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

import "testing"

// TestResult_Fmt tests that sub results that return a zero string return
// the appropriate error response.
func TestResult_Fmt_NonSuccess(t *testing.T) {
	result := Result{&subCommandMock{false}}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	output := result.Fmt()
	expectedResult := "Fatal error."
	if output != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'.", expectedResult, output)
	}
}

func TestResult_Fmt_Success(t *testing.T) {
	result := Result{&subCommandMock{true}}
	output := result.Fmt()
	expectedResult := ""
	if output != expectedResult {
		t.Fatalf("Expected '%s', but got '%s'.", expectedResult, output)
	}
}