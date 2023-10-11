package transaction

/******************************************
 * Trends API Methods
 * View hashtags that are currently being used more frequently than usual.
 * https://docs.joinmastodon.org/methods/trends/
 ******************************************/

// https://docs.joinmastodon.org/methods/trends/#tags
// GET /api/v1/trends
// Returns: []Tag
type GetTrends struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

// https://docs.joinmastodon.org/methods/trends/#statuses
// GET /api/v1/trends/statuses
// Returns: []Status
type GetTrends_Statuses struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

// https://docs.joinmastodon.org/methods/trends/#links
// GET /api/v1/trends/links
// Returns: []Link
type GetTrends_Links struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}
