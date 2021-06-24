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
    private List<BuildingObjectData> buildingDataObjects;
    private bool isMasterClient;
    private Building buildingPrefab;


    private void Awake()
    {
        var setup = SetupSetting.Instance;
        chunkSize = setup.chunkSize;
        seed = setup.seed;
        worldWidth = setup.worldWidth;
        worldHeight = setup.worldHeight;
        _defaultResourceObjectResourcesObject = setup.defaultResourceObjectResourcesObject;
        _defaultBiomObjectBlock = setup.defaultBiomObjectBlock;
        buildingPrefab = setup.buildingPrefab;

        blockDataObjects = Resources.LoadAll<ResourcesObjectData>("Blocks").ToList();
        biomDataObjects = Resources.LoadAll<BiomObjectData>("Blocks").ToList();
        buildingDataObjects = Resources.LoadAll<BuildingObjectData>("Building").ToList();

        isMasterClient = SetupSetting.Instance.isMasterClient;
    }

    public IEnumerator GenerateChunk(Chunk chunk)
    {
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
                            if (!CheckPerlinLevel(tilePosition, resourcesObject.perlinSpeed,
                                resourcesObject.perlinLevel))
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
                    #region biom set for tcp

                    BiomType biomType = chunk.GetTileChunkData(tilePosition).biomTypeType;
                    if (!biomType.Equals(_defaultBiomObjectBlock.type))
                        for (int i = 0; i < biomDataObjects.Count; i++)
                        {
                            BiomObjectData block = biomDataObjects[i] as BiomObjectData;
                            if (biomType.Equals(block.type))
                            {
                                biomObjectBlockData = block;
                                break;
                            }
                        }

                    #endregion

                    #region resources set for tcp

                    ResourceType resourceType = chunk.GetTileChunkData(tilePosition).resourceType;
                    if ((int) resourceType != 0 && (int) resourceType < 100)
                        for (int i = 0; i < blockDataObjects.Count; i++)
                        {
                            ResourcesObjectData resourcesObject = blockDataObjects[i] as ResourcesObjectData;
                            if (resourceType.Equals(resourcesObject.type))
                            {
                                resResourcesObjectData = resourcesObject;
                                break;
                            }
                        }

                    #endregion

                    #region build

                    BuildType buildType = chunk.GetTileChunkData(tilePosition).buildType;
                    if (buildType != BuildType.NONE)
                    {
                        for (int i = 0; i < buildingDataObjects.Count; i++)
                        {
                            BuildingObjectData buildingObject = buildingDataObjects[i] as BuildingObjectData;
                            if (buildType.Equals(buildingObject.type))
                            {
                                int xAct = 0;
                                int yAct = 0;
                                if (buildType == BuildType.BASE)
                                {
                                    xAct = -1;
                                    yAct = -1;
                                }

                                Vector3Int vpos = new Vector3Int(tilePosition.x + xAct, tilePosition.y + yAct, 0);
                                Building b = Instantiate(buildingPrefab, vpos, Quaternion.identity);
                                b.SetBuildType(buildingObject);
                                b.BuildDone(true);
                                break;
                            }
                        }
                    }

                    #endregion
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