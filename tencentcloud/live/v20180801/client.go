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
    "context"
    "errors"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddDelayLiveStream(request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    return c.AddDelayLiveStreamWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddDelayLiveStreamWithContext(ctx context.Context, request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewAddDelayLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDelayLiveStream require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_DOMAINADDED = "FailedOperation.DomainAdded"
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  FAILEDOPERATION_DOMAINNEEDREALNAME = "FailedOperation.DomainNeedRealName"
//  FAILEDOPERATION_DOMAINNEEDVERIFYOWNER = "FailedOperation.DomainNeedVerifyOwner"
//  FAILEDOPERATION_HOSTOUTLIMIT = "FailedOperation.HostOutLimit"
//  FAILEDOPERATION_PARENTDOMAINADDED = "FailedOperation.ParentDomainAdded"
//  FAILEDOPERATION_SUBDOMAINADDED = "FailedOperation.SubDomainAdded"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"
//  INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) AddLiveDomain(request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
    return c.AddLiveDomainWithContext(context.Background(), request)
}

// AddLiveDomain
// This API is used to add a domain name. Only one domain name can be submitted at a time, and it must have an ICP filing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_DOMAINADDED = "FailedOperation.DomainAdded"
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  FAILEDOPERATION_DOMAINNEEDREALNAME = "FailedOperation.DomainNeedRealName"
//  FAILEDOPERATION_DOMAINNEEDVERIFYOWNER = "FailedOperation.DomainNeedVerifyOwner"
//  FAILEDOPERATION_HOSTOUTLIMIT = "FailedOperation.HostOutLimit"
//  FAILEDOPERATION_PARENTDOMAINADDED = "FailedOperation.ParentDomainAdded"
//  FAILEDOPERATION_SUBDOMAINADDED = "FailedOperation.SubDomainAdded"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"
//  INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) AddLiveDomainWithContext(ctx context.Context, request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
    if request == nil {
        request = NewAddLiveDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLiveDomain require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"
func (c *Client) AddLiveWatermark(request *AddLiveWatermarkRequest) (response *AddLiveWatermarkResponse, err error) {
    return c.AddLiveWatermarkWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"
func (c *Client) AddLiveWatermarkWithContext(ctx context.Context, request *AddLiveWatermarkRequest) (response *AddLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewAddLiveWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewAuthenticateDomainOwnerRequest() (request *AuthenticateDomainOwnerRequest) {
    request = &AuthenticateDomainOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "AuthenticateDomainOwner")
    
    
    return
}

func NewAuthenticateDomainOwnerResponse() (response *AuthenticateDomainOwnerResponse) {
    response = &AuthenticateDomainOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AuthenticateDomainOwner
// This API is used to verify the ownership of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
func (c *Client) AuthenticateDomainOwner(request *AuthenticateDomainOwnerRequest) (response *AuthenticateDomainOwnerResponse, err error) {
    return c.AuthenticateDomainOwnerWithContext(context.Background(), request)
}

// AuthenticateDomainOwner
// This API is used to verify the ownership of a domain.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
func (c *Client) AuthenticateDomainOwnerWithContext(ctx context.Context, request *AuthenticateDomainOwnerRequest) (response *AuthenticateDomainOwnerResponse, err error) {
    if request == nil {
        request = NewAuthenticateDomainOwnerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuthenticateDomainOwner require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuthenticateDomainOwnerResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CancelCommonMixStream(request *CancelCommonMixStreamRequest) (response *CancelCommonMixStreamResponse, err error) {
    return c.CancelCommonMixStreamWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CancelCommonMixStreamWithContext(ctx context.Context, request *CancelCommonMixStreamRequest) (response *CancelCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCancelCommonMixStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelCommonMixStream require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateCommonMixStream(request *CreateCommonMixStreamRequest) (response *CreateCommonMixStreamResponse, err error) {
    return c.CreateCommonMixStreamWithContext(context.Background(), request)
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
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateCommonMixStreamWithContext(ctx context.Context, request *CreateCommonMixStreamRequest) (response *CreateCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCreateCommonMixStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCommonMixStream require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackRule(request *CreateLiveCallbackRuleRequest) (response *CreateLiveCallbackRuleResponse, err error) {
    return c.CreateLiveCallbackRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackRuleWithContext(ctx context.Context, request *CreateLiveCallbackRuleRequest) (response *CreateLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveCallbackRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a callback template. Up to 50 templates can be created. After the template ID is returned, you need to call the [CreateLiveCallbackRule](https://intl.cloud.tencent.com/document/product/267/32638?from_cn_redirect=1) API to bind the template ID to a domain name/path.
//
// <br>For information about callback protocols, see [How to Receive Event Notification](https://intl.cloud.tencent.com/document/product/267/32744?from_cn_redirect=1).
//
// Note: You need to specify at least one callback URL.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackTemplate(request *CreateLiveCallbackTemplateRequest) (response *CreateLiveCallbackTemplateResponse, err error) {
    return c.CreateLiveCallbackTemplateWithContext(context.Background(), request)
}

// CreateLiveCallbackTemplate
// This API is used to create a callback template. Up to 50 templates can be created. After the template ID is returned, you need to call the [CreateLiveCallbackRule](https://intl.cloud.tencent.com/document/product/267/32638?from_cn_redirect=1) API to bind the template ID to a domain name/path.
//
// <br>For information about callback protocols, see [How to Receive Event Notification](https://intl.cloud.tencent.com/document/product/267/32744?from_cn_redirect=1).
//
// Note: You need to specify at least one callback URL.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveCallbackTemplateWithContext(ctx context.Context, request *CreateLiveCallbackTemplateRequest) (response *CreateLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLivePullStreamTaskRequest() (request *CreateLivePullStreamTaskRequest) {
    request = &CreateLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLivePullStreamTask")
    
    
    return
}

func NewCreateLivePullStreamTaskResponse() (response *CreateLivePullStreamTaskResponse) {
    response = &CreateLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLivePullStreamTask
// This API is used to create a task to pull streams from video files or an external live streaming source and publish them to a specified destination URL.
//
// Notes:
//
// 1. By default, you can have at most 20 stream pulling tasks at a time. You can submit a ticket to raise the limit.
//
// 2. Only H.264 and H.265 are supported for video. If the source video is in a different format, please transcode it first.
//
// 3. Only AAC is supported for audio. If the source audio is in a different format, please transcode it first.
//
// 4. You can enable auto deletion in the console to delete expired tasks automatically.
//
// 5. The pull and relay feature is a paid feature. For its billing details, see [Relay](https://intl.cloud.tencent.com/document/product/267/53308?from_cn_redirect=1).
//
// 6. CSS is only responsible for pulling and relaying content. Please make sure that your content is authorized and complies with relevant laws and regulations. In case of copyright infringement or violation of laws or regulations, CSS will suspend its service for you and reserves the right to seek legal remedies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDMIXINPUTPARAM = "InvalidParameter.InvalidMixInputParam"
//  INVALIDPARAMETER_INVALIDOUTPUTPARAM = "InvalidParameter.InvalidOutputParam"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TASKNUMMORETHANLIMIT = "InvalidParameter.TaskNumMoreThanLimit"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateLivePullStreamTask(request *CreateLivePullStreamTaskRequest) (response *CreateLivePullStreamTaskResponse, err error) {
    return c.CreateLivePullStreamTaskWithContext(context.Background(), request)
}

// CreateLivePullStreamTask
// This API is used to create a task to pull streams from video files or an external live streaming source and publish them to a specified destination URL.
//
// Notes:
//
// 1. By default, you can have at most 20 stream pulling tasks at a time. You can submit a ticket to raise the limit.
//
// 2. Only H.264 and H.265 are supported for video. If the source video is in a different format, please transcode it first.
//
// 3. Only AAC is supported for audio. If the source audio is in a different format, please transcode it first.
//
// 4. You can enable auto deletion in the console to delete expired tasks automatically.
//
// 5. The pull and relay feature is a paid feature. For its billing details, see [Relay](https://intl.cloud.tencent.com/document/product/267/53308?from_cn_redirect=1).
//
// 6. CSS is only responsible for pulling and relaying content. Please make sure that your content is authorized and complies with relevant laws and regulations. In case of copyright infringement or violation of laws or regulations, CSS will suspend its service for you and reserves the right to seek legal remedies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDMIXINPUTPARAM = "InvalidParameter.InvalidMixInputParam"
//  INVALIDPARAMETER_INVALIDOUTPUTPARAM = "InvalidParameter.InvalidOutputParam"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TASKNUMMORETHANLIMIT = "InvalidParameter.TaskNumMoreThanLimit"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateLivePullStreamTaskWithContext(ctx context.Context, request *CreateLivePullStreamTaskRequest) (response *CreateLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewCreateLivePullStreamTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLivePullStreamTaskResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"
func (c *Client) CreateLiveRecord(request *CreateLiveRecordRequest) (response *CreateLiveRecordResponse, err error) {
    return c.CreateLiveRecordWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"
func (c *Client) CreateLiveRecordWithContext(ctx context.Context, request *CreateLiveRecordRequest) (response *CreateLiveRecordResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecord require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveRecordRule(request *CreateLiveRecordRuleRequest) (response *CreateLiveRecordRuleResponse, err error) {
    return c.CreateLiveRecordRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveRecordRuleWithContext(ctx context.Context, request *CreateLiveRecordRuleRequest) (response *CreateLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecordRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a recording template. You can create up to 50 templates. To use a template, you need to call the [CreateLiveRecordRule](https://intl.cloud.tencent.com/document/product/267/32615?from_cn_redirect=1) API to bind the template ID returned by this API to a stream.
//
// <br>More on recording: [Live Recording](https://intl.cloud.tencent.com/document/product/267/32739?from_cn_redirect=1)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    return c.CreateLiveRecordTemplateWithContext(context.Background(), request)
}

// CreateLiveRecordTemplate
// This API is used to create a recording template. You can create up to 50 templates. To use a template, you need to call the [CreateLiveRecordRule](https://intl.cloud.tencent.com/document/product/267/32615?from_cn_redirect=1) API to bind the template ID returned by this API to a stream.
//
// <br>More on recording: [Live Recording](https://intl.cloud.tencent.com/document/product/267/32739?from_cn_redirect=1)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveRecordTemplateWithContext(ctx context.Context, request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveSnapshotRule(request *CreateLiveSnapshotRuleRequest) (response *CreateLiveSnapshotRuleResponse, err error) {
    return c.CreateLiveSnapshotRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveSnapshotRuleWithContext(ctx context.Context, request *CreateLiveSnapshotRuleRequest) (response *CreateLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveSnapshotRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a screencapture template. After a template ID is returned, you need to call the [CreateLiveSnapshotRule](https://intl.cloud.tencent.com/document/product/267/32625?from_cn_redirect=1) API to bind the template ID to a stream. You can create up to 50 screencapture templates.
//
// <br>To learn more about the live screencapture feature, see [Live Screencapture](https://intl.cloud.tencent.com/document/product/267/32737?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveSnapshotTemplate(request *CreateLiveSnapshotTemplateRequest) (response *CreateLiveSnapshotTemplateResponse, err error) {
    return c.CreateLiveSnapshotTemplateWithContext(context.Background(), request)
}

// CreateLiveSnapshotTemplate
// This API is used to create a screencapture template. After a template ID is returned, you need to call the [CreateLiveSnapshotRule](https://intl.cloud.tencent.com/document/product/267/32625?from_cn_redirect=1) API to bind the template ID to a stream. You can create up to 50 screencapture templates.
//
// <br>To learn more about the live screencapture feature, see [Live Screencapture](https://intl.cloud.tencent.com/document/product/267/32737?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveSnapshotTemplateWithContext(ctx context.Context, request *CreateLiveSnapshotTemplateRequest) (response *CreateLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTimeShiftRuleRequest() (request *CreateLiveTimeShiftRuleRequest) {
    request = &CreateLiveTimeShiftRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTimeShiftRule")
    
    
    return
}

func NewCreateLiveTimeShiftRuleResponse() (response *CreateLiveTimeShiftRuleResponse) {
    response = &CreateLiveTimeShiftRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveTimeShiftRule
// This API is used to create a time shifting rule. You need to first call the [CreateLiveTranscodeTemplate](https://intl.cloud.tencent.com/document/product/267/86169?from_cn_redirect=1) API to create a time shifting template, and then call this API to bind the template ID returned to a stream.
//
// <br>More about time shifting: [Time Shifting](https://intl.cloud.tencent.com/document/product/267/86134?from_cn_redirect=1).
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveTimeShiftRule(request *CreateLiveTimeShiftRuleRequest) (response *CreateLiveTimeShiftRuleResponse, err error) {
    return c.CreateLiveTimeShiftRuleWithContext(context.Background(), request)
}

// CreateLiveTimeShiftRule
// This API is used to create a time shifting rule. You need to first call the [CreateLiveTranscodeTemplate](https://intl.cloud.tencent.com/document/product/267/86169?from_cn_redirect=1) API to create a time shifting template, and then call this API to bind the template ID returned to a stream.
//
// <br>More about time shifting: [Time Shifting](https://intl.cloud.tencent.com/document/product/267/86134?from_cn_redirect=1).
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) CreateLiveTimeShiftRuleWithContext(ctx context.Context, request *CreateLiveTimeShiftRuleRequest) (response *CreateLiveTimeShiftRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveTimeShiftRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTimeShiftRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveTimeShiftRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTimeShiftTemplateRequest() (request *CreateLiveTimeShiftTemplateRequest) {
    request = &CreateLiveTimeShiftTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTimeShiftTemplate")
    
    
    return
}

func NewCreateLiveTimeShiftTemplateResponse() (response *CreateLiveTimeShiftTemplateResponse) {
    response = &CreateLiveTimeShiftTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLiveTimeShiftTemplate
// This API is used to create a time shifting template.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTimeShiftTemplate(request *CreateLiveTimeShiftTemplateRequest) (response *CreateLiveTimeShiftTemplateResponse, err error) {
    return c.CreateLiveTimeShiftTemplateWithContext(context.Background(), request)
}

// CreateLiveTimeShiftTemplate
// This API is used to create a time shifting template.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTimeShiftTemplateWithContext(ctx context.Context, request *CreateLiveTimeShiftTemplateRequest) (response *CreateLiveTimeShiftTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveTimeShiftTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTimeShiftTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLiveTimeShiftTemplateResponse()
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
// This API is used to create a transcoding rule that binds a template ID to a stream. Up to 50 transcoding rules can be created in total. Before you call this API, you need to first call [CreateLiveTranscodeTemplate](https://intl.cloud.tencent.com/document/product/267/32646?from_cn_redirect=1) to get the template ID.
//
// <br>Related document: [Live Remuxing and Transcoding](https://intl.cloud.tencent.com/document/product/267/32736?from_cn_redirect=1).
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeRule(request *CreateLiveTranscodeRuleRequest) (response *CreateLiveTranscodeRuleResponse, err error) {
    return c.CreateLiveTranscodeRuleWithContext(context.Background(), request)
}

// CreateLiveTranscodeRule
// This API is used to create a transcoding rule that binds a template ID to a stream. Up to 50 transcoding rules can be created in total. Before you call this API, you need to first call [CreateLiveTranscodeTemplate](https://intl.cloud.tencent.com/document/product/267/32646?from_cn_redirect=1) to get the template ID.
//
// <br>Related document: [Live Remuxing and Transcoding](https://intl.cloud.tencent.com/document/product/267/32736?from_cn_redirect=1).
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeRuleWithContext(ctx context.Context, request *CreateLiveTranscodeRuleRequest) (response *CreateLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTranscodeRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a transcoding template. Up to 50 transcoding templates can be created in total. To use a template, you need to call [CreateLiveTranscodeRule](https://intl.cloud.tencent.com/document/product/267/32647?from_cn_redirect=1) to bind the template ID returned by this API to a stream.
//
// <br>For more information about transcoding, see [Live Remuxing and Transcoding](https://intl.cloud.tencent.com/document/product/267/32736?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
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
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeTemplate(request *CreateLiveTranscodeTemplateRequest) (response *CreateLiveTranscodeTemplateResponse, err error) {
    return c.CreateLiveTranscodeTemplateWithContext(context.Background(), request)
}

// CreateLiveTranscodeTemplate
// This API is used to create a transcoding template. Up to 50 transcoding templates can be created in total. To use a template, you need to call [CreateLiveTranscodeRule](https://intl.cloud.tencent.com/document/product/267/32647?from_cn_redirect=1) to bind the template ID returned by this API to a stream.
//
// <br>For more information about transcoding, see [Live Remuxing and Transcoding](https://intl.cloud.tencent.com/document/product/267/32736?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
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
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveTranscodeTemplateWithContext(ctx context.Context, request *CreateLiveTranscodeTemplateRequest) (response *CreateLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveWatermarkRule(request *CreateLiveWatermarkRuleRequest) (response *CreateLiveWatermarkRuleResponse, err error) {
    return c.CreateLiveWatermarkRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_CONFNOTFOUND = "InvalidParameter.ConfNotFound"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) CreateLiveWatermarkRuleWithContext(ctx context.Context, request *CreateLiveWatermarkRuleRequest) (response *CreateLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveWatermarkRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLiveWatermarkRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a recording task that starts and ends at specific times and records videos according to a specific recording template.
//
// - Prerequisites
//
// 1. Because recording files are saved in VOD, you must first activate VOD.
//
// 2. Storage and playback traffic fees are charged for storing and playing the videos recorded. For details, see [Purchase Guide](https://intl.cloud.tencent.com/document/product/266/2837).
//
// - Notes
//
// 1. If streaming is interrupted, the current recording will stop and a recording file will be generated. When streaming resumes, as long as it’s before the end time of the task, recording will start again.
//
// 2. Avoid creating recording tasks with overlapping time periods. If there are multiple tasks with overlapping time periods for the same stream, the system will start three recording tasks at most to avoid repeated recording.
//
// 3. Task records are kept for three months by the platform.
//
// 4. Do not use the new [CreateRecordTask](https://intl.cloud.tencent.com/document/product/267/45983?from_cn_redirect=1), [StopRecordTask](https://intl.cloud.tencent.com/document/product/267/45981?from_cn_redirect=1), and [DeleteRecordTask] APIs together with the old `CreateLiveRecord`, `StopLiveRecord`, and `DeleteLiveRecord` APIs.
//
// 5. Do not create recording tasks and push streams at the same time. After creating a recording task, we recommend you wait at least three seconds before pushing streams.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRecordTask(request *CreateRecordTaskRequest) (response *CreateRecordTaskResponse, err error) {
    return c.CreateRecordTaskWithContext(context.Background(), request)
}

// CreateRecordTask
// This API is used to create a recording task that starts and ends at specific times and records videos according to a specific recording template.
//
// - Prerequisites
//
// 1. Because recording files are saved in VOD, you must first activate VOD.
//
// 2. Storage and playback traffic fees are charged for storing and playing the videos recorded. For details, see [Purchase Guide](https://intl.cloud.tencent.com/document/product/266/2837).
//
// - Notes
//
// 1. If streaming is interrupted, the current recording will stop and a recording file will be generated. When streaming resumes, as long as it’s before the end time of the task, recording will start again.
//
// 2. Avoid creating recording tasks with overlapping time periods. If there are multiple tasks with overlapping time periods for the same stream, the system will start three recording tasks at most to avoid repeated recording.
//
// 3. Task records are kept for three months by the platform.
//
// 4. Do not use the new [CreateRecordTask](https://intl.cloud.tencent.com/document/product/267/45983?from_cn_redirect=1), [StopRecordTask](https://intl.cloud.tencent.com/document/product/267/45981?from_cn_redirect=1), and [DeleteRecordTask] APIs together with the old `CreateLiveRecord`, `StopLiveRecord`, and `DeleteLiveRecord` APIs.
//
// 5. Do not create recording tasks and push streams at the same time. After creating a recording task, we recommend you wait at least three seconds before pushing streams.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRecordTaskWithContext(ctx context.Context, request *CreateRecordTaskRequest) (response *CreateRecordTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecordTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScreenshotTaskRequest() (request *CreateScreenshotTaskRequest) {
    request = &CreateScreenshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "CreateScreenshotTask")
    
    
    return
}

func NewCreateScreenshotTaskResponse() (response *CreateScreenshotTaskResponse) {
    response = &CreateScreenshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScreenshotTask
// This API is used to create a screencapturing task that has a specific start and end time and takes screenshots according to the template configured.
//
// - Note
//
// 1. If the stream is interrupted, screencapturing will stop. However, the task will still be valid before the specified end time, and screencapturing will be performed as required after the stream is resumed.
//
// 2. Avoid creating screencapturing tasks with overlapping time periods. The system will execute at most three screencapturing tasks on the same stream at a time.
//
// 3. Task records are only kept for three months.
//
// 4. The new screencapturing APIs (CreateScreenshotTask/StopScreenshotTask/DeleteScreenshotTask) are not compatible with the legacy ones (CreateLiveInstantSnapshot/StopLiveInstantSnapshot). Do not mix them when you call APIs to manage screencapturing tasks.
//
// 5. If you create a screencapturing task and publish the stream at the same time, the task may fail to be executed at the specified time. After creating a screencapturing task, we recommend you wait at least three seconds before publishing the stream.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateScreenshotTask(request *CreateScreenshotTaskRequest) (response *CreateScreenshotTaskResponse, err error) {
    return c.CreateScreenshotTaskWithContext(context.Background(), request)
}

// CreateScreenshotTask
// This API is used to create a screencapturing task that has a specific start and end time and takes screenshots according to the template configured.
//
// - Note
//
// 1. If the stream is interrupted, screencapturing will stop. However, the task will still be valid before the specified end time, and screencapturing will be performed as required after the stream is resumed.
//
// 2. Avoid creating screencapturing tasks with overlapping time periods. The system will execute at most three screencapturing tasks on the same stream at a time.
//
// 3. Task records are only kept for three months.
//
// 4. The new screencapturing APIs (CreateScreenshotTask/StopScreenshotTask/DeleteScreenshotTask) are not compatible with the legacy ones (CreateLiveInstantSnapshot/StopLiveInstantSnapshot). Do not mix them when you call APIs to manage screencapturing tasks.
//
// 5. If you create a screencapturing task and publish the stream at the same time, the task may fail to be executed at the specified time. After creating a screencapturing task, we recommend you wait at least three seconds before publishing the stream.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateScreenshotTaskWithContext(ctx context.Context, request *CreateScreenshotTaskRequest) (response *CreateScreenshotTaskResponse, err error) {
    if request == nil {
        request = NewCreateScreenshotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScreenshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScreenshotTaskResponse()
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DeleteLiveCallbackRule(request *DeleteLiveCallbackRuleRequest) (response *DeleteLiveCallbackRuleResponse, err error) {
    return c.DeleteLiveCallbackRuleWithContext(context.Background(), request)
}

// DeleteLiveCallbackRule
// This API is used to delete a callback rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DeleteLiveCallbackRuleWithContext(ctx context.Context, request *DeleteLiveCallbackRuleRequest) (response *DeleteLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveCallbackRule require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveCallbackTemplate(request *DeleteLiveCallbackTemplateRequest) (response *DeleteLiveCallbackTemplateResponse, err error) {
    return c.DeleteLiveCallbackTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveCallbackTemplateWithContext(ctx context.Context, request *DeleteLiveCallbackTemplateRequest) (response *DeleteLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveCallbackTemplateResponse()
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
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_TAGUNBINDERROR = "FailedOperation.TagUnbindError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveDomain(request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    return c.DeleteLiveDomainWithContext(context.Background(), request)
}

// DeleteLiveDomain
// This API is used to delete an added LVB domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_JIFEINOENOUGHFUND = "FailedOperation.JiFeiNoEnoughFund"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  FAILEDOPERATION_TAGUNBINDERROR = "FailedOperation.TagUnbindError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveDomainWithContext(ctx context.Context, request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    if request == nil {
        request = NewDeleteLiveDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLivePullStreamTaskRequest() (request *DeleteLivePullStreamTaskRequest) {
    request = &DeleteLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLivePullStreamTask")
    
    
    return
}

func NewDeleteLivePullStreamTaskResponse() (response *DeleteLivePullStreamTaskResponse) {
    response = &DeleteLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLivePullStreamTask
// This API is used to delete a task created by `CreateLivePullStreamTask`.
//
// Notes:
//
// 1. For the `TaskId` request parameter, pass in the task ID returned by the `CreateLivePullStreamTask` API.
//
// 2. You can query the ID of a task using the `DescribeLivePullStreamTasks` API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteLivePullStreamTask(request *DeleteLivePullStreamTaskRequest) (response *DeleteLivePullStreamTaskResponse, err error) {
    return c.DeleteLivePullStreamTaskWithContext(context.Background(), request)
}

// DeleteLivePullStreamTask
// This API is used to delete a task created by `CreateLivePullStreamTask`.
//
// Notes:
//
// 1. For the `TaskId` request parameter, pass in the task ID returned by the `CreateLivePullStreamTask` API.
//
// 2. You can query the ID of a task using the `DescribeLivePullStreamTasks` API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteLivePullStreamTaskWithContext(ctx context.Context, request *DeleteLivePullStreamTaskRequest) (response *DeleteLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewDeleteLivePullStreamTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLivePullStreamTaskResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecord(request *DeleteLiveRecordRequest) (response *DeleteLiveRecordResponse, err error) {
    return c.DeleteLiveRecordWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordWithContext(ctx context.Context, request *DeleteLiveRecordRequest) (response *DeleteLiveRecordResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecord require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordRule(request *DeleteLiveRecordRuleRequest) (response *DeleteLiveRecordRuleResponse, err error) {
    return c.DeleteLiveRecordRuleWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordRuleWithContext(ctx context.Context, request *DeleteLiveRecordRuleRequest) (response *DeleteLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecordRule require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    return c.DeleteLiveRecordTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveRecordTemplateWithContext(ctx context.Context, request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotRule(request *DeleteLiveSnapshotRuleRequest) (response *DeleteLiveSnapshotRuleResponse, err error) {
    return c.DeleteLiveSnapshotRuleWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotRuleWithContext(ctx context.Context, request *DeleteLiveSnapshotRuleRequest) (response *DeleteLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveSnapshotRule require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotTemplate(request *DeleteLiveSnapshotTemplateRequest) (response *DeleteLiveSnapshotTemplateResponse, err error) {
    return c.DeleteLiveSnapshotTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveSnapshotTemplateWithContext(ctx context.Context, request *DeleteLiveSnapshotTemplateRequest) (response *DeleteLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTimeShiftRuleRequest() (request *DeleteLiveTimeShiftRuleRequest) {
    request = &DeleteLiveTimeShiftRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTimeShiftRule")
    
    
    return
}

func NewDeleteLiveTimeShiftRuleResponse() (response *DeleteLiveTimeShiftRuleResponse) {
    response = &DeleteLiveTimeShiftRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveTimeShiftRule
// This API is used to delete a time shifting rule.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftRule(request *DeleteLiveTimeShiftRuleRequest) (response *DeleteLiveTimeShiftRuleResponse, err error) {
    return c.DeleteLiveTimeShiftRuleWithContext(context.Background(), request)
}

// DeleteLiveTimeShiftRule
// This API is used to delete a time shifting rule.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftRuleWithContext(ctx context.Context, request *DeleteLiveTimeShiftRuleRequest) (response *DeleteLiveTimeShiftRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTimeShiftRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTimeShiftRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveTimeShiftRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTimeShiftTemplateRequest() (request *DeleteLiveTimeShiftTemplateRequest) {
    request = &DeleteLiveTimeShiftTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTimeShiftTemplate")
    
    
    return
}

func NewDeleteLiveTimeShiftTemplateResponse() (response *DeleteLiveTimeShiftTemplateResponse) {
    response = &DeleteLiveTimeShiftTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLiveTimeShiftTemplate
// This API is used to delete a time shifting template.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftTemplate(request *DeleteLiveTimeShiftTemplateRequest) (response *DeleteLiveTimeShiftTemplateResponse, err error) {
    return c.DeleteLiveTimeShiftTemplateWithContext(context.Background(), request)
}

// DeleteLiveTimeShiftTemplate
// This API is used to delete a time shifting template.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTimeShiftTemplateWithContext(ctx context.Context, request *DeleteLiveTimeShiftTemplateRequest) (response *DeleteLiveTimeShiftTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTimeShiftTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTimeShiftTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLiveTimeShiftTemplateResponse()
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
//  INVALIDPARAMETER_RULENOTFOUND = "InvalidParameter.RuleNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeRule(request *DeleteLiveTranscodeRuleRequest) (response *DeleteLiveTranscodeRuleResponse, err error) {
    return c.DeleteLiveTranscodeRuleWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_RULENOTFOUND = "InvalidParameter.RuleNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeRuleWithContext(ctx context.Context, request *DeleteLiveTranscodeRuleRequest) (response *DeleteLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTranscodeRule require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeTemplate(request *DeleteLiveTranscodeTemplateRequest) (response *DeleteLiveTranscodeTemplateResponse, err error) {
    return c.DeleteLiveTranscodeTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveTranscodeTemplateWithContext(ctx context.Context, request *DeleteLiveTranscodeTemplateRequest) (response *DeleteLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_CONFINUSED = "InvalidParameter.ConfInUsed"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) DeleteLiveWatermark(request *DeleteLiveWatermarkRequest) (response *DeleteLiveWatermarkResponse, err error) {
    return c.DeleteLiveWatermarkWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_CONFINUSED = "InvalidParameter.ConfInUsed"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) DeleteLiveWatermarkWithContext(ctx context.Context, request *DeleteLiveWatermarkRequest) (response *DeleteLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveWatermarkRule(request *DeleteLiveWatermarkRuleRequest) (response *DeleteLiveWatermarkRuleResponse, err error) {
    return c.DeleteLiveWatermarkRuleWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DeleteLiveWatermarkRuleWithContext(ctx context.Context, request *DeleteLiveWatermarkRuleRequest) (response *DeleteLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLiveWatermarkRule require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordTask(request *DeleteRecordTaskRequest) (response *DeleteRecordTaskResponse, err error) {
    return c.DeleteRecordTaskWithContext(context.Background(), request)
}

// DeleteRecordTask
// This API is used to delete a recording task configuration. The deletion does not affect running tasks and takes effect only for new pushes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordTaskWithContext(ctx context.Context, request *DeleteRecordTaskRequest) (response *DeleteRecordTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordTask require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to get the playback data of all streams at a specified time point (accurate to the minute).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeAllStreamPlayInfoList(request *DescribeAllStreamPlayInfoListRequest) (response *DescribeAllStreamPlayInfoListResponse, err error) {
    return c.DescribeAllStreamPlayInfoListWithContext(context.Background(), request)
}

// DescribeAllStreamPlayInfoList
// This API is used to get the playback data of all streams at a specified time point (accurate to the minute).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeAllStreamPlayInfoListWithContext(ctx context.Context, request *DescribeAllStreamPlayInfoListRequest) (response *DescribeAllStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAllStreamPlayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllStreamPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllStreamPlayInfoListResponse()
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
// This API is used to query the data of billable bandwidth and traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeBillBandwidthAndFluxList(request *DescribeBillBandwidthAndFluxListRequest) (response *DescribeBillBandwidthAndFluxListResponse, err error) {
    return c.DescribeBillBandwidthAndFluxListWithContext(context.Background(), request)
}

// DescribeBillBandwidthAndFluxList
// This API is used to query the data of billable bandwidth and traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeBillBandwidthAndFluxListWithContext(ctx context.Context, request *DescribeBillBandwidthAndFluxListRequest) (response *DescribeBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeBillBandwidthAndFluxListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillBandwidthAndFluxList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeConcurrentRecordStreamNum(request *DescribeConcurrentRecordStreamNumRequest) (response *DescribeConcurrentRecordStreamNumResponse, err error) {
    return c.DescribeConcurrentRecordStreamNumWithContext(context.Background(), request)
}

// DescribeConcurrentRecordStreamNum
// This API is used to query the number of concurrent recording channels, which is applicable to LCB and LVB.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeConcurrentRecordStreamNumWithContext(ctx context.Context, request *DescribeConcurrentRecordStreamNumRequest) (response *DescribeConcurrentRecordStreamNumResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentRecordStreamNumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrentRecordStreamNum require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeliverBandwidthList(request *DescribeDeliverBandwidthListRequest) (response *DescribeDeliverBandwidthListResponse, err error) {
    return c.DescribeDeliverBandwidthListWithContext(context.Background(), request)
}

// DescribeDeliverBandwidthList
// This API is used to query the billable bandwidth of live stream relaying in the last 3 months. The query period is up to 31 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeliverBandwidthListWithContext(ctx context.Context, request *DescribeDeliverBandwidthListRequest) (response *DescribeDeliverBandwidthListResponse, err error) {
    if request == nil {
        request = NewDescribeDeliverBandwidthListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeliverBandwidthList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupProIspPlayInfoList(request *DescribeGroupProIspPlayInfoListRequest) (response *DescribeGroupProIspPlayInfoListResponse, err error) {
    return c.DescribeGroupProIspPlayInfoListWithContext(context.Background(), request)
}

// DescribeGroupProIspPlayInfoList
// This API is used to query the downstream playback data by district and ISP.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupProIspPlayInfoListWithContext(ctx context.Context, request *DescribeGroupProIspPlayInfoListRequest) (response *DescribeGroupProIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupProIspPlayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupProIspPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeHttpStatusInfoList(request *DescribeHttpStatusInfoListRequest) (response *DescribeHttpStatusInfoListResponse, err error) {
    return c.DescribeHttpStatusInfoListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeHttpStatusInfoListWithContext(ctx context.Context, request *DescribeHttpStatusInfoListRequest) (response *DescribeHttpStatusInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeHttpStatusInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHttpStatusInfoList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackRules(request *DescribeLiveCallbackRulesRequest) (response *DescribeLiveCallbackRulesResponse, err error) {
    return c.DescribeLiveCallbackRulesWithContext(context.Background(), request)
}

// DescribeLiveCallbackRules
// This API is used to get the callback rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackRulesWithContext(ctx context.Context, request *DescribeLiveCallbackRulesRequest) (response *DescribeLiveCallbackRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCallbackRules require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCallbackTemplate(request *DescribeLiveCallbackTemplateRequest) (response *DescribeLiveCallbackTemplateResponse, err error) {
    return c.DescribeLiveCallbackTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCallbackTemplateWithContext(ctx context.Context, request *DescribeLiveCallbackTemplateRequest) (response *DescribeLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackTemplates(request *DescribeLiveCallbackTemplatesRequest) (response *DescribeLiveCallbackTemplatesResponse, err error) {
    return c.DescribeLiveCallbackTemplatesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveCallbackTemplatesWithContext(ctx context.Context, request *DescribeLiveCallbackTemplatesRequest) (response *DescribeLiveCallbackTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCallbackTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCert(request *DescribeLiveCertRequest) (response *DescribeLiveCertResponse, err error) {
    return c.DescribeLiveCertWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCertWithContext(ctx context.Context, request *DescribeLiveCertRequest) (response *DescribeLiveCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCert require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCerts(request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    return c.DescribeLiveCertsWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveCertsWithContext(ctx context.Context, request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveCerts require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveDelayInfoList(request *DescribeLiveDelayInfoListRequest) (response *DescribeLiveDelayInfoListResponse, err error) {
    return c.DescribeLiveDelayInfoListWithContext(context.Background(), request)
}

// DescribeLiveDelayInfoList
// This API is used to get the list of delayed playbacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveDelayInfoListWithContext(ctx context.Context, request *DescribeLiveDelayInfoListRequest) (response *DescribeLiveDelayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDelayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDelayInfoList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomain(request *DescribeLiveDomainRequest) (response *DescribeLiveDomainResponse, err error) {
    return c.DescribeLiveDomainWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainWithContext(ctx context.Context, request *DescribeLiveDomainRequest) (response *DescribeLiveDomainResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomain require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCert(request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    return c.DescribeLiveDomainCertWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCertWithContext(ctx context.Context, request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainCertBindingsRequest() (request *DescribeLiveDomainCertBindingsRequest) {
    request = &DescribeLiveDomainCertBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainCertBindings")
    
    
    return
}

func NewDescribeLiveDomainCertBindingsResponse() (response *DescribeLiveDomainCertBindingsResponse) {
    response = &DescribeLiveDomainCertBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveDomainCertBindings
// This API is used to query domains bound with certificates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCertBindings(request *DescribeLiveDomainCertBindingsRequest) (response *DescribeLiveDomainCertBindingsResponse, err error) {
    return c.DescribeLiveDomainCertBindingsWithContext(context.Background(), request)
}

// DescribeLiveDomainCertBindings
// This API is used to query domains bound with certificates.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainCertBindingsWithContext(ctx context.Context, request *DescribeLiveDomainCertBindingsRequest) (response *DescribeLiveDomainCertBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertBindingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainCertBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveDomainCertBindingsResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainReferer(request *DescribeLiveDomainRefererRequest) (response *DescribeLiveDomainRefererResponse, err error) {
    return c.DescribeLiveDomainRefererWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainRefererWithContext(ctx context.Context, request *DescribeLiveDomainRefererRequest) (response *DescribeLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRefererRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomainReferer require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomains(request *DescribeLiveDomainsRequest) (response *DescribeLiveDomainsResponse, err error) {
    return c.DescribeLiveDomainsWithContext(context.Background(), request)
}

// DescribeLiveDomains
// This API is used to query domain names by domain name status and type.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  RESOURCENOTFOUND_EMPTYDATA = "ResourceNotFound.EmptyData"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveDomainsWithContext(ctx context.Context, request *DescribeLiveDomainsRequest) (response *DescribeLiveDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveDomains require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveForbidStreamList(request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    return c.DescribeLiveForbidStreamListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveForbidStreamListWithContext(ctx context.Context, request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveForbidStreamListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveForbidStreamList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePlayAuthKey(request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    return c.DescribeLivePlayAuthKeyWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePlayAuthKeyWithContext(ctx context.Context, request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePlayAuthKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePlayAuthKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePullStreamTasksRequest() (request *DescribeLivePullStreamTasksRequest) {
    request = &DescribeLivePullStreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePullStreamTasks")
    
    
    return
}

func NewDescribeLivePullStreamTasksResponse() (response *DescribeLivePullStreamTasksResponse) {
    response = &DescribeLivePullStreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLivePullStreamTasks
// This API is used to query the stream pulling tasks created by `CreateLivePullStreamTask`.
//
// The tasks returned are sorted by last updated time in descending order.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTasks(request *DescribeLivePullStreamTasksRequest) (response *DescribeLivePullStreamTasksResponse, err error) {
    return c.DescribeLivePullStreamTasksWithContext(context.Background(), request)
}

// DescribeLivePullStreamTasks
// This API is used to query the stream pulling tasks created by `CreateLivePullStreamTask`.
//
// The tasks returned are sorted by last updated time in descending order.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTasksWithContext(ctx context.Context, request *DescribeLivePullStreamTasksRequest) (response *DescribeLivePullStreamTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLivePullStreamTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePullStreamTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLivePullStreamTasksResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePushAuthKey(request *DescribeLivePushAuthKeyRequest) (response *DescribeLivePushAuthKeyResponse, err error) {
    return c.DescribeLivePushAuthKeyWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLivePushAuthKeyWithContext(ctx context.Context, request *DescribeLivePushAuthKeyRequest) (response *DescribeLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePushAuthKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLivePushAuthKey require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordRules(request *DescribeLiveRecordRulesRequest) (response *DescribeLiveRecordRulesResponse, err error) {
    return c.DescribeLiveRecordRulesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordRulesWithContext(ctx context.Context, request *DescribeLiveRecordRulesRequest) (response *DescribeLiveRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordRules require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplate(request *DescribeLiveRecordTemplateRequest) (response *DescribeLiveRecordTemplateResponse, err error) {
    return c.DescribeLiveRecordTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplateWithContext(ctx context.Context, request *DescribeLiveRecordTemplateRequest) (response *DescribeLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplates(request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    return c.DescribeLiveRecordTemplatesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveRecordTemplatesWithContext(ctx context.Context, request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveRecordTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveSnapshotRules(request *DescribeLiveSnapshotRulesRequest) (response *DescribeLiveSnapshotRulesResponse, err error) {
    return c.DescribeLiveSnapshotRulesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveSnapshotRulesWithContext(ctx context.Context, request *DescribeLiveSnapshotRulesRequest) (response *DescribeLiveSnapshotRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveSnapshotRules require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplate(request *DescribeLiveSnapshotTemplateRequest) (response *DescribeLiveSnapshotTemplateResponse, err error) {
    return c.DescribeLiveSnapshotTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplateWithContext(ctx context.Context, request *DescribeLiveSnapshotTemplateRequest) (response *DescribeLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplates(request *DescribeLiveSnapshotTemplatesRequest) (response *DescribeLiveSnapshotTemplatesResponse, err error) {
    return c.DescribeLiveSnapshotTemplatesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveSnapshotTemplatesWithContext(ctx context.Context, request *DescribeLiveSnapshotTemplatesRequest) (response *DescribeLiveSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveSnapshotTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamEventList(request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    return c.DescribeLiveStreamEventListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamEventListWithContext(ctx context.Context, request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamEventList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamOnlineList(request *DescribeLiveStreamOnlineListRequest) (response *DescribeLiveStreamOnlineListResponse, err error) {
    return c.DescribeLiveStreamOnlineListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamOnlineListWithContext(ctx context.Context, request *DescribeLiveStreamOnlineListRequest) (response *DescribeLiveStreamOnlineListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamOnlineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamOnlineList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamPublishedList(request *DescribeLiveStreamPublishedListRequest) (response *DescribeLiveStreamPublishedListResponse, err error) {
    return c.DescribeLiveStreamPublishedListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamPublishedListWithContext(ctx context.Context, request *DescribeLiveStreamPublishedListRequest) (response *DescribeLiveStreamPublishedListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPublishedListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamPublishedList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamPushInfoList(request *DescribeLiveStreamPushInfoListRequest) (response *DescribeLiveStreamPushInfoListResponse, err error) {
    return c.DescribeLiveStreamPushInfoListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveStreamPushInfoListWithContext(ctx context.Context, request *DescribeLiveStreamPushInfoListRequest) (response *DescribeLiveStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPushInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamPushInfoList require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to get the stream status, which may be active, inactive, or disabled.
//
// 
//
// Notes:
//
// This API allows you to query the status of a stream in real time. Given external factors such as network jitter, note the following when you determine whether a host is online:
//
// 1. If possible, use your own logic of stream starting/stopping in a room, such as streaming signaling on the client and the online heartbeat of a host, to determine whether the host is online.
//
// 2. If your application does not provide the room management feature, use the following methods to determine whether a host is online:
//
// 2.1 Use the [live stream callback](https://intl.cloud.tencent.com/document/product/267/20388?from_cn_redirect=1).
//
// 2.2 Call [DescribeLiveStreamOnlineList](https://intl.cloud.tencent.com/document/api/267/20472?from_cn_redirect=1) on a regular basis (interval > 1 min).
//
// 2.3 Call this API.
//
// 2.4 A host is considered to be online if the result returned by any of the above methods indicates so. If an API call times out or a parsing error occurs, to minimize the impact on your business, CSS will also consider the host online.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamState(request *DescribeLiveStreamStateRequest) (response *DescribeLiveStreamStateResponse, err error) {
    return c.DescribeLiveStreamStateWithContext(context.Background(), request)
}

// DescribeLiveStreamState
// This API is used to get the stream status, which may be active, inactive, or disabled.
//
// 
//
// Notes:
//
// This API allows you to query the status of a stream in real time. Given external factors such as network jitter, note the following when you determine whether a host is online:
//
// 1. If possible, use your own logic of stream starting/stopping in a room, such as streaming signaling on the client and the online heartbeat of a host, to determine whether the host is online.
//
// 2. If your application does not provide the room management feature, use the following methods to determine whether a host is online:
//
// 2.1 Use the [live stream callback](https://intl.cloud.tencent.com/document/product/267/20388?from_cn_redirect=1).
//
// 2.2 Call [DescribeLiveStreamOnlineList](https://intl.cloud.tencent.com/document/api/267/20472?from_cn_redirect=1) on a regular basis (interval > 1 min).
//
// 2.3 Call this API.
//
// 2.4 A host is considered to be online if the result returned by any of the above methods indicates so. If an API call times out or a parsing error occurs, to minimize the impact on your business, CSS will also consider the host online.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamStateWithContext(ctx context.Context, request *DescribeLiveStreamStateRequest) (response *DescribeLiveStreamStateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveStreamState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveStreamStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftBillInfoListRequest() (request *DescribeLiveTimeShiftBillInfoListRequest) {
    request = &DescribeLiveTimeShiftBillInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftBillInfoList")
    
    
    return
}

func NewDescribeLiveTimeShiftBillInfoListResponse() (response *DescribeLiveTimeShiftBillInfoListResponse) {
    response = &DescribeLiveTimeShiftBillInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftBillInfoList
// This API is used to query the time-shift video length, time-shift days, and total time-shift period of push domains. The data returned is on a five-minute basis. You can use it for reconciliation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTimeShiftBillInfoList(request *DescribeLiveTimeShiftBillInfoListRequest) (response *DescribeLiveTimeShiftBillInfoListResponse, err error) {
    return c.DescribeLiveTimeShiftBillInfoListWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftBillInfoList
// This API is used to query the time-shift video length, time-shift days, and total time-shift period of push domains. The data returned is on a five-minute basis. You can use it for reconciliation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTimeShiftBillInfoListWithContext(ctx context.Context, request *DescribeLiveTimeShiftBillInfoListRequest) (response *DescribeLiveTimeShiftBillInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftBillInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftBillInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftBillInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftRulesRequest() (request *DescribeLiveTimeShiftRulesRequest) {
    request = &DescribeLiveTimeShiftRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftRules")
    
    
    return
}

func NewDescribeLiveTimeShiftRulesResponse() (response *DescribeLiveTimeShiftRulesResponse) {
    response = &DescribeLiveTimeShiftRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftRules
// This API is used to query time shifting rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftRules(request *DescribeLiveTimeShiftRulesRequest) (response *DescribeLiveTimeShiftRulesResponse, err error) {
    return c.DescribeLiveTimeShiftRulesWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftRules
// This API is used to query time shifting rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftRulesWithContext(ctx context.Context, request *DescribeLiveTimeShiftRulesRequest) (response *DescribeLiveTimeShiftRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTimeShiftTemplatesRequest() (request *DescribeLiveTimeShiftTemplatesRequest) {
    request = &DescribeLiveTimeShiftTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTimeShiftTemplates")
    
    
    return
}

func NewDescribeLiveTimeShiftTemplatesResponse() (response *DescribeLiveTimeShiftTemplatesResponse) {
    response = &DescribeLiveTimeShiftTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTimeShiftTemplates
// This API is used to query time shifting templates.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftTemplates(request *DescribeLiveTimeShiftTemplatesRequest) (response *DescribeLiveTimeShiftTemplatesResponse, err error) {
    return c.DescribeLiveTimeShiftTemplatesWithContext(context.Background(), request)
}

// DescribeLiveTimeShiftTemplates
// This API is used to query time shifting templates.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTimeShiftTemplatesWithContext(ctx context.Context, request *DescribeLiveTimeShiftTemplatesRequest) (response *DescribeLiveTimeShiftTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTimeShiftTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTimeShiftTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTimeShiftTemplatesResponse()
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
// This API is used to query the transcoding details of a particular day or a specific time period. Querying may fail if the amount of data queried is too large. In such cases, try shortening the time period.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeDetailInfo(request *DescribeLiveTranscodeDetailInfoRequest) (response *DescribeLiveTranscodeDetailInfoResponse, err error) {
    return c.DescribeLiveTranscodeDetailInfoWithContext(context.Background(), request)
}

// DescribeLiveTranscodeDetailInfo
// This API is used to query the transcoding details of a particular day or a specific time period. Querying may fail if the amount of data queried is too large. In such cases, try shortening the time period.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeDetailInfoWithContext(ctx context.Context, request *DescribeLiveTranscodeDetailInfoRequest) (response *DescribeLiveTranscodeDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeDetailInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeDetailInfo require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeRules(request *DescribeLiveTranscodeRulesRequest) (response *DescribeLiveTranscodeRulesResponse, err error) {
    return c.DescribeLiveTranscodeRulesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveTranscodeRulesWithContext(ctx context.Context, request *DescribeLiveTranscodeRulesRequest) (response *DescribeLiveTranscodeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeRules require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplate(request *DescribeLiveTranscodeTemplateRequest) (response *DescribeLiveTranscodeTemplateResponse, err error) {
    return c.DescribeLiveTranscodeTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplateWithContext(ctx context.Context, request *DescribeLiveTranscodeTemplateRequest) (response *DescribeLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplates(request *DescribeLiveTranscodeTemplatesRequest) (response *DescribeLiveTranscodeTemplatesResponse, err error) {
    return c.DescribeLiveTranscodeTemplatesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) DescribeLiveTranscodeTemplatesWithContext(ctx context.Context, request *DescribeLiveTranscodeTemplatesRequest) (response *DescribeLiveTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTotalInfoRequest() (request *DescribeLiveTranscodeTotalInfoRequest) {
    request = &DescribeLiveTranscodeTotalInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTotalInfo")
    
    
    return
}

func NewDescribeLiveTranscodeTotalInfoResponse() (response *DescribeLiveTranscodeTotalInfoResponse) {
    response = &DescribeLiveTranscodeTotalInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLiveTranscodeTotalInfo
// This API is used to query transcoding usage. You can use it to query data in the past three months.
//
// Notes:
//
// If the start time and end time are on the same day, the data returned will be on a 5-minute basis.
//
// If the start time and end time are not on the same day or if the data of specified domains is queried, the data returned will be on an hourly basis.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveTranscodeTotalInfo(request *DescribeLiveTranscodeTotalInfoRequest) (response *DescribeLiveTranscodeTotalInfoResponse, err error) {
    return c.DescribeLiveTranscodeTotalInfoWithContext(context.Background(), request)
}

// DescribeLiveTranscodeTotalInfo
// This API is used to query transcoding usage. You can use it to query data in the past three months.
//
// Notes:
//
// If the start time and end time are on the same day, the data returned will be on a 5-minute basis.
//
// If the start time and end time are not on the same day or if the data of specified domains is queried, the data returned will be on an hourly basis.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveTranscodeTotalInfoWithContext(ctx context.Context, request *DescribeLiveTranscodeTotalInfoRequest) (response *DescribeLiveTranscodeTotalInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTotalInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveTranscodeTotalInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveTranscodeTotalInfoResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermark(request *DescribeLiveWatermarkRequest) (response *DescribeLiveWatermarkResponse, err error) {
    return c.DescribeLiveWatermarkWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarkWithContext(ctx context.Context, request *DescribeLiveWatermarkRequest) (response *DescribeLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarkRules(request *DescribeLiveWatermarkRulesRequest) (response *DescribeLiveWatermarkRulesResponse, err error) {
    return c.DescribeLiveWatermarkRulesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarkRulesWithContext(ctx context.Context, request *DescribeLiveWatermarkRulesRequest) (response *DescribeLiveWatermarkRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveWatermarkRules require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarks(request *DescribeLiveWatermarksRequest) (response *DescribeLiveWatermarksResponse, err error) {
    return c.DescribeLiveWatermarksWithContext(context.Background(), request)
}

// DescribeLiveWatermarks
// This API is used to query the watermark list.
//
// error code that may be returned:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeLiveWatermarksWithContext(ctx context.Context, request *DescribeLiveWatermarksRequest) (response *DescribeLiveWatermarksResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveWatermarks require credential")
    }

    request.SetContext(ctx)
    
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
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeDetailInfoList(request *DescribePlayErrorCodeDetailInfoListRequest) (response *DescribePlayErrorCodeDetailInfoListResponse, err error) {
    return c.DescribePlayErrorCodeDetailInfoListWithContext(context.Background(), request)
}

// DescribePlayErrorCodeDetailInfoList
// This API is used to query the information of downstream playback error codes, i.e., the occurrences of each HTTP error code (4xx and 5xx) at a 1-minute granularity in a certain period of time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeDetailInfoListWithContext(ctx context.Context, request *DescribePlayErrorCodeDetailInfoListRequest) (response *DescribePlayErrorCodeDetailInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeDetailInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlayErrorCodeDetailInfoList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeSumInfoList(request *DescribePlayErrorCodeSumInfoListRequest) (response *DescribePlayErrorCodeSumInfoListResponse, err error) {
    return c.DescribePlayErrorCodeSumInfoListWithContext(context.Background(), request)
}

// DescribePlayErrorCodeSumInfoList
// This API is used to query the information of downstream playback error codes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribePlayErrorCodeSumInfoListWithContext(ctx context.Context, request *DescribePlayErrorCodeSumInfoListRequest) (response *DescribePlayErrorCodeSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeSumInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlayErrorCodeSumInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePlayErrorCodeSumInfoListResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeProvinceIspPlayInfoList(request *DescribeProvinceIspPlayInfoListRequest) (response *DescribeProvinceIspPlayInfoListResponse, err error) {
    return c.DescribeProvinceIspPlayInfoListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeProvinceIspPlayInfoListWithContext(ctx context.Context, request *DescribeProvinceIspPlayInfoListRequest) (response *DescribeProvinceIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProvinceIspPlayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProvinceIspPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProvinceIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTaskRequest() (request *DescribeRecordTaskRequest) {
    request = &DescribeRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeRecordTask")
    
    
    return
}

func NewDescribeRecordTaskResponse() (response *DescribeRecordTaskResponse) {
    response = &DescribeRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordTask
// This API is used to retrieve a list of recording tasks that were started and ended within a specified time range. 
//
// - Prerequisites: 
//
// 1. This API is only used to query recording tasks created by the CreateRecordTask interface. 
//
// 2. It cannot retrieve recording tasks that have been deleted by the DeleteRecordTask interface or tasks that have expired (platform retains for 3 months).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTask(request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    return c.DescribeRecordTaskWithContext(context.Background(), request)
}

// DescribeRecordTask
// This API is used to retrieve a list of recording tasks that were started and ended within a specified time range. 
//
// - Prerequisites: 
//
// 1. This API is only used to query recording tasks created by the CreateRecordTask interface. 
//
// 2. It cannot retrieve recording tasks that have been deleted by the DeleteRecordTask interface or tasks that have expired (platform retains for 3 months).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTaskWithContext(ctx context.Context, request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordTaskResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeScreenShotSheetNumList(request *DescribeScreenShotSheetNumListRequest) (response *DescribeScreenShotSheetNumListResponse, err error) {
    return c.DescribeScreenShotSheetNumListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeScreenShotSheetNumListWithContext(ctx context.Context, request *DescribeScreenShotSheetNumListRequest) (response *DescribeScreenShotSheetNumListResponse, err error) {
    if request == nil {
        request = NewDescribeScreenShotSheetNumListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScreenShotSheetNumList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamDayPlayInfoList(request *DescribeStreamDayPlayInfoListRequest) (response *DescribeStreamDayPlayInfoListResponse, err error) {
    return c.DescribeStreamDayPlayInfoListWithContext(context.Background(), request)
}

// DescribeStreamDayPlayInfoList
// This API is used to query the playback data of each stream at the day level, including the total traffic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamDayPlayInfoListWithContext(ctx context.Context, request *DescribeStreamDayPlayInfoListRequest) (response *DescribeStreamDayPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamDayPlayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamDayPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    return c.DescribeStreamPlayInfoListWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPlayInfoListWithContext(ctx context.Context, request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPlayInfoList require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to get the push data of a stream, including the audio/video frame rate, bitrate, elapsed time, and codec.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPushInfoList(request *DescribeStreamPushInfoListRequest) (response *DescribeStreamPushInfoListResponse, err error) {
    return c.DescribeStreamPushInfoListWithContext(context.Background(), request)
}

// DescribeStreamPushInfoList
// This API is used to get the push data of a stream, including the audio/video frame rate, bitrate, elapsed time, and codec.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeStreamPushInfoListWithContext(ctx context.Context, request *DescribeStreamPushInfoListRequest) (response *DescribeStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPushInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPushInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimeShiftRecordDetailRequest() (request *DescribeTimeShiftRecordDetailRequest) {
    request = &DescribeTimeShiftRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTimeShiftRecordDetail")
    
    
    return
}

func NewDescribeTimeShiftRecordDetailResponse() (response *DescribeTimeShiftRecordDetailResponse) {
    response = &DescribeTimeShiftRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimeShiftRecordDetail
// This API is used to query the time shifting details of a specific time period (up to 24 hours). You need to call `DescribeTimeShiftStreamList` first to get the request parameters of this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftRecordDetail(request *DescribeTimeShiftRecordDetailRequest) (response *DescribeTimeShiftRecordDetailResponse, err error) {
    return c.DescribeTimeShiftRecordDetailWithContext(context.Background(), request)
}

// DescribeTimeShiftRecordDetail
// This API is used to query the time shifting details of a specific time period (up to 24 hours). You need to call `DescribeTimeShiftStreamList` first to get the request parameters of this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftRecordDetailWithContext(ctx context.Context, request *DescribeTimeShiftRecordDetailRequest) (response *DescribeTimeShiftRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTimeShiftRecordDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimeShiftRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimeShiftRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimeShiftStreamListRequest() (request *DescribeTimeShiftStreamListRequest) {
    request = &DescribeTimeShiftStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTimeShiftStreamList")
    
    
    return
}

func NewDescribeTimeShiftStreamListResponse() (response *DescribeTimeShiftStreamListResponse) {
    response = &DescribeTimeShiftStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimeShiftStreamList
// This API is used to query the time shifted streams in a specific time period (up to 24 hours).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftStreamList(request *DescribeTimeShiftStreamListRequest) (response *DescribeTimeShiftStreamListResponse, err error) {
    return c.DescribeTimeShiftStreamListWithContext(context.Background(), request)
}

// DescribeTimeShiftStreamList
// This API is used to query the time shifted streams in a specific time period (up to 24 hours).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTimeShiftStreamListWithContext(ctx context.Context, request *DescribeTimeShiftStreamListRequest) (response *DescribeTimeShiftStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeTimeShiftStreamListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimeShiftStreamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimeShiftStreamListResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTopClientIpSumInfoList(request *DescribeTopClientIpSumInfoListRequest) (response *DescribeTopClientIpSumInfoListResponse, err error) {
    return c.DescribeTopClientIpSumInfoListWithContext(context.Background(), request)
}

// DescribeTopClientIpSumInfoList
// This API is used to query the information of the top n client IPs in a certain period of time (top 1,000 is supported currently).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTopClientIpSumInfoListWithContext(ctx context.Context, request *DescribeTopClientIpSumInfoListRequest) (response *DescribeTopClientIpSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeTopClientIpSumInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopClientIpSumInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopClientIpSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeTaskNumRequest() (request *DescribeTranscodeTaskNumRequest) {
    request = &DescribeTranscodeTaskNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "DescribeTranscodeTaskNum")
    
    
    return
}

func NewDescribeTranscodeTaskNumResponse() (response *DescribeTranscodeTaskNumResponse) {
    response = &DescribeTranscodeTaskNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeTaskNum
// This API is used to query the number of transcoding tasks.
//
// error code that may be returned:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTranscodeTaskNum(request *DescribeTranscodeTaskNumRequest) (response *DescribeTranscodeTaskNumResponse, err error) {
    return c.DescribeTranscodeTaskNumWithContext(context.Background(), request)
}

// DescribeTranscodeTaskNum
// This API is used to query the number of transcoding tasks.
//
// error code that may be returned:
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeTranscodeTaskNumWithContext(ctx context.Context, request *DescribeTranscodeTaskNumRequest) (response *DescribeTranscodeTaskNumResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeTaskNumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeTaskNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeTaskNumResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeUploadStreamNums(request *DescribeUploadStreamNumsRequest) (response *DescribeUploadStreamNumsResponse, err error) {
    return c.DescribeUploadStreamNumsWithContext(context.Background(), request)
}

// DescribeUploadStreamNums
// This API is used to query the number of LVB upstream channels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeUploadStreamNumsWithContext(ctx context.Context, request *DescribeUploadStreamNumsRequest) (response *DescribeUploadStreamNumsResponse, err error) {
    if request == nil {
        request = NewDescribeUploadStreamNumsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadStreamNums require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeVisitTopSumInfoList(request *DescribeVisitTopSumInfoListRequest) (response *DescribeVisitTopSumInfoListResponse, err error) {
    return c.DescribeVisitTopSumInfoListWithContext(context.Background(), request)
}

// DescribeVisitTopSumInfoList
// This API is used to query the information of the top n domain names or stream IDs in a certain period of time (top 1,000 is supported currently).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) DescribeVisitTopSumInfoListWithContext(ctx context.Context, request *DescribeVisitTopSumInfoListRequest) (response *DescribeVisitTopSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeVisitTopSumInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVisitTopSumInfoList require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to pause a live stream. The stream can be resumed if it is paused.
//
// Note: If you call this API to pause an inactive stream, the request will be considered successful.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DropLiveStream(request *DropLiveStreamRequest) (response *DropLiveStreamResponse, err error) {
    return c.DropLiveStreamWithContext(context.Background(), request)
}

// DropLiveStream
// This API is used to pause a live stream. The stream can be resumed if it is paused.
//
// Note: If you call this API to pause an inactive stream, the request will be considered successful.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DropLiveStreamWithContext(ctx context.Context, request *DropLiveStreamRequest) (response *DropLiveStreamResponse, err error) {
    if request == nil {
        request = NewDropLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropLiveStream require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) EnableLiveDomain(request *EnableLiveDomainRequest) (response *EnableLiveDomainResponse, err error) {
    return c.EnableLiveDomainWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) EnableLiveDomainWithContext(ctx context.Context, request *EnableLiveDomainRequest) (response *EnableLiveDomainResponse, err error) {
    if request == nil {
        request = NewEnableLiveDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableLiveDomain require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ForbidLiveDomain(request *ForbidLiveDomainRequest) (response *ForbidLiveDomainResponse, err error) {
    return c.ForbidLiveDomainWithContext(context.Background(), request)
}

// ForbidLiveDomain
// This API is used to disable an LVB domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ForbidLiveDomainWithContext(ctx context.Context, request *ForbidLiveDomainRequest) (response *ForbidLiveDomainResponse, err error) {
    if request == nil {
        request = NewForbidLiveDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidLiveDomain require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to disable a stream. You can set a time to resume the stream.
//
// Note:
//
// 1. As long as the correct stream name is passed in, the stream will be disabled successfully.
//
// 2. If you want a stream to be disabled only if the push domain, push path, and stream name match, please submit a ticket.
//
// 3. If you have configured domain groups, you must pass in the correct push domain in order to disable a stream.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    return c.ForbidLiveStreamWithContext(context.Background(), request)
}

// ForbidLiveStream
// This API is used to disable a stream. You can set a time to resume the stream.
//
// Note:
//
// 1. As long as the correct stream name is passed in, the stream will be disabled successfully.
//
// 2. If you want a stream to be disabled only if the push domain, push path, and stream name match, please submit a ticket.
//
// 3. If you have configured domain groups, you must pass in the correct push domain in order to disable a stream.
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ForbidLiveStreamWithContext(ctx context.Context, request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    if request == nil {
        request = NewForbidLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForbidLiveStream require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveCallbackTemplate(request *ModifyLiveCallbackTemplateRequest) (response *ModifyLiveCallbackTemplateResponse, err error) {
    return c.ModifyLiveCallbackTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveCallbackTemplateWithContext(ctx context.Context, request *ModifyLiveCallbackTemplateRequest) (response *ModifyLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveCallbackTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveCallbackTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainCertBindingsRequest() (request *ModifyLiveDomainCertBindingsRequest) {
    request = &ModifyLiveDomainCertBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainCertBindings")
    
    
    return
}

func NewModifyLiveDomainCertBindingsResponse() (response *ModifyLiveDomainCertBindingsResponse) {
    response = &ModifyLiveDomainCertBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveDomainCertBindings
// This API is used to bind a certificate to multiple playback domains and update the HTTPS configuration of the domains.
//
// If a self-owned certificate is used, it will be automatically uploaded to Tencent Cloud’s SSL Certificate Service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIGCDNFAILED = "FailedOperation.ConfigCDNFailed"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATENOTFOUND = "InvalidParameter.CrtDateNotFound"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTDOMAINNOTFOUND = "InvalidParameter.CrtDomainNotFound"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
//  INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_CRTDATENOTFOUND = "ResourceNotFound.CrtDateNotFound"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainCertBindings(request *ModifyLiveDomainCertBindingsRequest) (response *ModifyLiveDomainCertBindingsResponse, err error) {
    return c.ModifyLiveDomainCertBindingsWithContext(context.Background(), request)
}

// ModifyLiveDomainCertBindings
// This API is used to bind a certificate to multiple playback domains and update the HTTPS configuration of the domains.
//
// If a self-owned certificate is used, it will be automatically uploaded to Tencent Cloud’s SSL Certificate Service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIGCDNFAILED = "FailedOperation.ConfigCDNFailed"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATENOTFOUND = "InvalidParameter.CrtDateNotFound"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTDOMAINNOTFOUND = "InvalidParameter.CrtDomainNotFound"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
//  INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_CRTDATENOTFOUND = "ResourceNotFound.CrtDateNotFound"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainCertBindingsWithContext(ctx context.Context, request *ModifyLiveDomainCertBindingsRequest) (response *ModifyLiveDomainCertBindingsResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainCertBindingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveDomainCertBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveDomainCertBindingsResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainReferer(request *ModifyLiveDomainRefererRequest) (response *ModifyLiveDomainRefererResponse, err error) {
    return c.ModifyLiveDomainRefererWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveDomainRefererWithContext(ctx context.Context, request *ModifyLiveDomainRefererRequest) (response *ModifyLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainRefererRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveDomainReferer require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayAuthKey(request *ModifyLivePlayAuthKeyRequest) (response *ModifyLivePlayAuthKeyResponse, err error) {
    return c.ModifyLivePlayAuthKeyWithContext(context.Background(), request)
}

// ModifyLivePlayAuthKey
// This API is used to modify the playback authentication key.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayAuthKeyWithContext(ctx context.Context, request *ModifyLivePlayAuthKeyRequest) (response *ModifyLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayAuthKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePlayAuthKey require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayDomain(request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    return c.ModifyLivePlayDomainWithContext(context.Background(), request)
}

// ModifyLivePlayDomain
// This API is used to modify a playback domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePlayDomainWithContext(ctx context.Context, request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePlayDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePlayDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePullStreamTaskRequest() (request *ModifyLivePullStreamTaskRequest) {
    request = &ModifyLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePullStreamTask")
    
    
    return
}

func NewModifyLivePullStreamTaskResponse() (response *ModifyLivePullStreamTaskResponse) {
    response = &ModifyLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLivePullStreamTask
// This API is used to modify a stream pulling task. 
//
// 1. You cannot modify the destination URL. To publish to a new destination, please create a new task.
//
// 2. You cannot modify the source type. To use a different source type, please create a new task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyLivePullStreamTask(request *ModifyLivePullStreamTaskRequest) (response *ModifyLivePullStreamTaskResponse, err error) {
    return c.ModifyLivePullStreamTaskWithContext(context.Background(), request)
}

// ModifyLivePullStreamTask
// This API is used to modify a stream pulling task. 
//
// 1. You cannot modify the destination URL. To publish to a new destination, please create a new task.
//
// 2. You cannot modify the source type. To use a different source type, please create a new task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDBACKUPTOURL = "InvalidParameter.InvalidBackupToUrl"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyLivePullStreamTaskWithContext(ctx context.Context, request *ModifyLivePullStreamTaskRequest) (response *ModifyLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewModifyLivePullStreamTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLivePullStreamTaskResponse()
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePushAuthKey(request *ModifyLivePushAuthKeyRequest) (response *ModifyLivePushAuthKeyResponse, err error) {
    return c.ModifyLivePushAuthKeyWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLivePushAuthKeyWithContext(ctx context.Context, request *ModifyLivePushAuthKeyRequest) (response *ModifyLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePushAuthKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLivePushAuthKey require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveRecordTemplate(request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    return c.ModifyLiveRecordTemplateWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveRecordTemplateWithContext(ctx context.Context, request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveSnapshotTemplate(request *ModifyLiveSnapshotTemplateRequest) (response *ModifyLiveSnapshotTemplateResponse, err error) {
    return c.ModifyLiveSnapshotTemplateWithContext(context.Background(), request)
}

// ModifyLiveSnapshotTemplate
// This API is used to modify the screencapturing template configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"
//  FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"
//  FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveSnapshotTemplateWithContext(ctx context.Context, request *ModifyLiveSnapshotTemplateRequest) (response *ModifyLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveSnapshotTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveTimeShiftTemplateRequest() (request *ModifyLiveTimeShiftTemplateRequest) {
    request = &ModifyLiveTimeShiftTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveTimeShiftTemplate")
    
    
    return
}

func NewModifyLiveTimeShiftTemplateResponse() (response *ModifyLiveTimeShiftTemplateResponse) {
    response = &ModifyLiveTimeShiftTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLiveTimeShiftTemplate
// This API is used to modify a time shifting template.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveTimeShiftTemplate(request *ModifyLiveTimeShiftTemplateRequest) (response *ModifyLiveTimeShiftTemplateResponse, err error) {
    return c.ModifyLiveTimeShiftTemplateWithContext(context.Background(), request)
}

// ModifyLiveTimeShiftTemplate
// This API is used to modify a time shifting template.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
func (c *Client) ModifyLiveTimeShiftTemplateWithContext(ctx context.Context, request *ModifyLiveTimeShiftTemplateRequest) (response *ModifyLiveTimeShiftTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveTimeShiftTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveTimeShiftTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveTimeShiftTemplateResponse()
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
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveTranscodeTemplate(request *ModifyLiveTranscodeTemplateRequest) (response *ModifyLiveTranscodeTemplateResponse, err error) {
    return c.ModifyLiveTranscodeTemplateWithContext(context.Background(), request)
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
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"
//  INVALIDPARAMETER_PROCESSORALREADYEXIST = "InvalidParameter.ProcessorAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) ModifyLiveTranscodeTemplateWithContext(ctx context.Context, request *ModifyLiveTranscodeTemplateRequest) (response *ModifyLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveTranscodeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLiveTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRestartLivePullStreamTaskRequest() (request *RestartLivePullStreamTaskRequest) {
    request = &RestartLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("live", APIVersion, "RestartLivePullStreamTask")
    
    
    return
}

func NewRestartLivePullStreamTaskResponse() (response *RestartLivePullStreamTaskResponse) {
    response = &RestartLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartLivePullStreamTask
// Restart the running live streaming pull task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) RestartLivePullStreamTask(request *RestartLivePullStreamTaskRequest) (response *RestartLivePullStreamTaskResponse, err error) {
    return c.RestartLivePullStreamTaskWithContext(context.Background(), request)
}

// RestartLivePullStreamTask
// Restart the running live streaming pull task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) RestartLivePullStreamTaskWithContext(ctx context.Context, request *RestartLivePullStreamTaskRequest) (response *RestartLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewRestartLivePullStreamTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartLivePullStreamTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartLivePullStreamTaskResponse()
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
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeDelayLiveStream(request *ResumeDelayLiveStreamRequest) (response *ResumeDelayLiveStreamResponse, err error) {
    return c.ResumeDelayLiveStreamWithContext(context.Background(), request)
}

// ResumeDelayLiveStream
// This API is used to cancel the delay setting and recover the real-time display of the live streaming image.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
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
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeDelayLiveStreamWithContext(ctx context.Context, request *ResumeDelayLiveStreamRequest) (response *ResumeDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeDelayLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeDelayLiveStream require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ResumeLiveStream(request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    return c.ResumeLiveStreamWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ResumeLiveStreamWithContext(ctx context.Context, request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeLiveStream require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StopLiveRecord(request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    return c.StopLiveRecordWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) StopLiveRecordWithContext(ctx context.Context, request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    if request == nil {
        request = NewStopLiveRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLiveRecord require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to terminate an ongoing recording task and generate a recording file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopRecordTask(request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
    return c.StopRecordTaskWithContext(context.Background(), request)
}

// StopRecordTask
// This API is used to terminate an ongoing recording task and generate a recording file.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopRecordTaskWithContext(ctx context.Context, request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
    if request == nil {
        request = NewStopRecordTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRecordTask require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) UnBindLiveDomainCert(request *UnBindLiveDomainCertRequest) (response *UnBindLiveDomainCertResponse, err error) {
    return c.UnBindLiveDomainCertWithContext(context.Background(), request)
}

// UnBindLiveDomainCert
// This API is used to unbind a domain name certificate.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
func (c *Client) UnBindLiveDomainCertWithContext(ctx context.Context, request *UnBindLiveDomainCertRequest) (response *UnBindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewUnBindLiveDomainCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindLiveDomainCert require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) UpdateLiveWatermark(request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    return c.UpdateLiveWatermarkWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"
//  RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
//  RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) UpdateLiveWatermarkWithContext(ctx context.Context, request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewUpdateLiveWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLiveWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}
