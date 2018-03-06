
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
	"os"
)

// FileGateway is the interface to accessing the tag data source.
type FileGatewayI interface {
	AddFile(name string, filePath string) (int, error)
	RemoveFile(id int, filePath string) error
	ReadFile(id int, filePath string) (*FileDataDTO, error)
	SaveFile(id int, name string, filePath string) (error)
	GetFileIdByName(name string, filePath string) (int, error)
	GetAllFileIds(filePath string) ([]int, error)
}

// NewFileGateway creates a pointer to a FileGateway
func NewFileGateway() (*FileGateway) {
	return &FileGateway{}
}

// FileGateway is the implementation of FileGatewayI that provides access to the file
// data source.
type FileGateway struct {}

// AddFile adds a file with the given name to the data source found at the given file
// path. AddFile returns the id of the file in the data source if added successfully,
// or the error otherwise.
func (f FileGateway) AddFile(name string, filePath string) (int, error) {

	if _, err := os.Stat(name); os.IsNotExist(err) {
		return 0, &NoFileExistsError{name}
	}

	engine, err := getInstance(filePath)

	if err != nil {
		return 0, nil
	}

	stmt, err := engine.DataEngine.Prepare("INSERT INTO files(file_name) VALUES (?)")

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name)

	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// RemoveFile removes the file with the given id from the data source found at the given
// file path. RemoveFile returns any error encountered in the transaction.
func (f FileGateway) RemoveFile(id int, filePath string) error {

	engine, err := getInstance(filePath)

	if err != nil {
		return err
	}

	stmt, err := engine.DataEngine.Prepare("DELETE FROM files WHERE id=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// ReadFile reads the file with the given id from the data source found at the given path.
// ReadFile returns the data or any errors encountered during the transaction.
func (f FileGateway) ReadFile(id int, filePath string) (*FileDataDTO, error) {

	engine, err := getInstance(filePath)

	rows, err := engine.DataEngine.Query("SELECT * FROM files WHERE id=?", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	file := FileDataDTO{}

	rows.Next()
	err = rows.Scan(&file.Id, &file.Name)
	if err != nil {
		return nil, err
	}
	rows.Close()

	return &file, nil
}

// SaveFile updates a file with the given id in the data source found at the given path.
// SaveFile returns the any errors encountered in the transaction.
func (f FileGateway) SaveFile(id int, name string, filePath string) (error) {
	engine, err := getInstance(filePath)

	if err != nil {
		return err
	}

	stmt, err := engine.DataEngine.Prepare("update files set file_name=? where id=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, id)

	if err != nil {
		return err
	}

	return nil
}

// GetFileIdByName gets the id of the file with the given name in the data source at the given
// file path. GetFileIdByName returns the id and any errors encountered in the transaction.
func (f FileGateway) GetFileIdByName(name string, filePath string) (int, error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return 0, &UnexpectedError{err}
	}

	rows, err := engine.DataEngine.Query("SELECT id FROM files WHERE file_name=?", name)

	if err != nil {
		return 0, &UnexpectedError{err}
	}
	rows.Next()
	id := 0

	err = rows.Scan(&id)
	if err != nil {
		// it's closed
		return 0, &FileNotFoundError{name}
	}
	defer rows.Close()

	return id, nil
}

// GetAllFileIds returns all ids of files in the data source at the given file path.
func (f FileGateway) GetAllFileIds(filePath string) ([]int, error) {
	engine, err := getInstance(filePath)
	if err != nil {
		return nil, err
	}

	rows, err := engine.DataEngine.Query("SELECT id FROM files")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := []int{}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}