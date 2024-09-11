package hasher

import (
	"testing"
)

func TestHasher_Hash(t *testing.T) {
	tests := []struct {
		name     string
		mode     string
		input    string
		expected string
	}{
		{
			name:     "SHA256 hashing",
			mode:     TypeSha256,
			input:    "hello",
			expected: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			name:     "MD5 hashing",
			mode:     TypeMD5,
			input:    "world",
			expected: "7d793037a0760186574b0282f2f435e7",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hasher := New(test.mode)
			result := hasher.Hash(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
