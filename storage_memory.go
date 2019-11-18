package main

import (
	"fmt"
)

// StorageMemory in-memory data storage
type StorageMemory struct {
	urls map[string]RedirectUrl
}

// AddUrl inserts new redirect URL record
func (s *StorageMemory) AddUrl(url RedirectUrl) error {
	_, ok := s.urls[url.Hash]
	if ok {
		return fmt.Errorf("redirect URL with %s hash already ok", url.Hash)
	}

	s.urls[url.Hash] = url

	return nil
}

// GetUrlByHash returns redirect URL record by hash
func (s *StorageMemory) GetUrlByHash(hash string) (RedirectUrl, error) {
	url, ok := s.urls[hash]
	if !ok {
		return RedirectUrl{}, fmt.Errorf("redirect URL with %s hash not found", url.Hash)
	}

	return url, nil
}
