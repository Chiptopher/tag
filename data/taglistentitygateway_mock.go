
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

type mockTagListEntityGateway struct {
	mock.Mock
}

func MockNewTagListEntityGateway() (*mockTagListEntityGateway) {
	return &mockTagListEntityGateway{}
}

func (m mockTagListEntityGateway) AddTagListEntity(fileId int, tagId int, filePath string) (int, error) {
	args := m.Called(fileId, tagId, filePath)
	return args.Int(0), args.Error(1)
}

func (m mockTagListEntityGateway) ReadTags(fileId int, filePath string) ([]*TagListEntityDTO, error) {
	args := m.Called(fileId, filePath)
	return args.Get(0).([]*TagListEntityDTO), args.Error(1)
}

func (m mockTagListEntityGateway) ReadFiles(tagId int, filePath string) ([]TagListEntityDTO, error) {
	args := m.Called(tagId, filePath)
	return args.Get(0).([]TagListEntityDTO), args.Error(1)
}

func (m mockTagListEntityGateway) RemoveTagListEntity(fileId int, tagId int, filePath string) (error) {
	args := m.Called(fileId, tagId, filePath)
	return args.Error(0)
}