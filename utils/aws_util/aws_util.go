package aws_util

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3Client() *s3.S3 {
	// Hardcoded AWS credentials and region
	awsAccessKey := "00502605ba7f2890000000003"
	awsSecretKey := "K005L1tA6w5oQ0EbtdUeppUH4JFU/I4"
	awsRegion := "us-east-005"
	endpoint := "https://s3.us-east-005.backblazeb2.com"

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(awsRegion),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)
	return s3.New(newSession)
}

func getS3Bucket() string {
	return "testb23"
}

func UploadStrDataToS3(objectKey string, str string) error {
	data := []byte(str)
	return UploadDataToS3(objectKey, data)
}

func UploadDataToS3(objectKey string, data []byte) error {

	svc := getS3Client()

	// Upload the data
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(getS3Bucket()),
		Key:                aws.String(objectKey),
		Body:               bytes.NewReader(data),
		ContentType:        aws.String(getContentType(objectKey)), // Change as necessary based on your file type
		ContentDisposition: aws.String("inline"),                  // Set Content-Disposition to inline
	})

	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", objectKey, getS3Bucket())
	return nil
}

// getContentType returns the MIME type based on the file extension.
func getContentType(filename string) string {
	// Map of file extensions to their corresponding MIME types
	contentTypes := map[string]string{
		".txt":  "text/plain",
		".html": "text/html",
		".json": "application/json",
		".xml":  "application/xml",
		".pdf":  "application/pdf",
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".csv":  "text/csv",
		".zip":  "application/zip",
		".mp4":  "video/mp4",
		".mp3":  "audio/mpeg",
		// Add more types as needed
	}

	// Get the file extension
	ext := strings.ToLower(filename[strings.LastIndex(filename, "."):])

	// Return the corresponding content type, or a default if not found
	if contentType, exists := contentTypes[ext]; exists {
		return contentType
	}
	return "application/octet-stream" // Default content type for unknown extensions
}
