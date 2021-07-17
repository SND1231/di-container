package ftp

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

type Test struct{}

func (test Test) Upload(filepath string, bucketName string, objectKey string) (err error) {
	success := os.Getenv("IS_SUCCESS")
	if success == "1" {
		return nil
	}

	return awserr.New("ReadRequestBody", "unable to initialize upload", err)
}
