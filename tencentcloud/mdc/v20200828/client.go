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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateMediaConnectFlowRequest() (request *CreateMediaConnectFlowRequest) {
    request = &CreateMediaConnectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "CreateMediaConnectFlow")
    return
}

func NewCreateMediaConnectFlowResponse() (response *CreateMediaConnectFlowResponse) {
    response = &CreateMediaConnectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMediaConnectFlow
// This API is used to create the configuration of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
func (c *Client) CreateMediaConnectFlow(request *CreateMediaConnectFlowRequest) (response *CreateMediaConnectFlowResponse, err error) {
    if request == nil {
        request = NewCreateMediaConnectFlowRequest()
    }
    response = NewCreateMediaConnectFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaConnectOutputRequest() (request *CreateMediaConnectOutputRequest) {
    request = &CreateMediaConnectOutputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "CreateMediaConnectOutput")
    return
}

func NewCreateMediaConnectOutputResponse() (response *CreateMediaConnectOutputResponse) {
    response = &CreateMediaConnectOutputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMediaConnectOutput
// This API is used to create the output information of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) CreateMediaConnectOutput(request *CreateMediaConnectOutputRequest) (response *CreateMediaConnectOutputResponse, err error) {
    if request == nil {
        request = NewCreateMediaConnectOutputRequest()
    }
    response = NewCreateMediaConnectOutputResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaConnectFlowRequest() (request *DeleteMediaConnectFlowRequest) {
    request = &DeleteMediaConnectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DeleteMediaConnectFlow")
    return
}

func NewDeleteMediaConnectFlowResponse() (response *DeleteMediaConnectFlowResponse) {
    response = &DeleteMediaConnectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMediaConnectFlow
// This API is used to delete the configuration of a MediaConnect flow.
//
// error code that may be returned:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteMediaConnectFlow(request *DeleteMediaConnectFlowRequest) (response *DeleteMediaConnectFlowResponse, err error) {
    if request == nil {
        request = NewDeleteMediaConnectFlowRequest()
    }
    response = NewDeleteMediaConnectFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaConnectOutputRequest() (request *DeleteMediaConnectOutputRequest) {
    request = &DeleteMediaConnectOutputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DeleteMediaConnectOutput")
    return
}

func NewDeleteMediaConnectOutputResponse() (response *DeleteMediaConnectOutputResponse) {
    response = &DeleteMediaConnectOutputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMediaConnectOutput
// This API is used to delete the output configuration of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) DeleteMediaConnectOutput(request *DeleteMediaConnectOutputRequest) (response *DeleteMediaConnectOutputResponse, err error) {
    if request == nil {
        request = NewDeleteMediaConnectOutputRequest()
    }
    response = NewDeleteMediaConnectOutputResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaConnectFlowRequest() (request *DescribeMediaConnectFlowRequest) {
    request = &DescribeMediaConnectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeMediaConnectFlow")
    return
}

func NewDescribeMediaConnectFlowResponse() (response *DescribeMediaConnectFlowResponse) {
    response = &DescribeMediaConnectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMediaConnectFlow
// This API is used to query the configuration information of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
func (c *Client) DescribeMediaConnectFlow(request *DescribeMediaConnectFlowRequest) (response *DescribeMediaConnectFlowResponse, err error) {
    if request == nil {
        request = NewDescribeMediaConnectFlowRequest()
    }
    response = NewDescribeMediaConnectFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaConnectFlowsRequest() (request *DescribeMediaConnectFlowsRequest) {
    request = &DescribeMediaConnectFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "DescribeMediaConnectFlows")
    return
}

func NewDescribeMediaConnectFlowsResponse() (response *DescribeMediaConnectFlowsResponse) {
    response = &DescribeMediaConnectFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMediaConnectFlows
// This API is used to query the configuration information of multiple MediaConnect flows in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
func (c *Client) DescribeMediaConnectFlows(request *DescribeMediaConnectFlowsRequest) (response *DescribeMediaConnectFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeMediaConnectFlowsRequest()
    }
    response = NewDescribeMediaConnectFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaConnectFlowRequest() (request *ModifyMediaConnectFlowRequest) {
    request = &ModifyMediaConnectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyMediaConnectFlow")
    return
}

func NewModifyMediaConnectFlowResponse() (response *ModifyMediaConnectFlowResponse) {
    response = &ModifyMediaConnectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMediaConnectFlow
// This API is used to modify the configuration information of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyMediaConnectFlow(request *ModifyMediaConnectFlowRequest) (response *ModifyMediaConnectFlowResponse, err error) {
    if request == nil {
        request = NewModifyMediaConnectFlowRequest()
    }
    response = NewModifyMediaConnectFlowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaConnectInputRequest() (request *ModifyMediaConnectInputRequest) {
    request = &ModifyMediaConnectInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyMediaConnectInput")
    return
}

func NewModifyMediaConnectInputResponse() (response *ModifyMediaConnectInputResponse) {
    response = &ModifyMediaConnectInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMediaConnectInput
// This API is used to modify the input information of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUT = "InvalidParameter.Input"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyMediaConnectInput(request *ModifyMediaConnectInputRequest) (response *ModifyMediaConnectInputResponse, err error) {
    if request == nil {
        request = NewModifyMediaConnectInputRequest()
    }
    response = NewModifyMediaConnectInputResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaConnectOutputRequest() (request *ModifyMediaConnectOutputRequest) {
    request = &ModifyMediaConnectOutputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "ModifyMediaConnectOutput")
    return
}

func NewModifyMediaConnectOutputResponse() (response *ModifyMediaConnectOutputResponse) {
    response = &ModifyMediaConnectOutputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMediaConnectOutput
// This API is used to modify the output configuration of a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"
//  INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) ModifyMediaConnectOutput(request *ModifyMediaConnectOutputRequest) (response *ModifyMediaConnectOutputResponse, err error) {
    if request == nil {
        request = NewModifyMediaConnectOutputRequest()
    }
    response = NewModifyMediaConnectOutputResponse()
    err = c.Send(request, response)
    return
}

func NewStartMediaConnectFlowRequest() (request *StartMediaConnectFlowRequest) {
    request = &StartMediaConnectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "StartMediaConnectFlow")
    return
}

func NewStartMediaConnectFlowResponse() (response *StartMediaConnectFlowResponse) {
    response = &StartMediaConnectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMediaConnectFlow
// This API is used to start a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StartMediaConnectFlow(request *StartMediaConnectFlowRequest) (response *StartMediaConnectFlowResponse, err error) {
    if request == nil {
        request = NewStartMediaConnectFlowRequest()
    }
    response = NewStartMediaConnectFlowResponse()
    err = c.Send(request, response)
    return
}

func NewStopMediaConnectFlowRequest() (request *StopMediaConnectFlowRequest) {
    request = &StopMediaConnectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdc", APIVersion, "StopMediaConnectFlow")
    return
}

func NewStopMediaConnectFlowResponse() (response *StopMediaConnectFlowResponse) {
    response = &StopMediaConnectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMediaConnectFlow
// This API is used to stop a MediaConnect flow.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
func (c *Client) StopMediaConnectFlow(request *StopMediaConnectFlowRequest) (response *StopMediaConnectFlowResponse, err error) {
    if request == nil {
        request = NewStopMediaConnectFlowRequest()
    }
    response = NewStopMediaConnectFlowResponse()
    err = c.Send(request, response)
    return
}
