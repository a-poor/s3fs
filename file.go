package main

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/atomic"
)

var (
	ErrLockNotSupported     = errors.New("lock not supported by s3")
	ErrTruncateNotSupported = errors.New("truncate not supported by s3")
	ErrFileClosed           = errors.New("file is closed")
)

// s3File implements billy.File
type s3File struct {
	client  *s3.Client    // s3 skd client
	bucket  string        // S3 bucket name
	key     string        // Object key / filename
	closed  atomic.Bool   // Is the file closed?
	nWrites atomic.Uint64 // Tracks the number of writes
}

// Name returns the name of the file as presented to Open.
func (f *s3File) Name() string {
	return f.key
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
	return ErrTruncateNotSupported
}
