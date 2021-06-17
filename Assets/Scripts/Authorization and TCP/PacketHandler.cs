using UnityEngine;
using UnityEngine.SceneManagement;

public class PacketHandler
{
    private bool CheckForCorrectPacket()
    {
        //TODO: Add method
        return true;
    }

    public void Handle(string serverMessage)
    {
        if (!CheckForCorrectPacket())
            return;

        HandleHeader(new JSONHandler(serverMessage).GetStringObject(Packet.PacketKey.HEADER));
    }

    private void HandleHeader(string header)
    {
        JSONHandler jh = new JSONHandler(header);
        if (jh.GetInt(Packet.PacketKey.STATUS_CODE).Equals(Packet.StatusCode.OK_CODE))
        {
            HandleByID(jh.GetInt(Packet.PacketKey.ID));
        }
    }

    private void HandleByID(int id)
    {
        switch (id)
        {
            case Packet.SegmentID.AUTHORIZATION_CODE:
                CustomSceneManager.Instance.LoadScene();
              break;

            case Packet.SegmentID.GET_TILE_CODE:

                break;

            case Packet.SegmentID.GET_USER_CODE:

                break;
            case Packet.SegmentID.GET_CHUNK_CODE:
                Debug.Log("HANDLE CHUNK");
                break;
            case Packet.SegmentID.GET_UNITS_CODE:

                break;
            case Packet.SegmentID.GET_DATA_MAP_CODE:

                break;
            case Packet.SegmentID.GET_INVENTORY_CODE:

                break;
            case Packet.SegmentID.GET_RESOURCES_CODE:

                break;
        }
    }
}