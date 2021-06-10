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

package v20190319

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-03-19"

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


func NewCreateRecorderRequest() (request *CreateRecorderRequest) {
    request = &CreateRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateRecorder")
    return
}

func NewCreateRecorderResponse() (response *CreateRecorderResponse) {
    response = &CreateRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecorder
// This API is used to create resource recorders to detect and record resource configuration changes.
//
// error code that may be returned:
//  LIMITEXCEEDED_RECORDEROVERAMOUNT = "LimitExceeded.RecorderOverAmount"
func (c *Client) CreateRecorder(request *CreateRecorderRequest) (response *CreateRecorderResponse, err error) {
    if request == nil {
        request = NewCreateRecorderRequest()
    }
    response = NewCreateRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecorderRequest() (request *DeleteRecorderRequest) {
    request = &DeleteRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteRecorder")
    return
}

func NewDeleteRecorderResponse() (response *DeleteRecorderResponse) {
    response = &DeleteRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecorder
// This API is used to delete resource recorders. After deletion, resource configuration changes will not be recorded.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RECORDERNOTFOUND = "ResourceNotFound.RecorderNotFound"
func (c *Client) DeleteRecorder(request *DeleteRecorderRequest) (response *DeleteRecorderResponse, err error) {
    if request == nil {
        request = NewDeleteRecorderRequest()
    }
    response = NewDeleteRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiscoveredResourceRequest() (request *DescribeDiscoveredResourceRequest) {
    request = &DescribeDiscoveredResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeDiscoveredResource")
    return
}

func NewDescribeDiscoveredResourceResponse() (response *DescribeDiscoveredResourceResponse) {
    response = &DescribeDiscoveredResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiscoveredResource
// This API is used to view the basic information of discovered resources.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) DescribeDiscoveredResource(request *DescribeDiscoveredResourceRequest) (response *DescribeDiscoveredResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDiscoveredResourceRequest()
    }
    response = NewDescribeDiscoveredResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecorderRequest() (request *DescribeRecorderRequest) {
    request = &DescribeRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeRecorder")
    return
}

func NewDescribeRecorderResponse() (response *DescribeRecorderResponse) {
    response = &DescribeRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecorder
// This API is used to display current configurations and status of a recorder.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RECORDERNOTFOUND = "ResourceNotFound.RecorderNotFound"
func (c *Client) DescribeRecorder(request *DescribeRecorderRequest) (response *DescribeRecorderResponse, err error) {
    if request == nil {
        request = NewDescribeRecorderRequest()
    }
    response = NewDescribeRecorderResponse()
    err = c.Send(request, response)
    return
}

func NewGetConfigurationItemsRequest() (request *GetConfigurationItemsRequest) {
    request = &GetConfigurationItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "GetConfigurationItems")
    return
}

func NewGetConfigurationItemsResponse() (response *GetConfigurationItemsResponse) {
    response = &GetConfigurationItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetConfigurationItems
// This API is used to get the list of resource configuration items and display resource configuration changes in chronological order.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) GetConfigurationItems(request *GetConfigurationItemsRequest) (response *GetConfigurationItemsResponse, err error) {
    if request == nil {
        request = NewGetConfigurationItemsRequest()
    }
    response = NewGetConfigurationItemsResponse()
    err = c.Send(request, response)
    return
}

func NewListDiscoveredResourcesRequest() (request *ListDiscoveredResourcesRequest) {
    request = &ListDiscoveredResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListDiscoveredResources")
    return
}

func NewListDiscoveredResourcesResponse() (response *ListDiscoveredResourcesResponse) {
    response = &ListDiscoveredResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListDiscoveredResources
// This API is used to view the list of discovered resources.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (response *ListDiscoveredResourcesResponse, err error) {
    if request == nil {
        request = NewListDiscoveredResourcesRequest()
    }
    response = NewListDiscoveredResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewListSupportResourceTypesRequest() (request *ListSupportResourceTypesRequest) {
    request = &ListSupportResourceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListSupportResourceTypes")
    return
}

func NewListSupportResourceTypesResponse() (response *ListSupportResourceTypesResponse) {
    response = &ListSupportResourceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSupportResourceTypes
// This API is used to query the list of all CFA supported resource types.
//
// error code that may be returned:
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
func (c *Client) ListSupportResourceTypes(request *ListSupportResourceTypesRequest) (response *ListSupportResourceTypesResponse, err error) {
    if request == nil {
        request = NewListSupportResourceTypesRequest()
    }
    response = NewListSupportResourceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecorderRequest() (request *UpdateRecorderRequest) {
    request = &UpdateRecorderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "UpdateRecorder")
    return
}

func NewUpdateRecorderResponse() (response *UpdateRecorderResponse) {
    response = &UpdateRecorderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRecorder
// This API is used to modify the resources to monitor, recorder name, and other recorder configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_RECORDERNAMEREPEAT = "InvalidParameterValue.RecorderNameRepeat"
//  RESOURCENOTFOUND_RECORDERNOTFOUND = "ResourceNotFound.RecorderNotFound"
func (c *Client) UpdateRecorder(request *UpdateRecorderRequest) (response *UpdateRecorderResponse, err error) {
    if request == nil {
        request = NewUpdateRecorderRequest()
    }
    response = NewUpdateRecorderResponse()
    err = c.Send(request, response)
    return
}
