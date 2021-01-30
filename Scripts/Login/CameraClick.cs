using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CameraClick : MonoBehaviour
{
    ComponentGenerator generator;

    private bool alreadyGenerated = false;

    private GameObject gameContainer;

    private void Start()
    {
        generator = GameObject.Find("Generator").GetComponent<ComponentGenerator>();
        gameContainer = GameObject.Find("Game");
    }

    void Update()
    {
        if (!alreadyGenerated) GenerateInitialComponents();

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

    private void GenerateInitialComponents()
    {
        // Generamos un mundo aleatorio
        ComponentGenerator.selectedComponent = 4;
        GameObject t1 = ComponentGenerator.matrix[Random.Range(0, 12), Random.Range(0, 12)];
        ClickedTile(t1);

        // Generamos un mercado aleatorio
        ComponentGenerator.selectedComponent = 2;
        GameObject t2 = ComponentGenerator.matrix[Random.Range(0, 12), Random.Range(0, 12)];
        ClickedTile(t2);

        alreadyGenerated = true;
    }

    private void Clicked(GameObject gameObject)
    {
        if (gameObject.CompareTag("tiles"))
        {
            ClickedTile(gameObject);
        }
    }

    private void ClickedTile(GameObject tile)
    {
        Point point = tile.GetComponent<Point>();

        int col = point.getPoint()[0];
        int row = point.getPoint()[1];

        int index = ComponentGenerator.selectedComponent;

        Debug.Log(col.ToString() + ", " + row.ToString());

        if (ComponentGenerator.PlayerSet[col, row])
        {
            Debug.Log("Espacio ocupado...");
            return;
        }

        // Validamos que sea un espacio valido
        int[,] sizes = generator.ComponentSize;

        bool outBoundsCol = generator.cols <= col + sizes[index, 0] - 1;
        bool outBoundsRow = generator.rows <= row + sizes[index, 1] - 1;

        if (outBoundsCol || outBoundsRow)
        {
            Debug.Log("Espacios fuera de la matriz....");

            return;
        }

        // Agregamos la imagen y la reescalamos
        Debug.Log("Agregando imagen...");

        ComponentGenerator.PlayerMatrix[col, row] = index;
        ComponentGenerator.PlayerSet[col, row] = true;
        ComponentGenerator.PlayerSet[col + sizes[index, 0] - 1, row + sizes[index, 1] - 1] = true;

        GameObject img = (GameObject) Instantiate(Resources.Load("tileImg"), tile.transform);

        float width = sizes[index, 1];
        float height = sizes[index, 0];

        img.transform.localScale = new Vector2(width * 0.8f, height * 0.8f);

        img.transform.localPosition = new Vector2(width - 1.2f, height - 1.2f);

        SpriteRenderer renderer = img.GetComponent<SpriteRenderer>();

        renderer.sprite = generator.ComponentImages[index];
    }
}
