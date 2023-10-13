package object

// Marker represents the last read position within a user's timelines.
// https://docs.joinmastodon.org/entities/Marker/
type Marker struct {
	LastReadID string `json:"last_read_id"` // The ID of the most recently viewed entity.
	Version    int    `json:"version"`      // An incrementing counter, used for locking to prevent write conflicts.
	UpdatedAt  string `json:"updated_at"`   // The timestamp of when the marker was set.
}
