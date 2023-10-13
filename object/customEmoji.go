package object

// CustomEmoji represents a custom emoji.
// https://docs.joinmastodon.org/entities/CustomEmoji/
type CustomEmoji struct {
	ShortCode       string `json:"shortcode"`         // The name of the custom emoji.
	URL             string `json:"url"`               // A link to the custom emoji.
	StaticURL       string `json:"static_url"`        // A link to a static copy of the custom emoji.
	VisibleInPicker bool   `json:"visible_in_picker"` // Whether this Emoji should be visible in the picker or unlisted.
	Category        string `json:"category"`          // Used for sorting custom emoji in the picker.
}
