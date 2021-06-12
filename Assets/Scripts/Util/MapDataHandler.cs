using System.Collections.Generic;
using System.IO;
using System.Text;
using UnityEngine;
using LitJson;

public class MapDataHandler : MonoBehaviour
{
    private GameObject chunks;

    private void Awake()
    {
        chunks = SetupSetting.Instance.chunkRoot;
    }

    public void SaveMapToJson()
    {
        List<Chunk> chunks = ChunkLoadManager.Instance.chunks;
        var outputPath = @"JSONFILES/ChunkMap.json";
        StringBuilder sb = new StringBuilder();
        JsonWriter writer = new JsonWriter(sb);

        writer.WriteObjectStart();
        writer.WritePropertyName("map");
        writer.WriteArrayStart();
        foreach (var ch in chunks)
        {

            writer.WriteObjectStart();
            writer.WritePropertyName("chunk");
            writer.Write(JsonMapper.ToJson(ch.ChunkPosition));
            
            writer.WritePropertyName("tiles");
            writer.Write(JsonMapper.ToJson(ch.tileChunkLayer));
            writer.WriteObjectEnd();
            

        }
        writer.WriteArrayEnd();
        writer.WriteObjectEnd();
        File.WriteAllText (outputPath, sb.ToString().Replace("\\",""));
        Debug.Log("Print at " + outputPath);
    }


}