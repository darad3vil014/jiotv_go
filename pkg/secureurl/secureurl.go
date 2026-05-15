package secureurl

import (
	"net/url"
)

// EncryptURL returns the URL query-escaped (no encryption).
func EncryptURL(inputURL string) (string, error) {
	return url.QueryEscape(inputURL), nil
}

// DecryptURL returns the URL query-unescaped (no decryption).
func DecryptURL(encryptedURL string) (string, error) {
	return url.QueryUnescape(encryptedURL)
}

// Init is a no-op; encryption has been removed.
func Init() {}
