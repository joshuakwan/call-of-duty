package models

import "github.com/joshuakwan/call-of-duty/utils"

// Service defines a service
type Service struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	TeamID             string   `json:"team_id"`
	UserIDs            []string `json:"user_ids"`
	EscalationChainIDs string   `json:"escalation_chain_ids"`
}

// CreateService creates a new service
func CreateService(name, email string) *Service {
	return &Service{ID: utils.GenerateUUID(), Name: name}
}
