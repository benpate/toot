package object

// Reaction represents an emoji reaction to an Announcement.
// https://docs.joinmastodon.org/entities/Reaction/
type Reaction struct {
	Name      string `json:"name"`       // The emoji used for the reaction. Either a unicode emoji, or a custom emoji’s shortcode.
	Count     int    `json:"count"`      // The total number of users who have added this reaction.
	Me        bool   `json:"me"`         // If there is a currently authorized user: Have you added this reaction?
	URL       string `json:"url"`        // If the reaction is a custom emoji: A link to the custom emoji.
	StaticURL string `json:"static_url"` // If the reaction is a custom emoji: A link to a non-animated version of the custom emoji.
}
