package dao

import "github.com/OmarTariq612/codersquare-go/types"

type UserDAO interface {
	CreateUser(user *types.User) error
	GetUserByEmail(email string) *types.User
	GetUserByUsername(username string) *types.User
	GetUserByID(id string) *types.User
	DeleteUser(id string) error // added
}
