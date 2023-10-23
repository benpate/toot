package txn

/******************************************
 * Filters API Methods
 * Create and manage filters
 * https://docs.joinmastodon.org/methods/filters/
 ******************************************/

// https://docs.joinmastodon.org/methods/filters/#get
// GET /api/v2/filters
// Returns: Array of Filter
type GetFilters struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
}

// https://docs.joinmastodon.org/methods/filters/#get-one
// GET /api/v2/filters/:id
// Returns: Filter
type GetFilter struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/filters/#create
// POST /api/v2/filters
// Returns: Filter
// Create a filter group with the given parameters.
type PostFilter struct {
	Host               string   `header:"Host"`
	Authorization      string   `header:"Authorization"`
	Title              string   `form:"title"`
	Context            []string `form:"context"`
	FilterAction       string   `form:"filter_action"`
	ExpiresIn          int      `form:"expires_in"`
	KeywordsAttributes []struct {
		Host      string `header:"Host"`
		Keyword   string `form:"keyword"`
		WholeWord bool   `form:"whole_word"`
	} `form:"keywords_attributes"`
}

// https://docs.joinmastodon.org/methods/filters/#update
// PUT /api/v2/filters/:id
// Returns: Filter
// Update a filter group with the given parameters.
type PutFilter struct {
	Host               string   `header:"Host"`
	Authorization      string   `header:"Authorization"`
	ID                 string   `param:"id"`           // The ID of the Filter in the database.
	Title              string   `form:"title"`         // The name of the filter group.
	Context            []string `form:"context"`       // Where the filter should be applied. Specify at least one of home, notifications, public, thread, account.
	FilterAction       string   `form:"filter_action"` // The policy to be applied when the filter is matched. Specify warn or hide.
	ExpiresIn          int      `form:"expires_in"`    // How many seconds from now should the filter expire?
	KeywordsAttributes []struct {
		Host      string `header:"Host"`
		Keyword   string `form:"keyword"`    // A keyword to be added to the newly-created filter group.
		WholeWord bool   `form:"whole_word"` // Whether the keyword should consider word boundaries.
		ID        string `form:"id"`         // Provide the ID of an existing keyword to modify it, instead of creating a new keyword.
		Destroy   bool   `form:"_destroy"`   // If true, will remove the keyword with the given ID.
	} `form:"keywords_attributes"`
}

// https://docs.joinmastodon.org/methods/filters/#delete
// DELETE /api/v2/filters/:id
// Returns: Empty object
// Delete a filter group.
type DeleteFilter struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"` // The ID of the Filter in the database.
}

// https://docs.joinmastodon.org/methods/filters/#keywords-get
// GET /api/v2/filters/:filter_id/keywords
// Returns: Array of FilterKeyword
// List all keywords attached to the current filter group.
type GetFilter_Keywords struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	FilterID      string `param:"filter_id"` // The ID of the Filter in the database.
}

// https://docs.joinmastodon.org/methods/filters/#keywords-create
// POST /api/v2/filters/:filter_id/keywords
// Returns: FilterKeyword
// Add the given keyword to the specified filter group
type PostFilter_Keyword struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	FilterID      string `param:"filter_id"` // The ID of the Filter in the database.
	Keyword       string `form:"keyword"`    // The keyword to be added to the filter group.
	WholeWord     bool   `form:"whole_word"` // Whether the keyword should consider word boundaries.
}

// https://docs.joinmastodon.org/methods/filters/#keywords-get-one
// GET /api/v2/filters/keywords/:id
// Returns: FilterKeyword
// Get one filter keyword by the given id
type GetFilter_Keyword struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"` // The ID of the FilterKeyword in the database.
}

// https://docs.joinmastodon.org/methods/filters/#keywords-update
// PUT /api/v2/filters/keywords/:id
// Returns: FilterKeyword
// Update the given filter keyword.
type PutFilter_Keyword struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`        // The ID of the FilterKeyword in the database.
	Keyword       string `form:"keyword"`    // The keyword to be added to the filter group.
	WholeWord     bool   `form:"whole_word"` // Whether the keyword should consider word boundaries.
}

// https://docs.joinmastodon.org/methods/filters/#keywords-delete
// DELETE /api/v2/filters/keywords/:id
// Returns: Empty object
// Delete the given filter keyword.
type DeleteFilter_Keyword struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"` // The ID of the FilterKeyword in the database.
}

// https://docs.joinmastodon.org/methods/filters/#statuses-get
// GET /api/v2/filters/:filter_id/statuses
// Returns: Array of FilterStatus
// Obtain a list of all status filters within this filter group.
type GetFilter_Statuses struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	FilterID      string `param:"filter_id"` // The ID of the Filter in the database.
}

// https://docs.joinmastodon.org/methods/filters/#statuses-add
// POST /api/v2/filters/:filter_id/statuses
// Returns: FilterStatus
// Add a status filter to the current filter group.
type PostFilter_Status struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	FilterID      string `param:"filter_id"` // The ID of the Filter in the database.
}

// https://docs.joinmastodon.org/methods/filters/#statuses-get-one
// GET /api/v2/filters/statuses/:id
// Returns: FilterStatus
// Obtain a single status filter.
type GetFilter_Status struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"` // The ID of the FilterStatus in the database.
}

// https://docs.joinmastodon.org/methods/filters/#statuses-remove
// DELETE /api/v2/filters/statuses/:id
// Returns: FilterStatus
// Remove a status filter from the current filter group.
type DeleteFilter_Status struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"` // The ID of the FilterStatus in the database.
}

// https://docs.joinmastodon.org/methods/filters/#get-v1
// GET /api/v1/filters
// Returns List of V1::"Filter"
type GetFilters_V1 struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
}

// https://docs.joinmastodon.org/methods/filters/#get-one-v1
// GET /api/v1/filters/:id
// Returns V1::"Filter"
type GetFilter_V1 struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/filters/#create-v1
// POST /api/v1/filters
// Returns V1::"Filter"
type PostFilter_V1 struct {
	Host          string   `header:"Host"`
	Authorization string   `header:"Authorization"`
	Phrase        string   `form:"phrase"`
	Context       []string `form:"context"`
	Irreversible  bool     `form:"irreversible"`
	WholeWord     bool     `form:"whole_word"`
	ExpiresIn     int      `form:"expires_in"`
}

// https://docs.joinmastodon.org/methods/filters/#update-v1
// PUT /api/v1/filters/:id
// Returns V1::"Filter"
type PutFilter_V1 struct {
	Host          string   `header:"Host"`
	Authorization string   `header:"Authorization"`
	ID            string   `param:"id"`
	Phrase        string   `form:"phrase"`
	Context       []string `form:"context"`
	Irreversible  bool     `form:"irreversible"`
	WholeWord     bool     `form:"whole_word"`
	ExpiresIn     int      `form:"expires_in"`
}

// https://docs.joinmastodon.org/methods/filters/#delete-v1
// DELETE /api/v1/filters/:id
// Returns Empty object
type DeleteFilter_V1 struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}
