package transaction

/******************************************
 * Lists API Methods
 * View and manage lists. See also /api/v1/timelines/list/:id for loading a list timeline
 * https://docs.joinmastodon.org/methods/lists/
 ******************************************/

// https://docs.joinmastodon.org/methods/lists/#get
// GET /api/v1/lists
// Returns: Array of List
type GetLists struct{}

// https://docs.joinmastodon.org/methods/lists/#get-one
// GET /api/v1/lists/:id
// Returns: List
type GetList struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/lists/#create
// POST /api/v1/lists
// Returns: List
type PostList struct {
	Title         string `form:"title"`
	RepliesPolicy string `form:"replies_policy"`
	Exclusive     bool   `form:"exclusive"`
}

// https://docs.joinmastodon.org/methods/lists/#update
// PUT /api/v1/lists/:id
// Returns: List
type PutList struct {
	ID            string `param:"id"`
	Title         string `form:"title"`
	RepliesPolicy string `form:"replies_policy"`
}

// https://docs.joinmastodon.org/methods/lists/#delete
// DELETE /api/v1/lists/:id
type DeleteList struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/lists/#accounts
// GET /api/v1/lists/:id/accounts
// Returns: Array of Account
type GetList_Accounts struct {
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/lists/#accounts-add
// POST /api/v1/lists/:id/accounts
// Returns: Empty Struct
// Add accounts to the given list. Note that the user must be following these accounts.
type PostList_Accounts struct {
	ID         string   `param:"id"`
	AccountIDs []string `form:"account_ids"`
}

// https://docs.joinmastodon.org/methods/lists/#accounts-remove
// DELETE /api/v1/lists/:id/accounts
// Returns: Empty Struct
// Remove accounts from the given list.
type DeleteList_Accounts struct {
	ID         string   `param:"id"`
	AccountIDs []string `form:"account_ids"`
}
