using UnityEngine;

namespace Model.BuildData
{
    public abstract class BuildData
    {
        private string ownerId;
        private string ownerName;
        private int level;
        private Vector2Int position; 

        public string OwnerId
        {
            get => ownerId;
            set => ownerId = value;
        }

        public string OwnerName
        {
            get => ownerName;
            set => ownerName = value;
        }

        public int Level
        {
            get => level;
            set => level = value;
        }

        public Vector2Int Position
        {
            get => position;
            set => position = value;
        }
    }
}