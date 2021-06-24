
using System;

using Newtonsoft.Json;


[Serializable]
public class TileChunk
{
    [NonSerialized] public ResourceType resourceType;
    [NonSerialized] public BiomType biomTypeType;
    [NonSerialized] public BuildType buildType;
    public SerializableVector2Int pos;
    [NonSerialized]
    public SerializableVector2Int relPos;
    //копирует tiletype из-за проблем в сериализации tyletype
    public int rtype;
    public int btype;


    public int value = 1000;//DEFAULT
}