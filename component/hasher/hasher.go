package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

const (
	TypeSha256 = "sha256"
	TypeMD5    = "md5"
)

type Hasher struct {
	mode string
}

func New(mode string) *Hasher {
	return &Hasher{
		mode: mode,
	}
}

// accept "md5" and "sha256"
func (h *Hasher) Hash(input string) string {
	if h.mode == TypeSha256 {
		return sha256Hash(input)
	} else {
		return md5Hash(input)
	}
}

func sha256Hash(input string) string {
	// Create a SHA-256 hash object
	hasher := sha256.New()
	// Write the input string to the hash object
	hasher.Write([]byte(input))
	// Get the final hash sum
	hashSum := hasher.Sum(nil)
	// Convert the hash sum to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)

	return hashString
}

func md5Hash(input string) string {
	// Create an MD5 hash object
	hasher := md5.New()
	// Write the input string to the hash object
	hasher.Write([]byte(input))
	// Get the final hash sum
	hashSum := hasher.Sum(nil)
	// Convert the hash sum to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)

	return hashString
}
