package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Bucket string = "goexpert-bucket-alexduzi"
	s3Client *s3.Client
	wg       sync.WaitGroup
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

	uploadControl := make(chan struct{}, 100) // channel for upload control with 100 buffer
	errorFileUpload := make(chan string, 10)

	go func() {
		for fileName := range errorFileUpload {
			uploadControl <- struct{}{}
			wg.Add(1)
			go uploadS3Bucket(fileName, uploadControl, errorFileUpload)
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{} // add struct to control upload flow
		go uploadS3Bucket(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Done()
}

func uploadS3Bucket(fileName string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)

	f, err := os.Open(completeFileName)
	if err != nil {
		log.Fatalf("failed to open file, %v\n", err)
		<-uploadControl // removing one struct in channel for next upload, empty channel
		errorFileUpload <- completeFileName
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
		<-uploadControl // removing one struct in channel for next upload
		errorFileUpload <- completeFileName
		return
	}

	log.Printf("successfully uploaded object %s to bucket %s\n", fileName, s3Bucket)
	<-uploadControl // removing one struct in channel for next upload
}
