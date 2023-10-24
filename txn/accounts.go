package txn

/******************************************
 * Accounts API Methods
 * Methods concerning accounts and profiles
 * https://docs.joinmastodon.org/methods/accounts/
 ******************************************/

// https://docs.joinmastodon.org/methods/accounts/#create
// POST /api/v1/accounts
// Returns: Token
// Register an account
type PostAccount struct {
	Host      string `header:"Host"`
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	Agreement bool   `form:"agreement"`
	Locale    string `form:"locale"`
	Reason    string `form:"reason"`
}

// https://docs.joinmastodon.org/methods/accounts/#verify_credentials
// GET /api/v1/accounts/verify_credentials
// Returns: CredentialAccount
// Test to make sure that the user token works.
type GetAccount_VerifyCredentials struct {
	Host string `header:"Host"`
}

// https://docs.joinmastodon.org/methods/accounts/#update_credentials
// PATCH /api/v1/accounts/update_credentials
type PatchAccount_UpdateCredentials struct {
	Host         string `header:"Host"`
	DisplayName  string `form:"display_name"`
	Note         string `form:"note"`
	Avatar       string `form:"avatar"`
	Header       string `form:"header"`
	Locked       bool   `form:"locked"`
	Bot          bool   `form:"bot"`
	Discoverable bool   `form:"discoverable"`
}

// https://docs.joinmastodon.org/methods/accounts/#get
// GET /api/v1/accounts/:id
type GetAccount struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#statuses
// GET /api/v1/accounts/:id/statuses
type GetAccount_Statuses struct {
	Host           string `header:"Host"`
	Authorization  string `header:"Authorization"`
	ID             string `param:"id"`
	OnlyMedia      bool   `query:"only_media"`
	ExcludeReplies bool   `query:"exclude_replies"`
	ExcludeReblogs bool   `query:"exclude_reblogs"`
	Pinned         bool   `query:"pinned"`
	Tagged         string `query:"tagged"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/accounts/#followers
// GET /api/v1/accounts/:id/followers
type GetAccount_Followers struct {
	Host string `header:"Host"`
	ID   string `param:"id"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/accounts/#following
// GET /api/v1/accounts/:id/following
type GetAccount_Following struct {
	Host string `header:"Host"`
	ID   string `param:"id"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/accounts/#featured_tags
// GET /api/v1/accounts/:id/featured_tags
type GetAccount_FeaturedTags struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#lists
// GET /api/v1/accounts/:id/lists
type GetAccount_Lists struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#follow
// POST /api/v1/accounts/:id/follow
type PostAccount_Follow struct {
	Host      string   `header:"Host"`
	ID        string   `param:"id"`
	Reblogs   bool     `form:"reblogs"`
	Notify    bool     `form:"notify"`
	Languages []string `form:"languages"`
}

// https://docs.joinmastodon.org/methods/accounts/#unfollow
// POST /api/v1/accounts/:id/unfollow
type PostAccount_Unfollow struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
// POST /api/v1/accounts/:id/remove_from_followers
// Returns: Relationship
type PostAccount_RemoveFromFollowers struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#block
// POST /api/v1/accounts/:id/block
// Returns: Relationship
type PostAccount_Block struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#unblock
// POST /api/v1/accounts/:id/unblock
// Returns: Relationship
type PostAccount_Unblock struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#mute
// POST /api/v1/accounts/:id/mute
// Returns: Relationship
type PostAccount_Mute struct {
	Host          string `header:"Host"`
	ID            string `param:"id"`
	Notifications bool   `form:"notifications"`
	Duration      bool   `form:"duration"`
}

// https://docs.joinmastodon.org/methods/accounts/#unmute
// POST /api/v1/accounts/:id/unmute
// Returns: Relationship
type PostAccount_Unmute struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#pin
// POST /api/v1/accounts/:id/pin
// Returns: Relationship
type PostAccount_Pin struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#unpin
// POST /api/v1/accounts/:id/unpin
// Returns: Relationship
type PostAccount_Unpin struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#note
// POST /api/v1/accounts/:id/note
// Returns: Relationship
type PostAccount_Note struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	Comment string `form:"comment"`
}

// https://docs.joinmastodon.org/methods/accounts/#relationships
// GET /api/v1/accounts/relationships
// Returns: Array of Relationships
type GetAccount_Relationships struct {
	Host string   `header:"Host"`
	IDs  []string `query:"id[]"`
}

// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
// GET /api/v1/accounts/:id/familiar_followers
// Returns: Array of FamiliarFollower
type GetAccount_FamiliarFollowers struct {
	Host string `header:"Host"`
	ID   string `param:"id[]"`
}

// https://docs.joinmastodon.org/methods/accounts/#search
// GET /api/v1/accounts/search
// Returns: Array of Account
type GetAccount_Search struct {
	Host      string `header:"Host"`
	Q         string `query:"q"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
	Resolve   bool   `query:"resolve"`
	Following bool   `query:"following"`
}

// https://docs.joinmastodon.org/methods/accounts/#lookup
// GET /api/v1/accounts/lookup
// Returns: Account
type GetAccount_Lookup struct {
	Host string `header:"Host"`
	Acct string `query:"acct"`
}
