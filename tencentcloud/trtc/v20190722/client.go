// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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


func NewControlAIConversationRequest() (request *ControlAIConversationRequest) {
    request = &ControlAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ControlAIConversation")
    
    
    return
}

func NewControlAIConversationResponse() (response *ControlAIConversationResponse) {
    response = &ControlAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlAIConversation
// This API is used to provide service for server-side robot control.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) ControlAIConversation(request *ControlAIConversationRequest) (response *ControlAIConversationResponse, err error) {
    return c.ControlAIConversationWithContext(context.Background(), request)
}

// ControlAIConversation
// This API is used to provide service for server-side robot control.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) ControlAIConversationWithContext(ctx context.Context, request *ControlAIConversationRequest) (response *ControlAIConversationResponse, err error) {
    if request == nil {
        request = NewControlAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ControlAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudModerationRequest() (request *CreateCloudModerationRequest) {
    request = &CreateCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudModeration")
    
    
    return
}

func NewCreateCloudModerationResponse() (response *CreateCloudModerationResponse) {
    response = &CreateCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudModeration
// API description:
//
// This API is used to enable the cloud moderation feature to complete audio and video slicing, video frame extraction, and audio stream recording in the room, and submit them to the specified moderation supplier for completing the moderation.
//
// 
//
// This API is used to achieve the following goals:
//
// * This API is used to specify the moderation parameters (ModerationParams) to specify the detailed parameters required for moderation.
//
// * This API is used to specify the storage parameter (SliceStorageParams) to specify the cloud storage you want to upload the file complying with the moderation policy to. Currently, Tencent Cloud Object Storage (COS) and third-party AWS are supported.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudModeration(request *CreateCloudModerationRequest) (response *CreateCloudModerationResponse, err error) {
    return c.CreateCloudModerationWithContext(context.Background(), request)
}

// CreateCloudModeration
// API description:
//
// This API is used to enable the cloud moderation feature to complete audio and video slicing, video frame extraction, and audio stream recording in the room, and submit them to the specified moderation supplier for completing the moderation.
//
// 
//
// This API is used to achieve the following goals:
//
// * This API is used to specify the moderation parameters (ModerationParams) to specify the detailed parameters required for moderation.
//
// * This API is used to specify the storage parameter (SliceStorageParams) to specify the cloud storage you want to upload the file complying with the moderation policy to. Currently, Tencent Cloud Object Storage (COS) and third-party AWS are supported.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudModerationWithContext(ctx context.Context, request *CreateCloudModerationRequest) (response *CreateCloudModerationResponse, err error) {
    if request == nil {
        request = NewCreateCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudModerationResponse()
    err = c.Send(request, response)
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
// API description:
//
// This API is used to start an on-cloud recording task. It records the audio and video streams in a room and saves them to the specified cloud storage. You can use this API to record the streams in a room separately, or you can mix the streams first and then record the mixed stream.
//
// 
//
// You can use this API to perform the following operations:
//
// * Specify the anchors whose streams you want or do not want to record by using the `RecordParams` parameter
//
// * Specify the storage service you want to save recording files to by using the `StorageParams` parameter. Currently, you can save recording files to Tencent Cloud VOD or COS.
//
// * Specify transcoding settings for mixed-stream recording, including video resolution, video bitrate, frame rate, and audio quality, by using `MixTranscodeParams`
//
// * Specify the layout of different videos in mixed-stream recording mode or select an auto-arranged layout template
//
// 
//
// Key concepts:
//
// * Single-stream recording: Record the audio and video of each subscribed user (`UserId`) in a room and save the recording files to the storage you specify.
//
// Mixed-stream recording: Mix the audios and videos of subscribed users (`UserId`) in a room, record the mixed stream, and save the recording files to the storage you specify. After a recording task ends, you can go to the VOD console (https://console.tencentcloud.com/vod/media) or [COS console](https://console.cloud.tencent.com/cos/bucket) to view the recording files.
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
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudRecording(request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    return c.CreateCloudRecordingWithContext(context.Background(), request)
}

// CreateCloudRecording
// API description:
//
// This API is used to start an on-cloud recording task. It records the audio and video streams in a room and saves them to the specified cloud storage. You can use this API to record the streams in a room separately, or you can mix the streams first and then record the mixed stream.
//
// 
//
// You can use this API to perform the following operations:
//
// * Specify the anchors whose streams you want or do not want to record by using the `RecordParams` parameter
//
// * Specify the storage service you want to save recording files to by using the `StorageParams` parameter. Currently, you can save recording files to Tencent Cloud VOD or COS.
//
// * Specify transcoding settings for mixed-stream recording, including video resolution, video bitrate, frame rate, and audio quality, by using `MixTranscodeParams`
//
// * Specify the layout of different videos in mixed-stream recording mode or select an auto-arranged layout template
//
// 
//
// Key concepts:
//
// * Single-stream recording: Record the audio and video of each subscribed user (`UserId`) in a room and save the recording files to the storage you specify.
//
// Mixed-stream recording: Mix the audios and videos of subscribed users (`UserId`) in a room, record the mixed stream, and save the recording files to the storage you specify. After a recording task ends, you can go to the VOD console (https://console.tencentcloud.com/vod/media) or [COS console](https://console.cloud.tencent.com/cos/bucket) to view the recording files.
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
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudRecordingWithContext(ctx context.Context, request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    if request == nil {
        request = NewCreateCloudRecordingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudSliceTaskRequest() (request *CreateCloudSliceTaskRequest) {
    request = &CreateCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudSliceTask")
    
    
    return
}

func NewCreateCloudSliceTaskResponse() (response *CreateCloudSliceTaskResponse) {
    response = &CreateCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudSliceTask
// API description:
//
// This API is used to enable the cloud slicing feature, completing audio and video slicing tasks in the room, and uploading them to the specified cloud storage.
//
// This API is used to achieve the following goals:
//
// * This API is used to specify the slicing parameter (SliceParams) to define the blocklist or allowlist of the anchors that require slicing.
//
// * This API is used to specify the storage parameter (SliceStorageParams) to specify the cloud storage you want to upload to. Currently, Tencent Cloud Object Storage (COS) and third-party AWS are supported.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudSliceTask(request *CreateCloudSliceTaskRequest) (response *CreateCloudSliceTaskResponse, err error) {
    return c.CreateCloudSliceTaskWithContext(context.Background(), request)
}

// CreateCloudSliceTask
// API description:
//
// This API is used to enable the cloud slicing feature, completing audio and video slicing tasks in the room, and uploading them to the specified cloud storage.
//
// This API is used to achieve the following goals:
//
// * This API is used to specify the slicing parameter (SliceParams) to define the blocklist or allowlist of the anchors that require slicing.
//
// * This API is used to specify the storage parameter (SliceStorageParams) to specify the cloud storage you want to upload to. Currently, Tencent Cloud Object Storage (COS) and third-party AWS are supported.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudSliceTaskWithContext(ctx context.Context, request *CreateCloudSliceTaskRequest) (response *CreateCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewCreateCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudModerationRequest() (request *DeleteCloudModerationRequest) {
    request = &DeleteCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudModeration")
    
    
    return
}

func NewDeleteCloudModerationResponse() (response *DeleteCloudModerationResponse) {
    response = &DeleteCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudModeration
// This API is used to stop submission for moderation after the cloud moderation task is successfully started.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudModeration(request *DeleteCloudModerationRequest) (response *DeleteCloudModerationResponse, err error) {
    return c.DeleteCloudModerationWithContext(context.Background(), request)
}

// DeleteCloudModeration
// This API is used to stop submission for moderation after the cloud moderation task is successfully started.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudModerationWithContext(ctx context.Context, request *DeleteCloudModerationRequest) (response *DeleteCloudModerationResponse, err error) {
    if request == nil {
        request = NewDeleteCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudModerationResponse()
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
// This API is used to stop a recording task. If a task is stopped successfully, but the uploading of recording files has not completed, the backend will continue to upload the files and will notify you via a callback when the upload is completed.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudRecording(request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    return c.DeleteCloudRecordingWithContext(context.Background(), request)
}

// DeleteCloudRecording
// This API is used to stop a recording task. If a task is stopped successfully, but the uploading of recording files has not completed, the backend will continue to upload the files and will notify you via a callback when the upload is completed.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudRecordingWithContext(ctx context.Context, request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDeleteCloudRecordingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudSliceTaskRequest() (request *DeleteCloudSliceTaskRequest) {
    request = &DeleteCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudSliceTask")
    
    
    return
}

func NewDeleteCloudSliceTaskResponse() (response *DeleteCloudSliceTaskResponse) {
    response = &DeleteCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudSliceTask
// This API is used to stop the slicing task after it is started. Successfully stopping the slicing does not mean that all files are fully transmitted; if the transmission is not completed, the backend will continue to upload files. After the upload is successful, a notification is sent to the customer, prompting that all files have been transmitted, through the event callback.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudSliceTask(request *DeleteCloudSliceTaskRequest) (response *DeleteCloudSliceTaskResponse, err error) {
    return c.DeleteCloudSliceTaskWithContext(context.Background(), request)
}

// DeleteCloudSliceTask
// This API is used to stop the slicing task after it is started. Successfully stopping the slicing does not mean that all files are fully transmitted; if the transmission is not completed, the backend will continue to upload files. After the upload is successful, a notification is sent to the customer, prompting that all files have been transmitted, through the event callback.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudSliceTaskWithContext(ctx context.Context, request *DeleteCloudSliceTaskRequest) (response *DeleteCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewDeleteCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIConversationRequest() (request *DescribeAIConversationRequest) {
    request = &DescribeAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAIConversation")
    
    
    return
}

func NewDescribeAIConversationResponse() (response *DescribeAIConversationResponse) {
    response = &DescribeAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIConversation
// Describe the AI conversation task status
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAIConversation(request *DescribeAIConversationRequest) (response *DescribeAIConversationResponse, err error) {
    return c.DescribeAIConversationWithContext(context.Background(), request)
}

// DescribeAIConversation
// Describe the AI conversation task status
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAIConversationWithContext(ctx context.Context, request *DescribeAIConversationRequest) (response *DescribeAIConversationResponse, err error) {
    if request == nil {
        request = NewDescribeAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAITranscriptionRequest() (request *DescribeAITranscriptionRequest) {
    request = &DescribeAITranscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAITranscription")
    
    
    return
}

func NewDescribeAITranscriptionResponse() (response *DescribeAITranscriptionResponse) {
    response = &DescribeAITranscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAITranscription
// Describe AI transcription task status
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAITranscription(request *DescribeAITranscriptionRequest) (response *DescribeAITranscriptionResponse, err error) {
    return c.DescribeAITranscriptionWithContext(context.Background(), request)
}

// DescribeAITranscription
// Describe AI transcription task status
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAITranscriptionWithContext(ctx context.Context, request *DescribeAITranscriptionRequest) (response *DescribeAITranscriptionResponse, err error) {
    if request == nil {
        request = NewDescribeAITranscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeAITranscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAITranscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAITranscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallDetailInfoRequest() (request *DescribeCallDetailInfoRequest) {
    request = &DescribeCallDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCallDetailInfo")
    
    
    return
}

func NewDescribeCallDetailInfoResponse() (response *DescribeCallDetailInfoResponse) {
    response = &DescribeCallDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallDetailInfo
// This API (the old `DescribeCallDetail`) is used to query the user list and call quality data of a specified time range in the last 14 days. If `DataType` is not null, the data of up to six users during a period of up to one hour can be queried (the period can start and end on different days). If `DataType` is null, the data of up to 100 users can be returned per page (the value of `PageSize` cannot exceed 100). Six users are queried by default. The period queried cannot exceed four hours. This API is used to query call quality and is not recommended for billing purposes.
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/54481?product=pricing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetailInfo(request *DescribeCallDetailInfoRequest) (response *DescribeCallDetailInfoResponse, err error) {
    return c.DescribeCallDetailInfoWithContext(context.Background(), request)
}

// DescribeCallDetailInfo
// This API (the old `DescribeCallDetail`) is used to query the user list and call quality data of a specified time range in the last 14 days. If `DataType` is not null, the data of up to six users during a period of up to one hour can be queried (the period can start and end on different days). If `DataType` is null, the data of up to 100 users can be returned per page (the value of `PageSize` cannot exceed 100). Six users are queried by default. The period queried cannot exceed four hours. This API is used to query call quality and is not recommended for billing purposes.
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/54481?product=pricing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetailInfoWithContext(ctx context.Context, request *DescribeCallDetailInfoRequest) (response *DescribeCallDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCallDetailInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudModerationRequest() (request *DescribeCloudModerationRequest) {
    request = &DescribeCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudModeration")
    
    
    return
}

func NewDescribeCloudModerationResponse() (response *DescribeCloudModerationResponse) {
    response = &DescribeCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudModeration
// This API is used to query the status of the moderation task and information about the subscription blocklist and allowlist after the task is started, which is valid only when the task is in progress. An error will be returned if the task is exited.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudModeration(request *DescribeCloudModerationRequest) (response *DescribeCloudModerationResponse, err error) {
    return c.DescribeCloudModerationWithContext(context.Background(), request)
}

// DescribeCloudModeration
// This API is used to query the status of the moderation task and information about the subscription blocklist and allowlist after the task is started, which is valid only when the task is in progress. An error will be returned if the task is exited.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudModerationWithContext(ctx context.Context, request *DescribeCloudModerationRequest) (response *DescribeCloudModerationResponse, err error) {
    if request == nil {
        request = NewDescribeCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudModerationResponse()
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
// If a recording file is being uploaded to VOD, the response parameter `StorageFileList` will not contain the information of the recording file. Please listen for the recording file callback to get the information.
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
// If a recording file is being uploaded to VOD, the response parameter `StorageFileList` will not contain the information of the recording file. Please listen for the recording file callback to get the information.
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
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudSliceTaskRequest() (request *DescribeCloudSliceTaskRequest) {
    request = &DescribeCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudSliceTask")
    
    
    return
}

func NewDescribeCloudSliceTaskResponse() (response *DescribeCloudSliceTaskResponse) {
    response = &DescribeCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudSliceTask
// This API is used to query the status of the slicing task after it is started, which is valid only when the task is in progress. An error will be returned if the task is exited.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudSliceTask(request *DescribeCloudSliceTaskRequest) (response *DescribeCloudSliceTaskResponse, err error) {
    return c.DescribeCloudSliceTaskWithContext(context.Background(), request)
}

// DescribeCloudSliceTask
// This API is used to query the status of the slicing task after it is started, which is valid only when the task is in progress. An error will be returned if the task is exited.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudSliceTaskWithContext(ctx context.Context, request *DescribeCloudSliceTaskRequest) (response *DescribeCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMixTranscodingUsageRequest() (request *DescribeMixTranscodingUsageRequest) {
    request = &DescribeMixTranscodingUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeMixTranscodingUsage")
    
    
    return
}

func NewDescribeMixTranscodingUsageResponse() (response *DescribeMixTranscodingUsageResponse) {
    response = &DescribeMixTranscodingUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMixTranscodingUsage
// This API is used to query your usage of TRTC’s On-Cloud MixTranscoding service.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeMixTranscodingUsage(request *DescribeMixTranscodingUsageRequest) (response *DescribeMixTranscodingUsageResponse, err error) {
    return c.DescribeMixTranscodingUsageWithContext(context.Background(), request)
}

// DescribeMixTranscodingUsage
// This API is used to query your usage of TRTC’s On-Cloud MixTranscoding service.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeMixTranscodingUsageWithContext(ctx context.Context, request *DescribeMixTranscodingUsageRequest) (response *DescribeMixTranscodingUsageResponse, err error) {
    if request == nil {
        request = NewDescribeMixTranscodingUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeMixTranscodingUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMixTranscodingUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMixTranscodingUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingUsageRequest() (request *DescribeRecordingUsageRequest) {
    request = &DescribeRecordingUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRecordingUsage")
    
    
    return
}

func NewDescribeRecordingUsageResponse() (response *DescribeRecordingUsageResponse) {
    response = &DescribeRecordingUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordingUsage
// This API is used to query your TRTC recording usage.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordingUsage(request *DescribeRecordingUsageRequest) (response *DescribeRecordingUsageResponse, err error) {
    return c.DescribeRecordingUsageWithContext(context.Background(), request)
}

// DescribeRecordingUsage
// This API is used to query your TRTC recording usage.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordingUsageWithContext(ctx context.Context, request *DescribeRecordingUsageRequest) (response *DescribeRecordingUsageResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRecordingUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordingUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordingUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelayUsageRequest() (request *DescribeRelayUsageRequest) {
    request = &DescribeRelayUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRelayUsage")
    
    
    return
}

func NewDescribeRelayUsageResponse() (response *DescribeRelayUsageResponse) {
    response = &DescribeRelayUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRelayUsage
// This API is used to query your usage of TRTC’s relay to CDN service.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRelayUsage(request *DescribeRelayUsageRequest) (response *DescribeRelayUsageResponse, err error) {
    return c.DescribeRelayUsageWithContext(context.Background(), request)
}

// DescribeRelayUsage
// This API is used to query your usage of TRTC’s relay to CDN service.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRelayUsageWithContext(ctx context.Context, request *DescribeRelayUsageRequest) (response *DescribeRelayUsageResponse, err error) {
    if request == nil {
        request = NewDescribeRelayUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRelayUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelayUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRelayUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInfoRequest() (request *DescribeRoomInfoRequest) {
    request = &DescribeRoomInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRoomInfo")
    
    
    return
}

func NewDescribeRoomInfoResponse() (response *DescribeRoomInfoResponse) {
    response = &DescribeRoomInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomInfo
// This API (the old `DescribeRoomInformation`) is used to query the rooms of an application (`SDKAppID`) in the last 14 days. Up to 100 records can be returned per call (10 are returned by default).
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/54481
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInfo(request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    return c.DescribeRoomInfoWithContext(context.Background(), request)
}

// DescribeRoomInfo
// This API (the old `DescribeRoomInformation`) is used to query the rooms of an application (`SDKAppID`) in the last 14 days. Up to 100 records can be returned per call (10 are returned by default).
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/54481
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInfoWithContext(ctx context.Context, request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRoomInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScaleInfoRequest() (request *DescribeScaleInfoRequest) {
    request = &DescribeScaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeScaleInfo")
    
    
    return
}

func NewDescribeScaleInfoResponse() (response *DescribeScaleInfoResponse) {
    response = &DescribeScaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScaleInfo
// This API (the old `DescribeHistoryScale`) is used to query the daily number of rooms and users of an application (`SDKAppID`) in the last 14 days. Data for the current day cannot be queried.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeScaleInfo(request *DescribeScaleInfoRequest) (response *DescribeScaleInfoResponse, err error) {
    return c.DescribeScaleInfoWithContext(context.Background(), request)
}

// DescribeScaleInfo
// This API (the old `DescribeHistoryScale`) is used to query the daily number of rooms and users of an application (`SDKAppID`) in the last 14 days. Data for the current day cannot be queried.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeScaleInfoWithContext(ctx context.Context, request *DescribeScaleInfoRequest) (response *DescribeScaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeScaleInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeScaleInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScaleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamIngestRequest() (request *DescribeStreamIngestRequest) {
    request = &DescribeStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeStreamIngest")
    
    
    return
}

func NewDescribeStreamIngestResponse() (response *DescribeStreamIngestResponse) {
    response = &DescribeStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamIngest
// You can query the status of the Relay task.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeStreamIngest(request *DescribeStreamIngestRequest) (response *DescribeStreamIngestResponse, err error) {
    return c.DescribeStreamIngestWithContext(context.Background(), request)
}

// DescribeStreamIngest
// You can query the status of the Relay task.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeStreamIngestWithContext(ctx context.Context, request *DescribeStreamIngestRequest) (response *DescribeStreamIngestResponse, err error) {
    if request == nil {
        request = NewDescribeStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCMarketQualityDataRequest() (request *DescribeTRTCMarketQualityDataRequest) {
    request = &DescribeTRTCMarketQualityDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCMarketQualityData")
    
    
    return
}

func NewDescribeTRTCMarketQualityDataResponse() (response *DescribeTRTCMarketQualityDataResponse) {
    response = &DescribeTRTCMarketQualityDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCMarketQualityData
// Query TRTC Monitoring Dashboard - Data Dashboard Quality Metrics (including the following metrics)
//
// joinSuccessRate: Join channel success rate.
//
// joinSuccessIn5sRate: Join channel success rate within 5s.
//
// audioFreezeRate: Audio stutter rate.
//
// videoFreezeRate: Video stutter rate.
//
// networkDelay: Lag rate.
//
// Note:
//
// 1. To call the API, you need to activate the monitoring dashboard Standard Edition and Premium Edition, the monitoring dashboard Free Edition does not support calling. Monitoring dashboard version features and billing overview: https://trtc.io/document/54481.
//
// 2. The query time range depends on the monitoring dashboard function version, premium edition can query the last 30 days.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCMarketQualityData(request *DescribeTRTCMarketQualityDataRequest) (response *DescribeTRTCMarketQualityDataResponse, err error) {
    return c.DescribeTRTCMarketQualityDataWithContext(context.Background(), request)
}

// DescribeTRTCMarketQualityData
// Query TRTC Monitoring Dashboard - Data Dashboard Quality Metrics (including the following metrics)
//
// joinSuccessRate: Join channel success rate.
//
// joinSuccessIn5sRate: Join channel success rate within 5s.
//
// audioFreezeRate: Audio stutter rate.
//
// videoFreezeRate: Video stutter rate.
//
// networkDelay: Lag rate.
//
// Note:
//
// 1. To call the API, you need to activate the monitoring dashboard Standard Edition and Premium Edition, the monitoring dashboard Free Edition does not support calling. Monitoring dashboard version features and billing overview: https://trtc.io/document/54481.
//
// 2. The query time range depends on the monitoring dashboard function version, premium edition can query the last 30 days.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCMarketQualityDataWithContext(ctx context.Context, request *DescribeTRTCMarketQualityDataRequest) (response *DescribeTRTCMarketQualityDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCMarketQualityDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCMarketQualityData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCMarketQualityData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCMarketQualityDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCMarketScaleDataRequest() (request *DescribeTRTCMarketScaleDataRequest) {
    request = &DescribeTRTCMarketScaleDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCMarketScaleData")
    
    
    return
}

func NewDescribeTRTCMarketScaleDataResponse() (response *DescribeTRTCMarketScaleDataResponse) {
    response = &DescribeTRTCMarketScaleDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCMarketScaleData
// Query TRTC Monitoring Dashboard - Data Dashboard Scale Metrics (will return userCount, roomCount, peakCurrentUsers, peakCurrentChannels)
//
// - userCount: number of users in the call,
//
// - roomCount: number of rooms in the call, counted as one call channel from the time a user joins the channel to the time all users leave the channel.
//
// - peakCurrentChannels: peak number of channels online at the same time.
//
// - peakCurrentUsers: peak number of users online at the same time.
//
// Note:
//
// 1. To call the interface, you need to activate the monitoring dashboard Standard Edition and Premium Edition, the monitoring dashboard Free Edition does not support calling, for monitoring dashboard version features and billing overview: https://trtc.io/document/54481.
//
// 2. The query time range depends on the monitoring dashboard function version, premium edition can query up to 60 days.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCMarketScaleData(request *DescribeTRTCMarketScaleDataRequest) (response *DescribeTRTCMarketScaleDataResponse, err error) {
    return c.DescribeTRTCMarketScaleDataWithContext(context.Background(), request)
}

// DescribeTRTCMarketScaleData
// Query TRTC Monitoring Dashboard - Data Dashboard Scale Metrics (will return userCount, roomCount, peakCurrentUsers, peakCurrentChannels)
//
// - userCount: number of users in the call,
//
// - roomCount: number of rooms in the call, counted as one call channel from the time a user joins the channel to the time all users leave the channel.
//
// - peakCurrentChannels: peak number of channels online at the same time.
//
// - peakCurrentUsers: peak number of users online at the same time.
//
// Note:
//
// 1. To call the interface, you need to activate the monitoring dashboard Standard Edition and Premium Edition, the monitoring dashboard Free Edition does not support calling, for monitoring dashboard version features and billing overview: https://trtc.io/document/54481.
//
// 2. The query time range depends on the monitoring dashboard function version, premium edition can query up to 60 days.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCMarketScaleDataWithContext(ctx context.Context, request *DescribeTRTCMarketScaleDataRequest) (response *DescribeTRTCMarketScaleDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCMarketScaleDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCMarketScaleData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCMarketScaleData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCMarketScaleDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCRealTimeQualityDataRequest() (request *DescribeTRTCRealTimeQualityDataRequest) {
    request = &DescribeTRTCRealTimeQualityDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCRealTimeQualityData")
    
    
    return
}

func NewDescribeTRTCRealTimeQualityDataResponse() (response *DescribeTRTCRealTimeQualityDataResponse) {
    response = &DescribeTRTCRealTimeQualityDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCRealTimeQualityData
// Query TRTC Monitoring Dashboard - Real-Time Monitoring Quality Metrics (return the following metrics)
//
//  -Video stutter rate
//
//  -Audio stutter rate
//
//  Note:
//
//  1. To call the API, you need to activate the Monitoring Dashboard Standard Edition and Premium Edition. The Monitoring Dashboard Free Edition does not support calling. For monitoring dashboard version features and billing overview, please visit: https://trtc.io/document/54481.
//
//  2. The query time range depends on the monitoring dashboard function version. The premium edition can query up to 1 hours
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCRealTimeQualityData(request *DescribeTRTCRealTimeQualityDataRequest) (response *DescribeTRTCRealTimeQualityDataResponse, err error) {
    return c.DescribeTRTCRealTimeQualityDataWithContext(context.Background(), request)
}

// DescribeTRTCRealTimeQualityData
// Query TRTC Monitoring Dashboard - Real-Time Monitoring Quality Metrics (return the following metrics)
//
//  -Video stutter rate
//
//  -Audio stutter rate
//
//  Note:
//
//  1. To call the API, you need to activate the Monitoring Dashboard Standard Edition and Premium Edition. The Monitoring Dashboard Free Edition does not support calling. For monitoring dashboard version features and billing overview, please visit: https://trtc.io/document/54481.
//
//  2. The query time range depends on the monitoring dashboard function version. The premium edition can query up to 1 hours
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCRealTimeQualityDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeQualityDataRequest) (response *DescribeTRTCRealTimeQualityDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeQualityDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCRealTimeQualityData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCRealTimeQualityData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCRealTimeQualityDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCRealTimeScaleDataRequest() (request *DescribeTRTCRealTimeScaleDataRequest) {
    request = &DescribeTRTCRealTimeScaleDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCRealTimeScaleData")
    
    
    return
}

func NewDescribeTRTCRealTimeScaleDataResponse() (response *DescribeTRTCRealTimeScaleDataResponse) {
    response = &DescribeTRTCRealTimeScaleDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCRealTimeScaleData
// Query TRTC Monitoring Dashboard - Real-Time Monitoring Scale Metrics (the following metrics will be returned) -userCount (Online users) -roomCount (Online rooms) Note: 1. To call the interface, you need to activate the monitoring dashboard Standard Edition and Premium Edition, the monitoring dashboard Free Edition does not support calling. For monitoring dashboard version features and billing overview, please visit: https://trtc.io/document/54481. 2. The query time range depends on the function version of the monitoring dashboard. The premium edition can query the last 1 hours
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleData(request *DescribeTRTCRealTimeScaleDataRequest) (response *DescribeTRTCRealTimeScaleDataResponse, err error) {
    return c.DescribeTRTCRealTimeScaleDataWithContext(context.Background(), request)
}

// DescribeTRTCRealTimeScaleData
// Query TRTC Monitoring Dashboard - Real-Time Monitoring Scale Metrics (the following metrics will be returned) -userCount (Online users) -roomCount (Online rooms) Note: 1. To call the interface, you need to activate the monitoring dashboard Standard Edition and Premium Edition, the monitoring dashboard Free Edition does not support calling. For monitoring dashboard version features and billing overview, please visit: https://trtc.io/document/54481. 2. The query time range depends on the function version of the monitoring dashboard. The premium edition can query the last 1 hours
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeScaleDataRequest) (response *DescribeTRTCRealTimeScaleDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeScaleDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCRealTimeScaleData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCRealTimeScaleData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCRealTimeScaleDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcRoomUsageRequest() (request *DescribeTrtcRoomUsageRequest) {
    request = &DescribeTrtcRoomUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcRoomUsage")
    
    
    return
}

func NewDescribeTrtcRoomUsageResponse() (response *DescribeTrtcRoomUsageResponse) {
    response = &DescribeTrtcRoomUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrtcRoomUsage
// This API is used to query usage data grouped by room.
//
// - The queried period cannot exceed 24 hours. If the period spans two different days, the data returned may not be accurate due to a delay in data collection. You can make multiple calls to query the usage on different days.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is one call every 15 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTrtcRoomUsage(request *DescribeTrtcRoomUsageRequest) (response *DescribeTrtcRoomUsageResponse, err error) {
    return c.DescribeTrtcRoomUsageWithContext(context.Background(), request)
}

// DescribeTrtcRoomUsage
// This API is used to query usage data grouped by room.
//
// - The queried period cannot exceed 24 hours. If the period spans two different days, the data returned may not be accurate due to a delay in data collection. You can make multiple calls to query the usage on different days.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is one call every 15 seconds.
//
// error code that may be returned:
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTrtcRoomUsageWithContext(ctx context.Context, request *DescribeTrtcRoomUsageRequest) (response *DescribeTrtcRoomUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcRoomUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTrtcRoomUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrtcRoomUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrtcRoomUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcUsageRequest() (request *DescribeTrtcUsageRequest) {
    request = &DescribeTrtcUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcUsage")
    
    
    return
}

func NewDescribeTrtcUsageResponse() (response *DescribeTrtcUsageResponse) {
    response = &DescribeTrtcUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrtcUsage
// This API is used to query your TRTC audio/video duration.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcUsage(request *DescribeTrtcUsageRequest) (response *DescribeTrtcUsageResponse, err error) {
    return c.DescribeTrtcUsageWithContext(context.Background(), request)
}

// DescribeTrtcUsage
// This API is used to query your TRTC audio/video duration.
//
// - If the period queried is one day or shorter, the statistics returned are on a five-minute basis. If the period queried is longer than one day, the statistics returned are on a daily basis.
//
// - The period queried per request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - You can use this API to query your historical usage or to reconcile data, but we do not recommend you use it for crucial business logic.
//
// - The rate limit of this API is five calls per second.
//
// error code that may be returned:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcUsageWithContext(ctx context.Context, request *DescribeTrtcUsageRequest) (response *DescribeTrtcUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTrtcUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrtcUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrtcUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnusualEventRequest() (request *DescribeUnusualEventRequest) {
    request = &DescribeUnusualEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUnusualEvent")
    
    
    return
}

func NewDescribeUnusualEventResponse() (response *DescribeUnusualEventResponse) {
    response = &DescribeUnusualEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnusualEvent
// This API (the old `DescribeAbnormalEvent`) is used to query up to 20 random abnormal user experiences of an application (`SDKAppID`) in the last 14 days. The start and end time can be on two different days, but they cannot be more than one hour apart.
//
// For details about the error events, see https://intl.cloud.tencent.com/document/product/647/44916?from_cn_redirect=1
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUnusualEvent(request *DescribeUnusualEventRequest) (response *DescribeUnusualEventResponse, err error) {
    return c.DescribeUnusualEventWithContext(context.Background(), request)
}

// DescribeUnusualEvent
// This API (the old `DescribeAbnormalEvent`) is used to query up to 20 random abnormal user experiences of an application (`SDKAppID`) in the last 14 days. The start and end time can be on two different days, but they cannot be more than one hour apart.
//
// For details about the error events, see https://intl.cloud.tencent.com/document/product/647/44916?from_cn_redirect=1
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUnusualEventWithContext(ctx context.Context, request *DescribeUnusualEventRequest) (response *DescribeUnusualEventResponse, err error) {
    if request == nil {
        request = NewDescribeUnusualEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeUnusualEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnusualEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnusualEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserEventRequest() (request *DescribeUserEventRequest) {
    request = &DescribeUserEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUserEvent")
    
    
    return
}

func NewDescribeUserEventResponse() (response *DescribeUserEventResponse) {
    response = &DescribeUserEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserEvent
// This API (the old `DescribeDetailEvent`) is used to query the events of a call in the last 14 days, including user entry and exit, turning the camera on/off, etc.
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/54481?product=pricing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeUserEvent(request *DescribeUserEventRequest) (response *DescribeUserEventResponse, err error) {
    return c.DescribeUserEventWithContext(context.Background(), request)
}

// DescribeUserEvent
// This API (the old `DescribeDetailEvent`) is used to query the events of a call in the last 14 days, including user entry and exit, turning the camera on/off, etc.
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/54481?product=pricing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeUserEventWithContext(ctx context.Context, request *DescribeUserEventRequest) (response *DescribeUserEventResponse, err error) {
    if request == nil {
        request = NewDescribeUserEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeUserEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUserInfo")
    
    
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserInfo
// This API (the old `DescribeUserInformation`) is used to query the user list of a specified time range (up to four hours) in the last 14 days. The data of up to 100 users can be returned per page (six are returned by default).
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/60214?product=pricing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    return c.DescribeUserInfoWithContext(context.Background(), request)
}

// DescribeUserInfo
// This API (the old `DescribeUserInformation`) is used to query the user list of a specified time range (up to four hours) in the last 14 days. The data of up to 100 users can be returned per page (six are returned by default).
//
// **Note**:
//
// 1. You can use this API to query historical data or for reconciliation purposes, but we do not recommend you use it for crucial business logic.
//
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://trtc.io/document/60214?product=pricing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInfoWithContext(ctx context.Context, request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebRecordRequest() (request *DescribeWebRecordRequest) {
    request = &DescribeWebRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeWebRecord")
    
    
    return
}

func NewDescribeWebRecordResponse() (response *DescribeWebRecordResponse) {
    response = &DescribeWebRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebRecord
// Queries the status of a web-page recording task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeWebRecord(request *DescribeWebRecordRequest) (response *DescribeWebRecordResponse, err error) {
    return c.DescribeWebRecordWithContext(context.Background(), request)
}

// DescribeWebRecord
// Queries the status of a web-page recording task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeWebRecordWithContext(ctx context.Context, request *DescribeWebRecordRequest) (response *DescribeWebRecordResponse, err error) {
    if request == nil {
        request = NewDescribeWebRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeWebRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebRecordResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DismissRoom")
    
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
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DismissRoomByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoomByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudModerationRequest() (request *ModifyCloudModerationRequest) {
    request = &ModifyCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudModeration")
    
    
    return
}

func NewModifyCloudModerationResponse() (response *ModifyCloudModerationResponse) {
    response = &ModifyCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudModeration
// This API is used to update the subscription blocklist and allowlist after the cloud moderation task is successfully started.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudModeration(request *ModifyCloudModerationRequest) (response *ModifyCloudModerationResponse, err error) {
    return c.ModifyCloudModerationWithContext(context.Background(), request)
}

// ModifyCloudModeration
// This API is used to update the subscription blocklist and allowlist after the cloud moderation task is successfully started.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudModerationWithContext(ctx context.Context, request *ModifyCloudModerationRequest) (response *ModifyCloudModerationResponse, err error) {
    if request == nil {
        request = NewModifyCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudModerationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudSliceTaskRequest() (request *ModifyCloudSliceTaskRequest) {
    request = &ModifyCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudSliceTask")
    
    
    return
}

func NewModifyCloudSliceTaskResponse() (response *ModifyCloudSliceTaskResponse) {
    response = &ModifyCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudSliceTask
// This API is used to update the slicing task after it is started. It can be used to update the allowlist or blocklist for the specified subscription stream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudSliceTask(request *ModifyCloudSliceTaskRequest) (response *ModifyCloudSliceTaskResponse, err error) {
    return c.ModifyCloudSliceTaskWithContext(context.Background(), request)
}

// ModifyCloudSliceTask
// This API is used to update the slicing task after it is started. It can be used to update the allowlist or blocklist for the specified subscription stream.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudSliceTaskWithContext(ctx context.Context, request *ModifyCloudSliceTaskRequest) (response *ModifyCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewModifyCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudSliceTaskResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "RemoveUser")
    
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
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "RemoveUserByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewSetUserBlockedRequest() (request *SetUserBlockedRequest) {
    request = &SetUserBlockedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "SetUserBlocked")
    
    
    return
}

func NewSetUserBlockedResponse() (response *SetUserBlockedResponse) {
    response = &SetUserBlockedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetUserBlocked
// This API is used to disable or enable the audio and video of a user. It can be used by an anchor, room owner, or admin to block or unblock a user. It supports platforms including Android, iOS, Windows, macOS, web, and WeChat Mini Program. Use this API if the room ID is a number.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetUserBlocked(request *SetUserBlockedRequest) (response *SetUserBlockedResponse, err error) {
    return c.SetUserBlockedWithContext(context.Background(), request)
}

// SetUserBlocked
// This API is used to disable or enable the audio and video of a user. It can be used by an anchor, room owner, or admin to block or unblock a user. It supports platforms including Android, iOS, Windows, macOS, web, and WeChat Mini Program. Use this API if the room ID is a number.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetUserBlockedWithContext(ctx context.Context, request *SetUserBlockedRequest) (response *SetUserBlockedResponse, err error) {
    if request == nil {
        request = NewSetUserBlockedRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "SetUserBlocked")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetUserBlocked require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetUserBlockedResponse()
    err = c.Send(request, response)
    return
}

func NewSetUserBlockedByStrRoomIdRequest() (request *SetUserBlockedByStrRoomIdRequest) {
    request = &SetUserBlockedByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "SetUserBlockedByStrRoomId")
    
    
    return
}

func NewSetUserBlockedByStrRoomIdResponse() (response *SetUserBlockedByStrRoomIdResponse) {
    response = &SetUserBlockedByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetUserBlockedByStrRoomId
// This API allows an anchor, room owner, admin to mute/unmute a user. It can be used on platforms including Android, iOS, Windows, macOS, web, and WeChat Mini Program. Use this API when the room ID is a string.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetUserBlockedByStrRoomId(request *SetUserBlockedByStrRoomIdRequest) (response *SetUserBlockedByStrRoomIdResponse, err error) {
    return c.SetUserBlockedByStrRoomIdWithContext(context.Background(), request)
}

// SetUserBlockedByStrRoomId
// This API allows an anchor, room owner, admin to mute/unmute a user. It can be used on platforms including Android, iOS, Windows, macOS, web, and WeChat Mini Program. Use this API when the room ID is a string.
//
// error code that may be returned:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  FAILEDOPERATION_USERNOTEXIST = "FailedOperation.UserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetUserBlockedByStrRoomIdWithContext(ctx context.Context, request *SetUserBlockedByStrRoomIdRequest) (response *SetUserBlockedByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewSetUserBlockedByStrRoomIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "SetUserBlockedByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetUserBlockedByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetUserBlockedByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStartAIConversationRequest() (request *StartAIConversationRequest) {
    request = &StartAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartAIConversation")
    
    
    return
}

func NewStartAIConversationResponse() (response *StartAIConversationResponse) {
    response = &StartAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAIConversation
// Initiate AI conversation task, where the AI bot enters the TRTC room to engage in AI conversation with specified members in the room. This is suitable for scenarios such as intelligent customer service and AI language teachers. The TRTC AI conversation feature has built-in speech-to-text capabilities , allowing customers to flexibly specify third-party AI model (LLM) services and text-to-speech (TTS) services. For more [feature details](https://cloud.tencent.com/document/product/647/108901).
//
// error code that may be returned:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAIConversation(request *StartAIConversationRequest) (response *StartAIConversationResponse, err error) {
    return c.StartAIConversationWithContext(context.Background(), request)
}

// StartAIConversation
// Initiate AI conversation task, where the AI bot enters the TRTC room to engage in AI conversation with specified members in the room. This is suitable for scenarios such as intelligent customer service and AI language teachers. The TRTC AI conversation feature has built-in speech-to-text capabilities , allowing customers to flexibly specify third-party AI model (LLM) services and text-to-speech (TTS) services. For more [feature details](https://cloud.tencent.com/document/product/647/108901).
//
// error code that may be returned:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAIConversationWithContext(ctx context.Context, request *StartAIConversationRequest) (response *StartAIConversationResponse, err error) {
    if request == nil {
        request = NewStartAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewStartAITranscriptionRequest() (request *StartAITranscriptionRequest) {
    request = &StartAITranscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartAITranscription")
    
    
    return
}

func NewStartAITranscriptionResponse() (response *StartAITranscriptionResponse) {
    response = &StartAITranscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAITranscription
// Initiate the transcription bot. The backend will pull the stream through the bot to perform real-time speech recognition and deliver subtitles and transcription messages. The transcription bot supports two stream pulling modes, controlled by the `TranscriptionMode` field:
//
// - Pull the stream of the entire room.
//
// - Pull the stream of a specific user.
//
// 
//
// The server delivers subtitles and transcription messages in real-time through TRTC's custom messages, with `CmdId` fixed at 1. The client only needs to listen for the callback of custom messages. For example, see the [C++ callback](https://cloud.tencent.com/document/product/647/79637#4cd82f4edb24992a15a25187089e1565). Other clients, such as Android, Web, etc., can also be found at the same link.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  INVALIDPARAMETER_USERSIGNOTADMIN = "InvalidParameter.UserSigNotAdmin"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAITranscription(request *StartAITranscriptionRequest) (response *StartAITranscriptionResponse, err error) {
    return c.StartAITranscriptionWithContext(context.Background(), request)
}

// StartAITranscription
// Initiate the transcription bot. The backend will pull the stream through the bot to perform real-time speech recognition and deliver subtitles and transcription messages. The transcription bot supports two stream pulling modes, controlled by the `TranscriptionMode` field:
//
// - Pull the stream of the entire room.
//
// - Pull the stream of a specific user.
//
// 
//
// The server delivers subtitles and transcription messages in real-time through TRTC's custom messages, with `CmdId` fixed at 1. The client only needs to listen for the callback of custom messages. For example, see the [C++ callback](https://cloud.tencent.com/document/product/647/79637#4cd82f4edb24992a15a25187089e1565). Other clients, such as Android, Web, etc., can also be found at the same link.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  INVALIDPARAMETER_USERSIGNOTADMIN = "InvalidParameter.UserSigNotAdmin"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAITranscriptionWithContext(ctx context.Context, request *StartAITranscriptionRequest) (response *StartAITranscriptionResponse, err error) {
    if request == nil {
        request = NewStartAITranscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartAITranscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAITranscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAITranscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishCdnStreamRequest() (request *StartPublishCdnStreamRequest) {
    request = &StartPublishCdnStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartPublishCdnStream")
    
    
    return
}

func NewStartPublishCdnStreamResponse() (response *StartPublishCdnStreamResponse) {
    response = &StartPublishCdnStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartPublishCdnStream
// **API Description**
//
// 
//
// This API starts a stream mixing and relaying task. This API mixes multiple audio/video streams from a TRTC room into a single stream, encodes it, and then pushes it to CDN server or publishs it into the TRTC room. It also supports relaying a single stream from a TRTC room directly without transcoding.
//
// 
//
// After success, the API returns a globally unique TaskID. You will need this TaskId in later operations such as updating or stopping the task.
//
// 
//
// For more details, refer to the document:  [Feature Description](https://trtc.io/zh/document/47858?product=rtcengine ) and [FAQs](https://trtc.io/zh/document/36058?product=rtcengine&menulabel=core%20sdk&platform=web) .
//
// 
//
// Note: You can enable the relay to CDN in the console to monitor events under the CDN relay status. For callback details, see: [Relay to CDN Callback Description](https://trtc.io/zh/document/54913?product=rtcengine&menulabel=core%20sdk&platform=web ) .
//
// 
//
// Starting a relay task may incur the following fees:
//
// MCU stream mixing and transcoding fees: [See Cloud Stream Mixing and Transcoding Pricing](https://trtc.io/zh/document/47631 ) .
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartPublishCdnStream(request *StartPublishCdnStreamRequest) (response *StartPublishCdnStreamResponse, err error) {
    return c.StartPublishCdnStreamWithContext(context.Background(), request)
}

// StartPublishCdnStream
// **API Description**
//
// 
//
// This API starts a stream mixing and relaying task. This API mixes multiple audio/video streams from a TRTC room into a single stream, encodes it, and then pushes it to CDN server or publishs it into the TRTC room. It also supports relaying a single stream from a TRTC room directly without transcoding.
//
// 
//
// After success, the API returns a globally unique TaskID. You will need this TaskId in later operations such as updating or stopping the task.
//
// 
//
// For more details, refer to the document:  [Feature Description](https://trtc.io/zh/document/47858?product=rtcengine ) and [FAQs](https://trtc.io/zh/document/36058?product=rtcengine&menulabel=core%20sdk&platform=web) .
//
// 
//
// Note: You can enable the relay to CDN in the console to monitor events under the CDN relay status. For callback details, see: [Relay to CDN Callback Description](https://trtc.io/zh/document/54913?product=rtcengine&menulabel=core%20sdk&platform=web ) .
//
// 
//
// Starting a relay task may incur the following fees:
//
// MCU stream mixing and transcoding fees: [See Cloud Stream Mixing and Transcoding Pricing](https://trtc.io/zh/document/47631 ) .
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartPublishCdnStreamWithContext(ctx context.Context, request *StartPublishCdnStreamRequest) (response *StartPublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewStartPublishCdnStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartPublishCdnStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamIngestRequest() (request *StartStreamIngestRequest) {
    request = &StartStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartStreamIngest")
    
    
    return
}

func NewStartStreamIngestResponse() (response *StartStreamIngestResponse) {
    response = &StartStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartStreamIngest
// Push an online media stream to the TRTC room.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_NOTRTMPFUNCTION = "FailedOperation.NotRtmpFunction"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartStreamIngest(request *StartStreamIngestRequest) (response *StartStreamIngestResponse, err error) {
    return c.StartStreamIngestWithContext(context.Background(), request)
}

// StartStreamIngest
// Push an online media stream to the TRTC room.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_NOTRTMPFUNCTION = "FailedOperation.NotRtmpFunction"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartStreamIngestWithContext(ctx context.Context, request *StartStreamIngestRequest) (response *StartStreamIngestResponse, err error) {
    if request == nil {
        request = NewStartStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewStartWebRecordRequest() (request *StartWebRecordRequest) {
    request = &StartWebRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartWebRecord")
    
    
    return
}

func NewStartWebRecordResponse() (response *StartWebRecordResponse) {
    response = &StartWebRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartWebRecord
// This interface can be used to initiate a web-page recording task. In the interface parameters, specify the recording URL, recording resolution, recording result storage and other parameters. If there are parameter or API logic problems, the result will be returned immediately. If there are page problems, such as the page cannot be accessed, the result will be returned in the callback. Please pay attention.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartWebRecord(request *StartWebRecordRequest) (response *StartWebRecordResponse, err error) {
    return c.StartWebRecordWithContext(context.Background(), request)
}

// StartWebRecord
// This interface can be used to initiate a web-page recording task. In the interface parameters, specify the recording URL, recording resolution, recording result storage and other parameters. If there are parameter or API logic problems, the result will be returned immediately. If there are page problems, such as the page cannot be accessed, the result will be returned in the callback. Please pay attention.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartWebRecordWithContext(ctx context.Context, request *StartWebRecordRequest) (response *StartWebRecordResponse, err error) {
    if request == nil {
        request = NewStartWebRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartWebRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartWebRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopAIConversationRequest() (request *StopAIConversationRequest) {
    request = &StopAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopAIConversation")
    
    
    return
}

func NewStopAIConversationResponse() (response *StopAIConversationResponse) {
    response = &StopAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAIConversation
// Stop AI conversation task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAIConversation(request *StopAIConversationRequest) (response *StopAIConversationResponse, err error) {
    return c.StopAIConversationWithContext(context.Background(), request)
}

// StopAIConversation
// Stop AI conversation task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAIConversationWithContext(ctx context.Context, request *StopAIConversationRequest) (response *StopAIConversationResponse, err error) {
    if request == nil {
        request = NewStopAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewStopAITranscriptionRequest() (request *StopAITranscriptionRequest) {
    request = &StopAITranscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopAITranscription")
    
    
    return
}

func NewStopAITranscriptionResponse() (response *StopAITranscriptionResponse) {
    response = &StopAITranscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAITranscription
// Stop AI Transcription task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAITranscription(request *StopAITranscriptionRequest) (response *StopAITranscriptionResponse, err error) {
    return c.StopAITranscriptionWithContext(context.Background(), request)
}

// StopAITranscription
// Stop AI Transcription task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAITranscriptionWithContext(ctx context.Context, request *StopAITranscriptionRequest) (response *StopAITranscriptionResponse, err error) {
    if request == nil {
        request = NewStopAITranscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopAITranscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAITranscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAITranscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewStopPublishCdnStreamRequest() (request *StopPublishCdnStreamRequest) {
    request = &StopPublishCdnStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopPublishCdnStream")
    
    
    return
}

func NewStopPublishCdnStreamResponse() (response *StopPublishCdnStreamResponse) {
    response = &StopPublishCdnStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopPublishCdnStream
// This API is used to stop a relaying task.
//
// You can create a relay task before the anchor enters the room. When the relay task is finished, you need to call the stop interface actively. If you do not call the Stop Relay Task Interface, Tencent Cloud will automatically stop the mix relay task when all users participating in the mix have no data uploaded for a period of time exceeding the timeout (AgentParams.MaxIdleTime) set when starting the relay task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopPublishCdnStream(request *StopPublishCdnStreamRequest) (response *StopPublishCdnStreamResponse, err error) {
    return c.StopPublishCdnStreamWithContext(context.Background(), request)
}

// StopPublishCdnStream
// This API is used to stop a relaying task.
//
// You can create a relay task before the anchor enters the room. When the relay task is finished, you need to call the stop interface actively. If you do not call the Stop Relay Task Interface, Tencent Cloud will automatically stop the mix relay task when all users participating in the mix have no data uploaded for a period of time exceeding the timeout (AgentParams.MaxIdleTime) set when starting the relay task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopPublishCdnStreamWithContext(ctx context.Context, request *StopPublishCdnStreamRequest) (response *StopPublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewStopPublishCdnStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopPublishCdnStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopPublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopPublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamIngestRequest() (request *StopStreamIngestRequest) {
    request = &StopStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopStreamIngest")
    
    
    return
}

func NewStopStreamIngestResponse() (response *StopStreamIngestResponse) {
    response = &StopStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopStreamIngest
// Stop a Pull stream Relay task.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopStreamIngest(request *StopStreamIngestRequest) (response *StopStreamIngestResponse, err error) {
    return c.StopStreamIngestWithContext(context.Background(), request)
}

// StopStreamIngest
// Stop a Pull stream Relay task.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopStreamIngestWithContext(ctx context.Context, request *StopStreamIngestRequest) (response *StopStreamIngestResponse, err error) {
    if request == nil {
        request = NewStopStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewStopWebRecordRequest() (request *StopWebRecordRequest) {
    request = &StopWebRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopWebRecord")
    
    
    return
}

func NewStopWebRecordResponse() (response *StopWebRecordResponse) {
    response = &StopWebRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopWebRecord
// Stop an web-page recording task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopWebRecord(request *StopWebRecordRequest) (response *StopWebRecordResponse, err error) {
    return c.StopWebRecordWithContext(context.Background(), request)
}

// StopWebRecord
// Stop an web-page recording task
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopWebRecordWithContext(ctx context.Context, request *StopWebRecordRequest) (response *StopWebRecordResponse, err error) {
    if request == nil {
        request = NewStopWebRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopWebRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWebRecordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAIConversationRequest() (request *UpdateAIConversationRequest) {
    request = &UpdateAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdateAIConversation")
    
    
    return
}

func NewUpdateAIConversationResponse() (response *UpdateAIConversationResponse) {
    response = &UpdateAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAIConversation
// Update AI conversation task parameters
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
func (c *Client) UpdateAIConversation(request *UpdateAIConversationRequest) (response *UpdateAIConversationResponse, err error) {
    return c.UpdateAIConversationWithContext(context.Background(), request)
}

// UpdateAIConversation
// Update AI conversation task parameters
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
func (c *Client) UpdateAIConversationWithContext(ctx context.Context, request *UpdateAIConversationRequest) (response *UpdateAIConversationResponse, err error) {
    if request == nil {
        request = NewUpdateAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdateAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePublishCdnStreamRequest() (request *UpdatePublishCdnStreamRequest) {
    request = &UpdatePublishCdnStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdatePublishCdnStream")
    
    
    return
}

func NewUpdatePublishCdnStreamResponse() (response *UpdatePublishCdnStreamResponse) {
    response = &UpdatePublishCdnStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePublishCdnStream
// This API is used to change the parameters of a relaying task.
//
// Note: For details about how to use this API, see the `StartPublishCdnStream` document.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePublishCdnStream(request *UpdatePublishCdnStreamRequest) (response *UpdatePublishCdnStreamResponse, err error) {
    return c.UpdatePublishCdnStreamWithContext(context.Background(), request)
}

// UpdatePublishCdnStream
// This API is used to change the parameters of a relaying task.
//
// Note: For details about how to use this API, see the `StartPublishCdnStream` document.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePublishCdnStreamWithContext(ctx context.Context, request *UpdatePublishCdnStreamRequest) (response *UpdatePublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewUpdatePublishCdnStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdatePublishCdnStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateStreamIngestRequest() (request *UpdateStreamIngestRequest) {
    request = &UpdateStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdateStreamIngest")
    
    
    return
}

func NewUpdateStreamIngestResponse() (response *UpdateStreamIngestResponse) {
    response = &UpdateStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateStreamIngest
// You can update the StreamUrl of the Relay task.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) UpdateStreamIngest(request *UpdateStreamIngestRequest) (response *UpdateStreamIngestResponse, err error) {
    return c.UpdateStreamIngestWithContext(context.Background(), request)
}

// UpdateStreamIngest
// You can update the StreamUrl of the Relay task.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) UpdateStreamIngestWithContext(ctx context.Context, request *UpdateStreamIngestRequest) (response *UpdateStreamIngestResponse, err error) {
    if request == nil {
        request = NewUpdateStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdateStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateStreamIngestResponse()
    err = c.Send(request, response)
    return
}
