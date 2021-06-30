using TMPro;
using UnityEngine;
using UnityEngine.UI;

public class DetailInfo : Singleton<DetailInfo>
{
    [SerializeField] private TextMeshProUGUI equipmentText;
    [SerializeField] private TextMeshProUGUI titleText;
    [SerializeField] private TextMeshProUGUI descriptionText;
    [SerializeField] private Transform itemParent;
    [SerializeField] private VerticalLayoutGroup statsList;
    [SerializeField] private HorizontalLayoutGroup textStatPrefab;

    private InventoryItem currentInventoryItem;

    public void ShowItemDetail(InventoryItem inventoryItem)
    {
        titleText.text = inventoryItem.ItemObjectData.Title;
        descriptionText.text = inventoryItem.ItemObjectData.Description;

        if (inventoryItem.ItemObjectData as WeaponObjectData)
        {
            
        }

        if (currentInventoryItem != null)
            Destroy(currentInventoryItem.gameObject);

        currentInventoryItem = Instantiate(inventoryItem, itemParent.position, Quaternion.identity, itemParent);
        
        
        // currentInventoryItem.transform.SetParent(itemParent);
    }
}