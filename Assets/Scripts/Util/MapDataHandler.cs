using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using Leguar.TotalJSON;
using UnityEngine;

public class MapDataHandler : Singleton<MapDataHandler>
{
    private string outputPath;
    private GameObject chunks;
    public Dictionary<SerializableVector2Int, ChunkData> chunkDataMap;

    private void Awake()
    {
        outputPath = SetupSetting.Instance.outputPath;
        chunks = SetupSetting.Instance.chunkRoot;
    }

    private void Start()
    {
        if (!SetupSetting.Instance.isMasterClient)
            StartCoroutine(LoadMapFromJson());
    }

    public void SaveMapToJson()
    {
        StartCoroutine(SaveMapToJsonCor());
    }
    
    public IEnumerator SaveMapToJsonCor()
    {
        yield return StartCoroutine(PerformSaveMapToJsonCor());
    }

    public IEnumerator LoadMapFromJson()
    {
        yield return StartCoroutine(LoadMapFromJsonCor());
    }

    private IEnumerator PerformSaveMapToJsonCor()
    {
        List<Chunk> chunks = ChunkLoadManager.Instance.chunksToMasterMapSave;
        JSON mapJson = new JSON();

        JArray chunkArrayJson = new JArray();
        foreach (var ch in chunks)
        {
            JSON chunkJson = new JSON();
            JSON posJson = JSON.Serialize(ch.chunkData.Position);
            chunkJson.Add("chunk", posJson);

            JArray tilesArrayJson = new JArray();

            foreach (var tileCh in ch.chunkData.tileChunkLayer)
            {
                JSON tileChJson = JSON.Serialize(tileCh);
                tilesArrayJson.Add(tileChJson);
            }

            chunkJson.Add("tiles", tilesArrayJson);

            chunkArrayJson.Add(chunkJson);
        }

        mapJson.Add("chunkSize", SetupSetting.Instance.chunkSize);
        mapJson.Add("map", chunkArrayJson);
        File.WriteAllText(outputPath, mapJson.CreateString());
        Debug.Log("Print at " + outputPath);
        yield return null;
    }


    public IEnumerator LoadMapFromJsonCor()
    {
        chunkDataMap = new Dictionary<SerializableVector2Int, ChunkData>();
        Debug.Log("Start Read Text");
        String text = File.ReadAllText(outputPath);


        JSON textJson = JSON.ParseString(text);
        JArray mapJson = textJson.GetJArray("map");
        mapJson.SetProtected();
        mapJson.DebugInEditor("MapJson");
        Debug.Log("End Read Text");
        int chunkSize = textJson.GetInt("chunkSize");
        SetupSetting.Instance.chunkSize = chunkSize;
        
        foreach (JSON v in mapJson.Values)
        {
            SerializableVector2Int pos = v.GetJSON("chunk").Deserialize<SerializableVector2Int>();
            ChunkData ch = new ChunkData();
            ch.Position = pos;
            JArray tiles = v.GetJArray("tiles");

            TileChunk[,] tileChunkLayer = new TileChunk[chunkSize, chunkSize];
            foreach (JSON t in tiles.Values)
            {
                TileChunk tileChunk = t.Deserialize<TileChunk>();
                tileChunkLayer[tileChunk.relPos.x, tileChunk.relPos.y] = tileChunk;
            }

            ch.tileChunkLayer = tileChunkLayer;
            chunkDataMap.Add(pos, ch);
            Debug.Log(pos.x + " " + pos.y);
        }
        yield return null;
    }
}