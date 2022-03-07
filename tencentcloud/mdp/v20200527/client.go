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
    "context"
    "errors"
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


func NewBindNewLVBDomainWithChannelRequest() (request *BindNewLVBDomainWithChannelRequest) {
    request = &BindNewLVBDomainWithChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "BindNewLVBDomainWithChannel")
    
    
    return
}

func NewBindNewLVBDomainWithChannelResponse() (response *BindNewLVBDomainWithChannelResponse) {
    response = &BindNewLVBDomainWithChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindNewLVBDomainWithChannel
// This API is used to bind an LVB domain name to a channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
//  INVALIDPARAMETER_LVBDOMAIN = "InvalidParameter.LvbDomain"
func (c *Client) BindNewLVBDomainWithChannel(request *BindNewLVBDomainWithChannelRequest) (response *BindNewLVBDomainWithChannelResponse, err error) {
    return c.BindNewLVBDomainWithChannelWithContext(context.Background(), request)
}

// BindNewLVBDomainWithChannel
// This API is used to bind an LVB domain name to a channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
//  INVALIDPARAMETER_LVBDOMAIN = "InvalidParameter.LvbDomain"
func (c *Client) BindNewLVBDomainWithChannelWithContext(ctx context.Context, request *BindNewLVBDomainWithChannelRequest) (response *BindNewLVBDomainWithChannelResponse, err error) {
    if request == nil {
        request = NewBindNewLVBDomainWithChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindNewLVBDomainWithChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindNewLVBDomainWithChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageChannelRequest() (request *CreateStreamPackageChannelRequest) {
    request = &CreateStreamPackageChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageChannel")
    
    
    return
}

func NewCreateStreamPackageChannelResponse() (response *CreateStreamPackageChannelResponse) {
    response = &CreateStreamPackageChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamPackageChannel
// This API is used to create a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageChannel(request *CreateStreamPackageChannelRequest) (response *CreateStreamPackageChannelResponse, err error) {
    return c.CreateStreamPackageChannelWithContext(context.Background(), request)
}

// CreateStreamPackageChannel
// This API is used to create a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageChannelWithContext(ctx context.Context, request *CreateStreamPackageChannelRequest) (response *CreateStreamPackageChannelResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageChannelEndpointRequest() (request *CreateStreamPackageChannelEndpointRequest) {
    request = &CreateStreamPackageChannelEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageChannelEndpoint")
    
    
    return
}

func NewCreateStreamPackageChannelEndpointResponse() (response *CreateStreamPackageChannelEndpointResponse) {
    response = &CreateStreamPackageChannelEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamPackageChannelEndpoint
// This API is used to create an endpoint on a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) CreateStreamPackageChannelEndpoint(request *CreateStreamPackageChannelEndpointRequest) (response *CreateStreamPackageChannelEndpointResponse, err error) {
    return c.CreateStreamPackageChannelEndpointWithContext(context.Background(), request)
}

// CreateStreamPackageChannelEndpoint
// This API is used to create an endpoint on a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) CreateStreamPackageChannelEndpointWithContext(ctx context.Context, request *CreateStreamPackageChannelEndpointRequest) (response *CreateStreamPackageChannelEndpointResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageChannelEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageChannelEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageChannelEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageChannelEndpointsRequest() (request *DeleteStreamPackageChannelEndpointsRequest) {
    request = &DeleteStreamPackageChannelEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageChannelEndpoints")
    
    
    return
}

func NewDeleteStreamPackageChannelEndpointsResponse() (response *DeleteStreamPackageChannelEndpointsResponse) {
    response = &DeleteStreamPackageChannelEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamPackageChannelEndpoints
// This API is used to delete endpoints from a StreamPackage channel in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageChannelEndpoints(request *DeleteStreamPackageChannelEndpointsRequest) (response *DeleteStreamPackageChannelEndpointsResponse, err error) {
    return c.DeleteStreamPackageChannelEndpointsWithContext(context.Background(), request)
}

// DeleteStreamPackageChannelEndpoints
// This API is used to delete endpoints from a StreamPackage channel in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageChannelEndpointsWithContext(ctx context.Context, request *DeleteStreamPackageChannelEndpointsRequest) (response *DeleteStreamPackageChannelEndpointsResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageChannelEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageChannelEndpoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageChannelEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageChannelsRequest() (request *DeleteStreamPackageChannelsRequest) {
    request = &DeleteStreamPackageChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageChannels")
    
    
    return
}

func NewDeleteStreamPackageChannelsResponse() (response *DeleteStreamPackageChannelsResponse) {
    response = &DeleteStreamPackageChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamPackageChannels
// This API is used to delete StreamPackage channels in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DeleteStreamPackageChannels(request *DeleteStreamPackageChannelsRequest) (response *DeleteStreamPackageChannelsResponse, err error) {
    return c.DeleteStreamPackageChannelsWithContext(context.Background(), request)
}

// DeleteStreamPackageChannels
// This API is used to delete StreamPackage channels in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DeleteStreamPackageChannelsWithContext(ctx context.Context, request *DeleteStreamPackageChannelsRequest) (response *DeleteStreamPackageChannelsResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageChannelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageChannelRequest() (request *DescribeStreamPackageChannelRequest) {
    request = &DescribeStreamPackageChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageChannel")
    
    
    return
}

func NewDescribeStreamPackageChannelResponse() (response *DescribeStreamPackageChannelResponse) {
    response = &DescribeStreamPackageChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamPackageChannel
// This API is used to query the information of a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageChannel(request *DescribeStreamPackageChannelRequest) (response *DescribeStreamPackageChannelResponse, err error) {
    return c.DescribeStreamPackageChannelWithContext(context.Background(), request)
}

// DescribeStreamPackageChannel
// This API is used to query the information of a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageChannelWithContext(ctx context.Context, request *DescribeStreamPackageChannelRequest) (response *DescribeStreamPackageChannelResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageChannelsRequest() (request *DescribeStreamPackageChannelsRequest) {
    request = &DescribeStreamPackageChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageChannels")
    
    
    return
}

func NewDescribeStreamPackageChannelsResponse() (response *DescribeStreamPackageChannelsResponse) {
    response = &DescribeStreamPackageChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamPackageChannels
// This API is used to query the information of multiple StreamPackage channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageChannels(request *DescribeStreamPackageChannelsRequest) (response *DescribeStreamPackageChannelsResponse, err error) {
    return c.DescribeStreamPackageChannelsWithContext(context.Background(), request)
}

// DescribeStreamPackageChannels
// This API is used to query the information of multiple StreamPackage channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageChannelsWithContext(ctx context.Context, request *DescribeStreamPackageChannelsRequest) (response *DescribeStreamPackageChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageChannelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageChannelRequest() (request *ModifyStreamPackageChannelRequest) {
    request = &ModifyStreamPackageChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageChannel")
    
    
    return
}

func NewModifyStreamPackageChannelResponse() (response *ModifyStreamPackageChannelResponse) {
    response = &ModifyStreamPackageChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamPackageChannel
// This API is used to modify a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageChannel(request *ModifyStreamPackageChannelRequest) (response *ModifyStreamPackageChannelResponse, err error) {
    return c.ModifyStreamPackageChannelWithContext(context.Background(), request)
}

// ModifyStreamPackageChannel
// This API is used to modify a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageChannelWithContext(ctx context.Context, request *ModifyStreamPackageChannelRequest) (response *ModifyStreamPackageChannelResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageChannelEndpointRequest() (request *ModifyStreamPackageChannelEndpointRequest) {
    request = &ModifyStreamPackageChannelEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageChannelEndpoint")
    
    
    return
}

func NewModifyStreamPackageChannelEndpointResponse() (response *ModifyStreamPackageChannelEndpointResponse) {
    response = &ModifyStreamPackageChannelEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamPackageChannelEndpoint
// This API is used to modify an endpoint of a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyStreamPackageChannelEndpoint(request *ModifyStreamPackageChannelEndpointRequest) (response *ModifyStreamPackageChannelEndpointResponse, err error) {
    return c.ModifyStreamPackageChannelEndpointWithContext(context.Background(), request)
}

// ModifyStreamPackageChannelEndpoint
// This API is used to modify an endpoint of a StreamPackage channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyStreamPackageChannelEndpointWithContext(ctx context.Context, request *ModifyStreamPackageChannelEndpointRequest) (response *ModifyStreamPackageChannelEndpointResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageChannelEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageChannelEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageChannelEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageChannelInputAuthInfoRequest() (request *ModifyStreamPackageChannelInputAuthInfoRequest) {
    request = &ModifyStreamPackageChannelInputAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageChannelInputAuthInfo")
    
    
    return
}

func NewModifyStreamPackageChannelInputAuthInfoResponse() (response *ModifyStreamPackageChannelInputAuthInfoResponse) {
    response = &ModifyStreamPackageChannelInputAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamPackageChannelInputAuthInfo
// This API is used to modify the input authentication information of a StreamPackage channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONTYPE = "InvalidParameter.ActionType"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyStreamPackageChannelInputAuthInfo(request *ModifyStreamPackageChannelInputAuthInfoRequest) (response *ModifyStreamPackageChannelInputAuthInfoResponse, err error) {
    return c.ModifyStreamPackageChannelInputAuthInfoWithContext(context.Background(), request)
}

// ModifyStreamPackageChannelInputAuthInfo
// This API is used to modify the input authentication information of a StreamPackage channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONTYPE = "InvalidParameter.ActionType"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyStreamPackageChannelInputAuthInfoWithContext(ctx context.Context, request *ModifyStreamPackageChannelInputAuthInfoRequest) (response *ModifyStreamPackageChannelInputAuthInfoResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageChannelInputAuthInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageChannelInputAuthInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageChannelInputAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindCdnDomainWithChannelRequest() (request *UnbindCdnDomainWithChannelRequest) {
    request = &UnbindCdnDomainWithChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdp", APIVersion, "UnbindCdnDomainWithChannel")
    
    
    return
}

func NewUnbindCdnDomainWithChannelResponse() (response *UnbindCdnDomainWithChannelResponse) {
    response = &UnbindCdnDomainWithChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindCdnDomainWithChannel
// This API is used to unbind a CDN playback domain name from a channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CDNDOMAIN = "InvalidParameter.CdnDomain"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) UnbindCdnDomainWithChannel(request *UnbindCdnDomainWithChannelRequest) (response *UnbindCdnDomainWithChannelResponse, err error) {
    return c.UnbindCdnDomainWithChannelWithContext(context.Background(), request)
}

// UnbindCdnDomainWithChannel
// This API is used to unbind a CDN playback domain name from a channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CDNDOMAIN = "InvalidParameter.CdnDomain"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) UnbindCdnDomainWithChannelWithContext(ctx context.Context, request *UnbindCdnDomainWithChannelRequest) (response *UnbindCdnDomainWithChannelResponse, err error) {
    if request == nil {
        request = NewUnbindCdnDomainWithChannelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindCdnDomainWithChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindCdnDomainWithChannelResponse()
    err = c.Send(request, response)
    return
}
