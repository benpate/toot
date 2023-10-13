package object

// StatusEdit represents a revision of a status that has been edited.
// https://docs.joinmastodon.org/entities/StatusEdit/
type StatusEdit struct {
	Content          string            `json:"content"`           // The content of the status at this revision.
	SpoilerText      string            `json:"spoiler_text"`      // The content of the subject or content warning at this revision.
	Sensitive        bool              `json:"sensitive"`         // Whether the status was marked sensitive at this revision.
	CreatedAt        string            `json:"created_at"`        // The timestamp of when the revision was published. (ISO 8601 Datetime)
	Account          Account           `json:"account"`           // The account that published this revision.
	Poll             *Poll             `json:"poll"`              // The current state of the poll options at this revision. Note that edits changing the poll options will be collapsed together into one edit, since this action resets the poll.
	MediaAttachments []MediaAttachment `json:"media_attachments"` // The current state of the poll options at this revision. Note that edits changing the poll options will be collapsed together into one edit, since this action resets the poll.
	Emojis           []CustomEmoji     `json:"emojis"`            // Any custom emoji that are used in the current revision.
}
