package main

type Ak47 struct {
	Gun
}

func NewAk47() Igun {
	return &Ak47{
		Gun{
			Name:   "AK47",
			Damage: 30,
		},
	}
}
