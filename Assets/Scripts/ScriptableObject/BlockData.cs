
using System;
using UnityEngine;
using UnityEngine.Tilemaps;

[CreateAssetMenu(menuName = "Blocks/BlockData")]
public class BlockData : ScriptableObject
{
    public String title;
    public TileType type;
    public Tile tile;
    public float perlinLevel;
    public float perlinSpeed;

}
