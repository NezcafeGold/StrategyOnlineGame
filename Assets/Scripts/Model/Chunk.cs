using System;
using System.Collections;
using System.Runtime.CompilerServices;
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
    private bool isSet = false;

    // Start is called before the first frame update
    void Start()
    {
        int chunkSize = SetupSetting.Instance.chunkSize;
        resourcesTilemap = SetupSetting.Instance.resourcesTileMap;
        biomsTilemap = SetupSetting.Instance.biomsTileMap;

        chunkData = new ChunkData();
        chunkData.Position = new Vector2Int((int) transform.position.x, (int) transform.position.y);
        chunkData.ChunkPosition = new Vector2Int(
            chunkData.Position.x / chunkSize,
            chunkData.Position.y / chunkSize);
        chunkData.tileChunkLayer = new TileChunk[
            SetupSetting.Instance.chunkSize,
            SetupSetting.Instance.chunkSize];

        chunkCollider = GetComponent<BoxCollider2D>();
        chunkCollider.size = new Vector2(
            chunkSize,
            chunkSize);
        chunkCollider.offset = new Vector2(
            chunkSize / 2,
            chunkSize / 2);
        LoadChunk();
    }

    private void Update()
    {
        if (PlayerData.GetInstance().ChunkMap.ContainsKey(chunkData.Position))
        {
            SetData();
            isSet = true;
        }
    }

    private void LoadChunk()
    {
        if (SetupSetting.Instance.isMasterClient)
            StartCoroutine(GenerationManager.Instance.GenerateChunk(this));
        else
            LoadDataFromTCP();
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

    private void LoadDataFromTCP()
    {
        if (!PlayerData.GetInstance().ChunkMap.ContainsKey(chunkData.Position) && !isSendToTCP)
        {
            TCPClient.Instance.SendMessageTCP(new Packet(Packet.SegmentID.GET_CHUNK_ID,
                Packet.StatusCode.OK_CODE, Packet.Body.OfInt("x", chunkData.Position.x),
                Packet.Body.OfInt("y", chunkData.Position.y)).ToString());
            isSendToTCP = true;
        }
    }

    public void SetData()
    {
        if (isSet) return;
        ChunkData ch = PlayerData.GetInstance().ChunkMap[chunkData.Position];
        chunkData.Position = ch.Position;
        chunkData.tileChunkLayer = ch.tileChunkLayer;
        StartCoroutine(GenerationManager.Instance.GenerateChunk(this));
        isSet = true;
        Debug.Log("CHUNK FOUND" + ch.Position.x + " " + ch.Position.y);
    }
}