package structure

type Message struct {
	User	string	`json:"user"`
	Info	string	`json:"message,omitempty"`
}