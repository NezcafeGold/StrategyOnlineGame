using System.Collections;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class ResourcesManager : MonoBehaviour
{
    [SerializeField] private UIItem uiItem;
    private List<ResourcesData> resDataObjects;
    private bool isSet = false;

    private void Awake()
    {
        resDataObjects = Resources.LoadAll<ResourcesData>("Blocks").ToList();
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
            if (PlayerData.Instance.ResourcesDictionary.Count > 0 && !isSet)
            {
                foreach (var rdo in resDataObjects)
                {
                    if (rdo.type.Equals(ResourceType.NONE)) continue;
                    UIItem item = Instantiate(uiItem, transform, true);
                    item.iconSprite = rdo.sprite;
                    item.resType = rdo.type;
                    item.UpdateStat();
                    yield return null;
                }
                break;
            }

            yield return null;
        }

        yield return null;
    }
}