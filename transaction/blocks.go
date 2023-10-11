package transaction

/******************************************
 * Blocks API Methods
 * View your blocks. See also accounts/:id/blocks
 * https://docs.joinmastodon.org/methods/blocks/
 ******************************************/

// https://docs.joinmastodon.org/methods/blocks/#get
// GET /api/v1/blocks
// Returns: Array of Account
type GetBlocks struct {
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}
