
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
	"testing"
	"database/sql"
	"os"
	"github.com/Chiptopher/tag/test"
	"path/filepath"
)

func TestGetInstance(t *testing.T) {
	CreateDataFolder(test.GoPathTestFolder)

	engine, err := getInstance(test.GoPathTestFolder)

	os.RemoveAll(filepath.Join(test.GoPathTestFolder, DataFolderName))

	if err != nil {
		t.Fatalf("GetInstance errored: %s", err)
	}
	if engine.FilePath == "" || engine.DataEngine == nil{
		t.Fatalf("engine incorrectly initialized.")
	}
}

func TestGetInstance_CreatesOnlyOneInstance(t *testing.T) {

	CreateDataFolder(test.GoPathTestFolder)

	engine, _ := getInstance(test.GoPathTestFolder)
	engine2, _ := getInstance(test.GoPathTestFolder)

	os.RemoveAll(filepath.Join(test.GoPathTestFolder, DataFolderName))

	if engine != engine2 {
		t.Fatalf("engine does not create singleton engine.")
	}
}

func TestCreateDatabase(t *testing.T) {

	os.Remove(test.CreatePathInTestFolder(DatabaseName))

	actual, err := CreateDatabase(test.GoPathTestFolder)
	var expected = filepath.Join(test.GoPathTestFolder, DatabaseName)

	if actual != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}


	if err != nil {
		t.Fatalf("Create failed, %s.", err)
	}

	db, err := sql.Open("sqlite3", test.CreatePathInTestFolder(DatabaseName))
	if err != nil {
		t.Fatalf("Database file not created.")
	}
	_, err = db.Query("SELECT * FROM tags")
	if err != nil {
		t.Fatalf("tags table not created %s.", err)
	}
	_, err = db.Query("SELECT * FROM files")
	if err != nil {
		t.Fatalf("files table not created.")
	}
	_, err = db.Query("SELECT * FROM tag_list_entity")
	if err != nil {
		t.Fatalf("tag_list_entity table not created.")
	}
	err = db.Close()
	if err != nil {
		t.Fatalf("Couldn't close the DB.")
	}
}