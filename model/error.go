package model

// Error represents an error message.
// https://docs.joinmastodon.org/entities/Error/
type Error struct {
	Error            int    `json:"error"`             // The error message.
	ErrorDescription string `json:"error_description"` // A longer description of the error, mainly provided with the OAuth API.
}
