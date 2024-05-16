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

package v20180301

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-03-01"

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


func NewApplySdkVerificationTokenRequest() (request *ApplySdkVerificationTokenRequest) {
    request = &ApplySdkVerificationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ApplySdkVerificationToken")
    
    
    return
}

func NewApplySdkVerificationTokenResponse() (response *ApplySdkVerificationTokenResponse) {
    response = &ApplySdkVerificationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplySdkVerificationToken
// This API is used to apply for a token before calling the Identity Verification SDK service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
func (c *Client) ApplySdkVerificationToken(request *ApplySdkVerificationTokenRequest) (response *ApplySdkVerificationTokenResponse, err error) {
    return c.ApplySdkVerificationTokenWithContext(context.Background(), request)
}

// ApplySdkVerificationToken
// This API is used to apply for a token before calling the Identity Verification SDK service each time. This token is required for initiating the verification process and getting the result after the verification is completed.
func (c *Client) ApplySdkVerificationTokenWithContext(ctx context.Context, request *ApplySdkVerificationTokenRequest) (response *ApplySdkVerificationTokenResponse, err error) {
    if request == nil {
        request = NewApplySdkVerificationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplySdkVerificationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplySdkVerificationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewGetSdkVerificationResultRequest() (request *GetSdkVerificationResultRequest) {
    request = &GetSdkVerificationResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetSdkVerificationResult")
    
    
    return
}

func NewGetSdkVerificationResultResponse() (response *GetSdkVerificationResultResponse) {
    response = &GetSdkVerificationResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSdkVerificationResult
// This API is used to get the verification result with the corresponding token after the SDK-based verification is completed. The token is valid for three days after issuance and can be called multiple times.
func (c *Client) GetSdkVerificationResult(request *GetSdkVerificationResultRequest) (response *GetSdkVerificationResultResponse, err error) {
    return c.GetSdkVerificationResultWithContext(context.Background(), request)
}

// GetSdkVerificationResult
// This API is used to get the verification result with the corresponding token after the SDK-based verification is completed. The token is valid for three days after issuance and can be called multiple times.
func (c *Client) GetSdkVerificationResultWithContext(ctx context.Context, request *GetSdkVerificationResultRequest) (response *GetSdkVerificationResultResponse, err error) {
    if request == nil {
        request = NewGetSdkVerificationResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSdkVerificationResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSdkVerificationResultResponse()
    err = c.Send(request, response)
    return
}
