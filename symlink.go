// symlink.go implements the interface billy.Symlink

package main

import "os"

// Lstat returns a FileInfo describing the named file. If the file is a
// symbolic link, the returned FileInfo describes the symbolic link. Lstat
// makes no attempt to follow the link.
func (fi s3FileInfo) Lstat(filename string) (os.FileInfo, error) {
	return nil, nil
}

// Symlink creates a symbolic-link from link to target. target may be an
// absolute or relative path, and need not refer to an existing node.
// Parent directories of link are created as necessary.
func (fi s3FileInfo) Symlink(target, link string) error {
	return nil
}

// Readlink returns the target path of link.
func (fi s3FileInfo) Readlink(link string) (string, error) {
	return "", nil
}
