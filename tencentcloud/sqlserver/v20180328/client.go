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

// CloneDB
// This API is used to clone and rename databases of an instance. The clones are still in the instance from which they are cloned.
//
// error code that may be returned:
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// CreateAccount
// This API is used to create an instance account.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMEISILLEGAL = "InvalidParameterValue.AccountNameIsIllegal"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMEISKEYWORDS = "InvalidParameterValue.AccountNameIsKeyWords"
//  INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// CreateBackup
// This API is used to create a backup.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// CreateBackupMigration
// This API is used to create a backup import task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// CreateDB
// This API is used to create a database.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_CHARSETISILLEGAL = "InvalidParameterValue.CharsetIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PRIVILEGEISILLEGAL = "InvalidParameterValue.PrivilegeIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// CreateDBInstances
// This API is used to create an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// CreateIncrementalMigration
// This API is used to create an incremental backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCEINUSE_INCREMENTALMIGRATIONEXIST = "ResourceInUse.IncrementalMigrationExist"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BACKUPMIGRATIONRECOVERYTYPEERR = "ResourceUnavailable.BackupMigrationRecoveryTypeErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// CreateMigration
// This API is used to create a migration task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DeleteAccount
// This API is used to delete an instance account.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DeleteBackupMigration
// This API is used to delete a backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
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

// DeleteDB
// This API is used to drop a database.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DeleteIncrementalMigration
// This API is used to delete an incremental backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DeleteMigration
// This API is used to delete a migration task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeAccounts
// This API is used to pull the list of instance accounts.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeBackupCommand
// This API is used to query the commands of creating backups canonically.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeBackupMigration
// This API is used to create an incremental backup import task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeBackupUploadSize
// This API is used to query the size of uploaded backup files. It is valid if the backup file type is `COS_UPLOAD` (the file is stored in COS).
//
// error code that may be returned:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeBackups
// This API is used to query the list of backups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeDBCharsets
// This API is used to query the database character sets supported by an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// DescribeDBInstances
// This API is used to query the list of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeDBs
// This API is used to query the list of databases
//
// error code that may be returned:
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeFlowStatus
// This API is used to query flow status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeIncrementalMigration
// This API is used to query an incremental backup import task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeMigrationDetail
// This API is used to query migration task details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeMigrations
// This API is used to query the list of eligible migration tasks based on the entered criteria.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
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

// DescribeOrders
// This API is used to query order information.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYORDERFAILED = "FailedOperation.QueryOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeProductConfig
// This API is used to query purchasable specification configuration.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeRegions
// This API is used to query purchasable regions.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeRollbackTime
// This API is used to query the time range available for instance rollback.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeSlowlogs
// This API is used to get file information of slow query logs.
//
// error code that may be returned:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// DescribeUploadBackupInfo
// This API is used to query a backup upload permission.
//
// error code that may be returned:
//  FAILEDOPERATION_COSPROPERTIESERROR = "FailedOperation.CosPropertiesError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_STSERROR = "InternalError.StsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_UPLOADTYPEERROR = "UnsupportedOperation.UploadTypeError"
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

// DescribeZones
// This API is used to query currently purchasable AZs.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// InquiryPriceCreateDBInstances
// This API is used to query the price of requested instances.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BADGOODSNUM = "InvalidParameterValue.BadGoodsNum"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// InquiryPriceUpgradeDBInstance
// This API is used to query the upgrade price of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSTTYPENOTSUPPORTED = "InvalidParameterValue.CostTypeNotSupported"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCEEXPANDVOLUMELOW = "InvalidParameterValue.InstanceExpandVolumeLow"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyAccountPrivilege
// This API is used to modify instance account permissions.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyAccountRemark
// This API is used to modify account remarks.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyBackupMigration
// This API is used to modify a backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyBackupStrategy
// This API is used to modify the backup policy.
//
// error code that may be returned:
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyDBInstanceName
// This API is used to rename an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_INSTANCENAMEISILLEGAL = "InvalidParameterValue.InstanceNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyDBInstanceNetwork
// This API is used to switch a running instance from a VPC to another.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
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

// ModifyDBInstanceProject
// This API is used to modify the project to which a database instance belongs.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyDBName
// This API is used to rename a database.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyDBRemark
// This API is used to modify database remarks.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyIncrementalMigration
// This API is used to modify an incremental backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ModifyMigration
// This API is used to modify an existing migration task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// RecycleDBInstance
// This API is used to return a deactivated SQL Server instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// ResetAccountPassword
// This API is used to reset the account password of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
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

// RestartDBInstance
// This API is used to restart a database instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// RestoreInstance
// This API is used to restore an instance from a backup file.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// RollbackInstance
// This API is used to roll back an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// RunMigration
// This API is used to start running a migration task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// StartBackupMigration
// This API is used to start a backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// StartIncrementalMigration
// This API is used to start an incremental backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// TerminateDBInstance
// This API is used to isolate an instance to move it into a recycle bin.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

// UpgradeDBInstance
// This API is used to upgrade an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
