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

package v20180410

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-10"

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


func NewAcceptDirectConnectTunnelRequest() (request *AcceptDirectConnectTunnelRequest) {
    request = &AcceptDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "AcceptDirectConnectTunnel")
    return
}

func NewAcceptDirectConnectTunnelResponse() (response *AcceptDirectConnectTunnelResponse) {
    response = &AcceptDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to accept an application for a dedicated tunnel.
func (c *Client) AcceptDirectConnectTunnel(request *AcceptDirectConnectTunnelRequest) (response *AcceptDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewAcceptDirectConnectTunnelRequest()
    }
    response = NewAcceptDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectTunnelRequest() (request *CreateDirectConnectTunnelRequest) {
    request = &CreateDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "CreateDirectConnectTunnel")
    return
}

func NewCreateDirectConnectTunnelResponse() (response *CreateDirectConnectTunnelResponse) {
    response = &CreateDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a dedicated tunnel.
func (c *Client) CreateDirectConnectTunnel(request *CreateDirectConnectTunnelRequest) (response *CreateDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectTunnelRequest()
    }
    response = NewCreateDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectTunnelRequest() (request *DeleteDirectConnectTunnelRequest) {
    request = &DeleteDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DeleteDirectConnectTunnel")
    return
}

func NewDeleteDirectConnectTunnelResponse() (response *DeleteDirectConnectTunnelResponse) {
    response = &DeleteDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a dedicated tunnel.
func (c *Client) DeleteDirectConnectTunnel(request *DeleteDirectConnectTunnelRequest) (response *DeleteDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectTunnelRequest()
    }
    response = NewDeleteDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectTunnelsRequest() (request *DescribeDirectConnectTunnelsRequest) {
    request = &DescribeDirectConnectTunnelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnectTunnels")
    return
}

func NewDescribeDirectConnectTunnelsResponse() (response *DescribeDirectConnectTunnelsResponse) {
    response = &DescribeDirectConnectTunnelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of dedicated tunnels.
func (c *Client) DescribeDirectConnectTunnels(request *DescribeDirectConnectTunnelsRequest) (response *DescribeDirectConnectTunnelsResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectTunnelsRequest()
    }
    response = NewDescribeDirectConnectTunnelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectsRequest() (request *DescribeDirectConnectsRequest) {
    request = &DescribeDirectConnectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "DescribeDirectConnects")
    return
}

func NewDescribeDirectConnectsResponse() (response *DescribeDirectConnectsResponse) {
    response = &DescribeDirectConnectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of connections.
func (c *Client) DescribeDirectConnects(request *DescribeDirectConnectsRequest) (response *DescribeDirectConnectsResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectsRequest()
    }
    response = NewDescribeDirectConnectsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDirectConnectTunnelAttributeRequest() (request *ModifyDirectConnectTunnelAttributeRequest) {
    request = &ModifyDirectConnectTunnelAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "ModifyDirectConnectTunnelAttribute")
    return
}

func NewModifyDirectConnectTunnelAttributeResponse() (response *ModifyDirectConnectTunnelAttributeResponse) {
    response = &ModifyDirectConnectTunnelAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the dedicated tunnel attributes.
func (c *Client) ModifyDirectConnectTunnelAttribute(request *ModifyDirectConnectTunnelAttributeRequest) (response *ModifyDirectConnectTunnelAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectTunnelAttributeRequest()
    }
    response = NewModifyDirectConnectTunnelAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRejectDirectConnectTunnelRequest() (request *RejectDirectConnectTunnelRequest) {
    request = &RejectDirectConnectTunnelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dc", APIVersion, "RejectDirectConnectTunnel")
    return
}

func NewRejectDirectConnectTunnelResponse() (response *RejectDirectConnectTunnelResponse) {
    response = &RejectDirectConnectTunnelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reject an application for a dedicated tunnel.
func (c *Client) RejectDirectConnectTunnel(request *RejectDirectConnectTunnelRequest) (response *RejectDirectConnectTunnelResponse, err error) {
    if request == nil {
        request = NewRejectDirectConnectTunnelRequest()
    }
    response = NewRejectDirectConnectTunnelResponse()
    err = c.Send(request, response)
    return
}
