package database

import (
	"errors"
	"fmt"
	"github.com/tangent-c/tangentdb/env"
)

type Type struct {}

type dbColumn struct {
	name string
	type_ Type 
}

type dbTable struct {
	name string
	columnsCount int32
	columns []dbColumn
}

type TheDb struct {
	name string
	tableCount int32
	tables []dbTable
}

func (db *TheDb) addTable(name string, columns int32) error {
	if len(name) < 1 {
		return errors.New("empty name")
	}
	
	db.tables = append(db.tables, dbTable{name, columns, make([]dbColumn, columns)})

	return nil
}

func OpenDb(path string) (*TheDb, error) {
	if len(path) <= 0 {
		return &TheDb{}, errors.New(fmt.Errorf("failed to open db: %s", path).Error())
	}

	var theDb *TheDb
	var err error

	if !env.CreateDbIfNotExist {
		return nil, errors.New("Db does not exist")
	}

	if env.CreateDbIfNotExist {
		theDb, err = createDb(path)
		if err != nil {
			return nil, err
		}
	}

	theDb.name = ""

	// TODO: Implement the db open code, need to use file system abstractions later

	return theDb, nil
}

func createDb(path string) (*TheDb, error) {
	// TODO: dependent: storage layer

	
	return nil, nil
}