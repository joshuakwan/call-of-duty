package models

// Team defines a team of multiple users
type Team struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	UserIDs    []string `json:"user_ids"`
	ServiceIDs []string `json:"service_ids"`
}
