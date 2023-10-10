package model

// FilterKeyword represents a keyword that, if matched, should cause the filter action to be taken.
// https://docs.joinmastodon.org/entities/FilterKeyword/
type FilterKeyword struct {
	ID        string `json:"id"`         // The ID of the FilterKeyword in the database.
	Keyword   string `json:"keyword"`    // The phrase to be matched against.
	WholeWord bool   `json:"whole_word"` // Should the filter consider word boundaries? See implementation guidelines for filters. https://docs.joinmastodon.org/api/guidelines/#filters
}
