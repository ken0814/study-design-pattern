package main

type Gun struct {
	Name   string
	Damage int
}

func (g *Gun) setName(name string) {
	g.Name = name
}

func (g *Gun) setDamage(damage int) {
	g.Damage = damage
}

func (g *Gun) getName() string {
	return g.Name
}

func (g *Gun) getDamage() int {
	return g.Damage
}
