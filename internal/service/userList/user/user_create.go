package user

import (
	userlist "github.com/ozonmp/omp-bot/internal/model/userList"
)

func (u *DummyUserService) Create(users []userlist.User) error {
	u.Users = append(u.Users, users...)

	return nil
}
