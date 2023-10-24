package txn

/******************************************
 * Lists API Methods
 * View and manage lists. See also /api/v1/timelines/list/:id for loading a list timeline
 * https://docs.joinmastodon.org/methods/lists/
 ******************************************/

// https://docs.joinmastodon.org/methods/lists/#get
// GET /api/v1/lists
// Returns: Array of List
type GetLists struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
}

// https://docs.joinmastodon.org/methods/lists/#get-one
// GET /api/v1/lists/:id
// Returns: List
type GetList struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/lists/#create
// POST /api/v1/lists
// Returns: List
type PostList struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	Title         string `form:"title"`
	RepliesPolicy string `form:"replies_policy"`
	Exclusive     bool   `form:"exclusive"`
}

// https://docs.joinmastodon.org/methods/lists/#update
// PUT /api/v1/lists/:id
// Returns: List
type PutList struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
	Title         string `form:"title"`
	RepliesPolicy string `form:"replies_policy"`
}

// https://docs.joinmastodon.org/methods/lists/#delete
// DELETE /api/v1/lists/:id
type DeleteList struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/lists/#accounts
// GET /api/v1/lists/:id/accounts
// Returns: Array of Account
type GetList_Accounts struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
	MaxID         string `query:"max_id"`
	SinceID       string `query:"since_id"`
	MinID         string `query:"min_id"`
	Limit         int    `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetList_Accounts) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// https://docs.joinmastodon.org/methods/lists/#accounts-add
// POST /api/v1/lists/:id/accounts
// Returns: Empty Struct
// Add accounts to the given list. Note that the user must be following these accounts.
type PostList_Accounts struct {
	Host          string   `header:"Host"`
	Authorization string   `header:"Authorization"`
	ID            string   `param:"id"`
	AccountIDs    []string `form:"account_ids"`
}

// https://docs.joinmastodon.org/methods/lists/#accounts-remove
// DELETE /api/v1/lists/:id/accounts
// Returns: Empty Struct
// Remove accounts from the given list.
type DeleteList_Accounts struct {
	Host          string   `header:"Host"`
	Authorization string   `header:"Authorization"`
	ID            string   `param:"id"`
	AccountIDs    []string `form:"account_ids"`
}
