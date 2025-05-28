package utils

import (
	"bytes"
	"crypto/sha256"
	"specture/internal/config"
	"time"
	"fmt"
)

func SHA256(b []byte) []byte {
	s := sha256.Sum256(b)
	return s[:]
}

func SHA256STR(s string) string {
	str := SHA256([]byte(s))
	return fmt.Sprintf("%x", str)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func CompareBytes(a, b []byte) bool {
	if !bytes.Equal(a, b) {
		return false
	}
	return true
}

func ValidateTimestamp(timestamp int) bool {
	if config.GetExpire() > 0 && Abs(int(time.Now().Unix())-timestamp) > 300 {
		return false
	}
	return true
}
