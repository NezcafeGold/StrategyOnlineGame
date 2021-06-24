using UnityEngine;
using Plane = UnityEngine.Plane;
using Vector3 = UnityEngine.Vector3;


[RequireComponent(typeof(BoxCollider2D))]
public class Building : MonoBehaviour
{
    [SerializeField] private SpriteRenderer zone;
    [SerializeField] private BoxCollider2D buildCollider;
    [SerializeField] private SpriteRenderer buildSprite;
    
    public bool canDrag = true;
    public bool canBuild = false; 

    private Plane dragPlane;
    private Vector3 offset;

    private Camera myMainCamera;

    void Start()
    {
        myMainCamera = Camera.main;
        SetColorForZone(Vector3Int.RoundToInt(zone.transform.position));
        buildSprite.color = new Color(1f,1f,1f,0.5f);
        
    }

    void OnMouseDown()
    {
        if (canDrag)
        {
            Vector3 pos = transform.position;
            dragPlane = new Plane(myMainCamera.transform.forward, pos);
            Ray camRay = myMainCamera.ScreenPointToRay(Input.mousePosition);

            float planeDist;
            dragPlane.Raycast(camRay, out planeDist);
            offset = pos - camRay.GetPoint(planeDist);
        }
    }

    void OnMouseDrag()
    {
        if (canDrag)
        {
            Ray camRay = myMainCamera.ScreenPointToRay(Input.mousePosition);

            float planeDist;
            dragPlane.Raycast(camRay, out planeDist);
            transform.position = camRay.GetPoint(planeDist) + offset;
            Vector3Int zonePos = Vector3Int.RoundToInt(transform.position);
            zone.transform.position = zonePos;

            SetColorForZone(zonePos);
        }
    }

    private void SetColorForZone(Vector3Int zonePos)
    {
        if (CanBuild(zonePos))
        {
            canBuild = true;
            zone.color = new Color(0, 1, 0, 0.5f);
        }
        else
        {
            canBuild = false;
            zone.color = new Color(1, 0, 0, 0.5f);
        }
    }

    private void OnMouseUp()
    {
        if (canDrag)
        {
            Vector3Int pos = Vector3Int.RoundToInt(transform.position);
            gameObject.transform.position = pos;
            zone.transform.position = pos;
        }
    }

    //     * * * *
    //     * ? ! *       ! - this is center of zone in 4x4                           2 1
    //     * ? ? *       ? - place for build                                         3 4
    //     * * * *       Building can be place only on NONE tiles (DIRT for default)
    private bool CanBuild(Vector3Int pos)
    {
        if (!(pos.x - 1 > 0 && pos.y - 1 > 0))
            return false;

        Vector3Int pos1 = pos;
        Vector3Int pos2 = new Vector3Int(pos.x - 1, pos.y, pos.z);
        Vector3Int pos3 = new Vector3Int(pos.x - 1, pos.y - 1, pos.z);
        Vector3Int pos4 = new Vector3Int(pos.x, pos.y - 1, pos.z);

        Chunk ch1 = ChunkLoadManager.Instance.GetChunk(pos);
        Chunk ch3 = ChunkLoadManager.Instance.GetChunk(pos3);
        Chunk ch2 = ch1;
        Chunk ch4 = ch3;
        if (!ch1.Equals(ch3))
        {
            ch2 = ChunkLoadManager.Instance.GetChunk(pos2);
            ch4 = ChunkLoadManager.Instance.GetChunk(pos4);
        }

        return ch1.GetTileChunkData(pos1).resourceType.Equals(ResourceType.NONE) &&
               ch2.GetTileChunkData(pos2).resourceType.Equals(ResourceType.NONE) &&
               ch3.GetTileChunkData(pos3).resourceType.Equals(ResourceType.NONE) &&
               ch4.GetTileChunkData(pos4).resourceType.Equals(ResourceType.NONE);
    }

    public void BuildDone()
    {
        canDrag = false;
        buildSprite.color = new Color(1,1,1,1f);
        zone.color = new Color(1, 1, 1, 0.5f);
    }
    
}