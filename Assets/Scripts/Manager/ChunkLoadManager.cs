
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ChunkLoadManager : Singleton<ChunkLoadManager>
{
    [SerializeField] private bool isMasterClient = false;
    [SerializeField] private GameObject chunkPrefab;
    [SerializeField] private GameObject chunkRoot;
    [SerializeField] private int horizontalChunkVisible = 8;
    [SerializeField] private int verticalChunkVisible = 6;
    
    
    private BoxCollider2D boxColl;
    private bool isUpdatingChunks = false;
    private Rect loadBoundaries;

    public List<Chunk> chunks;


    void Start()
    {
        chunks = new List<Chunk>();
        boxColl = Camera.main.GetComponent<BoxCollider2D>();
        StartCoroutine(LoadChunks());
        StartCoroutine(UnloadChunks());
    }

    private IEnumerator LoadChunks()
    {
        while (true)
        {
            isUpdatingChunks = true;
            yield return StartCoroutine(PerformLoadChunks());
            isUpdatingChunks = false;
            yield return null;
        }
    }

    private IEnumerator UnloadChunks()
    {
        while (true)
        {
            if (!isUpdatingChunks)
                yield return StartCoroutine(PerformUnloadChunks());
            yield return null;
        }
    }

    private IEnumerator PerformUnloadChunks()
    {
        loadBoundaries = GetChunkLoadBounds();
        List<Chunk> chunksToUnload = new List<Chunk>();
        foreach (Transform child in chunkRoot.transform)
        {
            Chunk chunk = child.GetComponent<Chunk>();
            if (chunk != null)
            {
                if (!loadBoundaries.Contains(chunk.ChunkPosition))
                    chunksToUnload.Add(chunk);
            }
        }

        foreach (Chunk chunk in chunksToUnload)
        {
            while (isUpdatingChunks)
                yield return null;

            if (chunk != null)
                chunk.UnloadChunk();
            yield return null;
        }
    }

    private IEnumerator PerformLoadChunks()
    {
        //UpdateBounds();
        loadBoundaries = GetChunkLoadBounds();
        List<Chunk> chunksToLoad = new List<Chunk>();
        for (int h = (int) loadBoundaries.xMax; h >= (int) loadBoundaries.xMin; h--)
        {
            for (int v = (int) loadBoundaries.yMax; v >= (int) loadBoundaries.yMin; v--)
            {
                if ((h < 0 || h >= GenerationManager.Instance.worldWidth / GenerationManager.Instance.chunkSize) ||
                    (v < 0 || v >= GenerationManager.Instance.worldHeight / GenerationManager.Instance.chunkSize))
                    continue;
                Vector3Int chunkPosition = new Vector3Int(h, v, 0);
                Vector3Int worldPosition = new Vector3Int(
                    h * GenerationManager.Instance.chunkSize,
                    v * GenerationManager.Instance.chunkSize, 0);

                if (loadBoundaries.Contains(chunkPosition) && !GetChunk(worldPosition))
                {
                    Chunk ch = Instantiate(chunkPrefab, worldPosition, Quaternion.identity, chunkRoot.transform)
                        .GetComponent<Chunk>();
                    chunksToLoad.Add(ch);
                    yield return null;
                }
            }
        }

        if (chunksToLoad.Count > 0)
            chunks = chunksToLoad;
    }


    // Возвращает чанк в заданной позиции по бохколлайдеру))
    public Chunk GetChunk(Vector3Int position)
    {
        RaycastHit2D hit;
        float fl = Vector3Int.zero == position ? 0 : 0.5f;
        hit = Physics2D.Raycast(
            new Vector2(position.x + fl, position.y + fl),
            Vector2.zero, 0f);

        return hit ? hit.collider.GetComponent<Chunk>() : null;
    }

    private void UpdateBounds()
    {
        Debug.Log(boxColl.bounds.max.x.ToString());
    }

    private Rect GetChunkLoadBounds()
    {
        Vector3 regionStart = Camera.main.transform.position +
                              Vector3.left * horizontalChunkVisible +
                              Vector3.down * verticalChunkVisible;
        Vector3 regionEnd = Camera.main.transform.position +
                            Vector3.right * horizontalChunkVisible +
                            Vector3.up * verticalChunkVisible;

        int regionStartX = (int) regionStart.x / GenerationManager.Instance.chunkSize;
        int regionStartY = (int) regionStart.y / GenerationManager.Instance.chunkSize;
        int regionEndX = ((int) regionEnd.x + GenerationManager.Instance.chunkSize) /
                         GenerationManager.Instance.chunkSize;
        int regionEndY = ((int) regionEnd.y + GenerationManager.Instance.chunkSize) /
                         GenerationManager.Instance.chunkSize;
        Rect loadBoundaries =
            new Rect(regionStartX, regionStartY, regionEndX - regionStartX, regionEndY - regionStartY);

        return loadBoundaries;
    }
}