package txn

/******************************************
 * Timelines API Methods
 * Read and view timelines of statuses
 * https://docs.joinmastodon.org/methods/timelines/
 ******************************************/

// https://docs.joinmastodon.org/methods/timelines/#public
// GET /api/v1/timelines/public
// Returns: []Status
type GetTimeline_Public struct {
	Host      string `header:"Host"`
	Local     bool   `query:"local"`
	Remote    bool   `query:"remote"`
	OnlyMedia bool   `query:"only_media"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/timelines/#tag
// GET /api/v1/timelines/tag/:hashtag
// Returns: []Status
type GetTimeline_Hashtag struct {
	Host      string   `header:"Host"`
	Hashtag   string   `param:"hashtag"`
	Any       []string `query:"any"`
	All       []string `query:"all"`
	None      []string `query:"none"`
	Local     bool     `query:"local"`
	Remote    bool     `query:"remote"`
	OnlyMedia bool     `query:"only_media"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/timelines/#home
// GET /api/v1/timelines/home
// Returns: []Status
type GetTimeline_Home struct {
	Host string `header:"Host"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/timelines/#list
// GET /api/v1/timelines/list/:list_id
// Returns: []Status
type GetTimeline_List struct {
	Host   string `header:"Host"`
	ListID string `param:"list_id"`

	QueryPage
}
