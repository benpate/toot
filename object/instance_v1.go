package object

// https://docs.joinmastodon.org/entities/Instance_V1/
type Instance_V1 struct {
	URI              string `json:"uri"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Email            string `json:"email"`
	Version          string `json:"version"`
	URLs             struct {
		StreamingAPI string `json:"streaming_api"`
	}
	Stats struct {
		UserCount   int `json:"user_count"`
		StatusCount int `json:"status_count"`
		DomainCount int `json:"domain_count"`
	}
	Thumbnail        string                    `json:"thumbnail"`
	Languages        []string                  `json:"languages"`
	Registrations    bool                      `json:"registrations"`
	ApprovalRequired bool                      `json:"approval_required"`
	InvitesEnabled   bool                      `json:"invites_enabled"`
	Configuration    Instance_Configuration_V1 `json:"configuration"`
	ContactAccount   Account                   `json:"contact_account"`
	Rules            []Rule                    `json:"rules"`
}

type Instance_Configuration_V1 struct {
	Accounts         Instance_Accounts_V1 `json:"accounts"`
	Statuses         Instance_Statuses_V1 `json:"statuses"`
	MediaAttachments Instance_Media_V1    `json:"media_attachments"`
	Polls            Instance_Polls_V1    `json:"polls"`
}

type Instance_Accounts_V1 struct {
	MaxFeaturedTags int `json:"max_featured_tags"`
}

type Instance_Statuses_V1 struct {
	MaxMediaAttachments      int `json:"max_media_attachments"`
	CharactersReservedPerURL int `json:"characters_reserved_per_url"`
}

type Instance_Media_V1 struct {
	SupportedMimeTypes  []string `json:"supported_mime_types"`
	ImageSizeLimit      int      `json:"image_size_limit"`
	ImageMatrixLimit    int      `json:"image_matrix_limit"`
	VideoSizeLimit      int      `json:"video_size_limit"`
	VideoFrameRateLimit int      `json:"video_frame_rate_limit"`
	VideoMatrixLimit    int      `json:"video_matrix_limit"`
}

type Instance_Polls_V1 struct {
	MaxOptions             int `json:"max_options"`
	MaxCharactersPerOption int `json:"max_characters_per_option"`
	MinExpiration          int `json:"min_expiration"`
	MaxExpiration          int `json:"max_expiration"`
}
