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
	Host string `header:"Host"`

	QueryPage
}
