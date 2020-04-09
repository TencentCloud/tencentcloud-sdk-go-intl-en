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

package v20180228

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-02-28"

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


func NewAddLoginWhiteListRequest() (request *AddLoginWhiteListRequest) {
    request = &AddLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "AddLoginWhiteList")
    return
}

func NewAddLoginWhiteListResponse() (response *AddLoginWhiteListResponse) {
    response = &AddLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a whitelist rule.
func (c *Client) AddLoginWhiteList(request *AddLoginWhiteListRequest) (response *AddLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewAddLoginWhiteListRequest()
    }
    response = NewAddLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewAddMachineTagRequest() (request *AddMachineTagRequest) {
    request = &AddMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "AddMachineTag")
    return
}

func NewAddMachineTagResponse() (response *AddMachineTagResponse) {
    response = &AddMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a tag to a server.
func (c *Client) AddMachineTag(request *AddMachineTagRequest) (response *AddMachineTagResponse, err error) {
    if request == nil {
        request = NewAddMachineTagRequest()
    }
    response = NewAddMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProVersionRequest() (request *CloseProVersionRequest) {
    request = &CloseProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CloseProVersion")
    return
}

func NewCloseProVersionResponse() (response *CloseProVersionResponse) {
    response = &CloseProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to deactivate CWP Pro.
func (c *Client) CloseProVersion(request *CloseProVersionRequest) (response *CloseProVersionResponse, err error) {
    if request == nil {
        request = NewCloseProVersionRequest()
    }
    response = NewCloseProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpenPortTaskRequest() (request *CreateOpenPortTaskRequest) {
    request = &CreateOpenPortTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateOpenPortTask")
    return
}

func NewCreateOpenPortTaskResponse() (response *CreateOpenPortTaskResponse) {
    response = &CreateOpenPortTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a real-time port acquisition task.
func (c *Client) CreateOpenPortTask(request *CreateOpenPortTaskRequest) (response *CreateOpenPortTaskResponse, err error) {
    if request == nil {
        request = NewCreateOpenPortTaskRequest()
    }
    response = NewCreateOpenPortTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProcessTaskRequest() (request *CreateProcessTaskRequest) {
    request = &CreateProcessTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateProcessTask")
    return
}

func NewCreateProcessTaskResponse() (response *CreateProcessTaskResponse) {
    response = &CreateProcessTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a real-time process pulling task.
func (c *Client) CreateProcessTask(request *CreateProcessTaskRequest) (response *CreateProcessTaskResponse, err error) {
    if request == nil {
        request = NewCreateProcessTaskRequest()
    }
    response = NewCreateProcessTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsualLoginPlacesRequest() (request *CreateUsualLoginPlacesRequest) {
    request = &CreateUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "CreateUsualLoginPlaces")
    return
}

func NewCreateUsualLoginPlacesResponse() (response *CreateUsualLoginPlacesResponse) {
    response = &CreateUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add one or more usual login locations.
func (c *Client) CreateUsualLoginPlaces(request *CreateUsualLoginPlacesRequest) (response *CreateUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewCreateUsualLoginPlacesRequest()
    }
    response = NewCreateUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBruteAttacksRequest() (request *DeleteBruteAttacksRequest) {
    request = &DeleteBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteBruteAttacks")
    return
}

func NewDeleteBruteAttacksResponse() (response *DeleteBruteAttacksResponse) {
    response = &DeleteBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete brute force attack records.
func (c *Client) DeleteBruteAttacks(request *DeleteBruteAttacksRequest) (response *DeleteBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDeleteBruteAttacksRequest()
    }
    response = NewDeleteBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoginWhiteListRequest() (request *DeleteLoginWhiteListRequest) {
    request = &DeleteLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteLoginWhiteList")
    return
}

func NewDeleteLoginWhiteListResponse() (response *DeleteLoginWhiteListResponse) {
    response = &DeleteLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a whitelist rule.
func (c *Client) DeleteLoginWhiteList(request *DeleteLoginWhiteListRequest) (response *DeleteLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteLoginWhiteListRequest()
    }
    response = NewDeleteLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineRequest() (request *DeleteMachineRequest) {
    request = &DeleteMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachine")
    return
}

func NewDeleteMachineResponse() (response *DeleteMachineResponse) {
    response = &DeleteMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to uninstall the CWP agent.
func (c *Client) DeleteMachine(request *DeleteMachineRequest) (response *DeleteMachineResponse, err error) {
    if request == nil {
        request = NewDeleteMachineRequest()
    }
    response = NewDeleteMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineTagRequest() (request *DeleteMachineTagRequest) {
    request = &DeleteMachineTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMachineTag")
    return
}

func NewDeleteMachineTagResponse() (response *DeleteMachineTagResponse) {
    response = &DeleteMachineTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to remove a tag from a server.
func (c *Client) DeleteMachineTag(request *DeleteMachineTagRequest) (response *DeleteMachineTagResponse, err error) {
    if request == nil {
        request = NewDeleteMachineTagRequest()
    }
    response = NewDeleteMachineTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMaliciousRequestsRequest() (request *DeleteMaliciousRequestsRequest) {
    request = &DeleteMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMaliciousRequests")
    return
}

func NewDeleteMaliciousRequestsResponse() (response *DeleteMaliciousRequestsResponse) {
    response = &DeleteMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete malicious request records.
func (c *Client) DeleteMaliciousRequests(request *DeleteMaliciousRequestsRequest) (response *DeleteMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDeleteMaliciousRequestsRequest()
    }
    response = NewDeleteMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMalwaresRequest() (request *DeleteMalwaresRequest) {
    request = &DeleteMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteMalwares")
    return
}

func NewDeleteMalwaresResponse() (response *DeleteMalwaresResponse) {
    response = &DeleteMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete trojan records.
func (c *Client) DeleteMalwares(request *DeleteMalwaresRequest) (response *DeleteMalwaresResponse, err error) {
    if request == nil {
        request = NewDeleteMalwaresRequest()
    }
    response = NewDeleteMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNonlocalLoginPlacesRequest() (request *DeleteNonlocalLoginPlacesRequest) {
    request = &DeleteNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteNonlocalLoginPlaces")
    return
}

func NewDeleteNonlocalLoginPlacesResponse() (response *DeleteNonlocalLoginPlacesResponse) {
    response = &DeleteNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete unusual login location records.
func (c *Client) DeleteNonlocalLoginPlaces(request *DeleteNonlocalLoginPlacesRequest) (response *DeleteNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteNonlocalLoginPlacesRequest()
    }
    response = NewDeleteNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsualLoginPlacesRequest() (request *DeleteUsualLoginPlacesRequest) {
    request = &DeleteUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DeleteUsualLoginPlaces")
    return
}

func NewDeleteUsualLoginPlacesResponse() (response *DeleteUsualLoginPlacesResponse) {
    response = &DeleteUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete one or more usual login locations.
func (c *Client) DeleteUsualLoginPlaces(request *DeleteUsualLoginPlacesRequest) (response *DeleteUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDeleteUsualLoginPlacesRequest()
    }
    response = NewDeleteUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountStatisticsRequest() (request *DescribeAccountStatisticsRequest) {
    request = &DescribeAccountStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccountStatistics")
    return
}

func NewDescribeAccountStatisticsResponse() (response *DescribeAccountStatisticsResponse) {
    response = &DescribeAccountStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the account statistics list.
func (c *Client) DescribeAccountStatistics(request *DescribeAccountStatisticsRequest) (response *DescribeAccountStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountStatisticsRequest()
    }
    response = NewDescribeAccountStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the account list.
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentVulsRequest() (request *DescribeAgentVulsRequest) {
    request = &DescribeAgentVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAgentVuls")
    return
}

func NewDescribeAgentVulsResponse() (response *DescribeAgentVulsResponse) {
    response = &DescribeAgentVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of vulnerabilities on a single server.
func (c *Client) DescribeAgentVuls(request *DescribeAgentVulsRequest) (response *DescribeAgentVulsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVulsRequest()
    }
    response = NewDescribeAgentVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmAttributeRequest() (request *DescribeAlarmAttributeRequest) {
    request = &DescribeAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeAlarmAttribute")
    return
}

func NewDescribeAlarmAttributeResponse() (response *DescribeAlarmAttributeResponse) {
    response = &DescribeAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the alarm settings.
func (c *Client) DescribeAlarmAttribute(request *DescribeAlarmAttributeRequest) (response *DescribeAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmAttributeRequest()
    }
    response = NewDescribeAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBruteAttacksRequest() (request *DescribeBruteAttacksRequest) {
    request = &DescribeBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeBruteAttacks")
    return
}

func NewDescribeBruteAttacksResponse() (response *DescribeBruteAttacksResponse) {
    response = &DescribeBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the brute force attack event list.
func (c *Client) DescribeBruteAttacks(request *DescribeBruteAttacksRequest) (response *DescribeBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeBruteAttacksRequest()
    }
    response = NewDescribeBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentInfoRequest() (request *DescribeComponentInfoRequest) {
    request = &DescribeComponentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentInfo")
    return
}

func NewDescribeComponentInfoResponse() (response *DescribeComponentInfoResponse) {
    response = &DescribeComponentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the component information.
func (c *Client) DescribeComponentInfo(request *DescribeComponentInfoRequest) (response *DescribeComponentInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComponentInfoRequest()
    }
    response = NewDescribeComponentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentStatisticsRequest() (request *DescribeComponentStatisticsRequest) {
    request = &DescribeComponentStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponentStatistics")
    return
}

func NewDescribeComponentStatisticsResponse() (response *DescribeComponentStatisticsResponse) {
    response = &DescribeComponentStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the component statistics list.
func (c *Client) DescribeComponentStatistics(request *DescribeComponentStatisticsRequest) (response *DescribeComponentStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentStatisticsRequest()
    }
    response = NewDescribeComponentStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComponentsRequest() (request *DescribeComponentsRequest) {
    request = &DescribeComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeComponents")
    return
}

func NewDescribeComponentsResponse() (response *DescribeComponentsResponse) {
    response = &DescribeComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the component list.
func (c *Client) DescribeComponents(request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeComponentsRequest()
    }
    response = NewDescribeComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryAccountsRequest() (request *DescribeHistoryAccountsRequest) {
    request = &DescribeHistoryAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeHistoryAccounts")
    return
}

func NewDescribeHistoryAccountsResponse() (response *DescribeHistoryAccountsResponse) {
    response = &DescribeHistoryAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the account change history list.
func (c *Client) DescribeHistoryAccounts(request *DescribeHistoryAccountsRequest) (response *DescribeHistoryAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryAccountsRequest()
    }
    response = NewDescribeHistoryAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImpactedHostsRequest() (request *DescribeImpactedHostsRequest) {
    request = &DescribeImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeImpactedHosts")
    return
}

func NewDescribeImpactedHostsResponse() (response *DescribeImpactedHostsResponse) {
    response = &DescribeImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of servers affected by a vulnerability.
func (c *Client) DescribeImpactedHosts(request *DescribeImpactedHostsRequest) (response *DescribeImpactedHostsResponse, err error) {
    if request == nil {
        request = NewDescribeImpactedHostsRequest()
    }
    response = NewDescribeImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoginWhiteListRequest() (request *DescribeLoginWhiteListRequest) {
    request = &DescribeLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeLoginWhiteList")
    return
}

func NewDescribeLoginWhiteListResponse() (response *DescribeLoginWhiteListResponse) {
    response = &DescribeLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of login whitelist entries.
func (c *Client) DescribeLoginWhiteList(request *DescribeLoginWhiteListRequest) (response *DescribeLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeLoginWhiteListRequest()
    }
    response = NewDescribeLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineInfoRequest() (request *DescribeMachineInfoRequest) {
    request = &DescribeMachineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachineInfo")
    return
}

func NewDescribeMachineInfoResponse() (response *DescribeMachineInfoResponse) {
    response = &DescribeMachineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get server details.
func (c *Client) DescribeMachineInfo(request *DescribeMachineInfoRequest) (response *DescribeMachineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMachineInfoRequest()
    }
    response = NewDescribeMachineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMachines")
    return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
    response = &DescribeMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of servers in a specified region.
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaliciousRequestsRequest() (request *DescribeMaliciousRequestsRequest) {
    request = &DescribeMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMaliciousRequests")
    return
}

func NewDescribeMaliciousRequestsResponse() (response *DescribeMaliciousRequestsResponse) {
    response = &DescribeMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get malicious request data.
func (c *Client) DescribeMaliciousRequests(request *DescribeMaliciousRequestsRequest) (response *DescribeMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewDescribeMaliciousRequestsRequest()
    }
    response = NewDescribeMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMalwaresRequest() (request *DescribeMalwaresRequest) {
    request = &DescribeMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeMalwares")
    return
}

func NewDescribeMalwaresResponse() (response *DescribeMalwaresResponse) {
    response = &DescribeMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of trojan events.
func (c *Client) DescribeMalwares(request *DescribeMalwaresRequest) (response *DescribeMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeMalwaresRequest()
    }
    response = NewDescribeMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNonlocalLoginPlacesRequest() (request *DescribeNonlocalLoginPlacesRequest) {
    request = &DescribeNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeNonlocalLoginPlaces")
    return
}

func NewDescribeNonlocalLoginPlacesResponse() (response *DescribeNonlocalLoginPlacesResponse) {
    response = &DescribeNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get unusual login events.
func (c *Client) DescribeNonlocalLoginPlaces(request *DescribeNonlocalLoginPlacesRequest) (response *DescribeNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeNonlocalLoginPlacesRequest()
    }
    response = NewDescribeNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortStatisticsRequest() (request *DescribeOpenPortStatisticsRequest) {
    request = &DescribeOpenPortStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortStatistics")
    return
}

func NewDescribeOpenPortStatisticsResponse() (response *DescribeOpenPortStatisticsResponse) {
    response = &DescribeOpenPortStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the statistics on port usage
func (c *Client) DescribeOpenPortStatistics(request *DescribeOpenPortStatisticsRequest) (response *DescribeOpenPortStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortStatisticsRequest()
    }
    response = NewDescribeOpenPortStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortTaskStatusRequest() (request *DescribeOpenPortTaskStatusRequest) {
    request = &DescribeOpenPortTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPortTaskStatus")
    return
}

func NewDescribeOpenPortTaskStatusResponse() (response *DescribeOpenPortTaskStatusResponse) {
    response = &DescribeOpenPortTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the status of a real-time port pulling task.
func (c *Client) DescribeOpenPortTaskStatus(request *DescribeOpenPortTaskStatusRequest) (response *DescribeOpenPortTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortTaskStatusRequest()
    }
    response = NewDescribeOpenPortTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpenPortsRequest() (request *DescribeOpenPortsRequest) {
    request = &DescribeOpenPortsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOpenPorts")
    return
}

func NewDescribeOpenPortsResponse() (response *DescribeOpenPortsResponse) {
    response = &DescribeOpenPortsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the port list.
func (c *Client) DescribeOpenPorts(request *DescribeOpenPortsRequest) (response *DescribeOpenPortsResponse, err error) {
    if request == nil {
        request = NewDescribeOpenPortsRequest()
    }
    response = NewDescribeOpenPortsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewStatisticsRequest() (request *DescribeOverviewStatisticsRequest) {
    request = &DescribeOverviewStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeOverviewStatistics")
    return
}

func NewDescribeOverviewStatisticsResponse() (response *DescribeOverviewStatisticsResponse) {
    response = &DescribeOverviewStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the overview statistics of CWP under the current account.
func (c *Client) DescribeOverviewStatistics(request *DescribeOverviewStatisticsRequest) (response *DescribeOverviewStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewStatisticsRequest()
    }
    response = NewDescribeOverviewStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProVersionInfoRequest() (request *DescribeProVersionInfoRequest) {
    request = &DescribeProVersionInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProVersionInfo")
    return
}

func NewDescribeProVersionInfoResponse() (response *DescribeProVersionInfoResponse) {
    response = &DescribeProVersionInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the CWP Pro information.
func (c *Client) DescribeProVersionInfo(request *DescribeProVersionInfoRequest) (response *DescribeProVersionInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProVersionInfoRequest()
    }
    response = NewDescribeProVersionInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessStatisticsRequest() (request *DescribeProcessStatisticsRequest) {
    request = &DescribeProcessStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessStatistics")
    return
}

func NewDescribeProcessStatisticsResponse() (response *DescribeProcessStatisticsResponse) {
    response = &DescribeProcessStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the process statistics list.
func (c *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (response *DescribeProcessStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProcessStatisticsRequest()
    }
    response = NewDescribeProcessStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessTaskStatusRequest() (request *DescribeProcessTaskStatusRequest) {
    request = &DescribeProcessTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcessTaskStatus")
    return
}

func NewDescribeProcessTaskStatusResponse() (response *DescribeProcessTaskStatusResponse) {
    response = &DescribeProcessTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the status of a real-time process pulling task.
func (c *Client) DescribeProcessTaskStatus(request *DescribeProcessTaskStatusRequest) (response *DescribeProcessTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeProcessTaskStatusRequest()
    }
    response = NewDescribeProcessTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProcessesRequest() (request *DescribeProcessesRequest) {
    request = &DescribeProcessesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeProcesses")
    return
}

func NewDescribeProcessesResponse() (response *DescribeProcessesResponse) {
    response = &DescribeProcessesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the process list.
func (c *Client) DescribeProcesses(request *DescribeProcessesRequest) (response *DescribeProcessesResponse, err error) {
    if request == nil {
        request = NewDescribeProcessesRequest()
    }
    response = NewDescribeProcessesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityDynamicsRequest() (request *DescribeSecurityDynamicsRequest) {
    request = &DescribeSecurityDynamicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityDynamics")
    return
}

func NewDescribeSecurityDynamicsResponse() (response *DescribeSecurityDynamicsResponse) {
    response = &DescribeSecurityDynamicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the security event message data.
func (c *Client) DescribeSecurityDynamics(request *DescribeSecurityDynamicsRequest) (response *DescribeSecurityDynamicsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityDynamicsRequest()
    }
    response = NewDescribeSecurityDynamicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityTrendsRequest() (request *DescribeSecurityTrendsRequest) {
    request = &DescribeSecurityTrendsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeSecurityTrends")
    return
}

func NewDescribeSecurityTrendsResponse() (response *DescribeSecurityTrendsResponse) {
    response = &DescribeSecurityTrendsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the security event statistics.
func (c *Client) DescribeSecurityTrends(request *DescribeSecurityTrendsRequest) (response *DescribeSecurityTrendsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTrendsRequest()
    }
    response = NewDescribeSecurityTrendsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagMachinesRequest() (request *DescribeTagMachinesRequest) {
    request = &DescribeTagMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTagMachines")
    return
}

func NewDescribeTagMachinesResponse() (response *DescribeTagMachinesResponse) {
    response = &DescribeTagMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the information of servers associated with a specified tag.
func (c *Client) DescribeTagMachines(request *DescribeTagMachinesRequest) (response *DescribeTagMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeTagMachinesRequest()
    }
    response = NewDescribeTagMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
    request = &DescribeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeTags")
    return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
    response = &DescribeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get all server tags.
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsRequest()
    }
    response = NewDescribeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsualLoginPlacesRequest() (request *DescribeUsualLoginPlacesRequest) {
    request = &DescribeUsualLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeUsualLoginPlaces")
    return
}

func NewDescribeUsualLoginPlacesResponse() (response *DescribeUsualLoginPlacesResponse) {
    response = &DescribeUsualLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query usual login locations.
func (c *Client) DescribeUsualLoginPlaces(request *DescribeUsualLoginPlacesRequest) (response *DescribeUsualLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsualLoginPlacesRequest()
    }
    response = NewDescribeUsualLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulInfoRequest() (request *DescribeVulInfoRequest) {
    request = &DescribeVulInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulInfo")
    return
}

func NewDescribeVulInfoResponse() (response *DescribeVulInfoResponse) {
    response = &DescribeVulInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get vulnerability details.
func (c *Client) DescribeVulInfo(request *DescribeVulInfoRequest) (response *DescribeVulInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVulInfoRequest()
    }
    response = NewDescribeVulInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulScanResultRequest() (request *DescribeVulScanResultRequest) {
    request = &DescribeVulScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVulScanResult")
    return
}

func NewDescribeVulScanResultResponse() (response *DescribeVulScanResultResponse) {
    response = &DescribeVulScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the vulnerability detection result.
func (c *Client) DescribeVulScanResult(request *DescribeVulScanResultRequest) (response *DescribeVulScanResultResponse, err error) {
    if request == nil {
        request = NewDescribeVulScanResultRequest()
    }
    response = NewDescribeVulScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulsRequest() (request *DescribeVulsRequest) {
    request = &DescribeVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeVuls")
    return
}

func NewDescribeVulsResponse() (response *DescribeVulsResponse) {
    response = &DescribeVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the vulnerability list.
func (c *Client) DescribeVuls(request *DescribeVulsRequest) (response *DescribeVulsResponse, err error) {
    if request == nil {
        request = NewDescribeVulsRequest()
    }
    response = NewDescribeVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportBruteAttacksRequest() (request *DescribeWeeklyReportBruteAttacksRequest) {
    request = &DescribeWeeklyReportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportBruteAttacks")
    return
}

func NewDescribeWeeklyReportBruteAttacksResponse() (response *DescribeWeeklyReportBruteAttacksResponse) {
    response = &DescribeWeeklyReportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the brute force attack data in the weekly CWP Pro report.
func (c *Client) DescribeWeeklyReportBruteAttacks(request *DescribeWeeklyReportBruteAttacksRequest) (response *DescribeWeeklyReportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportBruteAttacksRequest()
    }
    response = NewDescribeWeeklyReportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportInfoRequest() (request *DescribeWeeklyReportInfoRequest) {
    request = &DescribeWeeklyReportInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportInfo")
    return
}

func NewDescribeWeeklyReportInfoResponse() (response *DescribeWeeklyReportInfoResponse) {
    response = &DescribeWeeklyReportInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the details in the weekly CWP Pro report.
func (c *Client) DescribeWeeklyReportInfo(request *DescribeWeeklyReportInfoRequest) (response *DescribeWeeklyReportInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportInfoRequest()
    }
    response = NewDescribeWeeklyReportInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportMalwaresRequest() (request *DescribeWeeklyReportMalwaresRequest) {
    request = &DescribeWeeklyReportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportMalwares")
    return
}

func NewDescribeWeeklyReportMalwaresResponse() (response *DescribeWeeklyReportMalwaresResponse) {
    response = &DescribeWeeklyReportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the trojan data in the weekly CWP Pro report.
func (c *Client) DescribeWeeklyReportMalwares(request *DescribeWeeklyReportMalwaresRequest) (response *DescribeWeeklyReportMalwaresResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportMalwaresRequest()
    }
    response = NewDescribeWeeklyReportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesRequest() (request *DescribeWeeklyReportNonlocalLoginPlacesRequest) {
    request = &DescribeWeeklyReportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportNonlocalLoginPlaces")
    return
}

func NewDescribeWeeklyReportNonlocalLoginPlacesResponse() (response *DescribeWeeklyReportNonlocalLoginPlacesResponse) {
    response = &DescribeWeeklyReportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the unusual login location data in the weekly CWP Pro report.
func (c *Client) DescribeWeeklyReportNonlocalLoginPlaces(request *DescribeWeeklyReportNonlocalLoginPlacesRequest) (response *DescribeWeeklyReportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportNonlocalLoginPlacesRequest()
    }
    response = NewDescribeWeeklyReportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportVulsRequest() (request *DescribeWeeklyReportVulsRequest) {
    request = &DescribeWeeklyReportVulsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReportVuls")
    return
}

func NewDescribeWeeklyReportVulsResponse() (response *DescribeWeeklyReportVulsResponse) {
    response = &DescribeWeeklyReportVulsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the vulnerability data in the weekly CWP Pro report.
func (c *Client) DescribeWeeklyReportVuls(request *DescribeWeeklyReportVulsRequest) (response *DescribeWeeklyReportVulsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportVulsRequest()
    }
    response = NewDescribeWeeklyReportVulsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWeeklyReportsRequest() (request *DescribeWeeklyReportsRequest) {
    request = &DescribeWeeklyReportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "DescribeWeeklyReports")
    return
}

func NewDescribeWeeklyReportsResponse() (response *DescribeWeeklyReportsResponse) {
    response = &DescribeWeeklyReportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the weekly report list.
func (c *Client) DescribeWeeklyReports(request *DescribeWeeklyReportsRequest) (response *DescribeWeeklyReportsResponse, err error) {
    if request == nil {
        request = NewDescribeWeeklyReportsRequest()
    }
    response = NewDescribeWeeklyReportsResponse()
    err = c.Send(request, response)
    return
}

func NewEditTagsRequest() (request *EditTagsRequest) {
    request = &EditTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "EditTags")
    return
}

func NewEditTagsResponse() (response *EditTagsResponse) {
    response = &EditTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add or edit tags.
func (c *Client) EditTags(request *EditTagsRequest) (response *EditTagsResponse, err error) {
    if request == nil {
        request = NewEditTagsRequest()
    }
    response = NewEditTagsResponse()
    err = c.Send(request, response)
    return
}

func NewExportBruteAttacksRequest() (request *ExportBruteAttacksRequest) {
    request = &ExportBruteAttacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportBruteAttacks")
    return
}

func NewExportBruteAttacksResponse() (response *ExportBruteAttacksResponse) {
    response = &ExportBruteAttacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to export brute force attack records into a CSV file.
func (c *Client) ExportBruteAttacks(request *ExportBruteAttacksRequest) (response *ExportBruteAttacksResponse, err error) {
    if request == nil {
        request = NewExportBruteAttacksRequest()
    }
    response = NewExportBruteAttacksResponse()
    err = c.Send(request, response)
    return
}

func NewExportMaliciousRequestsRequest() (request *ExportMaliciousRequestsRequest) {
    request = &ExportMaliciousRequestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMaliciousRequests")
    return
}

func NewExportMaliciousRequestsResponse() (response *ExportMaliciousRequestsResponse) {
    response = &ExportMaliciousRequestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to export the malicious request file into a CSV file for download.
func (c *Client) ExportMaliciousRequests(request *ExportMaliciousRequestsRequest) (response *ExportMaliciousRequestsResponse, err error) {
    if request == nil {
        request = NewExportMaliciousRequestsRequest()
    }
    response = NewExportMaliciousRequestsResponse()
    err = c.Send(request, response)
    return
}

func NewExportMalwaresRequest() (request *ExportMalwaresRequest) {
    request = &ExportMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportMalwares")
    return
}

func NewExportMalwaresResponse() (response *ExportMalwaresResponse) {
    response = &ExportMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to export trojan records into a CSV file.
func (c *Client) ExportMalwares(request *ExportMalwaresRequest) (response *ExportMalwaresResponse, err error) {
    if request == nil {
        request = NewExportMalwaresRequest()
    }
    response = NewExportMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewExportNonlocalLoginPlacesRequest() (request *ExportNonlocalLoginPlacesRequest) {
    request = &ExportNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ExportNonlocalLoginPlaces")
    return
}

func NewExportNonlocalLoginPlacesResponse() (response *ExportNonlocalLoginPlacesResponse) {
    response = &ExportNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to export unusual login location events into a CSV file.
func (c *Client) ExportNonlocalLoginPlaces(request *ExportNonlocalLoginPlacesRequest) (response *ExportNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewExportNonlocalLoginPlacesRequest()
    }
    response = NewExportNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreImpactedHostsRequest() (request *IgnoreImpactedHostsRequest) {
    request = &IgnoreImpactedHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "IgnoreImpactedHosts")
    return
}

func NewIgnoreImpactedHostsResponse() (response *IgnoreImpactedHostsResponse) {
    response = &IgnoreImpactedHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to ignore one or more vulnerabilities.
func (c *Client) IgnoreImpactedHosts(request *IgnoreImpactedHostsRequest) (response *IgnoreImpactedHostsResponse, err error) {
    if request == nil {
        request = NewIgnoreImpactedHostsRequest()
    }
    response = NewIgnoreImpactedHostsResponse()
    err = c.Send(request, response)
    return
}

func NewMisAlarmNonlocalLoginPlacesRequest() (request *MisAlarmNonlocalLoginPlacesRequest) {
    request = &MisAlarmNonlocalLoginPlacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "MisAlarmNonlocalLoginPlaces")
    return
}

func NewMisAlarmNonlocalLoginPlacesResponse() (response *MisAlarmNonlocalLoginPlacesResponse) {
    response = &MisAlarmNonlocalLoginPlacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to set the current location as the usual login location.
func (c *Client) MisAlarmNonlocalLoginPlaces(request *MisAlarmNonlocalLoginPlacesRequest) (response *MisAlarmNonlocalLoginPlacesResponse, err error) {
    if request == nil {
        request = NewMisAlarmNonlocalLoginPlacesRequest()
    }
    response = NewMisAlarmNonlocalLoginPlacesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmAttributeRequest() (request *ModifyAlarmAttributeRequest) {
    request = &ModifyAlarmAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAlarmAttribute")
    return
}

func NewModifyAlarmAttributeResponse() (response *ModifyAlarmAttributeResponse) {
    response = &ModifyAlarmAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify alarm settings.
func (c *Client) ModifyAlarmAttribute(request *ModifyAlarmAttributeRequest) (response *ModifyAlarmAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmAttributeRequest()
    }
    response = NewModifyAlarmAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoOpenProVersionConfigRequest() (request *ModifyAutoOpenProVersionConfigRequest) {
    request = &ModifyAutoOpenProVersionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyAutoOpenProVersionConfig")
    return
}

func NewModifyAutoOpenProVersionConfigResponse() (response *ModifyAutoOpenProVersionConfigResponse) {
    response = &ModifyAutoOpenProVersionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to set whether to automatically activate CWP Pro for newly added servers.
func (c *Client) ModifyAutoOpenProVersionConfig(request *ModifyAutoOpenProVersionConfigRequest) (response *ModifyAutoOpenProVersionConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoOpenProVersionConfigRequest()
    }
    response = NewModifyAutoOpenProVersionConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoginWhiteListRequest() (request *ModifyLoginWhiteListRequest) {
    request = &ModifyLoginWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyLoginWhiteList")
    return
}

func NewModifyLoginWhiteListResponse() (response *ModifyLoginWhiteListResponse) {
    response = &ModifyLoginWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to edit a whitelist rule.
func (c *Client) ModifyLoginWhiteList(request *ModifyLoginWhiteListRequest) (response *ModifyLoginWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyLoginWhiteListRequest()
    }
    response = NewModifyLoginWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProVersionRenewFlagRequest() (request *ModifyProVersionRenewFlagRequest) {
    request = &ModifyProVersionRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "ModifyProVersionRenewFlag")
    return
}

func NewModifyProVersionRenewFlagResponse() (response *ModifyProVersionRenewFlagResponse) {
    response = &ModifyProVersionRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the renewal flag of CWP Pro.
func (c *Client) ModifyProVersionRenewFlag(request *ModifyProVersionRenewFlagRequest) (response *ModifyProVersionRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyProVersionRenewFlagRequest()
    }
    response = NewModifyProVersionRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewOpenProVersionRequest() (request *OpenProVersionRequest) {
    request = &OpenProVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "OpenProVersion")
    return
}

func NewOpenProVersionResponse() (response *OpenProVersionResponse) {
    response = &OpenProVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to activate CWP Pro.
func (c *Client) OpenProVersion(request *OpenProVersionRequest) (response *OpenProVersionResponse, err error) {
    if request == nil {
        request = NewOpenProVersionRequest()
    }
    response = NewOpenProVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMalwaresRequest() (request *RecoverMalwaresRequest) {
    request = &RecoverMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RecoverMalwares")
    return
}

func NewRecoverMalwaresResponse() (response *RecoverMalwaresResponse) {
    response = &RecoverMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to recover isolated trojan files in a batch.
func (c *Client) RecoverMalwares(request *RecoverMalwaresRequest) (response *RecoverMalwaresResponse, err error) {
    if request == nil {
        request = NewRecoverMalwaresRequest()
    }
    response = NewRecoverMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewRescanImpactedHostRequest() (request *RescanImpactedHostRequest) {
    request = &RescanImpactedHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "RescanImpactedHost")
    return
}

func NewRescanImpactedHostResponse() (response *RescanImpactedHostResponse) {
    response = &RescanImpactedHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to re-scan for vulnerabilities.
func (c *Client) RescanImpactedHost(request *RescanImpactedHostRequest) (response *RescanImpactedHostResponse, err error) {
    if request == nil {
        request = NewRescanImpactedHostRequest()
    }
    response = NewRescanImpactedHostResponse()
    err = c.Send(request, response)
    return
}

func NewSeparateMalwaresRequest() (request *SeparateMalwaresRequest) {
    request = &SeparateMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "SeparateMalwares")
    return
}

func NewSeparateMalwaresResponse() (response *SeparateMalwaresResponse) {
    response = &SeparateMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to isolate trojans.
func (c *Client) SeparateMalwares(request *SeparateMalwaresRequest) (response *SeparateMalwaresResponse, err error) {
    if request == nil {
        request = NewSeparateMalwaresRequest()
    }
    response = NewSeparateMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMaliciousRequestRequest() (request *TrustMaliciousRequestRequest) {
    request = &TrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMaliciousRequest")
    return
}

func NewTrustMaliciousRequestResponse() (response *TrustMaliciousRequestResponse) {
    response = &TrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to trust a malicious request.
func (c *Client) TrustMaliciousRequest(request *TrustMaliciousRequestRequest) (response *TrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewTrustMaliciousRequestRequest()
    }
    response = NewTrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewTrustMalwaresRequest() (request *TrustMalwaresRequest) {
    request = &TrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "TrustMalwares")
    return
}

func NewTrustMalwaresResponse() (response *TrustMalwaresResponse) {
    response = &TrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to trust an identified trojan file.
func (c *Client) TrustMalwares(request *TrustMalwaresRequest) (response *TrustMalwaresResponse, err error) {
    if request == nil {
        request = NewTrustMalwaresRequest()
    }
    response = NewTrustMalwaresResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMaliciousRequestRequest() (request *UntrustMaliciousRequestRequest) {
    request = &UntrustMaliciousRequestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMaliciousRequest")
    return
}

func NewUntrustMaliciousRequestResponse() (response *UntrustMaliciousRequestResponse) {
    response = &UntrustMaliciousRequestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to untrust a malicious request.
func (c *Client) UntrustMaliciousRequest(request *UntrustMaliciousRequestRequest) (response *UntrustMaliciousRequestResponse, err error) {
    if request == nil {
        request = NewUntrustMaliciousRequestRequest()
    }
    response = NewUntrustMaliciousRequestResponse()
    err = c.Send(request, response)
    return
}

func NewUntrustMalwaresRequest() (request *UntrustMalwaresRequest) {
    request = &UntrustMalwaresRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("yunjing", APIVersion, "UntrustMalwares")
    return
}

func NewUntrustMalwaresResponse() (response *UntrustMalwaresResponse) {
    response = &UntrustMalwaresResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to untrust a trojan file.
func (c *Client) UntrustMalwares(request *UntrustMalwaresRequest) (response *UntrustMalwaresResponse, err error) {
    if request == nil {
        request = NewUntrustMalwaresRequest()
    }
    response = NewUntrustMalwaresResponse()
    err = c.Send(request, response)
    return
}
