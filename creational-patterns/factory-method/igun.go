package main

type Igun interface {
	setName(name string)
	setDamage(damage int)
	getName() string
	getDamage() int
}