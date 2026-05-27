package auth

import (
	"golang.org/x/oauth2"
)

const (
	keyringService = "optimus"
)

func RetrieveFromKeyring(clientID string) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StoreInKeyring(clientID string, t *oauth2.Token) error { _ = "STUB: not implemented"; return nil }
