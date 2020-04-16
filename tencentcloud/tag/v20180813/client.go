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

package v20180813

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-08-13"

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


func NewAddResourceTagRequest() (request *AddResourceTagRequest) {
    request = &AddResourceTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "AddResourceTag")
    return
}

func NewAddResourceTagResponse() (response *AddResourceTagResponse) {
    response = &AddResourceTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to associate resources with tags.
func (c *Client) AddResourceTag(request *AddResourceTagRequest) (response *AddResourceTagResponse, err error) {
    if request == nil {
        request = NewAddResourceTagRequest()
    }
    response = NewAddResourceTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagRequest() (request *CreateTagRequest) {
    request = &CreateTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "CreateTag")
    return
}

func NewCreateTagResponse() (response *CreateTagResponse) {
    response = &CreateTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a tag key and tag value pair.
func (c *Client) CreateTag(request *CreateTagRequest) (response *CreateTagResponse, err error) {
    if request == nil {
        request = NewCreateTagRequest()
    }
    response = NewCreateTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceTagRequest() (request *DeleteResourceTagRequest) {
    request = &DeleteResourceTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DeleteResourceTag")
    return
}

func NewDeleteResourceTagResponse() (response *DeleteResourceTagResponse) {
    response = &DeleteResourceTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unassociate tags and resources.
func (c *Client) DeleteResourceTag(request *DeleteResourceTagRequest) (response *DeleteResourceTagResponse, err error) {
    if request == nil {
        request = NewDeleteResourceTagRequest()
    }
    response = NewDeleteResourceTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagRequest() (request *DeleteTagRequest) {
    request = &DeleteTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DeleteTag")
    return
}

func NewDeleteTagResponse() (response *DeleteTagResponse) {
    response = &DeleteTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a tag key and tag value pair.
func (c *Client) DeleteTag(request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
    if request == nil {
        request = NewDeleteTagRequest()
    }
    response = NewDeleteTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsByResourceIdsRequest() (request *DescribeResourceTagsByResourceIdsRequest) {
    request = &DescribeResourceTagsByResourceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourceTagsByResourceIds")
    return
}

func NewDescribeResourceTagsByResourceIdsResponse() (response *DescribeResourceTagsByResourceIdsResponse) {
    response = &DescribeResourceTagsByResourceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query tag key and value pairs for existing resources.
func (c *Client) DescribeResourceTagsByResourceIds(request *DescribeResourceTagsByResourceIdsRequest) (response *DescribeResourceTagsByResourceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsByResourceIdsRequest()
    }
    response = NewDescribeResourceTagsByResourceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsByTagKeysRequest() (request *DescribeResourceTagsByTagKeysRequest) {
    request = &DescribeResourceTagsByTagKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourceTagsByTagKeys")
    return
}

func NewDescribeResourceTagsByTagKeysResponse() (response *DescribeResourceTagsByTagKeysResponse) {
    response = &DescribeResourceTagsByTagKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get resource tags based on tag keys.
func (c *Client) DescribeResourceTagsByTagKeys(request *DescribeResourceTagsByTagKeysRequest) (response *DescribeResourceTagsByTagKeysResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsByTagKeysRequest()
    }
    response = NewDescribeResourceTagsByTagKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByTagsRequest() (request *DescribeResourcesByTagsRequest) {
    request = &DescribeResourcesByTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourcesByTags")
    return
}

func NewDescribeResourcesByTagsResponse() (response *DescribeResourcesByTagsResponse) {
    response = &DescribeResourcesByTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query resources by tags.
func (c *Client) DescribeResourcesByTags(request *DescribeResourcesByTagsRequest) (response *DescribeResourcesByTagsResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByTagsRequest()
    }
    response = NewDescribeResourcesByTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagKeysRequest() (request *DescribeTagKeysRequest) {
    request = &DescribeTagKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTagKeys")
    return
}

func NewDescribeTagKeysResponse() (response *DescribeTagKeysResponse) {
    response = &DescribeTagKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query tag keys in an existing tag list.
func (c *Client) DescribeTagKeys(request *DescribeTagKeysRequest) (response *DescribeTagKeysResponse, err error) {
    if request == nil {
        request = NewDescribeTagKeysRequest()
    }
    response = NewDescribeTagKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
    request = &DescribeTagValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTagValues")
    return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
    response = &DescribeTagValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query tag values in an existing tag list.
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    if request == nil {
        request = NewDescribeTagValuesRequest()
    }
    response = NewDescribeTagValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
    request = &DescribeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTags")
    return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
    response = &DescribeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query existing tag lists.
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsRequest()
    }
    response = NewDescribeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceTagsRequest() (request *ModifyResourceTagsRequest) {
    request = &ModifyResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "ModifyResourceTags")
    return
}

func NewModifyResourceTagsResponse() (response *ModifyResourceTagsResponse) {
    response = &ModifyResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify all tags associated with a resource.
func (c *Client) ModifyResourceTags(request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourceTagsRequest()
    }
    response = NewModifyResourceTagsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceTagValueRequest() (request *UpdateResourceTagValueRequest) {
    request = &UpdateResourceTagValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tag", APIVersion, "UpdateResourceTagValue")
    return
}

func NewUpdateResourceTagValueResponse() (response *UpdateResourceTagValueResponse) {
    response = &UpdateResourceTagValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the values of tags associated with a resource (the tag key does not change).
func (c *Client) UpdateResourceTagValue(request *UpdateResourceTagValueRequest) (response *UpdateResourceTagValueResponse, err error) {
    if request == nil {
        request = NewUpdateResourceTagValueRequest()
    }
    response = NewUpdateResourceTagValueResponse()
    err = c.Send(request, response)
    return
}
