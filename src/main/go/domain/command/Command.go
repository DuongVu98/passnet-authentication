package command

type (
	BaseCommand interface {
	}
	RegisterCommand struct {
		Username    string
		Email       string
		Password    string
		FirstName   string
		LastName    string
		ProfileRole string
		BaseCommand
	}
)
