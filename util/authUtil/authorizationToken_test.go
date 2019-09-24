package authUtil

import (
	"fmt"
	"testing"
)

var tokenData string

func TestSign(t *testing.T) {
	if token, err := Sign(TokenInfo{Type: 1, Username: "1"}, "1246"); err == nil {
		fmt.Println(token)
		tokenData = token
	}
}

func TestParse(t *testing.T) {
	if tokenInfo, err := Parse(tokenData, "1246"); err == nil {
		fmt.Println(tokenInfo)
	}
}
