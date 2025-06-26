package command

type CreateUserCommand struct {
	Username  string
	Password  string
	Role      string
	CreatedBy string
	UpdatedBy string
}

func NewCreateUserCommand(username, password, role, createdBy, updatedBy string) *CreateUserCommand {
	return &CreateUserCommand{
		Username:  username,
		Password:  password,
		Role:      role,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}
}
