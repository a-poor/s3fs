package main

import (
	"context"
	"fmt"
	"log"
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

type S3FS struct {
	bucketName string
	prefix     string
}

func NewS3FS(bucketName, prefix string) *S3FS {
	return &S3FS{
		bucketName: bucketName,
		prefix:     prefix,
	}
}

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(BucketName)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	// manager := manager.NewDownloader(client)

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: &BucketName,
		// Prefix: aws.String(""),
		Delimiter: aws.String("/"),
	})

	for p := 0; paginator.HasMorePages(); p++ {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalln("error:", err)
		}
		// fmt.Printf("Page: %+v\n", page)

		fmt.Println("Listing prefixes...")
		for _, pre := range page.CommonPrefixes {
			fmt.Println(aws.ToString(pre.Prefix))
		}

		fmt.Println("Listing objects...")
		for i, obj := range page.Contents {
			fmt.Printf("[%2d] Page: %d, Bucket: %q, Key: %q\n", i, p, BucketName, aws.ToString(obj.Key))
			// fmt.Printf("%+v\n", obj)
		}
	}

}
