package object

// StatusSource represents a status's source as plain text.
// https://docs.joinmastodon.org/entities/StatusSource/
type StatusSource struct {
	ID          string `json:"id"`           // ID of the status in the database.
	Text        string `json:"text"`         // The plain text used to compose the status.
	SpoilerText string `json:"spoiler_text"` // The plain text used to compose the status’s subject or content warning.
}
