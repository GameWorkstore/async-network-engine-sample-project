using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using Asyncnetworkengine;
using Google.Protobuf;
using ProtobufData;
using UnityEngine;

public class MyMono : MonoBehaviour
{
    private const string SetEndpoint = "https://ryigmkqwt6.execute-api.us-east-1.amazonaws.com/default/lkssetdata";
    private const string GetEndpoint = "https://ryigmkqwt6.execute-api.us-east-1.amazonaws.com/default/lksgetdata";

    void Start()
    {
        /*AsyncNetworkEngine<SetUserRequest, SetUserResponse, GenericErrorResponse>.Cloud = AsyncNetworkEngineCloud.AWS;

        var rqt = new SetUserRequest()
        {
            User = new User
            {
               Id = "99123456",
               Name = "Hadoken",
               Coins = 250
            }
        };

        AsyncNetworkEngine<SetUserRequest, SetUserResponse, GenericErrorResponse>.Send(SetEndpoint, rqt, OnSetUserResponse);*/


        AsyncNetworkEngine<GetUserRequest, GetUserResponse, GenericErrorResponse>.Cloud = AsyncNetworkEngineCloud.AWS;

        var rqt = new GetUserRequest()
        {
            Id = "99123456"
        };

        AsyncNetworkEngine<GetUserRequest, GetUserResponse, GenericErrorResponse>.Send(GetEndpoint, rqt, OnGetUserResponse);
    }

    private void OnSetUserResponse(AsyncNetworkResult result, SetUserResponse response, GenericErrorResponse error)
    {
        if (result == AsyncNetworkResult.E_NETWORK)
        {
            //RETRY
        }
        Debug.Log("Result:" + result);
        Debug.Log("Response:" + (response != null ? response.ToString() : "null"));
        Debug.Log("Result:" + (error != null ? error.ToString() : "null"));
    }

    private void OnGetUserResponse(AsyncNetworkResult result, GetUserResponse response, GenericErrorResponse error)
    {
        Debug.Log("Result:" + result);
        Debug.Log("Response:" + (response != null? response.ToString() : "null"));
        Debug.Log("Result:" + (error != null? error.ToString() : "null"));
    }
}
