package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var BucketName string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	BucketName = os.Getenv("BUCKET_NAME")
}

func main() {
	fmt.Println(BucketName)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	get, err := client.ListObjectsV2(
		context.Background(),
		&s3.ListObjectsV2Input{
			Bucket:    &BucketName,
			Delimiter: aws.String("/"),
		},
	)
	if err != nil {
		panic(err)
	}
	for _, d := range get.CommonPrefixes {
		fmt.Println(aws.ToString(d.Prefix))
	}
	for _, f := range get.Contents {
		fmt.Println(aws.ToString(f.Key))
	}

	_, err = client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: &BucketName,
			Key:    aws.String("hello/"),
		},
	)
	if err != nil {
		panic(err)
	}

	s3fs, err := NewS3FS(client, BucketName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("s3fs.Root() = %q\n", s3fs.Root())

}
