

using System;
using System.Collections.Generic;
using UnityEditor;
using UnityEngine;

public class CustomSceneManager : MonoBehaviour
{
    [SerializeField] private SceneAsset scene;
    [SerializeField] private Color color;
    [SerializeField] private float time;
    [SerializeField] private GameObject[] gameObjects;

    private Queue<Action> jobs = new Queue<Action>();
    
    private void Awake()
    {
        Messenger.AddListener(GameEvent.AUTHORIZATION_SUCC, AddLoadJob);
    }

    private void OnDestroy()
    {
        Messenger.RemoveListener(GameEvent.AUTHORIZATION_SUCC, AddLoadJob);
    }

    private void Update()
    {
        while (jobs.Count > 0) 
            jobs.Dequeue().Invoke();
    }

    private void AddLoadJob()
    {
        jobs.Enqueue(LoadGameScene);
    }

    private void  LoadGameScene()
    {
        foreach (var g in gameObjects)
        {
            g.SetActive(false);
        }

        Initiate.Fade(scene.name, color, time);
    }
}