
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

import "github.com/stretchr/testify/mock"

type fileMapperMock struct {
	mock.Mock
}

func MockNewFileMapper() (*fileMapperMock) {
	return new(fileMapperMock)
}

func (m fileMapperMock) getFile(id int) (*File, error) {
	args := m.Called(id)
	return args.Get(0).(*File), args.Error(1)
}

func (m *fileMapperMock) SaveFile(file *File) (int, error) {
	args := m.Called(file)
	return args.Int(0), args.Error(1)
}

func (m fileMapperMock) RemoveFile(id int) (error) {
	args := m.Called(id)
	return args.Error(0)
}

func (m fileMapperMock) getFileByName(name string) (*File, error) {
	args := m.Called(name)
	return args.Get(0).(*File), args.Error(1)
}

func (m fileMapperMock) getCWD() (string) {
	args := m.Called()
	return args.Get(0).(string)
}