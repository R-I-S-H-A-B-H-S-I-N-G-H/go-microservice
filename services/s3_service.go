package services

import "rishabhsingh2305/spendings-app/utils/aws_util"

type S3Service struct {
}

func PushToS3(filePath string, data string) (string, error) {
	return aws_util.UploadStrDataToS3(filePath, data)
}
