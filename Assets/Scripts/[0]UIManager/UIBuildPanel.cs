using System;
using System.Collections.Generic;
using System.Linq;
using FantasyRPG;
using UnityEngine;
using UnityEngine.Serialization;


public class UIBuildPanel : PanelShop
{
    private List<BuildingObjectData> buildingObjectDatas;
    private BuildingManager buildingManager;
    [FormerlySerializedAs("building")] [SerializeField] private BuildingByMenu buildingByMenu;
    private Camera cam;
    private void Start()
    {
        buildingObjectDatas = Resources.LoadAll<BuildingObjectData>("Building").ToList();
        cam = Camera.main;
        buildingManager = SetupSetting.Instance.buildingManager;
    }

    public void SelectAndBuild(UIBuildItem uiBuildItem)
    {
        foreach (var bod in buildingObjectDatas)
        {
            if (bod.type.Equals(uiBuildItem.buildType))
            {
                Vector3 pos = cam.transform.position;
                BuildingByMenu b =Instantiate(buildingByMenu, new Vector3(pos.x, pos.y, 0), Quaternion.identity, buildingManager.transform);
                b.SetBuildType(bod);
                break;
            }
        }
        Close();
    }
}
