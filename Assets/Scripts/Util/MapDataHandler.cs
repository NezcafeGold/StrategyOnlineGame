using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Runtime.Serialization.Formatters.Binary;
using Leguar.TotalJSON;
using UnityEngine;

public class MapDataHandler : Singleton<MapDataHandler>
{
    private string jsonPath;
    private string dictionaryPath;
    private GameObject chunks;
    public SerializableDictionary<SerializableVector2Int, ChunkData> chunkDataMap;

    private void Awake()
    {
        if (!SetupSetting.Instance.isMasterClient) gameObject.SetActive(false);
        jsonPath = SetupSetting.Instance.jsonPath;
        chunks = SetupSetting.Instance.chunkRoot;
        dictionaryPath = SetupSetting.Instance.dictionaryPath;
    }

    private void Start()
    {
        //if (!SetupSetting.Instance.isMasterClient)
        //StartCoroutine(LoadMapFromJson());
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
            chunkJson.Add("pos", posJson); //pos:{x,y}

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
        mapJson.Add("seed", SetupSetting.Instance.seed);
        mapJson.Add("width", SetupSetting.Instance.worldWidth);
        mapJson.Add("height", SetupSetting.Instance.worldHeight);
        mapJson.Add("map", chunkArrayJson);
        File.WriteAllText(jsonPath, mapJson.CreateString());
        Debug.Log("Print at " + jsonPath);
        yield return null;
    }


    public IEnumerator LoadMapFromJsonCor()
    {
        chunkDataMap = new SerializableDictionary<SerializableVector2Int, ChunkData>();
        Debug.Log("Start Read JSON MAP");
        String text = File.ReadAllText(jsonPath);


        JSON textJson = JSON.ParseString(text);
        JArray mapJson = textJson.GetJArray("map");
        mapJson.SetProtected();
        mapJson.DebugInEditor("MapJson");
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
        }


//        BinaryFormatter formatter = new BinaryFormatter();
//
//        using (FileStream stream = new FileStream(dictionaryPath, FileMode.Create))
//        {
//            formatter.Serialize(stream, chunkDataMap);
//        }
        yield return null;
    }
}