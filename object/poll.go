package object

// Poll represents a poll attached to a status.
// https://docs.joinmastodon.org/entities/Poll/
type Poll struct {
	ID          string        `json:"id"`                     // The ID of the poll in the database.
	ExpiresAt   string        `json:"expires_at"`             // When the poll ends. (ISO 8601 Datetime)
	Expired     bool          `json:"expired"`                // Is the poll currently expired?
	Multiple    bool          `json:"multiple"`               // Does the poll allow multiple-choice answers?
	VotesCount  int           `json:"votes_count"`            // How many votes have been received.
	VotersCount int           `json:"voters_count,omitempty"` // How many unique accounts have voted on a multiple-choice poll.
	Options     []PollOption  `json:"options"`                // Possible answers for the poll.
	Emojis      []CustomEmoji `json:"emojis"`                 // Custom emoji to be used for rendering poll options.
	Voted       bool          `json:"voted,omitempty"`        // When called with a user token, has the authorized user voted?
	OwnVotes    []int         `json:"own_votes"`              // When called with a user token, which options has the authorized user chosen? Contains an array of index values for options.
}

// https://docs.joinmastodon.org/entities/Poll/#Option
type PollOption struct {
	Title      string `json:"title"`       // The text value of the poll option.
	VotesCount int    `json:"votes_count"` // The total number of received votes for this option.
}
