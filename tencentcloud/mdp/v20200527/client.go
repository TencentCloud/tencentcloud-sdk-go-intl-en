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

package v20200527

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-05-27"

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


func NewCreateMediaPackageChannelRequest() (request *CreateMediaPackageChannelRequest) {
    request = &CreateMediaPackageChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "CreateMediaPackageChannel")
    return
}

func NewCreateMediaPackageChannelResponse() (response *CreateMediaPackageChannelResponse) {
    response = &CreateMediaPackageChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMediaPackageChannel
// This API is used to create a media package channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateMediaPackageChannel(request *CreateMediaPackageChannelRequest) (response *CreateMediaPackageChannelResponse, err error) {
    if request == nil {
        request = NewCreateMediaPackageChannelRequest()
    }
    response = NewCreateMediaPackageChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaPackageChannelEndpointRequest() (request *CreateMediaPackageChannelEndpointRequest) {
    request = &CreateMediaPackageChannelEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "CreateMediaPackageChannelEndpoint")
    return
}

func NewCreateMediaPackageChannelEndpointResponse() (response *CreateMediaPackageChannelEndpointResponse) {
    response = &CreateMediaPackageChannelEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMediaPackageChannelEndpoint
// This API is used to create an endpoint of a media package channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) CreateMediaPackageChannelEndpoint(request *CreateMediaPackageChannelEndpointRequest) (response *CreateMediaPackageChannelEndpointResponse, err error) {
    if request == nil {
        request = NewCreateMediaPackageChannelEndpointRequest()
    }
    response = NewCreateMediaPackageChannelEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaPackageChannelEndpointsRequest() (request *DeleteMediaPackageChannelEndpointsRequest) {
    request = &DeleteMediaPackageChannelEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteMediaPackageChannelEndpoints")
    return
}

func NewDeleteMediaPackageChannelEndpointsResponse() (response *DeleteMediaPackageChannelEndpointsResponse) {
    response = &DeleteMediaPackageChannelEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMediaPackageChannelEndpoints
// This API is used to delete endpoints from a media package channel in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteMediaPackageChannelEndpoints(request *DeleteMediaPackageChannelEndpointsRequest) (response *DeleteMediaPackageChannelEndpointsResponse, err error) {
    if request == nil {
        request = NewDeleteMediaPackageChannelEndpointsRequest()
    }
    response = NewDeleteMediaPackageChannelEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaPackageChannelsRequest() (request *DeleteMediaPackageChannelsRequest) {
    request = &DeleteMediaPackageChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteMediaPackageChannels")
    return
}

func NewDeleteMediaPackageChannelsResponse() (response *DeleteMediaPackageChannelsResponse) {
    response = &DeleteMediaPackageChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMediaPackageChannels
// This API is used to delete media package channels in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DeleteMediaPackageChannels(request *DeleteMediaPackageChannelsRequest) (response *DeleteMediaPackageChannelsResponse, err error) {
    if request == nil {
        request = NewDeleteMediaPackageChannelsRequest()
    }
    response = NewDeleteMediaPackageChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaPackageChannelRequest() (request *DescribeMediaPackageChannelRequest) {
    request = &DescribeMediaPackageChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeMediaPackageChannel")
    return
}

func NewDescribeMediaPackageChannelResponse() (response *DescribeMediaPackageChannelResponse) {
    response = &DescribeMediaPackageChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMediaPackageChannel
// This API is used to query the information of a media package channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeMediaPackageChannel(request *DescribeMediaPackageChannelRequest) (response *DescribeMediaPackageChannelResponse, err error) {
    if request == nil {
        request = NewDescribeMediaPackageChannelRequest()
    }
    response = NewDescribeMediaPackageChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaPackageChannelsRequest() (request *DescribeMediaPackageChannelsRequest) {
    request = &DescribeMediaPackageChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeMediaPackageChannels")
    return
}

func NewDescribeMediaPackageChannelsResponse() (response *DescribeMediaPackageChannelsResponse) {
    response = &DescribeMediaPackageChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMediaPackageChannels
// This API is used to query the information list of media package channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeMediaPackageChannels(request *DescribeMediaPackageChannelsRequest) (response *DescribeMediaPackageChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeMediaPackageChannelsRequest()
    }
    response = NewDescribeMediaPackageChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaPackageChannelRequest() (request *ModifyMediaPackageChannelRequest) {
    request = &ModifyMediaPackageChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyMediaPackageChannel")
    return
}

func NewModifyMediaPackageChannelResponse() (response *ModifyMediaPackageChannelResponse) {
    response = &ModifyMediaPackageChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMediaPackageChannel
// This API is used to modify the information of a media package channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyMediaPackageChannel(request *ModifyMediaPackageChannelRequest) (response *ModifyMediaPackageChannelResponse, err error) {
    if request == nil {
        request = NewModifyMediaPackageChannelRequest()
    }
    response = NewModifyMediaPackageChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaPackageChannelEndpointRequest() (request *ModifyMediaPackageChannelEndpointRequest) {
    request = &ModifyMediaPackageChannelEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyMediaPackageChannelEndpoint")
    return
}

func NewModifyMediaPackageChannelEndpointResponse() (response *ModifyMediaPackageChannelEndpointResponse) {
    response = &ModifyMediaPackageChannelEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMediaPackageChannelEndpoint
// This API is used to modify an endpoint of a media package channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyMediaPackageChannelEndpoint(request *ModifyMediaPackageChannelEndpointRequest) (response *ModifyMediaPackageChannelEndpointResponse, err error) {
    if request == nil {
        request = NewModifyMediaPackageChannelEndpointRequest()
    }
    response = NewModifyMediaPackageChannelEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaPackageChannelInputAuthInfoRequest() (request *ModifyMediaPackageChannelInputAuthInfoRequest) {
    request = &ModifyMediaPackageChannelInputAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyMediaPackageChannelInputAuthInfo")
    return
}

func NewModifyMediaPackageChannelInputAuthInfoResponse() (response *ModifyMediaPackageChannelInputAuthInfoResponse) {
    response = &ModifyMediaPackageChannelInputAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMediaPackageChannelInputAuthInfo
// This API is used to modify the input authentication information of a media package channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONTYPE = "InvalidParameter.ActionType"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyMediaPackageChannelInputAuthInfo(request *ModifyMediaPackageChannelInputAuthInfoRequest) (response *ModifyMediaPackageChannelInputAuthInfoResponse, err error) {
    if request == nil {
        request = NewModifyMediaPackageChannelInputAuthInfoRequest()
    }
    response = NewModifyMediaPackageChannelInputAuthInfoResponse()
    err = c.Send(request, response)
    return
}
