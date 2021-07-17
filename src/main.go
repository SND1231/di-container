package main

import (
	"flag"
	"log"

	"github.com/SND1231/dip-test/factory"
)

func main() {
	flag.Parse()
	args := flag.Args()
	filepath := args[0]
	bucketName := args[1]
	objectKey := args[2]
	err := Execute(filepath, bucketName, objectKey)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}

func Execute(filepath string, bucketName string, objectKey string) error {

	ftp := factory.GetFtp()
	err := ftp.Upload(filepath, bucketName, objectKey)

	return err
}
