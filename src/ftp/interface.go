package ftp

type Ftp interface {
	Upload(filepath string, bucketName string, objectKey string) error
}
