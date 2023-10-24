package txn

/******************************************
 * Endorsements API Methods
 * Feature other profiles on your own profile. See also accounts/:id/{pin,unpin}
 * https://docs.joinmastodon.org/methods/endorsements/
 ******************************************/

// https://docs.joinmastodon.org/methods/endorsements/#get
// GET /api/v1/endorsements
// Returns: Array of Account
type GetEndorsements struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	MinID         string `query:"min_id"`
	SinceID       string `query:"since_id"`
	Limit         int    `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetEndorsements) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		MinID:   t.MinID,
		SinceID: t.SinceID,
		Limit:   t.Limit,
	}
}
