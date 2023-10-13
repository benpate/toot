package object

// https://docs.joinmastodon.org/entities/Tag/
type Tag struct {
	Name      string       `json:"name"`      //  The value of the hashtag after the # sign.
	URL       string       `json:"url"`       // A link to the hashtag on the instance.
	History   []TagHistory `json:"history"`   // Usage statistics for given days (typically the past week).
	Following bool         `json:"following"` // Whether the current token’s authorized user is following this tag.
}
