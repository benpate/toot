package transaction

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
type GetAccount_VerifyCredentials struct{}

// https://docs.joinmastodon.org/methods/accounts/#update_credentials
// PATCH /api/v1/accounts/update_credentials
type PatchAccount_UpdateCredentials struct {
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
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#statuses
// GET /api/v1/accounts/:id/statuses
type GetAccount_Statuses struct {
	ID             string `param:"id"`
	MaxID          string `query:"max_id"`
	SinceID        string `query:"since_id"`
	MinID          string `query:"min_id"`
	Limit          int    `query:"limit"`
	OnlyMedia      bool   `query:"only_media"`
	ExcludeReplies bool   `query:"exclude_replies"`
	ExcludeReblogs bool   `query:"exclude_reblogs"`
	Pinned         bool   `query:"pinned"`
	Tagged         string `query:"tagged"`
}

// https://docs.joinmastodon.org/methods/accounts/#followers
// GET /api/v1/accounts/:id/followers
type GetAccount_Followers struct {
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/accounts/#following
// GET /api/v1/accounts/:id/following
type GetAccount_Following struct {
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/accounts/#featured_tags
// GET /api/v1/accounts/:id/featured_tags
type GetAccount_FeaturedTags struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#lists
// GET /api/v1/accounts/:id/lists
type GetAccount_Lists struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#follow
// POST /api/v1/accounts/:id/follow
type PostAccount_Follow struct {
	ID        string   `param:"id"`
	Reblogs   bool     `form:"reblogs"`
	Notify    bool     `form:"notify"`
	Languages []string `form:"languages"`
}

// https://docs.joinmastodon.org/methods/accounts/#unfollow
// POST /api/v1/accounts/:id/unfollow
type PostAccount_Unfollow struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
// POST /api/v1/accounts/:id/remove_from_followers
// Returns: Relationship
type PostAccount_RemoveFromFollowers struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#block
// POST /api/v1/accounts/:id/block
// Returns: Relationship
type PostAccount_Block struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#unblock
// POST /api/v1/accounts/:id/unblock
// Returns: Relationship
type PostAccount_Unblock struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#mute
// POST /api/v1/accounts/:id/mute
// Returns: Relationship
type PostAccount_Mute struct {
	ID            string `param:"id"`
	Notifications bool   `form:"notifications"`
	Duration      bool   `form:"duration"`
}

// https://docs.joinmastodon.org/methods/accounts/#unmute
// POST /api/v1/accounts/:id/unmute
// Returns: Relationship
type PostAccount_Unmute struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#pin
// POST /api/v1/accounts/:id/pin
// Returns: Relationship
type PostAccount_Pin struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#unpin
// POST /api/v1/accounts/:id/unpin
// Returns: Relationship
type PostAccount_Unpin struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/accounts/#note
// POST /api/v1/accounts/:id/note
// Returns: Relationship
type PostAccount_Note struct {
	ID      string `param:"id"`
	Comment string `form:"comment"`
}

// https://docs.joinmastodon.org/methods/accounts/#relationships
// GET /api/v1/accounts/relationships
// Returns: Array of Relationships
type GetAccount_Relationships struct {
	IDs []string `query:"id[]"`
}

// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
// GET /api/v1/accounts/:id/familiar_followers
// Returns: Array of FamiliarFollower
type GetAccount_FamiliarFollowers struct {
	ID string `param:"id[]"`
}

// https://docs.joinmastodon.org/methods/accounts/#search
// GET /api/v1/accounts/search
// Returns: Array of Account
type GetAccount_Search struct {
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
	Acct string `query:"acct"`
}
