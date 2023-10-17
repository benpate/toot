package route

/******************************************
 * Timelines API Methods
 * Read and view timelines of statuses
 * https://docs.joinmastodon.org/methods/timelines/
 ******************************************/

// https://docs.joinmastodon.org/methods/timelines/#public
const GetTimeline_Public = "/api/v1/timelines/public"

// https://docs.joinmastodon.org/methods/timelines/#tag
const GetTimeline_Hashtag = "/api/v1/timelines/tag/:hashtag"

// https://docs.joinmastodon.org/methods/timelines/#home
const GetTimeline_Home = "/api/v1/timelines/home"

// https://docs.joinmastodon.org/methods/timelines/#list
const GetTimeline_List = "/api/v1/timelines/list/:list_id"
