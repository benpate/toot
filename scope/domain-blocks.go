package scope

/******************************************
 * Domain Blocks API Methods
 * Manage a User's blocked domains.
 * https://docs.joinmastodon.org/methods/domain_blocks/
 ******************************************/

// https://docs.joinmastodon.org/methods/domain_blocks/#get
const GetDomainBlocks = ReadBlocks

// https://docs.joinmastodon.org/methods/domain_blocks/#block
const PostDomainBlock = WriteBlocks

// https://docs.joinmastodon.org/methods/domain_blocks/#unblock
const DeleteDomainBlock = WriteBlocks
