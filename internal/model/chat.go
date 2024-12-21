package model

type Chat struct {
	Name      string
	Usernames []string
}

type Message struct {
	ChatID int64
	From   string
	Text   string
}
