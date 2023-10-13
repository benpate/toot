package transaction

/******************************************
 * Timelines API Methods
 * Read and view timelines of statuses
 * https://docs.joinmastodon.org/methods/timelines/
 ******************************************/

// https://docs.joinmastodon.org/methods/timelines/#public
// GET /api/v1/timelines/public
// Returns: []Status
type GetTimeline_Public struct {
	Local     bool   `query:"local"`
	Remote    bool   `query:"remote"`
	OnlyMedia bool   `query:"only_media"`
	MaxID     string `query:"max_id"`
	SinceID   string `query:"since_id"`
	MinID     string `query:"min_id"`
	Limit     int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/timelines/#tag
// GET /api/v1/timelines/tag/:hashtag
// Returns: []Status
type GetTimeline_Hashtag struct {
	Hashtag   string   `param:"hashtag"`
	Any       []string `query:"any"`
	All       []string `query:"all"`
	None      []string `query:"none"`
	Local     bool     `query:"local"`
	Remote    bool     `query:"remote"`
	OnlyMedia bool     `query:"only_media"`
	MaxID     string   `query:"max_id"`
	SinceID   string   `query:"since_id"`
	MinID     string   `query:"min_id"`
	Limit     int      `query:"limit"`
}

// https://docs.joinmastodon.org/methods/timelines/#home
// GET /api/v1/timelines/home
// Returns: []Status
type GetTimeline_Home struct {
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/timelines/#list
// GET /api/v1/timelines/list/:list_id
// Returns: []Status
type GetTimeline_List struct {
	ListID  string `param:"list_id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}
