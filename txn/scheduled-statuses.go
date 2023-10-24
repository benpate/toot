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
	Host string `header:"Host"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
// GET /api/v1/scheduled_statuses/:id
// Returns: ScheduledStatus
type GetScheduledStatus struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
// PUT /api/v1/scheduled_statuses/:id
// Returns: ScheduledStatus
type PutScheduledStatus struct {
	Host        string `header:"Host"`
	ID          string `param:"id"`
	ScheduledAt string `form:"scheduled_at"` // ISO 8601 Datetime
}

// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
// DELETE /api/v1/scheduled_statuses/:id
// Returns: Empty struct
type DeleteScheduledStatus struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}
