package scope

/******************************************
 * Tags API Methods
 * View information about or follow/unfollow hashtags
 * https://docs.joinmastodon.org/methods/tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/tags/#get
const GetTag = Public

// https://docs.joinmastodon.org/methods/tags/#follow
const PostTag_Follow = WriteFollows

// https://docs.joinmastodon.org/methods/tags/#unfollow
const PostTag_Unfollow = WriteFollows
