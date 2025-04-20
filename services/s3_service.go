package services

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/aws_util"
	"fmt"
)

type S3Service struct{}

var cdnService *CdnService

func PushToS3(filePath string, data string) (string, error) {
	_, err := aws_util.UploadStrDataToS3(filePath, data)
	if err != nil {
		return "", err
	}

	// purging cdn
	fmt.Println(cdnService.GetFullPath(filePath))
	var cdnPath = cdnService.GetFullPath(filePath)
	err = cdnService.Purge(cdnPath)
	if err != nil {
		return "", err
	}

	return cdnPath, err
}
