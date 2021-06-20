using System;
using System.Collections;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using UnityEngine;

public class TCPClient : Singleton<TCPClient>
{
    #region private members 	

    private TcpClient socketConnection;
    private Thread clientReceiveThread;
    private Thread pingThread;
    [SerializeField] private string host = "46.0.193.126";
    [SerializeField] private int port = 33001;
    private static long milliseconds;
    private const int pingTime = 10;
    private ConcurrentQueue<string> chunkQueue;
    public Queue<Action> actionsQueue;

    #endregion

    private void Awake()
    {
        DontDestroyOnLoad(gameObject);
        chunkQueue = new ConcurrentQueue<string>();
        actionsQueue = new Queue<Action>();
    }

    private void Update()
    {
        if (actionsQueue.Count > 0)
        {
            actionsQueue.Dequeue().Invoke();
        }
    }

    /// <summary> 	
    /// Setup socket connection. 	
    /// </summary> 	
    public void ConnectToTcpServer()
    {
        try
        {
            clientReceiveThread = new Thread(ListenForData) {IsBackground = true};
            pingThread = new Thread(PingTCP) {IsBackground = true};
            clientReceiveThread.Start();
            pingThread.Start();
            milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
            //StartCoroutine(ListenForData());
            //StartCoroutine(PingTCP());
            //StartCoroutine(HandlePacket());
        }
        catch (Exception e)
        {
            Debug.Log("On client connect exception " + e);
        }
    }

    /// <summary> 	
    /// Runs in background clientReceiveThread; Listens for incomming data. 	
    /// </summary>     
    private void ListenForData()
    {
        socketConnection = new TcpClient(host, port);
        // yield return null;
        Byte[] bytes = new Byte[10000];
        while (true)
        {
            if (socketConnection.Connected)
                using (NetworkStream stream = socketConnection.GetStream())
                {
                    int length;
                    // Read incomming stream into byte arrary. 					
                    while ((length = stream.Read(bytes, 0, bytes.Length)) != 0)
                    {
                        var incommingData = new byte[length];
                        Array.Copy(bytes, 0, incommingData, 0, length);
                        // Convert byte array to string message. 						
                        string serverMessage = Encoding.ASCII.GetString(incommingData);
                        Debug.Log("server message received as: " + serverMessage);

                        //HANDLE PACKET
                        //PacketHandler.Instance.actions.Enqueue(() => PacketHandler.Instance.Handle(serverMessage));
                        //chunkQueue.Enqueue(serverMessage);
                        Thread handleThread = new Thread(()=>HandlePacket(serverMessage));
                        handleThread.Start();
                    }
                }
            else
            {
                socketConnection.Close();
                break;
            }
        }
    }

    /// <summary> 	
    /// Send message to server using socket connection. 	
    /// </summary> 	
    public void SendMessageTCP(string clientMessage)
    {
        if (socketConnection == null)
        {
            return;
        }

        try
        {
            // Get a stream object for writing. 			
            if (!socketConnection.Connected)
            {
                socketConnection.Close();
            }

            NetworkStream stream = socketConnection.GetStream();
            if (stream.CanWrite)
            {
                milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
                Debug.Log("client sent message: " + clientMessage);
                // Convert string message to byte array.                 
                byte[] clientMessageAsByteArray = Encoding.UTF8.GetBytes(clientMessage.Replace("\u200B", ""));
                // Write byte array to socketConnection stream.                 
                stream.Write(clientMessageAsByteArray, 0, clientMessageAsByteArray.Length);
                //Debug.Log("Client sent his message - should be received by server");
            }
        }
        catch (SocketException socketException)
        {
            Debug.Log("Socket exception: " + socketException);
        }
        catch (Exception e)
        {
            Debug.Log(e);
        }
    }

    private void PingTCP()
    {
        while (true)
        {
            Thread.Sleep(1000);
            if (milliseconds + pingTime * 1000 < DateTimeOffset.Now.ToUnixTimeMilliseconds())
            {
                SendMessageTCP(new Packet(Packet.SegmentID.PING_ID, Packet.StatusCode.OK_CODE).WithoutUUID()
                    .ToString());
                milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
                Debug.Log(milliseconds);
            }
                
        }
    }

    private void HandlePacket(string servMessage)
    {
        new PacketHandler().Handle(servMessage);
    }
}