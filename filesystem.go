package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-git/go-billy/v5"
)

const (
	DefaultSeparator = "/"
)

type S3FS struct {
	client    *s3.Client
	bucket    string
	root      string
	separator string
}

// NewS3FS creates a new S3FS Filesystem.
func NewS3FS(client *s3.Client, bucket string) (billy.Filesystem, error) {
	// Check for a non-nil client
	if client == nil {
		return nil, fmt.Errorf("s3 client cannot be nil")
	}
	return &S3FS{
		client:    client,
		bucket:    bucket,
		root:      "",
		separator: DefaultSeparator,
	}, nil
}

// NewS3FSWithSeparator creates a new S3FS with the given separator.
func NewS3FSWithCustomSeparator(client *s3.Client, bucket, separator string) (billy.Filesystem, error) {
	// Check for a non-nil client
	if client == nil {
		return nil, fmt.Errorf("s3 client cannot be nil")
	}
	return &S3FS{
		client:    client,
		bucket:    bucket,
		root:      "",
		separator: separator,
	}, nil
}

// Capabilities returns the filesystem capabilities.
func (fs *S3FS) Capabilities() billy.Capability {
	return billy.ReadCapability | billy.WriteCapability
}
