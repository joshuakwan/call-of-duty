package models

import "../utils"

// Team defines a team of multiple users
type Team struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UserIDs     []string `json:"user_ids"`
	ServiceIDs  []string `json:"service_ids"`
}

// CreateTeam creates a new team
func CreateTeam(name, description string) *Team {
	return &Team{ID: utils.GenerateUUID(), Name: name, Description: description}
}

// AddNewService adds a service to the team
func (team *Team) AddNewService(service *Service) {
	service.TeamID = team.ID
	team.ServiceIDs = append(team.ServiceIDs, service.ID)
}

// RemoveService removes a service from the team
func (team *Team) RemoveService(service *Service) string {
	for i, serviceID := range team.ServiceIDs {
		if serviceID == service.ID {
			service.TeamID = ""
			team.ServiceIDs = utils.RemoveElement(team.ServiceIDs, i)
		}
	}
	return service.ID
}

// AddNewUser adds a user to the team
func (team *Team) AddNewUser(user *User) {
	team.UserIDs = append(team.UserIDs, user.ID)
	user.TeamIDs = append(user.TeamIDs, team.ID)
}

// RemoveUser removes a user from the team
func (team *Team) RemoveUser(user *User) string {
	for i, userID := range team.UserIDs {
		if userID == user.ID {
			user.removeFromTeam(team.ID)
			team.UserIDs = utils.RemoveElement(team.UserIDs, i)
			break
		}
	}
	return user.ID
}
