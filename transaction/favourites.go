package transaction

/******************************************
 * Favourites API Methods
 * View your favourites. See also /statuses/:id/{favourite,unfavourite}
 * https://docs.joinmastodon.org/methods/favourites/
 ******************************************/

// https://docs.joinmastodon.org/methods/favourites/#get
// GET /api/v1/favourites
// Returns: Array of Status
type GetFavourites struct {
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}
