package services

import "R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/aws_util"

type S3Service struct {
}

func PushToS3(filePath string, data string) (string, error) {
	return aws_util.UploadStrDataToS3(filePath, data)
}
