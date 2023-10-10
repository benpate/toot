package model

// https://docs.joinmastodon.org/entities/Token/
type Token struct {
	AccessToken string `json:"access_token"` // An OAuth token to be used for authorization.
	TokenType   string `json:"token_type"`   // The OAuth token type. Mastodon uses Bearer tokens.
	Scope       string `json:"scope"`        // The OAuth scopes granted by this token, space-separated.
	CreatedAt   int64  `json:"created_at"`   // When the token was generated. (Unix Timestamp)
}
