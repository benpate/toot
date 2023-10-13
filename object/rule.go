package object

// Rule represents a rule that server users should follow.
// https://docs.joinmastodon.org/entities/Rule/
type Rule struct {
	ID   string `json:"id"`   // An identifier for the rule.
	Text string `json:"text"` // The rule to be followed.
}
