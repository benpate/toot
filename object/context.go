package object

// Context represents the tree around a given status. Used for reconstructing threads of statuses.
// https://docs.joinmastodon.org/entities/Context/
type Context struct {
	Ancestors   []Status `json:"ancestors"`   // Parents in the thread.
	Descendants []Status `json:"descendants"` // Children in the thread.
}
