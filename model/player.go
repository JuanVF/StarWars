package model

type Player struct {
	Name           string
	Money          int64
	Steel          int64
	HasShield      bool
	isGraphVisible bool
	GunsList       map[Guns]int
	Matrix         [15][15]GameObject
}

func (p *Player) AddObject(object GameObject) {

}

func (p *Player) RemoveObject(object GameObject) {

}

func (p *Player) AddSteel(amount int64) {
	p.Steel += amount
}
