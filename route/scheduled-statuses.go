package route

/******************************************
 * Scheduled Statuses API Methods
 * Manage statuses that were scheduled to be published at a future date.
 * https://docs.joinmastodon.org/methods/scheduled_statuses/
 ******************************************/

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get
const GetScheduledStatuses = "/api/v1/scheduled_statuses"

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
const GetScheduledStatus = "/api/v1/scheduled_statuses/:id"

// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
const PutScheduledStatus = "/api/v1/scheduled_statuses/:id"

// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
const DeleteScheduledStatus = "/api/v1/scheduled_statuses/:id"
