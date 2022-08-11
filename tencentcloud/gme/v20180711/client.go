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

package v20180711

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-11"

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


func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApp
// This API is used to create a GME application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// This API is used to create a GME application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppStatisticsRequest() (request *DescribeAppStatisticsRequest) {
    request = &DescribeAppStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAppStatistics")
    
    
    return
}

func NewDescribeAppStatisticsResponse() (response *DescribeAppStatisticsResponse) {
    response = &DescribeAppStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppStatistics
// This API is used to get the usage statistics of a GME application, including those of voice chat, voice messaging and speech-to-text, phrase analysis, etc. The maximum query period is the past 30 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatistics(request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    return c.DescribeAppStatisticsWithContext(context.Background(), request)
}

// DescribeAppStatistics
// This API is used to get the usage statistics of a GME application, including those of voice chat, voice messaging and speech-to-text, phrase analysis, etc. The maximum query period is the past 30 days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatisticsWithContext(ctx context.Context, request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationDataRequest() (request *DescribeApplicationDataRequest) {
    request = &DescribeApplicationDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeApplicationData")
    
    
    return
}

func NewDescribeApplicationDataResponse() (response *DescribeApplicationDataResponse) {
    response = &DescribeApplicationDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationData
// This API (DescribeApplicationData) is used to query usage data details within 90 days.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationData(request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    return c.DescribeApplicationDataWithContext(context.Background(), request)
}

// DescribeApplicationData
// This API (DescribeApplicationData) is used to query usage data details within 90 days.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationDataWithContext(ctx context.Context, request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanResultListRequest() (request *DescribeScanResultListRequest) {
    request = &DescribeScanResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeScanResultList")
    
    
    return
}

func NewDescribeScanResultListResponse() (response *DescribeScanResultListResponse) {
    response = &DescribeScanResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanResultList
// This API is used to query the speech detection result. Up to 100 tasks can be added in the task query list.
//
// <p style="color:red">If the `Callback` field is not set when a speech detection task is submitted, this API will be needed to get the detection result.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultList(request *DescribeScanResultListRequest) (response *DescribeScanResultListResponse, err error) {
    return c.DescribeScanResultListWithContext(context.Background(), request)
}

// DescribeScanResultList
// This API is used to query the speech detection result. Up to 100 tasks can be added in the task query list.
//
// <p style="color:red">If the `Callback` field is not set when a speech detection task is submitted, this API will be needed to get the detection result.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultListWithContext(ctx context.Context, request *DescribeScanResultListRequest) (response *DescribeScanResultListResponse, err error) {
    if request == nil {
        request = NewDescribeScanResultListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanResultList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanResultListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppStatusRequest() (request *ModifyAppStatusRequest) {
    request = &ModifyAppStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "ModifyAppStatus")
    
    
    return
}

func NewModifyAppStatusResponse() (response *ModifyAppStatusResponse) {
    response = &ModifyAppStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAppStatus
// This API is used to change the status of an application's primary switch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatus(request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    return c.ModifyAppStatusWithContext(context.Background(), request)
}

// ModifyAppStatus
// This API is used to change the status of an application's primary switch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatusWithContext(ctx context.Context, request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    if request == nil {
        request = NewModifyAppStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppStatusResponse()
    err = c.Send(request, response)
    return
}

func NewScanVoiceRequest() (request *ScanVoiceRequest) {
    request = &ScanVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "ScanVoice")
    
    
    return
}

func NewScanVoiceResponse() (response *ScanVoiceResponse) {
    response = &ScanVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanVoice
// This API is used to submit a speech detection task. Up to 100 tasks can be added in the detection task list. Before using this API, please activate the speech analysis service in [Console > Service Configuration](https://console.cloud.tencent.com/gamegme/conf).
//
// </br></br>
//
// 
//
// <h4><b>About the trial:</b></h4>
//
// <li>You can try out the speech analysis service free of charge in <a href="https://console.cloud.tencent.com/gamegme/tryout">Console > Product Trial</a>.</li>
//
// </br>
//
// 
//
// <h4><b>API feature description:</b></h4>
//
// <li>This API checks audio streams or files for non-compliant content.</li>
//
// <li>The detection result can be obtained by setting the callback address (`Callback`) or calling the `DescribeScanResultList` API for polling.</li>
//
// <li>The scenario can be specified, such as abusive or pornographic.</li>
//
// <li>Detection tasks can be submitted in batches. Up to 100 tasks can be added in the detection task list.</li>
//
// </br>
//
// <h4><b>Audio file limit description:</b></h4>
//
// <li>Audio file size limit: 100 MB</li>
//
// <li>Audio file duration limit: 30 minutes</li>
//
// <li>Supported audio file formats: .wav, .m4a, .amr, .mp3, .aac, .wma, .ogg</li>
//
// </br>
//
// <h4><b>Audio stream limit description:</b></h4>
//
// <li>Supported audio stream formats: .m3u8, .flv</li>
//
// <li>Supported audio stream transfer protocols: RTMP, HTTP, HTTPS</li>
//
// <li>Audio stream duration limit: 4 hours</li>
//
// <li>Audio/video stream separation and audio stream analysis are supported</li>
//
// </br>
//
// <h4 id="Label_Value"><b>`Scenes` and `Label` parameter description:</b></h4>
//
// <p>When submitting a speech detection task, you need to specify the `Scenes` parameter. <font color="red">You are currently required to set the `Scenes` parameter to `["default"]`</font>. The detection result will contain the scenario specified at the time of request and detection result in the corresponding type.</p>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>Scenario</th>
//
// <th>Description</th>
//
// <th>Label</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Speech detection</td>
//
// <td>Speech detection type</td>
//
// <td>
//
// <p>normal: normal</p>
//
// <p>porn: pornographic</p>
//
// <p>abuse: abusive</p>
//
// <p>ad: advertising</p>
//
// <p>contraband: prohibited</p>
//
// <p>customized: custom keyword library. This feature is only available to beta users. To try it out, please <a href="https://intl.cloud.tencent.com/apply/p/8809fjcik56?from_cn_redirect=1">contact us</a>.</p>
//
// </td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// </br>
//
// <h4 id="Callback_Declare"><b>Callback description:</b></h4>
//
// <li>If the callback address parameter `Callback` (i.e., the URL of an HTTP(S) API) is specified in the request parameters, then the POST method should be supported and transferred data should be encoded with UTF-8.</li>
//
// <li>After the callback data is pushed, if the HTTP status code received is 200, the push is successful.</li>
//
// <li>HTTP header parameter description:</li>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>Item</th>
//
// <th>Type</th>
//
// <th>Required</th>
//
// <th>Description</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Signatue</td>
//
// <td>string</td>
//
// <td>Yes</td>
//
// <td>Signature. For more information, please see <a href="#Callback_Signatue">Signature generation description</a>.</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// <ul  id="Callback_Signatue">
//
// 	<li>Signature generation description:</li>
//
// 	<ul>
//
// 		<li>The HMAC-SH1 algorithm should be used, and the result should be encoded with Base64;</li>
//
// 		<li>The original signature string is the entire JSON content of POST and body (the length is subject to `Content-Length`);</li>
//
// 		<li>The signature key is the `SecretKey` of the application, which can be viewed in the console.</li>
//
// 	</ul>
//
// </ul>
//
// 
//
// <li>Below is a sample callback <font color="red">(for more information on the fields, please see the structure:
//
// <a href="https://intl.cloud.tencent.com/document/api/607/35375?from_cn_redirect=1#DescribeScanResult" target="_blank">DescribeScanResult</a>)</font>:</li>
//
// <pre><code>{
//
// 	"Code": 0,
//
// 	"DataId": "1400000000_test_data_id",
//
// 	"ScanFinishTime": 1566720906,
//
// 	"HitFlag": true,
//
// 	"Live": false,
//
// 	"Msg": "",
//
// 	"ScanPiece": [{
//
// 		"DumpUrl": "",
//
// 		"HitFlag": true,
//
// 		"MainType": "abuse",
//
// 		"RoomId": "123",
//
// 		"OpenId": "xxx",
//
// 		"Info":"",
//
// 		"Offset": 0,
//
// 		"Duration": 3400,
//
// 		"PieceStartTime":1574684231,
//
// 		"ScanDetail": [{
//
// 			"EndTime": 1110,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 1110
//
// 		}, {
//
// 			"EndTime": 1380,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 1560,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 2820,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 2490
//
// 		}]
//
// 	}],
//
// 	"ScanStartTime": 1566720905,
//
// 	"Scenes": [
//
// 		"default"
//
// 	],
//
// 	"Status": "Success",
//
// 	"TaskId": "xxx",
//
// 	"Url": "https://xxx/xxx.m4a"
//
// }
//
// </code></pre>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ScanVoice(request *ScanVoiceRequest) (response *ScanVoiceResponse, err error) {
    return c.ScanVoiceWithContext(context.Background(), request)
}

// ScanVoice
// This API is used to submit a speech detection task. Up to 100 tasks can be added in the detection task list. Before using this API, please activate the speech analysis service in [Console > Service Configuration](https://console.cloud.tencent.com/gamegme/conf).
//
// </br></br>
//
// 
//
// <h4><b>About the trial:</b></h4>
//
// <li>You can try out the speech analysis service free of charge in <a href="https://console.cloud.tencent.com/gamegme/tryout">Console > Product Trial</a>.</li>
//
// </br>
//
// 
//
// <h4><b>API feature description:</b></h4>
//
// <li>This API checks audio streams or files for non-compliant content.</li>
//
// <li>The detection result can be obtained by setting the callback address (`Callback`) or calling the `DescribeScanResultList` API for polling.</li>
//
// <li>The scenario can be specified, such as abusive or pornographic.</li>
//
// <li>Detection tasks can be submitted in batches. Up to 100 tasks can be added in the detection task list.</li>
//
// </br>
//
// <h4><b>Audio file limit description:</b></h4>
//
// <li>Audio file size limit: 100 MB</li>
//
// <li>Audio file duration limit: 30 minutes</li>
//
// <li>Supported audio file formats: .wav, .m4a, .amr, .mp3, .aac, .wma, .ogg</li>
//
// </br>
//
// <h4><b>Audio stream limit description:</b></h4>
//
// <li>Supported audio stream formats: .m3u8, .flv</li>
//
// <li>Supported audio stream transfer protocols: RTMP, HTTP, HTTPS</li>
//
// <li>Audio stream duration limit: 4 hours</li>
//
// <li>Audio/video stream separation and audio stream analysis are supported</li>
//
// </br>
//
// <h4 id="Label_Value"><b>`Scenes` and `Label` parameter description:</b></h4>
//
// <p>When submitting a speech detection task, you need to specify the `Scenes` parameter. <font color="red">You are currently required to set the `Scenes` parameter to `["default"]`</font>. The detection result will contain the scenario specified at the time of request and detection result in the corresponding type.</p>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>Scenario</th>
//
// <th>Description</th>
//
// <th>Label</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Speech detection</td>
//
// <td>Speech detection type</td>
//
// <td>
//
// <p>normal: normal</p>
//
// <p>porn: pornographic</p>
//
// <p>abuse: abusive</p>
//
// <p>ad: advertising</p>
//
// <p>contraband: prohibited</p>
//
// <p>customized: custom keyword library. This feature is only available to beta users. To try it out, please <a href="https://intl.cloud.tencent.com/apply/p/8809fjcik56?from_cn_redirect=1">contact us</a>.</p>
//
// </td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// </br>
//
// <h4 id="Callback_Declare"><b>Callback description:</b></h4>
//
// <li>If the callback address parameter `Callback` (i.e., the URL of an HTTP(S) API) is specified in the request parameters, then the POST method should be supported and transferred data should be encoded with UTF-8.</li>
//
// <li>After the callback data is pushed, if the HTTP status code received is 200, the push is successful.</li>
//
// <li>HTTP header parameter description:</li>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>Item</th>
//
// <th>Type</th>
//
// <th>Required</th>
//
// <th>Description</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Signatue</td>
//
// <td>string</td>
//
// <td>Yes</td>
//
// <td>Signature. For more information, please see <a href="#Callback_Signatue">Signature generation description</a>.</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// <ul  id="Callback_Signatue">
//
// 	<li>Signature generation description:</li>
//
// 	<ul>
//
// 		<li>The HMAC-SH1 algorithm should be used, and the result should be encoded with Base64;</li>
//
// 		<li>The original signature string is the entire JSON content of POST and body (the length is subject to `Content-Length`);</li>
//
// 		<li>The signature key is the `SecretKey` of the application, which can be viewed in the console.</li>
//
// 	</ul>
//
// </ul>
//
// 
//
// <li>Below is a sample callback <font color="red">(for more information on the fields, please see the structure:
//
// <a href="https://intl.cloud.tencent.com/document/api/607/35375?from_cn_redirect=1#DescribeScanResult" target="_blank">DescribeScanResult</a>)</font>:</li>
//
// <pre><code>{
//
// 	"Code": 0,
//
// 	"DataId": "1400000000_test_data_id",
//
// 	"ScanFinishTime": 1566720906,
//
// 	"HitFlag": true,
//
// 	"Live": false,
//
// 	"Msg": "",
//
// 	"ScanPiece": [{
//
// 		"DumpUrl": "",
//
// 		"HitFlag": true,
//
// 		"MainType": "abuse",
//
// 		"RoomId": "123",
//
// 		"OpenId": "xxx",
//
// 		"Info":"",
//
// 		"Offset": 0,
//
// 		"Duration": 3400,
//
// 		"PieceStartTime":1574684231,
//
// 		"ScanDetail": [{
//
// 			"EndTime": 1110,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 1110
//
// 		}, {
//
// 			"EndTime": 1380,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 1560,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 2820,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 2490
//
// 		}]
//
// 	}],
//
// 	"ScanStartTime": 1566720905,
//
// 	"Scenes": [
//
// 		"default"
//
// 	],
//
// 	"Status": "Success",
//
// 	"TaskId": "xxx",
//
// 	"Url": "https://xxx/xxx.m4a"
//
// }
//
// </code></pre>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ScanVoiceWithContext(ctx context.Context, request *ScanVoiceRequest) (response *ScanVoiceResponse, err error) {
    if request == nil {
        request = NewScanVoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanVoiceResponse()
    err = c.Send(request, response)
    return
}
