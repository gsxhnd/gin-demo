package authUtil

import (
	"crypto/sha256"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"wattime-admin-api/logger"
)

func PasswordEncrypt(password string) (string, error) {
	h := sha256.New()
	salt := viper.GetString("salt")
	logger.HandlerLogger().Debug(salt)
	s := password + salt
	_, err := io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil)), err
}
