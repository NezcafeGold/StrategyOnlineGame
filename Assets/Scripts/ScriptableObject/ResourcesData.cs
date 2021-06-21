
using System;
using UnityEngine;
using UnityEngine.Tilemaps;

[CreateAssetMenu(menuName = "Blocks/BlockData")]
public class ResourcesData : ScriptableObject
{
    public string title;
    public ResourceType type;
    public Tile tile;
    public float perlinLevel;
    public float perlinSpeed;
    public TileBase tilebas;
    public Sprite sprite;

}
