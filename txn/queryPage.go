package txn

// QueryPager wraps the QueryPage method, and is implemented by all "query" transactions
type QueryPager interface {

	// QueryPage extracts the query page data from an object
	QueryPage() QueryPage
}

// QueryPage represents the query parameters that are common to many Mastodon API calls.
// Many API methods allow you to paginate for more information, using parameters such as limit, max_id, min_id, and since_id.
// https://docs.joinmastodon.org/api/guidelines/#pagination
type QueryPage struct {
	Limit   int64  `query:"limit"`    // The maximum number of results to return. Usually, there is a default limit and a maximum limit; these will vary according to the API method.
	MaxID   string `query:"max_id"`   // All results returned will be lesser than this ID. In effect, sets an upper bound on results.
	MinID   string `query:"min_id"`   // Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward.
	SinceID string `query:"since_id"` // All results returned will be greater than this ID. In effect, sets a lower bound on results.
}
