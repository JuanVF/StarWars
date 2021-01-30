package controller

import (
	"github.com/gorilla/websocket"
)

var TurnBroadcast = make(chan bool)
var TurnList = make(map[int]*websocket.Conn)
var MAX_TURN = -1
var current_turn = 0

var AlreadySettingTurn = false
var SettingTurnBroadcast = make(chan bool)

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

	MAX_TURN = turn

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
