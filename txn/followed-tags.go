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
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	SinceID       string `query:"since_id"`
	MinID         string `query:"min_id"`
	Limit         int    `query:"limit"`
}
