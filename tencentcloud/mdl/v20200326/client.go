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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateMediaLiveChannelRequest() (request *CreateMediaLiveChannelRequest) {
    request = &CreateMediaLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateMediaLiveChannel")
    return
}

func NewCreateMediaLiveChannelResponse() (response *CreateMediaLiveChannelResponse) {
    response = &CreateMediaLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a media channel.
func (c *Client) CreateMediaLiveChannel(request *CreateMediaLiveChannelRequest) (response *CreateMediaLiveChannelResponse, err error) {
    if request == nil {
        request = NewCreateMediaLiveChannelRequest()
    }
    response = NewCreateMediaLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaLiveInputRequest() (request *CreateMediaLiveInputRequest) {
    request = &CreateMediaLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateMediaLiveInput")
    return
}

func NewCreateMediaLiveInputResponse() (response *CreateMediaLiveInputResponse) {
    response = &CreateMediaLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a media input.
func (c *Client) CreateMediaLiveInput(request *CreateMediaLiveInputRequest) (response *CreateMediaLiveInputResponse, err error) {
    if request == nil {
        request = NewCreateMediaLiveInputRequest()
    }
    response = NewCreateMediaLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMediaLiveInputSecurityGroupRequest() (request *CreateMediaLiveInputSecurityGroupRequest) {
    request = &CreateMediaLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "CreateMediaLiveInputSecurityGroup")
    return
}

func NewCreateMediaLiveInputSecurityGroupResponse() (response *CreateMediaLiveInputSecurityGroupResponse) {
    response = &CreateMediaLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an input security group. Up to 5 ones can be created.
func (c *Client) CreateMediaLiveInputSecurityGroup(request *CreateMediaLiveInputSecurityGroupRequest) (response *CreateMediaLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateMediaLiveInputSecurityGroupRequest()
    }
    response = NewCreateMediaLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaLiveChannelRequest() (request *DeleteMediaLiveChannelRequest) {
    request = &DeleteMediaLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteMediaLiveChannel")
    return
}

func NewDeleteMediaLiveChannelResponse() (response *DeleteMediaLiveChannelResponse) {
    response = &DeleteMediaLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a MediaLive channel.
func (c *Client) DeleteMediaLiveChannel(request *DeleteMediaLiveChannelRequest) (response *DeleteMediaLiveChannelResponse, err error) {
    if request == nil {
        request = NewDeleteMediaLiveChannelRequest()
    }
    response = NewDeleteMediaLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaLiveInputRequest() (request *DeleteMediaLiveInputRequest) {
    request = &DeleteMediaLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteMediaLiveInput")
    return
}

func NewDeleteMediaLiveInputResponse() (response *DeleteMediaLiveInputResponse) {
    response = &DeleteMediaLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a media input.
func (c *Client) DeleteMediaLiveInput(request *DeleteMediaLiveInputRequest) (response *DeleteMediaLiveInputResponse, err error) {
    if request == nil {
        request = NewDeleteMediaLiveInputRequest()
    }
    response = NewDeleteMediaLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMediaLiveInputSecurityGroupRequest() (request *DeleteMediaLiveInputSecurityGroupRequest) {
    request = &DeleteMediaLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DeleteMediaLiveInputSecurityGroup")
    return
}

func NewDeleteMediaLiveInputSecurityGroupResponse() (response *DeleteMediaLiveInputSecurityGroupResponse) {
    response = &DeleteMediaLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an input security group.
func (c *Client) DeleteMediaLiveInputSecurityGroup(request *DeleteMediaLiveInputSecurityGroupRequest) (response *DeleteMediaLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteMediaLiveInputSecurityGroupRequest()
    }
    response = NewDeleteMediaLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaLiveChannelRequest() (request *DescribeMediaLiveChannelRequest) {
    request = &DescribeMediaLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeMediaLiveChannel")
    return
}

func NewDescribeMediaLiveChannelResponse() (response *DescribeMediaLiveChannelResponse) {
    response = &DescribeMediaLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of a MediaLive channel.
func (c *Client) DescribeMediaLiveChannel(request *DescribeMediaLiveChannelRequest) (response *DescribeMediaLiveChannelResponse, err error) {
    if request == nil {
        request = NewDescribeMediaLiveChannelRequest()
    }
    response = NewDescribeMediaLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaLiveChannelsRequest() (request *DescribeMediaLiveChannelsRequest) {
    request = &DescribeMediaLiveChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeMediaLiveChannels")
    return
}

func NewDescribeMediaLiveChannelsResponse() (response *DescribeMediaLiveChannelsResponse) {
    response = &DescribeMediaLiveChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of MediaLive channels in batches.
func (c *Client) DescribeMediaLiveChannels(request *DescribeMediaLiveChannelsRequest) (response *DescribeMediaLiveChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeMediaLiveChannelsRequest()
    }
    response = NewDescribeMediaLiveChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaLiveInputRequest() (request *DescribeMediaLiveInputRequest) {
    request = &DescribeMediaLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeMediaLiveInput")
    return
}

func NewDescribeMediaLiveInputResponse() (response *DescribeMediaLiveInputResponse) {
    response = &DescribeMediaLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query a media input.
func (c *Client) DescribeMediaLiveInput(request *DescribeMediaLiveInputRequest) (response *DescribeMediaLiveInputResponse, err error) {
    if request == nil {
        request = NewDescribeMediaLiveInputRequest()
    }
    response = NewDescribeMediaLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaLiveInputSecurityGroupRequest() (request *DescribeMediaLiveInputSecurityGroupRequest) {
    request = &DescribeMediaLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeMediaLiveInputSecurityGroup")
    return
}

func NewDescribeMediaLiveInputSecurityGroupResponse() (response *DescribeMediaLiveInputSecurityGroupResponse) {
    response = &DescribeMediaLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query an input security group.
func (c *Client) DescribeMediaLiveInputSecurityGroup(request *DescribeMediaLiveInputSecurityGroupRequest) (response *DescribeMediaLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeMediaLiveInputSecurityGroupRequest()
    }
    response = NewDescribeMediaLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaLiveInputSecurityGroupsRequest() (request *DescribeMediaLiveInputSecurityGroupsRequest) {
    request = &DescribeMediaLiveInputSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeMediaLiveInputSecurityGroups")
    return
}

func NewDescribeMediaLiveInputSecurityGroupsResponse() (response *DescribeMediaLiveInputSecurityGroupsResponse) {
    response = &DescribeMediaLiveInputSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of input security groups in batches.
func (c *Client) DescribeMediaLiveInputSecurityGroups(request *DescribeMediaLiveInputSecurityGroupsRequest) (response *DescribeMediaLiveInputSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeMediaLiveInputSecurityGroupsRequest()
    }
    response = NewDescribeMediaLiveInputSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMediaLiveInputsRequest() (request *DescribeMediaLiveInputsRequest) {
    request = &DescribeMediaLiveInputsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "DescribeMediaLiveInputs")
    return
}

func NewDescribeMediaLiveInputsResponse() (response *DescribeMediaLiveInputsResponse) {
    response = &DescribeMediaLiveInputsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of media inputs in batches.
func (c *Client) DescribeMediaLiveInputs(request *DescribeMediaLiveInputsRequest) (response *DescribeMediaLiveInputsResponse, err error) {
    if request == nil {
        request = NewDescribeMediaLiveInputsRequest()
    }
    response = NewDescribeMediaLiveInputsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaLiveChannelRequest() (request *ModifyMediaLiveChannelRequest) {
    request = &ModifyMediaLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "ModifyMediaLiveChannel")
    return
}

func NewModifyMediaLiveChannelResponse() (response *ModifyMediaLiveChannelResponse) {
    response = &ModifyMediaLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the information of a MediaLive channel.
func (c *Client) ModifyMediaLiveChannel(request *ModifyMediaLiveChannelRequest) (response *ModifyMediaLiveChannelResponse, err error) {
    if request == nil {
        request = NewModifyMediaLiveChannelRequest()
    }
    response = NewModifyMediaLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaLiveInputRequest() (request *ModifyMediaLiveInputRequest) {
    request = &ModifyMediaLiveInputRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "ModifyMediaLiveInput")
    return
}

func NewModifyMediaLiveInputResponse() (response *ModifyMediaLiveInputResponse) {
    response = &ModifyMediaLiveInputResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update a media input.
func (c *Client) ModifyMediaLiveInput(request *ModifyMediaLiveInputRequest) (response *ModifyMediaLiveInputResponse, err error) {
    if request == nil {
        request = NewModifyMediaLiveInputRequest()
    }
    response = NewModifyMediaLiveInputResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMediaLiveInputSecurityGroupRequest() (request *ModifyMediaLiveInputSecurityGroupRequest) {
    request = &ModifyMediaLiveInputSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "ModifyMediaLiveInputSecurityGroup")
    return
}

func NewModifyMediaLiveInputSecurityGroupResponse() (response *ModifyMediaLiveInputSecurityGroupResponse) {
    response = &ModifyMediaLiveInputSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update an input security group.
func (c *Client) ModifyMediaLiveInputSecurityGroup(request *ModifyMediaLiveInputSecurityGroupRequest) (response *ModifyMediaLiveInputSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyMediaLiveInputSecurityGroupRequest()
    }
    response = NewModifyMediaLiveInputSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartMediaLiveChannelRequest() (request *StartMediaLiveChannelRequest) {
    request = &StartMediaLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "StartMediaLiveChannel")
    return
}

func NewStartMediaLiveChannelResponse() (response *StartMediaLiveChannelResponse) {
    response = &StartMediaLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to start a MediaLive channel.
func (c *Client) StartMediaLiveChannel(request *StartMediaLiveChannelRequest) (response *StartMediaLiveChannelResponse, err error) {
    if request == nil {
        request = NewStartMediaLiveChannelRequest()
    }
    response = NewStartMediaLiveChannelResponse()
    err = c.Send(request, response)
    return
}

func NewStopMediaLiveChannelRequest() (request *StopMediaLiveChannelRequest) {
    request = &StopMediaLiveChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mdl", APIVersion, "StopMediaLiveChannel")
    return
}

func NewStopMediaLiveChannelResponse() (response *StopMediaLiveChannelResponse) {
    response = &StopMediaLiveChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to stop a MediaLive channel.
func (c *Client) StopMediaLiveChannel(request *StopMediaLiveChannelRequest) (response *StopMediaLiveChannelResponse, err error) {
    if request == nil {
        request = NewStopMediaLiveChannelRequest()
    }
    response = NewStopMediaLiveChannelResponse()
    err = c.Send(request, response)
    return
}
