package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// HmacSha1 encode message with key with hmacsha1
func HmacSha1(key []byte, message []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(message)
	sha := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(sha))
}
