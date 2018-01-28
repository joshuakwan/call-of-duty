package models

// Service defines a service
type Service struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	TeamID              string   `json:"team_id"`
	EscalationPolicyIDs []string `json:"escalation_policy_ids"`
}
