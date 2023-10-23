package txn

/******************************************
 * OAuth API Methods
 * Generate and manage OAuth tokens
 * https://docs.joinmastodon.org/methods/oauth/
******************************************/

// https://docs.joinmastodon.org/methods/oauth/#authorize
// GET /oauth/authorize
// Returns: Authorization code
type GetOAuth_Authorize struct {
	Host         string `header:"Host"`
	ResponseType string `query:"response_type"`
	ClientID     string `query:"client_id"`
	RedirectURI  string `query:"redirect_uri"`
	Scope        string `query:"scope"`
	ForceLogin   bool   `query:"force_login"`
	Language     string `query:"language"`
}

// https://docs.joinmastodon.org/methods/oauth/#token
// POST /oauth/token
// Returns: Token
// Obtain an access token, to be used during API calls that are not public
type PostOAuth_Token struct {
	Host         string `header:"Host"`
	GrantType    string `form:"grant_type"`
	Code         string `form:"code"`
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	RedirectURI  string `form:"redirect_uri"`
	Scope        string `form:"scope"`
}

// https://docs.joinmastodon.org/methods/oauth/#revoke
// POST /oauth/revoke
// Returns: Empty struct
// Revoke an access token to make it no longer valid for use
type PostOAuth_Revoke struct {
	Host         string `header:"Host"`
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Token        string `form:"token"`
}
