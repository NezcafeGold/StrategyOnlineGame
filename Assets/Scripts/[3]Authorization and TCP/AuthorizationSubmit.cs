using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.SceneManagement;

public class AuthorizationSubmit : MonoBehaviour
{
    [SerializeField] private TMP_InputField emailText;
    [SerializeField] private TMP_InputField passwordText;
    private TCPClient tcpClient;

    private string email;

    private string password;

    private void Start()
    {
        tcpClient = TCPClient.Instance;
    }

    public void Submit()
    {
        email = emailText.text;
        password = passwordText.text;
        StartCoroutine(SendMessageCor());
    }

    private IEnumerator SendMessageCor()
    {
        tcpClient.actionsQueue.Enqueue(tcpClient.ConnectToTcpServer);
        yield return new WaitForSecondsRealtime(0.1f);
        tcpClient.SendMessageTCP(new Packet(Packet.SegmentID.AUTHORIZATION_ID,
            Packet.StatusCode.AUTHORIZATION_CODE,
            Packet.Body.Of("email", email),
            Packet.Body.Of("password", password)).ToString());
        yield return new WaitForSecondsRealtime(0.1f);
        tcpClient.SendMessageTCP(new Packet(Packet.SegmentID.GET_USER_ID,
            Packet.StatusCode.OK_CODE).ToString());
    }
}