using UnityEditor;
using UnityEngine;

public class CustomSceneManager : Singleton<CustomSceneManager>
{
    [SerializeField] private SceneAsset scene;
    [SerializeField] private Color color;
    [SerializeField] private float time;
    [SerializeField] private GameObject[] gameObjects;

    public void LoadScene()
    {
        foreach (var g in gameObjects)
        {
            g.SetActive(false);
        }

        Initiate.Fade(scene.name, color, time);
    }
}