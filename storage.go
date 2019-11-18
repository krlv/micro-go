package main

// Storage represents all possible actions available to deal with data
type Storage interface {
	AddUrl(RedirectUrl) error
	GetUrlByHash(string) error
}
