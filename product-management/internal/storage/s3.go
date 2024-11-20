package storage

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"product-management/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3
var bucketName string

// InitS3 initializes the AWS S3 client
func InitS3(cfg *configs.Config) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	s3Client = s3.New(sess)
	bucketName = cfg.S3Bucket
}

// UploadToS3 uploads a file to AWS S3 and returns the file's URL
func UploadToS3(file multipart.File, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(buf.Bytes()),
		ContentLength: aws.Int64(int64(buf.Len())),
		ACL:           aws.String(s3.ObjectCannedACLPublicRead),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %v", err)
	}

	// Return the public URL of the uploaded file
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName), nil
}

// CloseS3 is a placeholder for any potential cleanup logic (AWS SDK handles cleanup automatically)
func CloseS3() {
	// No explicit close for AWS SDK S3 client, but you can add any cleanup here if needed.
}
