package domain

type UserRepository interface {
	FindUserByEmail(email string) (*User, error)
}
