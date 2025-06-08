package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func generateSignature(secret string, payload string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
