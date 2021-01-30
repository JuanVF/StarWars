﻿using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class GameController : MonoBehaviour
{
    public static int lastTurn = -1;
    public static int lastCurrentTurn = -1;

    private Text playerTurn;
    private Text currentTurn;

    private GameObject chatContainer;

    void Start()
    {
        playerTurn = GameObject.Find("PlayerTurn").GetComponent<Text>();
        currentTurn = GameObject.Find("CurrentTurn").GetComponent<Text>();

        chatContainer = GameObject.Find("ChatContainer");

        notifyServer();
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

        if (Network.chatMessage != "")
        {
            AddChatMessage(Network.chatMessage);

            Network.chatMessage = "";
        }
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

    // Se notifica al server de que ya cambiamos de escena
    private void notifyServer()
    {
        Message message = new Message
        {
            idMessage = "STARTED"
        };

        Network.getInstance().SendMessage(message);
    }
}