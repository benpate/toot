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
	Host string `header:"Host"`

	QueryPage
}
