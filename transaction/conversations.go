package transaction

/******************************************
 * Conversations API Methods
 * Direct conversations with other participants.
 * (Currently, just threads containing a post with "direct" visibility.)
 * https://docs.joinmastodon.org/methods/conversations/
 ******************************************/

// https://docs.joinmastodon.org/methods/conversations/#get
// GET /api/v1/conversations
// Returns: Array of Conversation
type GetConversations struct {
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/conversations/#delete
// DELETE /api/v1/conversations/:id
// Returns: Empty struct
type DeleteConversation struct {
	ID string `uri:"id"`
}

// https://docs.joinmastodon.org/methods/conversations/#read
// POST /api/v1/conversations/:id/read
// Returns: Conversation
type PostConversationRead struct {
	ID string `uri:"id"`
}
