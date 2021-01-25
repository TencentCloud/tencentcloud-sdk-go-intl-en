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

package v20170312

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "AssociateSecurityGroups")
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to associate security groups with Tencent Cloud resources in batches.
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCloneAccountRequest() (request *CloneAccountRequest) {
    request = &CloneAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CloneAccount")
    return
}

func NewCloneAccountResponse() (response *CloneAccountResponse) {
    response = &CloneAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to clone an instance account.
func (c *Client) CloneAccount(request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    if request == nil {
        request = NewCloneAccountRequest()
    }
    response = NewCloneAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseDBExtranetAccessRequest() (request *CloseDBExtranetAccessRequest) {
    request = &CloseDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CloseDBExtranetAccess")
    return
}

func NewCloseDBExtranetAccessResponse() (response *CloseDBExtranetAccessResponse) {
    response = &CloseDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable public network access for a TencentDB instance, which will make the public IP address inaccessible. The `DescribeDCDBInstances` API will not return the public domain name and port information of the corresponding instance.
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAccountPrivilegesRequest() (request *CopyAccountPrivilegesRequest) {
    request = &CopyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CopyAccountPrivileges")
    return
}

func NewCopyAccountPrivilegesResponse() (response *CopyAccountPrivilegesResponse) {
    response = &CopyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to copy the permissions of a TencentDB account.
// Note: accounts with the same username but different hosts are different accounts. Permissions can only be copied between accounts with the same `Readonly` attribute.
func (c *Client) CopyAccountPrivileges(request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewCopyAccountPrivilegesRequest()
    }
    response = NewCopyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateAccount")
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a TencentDB account. Multiple accounts can be created for one instance. Accounts with the same username but different hosts are different accounts.
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DeleteAccount")
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a TencentDB account, which is uniquely identified by username and host.
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAccountPrivileges")
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the permissions of a TencentDB account.
// Note: accounts with the same username but different hosts are different accounts.
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
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of accounts of a specified TencentDB instance.
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupTimeRequest() (request *DescribeBackupTimeRequest) {
    request = &DescribeBackupTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeBackupTime")
    return
}

func NewDescribeBackupTimeResponse() (response *DescribeBackupTimeResponse) {
    response = &DescribeBackupTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the backup time of a TencentDB instance. The backend system will perform instance backup regularly according to this configuration.
func (c *Client) DescribeBackupTime(request *DescribeBackupTimeRequest) (response *DescribeBackupTimeResponse, err error) {
    if request == nil {
        request = NewDescribeBackupTimeRequest()
    }
    response = NewDescribeBackupTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the TencentDB instance list. It supports filtering instances by project ID, instance ID, private address, and instance name.
// If no filter is specified, 20 instances will be returned by default. Up to 100 instances can be returned for a single request.
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBLogFiles")
    return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
    response = &DescribeDBLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of various logs of a database, including cold backups, binlogs, errlogs, and slowlogs.
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    response = NewDescribeDBLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
    request = &DescribeDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBParameters")
    return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
    response = &DescribeDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the current parameter settings of a database.
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBParametersRequest()
    }
    response = NewDescribeDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBPerformanceRequest() (request *DescribeDBPerformanceRequest) {
    request = &DescribeDBPerformanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBPerformance")
    return
}

func NewDescribeDBPerformanceResponse() (response *DescribeDBPerformanceResponse) {
    response = &DescribeDBPerformanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the current performance data of a database instance.
func (c *Client) DescribeDBPerformance(request *DescribeDBPerformanceRequest) (response *DescribeDBPerformanceResponse, err error) {
    if request == nil {
        request = NewDescribeDBPerformanceRequest()
    }
    response = NewDescribeDBPerformanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBPerformanceDetailsRequest() (request *DescribeDBPerformanceDetailsRequest) {
    request = &DescribeDBPerformanceDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBPerformanceDetails")
    return
}

func NewDescribeDBPerformanceDetailsResponse() (response *DescribeDBPerformanceDetailsResponse) {
    response = &DescribeDBPerformanceDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the instance performance data details.
func (c *Client) DescribeDBPerformanceDetails(request *DescribeDBPerformanceDetailsRequest) (response *DescribeDBPerformanceDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeDBPerformanceDetailsRequest()
    }
    response = NewDescribeDBPerformanceDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBResourceUsageRequest() (request *DescribeDBResourceUsageRequest) {
    request = &DescribeDBResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBResourceUsage")
    return
}

func NewDescribeDBResourceUsageResponse() (response *DescribeDBResourceUsageResponse) {
    response = &DescribeDBResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the resource usage of a database instance.
func (c *Client) DescribeDBResourceUsage(request *DescribeDBResourceUsageRequest) (response *DescribeDBResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeDBResourceUsageRequest()
    }
    response = NewDescribeDBResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBResourceUsageDetailsRequest() (request *DescribeDBResourceUsageDetailsRequest) {
    request = &DescribeDBResourceUsageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBResourceUsageDetails")
    return
}

func NewDescribeDBResourceUsageDetailsResponse() (response *DescribeDBResourceUsageDetailsResponse) {
    response = &DescribeDBResourceUsageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the current performance data of a database instance.
func (c *Client) DescribeDBResourceUsageDetails(request *DescribeDBResourceUsageDetailsRequest) (response *DescribeDBResourceUsageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeDBResourceUsageDetailsRequest()
    }
    response = NewDescribeDBResourceUsageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSecurityGroups")
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the security group details of an instance.
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSlowLogsRequest() (request *DescribeDBSlowLogsRequest) {
    request = &DescribeDBSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSlowLogs")
    return
}

func NewDescribeDBSlowLogsResponse() (response *DescribeDBSlowLogsResponse) {
    response = &DescribeDBSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the slow query log list.
func (c *Client) DescribeDBSlowLogs(request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowLogsRequest()
    }
    response = NewDescribeDBSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabases")
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of databases of a TencentDB instance.
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeFlow")
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query flow status.
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogFileRetentionPeriodRequest() (request *DescribeLogFileRetentionPeriodRequest) {
    request = &DescribeLogFileRetentionPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeLogFileRetentionPeriod")
    return
}

func NewDescribeLogFileRetentionPeriodResponse() (response *DescribeLogFileRetentionPeriodResponse) {
    response = &DescribeLogFileRetentionPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the configured number of days for retention of database backup logs.
func (c *Client) DescribeLogFileRetentionPeriod(request *DescribeLogFileRetentionPeriodRequest) (response *DescribeLogFileRetentionPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeLogFileRetentionPeriodRequest()
    }
    response = NewDescribeLogFileRetentionPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeProjectSecurityGroups")
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the security group details of a project.
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "DisassociateSecurityGroups")
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unassociate security groups from instances in batches.
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
    request = &GrantAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "GrantAccountPrivileges")
    return
}

func NewGrantAccountPrivilegesResponse() (response *GrantAccountPrivilegesResponse) {
    response = &GrantAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to grant permissions to a TencentDB account.
// Note: accounts with the same username but different hosts are different accounts.
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewGrantAccountPrivilegesRequest()
    }
    response = NewGrantAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "InitDBInstances")
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to initialize TencentDB instances, including setting the default character set and table name case sensitivity.
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyAccountDescription")
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the remarks of a TencentDB account.
// Note: accounts with the same username but different hosts are different accounts.
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupTimeRequest() (request *ModifyBackupTimeRequest) {
    request = &ModifyBackupTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyBackupTime")
    return
}

func NewModifyBackupTimeResponse() (response *ModifyBackupTimeResponse) {
    response = &ModifyBackupTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to set the backup time of a TencentDB instance. The backend system will perform instance backup regularly according to this configuration.
func (c *Client) ModifyBackupTime(request *ModifyBackupTimeRequest) (response *ModifyBackupTimeResponse, err error) {
    if request == nil {
        request = NewModifyBackupTimeRequest()
    }
    response = NewModifyBackupTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename a TencentDB instance.
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBInstanceSecurityGroups")
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the security groups associated with TencentDB.
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstancesProjectRequest() (request *ModifyDBInstancesProjectRequest) {
    request = &ModifyDBInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBInstancesProject")
    return
}

func NewModifyDBInstancesProjectResponse() (response *ModifyDBInstancesProjectResponse) {
    response = &ModifyDBInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the project to which TencentDB instances belong.
func (c *Client) ModifyDBInstancesProject(request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
    request = &ModifyDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBParameters")
    return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
    response = &ModifyDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify database parameters.
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogFileRetentionPeriodRequest() (request *ModifyLogFileRetentionPeriodRequest) {
    request = &ModifyLogFileRetentionPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyLogFileRetentionPeriod")
    return
}

func NewModifyLogFileRetentionPeriodResponse() (response *ModifyLogFileRetentionPeriodResponse) {
    response = &ModifyLogFileRetentionPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the number of days for retention of database backup logs.
func (c *Client) ModifyLogFileRetentionPeriod(request *ModifyLogFileRetentionPeriodRequest) (response *ModifyLogFileRetentionPeriodResponse, err error) {
    if request == nil {
        request = NewModifyLogFileRetentionPeriodRequest()
    }
    response = NewModifyLogFileRetentionPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBExtranetAccessRequest() (request *OpenDBExtranetAccessRequest) {
    request = &OpenDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "OpenDBExtranetAccess")
    return
}

func NewOpenDBExtranetAccessResponse() (response *OpenDBExtranetAccessResponse) {
    response = &OpenDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable public network access for a TencentDB instance. After that, you can access the instance with the public domain name and port obtained through the `DescribeDCDBInstances` API.
func (c *Client) OpenDBExtranetAccess(request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mariadb", APIVersion, "ResetAccountPassword")
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reset the password of a TencentDB account.
// Note: accounts with the same username but different hosts are different accounts.
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}
