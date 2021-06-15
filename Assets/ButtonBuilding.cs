using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ButtonBuilding : MonoBehaviour
{
    [SerializeField] private DragToBuild d;
    [SerializeField] private GameObject buttons;
    [SerializeField] private Type type;

    enum Type
    {
        YES,
        NO
    }

    private void OnMouseDown()
    {
        if (type.Equals(Type.YES))
            OnPressedYes();
        else if (type.Equals(Type.NO))
            OnPressedNo();
    }

    public void OnPressedYes()
    {
        d.dragDisabled = false;
        buttons.SetActive(false);
    }

    public void OnPressedNo()
    {
        Destroy(d.gameObject);
    }
}