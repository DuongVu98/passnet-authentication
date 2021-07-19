package event

type UserRegisteredEvent struct {
	Uid            string
	Username       string
	Email          string
	FirstName      string
	LastName       string
	OrganizationId string
	DepartmentId   string
	CardId         string
	ProfileRole    string
}
