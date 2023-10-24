package txn

/******************************************
 * Favourites API Methods
 * View your favourites. See also /statuses/:id/{favourite,unfavourite}
 * https://docs.joinmastodon.org/methods/favourites/
 ******************************************/

// https://docs.joinmastodon.org/methods/favourites/#get
// GET /api/v1/favourites
// Returns: Array of Status
type GetFavourites struct {
	Host string `header:"Host"`

	QueryPage
}
