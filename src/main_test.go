package main

import (
	"os"
	"testing"

	. "github.com/SND1231/dip-test"
)

func TestCheckExecuteSuccess(t *testing.T) {
	os.Setenv("IS_SUCCESS", "1")
	err := Execute("file", "bucket", "key")

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckExecuteError(t *testing.T) {
	os.Setenv("IS_SUCCESS", "0")
	err := Execute("file", "bucket", "key")

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}
