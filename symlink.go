// symlink.go implements the interface billy.Symlink

package main

import (
	"errors"
	"os"
)

var (
	ErrSymLinkNotSupported = errors.New("symlink not supported by s3")
)

// Lstat returns a FileInfo describing the named file. If the file is a
// symbolic link, the returned FileInfo describes the symbolic link. Lstat
// makes no attempt to follow the link.
func (fs *S3FS) Lstat(filename string) (os.FileInfo, error) {
	return nil, nil
}

// Symlink creates a symbolic-link from link to target. target may be an
// absolute or relative path, and need not refer to an existing node.
// Parent directories of link are created as necessary.
func (fs *S3FS) Symlink(target, link string) error {
	return ErrSymLinkNotSupported
}

// Readlink returns the target path of link.
func (fs *S3FS) Readlink(link string) (string, error) {
	return "", nil
}
