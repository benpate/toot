package transaction

/******************************************
 * Reports API Methods
 * Report problematic users to your moderators
 * https://docs.joinmastodon.org/methods/reports/
 ******************************************/

// https://docs.joinmastodon.org/methods/reports/#post
// POST /api/v1/reports
// Returns: Report
type PostReport struct {
	AccountID string   `form:"account_id"`
	StatusIDs []string `form:"status_ids"`
	Comment   string   `form:"comment"`
	Forward   bool     `form:"forward"`
	Category  string   `form:"category"`
	RuleIDs   []string `form:"rule_ids"`
}
