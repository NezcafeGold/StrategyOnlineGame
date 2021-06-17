using Leguar.TotalJSON;
using UnityEngine;

public class JSONHandler
{
    private JSON textJson;

    public JSONHandler(string text)
    {
        textJson = JSON.ParseString(text);
    }

    public string GetString(string key)
    {
        string t = "";
        try
        {
            t = textJson.GetString(key);
        }
        catch (JSONKeyNotFoundException e)
        {
            Debug.Log(e);
        }

        return t;
    }


    public string GetStringObject(string key)
    {
        string t = "";
        try
        {
            t = textJson.GetJSON(key).CreateString();
        }
        catch (JSONKeyNotFoundException e)
        {
            Debug.Log(e);
        }

        return t;
    }

    public int GetInt(string key)
    {
        int t = 999999;
        try
        {
            t = textJson.GetInt(key);
        }
        catch (JSONKeyNotFoundException e)
        {
            Debug.Log(e);
        }

        return t;
    }
}