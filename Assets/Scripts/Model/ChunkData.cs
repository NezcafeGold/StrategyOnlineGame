using System;
using UnityEngine;

[Serializable]
public class ChunkData
{
    public SerializableVector2Int Position;
    [NonSerialized]
    public SerializableVector2Int ChunkPosition;
    public TileChunk[,] tileChunkLayer;
}