
using System;

[Serializable]
public class TileChunk
{
    [NonSerialized] public ResourceType resourceType;
    [NonSerialized] public BiomType biomTypeType;
    public SerializableVector2Int pos;
    [NonSerialized]
    public SerializableVector2Int relPos;
    //копирует tiletype из-за проблем в сериализации tyletype
    public int type;
}