package txn

/******************************************
 * Search API Methods
 * Search for content in accounts, statuses and hashtags
 * https://docs.joinmastodon.org/methods/search/
 ******************************************/

// https://docs.joinmastodon.org/methods/search/#v2
// GET /api/v2/search
// Returns: Search
type GetSearch struct {
	Authorization     string `header:"Authorization"`
	Q                 string `query:"q"`
	Type              string `query:"type"` // [accounts | hashtags | statuses]
	Resolve           bool   `query:"resolve"`
	Following         bool   `query:"following"`
	AccountID         string `query:"account_id"`
	ExcludeUnreviewed bool   `query:"exclude_unreviewed"`
	MaxID             string `query:"max_id"`
	MinID             string `query:"min_id"`
	Limit             int    `query:"limit"`
	Offset            int    `query:"offset"`
}
