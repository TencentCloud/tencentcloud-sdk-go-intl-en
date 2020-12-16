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

package v20170320

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2017-03-20"

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


func NewAddTimeWindowRequest() (request *AddTimeWindowRequest) {
    request = &AddTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "AddTimeWindow")
    return
}

func NewAddTimeWindowResponse() (response *AddTimeWindowResponse) {
    response = &AddTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (AddTimeWindow) is used to add a maintenance time window for a TencentDB instance, so as to specify when the instance can automatically perform access switch operations.
func (c *Client) AddTimeWindow(request *AddTimeWindowRequest) (response *AddTimeWindowResponse, err error) {
    if request == nil {
        request = NewAddTimeWindowRequest()
    }
    response = NewAddTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "AssociateSecurityGroups")
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (AssociateSecurityGroups) is used to bind security groups to instances in batches.
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewBalanceRoGroupLoadRequest() (request *BalanceRoGroupLoadRequest) {
    request = &BalanceRoGroupLoadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "BalanceRoGroupLoad")
    return
}

func NewBalanceRoGroupLoadResponse() (response *BalanceRoGroupLoadResponse) {
    response = &BalanceRoGroupLoadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rebalance the loads of instances in an RO group. Please note that the database connections to those instances will be interrupted transiently; therefore, you should ensure that your application can reconnect to the databases. This operation should be performed with caution.
func (c *Client) BalanceRoGroupLoad(request *BalanceRoGroupLoadRequest) (response *BalanceRoGroupLoadResponse, err error) {
    if request == nil {
        request = NewBalanceRoGroupLoadRequest()
    }
    response = NewBalanceRoGroupLoadResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWanServiceRequest() (request *CloseWanServiceRequest) {
    request = &CloseWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CloseWanService")
    return
}

func NewCloseWanServiceResponse() (response *CloseWanServiceResponse) {
    response = &CloseWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CloseWanService) is used to disable public network access for TencentDB instance, which will make public IP addresses inaccessible.
func (c *Client) CloseWanService(request *CloseWanServiceRequest) (response *CloseWanServiceResponse, err error) {
    if request == nil {
        request = NewCloseWanServiceRequest()
    }
    response = NewCloseWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountsRequest() (request *CreateAccountsRequest) {
    request = &CreateAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAccounts")
    return
}

func NewCreateAccountsResponse() (response *CreateAccountsResponse) {
    response = &CreateAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateAccounts) is used to create TencentDB accounts. The new account names, domain names, and passwords need to be specified, and account remarks can also be added.
func (c *Client) CreateAccounts(request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    if request == nil {
        request = NewCreateAccountsRequest()
    }
    response = NewCreateAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateBackup")
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateBackup) is used to create a TencentDB instance backup.
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloneInstanceRequest() (request *CreateCloneInstanceRequest) {
    request = &CreateCloneInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateCloneInstance")
    return
}

func NewCreateCloneInstanceResponse() (response *CreateCloneInstanceResponse) {
    response = &CreateCloneInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a clone of a specific instance, and roll back the clone by using a physical backup file of the instance or roll back the clone to a point in time.
func (c *Client) CreateCloneInstance(request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    if request == nil {
        request = NewCreateCloneInstanceRequest()
    }
    response = NewCreateCloneInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBImportJobRequest() (request *CreateDBImportJobRequest) {
    request = &CreateDBImportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBImportJob")
    return
}

func NewCreateDBImportJobResponse() (response *CreateDBImportJobResponse) {
    response = &CreateDBImportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateDBImportJob) is used to create a data import task for a TencentDB instance.
// 
// Note that the files for a data import task must be uploaded to Tencent Cloud in advance. You need to do so in the console.
func (c *Client) CreateDBImportJob(request *CreateDBImportJobRequest) (response *CreateDBImportJobResponse, err error) {
    if request == nil {
        request = NewCreateDBImportJobRequest()
    }
    response = NewCreateDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceHourRequest() (request *CreateDBInstanceHourRequest) {
    request = &CreateDBInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBInstanceHour")
    return
}

func NewCreateDBInstanceHourResponse() (response *CreateDBInstanceHourResponse) {
    response = &CreateDBInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a pay-as-you-go TencentDB instance (which can be a primary, disaster recovery, or read-only instance) by passing in information such as instance specifications, MySQL version number, and quantity.
// 
// This is an async API. You can also use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the instance details. If the `Status` value of an instance is 1 and `TaskStatus` is 0, the instance has been successfully delivered.
// 
// 1. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported instance specifications first and then use the [DescribeDBPrice](https://intl.cloud.tencent.com/document/api/236/18566?from_cn_redirect=1) API to query the prices of the supported instances;
// 2. You can create up to 100 instances at a time, with an instance validity period of up to 36 months;
// 3. MySQL v5.5, v5.6, and v5.7 are supported;
// 4. primary instances, read-only instances, and disaster recovery instances can be created;
// 5. If `Port`, `ParamList`, or `Password` is set in the input parameters, the instance will be initialized.
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeployGroupRequest() (request *CreateDeployGroupRequest) {
    request = &CreateDeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDeployGroup")
    return
}

func NewCreateDeployGroupResponse() (response *CreateDeployGroupResponse) {
    response = &CreateDeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a placement group for placing instances.
func (c *Client) CreateDeployGroup(request *CreateDeployGroupRequest) (response *CreateDeployGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeployGroupRequest()
    }
    response = NewCreateDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateParamTemplateRequest() (request *CreateParamTemplateRequest) {
    request = &CreateParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateParamTemplate")
    return
}

func NewCreateParamTemplateResponse() (response *CreateParamTemplateResponse) {
    response = &CreateParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateParamTemplate) is used to create a parameter template.
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    if request == nil {
        request = NewCreateParamTemplateRequest()
    }
    response = NewCreateParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoInstanceIpRequest() (request *CreateRoInstanceIpRequest) {
    request = &CreateRoInstanceIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateRoInstanceIp")
    return
}

func NewCreateRoInstanceIpResponse() (response *CreateRoInstanceIpResponse) {
    response = &CreateRoInstanceIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a VIP exclusive to a TencentDB read-only instance.
func (c *Client) CreateRoInstanceIp(request *CreateRoInstanceIpRequest) (response *CreateRoInstanceIpResponse, err error) {
    if request == nil {
        request = NewCreateRoInstanceIpRequest()
    }
    response = NewCreateRoInstanceIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountsRequest() (request *DeleteAccountsRequest) {
    request = &DeleteAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteAccounts")
    return
}

func NewDeleteAccountsResponse() (response *DeleteAccountsResponse) {
    response = &DeleteAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DeleteAccounts) is used to delete TencentDB accounts.
func (c *Client) DeleteAccounts(request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteAccountsRequest()
    }
    response = NewDeleteAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
    request = &DeleteBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteBackup")
    return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
    response = &DeleteBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a database backup. It can only delete manually initiated backups.
func (c *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    if request == nil {
        request = NewDeleteBackupRequest()
    }
    response = NewDeleteBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeployGroupsRequest() (request *DeleteDeployGroupsRequest) {
    request = &DeleteDeployGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteDeployGroups")
    return
}

func NewDeleteDeployGroupsResponse() (response *DeleteDeployGroupsResponse) {
    response = &DeleteDeployGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete placement groups by placement group ID (a placement group cannot be deleted if it contains resources).
func (c *Client) DeleteDeployGroups(request *DeleteDeployGroupsRequest) (response *DeleteDeployGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDeployGroupsRequest()
    }
    response = NewDeleteDeployGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParamTemplateRequest() (request *DeleteParamTemplateRequest) {
    request = &DeleteParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteParamTemplate")
    return
}

func NewDeleteParamTemplateResponse() (response *DeleteParamTemplateResponse) {
    response = &DeleteParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DeleteParamTemplate) is used to delete a parameter template.
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteParamTemplateRequest()
    }
    response = NewDeleteParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTimeWindowRequest() (request *DeleteTimeWindowRequest) {
    request = &DeleteTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteTimeWindow")
    return
}

func NewDeleteTimeWindowResponse() (response *DeleteTimeWindowResponse) {
    response = &DeleteTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DeleteTimeWindow) is used to delete a maintenance time window for a TencentDB instance. After it is deleted, the default maintenance time window will be 03:00-04:00, i.e., switch to a new instance will be performed during 03:00-04:00 by default.
func (c *Client) DeleteTimeWindow(request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    if request == nil {
        request = NewDeleteTimeWindowRequest()
    }
    response = NewDeleteTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAccountPrivileges")
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeAccountPrivileges) is used to query the information of TencentDB account permissions.
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeAccounts) is used to query information of all TencentDB accounts.
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAsyncRequestInfo")
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeAsyncRequestInfo) is used to query the async task execution result of a TencentDB instance.
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupConfigRequest() (request *DescribeBackupConfigRequest) {
    request = &DescribeBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupConfig")
    return
}

func NewDescribeBackupConfigResponse() (response *DescribeBackupConfigResponse) {
    response = &DescribeBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeBackupConfig) is used to query the configuration information of a TencentDB instance backup.
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDatabasesRequest() (request *DescribeBackupDatabasesRequest) {
    request = &DescribeBackupDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDatabases")
    return
}

func NewDescribeBackupDatabasesResponse() (response *DescribeBackupDatabasesResponse) {
    response = &DescribeBackupDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the databases contained in a backup file. It has been disused.
// After the legacy version becomes capable of full backup, if you want to download logical backup files by table, you need to use this API.
// The new API (CreateBackup) can specify the table to be backed up when a logical backup file is created, which can be downloaded directly.
func (c *Client) DescribeBackupDatabases(request *DescribeBackupDatabasesRequest) (response *DescribeBackupDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDatabasesRequest()
    }
    response = NewDescribeBackupDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupOverviewRequest() (request *DescribeBackupOverviewRequest) {
    request = &DescribeBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupOverview")
    return
}

func NewDescribeBackupOverviewResponse() (response *DescribeBackupOverviewResponse) {
    response = &DescribeBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the backup overview of a user. It will return the user's current total number of backups, total capacity used by backups, capacity in the free tier, and paid capacity (all capacity values are in bytes).
func (c *Client) DescribeBackupOverview(request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBackupOverviewRequest()
    }
    response = NewDescribeBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupSummariesRequest() (request *DescribeBackupSummariesRequest) {
    request = &DescribeBackupSummariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupSummaries")
    return
}

func NewDescribeBackupSummariesResponse() (response *DescribeBackupSummariesResponse) {
    response = &DescribeBackupSummariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the statistics of backups. It will return the capacity used by backups at the instance level and the number and used capacity of data backups and log backups of each instance (all capacity values are in bytes).
func (c *Client) DescribeBackupSummaries(request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupSummariesRequest()
    }
    response = NewDescribeBackupSummariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupTablesRequest() (request *DescribeBackupTablesRequest) {
    request = &DescribeBackupTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupTables")
    return
}

func NewDescribeBackupTablesResponse() (response *DescribeBackupTablesResponse) {
    response = &DescribeBackupTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the backup tables of the specified database. It has been disused.
// After the legacy version becomes capable of full backup, if you want to download logical backup files by table, you need to use this API.
// The new API (CreateBackup) can specify the table to be backed up when a logical backup file is created, which can be downloaded directly.
func (c *Client) DescribeBackupTables(request *DescribeBackupTablesRequest) (response *DescribeBackupTablesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupTablesRequest()
    }
    response = NewDescribeBackupTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
    request = &DescribeBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackups")
    return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
    response = &DescribeBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeBackups) is used to query the backups of a TencentDB instance.
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    response = NewDescribeBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogBackupOverviewRequest() (request *DescribeBinlogBackupOverviewRequest) {
    request = &DescribeBinlogBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBinlogBackupOverview")
    return
}

func NewDescribeBinlogBackupOverviewResponse() (response *DescribeBinlogBackupOverviewResponse) {
    response = &DescribeBinlogBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the log backup overview of a user in the current region.
func (c *Client) DescribeBinlogBackupOverview(request *DescribeBinlogBackupOverviewRequest) (response *DescribeBinlogBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogBackupOverviewRequest()
    }
    response = NewDescribeBinlogBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogsRequest() (request *DescribeBinlogsRequest) {
    request = &DescribeBinlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBinlogs")
    return
}

func NewDescribeBinlogsResponse() (response *DescribeBinlogsResponse) {
    response = &DescribeBinlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of binlog files of a TencentDB instance.
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogsRequest()
    }
    response = NewDescribeBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloneListRequest() (request *DescribeCloneListRequest) {
    request = &DescribeCloneListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCloneList")
    return
}

func NewDescribeCloneListResponse() (response *DescribeCloneListResponse) {
    response = &DescribeCloneListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the clone task list of an instance.
func (c *Client) DescribeCloneList(request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    if request == nil {
        request = NewDescribeCloneListRequest()
    }
    response = NewDescribeCloneListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBImportRecordsRequest() (request *DescribeDBImportRecordsRequest) {
    request = &DescribeDBImportRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBImportRecords")
    return
}

func NewDescribeDBImportRecordsResponse() (response *DescribeDBImportRecordsResponse) {
    response = &DescribeDBImportRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBImportRecords) is used to query the records of import tasks in a TencentDB instance.
func (c *Client) DescribeDBImportRecords(request *DescribeDBImportRecordsRequest) (response *DescribeDBImportRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBImportRecordsRequest()
    }
    response = NewDescribeDBImportRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceCharsetRequest() (request *DescribeDBInstanceCharsetRequest) {
    request = &DescribeDBInstanceCharsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceCharset")
    return
}

func NewDescribeDBInstanceCharsetResponse() (response *DescribeDBInstanceCharsetResponse) {
    response = &DescribeDBInstanceCharsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBInstanceCharset) is used to query the character set and its name of a TencentDB instance.
func (c *Client) DescribeDBInstanceCharset(request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceCharsetRequest()
    }
    response = NewDescribeDBInstanceCharsetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceConfigRequest() (request *DescribeDBInstanceConfigRequest) {
    request = &DescribeDBInstanceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceConfig")
    return
}

func NewDescribeDBInstanceConfigResponse() (response *DescribeDBInstanceConfigResponse) {
    response = &DescribeDBInstanceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBInstanceConfig) is used to query the configuration information of a TencentDB instance, such as its synchronization mode and deployment mode.
func (c *Client) DescribeDBInstanceConfig(request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceConfigRequest()
    }
    response = NewDescribeDBInstanceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceGTIDRequest() (request *DescribeDBInstanceGTIDRequest) {
    request = &DescribeDBInstanceGTIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceGTID")
    return
}

func NewDescribeDBInstanceGTIDResponse() (response *DescribeDBInstanceGTIDResponse) {
    response = &DescribeDBInstanceGTIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBInstanceGTID) is used to query whether GTID is activated for a TencentDB instance. Instances on or below version 5.5 are not supported.
func (c *Client) DescribeDBInstanceGTID(request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceGTIDRequest()
    }
    response = NewDescribeDBInstanceGTIDResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceInfoRequest() (request *DescribeDBInstanceInfoRequest) {
    request = &DescribeDBInstanceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceInfo")
    return
}

func NewDescribeDBInstanceInfoResponse() (response *DescribeDBInstanceInfoResponse) {
    response = &DescribeDBInstanceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the basic information of an instance (instance ID, instance name, and whether encryption is enabled).
func (c *Client) DescribeDBInstanceInfo(request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceInfoRequest()
    }
    response = NewDescribeDBInstanceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceRebootTimeRequest() (request *DescribeDBInstanceRebootTimeRequest) {
    request = &DescribeDBInstanceRebootTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceRebootTime")
    return
}

func NewDescribeDBInstanceRebootTimeResponse() (response *DescribeDBInstanceRebootTimeResponse) {
    response = &DescribeDBInstanceRebootTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBInstanceRebootTime) is used to query the estimated time needed for a TencentDB instance to restart.
func (c *Client) DescribeDBInstanceRebootTime(request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceRebootTimeRequest()
    }
    response = NewDescribeDBInstanceRebootTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBInstances) is used to query the list of TencentDB instances (which can be primary, disaster recovery, or read-only instances). It supports filtering instances by project ID, instance ID, access address, and instance status.
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBSecurityGroups")
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBSecurityGroups) is used to query the security group details of an instance.
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSwitchRecordsRequest() (request *DescribeDBSwitchRecordsRequest) {
    request = &DescribeDBSwitchRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBSwitchRecords")
    return
}

func NewDescribeDBSwitchRecordsResponse() (response *DescribeDBSwitchRecordsResponse) {
    response = &DescribeDBSwitchRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBSwitchRecords) is used to query the instance switch records.
func (c *Client) DescribeDBSwitchRecords(request *DescribeDBSwitchRecordsRequest) (response *DescribeDBSwitchRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSwitchRecordsRequest()
    }
    response = NewDescribeDBSwitchRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBZoneConfigRequest() (request *DescribeDBZoneConfigRequest) {
    request = &DescribeDBZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBZoneConfig")
    return
}

func NewDescribeDBZoneConfigResponse() (response *DescribeDBZoneConfigResponse) {
    response = &DescribeDBZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDBZoneConfig) is used to query the specifications of TencentDB instances purchasable in a region.
func (c *Client) DescribeDBZoneConfig(request *DescribeDBZoneConfigRequest) (response *DescribeDBZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBZoneConfigRequest()
    }
    response = NewDescribeDBZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataBackupOverviewRequest() (request *DescribeDataBackupOverviewRequest) {
    request = &DescribeDataBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDataBackupOverview")
    return
}

func NewDescribeDataBackupOverviewResponse() (response *DescribeDataBackupOverviewResponse) {
    response = &DescribeDataBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the data backup overview of a user in the current region.
func (c *Client) DescribeDataBackupOverview(request *DescribeDataBackupOverviewRequest) (response *DescribeDataBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDataBackupOverviewRequest()
    }
    response = NewDescribeDataBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDatabases")
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDatabases) is used to query the information of databases of a TencentDB instance.
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultParamsRequest() (request *DescribeDefaultParamsRequest) {
    request = &DescribeDefaultParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDefaultParams")
    return
}

func NewDescribeDefaultParamsResponse() (response *DescribeDefaultParamsResponse) {
    response = &DescribeDefaultParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDefaultParams) is used to query the list of default configurable parameters.
func (c *Client) DescribeDefaultParams(request *DescribeDefaultParamsRequest) (response *DescribeDefaultParamsResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultParamsRequest()
    }
    response = NewDescribeDefaultParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployGroupListRequest() (request *DescribeDeployGroupListRequest) {
    request = &DescribeDeployGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDeployGroupList")
    return
}

func NewDescribeDeployGroupListResponse() (response *DescribeDeployGroupListResponse) {
    response = &DescribeDeployGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of placement groups of a user. You can specify the placement group ID or name.
func (c *Client) DescribeDeployGroupList(request *DescribeDeployGroupListRequest) (response *DescribeDeployGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeDeployGroupListRequest()
    }
    response = NewDescribeDeployGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceMonitorInfoRequest() (request *DescribeDeviceMonitorInfoRequest) {
    request = &DescribeDeviceMonitorInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDeviceMonitorInfo")
    return
}

func NewDescribeDeviceMonitorInfoResponse() (response *DescribeDeviceMonitorInfoResponse) {
    response = &DescribeDeviceMonitorInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeDeviceMonitorInfo) is used to query the monitoring information of a TencentDB physical machine on the day. Currently, it only supports instances with 488 GB memory and 6 TB disk.
func (c *Client) DescribeDeviceMonitorInfo(request *DescribeDeviceMonitorInfoRequest) (response *DescribeDeviceMonitorInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceMonitorInfoRequest()
    }
    response = NewDescribeDeviceMonitorInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorLogDataRequest() (request *DescribeErrorLogDataRequest) {
    request = &DescribeErrorLogDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeErrorLogData")
    return
}

func NewDescribeErrorLogDataResponse() (response *DescribeErrorLogDataResponse) {
    response = &DescribeErrorLogDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of instance error logs by search criteria. You can only query error logs within a month.
func (c *Client) DescribeErrorLogData(request *DescribeErrorLogDataRequest) (response *DescribeErrorLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeErrorLogDataRequest()
    }
    response = NewDescribeErrorLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceParamRecords")
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeInstanceParamRecords) is used to query the parameter modification records of an instance.
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamRecordsRequest()
    }
    response = NewDescribeInstanceParamRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceParams")
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeInstanceParams) is used to query the list of parameters for an instance.
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplateInfoRequest() (request *DescribeParamTemplateInfoRequest) {
    request = &DescribeParamTemplateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeParamTemplateInfo")
    return
}

func NewDescribeParamTemplateInfoResponse() (response *DescribeParamTemplateInfoResponse) {
    response = &DescribeParamTemplateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeParamTemplateInfo) is used to query parameter template details.
func (c *Client) DescribeParamTemplateInfo(request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplateInfoRequest()
    }
    response = NewDescribeParamTemplateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplatesRequest() (request *DescribeParamTemplatesRequest) {
    request = &DescribeParamTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeParamTemplates")
    return
}

func NewDescribeParamTemplatesResponse() (response *DescribeParamTemplatesResponse) {
    response = &DescribeParamTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeParamTemplates) is used to query the list of parameter templates
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplatesRequest()
    }
    response = NewDescribeParamTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProjectSecurityGroups")
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeProjectSecurityGroups) is used to query the security group details of a project.
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoGroupsRequest() (request *DescribeRoGroupsRequest) {
    request = &DescribeRoGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRoGroups")
    return
}

func NewDescribeRoGroupsResponse() (response *DescribeRoGroupsResponse) {
    response = &DescribeRoGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of all RO groups of a TencentDB instance.
func (c *Client) DescribeRoGroups(request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRoGroupsRequest()
    }
    response = NewDescribeRoGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoMinScaleRequest() (request *DescribeRoMinScaleRequest) {
    request = &DescribeRoMinScaleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRoMinScale")
    return
}

func NewDescribeRoMinScaleResponse() (response *DescribeRoMinScaleResponse) {
    response = &DescribeRoMinScaleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the minimum specification of a read-only instance that can be purchased or upgraded to.
func (c *Client) DescribeRoMinScale(request *DescribeRoMinScaleRequest) (response *DescribeRoMinScaleResponse, err error) {
    if request == nil {
        request = NewDescribeRoMinScaleRequest()
    }
    response = NewDescribeRoMinScaleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackRangeTimeRequest() (request *DescribeRollbackRangeTimeRequest) {
    request = &DescribeRollbackRangeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRollbackRangeTime")
    return
}

func NewDescribeRollbackRangeTimeResponse() (response *DescribeRollbackRangeTimeResponse) {
    response = &DescribeRollbackRangeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeRollbackRangeTime) is used to query the time range available for instance rollback.
func (c *Client) DescribeRollbackRangeTime(request *DescribeRollbackRangeTimeRequest) (response *DescribeRollbackRangeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackRangeTimeRequest()
    }
    response = NewDescribeRollbackRangeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTaskDetailRequest() (request *DescribeRollbackTaskDetailRequest) {
    request = &DescribeRollbackTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRollbackTaskDetail")
    return
}

func NewDescribeRollbackTaskDetailResponse() (response *DescribeRollbackTaskDetailResponse) {
    response = &DescribeRollbackTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of a TencentDB instance rollback task.
func (c *Client) DescribeRollbackTaskDetail(request *DescribeRollbackTaskDetailRequest) (response *DescribeRollbackTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTaskDetailRequest()
    }
    response = NewDescribeRollbackTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogDataRequest() (request *DescribeSlowLogDataRequest) {
    request = &DescribeSlowLogDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSlowLogData")
    return
}

func NewDescribeSlowLogDataResponse() (response *DescribeSlowLogDataResponse) {
    response = &DescribeSlowLogDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to search for slow logs of an instance by criteria. You can only view slow logs within a month.
func (c *Client) DescribeSlowLogData(request *DescribeSlowLogDataRequest) (response *DescribeSlowLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogDataRequest()
    }
    response = NewDescribeSlowLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSlowLogs")
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeSlowLogs) is used to query the slow logs of a TencentDB instance.
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportedPrivilegesRequest() (request *DescribeSupportedPrivilegesRequest) {
    request = &DescribeSupportedPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSupportedPrivileges")
    return
}

func NewDescribeSupportedPrivilegesResponse() (response *DescribeSupportedPrivilegesResponse) {
    response = &DescribeSupportedPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeSupportedPrivileges) is used to query the information of TencentDB account permissions, including global permissions, database permissions, table permissions, and column permissions.
func (c *Client) DescribeSupportedPrivileges(request *DescribeSupportedPrivilegesRequest) (response *DescribeSupportedPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeSupportedPrivilegesRequest()
    }
    response = NewDescribeSupportedPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTables")
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeTables) is used to query the database tables of a TencentDB instance.
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsOfInstanceIdsRequest() (request *DescribeTagsOfInstanceIdsRequest) {
    request = &DescribeTagsOfInstanceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTagsOfInstanceIds")
    return
}

func NewDescribeTagsOfInstanceIdsResponse() (response *DescribeTagsOfInstanceIdsResponse) {
    response = &DescribeTagsOfInstanceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeTagsOfInstanceIds) is used to query the tag information of a TencentDB instance.
func (c *Client) DescribeTagsOfInstanceIds(request *DescribeTagsOfInstanceIdsRequest) (response *DescribeTagsOfInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsOfInstanceIdsRequest()
    }
    response = NewDescribeTagsOfInstanceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeTasks) is used to query the list of tasks for a TencentDB instance.
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimeWindowRequest() (request *DescribeTimeWindowRequest) {
    request = &DescribeTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTimeWindow")
    return
}

func NewDescribeTimeWindowResponse() (response *DescribeTimeWindowResponse) {
    response = &DescribeTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DescribeTimeWindow) is used to query the maintenance time window of a TencentDB instance.
func (c *Client) DescribeTimeWindow(request *DescribeTimeWindowRequest) (response *DescribeTimeWindowResponse, err error) {
    if request == nil {
        request = NewDescribeTimeWindowRequest()
    }
    response = NewDescribeTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadedFilesRequest() (request *DescribeUploadedFilesRequest) {
    request = &DescribeUploadedFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeUploadedFiles")
    return
}

func NewDescribeUploadedFilesResponse() (response *DescribeUploadedFilesResponse) {
    response = &DescribeUploadedFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of user-imported SQL files.
func (c *Client) DescribeUploadedFiles(request *DescribeUploadedFilesRequest) (response *DescribeUploadedFilesResponse, err error) {
    if request == nil {
        request = NewDescribeUploadedFilesRequest()
    }
    response = NewDescribeUploadedFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DisassociateSecurityGroups")
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DisassociateSecurityGroups) is used to unbind security groups from instances in batches.
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "InitDBInstances")
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (InitDBInstances) is used to initialize instances, including their password, default character set, and instance port number.
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "IsolateDBInstance")
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to isolate a TencentDB instance, which will no longer be accessible via IP and port. The isolated instance can be started up in the recycle bin. If it is isolated due to arrears, please top up your account as soon as possible.
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountDescription")
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyAccountDescription) is used to modify the remarks of a TencentDB instance account.
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPasswordRequest() (request *ModifyAccountPasswordRequest) {
    request = &ModifyAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountPassword")
    return
}

func NewModifyAccountPasswordResponse() (response *ModifyAccountPasswordResponse) {
    response = &ModifyAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyAccountPassword) is used to modify the password of a TencentDB instance account.
func (c *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (response *ModifyAccountPasswordResponse, err error) {
    if request == nil {
        request = NewModifyAccountPasswordRequest()
    }
    response = NewModifyAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountPrivileges")
    return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
    response = &ModifyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the permissions of a TencentDB instance account.
// 
// Note that when modifying account permissions, you need to pass in the full permission information of the account. You can [query the account permission information
// ](https://intl.cloud.tencent.com/document/api/236/17500?from_cn_redirect=1) first before modifying permissions.
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    response = NewModifyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAutoRenewFlag")
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the auto-renewal flag of a TencentDB instance.
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupConfigRequest() (request *ModifyBackupConfigRequest) {
    request = &ModifyBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyBackupConfig")
    return
}

func NewModifyBackupConfigResponse() (response *ModifyBackupConfigResponse) {
    response = &ModifyBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyBackupConfig) is used to modify the database backup configuration.
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    response = NewModifyBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyDBInstanceName) is used to rename a TencentDB instance.
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceProject")
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyDBInstanceProject) is used to modify the project to which a TencentDB instance belongs.
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyDBInstanceSecurityGroups) is used to modify the security groups bound to a TencentDB instance.
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceVipVportRequest() (request *ModifyDBInstanceVipVportRequest) {
    request = &ModifyDBInstanceVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceVipVport")
    return
}

func NewModifyDBInstanceVipVportResponse() (response *ModifyDBInstanceVipVportResponse) {
    response = &ModifyDBInstanceVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyDBInstanceVipVport) is used to modify the IP and port number of a TencentDB instance, switch from the basic network to VPC, or change VPC subnets.
func (c *Client) ModifyDBInstanceVipVport(request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceVipVportRequest()
    }
    response = NewModifyDBInstanceVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstanceParam")
    return
}

func NewModifyInstanceParamResponse() (response *ModifyInstanceParamResponse) {
    response = &ModifyInstanceParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyInstanceParam) is used to modify instance parameters.
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    response = NewModifyInstanceParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceTagRequest() (request *ModifyInstanceTagRequest) {
    request = &ModifyInstanceTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstanceTag")
    return
}

func NewModifyInstanceTagResponse() (response *ModifyInstanceTagResponse) {
    response = &ModifyInstanceTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyInstanceTag) is used to add, modify, or delete an instance tag.
func (c *Client) ModifyInstanceTag(request *ModifyInstanceTagRequest) (response *ModifyInstanceTagResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTagRequest()
    }
    response = NewModifyInstanceTagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNameOrDescByDpIdRequest() (request *ModifyNameOrDescByDpIdRequest) {
    request = &ModifyNameOrDescByDpIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyNameOrDescByDpId")
    return
}

func NewModifyNameOrDescByDpIdResponse() (response *ModifyNameOrDescByDpIdResponse) {
    response = &ModifyNameOrDescByDpIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the name or description of a placement group.
func (c *Client) ModifyNameOrDescByDpId(request *ModifyNameOrDescByDpIdRequest) (response *ModifyNameOrDescByDpIdResponse, err error) {
    if request == nil {
        request = NewModifyNameOrDescByDpIdRequest()
    }
    response = NewModifyNameOrDescByDpIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyParamTemplateRequest() (request *ModifyParamTemplateRequest) {
    request = &ModifyParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyParamTemplate")
    return
}

func NewModifyParamTemplateResponse() (response *ModifyParamTemplateResponse) {
    response = &ModifyParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyParamTemplate) is used to modify a parameter template.
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    if request == nil {
        request = NewModifyParamTemplateRequest()
    }
    response = NewModifyParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoGroupInfoRequest() (request *ModifyRoGroupInfoRequest) {
    request = &ModifyRoGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyRoGroupInfo")
    return
}

func NewModifyRoGroupInfoResponse() (response *ModifyRoGroupInfoResponse) {
    response = &ModifyRoGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the information of a TencentDB RO group, such as configuring an instance removal policy in case of excessive delay and setting read weights of RO instances.
func (c *Client) ModifyRoGroupInfo(request *ModifyRoGroupInfoRequest) (response *ModifyRoGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyRoGroupInfoRequest()
    }
    response = NewModifyRoGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTimeWindowRequest() (request *ModifyTimeWindowRequest) {
    request = &ModifyTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyTimeWindow")
    return
}

func NewModifyTimeWindowResponse() (response *ModifyTimeWindowResponse) {
    response = &ModifyTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyTimeWindow) is used to update the maintenance time window of a TencentDB instance.
func (c *Client) ModifyTimeWindow(request *ModifyTimeWindowRequest) (response *ModifyTimeWindowResponse, err error) {
    if request == nil {
        request = NewModifyTimeWindowRequest()
    }
    response = NewModifyTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedInstancesRequest() (request *OfflineIsolatedInstancesRequest) {
    request = &OfflineIsolatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "OfflineIsolatedInstances")
    return
}

func NewOfflineIsolatedInstancesResponse() (response *OfflineIsolatedInstancesResponse) {
    response = &OfflineIsolatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (OfflineIsolatedInstances) is used to deactivate isolated TencentDB instances immediately. The instances must be in isolated status, i.e., their `Status` value is 5 in the return of the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1).
// 
// This is an asynchronous API. There may be a delay in repossessing some resources. You can query the details by using the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) and specifying the InstanceId and the `Status` value as [5, 6, 7]. If the returned instance is empty, then all its resources have been released.
// 
// Note that once an instance is deactivated, its resources and data will not be recoverable. Please do so with caution.
func (c *Client) OfflineIsolatedInstances(request *OfflineIsolatedInstancesRequest) (response *OfflineIsolatedInstancesResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedInstancesRequest()
    }
    response = NewOfflineIsolatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBInstanceGTIDRequest() (request *OpenDBInstanceGTIDRequest) {
    request = &OpenDBInstanceGTIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "OpenDBInstanceGTID")
    return
}

func NewOpenDBInstanceGTIDResponse() (response *OpenDBInstanceGTIDResponse) {
    response = &OpenDBInstanceGTIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (OpenDBInstanceGTID) is used to enable GTID for a TencentDB instance. Only instances on or above version 5.6 are supported.
func (c *Client) OpenDBInstanceGTID(request *OpenDBInstanceGTIDRequest) (response *OpenDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewOpenDBInstanceGTIDRequest()
    }
    response = NewOpenDBInstanceGTIDResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWanServiceRequest() (request *OpenWanServiceRequest) {
    request = &OpenWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "OpenWanService")
    return
}

func NewOpenWanServiceResponse() (response *OpenWanServiceResponse) {
    response = &OpenWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (OpenWanService) is used to enable public network access for an instance.
// 
// Note that before enabling public network access, you need to first [initialize the instance](https://intl.cloud.tencent.com/document/api/236/15873?from_cn_redirect=1).
func (c *Client) OpenWanService(request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    if request == nil {
        request = NewOpenWanServiceRequest()
    }
    response = NewOpenWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseIsolatedDBInstancesRequest() (request *ReleaseIsolatedDBInstancesRequest) {
    request = &ReleaseIsolatedDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ReleaseIsolatedDBInstances")
    return
}

func NewReleaseIsolatedDBInstancesResponse() (response *ReleaseIsolatedDBInstancesResponse) {
    response = &ReleaseIsolatedDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to deisolate an isolated TencentDB instance.
func (c *Client) ReleaseIsolatedDBInstances(request *ReleaseIsolatedDBInstancesRequest) (response *ReleaseIsolatedDBInstancesResponse, err error) {
    if request == nil {
        request = NewReleaseIsolatedDBInstancesRequest()
    }
    response = NewReleaseIsolatedDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstancesRequest() (request *RestartDBInstancesRequest) {
    request = &RestartDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "RestartDBInstances")
    return
}

func NewRestartDBInstancesResponse() (response *RestartDBInstancesResponse) {
    response = &RestartDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (RestartDBInstances) is used to restart TencentDB instances.
// 
// Note:
// 1. This API only supports restarting primary instances.
// 2. The instance status must be normal, and no other async tasks are in progress.
func (c *Client) RestartDBInstances(request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    if request == nil {
        request = NewRestartDBInstancesRequest()
    }
    response = NewRestartDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartBatchRollbackRequest() (request *StartBatchRollbackRequest) {
    request = &StartBatchRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "StartBatchRollback")
    return
}

func NewStartBatchRollbackResponse() (response *StartBatchRollbackResponse) {
    response = &StartBatchRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (StartBatchRollback) is used to roll back the tables of a TencentDB instance in batches.
func (c *Client) StartBatchRollback(request *StartBatchRollbackRequest) (response *StartBatchRollbackResponse, err error) {
    if request == nil {
        request = NewStartBatchRollbackRequest()
    }
    response = NewStartBatchRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewStopDBImportJobRequest() (request *StopDBImportJobRequest) {
    request = &StopDBImportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "StopDBImportJob")
    return
}

func NewStopDBImportJobResponse() (response *StopDBImportJobResponse) {
    response = &StopDBImportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (StopDBImportJob) is used to stop a data import task.
func (c *Client) StopDBImportJob(request *StopDBImportJobRequest) (response *StopDBImportJobResponse, err error) {
    if request == nil {
        request = NewStopDBImportJobRequest()
    }
    response = NewStopDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopRollbackRequest() (request *StopRollbackRequest) {
    request = &StopRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "StopRollback")
    return
}

func NewStopRollbackResponse() (response *StopRollbackResponse) {
    response = &StopRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to cancel a rollback task in progress, and returns an async task ID. You can use the `DescribeAsyncRequestInfo` API to query the result of cancellation.
func (c *Client) StopRollback(request *StopRollbackRequest) (response *StopRollbackResponse, err error) {
    if request == nil {
        request = NewStopRollbackRequest()
    }
    response = NewStopRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchForUpgradeRequest() (request *SwitchForUpgradeRequest) {
    request = &SwitchForUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchForUpgrade")
    return
}

func NewSwitchForUpgradeResponse() (response *SwitchForUpgradeResponse) {
    response = &SwitchForUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (SwitchForUpgrade) is used to switch to a new instance. You can initiate this process when the primary instance being upgraded is pending switch.
func (c *Client) SwitchForUpgrade(request *SwitchForUpgradeRequest) (response *SwitchForUpgradeResponse, err error) {
    if request == nil {
        request = NewSwitchForUpgradeRequest()
    }
    response = NewSwitchForUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeDBInstance")
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to upgrade or downgrade a TencentDB instance, which can be a primary instance, disaster recovery instance, or read-only instance.
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceEngineVersionRequest() (request *UpgradeDBInstanceEngineVersionRequest) {
    request = &UpgradeDBInstanceEngineVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeDBInstanceEngineVersion")
    return
}

func NewUpgradeDBInstanceEngineVersionResponse() (response *UpgradeDBInstanceEngineVersionResponse) {
    response = &UpgradeDBInstanceEngineVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (UpgradeDBInstanceEngineVersion) is used to upgrade the version of a TencentDB instance, which can be a primary instance, disaster recovery instance, or read-only instance.
func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceEngineVersionRequest()
    }
    response = NewUpgradeDBInstanceEngineVersionResponse()
    err = c.Send(request, response)
    return
}
