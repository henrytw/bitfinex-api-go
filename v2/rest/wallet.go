package rest

import (
	"github.com/henrytw/bitfinex-api-go/v2"
	"path"
)

// WalletService manages data flow for the Wallet API endpoint
type WalletService struct {
	requestFactory
	Synchronous
}

// All returns all orders for the authenticated account.
func (s *WalletService) Wallet() (*bitfinex.WalletSnapshot, error) {
	req, err := s.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("wallets"))
	if err != nil {
		return nil, err
	}
	raw, err := s.Request(req)
	if err != nil {
		return nil, err
	}

	os, err := bitfinex.NewWalletSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}

	return os, nil
}
