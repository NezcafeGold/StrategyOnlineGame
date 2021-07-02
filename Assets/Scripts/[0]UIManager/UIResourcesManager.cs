using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class UIResourcesManager : MonoBehaviour
{
    [SerializeField] private UIResourceItem uiResourceItem;
    private List<ResourcesObjectData> resDataObjects;
    private bool isSet = false;

    private void Awake()
    {
        resDataObjects = Resources.LoadAll<ResourcesObjectData>("Blocks").ToList();
    }

    // Start is called before the first frame update
    void Start()
    {
        StartCoroutine(StatSetter());
    }

    private IEnumerator StatSetter()
    {
        while (true)
        {
            if (PlayerData.GetInstance().ResourcesDictionary.Count > 0 && !isSet)
            {
                foreach (var rdo in resDataObjects)
                {
                    if (rdo.type.Equals(ResourceType.NONE)) continue;
                    UIResourceItem resourceItem = Instantiate(uiResourceItem, transform, true);
                    resourceItem.iconSprite = rdo.sprite;
                    resourceItem.resType = rdo.type;
                    resourceItem.UpdateStat();
                    yield return null;
                }
                break;
            }

            yield return null;
        }

        yield return null;
    }
}