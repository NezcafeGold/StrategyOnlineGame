using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SetupSetting : Singleton<SetupSetting>
{
    
    public bool isMasterClient = false;
    
    [Header("ChunkLoadManager")]
    public GameObject chunkPrefab;
    public GameObject chunkRoot;
    public int horizontalChunkVisible = 120;
    public int verticalChunkVisible = 80;
    
    [Header("GenerationManager")]
    public int chunkSize = 40;
    public int seed = 123456;
    public int worldWidth = 256;
    public int worldHeight = 256;
    public BlockData defaultBlock;
}
