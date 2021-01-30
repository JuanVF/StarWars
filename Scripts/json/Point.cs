using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Point : MonoBehaviour
{
    private int col;
    private int row;

    public void setPoint(int _col, int _row)
    {
        col = _col;
        row = _row;
    }

    public int[] getPoint()
    {
        return new int[] { col, row };
    }
}
