package model

// https://docs.joinmastodon.org/entities/Account/#source
type AccountSource struct {
	Note                string         `json:"note"`                  // Profile bio, in plain-text instead of in HTML.
	Fields              []AccountField `json:"fields"`                // Metadata about the account.
	Privacy             string         `json:"privacy"`               // The default post privacy to be used for new statuses. [public | unlisted | private | direct]
	Sensitive           bool           `json:"sensitive"`             // Whether new statuses should be marked sensitive by default.
	Language            string         `json:"language"`              // The default posting language for new statuses. (ISO 639-1 language two-letter code) or empty string
	FollowRequestsCount int            `json:"follow_requests_count"` // The number of pending follow requests.
}
