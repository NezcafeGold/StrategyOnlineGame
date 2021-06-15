using UnityEngine;

[CreateAssetMenu(menuName = "Blocks/Building")]
public class BuildingData : ScriptableObject
{
    public enum BuildingType
    {
        NONE,
        HOUSE
    }
    public GameObject prefab;
    public int buildTime;
    public Sprite sprite;
    public BuildingType type;

}
