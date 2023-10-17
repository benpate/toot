package route

/******************************************
 * Instance API Methods
 * Discover information about a Mastodon website
 * https://docs.joinmastodon.org/methods/instance/
 ******************************************/

// https://docs.joinmastodon.org/methods/instance/#v2
const GetInstance = "/api/v2/instance"

// https://docs.joinmastodon.org/methods/instance/#peers
const GetInstance_Peers = "/api/v1/instance/peers"

// https://docs.joinmastodon.org/methods/instance/#activity
const GetInstance_Activity = "/api/v1/instance/activity"

// https://docs.joinmastodon.org/methods/instance/#rules
const GetInstance_Rules = "/api/v1/instance/rules"

// https://docs.joinmastodon.org/methods/instance/#domain_blocks
const GetInstance_DomainBlocks = "/api/v1/instance/domain_blocks"

// https://docs.joinmastodon.org/methods/instance/#extended_description
const GetInstance_ExtendedDescription = "/api/v1/instance/extended_description"
