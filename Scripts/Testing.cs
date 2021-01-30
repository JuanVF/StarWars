using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using CodeMonkey.Utils;

public class Testing : MonoBehaviour{

    private Grid grid;

    private void Start(){
        grid = new Grid(10, 10,10f, new Vector3(20,0);
    }

    private void Update(){
        if (Input.GetMouseButtonDown(0)){
            grid.SetValue(UtilsClass.GetMouseWorldPosition(), 55);
        }
        if (Input.GetMouseButtonDown(1)){
            Debug.Log(grid.GetValue(UtilsClass.GetMouseWorldPosition()));
        }
    }

}
