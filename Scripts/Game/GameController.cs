using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class GameController : MonoBehaviour
{
    public static int lastTurn = -1;
    public static int lastCurrentTurn = -1;
    public static int lastcurrentMoney = 0;
    public static int lastcurrentSteel = -1;

    public static string lastEnemyName = "";

    private Text playerTurn;
    private Text currentTurn;
    private Text currentMoney;
    private Text currentSteel;
    private Text enemyName;

    private GameObject chatContainer;
    private GameObject attacksContainer;

    private Button chatButton;
    private Text chatText;

    public Sprite[] attacks;
    public string[] attacksName;

    void Start()
    {
        playerTurn = GameObject.Find("PlayerTurn").GetComponent<Text>();
        currentTurn = GameObject.Find("CurrentTurn").GetComponent<Text>();
        chatButton = GameObject.Find("ChatButton").GetComponent<Button>();
        chatText = GameObject.Find("ChatText").GetComponent<Text>();
        currentSteel = GameObject.Find("PlayerSteel").GetComponent<Text>();
        currentMoney = GameObject.Find("PlayerMoney").GetComponent<Text>();
        enemyName = GameObject.Find("EnemyName").GetComponent<Text>();
        attacksContainer = GameObject.Find("AttackContainer");

        chatContainer = GameObject.Find("ChatContainer");

        notifyServer();

        chatButton.onClick.AddListener(delegate {
            OnBtnChatClicked();
        });

        GenerateAttacks();
    }

    void Update()
    {
        if (lastTurn != Network.turn)
        {
            SetTurn();
            lastTurn = Network.turn;
        }

        if (lastCurrentTurn != Network.currentTurn)
        {
            SetCurrentTurn();
            lastCurrentTurn = Network.currentTurn;
        }

        if (lastcurrentMoney != Network.money)
        {
            SetCurrentMoney();
            lastcurrentMoney = (int) Network.money;
        }

        if (lastcurrentSteel != Network.steel)
        {
            SetCurrentSteel();
            lastcurrentSteel = (int) Network.steel;
        }

        if (lastEnemyName != Network.enemyName)
        {
            SetCurrentEnemy();

            lastEnemyName = Network.enemyName;
        }

        if (Network.chatMessage != "")
        {
            AddChatMessage(Network.chatMessage);

            Network.chatMessage = "";
        }
    }

    private void OnBtnChatClicked()
    {
        if (chatText.text == "")
        {
            AddChatMessage("Tu mensaje esta vacio...");

            return;
        }

        Network.getInstance().SendChatMessage(chatText.text);
    }

    private void AddChatMessage(string message)
    {
        GameObject reference = (GameObject)Instantiate(Resources.Load("ChatPrefab"), chatContainer.transform);

        reference.GetComponent<Text>().text = message;
    }

    // Asigna el turno del jugador en pantalla
    private void SetTurn()
    {
        playerTurn.text = "Turno del jugador: " + Network.turn;
    }

    // Asigna el turno actual en pantalla
    private void SetCurrentTurn()
    {
        currentTurn.text = "Turno actual: " + Network.currentTurn;
    }

    // Asigna el turno actual en pantalla
    private void SetCurrentMoney()
    {
        currentMoney.text = "Acero: " + Network.steel;
    }

    // Asigna el turno actual en pantalla
    private void SetCurrentSteel()
    {
        currentSteel.text = "Dinero: $" + Network.money;
    }

    // Asigna el turno actual en pantalla
    private void SetCurrentEnemy()
    {
        enemyName.text = "Enemigo: " + Network.enemyName;
    }

    // Se notifica al server de que ya cambiamos de escena
    private void notifyServer()
    {
        Message message = new Message
        {
            idMessage = "STARTED"
        };

        Network.getInstance().SendMessage(message);
    }

    // Se generan los ataques
    private void GenerateAttacks()
    {
        GameObject reference = (GameObject) Instantiate(Resources.Load("Attack"), attacksContainer.transform);
    }
}
