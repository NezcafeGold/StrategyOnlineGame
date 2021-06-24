using System.Collections.Generic;
using Model.BuildData;
using UnityEngine;

public class BuildingManager : Singleton<BuildingManager>
{
    private Dictionary<Vector2Int, BuildData> buildDict;
    [SerializeField] private GameObject playerBase;


    private void Start()
    {
        buildDict = new Dictionary<Vector2Int, BuildData>();
        Vector2Int basePos = PlayerData.GetInstance().baseData.Position;
        GameObject g = Instantiate(playerBase, transform, true);
        g.transform.position = new Vector3(basePos.x-1, basePos.y-1, 0);

        BaseData baseData = new BaseData();
        baseData.Position = basePos;
        //baseData.OwnerId = 
    }
}