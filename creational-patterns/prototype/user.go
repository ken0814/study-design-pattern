package main

type User struct {
	Name  string
	Age   int
	Email string
}

// 淺拷貝
func (u *User) Clone() *User {
	return &User{
		Name:  u.Name,
		Age:   u.Age,
		Email: u.Email,
	}
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) setEmail(email string) {
	u.Email = email
}