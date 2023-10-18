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


func NewDescribeCaptchaResultRequest() (request *DescribeCaptchaResultRequest) {
    request = &DescribeCaptchaResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaResult")
    
    
    return
}

func NewDescribeCaptchaResultResponse() (response *DescribeCaptchaResultResponse) {
    response = &DescribeCaptchaResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCaptchaResult
// This API is used to query the result of CAPTCHA ticket verification (web and app).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaResult(request *DescribeCaptchaResultRequest) (response *DescribeCaptchaResultResponse, err error) {
    return c.DescribeCaptchaResultWithContext(context.Background(), request)
}

// DescribeCaptchaResult
// This API is used to query the result of CAPTCHA ticket verification (web and app).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_ERRAUTH = "UnauthorizedOperation.ErrAuth"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
func (c *Client) DescribeCaptchaResultWithContext(ctx context.Context, request *DescribeCaptchaResultRequest) (response *DescribeCaptchaResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCaptchaResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCaptchaResultResponse()
    err = c.Send(request, response)
    return
}
