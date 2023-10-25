package scope

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// https://docs.joinmastodon.org/methods/notifications/#get
const GetNotifications = ReadNotifications

// https://docs.joinmastodon.org/methods/notifications/#get-one
const GetNotification = ReadNotifications

// https://docs.joinmastodon.org/methods/notifications/#clear
const PostNotifications_Clear = WriteNotifications

// https://docs.joinmastodon.org/methods/notifications/#dismiss
const PostNotification_Dismiss = WriteNotifications
