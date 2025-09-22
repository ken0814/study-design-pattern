package main

import (
	"fmt"
)

func main() {
	fmt.Println("Factory Method Pattern 工廠方法模式")
	ak47 := NewAk47()
	musket := NewMusket()
	fmt.Printf("Gun: %s, Damage: %d\n", ak47.getName(), ak47.getDamage())
	fmt.Printf("Gun: %s, Damage: %d\n", musket.getName(), musket.getDamage())
}