package transaction

/******************************************
 * Tags API Methods
 * View information about or follow/unfollow hashtags
 * https://docs.joinmastodon.org/methods/tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/tags/#get
// GET /api/v1/tags/:id
// Returns: Tag
type GetTag struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/tags/#follow
// POST /api/v1/tags/:id/follow
// Returns: Tag
type PostTag_Follow struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/tags/#unfollow
// POST /api/v1/tags/:id/unfollow
// Returns: Tag
type PostTag_Unfollow struct {
	ID string `param:"id"`
}
