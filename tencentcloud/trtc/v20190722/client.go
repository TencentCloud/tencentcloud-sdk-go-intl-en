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
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallDetailInfoResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRecordingResponse()
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
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
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
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScaleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScaleInfoResponse()
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
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
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
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
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
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
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
// 2. If you need to call this API, please upgrade the monitoring dashboard version to "Standard". For more details, please refer to: https://www.tencentcloud.com/document/product/647/54481.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
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
func (c *Client) SetUserBlockedByStrRoomIdWithContext(ctx context.Context, request *SetUserBlockedByStrRoomIdRequest) (response *SetUserBlockedByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewSetUserBlockedByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetUserBlockedByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetUserBlockedByStrRoomIdResponse()
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
// This API is used to mix streams and relay the mixed stream to CDNs. You can use this API to do the following:
//
// 1. Publish (also known as "relay") the audio/video stream of one anchor to CDNs. For details, see example 2 (starting a task to relay the audio and video of a stream) and example 3 (starting a task to relay only the audio of a stream).
//
// 2. Mix the streams of multiple anchors in a room or in different rooms and publish the mixed stream to CDNs. You can use `AudioParams.SubscribeAudioList` to specify the users whose audios are mixed, and use `VideoParams.LayoutParams` to specify the layout of the anchors’ videos. For details, see example 1 (mixing streams and publishing the mixed stream to a CDN).
//
// 3. Mix multiple streams in a room according to a template and publish the mixed stream to CDNs. The TRTC backend will detect the change of anchors in the room and adjust the video layout automatically according to the stream mixing template. The following template types are supported:
//
//      - Floating: The entire screen is covered by the video of the first user who enters the room, and the videos of other users are displayed as small windows in rows in the bottom-left corner in room entry sequence. The screen allows up to 4 rows of 4 small videos, which float over the large video. Up to 1 large and 15 small videos can be displayed.
//
//      - Grid: The videos of all users split the screen evenly. The more the users, the smaller the video dimensions. Up to 16 videos are supported.
//
//      - Screen sharing: This is designed for video conferencing and online education. The shared screen (or camera video of the anchor) is always displayed as the large video, which occupies the left half of the screen, and the videos of other users occupy the right half in up to two columns of up to eight small videos each. Up to 1 large video and 15 small videos can be displayed. If the upstream aspect ratio does not match the output, the large video on the left will be scaled and displayed in whole, while the small videos on the right will be cropped.
//
// 4. Publish audio/video streams to up to 10 CDNs at a time. You can use `PublishCdnParams.PublishCdnUrl` to specify the URLs of the CDNs to publish to. To publish to Tencent Cloud’s CDN, set `PublishCdnParams.IsTencentCdn` to 1.
//
// 5. Configure a server-side callback to have Tencent Cloud send relay status updates to your server in the form of HTTP/HTTPS POST requests. To use the callback for relay events, please contact Technical Support.
//
// 6. The API supports four regions: Guangzhou, Shanghai, Beijing, and Hong Kong. If you relay to regions outside the Chinese mainland, select Hong Kong.
//
//  
//
// Notes:
//
// 1. **Because On-Cloud MixTranscoding is a paid feature, calling this API will incur MixTranscoding fees. For details, see [Billing of MixTranscoding and Relay to CDN](https://intl.cloud.tencent.com/document/product/647/49446?from_cn_redirect=1).**
//
// 2. **Relaying to third-party CDNs will incur relaying fees. For details, see [Billing of MixTranscoding and Relay to CDN](https://intl.cloud.tencent.com/document/product/647/82155?from_cn_redirect=1).**
//
// 
//
// Others:
//
// 1. You need to first call `StartPublishCdnStream` to start a relay task and get the task ID before you can use the `UpdatePublishCdnStream` API to modify the task and `StopPublishCdnStream` to stop the task.
//
// 2. To ensure the stability of relaying, you cannot switch between relaying audio only, relaying audio and video, and relaying video only for the same task.
//
// 3. To ensure the stability of relaying, you cannot change the video codec, audio codec, audio sample rate, audio bitrate, or sound channels using the `UpdatePublishCdnStream` API. We recommend you pass in all the other parameters when calling `UpdatePublishCdnStream`. If you only want to enable/disable transcoding, make sure you pass in all the other parameters.
//
// 4. When you relay a single stream, specify both `AudioParams` and `VideoParams` to publish both audio and video, and specify only `AudioParams` to publish audio only. You cannot switch between the two modes during the relaying process. For `VideoParams`, set `Width`, `Height`, `Fps`, `Bitrate`, and `Gop` according to the actual settings used for publishing.
//
// 5. The `SequenceNumber` parameter is required when you call `UpdatePublishCdnStream` to change the relaying parameters. It ensures that multiple requests for the same relaying task are in the correct order. The value of `SequenceNumber` increases each time a new request is made for the same task. If `InternalError` is returned, try again using the same `SequenceNumber`. You don’t need to handle the `FailedOperation.OutdateRequest` error.
//
// 6. You can create a relay task before anchors enter a room, in which case you need to manually call `StopPublishCdnStream` to stop the task. If you don’t, after all the users whose streams are mixed leave the room, the TRTC backend will wait for the timeout period (`AgentParams.MaxIdleTime`) to elapse before stopping the relaying task.
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
// This API is used to mix streams and relay the mixed stream to CDNs. You can use this API to do the following:
//
// 1. Publish (also known as "relay") the audio/video stream of one anchor to CDNs. For details, see example 2 (starting a task to relay the audio and video of a stream) and example 3 (starting a task to relay only the audio of a stream).
//
// 2. Mix the streams of multiple anchors in a room or in different rooms and publish the mixed stream to CDNs. You can use `AudioParams.SubscribeAudioList` to specify the users whose audios are mixed, and use `VideoParams.LayoutParams` to specify the layout of the anchors’ videos. For details, see example 1 (mixing streams and publishing the mixed stream to a CDN).
//
// 3. Mix multiple streams in a room according to a template and publish the mixed stream to CDNs. The TRTC backend will detect the change of anchors in the room and adjust the video layout automatically according to the stream mixing template. The following template types are supported:
//
//      - Floating: The entire screen is covered by the video of the first user who enters the room, and the videos of other users are displayed as small windows in rows in the bottom-left corner in room entry sequence. The screen allows up to 4 rows of 4 small videos, which float over the large video. Up to 1 large and 15 small videos can be displayed.
//
//      - Grid: The videos of all users split the screen evenly. The more the users, the smaller the video dimensions. Up to 16 videos are supported.
//
//      - Screen sharing: This is designed for video conferencing and online education. The shared screen (or camera video of the anchor) is always displayed as the large video, which occupies the left half of the screen, and the videos of other users occupy the right half in up to two columns of up to eight small videos each. Up to 1 large video and 15 small videos can be displayed. If the upstream aspect ratio does not match the output, the large video on the left will be scaled and displayed in whole, while the small videos on the right will be cropped.
//
// 4. Publish audio/video streams to up to 10 CDNs at a time. You can use `PublishCdnParams.PublishCdnUrl` to specify the URLs of the CDNs to publish to. To publish to Tencent Cloud’s CDN, set `PublishCdnParams.IsTencentCdn` to 1.
//
// 5. Configure a server-side callback to have Tencent Cloud send relay status updates to your server in the form of HTTP/HTTPS POST requests. To use the callback for relay events, please contact Technical Support.
//
// 6. The API supports four regions: Guangzhou, Shanghai, Beijing, and Hong Kong. If you relay to regions outside the Chinese mainland, select Hong Kong.
//
//  
//
// Notes:
//
// 1. **Because On-Cloud MixTranscoding is a paid feature, calling this API will incur MixTranscoding fees. For details, see [Billing of MixTranscoding and Relay to CDN](https://intl.cloud.tencent.com/document/product/647/49446?from_cn_redirect=1).**
//
// 2. **Relaying to third-party CDNs will incur relaying fees. For details, see [Billing of MixTranscoding and Relay to CDN](https://intl.cloud.tencent.com/document/product/647/82155?from_cn_redirect=1).**
//
// 
//
// Others:
//
// 1. You need to first call `StartPublishCdnStream` to start a relay task and get the task ID before you can use the `UpdatePublishCdnStream` API to modify the task and `StopPublishCdnStream` to stop the task.
//
// 2. To ensure the stability of relaying, you cannot switch between relaying audio only, relaying audio and video, and relaying video only for the same task.
//
// 3. To ensure the stability of relaying, you cannot change the video codec, audio codec, audio sample rate, audio bitrate, or sound channels using the `UpdatePublishCdnStream` API. We recommend you pass in all the other parameters when calling `UpdatePublishCdnStream`. If you only want to enable/disable transcoding, make sure you pass in all the other parameters.
//
// 4. When you relay a single stream, specify both `AudioParams` and `VideoParams` to publish both audio and video, and specify only `AudioParams` to publish audio only. You cannot switch between the two modes during the relaying process. For `VideoParams`, set `Width`, `Height`, `Fps`, `Bitrate`, and `Gop` according to the actual settings used for publishing.
//
// 5. The `SequenceNumber` parameter is required when you call `UpdatePublishCdnStream` to change the relaying parameters. It ensures that multiple requests for the same relaying task are in the correct order. The value of `SequenceNumber` increases each time a new request is made for the same task. If `InternalError` is returned, try again using the same `SequenceNumber`. You don’t need to handle the `FailedOperation.OutdateRequest` error.
//
// 6. You can create a relay task before anchors enter a room, in which case you need to manually call `StopPublishCdnStream` to stop the task. If you don’t, after all the users whose streams are mixed leave the room, the TRTC backend will wait for the timeout period (`AgentParams.MaxIdleTime`) to elapse before stopping the relaying task.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishCdnStreamResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopPublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopPublishCdnStreamResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}
