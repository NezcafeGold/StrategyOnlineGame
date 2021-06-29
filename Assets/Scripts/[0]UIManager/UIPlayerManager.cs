using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;

public class UIPlayerManager : MonoBehaviour
{

    [SerializeField] private TextMeshProUGUI text;
    private bool isSet = false;
    // Start is called before the first frame update
    void Start()
    {
        text.text = PlayerData.GetInstance().Nickname;
    }

    public void ShowPlayerBase()
    {
        CameraMoveManager.Instance.SetPosition(PlayerData.GetInstance().baseData.Position);
    }
}
