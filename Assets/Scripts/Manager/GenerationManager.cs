using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class GenerationManager : Singleton<GenerationManager>
{
    private int chunkSize = 40;
    private int seed = 123456;
    private int worldWidth = 256;
    private int worldHeight = 256;
    private BlockData defaultBlock;

    private Vector2 perlinOffset;
    private List<ScriptableObject> scriptableObjects;


    private void Awake()
    {
        var setup = SetupSetting.Instance;
        chunkSize = setup.chunkSize;
        seed = setup.seed;
        worldWidth = setup.worldWidth;
        worldHeight = setup.worldHeight;
        defaultBlock = setup.defaultBlock;

        scriptableObjects = Resources.LoadAll<ScriptableObject>("Blocks").ToList();
    }

    public IEnumerator GenerateChunk(Chunk chunk)
    {
        ChunkData ch = null;
        if (!SetupSetting.Instance.isMasterClient)
        {
           // ch = PlayerData.Instance.GetChunk(chunk.chunkData.Position);
        }

        //TODO: TCP
        for (int v = 0; v < chunkSize; v++)
        {
            for (int h = 0; h < chunkSize; h++)
            {
                Vector3Int tilePosition =
                    new Vector3Int(chunk.chunkData.Position.x + h, chunk.chunkData.Position.y + v, 0);
                if ((tilePosition.x < 0 || tilePosition.x >= worldWidth) ||
                    (tilePosition.y < 0 || tilePosition.y >= worldHeight))
                    continue;

                BlockData blockData = defaultBlock;

                if (SetupSetting.Instance.isMasterClient)
                {
                    // Бегаем по блокам и проверяем шанс
                    for (int i = 0; i < scriptableObjects.Count; i++)
                    {
                        BlockData block = scriptableObjects[i] as BlockData;
                        if (block != defaultBlock)
                        {
                            if (!CheckPerlinLevel(tilePosition, block.perlinSpeed, block.perlinLevel))
                            {
                                blockData = block;
                                break;
                            }
                        }
                    }
                }
                else
                {
                    for (int i = 0; i < scriptableObjects.Count; i++)
                    {
                        BlockData block = scriptableObjects[i] as BlockData;
                        if (block.type.Equals(chunk.chunkData.tileChunkLayer[h, v].tileType))
                        {
                            blockData = block;
                            break;
                        }
                    }
                }

                chunk.SetChunkTile(tilePosition, defaultBlock.tile);
                chunk.SetTileChunkData(tilePosition, TileType.NONE);
                if (blockData.tile != defaultBlock.tile)
                {
                    chunk.SetChunkTile(tilePosition, blockData.tile, true);
                    chunk.SetTileChunkData(tilePosition, blockData.type);
                }
            }
        }

        yield return null;
    }


    public bool CheckPerlinLevel(Vector3Int tilePosition, float perlinSpeed, float perlinLevel)
    {
        return (Mathf.PerlinNoise(
                    perlinOffset.x + tilePosition.x * perlinSpeed,
                    perlinOffset.y + tilePosition.y * perlinSpeed) +
                Mathf.PerlinNoise(
                    perlinOffset.x - tilePosition.x * perlinSpeed,
                    perlinOffset.y - tilePosition.y * perlinSpeed)) / 2f >= perlinLevel;
    }
}