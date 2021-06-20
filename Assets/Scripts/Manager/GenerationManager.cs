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
    private BlockData defaultResourceBlock;
    private BiomData defaultBiomBlock;

    private Vector2 perlinOffset;
    private List<BlockData> blockDataObjects;
    private List<BiomData> biomDataObjects;
    private bool isMasterClient;


    private void Awake()
    {
        var setup = SetupSetting.Instance;
        chunkSize = setup.chunkSize;
        seed = setup.seed;
        worldWidth = setup.worldWidth;
        worldHeight = setup.worldHeight;
        defaultResourceBlock = setup.defaultResourceBlock;
        defaultBiomBlock = setup.defaultBiomBlock;

        blockDataObjects = Resources.LoadAll<BlockData>("Blocks").ToList();
        biomDataObjects = Resources.LoadAll<BiomData>("Blocks").ToList();

        isMasterClient = SetupSetting.Instance.isMasterClient;
    }

    public IEnumerator GenerateChunk(Chunk chunk, bool isBiom = false)
    {
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


                BlockData resBlockData = defaultResourceBlock;
                BiomData biomBlockData = defaultBiomBlock;


                if (isMasterClient)
                {
                    #region generation resources

                    // Бегаем по блокам и проверяем шанс
                    for (int i = 0; i < blockDataObjects.Count; i++)
                    {
                        BlockData block = blockDataObjects[i] as BlockData;
                        if (block != defaultBiomBlock)
                        {
                            if (!CheckPerlinLevel(tilePosition, block.perlinSpeed, block.perlinLevel))
                            {
                                resBlockData = block;
                                break;
                            }
                        }
                    }

                    #endregion

                    #region generation bioms

                    for (int i = 0; i < biomDataObjects.Count; i++)
                    {
                        BiomData block = biomDataObjects[i] as BiomData;
                        if (block != defaultBiomBlock)
                        {
                            if (!CheckPerlinLevel(tilePosition, block.perlinSpeed, block.perlinLevel))
                            {
                                biomBlockData = block;
                                break;
                            }
                        }
                    }

                    #endregion
                }
                else
                {
                    for (int i = 0; i < biomDataObjects.Count; i++)
                    {
                        BiomData block = biomDataObjects[i] as BiomData;
                        if (chunk.GetTileChunkData(tilePosition).resourceType.Equals(block.type))
                        {
                            biomBlockData = block;
                            break;
                        }
                    }

                    for (int i = 0; i < blockDataObjects.Count; i++)
                    {
                        BlockData block = blockDataObjects[i] as BlockData;
                        if (chunk.GetTileChunkData(tilePosition).resourceType.Equals(block.type))
                        {
                            resBlockData = block;
                            break;
                        }
                    }
                }


                chunk.SetChunkTile(tilePosition, biomBlockData.tile);
                
                chunk.SetTileChunkData(tilePosition, resBlockData.type, biomBlockData.type); 

                if (resBlockData.tilebas != null)
                    chunk.SetChunkTile(tilePosition, resBlockData.tilebas, true);
                else
                    chunk.SetChunkTile(tilePosition, resBlockData.tile, true);
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