package entity

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	// Password always keeps the hashes password
	Password string
}
