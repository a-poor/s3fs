package main

import (
	"os"
	"time"
)

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
