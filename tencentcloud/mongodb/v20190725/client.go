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

package v20190725

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-07-25"

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


func NewAssignProjectRequest() (request *AssignProjectRequest) {
    request = &AssignProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "AssignProject")
    return
}

func NewAssignProjectResponse() (response *AssignProjectResponse) {
    response = &AssignProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to specify the project to which a TencentDB instance belongs.
func (c *Client) AssignProject(request *AssignProjectRequest) (response *AssignProjectResponse, err error) {
    if request == nil {
        request = NewAssignProjectRequest()
    }
    response = NewAssignProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstance")
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a monthly subscription TencentDB for MongoDB instance. The purchasable specifications supported by this API can be obtained through the `DescribeSpecInfo` API.
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    response = NewCreateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceHourRequest() (request *CreateDBInstanceHourRequest) {
    request = &CreateDBInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstanceHour")
    return
}

func NewCreateDBInstanceHourResponse() (response *CreateDBInstanceHourResponse) {
    response = &CreateDBInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a pay-as-you-go TencentDB for MongoDB instance.
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupAccessRequest() (request *DescribeBackupAccessRequest) {
    request = &DescribeBackupAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeBackupAccess")
    return
}

func NewDescribeBackupAccessResponse() (response *DescribeBackupAccessResponse) {
    response = &DescribeBackupAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the permission to download a backup file. The specific backup file information can be obtained through the DescribeDBBackups API.
func (c *Client) DescribeBackupAccess(request *DescribeBackupAccessRequest) (response *DescribeBackupAccessResponse, err error) {
    if request == nil {
        request = NewDescribeBackupAccessRequest()
    }
    response = NewDescribeBackupAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientConnectionsRequest() (request *DescribeClientConnectionsRequest) {
    request = &DescribeClientConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeClientConnections")
    return
}

func NewDescribeClientConnectionsResponse() (response *DescribeClientConnectionsResponse) {
    response = &DescribeClientConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the client connection information of an instance, including the IP and number of connections. Currently, only instances of MongoDB 3.2 are supported.
func (c *Client) DescribeClientConnections(request *DescribeClientConnectionsRequest) (response *DescribeClientConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeClientConnectionsRequest()
    }
    response = NewDescribeClientConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
    request = &DescribeDBBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBBackups")
    return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
    response = &DescribeDBBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of instance backups. Currently, only backups in the last 7 days can be queried.
func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBBackupsRequest()
    }
    response = NewDescribeDBBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceDealRequest() (request *DescribeDBInstanceDealRequest) {
    request = &DescribeDBInstanceDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceDeal")
    return
}

func NewDescribeDBInstanceDealResponse() (response *DescribeDBInstanceDealResponse) {
    response = &DescribeDBInstanceDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get details of purchase, renewal, and specification adjustment orders of a MongoDB instance.
func (c *Client) DescribeDBInstanceDeal(request *DescribeDBInstanceDealRequest) (response *DescribeDBInstanceDealResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceDealRequest()
    }
    response = NewDescribeDBInstanceDealResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of TencentDB instances (which can be main, disaster recovery, or read-only instances). It supports filtering instances by project ID, instance ID, and instance status.
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogPatternsRequest() (request *DescribeSlowLogPatternsRequest) {
    request = &DescribeSlowLogPatternsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogPatterns")
    return
}

func NewDescribeSlowLogPatternsResponse() (response *DescribeSlowLogPatternsResponse) {
    response = &DescribeSlowLogPatternsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the slow log statistics of a database instance.
func (c *Client) DescribeSlowLogPatterns(request *DescribeSlowLogPatternsRequest) (response *DescribeSlowLogPatternsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogPatternsRequest()
    }
    response = NewDescribeSlowLogPatternsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogs")
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the slow log information of a TencentDB instance. Only slow logs for the last 7 days can be queried.
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecInfoRequest() (request *DescribeSpecInfoRequest) {
    request = &DescribeSpecInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSpecInfo")
    return
}

func NewDescribeSpecInfoResponse() (response *DescribeSpecInfoResponse) {
    response = &DescribeSpecInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the purchasable instance specifications.
func (c *Client) DescribeSpecInfo(request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSpecInfoRequest()
    }
    response = NewDescribeSpecInfoResponse()
    err = c.Send(request, response)
    return
}

func NewFlushInstanceRouterConfigRequest() (request *FlushInstanceRouterConfigRequest) {
    request = &FlushInstanceRouterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "FlushInstanceRouterConfig")
    return
}

func NewFlushInstanceRouterConfigResponse() (response *FlushInstanceRouterConfigResponse) {
    response = &FlushInstanceRouterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to run the `FlushRouterConfig` command on all mongos instances.
func (c *Client) FlushInstanceRouterConfig(request *FlushInstanceRouterConfigRequest) (response *FlushInstanceRouterConfigResponse, err error) {
    if request == nil {
        request = NewFlushInstanceRouterConfigRequest()
    }
    response = NewFlushInstanceRouterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateDBInstancesRequest() (request *InquirePriceCreateDBInstancesRequest) {
    request = &InquirePriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceCreateDBInstances")
    return
}

func NewInquirePriceCreateDBInstancesResponse() (response *InquirePriceCreateDBInstancesResponse) {
    response = &InquirePriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query price of instance creation. The `region` parameter must be passed in this API, otherwise verification will fail. This API allows queries only for purchasable instance specifications.
func (c *Client) InquirePriceCreateDBInstances(request *InquirePriceCreateDBInstancesRequest) (response *InquirePriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDBInstancesRequest()
    }
    response = NewInquirePriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyDBInstanceSpecRequest() (request *InquirePriceModifyDBInstanceSpecRequest) {
    request = &InquirePriceModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceModifyDBInstanceSpec")
    return
}

func NewInquirePriceModifyDBInstanceSpecResponse() (response *InquirePriceModifyDBInstanceSpecResponse) {
    response = &InquirePriceModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query price of instance specification adjustment.
func (c *Client) InquirePriceModifyDBInstanceSpec(request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDBInstanceSpecRequest()
    }
    response = NewInquirePriceModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewDBInstancesRequest() (request *InquirePriceRenewDBInstancesRequest) {
    request = &InquirePriceRenewDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceRenewDBInstances")
    return
}

func NewInquirePriceRenewDBInstancesResponse() (response *InquirePriceRenewDBInstancesResponse) {
    response = &InquirePriceRenewDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the renewal price of a monthly subscription instance.
func (c *Client) InquirePriceRenewDBInstances(request *InquirePriceRenewDBInstancesRequest) (response *InquirePriceRenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewDBInstancesRequest()
    }
    response = NewInquirePriceRenewDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "IsolateDBInstance")
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to isolate a pay-as-you-go TencentDB for MongoDB instance. An isolated instance is retained in the recycle bin and data can no longer be written to it. After it is isolated for a certain period of time, it will be completely deleted. For the retention period in the recycle bin, please see the terms of service for pay-as-you-go billing. Isolated pay-as-you-go instances cannot be recovered, so please proceed with caution.
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
    request = &ModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceSpec")
    return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
    response = &ModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to adjust the specification configuration of a TencentDB for MongoDB. The purchasable specifications supported by the API can be obtained through the DescribeSpecInfo API.
func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSpecRequest()
    }
    response = NewModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedDBInstanceRequest() (request *OfflineIsolatedDBInstanceRequest) {
    request = &OfflineIsolatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "OfflineIsolatedDBInstance")
    return
}

func NewOfflineIsolatedDBInstanceResponse() (response *OfflineIsolatedDBInstanceResponse) {
    response = &OfflineIsolatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to deactivate isolated TencentDB instances immediately. The instances must be in isolated status.
func (c *Client) OfflineIsolatedDBInstance(request *OfflineIsolatedDBInstanceRequest) (response *OfflineIsolatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedDBInstanceRequest()
    }
    response = NewOfflineIsolatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
    request = &RenameInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "RenameInstance")
    return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
    response = &RenameInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename a TencentDB instance.
func (c *Client) RenameInstance(request *RenameInstanceRequest) (response *RenameInstanceResponse, err error) {
    if request == nil {
        request = NewRenameInstanceRequest()
    }
    response = NewRenameInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBInstancesRequest() (request *RenewDBInstancesRequest) {
    request = &RenewDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mongodb", APIVersion, "RenewDBInstances")
    return
}

func NewRenewDBInstancesResponse() (response *RenewDBInstancesResponse) {
    response = &RenewDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to renew a monthly subscription TencentDB instance. Only monthly subscription instances are supported, while pay-as-you-go instances do not need to be renewed.
func (c *Client) RenewDBInstances(request *RenewDBInstancesRequest) (response *RenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewRenewDBInstancesRequest()
    }
    response = NewRenewDBInstancesResponse()
    err = c.Send(request, response)
    return
}
