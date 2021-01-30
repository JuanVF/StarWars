using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using WebSocketSharp;
using Newtonsoft.Json;
using System.Threading;

public class Network
{
    private static Network network;

    private WebSocket ws;

    private static Message messageAvailable;
    private static bool isConnected = false;

    public string name { get; set; }
    public float[,] matrix { get; set; }
    public int amount_users { get; set; }

    public static int[] graphConnections = null;

    // Start is called before the first frame update
    void Start()
    {
        ws = new WebSocket("ws://localhost:3000/api/login");

        ws.OnMessage += (sender, e) =>
        {
            Message message = JsonConvert.DeserializeObject<Message>(e.Data);

            messageAvailable = message;
        };

        ws.OnError += (sender, e) =>
        {
            isConnected = false;
            Debug.LogError("WS Error: " + e.Message);

            ws.Close();
        };

        ws.Connect();

        isConnected = true;

        Thread listener = new Thread(new ThreadStart(Listener));

        listener.Start();
    }

    // Este va a ser el listener del juego
    private void Listener()
    {
        while (isConnected)
        {
            if (messageAvailable != null)
            {
                Debug.Log(messageAvailable.idMessage + " : " + messageAvailable.text);

                React();
            }
        }
    }

    private void React()
    {
        switch (messageAvailable.idMessage)
        {
            case "ADMIN":
                Debug.Log("Eres administrador...");

                Message admMsg = new Message
                {
                    idMessage = "ADMIN",
                    number = 2
                };

                SendMessage(admMsg);

                break;
            case "REQUESTDATA":
                Debug.Log("Enviando datos...");

                Message dataMsg = new Message
                {
                    idMessage = "REQUESTDATA",
                    text = name,
                    matrix = matrix
                };

                SendMessage(dataMsg);

                break;
            default:
                Debug.Log("Mensaje recibido no soportado...");
                break;
        }

        messageAvailable = null;
    }

    public static Network getInstance()
    {
        if (network == null)
        {
            network = new Network();
            network.Start();
        }

        return network;
    }

    public void SendMessage(Message message)
    {
        string json = JsonConvert.SerializeObject(message);

        ws.Send(json);
    }

    public void SendChatMessage(string msg)
    {
        Message message = new Message
        {
            username = "Jugador X", // Lo cambiaremos cuando agregue un login funcional
            text = msg,
            idMessage = "CHAT"
        };

        SendMessage(message);
    }
}
