using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;


public class PlayerGrid : MonoBehaviour
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


    // Start is called before the first frame update
    void Start()
    {
        ComponentSize = new int[,] { { 1, 2 }, { 1, 1 }, { 1, 2 }, { 1, 2 }, { 2, 2 }, { 1, 2 } };

        PlayerComponents = new GameObject[cols, rows];

        ComponentsContainer = GameObject.Find("PlayerGrid").GetComponent<Transform>();

        GenerateMatrix();
    }

    // Update is called once per frame
    void Update(){    
        if (Network.matrix != null)
        {
            GenerateComponents();

            Network.matrix = null;
        }

        if (Network.graphConnections != null)
        {
            GenerateGraphPoints();

            Network.graphConnections = null;
        }
    }

    // Actualizamos los sprites de la matriz
    private void GenerateComponents()
    {
        for (int col = 0; col < cols; col++)
        {
            for (int row = 0; row < rows; row++)
            {
                if (Network.matrix[col, row] == -1 || Network.matrix[col, row] >= ComponentImages.Length) continue;

                AddComponent(matrix[col, row], (int) Network.matrix[col, row]);
            }
        }
    }

    // Generamos las lineas
    private void GenerateGraphPoints()
    {
        // Validamos que lleguen puntos correctos
        if (Network.graphConnections.Length % 2 != 0)
        {
            Debug.Log("Data de grafos incorrecta...");

            return;
        }

        GameObject refLine = (GameObject)Instantiate(Resources.Load("Line"), ComponentsContainer);

        // Generamos las lineas
        for (int i = 0; i < Network.graphConnections.Length; i += 4)
        {
            int col1 = (int) Network.graphConnections[i];
            int row1 = (int) Network.graphConnections[i+1];
            int col2 = (int) Network.graphConnections[i+2];
            int row2 = (int) Network.graphConnections[i+3];

            GameObject tile1 = matrix[col1, row1];
            GameObject tile2 = matrix[col2, row2];

            GameObject currentLine = Instantiate(refLine);

            LineRenderer line = currentLine.GetComponent<LineRenderer>();

            line.SetPosition(0, tile1.transform.position);
            line.SetPosition(1, tile2.transform.position);
        }

        Destroy(refLine);
    }

    private void GenerateMatrix()
    {
        GameObject reference = (GameObject)Instantiate(Resources.Load("tile"));

        for (int row = 0; row < rows; row++)
            for (int col = 0; col < cols; col++)
            {
                GameObject tile = Instantiate(reference, ComponentsContainer);
                Point point = tile.GetComponent<Point>();

                float posX = col * tileSize + ComponentsContainer.position.x;
                float posY = row * tileSize + ComponentsContainer.position.y;

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


    private void AddComponent(GameObject tile, int index)
    {
        Point point = tile.GetComponent<Point>();

        int col = point.getPoint()[0];
        int row = point.getPoint()[1];

        Debug.Log(col.ToString() + ", " + row.ToString());

        if (PlayerSet[col, row])
        {
            Debug.Log("Espacio ocupado...");
            return;
        }

        // Validamos que sea un espacio valido
        int[,] sizes = ComponentSize;

        bool outBoundsCol = cols <= col + sizes[index, 0] - 1;
        bool outBoundsRow = rows <= row + sizes[index, 1] - 1;

        if (outBoundsCol || outBoundsRow)
        {
            Debug.Log("Espacios fuera de la matriz....");

            return;
        }

        // Agregamos la imagen y la reescalamos
        Debug.Log("Agregando imagen...");

        PlayerMatrix[col, row] = index;
        PlayerSet[col, row] = true;
        PlayerSet[col + sizes[index, 0] - 1, row + sizes[index, 1] - 1] = true;

        GameObject img = (GameObject)Instantiate(Resources.Load("tileImg"), tile.transform);

        float width = sizes[index, 1];
        float height = sizes[index, 0];

        img.transform.localScale = new Vector2(width * 0.8f, height * 0.8f);

        img.transform.localPosition = new Vector2(width - 1.2f, height - 1.2f);

        SpriteRenderer renderer = img.GetComponent<SpriteRenderer>();

        renderer.sprite = ComponentImages[index];
    }
}
