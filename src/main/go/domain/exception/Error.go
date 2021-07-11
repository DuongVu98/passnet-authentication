package exception

type InvalidCommandException struct {
	error
}

func (ex InvalidCommandException) Error() string {
	return "Invalid input command"
}
