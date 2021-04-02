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

package v20200820

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-08-20"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCancelMatchingRequest() (request *CancelMatchingRequest) {
    request = &CancelMatchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CancelMatching")
    return
}

func NewCancelMatchingResponse() (response *CancelMatchingResponse) {
    response = &CancelMatchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to cancel matching.
func (c *Client) CancelMatching(request *CancelMatchingRequest) (response *CancelMatchingResponse, err error) {
    if request == nil {
        request = NewCancelMatchingRequest()
    }
    response = NewCancelMatchingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchingProgressRequest() (request *DescribeMatchingProgressRequest) {
    request = &DescribeMatchingProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatchingProgress")
    return
}

func NewDescribeMatchingProgressResponse() (response *DescribeMatchingProgressResponse) {
    response = &DescribeMatchingProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the matching progress.
func (c *Client) DescribeMatchingProgress(request *DescribeMatchingProgressRequest) (response *DescribeMatchingProgressResponse, err error) {
    if request == nil {
        request = NewDescribeMatchingProgressRequest()
    }
    response = NewDescribeMatchingProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenRequest() (request *DescribeTokenRequest) {
    request = &DescribeTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeToken")
    return
}

func NewDescribeTokenResponse() (response *DescribeTokenResponse) {
    response = &DescribeTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the token of a Matchcode, which is used for verified the pushed message.
func (c *Client) DescribeToken(request *DescribeTokenRequest) (response *DescribeTokenResponse, err error) {
    if request == nil {
        request = NewDescribeTokenRequest()
    }
    response = NewDescribeTokenResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTokenRequest() (request *ModifyTokenRequest) {
    request = &ModifyTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "ModifyToken")
    return
}

func NewModifyTokenResponse() (response *ModifyTokenResponse) {
    response = &ModifyTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the token of a Matchcode.
func (c *Client) ModifyToken(request *ModifyTokenRequest) (response *ModifyTokenResponse, err error) {
    if request == nil {
        request = NewModifyTokenRequest()
    }
    response = NewModifyTokenResponse()
    err = c.Send(request, response)
    return
}

func NewStartMatchingRequest() (request *StartMatchingRequest) {
    request = &StartMatchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "StartMatching")
    return
}

func NewStartMatchingResponse() (response *StartMatchingResponse) {
    response = &StartMatchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to pass in one player or multiple players to initiate match. Players within the same request will be assigned to the same game session.
func (c *Client) StartMatching(request *StartMatchingRequest) (response *StartMatchingResponse, err error) {
    if request == nil {
        request = NewStartMatchingRequest()
    }
    response = NewStartMatchingResponse()
    err = c.Send(request, response)
    return
}
