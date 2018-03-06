
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

package data

import "github.com/stretchr/testify/mock"

type fileGatewayMock struct{
	mock.Mock
}

func MockNewFileGateway() (*fileGatewayMock) {
	return new(fileGatewayMock)
}

func (m *fileGatewayMock) AddFile(name string, filePath string) (int, error) {
	args := m.Called(name, filePath)
	return args.Int(0), nil
}

func (m *fileGatewayMock) RemoveFile(id int, filePath string) error {
	args := m.Called(id, filePath)
	return args.Error(0)
}

func (m *fileGatewayMock) ReadFile(id int, filePath string) (*FileDataDTO, error) {
	args := m.Called(id, filePath)
	result := args.Get(0).(*FileDataDTO)
	return result, args.Error(1)
}

func (m *fileGatewayMock) SaveFile(id int, name string, filePath string) (error) {
	args := m.Called(id, name, filePath)
	return args.Error(0)
}

func (m *fileGatewayMock) GetFileIdByName(name string, filePath string) (int, error) {
	args := m.Called(name, filePath)
	return args.Int(0), args.Error(1)
}

func (m *fileGatewayMock) GetAllFileIds(filePath string) ([]int, error) {
	args := m.Called(filePath)
	return args.Get(0).([]int), args.Error(1)
}