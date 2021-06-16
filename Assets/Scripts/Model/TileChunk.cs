
using System;

[Serializable]
public class TileChunk
{
    [NonSerialized] public TileType tileType;
    public SerializableVector2Int pos;
    [NonSerialized]
    public SerializableVector2Int relPos;
    //копирует tiletype из-за проблем в сериализации tyletype
    public int type;
}