// dir.go implements the interface billy.Dir

package main

import "os"

// ReadDir reads the directory named by dirname and returns a list of
// directory entries sorted by filename.
func (fs *S3FS) ReadDir(path string) ([]os.FileInfo, error) {
	return nil, nil
}

// MkdirAll creates a directory named path, along with any necessary
// parents, and returns nil, or else returns an error. The permission bits
// perm are used for all directories that MkdirAll creates. If path is/
// already a directory, MkdirAll does nothing and returns nil.
func (fs *S3FS) MkdirAll(filename string, perm os.FileMode) error {
	return nil
}
