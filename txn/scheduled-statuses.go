package txn

/******************************************
 * Scheduled Statuses API Methods
 * Manage statuses that were scheduled to be published at a future date.
 * https://docs.joinmastodon.org/methods/scheduled_statuses/
 ******************************************/

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get
// GET /api/v1/scheduled_statuses
// Returns: []ScheduledStatus
type GetScheduledStatuses struct {
	Authorization string `header:"Authorization"`
	MaxID         string `query:"max_id"`
	SinceID       string `query:"since_id"`
	MinID         string `query:"min_id"`
	Limit         int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
// GET /api/v1/scheduled_statuses/:id
// Returns: ScheduledStatus
type GetScheduledStatus struct {
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
// PUT /api/v1/scheduled_statuses/:id
// Returns: ScheduledStatus
type PutScheduledStatus struct {
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
	ScheduledAt   string `form:"scheduled_at"` // ISO 8601 Datetime
}

// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
// DELETE /api/v1/scheduled_statuses/:id
// Returns: Empty struct
type DeleteScheduledStatus struct {
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}
