package models

type User struct {
	ID       int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"size:255;not null" json:"name"`
	Email    string  `gorm:"size:255;unique;not null" json:"email"`
	Password string  `gorm:"size:255;not null" json:"password"`
	Balance  float64 `gorm:"type:decimal(10,2);default:0" json:"balance"`
}

type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}
