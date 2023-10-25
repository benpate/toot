package scope

/******************************************
 * Statuses API Methods
 * Publish, interact, and view information about statuses
 * https://docs.joinmastodon.org/methods/statuses/
 ******************************************/

// https://docs.joinmastodon.org/methods/statuses/#create
const PostStatus = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#get
const GetStatus = ReadStatuses

// https://docs.joinmastodon.org/methods/statuses/#delete
const DeleteStatus = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#context
const GetStatus_Context = ReadStatuses

// https://docs.joinmastodon.org/methods/statuses/#translate
const PostStatus_Translate = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#reblogged_by
const GetStatus_RebloggedBy = ReadStatuses

// https://docs.joinmastodon.org/methods/statuses/#favourited_by
const GetStatus_FavouritedBy = ReadStatuses

// https://docs.joinmastodon.org/methods/statuses/#favourite
const PostStatus_Favourite = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#unfavourite
const PostStatus_Unfavourite = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#reblog
const PostStatus_Reblog = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#unreblog
const PostStatus_Unreblog = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#bookmark
const PostStatus_Bookmark = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#unbookmark
const PostStatus_Unbookmark = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#mute
const PostStatus_Mute = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#unmute
const PostStatus_Unmute = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#pin
const PostStatus_Pin = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#unpin
const PostStatus_Unpin = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#edit
const PutStatus = WriteStatuses

// https://docs.joinmastodon.org/methods/statuses/#history
const GetStatus_History = ReadStatuses

// https://docs.joinmastodon.org/methods/statuses/#source
const GetStatus_Source = ReadStatuses
