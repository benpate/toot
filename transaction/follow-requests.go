package transaction

/******************************************
 * Follow Requests API Methods
 * View and manage follow requests
 * https://docs.joinmastodon.org/methods/follow_requests/
 ******************************************/

// https://docs.joinmastodon.org/methods/follow_requests/#get
// GET /api/v1/follow_requests
// Returns: Array of Account
type GetFollowRequests struct {
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/follow_requests/#accept
// POST /api/v1/follow_requests/:account_id/authorize
// Returns: Relationship
type PostFollowRequests_Authorize struct {
	AccountID string `param:"account_id"`
}

// https://docs.joinmastodon.org/methods/follow_requests/#reject
// POST /api/v1/follow_requests/:account_id/reject
// Returns: Relationship
type PostFollowRequests_Reject struct {
	AccountID string `param:"account_id"`
}
