package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DownloadFromS3(svc *s3.S3, bucketName, localDir string) error {
	var totalSize int64
	var continuationToken *string

	for {
		result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
			Bucket:            aws.String(bucketName),
			ContinuationToken: continuationToken,
		})
		if err != nil {
			return fmt.Errorf("failed to list objects: %w", err)
		}

		if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create local directory: %w", err)
		}

		for _, item := range result.Contents {
			key := *item.Key
			if strings.HasSuffix(key, "/") {
				continue
			}

			log.Printf("Downloading: %s\n", *item.Key)
			totalSize += *item.Size

			filePath := filepath.Join(localDir, *item.Key)
			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory for file %s: %w", filePath, err)
			}

			outFile, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", filePath, err)
			}
			defer outFile.Close()

			resultObj, err := svc.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucketName),
				Key:    item.Key,
			})
			if err != nil {
				return fmt.Errorf("failed to download file %s: %w", *item.Key, err)
			}
			defer resultObj.Body.Close()

			_, err = io.Copy(outFile, resultObj.Body)
			if err != nil {
				return fmt.Errorf("failed to write file %s: %w", filePath, err)
			}

			log.Printf("Downloaded %s to %s\n", *item.Key, filePath)
		}

		if *result.IsTruncated {
			continuationToken = result.NextContinuationToken
		} else {
			break
		}
	}

	log.Printf("Total size of downloaded files: %d bytes", totalSize)

	return nil
}
