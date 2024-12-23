package model

// UserToCreate модель юзера при его создании
type UserToCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
}

// UserToLogin модель юзера при авторизации
type UserToLogin struct {
	Name     string
	Password string
}
