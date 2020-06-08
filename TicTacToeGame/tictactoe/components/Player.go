package components

type Player struct {
	Name string
	Mark string
}

func NewPlayer(name, mark string) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}
}
