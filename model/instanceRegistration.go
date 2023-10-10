package model

// Information about registering for this website.
// https://docs.joinmastodon.org/entities/Instance/#registrations
type InstanceRegistrations struct {
	Enabled          bool   `json:"enabled"`           // Whether registrations are enabled.
	ApprovalRequired bool   `json:"approval_required"` // Whether registrations require moderator approval.
	Message          string `json:"message,omitempty"` // A custom message to be shown when registrations are closed.
}
