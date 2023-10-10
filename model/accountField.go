package model

// https://docs.joinmastodon.org/entities/Account/#Field
type AccountField struct {
	Name       string `json:"name"`                  // The key of a given field’s key-value pair.
	Value      string `json:"value"`                 // The value associated with the name key.
	VerifiedAt string `json:"verified_at,omitempty"` // Timestamp of when the server verified a URL value for a rel=“me” link. NULLABLE String (ISO 8601 Datetime) if value is a verified URL. Otherwise, null.
}
