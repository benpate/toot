package transaction

/******************************************
 * OEmbed API Methods
 * For generating OEmbed previews
 * https://docs.joinmastodon.org/methods/oembed/
 ******************************************/

// https://docs.joinmastodon.org/methods/oembed/#get
// GET /api/oembed
// Returns: OEmbed metadata
type GetOEmbed struct {
	URL       string `query:"url"`
	MaxWidth  int    `query:"maxwidth"`
	MaxHeight int    `query:"maxheight"`
}
