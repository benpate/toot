package model

// Instance represents the software instance of Mastodon running on this domain.
// https://docs.joinmastodon.org/entities/Instance/
type Instance struct {
	Domain        string                `json:"domain"`        // The domain name of the instance.
	Title         string                `json:"title"`         // The title of the website.
	Version       string                `json:"version"`       // The version of Mastodon installed on the instance.
	SourceURL     string                `json:"source_url"`    // The URL for the source code of the software running on this instance, in keeping with AGPL license requirements.
	Description   string                `json:"description"`   // A short, plain-text description defined by the admin.
	Usage         InstanceUsage         `json:"usage"`         // Usage data for this instance.
	Thumbnail     InstanceThumbnail     `json:"thumbnail"`     // An image used to represent this instance.
	Languages     []string              `json:"languages"`     // Primary languages of the website and its staff. (ISO 639-1 two letter code)
	Configuration InstanceConfiguration `json:"configuration"` // Configured values and limits for this website.
	Registrations InstanceRegistrations `json:"registrations"` // Information about registering for this website.
	Contact       InstanceContact       `json:"contact"`       // Hints related to contacting a representative of the website.
	Rules         []Rule                `json:"rules"`         // An itemized list of rules for this website.
}
