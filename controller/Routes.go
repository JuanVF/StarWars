package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JuanVF/StarWars/model"
	"github.com/JuanVF/StarWars/utils"
	"github.com/gorilla/websocket"
)

var SERVER_PORT = ":3000"
var LoginUpgrader = websocket.Upgrader{}

var Clients = make(map[*websocket.Conn]*model.Player)
var Broadcast = make(chan *NetworkPackage)

var MAX_USERS = -1

func StartServer() error {
	fmt.Println("Servidor iniciado en: " + SERVER_PORT)

	// Servimos la ruta para loggearse
	http.HandleFunc("/api/login", HandlerLogin)

	// Abrimos el broadcast
	go BroadcastHandler()

	if err := http.ListenAndServe(SERVER_PORT, nil); err != nil {
		return fmt.Errorf("Server error: %s", err)
	}

	return nil
}

// Handler de la ruta login
func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ws, err := LoginUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)

		return
	}

	defer ws.Close()

	if MAX_USERS != -1 && len(Clients) >= MAX_USERS {
		return
	}

	Clients[ws] = &model.Player{}

	log.Printf(utils.ReadUserIP(r) + ", Se ha conectado...")
	log.Printf("Cantidad de usuarios: #%d", len(Clients))

	// Si solo hay un cliente conectado el va a ser admin
	if len(Clients) == 1 {
		AssignAdmin(ws)
	} else {
		AssignPlayer(ws)
	}

	PlayerListener(ws)
}

// Listener para escuchar los mensajes del jugador
func PlayerListener(ws *websocket.Conn) {
	for {
		var message Message

		err := ws.ReadJSON(&message)

		if err != nil {
			log.Printf("Player error: %v", err)

			ws.Close()
			delete(Clients, ws)

			return
		}

		fmt.Printf("Player sended: %v\n", message.IdMessage)

		Broadcast <- CreatePackageMsg(&message, ws)
	}

}

// Aqui se distribuyen los mensajes
func BroadcastHandler() {
	for {
		msg := <-Broadcast

		if msg == nil {
			continue
		}

		pack := CreatePackage(msg, msg.To)

		if !pack.Response || pack.Msg == nil {
			continue
		}

		if pack.To == nil {
			SendToAll(pack.Msg)
		} else {
			SendTo(pack)
		}
	}
}
