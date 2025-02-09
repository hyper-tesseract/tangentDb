package storage

import (
	"os"
	"syscall"
	"github.com/tangent-c/tangentdb/env"
)

// RWFile implements the disk IO operations
// Since Direct IO is being used, so is responsible for all the
// allignments needed to perform page alligned reads and writes
// TODO: Later assess the need for sequential/prefetch behavior
// optimizations and implmentation

type RWFile struct {
	name string
	file *os.File
	perm os.FileMode
}

func (f *RWFile) Read(result []byte, offset int64) error {
	bytes, err := syscall.Pread(int(f.file.Fd()), result, offset)
	if err != nil {
		return err
	}

	return nil
}

func (f *RWFile) Write(result *[]byte) error {
	_, err := syscall.Write(int(f.file.Fd()), *result)

	if err != nil {
		return err
	}

	return nil
}

func (f *RWFile) PrepareWrite() {

}

func (f *RWFile) Close() {

}