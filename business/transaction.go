
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

import "github.com/Chiptopher/tag/logic"

// TransactionI is an interface for managing changes that occur during the
// life of a system.
type TransactionI interface {
	Commit() (error)
	GetTagIdentityMap() (logic.TagIdentityMapI)
	GetTagMapper() (logic.TagMapperI)
	GetFileIdentityMap() (logic.FileIdentityMapI)
	GetFileMapper() (logic.FileMapperI)
	GetFileById(id int) (*logic.File, error)
	GetFileByName(name string) (*logic.File, error)
	GetAllFiles() ([]*logic.File, error)
	GetTagById(id int) (*logic.Tag, error)
	GetTagByName(name string) (*logic.Tag, error)
	RegisterFile(file *logic.File)
	GetCWD() (string)
}

// NewTransaction creates a pointer to a Transaction.
func NewTransaction(cwd string) (*Transaction) {
	tagMapper := logic.NewTagMapper(cwd)
	tagMap := logic.NewTagIdentityMap(tagMapper)
	fileMapper := logic.NewFileMapper(tagMap, cwd)
	fileMap := logic.NewFileIdentityMap(fileMapper)
	m := make(map[int]*logic.File)
	t := &Transaction{tagMapper, fileMapper, tagMap, fileMap, nil, []*logic.File{}, make(map[int]*logic.Tag), []*logic.Tag{}, cwd}
	t.fileOriginal = m
	return t
}

// Transaction is the implementation of TransactionI.
type Transaction struct {
	tagMapper    logic.TagMapperI
	fileMapper   logic.FileMapperI
	tagMap       logic.TagIdentityMapI
	fileMap      logic.FileIdentityMapI
	fileOriginal map[int]*logic.File
	addedFiles   []*logic.File
	tagOriginal  map[int]*logic.Tag
	addedTags    []*logic.Tag
	cwd          string
}

// Commit commits changes that have occured during the life of the
// transaction.
func (t *Transaction) Commit() (error) {
	for _, file := range t.addedFiles {
		_, err := t.fileMapper.SaveFile(file)
		if err != nil {
			return err
		}
	}

	fileMap := t.fileMap.GetMap()

	for _, f := range fileMap {
		_, err := t.fileMapper.SaveFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetTagIdentityMap gets the TagIdentityMapI associated with this Transaction.
func (t *Transaction) GetTagIdentityMap() (logic.TagIdentityMapI) {
	return t.tagMap
}

// GetTagMapper reports the TagMapperI associated with this Transaction.
func (t *Transaction) GetTagMapper() (logic.TagMapperI) {
	return t.tagMapper
}

// GetFileIdentityMap reports the FileIdentityMapI associated with this Transaction.
func (t *Transaction) GetFileIdentityMap() (logic.FileIdentityMapI) {
	return t.fileMap
}

// GetFileMapper reports the FileMapperI associated with this Transaction.
func (t *Transaction) GetFileMapper() (logic.FileMapperI) {
	return t.fileMapper
}

// GetFileById gets the File with the given id, and any errors encountered.
func (t *Transaction) GetFileById(id int) (*logic.File, error) {
	return t.loadFile(t.fileMap.GetFileById(id))
}

// GetFileByName gets the File with the given name, and any errors encountered.
func (t *Transaction) GetFileByName(name string) (*logic.File, error) {
	return t.loadFile(t.fileMap.GetFileByName(name))

}

// GetAllFiles reports all the Files in the system.
func (t *Transaction) GetAllFiles() ([]*logic.File, error) {
	return t.fileMap.GetAllFiles()
}

// GetTagById gets the Tag with the given id, and any errors encountered.
func (t *Transaction) GetTagById(id int) (*logic.Tag, error) {
	return t.loadTag(t.tagMap.GetTagById(id))
}

// GetTagByName gets the Tag with the given name, and any errors encountered.
func (t *Transaction) GetTagByName(name string) (*logic.Tag, error) {
	return t.loadTag(t.tagMap.GetTagByName(name))
}

func (t *Transaction) loadTag(tag *logic.Tag, err error) (*logic.Tag, error) {
	if err != nil {
		return nil, err
	}

	if t.tagOriginal[tag.Id()] == nil {
		t.tagOriginal[tag.Id()] = tag.Copy()
	}

	return tag, err
}

func (t *Transaction) loadFile(file *logic.File, err error) (*logic.File, error) {

	if err != nil {
		return nil, err
	}

	if t.fileOriginal[file.Id()] == nil {
		t.fileOriginal[file.Id()] = file.Copy()
	}

	return file, nil
}

// RegisterFile registers the given File with the transaction as a created File.
func (t *Transaction) RegisterFile(file *logic.File) {
	t.addedFiles = append(t.addedFiles, file)
}

// GetCWD gets the directory the transaction is being executed under.
func (t *Transaction) GetCWD() (string) {
	return t.cwd
}
