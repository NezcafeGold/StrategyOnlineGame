using System;
using System.Collections.Generic;
using UnityEngine;

public class PlayerData : Singleton<PlayerData>
{
    public SerializableVector2Int SpawnCoord = new SerializableVector2Int(0, 0);
    public Dictionary<SerializableVector2Int, ChunkData> ChunkMap;
    private Queue<Action> chunkQueue;

    private string name;
    private Dictionary<ResourceType, int> resourcesDictionary;
    
    private void Awake()
    {
        ChunkMap = new Dictionary<SerializableVector2Int, ChunkData>();
        chunkQueue = new Queue<Action>();
        DontDestroyOnLoad(this);
    }

    private void UodateData()
    {
        
    }
}