package route

/******************************************
 * Domain Blocks API Methods
 * Manage a User's blocked domains.
 * https://docs.joinmastodon.org/methods/domain_blocks/
 ******************************************/

// https://docs.joinmastodon.org/methods/domain_blocks/#get
const GetDomainBlocks = "/api/v1/domain_blocks"

// https://docs.joinmastodon.org/methods/domain_blocks/#block
const PostDomainBlock = "/api/v1/domain_blocks"

// https://docs.joinmastodon.org/methods/domain_blocks/#unblock
const DeleteDomainBlock = "/api/v1/domain_blocks"
