using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.SceneManagement;

public class UIController : MonoBehaviour
{
    private Button playButton;
    private Text playerName;
    private Dropdown dpAmountUsers;

    private bool alreadyLogged = false;

    void Start()
    {
        playButton = GameObject.Find("ButtonPlay").GetComponent<Button>();
        playerName = GameObject.Find("TextPlayerName").GetComponent<Text>();
        dpAmountUsers = GameObject.Find("dpAdmin").GetComponent<Dropdown>();

        dpAmountUsers.onValueChanged.AddListener(delegate {
            Network.amount_users = dpAmountUsers.value + 2;
        });

        playButton.onClick.AddListener(delegate {
            OnPlayButtonClicked();
        });
    }

    // Nos conectamos con el server...
    void OnPlayButtonClicked()
    {
        if (alreadyLogged) return; // Evitamos que se loggee dos veces

        Network.matrix = ComponentGenerator.PlayerMatrix;
        Network.name = playerName.text;
        Network.getInstance();

        alreadyLogged = true;
    }

    // Update is called once per frame
    void Update()
    {
        if (Network.canStart)
        {
            SceneManager.LoadScene(1);

            Network.canStart = false;
        }
    }
}
