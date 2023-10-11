package transaction

/******************************************
 * Statuses API Methods
 * Publish, interact, and view information about statuses
 * https://docs.joinmastodon.org/methods/statuses/
 ******************************************/

// https://docs.joinmastodon.org/methods/statuses/#create
// POST /api/v1/statuses
// Returns: Status
type PostStatus struct {
	Status   string   `form:"status"`
	MediaIDs []string `form:"media_ids"`
	Poll     struct {
		Options    []string `form:"options"`
		ExpiresIn  int      `form:"expires_in"`
		Multiple   bool     `form:"multiple"`
		HideTotals bool     `form:"hide_totals"`
	} `form:"poll"`
	InReplyToID string `form:"in_reply_to_id"`
	Sensitive   bool   `form:"sensitive"`
	SpoilerText string `form:"spoiler_text"`
	Visibility  string `form:"visibility"`   // [public | unlisted | private | direct]
	Language    string `form:"language"`     // ISO 629 2-letter language code
	ScheduledAt string `form:"scheduled_at"` // ISO 8601 Datetime
}

// https://docs.joinmastodon.org/methods/statuses/#get
// GET /api/v1/statuses/:id
// Returns: Status
type GetStatus struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#delete
// DELETE /api/v1/statuses/:id
type DeleteStatus struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#context
// GET /api/v1/statuses/:id/context
// Returns: Context

type GetStatus_Context struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#translate
// POST /api/v1/statuses/:id/translate
// Returns: Status
type PostStatus_Translate struct {
	ID   string `param:"id"`
	Lang string `form:"lang"`
}

// https://docs.joinmastodon.org/methods/statuses/#reblogged_by
// GET /api/v1/statuses/:id/reblogged_by
// Returns: []Account
type GetStatus_RebloggedBy struct {
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/statuses/#favourited_by
// GET /api/v1/statuses/:id/favourited_by
// Returns: []Account
type GetStatus_FavouritedBy struct {
	ID      string `param:"id"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	Limit   int    `query:"limit"`
}

// https://docs.joinmastodon.org/methods/statuses/#favourite
// POST /api/v1/statuses/:id/favourite
// Returns: Status
type PostStatus_Favourite struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#unfavourite
// POST /api/v1/statuses/:id/unfavourite
// Returns: Status
type PostStatus_Unfavourite struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#reblog
// POST /api/v1/statuses/:id/reblog
// Returns: Status
type PostStatus_Reblog struct {
	ID         string `param:"id"`
	Visibility string `form:"visibility"` // [public | unlisted | private]
}

// https://docs.joinmastodon.org/methods/statuses/#unreblog
// POST /api/v1/statuses/:id/unreblog
// Returns: Status
type PostStatus_Unreblog struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#bookmark
// POST /api/v1/statuses/:id/bookmark
// Returns: Status
type PostStatus_Bookmark struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#unbookmark
// POST /api/v1/statuses/:id/unbookmark
// Returns: Status
type PostStatus_Unbookmark struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#mute
// POST /api/v1/statuses/:id/mute
// Returns: Status
type PostStatus_Mute struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#unmute
// POST /api/v1/statuses/:id/unmute
// Returns: Status
type PostStatus_Unmute struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#pin
// POST /api/v1/statuses/:id/pin
// Returns: Status
type PostStatus_Pin struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#unpin
// POST /api/v1/statuses/:id/unpin
// Returns: Status
type PostStatus_Unpin struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#edit
// PUT /api/v1/statuses/:id
// Returns: Status
type PutStatus struct {
	ID          string   `param:"id"`
	Status      string   `form:"status"`
	SpoilerText string   `form:"spoiler_text"`
	Sensitive   bool     `form:"sensitive"`
	Language    string   `form:"language"`
	MediaIDs    []string `form:"media_ids[]"`
	Poll        struct {
		Options    []string `form:"options[]"`
		ExpiresIn  int      `form:"expires_in"`
		Multiple   bool     `form:"multiple"`
		HideTotals bool     `form:"hide_totals"`
	} `form:"poll"`
}

// https://docs.joinmastodon.org/methods/statuses/#history
// GET /api/v1/statuses/:id/history
// Returns: []StatusEdit
type GetStatus_History struct {
	ID string `param:"id"`
}

// https://docs.joinmastodon.org/methods/statuses/#source
// GET /api/v1/statuses/:id/source
// Returns: StatusSource
type GetStatus_Source struct {
	ID string `param:"id"`
}
