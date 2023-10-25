package route

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// https://docs.joinmastodon.org/methods/notifications/#get
const GetNotifications = "/api/v1/notifications"

// https://docs.joinmastodon.org/methods/notifications/#get-one
const GetNotification = "/api/v1/notifications/:id"

// https://docs.joinmastodon.org/methods/notifications/#clear
const PostNotifications_Clear = "/api/v1/notifications/clear"

// https://docs.joinmastodon.org/methods/notifications/#dismiss
const PostNotification_Dismiss = "/api/v1/notifications/dismiss"
