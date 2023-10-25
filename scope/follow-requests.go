package scope

/******************************************
 * Follow Requests API Methods
 * View and manage follow requests
 * https://docs.joinmastodon.org/methods/follow_requests/
 ******************************************/

// https://docs.joinmastodon.org/methods/follow_requests/#get
const GetFollowRequests = ReadFollows

// https://docs.joinmastodon.org/methods/follow_requests/#accept
const PostFollowRequest_Authorize = WriteFollows

// https://docs.joinmastodon.org/methods/follow_requests/#reject
const PostFollowRequest_Reject = WriteFollows
