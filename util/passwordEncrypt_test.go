package util

import "testing"

func TestPasswordEncrypt(t *testing.T) {
	h, _ := PasswordEncrypt("123")
	t.Log(h)
	t.Log(len(h))
}
