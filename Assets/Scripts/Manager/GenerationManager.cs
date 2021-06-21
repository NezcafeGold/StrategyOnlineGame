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
    private ResourcesData defaultResourceResources;
    private BiomData defaultBiomBlock;

    private Vector2 perlinOffset;
    private List<ResourcesData> blockDataObjects;
    private List<BiomData> biomDataObjects;
    private bool isMasterClient;


    private void Awake()
    {
        var setup = SetupSetting.Instance;
        chunkSize = setup.chunkSize;
        seed = setup.seed;
        worldWidth = setup.worldWidth;
        worldHeight = setup.worldHeight;
        defaultResourceResources = setup.defaultResourceResources;
        defaultBiomBlock = setup.defaultBiomBlock;

        blockDataObjects = Resources.LoadAll<ResourcesData>("Blocks").ToList();
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


                ResourcesData resResourcesData = defaultResourceResources;
                BiomData biomBlockData = defaultBiomBlock;


                if (isMasterClient)
                {
                    #region generation resources

                    // Бегаем по блокам и проверяем шанс
                    for (int i = 0; i < blockDataObjects.Count; i++)
                    {
                        ResourcesData resources = blockDataObjects[i] as ResourcesData;
                        if (resources != defaultBiomBlock)
                        {
                            if (!CheckPerlinLevel(tilePosition, resources.perlinSpeed, resources.perlinLevel))
                            {
                                resResourcesData = resources;
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
                        if (chunk.GetTileChunkData(tilePosition).biomTypeType.Equals(block.type))
                        {
                            biomBlockData = block;
                            break;
                        }
                    }

                    for (int i = 0; i < blockDataObjects.Count; i++)
                    {
                        ResourcesData resources = blockDataObjects[i] as ResourcesData;
                        if (chunk.GetTileChunkData(tilePosition).resourceType.Equals(resources.type))
                        {
                            resResourcesData = resources;
                            break;
                        }
                    }
                }

                chunk.SetChunkTile(tilePosition, biomBlockData.tile);

                chunk.SetTileChunkData(tilePosition, resResourcesData.type, biomBlockData.type);

                if (resResourcesData.tilebas != null)
                    chunk.SetChunkTile(tilePosition, resResourcesData.tilebas, true);
                else
                    chunk.SetChunkTile(tilePosition, resResourcesData.tile, true);
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