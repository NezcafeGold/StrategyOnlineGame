using UnityEngine;

[CreateAssetMenu(menuName = "Blocks/Building")]
public class BuildingObjectData : ScriptableObject
{
    public int buildTime;
    public Sprite sprite;
    public BuildType type;
}
