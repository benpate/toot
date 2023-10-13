package object

// ExtendedDescription represents an extended description for the instance, to be shown on its about page.
// https://docs.joinmastodon.org/entities/ExtendedDescription/
type ExtendedDescription struct {
	UpdatedAt string `json:"updated_at"` // A timestamp of when the extended description was last updated. (ISO 8601 Datetime)
	Content   string `json:"content"`    // The rendered HTML content of the extended description.
}
