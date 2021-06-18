using System;
using System.Collections;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using UnityEngine;

public class TCPClient : Singleton<TCPClient>
{
    #region private members 	

    private TcpClient socketConnection;
    private Thread clientReceiveThread;
    [SerializeField] private string host = "46.0.193.126";
    [SerializeField] private int port = 33001;
    private long milliseconds;
    private int pingTime = 10;

    #endregion

    private void Awake()
    {
        DontDestroyOnLoad(gameObject);
    }

    /// <summary> 	
    /// Setup socket connection. 	
    /// </summary> 	
    public void ConnectToTcpServer()
    {
        try
        {
            clientReceiveThread = new Thread(new ThreadStart(ListenForData));
            clientReceiveThread.IsBackground = true;
            clientReceiveThread.Start();
            milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
            StartCoroutine(PingTCP());
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
            Byte[] bytes = new Byte[10000];
            while (true)
            {
                try
                {
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
                            PacketHandler.Instance().Handle(serverMessage);
                        }
                    }
                }
                catch (Exception e)
                {
                    if (e is Leguar.TotalJSON.JValueTypeException)
                        continue;
                    Debug.Log(e);
                    socketConnection.Close();
                    Debug.Log("CLOSE!");
                    break;
                }
            }
        }
        catch (SocketException socketException)
        {
            Debug.Log("Socket exception: " + socketException);
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
            NetworkStream stream = socketConnection.GetStream();
            if (stream.CanWrite)
            {
                milliseconds = DateTimeOffset.Now.ToUnixTimeMilliseconds();
                Debug.Log("client sent message: " + clientMessage);
                // Convert string message to byte array.                 
                byte[] clientMessageAsByteArray = Encoding.UTF8.GetBytes(clientMessage.Replace("\u200B", ""));
                // Write byte array to socketConnection stream.                 
                stream.Write(clientMessageAsByteArray, 0, clientMessageAsByteArray.Length);
                Debug.Log("Client sent his message - should be received by server");
            }
        }
        catch (SocketException socketException)
        {
            Debug.Log("Socket exception: " + socketException);
        }
        catch (ObjectDisposedException e)
        {
            Debug.Log(e);
        }
    }

    private IEnumerator PingTCP()
    {
        while (true)
        {
            yield return new WaitForSecondsRealtime(pingTime);
            if (milliseconds > pingTime * 1000)
                SendMessageTCP(new Packet(Packet.SegmentID.PING_ID, Packet.StatusCode.OK_CODE).WithoutUUID()
                    .ToString());
        }
    }
}