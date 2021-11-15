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

package v20200326

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-03-26"

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


func NewCreateStreamLiveChannelRequest() (request *CreateStreamLiveChannelRequest) {
    request = &CreateStreamLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateStreamLiveChannel")
    
    return
}

func NewCreateStreamLiveChannelResponse() (response *CreateStreamLiveChannelResponse) {
    response = &CreateStreamLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLiveChannel
// This API is used to create a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AVTEMPLATES = "InvalidParameter.AVTemplates"
//  INVALIDPARAMETER_ALREADYASSOCIATEDCHANNEL = "InvalidParameter.AlreadyAssociatedChannel"
//  INVALIDPARAMETER_ATTACHEDINPUTS = "InvalidParameter.AttachedInputs"
//  INVALIDPARAMETER_AUDIOTEMPLATES = "InvalidParameter.AudioTemplates"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_VIDEOTEMPLATES = "InvalidParameter.VideoTemplates"
func (c *Client) CreateStreamLiveChannel(request *CreateStreamLiveChannelRequest) (response *CreateStreamLiveChannelResponse, err error) {
    if request == nil {
        request = NewCreateStreamLiveChannelRequest()
    }
    response = NewCreateStreamLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLiveInputRequest() (request *CreateStreamLiveInputRequest) {
    request = &CreateStreamLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateStreamLiveInput")
    
    return
}

func NewCreateStreamLiveInputResponse() (response *CreateStreamLiveInputResponse) {
    response = &CreateStreamLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLiveInput
// This API is used to create a StreamLive input.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_INPUTSETTINGS = "InvalidParameter.InputSettings"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SECURITYGROUPS = "InvalidParameter.SecurityGroups"
//  INVALIDPARAMETER_TYPE = "InvalidParameter.Type"
func (c *Client) CreateStreamLiveInput(request *CreateStreamLiveInputRequest) (response *CreateStreamLiveInputResponse, err error) {
    if request == nil {
        request = NewCreateStreamLiveInputRequest()
    }
    response = NewCreateStreamLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLiveInputSecurityGroupRequest() (request *CreateStreamLiveInputSecurityGroupRequest) {
    request = &CreateStreamLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateStreamLiveInputSecurityGroup")
    
    return
}

func NewCreateStreamLiveInputSecurityGroupResponse() (response *CreateStreamLiveInputSecurityGroupResponse) {
    response = &CreateStreamLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLiveInputSecurityGroup
// This API is used to create an input security group. Up to 5 security groups are allowed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
func (c *Client) CreateStreamLiveInputSecurityGroup(request *CreateStreamLiveInputSecurityGroupRequest) (response *CreateStreamLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateStreamLiveInputSecurityGroupRequest()
    }
    response = NewCreateStreamLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStreamLivePlanRequest() (request *CreateStreamLivePlanRequest) {
    request = &CreateStreamLivePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateStreamLivePlan")
    
    return
}

func NewCreateStreamLivePlanResponse() (response *CreateStreamLivePlanResponse) {
    response = &CreateStreamLivePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStreamLivePlan
// This API is used to create an event in the plan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PLAN = "InvalidParameter.Plan"
func (c *Client) CreateStreamLivePlan(request *CreateStreamLivePlanRequest) (response *CreateStreamLivePlanResponse, err error) {
    if request == nil {
        request = NewCreateStreamLivePlanRequest()
    }
    response = NewCreateStreamLivePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLiveChannelRequest() (request *DeleteStreamLiveChannelRequest) {
    request = &DeleteStreamLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteStreamLiveChannel")
    
    return
}

func NewDeleteStreamLiveChannelResponse() (response *DeleteStreamLiveChannelResponse) {
    response = &DeleteStreamLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLiveChannel
// This API is used to delete a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATE = "InvalidParameter.State"
//  INVALIDPARAMETER_STATEERROR = "InvalidParameter.StateError"
func (c *Client) DeleteStreamLiveChannel(request *DeleteStreamLiveChannelRequest) (response *DeleteStreamLiveChannelResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLiveChannelRequest()
    }
    response = NewDeleteStreamLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLiveInputRequest() (request *DeleteStreamLiveInputRequest) {
    request = &DeleteStreamLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteStreamLiveInput")
    
    return
}

func NewDeleteStreamLiveInputResponse() (response *DeleteStreamLiveInputResponse) {
    response = &DeleteStreamLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLiveInput
// This API is used to delete a StreamLive input.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALREADYASSOCIATEDCHANNEL = "InvalidParameter.AlreadyAssociatedChannel"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DeleteStreamLiveInput(request *DeleteStreamLiveInputRequest) (response *DeleteStreamLiveInputResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLiveInputRequest()
    }
    response = NewDeleteStreamLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLiveInputSecurityGroupRequest() (request *DeleteStreamLiveInputSecurityGroupRequest) {
    request = &DeleteStreamLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteStreamLiveInputSecurityGroup")
    
    return
}

func NewDeleteStreamLiveInputSecurityGroupResponse() (response *DeleteStreamLiveInputSecurityGroupResponse) {
    response = &DeleteStreamLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLiveInputSecurityGroup
// This API is used to delete an input security group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALREADYASSOCIATEDINPUT = "InvalidParameter.AlreadyAssociatedInput"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DeleteStreamLiveInputSecurityGroup(request *DeleteStreamLiveInputSecurityGroupRequest) (response *DeleteStreamLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLiveInputSecurityGroupRequest()
    }
    response = NewDeleteStreamLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStreamLivePlanRequest() (request *DeleteStreamLivePlanRequest) {
    request = &DeleteStreamLivePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteStreamLivePlan")
    
    return
}

func NewDeleteStreamLivePlanResponse() (response *DeleteStreamLivePlanResponse) {
    response = &DeleteStreamLivePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStreamLivePlan
// This API is used to delete a StreamLive event.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_PLAN = "InvalidParameter.Plan"
func (c *Client) DeleteStreamLivePlan(request *DeleteStreamLivePlanRequest) (response *DeleteStreamLivePlanResponse, err error) {
    if request == nil {
        request = NewDeleteStreamLivePlanRequest()
    }
    response = NewDeleteStreamLivePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveChannelRequest() (request *DescribeStreamLiveChannelRequest) {
    request = &DescribeStreamLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveChannel")
    
    return
}

func NewDescribeStreamLiveChannelResponse() (response *DescribeStreamLiveChannelResponse) {
    response = &DescribeStreamLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveChannel
// This API is used to query a StreamLive channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLiveChannel(request *DescribeStreamLiveChannelRequest) (response *DescribeStreamLiveChannelResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveChannelRequest()
    }
    response = NewDescribeStreamLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveChannelAlertsRequest() (request *DescribeStreamLiveChannelAlertsRequest) {
    request = &DescribeStreamLiveChannelAlertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveChannelAlerts")
    
    return
}

func NewDescribeStreamLiveChannelAlertsResponse() (response *DescribeStreamLiveChannelAlertsResponse) {
    response = &DescribeStreamLiveChannelAlertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveChannelAlerts
// This API is used to query the alarm information of a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLiveChannelAlerts(request *DescribeStreamLiveChannelAlertsRequest) (response *DescribeStreamLiveChannelAlertsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveChannelAlertsRequest()
    }
    response = NewDescribeStreamLiveChannelAlertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveChannelInputStatisticsRequest() (request *DescribeStreamLiveChannelInputStatisticsRequest) {
    request = &DescribeStreamLiveChannelInputStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveChannelInputStatistics")
    
    return
}

func NewDescribeStreamLiveChannelInputStatisticsResponse() (response *DescribeStreamLiveChannelInputStatisticsResponse) {
    response = &DescribeStreamLiveChannelInputStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveChannelInputStatistics
// This API is used to query input statistics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
func (c *Client) DescribeStreamLiveChannelInputStatistics(request *DescribeStreamLiveChannelInputStatisticsRequest) (response *DescribeStreamLiveChannelInputStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveChannelInputStatisticsRequest()
    }
    response = NewDescribeStreamLiveChannelInputStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveChannelLogsRequest() (request *DescribeStreamLiveChannelLogsRequest) {
    request = &DescribeStreamLiveChannelLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveChannelLogs")
    
    return
}

func NewDescribeStreamLiveChannelLogsResponse() (response *DescribeStreamLiveChannelLogsResponse) {
    response = &DescribeStreamLiveChannelLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveChannelLogs
// This API is used to query StreamLive channel logs, such as push event logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
func (c *Client) DescribeStreamLiveChannelLogs(request *DescribeStreamLiveChannelLogsRequest) (response *DescribeStreamLiveChannelLogsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveChannelLogsRequest()
    }
    response = NewDescribeStreamLiveChannelLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveChannelOutputStatisticsRequest() (request *DescribeStreamLiveChannelOutputStatisticsRequest) {
    request = &DescribeStreamLiveChannelOutputStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveChannelOutputStatistics")
    
    return
}

func NewDescribeStreamLiveChannelOutputStatisticsResponse() (response *DescribeStreamLiveChannelOutputStatisticsResponse) {
    response = &DescribeStreamLiveChannelOutputStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveChannelOutputStatistics
// This API is used to query the output statistics of a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
func (c *Client) DescribeStreamLiveChannelOutputStatistics(request *DescribeStreamLiveChannelOutputStatisticsRequest) (response *DescribeStreamLiveChannelOutputStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveChannelOutputStatisticsRequest()
    }
    response = NewDescribeStreamLiveChannelOutputStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveChannelsRequest() (request *DescribeStreamLiveChannelsRequest) {
    request = &DescribeStreamLiveChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveChannels")
    
    return
}

func NewDescribeStreamLiveChannelsResponse() (response *DescribeStreamLiveChannelsResponse) {
    response = &DescribeStreamLiveChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveChannels
// This API is used to query StreamLive channels in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLiveChannels(request *DescribeStreamLiveChannelsRequest) (response *DescribeStreamLiveChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveChannelsRequest()
    }
    response = NewDescribeStreamLiveChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveInputRequest() (request *DescribeStreamLiveInputRequest) {
    request = &DescribeStreamLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveInput")
    
    return
}

func NewDescribeStreamLiveInputResponse() (response *DescribeStreamLiveInputResponse) {
    response = &DescribeStreamLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveInput
// This API is used to query a StreamLive input.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLiveInput(request *DescribeStreamLiveInputRequest) (response *DescribeStreamLiveInputResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveInputRequest()
    }
    response = NewDescribeStreamLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveInputSecurityGroupRequest() (request *DescribeStreamLiveInputSecurityGroupRequest) {
    request = &DescribeStreamLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveInputSecurityGroup")
    
    return
}

func NewDescribeStreamLiveInputSecurityGroupResponse() (response *DescribeStreamLiveInputSecurityGroupResponse) {
    response = &DescribeStreamLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveInputSecurityGroup
// This API is used to query an input security group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLiveInputSecurityGroup(request *DescribeStreamLiveInputSecurityGroupRequest) (response *DescribeStreamLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveInputSecurityGroupRequest()
    }
    response = NewDescribeStreamLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveInputSecurityGroupsRequest() (request *DescribeStreamLiveInputSecurityGroupsRequest) {
    request = &DescribeStreamLiveInputSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveInputSecurityGroups")
    
    return
}

func NewDescribeStreamLiveInputSecurityGroupsResponse() (response *DescribeStreamLiveInputSecurityGroupsResponse) {
    response = &DescribeStreamLiveInputSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveInputSecurityGroups
// This API is used to query input security groups in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLiveInputSecurityGroups(request *DescribeStreamLiveInputSecurityGroupsRequest) (response *DescribeStreamLiveInputSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveInputSecurityGroupsRequest()
    }
    response = NewDescribeStreamLiveInputSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveInputsRequest() (request *DescribeStreamLiveInputsRequest) {
    request = &DescribeStreamLiveInputsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveInputs")
    
    return
}

func NewDescribeStreamLiveInputsResponse() (response *DescribeStreamLiveInputsResponse) {
    response = &DescribeStreamLiveInputsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveInputs
// This API is used to query StreamLive inputs in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLiveInputs(request *DescribeStreamLiveInputsRequest) (response *DescribeStreamLiveInputsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveInputsRequest()
    }
    response = NewDescribeStreamLiveInputsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLivePlansRequest() (request *DescribeStreamLivePlansRequest) {
    request = &DescribeStreamLivePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLivePlans")
    
    return
}

func NewDescribeStreamLivePlansResponse() (response *DescribeStreamLivePlansResponse) {
    response = &DescribeStreamLivePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLivePlans
// This API is used to query the events in the plan in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
func (c *Client) DescribeStreamLivePlans(request *DescribeStreamLivePlansRequest) (response *DescribeStreamLivePlansResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLivePlansRequest()
    }
    response = NewDescribeStreamLivePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamLiveRegionsRequest() (request *DescribeStreamLiveRegionsRequest) {
    request = &DescribeStreamLiveRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeStreamLiveRegions")
    
    return
}

func NewDescribeStreamLiveRegionsResponse() (response *DescribeStreamLiveRegionsResponse) {
    response = &DescribeStreamLiveRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamLiveRegions
// This API is used to query all StreamLive regions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeStreamLiveRegions(request *DescribeStreamLiveRegionsRequest) (response *DescribeStreamLiveRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeStreamLiveRegionsRequest()
    }
    response = NewDescribeStreamLiveRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLiveChannelRequest() (request *ModifyStreamLiveChannelRequest) {
    request = &ModifyStreamLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "ModifyStreamLiveChannel")
    
    return
}

func NewModifyStreamLiveChannelResponse() (response *ModifyStreamLiveChannelResponse) {
    response = &ModifyStreamLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLiveChannel
// This API is used to modify a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AVTEMPLATES = "InvalidParameter.AVTemplates"
//  INVALIDPARAMETER_ALREADYASSOCIATEDCHANNEL = "InvalidParameter.AlreadyAssociatedChannel"
//  INVALIDPARAMETER_ATTACHEDINPUTS = "InvalidParameter.AttachedInputs"
//  INVALIDPARAMETER_AUDIOTEMPLATES = "InvalidParameter.AudioTemplates"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"
//  INVALIDPARAMETER_VIDEOTEMPLATES = "InvalidParameter.VideoTemplates"
func (c *Client) ModifyStreamLiveChannel(request *ModifyStreamLiveChannelRequest) (response *ModifyStreamLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyStreamLiveChannelRequest()
    }
    response = NewModifyStreamLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLiveInputRequest() (request *ModifyStreamLiveInputRequest) {
    request = &ModifyStreamLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "ModifyStreamLiveInput")
    
    return
}

func NewModifyStreamLiveInputResponse() (response *ModifyStreamLiveInputResponse) {
    response = &ModifyStreamLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLiveInput
// This API is used to modify a StreamLive input.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_INPUTSETTINGS = "InvalidParameter.InputSettings"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_SECURITYGROUPS = "InvalidParameter.SecurityGroups"
//  INVALIDPARAMETER_STATEERROR = "InvalidParameter.StateError"
func (c *Client) ModifyStreamLiveInput(request *ModifyStreamLiveInputRequest) (response *ModifyStreamLiveInputResponse, err error) {
    if request == nil {
        request = NewModifyStreamLiveInputRequest()
    }
    response = NewModifyStreamLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStreamLiveInputSecurityGroupRequest() (request *ModifyStreamLiveInputSecurityGroupRequest) {
    request = &ModifyStreamLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "ModifyStreamLiveInputSecurityGroup")
    
    return
}

func NewModifyStreamLiveInputSecurityGroupResponse() (response *ModifyStreamLiveInputSecurityGroupResponse) {
    response = &ModifyStreamLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyStreamLiveInputSecurityGroup
// This API is used to modify an input security group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
func (c *Client) ModifyStreamLiveInputSecurityGroup(request *ModifyStreamLiveInputSecurityGroupRequest) (response *ModifyStreamLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyStreamLiveInputSecurityGroupRequest()
    }
    response = NewModifyStreamLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamLiveChannelRequest() (request *StartStreamLiveChannelRequest) {
    request = &StartStreamLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "StartStreamLiveChannel")
    
    return
}

func NewStartStreamLiveChannelResponse() (response *StartStreamLiveChannelResponse) {
    response = &StartStreamLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartStreamLiveChannel
// This API is used to start a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATEERROR = "InvalidParameter.StateError"
func (c *Client) StartStreamLiveChannel(request *StartStreamLiveChannelRequest) (response *StartStreamLiveChannelResponse, err error) {
    if request == nil {
        request = NewStartStreamLiveChannelRequest()
    }
    response = NewStartStreamLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamLiveChannelRequest() (request *StopStreamLiveChannelRequest) {
    request = &StopStreamLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "StopStreamLiveChannel")
    
    return
}

func NewStopStreamLiveChannelResponse() (response *StopStreamLiveChannelResponse) {
    response = &StopStreamLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopStreamLiveChannel
// This API is used to stop a StreamLive channel.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ID = "InvalidParameter.Id"
//  INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"
//  INVALIDPARAMETER_STATEERROR = "InvalidParameter.StateError"
func (c *Client) StopStreamLiveChannel(request *StopStreamLiveChannelRequest) (response *StopStreamLiveChannelResponse, err error) {
    if request == nil {
        request = NewStopStreamLiveChannelRequest()
    }
    response = NewStopStreamLiveChannelResponse()
    err = c.Send(request, response)
    return
}
