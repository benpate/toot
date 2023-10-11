// Package toot implements a Mastodon API server, which opens up other applications to Mastodon clients.
package toot

import (
	"github.com/benpate/toot/model"
	"github.com/benpate/toot/transaction"
)

type API struct {
	PostAccounts                    func(transaction.PostAccounts) (model.Token, error)
	GetAccounts_VerifyCredentials   func(transaction.GetAccounts_VerifyCredentials) (model.Account, error)
	PatchAccounts_UpdateCredentials func(transaction.PatchAccounts_UpdateCredentials) (model.Account, error)
	GetAccount                      func(transaction.GetAccount) (model.Account, error)
	GetAccount_Statuses             func(transaction.GetAccount_Statuses) ([]model.Status, error)
	GetAccount_Followers            func(transaction.GetAccount_Followers) ([]model.Account, error)
	GetAccount_Following            func(transaction.GetAccount_Following) ([]model.Account, error)
}

func New() API {
	return API{}
}
