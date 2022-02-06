// chroot.go implements the interface billy.Chroot

package main

import "github.com/go-git/go-billy/v5"

// Chroot returns a new filesystem from the same type where the new root is
// the given path. Files outside of the designated directory tree cannot be
// accessed.
func (fs *S3FS) Chroot(path string) (billy.Filesystem, error) {
	return nil, nil
}

// Root returns the root path of the filesystem.
func (fs *S3FS) Root() string {
	return fs.root
}
