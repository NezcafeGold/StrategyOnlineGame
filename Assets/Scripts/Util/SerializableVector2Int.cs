using System;
using UnityEngine;

/// <summary>
/// Since unity doesn't flag the Vector2Int as serializable, we
/// need to create our own version. This one will automatically convert
/// between Vector2Int and SerializableVector2Int
/// </summary>
[Serializable]
public struct SerializableVector2Int
{
    /// <summary>
    /// x component
    /// </summary>
    public int x;
     
    /// <summary>
    /// y component
    /// </summary>
    public int y;

    /// <summary>
    /// Constructor
    /// </summary>
    /// <param name="rX"></param>
    /// <param name="rY"></param>
    public SerializableVector2Int(int rX, int rY)
    {
        x = rX;
        y = rY;
    }
     
    /// <summary>
    /// Returns a string representation of the object
    /// </summary>
    /// <returns></returns>
    public override string ToString()
    {
        return String.Format("[{0}, {1}]", x, y);
    }
     
    /// <summary>
    /// Automatic conversion from SerializableVector2Int to Vector2Int
    /// </summary>
    /// <param name="rValue"></param>
    /// <returns></returns>
    public static implicit operator Vector2(SerializableVector2Int rValue)
    {
        return new Vector2Int(rValue.x, rValue.y);
    }
     
    /// <summary>
    /// Automatic conversion from Vector2Int to SerializableVector2Int
    /// </summary>
    /// <param name="rValue"></param>
    /// <returns></returns>
    public static implicit operator SerializableVector2Int(Vector2Int rValue)
    {
        return new SerializableVector2Int(rValue.x, rValue.y);
    }
}