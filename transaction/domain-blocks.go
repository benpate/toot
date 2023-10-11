package transaction

/******************************************
 * Domain Blocks API Methods
 * Manage a User's blocked domains.
 * https://docs.joinmastodon.org/methods/domain_blocks/
 ******************************************/

// https://docs.joinmastodon.org/methods/domain_blocks/#get
// GET /api/v1/domain_blocks
// Returns: Array of String
type GetDomainBlocks struct {
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/domain_blocks/#block
// POST /api/v1/domain_blocks
// Returns: Empty struct
type PostDomainBlock struct {
	Domain string `form:"domain"`
}

// https://docs.joinmastodon.org/methods/domain_blocks/#unblock
// DELETE /api/v1/domain_blocks
// Returns: Empty struct
type DeleteDomainBlock struct {
	Domain string `form:"domain"`
}
