package route

/******************************************
 * Announcements API Methods
 * For announcements set by administration
 * https://docs.joinmastodon.org/methods/announcements/
 ******************************************/

// https://docs.joinmastodon.org/methods/announcements/#get
const GetAnnouncements = "/api/v1/announcements"

// https://docs.joinmastodon.org/methods/announcements/#dismiss
const PostAnnoucement_Dismis = "/api/v1/announcements/:id/dismiss"

// https://docs.joinmastodon.org/methods/announcements/#put-reactions
const PutAnnouncement_Reaction = "/api/v1/announcements/:id/reactions/:name"

// https://docs.joinmastodon.org/methods/announcements/#delete-reactions
const DeleteAnnouncement_Reaction = "/api/v1/announcements/:id/reactions/:name"
