package main

type Musket struct {
	Gun
}

func NewMusket() Igun {
	return &Musket{
		Gun{
			Name:   "Musket",
			Damage: 25,
		},
	}
}
