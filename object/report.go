package object

// Report represents reports filed against users and/or statuses, to be taken action on by moderators.
// https://docs.joinmastodon.org/entities/Report/
type Report struct {
	ID            string   `json:"id"`                   // The ID of the report in the database.
	ActionTaken   bool     `json:"action_taken"`         // Whether an action was taken yet.
	ActionTakenAt string   `json:"action_taken_at"`      // When an action was taken against the report. (ISO 8601 datetime)
	Category      string   `json:"category"`             // The generic reason for the report. (spam | violoation |  other)
	Comment       string   `json:"comment"`              // The reason for the report.
	Forwarded     bool     `json:"forwarded"`            // Whether the report was forwarded to a remote domain.
	CreatedAt     string   `json:"created_at"`           // When the report was created.
	StatusIDs     []string `json:"status_ids,omitempty"` // IDs of statuses that have been attached to this report for additional context.
	RuleIDs       []string `json:"rule_ids,omitempty"`   // IDs of the rules that have been cited as a violation by this report.
	TargetAcount  Account  `json:"target_account"`       // The account that was reported.
}
