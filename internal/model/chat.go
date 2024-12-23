package model

// Chat модель чата
type Chat struct {
	Name      string
	Usernames []string
}

// Message модель сообщения
type Message struct {
	ChatID int64
	From   string
	Text   string
}
