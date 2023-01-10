package user

type Store interface {
	Create(u *User) error
}
