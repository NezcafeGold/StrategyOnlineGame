using UnityEngine;
using System.Collections;
using System;
using System.IO;
using System.Net.Sockets;

public class TcpClientAndroid : MonoBehaviour
{
    bool socketReady = false;

    TcpClient mySocket;
    NetworkStream theStream;
    StreamWriter theWriter;
    StreamReader theReader;
    [SerializeField] private string host = "46.0.193.126";
    [SerializeField] private int port = 33001;

    // Use this for initialization
    void Start()
    {
    }

    // Update is called once per frame
    void Update()
    {
    }

    public void setupSocket()
    {
        try
        {
            mySocket = new TcpClient(host, port);
            theStream = mySocket.GetStream();
            theWriter = new StreamWriter(theStream);
            theReader = new StreamReader(theStream);
            socketReady = true;
        }
        catch (Exception e)
        {
            Debug.Log("Socket error:" + e);
        }
    }

    public void SendMessageTCP(string theLine)
    {
        if (!socketReady)
            return;
        String tmpString = theLine + "\r\n";
        theWriter.Write(tmpString);
        theWriter.Flush();
    }

    public void ListenForData()
    {
        string receivedData;
        if (!socketReady)
            receivedData = "";
        if (theStream.DataAvailable)
            receivedData = theReader.ReadLine();
        
        
    }

    public void closeSocket()
    {
        if (!socketReady)
            return;
        theWriter.Close();
        theReader.Close();
        mySocket.Close();
        socketReady = false;
    }

    public void maintainConnection()
    {
        if (!theStream.CanRead)
        {
            setupSocket();
        }
    }
}