package main

import (
	"fmt"
)

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

// 深拷貝
func (u *User) DeepClone() *User {
	clone := *u
	return &clone
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) setEmail(email string) {
	u.Email = email
}

func main() {
	fmt.Println("Prototype Pattern 原形模式")
	original := &User{Name: "ken", Age: 26, Email: "ken@example.com"}
	clone := original.Clone()
	clone.setName("joyce")
	clone.setEmail("joyce@example.com")
	fmt.Printf("Original point:%p\nOriginal value: %+v\n", original, original)
	fmt.Printf("Clone point:%p\nClone value: %+v\n", clone, clone)
}
