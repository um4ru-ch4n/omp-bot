package user

func (u *DummyUserService) Count() uint64 {
	return uint64(len(u.Users))
}
