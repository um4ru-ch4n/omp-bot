package user

import userlist "github.com/ozonmp/omp-bot/internal/model/userList"

type DummyUserService struct {
	Type  string
	Users []userlist.User
}

func NewDummyUserService() *DummyUserService {
	return &DummyUserService{
		Type:  "user",
		Users: make([]userlist.User, 0),
	}
}
