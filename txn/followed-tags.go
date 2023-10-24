package txn

/******************************************
 * Followed Tags API Methods
 * View your followed hashtags
 * https://docs.joinmastodon.org/methods/followed_tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/followed_tags/#get
// GET /api/v1/followed_tags
// Returns: Array of Tag
type GetFollowedTags struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetFollowedTags) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}
