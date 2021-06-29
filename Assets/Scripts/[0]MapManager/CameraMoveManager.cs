using UnityEngine;

public class CameraMoveManager : Singleton<CameraMoveManager>
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
    //private int speed = 5;

    void Awake()
    {
        cam = Camera.main;
        camSize = cam.orthographicSize;
        maxX = SetupSetting.Instance.worldWidth;
        maxY = SetupSetting.Instance.worldHeight;
        Vector2Int v = PlayerData.GetInstance().baseData.Position;
        cam.transform.position = new Vector3Int(v.x, v.y, -10);
        dragVector = new Vector3Int(v.x, v.y, -10);
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
#if UNITY_STANDALONE_WIN
        if (Input.GetKeyDown(keyCode) && !isDrag)
        {
            mouseClickPos = cam.ScreenToWorldPoint(Input.mousePosition);
            isDrag = true;
        }
#endif

#if UNITY_ANDROID

        if (Input.touchCount > 0 && Input.GetTouch(0).phase == TouchPhase.Moved)
        {
            Vector2 touchDeltaPosition = Input.GetTouch(0).deltaPosition;
            transform.Translate(-touchDeltaPosition.x * speed, -touchDeltaPosition.y * speed, 0);
            isDrag = true;
        }


#endif


        if (isDrag)
        {
            mouseCurrentPos = cam.ScreenToWorldPoint(Input.mousePosition);
            var distance = mouseCurrentPos - mouseClickPos;

            dragVector +=
                new Vector3(-distance.x, -distance.y, 0);

            CheckBorders();
        }

#if UNITY_STANDALONE_WIN
        if (Input.GetKeyUp(keyCode))
            isDrag = false;

#endif

#if UNITY_ANDROID
        if (Input.touchCount > 0 && Input.GetTouch(0).phase == TouchPhase.Ended)
            isDrag = false;
#endif
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

    public void SetPosition(Vector2Int position)
    {
        cam.transform.position = new Vector3Int(position.x, position.y, -10);
        dragVector = new Vector3Int(position.x, position.y, -10);
    }
}