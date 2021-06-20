using System;
using System.Collections;
using UnityEngine;
using UnityEngine.Tilemaps;

[Serializable]
public class Chunk : MonoBehaviour
{
    public ChunkData chunkData;

    //public SerializableVector2Int Position;
    [SerializeField] private Tilemap resourcesTilemap, biomsTilemap;
    private BoxCollider2D chunkCollider;
    private bool isUnloading = false;
    private bool isSendToTCP = false;

    // Start is called before the first frame update
    void Start()
    {
        resourcesTilemap = SetupSetting.Instance.resourcesTileMap;
        biomsTilemap = SetupSetting.Instance.biomsTileMap;

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
        if (SetupSetting.Instance.isMasterClient)
            StartCoroutine(GenerationManager.Instance.GenerateChunk(this));
        else
            StartCoroutine(LoadDataFromTCP());
    }


    public IEnumerator UnloadChunkCor()
    {
        for (int x = chunkData.Position.x; x < chunkData.Position.x + SetupSetting.Instance.chunkSize; x++)
        {
            for (int y = chunkData.Position.y; y < chunkData.Position.y + SetupSetting.Instance.chunkSize; y++)
            {
                resourcesTilemap.SetTile(new Vector3Int(x, y, 0), null);
                biomsTilemap.SetTile(new Vector3Int(x, y, 0), null);
            }
        }

        //BigTileMapManager.Instance.UnloadTiles(chunkData.Position);
        yield return null;
        isUnloading = true;
        Destroy(gameObject);
    }

    public void SetChunkTile(Vector3Int tilePosition, TileBase blockTile, bool isLayerTileMap = false)
    {
        if (isUnloading)
            return;
        Vector3Int relativePosition = tilePosition - new Vector3Int(chunkData.Position.x, chunkData.Position.y, 0);

        if (isLayerTileMap)
        {
            resourcesTilemap.SetTile(tilePosition, blockTile);
        }

        else
            biomsTilemap.SetTile(tilePosition, blockTile);
    }

    public void SetTileChunkData(Vector3Int position, ResourceType resType, BiomType biomType)
    {
        if (isUnloading)
            return;

        Vector3Int relativePosition = position - new Vector3Int(chunkData.Position.x, chunkData.Position.y, 0);
        TileChunk tileChunk = new TileChunk();

        tileChunk.pos = (Vector2Int) position;
        tileChunk.relPos = (Vector2Int) relativePosition;

        tileChunk.rtype = (int) resType;
        tileChunk.btype = (int) biomType;
        tileChunk.resourceType = resType;
        tileChunk.biomTypeType = biomType;
        chunkData.tileChunkLayer[relativePosition.x, relativePosition.y] = tileChunk;
    }

    public TileChunk GetTileChunkData(Vector3Int position)
    {
        if (isUnloading)
            return null;

        Vector3Int relativePosition = position - new Vector3Int(chunkData.Position.x, chunkData.Position.y, 0);
        return chunkData.tileChunkLayer[relativePosition.x, relativePosition.y];
    }

    private IEnumerator LoadDataFromTCP()
    {
        while (true)
        {
            if (!PlayerData.Instance.ChunkMap.ContainsKey(chunkData.Position) && !isSendToTCP)
            {
                TCPClient.Instance.SendMessageTCP(new Packet(Packet.SegmentID.GET_CHUNK_ID,
                    Packet.StatusCode.OK_CODE, Packet.Body.OfInt("x", chunkData.Position.x),
                    Packet.Body.OfInt("y", chunkData.Position.y)).ToString());
                isSendToTCP = true;
            }
            else
            {
                if (!PlayerData.Instance.ChunkMap.ContainsKey(chunkData.Position))
                    continue;
                ChunkData ch = PlayerData.Instance.ChunkMap[chunkData.Position];
                chunkData.Position = ch.Position;
                chunkData.tileChunkLayer = ch.tileChunkLayer;
                yield return null;
                StartCoroutine(GenerationManager.Instance.GenerateChunk(this));
                
                Debug.Log("CHUNK FOUND" + ch.Position.x + " " + ch.Position.y);
                break;
            }

            yield return null;
        }
    }
}