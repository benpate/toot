package object

// InstanceConfiguration contains configured values and limits for this website.
// https://docs.joinmastodon.org/entities/Instance/#configuration
type InstanceConfiguration struct {
	URLs     InstanceConfigurationURLs     `json:"urls"` // URLs of interest for clients apps.
	Accounts InstanceConfigurationAccounts `json:"accounts"`
}

// URLs of interest for clients apps.
// https://docs.joinmastodon.org/entities/Instance/#urls
type InstanceConfigurationURLs struct {
	StreamingAPI string `json:"streaming_api"` // The Websockets URL for connecting to the streaming API.
}

// Limits related to accounts.
// https://docs.joinmastodon.org/entities/Instance/#accounts
type InstanceConfigurationAccounts struct {
	MaxFeaturedTags int `json:"max_featured_tags"` // The maximum number of featured tags allowed for each account.
}

// Limits related to authoring statuses.
// https://docs.joinmastodon.org/entities/Instance/#statuses
type InstanceConfigurationStatus struct {
	MaxMediaAttachments      int `json:"max_media_attachments"`         // The maximum number of media attachments that can be added to a status.
	CharactersReservedPerURL int `json:"characters_reserved_per_media"` // Each URL in a status will be assumed to be exactly this many characters.
}

// Hints for which attachments will be accepted.
// https://docs.joinmastodon.org/entities/Instance/#media_attachments
type InstanceConfigurationMediaAttachments struct {
	SupportedMimeTypes []string `json:"mime_types"`         // Contains MIME types that can be uploaded.
	ImageSizeLimit     int      `json:"image_size_limit"`   // The maximum size of any uploaded image, in bytes.
	ImageMatrixLimit   int      `json:"image_matrix_limit"` // The maximum number of pixels (width times height) for image uploads.
	VideoSizeLimit     int      `json:"video_size_limit"`   // The maximum size of any uploaded video, in bytes.
	VideoFrameLimit    int      `json:"video_frame_limit"`  // The maximum number of pixels (width times height) for video uploads.
}

//	Limits related to polls.
//
// https://docs.joinmastodon.org/entities/Instance/#polls
type InstanceConfigurationPolls struct {
	MaxOptions    int `json:"max_options"`    // Each poll is allowed to have up to this many options.
	MinExpiration int `json:"min_expiration"` // The shortest allowed poll duration, in seconds.
	MaxExpiration int `json:"max_expiration"` // The longest allowed poll duration, in seconds.
}

// Hints related to translation.
// https://docs.joinmastodon.org/entities/Instance/#translation
type InstanceConfigurationTranslation struct {
	Enabled bool `json:"enabled"` // Whether the Translations API is available on this instance.
}
