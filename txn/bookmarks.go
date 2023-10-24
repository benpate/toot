package txn

/******************************************
 * Bookmarks API Methods
 * View your bookmarks. See also statuses/:id/(bookmark,unbookmark)
 * https://docs.joinmastodon.org/methods/bookmarks/
 ******************************************/

// https://docs.joinmastodon.org/methods/bookmarks/#get
// GET /api/v1/bookmarks
// Returns: Array of Status
// Statuses the user has bookmarked.
type GetBookmarks struct {
	Host string `header:"Host"`

	QueryPage
}
