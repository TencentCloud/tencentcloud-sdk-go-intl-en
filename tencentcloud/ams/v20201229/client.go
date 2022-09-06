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

package v20201229

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-12-29"

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


func NewCancelTaskRequest() (request *CancelTaskRequest) {
    request = &CancelTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ams", APIVersion, "CancelTask")
    
    
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelTask
// This API is used to cancel a moderation task. If it returns `RequestId`, the task has been canceled successfully.<br>Default API request rate limit: **20 requests/sec**.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    return c.CancelTaskWithContext(context.Background(), request)
}

// CancelTask
// This API is used to cancel a moderation task. If it returns `RequestId`, the task has been canceled successfully.<br>Default API request rate limit: **20 requests/sec**.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAudioModerationTaskRequest() (request *CreateAudioModerationTaskRequest) {
    request = &CreateAudioModerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ams", APIVersion, "CreateAudioModerationTask")
    
    
    return
}

func NewCreateAudioModerationTaskResponse() (response *CreateAudioModerationTaskResponse) {
    response = &CreateAudioModerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAudioModerationTask
// This API is used to submit audio content (such as an audio file or stream URL) for smart moderation. Before using it, you need to log in to the console with the Tencent Cloud root account [to activate AMS](https://console.cloud.tencent.com/cms/audio/package) and adjust the business configuration.<br>
//
// 
//
// ### Feature use instructions
//
// - Go to the "[CMS console - AMS](https://console.cloud.tencent.com/cms/audio/package)" to activate AMS.
//
// - Default API request rate limit: **20 requests/sec**. When this limit is exceeded, requests for async moderation tasks (audio on demand) will automatically join the queue of requests pending moderation, while an error will be reported for sync moderation tasks (audio live streaming).
//
// 
//
// ### API feature description
//
// - It can detect audio streams or files for non-compliant content;
//
// - You can set the callback address (Callback) to get the detection result (for moderation tasks in progress, the maximum callback time is the configured **segment length + 2s). You can also call the API for querying the audio detection result to poll the detection result;
//
// - It can recognize various types of non-compliant content, including vulgarity, abuse, porn, and advertising;
//
// - You can submit **up to 10** detection tasks at a time.
//
// 
//
// ### Call description for audio file
//
// - Supported audio file size: **< 500 MB**;
//
// - Supported audio file duration: **< 1 hour**;
//
// - Supported audio bitrate: 128–256 Kbps;
//
// - Supported audio file formats: WAV, MP3, AAC, FLAC, AMR, 3GP, M4A, WMA, OGG, and APE;
//
// - **When the input is a video file**, the audio track can be extracted from it for audio content moderation.
//
// 
//
// ### Call description for audio stream
//
// - Supported audio stream duration: **< 3 hours**;
//
// - Supported audio bitrate: 128–256 Kbps;
//
// - Supported audio stream transfer protocols: RTMP, HTTP, and HTTPS;
//
// - Supported audio stream formats: RTP, SRTP, RTMP, RTMPS, MMSH, MMST, HLS, HTTP, TCP, HTTPS, and M3U8;
//
// - **When the input is a video stream**, the audio track can be extracted from it for audio content moderation.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAudioModerationTask(request *CreateAudioModerationTaskRequest) (response *CreateAudioModerationTaskResponse, err error) {
    return c.CreateAudioModerationTaskWithContext(context.Background(), request)
}

// CreateAudioModerationTask
// This API is used to submit audio content (such as an audio file or stream URL) for smart moderation. Before using it, you need to log in to the console with the Tencent Cloud root account [to activate AMS](https://console.cloud.tencent.com/cms/audio/package) and adjust the business configuration.<br>
//
// 
//
// ### Feature use instructions
//
// - Go to the "[CMS console - AMS](https://console.cloud.tencent.com/cms/audio/package)" to activate AMS.
//
// - Default API request rate limit: **20 requests/sec**. When this limit is exceeded, requests for async moderation tasks (audio on demand) will automatically join the queue of requests pending moderation, while an error will be reported for sync moderation tasks (audio live streaming).
//
// 
//
// ### API feature description
//
// - It can detect audio streams or files for non-compliant content;
//
// - You can set the callback address (Callback) to get the detection result (for moderation tasks in progress, the maximum callback time is the configured **segment length + 2s). You can also call the API for querying the audio detection result to poll the detection result;
//
// - It can recognize various types of non-compliant content, including vulgarity, abuse, porn, and advertising;
//
// - You can submit **up to 10** detection tasks at a time.
//
// 
//
// ### Call description for audio file
//
// - Supported audio file size: **< 500 MB**;
//
// - Supported audio file duration: **< 1 hour**;
//
// - Supported audio bitrate: 128–256 Kbps;
//
// - Supported audio file formats: WAV, MP3, AAC, FLAC, AMR, 3GP, M4A, WMA, OGG, and APE;
//
// - **When the input is a video file**, the audio track can be extracted from it for audio content moderation.
//
// 
//
// ### Call description for audio stream
//
// - Supported audio stream duration: **< 3 hours**;
//
// - Supported audio bitrate: 128–256 Kbps;
//
// - Supported audio stream transfer protocols: RTMP, HTTP, and HTTPS;
//
// - Supported audio stream formats: RTP, SRTP, RTMP, RTMPS, MMSH, MMST, HLS, HTTP, TCP, HTTPS, and M3U8;
//
// - **When the input is a video stream**, the audio track can be extracted from it for audio content moderation.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAudioModerationTaskWithContext(ctx context.Context, request *CreateAudioModerationTaskRequest) (response *CreateAudioModerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateAudioModerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAudioModerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAudioModerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ams", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// This API is used to view the details of an audio moderation task, including task status, detection result, recognized text content of the audio file, maliciousness tag that corresponds to the detection result, and suggested operation. For the specific output content, see the sample output parameters.<br>Default API request rate limit: **100 requests/sec**.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// This API is used to view the details of an audio moderation task, including task status, detection result, recognized text content of the audio file, maliciousness tag that corresponds to the detection result, and suggested operation. For the specific output content, see the sample output parameters.<br>Default API request rate limit: **100 requests/sec**.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ams", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// This API is used to view the list of moderation tasks. You can also filter moderation tasks by multiple types of business information, such as business type, moderation result, and task status. The output content of the task list includes the total number of queried tasks, task name, task status, audio moderation type, maliciousness tag of the detection result, and suggested operation. For the specific output content, see the sample output parameters.<br>Default API request rate limit: **20 requests/sec**.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// This API is used to view the list of moderation tasks. You can also filter moderation tasks by multiple types of business information, such as business type, moderation result, and task status. The output content of the task list includes the total number of queried tasks, task name, task status, audio moderation type, maliciousness tag of the detection result, and suggested operation. For the specific output content, see the sample output parameters.<br>Default API request rate limit: **20 requests/sec**.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}
