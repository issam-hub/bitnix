package request

import (
	"regexp"
	"slices"
	"strings"
	"user-service/internal/application/command"
	"user-service/internal/validator"
)

var (
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zAZ0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	passwordRegex = regexp.MustCompile(`^[A-Za-z\d@$!%*?#&]{8,}$`)
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (req *RegisterUserRequest) Validate() validator.ValidationErrors {
	var errors validator.ValidationErrors

	if strings.TrimSpace(req.Username) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "username",
			Message: "username is required and cannot be empty",
		})
	}

	if strings.TrimSpace(req.Email) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "email",
			Message: "email is required",
		})
	} else if !emailRegex.MatchString(req.Email) {
		errors = append(errors, validator.ValidationError{
			Field:   "email",
			Message: "email is invalid",
		})
	}

	if strings.TrimSpace(req.Password) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "password",
			Message: "password is required",
		})
	} else if !passwordRegex.MatchString(req.Password) {
		errors = append(errors, validator.ValidationError{
			Field:   "password",
			Message: "password is invalid",
		})
	}

	allowedRoles := []string{"developer", "buyer", "admin"}

	if !slices.Contains(allowedRoles, req.Role) {
		errors = append(errors, validator.ValidationError{
			Field:   "role",
			Message: "wrong role",
		})
	}

	return errors
}

func (req *RegisterUserRequest) ToRegisterUserCommand() (*command.RegisterUserCommand, error) {
	validationErrors := req.Validate()
	if validationErrors.HasErrors() {
		return nil, validationErrors
	}

	return &command.RegisterUserCommand{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}, nil
}
