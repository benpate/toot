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
	Host              string `header:"Host"`
	Q                 string `query:"q"`
	Type              string `query:"type"` // [accounts | hashtags | statuses]
	Resolve           bool   `query:"resolve"`
	Following         bool   `query:"following"`
	AccountID         string `query:"account_id"`
	ExcludeUnreviewed bool   `query:"exclude_unreviewed"`
	MaxID             string `query:"max_id"`
	MinID             string `query:"min_id"`
	SinceID           string `query:"since_id"`
	Limit             int    `query:"limit"`
	Offset            int    `query:"offset"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetSearch) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}
