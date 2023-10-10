package model

// Hints related to contacting a representative of the website.
// https://docs.joinmastodon.org/entities/Instance/#contact
type InstanceContact struct {
	Email   string  `json:"email"`   // An email address that can be messaged regarding inquiries or issues.
	Account Account `json:"account"` // An account that can be contacted natively over the network regarding inquiries or issues.
}
