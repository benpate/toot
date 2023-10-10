package model

// https://docs.joinmastodon.org/entities/Status/#Tag
type StatusTag struct {
	Name string `json:"name"` // The value of the hashtag after the # sign.
	URL  string `json:"url"`  // A link to the hashtag on the instance.
}
