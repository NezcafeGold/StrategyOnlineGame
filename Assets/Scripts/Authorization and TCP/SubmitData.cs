using System;
using System.Text;
using TMPro;
using UnityEditor;
using UnityEngine;

public class SubmitData : TCPClient
{
    [SerializeField] private TextMeshProUGUI emailText;
    [SerializeField] private TextMeshProUGUI passwordText;
//    [SerializeField] private TCPClient tcpClient;

    private string email;

    private string password;


    public void Submit()
    {
        //tcpClient.CloseClient();
        email = emailText.text;
        password = passwordText.text;
        StartClient();
        SendMessageToServer(new Packet(Packet.SegmentID.AUTHORIZATION_CODE,
            Packet.StatusCode.AUTHORIZATION_CODE,
            Packet.Body.Of("email", email),
            Packet.Body.Of("password", password)).ToString());
    }

    protected override void OnMessageReceived(string receivedMessage)
    {
        base.OnMessageReceived(receivedMessage);
        SendMessageToServer(new Packet(Packet.SegmentID.GET_USER_CODE, Packet.StatusCode.OK_CODE).ToString());
    }
}