package models

var testInputUsers = []struct {
	name  string
	email string
}{
	{"abe", "abe@example.com"},
	{"bobby", "bobby@example.com"},
	{"carl", "carl@example.com"},
	{"dick", "dick@example.com"},
	{"erick", "erick@example.com"},
	{"fredric", "fredric@example.com"},
	{"gabriel", "gabriel@example.com"},
	{"hana", "hana@example.com"},
	{"ibe", "ibe@example.com"},
	{"jojo", "jojo@example.com"},
}

var testInputServices = []string{"Service A", "Service B", "Service C", "Service D"}

func createTestUsers() []*User {
	var users []*User
	for _, input := range testInputUsers {
		users = append(users, CreateUser(input.name, input.email))
	}

	return users
}

func createTestServices() []*Service {
	var services []*Service
	for _, input := range testInputServices {
		services = append(services, CreateService(input))
	}
	return services
}

func createTestTeam() *Team {
	teamName := "Team A"
	teamDescription := "A Testing Team"
	return &Team{Name: teamName, Description: teamDescription}
}
