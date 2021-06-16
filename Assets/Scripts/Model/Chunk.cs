using System;
using UnityEngine;
using UnityEngine.Tilemaps;

[Serializable]
public class Chunk : MonoBehaviour
{

    public ChunkData chunkData;
    //public SerializableVector2Int Position;
    [SerializeField] private Tilemap baseTileMap, layerTileMap;
    private BoxCollider2D chunkCollider;
    private bool isUnloading = false;

    // Start is called before the first frame update
    void Start()
    {
        chunkData = new ChunkData();
        chunkData.Position = new Vector2Int((int) transform.position.x, (int) transform.position.y);
        chunkData.ChunkPosition = new Vector2Int(
            chunkData.Position.x / SetupSetting.Instance.chunkSize,
            chunkData.Position.y / SetupSetting.Instance.chunkSize);
        chunkData.tileChunkLayer = new TileChunk[
            SetupSetting.Instance.chunkSize,
            SetupSetting.Instance.chunkSize];

        chunkCollider = GetComponent<BoxCollider2D>();
        chunkCollider.size = new Vector2(
            SetupSetting.Instance.chunkSize,
            SetupSetting.Instance.chunkSize);
        chunkCollider.offset = new Vector2(
            SetupSetting.Instance.chunkSize / 2,
            SetupSetting.Instance.chunkSize / 2);
        LoadChunk();
    }

    private void LoadChunk()
    {
            StartCoroutine(GenerationManager.Instance.GenerateChunk(this));
    }

    public void UnloadChunk()
    {
        Destroy(gameObject);
        isUnloading = true;
    }

    public void SetChunkTile(Vector3Int tilePosition, Tile blockTile, bool isLayerTileMap = false)
    {
        if (isUnloading)
            return;
        Vector3Int relativePosition = tilePosition - new Vector3Int(chunkData.Position.x, chunkData.Position.y, 0);

        if (isLayerTileMap)
            layerTileMap.SetTile(relativePosition, blockTile);
        else
            baseTileMap.SetTile(relativePosition, blockTile);
    }

    public void SetTileChunkData(Vector3Int position, TileType type)
    {
        if (isUnloading)
            return;

        Vector3Int relativePosition = position - new Vector3Int(chunkData.Position.x, chunkData.Position.y, 0);
        TileChunk tileChunk = new TileChunk();
        tileChunk.tileType = type;
        tileChunk.type = (int) type;
        tileChunk.pos = (Vector2Int) position;
        tileChunk.relPos = (Vector2Int) relativePosition;

        chunkData.tileChunkLayer[relativePosition.x, relativePosition.y] = tileChunk;
    }

    public TileChunk GetTileChunkData(Vector3Int position)
    {
        if (isUnloading)
            return null;

        Vector3Int relativePosition = position - new Vector3Int(chunkData.Position.x, chunkData.Position.y, 0);
        return chunkData.tileChunkLayer[relativePosition.x, relativePosition.y];
    }    
}