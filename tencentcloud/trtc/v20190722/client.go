// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190722

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateCloudRecordingRequest() (request *CreateCloudRecordingRequest) {
    request = &CreateCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudRecording")
    
    
    return
}

func NewCreateCloudRecordingResponse() (response *CreateCloudRecordingResponse) {
    response = &CreateCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudRecording
// ### API description
//
// This API is used to start an on-cloud recording task. It records the audio and video streams in a room and saves them to the specified cloud storage. You can use this API to record the streams in a room separately, or you can mix the streams first and then record the mixed stream.
//
// 
//
// ### You can use this API to perform the following operations:
//
// * Specify the anchors whose streams you want or do not want to record by using the `RecordParams` parameter
//
// * Specify the storage service you want to save recording files to by using the `StorageParams` parameter
//
// * Specify transcoding settings for mixed-stream recording, including video resolution, video bitrate, frame rate, and audio quality, by using `MixTranscodeParams`
//
// * Specify the layout of different videos in mixed-stream recording mode or select an auto-arranged layout template
//
// 
//
// ### Key concepts
//
// * Single-stream recording: Record the audio and video of each subscribed user (`UserId`) in a room and save the recording files (M3U8/TS) to the cloud
//
// * Mixed-stream recording: Mix the audios and videos of subscribed users (`UserId`) in a room, record the mixed stream, and save the recording files (M3U8/TS) to the cloud
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"
//  MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"
//  MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"
//  MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
func (c *Client) CreateCloudRecording(request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    return c.CreateCloudRecordingWithContext(context.Background(), request)
}

// CreateCloudRecording
// ### API description
//
// This API is used to start an on-cloud recording task. It records the audio and video streams in a room and saves them to the specified cloud storage. You can use this API to record the streams in a room separately, or you can mix the streams first and then record the mixed stream.
//
// 
//
// ### You can use this API to perform the following operations:
//
// * Specify the anchors whose streams you want or do not want to record by using the `RecordParams` parameter
//
// * Specify the storage service you want to save recording files to by using the `StorageParams` parameter
//
// * Specify transcoding settings for mixed-stream recording, including video resolution, video bitrate, frame rate, and audio quality, by using `MixTranscodeParams`
//
// * Specify the layout of different videos in mixed-stream recording mode or select an auto-arranged layout template
//
// 
//
// ### Key concepts
//
// * Single-stream recording: Record the audio and video of each subscribed user (`UserId`) in a room and save the recording files (M3U8/TS) to the cloud
//
// * Mixed-stream recording: Mix the audios and videos of subscribed users (`UserId`) in a room, record the mixed stream, and save the recording files (M3U8/TS) to the cloud
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"
//  MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"
//  MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"
//  MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
func (c *Client) CreateCloudRecordingWithContext(ctx context.Context, request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    if request == nil {
        request = NewCreateCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudRecordingRequest() (request *DeleteCloudRecordingRequest) {
    request = &DeleteCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudRecording")
    
    
    return
}

func NewDeleteCloudRecordingResponse() (response *DeleteCloudRecordingResponse) {
    response = &DeleteCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudRecording
// This API is used to stop a recording task. It works only when a task is in progress. If the task has already ended when this API is called, an error will be returned. If a task is stopped successfully, but the uploading of recording files has not completed, the backend will continue to upload the files and will notify you via a callback when the upload is completed.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DeleteCloudRecording(request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    return c.DeleteCloudRecordingWithContext(context.Background(), request)
}

// DeleteCloudRecording
// This API is used to stop a recording task. It works only when a task is in progress. If the task has already ended when this API is called, an error will be returned. If a task is stopped successfully, but the uploading of recording files has not completed, the backend will continue to upload the files and will notify you via a callback when the upload is completed.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DeleteCloudRecordingWithContext(ctx context.Context, request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDeleteCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRecordingRequest() (request *DescribeCloudRecordingRequest) {
    request = &DescribeCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudRecording")
    
    
    return
}

func NewDescribeCloudRecordingResponse() (response *DescribeCloudRecordingResponse) {
    response = &DescribeCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudRecording
// This API is used to query the status of a recording task after it starts. It works only when a task is in progress. If the task has already ended when this API is called, an error will be returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRecording(request *DescribeCloudRecordingRequest) (response *DescribeCloudRecordingResponse, err error) {
    return c.DescribeCloudRecordingWithContext(context.Background(), request)
}

// DescribeCloudRecording
// This API is used to query the status of a recording task after it starts. It works only when a task is in progress. If the task has already ended when this API is called, an error will be returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRecordingWithContext(ctx context.Context, request *DescribeCloudRecordingRequest) (response *DescribeCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomRequest() (request *DismissRoomRequest) {
    request = &DismissRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoom")
    
    
    return
}

func NewDismissRoomResponse() (response *DismissRoomResponse) {
    response = &DismissRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DismissRoom
// This API is used to remove all users from a room and dismiss the room. It supports all platforms. For Android, iOS, Windows, and macOS, the TRTC SDK needs to be upgraded to v6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoom(request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    return c.DismissRoomWithContext(context.Background(), request)
}

// DismissRoom
// This API is used to remove all users from a room and dismiss the room. It supports all platforms. For Android, iOS, Windows, and macOS, the TRTC SDK needs to be upgraded to v6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomWithContext(ctx context.Context, request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    if request == nil {
        request = NewDismissRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomByStrRoomIdRequest() (request *DismissRoomByStrRoomIdRequest) {
    request = &DismissRoomByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoomByStrRoomId")
    
    
    return
}

func NewDismissRoomByStrRoomIdResponse() (response *DismissRoomByStrRoomIdResponse) {
    response = &DismissRoomByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DismissRoomByStrRoomId
// This API is used to remove all users from a room and close the room. It works on all platforms. For Android, iOS, Windows, and macOS, you need to update the TRTC SDK to version 6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomByStrRoomId(request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    return c.DismissRoomByStrRoomIdWithContext(context.Background(), request)
}

// DismissRoomByStrRoomId
// This API is used to remove all users from a room and close the room. It works on all platforms. For Android, iOS, Windows, and macOS, you need to update the TRTC SDK to version 6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomByStrRoomIdWithContext(ctx context.Context, request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewDismissRoomByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoomByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudRecordingRequest() (request *ModifyCloudRecordingRequest) {
    request = &ModifyCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudRecording")
    
    
    return
}

func NewModifyCloudRecordingResponse() (response *ModifyCloudRecordingResponse) {
    response = &ModifyCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudRecording
// This API is used to modify a recording task. It works only when a task is in progress. If the task has already ended when this API is called, an error will be returned. You need to specify all the parameters for each request instead of just the ones you want to modify.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudRecording(request *ModifyCloudRecordingRequest) (response *ModifyCloudRecordingResponse, err error) {
    return c.ModifyCloudRecordingWithContext(context.Background(), request)
}

// ModifyCloudRecording
// This API is used to modify a recording task. It works only when a task is in progress. If the task has already ended when this API is called, an error will be returned. You need to specify all the parameters for each request instead of just the ones you want to modify.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudRecordingWithContext(ctx context.Context, request *ModifyCloudRecordingRequest) (response *ModifyCloudRecordingResponse, err error) {
    if request == nil {
        request = NewModifyCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserRequest() (request *RemoveUserRequest) {
    request = &RemoveUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUser")
    
    
    return
}

func NewRemoveUserResponse() (response *RemoveUserResponse) {
    response = &RemoveUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveUser
// This API is used to remove a user from a room. It is applicable to scenarios where the anchor, room owner, or admin wants to kick out a user. It supports all platforms. For Android, iOS, Windows, and macOS, the TRTC SDK needs to be upgraded to v6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUser(request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    return c.RemoveUserWithContext(context.Background(), request)
}

// RemoveUser
// This API is used to remove a user from a room. It is applicable to scenarios where the anchor, room owner, or admin wants to kick out a user. It supports all platforms. For Android, iOS, Windows, and macOS, the TRTC SDK needs to be upgraded to v6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserWithContext(ctx context.Context, request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    if request == nil {
        request = NewRemoveUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserByStrRoomIdRequest() (request *RemoveUserByStrRoomIdRequest) {
    request = &RemoveUserByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUserByStrRoomId")
    
    
    return
}

func NewRemoveUserByStrRoomIdResponse() (response *RemoveUserByStrRoomIdResponse) {
    response = &RemoveUserByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveUserByStrRoomId
// This API is used to remove a user from a room. It allows the anchor, room owner, or admin to kick out a user, and works on all platforms. For Android, iOS, Windows, and macOS, you need to update the TRTC SDK to version 6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserByStrRoomId(request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    return c.RemoveUserByStrRoomIdWithContext(context.Background(), request)
}

// RemoveUserByStrRoomId
// This API is used to remove a user from a room. It allows the anchor, room owner, or admin to kick out a user, and works on all platforms. For Android, iOS, Windows, and macOS, you need to update the TRTC SDK to version 6.6 or above.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserByStrRoomIdWithContext(ctx context.Context, request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewRemoveUserByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}
