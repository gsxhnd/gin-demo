package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Md5File(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func GetFileExt(fileName string) (string, error) {
	a := strings.Split(fileName, ".")
	if len(a) == 1 {
		return "", errors.New("no ext")
	}
	return a[len(a)-1], nil
}
