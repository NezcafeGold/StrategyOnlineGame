using System.Collections.Generic;
using Model.BuildData;
using UnityEngine;

public class BuildingManager : Singleton<BuildingManager>
{
    [SerializeField] private GameObject playerBase;
    private PlayerData playerData;

    private void Start()
    {
        playerData = PlayerData.GetInstance();
        Vector2Int basePos = playerData.baseData.Position;
        GameObject g = Instantiate(playerBase, transform, true);
        g.transform.position = new Vector3(basePos.x - 1, basePos.y - 1, 0);
        BaseData baseData = new BaseData();
        baseData.Position = basePos;
    }

    public void AddNewBuilding(BuildData buildData, bool isLoaded = false)
    {
        if (playerData.BuildMap.ContainsKey(buildData.Position))
            playerData.BuildMap.Remove(buildData.Position);
        playerData.BuildMap.Add(buildData.Position, buildData);

        if (!isLoaded)
            SendToTcp(buildData);
    }

    private void SendToTcp(BuildData buildData) //КОРУТИНУ СДЕЛАЙ
    {
        Vector2Int v = buildData.Position;

        switch (buildData.BuildType)
        {
            case BuildType.FARM:
                TCPClient.Instance.SendMessageTCP(new Packet(Packet.SegmentID.SET_FARM_ID,
                    Packet.StatusCode.OK_CODE, Packet.Body.OfInt("x", v.x), Packet.Body.OfInt("y", v.y)).ToString());
                break;
        }
    }
}