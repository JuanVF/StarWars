using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using System.Text.RegularExpressions;

public class GameController : MonoBehaviour
{
    public static int lastTurn = -1;
    public static int lastCurrentTurn = -1;
    public static int lastcurrentMoney = 0;
    public static int lastcurrentSteel = -1;

    public static string lastEnemyName = "";

    public static int attackSelected = -1;

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
    public string[] attacksPrice;

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
    private static string[] ParseMessage(string message)
    {

        var match = Regex.Matches(message, @"(\w)+");
        string[] msg = new string[match.Count];

        for (int i = 1; i < match.Count; i++)
        {
            msg[i - 1] = match[i].Value.ToLower();
        }
        return msg;

    }

    private static string joinStr(string[] message)
    {
        string builded = "";

        for (int i = 0; i < message.Length; i++)
        {
            builded += message[i] + " ";
        }

        return builded;
    }

    public void messageSender(string arr)
    {
        var match = Regex.Matches(arr, @"(\w)+");
        string key = match[0].Value;
        string[] msg = ParseMessage(arr);

        Message message = null;

        switch (key.ToUpper())
        {
            case "CHAT":
                Network.getInstance().SendChatMessage(joinStr(msg));
                break;

            case "ARMORY":
                message = new Message
                {
                    idMessage = "ARMORY",
                    texts = msg
                };

                Network.getInstance().SendMessage(message);
                break;

            case "MARKET":
                message = new Message
                {
                    idMessage = "MARKET",
                    texts = msg
                };

                Network.getInstance().SendMessage(message);
                break;

            case "PLAYER_INFO":
                message = new Message
                {
                    idMessage = "PLAYER_INFO",
                    texts = msg
                };

                Network.getInstance().SendMessage(message);
                break;
        }


    }

    private void OnBtnChatClicked()
    {
        if (chatText.text == "")
        {
            AddChatMessage("Tu mensaje esta vacio...");

            return;
        }

        messageSender(chatText.text);
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
        currentMoney.text = "Dinero: $" + Network.money;
    }

    // Asigna el turno actual en pantalla
    private void SetCurrentSteel()
    {
        currentSteel.text = "Acero: " + Network.steel;
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
        GameObject reference = (GameObject) Instantiate(Resources.Load("Attack"));

        for (int i = 0; i < attacksName.Length; i++)
        {
            GameObject attack = Instantiate(reference, attacksContainer.transform);
            Image image = attack.transform.Find("AttackImage").GetComponent<Image>();
            Text name = attack.transform.Find("AttackName").GetComponent<Text>();
            Text price = attack.transform.Find("AttackPrice").GetComponent<Text>();
            Button button = attack.transform.Find("AttackButton").GetComponent<Button>();

            image.sprite = attacks[i];
            name.text = attacksName[i];
            price.text = "Precio: " + attacksPrice[i] + "Kg";

            int tmpI = i;

            button.onClick.AddListener(delegate {
                if (attackSelected != tmpI)
                {
                    GameCamara.toAttack = null;
                }

                attackSelected = tmpI;
            });
        }
    }
}
