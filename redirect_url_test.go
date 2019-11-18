package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRedirectUrl(t *testing.T) {
	url := NewRedirectUrl("https://example.com")
	assert.NotNil(t, url.Hash)
	assert.Equal(t, "https://example.com", url.OriginalUrl)
	assert.Equal(t, 0, url.RedirectCount)
	assert.True(t, url.CreationDate.Before(time.Now()))
	assert.True(t, url.ExpirationDate.After(time.Now()))
}

func Test_generateUrlHash(t *testing.T) {
	hash := generateUrlHash("https://example.com")
	assert.Equal(t, 8, len(hash))
}
