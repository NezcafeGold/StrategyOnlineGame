
using System;

public class TileChunk
{
    [NonSerialized] public TileType tileType;
    public SerializableVector2Int pos;
    public SerializableVector2Int relPos;
    //копирует tiletype из-за проблем в сериализации tyletype
    public String type;
}