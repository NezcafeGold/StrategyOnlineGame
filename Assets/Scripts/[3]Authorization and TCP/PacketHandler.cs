using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Threading;
using Leguar.TotalJSON;
using Model.BuildData;
using UnityEngine;

public class PacketHandler
{
    private bool CheckForCorrectPacket()
    {
        //TODO: Add method
        return true;
    }

    public void Handle(string serverMessage)
    {
        try
        {
            HandleHeader(JSON.ParseString(serverMessage).GetJSON(Packet.PacketKey.HEADER), serverMessage);
        }
        catch (Exception e)
        {
            Debug.Log("Error with parse " + serverMessage + " " + e);
        }
    }


    private void HandleHeader(JSON header, string serverMessage)
    {
        if (header.GetInt(Packet.PacketKey.STATUS_CODE).Equals(Packet.StatusCode.OK_CODE))
        {
            HandleByID(header.GetInt(Packet.PacketKey.ID), serverMessage);
        }
    }

    private void HandleByID(int id, string serverMessage)
    {
        try
        {
            switch (id)
            {
                case Packet.SegmentID.AUTHORIZATION_ID:
                    break;

                case Packet.SegmentID.GET_TILE_ID:

                    break;

                case Packet.SegmentID.GET_USER_ID:
                    HandleUserData(serverMessage);
                    break;
                case Packet.SegmentID.GET_CHUNK_ID:
                    HandleChunk(serverMessage);
                    break;
                case Packet.SegmentID.GET_UNITS_ID:

                    break;
                case Packet.SegmentID.GET_DATA_MAP_ID:

                    break;
                case Packet.SegmentID.GET_INVENTORY_ID:

                    break;
                case Packet.SegmentID.GET_RESOURCES_ID:

                    break;
                default:

                    break;
            }
        }
        catch (Exception e)
        {
            Debug.Log(e);
        }
    }

    private void HandleUserData(string serverMessage)
    {
        try
        {
            PlayerData pd = PlayerData.GetInstance();
            JSON bodyJs = JSON.ParseString(serverMessage).GetJSON(Packet.PacketKey.BODY);
            JSON userJs = bodyJs.GetJSON("user");
            pd.Nickname = userJs.GetString("nickname");
            JSON playerJs = userJs.GetJSON("player");
            pd.Level = playerJs.GetInt("level");
            pd.Experience = playerJs.GetInt("experience");
            JSON resources = playerJs.GetJSON("resources");
            Dictionary<ResourceType, int> resourcesDictionary = new Dictionary<ResourceType, int>();
            resourcesDictionary.Add(ResourceType.COAL, resources.GetInt("coal"));
            resourcesDictionary.Add(ResourceType.SULFUR, resources.GetInt("sulfur"));
            resourcesDictionary.Add(ResourceType.COPPER, resources.GetInt("copper"));
            resourcesDictionary.Add(ResourceType.IRON, resources.GetInt("iron"));
            resourcesDictionary.Add(ResourceType.GOLD, resources.GetInt("gold"));
            resourcesDictionary.Add(ResourceType.URANUS, resources.GetInt("uranus"));
            resourcesDictionary.Add(ResourceType.STONE, resources.GetInt("stone"));
            resourcesDictionary.Add(ResourceType.WOOD, resources.GetInt("wood"));
            resourcesDictionary.Add(ResourceType.FOOD, resources.GetInt("food"));
            pd.ResourcesDictionary = resourcesDictionary;
            JSON baseJs = playerJs.GetJSON("base");
            PlayerBaseData baseData = new PlayerBaseData();
            baseData.OwnerId = baseJs.GetString("OwnerID");
            baseData.OwnerName = baseJs.GetString("Owner");
            baseData.Level = baseJs.GetInt("level");
            string[] v = baseJs.GetString("Coordinates").Split('_');
            baseData.Position = new Vector2Int(int.Parse(v[0]), int.Parse(v[1]));
            pd.baseData = baseData;
        }
        catch (Exception e)
        {
            Debug.Log("Cant parse GET_USER_ID " + e);
        }

        Messenger.Broadcast(GameEvent.AUTHORIZATION_SUCC);
    }


    private void HandleChunk(string serverMessage)
    {
        try
        {
            JSON bodyJs = JSON.ParseString(serverMessage).GetJSON(Packet.PacketKey.BODY);
            JSON chunkJs = bodyJs.GetJSON("chunk");
            JArray tilesArray = chunkJs.GetJArray("tiles");
            int chX = chunkJs.GetJSON("pos").GetInt("x");
            int chY = chunkJs.GetJSON("pos").GetInt("y");
            SerializableVector2Int pos = new SerializableVector2Int(chX, chY);
            ChunkData ch = new ChunkData();
            TileChunk[,]
                tiles = new TileChunk[SetupSetting.Instance.chunkSize,
                    SetupSetting.Instance.chunkSize]; //TODO получать из playerData
            foreach (JSON v in tilesArray.Values)
            {
                try
                {
                    TileChunk tile = new TileChunk();
                    JSON posTile = v.GetJSON("pos");
                    int x = posTile.GetInt("x");
                    int y = posTile.GetInt("y");
                    SerializableVector2Int posTileVect = new SerializableVector2Int(x, y);

                    BiomType btype = (BiomType) v.GetInt("btype");
                    tile.biomTypeType = btype;
                    tile.pos = posTileVect;

                    int resOrBuild = v.GetInt("rtype");
                    if (resOrBuild < 100)
                    {
                        ResourceType rtype = (ResourceType) resOrBuild;
                        tile.resourceType = rtype;
                    }
                    else if (resOrBuild >= 100)
                    {
                        BuildType rtype = (BuildType) resOrBuild;
                        tile.buildType = rtype;
                    }

                    tiles[x - chX, y - chY] = tile;
                }
                catch (Exception e)
                {
                    Debug.Log("Error when parse tile from TCP" + e);
                    break;
                }
            }

            ch.Position = pos;
            ch.tileChunkLayer = tiles;
            if (PlayerData.GetInstance().ChunkMap.ContainsKey(ch.Position))
                PlayerData.GetInstance().ChunkMap.Remove(ch.Position);
            PlayerData.GetInstance().ChunkMap.Add(ch.Position, ch);
            //ChunkLoadManager.chunkQueue.Enqueue(ch); 
        }
        catch (Exception e)
        {
            Debug.Log("Error when parse CHUNK from TCP" + e);
        }
    }
}