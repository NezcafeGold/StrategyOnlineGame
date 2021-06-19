using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Threading;
using Leguar.TotalJSON;
using UnityEngine;

public class PacketHandler
{
    private static PacketHandler instance;
    private static Thread th;

    public static Queue<string> packets = new Queue<string>();
    public static Queue<Action> actions = new Queue<Action>();


    public static PacketHandler Instance()
    {
        if (instance == null)
        {
            instance = new PacketHandler();
            th = new Thread(UpdateTick);
            th.Start();
        }

        if (!th.IsAlive)
            th.Start();

        return instance;
    }


    private bool CheckForCorrectPacket()
    {
        //TODO: Add method
        return true;
    }

    public void Handle(string serverMessage)
    {
        if (!CheckForCorrectPacket())
            return;

        packets.Enqueue(serverMessage);
    }

    private static void UpdateTick()
    {
        while (true)
        {
            while (actions.Count > 0)
                actions.Dequeue().Invoke();

            //TODO: КОРУТИНКА?
            while (packets.Count > 0)
            {
                string serverMessage = packets.Dequeue();
                HandleHeader(new JSONHandler(serverMessage).GetStringObject(Packet.PacketKey.HEADER),
                    serverMessage);
            }
        }
    }

    private static void HandleHeader(string header, string serverMessage)
    {
        JSONHandler jh = new JSONHandler(header);
        if (jh.GetInt(Packet.PacketKey.STATUS_CODE).Equals(Packet.StatusCode.OK_CODE))
        {
            HandleByID(jh.GetInt(Packet.PacketKey.ID), serverMessage);
        }
    }

    private static void HandleByID(int id, string serverMessage)
    {
        try
        {
            switch (id)
            {
                case Packet.SegmentID.AUTHORIZATION_ID:
                    Messenger.Broadcast(GameEvent.AUTHORIZATION_SUCC);
                    break;

                case Packet.SegmentID.GET_TILE_ID:

                    break;

                case Packet.SegmentID.GET_USER_ID:

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
            }
        }
        catch (Exception e)
        {
            Debug.Log(e);
        }
    }

    private static void HandleChunk(string serverMessage)
    {
        try
        {
            String body = new JSONHandler(serverMessage).GetStringObject(Packet.PacketKey.BODY);
            JSON bodyJs = JSON.ParseString(body);
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
                TileChunk tile = new TileChunk();
                JSON posTile = v.GetJSON("pos");

                int x = posTile.GetInt("x");
                int y = posTile.GetInt("y");
                SerializableVector2Int posTileVect = new SerializableVector2Int(x, y);
                ResourceType type = (ResourceType) v.GetInt("type");

                tile.pos = posTileVect;
                tile.resourceType = type;
                tiles[x - chX, y - chY] = tile;
            }

            ch.Position = pos;
            ch.tileChunkLayer = tiles;
            PlayerData.Instance.ChunkMap.Add(ch.Position, ch);
        }
        catch (Exception e)
        {
            Debug.Log("Error when parse CHUNK from TCP" + e);
        }
    }
}