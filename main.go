package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

const file = "testfile.txt"
const content = "This is our content file for 4sysops!"
const bucketName = "aws-s3-bucket-test-001"

func main() {
	// Create local file to upload to newly created bucket
	localFile, err := CreateLocalFile(file, content)
	if err != nil {
		log.Fatal(err)
	}

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithSharedConfigProfile("sgnadmin"))
	if err != nil {
		log.Fatal(err)
	}

	News3Bucket(cfg, bucketName)
	UploadFileToS3(cfg, bucketName, localFile)
}
