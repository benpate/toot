package txn

/******************************************
 * Mutes API Methods
 * View your mutes. See also /accounts/:id/{mute,unmute}
 * https://docs.joinmastodon.org/methods/mutes/
 ******************************************/

// https://docs.joinmastodon.org/methods/mutes/#get
// GET /api/v1/mutes
// Returns: Array of Account
type GetMutes struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	MinID         string `query:"min_id"`
	SinceID       string `query:"since_id"`
	Limit         int    `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetMutes) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}
