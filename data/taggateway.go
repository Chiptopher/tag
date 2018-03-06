
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

import (
	_ "github.com/mattn/go-sqlite3"
)


// TagGatewayI is the interface for the file data source.
type TagGatewayI interface {
	AddTag(Name string, filePath string) (int, error)
	RemoveTag(id int, filePath string) (error)
	ReadTag(id int, filePath string) (*TagDataDTO, error)
	SaveTag(id int, name string, filePath string) (error)
	GetTagIdByName(name string, filePath string) (int, error)
}

// NewTagGatewy provides a pointer to a TagGateway.
func NewTagGateway() (*TagGateway) {
	return &TagGateway{}
}

// TagGateway is the implementation of TagGatewayI that provides access to the tag
// data source.
type TagGateway struct {}

// AddTag adds a tag to the data source found at the file path. It returns the id of the
// tag in the data source, and an error if one is encountered in the transaction.
func (t TagGateway) AddTag(Name string, filePath string) (int, error) {
	engine, err := getInstance(filePath)
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	stmt, err := engine.DataEngine.Prepare("INSERT INTO tags(name) VALUES (?)")

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(Name)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// RemoveTag removes the tag with the given id from the data source found at the file
// path. RemoveTag returns any error encountered in the transaction.
func (t TagGateway) RemoveTag(id int, filePath string) error {

	engine, err := getInstance(filePath)

	if err != nil {
		return err
	}

	stmt, err := engine.DataEngine.Prepare("DELETE FROM tags where id=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// ReadTag reads the tag with the given id from the data source at the given file path.
// ReadTag returns any error encountered in the transaction.
func (t TagGateway) ReadTag(id int, filePath string) (*TagDataDTO, error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return nil, nil
	}

	rows, err := engine.DataEngine.Query("SELECT * FROM tags WHERE id=?", id)

	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	tag := TagDataDTO{}
	rows.Next()
	rows.Scan(&tag.Id, &tag.Name)

	rows.Close()

	return &tag, nil
}

// SaveTag saves updates to the tag with the given id at the given file path. SaveTag
// returns any error encountered in the transaction.
func (t TagGateway) SaveTag(id int, name string, filePath string) (error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return err
	}

	stmt, err := engine.DataEngine.Prepare("update tags set name=? where id=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, id)

	if err != nil {
		return err
	}

	return nil
}

// GetTagIdByName reads the id of the tag with the given name from the data source at the
// given file path.
func (t TagGateway) GetTagIdByName(name string, filePath string) (int, error) {

	engine, err := getInstance(filePath)
	rows, err := engine.DataEngine.Query("SELECT id FROM tags WHERE name=?", name)

	if err != nil {
		return 0, &UnexpectedError{err}
	}
	defer rows.Close()

	rows.Next()
	id := 0

	err = rows.Scan(&id)
	if err != nil {
		return 0, &TagNotFoundError{filePath}
	}
	rows.Close()

	return id, nil
}