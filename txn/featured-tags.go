package txn

/******************************************
 * Featured Tags API Methods
 * Feature tags that you use frequently on your profile
 * https://docs.joinmastodon.org/methods/featured_tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/featured_tags/#get
// GET /api/v1/featured_tags
// Returns: Array of FeaturedTag
type GetFeaturedTags struct {
	Host string `header:"Host"`
}

// https://docs.joinmastodon.org/methods/featured_tags/#feature
// POST /api/v1/featured_tags
// Returns: FeaturedTag
type PostFeaturedTag struct {
	Host string `header:"Host"`
	Name string `form:"name"`
}

// https://docs.joinmastodon.org/methods/featured_tags/#unfeature
// DELETE /api/v1/featured_tags/:id
// Returns: Empty object
type DeleteFeaturedTag struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/featured_tags/#suggestions
// GET /api/v1/featured_tags/suggestions
// Returns: Array of FeaturedTag
type GetFeaturedTags_Suggestions struct {
	Host string `header:"Host"`
}
