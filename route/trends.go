package route

/******************************************
 * Trends API Methods
 * View hashtags that are currently being used more frequently than usual.
 * https://docs.joinmastodon.org/methods/trends/
 ******************************************/

// https://docs.joinmastodon.org/methods/trends/#tags
const GetTrends = "/api/v1/trends"

// https://docs.joinmastodon.org/methods/trends/#statuses
const GetTrends_Statuses = "/api/v1/trends/statuses"

// https://docs.joinmastodon.org/methods/trends/#links
const GetTrends_Links = "/api/v1/trends/links"
