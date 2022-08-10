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

package v20190919

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-19"

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


func NewCreateTranscodeRequest() (request *CreateTranscodeRequest) {
    request = &CreateTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "CreateTranscode")
    
    
    return
}

func NewCreateTranscodeResponse() (response *CreateTranscodeResponse) {
    response = &CreateTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTranscode
// This API is used to create a document transcoding task.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateTranscode(request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    return c.CreateTranscodeWithContext(context.Background(), request)
}

// CreateTranscode
// This API is used to create a document transcoding task.
//
// error code that may be returned:
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateTranscodeWithContext(ctx context.Context, request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordRequest() (request *DescribeOnlineRecordRequest) {
    request = &DescribeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecord")
    
    
    return
}

func NewDescribeOnlineRecordResponse() (response *DescribeOnlineRecordResponse) {
    response = &DescribeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOnlineRecord
// This API is used to query the status and result of a real-time recording task.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecord(request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    return c.DescribeOnlineRecordWithContext(context.Background(), request)
}

// DescribeOnlineRecord
// This API is used to query the status and result of a real-time recording task.
//
// error code that may be returned:
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordWithContext(ctx context.Context, request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordCallbackRequest() (request *DescribeOnlineRecordCallbackRequest) {
    request = &DescribeOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecordCallback")
    
    
    return
}

func NewDescribeOnlineRecordCallbackResponse() (response *DescribeOnlineRecordCallbackResponse) {
    response = &DescribeOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOnlineRecordCallback
// This API is used to query the real-time recording callback address.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordCallback(request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    return c.DescribeOnlineRecordCallbackWithContext(context.Background(), request)
}

// DescribeOnlineRecordCallback
// This API is used to query the real-time recording callback address.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordCallbackWithContext(ctx context.Context, request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeRequest() (request *DescribeTranscodeRequest) {
    request = &DescribeTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscode")
    
    
    return
}

func NewDescribeTranscodeResponse() (response *DescribeTranscodeResponse) {
    response = &DescribeTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTranscode
// This API is used to query the progress and result of a document transcoding task.
//
// error code that may be returned:
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscode(request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    return c.DescribeTranscodeWithContext(context.Background(), request)
}

// DescribeTranscode
// This API is used to query the progress and result of a document transcoding task.
//
// error code that may be returned:
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeWithContext(ctx context.Context, request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeCallbackRequest() (request *DescribeTranscodeCallbackRequest) {
    request = &DescribeTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeCallback")
    
    
    return
}

func NewDescribeTranscodeCallbackResponse() (response *DescribeTranscodeCallbackResponse) {
    response = &DescribeTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTranscodeCallback
// This API is used to query the document transcoding callback address.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeCallback(request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    return c.DescribeTranscodeCallbackWithContext(context.Background(), request)
}

// DescribeTranscodeCallback
// This API is used to query the document transcoding callback address.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeCallbackWithContext(ctx context.Context, request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewPauseOnlineRecordRequest() (request *PauseOnlineRecordRequest) {
    request = &PauseOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "PauseOnlineRecord")
    
    
    return
}

func NewPauseOnlineRecordResponse() (response *PauseOnlineRecordResponse) {
    response = &PauseOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PauseOnlineRecord
// This API is used to pause real-time recording.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) PauseOnlineRecord(request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    return c.PauseOnlineRecordWithContext(context.Background(), request)
}

// PauseOnlineRecord
// This API is used to pause real-time recording.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) PauseOnlineRecordWithContext(ctx context.Context, request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    if request == nil {
        request = NewPauseOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewResumeOnlineRecordRequest() (request *ResumeOnlineRecordRequest) {
    request = &ResumeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "ResumeOnlineRecord")
    
    
    return
}

func NewResumeOnlineRecordResponse() (response *ResumeOnlineRecordResponse) {
    response = &ResumeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeOnlineRecord
// This API is used to resume real-time recording.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) ResumeOnlineRecord(request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    return c.ResumeOnlineRecordWithContext(context.Background(), request)
}

// ResumeOnlineRecord
// This API is used to resume real-time recording.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) ResumeOnlineRecordWithContext(ctx context.Context, request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewResumeOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackRequest() (request *SetOnlineRecordCallbackRequest) {
    request = &SetOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallback")
    
    
    return
}

func NewSetOnlineRecordCallbackResponse() (response *SetOnlineRecordCallbackResponse) {
    response = &SetOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetOnlineRecordCallback
// This API is used to set the real-time recording callback address. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40258?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallback(request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    return c.SetOnlineRecordCallbackWithContext(context.Background(), request)
}

// SetOnlineRecordCallback
// This API is used to set the real-time recording callback address. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40258?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackWithContext(ctx context.Context, request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOnlineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackKeyRequest() (request *SetOnlineRecordCallbackKeyRequest) {
    request = &SetOnlineRecordCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallbackKey")
    
    
    return
}

func NewSetOnlineRecordCallbackKeyResponse() (response *SetOnlineRecordCallbackKeyResponse) {
    response = &SetOnlineRecordCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetOnlineRecordCallbackKey
// This API is used to set the callback authentication key for real-time recording. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackKey(request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    return c.SetOnlineRecordCallbackKeyWithContext(context.Background(), request)
}

// SetOnlineRecordCallbackKey
// This API is used to set the callback authentication key for real-time recording. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackKeyWithContext(ctx context.Context, request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOnlineRecordCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOnlineRecordCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackRequest() (request *SetTranscodeCallbackRequest) {
    request = &SetTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallback")
    
    
    return
}

func NewSetTranscodeCallbackResponse() (response *SetTranscodeCallbackResponse) {
    response = &SetTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetTranscodeCallback
// This API is used to set the callback address for document transcoding. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40260?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallback(request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    return c.SetTranscodeCallbackWithContext(context.Background(), request)
}

// SetTranscodeCallback
// This API is used to set the callback address for document transcoding. For the callback format, please [see here](https://intl.cloud.tencent.com/document/product/1137/40260?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackWithContext(ctx context.Context, request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTranscodeCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackKeyRequest() (request *SetTranscodeCallbackKeyRequest) {
    request = &SetTranscodeCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallbackKey")
    
    
    return
}

func NewSetTranscodeCallbackKeyResponse() (response *SetTranscodeCallbackKeyResponse) {
    response = &SetTranscodeCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetTranscodeCallbackKey
// This API is used to set the callback authentication key for document transcoding. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackKey(request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    return c.SetTranscodeCallbackKeyWithContext(context.Background(), request)
}

// SetTranscodeCallbackKey
// This API is used to set the callback authentication key for document transcoding. For more information, see [Event Notification](https://intl.cloud.tencent.com/document/product/1137/40257?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackKeyWithContext(ctx context.Context, request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTranscodeCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTranscodeCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewStartOnlineRecordRequest() (request *StartOnlineRecordRequest) {
    request = &StartOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "StartOnlineRecord")
    
    
    return
}

func NewStartOnlineRecordResponse() (response *StartOnlineRecordResponse) {
    response = &StartOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartOnlineRecord
// This API is used to start a real-time recording task.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartOnlineRecord(request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    return c.StartOnlineRecordWithContext(context.Background(), request)
}

// StartOnlineRecord
// This API is used to start a real-time recording task.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartOnlineRecordWithContext(ctx context.Context, request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStartOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopOnlineRecordRequest() (request *StopOnlineRecordRequest) {
    request = &StopOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiw", APIVersion, "StopOnlineRecord")
    
    
    return
}

func NewStopOnlineRecordResponse() (response *StopOnlineRecordResponse) {
    response = &StopOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopOnlineRecord
// This API is used to stop real-time recording.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopOnlineRecord(request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    return c.StopOnlineRecordWithContext(context.Background(), request)
}

// StopOnlineRecord
// This API is used to stop real-time recording.
//
// error code that may be returned:
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopOnlineRecordWithContext(ctx context.Context, request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStopOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopOnlineRecordResponse()
    err = c.Send(request, response)
    return
}
