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

package v20220802

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-08-02"

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


func NewDescribeDiscoveredResourceRequest() (request *DescribeDiscoveredResourceRequest) {
    request = &DescribeDiscoveredResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "DescribeDiscoveredResource")
    
    
    return
}

func NewDescribeDiscoveredResourceResponse() (response *DescribeDiscoveredResourceResponse) {
    response = &DescribeDiscoveredResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDiscoveredResource
// Resource details.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) DescribeDiscoveredResource(request *DescribeDiscoveredResourceRequest) (response *DescribeDiscoveredResourceResponse, err error) {
    return c.DescribeDiscoveredResourceWithContext(context.Background(), request)
}

// DescribeDiscoveredResource
// Resource details.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) DescribeDiscoveredResourceWithContext(ctx context.Context, request *DescribeDiscoveredResourceRequest) (response *DescribeDiscoveredResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDiscoveredResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "DescribeDiscoveredResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiscoveredResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiscoveredResourceResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregateConfigRulesRequest() (request *ListAggregateConfigRulesRequest) {
    request = &ListAggregateConfigRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateConfigRules")
    
    
    return
}

func NewListAggregateConfigRulesResponse() (response *ListAggregateConfigRulesResponse) {
    response = &ListAggregateConfigRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateConfigRules
// This API is used to get the account group rule list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateConfigRules(request *ListAggregateConfigRulesRequest) (response *ListAggregateConfigRulesResponse, err error) {
    return c.ListAggregateConfigRulesWithContext(context.Background(), request)
}

// ListAggregateConfigRules
// This API is used to get the account group rule list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateConfigRulesWithContext(ctx context.Context, request *ListAggregateConfigRulesRequest) (response *ListAggregateConfigRulesResponse, err error) {
    if request == nil {
        request = NewListAggregateConfigRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregateConfigRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateConfigRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateConfigRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListAggregateDiscoveredResourcesRequest() (request *ListAggregateDiscoveredResourcesRequest) {
    request = &ListAggregateDiscoveredResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListAggregateDiscoveredResources")
    
    
    return
}

func NewListAggregateDiscoveredResourcesResponse() (response *ListAggregateDiscoveredResourcesResponse) {
    response = &ListAggregateDiscoveredResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAggregateDiscoveredResources
// Account Group access the list of resources.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateDiscoveredResources(request *ListAggregateDiscoveredResourcesRequest) (response *ListAggregateDiscoveredResourcesResponse, err error) {
    return c.ListAggregateDiscoveredResourcesWithContext(context.Background(), request)
}

// ListAggregateDiscoveredResources
// Account Group access the list of resources.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
func (c *Client) ListAggregateDiscoveredResourcesWithContext(ctx context.Context, request *ListAggregateDiscoveredResourcesRequest) (response *ListAggregateDiscoveredResourcesResponse, err error) {
    if request == nil {
        request = NewListAggregateDiscoveredResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListAggregateDiscoveredResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAggregateDiscoveredResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAggregateDiscoveredResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewListConfigRulesRequest() (request *ListConfigRulesRequest) {
    request = &ListConfigRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListConfigRules")
    
    
    return
}

func NewListConfigRulesResponse() (response *ListConfigRulesResponse) {
    response = &ListConfigRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListConfigRules
// This API is used to get the rule list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListConfigRules(request *ListConfigRulesRequest) (response *ListConfigRulesResponse, err error) {
    return c.ListConfigRulesWithContext(context.Background(), request)
}

// ListConfigRules
// This API is used to get the rule list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListConfigRulesWithContext(ctx context.Context, request *ListConfigRulesRequest) (response *ListConfigRulesResponse, err error) {
    if request == nil {
        request = NewListConfigRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListConfigRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListConfigRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListConfigRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListDiscoveredResourcesRequest() (request *ListDiscoveredResourcesRequest) {
    request = &ListDiscoveredResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "ListDiscoveredResources")
    
    
    return
}

func NewListDiscoveredResourcesResponse() (response *ListDiscoveredResourcesResponse) {
    response = &ListDiscoveredResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDiscoveredResources
// This API is used to get the resource list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (response *ListDiscoveredResourcesResponse, err error) {
    return c.ListDiscoveredResourcesWithContext(context.Background(), request)
}

// ListDiscoveredResources
// This API is used to get the resource list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDiscoveredResourcesWithContext(ctx context.Context, request *ListDiscoveredResourcesRequest) (response *ListDiscoveredResourcesResponse, err error) {
    if request == nil {
        request = NewListDiscoveredResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "ListDiscoveredResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDiscoveredResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDiscoveredResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewPutEvaluationsRequest() (request *PutEvaluationsRequest) {
    request = &PutEvaluationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("config", APIVersion, "PutEvaluations")
    
    
    return
}

func NewPutEvaluationsResponse() (response *PutEvaluationsResponse) {
    response = &PutEvaluationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PutEvaluations
// This API is used to report custom rule evaluation results.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) PutEvaluations(request *PutEvaluationsRequest) (response *PutEvaluationsResponse, err error) {
    return c.PutEvaluationsWithContext(context.Background(), request)
}

// PutEvaluations
// This API is used to report custom rule evaluation results.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"
func (c *Client) PutEvaluationsWithContext(ctx context.Context, request *PutEvaluationsRequest) (response *PutEvaluationsResponse, err error) {
    if request == nil {
        request = NewPutEvaluationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "config", APIVersion, "PutEvaluations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutEvaluations require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutEvaluationsResponse()
    err = c.Send(request, response)
    return
}
