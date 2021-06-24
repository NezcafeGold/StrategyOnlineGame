using System;
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
    private ResourcesObjectData _defaultResourceObjectResourcesObject;
    private BiomObjectData _defaultBiomObjectBlock;

    private Vector2 perlinOffset;
    private List<ResourcesObjectData> blockDataObjects;
    private List<BiomObjectData> biomDataObjects;
    private bool isMasterClient;


    private void Awake()
    {
        var setup = SetupSetting.Instance;
        chunkSize = setup.chunkSize;
        seed = setup.seed;
        worldWidth = setup.worldWidth;
        worldHeight = setup.worldHeight;
        _defaultResourceObjectResourcesObject = setup.defaultResourceObjectResourcesObject;
        _defaultBiomObjectBlock = setup.defaultBiomObjectBlock;

        blockDataObjects = Resources.LoadAll<ResourcesObjectData>("Blocks").ToList();
        biomDataObjects = Resources.LoadAll<BiomObjectData>("Blocks").ToList();

        isMasterClient = SetupSetting.Instance.isMasterClient;
    }

    public IEnumerator GenerateChunk(Chunk chunk)
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


                ResourcesObjectData resResourcesObjectData = _defaultResourceObjectResourcesObject;
                BiomObjectData biomObjectBlockData = _defaultBiomObjectBlock;


                if (isMasterClient)
                {
                    #region generation resources

                    // Бегаем по блокам и проверяем шанс
                    for (int i = 0; i < blockDataObjects.Count; i++)
                    {
                        ResourcesObjectData resourcesObject = blockDataObjects[i] as ResourcesObjectData;
                        if (resourcesObject != _defaultResourceObjectResourcesObject)
                        {
                            if (!CheckPerlinLevel(tilePosition, resourcesObject.perlinSpeed, resourcesObject.perlinLevel))
                            {
                                resResourcesObjectData = resourcesObject;
                                break;
                            }
                        }
                    }

                    #endregion

                    #region generation bioms

                    for (int i = 0; i < biomDataObjects.Count; i++)
                    {
                        BiomObjectData block = biomDataObjects[i] as BiomObjectData;
                        if (block != _defaultBiomObjectBlock)
                        {
                            if (!CheckPerlinLevel(tilePosition, block.perlinSpeed, block.perlinLevel))
                            {
                                biomObjectBlockData = block;
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
                        BiomObjectData block = biomDataObjects[i] as BiomObjectData;
                        if (chunk.GetTileChunkData(tilePosition).biomTypeType.Equals(block.type))
                        {
                            biomObjectBlockData = block;
                            break;
                        }
                    }

                    for (int i = 0; i < blockDataObjects.Count; i++)
                    {
                        ResourcesObjectData resourcesObject = blockDataObjects[i] as ResourcesObjectData;
                        if (chunk.GetTileChunkData(tilePosition).resourceType.Equals(resourcesObject.type))
                        {
                            resResourcesObjectData = resourcesObject;
                            break;
                        }
                    }
                }

                chunk.SetChunkTile(tilePosition, biomObjectBlockData.tile);
                chunk.SetTileChunkData(tilePosition, resResourcesObjectData.type, biomObjectBlockData.type);

                if (resResourcesObjectData.tilebas != null)
                    chunk.SetChunkTile(tilePosition, resResourcesObjectData.tilebas, true);
                else
                    chunk.SetChunkTile(tilePosition, resResourcesObjectData.tile, true);
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