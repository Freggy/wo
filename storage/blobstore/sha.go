package blobstore

import (
	"crypto/sha256"
	"encoding/hex"
)

func sha256digest(data []byte) string {
	sha := sha256.New()
	digest := sha.Sum(data)
	return hex.EncodeToString(digest)
}
