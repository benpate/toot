package route

/******************************************
 * Conversations API Methods
 * Direct conversations with other participants.
 * (Currently, just threads containing a post with "direct" visibility.)
 * https://docs.joinmastodon.org/methods/conversations/
 ******************************************/

// https://docs.joinmastodon.org/methods/conversations/#get
const GetConversations = "/api/v1/conversations"

// https://docs.joinmastodon.org/methods/conversations/#delete
const DeleteConversation = "/api/v1/conversations/:id"

// https://docs.joinmastodon.org/methods/conversations/#read
const PostConversationRead = "/api/v1/conversations/:id/read"
