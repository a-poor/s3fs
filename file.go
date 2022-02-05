package main

import (
	"errors"
	"os"
	"time"
)

var (
	ErrLockNotSupported = errors.New("lock not supported by s3")
)

// s3File implements billy.File
type s3File struct {
}

// Name returns the name of the file as presented to Open.
func (f *s3File) Name() string {
	return ""
}

// Write implements os.Writer for billy.File
func (f *s3File) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Read implements os.Reader for billy.File
func (f *s3File) Read(p []byte) (n int, err error) {
	return 0, nil
}

// ReadAt implements io.ReaderAt for billy.File
func (f *s3File) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

// Seek implements io.Seeker for billy.File
func (f *s3File) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// Close implements io.Closer for billy.File
func (f *s3File) Close() error {
	return nil
}

// Lock locks the file like e.g. flock. It protects against access from
// other processes.
func (f *s3File) Lock() error {
	return ErrLockNotSupported
}

// Unlock unlocks the file.
func (f *s3File) Unlock() error {
	return ErrLockNotSupported
}

// Truncate the file.
func (f *s3File) Truncate(size int64) error {
	return nil
}

// s3FileInfo implements os.FileInfo
type s3FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi s3FileInfo) Name() string {
	return fi.name
}

func (fi s3FileInfo) Size() int64 {
	return fi.size
}

func (fi s3FileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi s3FileInfo) IsDir() bool {
	return fi.mode.IsDir()
}

func (fi s3FileInfo) Sys() interface{} {
	return nil
}
