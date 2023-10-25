package scope

/******************************************
 * Timelines API Methods
 * Read and view timelines of statuses
 * https://docs.joinmastodon.org/methods/timelines/
 ******************************************/

// https://docs.joinmastodon.org/methods/timelines/#public
const GetTimeline_Public = ReadStatuses

// https://docs.joinmastodon.org/methods/timelines/#tag
const GetTimeline_Hashtag = ReadStatuses

// https://docs.joinmastodon.org/methods/timelines/#home
const GetTimeline_Home = ReadStatuses

// https://docs.joinmastodon.org/methods/timelines/#list
const GetTimeline_List = ReadStatuses
