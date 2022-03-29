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
    request.Init().WithApiInfo("vm", APIVersion, "CancelTask")
    
    
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelTask
// This API is used to cancel a moderation task. It will return the `TaskId` of the task after the task is canceled successfully.<br>
//
// 
//
// Default API request rate limit: **20 requests/sec**.
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
// This API is used to cancel a moderation task. It will return the `TaskId` of the task after the task is canceled successfully.<br>
//
// 
//
// Default API request rate limit: **20 requests/sec**.
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

func NewCreateVideoModerationTaskRequest() (request *CreateVideoModerationTaskRequest) {
    request = &CreateVideoModerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vm", APIVersion, "CreateVideoModerationTask")
    
    
    return
}

func NewCreateVideoModerationTaskResponse() (response *CreateVideoModerationTaskResponse) {
    response = &CreateVideoModerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVideoModerationTask
// This API is used to submit a video file or stream for smart moderation. Before using it, you need to log in to the console with the Tencent Cloud root account [to activate VM](https://console.cloud.tencent.com/cms/video/package) and adjust the business configuration.<br>
//
// ### Feature use instructions
//
// 
//
// - Go to the "[CMS console - VM](https://console.cloud.tencent.com/cms/video/package)" to activate AMS.
//
// 
//
// - This API is a paid API. For its billing mode, see [VM Pricing](https://intl.cloud.tencent.com/product/vm/pricing?from_cn_redirect=1).
//
// 
//
// - Default API request rate limit: **20 requests/sec**. When this limit is exceeded, requests for async moderation tasks (video on demand) will automatically join the queue of requests pending moderation, while an error will be reported for sync moderation tasks (video live streaming).
//
// - Default limit on the number of concurrent moderation channels: 10. When this limit is exceeded, requests for async moderation tasks (video on demand) will automatically join the queue of requests pending moderation, while an error will be reported for sync moderation tasks (video live streaming).
//
// 
//
// 
//
// ### API feature description
//
// 
//
// - It can automatically detect video files or streams and recognize non-compliant content in them based on the deep learning technology from the perspectives of OCR-based text recognition, object detection (such as object, advertising logo, and QR code), image recognition, and audio moderation;
//
// - It allows you to set the callback address (Callback) to get the detection result or call the API for viewing task details to get the details of the detection result through polling. For normal video moderation tasks, if non-compliant content is contained, the captured frames will be called back within **3s**, and the audio segments will be called back within the configured **segment duration + 2s**; for queued moderation tasks, the callback time will be equal to the sum of the callback time for normal moderation and waiting time;
//
// - The API for viewing the moderation task list can be called to query the task queue. You can filter moderation tasks by multiple types of business information, such as business type, moderation result, and task status;
//
// - It can recognize various non-compliant scenarios, including vulgarity, abuse, pornography, and advertising;
//
// - It allows you to customize moderation policies based on different business scenarios;
//
// - You can customize blocklist/allowlist dictionaries and image libraries to filter non-compliant content of custom types (currently, only blocklist configuration is supported);
//
// - You can customize the moderation task priority, so that when multiple tasks are queuing, the task priority will be automatically adjusted according to the configuration;
//
// - You can submit detection tasks in batches and **create up to 10 tasks at a time**;
//
// 
//
// ### Call description for video file
//
// 
//
// - Supported video file size: **< 3 GB**
//
// - Supported video file resolution: **the optimal resolution is 1920x1080 (1080p)**. For video files of less than 300 MB in size, their resolution can be greater than 1080p; for video files of a greater size, you can call [MPS](https://intl.cloud.tencent.com/product/mps/details?from_cn_redirect=1) to transcode them before submitting them for moderation;
//
// - Supported video file formats: FLV, MKV, MP4, RMVB, AVI, WMV, 3GP, TS, MOV, RM, MPEG, and WMF;
//
// - Supported video file access methods: URL (over HTTP/HTTPS) and Tencent Cloud COS;
//
// - If you pass in the access URL of a video file, you need to **limit its header file read time to 3 seconds**. To ensure the stability and reliability of the video to be detected, we recommend you use Tencent Cloud COS for storage or CDN for caching;
//
// - You can configure whether to enable audio moderation, and if it is not enabled, only the image content of video files will be moderated.
//
// 
//
// ### Call description for video stream
//
// 
//
// - Supported video stream duration: **< 5 hours**
//
// - Supported video stream resolution: **1920x1080 (1080p)**. For videos with a higher resolution, you can call [live transcoding](https://intl.cloud.tencent.com/document/product/267/39889?from_cn_redirect=1) to transcode them before submitting them for moderation;
//
// - Supported video stream formats: mainstream video stream codecs such as RMTP and FLV;
//
// - Supported video transfer protocols: HTTP, HTTPS, and RTMP;
//
// - You can configure whether to enable audio moderation, and if it is not enabled, only the image content of video streams will be moderated.
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
func (c *Client) CreateVideoModerationTask(request *CreateVideoModerationTaskRequest) (response *CreateVideoModerationTaskResponse, err error) {
    return c.CreateVideoModerationTaskWithContext(context.Background(), request)
}

// CreateVideoModerationTask
// This API is used to submit a video file or stream for smart moderation. Before using it, you need to log in to the console with the Tencent Cloud root account [to activate VM](https://console.cloud.tencent.com/cms/video/package) and adjust the business configuration.<br>
//
// ### Feature use instructions
//
// 
//
// - Go to the "[CMS console - VM](https://console.cloud.tencent.com/cms/video/package)" to activate AMS.
//
// 
//
// - This API is a paid API. For its billing mode, see [VM Pricing](https://intl.cloud.tencent.com/product/vm/pricing?from_cn_redirect=1).
//
// 
//
// - Default API request rate limit: **20 requests/sec**. When this limit is exceeded, requests for async moderation tasks (video on demand) will automatically join the queue of requests pending moderation, while an error will be reported for sync moderation tasks (video live streaming).
//
// - Default limit on the number of concurrent moderation channels: 10. When this limit is exceeded, requests for async moderation tasks (video on demand) will automatically join the queue of requests pending moderation, while an error will be reported for sync moderation tasks (video live streaming).
//
// 
//
// 
//
// ### API feature description
//
// 
//
// - It can automatically detect video files or streams and recognize non-compliant content in them based on the deep learning technology from the perspectives of OCR-based text recognition, object detection (such as object, advertising logo, and QR code), image recognition, and audio moderation;
//
// - It allows you to set the callback address (Callback) to get the detection result or call the API for viewing task details to get the details of the detection result through polling. For normal video moderation tasks, if non-compliant content is contained, the captured frames will be called back within **3s**, and the audio segments will be called back within the configured **segment duration + 2s**; for queued moderation tasks, the callback time will be equal to the sum of the callback time for normal moderation and waiting time;
//
// - The API for viewing the moderation task list can be called to query the task queue. You can filter moderation tasks by multiple types of business information, such as business type, moderation result, and task status;
//
// - It can recognize various non-compliant scenarios, including vulgarity, abuse, pornography, and advertising;
//
// - It allows you to customize moderation policies based on different business scenarios;
//
// - You can customize blocklist/allowlist dictionaries and image libraries to filter non-compliant content of custom types (currently, only blocklist configuration is supported);
//
// - You can customize the moderation task priority, so that when multiple tasks are queuing, the task priority will be automatically adjusted according to the configuration;
//
// - You can submit detection tasks in batches and **create up to 10 tasks at a time**;
//
// 
//
// ### Call description for video file
//
// 
//
// - Supported video file size: **< 3 GB**
//
// - Supported video file resolution: **the optimal resolution is 1920x1080 (1080p)**. For video files of less than 300 MB in size, their resolution can be greater than 1080p; for video files of a greater size, you can call [MPS](https://intl.cloud.tencent.com/product/mps/details?from_cn_redirect=1) to transcode them before submitting them for moderation;
//
// - Supported video file formats: FLV, MKV, MP4, RMVB, AVI, WMV, 3GP, TS, MOV, RM, MPEG, and WMF;
//
// - Supported video file access methods: URL (over HTTP/HTTPS) and Tencent Cloud COS;
//
// - If you pass in the access URL of a video file, you need to **limit its header file read time to 3 seconds**. To ensure the stability and reliability of the video to be detected, we recommend you use Tencent Cloud COS for storage or CDN for caching;
//
// - You can configure whether to enable audio moderation, and if it is not enabled, only the image content of video files will be moderated.
//
// 
//
// ### Call description for video stream
//
// 
//
// - Supported video stream duration: **< 5 hours**
//
// - Supported video stream resolution: **1920x1080 (1080p)**. For videos with a higher resolution, you can call [live transcoding](https://intl.cloud.tencent.com/document/product/267/39889?from_cn_redirect=1) to transcode them before submitting them for moderation;
//
// - Supported video stream formats: mainstream video stream codecs such as RMTP and FLV;
//
// - Supported video transfer protocols: HTTP, HTTPS, and RTMP;
//
// - You can configure whether to enable audio moderation, and if it is not enabled, only the image content of video streams will be moderated.
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
func (c *Client) CreateVideoModerationTaskWithContext(ctx context.Context, request *CreateVideoModerationTaskRequest) (response *CreateVideoModerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoModerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoModerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoModerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vm", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// This API is used to poll the details of the detection result.<br>
//
// 
//
// Default API request rate limit: **200 requests/sec**.
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
// This API is used to poll the details of the detection result.<br>
//
// 
//
// Default API request rate limit: **200 requests/sec**.
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
    request.Init().WithApiInfo("vm", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// This API is used to query the task queue. You can filter moderation tasks by multiple types of business information, such as business type, moderation result, and task status.<br>
//
// 
//
// Default API request rate limit: **20 requests/sec**.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
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
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// This API is used to query the task queue. You can filter moderation tasks by multiple types of business information, such as business type, moderation result, and task status.<br>
//
// 
//
// Default API request rate limit: **20 requests/sec**.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
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
