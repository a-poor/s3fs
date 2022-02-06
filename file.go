package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/atomic"
)

const (
	ModeMultipartUpload os.FileMode = fs.ModePerm + 1 // Custom os.FileMode for S3 multipart upload
)

var (
	ErrLockNotSupported      = errors.New("lock not supported by s3")
	ErrTruncateNotSupported  = errors.New("truncate not supported by s3")
	ErrFileClosed            = errors.New("file is closed")
	ErrCantWriteToReadOnly   = errors.New("can't write to read-only file")
	ErrCantReadFromWriteOnly = errors.New("can't read from write-only file")
)

// s3ReadFile implements billy.File for S3, and represents a file opened in read mode.
//
// Upon creation, the file is loaded from S3.
type s3ReadFile struct {
	client *s3.Client    // s3 skd client
	bucket string        // S3 bucket name
	key    string        // File object's key in S3
	closed bool          // Is the file closed?
	fbody  io.ReadCloser // File's Body from AWS S3 SDK response
}

// newS3ReadFile creates a new s3ReadFile.
func newS3ReadFile(client *s3.Client, bucket, key string) (*s3ReadFile, error) {
	// Check if the file exists

	// Create the context
	ctx := context.TODO() // TODO: How can user-supplied contexts be supported?

	// Run the GetObject operation
	res, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to perform GetObject operation: %w", err)
	}

	// Return the file
	return &s3ReadFile{
		client: client,
		bucket: bucket,
		key:    key,
		fbody:  res.Body,
	}, nil
}

// Name returns the name of the file as presented to Open.
func (f *s3ReadFile) Name() string {
	return f.key
}

// Write implements os.Writer for billy.File
func (f *s3ReadFile) Write(p []byte) (n int, err error) {
	return 0, ErrCantWriteToReadOnly
}

// Read implements os.Reader for billy.File
func (f *s3ReadFile) Read(p []byte) (n int, err error) {
	return f.fbody.Read(p)
}

// ReadAt implements io.ReaderAt for billy.File
func (f *s3ReadFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, errors.New("unsupported")
}

// Seek implements io.Seeker for billy.File
func (f *s3ReadFile) Seek(offset int64, whence int) (int64, error) {
	return 0, errors.New("unsupported")
}

// Close implements io.Closer for billy.File
func (f *s3ReadFile) Close() error {
	// Was the file already closed?
	if f.closed {
		return ErrFileClosed
	}

	// Mark the file as closed
	defer func() { f.closed = true }()

	// Close the underlying file
	err := f.fbody.Close()

	// Was there an error returned?
	// TODO: Check or wrap the error?
	if err != nil {
		return err
	}
	return nil
}

// Lock locks the file like e.g. flock. It protects against access from
// other processes.
func (f *s3ReadFile) Lock() error {
	return ErrLockNotSupported
}

// Unlock unlocks the file.
func (f *s3ReadFile) Unlock() error {
	return ErrLockNotSupported
}

// Truncate the file.
func (f *s3ReadFile) Truncate(size int64) error {
	return ErrTruncateNotSupported
}

// s3WriteFile implements billy.File
type s3WriteFile struct {
	client *s3.Client   // s3 skd client
	bucket string       // S3 bucket name
	key    string       // File object's key in S3
	closed bool         // Is the file closed?
	buf    bytes.Buffer // Buffer for storing the file before it's uploaded
}

// newS3WriteFile creates a new s3ReadFile.
func newS3WriteFile(client *s3.Client, bucket, key string) (*s3WriteFile, error) {
	return &s3WriteFile{
		client: client,
		bucket: bucket,
		key:    key,
	}, nil
}

// Name returns the name of the file as presented to Open.
func (f *s3WriteFile) Name() string {
	return f.key
}

// Write implements os.Writer for billy.File
func (f *s3WriteFile) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Read implements os.Reader for billy.File
func (f *s3WriteFile) Read(p []byte) (n int, err error) {
	return 0, ErrCantReadFromWriteOnly
}

// ReadAt implements io.ReaderAt for billy.File
func (f *s3WriteFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

// Seek implements io.Seeker for billy.File
func (f *s3WriteFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// Close implements io.Closer for billy.File
func (f *s3WriteFile) Close() error {
	if f.closed {
		return ErrFileClosed
	}
	return nil
}

// Lock locks the file like e.g. flock. It protects against access from
// other processes.
func (f *s3WriteFile) Lock() error {
	return ErrLockNotSupported
}

// Unlock unlocks the file.
func (f *s3WriteFile) Unlock() error {
	return ErrLockNotSupported
}

// Truncate the file.
func (f *s3WriteFile) Truncate(size int64) error {
	return ErrTruncateNotSupported
}

// s3MultipartUploadFile implements billy.File
type s3MultipartUploadFile struct {
	client  *s3.Client    // s3 skd client
	bucket  string        // S3 bucket name
	key     string        // File object's key in S3
	closed  bool          // Is the file closed?
	uploadN atomic.Uint32 // Counter tracking the number of uploads
}

// newS3MultipartUploadFile creates a new s3ReadFile.
func newS3MultipartUploadFile(client *s3.Client, bucket, key string) (*s3MultipartUploadFile, error) {
	return &s3MultipartUploadFile{
		client: client,
		bucket: bucket,
		key:    key,
	}, nil
}

// Name returns the name of the file as presented to Open.
func (f *s3MultipartUploadFile) Name() string {
	return f.key
}

// Write implements os.Writer for billy.File
func (f *s3MultipartUploadFile) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Read implements os.Reader for billy.File
func (f *s3MultipartUploadFile) Read(p []byte) (n int, err error) {
	return 0, ErrCantReadFromWriteOnly
}

// ReadAt implements io.ReaderAt for billy.File
func (f *s3MultipartUploadFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

// Seek implements io.Seeker for billy.File
func (f *s3MultipartUploadFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// Close implements io.Closer for billy.File
func (f *s3MultipartUploadFile) Close() error {
	if f.closed {
		return ErrFileClosed
	}
	return nil
}

// Lock locks the file like e.g. flock. It protects against access from
// other processes.
func (f *s3MultipartUploadFile) Lock() error {
	return ErrLockNotSupported
}

// Unlock unlocks the file.
func (f *s3MultipartUploadFile) Unlock() error {
	return ErrLockNotSupported
}

// Truncate the file.
func (f *s3MultipartUploadFile) Truncate(size int64) error {
	return ErrTruncateNotSupported
}
