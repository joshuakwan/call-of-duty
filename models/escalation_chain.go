package models

// EscalationChain loop policy

// ChainLayer defines a call layer of an escalation chain
// consists of users and calender
type ChainLayer struct {
	ID                string `json:"id"`
	EscalationChainID string `json:"escalation_chain_id"`
	CalendarEntries   []CalendarEntry
	PreDelaySeconds   int         // seconds to delay before processing the calendar entries
	PostDelaySeconds  int         // seconds to delay after going to the next layers
	PreviousLayer     *ChainLayer // pointer to previous layer
	NextLayer         *ChainLayer // pointer to next layer
}

// LoopPolicy defines the policy to loop an escalation chain
type LoopPolicy int

// Enumeration for LoopPolicy
const (
	LOOP LoopPolicy = 1 + iota
	ONCE
)

// EscalationChain defines an escalation chain
type EscalationChain struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ServiceID string `json:"service_id"`
	Policy    LoopPolicy
	Head      *ChainLayer // entry pointer of the chain
}
