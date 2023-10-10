package model

// https://docs.joinmastodon.org/entities/Instance/#thumbnail
type InstanceThumbnail struct {
	URL      string                    `json:"url"`                // The URL for the thumbnail image.
	Blurhash string                    `json:"blurhash,omitempty"` // A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	Versions InstanceThumbnailVersions `json:"versions,omitempty"` // Links to scaled resolution images, for high DPI screens.
}

// https://docs.joinmastodon.org/entities/Instance/#thumbnail-versions
type InstanceThumbnailVersions struct {
	At1X string `json:"@1x,omitempty"` // The URL for the thumbnail image at 1x resolution.
	At2X string `json:"@2x,omitempty"` // The URL for the thumbnail image at 2x resolution.
}
