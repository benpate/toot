package transaction

/******************************************
 * Suggestions API Methods
 * Server-generated suggestions on who to follow, based on previous
 * positive interactions
 * https://docs.joinmastodon.org/methods/suggestions/
 ******************************************/

// https://docs.joinmastodon.org/methods/suggestions/#v2
// GET /api/v2/suggestions
// Returns: []Suggestion
type GetSuggestions struct {
	Limit int `query:"limit"`
}

// https://docs.joinmastodon.org/methods/suggestions/#remove
// DELETE /api/v1/suggestions/:account_id
// Returns Empty Struct
type DeleteSuggestion struct {
	AccountID string `param:"account_id"`
}
