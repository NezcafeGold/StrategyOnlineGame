
using UnityEngine;
using UnityEngine.Tilemaps;

[CreateAssetMenu(menuName = "Blocks/BiomsData")]
public class BiomObjectData : ScriptableObject
{
    public string title;
    public BiomType type;
    public Tile tile;
    public float perlinLevel;
    public float perlinSpeed;
}