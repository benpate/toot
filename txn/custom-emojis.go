package txn

/******************************************
 * Custom Emoji API Methods
 * Each site can define and upload its own custom emoji to be attached to profiles or statuses.
 * https://docs.joinmastodon.org/methods/custom_emojis/
 ******************************************/

// https://docs.joinmastodon.org/methods/custom_emojis/#get
// GET /api/v1/custom_emojis
// Returns: Array of CustomEmoji
// Returns custom emojis that are available on the server.
type GetCustomEmojis struct{}
