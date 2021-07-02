using System;
using System.Collections;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.IO;
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
    private bool isPingTcp = true;
    private ConcurrentQueue<string> chunkQueue;
    public Queue<Action> actionsQueue;

    #endregion

    public static bool isRunning = false;

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

    private void OnDestroy()
    {
        isPingTcp = false;
        if (clientReceiveThread.ThreadState == ThreadState.WaitSleepJoin)
            clientReceiveThread.Interrupt();
        if (pingThread.ThreadState == ThreadState.WaitSleepJoin)
            pingThread.Interrupt();
        isRunning = false;
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
            isRunning = true;
            Debug.Log("Server has connected");
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
        try
        {
            socketConnection = new TcpClient(host, port);
        }
        catch (Exception e)
        {
            Debug.Log("Cant establish connect with " + host + " " + port + " " + e);
        }

        // yield return null;
        Byte[] bytes = new Byte[20000];
        while (true)
        {
            if (socketConnection.Connected)
                using (StreamReader reader = new StreamReader(socketConnection.GetStream()))
                {
                    int length;
                    // Read incomming stream into byte arrary. 				
                    string serverMessage;
                    while (((serverMessage = reader.ReadLine()) !=""))
                    {
                        //var incommingData = new byte[length];
                        //Array.Copy(bytes, 0, incommingData, 0, length);
                        // Convert byte array to string message. 						
                        //string serverMessage = Encoding.ASCII.GetString(incommingData);
                        Debug.Log("server message received as: " + serverMessage);


                        var message = serverMessage;
                        Thread handleThread = new Thread(() => HandlePacket(message));
                        handleThread.Start();
                    }
                }
            else
            {
                socketConnection.Close();
                isRunning = false;
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
            StreamWriter writer = new StreamWriter(stream);
            if (stream.CanWrite)
            {
                milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();

                // Convert string message to byte array.                 
                //byte[] clientMessageAsByteArray = Encoding.UTF8.GetBytes(clientMessage.Replace("\u200B", ""));
                // Write byte array to socketConnection stream.                 
                writer.Write(clientMessage);
                writer.Flush();
                Debug.Log("client sent message: " + clientMessage);
                //Debug.Log("Client sent his message - should be received by server");
            }
        }
        catch (SocketException socketException)
        {
            Debug.Log("Socket exception: " + socketException);
            socketConnection.Close();
        }
        catch (Exception e)
        {
            Debug.Log(e);
            isRunning = false;
            socketConnection.Close();
        }
    }

    private void PingTCP()
    {
        while (isPingTcp)
        {
            Thread.Sleep(1000);
            if (milliseconds + pingTime * 1000 < DateTimeOffset.Now.ToUnixTimeMilliseconds())
            {
                SendMessageTCP(new Packet(Packet.SegmentID.PING_ID, Packet.StatusCode.OK_CODE).WithoutUUID()
                    .ToString());
                milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
                Debug.Log("PING!");
            }
        }
    }

    private void HandlePacket(string servMessage)
    {
        new PacketHandler().Handle(servMessage);
    }
}