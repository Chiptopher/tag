
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

package logic

import (
	"testing"
)

func TestTagToString(t *testing.T) {
	actual := Tag{1, "name", entityStateClean}.ToString()
	expected := "Tag: 1, name"

	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestTag_Copy(t *testing.T) {

	tag := NewTag(1, "name")
	result := tag.Copy()

	if tag == result {
		t.Fatalf("Should not reference the same thing.")
	}

	if result.id != 1 {
		t.Fatalf("Expected %d but got %d", 1, result.id)
	}

	if result.name != "name" {
		t.Fatalf("Expected %s but got %s", "name", result.name)
	}
}

func TestTag_SetName_ChangesEntityState(t *testing.T) {

	tag := NewTag(1, "name")
	tag.SetName("Name2")

	if tag.getEntityState() != entityStateDirty {
		t.Fatalf("entity state should change.")
	}
}