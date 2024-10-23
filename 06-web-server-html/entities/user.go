package entities

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Address  string
	Phone    string
}

var Users = []*User{
	NewUser(1, "Reza", "reza@gmail.com", "123456", "jakarta", "0812345678"),
	NewUser(2, "James", "james@gmail.com", "123456", "surabaya", "0812345678"),
	NewUser(3, "Foo", "foo@gmail.com", "123456", "bandung", "0812345678"),
}

func NewUser(id int, name string, email string, password string, address string, phone string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
		Phone:    phone,
	}
}

func (u *User) FindByEmail(email string) *User {
	for _, user := range Users {
		if user.Email == email {
			return user
		}
	}
	return nil
}
