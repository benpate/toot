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
	Authorization string   `header:"Authorization"`
	Timeline      []string `query:"timeline[]"`
}

// https://docs.joinmastodon.org/methods/markers/#create
// POST /api/v1/markers
// Returns: Marker
type PostMarker struct {
	Authorization string `header:"Authorization"`
	Home          struct {
		LastReadID string `form:"last_read_id"`
	} `form:"home"`
	Notifications struct {
		LastReadID string `form:"last_read_id"`
	} `form:"notifications"`
}
