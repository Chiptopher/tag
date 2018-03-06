
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
	"github.com/stretchr/testify/mock"
	"github.com/Chiptopher/tag/logic"
)

func MockNewTransaction() (*MockTransaction) {
	return &MockTransaction{}
}

type MockTransaction struct {
	mock.Mock
}

func (m *MockTransaction) Commit() (error) {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTransaction) GetTagIdentityMap() (logic.TagIdentityMapI) {
	args := m.Called()
	return args.Get(0).(logic.TagIdentityMapI)
}

func (m *MockTransaction) GetTagMapper() (logic.TagMapperI) {
	args := m.Called()
	return args.Get(0).(logic.TagMapperI)
}

func (m *MockTransaction) GetFileIdentityMap() (logic.FileIdentityMapI) {
	args := m.Called()
	return args.Get(0).(logic.FileIdentityMapI)
}
func (m *MockTransaction) GetFileMapper() (logic.FileMapperI) {
	args := m.Called()
	return args.Get(0).(logic.FileMapperI)
}

func (m *MockTransaction) GetFileById(id int) (*logic.File, error) {
	args := m.Called(id)
	return args.Get(0).(*logic.File), args.Error(1)
}
func (m *MockTransaction) GetFileByName(name string) (*logic.File, error) {
	args := m.Called(name)
	file := args.Get(0)
	if file == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*logic.File), args.Error(1)
}

func (m *MockTransaction) GetAllFiles() ([]*logic.File, error) {
	args := m.Called()
	return args.Get(0).([]*logic.File), args.Error(1)
}

func (m *MockTransaction) GetTagById(id int) (*logic.Tag, error) {
	args := m.Called(id)
	return args.Get(0).(*logic.Tag), args.Error(1)
}

func (m *MockTransaction) GetTagByName(name string) (*logic.Tag, error) {
	args := m.Called(name)
	return args.Get(0).(*logic.Tag), args.Error(1)
}

func (m *MockTransaction) RegisterFile(file *logic.File) {
	m.Called(file)
}

func (m *MockTransaction) RegisterTag(tag *logic.Tag) {
	m.Called(tag)
}

func (m *MockTransaction) GetCWD() (string) {
	return m.Called().Get(0).(string)
}