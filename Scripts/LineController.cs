using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LineController : MonoBehaviour
{
    private LineRenderer line;
    private Vector3 mousepos;
    private Material material;
    private int currLines = 0;

    void Start()
    {

    }



    void Update()
    {
        if (Input.GetMouseButtonDown(0))
        {
            if (line = null)
            {
                createLine();
            }

            mousepos = Camera.main.ScreenToWorldPoint(Input.mousePosition);
            mousepos.z = 0;
            line.SetPosition(0, mousepos);
            line.SetPosition(1, mousepos);  
        }
        else if (Input.GetMouseButtonUp(0) && line)
        {

            mousepos = Camera.main.ScreenToWorldPoint(Input.mousePosition);
            mousepos.z = 0;
            line.SetPosition(1, mousepos);
            line = null;
            currLines++;
        }
        else if (Input.GetMouseButton(0) && line)
        {
            mousepos = Camera.main.ScreenToWorldPoint(Input.mousePosition);
            mousepos.z = 0;
            line.SetPosition(1, mousepos);
        }

    }

    void createLine()
    {
        line = new GameObject("Line" + currLines).AddComponent<LineRenderer>();
        line.material = material;
        line.positionCount = 2;
        line.startWidth = 0.15f;
        line.endWidth = 0.15f;
        line.useWorldSpace = true;
        line.numCapVertices = 50;
    }

}
