package model

// FilterStatus represents a status ID that, if matched, should cause the filter action to be taken.
// https://docs.joinmastodon.org/entities/FilterStatus/
type FilterStatus struct {
	ID       string `json:"id"`        // The ID of the FilterStatus in the database.
	StatusID string `json:"status_id"` // The ID of the Status that will be filtered.
}
