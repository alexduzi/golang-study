package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Bucket  string = "goexpert-bucket-alexduzi"
	objectKey string
	s3Client  *s3.Client
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithSharedConfigProfile("profile_name"))

	if err != nil {
		log.Fatalf("failed to load AWS config, %v\n", err)
		panic(err)
	}

	s3Client = s3.NewFromConfig(cfg)
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		log.Fatalf("failed to open tmp folder, %v\n", err)
		panic(err)
	}
	defer dir.Close()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading directory: %s\n", err)
			continue
		}
		uploadS3Bucket(files[0].Name())
	}
}

func uploadS3Bucket(fileName string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)

	f, err := os.Open(completeFileName)
	if err != nil {
		log.Fatalf("failed to open file, %v\n", err)
		return
	}
	defer f.Close()

	log.Printf("uploading object %s to bucket %s\n", completeFileName, s3Bucket)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})

	if err != nil {
		log.Fatalf("failed to put object, %v\n", err)
		return
	}

	log.Printf("successfully uploaded object %s to bucket %s\n", objectKey, s3Bucket)
}
