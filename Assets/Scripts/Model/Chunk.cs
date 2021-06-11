using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Tilemaps;

public class Chunk : MonoBehaviour
{
    public Vector3Int Position { get; private set; }
    public Vector3Int ChunkPosition { get; private set; }
    public TileChunk [,] baseTileType, tileChunkLayer;
    public Tilemap baseTileMap, layerTileMap;
    private BoxCollider2D chunkCollider;
    private bool isUnloading = false;

    // Start is called before the first frame update
    void Start()
    {
        Position = new Vector3Int((int)transform.position.x, (int)transform.position.y, 0);
        ChunkPosition = new Vector3Int(
            Position.x / GenerationManager.Instance.chunkSize,
            Position.y / GenerationManager.Instance.chunkSize, 0);
        baseTileType = new TileChunk[
            GenerationManager.Instance.chunkSize, 
            GenerationManager.Instance.chunkSize];
        tileChunkLayer = new TileChunk[
            GenerationManager.Instance.chunkSize, 
            GenerationManager.Instance.chunkSize];
        
        chunkCollider = GetComponent<BoxCollider2D>();
        chunkCollider.size = new Vector2(
            GenerationManager.Instance.chunkSize,
            GenerationManager.Instance.chunkSize);
        chunkCollider.offset = new Vector2(
            GenerationManager.Instance.chunkSize / 2,
            GenerationManager.Instance.chunkSize / 2);
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
        Vector3Int relativePosition = tilePosition - Position;
        
        if (isLayerTileMap)
            layerTileMap.SetTile(relativePosition, blockTile);
        else
            baseTileMap.SetTile(relativePosition, blockTile);
            
    }
    
    public void SetTileChunkData(Vector3Int position, TileType type)
    {
        if (isUnloading)
            return;

        Vector3Int relativePosition = position - Position;
        TileChunk tileChunk = new TileChunk();
        tileChunk.TileType = type;
        tileChunk.position = (Vector2Int) position;
        tileChunk.relativePosition = (Vector2Int) relativePosition;

        tileChunkLayer[relativePosition.x, relativePosition.y] = tileChunk;
    }
    
    public TileChunk GetTileChunkData(Vector3Int position)
    {
        if (isUnloading)
            return null;

        Vector3Int relativePosition = position - Position;
        return tileChunkLayer[relativePosition.x, relativePosition.y];
    }
    
}
