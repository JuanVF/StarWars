package controller

import (
	"fmt"

	"github.com/JuanVF/StarWars/utils"
	"github.com/gorilla/websocket"
)

var StartBroadcast = make(chan int)

// Asignamos el admin
func AssignAdmin(client *websocket.Conn) {
	fmt.Printf("Servidor: Asignando administrador...\n")

	pack := NetworkPackage{
		ID: utils.FIRST_USER,
	}

	SendTo(CreatePackage(&pack, client))
}

// Asignamos un jugador normal al flujo de juego
func AssignPlayer(client *websocket.Conn) {
	fmt.Printf("Servidor: Asignando nuevo jugador....\n")

	pack := NetworkPackage{
		ID: utils.LOGIN,
	}

	SendTo(CreatePackage(&pack, client))
}

// Creamos un nuevo jugador
func CreatePlayer(pack *NetworkPackage) {
	fmt.Println("\n\nServidor: Creando nuevo perfil de jugador...")

	fmt.Printf("Nombre recibido: %v\n", pack.Msg.Text)
	Clients[pack.To].Name = pack.Msg.Text

	fmt.Printf("Matriz recibida: \n %v\n", pack.Msg.Matrix)
	Clients[pack.To].GenerateMatrix(pack.Msg.Matrix)

	fmt.Println("Matriz generada:")
	Clients[pack.To].PrintMatrix()

	fmt.Println("Generando grafos:")
	Clients[pack.To].GenerateGraph()

	Clients[pack.To].GetGraphPoints()

	Clients[pack.To].Money = int64(pack.Msg.Numbers[0])

	fmt.Printf("\n\n")

	StartBroadcast <- 1
}

// Esta funcion va a terminar su ejecucion cuando se conecten todos los jugadore
func StartListener() {
	counter := 0

	fmt.Printf("Esperando %d jugadores \n", MAX_USERS)

	for {
		if counter >= MAX_USERS {
			fmt.Printf("Start Listener: El juego puede iniciar\n")
			pack := NetworkPackage{}

			pack.Msg = &Message{
				IdMessage: "STARTED",
			}

			SendToAll(pack.Msg)

			// Iniciamos todos los listeners de los players
			InitPlayersListener()

			return
		}

		amount := <-StartBroadcast

		counter += amount
		fmt.Printf("Start Listener: cantidad de jugadores: %d\n", counter)
	}
}
