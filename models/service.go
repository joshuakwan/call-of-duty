package models

// Service defines a service
type Service struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	TeamID             string   `json:"team_id"`
	UserIDs            []string `json:"user_ids"`
	EscalationChainIDs string   `json:"escalation_chain_ids"`
}
