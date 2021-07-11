package aggregate

type (
	User struct {
		Uid        UserId
		Username   Username
		Email      Email
		Profile    UserProfile
		Credential UserCredential
	}

	UserId struct {
		Value string
	}
	Username struct {
		Value string
	}
	Email struct {
		Value string
	}
	UserProfile struct {
		FirstName string
		LastName  string
	}
	UserCredential struct {
		Password string
	}
)
