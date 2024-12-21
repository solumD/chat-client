package model

type UserToCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
}

type UserToLogin struct {
	Name     string
	Password string
}
