using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UIController : MonoBehaviour
{
    private Button playButton;
    private Text playerName;

    void Start()
    {
        playButton = GameObject.Find("ButtonPlay").GetComponent<Button>();
        playerName = GameObject.Find("TextPlayerName").GetComponent<Text>();

        playButton.onClick.AddListener(delegate {
            OnPlayButtonClicked();
        });
    }

    void OnPlayButtonClicked()
    {
        Network.getInstance().matrix = ComponentGenerator.PlayerMatrix;
        Network.getInstance().name = playerName.text;
    }

    // Update is called once per frame
    void Update()
    {
        if (Network.graphConnections != null)
        {
            conectarLineas();

            Network.graphConnections = null;
        }
    }

    private void conectarLineas()
    {
        // TODO SU CODIGO
    }
}
