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


func NewImageTranslateLLMRequest() (request *ImageTranslateLLMRequest) {
    request = &ImageTranslateLLMRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tmt", APIVersion, "ImageTranslateLLM")
    
    
    return
}

func NewImageTranslateLLMResponse() (response *ImageTranslateLLMResponse) {
    response = &ImageTranslateLLMResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageTranslateLLM
// This API is used to provide translation service for images in 18 languages. It can automatically recognize text content in images and translate it into the target language. The recognized text is translated line by line, and a version that supports paragraph translation will be offered subsequently.
//
// 
//
// -Input image format: png, jpg, jpeg and other common image formats. gif animation is not supported.
//
// -Output image format: jpg.
//
// 
//
// Notification: For general developers, we recommend prioritizing SDK integration to simplify development. For SDK usage introduction, directly view the 5. Developer Resources part.
//
// error code that may be returned:
//  FAILEDOPERATION_DECODEERR = "FailedOperation.DecodeErr"
//  FAILEDOPERATION_DOWNLOADERR = "FailedOperation.DownloadErr"
//  FAILEDOPERATION_ERRORUSERAREA = "FailedOperation.ErrorUserArea"
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_STOPUSING = "FailedOperation.StopUsing"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) ImageTranslateLLM(request *ImageTranslateLLMRequest) (response *ImageTranslateLLMResponse, err error) {
    return c.ImageTranslateLLMWithContext(context.Background(), request)
}

// ImageTranslateLLM
// This API is used to provide translation service for images in 18 languages. It can automatically recognize text content in images and translate it into the target language. The recognized text is translated line by line, and a version that supports paragraph translation will be offered subsequently.
//
// 
//
// -Input image format: png, jpg, jpeg and other common image formats. gif animation is not supported.
//
// -Output image format: jpg.
//
// 
//
// Notification: For general developers, we recommend prioritizing SDK integration to simplify development. For SDK usage introduction, directly view the 5. Developer Resources part.
//
// error code that may be returned:
//  FAILEDOPERATION_DECODEERR = "FailedOperation.DecodeErr"
//  FAILEDOPERATION_DOWNLOADERR = "FailedOperation.DownloadErr"
//  FAILEDOPERATION_ERRORUSERAREA = "FailedOperation.ErrorUserArea"
//  FAILEDOPERATION_NOFREEAMOUNT = "FailedOperation.NoFreeAmount"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_STOPUSING = "FailedOperation.StopUsing"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeout"
//  INTERNALERROR_ERRORUNKNOWN = "InternalError.ErrorUnknown"
//  INTERNALERROR_REQUESTFAILED = "InternalError.RequestFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ACTIONNOTFOUND = "UnauthorizedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDTARGETLANGUAGE = "UnsupportedOperation.UnSupportedTargetLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDLANGUAGE = "UnsupportedOperation.UnsupportedLanguage"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSOURCELANGUAGE = "UnsupportedOperation.UnsupportedSourceLanguage"
func (c *Client) ImageTranslateLLMWithContext(ctx context.Context, request *ImageTranslateLLMRequest) (response *ImageTranslateLLMResponse, err error) {
    if request == nil {
        request = NewImageTranslateLLMRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tmt", APIVersion, "ImageTranslateLLM")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageTranslateLLM require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageTranslateLLMResponse()
    err = c.Send(request, response)
    return
}
