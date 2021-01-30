using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.EventSystems;

public class ComponentGenerator : MonoBehaviour
{
    public Sprite[] ComponentImages;
    public string[] ComponentNames;
    public int[] ComponentPrice;
    public int[,] ComponentSize;

    public GameObject[,] PlayerComponents;

    public int cols = 15;
    public int rows = 15;

    public float tileSize = 1;

    private Transform ComponentsContainer;

    public static int selectedComponent = 0;
    public static float[,] PlayerMatrix = new float[15, 15];
    public static bool[,] PlayerSet = new bool[15, 15];

    public static GameObject[,] matrix = new GameObject[15, 15];

    void Start()
    {
        ComponentSize = new int[,] { { 1, 2 }, { 1, 1 }, { 1, 2 }, { 1, 2 }, { 2, 2 }, { 1, 2 } };

        PlayerComponents = new GameObject[cols, rows];

        ComponentsContainer = GameObject.Find("ComponentsContainer").GetComponent<Transform>();

        GenerateComponents();

        GenerateMatrix();
    }

    void Update()
    {
    }

    // Genera la matriz de juego inicial
    private void GenerateMatrix()
    {
        GameObject reference = (GameObject) Instantiate(Resources.Load("tile"));

        for (int row = 0; row < rows; row++)
            for (int col = 0; col < cols; col++)
            {
                GameObject tile = Instantiate(reference, transform);
                Point point = tile.GetComponent<Point>();

                float posX = col * tileSize;
                float posY = row * tileSize;

                tile.transform.position = new Vector2(posX, posY);

                PlayerComponents[col, row] = tile;
                PlayerMatrix[col, row] = -1; // Esto es solo para indicar que no hay nada en la matriz
                PlayerSet[col, row] = false; // Esto es solo para indicar que no hay nada en la matriz

                int tmpCol = col;
                int tmpRow = row;

                point.setPoint(tmpCol, tmpRow);

                matrix[col, row] = tile;
            }

        Destroy(reference);
    }

    // Genera la lista de componentes disponibles
    private void GenerateComponents()
    {
        GameObject reference = (GameObject) Instantiate(Resources.Load("Component"));

        for (int i = 0; i < ComponentImages.Length; i++)
        {
            GameObject component = Instantiate(reference, ComponentsContainer);

            Text name = component.transform.Find("ComponentName").GetComponent<Text>();
            Image image = component.transform.Find("ComponentImage").GetComponent<Image>();
            Button button = component.transform.Find("ComponentButton").GetComponent<Button>();

            name.text = ComponentNames[i];
            image.sprite = ComponentImages[i];

            int tmpIndex = i;

            button.onClick.AddListener(delegate {
                selectedComponent = tmpIndex;
            });
        }

        Destroy(reference);
    }
}
