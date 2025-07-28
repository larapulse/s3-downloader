package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	bucket := os.Getenv("AWS_BUCKET")
	localDir := "storage"

	// Create the storage directory if it doesn't exist
	if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create storage directory: %v", err)
	}

	// Initialize a session that the SDK will use to load configuration,
	// credentials, and region from the environment variables or shared config file.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	svc := s3.New(sess)

	if err := DownloadFromS3(svc, bucket, localDir); err != nil {
		log.Fatalf("failed to download from S3: %v", err)
	}

	fmt.Println("Download completed successfully.")
}
