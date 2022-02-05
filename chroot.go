// chroot.go implements the interface billy.Chroot

package main

import "github.com/go-git/go-billy/v5"

// Chroot returns a new filesystem from the same type where the new root is
// the given path. Files outside of the designated directory tree cannot be
// accessed.
func (fi s3FileInfo) Chroot(path string) (billy.Filesystem, error) {
	return nil, nil
}

// Root returns the root path of the filesystem.
func (fi s3FileInfo) Root() string {
	return ""
}
