package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func News3Bucket(cfg aws.Config, bn string) (*s3.CreateBucketOutput, error) {
	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Create a bucket
	bucket, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bn),
		ACL:    types.BucketCannedACLPrivate,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintEuWest1,
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("Bucket created. Location: %s\n", *bucket.Location)
	return bucket, nil
}

func UploadFileToS3(cfg aws.Config, bn string, localfile *os.File) (*s3.PutObjectOutput, error) {
	// Place an object in a bucket.
	fmt.Println("Upload an object to the bucket")

	f, err := os.Open(localfile.Name())
	if err != nil {
		panic("Couldn't open local file")
	}
	client := s3.NewFromConfig(cfg)

	upload, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bn),
		Key:    aws.String(localfile.Name()),
		Body:   f,
	})

	f.Close()
	if err != nil {
		return nil, err
	}

	return upload, nil
}
