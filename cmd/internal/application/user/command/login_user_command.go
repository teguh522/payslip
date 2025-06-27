package command

type LoginUserCommand struct {
	Username string
	Password string
}

func NewLoginUserCommand(username, password string) *LoginUserCommand {
	return &LoginUserCommand{
		Username: username,
		Password: password,
	}
}
