package model

type Temple struct {
	owner         *Player
	price         int64
	componentType int64
}

// Funciones de la interfaz iGameObject
func (t *Temple) OnStart() {

}

func (t *Temple) Run() {

}

func (t *Temple) OnHit(player *Player) {

}

// Funciones de la interfaz iFactory
func (t *Temple) DoAction() {

}

// Funciones de la interfaz iComponent
func (t *Temple) SetPrice(int64) {

}

func (t *Temple) GetPrice() int64 {
	return t.price
}

func (t *Temple) SetPlayer(owner *Player) {

}

func (t *Temple) GetPlayer() *Player {
	return t.owner
}

func (t *Temple) GetType() int64 {
	return t.componentType
}

func (t *Temple) SetType(ComponentType int64) {

}
