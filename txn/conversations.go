package txn

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
	Host string `header:"Host"`

	QueryPage
}

// https://docs.joinmastodon.org/methods/conversations/#delete
// DELETE /api/v1/conversations/:id
// Returns: Empty struct
type DeleteConversation struct {
	Host string `header:"Host"`
	ID   string `uri:"id"`
}

// https://docs.joinmastodon.org/methods/conversations/#read
// POST /api/v1/conversations/:id/read
// Returns: Conversation
type PostConversationRead struct {
	Host string `header:"Host"`
	ID   string `uri:"id"`
}
