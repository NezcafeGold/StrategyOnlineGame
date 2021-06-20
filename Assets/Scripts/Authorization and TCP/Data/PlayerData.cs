using System;
using System.Collections.Generic;
using UnityEngine;

public class PlayerData : Singleton<PlayerData>
{
    public SerializableVector2Int SpawnCoord = new SerializableVector2Int(0, 0);
    public Dictionary<SerializableVector2Int, ChunkData> ChunkMap;
    private Queue<Action> chunkQueue;

    private void Awake()
    {
        ChunkMap = new Dictionary<SerializableVector2Int, ChunkData>();
        chunkQueue = new Queue<Action>();
        DontDestroyOnLoad(this);
    }

    private void Update()
    {
        if (chunkQueue.Count > 0)
            chunkQueue.Dequeue().Invoke();
    }

    public void AddToChunkQueue(SerializableVector2Int pos)
    {
        chunkQueue.Enqueue(() => GetChunk(pos));
    }
    
        
    private ChunkData GetChunk(SerializableVector2Int pos)
    {
        long milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
        if (ChunkMap.ContainsKey(pos)) return ChunkMap[pos];
        TCPClient.Instance.SendMessageTCP(new Packet(Packet.SegmentID.GET_CHUNK_ID,
            Packet.StatusCode.OK_CODE, Packet.Body.OfInt("x", pos.x), Packet.Body.OfInt("y", pos.y)).ToString());
        while (!ChunkMap.ContainsKey(pos))
        {
            if (milliseconds + 10000 < DateTimeOffset.Now.ToUnixTimeMilliseconds())
                Debug.Log("CANT FIND CHUNK FROM TCP");
                return null;
        }

        return ChunkMap[pos];
    }
}