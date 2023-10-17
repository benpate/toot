package route

/******************************************
 * Suggestions API Methods
 * Server-generated suggestions on who to follow, based on previous
 * positive interactions
 * https://docs.joinmastodon.org/methods/suggestions/
 ******************************************/

// https://docs.joinmastodon.org/methods/suggestions/#v2
const GetSuggestions = "/api/v2/suggestions"

// https://docs.joinmastodon.org/methods/suggestions/#remove
const DeleteSuggestion = "/api/v1/suggestions/:account_id"
