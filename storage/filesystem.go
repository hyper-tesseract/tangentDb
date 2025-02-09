package storage

import (
	"errors"
	"os"
	"syscall"

	"github.com/tangent-c/tangentdb/common"
)

type FileSystem interface {
	OpenFile(name string, create bool)
	deleteFile(name string)
	lockFile(name string)
	unlockFile(name string)

	CreateDir(name string)
	syncDir(path string)
	deleteDir(name string)
}


type PosixFileSystem struct {}


func (p *PosixFileSystem) CreateDir(filePath string) error {
	if !common.IsValidPath(filePath) {
		return errors.New("Invalid path")
	}

	err := os.Mkdir(filePath, os.ModeDir)
	if err != nil {
		return err
	}

	return nil
}

func (p *PosixFileSystem) OpenFile(filePath string, create bool, perm os.FileMode) (*RWFile, error) {
	if !common.IsValidPath(filePath) {
		return nil, errors.New("invalid path")
	}

	flags := syscall.O_DIRECT
	if create {
		flags |= os.O_CREATE
	}

	fd, err := os.OpenFile(filePath, flags, perm)
	if err != nil {
		return nil, err
	}

	result := &RWFile{filePath, fd, perm}

	return result, nil
}



type Directory struct {
	path string
}

func (d *Directory) SyncDir() {
	
}