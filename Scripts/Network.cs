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

    public static string name { get; set; }
    public static float money = 4000;
    public static float steel = 0;

    public static float[,] matrix { get; set; }
    public static int amount_users = 2;

    public static float[] graphConnections = null;
    public static bool canStart = false;

    public static int turn = -1;
    public static int currentTurn = -1;
    public static bool isTurn = false;

    public static string enemyName = "";
    public static int enemyTurn = -1;
    public static float[,] enemyMatrix = null;
    public static float[] enemyGraph = null;

    public static string chatMessage = "";

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
        Message message;

        switch (messageAvailable.idMessage)
        {
            case "ADMIN":
                Debug.Log("Eres administrador...");

                message = new Message
                {
                    idMessage = "ADMIN",
                    number = amount_users
                };

                SendMessage(message);

                break;
            case "REQUESTDATA":
                Debug.Log("Enviando datos...");

                message = new Message
                {
                    idMessage = "REQUESTDATA",
                    text = name,
                    matrix = matrix,
                    numbers = new float[] {money}
                };

                SendMessage(message);

                break;
            case "STARTED":
                Debug.Log("Cambiando de escena");

                canStart = true;

                break;
            case "RECEIVEMATRIX":
                Debug.Log("Matriz recibida...");

                matrix = messageAvailable.matrix;

                Debug.Log("Tamano: " + matrix.Length);

                message = new Message
                {
                    idMessage = "SEND_MATRIX"
                };

                SendMessage(message);

                break;
            case "RECEIVEPOINTS":
                Debug.Log("Recibiendo puntos de grafos...");
                Debug.Log("Puntos recibidos: " + messageAvailable.numbers.Length);
                string points = "[ ";

                graphConnections = messageAvailable.numbers;

                //  Imprimos los puntos recibidos...
                for (int i = 0; i < messageAvailable.numbers.Length; i++)
                {
                    points += " [ " + graphConnections[i] + " ] ";
                }

                points += " ]";

                Debug.Log(points);

                message = new Message
                {
                    idMessage = "PLAYER_INIT"
                };

                SendMessage(message);

                break;
            case "PLAYER_INIT":
                money = messageAvailable.numbers[0];
                steel = messageAvailable.numbers[1];

                message = new Message
                {
                    idMessage = "ENEMY_INIT"
                };

                SendMessage(message);

                break;
            case "ENEMY_INIT":
                enemyTurn = (int) messageAvailable.number;
                enemyName = messageAvailable.text;
                enemyMatrix = messageAvailable.matrix;

                Debug.Log("Datos del enemigo: ");
                Debug.Log("Nombre: " + enemyName);
                Debug.Log("ID: " + enemyTurn);

                message = new Message
                {
                    idMessage = "REQUEST_TURN"
                };

                SendMessage(message);

                break;
            case "GRAPH_VISIBLE":
                if (enemyName != messageAvailable.text) return;

                Debug.Log("Un enemigo tiene el grafo visible");

                enemyName = messageAvailable.text;
                enemyMatrix = messageAvailable.matrix;
                enemyGraph = messageAvailable.numbers;

                break;
            case "ASSIGN_TURN":
                Debug.Log("Recibiendo turno: " + messageAvailable.number);

                turn = (int) messageAvailable.number;

                break;
            case "TURN":
                Debug.Log("Es el turno: " + messageAvailable.number);

                // Asignamos el turno actual y verificamos si es el turno
                currentTurn = (int) messageAvailable.number;
                isTurn = turn == currentTurn;

                if (isTurn)
                {
                    chatMessage = "Es tu turno!";
                }

                break;
            case "CHAT":
                Debug.Log("Server: Mensaje = " + messageAvailable.text);

                chatMessage = messageAvailable.text;

                break;
            case "USER_INFO":
                Debug.Log("Datos del usuario actualizados...");

                steel = messageAvailable.numbers[0];
                money = messageAvailable.numbers[1];

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
            username = name,
            text = msg,
            idMessage = "CHAT"
        };

        SendMessage(message);
    }
}
