package txn

/******************************************
 * Follow Requests API Methods
 * View and manage follow requests
 * https://docs.joinmastodon.org/methods/follow_requests/
 ******************************************/

// https://docs.joinmastodon.org/methods/follow_requests/#get
// GET /api/v1/follow_requests
// Returns: Array of Account
type GetFollowRequests struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	MinID   string `query:"min_id"`
	SinceID string `query:"since_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetFollowRequests) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// https://docs.joinmastodon.org/methods/follow_requests/#accept
// POST /api/v1/follow_requests/:account_id/authorize
// Returns: Relationship
type PostFollowRequest_Authorize struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
}

// https://docs.joinmastodon.org/methods/follow_requests/#reject
// POST /api/v1/follow_requests/:account_id/reject
// Returns: Relationship
type PostFollowRequest_Reject struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
}
