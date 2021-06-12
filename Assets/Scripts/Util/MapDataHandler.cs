


using System.Collections.Generic;
using System.IO;
using System.Linq;
using Leguar.TotalJSON;
using UnityEngine;

public class MapDataHandler : MonoBehaviour
{
    private string outputPath;
    private GameObject chunks;
    //private Dictionary<SerializableVector2Int, Chunk>  

    private void Awake()
    {
        outputPath = SetupSetting.Instance.outputPath;
        chunks = SetupSetting.Instance.chunkRoot;
        LoadMapFromJson();
    }

    public void SaveMapToJson()
    {
        List<Chunk> chunks = ChunkLoadManager.Instance.chunks;
        JSON mapJson = new JSON();
        
        JArray chunkArrayJson = new JArray();
        foreach (var ch in chunks)
        {
            JSON chunkJson = new JSON();
            JSON posJson = JSON.Serialize(ch.Position);
            chunkJson.Add("chunk", posJson);
            

            JArray tilesArrayJson = new JArray();

            foreach (var tileCh in ch.tileChunkLayer)
            {
                JSON tileChJson = JSON.Serialize(tileCh);
                tilesArrayJson.Add(tileChJson);
                
            }
            chunkJson.Add("tiles", tilesArrayJson);
            
            chunkArrayJson.Add(chunkJson);
        }
       mapJson.Add("map", chunkArrayJson);
       File.WriteAllText(outputPath, mapJson.CreateString());
       Debug.Log("Print at " + outputPath);
    }



    public void LoadMapFromJson()
    {
       
    }

}