
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class GenerationManager : Singleton<GenerationManager>
{
    private List<ScriptableObject> scriptableObjects;
    public int chunkSize = 40;
    [SerializeField] private int seed = 123456;
    public int worldWidth = 256;
    public int worldHeight = 256;
    private Vector2 perlinOffset;
    private readonly float perlinOffsetMax = 10000f;

    public BlockData defaultBlock;


    private void Awake()
    {
        scriptableObjects = Resources.LoadAll<ScriptableObject>("Blocks").ToList();
    }

    public IEnumerator GenerateChunk(Chunk chunk)
    {
        for (int v = 0; v < chunkSize; v++)
        {
            for (int h = 0; h < chunkSize; h++)
            {
                Vector3Int tilePosition = new Vector3Int(chunk.Position.x + h, chunk.Position.y + v, 0);
                if ((tilePosition.x < 0 || tilePosition.x >= worldWidth) ||
                    (tilePosition.y < 0 || tilePosition.y >= worldHeight))
                    continue;

                BlockData blockData = defaultBlock;


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