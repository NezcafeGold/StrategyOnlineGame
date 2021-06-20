using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Tilemaps;

public class BigTileMapManager : Singleton<BigTileMapManager>
{
    private int chunkSize;
    private Tilemap resourcesTilemap, biomsTilemap;
    private Queue<Action> unloadActions;

    private void Awake()
    {
        resourcesTilemap = SetupSetting.Instance.resourcesTileMap;
        biomsTilemap = SetupSetting.Instance.biomsTileMap;
        chunkSize = SetupSetting.Instance.chunkSize;
        unloadActions = new Queue<Action>();
    }

    private void Update()
    {
        if (unloadActions.Count > 0)
            unloadActions.Dequeue().Invoke();
    }


    public void UnloadTiles(SerializableVector2Int pos)
    {
        unloadActions.Enqueue(() => StartCoroutine(UnloadTilesCor(pos)));
    }

    private IEnumerator UnloadTilesCor(SerializableVector2Int pos)
    {
        for (int x = pos.x; x < pos.x + chunkSize; x++)
        {
            for (int y = pos.y; y < pos.y + chunkSize; y++)
            {
                resourcesTilemap.SetTile(new Vector3Int(x, y, 0), null);
                biomsTilemap.SetTile(new Vector3Int(x, y, 0), null);
                yield return null;
            }
        }
    }
}