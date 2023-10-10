package model

// FamiliarFollowers represents a subset of your follows who also follow some other user.
// https://docs.joinmastodon.org/entities/FamiliarFollowers/
type FamiliarFollowers []FamiliarFollower

// https://docs.joinmastodon.org/entities/FamiliarFollowers/#attributes
type FamiliarFollower struct {
	ID       string    `json:"id"`       // The ID of the Account in the database.
	Accounts []Account `json:"accounts"` // Accounts you follow that also follow this account.
}
