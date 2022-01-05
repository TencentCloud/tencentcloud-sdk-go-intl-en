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


func NewCreatePictureRequest() (request *CreatePictureRequest) {
    request = &CreatePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreatePicture")
    
    
    return
}

func NewCreatePictureResponse() (response *CreatePictureResponse) {
    response = &CreatePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePicture
// This API is used to add custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to add such images frequently, we recommend you add them in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"
//  INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePicture(request *CreatePictureRequest) (response *CreatePictureResponse, err error) {
    if request == nil {
        request = NewCreatePictureRequest()
    }
    
    response = NewCreatePictureResponse()
    err = c.Send(request, response)
    return
}

// CreatePicture
// This API is used to add custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to add such images frequently, we recommend you add them in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"
//  INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePictureWithContext(ctx context.Context, request *CreatePictureRequest) (response *CreatePictureResponse, err error) {
    if request == nil {
        request = NewCreatePictureRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePictureResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTroubleInfoRequest() (request *CreateTroubleInfoRequest) {
    request = &CreateTroubleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreateTroubleInfo")
    
    
    return
}

func NewCreateTroubleInfoResponse() (response *CreateTroubleInfoResponse) {
    response = &CreateTroubleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTroubleInfo
// This API is used to create exception information.
//
// error code that may be returned:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDFAIL = "InternalError.BackendFail"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateTroubleInfo(request *CreateTroubleInfoRequest) (response *CreateTroubleInfoResponse, err error) {
    if request == nil {
        request = NewCreateTroubleInfoRequest()
    }
    
    response = NewCreateTroubleInfoResponse()
    err = c.Send(request, response)
    return
}

// CreateTroubleInfo
// This API is used to create exception information.
//
// error code that may be returned:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDFAIL = "InternalError.BackendFail"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateTroubleInfoWithContext(ctx context.Context, request *CreateTroubleInfoRequest) (response *CreateTroubleInfoResponse, err error) {
    if request == nil {
        request = NewCreateTroubleInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTroubleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePictureRequest() (request *DeletePictureRequest) {
    request = &DeletePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DeletePicture")
    
    
    return
}

func NewDeletePictureResponse() (response *DeletePictureResponse) {
    response = &DeletePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePicture
// This API is used to delete custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to delete such images frequently, we recommend you delete them in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeletePicture(request *DeletePictureRequest) (response *DeletePictureResponse, err error) {
    if request == nil {
        request = NewDeletePictureRequest()
    }
    
    response = NewDeletePictureResponse()
    err = c.Send(request, response)
    return
}

// DeletePicture
// This API is used to delete custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to delete such images frequently, we recommend you delete them in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeletePictureWithContext(ctx context.Context, request *DeletePictureRequest) (response *DeletePictureResponse, err error) {
    if request == nil {
        request = NewDeletePictureRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalEventRequest() (request *DescribeAbnormalEventRequest) {
    request = &DescribeAbnormalEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAbnormalEvent")
    
    
    return
}

func NewDescribeAbnormalEventResponse() (response *DescribeAbnormalEventResponse) {
    response = &DescribeAbnormalEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalEvent
// This API is used to query exception occurrences under a specified `SDKAppID` and return the exception IDs and possible causes. It queries data in last 15 days, and the query period is up to 1 hour, which can start and end on different days. For more information about exceptions, please see the exception event ID mapping table: https://intl.cloud.tencent.com/document/product/647/37906.
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
func (c *Client) DescribeAbnormalEvent(request *DescribeAbnormalEventRequest) (response *DescribeAbnormalEventResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalEventRequest()
    }
    
    response = NewDescribeAbnormalEventResponse()
    err = c.Send(request, response)
    return
}

// DescribeAbnormalEvent
// This API is used to query exception occurrences under a specified `SDKAppID` and return the exception IDs and possible causes. It queries data in last 15 days, and the query period is up to 1 hour, which can start and end on different days. For more information about exceptions, please see the exception event ID mapping table: https://intl.cloud.tencent.com/document/product/647/37906.
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
func (c *Client) DescribeAbnormalEventWithContext(ctx context.Context, request *DescribeAbnormalEventRequest) (response *DescribeAbnormalEventResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalEventRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAbnormalEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallDetailRequest() (request *DescribeCallDetailRequest) {
    request = &DescribeCallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCallDetail")
    
    
    return
}

func NewDescribeCallDetailResponse() (response *DescribeCallDetailResponse) {
    response = &DescribeCallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCallDetail
// This API is used to query the user list and call quality data of a specified time range in the last 14 days. When `DataType` is not null, data of up to 6 users during a period of up to 1 hour can be queried each time, and the period can start on one day and end on the next. When `DataType` and `UserIds` are null, 6 users are queried by default, and data of up to 100 users can be displayed on each page (`PageSize` set to 100 or smaller). This API is used to query call quality and is not recommended for billing.
//
// **Note**: You can use this API to query or check historical data, but not for real-time key business logic.
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
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetail(request *DescribeCallDetailRequest) (response *DescribeCallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailRequest()
    }
    
    response = NewDescribeCallDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeCallDetail
// This API is used to query the user list and call quality data of a specified time range in the last 14 days. When `DataType` is not null, data of up to 6 users during a period of up to 1 hour can be queried each time, and the period can start on one day and end on the next. When `DataType` and `UserIds` are null, 6 users are queried by default, and data of up to 100 users can be displayed on each page (`PageSize` set to 100 or smaller). This API is used to query call quality and is not recommended for billing.
//
// **Note**: You can use this API to query or check historical data, but not for real-time key business logic.
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
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetailWithContext(ctx context.Context, request *DescribeCallDetailRequest) (response *DescribeCallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetailEventRequest() (request *DescribeDetailEventRequest) {
    request = &DescribeDetailEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeDetailEvent")
    
    
    return
}

func NewDescribeDetailEventResponse() (response *DescribeDetailEventResponse) {
    response = &DescribeDetailEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDetailEvent
// This API is used to query a user’s activity details such as room entry/exit and video enablement/disablement during a call. It can query data for the last 14 days.
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
func (c *Client) DescribeDetailEvent(request *DescribeDetailEventRequest) (response *DescribeDetailEventResponse, err error) {
    if request == nil {
        request = NewDescribeDetailEventRequest()
    }
    
    response = NewDescribeDetailEventResponse()
    err = c.Send(request, response)
    return
}

// DescribeDetailEvent
// This API is used to query a user’s activity details such as room entry/exit and video enablement/disablement during a call. It can query data for the last 14 days.
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
func (c *Client) DescribeDetailEventWithContext(ctx context.Context, request *DescribeDetailEventRequest) (response *DescribeDetailEventResponse, err error) {
    if request == nil {
        request = NewDescribeDetailEventRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDetailEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryScaleRequest() (request *DescribeHistoryScaleRequest) {
    request = &DescribeHistoryScaleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeHistoryScale")
    
    
    return
}

func NewDescribeHistoryScaleResponse() (response *DescribeHistoryScaleResponse) {
    response = &DescribeHistoryScaleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHistoryScale
// This API is used to query the daily numbers of rooms and users under a specified `SDKAppID`. It can query data once per minute for the last 14 days. If a day has not ended, the numbers of rooms and users on the day cannot be queried. 
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
func (c *Client) DescribeHistoryScale(request *DescribeHistoryScaleRequest) (response *DescribeHistoryScaleResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryScaleRequest()
    }
    
    response = NewDescribeHistoryScaleResponse()
    err = c.Send(request, response)
    return
}

// DescribeHistoryScale
// This API is used to query the daily numbers of rooms and users under a specified `SDKAppID`. It can query data once per minute for the last 14 days. If a day has not ended, the numbers of rooms and users on the day cannot be queried. 
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
func (c *Client) DescribeHistoryScaleWithContext(ctx context.Context, request *DescribeHistoryScaleRequest) (response *DescribeHistoryScaleResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryScaleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeHistoryScaleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePictureRequest() (request *DescribePictureRequest) {
    request = &DescribePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribePicture")
    
    
    return
}

func NewDescribePictureResponse() (response *DescribePictureResponse) {
    response = &DescribePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePicture
// This API is used to query the information of custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to query such information frequently, we recommend you query it in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePicture(request *DescribePictureRequest) (response *DescribePictureResponse, err error) {
    if request == nil {
        request = NewDescribePictureRequest()
    }
    
    response = NewDescribePictureResponse()
    err = c.Send(request, response)
    return
}

// DescribePicture
// This API is used to query the information of custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to query such information frequently, we recommend you query it in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePictureWithContext(ctx context.Context, request *DescribePictureRequest) (response *DescribePictureResponse, err error) {
    if request == nil {
        request = NewDescribePictureRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordStatisticRequest() (request *DescribeRecordStatisticRequest) {
    request = &DescribeRecordStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRecordStatistic")
    
    
    return
}

func NewDescribeRecordStatisticResponse() (response *DescribeRecordStatisticResponse) {
    response = &DescribeRecordStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordStatistic
// This API is used to query billable on-cloud recording durations.
//
// 
//
// - If the period queried is 1 day or shorter, the statistics returned are on a 5-minute basis. If the period queried is longer than 1 day, the statistics returned are on a daily basis.
//
// - The period queried in a request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - In the daily pay-as-you-go mode, bills for a day are generated the next morning. Therefore, we recommend you query the statistics after 9 AM the next day.
//
// error code that may be returned:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPID = "InvalidParameter.AppId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordStatistic(request *DescribeRecordStatisticRequest) (response *DescribeRecordStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStatisticRequest()
    }
    
    response = NewDescribeRecordStatisticResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordStatistic
// This API is used to query billable on-cloud recording durations.
//
// 
//
// - If the period queried is 1 day or shorter, the statistics returned are on a 5-minute basis. If the period queried is longer than 1 day, the statistics returned are on a daily basis.
//
// - The period queried in a request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - In the daily pay-as-you-go mode, bills for a day are generated the next morning. Therefore, we recommend you query the statistics after 9 AM the next day.
//
// error code that may be returned:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPID = "InvalidParameter.AppId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordStatisticWithContext(ctx context.Context, request *DescribeRecordStatisticRequest) (response *DescribeRecordStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStatisticRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInformationRequest() (request *DescribeRoomInformationRequest) {
    request = &DescribeRoomInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRoomInformation")
    
    
    return
}

func NewDescribeRoomInformationResponse() (response *DescribeRoomInformationResponse) {
    response = &DescribeRoomInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoomInformation
// This API is used to query the room list of an `SDKAppID` in the last 14 days. It returns 10 calls by default and can return up to 100 calls per query.
//
// **Note**: You can use this API to query or check historical data, but not for real-time key business logic.
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
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInformation(request *DescribeRoomInformationRequest) (response *DescribeRoomInformationResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInformationRequest()
    }
    
    response = NewDescribeRoomInformationResponse()
    err = c.Send(request, response)
    return
}

// DescribeRoomInformation
// This API is used to query the room list of an `SDKAppID` in the last 14 days. It returns 10 calls by default and can return up to 100 calls per query.
//
// **Note**: You can use this API to query or check historical data, but not for real-time key business logic.
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
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInformationWithContext(ctx context.Context, request *DescribeRoomInformationRequest) (response *DescribeRoomInformationResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRoomInformationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcInteractiveTimeRequest() (request *DescribeTrtcInteractiveTimeRequest) {
    request = &DescribeTrtcInteractiveTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcInteractiveTime")
    
    
    return
}

func NewDescribeTrtcInteractiveTimeResponse() (response *DescribeTrtcInteractiveTimeResponse) {
    response = &DescribeTrtcInteractiveTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrtcInteractiveTime
// This API is used to query billable audio/video interaction durations.
//
// - If the period queried is 1 day or shorter, the statistics returned are on a 5-minute basis. If the period queried is longer than 1 day, the statistics returned are on a daily basis.
//
// - The period queried in a request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - In the daily pay-as-you-go mode, bills for a day are generated the next morning. Therefore, we recommend you query the statistics after 9 AM the next day.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcInteractiveTime(request *DescribeTrtcInteractiveTimeRequest) (response *DescribeTrtcInteractiveTimeResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcInteractiveTimeRequest()
    }
    
    response = NewDescribeTrtcInteractiveTimeResponse()
    err = c.Send(request, response)
    return
}

// DescribeTrtcInteractiveTime
// This API is used to query billable audio/video interaction durations.
//
// - If the period queried is 1 day or shorter, the statistics returned are on a 5-minute basis. If the period queried is longer than 1 day, the statistics returned are on a daily basis.
//
// - The period queried in a request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - In the daily pay-as-you-go mode, bills for a day are generated the next morning. Therefore, we recommend you query the statistics after 9 AM the next day.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcInteractiveTimeWithContext(ctx context.Context, request *DescribeTrtcInteractiveTimeRequest) (response *DescribeTrtcInteractiveTimeResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcInteractiveTimeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTrtcInteractiveTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcMcuTranscodeTimeRequest() (request *DescribeTrtcMcuTranscodeTimeRequest) {
    request = &DescribeTrtcMcuTranscodeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcMcuTranscodeTime")
    
    
    return
}

func NewDescribeTrtcMcuTranscodeTimeResponse() (response *DescribeTrtcMcuTranscodeTimeResponse) {
    response = &DescribeTrtcMcuTranscodeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrtcMcuTranscodeTime
// This API is used to query billable relaying and transcoding durations.
//
// - If the period queried is 1 day or shorter, the statistics returned are on a 5-minute basis. If the period queried is longer than 1 day, the statistics returned are on a daily basis.
//
// - The period queried in a request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - In the daily pay-as-you-go mode, bills for a day are generated the next morning. Therefore, we recommend you query the statistics after 9 AM the next day.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcMcuTranscodeTime(request *DescribeTrtcMcuTranscodeTimeRequest) (response *DescribeTrtcMcuTranscodeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcMcuTranscodeTimeRequest()
    }
    
    response = NewDescribeTrtcMcuTranscodeTimeResponse()
    err = c.Send(request, response)
    return
}

// DescribeTrtcMcuTranscodeTime
// This API is used to query billable relaying and transcoding durations.
//
// - If the period queried is 1 day or shorter, the statistics returned are on a 5-minute basis. If the period queried is longer than 1 day, the statistics returned are on a daily basis.
//
// - The period queried in a request cannot be longer than 31 days.
//
// - If you query the statistics of the current day, the statistics returned may be inaccurate due to the delay in data collection.
//
// - In the daily pay-as-you-go mode, bills for a day are generated the next morning. Therefore, we recommend you query the statistics after 9 AM the next day.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcMcuTranscodeTimeWithContext(ctx context.Context, request *DescribeTrtcMcuTranscodeTimeRequest) (response *DescribeTrtcMcuTranscodeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcMcuTranscodeTimeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTrtcMcuTranscodeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInformationRequest() (request *DescribeUserInformationRequest) {
    request = &DescribeUserInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUserInformation")
    
    
    return
}

func NewDescribeUserInformationResponse() (response *DescribeUserInformationResponse) {
    response = &DescribeUserInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserInformation
// This API is used to query the user list of a specified time range (up to 4 hours) in the last 14 days. Data of 6 users is displayed on each page by default, and data of up to 100 users can be displayed on each page (`PageSize` set to 100 or smaller).
//
// **Note**: You can use this API to query or check historical data, but not for real-time key business logic.
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
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInformation(request *DescribeUserInformationRequest) (response *DescribeUserInformationResponse, err error) {
    if request == nil {
        request = NewDescribeUserInformationRequest()
    }
    
    response = NewDescribeUserInformationResponse()
    err = c.Send(request, response)
    return
}

// DescribeUserInformation
// This API is used to query the user list of a specified time range (up to 4 hours) in the last 14 days. Data of 6 users is displayed on each page by default, and data of up to 100 users can be displayed on each page (`PageSize` set to 100 or smaller).
//
// **Note**: You can use this API to query or check historical data, but not for real-time key business logic.
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
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInformationWithContext(ctx context.Context, request *DescribeUserInformationRequest) (response *DescribeUserInformationResponse, err error) {
    if request == nil {
        request = NewDescribeUserInformationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUserInformationResponse()
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
    if request == nil {
        request = NewDismissRoomRequest()
    }
    
    response = NewDismissRoomResponse()
    err = c.Send(request, response)
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
func (c *Client) DismissRoomWithContext(ctx context.Context, request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    if request == nil {
        request = NewDismissRoomRequest()
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
    if request == nil {
        request = NewDismissRoomByStrRoomIdRequest()
    }
    
    response = NewDismissRoomByStrRoomIdResponse()
    err = c.Send(request, response)
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
func (c *Client) DismissRoomByStrRoomIdWithContext(ctx context.Context, request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewDismissRoomByStrRoomIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPictureRequest() (request *ModifyPictureRequest) {
    request = &ModifyPictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyPicture")
    
    
    return
}

func NewModifyPictureResponse() (response *ModifyPictureResponse) {
    response = &ModifyPictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPicture
// This API is used to modify custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to modify such images frequently, we recommend you modify them in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyPicture(request *ModifyPictureRequest) (response *ModifyPictureResponse, err error) {
    if request == nil {
        request = NewModifyPictureRequest()
    }
    
    response = NewModifyPictureResponse()
    err = c.Send(request, response)
    return
}

// ModifyPicture
// This API is used to modify custom background or watermark images during [On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/16827?from_cn_redirect=1). If you do not need to modify such images frequently, we recommend you modify them in the console via **Application Management** > **[Material Management](https://intl.cloud.tencent.com/document/product/647/50769?from_cn_redirect=1)**.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyPictureWithContext(ctx context.Context, request *ModifyPictureRequest) (response *ModifyPictureResponse, err error) {
    if request == nil {
        request = NewModifyPictureRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPictureResponse()
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
    if request == nil {
        request = NewRemoveUserRequest()
    }
    
    response = NewRemoveUserResponse()
    err = c.Send(request, response)
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
func (c *Client) RemoveUserWithContext(ctx context.Context, request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    if request == nil {
        request = NewRemoveUserRequest()
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
    if request == nil {
        request = NewRemoveUserByStrRoomIdRequest()
    }
    
    response = NewRemoveUserByStrRoomIdResponse()
    err = c.Send(request, response)
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
func (c *Client) RemoveUserByStrRoomIdWithContext(ctx context.Context, request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewRemoveUserByStrRoomIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewRemoveUserByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStartMCUMixTranscodeRequest() (request *StartMCUMixTranscodeRequest) {
    request = &StartMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscode")
    
    
    return
}

func NewStartMCUMixTranscodeResponse() (response *StartMCUMixTranscodeResponse) {
    response = &StartMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMCUMixTranscode
// This API is used to enable On-Cloud MixTranscoding and specify the position of each channel of image in stream mixing.
//
// 
//
// There may be multiple audio/video streams in a TRTC room. You can call this API to request the Tencent Cloud server to mix the audio and video and specify the position of each image to produce just one audio/video stream for recording and playback. The mixing stops automatically after a room is terminated.
//
// 
//
// You can use this API to perform the following operations:
//
// - Set image and audio quality parameters of the final live stream, including video resolution, video bitrate, video frame rate, and audio quality.
//
// - Set the layout, i.e., the position of each image. You only need to set it once when enabling On-Cloud MixTranscoding, and the layout engine will automatically arrange images as configured.
//
// - Set the names of recording files for future playback.
//
// - Set the stream ID for CDN live streaming.
//
// 
//
// Currently, On-Cloud MixTranscoding supports the following layout templates:
//
// - Floating: The entire screen is covered by the video image of the first user who enters the room, and the images of other users are displayed as small images in rows in the bottom-left corner in room entry sequence. The screen allows up to 4 rows of 4 small images, which float over the big image. Up to 1 big image and 15 small images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Grid: The images of all users split the entire screen evenly. The more users, the smaller the image dimensions. Up to 16 images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Screen sharing: This is designed for video conferencing and online education. The shared screen (or camera image of the anchor) is always displayed as the big image, which occupies the left half of the screen, and the images of other users occupy the right half in up to 2 columns of up to 8 small images each. Up to 1 big image and 15 small images can be displayed. If the upstream aspect ratio does not match the output, the big image on the left will be scaled and displayed in whole, while the small images on the right will be cropped.
//
// - Picture-in-picture: This template mixes the big and small images or big image of a user with the audio of other users. The small image floats over the big image. You can specify the user whose small and big images are displayed, as well as the position of the small image. This template can display at most 2 images.
//
// - Custom: This is designed for cases where you want to specify the image positions of users in the mixed stream or preset image positions. If users are assigned to preset positions, the layout engine will reserve the positions for the users; if not, users will occupy the positions in room entry sequence. Once all preset positions are occupied, TRTC will stop mixing the audio and video of other users. If the place-holding feature is enabled for a custom template (by setting `PlaceHolderMode` in `LayoutParams` to `1`) and a user for whom a place is held is not sending video, the position will show the specified placeholder image (`PlaceImageId`).
//
// 
//
// Notes:
//
// 1. **As On-Cloud MixTranscoding is a paid feature, you will be charged for calling this API. For details, see [Billing of On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/49446?from_cn_redirect=1).**
//
// 2. You can call this API only if your application is created on or after January 9, 2020. Applications created before use the stream mixing service of CSS by default. If you want to switch to MCU On-Cloud MixTranscoding, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// 3. You cannot use the server and client stream mixing APIs at the same time.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscode(request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeRequest()
    }
    
    response = NewStartMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

// StartMCUMixTranscode
// This API is used to enable On-Cloud MixTranscoding and specify the position of each channel of image in stream mixing.
//
// 
//
// There may be multiple audio/video streams in a TRTC room. You can call this API to request the Tencent Cloud server to mix the audio and video and specify the position of each image to produce just one audio/video stream for recording and playback. The mixing stops automatically after a room is terminated.
//
// 
//
// You can use this API to perform the following operations:
//
// - Set image and audio quality parameters of the final live stream, including video resolution, video bitrate, video frame rate, and audio quality.
//
// - Set the layout, i.e., the position of each image. You only need to set it once when enabling On-Cloud MixTranscoding, and the layout engine will automatically arrange images as configured.
//
// - Set the names of recording files for future playback.
//
// - Set the stream ID for CDN live streaming.
//
// 
//
// Currently, On-Cloud MixTranscoding supports the following layout templates:
//
// - Floating: The entire screen is covered by the video image of the first user who enters the room, and the images of other users are displayed as small images in rows in the bottom-left corner in room entry sequence. The screen allows up to 4 rows of 4 small images, which float over the big image. Up to 1 big image and 15 small images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Grid: The images of all users split the entire screen evenly. The more users, the smaller the image dimensions. Up to 16 images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Screen sharing: This is designed for video conferencing and online education. The shared screen (or camera image of the anchor) is always displayed as the big image, which occupies the left half of the screen, and the images of other users occupy the right half in up to 2 columns of up to 8 small images each. Up to 1 big image and 15 small images can be displayed. If the upstream aspect ratio does not match the output, the big image on the left will be scaled and displayed in whole, while the small images on the right will be cropped.
//
// - Picture-in-picture: This template mixes the big and small images or big image of a user with the audio of other users. The small image floats over the big image. You can specify the user whose small and big images are displayed, as well as the position of the small image. This template can display at most 2 images.
//
// - Custom: This is designed for cases where you want to specify the image positions of users in the mixed stream or preset image positions. If users are assigned to preset positions, the layout engine will reserve the positions for the users; if not, users will occupy the positions in room entry sequence. Once all preset positions are occupied, TRTC will stop mixing the audio and video of other users. If the place-holding feature is enabled for a custom template (by setting `PlaceHolderMode` in `LayoutParams` to `1`) and a user for whom a place is held is not sending video, the position will show the specified placeholder image (`PlaceImageId`).
//
// 
//
// Notes:
//
// 1. **As On-Cloud MixTranscoding is a paid feature, you will be charged for calling this API. For details, see [Billing of On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/49446?from_cn_redirect=1).**
//
// 2. You can call this API only if your application is created on or after January 9, 2020. Applications created before use the stream mixing service of CSS by default. If you want to switch to MCU On-Cloud MixTranscoding, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// 3. You cannot use the server and client stream mixing APIs at the same time.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeWithContext(ctx context.Context, request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStartMCUMixTranscodeByStrRoomIdRequest() (request *StartMCUMixTranscodeByStrRoomIdRequest) {
    request = &StartMCUMixTranscodeByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscodeByStrRoomId")
    
    
    return
}

func NewStartMCUMixTranscodeByStrRoomIdResponse() (response *StartMCUMixTranscodeByStrRoomIdResponse) {
    response = &StartMCUMixTranscodeByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMCUMixTranscodeByStrRoomId
// This API is used to enable On-Cloud MixTranscoding and specify the position of each channel of image in stream mixing.
//
// 
//
// There may be multiple channels of audio/video streams in a TRTC room. You can call this API to request the Tencent Cloud server to mix multiple channels of video images and audio into one channel and specify the position of each image so as to produce only one channel of audio/video stream for recording and live streaming.
//
// 
//
// You can use this API to perform the following operations:
//
// - Set image and audio quality parameters of the mixed stream, including video resolution, bitrate, frame rate, and audio quality.
//
// - Set the layout, i.e., the position of each channel of image. You only need to set it once when enabling On-Cloud MixTranscoding, and the layout engine will automatically arrange images as configured.
//
// - Set the names of recording files for future playback.
//
// - Set the stream ID for CDN live streaming.
//
// 
//
// Currently, On-Cloud MixTranscoding supports the following layout templates:
//
// - Floating: the entire screen is covered by the video image of the first user who enters the room, and the images of other users are displayed as small images in horizontal rows in the bottom-left corner in room entry sequence. The screen can accommodate up to 4 rows of 4 small images, which float over the big image. Up to 1 big image and 15 small images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Grid: the images of all users split the screen evenly. The more the users, the smaller the image dimensions. Up to 16 images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Screen sharing: this template is designed for video conferencing and online classes. The shared screen (or camera image of the anchor) is always displayed as the big image, which occupies the left half of the screen, and the images of other users occupy the right half in up to 2 columns of a maximum of 8 small images each. Up to 1 big image and 15 small images can be displayed. If the aspect ratio of upstream images does not match that of output images, the big image on the left will be scaled and displayed in whole, while the small images on the right will be cropped.
//
// - Picture-in-picture: this template mixes the big and small images or big image of a user with the audio of other users. The small image floats over the big image. You can specify the user whose big and small images are displayed and the position of the small image.
//
// - Custom: you can use custom templates to specify the image positions of users in mixed streams or preset image positions. If users are assigned to preset positions, the layout engine will reserve the positions for the users; if not, users will occupy the positions in room entry sequence. Once all preset positions are occupied, TRTC will stop mixing the audio and images of other users. If the placeholding feature is enabled for a custom template (`PlaceHolderMode` in `LayoutParams` is set to 1), but a user for whom a place is reserved is not sending video data, the position will show the corresponding placeholder image (`PlaceImageId`).
//
// 
//
// Notes:
//
// 1. **As On-Cloud MixTranscoding is a paid feature, you will be charged for calling this API. For details, see [Billing of On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/49446?from_cn_redirect=1).**
//
// 2. You can call this API only if your application is created on or after January 9, 2020. Applications created before use the stream mixing service of CSS by default. If you want to switch to MCU On-Cloud MixTranscoding, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// 3. You cannot use the server and client stream mixing APIs at the same time.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeByStrRoomId(request *StartMCUMixTranscodeByStrRoomIdRequest) (response *StartMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeByStrRoomIdRequest()
    }
    
    response = NewStartMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

// StartMCUMixTranscodeByStrRoomId
// This API is used to enable On-Cloud MixTranscoding and specify the position of each channel of image in stream mixing.
//
// 
//
// There may be multiple channels of audio/video streams in a TRTC room. You can call this API to request the Tencent Cloud server to mix multiple channels of video images and audio into one channel and specify the position of each image so as to produce only one channel of audio/video stream for recording and live streaming.
//
// 
//
// You can use this API to perform the following operations:
//
// - Set image and audio quality parameters of the mixed stream, including video resolution, bitrate, frame rate, and audio quality.
//
// - Set the layout, i.e., the position of each channel of image. You only need to set it once when enabling On-Cloud MixTranscoding, and the layout engine will automatically arrange images as configured.
//
// - Set the names of recording files for future playback.
//
// - Set the stream ID for CDN live streaming.
//
// 
//
// Currently, On-Cloud MixTranscoding supports the following layout templates:
//
// - Floating: the entire screen is covered by the video image of the first user who enters the room, and the images of other users are displayed as small images in horizontal rows in the bottom-left corner in room entry sequence. The screen can accommodate up to 4 rows of 4 small images, which float over the big image. Up to 1 big image and 15 small images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Grid: the images of all users split the screen evenly. The more the users, the smaller the image dimensions. Up to 16 images can be displayed. A user sending audio only will still occupy an image spot.
//
// - Screen sharing: this template is designed for video conferencing and online classes. The shared screen (or camera image of the anchor) is always displayed as the big image, which occupies the left half of the screen, and the images of other users occupy the right half in up to 2 columns of a maximum of 8 small images each. Up to 1 big image and 15 small images can be displayed. If the aspect ratio of upstream images does not match that of output images, the big image on the left will be scaled and displayed in whole, while the small images on the right will be cropped.
//
// - Picture-in-picture: this template mixes the big and small images or big image of a user with the audio of other users. The small image floats over the big image. You can specify the user whose big and small images are displayed and the position of the small image.
//
// - Custom: you can use custom templates to specify the image positions of users in mixed streams or preset image positions. If users are assigned to preset positions, the layout engine will reserve the positions for the users; if not, users will occupy the positions in room entry sequence. Once all preset positions are occupied, TRTC will stop mixing the audio and images of other users. If the placeholding feature is enabled for a custom template (`PlaceHolderMode` in `LayoutParams` is set to 1), but a user for whom a place is reserved is not sending video data, the position will show the corresponding placeholder image (`PlaceImageId`).
//
// 
//
// Notes:
//
// 1. **As On-Cloud MixTranscoding is a paid feature, you will be charged for calling this API. For details, see [Billing of On-Cloud MixTranscoding](https://intl.cloud.tencent.com/document/product/647/49446?from_cn_redirect=1).**
//
// 2. You can call this API only if your application is created on or after January 9, 2020. Applications created before use the stream mixing service of CSS by default. If you want to switch to MCU On-Cloud MixTranscoding, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// 3. You cannot use the server and client stream mixing APIs at the same time.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeByStrRoomIdWithContext(ctx context.Context, request *StartMCUMixTranscodeByStrRoomIdRequest) (response *StartMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeByStrRoomIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeRequest() (request *StopMCUMixTranscodeRequest) {
    request = &StopMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscode")
    
    
    return
}

func NewStopMCUMixTranscodeResponse() (response *StopMCUMixTranscodeResponse) {
    response = &StopMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMCUMixTranscode
// This API is used to end On-Cloud MixTranscoding.
//
// error code that may be returned:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscode(request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeRequest()
    }
    
    response = NewStopMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

// StopMCUMixTranscode
// This API is used to end On-Cloud MixTranscoding.
//
// error code that may be returned:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeWithContext(ctx context.Context, request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeByStrRoomIdRequest() (request *StopMCUMixTranscodeByStrRoomIdRequest) {
    request = &StopMCUMixTranscodeByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscodeByStrRoomId")
    
    
    return
}

func NewStopMCUMixTranscodeByStrRoomIdResponse() (response *StopMCUMixTranscodeByStrRoomIdResponse) {
    response = &StopMCUMixTranscodeByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMCUMixTranscodeByStrRoomId
// This API is used to stop On-Cloud MixTranscoding.
//
// error code that may be returned:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeByStrRoomId(request *StopMCUMixTranscodeByStrRoomIdRequest) (response *StopMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeByStrRoomIdRequest()
    }
    
    response = NewStopMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

// StopMCUMixTranscodeByStrRoomId
// This API is used to stop On-Cloud MixTranscoding.
//
// error code that may be returned:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeByStrRoomIdWithContext(ctx context.Context, request *StopMCUMixTranscodeByStrRoomIdRequest) (response *StopMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeByStrRoomIdRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}
