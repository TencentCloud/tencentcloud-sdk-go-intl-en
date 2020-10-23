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

package v20191012

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-10-12"

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


func NewAddEcdnDomainRequest() (request *AddEcdnDomainRequest) {
    request = &AddEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "AddEcdnDomain")
    return
}

func NewAddEcdnDomainResponse() (response *AddEcdnDomainResponse) {
    response = &AddEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an acceleration domain name.
func (c *Client) AddEcdnDomain(request *AddEcdnDomainRequest) (response *AddEcdnDomainResponse, err error) {
    if request == nil {
        request = NewAddEcdnDomainRequest()
    }
    response = NewAddEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEcdnDomainRequest() (request *DeleteEcdnDomainRequest) {
    request = &DeleteEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DeleteEcdnDomain")
    return
}

func NewDeleteEcdnDomainResponse() (response *DeleteEcdnDomainResponse) {
    response = &DeleteEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a specified acceleration domain name. The acceleration domain name to be deleted must be in disabled status.
func (c *Client) DeleteEcdnDomain(request *DeleteEcdnDomainRequest) (response *DeleteEcdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteEcdnDomainRequest()
    }
    response = NewDeleteEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomains")
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the basic information of a CDN domain name, including the project ID, status, business type, creation time, update time, etc.
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsConfigRequest() (request *DescribeDomainsConfigRequest) {
    request = &DescribeDomainsConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomainsConfig")
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the detailed configuration information of a CDN acceleration domain name.
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainLogsRequest() (request *DescribeEcdnDomainLogsRequest) {
    request = &DescribeEcdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainLogs")
    return
}

func NewDescribeEcdnDomainLogsResponse() (response *DescribeEcdnDomainLogsResponse) {
    response = &DescribeEcdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the access log download link of a domain name.
func (c *Client) DescribeEcdnDomainLogs(request *DescribeEcdnDomainLogsRequest) (response *DescribeEcdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainLogsRequest()
    }
    response = NewDescribeEcdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainStatisticsRequest() (request *DescribeEcdnDomainStatisticsRequest) {
    request = &DescribeEcdnDomainStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainStatistics")
    return
}

func NewDescribeEcdnDomainStatisticsResponse() (response *DescribeEcdnDomainStatisticsResponse) {
    response = &DescribeEcdnDomainStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the statistical metrics of domain name access within a specified time period.
func (c *Client) DescribeEcdnDomainStatistics(request *DescribeEcdnDomainStatisticsRequest) (response *DescribeEcdnDomainStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainStatisticsRequest()
    }
    response = NewDescribeEcdnDomainStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnStatisticsRequest() (request *DescribeEcdnStatisticsRequest) {
    request = &DescribeEcdnStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnStatistics")
    return
}

func NewDescribeEcdnStatisticsResponse() (response *DescribeEcdnStatisticsResponse) {
    response = &DescribeEcdnStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query ECDN real-time access monitoring data and supports the following metrics:
// 
// + Traffic (in bytes)
// + Bandwidth (in bps)
// + Number of requests
// + Response time (in ms)
// + Number of 2xx status codes and details of status codes starting with 2
// + Number of 3xx status codes and details of status codes starting with 3
// + Number of 4xx status codes and details of status codes starting with 4
// + Number of 5xx status codes and details of status codes starting with 5
func (c *Client) DescribeEcdnStatistics(request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnStatisticsRequest()
    }
    response = NewDescribeEcdnStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpStatusRequest() (request *DescribeIpStatusRequest) {
    request = &DescribeIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeIpStatus")
    return
}

func NewDescribeIpStatusResponse() (response *DescribeIpStatusResponse) {
    response = &DescribeIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the detailed node information of the acceleration platform to which the domain name is connected.
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    response = NewDescribeIpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeQuotaRequest() (request *DescribePurgeQuotaRequest) {
    request = &DescribePurgeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeQuota")
    return
}

func NewDescribePurgeQuotaResponse() (response *DescribePurgeQuotaResponse) {
    response = &DescribePurgeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the usage quota of the purge API.
func (c *Client) DescribePurgeQuota(request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePurgeQuotaRequest()
    }
    response = NewDescribePurgeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeTasks")
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the submission history of purge tasks and their execution progress.
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewPurgePathCacheRequest() (request *PurgePathCacheRequest) {
    request = &PurgePathCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgePathCache")
    return
}

func NewPurgePathCacheResponse() (response *PurgePathCacheResponse) {
    response = &PurgePathCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to purge cache directories in batches. One purge task ID will be returned for each submission.
func (c *Client) PurgePathCache(request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    if request == nil {
        request = NewPurgePathCacheRequest()
    }
    response = NewPurgePathCacheResponse()
    err = c.Send(request, response)
    return
}

func NewPurgeUrlsCacheRequest() (request *PurgeUrlsCacheRequest) {
    request = &PurgeUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgeUrlsCache")
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to batch purge URLs. One purge task ID will be returned for each submission.
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewStartEcdnDomainRequest() (request *StartEcdnDomainRequest) {
    request = &StartEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "StartEcdnDomain")
    return
}

func NewStartEcdnDomainResponse() (response *StartEcdnDomainResponse) {
    response = &StartEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable an acceleration domain name. The domain name to be enabled must be in deactivated status.
func (c *Client) StartEcdnDomain(request *StartEcdnDomainRequest) (response *StartEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStartEcdnDomainRequest()
    }
    response = NewStartEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopEcdnDomainRequest() (request *StopEcdnDomainRequest) {
    request = &StopEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "StopEcdnDomain")
    return
}

func NewStopEcdnDomainResponse() (response *StopEcdnDomainResponse) {
    response = &StopEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable an acceleration domain name. The domain name to be disabled must be in enabled or deploying status.
func (c *Client) StopEcdnDomain(request *StopEcdnDomainRequest) (response *StopEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStopEcdnDomainRequest()
    }
    response = NewStopEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDomainConfigRequest() (request *UpdateDomainConfigRequest) {
    request = &UpdateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "UpdateDomainConfig")
    return
}

func NewUpdateDomainConfigResponse() (response *UpdateDomainConfigResponse) {
    response = &UpdateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the configuration information of an ECDN acceleration domain name.
// Note: if you need to update a complex configuration item, you must pass in all attributes of the entire object, and the default values will be used for the attributes that are not passed in. You are recommended to get the configuration attribute through the query API first and then directly modify and pass it to this API. Due to the special nature of the certificate for HTTPS configuration, you do not need to pass in the certificate and key fields during the update.
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}
