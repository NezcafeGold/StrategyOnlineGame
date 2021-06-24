using UnityEngine;

[CreateAssetMenu(menuName = "Blocks/Building")]
public class BuildingObjectData : ScriptableObject
{
    public GameObject prefab;
    public int buildTime;
    public Sprite sprite;
    public BuildType type;

}
