package object

// Notification represents a notification of an event relevant to the user.
// https://docs.joinmastodon.org/entities/Notification/
type Notification struct {
	ID        string       `json:"id"`               // The id of the notification in the database.
	Type      string       `json:"type"`             // The type of event that resulted in the notification. [mention | status | reblog | follow | follow_request |favourite | poll | update | admin.sign_up | admin.report ]
	CreatedAt string       `json:"created_at"`       // The timestamp of the notification. (ISO 8601 Datetime)
	Account   Account      `json:"account"`          // The account that performed the action that generated the notification.
	Status    *Status      `json:"status,omitempty"` // Status that was the object of the notification. Attached when type of the notification is favourite, reblog, status, mention, poll, or update.
	Report    *AdminReport `json:"report,omitempty"` // Report that was the object of the notification. Attached when type of the notification is admin.report.
}
