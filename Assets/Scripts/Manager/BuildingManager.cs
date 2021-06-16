

using UnityEngine;

public class BuildingManager : MonoBehaviour
{
    public BuildingData.BuildingType type;

    private BuildingData buildPrefab;
    // Start is called before the first frame update
    void Start()
    {
        foreach (var b in SetupSetting.Instance.buildingList)
        {
            if (b.type == type)
            {
                buildPrefab = b;
            }
        }
    }


    public void PlaceBuild()
    {  
        GameObject b = Instantiate(buildPrefab.prefab, Camera.main.ScreenToWorldPoint(new Vector2(Screen.width / 2, Screen.height / 2)), Quaternion.identity);
        b.transform.position = new Vector3(b.transform.position.x, b.transform.position.y, -1);
        b.transform.SetParent(transform);
    }
}
