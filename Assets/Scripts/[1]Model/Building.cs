using System;
using Model.BuildData;
using TMPro;
using UnityEngine;

public class Building : MonoBehaviour
{
    [SerializeField] protected SpriteRenderer zone;
    [SerializeField] protected BoxCollider2D buildCollider;
    [SerializeField] protected SpriteRenderer buildSprite;
    [SerializeField] protected TextMeshProUGUI ownerName;
    public BuildData buildData;
    protected Camera myMainCamera;
    private bool isSet = false;

    private void Awake()
    {
        Messenger.AddListener<BuildData>(GameEvent.UPDATE_BUILD_INFO, SetBuildDataFromManager);
    }

    private void OnDestroy()
    {
        Messenger.RemoveListener<BuildData>(GameEvent.UPDATE_BUILD_INFO, SetBuildDataFromManager);
    }

    protected void Start()
    {
        myMainCamera = Camera.main;
    }

    private void Update()
    {
        if (!isSet && buildData != null && buildData.OwnerName != null)
        {
            //ownerName.text = buildData.OwnerName;
        }
    }


    public void SetBuildType(BuildingObjectData bod)
    {
        switch (bod.type)
        {
            case BuildType.BASE:
                buildData = new BaseData();
                break;

            case BuildType.FARM:
                buildData = new FarmData();
                break;
            default:

                break;
        }

        try
        {
            buildSprite.sprite = bod.sprite;
            buildData.BuildType = bod.type;

            Vector3Int vec = Vector3Int.CeilToInt(buildCollider.transform.position);
            Vector2Int vec2Int = new Vector2Int(vec.x, vec.y);
            buildData.Position = vec2Int;
            buildData.OwnerName = PlayerData.GetInstance().Nickname;

            transform.SetParent(SetupSetting.Instance.buildingManager.transform);

            //
            //
        }
        catch (Exception e)
        {
            Debug.Log("Cant set type for " + bod.type + e);
        }
    }

    public void BuildDone()
    {
        StartCoroutine(BuildingManager.Instance.UpdateBuildMap(buildData));
    }

    public void AskData()
    {
        BuildingManager.Instance.AskDataTcp(buildData);
    }

    public void SetBuildDataFromManager(BuildData buildData)
    {
        if (buildData.Position == this.buildData.Position)
        {
            this.buildData = buildData;
            ownerName.text = buildData.OwnerName;

            if (PlayerData.GetInstance().Nickname == buildData.OwnerName)
            {
                zone.color = new Color(0.5f, 0.5f, 0, 0.5f);
            }
            else zone.color = new Color(0.5f, 0, .5f, 0.5f);
        }
    }
}