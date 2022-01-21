// Package hash contains a hashing function for generating deterministic shortcodes.
package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// GetHash generates 128 bit hash of given text.
func getHash(text string) [16]byte {
	return md5.Sum([]byte(text))
}

// GenerateShortCode returns first 8 characters of above hash.
func GenerateShortCode(text string) string {
	hash := getHash(text)
	return hex.EncodeToString(hash[:])[:8]
}
