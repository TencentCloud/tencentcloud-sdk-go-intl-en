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

package v20180801

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-08-01"

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


func NewAddDelayLiveStreamRequest() (request *AddDelayLiveStreamRequest) {
    request = &AddDelayLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "AddDelayLiveStream")
    
    return
}

func NewAddDelayLiveStreamResponse() (response *AddDelayLiveStreamResponse) {
    response = &AddDelayLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddDelayLiveStream
// This API is used to set a delay in playing the images of large live streaming events, in case of emergencies.
//
// 
//
// Note: if you are to set the delay before stream push, set it at least 5 minutes before the push. This API supports configuration only by stream.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddDelayLiveStream(request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewAddDelayLiveStreamRequest()
    }
    response = NewAddDelayLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewAddLiveDomainRequest() (request *AddLiveDomainRequest) {
    request = &AddLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "AddLiveDomain")
    
    return
}

func NewAddLiveDomainResponse() (response *AddLiveDomainResponse) {
    response = &AddLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddLiveDomain
// This API is used to add a domain name. Only one domain name can be submitted at a time, and it must have an ICP filing.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_HOSTOUTLIMIT = "FailedOperation.HostOutLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHINESECHARACTERDETECTED = "InternalError.ChineseCharacterDetected"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINALREADYEXIST = "InternalError.DomainAlreadyExist"
//  INTERNALERROR_DOMAINFORMATERROR = "InternalError.DomainFormatError"
//  INTERNALERROR_DOMAINGSLBFAIL = "InternalError.DomainGslbFail"
//  INTERNALERROR_DOMAINISFAMOUS = "InternalError.DomainIsFamous"
//  INTERNALERROR_DOMAINISLIMITED = "InternalError.DomainIsLimited"
//  INTERNALERROR_DOMAINNORECORD = "InternalError.DomainNoRecord"
//  INTERNALERROR_DOMAINTOOLONG = "InternalError.DomainTooLong"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"
//  INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
func (c *Client) AddLiveDomain(request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
    if request == nil {
        request = NewAddLiveDomainRequest()
    }
    response = NewAddLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewAddLiveWatermarkRequest() (request *AddLiveWatermarkRequest) {
    request = &AddLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "AddLiveWatermark")
    
    return
}

func NewAddLiveWatermarkResponse() (response *AddLiveWatermarkResponse) {
    response = &AddLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddLiveWatermark
// After a watermark is added and a watermark ID is successfully returned, you need to call the [CreateLiveWatermarkRule](https://intl.cloud.tencent.com/document/product/267/32629?from_cn_redirect=1) API to bind the watermark ID to a stream.
//
// After the number of watermarks exceeds the upper limit of 100, to add a new watermark, you must delete an old one first.
//
// error code that may be returned:
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_WATERMARKADDERROR = "InternalError.WatermarkAddError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"
func (c *Client) AddLiveWatermark(request *AddLiveWatermarkRequest) (response *AddLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewAddLiveWatermarkRequest()
    }
    response = NewAddLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewBindLiveDomainCertRequest() (request *BindLiveDomainCertRequest) {
    request = &BindLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "BindLiveDomainCert")
    
    return
}

func NewBindLiveDomainCertResponse() (response *BindLiveDomainCertResponse) {
    response = &BindLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindLiveDomainCert
// This API is used to bind a domain name certificate.
//
// Note: you need to call the `CreateLiveCert` API first to add a certificate. After getting the certificate ID, call this API for binding.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) BindLiveDomainCert(request *BindLiveDomainCertRequest) (response *BindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewBindLiveDomainCertRequest()
    }
    response = NewBindLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCommonMixStreamRequest() (request *CancelCommonMixStreamRequest) {
    request = &CancelCommonMixStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CancelCommonMixStream")
    
    return
}

func NewCancelCommonMixStreamResponse() (response *CancelCommonMixStreamResponse) {
    response = &CancelCommonMixStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelCommonMixStream
// This API is used to cancel a stream mix. It can be used basically in the same way as `mix_streamv2.cancel_mix_stream`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
func (c *Client) CancelCommonMixStream(request *CancelCommonMixStreamRequest) (response *CancelCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCancelCommonMixStreamRequest()
    }
    response = NewCancelCommonMixStreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCommonMixStreamRequest() (request *CreateCommonMixStreamRequest) {
    request = &CreateCommonMixStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateCommonMixStream")
    
    return
}

func NewCreateCommonMixStreamResponse() (response *CreateCommonMixStreamResponse) {
    response = &CreateCommonMixStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCommonMixStream
// This API is used to create a general stream mix. It can be used basically in the same way as the legacy `mix_streamv2.start_mix_stream_advanced` API.
//
// Note: currently, up to 16 streams can be mixed.
//
// Best practice: https://intl.cloud.tencent.com/document/product/267/45566?from_cn_redirect=1
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_GETPICTUREURLERROR = "FailedOperation.GetPictureUrlError"
//  FAILEDOPERATION_GETSTREAMRESOLUTIONERROR = "FailedOperation.GetStreamResolutionError"
//  FAILEDOPERATION_PROCESSMIXERROR = "FailedOperation.ProcessMixError"
//  FAILEDOPERATION_STREAMNOTEXIST = "FailedOperation.StreamNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JIFEIOTHERERROR = "InternalError.JiFeiOtherError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_INPUTNUMLIMITEXCEEDED = "InvalidParameter.InputNumLimitExceeded"
//  INVALIDPARAMETER_INVALIDBACKGROUDRESOLUTION = "InvalidParameter.InvalidBackgroudResolution"
//  INVALIDPARAMETER_INVALIDBITRATE = "InvalidParameter.InvalidBitrate"
//  INVALIDPARAMETER_INVALIDCROPPARAM = "InvalidParameter.InvalidCropParam"
//  INVALIDPARAMETER_INVALIDLAYERPARAM = "InvalidParameter.InvalidLayerParam"
//  INVALIDPARAMETER_INVALIDOUTPUTSTREAMID = "InvalidParameter.InvalidOutputStreamID"
//  INVALIDPARAMETER_INVALIDOUTPUTTYPE = "InvalidParameter.InvalidOutputType"
//  INVALIDPARAMETER_INVALIDPICTUREID = "InvalidParameter.InvalidPictureID"
//  INVALIDPARAMETER_INVALIDROUNDRECTRADIUS = "InvalidParameter.InvalidRoundRectRadius"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
//  INVALIDPARAMETER_SESSIONOUTPUTSTREAMCHANGED = "InvalidParameter.SessionOutputStreamChanged"
//  INVALIDPARAMETER_TEMPLATENOTMATCHINPUTNUM = "InvalidParameter.TemplateNotMatchInputNum"
func (c *Client) CreateCommonMixStream(request *CreateCommonMixStreamRequest) (response *CreateCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCreateCommonMixStreamRequest()
    }
    response = NewCreateCommonMixStreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCallbackRuleRequest() (request *CreateLiveCallbackRuleRequest) {
    request = &CreateLiveCallbackRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCallbackRule")
    
    return
}

func NewCreateLiveCallbackRuleResponse() (response *CreateLiveCallbackRuleResponse) {
    response = &CreateLiveCallbackRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveCallbackRule
// To create a callback rule, you need to first call the [CreateLiveCallbackTemplate](https://intl.cloud.tencent.com/document/product/267/32637?from_cn_redirect=1) API to create a callback template and bind the returned template ID to the domain name/path.
//
// <br>Callback protocol-related document: [Event Message Notification](https://intl.cloud.tencent.com/document/product/267/32744?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackRule(request *CreateLiveCallbackRuleRequest) (response *CreateLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackRuleRequest()
    }
    response = NewCreateLiveCallbackRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCallbackTemplateRequest() (request *CreateLiveCallbackTemplateRequest) {
    request = &CreateLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCallbackTemplate")
    
    return
}

func NewCreateLiveCallbackTemplateResponse() (response *CreateLiveCallbackTemplateResponse) {
    response = &CreateLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveCallbackTemplate
// After a callback template is created and a template ID is successfully returned, you need to call the [CreateLiveCallbackRule](https://intl.cloud.tencent.com/document/product/267/32638?from_cn_redirect=1) API and bind the template ID to the domain name/path.
//
// <br>Callback protocol document: [Event Message Notification](https://intl.cloud.tencent.com/document/product/267/32744?from_cn_redirect=1).
//
// Note: at least enter one callback URL.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackTemplate(request *CreateLiveCallbackTemplateRequest) (response *CreateLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackTemplateRequest()
    }
    response = NewCreateLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCertRequest() (request *CreateLiveCertRequest) {
    request = &CreateLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCert")
    
    return
}

func NewCreateLiveCertResponse() (response *CreateLiveCertResponse) {
    response = &CreateLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveCert
// This API is used to add a certificate.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
//  INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateLiveCert(request *CreateLiveCertRequest) (response *CreateLiveCertResponse, err error) {
    if request == nil {
        request = NewCreateLiveCertRequest()
    }
    response = NewCreateLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordRequest() (request *CreateLiveRecordRequest) {
    request = &CreateLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecord")
    
    return
}

func NewCreateLiveRecordResponse() (response *CreateLiveRecordResponse) {
    response = &CreateLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveRecord
// - Prerequisites
//
//   1. Recording files are stored on the VOD platform, so if you need to use the recording feature, you must first activate the VOD service.
//
//   2. After the recording files are stored, applicable fees (including storage fees and downstream playback traffic fees) will be charged according to the VOD billing mode. For more information, please see the [corresponding document](https://intl.cloud.tencent.com/document/product/266/2838?from_cn_redirect=1).
//
// 
//
// - Mode description
//
//   This API supports two recording modes:
//
//   1. Scheduled recording mode **(default mode)**.
//
//     The start time and end time need to be passed in, according to which the recording task will start and end automatically. Before the set end time expires (provided that `StopLiveRecord` is not called to terminate the task prematurely), the recording task is valid and will be started even after the push is interrupted and restarted multiple times.
//
//   2. Real-time video recording mode.
//
//     The start time passed in will be ignored, and recording will be started immediately after the recording task is created. The recording duration can be up to 30 minutes. If the end time passed in is more than 30 minutes after the current time, it will be counted as 30 minutes. Real-time video recording is mainly used for recording highlight scenes, and you are recommended to keep the duration within 5 minutes.
//
// 
//
// - Precautions
//
//   1. The API call timeout period should be set to more than 3 seconds; otherwise, retries and calls by different start/end time values may result in repeated recording tasks and thus incur additional recording fees.
//
//   2. Subject to the audio and video file formats (FLV/MP4/HLS), the video codec of H.264 and audio codec of AAC are supported.
//
//   3. In order to avoid malicious or unintended frequent API requests, the maximum number of tasks that can be created in scheduled recording mode is limited: up to 4,000 tasks can be created per day (excluding deleted ones), and up to 400 ones can run concurrently. If you need a higher upper limit, please submit a ticket for application.
//
//   4. This calling method does not support recording streams outside Mainland China for the time being.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  LIMITEXCEEDED_MAXIMUMTASK = "LimitExceeded.MaximumTask"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"
func (c *Client) CreateLiveRecord(request *CreateLiveRecordRequest) (response *CreateLiveRecordResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRequest()
    }
    response = NewCreateLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordRuleRequest() (request *CreateLiveRecordRuleRequest) {
    request = &CreateLiveRecordRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecordRule")
    
    return
}

func NewCreateLiveRecordRuleResponse() (response *CreateLiveRecordRuleResponse) {
    response = &CreateLiveRecordRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveRecordRule
// To create a recording rule, you need to first call the [CreateLiveRecordTemplate](https://intl.cloud.tencent.com/document/product/267/32614?from_cn_redirect=1) API to create a recording template and bind the returned template ID to the stream.
//
// <br>Recording-related document: [LVB Recording](https://intl.cloud.tencent.com/document/product/267/32739?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateLiveRecordRule(request *CreateLiveRecordRuleRequest) (response *CreateLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRuleRequest()
    }
    response = NewCreateLiveRecordRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordTemplateRequest() (request *CreateLiveRecordTemplateRequest) {
    request = &CreateLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecordTemplate")
    
    return
}

func NewCreateLiveRecordTemplateResponse() (response *CreateLiveRecordTemplateResponse) {
    response = &CreateLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveRecordTemplate
// After a recording template is created and a template ID is successfully returned, you need to call the [CreateLiveRecordRule](https://intl.cloud.tencent.com/document/product/267/32615?from_cn_redirect=1) API and bind the template ID to the stream.
//
// <br>Recording-related document: [LVB Recording](https://intl.cloud.tencent.com/document/product/267/32739?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordTemplateRequest()
    }
    response = NewCreateLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveSnapshotRuleRequest() (request *CreateLiveSnapshotRuleRequest) {
    request = &CreateLiveSnapshotRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveSnapshotRule")
    
    return
}

func NewCreateLiveSnapshotRuleResponse() (response *CreateLiveSnapshotRuleResponse) {
    response = &CreateLiveSnapshotRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveSnapshotRule
// This API is used to create a screencapturing rule. You need to first call the [CreateLiveSnapshotTemplate](https://intl.cloud.tencent.com/document/product/267/32624?from_cn_redirect=1) API to create a screencapturing template to bind the returned template ID to the stream.
//
// <br>Screencapturing document: [LVB Screencapturing](https://intl.cloud.tencent.com/document/product/267/32737?from_cn_redirect=1).
//
// Note: only one screencapturing template can be associated with one domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateLiveSnapshotRule(request *CreateLiveSnapshotRuleRequest) (response *CreateLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotRuleRequest()
    }
    response = NewCreateLiveSnapshotRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveSnapshotTemplateRequest() (request *CreateLiveSnapshotTemplateRequest) {
    request = &CreateLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveSnapshotTemplate")
    
    return
}

func NewCreateLiveSnapshotTemplateResponse() (response *CreateLiveSnapshotTemplateResponse) {
    response = &CreateLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveSnapshotTemplate
// After a screencapturing template is created and a template ID is successfully returned, you need to call the [CreateLiveSnapshotRule](https://intl.cloud.tencent.com/document/product/267/32625?from_cn_redirect=1) API and bind the template ID to the stream.
//
// <br>Screencapturing-related document: [LVB Screencapturing](https://intl.cloud.tencent.com/document/product/267/32737?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveSnapshotTemplate(request *CreateLiveSnapshotTemplateRequest) (response *CreateLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotTemplateRequest()
    }
    response = NewCreateLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTranscodeRuleRequest() (request *CreateLiveTranscodeRuleRequest) {
    request = &CreateLiveTranscodeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTranscodeRule")
    
    return
}

func NewCreateLiveTranscodeRuleResponse() (response *CreateLiveTranscodeRuleResponse) {
    response = &CreateLiveTranscodeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveTranscodeRule
// To create a transcoding rule, you need to first call the [CreateLiveTranscodeTemplate](https://intl.cloud.tencent.com/document/product/267/32646?from_cn_redirect=1) API to create a transcoding template and bind the returned template ID to the stream.
//
// <br>Transcoding-related document: [LVB Remuxing and Transcoding](https://intl.cloud.tencent.com/document/product/267/32736?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeRule(request *CreateLiveTranscodeRuleRequest) (response *CreateLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeRuleRequest()
    }
    response = NewCreateLiveTranscodeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTranscodeTemplateRequest() (request *CreateLiveTranscodeTemplateRequest) {
    request = &CreateLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTranscodeTemplate")
    
    return
}

func NewCreateLiveTranscodeTemplateResponse() (response *CreateLiveTranscodeTemplateResponse) {
    response = &CreateLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveTranscodeTemplate
// After a transcoding template is created and a template ID is successfully returned, you need to call the [CreateLiveTranscodeRule](https://intl.cloud.tencent.com/document/product/267/32647?from_cn_redirect=1) API and bind the returned template ID to the stream.
//
// <br>Transcoding-related document: [LVB Remuxing and Transcoding](https://intl.cloud.tencent.com/document/product/267/32736?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeTemplate(request *CreateLiveTranscodeTemplateRequest) (response *CreateLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeTemplateRequest()
    }
    response = NewCreateLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveWatermarkRuleRequest() (request *CreateLiveWatermarkRuleRequest) {
    request = &CreateLiveWatermarkRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveWatermarkRule")
    
    return
}

func NewCreateLiveWatermarkRuleResponse() (response *CreateLiveWatermarkRuleResponse) {
    response = &CreateLiveWatermarkRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLiveWatermarkRule
// To create a watermarking rule, you need to first call the [AddLiveWatermark](https://intl.cloud.tencent.com/document/product/267/30154?from_cn_redirect=1) API to add a watermark and bind the returned watermark ID to the stream.
//
// error code that may be returned:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveWatermarkRule(request *CreateLiveWatermarkRuleRequest) (response *CreateLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveWatermarkRuleRequest()
    }
    response = NewCreateLiveWatermarkRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordTaskRequest() (request *CreateRecordTaskRequest) {
    request = &CreateRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateRecordTask")
    
    return
}

func NewCreateRecordTaskResponse() (response *CreateRecordTaskResponse) {
    response = &CreateRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecordTask
// This API is used to create a recording task that starts and ends at specified times and records by using the configuration corresponding to a specified recording template ID.
//
// - Prerequisites
//
// 1. Recording files are stored on the VOD platform, so if you need to use the recording feature, you must first activate the VOD service.
//
// 2. After the recording files are stored, applicable fees (including storage fees and downstream playback traffic fees) are charged according to the VOD billing method. For details, see the [corresponding document](https://intl.cloud.tencent.com/document/product/266/2837?from_cn_redirect=1).
//
// - Precautions
//
// 1. An interruption will end the current recording and generate a recording file. The task will still be valid before the end time expires, and the stream will be recorded within this period as long as it is pushed, regardless of whether the push is interrupted or restarted multiple times.
//
// 2. Avoid creating recording tasks with overlapping time periods. If there are multiple tasks with overlapping time periods for the same stream, the system will start three recording tasks at most to avoid repeated recording.
//
// 3. The record of a created recording task will be retained for 3 months on the platform.
//
// 4. The current recording task management APIs (CreateRecordTask/StopRecordTask/DeleteRecordTask) are not compatible with the legacy APIs (CreateLiveRecord/StopLiveRecord/DeleteLiveRecord), and they cannot be used together.
//
// 5. Do not create recording tasks and push streams at the same time, or recording tasks might not take effect and be delayed. Wait at least 3 seconds between each action.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRecordTask(request *CreateRecordTaskRequest) (response *CreateRecordTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecordTaskRequest()
    }
    response = NewCreateRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCallbackRuleRequest() (request *DeleteLiveCallbackRuleRequest) {
    request = &DeleteLiveCallbackRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCallbackRule")
    
    return
}

func NewDeleteLiveCallbackRuleResponse() (response *DeleteLiveCallbackRuleResponse) {
    response = &DeleteLiveCallbackRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveCallbackRule
// This API is used to delete a callback rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveCallbackRule(request *DeleteLiveCallbackRuleRequest) (response *DeleteLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackRuleRequest()
    }
    response = NewDeleteLiveCallbackRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCallbackTemplateRequest() (request *DeleteLiveCallbackTemplateRequest) {
    request = &DeleteLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCallbackTemplate")
    
    return
}

func NewDeleteLiveCallbackTemplateResponse() (response *DeleteLiveCallbackTemplateResponse) {
    response = &DeleteLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveCallbackTemplate
// This API deletes a callback template.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveCallbackTemplate(request *DeleteLiveCallbackTemplateRequest) (response *DeleteLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackTemplateRequest()
    }
    response = NewDeleteLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCertRequest() (request *DeleteLiveCertRequest) {
    request = &DeleteLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCert")
    
    return
}

func NewDeleteLiveCertResponse() (response *DeleteLiveCertResponse) {
    response = &DeleteLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveCert
// This API is used to delete a certificate corresponding to the domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
func (c *Client) DeleteLiveCert(request *DeleteLiveCertRequest) (response *DeleteLiveCertResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCertRequest()
    }
    response = NewDeleteLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveDomainRequest() (request *DeleteLiveDomainRequest) {
    request = &DeleteLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveDomain")
    
    return
}

func NewDeleteLiveDomainResponse() (response *DeleteLiveDomainResponse) {
    response = &DeleteLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveDomain
// This API is used to delete an added LVB domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_JIFEINOENOUGHFUND = "FailedOperation.JiFeiNoEnoughFund"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
func (c *Client) DeleteLiveDomain(request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    if request == nil {
        request = NewDeleteLiveDomainRequest()
    }
    response = NewDeleteLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordRequest() (request *DeleteLiveRecordRequest) {
    request = &DeleteLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecord")
    
    return
}

func NewDeleteLiveRecordResponse() (response *DeleteLiveRecordResponse) {
    response = &DeleteLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveRecord
// Note: The `DeleteLiveRecord` API is only used to delete the record of recording tasks but not stop recording or deleting an ongoing recording task. If you need to stop a recording task, please use the [StopLiveRecord](https://intl.cloud.tencent.com/document/product/267/30146?from_cn_redirect=1) API.
//
// error code that may be returned:
//  FAILEDOPERATION_ALTERTASKSTATE = "FailedOperation.AlterTaskState"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
func (c *Client) DeleteLiveRecord(request *DeleteLiveRecordRequest) (response *DeleteLiveRecordResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRequest()
    }
    response = NewDeleteLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordRuleRequest() (request *DeleteLiveRecordRuleRequest) {
    request = &DeleteLiveRecordRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecordRule")
    
    return
}

func NewDeleteLiveRecordRuleResponse() (response *DeleteLiveRecordRuleResponse) {
    response = &DeleteLiveRecordRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveRecordRule
// This API is used to delete a recording rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveRecordRule(request *DeleteLiveRecordRuleRequest) (response *DeleteLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRuleRequest()
    }
    response = NewDeleteLiveRecordRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordTemplateRequest() (request *DeleteLiveRecordTemplateRequest) {
    request = &DeleteLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecordTemplate")
    
    return
}

func NewDeleteLiveRecordTemplateResponse() (response *DeleteLiveRecordTemplateResponse) {
    response = &DeleteLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveRecordTemplate
// This API is used to delete a recording template.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordTemplateRequest()
    }
    response = NewDeleteLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveSnapshotRuleRequest() (request *DeleteLiveSnapshotRuleRequest) {
    request = &DeleteLiveSnapshotRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveSnapshotRule")
    
    return
}

func NewDeleteLiveSnapshotRuleResponse() (response *DeleteLiveSnapshotRuleResponse) {
    response = &DeleteLiveSnapshotRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveSnapshotRule
// This API is used to delete a screencapturing rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveSnapshotRule(request *DeleteLiveSnapshotRuleRequest) (response *DeleteLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotRuleRequest()
    }
    response = NewDeleteLiveSnapshotRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveSnapshotTemplateRequest() (request *DeleteLiveSnapshotTemplateRequest) {
    request = &DeleteLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveSnapshotTemplate")
    
    return
}

func NewDeleteLiveSnapshotTemplateResponse() (response *DeleteLiveSnapshotTemplateResponse) {
    response = &DeleteLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveSnapshotTemplate
// This API is used to delete a screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveSnapshotTemplate(request *DeleteLiveSnapshotTemplateRequest) (response *DeleteLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotTemplateRequest()
    }
    response = NewDeleteLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTranscodeRuleRequest() (request *DeleteLiveTranscodeRuleRequest) {
    request = &DeleteLiveTranscodeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTranscodeRule")
    
    return
}

func NewDeleteLiveTranscodeRuleResponse() (response *DeleteLiveTranscodeRuleResponse) {
    response = &DeleteLiveTranscodeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveTranscodeRule
// This API is used to delete a transcoding rule.
//
// `DomainName+AppName+StreamName+TemplateId` uniquely identifies a single transcoding rule. If you need to delete it, strong match is required. `TemplateId` is required. Even if other parameters are empty, you still need to pass in an empty string to make a strong match.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveTranscodeRule(request *DeleteLiveTranscodeRuleRequest) (response *DeleteLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeRuleRequest()
    }
    response = NewDeleteLiveTranscodeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTranscodeTemplateRequest() (request *DeleteLiveTranscodeTemplateRequest) {
    request = &DeleteLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTranscodeTemplate")
    
    return
}

func NewDeleteLiveTranscodeTemplateResponse() (response *DeleteLiveTranscodeTemplateResponse) {
    response = &DeleteLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveTranscodeTemplate
// This API is used to delete a transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveTranscodeTemplate(request *DeleteLiveTranscodeTemplateRequest) (response *DeleteLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeTemplateRequest()
    }
    response = NewDeleteLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveWatermarkRequest() (request *DeleteLiveWatermarkRequest) {
    request = &DeleteLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveWatermark")
    
    return
}

func NewDeleteLiveWatermarkResponse() (response *DeleteLiveWatermarkResponse) {
    response = &DeleteLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveWatermark
// This API is used to delete a watermark.
//
// error code that may be returned:
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETWATERMARKERROR = "InternalError.GetWatermarkError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) DeleteLiveWatermark(request *DeleteLiveWatermarkRequest) (response *DeleteLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRequest()
    }
    response = NewDeleteLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveWatermarkRuleRequest() (request *DeleteLiveWatermarkRuleRequest) {
    request = &DeleteLiveWatermarkRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveWatermarkRule")
    
    return
}

func NewDeleteLiveWatermarkRuleResponse() (response *DeleteLiveWatermarkRuleResponse) {
    response = &DeleteLiveWatermarkRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLiveWatermarkRule
// This API is used to delete a watermarking rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteLiveWatermarkRule(request *DeleteLiveWatermarkRuleRequest) (response *DeleteLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRuleRequest()
    }
    response = NewDeleteLiveWatermarkRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordTaskRequest() (request *DeleteRecordTaskRequest) {
    request = &DeleteRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteRecordTask")
    
    return
}

func NewDeleteRecordTaskResponse() (response *DeleteRecordTaskResponse) {
    response = &DeleteRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordTask
// This API is used to delete a recording task configuration. The deletion does not affect running tasks and takes effect only for new pushes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordTask(request *DeleteRecordTaskRequest) (response *DeleteRecordTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordTaskRequest()
    }
    response = NewDeleteRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllStreamPlayInfoListRequest() (request *DescribeAllStreamPlayInfoListRequest) {
    request = &DescribeAllStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeAllStreamPlayInfoList")
    
    return
}

func NewDescribeAllStreamPlayInfoListResponse() (response *DescribeAllStreamPlayInfoListResponse) {
    response = &DescribeAllStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllStreamPlayInfoList
// This API is used to query the downstream information of all streams at a specified point in time (at a 1-minute granularity).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllStreamPlayInfoList(request *DescribeAllStreamPlayInfoListRequest) (response *DescribeAllStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAllStreamPlayInfoListRequest()
    }
    response = NewDescribeAllStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAreaBillBandwidthAndFluxListRequest() (request *DescribeAreaBillBandwidthAndFluxListRequest) {
    request = &DescribeAreaBillBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeAreaBillBandwidthAndFluxList")
    
    return
}

func NewDescribeAreaBillBandwidthAndFluxListResponse() (response *DescribeAreaBillBandwidthAndFluxListResponse) {
    response = &DescribeAreaBillBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAreaBillBandwidthAndFluxList
// This API is used to query the billable LVB bandwidth and traffic data outside Chinese mainland.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAreaBillBandwidthAndFluxList(request *DescribeAreaBillBandwidthAndFluxListRequest) (response *DescribeAreaBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeAreaBillBandwidthAndFluxListRequest()
    }
    response = NewDescribeAreaBillBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillBandwidthAndFluxListRequest() (request *DescribeBillBandwidthAndFluxListRequest) {
    request = &DescribeBillBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeBillBandwidthAndFluxList")
    
    return
}

func NewDescribeBillBandwidthAndFluxListResponse() (response *DescribeBillBandwidthAndFluxListResponse) {
    response = &DescribeBillBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillBandwidthAndFluxList
// This API is used to query the data of billable LVB bandwidth and traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillBandwidthAndFluxList(request *DescribeBillBandwidthAndFluxListRequest) (response *DescribeBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeBillBandwidthAndFluxListRequest()
    }
    response = NewDescribeBillBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrentRecordStreamNumRequest() (request *DescribeConcurrentRecordStreamNumRequest) {
    request = &DescribeConcurrentRecordStreamNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeConcurrentRecordStreamNum")
    
    return
}

func NewDescribeConcurrentRecordStreamNumResponse() (response *DescribeConcurrentRecordStreamNumResponse) {
    response = &DescribeConcurrentRecordStreamNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConcurrentRecordStreamNum
// This API is used to query the number of concurrent recording channels, which is applicable to LCB and LVB.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeConcurrentRecordStreamNum(request *DescribeConcurrentRecordStreamNumRequest) (response *DescribeConcurrentRecordStreamNumResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentRecordStreamNumRequest()
    }
    response = NewDescribeConcurrentRecordStreamNumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliverBandwidthListRequest() (request *DescribeDeliverBandwidthListRequest) {
    request = &DescribeDeliverBandwidthListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeDeliverBandwidthList")
    
    return
}

func NewDescribeDeliverBandwidthListResponse() (response *DescribeDeliverBandwidthListResponse) {
    response = &DescribeDeliverBandwidthListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeliverBandwidthList
// This API is used to query the billable bandwidth of live stream relaying in the last 3 months. The query period is up to 31 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeliverBandwidthList(request *DescribeDeliverBandwidthListRequest) (response *DescribeDeliverBandwidthListResponse, err error) {
    if request == nil {
        request = NewDescribeDeliverBandwidthListRequest()
    }
    response = NewDescribeDeliverBandwidthListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupProIspPlayInfoListRequest() (request *DescribeGroupProIspPlayInfoListRequest) {
    request = &DescribeGroupProIspPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeGroupProIspPlayInfoList")
    
    return
}

func NewDescribeGroupProIspPlayInfoListResponse() (response *DescribeGroupProIspPlayInfoListResponse) {
    response = &DescribeGroupProIspPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupProIspPlayInfoList
// This API is used to query the downstream playback data by district and ISP.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupProIspPlayInfoList(request *DescribeGroupProIspPlayInfoListRequest) (response *DescribeGroupProIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupProIspPlayInfoListRequest()
    }
    response = NewDescribeGroupProIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHttpStatusInfoListRequest() (request *DescribeHttpStatusInfoListRequest) {
    request = &DescribeHttpStatusInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeHttpStatusInfoList")
    
    return
}

func NewDescribeHttpStatusInfoListResponse() (response *DescribeHttpStatusInfoListResponse) {
    response = &DescribeHttpStatusInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHttpStatusInfoList
// This API is used to query the number of each playback HTTP status code at a 5-minute granularity in a certain period of time.
//
// Note: data can be queried one hour after it is generated. For example, data between 10:00 and 10:59 cannot be queried until 12:00.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHttpStatusInfoList(request *DescribeHttpStatusInfoListRequest) (response *DescribeHttpStatusInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeHttpStatusInfoListRequest()
    }
    response = NewDescribeHttpStatusInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackRulesRequest() (request *DescribeLiveCallbackRulesRequest) {
    request = &DescribeLiveCallbackRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackRules")
    
    return
}

func NewDescribeLiveCallbackRulesResponse() (response *DescribeLiveCallbackRulesResponse) {
    response = &DescribeLiveCallbackRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveCallbackRules
// This API is used to get the callback rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackRules(request *DescribeLiveCallbackRulesRequest) (response *DescribeLiveCallbackRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackRulesRequest()
    }
    response = NewDescribeLiveCallbackRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackTemplateRequest() (request *DescribeLiveCallbackTemplateRequest) {
    request = &DescribeLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackTemplate")
    
    return
}

func NewDescribeLiveCallbackTemplateResponse() (response *DescribeLiveCallbackTemplateResponse) {
    response = &DescribeLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveCallbackTemplate
// This API is used to get a single callback template.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveCallbackTemplate(request *DescribeLiveCallbackTemplateRequest) (response *DescribeLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplateRequest()
    }
    response = NewDescribeLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackTemplatesRequest() (request *DescribeLiveCallbackTemplatesRequest) {
    request = &DescribeLiveCallbackTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackTemplates")
    
    return
}

func NewDescribeLiveCallbackTemplatesResponse() (response *DescribeLiveCallbackTemplatesResponse) {
    response = &DescribeLiveCallbackTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveCallbackTemplates
// This API is used to get the callback template list.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackTemplates(request *DescribeLiveCallbackTemplatesRequest) (response *DescribeLiveCallbackTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplatesRequest()
    }
    response = NewDescribeLiveCallbackTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCertRequest() (request *DescribeLiveCertRequest) {
    request = &DescribeLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCert")
    
    return
}

func NewDescribeLiveCertResponse() (response *DescribeLiveCertResponse) {
    response = &DescribeLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveCert
// This API is used to get certificate information.
//
// error code that may be returned:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) DescribeLiveCert(request *DescribeLiveCertRequest) (response *DescribeLiveCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertRequest()
    }
    response = NewDescribeLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCertsRequest() (request *DescribeLiveCertsRequest) {
    request = &DescribeLiveCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCerts")
    
    return
}

func NewDescribeLiveCertsResponse() (response *DescribeLiveCertsResponse) {
    response = &DescribeLiveCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveCerts
// This API is used to get the certificate information list.
//
// error code that may be returned:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) DescribeLiveCerts(request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertsRequest()
    }
    response = NewDescribeLiveCertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDelayInfoListRequest() (request *DescribeLiveDelayInfoListRequest) {
    request = &DescribeLiveDelayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDelayInfoList")
    
    return
}

func NewDescribeLiveDelayInfoListResponse() (response *DescribeLiveDelayInfoListResponse) {
    response = &DescribeLiveDelayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDelayInfoList
// This API is used to get the list of delayed playbacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveDelayInfoList(request *DescribeLiveDelayInfoListRequest) (response *DescribeLiveDelayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDelayInfoListRequest()
    }
    response = NewDescribeLiveDelayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainRequest() (request *DescribeLiveDomainRequest) {
    request = &DescribeLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomain")
    
    return
}

func NewDescribeLiveDomainResponse() (response *DescribeLiveDomainResponse) {
    response = &DescribeLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomain
// This API is used to query the LVB domain name information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) DescribeLiveDomain(request *DescribeLiveDomainRequest) (response *DescribeLiveDomainResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRequest()
    }
    response = NewDescribeLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainCertRequest() (request *DescribeLiveDomainCertRequest) {
    request = &DescribeLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainCert")
    
    return
}

func NewDescribeLiveDomainCertResponse() (response *DescribeLiveDomainCertResponse) {
    response = &DescribeLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomainCert
// This API is used to get the domain name certificate information.
//
// error code that may be returned:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) DescribeLiveDomainCert(request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertRequest()
    }
    response = NewDescribeLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainPlayInfoListRequest() (request *DescribeLiveDomainPlayInfoListRequest) {
    request = &DescribeLiveDomainPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainPlayInfoList")
    
    return
}

func NewDescribeLiveDomainPlayInfoListResponse() (response *DescribeLiveDomainPlayInfoListResponse) {
    response = &DescribeLiveDomainPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomainPlayInfoList
// This API is used to query the real-time downstream playback data at the domain name level. As it takes certain time to process data, the API queries quasi-real-time data generated 4 minutes ago by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveDomainPlayInfoList(request *DescribeLiveDomainPlayInfoListRequest) (response *DescribeLiveDomainPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainPlayInfoListRequest()
    }
    response = NewDescribeLiveDomainPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainRefererRequest() (request *DescribeLiveDomainRefererRequest) {
    request = &DescribeLiveDomainRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainReferer")
    
    return
}

func NewDescribeLiveDomainRefererResponse() (response *DescribeLiveDomainRefererResponse) {
    response = &DescribeLiveDomainRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomainReferer
// This API is used to query referer allowlist/blocklist configuration of a live streaming domain name.
//
// Referer information is included in HTTP requests. After you enable referer configuration, live streams using RTMP or WebRTC for playback will not authenticate the referer and can be played back normally. To make the referer configuration effective, the HTTP-FLV or HLS protocol is recommended for playback.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) DescribeLiveDomainReferer(request *DescribeLiveDomainRefererRequest) (response *DescribeLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRefererRequest()
    }
    response = NewDescribeLiveDomainRefererResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainsRequest() (request *DescribeLiveDomainsRequest) {
    request = &DescribeLiveDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomains")
    
    return
}

func NewDescribeLiveDomainsResponse() (response *DescribeLiveDomainsResponse) {
    response = &DescribeLiveDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomains
// This API is used to query domain names by domain name status and type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
func (c *Client) DescribeLiveDomains(request *DescribeLiveDomainsRequest) (response *DescribeLiveDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainsRequest()
    }
    response = NewDescribeLiveDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveForbidStreamListRequest() (request *DescribeLiveForbidStreamListRequest) {
    request = &DescribeLiveForbidStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveForbidStreamList")
    
    return
}

func NewDescribeLiveForbidStreamListResponse() (response *DescribeLiveForbidStreamListResponse) {
    response = &DescribeLiveForbidStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveForbidStreamList
// This API is used to get the list of disabled streams.
//
// 
//
// Note: this API is used for query only and should not be relied too much upon in important business scenarios.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveForbidStreamList(request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveForbidStreamListRequest()
    }
    response = NewDescribeLiveForbidStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePlayAuthKeyRequest() (request *DescribeLivePlayAuthKeyRequest) {
    request = &DescribeLivePlayAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePlayAuthKey")
    
    return
}

func NewDescribeLivePlayAuthKeyResponse() (response *DescribeLivePlayAuthKeyResponse) {
    response = &DescribeLivePlayAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLivePlayAuthKey
// This API is used to query the playback authentication key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
func (c *Client) DescribeLivePlayAuthKey(request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePlayAuthKeyRequest()
    }
    response = NewDescribeLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePushAuthKeyRequest() (request *DescribeLivePushAuthKeyRequest) {
    request = &DescribeLivePushAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePushAuthKey")
    
    return
}

func NewDescribeLivePushAuthKeyResponse() (response *DescribeLivePushAuthKeyResponse) {
    response = &DescribeLivePushAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLivePushAuthKey
// This API is used to query the LVB push authentication key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
func (c *Client) DescribeLivePushAuthKey(request *DescribeLivePushAuthKeyRequest) (response *DescribeLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePushAuthKeyRequest()
    }
    response = NewDescribeLivePushAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordRulesRequest() (request *DescribeLiveRecordRulesRequest) {
    request = &DescribeLiveRecordRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordRules")
    
    return
}

func NewDescribeLiveRecordRulesResponse() (response *DescribeLiveRecordRulesResponse) {
    response = &DescribeLiveRecordRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveRecordRules
// This API is used to get the list of recording rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveRecordRules(request *DescribeLiveRecordRulesRequest) (response *DescribeLiveRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordRulesRequest()
    }
    response = NewDescribeLiveRecordRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplateRequest() (request *DescribeLiveRecordTemplateRequest) {
    request = &DescribeLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordTemplate")
    
    return
}

func NewDescribeLiveRecordTemplateResponse() (response *DescribeLiveRecordTemplateResponse) {
    response = &DescribeLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveRecordTemplate
// This API is used to get a single recording template.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveRecordTemplate(request *DescribeLiveRecordTemplateRequest) (response *DescribeLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplateRequest()
    }
    response = NewDescribeLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplatesRequest() (request *DescribeLiveRecordTemplatesRequest) {
    request = &DescribeLiveRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordTemplates")
    
    return
}

func NewDescribeLiveRecordTemplatesResponse() (response *DescribeLiveRecordTemplatesResponse) {
    response = &DescribeLiveRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveRecordTemplates
// This API is used to get the recording template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveRecordTemplates(request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplatesRequest()
    }
    response = NewDescribeLiveRecordTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotRulesRequest() (request *DescribeLiveSnapshotRulesRequest) {
    request = &DescribeLiveSnapshotRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotRules")
    
    return
}

func NewDescribeLiveSnapshotRulesResponse() (response *DescribeLiveSnapshotRulesResponse) {
    response = &DescribeLiveSnapshotRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveSnapshotRules
// This API is used to get the screencapturing rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveSnapshotRules(request *DescribeLiveSnapshotRulesRequest) (response *DescribeLiveSnapshotRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotRulesRequest()
    }
    response = NewDescribeLiveSnapshotRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotTemplateRequest() (request *DescribeLiveSnapshotTemplateRequest) {
    request = &DescribeLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotTemplate")
    
    return
}

func NewDescribeLiveSnapshotTemplateResponse() (response *DescribeLiveSnapshotTemplateResponse) {
    response = &DescribeLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveSnapshotTemplate
// This API is used to get a single screencapturing template.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveSnapshotTemplate(request *DescribeLiveSnapshotTemplateRequest) (response *DescribeLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplateRequest()
    }
    response = NewDescribeLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotTemplatesRequest() (request *DescribeLiveSnapshotTemplatesRequest) {
    request = &DescribeLiveSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotTemplates")
    
    return
}

func NewDescribeLiveSnapshotTemplatesResponse() (response *DescribeLiveSnapshotTemplatesResponse) {
    response = &DescribeLiveSnapshotTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveSnapshotTemplates
// This API is used to get the screencapturing template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplates(request *DescribeLiveSnapshotTemplatesRequest) (response *DescribeLiveSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplatesRequest()
    }
    response = NewDescribeLiveSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamEventListRequest() (request *DescribeLiveStreamEventListRequest) {
    request = &DescribeLiveStreamEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamEventList")
    
    return
}

func NewDescribeLiveStreamEventListResponse() (response *DescribeLiveStreamEventListResponse) {
    response = &DescribeLiveStreamEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveStreamEventList
// This API is used to query stream push/interruption events.<br>
//
// 
//
// Notes:
//
// 1. This API is used to query stream push/interruption records, and should not be relied too much upon in important business scenarios.
//
// 2. You can use the `IsFilter` parameter of this API to filter and get required push records.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamEventList(request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamEventListRequest()
    }
    response = NewDescribeLiveStreamEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamOnlineListRequest() (request *DescribeLiveStreamOnlineListRequest) {
    request = &DescribeLiveStreamOnlineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamOnlineList")
    
    return
}

func NewDescribeLiveStreamOnlineListResponse() (response *DescribeLiveStreamOnlineListResponse) {
    response = &DescribeLiveStreamOnlineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveStreamOnlineList
// This API is used to get the list of ongoing live streams. It queries the information of live streams after they are pushed successfully. 
//
// 
//
// Notes:
//
// 1. This API is used to query the list of active live streams only, and should not be relied too much upon in important business scenarios.
//
// 2. This API can query up to 20,000 streams. To query more streams, please contact our after-sales service team.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamOnlineList(request *DescribeLiveStreamOnlineListRequest) (response *DescribeLiveStreamOnlineListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamOnlineListRequest()
    }
    response = NewDescribeLiveStreamOnlineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamPublishedListRequest() (request *DescribeLiveStreamPublishedListRequest) {
    request = &DescribeLiveStreamPublishedListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamPublishedList")
    
    return
}

func NewDescribeLiveStreamPublishedListResponse() (response *DescribeLiveStreamPublishedListResponse) {
    response = &DescribeLiveStreamPublishedListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveStreamPublishedList
// This API is used to return the list of pushed streams. <br>
//
// Note: Up to 10,000 entries can be queried per page. More data can be obtained by adjusting the query time range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamPublishedList(request *DescribeLiveStreamPublishedListRequest) (response *DescribeLiveStreamPublishedListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPublishedListRequest()
    }
    response = NewDescribeLiveStreamPublishedListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamPushInfoListRequest() (request *DescribeLiveStreamPushInfoListRequest) {
    request = &DescribeLiveStreamPushInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamPushInfoList")
    
    return
}

func NewDescribeLiveStreamPushInfoListResponse() (response *DescribeLiveStreamPushInfoListResponse) {
    response = &DescribeLiveStreamPushInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveStreamPushInfoList
// This API is used to query the push information of all real-time streams, including client IP, server IP, frame rate, bitrate, domain name, and push start time.
//
// error code that may be returned:
//  FAILEDOPERATION_HASNOTLIVINGSTREAM = "FailedOperation.HasNotLivingStream"
//  FAILEDOPERATION_QUERYUPLOADINFOFAILED = "FailedOperation.QueryUploadInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveStreamPushInfoList(request *DescribeLiveStreamPushInfoListRequest) (response *DescribeLiveStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPushInfoListRequest()
    }
    response = NewDescribeLiveStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamStateRequest() (request *DescribeLiveStreamStateRequest) {
    request = &DescribeLiveStreamStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamState")
    
    return
}

func NewDescribeLiveStreamStateResponse() (response *DescribeLiveStreamStateResponse) {
    response = &DescribeLiveStreamStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveStreamState
// This API is used to get the stream status, such as active, inactive, or disabled.
//
// Note: this API is used to query the stream status, and should not be relied too much upon in important business scenarios.
//
// 
//
// Notes:
//
// 1. You should not rely on the results returned by this API to initiate/interrupt live streams.
//
// 2. The application can get and store the status of live rooms via [Stream Push and Interruption Event Notification](https://intl.cloud.tencent.com/document/product/267/47025?from_cn_redirect=1).
//
// 3. You can use the [DescribeLiveStreamOnlineList](https://intl.cloud.tencent.com/document/product/267/20472?from_cn_redirect=1) API to regularly (with the interval larger than 1 minute) check the status of live rooms monitored by the application.
//
// 4. If you find that a stream is inactive using the stream status query API, you can use other above-mentioned methods to check its status.
//
// 5. If access or resolution errors occur when you use the API to query, you can regard the stream as active, and do not perform operations on the application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamState(request *DescribeLiveStreamStateRequest) (response *DescribeLiveStreamStateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamStateRequest()
    }
    response = NewDescribeLiveStreamStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeDetailInfoRequest() (request *DescribeLiveTranscodeDetailInfoRequest) {
    request = &DescribeLiveTranscodeDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeDetailInfo")
    
    return
}

func NewDescribeLiveTranscodeDetailInfoResponse() (response *DescribeLiveTranscodeDetailInfoResponse) {
    response = &DescribeLiveTranscodeDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveTranscodeDetailInfo
// This API is used to query the details of transcoding on a specified day or in a specified period of time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLiveTranscodeDetailInfo(request *DescribeLiveTranscodeDetailInfoRequest) (response *DescribeLiveTranscodeDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeDetailInfoRequest()
    }
    response = NewDescribeLiveTranscodeDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeRulesRequest() (request *DescribeLiveTranscodeRulesRequest) {
    request = &DescribeLiveTranscodeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeRules")
    
    return
}

func NewDescribeLiveTranscodeRulesResponse() (response *DescribeLiveTranscodeRulesResponse) {
    response = &DescribeLiveTranscodeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveTranscodeRules
// This API is used to get the list of transcoding rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveTranscodeRules(request *DescribeLiveTranscodeRulesRequest) (response *DescribeLiveTranscodeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeRulesRequest()
    }
    response = NewDescribeLiveTranscodeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTemplateRequest() (request *DescribeLiveTranscodeTemplateRequest) {
    request = &DescribeLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTemplate")
    
    return
}

func NewDescribeLiveTranscodeTemplateResponse() (response *DescribeLiveTranscodeTemplateResponse) {
    response = &DescribeLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveTranscodeTemplate
// This API is used to get a single transcoding template.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplate(request *DescribeLiveTranscodeTemplateRequest) (response *DescribeLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplateRequest()
    }
    response = NewDescribeLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTemplatesRequest() (request *DescribeLiveTranscodeTemplatesRequest) {
    request = &DescribeLiveTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTemplates")
    
    return
}

func NewDescribeLiveTranscodeTemplatesResponse() (response *DescribeLiveTranscodeTemplatesResponse) {
    response = &DescribeLiveTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveTranscodeTemplates
// This API is used to get the transcoding template list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplates(request *DescribeLiveTranscodeTemplatesRequest) (response *DescribeLiveTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplatesRequest()
    }
    response = NewDescribeLiveTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarkRequest() (request *DescribeLiveWatermarkRequest) {
    request = &DescribeLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermark")
    
    return
}

func NewDescribeLiveWatermarkResponse() (response *DescribeLiveWatermarkResponse) {
    response = &DescribeLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveWatermark
// This API is used to get the information of a single watermark.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveWatermark(request *DescribeLiveWatermarkRequest) (response *DescribeLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRequest()
    }
    response = NewDescribeLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarkRulesRequest() (request *DescribeLiveWatermarkRulesRequest) {
    request = &DescribeLiveWatermarkRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermarkRules")
    
    return
}

func NewDescribeLiveWatermarkRulesResponse() (response *DescribeLiveWatermarkRulesResponse) {
    response = &DescribeLiveWatermarkRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveWatermarkRules
// This API is used to get the watermarking rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveWatermarkRules(request *DescribeLiveWatermarkRulesRequest) (response *DescribeLiveWatermarkRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRulesRequest()
    }
    response = NewDescribeLiveWatermarkRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarksRequest() (request *DescribeLiveWatermarksRequest) {
    request = &DescribeLiveWatermarksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermarks")
    
    return
}

func NewDescribeLiveWatermarksResponse() (response *DescribeLiveWatermarksResponse) {
    response = &DescribeLiveWatermarksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveWatermarks
// This API is used to query the watermark list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveWatermarks(request *DescribeLiveWatermarksRequest) (response *DescribeLiveWatermarksResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarksRequest()
    }
    response = NewDescribeLiveWatermarksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayErrorCodeDetailInfoListRequest() (request *DescribePlayErrorCodeDetailInfoListRequest) {
    request = &DescribePlayErrorCodeDetailInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribePlayErrorCodeDetailInfoList")
    
    return
}

func NewDescribePlayErrorCodeDetailInfoListResponse() (response *DescribePlayErrorCodeDetailInfoListResponse) {
    response = &DescribePlayErrorCodeDetailInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlayErrorCodeDetailInfoList
// This API is used to query the information of downstream playback error codes, i.e., the occurrences of each HTTP error code (4xx and 5xx) at a 1-minute granularity in a certain period of time.
//
// 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePlayErrorCodeDetailInfoList(request *DescribePlayErrorCodeDetailInfoListRequest) (response *DescribePlayErrorCodeDetailInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeDetailInfoListRequest()
    }
    response = NewDescribePlayErrorCodeDetailInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayErrorCodeSumInfoListRequest() (request *DescribePlayErrorCodeSumInfoListRequest) {
    request = &DescribePlayErrorCodeSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribePlayErrorCodeSumInfoList")
    
    return
}

func NewDescribePlayErrorCodeSumInfoListResponse() (response *DescribePlayErrorCodeSumInfoListResponse) {
    response = &DescribePlayErrorCodeSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlayErrorCodeSumInfoList
// This API is used to query the information of downstream playback error codes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePlayErrorCodeSumInfoList(request *DescribePlayErrorCodeSumInfoListRequest) (response *DescribePlayErrorCodeSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeSumInfoListRequest()
    }
    response = NewDescribePlayErrorCodeSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProIspPlaySumInfoListRequest() (request *DescribeProIspPlaySumInfoListRequest) {
    request = &DescribeProIspPlaySumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeProIspPlaySumInfoList")
    
    return
}

func NewDescribeProIspPlaySumInfoListResponse() (response *DescribeProIspPlaySumInfoListResponse) {
    response = &DescribeProIspPlaySumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProIspPlaySumInfoList
// This API is used to query the average traffic per second, total traffic, and number of total requests by country/region, district, and ISP in a certain period of time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProIspPlaySumInfoList(request *DescribeProIspPlaySumInfoListRequest) (response *DescribeProIspPlaySumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProIspPlaySumInfoListRequest()
    }
    response = NewDescribeProIspPlaySumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProvinceIspPlayInfoListRequest() (request *DescribeProvinceIspPlayInfoListRequest) {
    request = &DescribeProvinceIspPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeProvinceIspPlayInfoList")
    
    return
}

func NewDescribeProvinceIspPlayInfoListResponse() (response *DescribeProvinceIspPlayInfoListResponse) {
    response = &DescribeProvinceIspPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProvinceIspPlayInfoList
// This API is used to query the downstream playback data of a specified ISP in a specified district, including bandwidth, traffic, number of requests, and number of concurrent connections.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_HASNOTLIVINGSTREAM = "InternalError.HasNotLivingStream"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeProvinceIspPlayInfoList(request *DescribeProvinceIspPlayInfoListRequest) (response *DescribeProvinceIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProvinceIspPlayInfoListRequest()
    }
    response = NewDescribeProvinceIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenShotSheetNumListRequest() (request *DescribeScreenShotSheetNumListRequest) {
    request = &DescribeScreenShotSheetNumListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeScreenShotSheetNumList")
    
    return
}

func NewDescribeScreenShotSheetNumListResponse() (response *DescribeScreenShotSheetNumListResponse) {
    response = &DescribeScreenShotSheetNumListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenShotSheetNumList
// The API is used to query the number of screenshots as an LVB value-added service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScreenShotSheetNumList(request *DescribeScreenShotSheetNumListRequest) (response *DescribeScreenShotSheetNumListResponse, err error) {
    if request == nil {
        request = NewDescribeScreenShotSheetNumListRequest()
    }
    response = NewDescribeScreenShotSheetNumListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamDayPlayInfoListRequest() (request *DescribeStreamDayPlayInfoListRequest) {
    request = &DescribeStreamDayPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamDayPlayInfoList")
    
    return
}

func NewDescribeStreamDayPlayInfoListResponse() (response *DescribeStreamDayPlayInfoListResponse) {
    response = &DescribeStreamDayPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamDayPlayInfoList
// This API is used to query the playback data of each stream at the day level, including the total traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeStreamDayPlayInfoList(request *DescribeStreamDayPlayInfoListRequest) (response *DescribeStreamDayPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamDayPlayInfoListRequest()
    }
    response = NewDescribeStreamDayPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPlayInfoListRequest() (request *DescribeStreamPlayInfoListRequest) {
    request = &DescribeStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamPlayInfoList")
    
    return
}

func NewDescribeStreamPlayInfoListResponse() (response *DescribeStreamPlayInfoListResponse) {
    response = &DescribeStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamPlayInfoList
// This API is used to query the playback data. It supports querying the playback details by stream name and aggregated data by playback domain name. Data in the last 4 minutes or so cannot be queried due to delay.
//
// Note: to query by `AppName`, you need to submit a ticket first. After your application succeeds, it will take about 5 business days (subject to the time in the reply) for the configuration to take effect.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    response = NewDescribeStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPushInfoListRequest() (request *DescribeStreamPushInfoListRequest) {
    request = &DescribeStreamPushInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamPushInfoList")
    
    return
}

func NewDescribeStreamPushInfoListResponse() (response *DescribeStreamPushInfoListResponse) {
    response = &DescribeStreamPushInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamPushInfoList
// This API is used to query the upstream push quality data by stream ID, including frame rate, bitrate, elapsed time, and codec of audio and video files.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStreamPushInfoList(request *DescribeStreamPushInfoListRequest) (response *DescribeStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPushInfoListRequest()
    }
    response = NewDescribeStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopClientIpSumInfoListRequest() (request *DescribeTopClientIpSumInfoListRequest) {
    request = &DescribeTopClientIpSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeTopClientIpSumInfoList")
    
    return
}

func NewDescribeTopClientIpSumInfoListResponse() (response *DescribeTopClientIpSumInfoListResponse) {
    response = &DescribeTopClientIpSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopClientIpSumInfoList
// This API is used to query the information of the top n client IPs in a certain period of time (top 1,000 is supported currently).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTopClientIpSumInfoList(request *DescribeTopClientIpSumInfoListRequest) (response *DescribeTopClientIpSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeTopClientIpSumInfoListRequest()
    }
    response = NewDescribeTopClientIpSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadStreamNumsRequest() (request *DescribeUploadStreamNumsRequest) {
    request = &DescribeUploadStreamNumsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeUploadStreamNums")
    
    return
}

func NewDescribeUploadStreamNumsResponse() (response *DescribeUploadStreamNumsResponse) {
    response = &DescribeUploadStreamNumsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUploadStreamNums
// This API is used to query the number of LVB upstream channels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUploadStreamNums(request *DescribeUploadStreamNumsRequest) (response *DescribeUploadStreamNumsResponse, err error) {
    if request == nil {
        request = NewDescribeUploadStreamNumsRequest()
    }
    response = NewDescribeUploadStreamNumsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVisitTopSumInfoListRequest() (request *DescribeVisitTopSumInfoListRequest) {
    request = &DescribeVisitTopSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeVisitTopSumInfoList")
    
    return
}

func NewDescribeVisitTopSumInfoListResponse() (response *DescribeVisitTopSumInfoListResponse) {
    response = &DescribeVisitTopSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVisitTopSumInfoList
// This API is used to query the information of the top n domain names or stream IDs in a certain period of time (top 1,000 is supported currently).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVisitTopSumInfoList(request *DescribeVisitTopSumInfoListRequest) (response *DescribeVisitTopSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeVisitTopSumInfoListRequest()
    }
    response = NewDescribeVisitTopSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDropLiveStreamRequest() (request *DropLiveStreamRequest) {
    request = &DropLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DropLiveStream")
    
    return
}

func NewDropLiveStreamResponse() (response *DropLiveStreamResponse) {
    response = &DropLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DropLiveStream
// This API is used to disconnect the push connection, which can be resumed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_STREAMNOTALIVE = "ResourceNotFound.StreamNotAlive"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DropLiveStream(request *DropLiveStreamRequest) (response *DropLiveStreamResponse, err error) {
    if request == nil {
        request = NewDropLiveStreamRequest()
    }
    response = NewDropLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewEnableLiveDomainRequest() (request *EnableLiveDomainRequest) {
    request = &EnableLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "EnableLiveDomain")
    
    return
}

func NewEnableLiveDomainResponse() (response *EnableLiveDomainResponse) {
    response = &EnableLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableLiveDomain
// This API is used to enable a disabled LVB domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
func (c *Client) EnableLiveDomain(request *EnableLiveDomainRequest) (response *EnableLiveDomainResponse, err error) {
    if request == nil {
        request = NewEnableLiveDomainRequest()
    }
    response = NewEnableLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveDomainRequest() (request *ForbidLiveDomainRequest) {
    request = &ForbidLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ForbidLiveDomain")
    
    return
}

func NewForbidLiveDomainResponse() (response *ForbidLiveDomainResponse) {
    response = &ForbidLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ForbidLiveDomain
// This API is used to disable an LVB domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ForbidLiveDomain(request *ForbidLiveDomainRequest) (response *ForbidLiveDomainResponse, err error) {
    if request == nil {
        request = NewForbidLiveDomainRequest()
    }
    response = NewForbidLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveStreamRequest() (request *ForbidLiveStreamRequest) {
    request = &ForbidLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ForbidLiveStream")
    
    return
}

func NewForbidLiveStreamResponse() (response *ForbidLiveStreamResponse) {
    response = &ForbidLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ForbidLiveStream
// This API is used to forbid the push of a specific stream. You can preset a time point to resume the stream.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    if request == nil {
        request = NewForbidLiveStreamRequest()
    }
    response = NewForbidLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveCallbackTemplateRequest() (request *ModifyLiveCallbackTemplateRequest) {
    request = &ModifyLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveCallbackTemplate")
    
    return
}

func NewModifyLiveCallbackTemplateResponse() (response *ModifyLiveCallbackTemplateResponse) {
    response = &ModifyLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveCallbackTemplate
// This API is used to modify a callback template.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveCallbackTemplate(request *ModifyLiveCallbackTemplateRequest) (response *ModifyLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveCallbackTemplateRequest()
    }
    response = NewModifyLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveCertRequest() (request *ModifyLiveCertRequest) {
    request = &ModifyLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveCert")
    
    return
}

func NewModifyLiveCertResponse() (response *ModifyLiveCertResponse) {
    response = &ModifyLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveCert
// This API is used to modify a certificate.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
func (c *Client) ModifyLiveCert(request *ModifyLiveCertRequest) (response *ModifyLiveCertResponse, err error) {
    if request == nil {
        request = NewModifyLiveCertRequest()
    }
    response = NewModifyLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainCertRequest() (request *ModifyLiveDomainCertRequest) {
    request = &ModifyLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainCert")
    
    return
}

func NewModifyLiveDomainCertResponse() (response *ModifyLiveDomainCertResponse) {
    response = &ModifyLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveDomainCert
// This API is used to modify the domain name and certificate binding information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) ModifyLiveDomainCert(request *ModifyLiveDomainCertRequest) (response *ModifyLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainCertRequest()
    }
    response = NewModifyLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainRefererRequest() (request *ModifyLiveDomainRefererRequest) {
    request = &ModifyLiveDomainRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainReferer")
    
    return
}

func NewModifyLiveDomainRefererResponse() (response *ModifyLiveDomainRefererResponse) {
    response = &ModifyLiveDomainRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveDomainReferer
// This API is used to configure referer allowlist/blocklist of a live streaming domain name.
//
// Referer information is included in HTTP requests. After you enable referer configuration, live streams using RTMP or WebRTC for playback will not authenticate the referer and can be played back normally. To make the referer configuration effective, the HTTP-FLV or HLS protocol is recommended for playback.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) ModifyLiveDomainReferer(request *ModifyLiveDomainRefererRequest) (response *ModifyLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainRefererRequest()
    }
    response = NewModifyLiveDomainRefererResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePlayAuthKeyRequest() (request *ModifyLivePlayAuthKeyRequest) {
    request = &ModifyLivePlayAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePlayAuthKey")
    
    return
}

func NewModifyLivePlayAuthKeyResponse() (response *ModifyLivePlayAuthKeyResponse) {
    response = &ModifyLivePlayAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLivePlayAuthKey
// This API is used to modify the playback authentication key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
func (c *Client) ModifyLivePlayAuthKey(request *ModifyLivePlayAuthKeyRequest) (response *ModifyLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayAuthKeyRequest()
    }
    response = NewModifyLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePlayDomainRequest() (request *ModifyLivePlayDomainRequest) {
    request = &ModifyLivePlayDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePlayDomain")
    
    return
}

func NewModifyLivePlayDomainResponse() (response *ModifyLivePlayDomainResponse) {
    response = &ModifyLivePlayDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLivePlayDomain
// This API is used to modify a playback domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) ModifyLivePlayDomain(request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayDomainRequest()
    }
    response = NewModifyLivePlayDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePushAuthKeyRequest() (request *ModifyLivePushAuthKeyRequest) {
    request = &ModifyLivePushAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePushAuthKey")
    
    return
}

func NewModifyLivePushAuthKeyResponse() (response *ModifyLivePushAuthKeyResponse) {
    response = &ModifyLivePushAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLivePushAuthKey
// This API is used to modify the LVB push authentication key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
func (c *Client) ModifyLivePushAuthKey(request *ModifyLivePushAuthKeyRequest) (response *ModifyLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePushAuthKeyRequest()
    }
    response = NewModifyLivePushAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveRecordTemplateRequest() (request *ModifyLiveRecordTemplateRequest) {
    request = &ModifyLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveRecordTemplate")
    
    return
}

func NewModifyLiveRecordTemplateResponse() (response *ModifyLiveRecordTemplateResponse) {
    response = &ModifyLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveRecordTemplate
// This API is used to modify the recording template configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLiveRecordTemplate(request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordTemplateRequest()
    }
    response = NewModifyLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveSnapshotTemplateRequest() (request *ModifyLiveSnapshotTemplateRequest) {
    request = &ModifyLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveSnapshotTemplate")
    
    return
}

func NewModifyLiveSnapshotTemplateResponse() (response *ModifyLiveSnapshotTemplateResponse) {
    response = &ModifyLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveSnapshotTemplate
// This API is used to modify the screencapturing template configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLiveSnapshotTemplate(request *ModifyLiveSnapshotTemplateRequest) (response *ModifyLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveSnapshotTemplateRequest()
    }
    response = NewModifyLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveTranscodeTemplateRequest() (request *ModifyLiveTranscodeTemplateRequest) {
    request = &ModifyLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveTranscodeTemplate")
    
    return
}

func NewModifyLiveTranscodeTemplateResponse() (response *ModifyLiveTranscodeTemplateResponse) {
    response = &ModifyLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveTranscodeTemplate
// This API is used to modify the transcoding template configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLiveTranscodeTemplate(request *ModifyLiveTranscodeTemplateRequest) (response *ModifyLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveTranscodeTemplateRequest()
    }
    response = NewModifyLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewResumeDelayLiveStreamRequest() (request *ResumeDelayLiveStreamRequest) {
    request = &ResumeDelayLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ResumeDelayLiveStream")
    
    return
}

func NewResumeDelayLiveStreamResponse() (response *ResumeDelayLiveStreamResponse) {
    response = &ResumeDelayLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeDelayLiveStream
// This API is used to cancel the delay setting and recover the real-time display of the live streaming image.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeDelayLiveStream(request *ResumeDelayLiveStreamRequest) (response *ResumeDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeDelayLiveStreamRequest()
    }
    response = NewResumeDelayLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewResumeLiveStreamRequest() (request *ResumeLiveStreamRequest) {
    request = &ResumeLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ResumeLiveStream")
    
    return
}

func NewResumeLiveStreamResponse() (response *ResumeLiveStreamResponse) {
    response = &ResumeLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeLiveStream
// This API is used to resume the push of a specific stream.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ResumeLiveStream(request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeLiveStreamRequest()
    }
    response = NewResumeLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStopLiveRecordRequest() (request *StopLiveRecordRequest) {
    request = &StopLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "StopLiveRecord")
    
    return
}

func NewStopLiveRecordResponse() (response *StopLiveRecordResponse) {
    response = &StopLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopLiveRecord
// Note: Recording files are stored on the VOD platform. To use the recording feature, you need to activate a VOD account and ensure that it is available. After the recording files are stored, applicable fees (including storage fees and downstream playback traffic fees) will be charged according to the VOD billing method. For more information, please see the corresponding document.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
func (c *Client) StopLiveRecord(request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    if request == nil {
        request = NewStopLiveRecordRequest()
    }
    response = NewStopLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopRecordTaskRequest() (request *StopRecordTaskRequest) {
    request = &StopRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "StopRecordTask")
    
    return
}

func NewStopRecordTaskResponse() (response *StopRecordTaskResponse) {
    response = &StopRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopRecordTask
// This API is used to end a recording prematurely and terminate an ongoing recording task. After the task is successfully terminated, it will not restart.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopRecordTask(request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
    if request == nil {
        request = NewStopRecordTaskRequest()
    }
    response = NewStopRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindLiveDomainCertRequest() (request *UnBindLiveDomainCertRequest) {
    request = &UnBindLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "UnBindLiveDomainCert")
    
    return
}

func NewUnBindLiveDomainCertResponse() (response *UnBindLiveDomainCertResponse) {
    response = &UnBindLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindLiveDomainCert
// This API is used to unbind a domain name certificate.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UnBindLiveDomainCert(request *UnBindLiveDomainCertRequest) (response *UnBindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewUnBindLiveDomainCertRequest()
    }
    response = NewUnBindLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLiveWatermarkRequest() (request *UpdateLiveWatermarkRequest) {
    request = &UpdateLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "UpdateLiveWatermark")
    
    return
}

func NewUpdateLiveWatermarkResponse() (response *UpdateLiveWatermarkResponse) {
    response = &UpdateLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateLiveWatermark
// This API is used to update a watermark.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_WATERMARKEDITERROR = "InternalError.WatermarkEditError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) UpdateLiveWatermark(request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewUpdateLiveWatermarkRequest()
    }
    response = NewUpdateLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}
