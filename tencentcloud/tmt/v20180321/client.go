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

package v20180321

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-03-21"

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


func NewTextTranslateRequest() (request *TextTranslateRequest) {
    request = &TextTranslateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "TextTranslate")
    
    
    return
}

func NewTextTranslateResponse() (response *TextTranslateResponse) {
    response = &TextTranslateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextTranslate
// This API is used to translate text in multiple language pairs, such as Chinese-English.<br />
//
// Note: We recommend that you simplify your development with the SDK integration mode. For how to use the SDK, see Section 5 "Developer Resources".
//
// error code that may be returned:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) TextTranslate(request *TextTranslateRequest) (response *TextTranslateResponse, err error) {
    return c.TextTranslateWithContext(context.Background(), request)
}

// TextTranslate
// This API is used to translate text in multiple language pairs, such as Chinese-English.<br />
//
// Note: We recommend that you simplify your development with the SDK integration mode. For how to use the SDK, see Section 5 "Developer Resources".
//
// error code that may be returned:
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LIMITEDACCESSFREQUENCY = "LimitExceeded.LimitedAccessFrequency"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) TextTranslateWithContext(ctx context.Context, request *TextTranslateRequest) (response *TextTranslateResponse, err error) {
    if request == nil {
        request = NewTextTranslateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextTranslate require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextTranslateResponse()
    err = c.Send(request, response)
    return
}
