using System;
using UnityEngine;
using UnityEngine.Tilemaps;

[Serializable]
public class Chunk : MonoBehaviour
{
    public SerializableVector2Int Position;
    public SerializableVector2Int ChunkPosition;
    public TileChunk[,] baseTileType, tileChunkLayer;

    [SerializeField] private Tilemap baseTileMap, layerTileMap;
    private BoxCollider2D chunkCollider;
    private bool isUnloading = false;

    // Start is called before the first frame update
    void Start()
    {
        Position = new Vector2Int((int) transform.position.x, (int) transform.position.y);
        ChunkPosition = new Vector2Int(
            Position.x / SetupSetting.Instance.chunkSize,
            Position.y / SetupSetting.Instance.chunkSize);
        baseTileType = new TileChunk[
            SetupSetting.Instance.chunkSize,
            SetupSetting.Instance.chunkSize];
        tileChunkLayer = new TileChunk[
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
        if (SetupSetting.Instance.isMasterClient)
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
        Vector3Int relativePosition = tilePosition - new Vector3Int(Position.x, Position.y, 0);

        if (isLayerTileMap)
            layerTileMap.SetTile(relativePosition, blockTile);
        else
            baseTileMap.SetTile(relativePosition, blockTile);
    }

    public void SetTileChunkData(Vector3Int position, TileType type)
    {
        if (isUnloading)
            return;

        Vector3Int relativePosition = position - new Vector3Int(Position.x, Position.y, 0);
        TileChunk tileChunk = new TileChunk();
        tileChunk.TileType = type;
        tileChunk.pos = (Vector2Int) position;
        tileChunk.relPos = (Vector2Int) relativePosition;

        tileChunkLayer[relativePosition.x, relativePosition.y] = tileChunk;
    }

    public TileChunk GetTileChunkData(Vector3Int position)
    {
        if (isUnloading)
            return null;

        Vector3Int relativePosition = position - new Vector3Int(Position.x, Position.y, 0);
        return tileChunkLayer[relativePosition.x, relativePosition.y];
    }

//    public struct TileData
//    {
//        public TileType type;
//        public SerializableVector2Int pos;
//
//        public SerializableVector2Int relPos;
////        public int xPos;
////        public int yPos;
////        public int xRelPos;
////        public int yRelPos;
//    }
}