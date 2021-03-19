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

package v20180328

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-03-28"

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


func NewCloneDBRequest() (request *CloneDBRequest) {
    request = &CloneDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CloneDB")
    return
}

func NewCloneDBResponse() (response *CloneDBResponse) {
    response = &CloneDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to clone and rename databases of an instance. The clones are still in the instance from which they are cloned.
func (c *Client) CloneDB(request *CloneDBRequest) (response *CloneDBResponse, err error) {
    if request == nil {
        request = NewCloneDBRequest()
    }
    response = NewCloneDBResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateAccount")
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an instance account.
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBackup")
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a backup.
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupMigrationRequest() (request *CreateBackupMigrationRequest) {
    request = &CreateBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBackupMigration")
    return
}

func NewCreateBackupMigrationResponse() (response *CreateBackupMigrationResponse) {
    response = &CreateBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a backup import task.
func (c *Client) CreateBackupMigration(request *CreateBackupMigrationRequest) (response *CreateBackupMigrationResponse, err error) {
    if request == nil {
        request = NewCreateBackupMigrationRequest()
    }
    response = NewCreateBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBRequest() (request *CreateDBRequest) {
    request = &CreateDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDB")
    return
}

func NewCreateDBResponse() (response *CreateDBResponse) {
    response = &CreateDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a database.
func (c *Client) CreateDB(request *CreateDBRequest) (response *CreateDBResponse, err error) {
    if request == nil {
        request = NewCreateDBRequest()
    }
    response = NewCreateDBResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstancesRequest() (request *CreateDBInstancesRequest) {
    request = &CreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDBInstances")
    return
}

func NewCreateDBInstancesResponse() (response *CreateDBInstancesResponse) {
    response = &CreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an instance.
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    response = NewCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIncrementalMigrationRequest() (request *CreateIncrementalMigrationRequest) {
    request = &CreateIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateIncrementalMigration")
    return
}

func NewCreateIncrementalMigrationResponse() (response *CreateIncrementalMigrationResponse) {
    response = &CreateIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an incremental backup import task.
func (c *Client) CreateIncrementalMigration(request *CreateIncrementalMigrationRequest) (response *CreateIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewCreateIncrementalMigrationRequest()
    }
    response = NewCreateIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrationRequest() (request *CreateMigrationRequest) {
    request = &CreateMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateMigration")
    return
}

func NewCreateMigrationResponse() (response *CreateMigrationResponse) {
    response = &CreateMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a migration task.
func (c *Client) CreateMigration(request *CreateMigrationRequest) (response *CreateMigrationResponse, err error) {
    if request == nil {
        request = NewCreateMigrationRequest()
    }
    response = NewCreateMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteAccount")
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an instance account.
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupMigrationRequest() (request *DeleteBackupMigrationRequest) {
    request = &DeleteBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteBackupMigration")
    return
}

func NewDeleteBackupMigrationResponse() (response *DeleteBackupMigrationResponse) {
    response = &DeleteBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a backup import task.
func (c *Client) DeleteBackupMigration(request *DeleteBackupMigrationRequest) (response *DeleteBackupMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteBackupMigrationRequest()
    }
    response = NewDeleteBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBRequest() (request *DeleteDBRequest) {
    request = &DeleteDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDB")
    return
}

func NewDeleteDBResponse() (response *DeleteDBResponse) {
    response = &DeleteDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to drop a database.
func (c *Client) DeleteDB(request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
    if request == nil {
        request = NewDeleteDBRequest()
    }
    response = NewDeleteDBResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIncrementalMigrationRequest() (request *DeleteIncrementalMigrationRequest) {
    request = &DeleteIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteIncrementalMigration")
    return
}

func NewDeleteIncrementalMigrationResponse() (response *DeleteIncrementalMigrationResponse) {
    response = &DeleteIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an incremental backup import task.
func (c *Client) DeleteIncrementalMigration(request *DeleteIncrementalMigrationRequest) (response *DeleteIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteIncrementalMigrationRequest()
    }
    response = NewDeleteIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrationRequest() (request *DeleteMigrationRequest) {
    request = &DeleteMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteMigration")
    return
}

func NewDeleteMigrationResponse() (response *DeleteMigrationResponse) {
    response = &DeleteMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a migration task.
func (c *Client) DeleteMigration(request *DeleteMigrationRequest) (response *DeleteMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteMigrationRequest()
    }
    response = NewDeleteMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to pull the list of instance accounts.
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupCommandRequest() (request *DescribeBackupCommandRequest) {
    request = &DescribeBackupCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupCommand")
    return
}

func NewDescribeBackupCommandResponse() (response *DescribeBackupCommandResponse) {
    response = &DescribeBackupCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the commands of creating backups canonically.
func (c *Client) DescribeBackupCommand(request *DescribeBackupCommandRequest) (response *DescribeBackupCommandResponse, err error) {
    if request == nil {
        request = NewDescribeBackupCommandRequest()
    }
    response = NewDescribeBackupCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupMigrationRequest() (request *DescribeBackupMigrationRequest) {
    request = &DescribeBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupMigration")
    return
}

func NewDescribeBackupMigrationResponse() (response *DescribeBackupMigrationResponse) {
    response = &DescribeBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an incremental backup import task.
func (c *Client) DescribeBackupMigration(request *DescribeBackupMigrationRequest) (response *DescribeBackupMigrationResponse, err error) {
    if request == nil {
        request = NewDescribeBackupMigrationRequest()
    }
    response = NewDescribeBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupUploadSizeRequest() (request *DescribeBackupUploadSizeRequest) {
    request = &DescribeBackupUploadSizeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupUploadSize")
    return
}

func NewDescribeBackupUploadSizeResponse() (response *DescribeBackupUploadSizeResponse) {
    response = &DescribeBackupUploadSizeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the size of uploaded backup files. It is valid if the backup file type is `COS_UPLOAD` (the file is stored in COS).
func (c *Client) DescribeBackupUploadSize(request *DescribeBackupUploadSizeRequest) (response *DescribeBackupUploadSizeResponse, err error) {
    if request == nil {
        request = NewDescribeBackupUploadSizeRequest()
    }
    response = NewDescribeBackupUploadSizeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
    request = &DescribeBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackups")
    return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
    response = &DescribeBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of backups.
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    response = NewDescribeBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCharsetsRequest() (request *DescribeDBCharsetsRequest) {
    request = &DescribeDBCharsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBCharsets")
    return
}

func NewDescribeDBCharsetsResponse() (response *DescribeDBCharsetsResponse) {
    response = &DescribeDBCharsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the database character sets supported by an instance.
func (c *Client) DescribeDBCharsets(request *DescribeDBCharsetsRequest) (response *DescribeDBCharsetsResponse, err error) {
    if request == nil {
        request = NewDescribeDBCharsetsRequest()
    }
    response = NewDescribeDBCharsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of instances.
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBsRequest() (request *DescribeDBsRequest) {
    request = &DescribeDBsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBs")
    return
}

func NewDescribeDBsResponse() (response *DescribeDBsResponse) {
    response = &DescribeDBsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of databases
func (c *Client) DescribeDBs(request *DescribeDBsRequest) (response *DescribeDBsResponse, err error) {
    if request == nil {
        request = NewDescribeDBsRequest()
    }
    response = NewDescribeDBsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowStatusRequest() (request *DescribeFlowStatusRequest) {
    request = &DescribeFlowStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeFlowStatus")
    return
}

func NewDescribeFlowStatusResponse() (response *DescribeFlowStatusResponse) {
    response = &DescribeFlowStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query flow status.
func (c *Client) DescribeFlowStatus(request *DescribeFlowStatusRequest) (response *DescribeFlowStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFlowStatusRequest()
    }
    response = NewDescribeFlowStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIncrementalMigrationRequest() (request *DescribeIncrementalMigrationRequest) {
    request = &DescribeIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeIncrementalMigration")
    return
}

func NewDescribeIncrementalMigrationResponse() (response *DescribeIncrementalMigrationResponse) {
    response = &DescribeIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query an incremental backup import task.
func (c *Client) DescribeIncrementalMigration(request *DescribeIncrementalMigrationRequest) (response *DescribeIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewDescribeIncrementalMigrationRequest()
    }
    response = NewDescribeIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDetailRequest() (request *DescribeMigrationDetailRequest) {
    request = &DescribeMigrationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrationDetail")
    return
}

func NewDescribeMigrationDetailResponse() (response *DescribeMigrationDetailResponse) {
    response = &DescribeMigrationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query migration task details.
func (c *Client) DescribeMigrationDetail(request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDetailRequest()
    }
    response = NewDescribeMigrationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationsRequest() (request *DescribeMigrationsRequest) {
    request = &DescribeMigrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrations")
    return
}

func NewDescribeMigrationsResponse() (response *DescribeMigrationsResponse) {
    response = &DescribeMigrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of eligible migration tasks based on the entered criteria.
func (c *Client) DescribeMigrations(request *DescribeMigrationsRequest) (response *DescribeMigrationsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationsRequest()
    }
    response = NewDescribeMigrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeOrders")
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query order information.
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductConfigRequest() (request *DescribeProductConfigRequest) {
    request = &DescribeProductConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeProductConfig")
    return
}

func NewDescribeProductConfigResponse() (response *DescribeProductConfigResponse) {
    response = &DescribeProductConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query purchasable specification configuration.
func (c *Client) DescribeProductConfig(request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    if request == nil {
        request = NewDescribeProductConfigRequest()
    }
    response = NewDescribeProductConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query purchasable regions.
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeRequest() (request *DescribeRollbackTimeRequest) {
    request = &DescribeRollbackTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRollbackTime")
    return
}

func NewDescribeRollbackTimeResponse() (response *DescribeRollbackTimeResponse) {
    response = &DescribeRollbackTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the time range available for instance rollback.
func (c *Client) DescribeRollbackTime(request *DescribeRollbackTimeRequest) (response *DescribeRollbackTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRequest()
    }
    response = NewDescribeRollbackTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowlogsRequest() (request *DescribeSlowlogsRequest) {
    request = &DescribeSlowlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeSlowlogs")
    return
}

func NewDescribeSlowlogsResponse() (response *DescribeSlowlogsResponse) {
    response = &DescribeSlowlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get file information of slow query logs.
func (c *Client) DescribeSlowlogs(request *DescribeSlowlogsRequest) (response *DescribeSlowlogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowlogsRequest()
    }
    response = NewDescribeSlowlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadBackupInfoRequest() (request *DescribeUploadBackupInfoRequest) {
    request = &DescribeUploadBackupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeUploadBackupInfo")
    return
}

func NewDescribeUploadBackupInfoResponse() (response *DescribeUploadBackupInfoResponse) {
    response = &DescribeUploadBackupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query a backup upload permission.
func (c *Client) DescribeUploadBackupInfo(request *DescribeUploadBackupInfoRequest) (response *DescribeUploadBackupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadBackupInfoRequest()
    }
    response = NewDescribeUploadBackupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeZones")
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query currently purchasable AZs.
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateDBInstancesRequest() (request *InquiryPriceCreateDBInstancesRequest) {
    request = &InquiryPriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceCreateDBInstances")
    return
}

func NewInquiryPriceCreateDBInstancesResponse() (response *InquiryPriceCreateDBInstancesResponse) {
    response = &InquiryPriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the price of requested instances.
func (c *Client) InquiryPriceCreateDBInstances(request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDBInstancesRequest()
    }
    response = NewInquiryPriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeDBInstanceRequest() (request *InquiryPriceUpgradeDBInstanceRequest) {
    request = &InquiryPriceUpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceUpgradeDBInstance")
    return
}

func NewInquiryPriceUpgradeDBInstanceResponse() (response *InquiryPriceUpgradeDBInstanceResponse) {
    response = &InquiryPriceUpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the upgrade price of an instance.
func (c *Client) InquiryPriceUpgradeDBInstance(request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeDBInstanceRequest()
    }
    response = NewInquiryPriceUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegeRequest() (request *ModifyAccountPrivilegeRequest) {
    request = &ModifyAccountPrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyAccountPrivilege")
    return
}

func NewModifyAccountPrivilegeResponse() (response *ModifyAccountPrivilegeResponse) {
    response = &ModifyAccountPrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify instance account permissions.
func (c *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (response *ModifyAccountPrivilegeResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegeRequest()
    }
    response = NewModifyAccountPrivilegeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountRemarkRequest() (request *ModifyAccountRemarkRequest) {
    request = &ModifyAccountRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyAccountRemark")
    return
}

func NewModifyAccountRemarkResponse() (response *ModifyAccountRemarkResponse) {
    response = &ModifyAccountRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify account remarks.
func (c *Client) ModifyAccountRemark(request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    if request == nil {
        request = NewModifyAccountRemarkRequest()
    }
    response = NewModifyAccountRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupMigrationRequest() (request *ModifyBackupMigrationRequest) {
    request = &ModifyBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupMigration")
    return
}

func NewModifyBackupMigrationResponse() (response *ModifyBackupMigrationResponse) {
    response = &ModifyBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a backup import task.
func (c *Client) ModifyBackupMigration(request *ModifyBackupMigrationRequest) (response *ModifyBackupMigrationResponse, err error) {
    if request == nil {
        request = NewModifyBackupMigrationRequest()
    }
    response = NewModifyBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupStrategyRequest() (request *ModifyBackupStrategyRequest) {
    request = &ModifyBackupStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupStrategy")
    return
}

func NewModifyBackupStrategyResponse() (response *ModifyBackupStrategyResponse) {
    response = &ModifyBackupStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the backup policy.
func (c *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
    if request == nil {
        request = NewModifyBackupStrategyRequest()
    }
    response = NewModifyBackupStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename an instance.
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNetworkRequest() (request *ModifyDBInstanceNetworkRequest) {
    request = &ModifyDBInstanceNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceNetwork")
    return
}

func NewModifyDBInstanceNetworkResponse() (response *ModifyDBInstanceNetworkResponse) {
    response = &ModifyDBInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to switch a running instance from a VPC to another.
func (c *Client) ModifyDBInstanceNetwork(request *ModifyDBInstanceNetworkRequest) (response *ModifyDBInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNetworkRequest()
    }
    response = NewModifyDBInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceProject")
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the project to which a database instance belongs.
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBNameRequest() (request *ModifyDBNameRequest) {
    request = &ModifyDBNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBName")
    return
}

func NewModifyDBNameResponse() (response *ModifyDBNameResponse) {
    response = &ModifyDBNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename a database.
func (c *Client) ModifyDBName(request *ModifyDBNameRequest) (response *ModifyDBNameResponse, err error) {
    if request == nil {
        request = NewModifyDBNameRequest()
    }
    response = NewModifyDBNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBRemarkRequest() (request *ModifyDBRemarkRequest) {
    request = &ModifyDBRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBRemark")
    return
}

func NewModifyDBRemarkResponse() (response *ModifyDBRemarkResponse) {
    response = &ModifyDBRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify database remarks.
func (c *Client) ModifyDBRemark(request *ModifyDBRemarkRequest) (response *ModifyDBRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDBRemarkRequest()
    }
    response = NewModifyDBRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIncrementalMigrationRequest() (request *ModifyIncrementalMigrationRequest) {
    request = &ModifyIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyIncrementalMigration")
    return
}

func NewModifyIncrementalMigrationResponse() (response *ModifyIncrementalMigrationResponse) {
    response = &ModifyIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify an incremental backup import task.
func (c *Client) ModifyIncrementalMigration(request *ModifyIncrementalMigrationRequest) (response *ModifyIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewModifyIncrementalMigrationRequest()
    }
    response = NewModifyIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationRequest() (request *ModifyMigrationRequest) {
    request = &ModifyMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyMigration")
    return
}

func NewModifyMigrationResponse() (response *ModifyMigrationResponse) {
    response = &ModifyMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify an existing migration task.
func (c *Client) ModifyMigration(request *ModifyMigrationRequest) (response *ModifyMigrationResponse, err error) {
    if request == nil {
        request = NewModifyMigrationRequest()
    }
    response = NewModifyMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleDBInstanceRequest() (request *RecycleDBInstanceRequest) {
    request = &RecycleDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RecycleDBInstance")
    return
}

func NewRecycleDBInstanceResponse() (response *RecycleDBInstanceResponse) {
    response = &RecycleDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to return a deactivated SQL Server instance.
func (c *Client) RecycleDBInstance(request *RecycleDBInstanceRequest) (response *RecycleDBInstanceResponse, err error) {
    if request == nil {
        request = NewRecycleDBInstanceRequest()
    }
    response = NewRecycleDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ResetAccountPassword")
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reset the account password of an instance.
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
    request = &RestartDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RestartDBInstance")
    return
}

func NewRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
    response = &RestartDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to restart a database instance.
func (c *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestartDBInstanceRequest()
    }
    response = NewRestartDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreInstanceRequest() (request *RestoreInstanceRequest) {
    request = &RestoreInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RestoreInstance")
    return
}

func NewRestoreInstanceResponse() (response *RestoreInstanceResponse) {
    response = &RestoreInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to restore an instance from a backup file.
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    if request == nil {
        request = NewRestoreInstanceRequest()
    }
    response = NewRestoreInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackInstanceRequest() (request *RollbackInstanceRequest) {
    request = &RollbackInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RollbackInstance")
    return
}

func NewRollbackInstanceResponse() (response *RollbackInstanceResponse) {
    response = &RollbackInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to roll back an instance.
func (c *Client) RollbackInstance(request *RollbackInstanceRequest) (response *RollbackInstanceResponse, err error) {
    if request == nil {
        request = NewRollbackInstanceRequest()
    }
    response = NewRollbackInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunMigrationRequest() (request *RunMigrationRequest) {
    request = &RunMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RunMigration")
    return
}

func NewRunMigrationResponse() (response *RunMigrationResponse) {
    response = &RunMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to start running a migration task.
func (c *Client) RunMigration(request *RunMigrationRequest) (response *RunMigrationResponse, err error) {
    if request == nil {
        request = NewRunMigrationRequest()
    }
    response = NewRunMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartBackupMigrationRequest() (request *StartBackupMigrationRequest) {
    request = &StartBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartBackupMigration")
    return
}

func NewStartBackupMigrationResponse() (response *StartBackupMigrationResponse) {
    response = &StartBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to start a backup import task.
func (c *Client) StartBackupMigration(request *StartBackupMigrationRequest) (response *StartBackupMigrationResponse, err error) {
    if request == nil {
        request = NewStartBackupMigrationRequest()
    }
    response = NewStartBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartIncrementalMigrationRequest() (request *StartIncrementalMigrationRequest) {
    request = &StartIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartIncrementalMigration")
    return
}

func NewStartIncrementalMigrationResponse() (response *StartIncrementalMigrationResponse) {
    response = &StartIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to start an incremental backup import task.
func (c *Client) StartIncrementalMigration(request *StartIncrementalMigrationRequest) (response *StartIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewStartIncrementalMigrationRequest()
    }
    response = NewStartIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDBInstanceRequest() (request *TerminateDBInstanceRequest) {
    request = &TerminateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "TerminateDBInstance")
    return
}

func NewTerminateDBInstanceResponse() (response *TerminateDBInstanceResponse) {
    response = &TerminateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to isolate an instance to move it into a recycle bin.
func (c *Client) TerminateDBInstance(request *TerminateDBInstanceRequest) (response *TerminateDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDBInstanceRequest()
    }
    response = NewTerminateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "UpgradeDBInstance")
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to upgrade an instance.
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
