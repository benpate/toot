package txn

/******************************************
* Preferences API Methods
* Preferred common behaviors to be shared across clients.
* https://docs.joinmastodon.org/methods/preferences/
******************************************/

// https://docs.joinmastodon.org/methods/preferences/#get
// GET /api/v1/preferences
// Returns: Preferences
type GetPreferences struct {
	Authorization string `header:"Authorization"`
}
