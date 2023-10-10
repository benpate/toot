package model

// ScheduledStatus represents a status that will be published at a future scheduled date.
// https://docs.joinmastodon.org/entities/ScheduledStatus/
type ScheduledStatus struct {
	ID               string            `json:"id"`
	ScheduledAt      string            `json:"scheduled_at"`
	Params           map[string]any    `json:"params"`
	MediaAttachments []MediaAttachment `json:"media_attachments"`
}
