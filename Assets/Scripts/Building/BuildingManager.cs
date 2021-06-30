using System;
using System.Collections;
using System.Collections.Generic;
using Model.BuildData;
using UnityEngine;

public class BuildingManager : Singleton<BuildingManager>
{
    [SerializeField] private GameObject playerBase;
    private PlayerData playerData;
    public Queue<BuildData> AddBuildingQueue = new Queue<BuildData>();
    public Dictionary<SerializableVector2Int, BuildData> BuildMap = new Dictionary<SerializableVector2Int, BuildData>();

    private void Start()
    {
        playerData = PlayerData.GetInstance();
        //  Vector2Int basePos = playerData.baseData.Position;
        // GameObject g = Instantiate(playerBase, transform, true);
        // g.transform.position = new Vector3(basePos.x - 1, basePos.y - 1, 0);
        //  BaseData baseData = new BaseData();
        //  baseData.Position = basePos;
    }

    private void Update()
    {
        if(AddBuildingQueue.Count>0)
            StartCoroutine(UpdateBuildMap(AddBuildingQueue.Dequeue(), true));
    }

    public IEnumerator UpdateBuildMap(BuildData buildData, bool isLoaded = false)
    {
        if (BuildMap.ContainsKey(buildData.Position))
            BuildMap.Remove(buildData.Position);
        BuildMap.Add(buildData.Position, buildData);
        Messenger.Broadcast(GameEvent.UPDATE_BUILD_INFO, buildData);
        yield return null;

        if (!isLoaded)
            SendToTcp(buildData);
        yield return null;
    }

    public void SendToTcp(BuildData buildData) //КОРУТИНУ СДЕЛАЙ
    {
        Vector2Int v = buildData.Position;
        try
        {
            int segmentId = 9999;
            switch (buildData.BuildType)
            {
                case BuildType.FARM:
                    segmentId = Packet.SegmentID.SET_FARM_ID;
                    break;
                
            }

            HandleTCP(segmentId, v);
        }
        catch (Exception e)
        {
            Debug.Log("Cant Send To Tcp Info About Building " + e, this);
        }
    }

    public void AskDataTcp(BuildData buildData)
    {
        Vector2Int v = buildData.Position;
        try
        {
            int segmentId = Packet.SegmentID.GET_INFO_BASE_ID;
            switch (buildData.BuildType)
            {
                case BuildType.FARM:
                    segmentId = Packet.SegmentID.GET_INFO_FARM_ID;
                    break;
                case BuildType.BASE:
                    segmentId = Packet.SegmentID.GET_INFO_BASE_ID;
                    break;
            }

            HandleTCP(segmentId, v);
        }
        catch (Exception e)
        {
            Debug.Log("Cant Ask Data From Tcp Info About Building " + e, this);
        }
    }

    private void HandleTCP(int segmentId, Vector2Int v)
    {
        TCPClient.Instance.SendMessageTCP(new Packet(segmentId,
                Packet.StatusCode.OK_CODE, Packet.Body.OfInt("x", v.x), Packet.Body.OfInt("y", v.y))
            .ToString());
    }
}