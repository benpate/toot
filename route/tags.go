package route

/******************************************
 * Tags API Methods
 * View information about or follow/unfollow hashtags
 * https://docs.joinmastodon.org/methods/tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/tags/#get
const GetTag = "/api/v1/tags/:id"

// https://docs.joinmastodon.org/methods/tags/#follow
const PostTag_Follow = "/api/v1/tags/:id/follow"

// https://docs.joinmastodon.org/methods/tags/#unfollow
const PostTag_Unfollow = "/api/v1/tags/:id/unfollow"
