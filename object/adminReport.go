package object

// Admin::Report represents admin-level information about a filed report.
// https://docs.joinmastodon.org/entities/Admin_Report/
type AdminReport struct {
	ID                   string        `json:"id"`                                // The ID of the report in the database.
	ActionTaken          bool          `json:"action_taken"`                      // Whether an action was taken to resolve this report.
	ActionTakenAt        string        `json:"action_taken_at"`                   // When an action was taken, if this report is currently resolved. (ISO 8601 Datetime)
	Category             string        `json:"category"`                          // The category under which the report is classified. [spam | violation | other]
	Comment              string        `json:"comment"`                           // An optional reason for reporting.
	Forwarded            bool          `json:"forwarded"`                         // Whether a report was forwarded to a remote instance.
	CreatedAt            string        `json:"created_at"`                        // The time the report was filed. (ISO 8601 Datetime)
	UpdatedAt            string        `json:"updated_at"`                        // The time of last action on this report. (ISO 8601 Datetime)
	Account              AdminAccount  `json:"account"`                           // The account which filed the report.
	TargetAccount        AdminAccount  `json:"target_account"`                    // The account being reported.
	AssignedAccount      *AdminAccount `json:"assigned_account,omitempty"`        // The account of the moderator assigned to this report.
	ActionTakenByAccount *AdminAccount `json:"action_taken_by_account,omitempty"` // The account of the moderator who handled the report.
	Statuses             []Status      `json:"statuses"`                          // Statuses attached to the report, for context.
	Rules                []Rule        `json:"rules"`                             // Rules attached to the report, for context.
}
