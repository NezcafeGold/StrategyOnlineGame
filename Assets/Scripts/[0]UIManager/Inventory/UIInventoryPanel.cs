using System;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;
using UnityEngine.UI;

public class UIInventoryPanel : PanelBase
{
    private List<WeaponObjectData> weapons;
    [SerializeField] private InventoryItem prefabItem;
    [SerializeField] private GridLayoutGroup parentObject;
    [SerializeField] private DetailInfo detailInfo;

    private void Awake()
    {
        weapons = Resources.LoadAll<WeaponObjectData>("Inventory").ToList();
    }

    private void Start()
    {
        InventoryItem firstItem = null;
        foreach (var wep in weapons)
        {
            InventoryItem inventoryItem =
                Instantiate(prefabItem, Vector3.zero, Quaternion.identity, parentObject.transform);
            if (firstItem == null)
                firstItem = inventoryItem;
            inventoryItem.SetData(wep);
        }

        detailInfo.ShowItemDetail(firstItem);
    }
}