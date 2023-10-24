package txn

/******************************************
 * Blocks API Methods
 * View your blocks. See also accounts/:id/blocks
 * https://docs.joinmastodon.org/methods/blocks/
 ******************************************/

// https://docs.joinmastodon.org/methods/blocks/#get
// GET /api/v1/blocks
// Returns: Array of Account
type GetBlocks struct {
	Host string `header:"Host"`

	QueryPage
}
