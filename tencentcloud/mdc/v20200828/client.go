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

package v20200828

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-08-28"

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


func NewCreateStreamLinkFlowRequest() (request *CreateStreamLinkFlowRequest) {
    request = &CreateStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "CreateStreamLinkFlow")
    
    
    return
}

func NewCreateStreamLinkFlowResponse() (response *CreateStreamLinkFlowResponse) {
    response = &CreateStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLinkFlow
// This API is used to create a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlow(request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkFlowRequest()
    }
    
    response = NewCreateStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

// CreateStreamLinkFlow
// This API is used to create a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateStreamLinkFlowWithContext(ctx context.Context, request *CreateStreamLinkFlowRequest) (response *CreateStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewCreateStreamLinkFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkFlowRequest() (request *DeleteStreamLinkFlowRequest) {
    request = &DeleteStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DeleteStreamLinkFlow")
    
    
    return
}

func NewDeleteStreamLinkFlowResponse() (response *DeleteStreamLinkFlowResponse) {
    response = &DeleteStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLinkFlow
// This API is used to delete a StreamLink flow.
//
// error code that may be returned:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkFlow(request *DeleteStreamLinkFlowRequest) (response *DeleteStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkFlowRequest()
    }
    
    response = NewDeleteStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

// DeleteStreamLinkFlow
// This API is used to delete a StreamLink flow.
//
// error code that may be returned:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkFlowWithContext(ctx context.Context, request *DeleteStreamLinkFlowRequest) (response *DeleteStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLinkOutputRequest() (request *DeleteStreamLinkOutputRequest) {
    request = &DeleteStreamLinkOutputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DeleteStreamLinkOutput")
    
    
    return
}

func NewDeleteStreamLinkOutputResponse() (response *DeleteStreamLinkOutputResponse) {
    response = &DeleteStreamLinkOutputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLinkOutput
// This API is used to delete an output of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkOutput(request *DeleteStreamLinkOutputRequest) (response *DeleteStreamLinkOutputResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkOutputRequest()
    }
    
    response = NewDeleteStreamLinkOutputResponse()
    err = c.Send(request, response)
    return
}

// DeleteStreamLinkOutput
// This API is used to delete an output of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteStreamLinkOutputWithContext(ctx context.Context, request *DeleteStreamLinkOutputRequest) (response *DeleteStreamLinkOutputResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLinkOutputRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteStreamLinkOutputResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowRequest() (request *DescribeStreamLinkFlowRequest) {
    request = &DescribeStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlow")
    
    
    return
}

func NewDescribeStreamLinkFlowResponse() (response *DescribeStreamLinkFlowResponse) {
    response = &DescribeStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlow
// This API is used to query the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeStreamLinkFlow(request *DescribeStreamLinkFlowRequest) (response *DescribeStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowRequest()
    }
    
    response = NewDescribeStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

// DescribeStreamLinkFlow
// This API is used to query the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeStreamLinkFlowWithContext(ctx context.Context, request *DescribeStreamLinkFlowRequest) (response *DescribeStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkFlowsRequest() (request *DescribeStreamLinkFlowsRequest) {
    request = &DescribeStreamLinkFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkFlows")
    
    
    return
}

func NewDescribeStreamLinkFlowsResponse() (response *DescribeStreamLinkFlowsResponse) {
    response = &DescribeStreamLinkFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkFlows
// This API is used to query the configuration information of multiple StreamLink flows in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkFlows(request *DescribeStreamLinkFlowsRequest) (response *DescribeStreamLinkFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowsRequest()
    }
    
    response = NewDescribeStreamLinkFlowsResponse()
    err = c.Send(request, response)
    return
}

// DescribeStreamLinkFlows
// This API is used to query the configuration information of multiple StreamLink flows in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeStreamLinkFlowsWithContext(ctx context.Context, request *DescribeStreamLinkFlowsRequest) (response *DescribeStreamLinkFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkFlowsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLinkRegionsRequest() (request *DescribeStreamLinkRegionsRequest) {
    request = &DescribeStreamLinkRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeStreamLinkRegions")
    
    
    return
}

func NewDescribeStreamLinkRegionsResponse() (response *DescribeStreamLinkRegionsResponse) {
    response = &DescribeStreamLinkRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLinkRegions
// This API is used to query all StreamLink regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkRegions(request *DescribeStreamLinkRegionsRequest) (response *DescribeStreamLinkRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkRegionsRequest()
    }
    
    response = NewDescribeStreamLinkRegionsResponse()
    err = c.Send(request, response)
    return
}

// DescribeStreamLinkRegions
// This API is used to query all StreamLink regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLinkRegionsWithContext(ctx context.Context, request *DescribeStreamLinkRegionsRequest) (response *DescribeStreamLinkRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLinkRegionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStreamLinkRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLinkFlowRequest() (request *ModifyStreamLinkFlowRequest) {
    request = &ModifyStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyStreamLinkFlow")
    
    
    return
}

func NewModifyStreamLinkFlowResponse() (response *ModifyStreamLinkFlowResponse) {
    response = &ModifyStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLinkFlow
// This API is used to modify the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkFlow(request *ModifyStreamLinkFlowRequest) (response *ModifyStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkFlowRequest()
    }
    
    response = NewModifyStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

// ModifyStreamLinkFlow
// This API is used to modify the configuration information of a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyStreamLinkFlowWithContext(ctx context.Context, request *ModifyStreamLinkFlowRequest) (response *ModifyStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewModifyStreamLinkFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamLinkFlowRequest() (request *StartStreamLinkFlowRequest) {
    request = &StartStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "StartStreamLinkFlow")
    
    
    return
}

func NewStartStreamLinkFlowResponse() (response *StartStreamLinkFlowResponse) {
    response = &StartStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartStreamLinkFlow
// This API is used to start a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartStreamLinkFlow(request *StartStreamLinkFlowRequest) (response *StartStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStartStreamLinkFlowRequest()
    }
    
    response = NewStartStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

// StartStreamLinkFlow
// This API is used to start a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartStreamLinkFlowWithContext(ctx context.Context, request *StartStreamLinkFlowRequest) (response *StartStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStartStreamLinkFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamLinkFlowRequest() (request *StopStreamLinkFlowRequest) {
    request = &StopStreamLinkFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "StopStreamLinkFlow")
    
    
    return
}

func NewStopStreamLinkFlowResponse() (response *StopStreamLinkFlowResponse) {
    response = &StopStreamLinkFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopStreamLinkFlow
// This API is used to stop a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopStreamLinkFlow(request *StopStreamLinkFlowRequest) (response *StopStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStopStreamLinkFlowRequest()
    }
    
    response = NewStopStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}

// StopStreamLinkFlow
// This API is used to stop a StreamLink flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopStreamLinkFlowWithContext(ctx context.Context, request *StopStreamLinkFlowRequest) (response *StopStreamLinkFlowResponse, err error) {
    if request == nil {
        request = NewStopStreamLinkFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopStreamLinkFlowResponse()
    err = c.Send(request, response)
    return
}
