package model

// Conversation represents a conversation with "direct message" visibility.
// https://docs.joinmastodon.org/entities/Conversation/
type Conversation struct {
	ID         string    `json:"id"`                    // The ID of the conversation in the database.
	Unread     bool      `json:"unread"`                // Is the conversation currently marked as unread?
	Accounts   []Account `json:"accounts"`              // Participants in the conversation.
	LastStatus Status    `json:"last_status,omitempty"` // The last status in the conversation.
}
