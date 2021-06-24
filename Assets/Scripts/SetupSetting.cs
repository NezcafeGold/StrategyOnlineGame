using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Serialization;
using UnityEngine.Tilemaps;

public class SetupSetting : Singleton<SetupSetting>
{
    
    public bool isMasterClient = false;
    
    [Header("ChunkLoadManager")]
    public GameObject chunkPrefab;
    public GameObject chunkRoot;
    public int horizontalChunkVisible = 120;
    public int verticalChunkVisible = 80;
    
    [Header("GenerationManager")]
    public  int chunkSize = 16;
    public int seed = 123456;
    public int worldWidth = 256;
    public int worldHeight = 256;
    [FormerlySerializedAs("defaultResourceResources")] public ResourcesObjectData defaultResourceObjectResourcesObject;
    [FormerlySerializedAs("defaultBiomBlock")] public BiomObjectData defaultBiomObjectBlock;

    [Header("DATA")]
    public string jsonPath = @"JSONFILES/ChunkMap.json";
    public string dictionaryPath = @"JSONFILES/ChunMap.dat";

    [Header("Building")] public List<BuildingObjectData> buildingList;
     public Building buildingPrefab;
     public BuildingManager buildingManager;

    [Header("TileMaps")]
    public Tilemap resourcesTileMap;
    public Tilemap biomsTileMap;
}
