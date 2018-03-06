
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


// TagListEntityGatewayI is the interface for the tag list entity data source.
type TagListEntityGatewayI interface {
	AddTagListEntity(fileId int, tagId int, filePath string) (int, error)
	ReadTags(fileId int, filePath string) ([]*TagListEntityDTO, error)
	ReadFiles(tagId int, filePath string) ([]TagListEntityDTO, error)
	RemoveTagListEntity(fileId int, tagId int, filePath string) (error)
}

// NewTagListEntityGateway provies a pointer to a TagListEntityGateway.
func NewTagListEntityGateway() (*TagListEntityGateway) {
	return &TagListEntityGateway{}
}

// TagListEnttiyGateway is the implementation of the TagListEntityGateway
// that provides access to the tag list entity gateway data source.
type TagListEntityGateway struct {}

// AddTagListEntity adds a tag list entity to the data source found at the
// file path. AddTagListEntity returns the id of the entity in the data source,
//  and an error if on is encountered in the transaction.
func (t TagListEntityGateway) AddTagListEntity(fileId int, tagId int, filePath string) (int, error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return 0, err
	}

	stmt, err := engine.DataEngine.Prepare("INSERT INTO tag_list_entity(file_id, tag_id) VALUES (?,?)")

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(fileId, tagId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// ReadTags reads the tag list entities with the given file id, and returns any errors that
// are encountered in the transaction.
func (t TagListEntityGateway) ReadTags(fileId int, filePath string) ([]*TagListEntityDTO, error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return nil, err
	}

	rows, err := engine.DataEngine.Query("SELECT * FROM tag_list_entity WHERE file_id=?", fileId)

	if err != nil {
		return nil, err
	}

	var result []*TagListEntityDTO

	for rows.Next() {
		var id int
		var fileId int
		var tagId int

		err = rows.Scan(&id, &fileId, &tagId)
		if err != nil {
			return nil, err
		}

		result = append(result, &TagListEntityDTO{id, fileId, tagId})
	}

	return result, nil
}

// ReadFiles reads the tag list entities with the given tag id, and returns any errors that
// are encountered in the transaction.
func (t TagListEntityGateway) ReadFiles(tagId int, filePath string) ([]TagListEntityDTO, error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return nil, err
	}

	rows, err := engine.DataEngine.Query("SELECT * FROM tag_list_entity WHERE tag_id=?", tagId)

	if err != nil {
		return nil, err
	}

	var result []TagListEntityDTO

	for rows.Next() {
		var id int
		var fileId int
		var tagId int

		err = rows.Scan(&id, &fileId, &tagId)
		if err != nil {
			return nil, err
		}

		result = append(result, TagListEntityDTO{id, fileId, tagId})
	}

	return result, nil
}

// RemoveTagListEntity removes the tag list entity from the data source, and returns any errors
// encountered in the transaction.
func (t TagListEntityGateway) RemoveTagListEntity(fileId int, tagId int, filePath string) (error) {

	engine, err := getInstance(filePath)

	if err != nil {
		return err
	}

	stmt, err := engine.DataEngine.Prepare("DELETE FROM tag_list_entity WHERE file_id=? AND tag_id=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(fileId, tagId)

	if err != nil {
		return err
	}

	return nil
}
