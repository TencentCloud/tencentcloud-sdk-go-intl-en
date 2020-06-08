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

package v20180808

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewBindEnvironmentRequest() (request *BindEnvironmentRequest) {
    request = &BindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindEnvironment")
    return
}

func NewBindEnvironmentResponse() (response *BindEnvironmentResponse) {
    response = &BindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to bind a usage plan to a service environment.
// After you publish a service to an environment, if the API requires authentication and can be called only when it is bound to a usage plan, you can use this API to bind a usage plan to the specified environment.
// Currently, a usage plan can be bound to an API; however, under the same service, usage plans bound to a service and usage plans bound to an API cannot coexist. Therefore, in an environment to which a service-level usage plan has already been bound, please use the `DemoteServiceUsagePlan` API to demote it.
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    if request == nil {
        request = NewBindEnvironmentRequest()
    }
    response = NewBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewBindIPStrategyRequest() (request *BindIPStrategyRequest) {
    request = &BindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindIPStrategy")
    return
}

func NewBindIPStrategyResponse() (response *BindIPStrategyResponse) {
    response = &BindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to bind an IP policy to an API.
func (c *Client) BindIPStrategy(request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    if request == nil {
        request = NewBindIPStrategyRequest()
    }
    response = NewBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewBindSecretIdsRequest() (request *BindSecretIdsRequest) {
    request = &BindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSecretIds")
    return
}

func NewBindSecretIdsResponse() (response *BindSecretIdsResponse) {
    response = &BindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to bind a key to a usage plan.
// You can bind a key to a usage plan and bind the usage plan to an environment where a service is published, so that callers can use the key to call APIs in the service. You can use this API to bind a key to a usage plan.
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    if request == nil {
        request = NewBindSecretIdsRequest()
    }
    response = NewBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewBindSubDomainRequest() (request *BindSubDomainRequest) {
    request = &BindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSubDomain")
    return
}

func NewBindSubDomainResponse() (response *BindSubDomainResponse) {
    response = &BindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to bind a custom domain name to a service.
// Each service in API Gateway provides a default domain name for users to call. If you want to use your own domain name, you can bind a custom domain name to the target service. After getting the ICP filing and configuring the CNAME record between the custom and default domain names, you can directly call the custom domain name.
func (c *Client) BindSubDomain(request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
    if request == nil {
        request = NewBindSubDomainRequest()
    }
    response = NewBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRequest() (request *CreateApiRequest) {
    request = &CreateApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApi")
    return
}

func NewCreateApiResponse() (response *CreateApiResponse) {
    response = &CreateApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an API. Before creating an API, you need to create a service, as each API belongs to a certain service.
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
    if request == nil {
        request = NewCreateApiRequest()
    }
    response = NewCreateApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiKey")
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an API key pair.
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPStrategyRequest() (request *CreateIPStrategyRequest) {
    request = &CreateIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateIPStrategy")
    return
}

func NewCreateIPStrategyResponse() (response *CreateIPStrategyResponse) {
    response = &CreateIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a service IP policy.
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIPStrategyRequest()
    }
    response = NewCreateIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateService")
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a service.
// The maximum unit in API Gateway is service. Multiple APIs can be created in one service, and each service has a default domain name for users to call. You can also bind your own custom domain name to a service.
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsagePlanRequest() (request *CreateUsagePlanRequest) {
    request = &CreateUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateUsagePlan")
    return
}

func NewCreateUsagePlanResponse() (response *CreateUsagePlanResponse) {
    response = &CreateUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a usage plan.
// To use API Gateway, you need to create a usage plan and bind it to a service environment.
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    if request == nil {
        request = NewCreateUsagePlanRequest()
    }
    response = NewCreateUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiRequest() (request *DeleteApiRequest) {
    request = &DeleteApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApi")
    return
}

func NewDeleteApiResponse() (response *DeleteApiResponse) {
    response = &DeleteApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a created API.
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    if request == nil {
        request = NewDeleteApiRequest()
    }
    response = NewDeleteApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiKey")
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an API key pair.
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIPStrategyRequest() (request *DeleteIPStrategyRequest) {
    request = &DeleteIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteIPStrategy")
    return
}

func NewDeleteIPStrategyResponse() (response *DeleteIPStrategyResponse) {
    response = &DeleteIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a service IP policy.
func (c *Client) DeleteIPStrategy(request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteIPStrategyRequest()
    }
    response = NewDeleteIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
    request = &DeleteServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteService")
    return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
    response = &DeleteServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a service in API Gateway.
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
    response = NewDeleteServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceSubDomainMappingRequest() (request *DeleteServiceSubDomainMappingRequest) {
    request = &DeleteServiceSubDomainMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteServiceSubDomainMapping")
    return
}

func NewDeleteServiceSubDomainMappingResponse() (response *DeleteServiceSubDomainMappingResponse) {
    response = &DeleteServiceSubDomainMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a custom domain name mapping in a service environment.
// You can use this API if you use a custom domain name and custom mapping. Please note that if you delete all mappings in all environments, a failure will be returned when this API is called.
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    if request == nil {
        request = NewDeleteServiceSubDomainMappingRequest()
    }
    response = NewDeleteServiceSubDomainMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsagePlanRequest() (request *DeleteUsagePlanRequest) {
    request = &DeleteUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUsagePlan")
    return
}

func NewDeleteUsagePlanResponse() (response *DeleteUsagePlanResponse) {
    response = &DeleteUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a usage plan.
func (c *Client) DeleteUsagePlan(request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
    if request == nil {
        request = NewDeleteUsagePlanRequest()
    }
    response = NewDeleteUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDemoteServiceUsagePlanRequest() (request *DemoteServiceUsagePlanRequest) {
    request = &DemoteServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DemoteServiceUsagePlan")
    return
}

func NewDemoteServiceUsagePlanResponse() (response *DemoteServiceUsagePlanResponse) {
    response = &DemoteServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to demote a usage plan of a service in an environment to the API level.
// To make authentication and throttling for a service take effect, you need to bind a usage plan to it. This API is used to query all usage plans bound to the service.
// This operation will be denied if there are no APIs under a service.
// This operation will also be denied if the current environment has not been published.
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDemoteServiceUsagePlanRequest()
    }
    response = NewDemoteServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRequest() (request *DescribeApiRequest) {
    request = &DescribeApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApi")
    return
}

func NewDescribeApiResponse() (response *DescribeApiResponse) {
    response = &DescribeApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the detailed information of an API deployed in API Gateway.
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    if request == nil {
        request = NewDescribeApiRequest()
    }
    response = NewDescribeApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiEnvironmentStrategyRequest() (request *DescribeApiEnvironmentStrategyRequest) {
    request = &DescribeApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentStrategy")
    return
}

func NewDescribeApiEnvironmentStrategyResponse() (response *DescribeApiEnvironmentStrategyResponse) {
    response = &DescribeApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to display the throttling policies bound to an API.
func (c *Client) DescribeApiEnvironmentStrategy(request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeApiEnvironmentStrategyRequest()
    }
    response = NewDescribeApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
    request = &DescribeApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKey")
    return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
    response = &DescribeApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query key details.
// After creating an API key, you can query its details using this API which will display its key.
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
    response = NewDescribeApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUsagePlanRequest() (request *DescribeApiUsagePlanRequest) {
    request = &DescribeApiUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiUsagePlan")
    return
}

func NewDescribeApiUsagePlanResponse() (response *DescribeApiUsagePlanResponse) {
    response = &DescribeApiUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of API usage plans in a service.
// To make authentication and throttling for a service takes effect, you need to bind usage plans to it. This API is used to query all usage plans bound to a service and APIs under it.
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeApiUsagePlanRequest()
    }
    response = NewDescribeApiUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogSearchRequest() (request *DescribeLogSearchRequest) {
    request = &DescribeLogSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogSearch")
    return
}

func NewDescribeLogSearchResponse() (response *DescribeLogSearchResponse) {
    response = &DescribeLogSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to search for logs.
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    response = NewDescribeLogSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
    request = &DescribeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeService")
    return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
    response = &DescribeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of a service, such as its description, domain name, protocol, creation time, and releases.
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
    response = NewDescribeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentListRequest() (request *DescribeServiceEnvironmentListRequest) {
    request = &DescribeServiceEnvironmentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentList")
    return
}

func NewDescribeServiceEnvironmentListResponse() (response *DescribeServiceEnvironmentListResponse) {
    response = &DescribeServiceEnvironmentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of environments under a service. All environments and their statuses under the queried service will be returned.
func (c *Client) DescribeServiceEnvironmentList(request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentListRequest()
    }
    response = NewDescribeServiceEnvironmentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryRequest() (request *DescribeServiceEnvironmentReleaseHistoryRequest) {
    request = &DescribeServiceEnvironmentReleaseHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentReleaseHistory")
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryResponse() (response *DescribeServiceEnvironmentReleaseHistoryResponse) {
    response = &DescribeServiceEnvironmentReleaseHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the release history of a service environment.
// A service can only be used when it is published to an environment after creation. This API is used to query the release history of an environment under a service.
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
    }
    response = NewDescribeServiceEnvironmentReleaseHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentStrategyRequest() (request *DescribeServiceEnvironmentStrategyRequest) {
    request = &DescribeServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentStrategy")
    return
}

func NewDescribeServiceEnvironmentStrategyResponse() (response *DescribeServiceEnvironmentStrategyResponse) {
    response = &DescribeServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to display service throttling policies.
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentStrategyRequest()
    }
    response = NewDescribeServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceReleaseVersionRequest() (request *DescribeServiceReleaseVersionRequest) {
    request = &DescribeServiceReleaseVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceReleaseVersion")
    return
}

func NewDescribeServiceReleaseVersionResponse() (response *DescribeServiceReleaseVersionResponse) {
    response = &DescribeServiceReleaseVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of all published versions under a service.
// A service is generally published on several versions. This API can be used to query the published versions.
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    if request == nil {
        request = NewDescribeServiceReleaseVersionRequest()
    }
    response = NewDescribeServiceReleaseVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainMappingsRequest() (request *DescribeServiceSubDomainMappingsRequest) {
    request = &DescribeServiceSubDomainMappingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomainMappings")
    return
}

func NewDescribeServiceSubDomainMappingsResponse() (response *DescribeServiceSubDomainMappingsResponse) {
    response = &DescribeServiceSubDomainMappingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the path mappings of a custom domain name.
// In API Gateway, you can bind a custom domain name to a service and map the paths for it. You can custom different path mappings to up to 3 environments under the service. This API is used to query the list of path mappings of a custom domain name bound to a service.
func (c *Client) DescribeServiceSubDomainMappings(request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainMappingsRequest()
    }
    response = NewDescribeServiceSubDomainMappingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainsRequest() (request *DescribeServiceSubDomainsRequest) {
    request = &DescribeServiceSubDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomains")
    return
}

func NewDescribeServiceSubDomainsResponse() (response *DescribeServiceSubDomainsResponse) {
    response = &DescribeServiceSubDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of custom domain names.
// In API Gateway, you can bind custom domain names to a service for service call. This API is used to query the list of custom domain names bound to a service.
func (c *Client) DescribeServiceSubDomains(request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainsRequest()
    }
    response = NewDescribeServiceSubDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceUsagePlanRequest() (request *DescribeServiceUsagePlanRequest) {
    request = &DescribeServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceUsagePlan")
    return
}

func NewDescribeServiceUsagePlanResponse() (response *DescribeServiceUsagePlanResponse) {
    response = &DescribeServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of usage plans in a service.
// To make authentication and throttling for a service take effect, you need to bind a usage plan to it. This API is used to query all usage plans bound to the service.
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeServiceUsagePlanRequest()
    }
    response = NewDescribeServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanRequest() (request *DescribeUsagePlanRequest) {
    request = &DescribeUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlan")
    return
}

func NewDescribeUsagePlanResponse() (response *DescribeUsagePlanResponse) {
    response = &DescribeUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of a usage plan, such as its name, QPS, creation time, and bound environment.
func (c *Client) DescribeUsagePlan(request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanRequest()
    }
    response = NewDescribeUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanEnvironmentsRequest() (request *DescribeUsagePlanEnvironmentsRequest) {
    request = &DescribeUsagePlanEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanEnvironments")
    return
}

func NewDescribeUsagePlanEnvironmentsResponse() (response *DescribeUsagePlanEnvironmentsResponse) {
    response = &DescribeUsagePlanEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of environments bound to a usage plan.
// After binding a usage plan to environments, you can use this API to query all service environments bound to the usage plan.
func (c *Client) DescribeUsagePlanEnvironments(request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanEnvironmentsRequest()
    }
    response = NewDescribeUsagePlanEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanSecretIdsRequest() (request *DescribeUsagePlanSecretIdsRequest) {
    request = &DescribeUsagePlanSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanSecretIds")
    return
}

func NewDescribeUsagePlanSecretIdsResponse() (response *DescribeUsagePlanSecretIdsResponse) {
    response = &DescribeUsagePlanSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of keys bound to a usage plan.
// In API Gateway, a usage plan can be bound to multiple key pairs. You can use this API to query the list of keys bound to it.
func (c *Client) DescribeUsagePlanSecretIds(request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanSecretIdsRequest()
    }
    response = NewDescribeUsagePlanSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableApiKeyRequest() (request *DisableApiKeyRequest) {
    request = &DisableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "DisableApiKey")
    return
}

func NewDisableApiKeyResponse() (response *DisableApiKeyResponse) {
    response = &DisableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable an API key pair.
func (c *Client) DisableApiKey(request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
    if request == nil {
        request = NewDisableApiKeyRequest()
    }
    response = NewDisableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableApiKeyRequest() (request *EnableApiKeyRequest) {
    request = &EnableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "EnableApiKey")
    return
}

func NewEnableApiKeyResponse() (response *EnableApiKeyResponse) {
    response = &EnableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable a disabled API key pair.
func (c *Client) EnableApiKey(request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
    if request == nil {
        request = NewEnableApiKeyRequest()
    }
    response = NewEnableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateApiDocumentRequest() (request *GenerateApiDocumentRequest) {
    request = &GenerateApiDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "GenerateApiDocument")
    return
}

func NewGenerateApiDocumentResponse() (response *GenerateApiDocumentResponse) {
    response = &GenerateApiDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to automatically generate API documentation and SDKs. A document and an SDK will be generated for each environment under each service, respectively.
func (c *Client) GenerateApiDocument(request *GenerateApiDocumentRequest) (response *GenerateApiDocumentResponse, err error) {
    if request == nil {
        request = NewGenerateApiDocumentRequest()
    }
    response = NewGenerateApiDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiRequest() (request *ModifyApiRequest) {
    request = &ModifyApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApi")
    return
}

func NewModifyApiResponse() (response *ModifyApiResponse) {
    response = &ModifyApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify an API. You can call this API to edit/modify a configured API. The modified API takes effect only after its service is published to the corresponding environment again.
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    if request == nil {
        request = NewModifyApiRequest()
    }
    response = NewModifyApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiEnvironmentStrategyRequest() (request *ModifyApiEnvironmentStrategyRequest) {
    request = &ModifyApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiEnvironmentStrategy")
    return
}

func NewModifyApiEnvironmentStrategyResponse() (response *ModifyApiEnvironmentStrategyResponse) {
    response = &ModifyApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify an API throttling policy.
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyApiEnvironmentStrategyRequest()
    }
    response = NewModifyApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiIncrementRequest() (request *ModifyApiIncrementRequest) {
    request = &ModifyApiIncrementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiIncrement")
    return
}

func NewModifyApiIncrementResponse() (response *ModifyApiIncrementResponse) {
    response = &ModifyApiIncrementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to incrementally update an API and mainly called by programs (different from `ModifyApi`, which requires full API parameters be passed in and is suitable for being used with the console).
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
    response = NewModifyApiIncrementResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIPStrategyRequest() (request *ModifyIPStrategyRequest) {
    request = &ModifyIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyIPStrategy")
    return
}

func NewModifyIPStrategyResponse() (response *ModifyIPStrategyResponse) {
    response = &ModifyIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a service IP policy.
func (c *Client) ModifyIPStrategy(request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
    if request == nil {
        request = NewModifyIPStrategyRequest()
    }
    response = NewModifyIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
    request = &ModifyServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyService")
    return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
    response = &ModifyServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the relevant information of a service. After a service is created, its name, description, and service type all can be modified.
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    if request == nil {
        request = NewModifyServiceRequest()
    }
    response = NewModifyServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceEnvironmentStrategyRequest() (request *ModifyServiceEnvironmentStrategyRequest) {
    request = &ModifyServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentStrategy")
    return
}

func NewModifyServiceEnvironmentStrategyResponse() (response *ModifyServiceEnvironmentStrategyResponse) {
    response = &ModifyServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a service throttling policy.
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentStrategyRequest()
    }
    response = NewModifyServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubDomainRequest() (request *ModifySubDomainRequest) {
    request = &ModifySubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifySubDomain")
    return
}

func NewModifySubDomainResponse() (response *ModifySubDomainResponse) {
    response = &ModifySubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the path mapping in the custom domain name settings of the service. The path mapping rule can be modified before it is bound to the custom domain name.
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    if request == nil {
        request = NewModifySubDomainRequest()
    }
    response = NewModifySubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUsagePlanRequest() (request *ModifyUsagePlanRequest) {
    request = &ModifyUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUsagePlan")
    return
}

func NewModifyUsagePlanResponse() (response *ModifyUsagePlanResponse) {
    response = &ModifyUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the name, description, and QPS of a usage plan.
func (c *Client) ModifyUsagePlan(request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
    if request == nil {
        request = NewModifyUsagePlanRequest()
    }
    response = NewModifyUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseServiceRequest() (request *ReleaseServiceRequest) {
    request = &ReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "ReleaseService")
    return
}

func NewReleaseServiceResponse() (response *ReleaseServiceResponse) {
    response = &ReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to publish a service.
// An API Gateway service can only be called when it is published to an environment and takes effect after creation. This API is used to publish a service to an environment such as the `release` environment.
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    if request == nil {
        request = NewReleaseServiceRequest()
    }
    response = NewReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindEnvironmentRequest() (request *UnBindEnvironmentRequest) {
    request = &UnBindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindEnvironment")
    return
}

func NewUnBindEnvironmentResponse() (response *UnBindEnvironmentResponse) {
    response = &UnBindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unbind a usage plan from a specified environment.
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    if request == nil {
        request = NewUnBindEnvironmentRequest()
    }
    response = NewUnBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindIPStrategyRequest() (request *UnBindIPStrategyRequest) {
    request = &UnBindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindIPStrategy")
    return
}

func NewUnBindIPStrategyResponse() (response *UnBindIPStrategyResponse) {
    response = &UnBindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unbind an IP policy from a service.
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    if request == nil {
        request = NewUnBindIPStrategyRequest()
    }
    response = NewUnBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSecretIdsRequest() (request *UnBindSecretIdsRequest) {
    request = &UnBindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSecretIds")
    return
}

func NewUnBindSecretIdsResponse() (response *UnBindSecretIdsResponse) {
    response = &UnBindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unbind a key from a usage plan.
func (c *Client) UnBindSecretIds(request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    if request == nil {
        request = NewUnBindSecretIdsRequest()
    }
    response = NewUnBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSubDomainRequest() (request *UnBindSubDomainRequest) {
    request = &UnBindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSubDomain")
    return
}

func NewUnBindSubDomainResponse() (response *UnBindSubDomainResponse) {
    response = &UnBindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unbind a custom domain name.
// After binding a custom domain name to a service using API Gateway, you can use this API to unbind it.
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    if request == nil {
        request = NewUnBindSubDomainRequest()
    }
    response = NewUnBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUnReleaseServiceRequest() (request *UnReleaseServiceRequest) {
    request = &UnReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UnReleaseService")
    return
}

func NewUnReleaseServiceResponse() (response *UnReleaseServiceResponse) {
    response = &UnReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to deactivate a service.
// Only when a service is published to an environment can its APIs be called by callers. You can call this API to deactivate a service from the release environment. Once deactivated, the service cannot be called.
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    if request == nil {
        request = NewUnReleaseServiceRequest()
    }
    response = NewUnReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiKeyRequest() (request *UpdateApiKeyRequest) {
    request = &UpdateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiKey")
    return
}

func NewUpdateApiKeyResponse() (response *UpdateApiKeyResponse) {
    response = &UpdateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update a created API key pair.
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiKeyRequest()
    }
    response = NewUpdateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
    request = &UpdateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateService")
    return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
    response = &UpdateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to switch the running version published in an environment of a service to a specified version. After you create a service by using API Gateway and publish it to an environment, multiple versions will be generated during development. In this case, you can call this API to switch versions.
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
