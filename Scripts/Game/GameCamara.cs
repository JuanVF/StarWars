using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameCamara : MonoBehaviour
{
    private bool alradyAttacked = false;

    public static float[] toAttack;

    private void Start()
    {
    }

    void Update()
    {
        if (!Input.GetMouseButtonDown(0))
            return;

        RaycastHit hit;
        Ray ray = Camera.main.ScreenPointToRay(Input.mousePosition);

        if (Physics.Raycast(ray, out hit, 100.0f))
        {
            if (hit.transform != null)
            {
                Clicked(hit.transform.gameObject);
            }
        }
    }

    private void Clicked(GameObject gameObject)
    {
        if (gameObject.CompareTag("tiles"))
        {
            int col = gameObject.GetComponent<Point>().getPoint()[0];
            int row = gameObject.GetComponent<Point>().getPoint()[1];

            Debug.Log("Atacando: " + col + ", " + row);

            Attack(col, row);
        }
    }

    private void Attack(int col, int row)
    {
        if (GameController.attackSelected == -1) return;

        toAttack = append(toAttack, col);
        toAttack = append(toAttack, row);

        int amount = toAttack.Length / 2;
        int total = -1;

        switch (GameController.attackSelected)
        {
            case 0:
                total = 1;
                break;
            case 1:
                total = 1;
                break;
            case 2:
                total = 3;
                break;
            case 3:
                total = 10;
                break;
            default:
                GameController.attackSelected = -1;
                break;
        }

        // Si ya selecciono todos sus ataques los enviamos
        if (amount == total)
        {
            Message message = new Message
            {
                idMessage = "ATTACK",
                number = GameController.attackSelected,
                numbers = toAttack,
                text = Network.enemyName
            };

            Network.getInstance().SendMessage(message);

            Network.chatMessage = "Atacando...";
            GameController.attackSelected = -1;
        }
    }

    private float[] append(float[] arr, float ele)
    {
        if (toAttack == null) return new float[] { ele };

        float[] attack = new float[arr.Length + 1];

        for (int i = 0; i < arr.Length; i++)
        {
            attack[i] = arr[i];
        }

        attack[attack.Length - 1] = ele;

        return attack;
    }
}
