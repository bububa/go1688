package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func HmacSha1(key []byte, message string) string {
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	sha := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(sha))
}
