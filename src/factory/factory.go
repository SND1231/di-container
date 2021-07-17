package factory

import (
	"os"

	"github.com/SND1231/dip-test/ftp"
)

func GetFtp() ftp.Ftp {
	testMode := os.Getenv("TESTMODE")
	if testMode == "test" {
		return ftp.Test{}
	}
	return ftp.S3{}
}
