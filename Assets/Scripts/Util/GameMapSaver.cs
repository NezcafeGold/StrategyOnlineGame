using System;
using System.Collections.Generic;
using System.IO;
using Leguar.TotalJSON;
using UnityEngine;

public class GameMapSaver : MonoBehaviour
{
    [SerializeField] private GameObject chunks;
    public void SaveMapToJson()
    {
        List<Chunk> chunks = ChunkLoadManager.Instance.chunks;
        foreach (var ch in chunks)
        {
            
            var outputPath = @"JSONFILES/Chunk " + ch.Position.x + " " + ch.Position.y + ".txt";

            JSON json = new JSON();
            json.Add("xPos", ch.Position.x);
            json.Add("yPos", ch.Position.y);
            
            JArray tilesArray = new JArray();

            foreach (var tileChunk in ch.tileChunkLayer)
            {
                Tile tile = new Tile();
                tile.type = tileChunk.TileType.ToString();
                tile.xPos = tileChunk.position.x;
                tile.yPos = tileChunk.position.y;
                tile.xRelPos = tileChunk.relativePosition.x;
                tile.yRelPos = tileChunk.relativePosition.y;
                tilesArray.Add(JSON.Serialize(tile));
            }

            json.Add("tiles", tilesArray);

           
            //json.Add("Tiles", JSON.Serialize(ch.tileChunkLayer));

            
            File.WriteAllText (outputPath, json.CreatePrettyString());
            Debug.Log("Print at " + outputPath);
        }
    }

    class Tile {
        public string type;
        public int xPos;
        public int yPos;
        public int xRelPos;
        public int yRelPos;
    }
}