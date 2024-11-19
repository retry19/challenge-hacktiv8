package database

type Username string

type Email string

type Password string

type User struct {
	BaseModel
	Username Username `json:"username" gorm:"column:username;not null;uniqueIndex" validate:"required"`
	Email    Email    `json:"email" gorm:"column:email;not null;uniqueIndex" validate:"required,email"`
	Password Password `json:"password,omitempty" gorm:"column:password;not null" validate:"required,min=6"`
	Age      uint     `json:"age" gorm:"column:age;not null" validate:"required,min=8"`
}

type LoginPayload struct {
	Username Username `json:"username" validate:"required_without=Email"`
	Email    Email    `json:"email" validate:"required_without=Username,email"`
	Password Password `json:"password" validate:"required,min=6"`
}
