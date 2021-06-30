using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

public class InventoryItem : MonoBehaviour, IPointerClickHandler
{
    [SerializeField] private Image itemSprite;
    [NonSerialized] public ItemObjectData ItemObjectData;

    public void SetData(ItemObjectData itemObjectData)
    {
        itemSprite.sprite = itemObjectData.Sprite;
        this.ItemObjectData = itemObjectData;
    }

    private void OnMouseDown()
    {
        
    }
    
    public void OnPointerClick(PointerEventData eventData)
    {
        
        DetailInfo.Instance.ShowItemDetail(this);
    }
}