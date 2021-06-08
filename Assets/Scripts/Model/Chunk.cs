using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Tilemaps;

public class Chunk : MonoBehaviour
{
    public Vector3Int Position { get; private set; }
    public Vector3Int ChunkPosition { get; private set; }
    public TileType[,] tileType;
    public Tilemap tileMap;
    private BoxCollider2D chunkCollider;
    private bool isUnloading = false;
    public enum TileType
    {
        DIRT,
        WOOD,
        STONE
    }
    // Start is called before the first frame update
    void Start()
    {
        Position = new Vector3Int((int)transform.position.x, (int)transform.position.y, 0);
        ChunkPosition = new Vector3Int(
            Position.x / GenerationManager.Instance.chunkSize,
            Position.y / GenerationManager.Instance.chunkSize, 0);
        tileType = new TileType[
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

    public void SetChunkTile(Vector3Int tilePosition, Tile blockTile)
    {
        if (isUnloading)
            return;
        Vector3Int relativePosition = tilePosition - Position;
        tileMap.SetTile(relativePosition, blockTile);
    }
}
