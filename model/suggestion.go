package model

// https://docs.joinmastodon.org/entities/Suggestion/
type Suggestion struct {
	Source  string  `json:"source"`  // The reason this account is being suggested. [staff | past_interactions | global]
	Account Account `json:"account"` // The account being recommended to follow.
}
