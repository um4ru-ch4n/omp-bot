package user

import "fmt"

func (u *DummyUserService) Remove(userID uint64) (bool, error) {
	if len(u.Users) < int(userID) {
		return false, fmt.Errorf("no user with such index")
	}

	if userID == uint64(len(u.Users)) {
		u.Users = u.Users[:userID-1]
		return true, nil
	}

	u.Users = append(u.Users[:userID-1], u.Users[userID:]...)

	return true, nil
}
