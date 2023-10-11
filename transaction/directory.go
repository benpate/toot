package transaction

/******************************************
 * Directory API Methods
 * A directory of profiles that your website is aware of.
 * https://docs.joinmastodon.org/methods/directory/
 ******************************************/

// https://docs.joinmastodon.org/methods/directory/#get
// GET /api/v1/directory
// Returns: Array of Account
// List accounts visible in the directory.
type GetDirectory struct {
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
	Order  string `query:"order"`
	Local  bool   `query:"local"`
}
