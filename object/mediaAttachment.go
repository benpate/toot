package object

// MediaAttachment represents a file or media attachment that can be added to a status.
// https://docs.joinmastodon.org/entities/MediaAttachment/
type MediaAttachment struct {
	ID          string         `json:"id"`                   // The ID of the attachment in the database.
	Type        string         `json:"type"`                 // The type of the attachment. [ image | gifv | video | audio ]
	URL         string         `json:"url"`                  // The location of the original full-size attachment.
	PreviewURL  string         `json:"preview_url"`          // The location of a scaled-down preview of the attachment.
	RemoteURL   string         `json:"remote_url,omitempty"` // The location of the full-size original attachment on the remote website.
	Meta        map[string]any `json:"meta"`                 // Metadata returned by Paperclip. https://docs.joinmastodon.org/api/guidelines/#focal-points
	Description string         `json:"description"`          // Alternate text that describes what is in the media attachment, to be used for the visually impaired or when media attachments do not load.
	Blurhash    string         `json:"blurhash"`             // A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.  https://github.com/woltapp/blurhash
}
