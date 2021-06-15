using UnityEngine;
using UnityEngine.EventSystems;


[RequireComponent(typeof(BoxCollider2D))]
public class DragToBuild : MonoBehaviour
{

    [SerializeField]private SpriteRenderer zone;
    public bool dragDisabled = true;

    private Plane dragPlane;
    private Vector3 offset;

    private Camera myMainCamera;

    void Start()
    {
        myMainCamera = Camera.main; 
    }

    void OnMouseDown()
    {
        if (dragDisabled)
        {
            dragPlane = new Plane(myMainCamera.transform.forward, transform.position);
            Ray camRay = myMainCamera.ScreenPointToRay(Input.mousePosition);

            float planeDist;
            dragPlane.Raycast(camRay, out planeDist);
            offset = transform.position - camRay.GetPoint(planeDist);
        }
    }

    void OnMouseDrag()
    {
        if (dragDisabled)
        {
            Ray camRay = myMainCamera.ScreenPointToRay(Input.mousePosition);

            float planeDist;
            dragPlane.Raycast(camRay, out planeDist);
            transform.position = camRay.GetPoint(planeDist) + offset;
        }
    }


}