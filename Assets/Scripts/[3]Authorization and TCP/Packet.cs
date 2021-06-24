using System;
using System.Collections.Generic;

public class Packet
{
    private int id;
    private int statusCode;
    private string uuid;
    private string body;

    public Packet(int id, int statusCode, params KeyValuePair<string, string>[] pairs)
    {
        MakeDefault(id, statusCode);
        string text = "";
        foreach (var pair in pairs)
        {
            text += $@"""{pair.Key}"":""{pair.Value}"",";
        }

        body = text.TrimEnd(',');
    }

    public Packet(int id, int statusCode, params KeyValuePair<string, int>[] pairs)
    {
        MakeDefault(id, statusCode);
        string text = "";
        foreach (var pair in pairs)
        {
            text += $@"""{pair.Key}"":{pair.Value},";
        }

        body = text.TrimEnd(',');
    }

    public Packet(int id, int statusCode)
    {
        MakeDefault(id, statusCode);
    }

    private void MakeDefault(int id, int statusCode)
    {
        this.id = id;
        this.statusCode = statusCode;
        uuid = Guid.NewGuid().ToString();
    }

    public static class Body
    {
        public static KeyValuePair<string, string> Of(string key, string value)
        {
            return new KeyValuePair<string, string>(key, value);
        }

        public static KeyValuePair<string, int> OfInt(string key, int value)
        {
            return new KeyValuePair<string, int>(key, value);
        }
    }

    public Packet WithoutUUID()
    {
        uuid = null;
        return this;
    }

    public override string ToString()
    {
        if (uuid != null)
            return
                $@"{{""header"":{{""id"":{id},""status_code"":{statusCode}, ""uuid"":""{uuid}""}},""body"":{{{body}}}}}"
                + Environment.NewLine;

        return
            $@"{{""header"":{{""id"":{id},""status_code"":{statusCode}}},""body"":{{{body}}}}}"
            + Environment.NewLine;
    }

    public class PacketKey
    {
        public const string HEADER = "header";
        public const string BODY = "body";
        public const string ID = "id";
        public const string STATUS_CODE = "status_code";
    }

    public class StatusCode
    {
        public const int OK_CODE = 0;

        public const int ERROR_CODE = 1;

        public const int AUTHORIZATION_CODE = 2;

        public const int ERROR_AUTHORIZATION_CODE = 3;
    }

    public class SegmentID
    {
        public static readonly int PING_ID = 0;

        public static readonly int SHUTDOWN_ID = 1;
        
        public const int AUTHORIZATION_ID = 2;

        public const int GET_USER_ID = 100;

        public const int GET_INVENTORY_ID = 101;

        public const int GET_UNITS_ID = 102;

        public const int GET_RESOURCES_ID = 103;

        public const int GET_CHUNK_ID = 200;

        public const int GET_TILE_ID = 201;

        public const int GET_DATA_MAP_ID = 202;
        
        
        
        
        
        
        
        
        
        
        
        public const int SET_FARM_ID = 301;
        
    }
}