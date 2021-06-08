
using System;
using UnityEngine;
using UnityEngine.Tilemaps;

[CreateAssetMenu(menuName = "Blocks/BlockData")]
public class BlockData : ScriptableObject
{
    public String title;
    public Chunk.TileType type;
    public Tile tile;
    public Sprite sprite;
    public float perlinLevel;
    public float perlinSpeed;

}
