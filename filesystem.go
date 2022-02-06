package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-git/go-billy/v5"
)

type S3FS struct {
	client     *s3.Client
	bucketName string
	root       string
}

func NewS3FS(client *s3.Client, bucketName string) (*S3FS, error) {
	// Check for a non-nil client
	if client == nil {
		return nil, fmt.Errorf("s3 client cannot be nil")
	}
	return &S3FS{
		client:     client,
		bucketName: bucketName,
		root:       "",
	}, nil
}

// Capabilities returns the filesystem capabilities.
func (fs *S3FS) Capabilities() billy.Capability {
	return billy.ReadCapability | billy.WriteCapability
}
