package scope

/******************************************
 * Featured Tags API Methods
 * Feature tags that you use frequently on your profile
 * https://docs.joinmastodon.org/methods/featured_tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/featured_tags/#get
const GetFeaturedTags = ReadAccounts

// https://docs.joinmastodon.org/methods/featured_tags/#feature
const PostFeaturedTag = WriteAccounts

// https://docs.joinmastodon.org/methods/featured_tags/#unfeature
const DeleteFeaturedTag = WriteAccounts

// https://docs.joinmastodon.org/methods/featured_tags/#suggestions
const GetFeaturedTags_Suggestions = ReadAccounts
