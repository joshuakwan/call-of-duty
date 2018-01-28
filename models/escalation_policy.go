package models

// EscalationPolicy loop policy
const (
	LOOP = iota
	ONCE
)

// PolicyLayer defines a call layer of an escalation policy
// consists of users and calender
type PolicyLayer struct {
	ID                 string   `json:"id"`
	EscalationPolicyID string   `json:"escalation_policy_id"`
	UserIDs            []string `json:"user_ids"`
}

// EscalationPolicy defines an escalation policy
type EscalationPolicy struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ServiceID  string `json:"service_id"`
	LoopPolicy string `json:"loop_policy"`
}
