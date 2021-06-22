
using UnityEngine;

public class CameraMoveManager : MonoBehaviour
{
    [SerializeField] private float minZoom = 10f;
    [SerializeField] private float maxZoom = 50f;
    [SerializeField] private KeyCode keyCode = KeyCode.Mouse2; //колесико

    private Camera cam;
    private float camSize;
    private float zoomDelta;
    private Vector2 mouseClickPos;
    private Vector2 mouseCurrentPos;
    private bool isDrag = false;
    private Vector3 dragVector;
    private int maxX, maxY;

    void Awake()
    {
        cam = Camera.main;
        camSize = cam.orthographicSize;
        maxX = SetupSetting.Instance.worldWidth;
        maxY = SetupSetting.Instance.worldHeight;        
        Vector2Int v = PlayerData.GetInstance().baseData.Position;
        cam.transform.position = new Vector3Int(v.x, v.x, -10);
        dragVector = new Vector3Int(v.x, v.x, -10);
    }

    void LateUpdate()
    {
        CamZoom();
        CamDrag();
    }

    private void CamZoom()
    {
        zoomDelta = camSize - Input.mouseScrollDelta.y;
        camSize = Mathf.Clamp(zoomDelta, minZoom, maxZoom);
        cam.orthographicSize = camSize;
        CheckBorders();
    }

    private void CamDrag()
    {
        if (Input.GetKeyDown(keyCode) && !isDrag)
        {
            mouseClickPos = cam.ScreenToWorldPoint(Input.mousePosition);
            isDrag = true;
        }


        if (isDrag)
        {
            mouseCurrentPos = cam.ScreenToWorldPoint(Input.mousePosition);
            var distance = mouseCurrentPos - mouseClickPos;

            dragVector +=
                new Vector3(-distance.x, -distance.y, 0);

            CheckBorders();
        }

        if (Input.GetKeyUp(keyCode))
            isDrag = false;
    }

    private void CheckBorders()
    {
        float horizExtent = camSize * Screen.width / Screen.height;
        float vertExtent = cam.orthographicSize;

        if (0 + horizExtent > dragVector.x) dragVector.x = 0 + horizExtent;
        if (maxX - horizExtent < dragVector.x) dragVector.x = maxX - horizExtent;
        if (0 + vertExtent > dragVector.y) dragVector.y = 0 + vertExtent;
        if (maxY - vertExtent < dragVector.y) dragVector.y = maxY - vertExtent;
        transform.position = new Vector3(dragVector.x, dragVector.y, transform.position.z);
    }
}