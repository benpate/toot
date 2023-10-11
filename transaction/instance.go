package transaction

/******************************************
 * Instance API Methods
 * Discover information about a Mastodon website
 * https://docs.joinmastodon.org/methods/instance/
 ******************************************/

// https://docs.joinmastodon.org/methods/instance/#v2
// GET /api/v2/instance
// Returns: Instance
type GetInstance struct{}

// https://docs.joinmastodon.org/methods/instance/#peers
// GET /api/v1/instance/peers
// Returns: []string
// Domains that this instance is aware of.
type GetInstance_Peers struct{}

// https://docs.joinmastodon.org/methods/instance/#activity
// GET /api/v1/instance/activity
// Returns: []map[string]any
// Instance activity over the last 3 months, binned weekly.
type GetInstance_Activity struct{}

// https://docs.joinmastodon.org/methods/instance/#rules
// GET /api/v1/instance/rules
// Returns: []Rule
// Rules that the users of this service should follow.
type GetInstance_Rules struct{}

// https://docs.joinmastodon.org/methods/instance/#domain_blocks
// GET /api/v1/instance/domain_blocks
// Returns: []DomainBlock
// Obtain a list of domains that have been blocked.
type GetInstance_DomainBlocks struct{}

// https://docs.joinmastodon.org/methods/instance/#extended_description
// GET /api/v1/instance/extended_description
// Returns: ExtendedDescription
// Obtain an extended description of this server
type GetInstance_ExtendedDescription struct{}
