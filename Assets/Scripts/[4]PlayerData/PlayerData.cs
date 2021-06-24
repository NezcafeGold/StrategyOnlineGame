using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using Model.BuildData;
using UnityEngine;

public class PlayerData 
{
    public SerializableVector2Int SpawnCoord = new SerializableVector2Int(0, 0);
    public Dictionary<SerializableVector2Int, ChunkData> ChunkMap = new Dictionary<SerializableVector2Int, ChunkData>();
    private Queue<Action> chunkQueue = new Queue<Action>();

    public string Nickname;
    public int Level;
    public int Experience;
    public Dictionary<ResourceType, int> ResourcesDictionary = new Dictionary<ResourceType, int>();
    public InventoryData InventoryData;
    public PlayerBaseData baseData;
    
    
    private bool isLoad = false;
    
    private static readonly PlayerData instance = new PlayerData();
    public string Date { get; private set; }
 
    private PlayerData()
    {
        Date = DateTime.Now.TimeOfDay.ToString();
    }
 
    public static PlayerData GetInstance()
    {
        return instance;
    }
    
    public int GetValueForType(ResourceType resType)
    {
        foreach (var k in ResourcesDictionary.Keys)
        {
            if (resType.Equals(k))
                return ResourcesDictionary[k];
        }
        return 0;
    }

    public void LoadTrue()
    {
        Messenger.Broadcast(GameEvent.UPDATE_USER_STATS);
    }

    public  bool isLoadDone()
    {
        return isLoad;
    }
}