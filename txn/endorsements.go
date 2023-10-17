package txn

/******************************************
 * Endorsements API Methods
 * Feature other profiles on your own profile. See also accounts/:id/{pin,unpin}
 * https://docs.joinmastodon.org/methods/endorsements/
 ******************************************/

// https://docs.joinmastodon.org/methods/endorsements/#get
// GET /api/v1/endorsements
// Returns: Array of Account
type GetEndorsements struct {
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	SinceID       string `query:"since_id"`
	Limit         int    `query:"limit"`
}
