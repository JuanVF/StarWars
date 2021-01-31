package controller

import (
	"fmt"

	"github.com/JuanVF/StarWars/model"
	"github.com/JuanVF/StarWars/utils"
	"github.com/gorilla/websocket"
)

var TurnBroadcast = make(chan bool)
var TurnList = make(map[int]*websocket.Conn)
var TurnListP = make(map[*websocket.Conn]int)
var MAX_TURN = -1
var current_turn = 0

var AlreadySettingTurn = false
var SettingTurnBroadcast = make(chan bool)

var ChatBroadcast = make(chan string)

// Espera que todos los jugadores pregunten por su turno para asignarlo
func WaitToAssign() {
	total := 0

	for {
		if total == len(Clients) {
			AssignTurns()

			break
		}

		_ = <-SettingTurnBroadcast

		total++
	}
}

// Asigna los turnos a cada jugador y se lo envia
func AssignTurns() {
	var turn int = 0

	for client := range Clients {
		TurnList[turn] = client
		TurnListP[client] = turn

		// Le enviamos el turno
		pack := NetworkPackage{
			To: client,
			Msg: &Message{
				IdMessage: "ASSIGN_TURN",
				Number:    float64(turn),
			},
		}

		SendTo(&pack)

		turn++
	}

	MAX_TURN = turn - 1

	fmt.Printf("Maximo de turnos: %v\n", MAX_TURN-1)

	// Le enviamos el turno
	pack := NetworkPackage{
		To: TurnList[0],
		Msg: &Message{
			IdMessage: "TURN",
			Number:    float64(current_turn),
		},
	}

	SendToAll(pack.Msg)

	// Activamos el listener de turnos
	go TurnListener()
}

// Listener de los turnos
//se ejecuta cuando todos los jugadores han enviado el started
func TurnListener() {
	for {
		_ = <-TurnBroadcast

		current_turn++

		if current_turn > MAX_TURN {
			current_turn = 0
		}

		// Le enviamos el turno
		pack := NetworkPackage{
			To: TurnList[current_turn],
			Msg: &Message{
				IdMessage: "TURN",
				Number:    float64(current_turn),
			},
		}

		SendToAll(pack.Msg)
	}
}

// Valida si es el turno de un jugador, si no es el turno se le notifica
func isPlayerTurn(client *websocket.Conn) bool {
	if TurnListP[client] == current_turn {
		return true
	}

	SendToPlayerChat("Server: No es su turno...", client)

	return false
}

// Dado un nombre lo buscamos en los clientes y lo retornamos
func GetPlayer(name string) *model.Player {
	for client := range Clients {
		if name == Clients[client].Name {
			return Clients[client]
		}
	}

	return nil
}

// Retornamos el socket de un jugador
func GetSocket(name string) *websocket.Conn {
	for client := range Clients {
		if Clients[client].Name != Clients[client].Name {
			return client
		}
	}

	return nil
}

// Buscamos el primer enemigo que no sea nosotros
func GetEnemieName(player *websocket.Conn) string {
	for client := range Clients {
		if Clients[client].Name != Clients[player].Name {
			return Clients[client].Name
		}
	}

	return ""
}

// Esto deja ejecutando cada uno de los listener de los jugadores
func InitPlayersListener() {
	for client := range Clients {
		Clients[client].FactoryChan = make(chan string)

		go func(client *websocket.Conn) {
			for {
				msg := <-Clients[client].FactoryChan

				if msg == "SERVER_GRAFO_VISIBLE" {
					pack := &Message{
						IdMessage: "GRAPH_VISIBLE",
						Matrix:    Clients[client].GetMatrix(),
						Numbers:   Clients[client].GetGraphPoints(),
						Text:      Clients[client].Name,
					}

					SendToAll(pack)

					netw := &NetworkPackage{
						Msg: &Message{
							IdMessage: "USER_INFO",
							Numbers:   []float64{float64(Clients[client].Steel), float64(Clients[client].Money)},
						},
						To: client,
					}

					SendTo(netw)
				} else if msg == "SERVER_STEEL" {
					fmt.Printf("Acero del jugador: %d\n", Clients[client].Steel)

					netw := &NetworkPackage{
						Msg: &Message{
							IdMessage: "USER_INFO",
							Numbers:   []float64{float64(Clients[client].Steel), float64(Clients[client].Money)},
						},
						To: client,
					}

					SendTo(netw)
				} else {
					SendToPlayerChat(msg, client)
				}
			}
		}(client)
	}
}

// Este listener recibe mensajes del chat y los distribuye a todos los jugadores
func ChatListener() {
	for {
		msg := <-ChatBroadcast

		SendToChat(msg)
	}
}

// Permite a un jugador realizar un ataque
func DoAttack(pack *NetworkPackage) {
	if !isPlayerTurn(pack.To) {
		return
	}

	var attackType model.Guns
	var attackTo *model.Player = GetPlayer(pack.Msg.Text)

	// El jugador a atacar no existe...
	if attackTo == nil {
		SendToPlayerChat("Server: ese jugador no existe...", pack.To)

		return
	}

	//gunType := ""

	// Determinamos el tipo de arma
	switch pack.Msg.Number {
	case utils.MISSILE:
		attackType = &model.Missile{}
		//gunType = "missile"
	case utils.MULTISHOT:
		attackType = &model.MultiShot{}
		//gunType = "multishot"
	case utils.BOMB:
		attackType = &model.Bomb{}
		//gunType = "bomb"
	case utils.COMBOSHOT:
		attackType = &model.ComboShot{}
		//gunType = "comboshot"
	default:
		return
	}

	// Verificamos que el jugador posea el arma con la que va a atacar
	/*gunList := Clients[pack.To].GunsList

	if gunList[gunType] == 0 {
		SendToPlayerChat("Server: No tienes ese tipo de arma...", pack.To)

		return
	}*/

	if attackTo.HasShield != 0 {
		SendToPlayerChat("Server: El jugador tiene un escudo activo...", pack.To)
		attackTo.HasShield--
	}

	// Atacamos al jugador y enviamos el resultado del ataque al jugador
	SendToChat("Server: " + attackType.Shot(Clients[pack.To], attackTo, pack.Msg.Numbers))

	TurnBroadcast <- true
}

// Permite a un jugador comprar armas
func BuyArmory(pack *NetworkPackage) {
	var gun model.Guns

	fmt.Printf("Mensaje: %v", pack.Msg)

	arType := pack.Msg.Texts[0]

	switch arType {
	case "missil":
		gun = &model.Missile{}
	case "bomb":
		gun = &model.Bomb{}
	case "comboshot":
		gun = &model.ComboShot{}
	case "multishot":
		gun = &model.MultiShot{}
	default:
		return
	}

	if Clients[pack.To].Steel < gun.GetPrice() {
		SendToPlayerChat("ARMERIA: No tienes suficiente acero...", pack.To)
		return
	}

	Clients[pack.To].Steel -= gun.GetPrice()

	// Evitamos un hashmap nulo
	if Clients[pack.To].GunsList == nil {
		Clients[pack.To].GunsList = make(map[string]int)
	}

	Clients[pack.To].GunsList[arType]++
	SendToPlayerChat("Has comprado el arma!", pack.To)
}
