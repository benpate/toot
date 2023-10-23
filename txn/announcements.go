package txn

/******************************************
 * Announcements API Methods
 * For announcements set by administration
 * https://docs.joinmastodon.org/methods/announcements/
 ******************************************/

// https://docs.joinmastodon.org/methods/announcements/#get
// GET /api/v1/announcements
// Returns: Array of Announcement
type GetAnnouncements struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	WithDismissed bool   `query:"with_dismissed"`
}

// https://docs.joinmastodon.org/methods/announcements/#dismiss
// POST /api/v1/announcements/:id/dismiss
// Returns: Empty struct
type PostAnnouncement_Dismiss struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/announcements/#put-reactions
// PUT /api/v1/announcements/:id/reactions/:name
// Returns: Empty struct
type PutAnnouncement_Reaction struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
	Name          string `param:"name"`
}

// https://docs.joinmastodon.org/methods/announcements/#delete-reactions
// DELETE /api/v1/announcements/:id/reactions/:name
// Returns: Empty struct
type DeleteAnnouncement_Reaction struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
	Name          string `param:"name"`
}
