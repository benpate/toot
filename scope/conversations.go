package scope

/******************************************
 * Conversations API Methods
 * Direct conversations with other participants.
 * (Currently, just threads containing a post with "direct" visibility.)
 * https://docs.joinmastodon.org/methods/conversations/
 ******************************************/

// https://docs.joinmastodon.org/methods/conversations/#get
const GetConversations = ReadStatuses

// https://docs.joinmastodon.org/methods/conversations/#delete
const DeleteConversation = WriteConversations

// https://docs.joinmastodon.org/methods/conversations/#read
const PostConversationRead = WriteConversations
