package main

import (
	"fmt"
)

func main() {
	fmt.Println("Prototype Pattern 原形模式")
	original := &User{Name: "ken", Age: 26, Email: "ken@example.com"}
	clone := original.Clone()
	clone.setName("joyce")
	clone.setEmail("joyce@example.com")
	fmt.Printf("Original point:%p\nOriginal value: %+v\n", original, original)
	fmt.Printf("Clone point:%p\nClone value: %+v\n", clone, clone)
}
