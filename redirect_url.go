package main

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
	"time"
)

// Expiration interval in months
const expirationIntervalMonths = 1

// RedirectUrl defines the properties of a shortened URL
type RedirectUrl struct {
	Hash           string
	OriginalUrl    string
	RedirectCount  int
	CreationDate   time.Time
	ExpirationDate time.Time
}

// NewRedirectUrl creates a new Redirect entity
func NewRedirectUrl(url string) RedirectUrl {
	return RedirectUrl{
		Hash:           generateUrlHash(url),
		OriginalUrl:    url,
		RedirectCount:  0,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().AddDate(0, expirationIntervalMonths, 0),
	}
}

// generateUrlHash generates short URL hash
func generateUrlHash(url string) string {
	hash := md5.New()
	hash.Write([]byte(url))

	hashedUrl := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	hashedUrl = hashedUrl[:8]
	hashedUrl = strings.ReplaceAll(hashedUrl, "+", "-")
	hashedUrl = strings.ReplaceAll(hashedUrl, "/", "_")

	return hashedUrl
}
