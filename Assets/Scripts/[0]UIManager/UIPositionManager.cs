using System;
using TMPro;
using UnityEngine;
using UnityEngine.EventSystems;

public class UIPositionManager : MonoBehaviour
{
    [SerializeField] private TextMeshProUGUI position;

    private void Update()
    {
        ShowChunkPosition();
    }

    private void ShowChunkPosition()
    {
        try
        {
            if (Input.GetMouseButtonDown(0))
            {
                if (!EventSystem.current.IsPointerOverGameObject())
                {
                    RaycastHit2D ray = Physics2D.GetRayIntersection(Camera.main.ScreenPointToRay(Input.mousePosition),
                        100, LayerMask.GetMask("Chunk"));
                    if (ray.collider == null || !ray.collider.gameObject.CompareTag("Chunk"))
                        return;
                    Vector3Int mousePos = Vector3Int.FloorToInt(Camera.main.ScreenToWorldPoint(Input.mousePosition));
                    Chunk chunk = ChunkLoadManager.Instance.GetChunk(mousePos);
                    if (chunk != null)
                    {
                        TileChunk tileChunk = chunk.GetTileChunkData(mousePos);
                        if (tileChunk == null)
                            Debug.Log("Tile is empty");
                        else
                            position.SetText(tileChunk.resourceType + "\n " + mousePos + "\n Chunk " +
                                             chunk.chunkData.Position);
                    }
                }
            }
        }
        catch (Exception e)
        {
            Debug.Log("Cant show position" + e);
        }
    }
}