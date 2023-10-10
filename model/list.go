package model

// List represents a list of some users that the authenticated user follows.
// https://docs.joinmastodon.org/entities/List/
type List struct {
	ID            string `json:"id"`             // The internal database ID of the list.
	Title         string `json:"title"`          // The user-defined title of the list.
	RepliesPolicy string `json:"replies_policy"` // Which replies should be shown in the list. [followed | list | none]
}
