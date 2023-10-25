package scope

/******************************************
 * Scheduled Statuses API Methods
 * Manage statuses that were scheduled to be published at a future date.
 * https://docs.joinmastodon.org/methods/scheduled_statuses/
 ******************************************/

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get
const GetScheduledStatuses = ReadStatuses

// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
const GetScheduledStatus = ReadStatuses

// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
const PutScheduledStatus = WriteStatuses

// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
const DeleteScheduledStatus = WriteStatuses
