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


func NewBindLinearAssemblyCDNDomainWithChannelRequest() (request *BindLinearAssemblyCDNDomainWithChannelRequest) {
    request = &BindLinearAssemblyCDNDomainWithChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "BindLinearAssemblyCDNDomainWithChannel")
    
    
    return
}

func NewBindLinearAssemblyCDNDomainWithChannelResponse() (response *BindLinearAssemblyCDNDomainWithChannelResponse) {
    response = &BindLinearAssemblyCDNDomainWithChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindLinearAssemblyCDNDomainWithChannel
// Linear Assembly channel is bound to CDN playback domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CDNDOMAIN = "InvalidParameter.CdnDomain"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
func (c *Client) BindLinearAssemblyCDNDomainWithChannel(request *BindLinearAssemblyCDNDomainWithChannelRequest) (response *BindLinearAssemblyCDNDomainWithChannelResponse, err error) {
    return c.BindLinearAssemblyCDNDomainWithChannelWithContext(context.Background(), request)
}

// BindLinearAssemblyCDNDomainWithChannel
// Linear Assembly channel is bound to CDN playback domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CDNDOMAIN = "InvalidParameter.CdnDomain"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
func (c *Client) BindLinearAssemblyCDNDomainWithChannelWithContext(ctx context.Context, request *BindLinearAssemblyCDNDomainWithChannelRequest) (response *BindLinearAssemblyCDNDomainWithChannelResponse, err error) {
    if request == nil {
        request = NewBindLinearAssemblyCDNDomainWithChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "BindLinearAssemblyCDNDomainWithChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindLinearAssemblyCDNDomainWithChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindLinearAssemblyCDNDomainWithChannelResponse()
    err = c.Send(request, response)
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "BindNewLVBDomainWithChannel")
    
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageChannel")
    
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
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
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
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) CreateStreamPackageChannelEndpointWithContext(ctx context.Context, request *CreateStreamPackageChannelEndpointRequest) (response *CreateStreamPackageChannelEndpointResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageChannelEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageChannelEndpoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageChannelEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageChannelEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageHarvestJobRequest() (request *CreateStreamPackageHarvestJobRequest) {
    request = &CreateStreamPackageHarvestJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageHarvestJob")
    
    
    return
}

func NewCreateStreamPackageHarvestJobResponse() (response *CreateStreamPackageHarvestJobResponse) {
    response = &CreateStreamPackageHarvestJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamPackageHarvestJob
// Create HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageHarvestJob(request *CreateStreamPackageHarvestJobRequest) (response *CreateStreamPackageHarvestJobResponse, err error) {
    return c.CreateStreamPackageHarvestJobWithContext(context.Background(), request)
}

// CreateStreamPackageHarvestJob
// Create HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageHarvestJobWithContext(ctx context.Context, request *CreateStreamPackageHarvestJobRequest) (response *CreateStreamPackageHarvestJobResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageHarvestJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageHarvestJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageHarvestJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageHarvestJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageLinearAssemblyChannelRequest() (request *CreateStreamPackageLinearAssemblyChannelRequest) {
    request = &CreateStreamPackageLinearAssemblyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageLinearAssemblyChannel")
    
    
    return
}

func NewCreateStreamPackageLinearAssemblyChannelResponse() (response *CreateStreamPackageLinearAssemblyChannelResponse) {
    response = &CreateStreamPackageLinearAssemblyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamPackageLinearAssemblyChannel
// Create a linear assembly channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageLinearAssemblyChannel(request *CreateStreamPackageLinearAssemblyChannelRequest) (response *CreateStreamPackageLinearAssemblyChannelResponse, err error) {
    return c.CreateStreamPackageLinearAssemblyChannelWithContext(context.Background(), request)
}

// CreateStreamPackageLinearAssemblyChannel
// Create a linear assembly channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageLinearAssemblyChannelWithContext(ctx context.Context, request *CreateStreamPackageLinearAssemblyChannelRequest) (response *CreateStreamPackageLinearAssemblyChannelResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageLinearAssemblyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageLinearAssemblyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageLinearAssemblyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageLinearAssemblyChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageLinearAssemblyProgramRequest() (request *CreateStreamPackageLinearAssemblyProgramRequest) {
    request = &CreateStreamPackageLinearAssemblyProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageLinearAssemblyProgram")
    
    
    return
}

func NewCreateStreamPackageLinearAssemblyProgramResponse() (response *CreateStreamPackageLinearAssemblyProgramResponse) {
    response = &CreateStreamPackageLinearAssemblyProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamPackageLinearAssemblyProgram
// Create a linear assembly program.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageLinearAssemblyProgram(request *CreateStreamPackageLinearAssemblyProgramRequest) (response *CreateStreamPackageLinearAssemblyProgramResponse, err error) {
    return c.CreateStreamPackageLinearAssemblyProgramWithContext(context.Background(), request)
}

// CreateStreamPackageLinearAssemblyProgram
// Create a linear assembly program.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageLinearAssemblyProgramWithContext(ctx context.Context, request *CreateStreamPackageLinearAssemblyProgramRequest) (response *CreateStreamPackageLinearAssemblyProgramResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageLinearAssemblyProgramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageLinearAssemblyProgram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageLinearAssemblyProgram require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageLinearAssemblyProgramResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageSSAIChannelRequest() (request *CreateStreamPackageSSAIChannelRequest) {
    request = &CreateStreamPackageSSAIChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageSSAIChannel")
    
    
    return
}

func NewCreateStreamPackageSSAIChannelResponse() (response *CreateStreamPackageSSAIChannelResponse) {
    response = &CreateStreamPackageSSAIChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamPackageSSAIChannel
// CreateStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) CreateStreamPackageSSAIChannel(request *CreateStreamPackageSSAIChannelRequest) (response *CreateStreamPackageSSAIChannelResponse, err error) {
    return c.CreateStreamPackageSSAIChannelWithContext(context.Background(), request)
}

// CreateStreamPackageSSAIChannel
// CreateStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) CreateStreamPackageSSAIChannelWithContext(ctx context.Context, request *CreateStreamPackageSSAIChannelRequest) (response *CreateStreamPackageSSAIChannelResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageSSAIChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageSSAIChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageSSAIChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageSSAIChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageSourceRequest() (request *CreateStreamPackageSourceRequest) {
    request = &CreateStreamPackageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageSource")
    
    
    return
}

func NewCreateStreamPackageSourceResponse() (response *CreateStreamPackageSourceResponse) {
    response = &CreateStreamPackageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamPackageSource
// Create channel linear assembly Source.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageSource(request *CreateStreamPackageSourceRequest) (response *CreateStreamPackageSourceResponse, err error) {
    return c.CreateStreamPackageSourceWithContext(context.Background(), request)
}

// CreateStreamPackageSource
// Create channel linear assembly Source.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageSourceWithContext(ctx context.Context, request *CreateStreamPackageSourceRequest) (response *CreateStreamPackageSourceResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamPackageSourceLocationRequest() (request *CreateStreamPackageSourceLocationRequest) {
    request = &CreateStreamPackageSourceLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "CreateStreamPackageSourceLocation")
    
    
    return
}

func NewCreateStreamPackageSourceLocationResponse() (response *CreateStreamPackageSourceLocationResponse) {
    response = &CreateStreamPackageSourceLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStreamPackageSourceLocation
// Create Linear Assembly SourceLocation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageSourceLocation(request *CreateStreamPackageSourceLocationRequest) (response *CreateStreamPackageSourceLocationResponse, err error) {
    return c.CreateStreamPackageSourceLocationWithContext(context.Background(), request)
}

// CreateStreamPackageSourceLocation
// Create Linear Assembly SourceLocation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) CreateStreamPackageSourceLocationWithContext(ctx context.Context, request *CreateStreamPackageSourceLocationRequest) (response *CreateStreamPackageSourceLocationResponse, err error) {
    if request == nil {
        request = NewCreateStreamPackageSourceLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "CreateStreamPackageSourceLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStreamPackageSourceLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStreamPackageSourceLocationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageChannelEndpoints")
    
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageHarvestJobRequest() (request *DeleteStreamPackageHarvestJobRequest) {
    request = &DeleteStreamPackageHarvestJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageHarvestJob")
    
    
    return
}

func NewDeleteStreamPackageHarvestJobResponse() (response *DeleteStreamPackageHarvestJobResponse) {
    response = &DeleteStreamPackageHarvestJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageHarvestJob
// Delete HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DeleteStreamPackageHarvestJob(request *DeleteStreamPackageHarvestJobRequest) (response *DeleteStreamPackageHarvestJobResponse, err error) {
    return c.DeleteStreamPackageHarvestJobWithContext(context.Background(), request)
}

// DeleteStreamPackageHarvestJob
// Delete HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DeleteStreamPackageHarvestJobWithContext(ctx context.Context, request *DeleteStreamPackageHarvestJobRequest) (response *DeleteStreamPackageHarvestJobResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageHarvestJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageHarvestJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageHarvestJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageHarvestJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageHarvestJobsRequest() (request *DeleteStreamPackageHarvestJobsRequest) {
    request = &DeleteStreamPackageHarvestJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageHarvestJobs")
    
    
    return
}

func NewDeleteStreamPackageHarvestJobsResponse() (response *DeleteStreamPackageHarvestJobsResponse) {
    response = &DeleteStreamPackageHarvestJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageHarvestJobs
// Deleting HarvestJobs in Batch.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DeleteStreamPackageHarvestJobs(request *DeleteStreamPackageHarvestJobsRequest) (response *DeleteStreamPackageHarvestJobsResponse, err error) {
    return c.DeleteStreamPackageHarvestJobsWithContext(context.Background(), request)
}

// DeleteStreamPackageHarvestJobs
// Deleting HarvestJobs in Batch.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DeleteStreamPackageHarvestJobsWithContext(ctx context.Context, request *DeleteStreamPackageHarvestJobsRequest) (response *DeleteStreamPackageHarvestJobsResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageHarvestJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageHarvestJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageHarvestJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageHarvestJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageLinearAssemblyChannelRequest() (request *DeleteStreamPackageLinearAssemblyChannelRequest) {
    request = &DeleteStreamPackageLinearAssemblyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageLinearAssemblyChannel")
    
    
    return
}

func NewDeleteStreamPackageLinearAssemblyChannelResponse() (response *DeleteStreamPackageLinearAssemblyChannelResponse) {
    response = &DeleteStreamPackageLinearAssemblyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageLinearAssemblyChannel
// Delete channel linear assemblyChannel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyChannel(request *DeleteStreamPackageLinearAssemblyChannelRequest) (response *DeleteStreamPackageLinearAssemblyChannelResponse, err error) {
    return c.DeleteStreamPackageLinearAssemblyChannelWithContext(context.Background(), request)
}

// DeleteStreamPackageLinearAssemblyChannel
// Delete channel linear assemblyChannel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyChannelWithContext(ctx context.Context, request *DeleteStreamPackageLinearAssemblyChannelRequest) (response *DeleteStreamPackageLinearAssemblyChannelResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageLinearAssemblyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageLinearAssemblyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageLinearAssemblyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageLinearAssemblyChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageLinearAssemblyChannelsRequest() (request *DeleteStreamPackageLinearAssemblyChannelsRequest) {
    request = &DeleteStreamPackageLinearAssemblyChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageLinearAssemblyChannels")
    
    
    return
}

func NewDeleteStreamPackageLinearAssemblyChannelsResponse() (response *DeleteStreamPackageLinearAssemblyChannelsResponse) {
    response = &DeleteStreamPackageLinearAssemblyChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageLinearAssemblyChannels
// Delete channels in batches and linearly assemble channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyChannels(request *DeleteStreamPackageLinearAssemblyChannelsRequest) (response *DeleteStreamPackageLinearAssemblyChannelsResponse, err error) {
    return c.DeleteStreamPackageLinearAssemblyChannelsWithContext(context.Background(), request)
}

// DeleteStreamPackageLinearAssemblyChannels
// Delete channels in batches and linearly assemble channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyChannelsWithContext(ctx context.Context, request *DeleteStreamPackageLinearAssemblyChannelsRequest) (response *DeleteStreamPackageLinearAssemblyChannelsResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageLinearAssemblyChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageLinearAssemblyChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageLinearAssemblyChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageLinearAssemblyChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageLinearAssemblyProgramRequest() (request *DeleteStreamPackageLinearAssemblyProgramRequest) {
    request = &DeleteStreamPackageLinearAssemblyProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageLinearAssemblyProgram")
    
    
    return
}

func NewDeleteStreamPackageLinearAssemblyProgramResponse() (response *DeleteStreamPackageLinearAssemblyProgramResponse) {
    response = &DeleteStreamPackageLinearAssemblyProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageLinearAssemblyProgram
// Delete Channel Linear Assembly Program.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyProgram(request *DeleteStreamPackageLinearAssemblyProgramRequest) (response *DeleteStreamPackageLinearAssemblyProgramResponse, err error) {
    return c.DeleteStreamPackageLinearAssemblyProgramWithContext(context.Background(), request)
}

// DeleteStreamPackageLinearAssemblyProgram
// Delete Channel Linear Assembly Program.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyProgramWithContext(ctx context.Context, request *DeleteStreamPackageLinearAssemblyProgramRequest) (response *DeleteStreamPackageLinearAssemblyProgramResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageLinearAssemblyProgramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageLinearAssemblyProgram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageLinearAssemblyProgram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageLinearAssemblyProgramResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageLinearAssemblyProgramsRequest() (request *DeleteStreamPackageLinearAssemblyProgramsRequest) {
    request = &DeleteStreamPackageLinearAssemblyProgramsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageLinearAssemblyPrograms")
    
    
    return
}

func NewDeleteStreamPackageLinearAssemblyProgramsResponse() (response *DeleteStreamPackageLinearAssemblyProgramsResponse) {
    response = &DeleteStreamPackageLinearAssemblyProgramsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageLinearAssemblyPrograms
// Batch deletion of channels linear assembly program.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyPrograms(request *DeleteStreamPackageLinearAssemblyProgramsRequest) (response *DeleteStreamPackageLinearAssemblyProgramsResponse, err error) {
    return c.DeleteStreamPackageLinearAssemblyProgramsWithContext(context.Background(), request)
}

// DeleteStreamPackageLinearAssemblyPrograms
// Batch deletion of channels linear assembly program.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageLinearAssemblyProgramsWithContext(ctx context.Context, request *DeleteStreamPackageLinearAssemblyProgramsRequest) (response *DeleteStreamPackageLinearAssemblyProgramsResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageLinearAssemblyProgramsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageLinearAssemblyPrograms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageLinearAssemblyPrograms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageLinearAssemblyProgramsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageSSAIChannelRequest() (request *DeleteStreamPackageSSAIChannelRequest) {
    request = &DeleteStreamPackageSSAIChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageSSAIChannel")
    
    
    return
}

func NewDeleteStreamPackageSSAIChannelResponse() (response *DeleteStreamPackageSSAIChannelResponse) {
    response = &DeleteStreamPackageSSAIChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageSSAIChannel
// DeleteStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) DeleteStreamPackageSSAIChannel(request *DeleteStreamPackageSSAIChannelRequest) (response *DeleteStreamPackageSSAIChannelResponse, err error) {
    return c.DeleteStreamPackageSSAIChannelWithContext(context.Background(), request)
}

// DeleteStreamPackageSSAIChannel
// DeleteStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) DeleteStreamPackageSSAIChannelWithContext(ctx context.Context, request *DeleteStreamPackageSSAIChannelRequest) (response *DeleteStreamPackageSSAIChannelResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageSSAIChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageSSAIChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageSSAIChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageSSAIChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageSourceRequest() (request *DeleteStreamPackageSourceRequest) {
    request = &DeleteStreamPackageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageSource")
    
    
    return
}

func NewDeleteStreamPackageSourceResponse() (response *DeleteStreamPackageSourceResponse) {
    response = &DeleteStreamPackageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageSource
// Delete channel linear assembly Source.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageSource(request *DeleteStreamPackageSourceRequest) (response *DeleteStreamPackageSourceResponse, err error) {
    return c.DeleteStreamPackageSourceWithContext(context.Background(), request)
}

// DeleteStreamPackageSource
// Delete channel linear assembly Source.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageSourceWithContext(ctx context.Context, request *DeleteStreamPackageSourceRequest) (response *DeleteStreamPackageSourceResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamPackageSourceLocationRequest() (request *DeleteStreamPackageSourceLocationRequest) {
    request = &DeleteStreamPackageSourceLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DeleteStreamPackageSourceLocation")
    
    
    return
}

func NewDeleteStreamPackageSourceLocationResponse() (response *DeleteStreamPackageSourceLocationResponse) {
    response = &DeleteStreamPackageSourceLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStreamPackageSourceLocation
// Batch delete media packaging SourceLocation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageSourceLocation(request *DeleteStreamPackageSourceLocationRequest) (response *DeleteStreamPackageSourceLocationResponse, err error) {
    return c.DeleteStreamPackageSourceLocationWithContext(context.Background(), request)
}

// DeleteStreamPackageSourceLocation
// Batch delete media packaging SourceLocation.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DeleteStreamPackageSourceLocationWithContext(ctx context.Context, request *DeleteStreamPackageSourceLocationRequest) (response *DeleteStreamPackageSourceLocationResponse, err error) {
    if request == nil {
        request = NewDeleteStreamPackageSourceLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DeleteStreamPackageSourceLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStreamPackageSourceLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStreamPackageSourceLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLinearAssemblyCDNDomainWithChannelRequest() (request *DescribeLinearAssemblyCDNDomainWithChannelRequest) {
    request = &DescribeLinearAssemblyCDNDomainWithChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeLinearAssemblyCDNDomainWithChannel")
    
    
    return
}

func NewDescribeLinearAssemblyCDNDomainWithChannelResponse() (response *DescribeLinearAssemblyCDNDomainWithChannelResponse) {
    response = &DescribeLinearAssemblyCDNDomainWithChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLinearAssemblyCDNDomainWithChannel
// Query the CDN domain name associated with the LinearAssembly channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
func (c *Client) DescribeLinearAssemblyCDNDomainWithChannel(request *DescribeLinearAssemblyCDNDomainWithChannelRequest) (response *DescribeLinearAssemblyCDNDomainWithChannelResponse, err error) {
    return c.DescribeLinearAssemblyCDNDomainWithChannelWithContext(context.Background(), request)
}

// DescribeLinearAssemblyCDNDomainWithChannel
// Query the CDN domain name associated with the LinearAssembly channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
func (c *Client) DescribeLinearAssemblyCDNDomainWithChannelWithContext(ctx context.Context, request *DescribeLinearAssemblyCDNDomainWithChannelRequest) (response *DescribeLinearAssemblyCDNDomainWithChannelResponse, err error) {
    if request == nil {
        request = NewDescribeLinearAssemblyCDNDomainWithChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeLinearAssemblyCDNDomainWithChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLinearAssemblyCDNDomainWithChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLinearAssemblyCDNDomainWithChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLinearAssemblyCDNDomainWithChannelsRequest() (request *DescribeLinearAssemblyCDNDomainWithChannelsRequest) {
    request = &DescribeLinearAssemblyCDNDomainWithChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeLinearAssemblyCDNDomainWithChannels")
    
    
    return
}

func NewDescribeLinearAssemblyCDNDomainWithChannelsResponse() (response *DescribeLinearAssemblyCDNDomainWithChannelsResponse) {
    response = &DescribeLinearAssemblyCDNDomainWithChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLinearAssemblyCDNDomainWithChannels
// Query the CDN domain names associated with all LinearAssembly channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLinearAssemblyCDNDomainWithChannels(request *DescribeLinearAssemblyCDNDomainWithChannelsRequest) (response *DescribeLinearAssemblyCDNDomainWithChannelsResponse, err error) {
    return c.DescribeLinearAssemblyCDNDomainWithChannelsWithContext(context.Background(), request)
}

// DescribeLinearAssemblyCDNDomainWithChannels
// Query the CDN domain names associated with all LinearAssembly channels.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeLinearAssemblyCDNDomainWithChannelsWithContext(ctx context.Context, request *DescribeLinearAssemblyCDNDomainWithChannelsRequest) (response *DescribeLinearAssemblyCDNDomainWithChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeLinearAssemblyCDNDomainWithChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeLinearAssemblyCDNDomainWithChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLinearAssemblyCDNDomainWithChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLinearAssemblyCDNDomainWithChannelsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageChannel")
    
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageHarvestJobRequest() (request *DescribeStreamPackageHarvestJobRequest) {
    request = &DescribeStreamPackageHarvestJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageHarvestJob")
    
    
    return
}

func NewDescribeStreamPackageHarvestJobResponse() (response *DescribeStreamPackageHarvestJobResponse) {
    response = &DescribeStreamPackageHarvestJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageHarvestJob
// Query HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DescribeStreamPackageHarvestJob(request *DescribeStreamPackageHarvestJobRequest) (response *DescribeStreamPackageHarvestJobResponse, err error) {
    return c.DescribeStreamPackageHarvestJobWithContext(context.Background(), request)
}

// DescribeStreamPackageHarvestJob
// Query HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DescribeStreamPackageHarvestJobWithContext(ctx context.Context, request *DescribeStreamPackageHarvestJobRequest) (response *DescribeStreamPackageHarvestJobResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageHarvestJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageHarvestJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageHarvestJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageHarvestJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageHarvestJobsRequest() (request *DescribeStreamPackageHarvestJobsRequest) {
    request = &DescribeStreamPackageHarvestJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageHarvestJobs")
    
    
    return
}

func NewDescribeStreamPackageHarvestJobsResponse() (response *DescribeStreamPackageHarvestJobsResponse) {
    response = &DescribeStreamPackageHarvestJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageHarvestJobs
// Batch query HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DescribeStreamPackageHarvestJobs(request *DescribeStreamPackageHarvestJobsRequest) (response *DescribeStreamPackageHarvestJobsResponse, err error) {
    return c.DescribeStreamPackageHarvestJobsWithContext(context.Background(), request)
}

// DescribeStreamPackageHarvestJobs
// Batch query HarvestJob.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) DescribeStreamPackageHarvestJobsWithContext(ctx context.Context, request *DescribeStreamPackageHarvestJobsRequest) (response *DescribeStreamPackageHarvestJobsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageHarvestJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageHarvestJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageHarvestJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageHarvestJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageLinearAssemblyChannelRequest() (request *DescribeStreamPackageLinearAssemblyChannelRequest) {
    request = &DescribeStreamPackageLinearAssemblyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageLinearAssemblyChannel")
    
    
    return
}

func NewDescribeStreamPackageLinearAssemblyChannelResponse() (response *DescribeStreamPackageLinearAssemblyChannelResponse) {
    response = &DescribeStreamPackageLinearAssemblyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageLinearAssemblyChannel
// Query channel linear assembly Channel information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageLinearAssemblyChannel(request *DescribeStreamPackageLinearAssemblyChannelRequest) (response *DescribeStreamPackageLinearAssemblyChannelResponse, err error) {
    return c.DescribeStreamPackageLinearAssemblyChannelWithContext(context.Background(), request)
}

// DescribeStreamPackageLinearAssemblyChannel
// Query channel linear assembly Channel information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageLinearAssemblyChannelWithContext(ctx context.Context, request *DescribeStreamPackageLinearAssemblyChannelRequest) (response *DescribeStreamPackageLinearAssemblyChannelResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageLinearAssemblyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageLinearAssemblyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageLinearAssemblyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageLinearAssemblyChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageLinearAssemblyChannelAlertsRequest() (request *DescribeStreamPackageLinearAssemblyChannelAlertsRequest) {
    request = &DescribeStreamPackageLinearAssemblyChannelAlertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageLinearAssemblyChannelAlerts")
    
    
    return
}

func NewDescribeStreamPackageLinearAssemblyChannelAlertsResponse() (response *DescribeStreamPackageLinearAssemblyChannelAlertsResponse) {
    response = &DescribeStreamPackageLinearAssemblyChannelAlertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageLinearAssemblyChannelAlerts
// Query linear assembly channel alarm information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DescribeStreamPackageLinearAssemblyChannelAlerts(request *DescribeStreamPackageLinearAssemblyChannelAlertsRequest) (response *DescribeStreamPackageLinearAssemblyChannelAlertsResponse, err error) {
    return c.DescribeStreamPackageLinearAssemblyChannelAlertsWithContext(context.Background(), request)
}

// DescribeStreamPackageLinearAssemblyChannelAlerts
// Query linear assembly channel alarm information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DescribeStreamPackageLinearAssemblyChannelAlertsWithContext(ctx context.Context, request *DescribeStreamPackageLinearAssemblyChannelAlertsRequest) (response *DescribeStreamPackageLinearAssemblyChannelAlertsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageLinearAssemblyChannelAlertsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageLinearAssemblyChannelAlerts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageLinearAssemblyChannelAlerts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageLinearAssemblyChannelAlertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageLinearAssemblyChannelsRequest() (request *DescribeStreamPackageLinearAssemblyChannelsRequest) {
    request = &DescribeStreamPackageLinearAssemblyChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageLinearAssemblyChannels")
    
    
    return
}

func NewDescribeStreamPackageLinearAssemblyChannelsResponse() (response *DescribeStreamPackageLinearAssemblyChannelsResponse) {
    response = &DescribeStreamPackageLinearAssemblyChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageLinearAssemblyChannels
// Query channel linear assembly Channel information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageLinearAssemblyChannels(request *DescribeStreamPackageLinearAssemblyChannelsRequest) (response *DescribeStreamPackageLinearAssemblyChannelsResponse, err error) {
    return c.DescribeStreamPackageLinearAssemblyChannelsWithContext(context.Background(), request)
}

// DescribeStreamPackageLinearAssemblyChannels
// Query channel linear assembly Channel information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageLinearAssemblyChannelsWithContext(ctx context.Context, request *DescribeStreamPackageLinearAssemblyChannelsRequest) (response *DescribeStreamPackageLinearAssemblyChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageLinearAssemblyChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageLinearAssemblyChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageLinearAssemblyChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageLinearAssemblyChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageLinearAssemblyProgramRequest() (request *DescribeStreamPackageLinearAssemblyProgramRequest) {
    request = &DescribeStreamPackageLinearAssemblyProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageLinearAssemblyProgram")
    
    
    return
}

func NewDescribeStreamPackageLinearAssemblyProgramResponse() (response *DescribeStreamPackageLinearAssemblyProgramResponse) {
    response = &DescribeStreamPackageLinearAssemblyProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageLinearAssemblyProgram
// Query channel linear assembly program information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageLinearAssemblyProgram(request *DescribeStreamPackageLinearAssemblyProgramRequest) (response *DescribeStreamPackageLinearAssemblyProgramResponse, err error) {
    return c.DescribeStreamPackageLinearAssemblyProgramWithContext(context.Background(), request)
}

// DescribeStreamPackageLinearAssemblyProgram
// Query channel linear assembly program information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageLinearAssemblyProgramWithContext(ctx context.Context, request *DescribeStreamPackageLinearAssemblyProgramRequest) (response *DescribeStreamPackageLinearAssemblyProgramResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageLinearAssemblyProgramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageLinearAssemblyProgram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageLinearAssemblyProgram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageLinearAssemblyProgramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageLinearAssemblyProgramSchedulesRequest() (request *DescribeStreamPackageLinearAssemblyProgramSchedulesRequest) {
    request = &DescribeStreamPackageLinearAssemblyProgramSchedulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageLinearAssemblyProgramSchedules")
    
    
    return
}

func NewDescribeStreamPackageLinearAssemblyProgramSchedulesResponse() (response *DescribeStreamPackageLinearAssemblyProgramSchedulesResponse) {
    response = &DescribeStreamPackageLinearAssemblyProgramSchedulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageLinearAssemblyProgramSchedules
// Query channel linear assembly Programl assembly scheduling information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageLinearAssemblyProgramSchedules(request *DescribeStreamPackageLinearAssemblyProgramSchedulesRequest) (response *DescribeStreamPackageLinearAssemblyProgramSchedulesResponse, err error) {
    return c.DescribeStreamPackageLinearAssemblyProgramSchedulesWithContext(context.Background(), request)
}

// DescribeStreamPackageLinearAssemblyProgramSchedules
// Query channel linear assembly Programl assembly scheduling information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageLinearAssemblyProgramSchedulesWithContext(ctx context.Context, request *DescribeStreamPackageLinearAssemblyProgramSchedulesRequest) (response *DescribeStreamPackageLinearAssemblyProgramSchedulesResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageLinearAssemblyProgramSchedulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageLinearAssemblyProgramSchedules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageLinearAssemblyProgramSchedules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageLinearAssemblyProgramSchedulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageLinearAssemblyProgramsRequest() (request *DescribeStreamPackageLinearAssemblyProgramsRequest) {
    request = &DescribeStreamPackageLinearAssemblyProgramsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageLinearAssemblyPrograms")
    
    
    return
}

func NewDescribeStreamPackageLinearAssemblyProgramsResponse() (response *DescribeStreamPackageLinearAssemblyProgramsResponse) {
    response = &DescribeStreamPackageLinearAssemblyProgramsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageLinearAssemblyPrograms
// Query channel linear assembly Programl information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageLinearAssemblyPrograms(request *DescribeStreamPackageLinearAssemblyProgramsRequest) (response *DescribeStreamPackageLinearAssemblyProgramsResponse, err error) {
    return c.DescribeStreamPackageLinearAssemblyProgramsWithContext(context.Background(), request)
}

// DescribeStreamPackageLinearAssemblyPrograms
// Query channel linear assembly Programl information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageLinearAssemblyProgramsWithContext(ctx context.Context, request *DescribeStreamPackageLinearAssemblyProgramsRequest) (response *DescribeStreamPackageLinearAssemblyProgramsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageLinearAssemblyProgramsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageLinearAssemblyPrograms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageLinearAssemblyPrograms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageLinearAssemblyProgramsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSSAIChannelRequest() (request *DescribeStreamPackageSSAIChannelRequest) {
    request = &DescribeStreamPackageSSAIChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSSAIChannel")
    
    
    return
}

func NewDescribeStreamPackageSSAIChannelResponse() (response *DescribeStreamPackageSSAIChannelResponse) {
    response = &DescribeStreamPackageSSAIChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSSAIChannel
// DescribeStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) DescribeStreamPackageSSAIChannel(request *DescribeStreamPackageSSAIChannelRequest) (response *DescribeStreamPackageSSAIChannelResponse, err error) {
    return c.DescribeStreamPackageSSAIChannelWithContext(context.Background(), request)
}

// DescribeStreamPackageSSAIChannel
// DescribeStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) DescribeStreamPackageSSAIChannelWithContext(ctx context.Context, request *DescribeStreamPackageSSAIChannelRequest) (response *DescribeStreamPackageSSAIChannelResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSSAIChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSSAIChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSSAIChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSSAIChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSSAIChannelsRequest() (request *DescribeStreamPackageSSAIChannelsRequest) {
    request = &DescribeStreamPackageSSAIChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSSAIChannels")
    
    
    return
}

func NewDescribeStreamPackageSSAIChannelsResponse() (response *DescribeStreamPackageSSAIChannelsResponse) {
    response = &DescribeStreamPackageSSAIChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSSAIChannels
// DescribeStreamPackageSSAIChannels
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) DescribeStreamPackageSSAIChannels(request *DescribeStreamPackageSSAIChannelsRequest) (response *DescribeStreamPackageSSAIChannelsResponse, err error) {
    return c.DescribeStreamPackageSSAIChannelsWithContext(context.Background(), request)
}

// DescribeStreamPackageSSAIChannels
// DescribeStreamPackageSSAIChannels
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) DescribeStreamPackageSSAIChannelsWithContext(ctx context.Context, request *DescribeStreamPackageSSAIChannelsRequest) (response *DescribeStreamPackageSSAIChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSSAIChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSSAIChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSSAIChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSSAIChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSSAIUsageRequest() (request *DescribeStreamPackageSSAIUsageRequest) {
    request = &DescribeStreamPackageSSAIUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSSAIUsage")
    
    
    return
}

func NewDescribeStreamPackageSSAIUsageResponse() (response *DescribeStreamPackageSSAIUsageResponse) {
    response = &DescribeStreamPackageSSAIUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSSAIUsage
// This API is used to query SSAI ad replacement usage.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETSSAICHANNELSFAILED = "InternalError.GetSSAIChannelsFailed"
//  INTERNALERROR_GETSSAIUSAGEPARTIALFAILED = "InternalError.GetSSAIUsagePartialFailed"
//  INVALIDPARAMETER_INVALIDPARAMETERCHANNELID = "InvalidParameter.InvalidParameterChannelId"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamPackageSSAIUsage(request *DescribeStreamPackageSSAIUsageRequest) (response *DescribeStreamPackageSSAIUsageResponse, err error) {
    return c.DescribeStreamPackageSSAIUsageWithContext(context.Background(), request)
}

// DescribeStreamPackageSSAIUsage
// This API is used to query SSAI ad replacement usage.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETSSAICHANNELSFAILED = "InternalError.GetSSAIChannelsFailed"
//  INTERNALERROR_GETSSAIUSAGEPARTIALFAILED = "InternalError.GetSSAIUsagePartialFailed"
//  INVALIDPARAMETER_INVALIDPARAMETERCHANNELID = "InvalidParameter.InvalidParameterChannelId"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) DescribeStreamPackageSSAIUsageWithContext(ctx context.Context, request *DescribeStreamPackageSSAIUsageRequest) (response *DescribeStreamPackageSSAIUsageResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSSAIUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSSAIUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSSAIUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSSAIUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSourceRequest() (request *DescribeStreamPackageSourceRequest) {
    request = &DescribeStreamPackageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSource")
    
    
    return
}

func NewDescribeStreamPackageSourceResponse() (response *DescribeStreamPackageSourceResponse) {
    response = &DescribeStreamPackageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSource
// Query channel linear assembly Source information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageSource(request *DescribeStreamPackageSourceRequest) (response *DescribeStreamPackageSourceResponse, err error) {
    return c.DescribeStreamPackageSourceWithContext(context.Background(), request)
}

// DescribeStreamPackageSource
// Query channel linear assembly Source information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageSourceWithContext(ctx context.Context, request *DescribeStreamPackageSourceRequest) (response *DescribeStreamPackageSourceResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSourceAlertsRequest() (request *DescribeStreamPackageSourceAlertsRequest) {
    request = &DescribeStreamPackageSourceAlertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSourceAlerts")
    
    
    return
}

func NewDescribeStreamPackageSourceAlertsResponse() (response *DescribeStreamPackageSourceAlertsResponse) {
    response = &DescribeStreamPackageSourceAlertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSourceAlerts
// Query channel linear assembly Source alarm information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DescribeStreamPackageSourceAlerts(request *DescribeStreamPackageSourceAlertsRequest) (response *DescribeStreamPackageSourceAlertsResponse, err error) {
    return c.DescribeStreamPackageSourceAlertsWithContext(context.Background(), request)
}

// DescribeStreamPackageSourceAlerts
// Query channel linear assembly Source alarm information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DescribeStreamPackageSourceAlertsWithContext(ctx context.Context, request *DescribeStreamPackageSourceAlertsRequest) (response *DescribeStreamPackageSourceAlertsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSourceAlertsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSourceAlerts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSourceAlerts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSourceAlertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSourceLocationRequest() (request *DescribeStreamPackageSourceLocationRequest) {
    request = &DescribeStreamPackageSourceLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSourceLocation")
    
    
    return
}

func NewDescribeStreamPackageSourceLocationResponse() (response *DescribeStreamPackageSourceLocationResponse) {
    response = &DescribeStreamPackageSourceLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSourceLocation
// Query channel linear assembly sourceLocation information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageSourceLocation(request *DescribeStreamPackageSourceLocationRequest) (response *DescribeStreamPackageSourceLocationResponse, err error) {
    return c.DescribeStreamPackageSourceLocationWithContext(context.Background(), request)
}

// DescribeStreamPackageSourceLocation
// Query channel linear assembly sourceLocation information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamPackageSourceLocationWithContext(ctx context.Context, request *DescribeStreamPackageSourceLocationRequest) (response *DescribeStreamPackageSourceLocationResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSourceLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSourceLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSourceLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSourceLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSourceLocationAlertsRequest() (request *DescribeStreamPackageSourceLocationAlertsRequest) {
    request = &DescribeStreamPackageSourceLocationAlertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSourceLocationAlerts")
    
    
    return
}

func NewDescribeStreamPackageSourceLocationAlertsResponse() (response *DescribeStreamPackageSourceLocationAlertsResponse) {
    response = &DescribeStreamPackageSourceLocationAlertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSourceLocationAlerts
// Query channel linear assembly Location alarm information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DescribeStreamPackageSourceLocationAlerts(request *DescribeStreamPackageSourceLocationAlertsRequest) (response *DescribeStreamPackageSourceLocationAlertsResponse, err error) {
    return c.DescribeStreamPackageSourceLocationAlertsWithContext(context.Background(), request)
}

// DescribeStreamPackageSourceLocationAlerts
// Query channel linear assembly Location alarm information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) DescribeStreamPackageSourceLocationAlertsWithContext(ctx context.Context, request *DescribeStreamPackageSourceLocationAlertsRequest) (response *DescribeStreamPackageSourceLocationAlertsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSourceLocationAlertsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSourceLocationAlerts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSourceLocationAlerts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSourceLocationAlertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSourceLocationsRequest() (request *DescribeStreamPackageSourceLocationsRequest) {
    request = &DescribeStreamPackageSourceLocationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSourceLocations")
    
    
    return
}

func NewDescribeStreamPackageSourceLocationsResponse() (response *DescribeStreamPackageSourceLocationsResponse) {
    response = &DescribeStreamPackageSourceLocationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSourceLocations
// Query channel linear assembly SourceLocation information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageSourceLocations(request *DescribeStreamPackageSourceLocationsRequest) (response *DescribeStreamPackageSourceLocationsResponse, err error) {
    return c.DescribeStreamPackageSourceLocationsWithContext(context.Background(), request)
}

// DescribeStreamPackageSourceLocations
// Query channel linear assembly SourceLocation information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageSourceLocationsWithContext(ctx context.Context, request *DescribeStreamPackageSourceLocationsRequest) (response *DescribeStreamPackageSourceLocationsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSourceLocationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSourceLocations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSourceLocations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSourceLocationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPackageSourcesRequest() (request *DescribeStreamPackageSourcesRequest) {
    request = &DescribeStreamPackageSourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "DescribeStreamPackageSources")
    
    
    return
}

func NewDescribeStreamPackageSourcesResponse() (response *DescribeStreamPackageSourcesResponse) {
    response = &DescribeStreamPackageSourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamPackageSources
// Query channel linear assembly Source information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageSources(request *DescribeStreamPackageSourcesRequest) (response *DescribeStreamPackageSourcesResponse, err error) {
    return c.DescribeStreamPackageSourcesWithContext(context.Background(), request)
}

// DescribeStreamPackageSources
// Query channel linear assembly Source information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_POINTS = "InvalidParameter.Points"
func (c *Client) DescribeStreamPackageSourcesWithContext(ctx context.Context, request *DescribeStreamPackageSourcesRequest) (response *DescribeStreamPackageSourcesResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPackageSourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "DescribeStreamPackageSources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamPackageSources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamPackageSourcesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageChannel")
    
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
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
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
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
//  INVALIDPARAMETER_URL = "InvalidParameter.Url"
func (c *Client) ModifyStreamPackageChannelEndpointWithContext(ctx context.Context, request *ModifyStreamPackageChannelEndpointRequest) (response *ModifyStreamPackageChannelEndpointResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageChannelEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageChannelEndpoint")
    
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageChannelInputAuthInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageChannelInputAuthInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageChannelInputAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageLinearAssemblyChannelRequest() (request *ModifyStreamPackageLinearAssemblyChannelRequest) {
    request = &ModifyStreamPackageLinearAssemblyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageLinearAssemblyChannel")
    
    
    return
}

func NewModifyStreamPackageLinearAssemblyChannelResponse() (response *ModifyStreamPackageLinearAssemblyChannelResponse) {
    response = &ModifyStreamPackageLinearAssemblyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamPackageLinearAssemblyChannel
// Modify channel linear assembly Channel configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageLinearAssemblyChannel(request *ModifyStreamPackageLinearAssemblyChannelRequest) (response *ModifyStreamPackageLinearAssemblyChannelResponse, err error) {
    return c.ModifyStreamPackageLinearAssemblyChannelWithContext(context.Background(), request)
}

// ModifyStreamPackageLinearAssemblyChannel
// Modify channel linear assembly Channel configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageLinearAssemblyChannelWithContext(ctx context.Context, request *ModifyStreamPackageLinearAssemblyChannelRequest) (response *ModifyStreamPackageLinearAssemblyChannelResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageLinearAssemblyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageLinearAssemblyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageLinearAssemblyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageLinearAssemblyChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageLinearAssemblyProgramRequest() (request *ModifyStreamPackageLinearAssemblyProgramRequest) {
    request = &ModifyStreamPackageLinearAssemblyProgramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageLinearAssemblyProgram")
    
    
    return
}

func NewModifyStreamPackageLinearAssemblyProgramResponse() (response *ModifyStreamPackageLinearAssemblyProgramResponse) {
    response = &ModifyStreamPackageLinearAssemblyProgramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamPackageLinearAssemblyProgram
// Modify channel linear assembly Program configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageLinearAssemblyProgram(request *ModifyStreamPackageLinearAssemblyProgramRequest) (response *ModifyStreamPackageLinearAssemblyProgramResponse, err error) {
    return c.ModifyStreamPackageLinearAssemblyProgramWithContext(context.Background(), request)
}

// ModifyStreamPackageLinearAssemblyProgram
// Modify channel linear assembly Program configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageLinearAssemblyProgramWithContext(ctx context.Context, request *ModifyStreamPackageLinearAssemblyProgramRequest) (response *ModifyStreamPackageLinearAssemblyProgramResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageLinearAssemblyProgramRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageLinearAssemblyProgram")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageLinearAssemblyProgram require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageLinearAssemblyProgramResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageSSAIChannelRequest() (request *ModifyStreamPackageSSAIChannelRequest) {
    request = &ModifyStreamPackageSSAIChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageSSAIChannel")
    
    
    return
}

func NewModifyStreamPackageSSAIChannelResponse() (response *ModifyStreamPackageSSAIChannelResponse) {
    response = &ModifyStreamPackageSSAIChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamPackageSSAIChannel
// ModifyStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) ModifyStreamPackageSSAIChannel(request *ModifyStreamPackageSSAIChannelRequest) (response *ModifyStreamPackageSSAIChannelResponse, err error) {
    return c.ModifyStreamPackageSSAIChannelWithContext(context.Background(), request)
}

// ModifyStreamPackageSSAIChannel
// ModifyStreamPackageSSAIChannel
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUTHINFO = "InvalidParameter.AuthInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SSAIINFO = "InvalidParameter.SSAIInfo"
func (c *Client) ModifyStreamPackageSSAIChannelWithContext(ctx context.Context, request *ModifyStreamPackageSSAIChannelRequest) (response *ModifyStreamPackageSSAIChannelResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageSSAIChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageSSAIChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageSSAIChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageSSAIChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageSourceRequest() (request *ModifyStreamPackageSourceRequest) {
    request = &ModifyStreamPackageSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageSource")
    
    
    return
}

func NewModifyStreamPackageSourceResponse() (response *ModifyStreamPackageSourceResponse) {
    response = &ModifyStreamPackageSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamPackageSource
// Modify channel linear assembly Source configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageSource(request *ModifyStreamPackageSourceRequest) (response *ModifyStreamPackageSourceResponse, err error) {
    return c.ModifyStreamPackageSourceWithContext(context.Background(), request)
}

// ModifyStreamPackageSource
// Modify channel linear assembly Source configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageSourceWithContext(ctx context.Context, request *ModifyStreamPackageSourceRequest) (response *ModifyStreamPackageSourceResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageSourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamPackageSourceLocationRequest() (request *ModifyStreamPackageSourceLocationRequest) {
    request = &ModifyStreamPackageSourceLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "ModifyStreamPackageSourceLocation")
    
    
    return
}

func NewModifyStreamPackageSourceLocationResponse() (response *ModifyStreamPackageSourceLocationResponse) {
    response = &ModifyStreamPackageSourceLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStreamPackageSourceLocation
// Modify channel linear assembly SourceLocation configuration
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageSourceLocation(request *ModifyStreamPackageSourceLocationRequest) (response *ModifyStreamPackageSourceLocationResponse, err error) {
    return c.ModifyStreamPackageSourceLocationWithContext(context.Background(), request)
}

// ModifyStreamPackageSourceLocation
// Modify channel linear assembly SourceLocation configuration
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) ModifyStreamPackageSourceLocationWithContext(ctx context.Context, request *ModifyStreamPackageSourceLocationRequest) (response *ModifyStreamPackageSourceLocationResponse, err error) {
    if request == nil {
        request = NewModifyStreamPackageSourceLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "ModifyStreamPackageSourceLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStreamPackageSourceLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStreamPackageSourceLocationResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamPackageLinearAssemblyChannelRequest() (request *StartStreamPackageLinearAssemblyChannelRequest) {
    request = &StartStreamPackageLinearAssemblyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "StartStreamPackageLinearAssemblyChannel")
    
    
    return
}

func NewStartStreamPackageLinearAssemblyChannelResponse() (response *StartStreamPackageLinearAssemblyChannelResponse) {
    response = &StartStreamPackageLinearAssemblyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartStreamPackageLinearAssemblyChannel
// Start Linear Assembly Channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) StartStreamPackageLinearAssemblyChannel(request *StartStreamPackageLinearAssemblyChannelRequest) (response *StartStreamPackageLinearAssemblyChannelResponse, err error) {
    return c.StartStreamPackageLinearAssemblyChannelWithContext(context.Background(), request)
}

// StartStreamPackageLinearAssemblyChannel
// Start Linear Assembly Channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) StartStreamPackageLinearAssemblyChannelWithContext(ctx context.Context, request *StartStreamPackageLinearAssemblyChannelRequest) (response *StartStreamPackageLinearAssemblyChannelResponse, err error) {
    if request == nil {
        request = NewStartStreamPackageLinearAssemblyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "StartStreamPackageLinearAssemblyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStreamPackageLinearAssemblyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStreamPackageLinearAssemblyChannelResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamPackageLinearAssemblyChannelRequest() (request *StopStreamPackageLinearAssemblyChannelRequest) {
    request = &StopStreamPackageLinearAssemblyChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "StopStreamPackageLinearAssemblyChannel")
    
    
    return
}

func NewStopStreamPackageLinearAssemblyChannelResponse() (response *StopStreamPackageLinearAssemblyChannelResponse) {
    response = &StopStreamPackageLinearAssemblyChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopStreamPackageLinearAssemblyChannel
// Stop linear assembly channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) StopStreamPackageLinearAssemblyChannel(request *StopStreamPackageLinearAssemblyChannelRequest) (response *StopStreamPackageLinearAssemblyChannelResponse, err error) {
    return c.StopStreamPackageLinearAssemblyChannelWithContext(context.Background(), request)
}

// StopStreamPackageLinearAssemblyChannel
// Stop linear assembly channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CACHEINFO = "InvalidParameter.CacheInfo"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"
func (c *Client) StopStreamPackageLinearAssemblyChannelWithContext(ctx context.Context, request *StopStreamPackageLinearAssemblyChannelRequest) (response *StopStreamPackageLinearAssemblyChannelResponse, err error) {
    if request == nil {
        request = NewStopStreamPackageLinearAssemblyChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "StopStreamPackageLinearAssemblyChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopStreamPackageLinearAssemblyChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopStreamPackageLinearAssemblyChannelResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "UnbindCdnDomainWithChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindCdnDomainWithChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindCdnDomainWithChannelResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindLinearAssemblyCDNDomainWithChannelRequest() (request *UnbindLinearAssemblyCDNDomainWithChannelRequest) {
    request = &UnbindLinearAssemblyCDNDomainWithChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mdp", APIVersion, "UnbindLinearAssemblyCDNDomainWithChannel")
    
    
    return
}

func NewUnbindLinearAssemblyCDNDomainWithChannelResponse() (response *UnbindLinearAssemblyCDNDomainWithChannelResponse) {
    response = &UnbindLinearAssemblyCDNDomainWithChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindLinearAssemblyCDNDomainWithChannel
// Unbind LinearAssembly channel with CDN domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CDNDOMAIN = "InvalidParameter.CdnDomain"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) UnbindLinearAssemblyCDNDomainWithChannel(request *UnbindLinearAssemblyCDNDomainWithChannelRequest) (response *UnbindLinearAssemblyCDNDomainWithChannelResponse, err error) {
    return c.UnbindLinearAssemblyCDNDomainWithChannelWithContext(context.Background(), request)
}

// UnbindLinearAssemblyCDNDomainWithChannel
// Unbind LinearAssembly channel with CDN domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CDNDOMAIN = "InvalidParameter.CdnDomain"
//  INVALIDPARAMETER_CHANNELID = "InvalidParameter.ChannelId"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) UnbindLinearAssemblyCDNDomainWithChannelWithContext(ctx context.Context, request *UnbindLinearAssemblyCDNDomainWithChannelRequest) (response *UnbindLinearAssemblyCDNDomainWithChannelResponse, err error) {
    if request == nil {
        request = NewUnbindLinearAssemblyCDNDomainWithChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mdp", APIVersion, "UnbindLinearAssemblyCDNDomainWithChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindLinearAssemblyCDNDomainWithChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindLinearAssemblyCDNDomainWithChannelResponse()
    err = c.Send(request, response)
    return
}
