package geetest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(src, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}
