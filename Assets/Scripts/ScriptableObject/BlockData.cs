
using System;
using UnityEngine;
using UnityEngine.Tilemaps;

[CreateAssetMenu(menuName = "Blocks/BlockData")]
public class BlockData : ScriptableObject
{
    public string title;
    public ResourceType type;
    public Tile tile;
    public float perlinLevel;
    public float perlinSpeed;

}
