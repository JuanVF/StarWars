package model

type Missile struct {
	price int64
}

// Funciones de la iGuns
func (m *Missile) Shot(player *Player, x int64, y int64) {

}

func (m *Missile) GetPrice() int64 {
	return m.price
}

func (m *Missile) SetPrice(price int64) {

}
