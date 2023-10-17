package txn

/******************************************
 * Apps API Methods
 * Register client applications that can be used to obtain OAuth tokens
 * https://docs.joinmastodon.org/methods/apps/
 ******************************************/

// https://docs.joinmastodon.org/methods/apps/#create
// POST /api/v1/apps
// Returns: Application
type PostApplication struct {
	ClientName   string `form:"client_name"`
	RedirectURIs string `form:"redirect_uris"`
	Scopes       string `form:"scopes"`
	Website      string `form:"website"`
}

// https://docs.joinmastodon.org/methods/apps/#verify_credentials
// GET /api/v1/apps/verify_credentials
// Returns: Application
type GetApplication_VerifyCredentials struct {
	Authorization string `header:"Authorization"`
}
