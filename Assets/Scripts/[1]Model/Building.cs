
    using Model.BuildData;
    using UnityEngine;

    public class Building : MonoBehaviour
    {
        [SerializeField] protected SpriteRenderer zone;
        [SerializeField] protected BoxCollider2D buildCollider;
        [SerializeField] protected SpriteRenderer buildSprite;
        protected BuildData buildData;
        protected Camera myMainCamera;
        
        protected void Start()
        {
            myMainCamera = Camera.main;
            
        }
        

        
        public void SetBuildType(BuildingObjectData bod)
        {
            switch (bod.type)
            {
                case BuildType.FARM:
                    buildData = new FarmData();
                
                    break;
            }
        
            buildSprite.sprite = bod.sprite;
            buildData.BuildType = bod.type;
        }
        
        public void BuildDone(bool isLoaded = false)
        {
            Vector3Int vec = Vector3Int.CeilToInt(buildCollider.transform.position);
            Vector2Int vec2Int = new Vector2Int(vec.x, vec.y);
            buildData.Position = vec2Int;
            buildData.OwnerName = PlayerData.GetInstance().Nickname;
            transform.SetParent(SetupSetting.Instance.buildingManager.transform);
            BuildingManager.Instance.AddNewBuilding(buildData, isLoaded);
        }
    }
