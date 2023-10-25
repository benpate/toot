package route

/******************************************
 * Statuses API Methods
 * Publish, interact, and view information about statuses
 * https://docs.joinmastodon.org/methods/statuses/
 ******************************************/

// https://docs.joinmastodon.org/methods/statuses/#create
const PostStatus = "/api/v1/statuses"

// https://docs.joinmastodon.org/methods/statuses/#get
const GetStatus = "/api/v1/statuses/:id"

// https://docs.joinmastodon.org/methods/statuses/#delete
const DeleteStatus = "/api/v1/statuses/:id"

// https://docs.joinmastodon.org/methods/statuses/#context
const GetStatus_Context = "/api/v1/statuses/:id/context"

// https://docs.joinmastodon.org/methods/statuses/#translate
const PostStatus_Translate = "/api/v1/statuses/:id/translate"

// https://docs.joinmastodon.org/methods/statuses/#reblogged_by
const GetStatus_RebloggedBy = "/api/v1/statuses/:id/reblogged_by"

// https://docs.joinmastodon.org/methods/statuses/#favourited_by
const GetStatus_FavouritedBy = "/api/v1/statuses/:id/favourited_by"

// https://docs.joinmastodon.org/methods/statuses/#favourite
const PostStatus_Favourite = "/api/v1/statuses/:id/favourite"

// https://docs.joinmastodon.org/methods/statuses/#unfavourite
const PostStatus_Unfavourite = "/api/v1/statuses/:id/unfavourite"

// https://docs.joinmastodon.org/methods/statuses/#reblog
const PostStatus_Reblog = "/api/v1/statuses/:id/reblog"

// https://docs.joinmastodon.org/methods/statuses/#unreblog
const PostStatus_Unreblog = "/api/v1/statuses/:id/unreblog"

// https://docs.joinmastodon.org/methods/statuses/#bookmark
const PostStatus_Bookmark = "/api/v1/statuses/:id/bookmark"

// https://docs.joinmastodon.org/methods/statuses/#unbookmark
const PostStatus_Unbookmark = "/api/v1/statuses/:id/unbookmark"

// https://docs.joinmastodon.org/methods/statuses/#mute
const PostStatus_Mute = "/api/v1/statuses/:id/mute"

// https://docs.joinmastodon.org/methods/statuses/#unmute
const PostStatus_Unmute = "/api/v1/statuses/:id/unmute"

// https://docs.joinmastodon.org/methods/statuses/#pin
const PostStatus_Pin = "/api/v1/statuses/:id/pin"

// https://docs.joinmastodon.org/methods/statuses/#unpin
const PostStatus_Unpin = "/api/v1/statuses/:id/unpin"

// https://docs.joinmastodon.org/methods/statuses/#edit
const PutStatus = "/api/v1/statuses/:id"

// https://docs.joinmastodon.org/methods/statuses/#history
const GetStatus_History = "/api/v1/statuses/:id/history"

// https://docs.joinmastodon.org/methods/statuses/#source
const GetStatus_Source = "/api/v1/statuses/:id/source"
