package txn

/******************************************
 * Mutes API Methods
 * View your mutes. See also /accounts/:id/{mute,unmute}
 * https://docs.joinmastodon.org/methods/mutes/
 ******************************************/

// https://docs.joinmastodon.org/methods/mutes/#get
// GET /api/v1/mutes
// Returns: Array of Account
type GetMutes struct {
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	SinceID       string `query:"since_id"`
	Limit         int    `query:"limit"`
}
