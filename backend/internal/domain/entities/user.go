package entities

import "github.com/google/uuid"

type UserRole string

const (
	RoleBuyer UserRole = "buyer"
	RoleDev   UserRole = "developer"
	RoleBoth  UserRole = "both"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
	Role  UserRole
}
