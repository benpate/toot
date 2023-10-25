package scope

/******************************************
 * Announcements API Methods
 * For announcements set by administration
 * https://docs.joinmastodon.org/methods/announcements/
 ******************************************/

// https://docs.joinmastodon.org/methods/announcements/#get
const GetAnnouncements = "*"

// https://docs.joinmastodon.org/methods/announcements/#dismiss
const PostAnnoucement_Dismis = "write:accounts"

// https://docs.joinmastodon.org/methods/announcements/#put-reactions
const PutAnnouncement_Reaction = "write:favourites"

// https://docs.joinmastodon.org/methods/announcements/#delete-reactions
const DeleteAnnouncement_Reaction = "write:favourites"
