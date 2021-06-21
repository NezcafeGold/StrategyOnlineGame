using System;
using System.Collections.Generic;
using System.Net.NetworkInformation;
using System.Net.Sockets;
using UnityEngine;

public class PlayerData : Singleton<PlayerData>
{
    public SerializableVector2Int SpawnCoord = new SerializableVector2Int(0, 0);
    public Dictionary<SerializableVector2Int, ChunkData> ChunkMap;
    private Queue<Action> chunkQueue;

    [SerializeField] public string Nickname;
    [SerializeField] public int Level;
    [SerializeField] public int Experience;
    [SerializeField] public Dictionary<ResourceType, int> ResourcesDictionary;
    [SerializeField] public InventoryData InventoryData;


    private void Awake()
    {
        ResourcesDictionary = new Dictionary<ResourceType, int>();
        ChunkMap = new Dictionary<SerializableVector2Int, ChunkData>();
        chunkQueue = new Queue<Action>();
        DontDestroyOnLoad(this);
        AskData();
    }
    

    private void AskData()
    {
        TCPClient.Instance.SendMessageTCP(new Packet(Packet.SegmentID.GET_USER_ID,
            Packet.StatusCode.OK_CODE).ToString());
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
}