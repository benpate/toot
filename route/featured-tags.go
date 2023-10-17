package route

/******************************************
 * Featured Tags API Methods
 * Feature tags that you use frequently on your profile
 * https://docs.joinmastodon.org/methods/featured_tags/
 ******************************************/

// https://docs.joinmastodon.org/methods/featured_tags/#get
const GetFeaturedTags = "/api/v1/featured_tags"

// https://docs.joinmastodon.org/methods/featured_tags/#feature
const PostFeaturedTag = "/api/v1/featured_tags"

// https://docs.joinmastodon.org/methods/featured_tags/#unfeature
const DeleteFeaturedTag = "/api/v1/featured_tags/:id"

// https://docs.joinmastodon.org/methods/featured_tags/#suggestions
const GetFeaturedTags_Suggestions = "/api/v1/featured_tags/suggestions"
