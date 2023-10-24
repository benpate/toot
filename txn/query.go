package txn

// https://docs.joinmastodon.org/api/guidelines/#pagination
type QueryPage struct {
	Limit   int    `query:"limit"`    // The maximum number of results to return. Usually, there is a default limit and a maximum limit; these will vary according to the API method.
	MaxID   string `query:"max_id"`   // All results returned will be lesser than this ID. In effect, sets an upper bound on results.
	MinID   string `query:"min_id"`   // Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward.
	SinceID string `query:"since_id"` // All results returned will be greater than this ID. In effect, sets a lower bound on results.
}
