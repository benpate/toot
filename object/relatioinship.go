package object

// Relationship represents the relationship between accounts, such as following / blocking / muting / etc.
// https://docs.joinmastodon.org/entities/Relationship/
type Relationship struct {
	ID                  string   `json:"id"`                   // The account ID.
	Following           bool     `json:"following"`            // Are you following this user?
	ShowingReblogs      bool     `json:"showing_reblogs"`      // Are you receiving this user’s boosts in your home timeline?
	Notifying           bool     `json:"notifying"`            // Have you enabled notifications for this user?
	Languages           []string `json:"languages"`            // Which languages are you following from this user? ISO 8601 2-letter language codes
	FollowedBy          bool     `json:"followed_by"`          // Are you followed by this user?
	Blocking            bool     `json:"blocking"`             // Are you blocking this user?
	BlockedBy           bool     `json:"blocked_by"`           // Is this user blocking you?
	Muting              bool     `json:"muting"`               // Are you muting this user?
	MutingNotifications bool     `json:"muting_notifications"` // Are you muting notifications from this user?
	Requested           bool     `json:"requested"`            // Do you have a pending follow request for this user?
	DomainBlocking      bool     `json:"domain_blocking"`      // Are you blocking this user’s domain?
	Endorsed            bool     `json:"endorsed"`             // Are you featuring this user on your profile?
	Note                string   `json:"note"`                 // The user’s profile bio
}
