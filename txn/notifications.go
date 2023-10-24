package txn

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// https://docs.joinmastodon.org/methods/notifications/#get
// GET /api/v1/notifications
// Returns: []Notification
type GetNotifications struct {
	Host         string   `header:"Host"`
	MaxID        string   `query:"max_id"`
	SinceID      string   `query:"since_id"`
	MinID        string   `query:"min_id"`
	Limit        int64    `query:"limit"`
	Types        []string `query:"types"`
	ExcludeTypes []string `query:"exclude_types"`
	AccountID    string   `query:"account_id"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetNotifications) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// https://docs.joinmastodon.org/methods/notifications/#get-one
// GET /api/v1/notifications/:id
// Returns: Notification
type GetNotification struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/notifications/#clear
// POST /api/v1/notifications/clear
// Returns: Empty Struct
type PostNotifications_Clear struct {
	Host string `header:"Host"`
}

// https://docs.joinmastodon.org/methods/notifications/#dismiss
// POST /api/v1/notifications/dismiss
// Returns: Empty Struct
type PostNotification_Dismiss struct {
	Host string `header:"Host"`
	ID   string `form:"id"`
}
