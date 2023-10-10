package model

// https://docs.joinmastodon.org/entities/Instance/#usage
type InstanceUsage struct {
	Users InstanceUsageUsers `json:"users"` // Usage data related to users on this instance.
}

// https://docs.joinmastodon.org/entities/Instance/#active_month
type InstanceUsageUsers struct {
	ActiveMonth int `json:"active_month"` // The number of active users in the past 4 weeks.
}
