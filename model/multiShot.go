package model

type MultiShot struct {
	price int64
}

// Funciones de la iGuns
func (m *MultiShot) Shot(player *Player, x int64, y int64) {

}

func (m *MultiShot) GetPrice() int64 {
	return m.price
}

func (m *MultiShot) SetPrice(price int64) {

}
