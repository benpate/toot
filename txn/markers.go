package txn

/******************************************
 * Markers API Methods
 * Save and restore your position in timelines
 * https://docs.joinmastodon.org/methods/markers/
 ******************************************/

// https://docs.joinmastodon.org/methods/markers/#get
// GET /api/v1/markers
// Returns: Marker
type GetMarkers struct {
	Host     string   `header:"Host"`
	Timeline []string `query:"timeline[]"`
}

// https://docs.joinmastodon.org/methods/markers/#create
// POST /api/v1/markers
// Returns: Marker
type PostMarker struct {
	Host string `header:"Host"`
	Home struct {
		Host       string `header:"Host"`
		LastReadID string `form:"last_read_id"`
	} `form:"home"`
	Notifications struct {
		Host       string `header:"Host"`
		LastReadID string `form:"last_read_id"`
	} `form:"notifications"`
}
