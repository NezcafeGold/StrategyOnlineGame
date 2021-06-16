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
        this.id = id;
        this.statusCode = statusCode;
        uuid = Guid.NewGuid().ToString();

        string text = "";
        foreach (var pair in pairs)
        {
            text += $@"""{pair.Key}"":""{pair.Value}"",";
        }

        body = text.TrimEnd(',');
    }

    public static class Body
    {
        public static KeyValuePair<string, string> Of(string key, string value)
        {
            return new KeyValuePair<string, string>(key, value);
        }
    }

    public override string ToString()
    {
        return $@"{{""header"":{{""id"":{id},""status_code"":{statusCode}, ""uuid"":""{uuid}""}},""body"":{{{body}}}}}"
               + Environment.NewLine;
    }

    public class StatusCode
    {
        public static readonly int OK_CODE = 0;

        public static readonly int ERROR_CODE = 1;

        public static readonly int AUTHORIZATION_CODE = 2;

        public static readonly int ERROR_AUTHORIZATION_CODE = 3;
    }

    public class SegmentID
    {
        public static readonly int AUTHORIZATION_CODE = 0;

        public static readonly int GET_USER_CODE = 1;

        public static readonly int GET_INVENTORY_CODE = 2;

        public static readonly int GET_UNITS_CODE = 3;

        public static readonly int GET_RESOURCES_CODE = 4;

        public static readonly int GET_CHUNK_CODE = 5;

        public static readonly int GET_TILE_CODE = 6;

        public static readonly int GET_DATA_MAP_CODE = 7;
    }
}