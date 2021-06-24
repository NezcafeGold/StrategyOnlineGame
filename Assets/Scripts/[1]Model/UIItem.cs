
using TMPro;
using UnityEngine;
using UnityEngine.UI;

public class UIItem : MonoBehaviour
{
    [SerializeField] private TextMeshProUGUI valueUI;
    [SerializeField] private Image iconUI;

    public ResourceType resType;
    public int Value;
    public Sprite iconSprite;

    

    private void Awake()
    {
        Messenger.MarkAsPermanent(GameEvent.UPDATE_USER_STATS);
        Messenger.AddListener(GameEvent.UPDATE_USER_STATS, UpdateStat);
    }

    private void OnDestroy()
    {
        Messenger.RemoveListener(GameEvent.UPDATE_USER_STATS, UpdateStat);
    }

    public void UpdateStat()
    {
        valueUI.SetText(Value.ToString());
        Value = PlayerData.GetInstance().GetValueForType(resType);
        valueUI.SetText(Value.ToString());
        iconUI.sprite = iconSprite;
    }


}