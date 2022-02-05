// basic.go implements the interface billy.Basic

package main

import (
	"os"

	"github.com/go-git/go-billy/v5"
)

// Create implements billy.Basic
// Create creates the named file with mode 0666 (before umask), truncating
// it if it already exists. If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
func (fs *S3FS) Create(filename string) (billy.File, error) {
	return nil, nil
}

// Open opens the named file for reading. If successful, methods on the
// returned file can be used for reading; the associated file descriptor has
// mode O_RDONLY.
func (fs *S3FS) Open(filename string) (billy.File, error) {
	return nil, nil
}

// OpenFile is the generalized open call; most users will use Open or Create
// instead. It opens the named file with specified flag (O_RDONLY etc.) and
// perm, (0666 etc.) if applicable. If successful, methods on the returned
// File can be used for I/O.
func (fs *S3FS) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	return nil, nil
}

// Stat returns a FileInfo describing the named file.
func (fs *S3FS) Stat(filename string) (os.FileInfo, error) {
	return nil, nil
}

// Rename renames (moves) oldpath to newpath. If newpath already exists and
// is not a directory, Rename replaces it. OS-specific restrictions may
// apply when oldpath and newpath are in different directories.
func (fs *S3FS) Rename(oldpath, newpath string) error {
	return nil
}

// Remove removes the named file or directory.
func (fs *S3FS) Remove(filename string) error {
	return nil
}

// Join joins any number of path elements into a single path, adding a
// Separator if necessary. Join calls filepath.Clean on the result; in
// particular, all empty strings are ignored. On Windows, the result is a
// UNC path if and only if the first path element is a UNC path.
func (fs *S3FS) Join(elem ...string) string {
	return ""
}