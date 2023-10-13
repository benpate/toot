package object

// https://docs.joinmastodon.org/entities/Announcement/#Status
type AnnouncementStatus struct {
	ID  string `json:"id"`  // The ID of an attached Status in the database.
	URL string `json:"url"` // The URL of an attached Status.
}
