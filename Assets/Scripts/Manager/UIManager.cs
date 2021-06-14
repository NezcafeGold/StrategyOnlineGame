using TMPro;
using UnityEngine;
using UnityEngine.EventSystems;

public class UIManager : MonoBehaviour
{

    [SerializeField] private TextMeshProUGUI position;
    
    private void Update()
    {
        ShowChunkPosition();
    }

    private void ShowChunkPosition()
    {
        if (Input.GetMouseButtonDown(0) && !EventSystem.current.IsPointerOverGameObject())
        {
            RaycastHit2D ray = Physics2D.GetRayIntersection(Camera.main.ScreenPointToRay(Input.mousePosition));

            if (!ray.collider.gameObject.CompareTag("Chunk"))
                return;

            Vector3Int mousePos = Vector3Int.FloorToInt(Camera.main.ScreenToWorldPoint(Input.mousePosition));
            Chunk chunk = ChunkLoadManager.Instance.GetChunk(mousePos);
            if (chunk != null)
            {
                TileChunk tileChunk = chunk.GetTileChunkData(mousePos);
                position.SetText(tileChunk.tileType + "\n "+ mousePos + "\n Chunk " + chunk.chunkData.Position );
          
            }
        }
    }
}