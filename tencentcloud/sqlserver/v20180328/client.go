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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
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
    return c.CloneDBWithContext(context.Background(), request)
}

// CloneDB
// This API is used to clone and rename databases of an instance. The clones are still in the instance from which they are cloned.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
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
func (c *Client) CloneDBWithContext(ctx context.Context, request *CloneDBRequest) (response *CloneDBResponse, err error) {
    if request == nil {
        request = NewCloneDBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneDB require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneDBResponse()
    err = c.Send(request, response)
    return
}

func NewCloseInterCommunicationRequest() (request *CloseInterCommunicationRequest) {
    request = &CloseInterCommunicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CloseInterCommunication")
    
    
    return
}

func NewCloseInterCommunicationResponse() (response *CloseInterCommunicationResponse) {
    response = &CloseInterCommunicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseInterCommunication
// This API is used to disable instance interconnection.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseInterCommunication(request *CloseInterCommunicationRequest) (response *CloseInterCommunicationResponse, err error) {
    return c.CloseInterCommunicationWithContext(context.Background(), request)
}

// CloseInterCommunication
// This API is used to disable instance interconnection.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseInterCommunicationWithContext(ctx context.Context, request *CloseInterCommunicationRequest) (response *CloseInterCommunicationResponse, err error) {
    if request == nil {
        request = NewCloseInterCommunicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseInterCommunication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseInterCommunicationResponse()
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
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// This API is used to create a backup.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackup require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateBackupMigrationWithContext(context.Background(), request)
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
func (c *Client) CreateBackupMigrationWithContext(ctx context.Context, request *CreateBackupMigrationRequest) (response *CreateBackupMigrationResponse, err error) {
    if request == nil {
        request = NewCreateBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBusinessDBInstancesRequest() (request *CreateBusinessDBInstancesRequest) {
    request = &CreateBusinessDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBusinessDBInstances")
    
    
    return
}

func NewCreateBusinessDBInstancesResponse() (response *CreateBusinessDBInstancesResponse) {
    response = &CreateBusinessDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBusinessDBInstances
// This API is used to create a business intelligence service instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessDBInstances(request *CreateBusinessDBInstancesRequest) (response *CreateBusinessDBInstancesResponse, err error) {
    return c.CreateBusinessDBInstancesWithContext(context.Background(), request)
}

// CreateBusinessDBInstances
// This API is used to create a business intelligence service instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessDBInstancesWithContext(ctx context.Context, request *CreateBusinessDBInstancesRequest) (response *CreateBusinessDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateBusinessDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBusinessDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBusinessDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBusinessIntelligenceFileRequest() (request *CreateBusinessIntelligenceFileRequest) {
    request = &CreateBusinessIntelligenceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBusinessIntelligenceFile")
    
    
    return
}

func NewCreateBusinessIntelligenceFileResponse() (response *CreateBusinessIntelligenceFileResponse) {
    response = &CreateBusinessIntelligenceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBusinessIntelligenceFile
// This API is used to add a business intelligence service file.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessIntelligenceFile(request *CreateBusinessIntelligenceFileRequest) (response *CreateBusinessIntelligenceFileResponse, err error) {
    return c.CreateBusinessIntelligenceFileWithContext(context.Background(), request)
}

// CreateBusinessIntelligenceFile
// This API is used to add a business intelligence service file.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessIntelligenceFileWithContext(ctx context.Context, request *CreateBusinessIntelligenceFileRequest) (response *CreateBusinessIntelligenceFileResponse, err error) {
    if request == nil {
        request = NewCreateBusinessIntelligenceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBusinessIntelligenceFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudDBInstancesRequest() (request *CreateCloudDBInstancesRequest) {
    request = &CreateCloudDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateCloudDBInstances")
    
    
    return
}

func NewCreateCloudDBInstancesResponse() (response *CreateCloudDBInstancesResponse) {
    response = &CreateCloudDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudDBInstances
// This API is used to create a high-availability instance of cloud disk edition.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateCloudDBInstances(request *CreateCloudDBInstancesRequest) (response *CreateCloudDBInstancesResponse, err error) {
    return c.CreateCloudDBInstancesWithContext(context.Background(), request)
}

// CreateCloudDBInstances
// This API is used to create a high-availability instance of cloud disk edition.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateCloudDBInstancesWithContext(ctx context.Context, request *CreateCloudDBInstancesRequest) (response *CreateCloudDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateCloudDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudReadOnlyDBInstancesRequest() (request *CreateCloudReadOnlyDBInstancesRequest) {
    request = &CreateCloudReadOnlyDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateCloudReadOnlyDBInstances")
    
    
    return
}

func NewCreateCloudReadOnlyDBInstancesResponse() (response *CreateCloudReadOnlyDBInstancesResponse) {
    response = &CreateCloudReadOnlyDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudReadOnlyDBInstances
// This API is used to add a read-only replica instance of cloud disk edition.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_ROGROUPNAMEISILLEGAL = "InvalidParameterValue.RoGroupNameIsIllegal"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateCloudReadOnlyDBInstances(request *CreateCloudReadOnlyDBInstancesRequest) (response *CreateCloudReadOnlyDBInstancesResponse, err error) {
    return c.CreateCloudReadOnlyDBInstancesWithContext(context.Background(), request)
}

// CreateCloudReadOnlyDBInstances
// This API is used to add a read-only replica instance of cloud disk edition.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_ROGROUPNAMEISILLEGAL = "InvalidParameterValue.RoGroupNameIsIllegal"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateCloudReadOnlyDBInstancesWithContext(ctx context.Context, request *CreateCloudReadOnlyDBInstancesRequest) (response *CreateCloudReadOnlyDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateCloudReadOnlyDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudReadOnlyDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudReadOnlyDBInstancesResponse()
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
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
    return c.CreateDBWithContext(context.Background(), request)
}

// CreateDB
// This API is used to create a database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
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
func (c *Client) CreateDBWithContext(ctx context.Context, request *CreateDBRequest) (response *CreateDBResponse, err error) {
    if request == nil {
        request = NewCreateDBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDB require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    return c.CreateDBInstancesWithContext(context.Background(), request)
}

// CreateDBInstances
// This API is used to create an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDBInstancesWithContext(ctx context.Context, request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIncrementalMigration(request *CreateIncrementalMigrationRequest) (response *CreateIncrementalMigrationResponse, err error) {
    return c.CreateIncrementalMigrationWithContext(context.Background(), request)
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
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIncrementalMigrationWithContext(ctx context.Context, request *CreateIncrementalMigrationRequest) (response *CreateIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewCreateIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  INVALIDPARAMETERVALUE_ONCVMTYPENOTSUPPORTED = "InvalidParameterValue.OnCvmTypeNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateMigration(request *CreateMigrationRequest) (response *CreateMigrationResponse, err error) {
    return c.CreateMigrationWithContext(context.Background(), request)
}

// CreateMigration
// This API is used to create a migration task.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  INVALIDPARAMETERVALUE_ONCVMTYPENOTSUPPORTED = "InvalidParameterValue.OnCvmTypeNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateMigrationWithContext(ctx context.Context, request *CreateMigrationRequest) (response *CreateMigrationResponse, err error) {
    if request == nil {
        request = NewCreateMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigration require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAccountWithContext(context.Background(), request)
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
func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccount require credential")
    }

    request.SetContext(ctx)
    
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackupMigration(request *DeleteBackupMigrationRequest) (response *DeleteBackupMigrationResponse, err error) {
    return c.DeleteBackupMigrationWithContext(context.Background(), request)
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackupMigrationWithContext(ctx context.Context, request *DeleteBackupMigrationRequest) (response *DeleteBackupMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBusinessIntelligenceFileRequest() (request *DeleteBusinessIntelligenceFileRequest) {
    request = &DeleteBusinessIntelligenceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteBusinessIntelligenceFile")
    
    
    return
}

func NewDeleteBusinessIntelligenceFileResponse() (response *DeleteBusinessIntelligenceFileResponse) {
    response = &DeleteBusinessIntelligenceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBusinessIntelligenceFile
// This API is used to delete a business intelligence service file.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBusinessIntelligenceFile(request *DeleteBusinessIntelligenceFileRequest) (response *DeleteBusinessIntelligenceFileResponse, err error) {
    return c.DeleteBusinessIntelligenceFileWithContext(context.Background(), request)
}

// DeleteBusinessIntelligenceFile
// This API is used to delete a business intelligence service file.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBusinessIntelligenceFileWithContext(ctx context.Context, request *DeleteBusinessIntelligenceFileRequest) (response *DeleteBusinessIntelligenceFileResponse, err error) {
    if request == nil {
        request = NewDeleteBusinessIntelligenceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBusinessIntelligenceFileResponse()
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
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDB(request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
    return c.DeleteDBWithContext(context.Background(), request)
}

// DeleteDB
// This API is used to drop a database.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDBWithContext(ctx context.Context, request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
    if request == nil {
        request = NewDeleteDBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDB require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteIncrementalMigrationWithContext(context.Background(), request)
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
func (c *Client) DeleteIncrementalMigrationWithContext(ctx context.Context, request *DeleteIncrementalMigrationRequest) (response *DeleteIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
    return c.DeleteMigrationWithContext(context.Background(), request)
}

// DeleteMigration
// This API is used to delete a migration task.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteMigrationWithContext(ctx context.Context, request *DeleteMigrationRequest) (response *DeleteMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMigration require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// This API is used to pull the list of instance accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeBackupCommandWithContext(context.Background(), request)
}

// DescribeBackupCommand
// This API is used to query the commands of creating backups canonically.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupCommandWithContext(ctx context.Context, request *DescribeBackupCommandRequest) (response *DescribeBackupCommandResponse, err error) {
    if request == nil {
        request = NewDescribeBackupCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupFilesRequest() (request *DescribeBackupFilesRequest) {
    request = &DescribeBackupFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupFiles")
    
    
    return
}

func NewDescribeBackupFilesResponse() (response *DescribeBackupFilesResponse) {
    response = &DescribeBackupFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupFiles
// This API is used to query the details of an unarchived backup.
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
func (c *Client) DescribeBackupFiles(request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
    return c.DescribeBackupFilesWithContext(context.Background(), request)
}

// DescribeBackupFiles
// This API is used to query the details of an unarchived backup.
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
func (c *Client) DescribeBackupFilesWithContext(ctx context.Context, request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupFilesResponse()
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
    return c.DescribeBackupMigrationWithContext(context.Background(), request)
}

// DescribeBackupMigration
// This API is used to create an incremental backup import task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupMigrationWithContext(ctx context.Context, request *DescribeBackupMigrationRequest) (response *DescribeBackupMigrationResponse, err error) {
    if request == nil {
        request = NewDescribeBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupMigration require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeBackupUploadSizeWithContext(context.Background(), request)
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
func (c *Client) DescribeBackupUploadSizeWithContext(ctx context.Context, request *DescribeBackupUploadSizeRequest) (response *DescribeBackupUploadSizeResponse, err error) {
    if request == nil {
        request = NewDescribeBackupUploadSizeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupUploadSize require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_COSERROR = "FailedOperation.CosError"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    return c.DescribeBackupsWithContext(context.Background(), request)
}

// DescribeBackups
// This API is used to query the list of backups.
//
// error code that may be returned:
//  FAILEDOPERATION_COSERROR = "FailedOperation.CosError"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupsWithContext(ctx context.Context, request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBusinessIntelligenceFileRequest() (request *DescribeBusinessIntelligenceFileRequest) {
    request = &DescribeBusinessIntelligenceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBusinessIntelligenceFile")
    
    
    return
}

func NewDescribeBusinessIntelligenceFileResponse() (response *DescribeBusinessIntelligenceFileResponse) {
    response = &DescribeBusinessIntelligenceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBusinessIntelligenceFile
// This API is used to query the files required by business intelligence service.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBusinessIntelligenceFile(request *DescribeBusinessIntelligenceFileRequest) (response *DescribeBusinessIntelligenceFileResponse, err error) {
    return c.DescribeBusinessIntelligenceFileWithContext(context.Background(), request)
}

// DescribeBusinessIntelligenceFile
// This API is used to query the files required by business intelligence service.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBusinessIntelligenceFileWithContext(ctx context.Context, request *DescribeBusinessIntelligenceFileRequest) (response *DescribeBusinessIntelligenceFileResponse, err error) {
    if request == nil {
        request = NewDescribeBusinessIntelligenceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBusinessIntelligenceFileResponse()
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
    return c.DescribeDBCharsetsWithContext(context.Background(), request)
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
func (c *Client) DescribeDBCharsetsWithContext(ctx context.Context, request *DescribeDBCharsetsRequest) (response *DescribeDBCharsetsResponse, err error) {
    if request == nil {
        request = NewDescribeDBCharsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCharsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCharsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceInterRequest() (request *DescribeDBInstanceInterRequest) {
    request = &DescribeDBInstanceInterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstanceInter")
    
    
    return
}

func NewDescribeDBInstanceInterResponse() (response *DescribeDBInstanceInterResponse) {
    response = &DescribeDBInstanceInterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceInter
// This API is used to query the information of the interconnected instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstanceInter(request *DescribeDBInstanceInterRequest) (response *DescribeDBInstanceInterResponse, err error) {
    return c.DescribeDBInstanceInterWithContext(context.Background(), request)
}

// DescribeDBInstanceInter
// This API is used to query the information of the interconnected instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstanceInterWithContext(ctx context.Context, request *DescribeDBInstanceInterRequest) (response *DescribeDBInstanceInterResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceInterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceInter require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceInterResponse()
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// This API is used to query the list of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesAttributeRequest() (request *DescribeDBInstancesAttributeRequest) {
    request = &DescribeDBInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstancesAttribute")
    
    
    return
}

func NewDescribeDBInstancesAttributeResponse() (response *DescribeDBInstancesAttributeResponse) {
    response = &DescribeDBInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstancesAttribute
// This API is used to query the attributes of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstancesAttribute(request *DescribeDBInstancesAttributeRequest) (response *DescribeDBInstancesAttributeResponse, err error) {
    return c.DescribeDBInstancesAttributeWithContext(context.Background(), request)
}

// DescribeDBInstancesAttribute
// This API is used to query the attributes of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstancesAttributeWithContext(ctx context.Context, request *DescribeDBInstancesAttributeRequest) (response *DescribeDBInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstancesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesAttributeResponse()
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBs(request *DescribeDBsRequest) (response *DescribeDBsResponse, err error) {
    return c.DescribeDBsWithContext(context.Background(), request)
}

// DescribeDBs
// This API is used to query the list of databases
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBsWithContext(ctx context.Context, request *DescribeDBsRequest) (response *DescribeDBsResponse, err error) {
    if request == nil {
        request = NewDescribeDBsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBsNormalRequest() (request *DescribeDBsNormalRequest) {
    request = &DescribeDBsNormalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBsNormal")
    
    
    return
}

func NewDescribeDBsNormalResponse() (response *DescribeDBsNormalResponse) {
    response = &DescribeDBsNormalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBsNormal
// This API is used to query database configurations. It does not return information of the accounts that have permissions to operate the database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBsNormal(request *DescribeDBsNormalRequest) (response *DescribeDBsNormalResponse, err error) {
    return c.DescribeDBsNormalWithContext(context.Background(), request)
}

// DescribeDBsNormal
// This API is used to query database configurations. It does not return information of the accounts that have permissions to operate the database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBsNormalWithContext(ctx context.Context, request *DescribeDBsNormalRequest) (response *DescribeDBsNormalResponse, err error) {
    if request == nil {
        request = NewDescribeDBsNormalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBsNormal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBsNormalResponse()
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
    return c.DescribeFlowStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeFlowStatusWithContext(ctx context.Context, request *DescribeFlowStatusRequest) (response *DescribeFlowStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFlowStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeIncrementalMigrationWithContext(context.Background(), request)
}

// DescribeIncrementalMigration
// This API is used to query an incremental backup import task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIncrementalMigrationWithContext(ctx context.Context, request *DescribeIncrementalMigrationRequest) (response *DescribeIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewDescribeIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceByOrdersRequest() (request *DescribeInstanceByOrdersRequest) {
    request = &DescribeInstanceByOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceByOrders")
    
    
    return
}

func NewDescribeInstanceByOrdersResponse() (response *DescribeInstanceByOrdersResponse) {
    response = &DescribeInstanceByOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceByOrders
// This API is used to query the instance ID by the order number.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
func (c *Client) DescribeInstanceByOrders(request *DescribeInstanceByOrdersRequest) (response *DescribeInstanceByOrdersResponse, err error) {
    return c.DescribeInstanceByOrdersWithContext(context.Background(), request)
}

// DescribeInstanceByOrders
// This API is used to query the instance ID by the order number.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
func (c *Client) DescribeInstanceByOrdersWithContext(ctx context.Context, request *DescribeInstanceByOrdersRequest) (response *DescribeInstanceByOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceByOrdersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceByOrders require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceByOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceParamRecords")
    
    
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParamRecords
// This API is used to query the parameter modification records of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// This API is used to query the parameter modification records of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamRecordsWithContext(ctx context.Context, request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParamRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParams
// This API is used to query the parameter list of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// This API is used to query the parameter list of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationDetail(request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    return c.DescribeMigrationDetailWithContext(context.Background(), request)
}

// DescribeMigrationDetail
// This API is used to query migration task details.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationDetailWithContext(ctx context.Context, request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationDetail require credential")
    }

    request.SetContext(ctx)
    
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrations(request *DescribeMigrationsRequest) (response *DescribeMigrationsResponse, err error) {
    return c.DescribeMigrationsWithContext(context.Background(), request)
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationsWithContext(ctx context.Context, request *DescribeMigrationsRequest) (response *DescribeMigrationsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrations require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeOrdersWithContext(context.Background(), request)
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
func (c *Client) DescribeOrdersWithContext(ctx context.Context, request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrders require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProductConfig(request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    return c.DescribeProductConfigWithContext(context.Background(), request)
}

// DescribeProductConfig
// This API is used to query purchasable specification configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProductConfigWithContext(ctx context.Context, request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    if request == nil {
        request = NewDescribeProductConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductConfig require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// This API is used to query purchasable regions.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeRollbackTimeWithContext(context.Background(), request)
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
func (c *Client) DescribeRollbackTimeWithContext(ctx context.Context, request *DescribeRollbackTimeRequest) (response *DescribeRollbackTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTime require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSlowlogs(request *DescribeSlowlogsRequest) (response *DescribeSlowlogsResponse, err error) {
    return c.DescribeSlowlogsWithContext(context.Background(), request)
}

// DescribeSlowlogs
// This API is used to get file information of slow query logs.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSlowlogsWithContext(ctx context.Context, request *DescribeSlowlogsRequest) (response *DescribeSlowlogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowlogs require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeUploadBackupInfoWithContext(context.Background(), request)
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
func (c *Client) DescribeUploadBackupInfoWithContext(ctx context.Context, request *DescribeUploadBackupInfoRequest) (response *DescribeUploadBackupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadBackupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadBackupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUploadBackupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeXEventsRequest() (request *DescribeXEventsRequest) {
    request = &DescribeXEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeXEvents")
    
    
    return
}

func NewDescribeXEventsResponse() (response *DescribeXEventsResponse) {
    response = &DescribeXEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeXEvents
// This API is used to query the list of extended events.
//
// error code that may be returned:
//  FAILEDOPERATION_COSERROR = "FailedOperation.CosError"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeXEvents(request *DescribeXEventsRequest) (response *DescribeXEventsResponse, err error) {
    return c.DescribeXEventsWithContext(context.Background(), request)
}

// DescribeXEvents
// This API is used to query the list of extended events.
//
// error code that may be returned:
//  FAILEDOPERATION_COSERROR = "FailedOperation.CosError"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeXEventsWithContext(ctx context.Context, request *DescribeXEventsRequest) (response *DescribeXEventsResponse, err error) {
    if request == nil {
        request = NewDescribeXEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeXEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeXEventsResponse()
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
    return c.DescribeZonesWithContext(context.Background(), request)
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
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
    return c.InquiryPriceCreateDBInstancesWithContext(context.Background(), request)
}

// InquiryPriceCreateDBInstances
// This API is used to query the price of requested instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
func (c *Client) InquiryPriceCreateDBInstancesWithContext(ctx context.Context, request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to query the upgrade prices of a monthly subscribed instance
//
// .
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
    return c.InquiryPriceUpgradeDBInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpgradeDBInstance
// This API is used to query the upgrade prices of a monthly subscribed instance
//
// .
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
func (c *Client) InquiryPriceUpgradeDBInstanceWithContext(ctx context.Context, request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (response *ModifyAccountPrivilegeResponse, err error) {
    return c.ModifyAccountPrivilegeWithContext(context.Background(), request)
}

// ModifyAccountPrivilege
// This API is used to modify instance account permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilegeWithContext(ctx context.Context, request *ModifyAccountPrivilegeRequest) (response *ModifyAccountPrivilegeResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivilege require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyAccountRemarkWithContext(context.Background(), request)
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
func (c *Client) ModifyAccountRemarkWithContext(ctx context.Context, request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    if request == nil {
        request = NewModifyAccountRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountRemark require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyBackupMigrationWithContext(context.Background(), request)
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
func (c *Client) ModifyBackupMigrationWithContext(ctx context.Context, request *ModifyBackupMigrationRequest) (response *ModifyBackupMigrationResponse, err error) {
    if request == nil {
        request = NewModifyBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupMigration require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
    return c.ModifyBackupStrategyWithContext(context.Background(), request)
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
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupStrategyWithContext(ctx context.Context, request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
    if request == nil {
        request = NewModifyBackupStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBEncryptAttributesRequest() (request *ModifyDBEncryptAttributesRequest) {
    request = &ModifyDBEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBEncryptAttributes")
    
    
    return
}

func NewModifyDBEncryptAttributesResponse() (response *ModifyDBEncryptAttributesResponse) {
    response = &ModifyDBEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBEncryptAttributes
// This API is used to u200denable or disable TDE of a database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBEncryptAttributes(request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    return c.ModifyDBEncryptAttributesWithContext(context.Background(), request)
}

// ModifyDBEncryptAttributes
// This API is used to u200denable or disable TDE of a database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBEncryptAttributesWithContext(ctx context.Context, request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBEncryptAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBEncryptAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBEncryptAttributesResponse()
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
    return c.ModifyDBInstanceNameWithContext(context.Background(), request)
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
func (c *Client) ModifyDBInstanceNameWithContext(ctx context.Context, request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceName require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_VPCERROR = "FailedOperation.VPCError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBInstanceNetwork(request *ModifyDBInstanceNetworkRequest) (response *ModifyDBInstanceNetworkResponse, err error) {
    return c.ModifyDBInstanceNetworkWithContext(context.Background(), request)
}

// ModifyDBInstanceNetwork
// This API is used to switch a running instance from a VPC to another.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_VPCERROR = "FailedOperation.VPCError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBInstanceNetworkWithContext(ctx context.Context, request *ModifyDBInstanceNetworkRequest) (response *ModifyDBInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNetwork require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    return c.ModifyDBInstanceProjectWithContext(context.Background(), request)
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
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceProjectWithContext(ctx context.Context, request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceProject require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyDBNameWithContext(context.Background(), request)
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
func (c *Client) ModifyDBNameWithContext(ctx context.Context, request *ModifyDBNameRequest) (response *ModifyDBNameResponse, err error) {
    if request == nil {
        request = NewModifyDBNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBName require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyDBRemarkWithContext(context.Background(), request)
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
func (c *Client) ModifyDBRemarkWithContext(ctx context.Context, request *ModifyDBRemarkRequest) (response *ModifyDBRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDBRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseCDCRequest() (request *ModifyDatabaseCDCRequest) {
    request = &ModifyDatabaseCDCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseCDC")
    
    
    return
}

func NewModifyDatabaseCDCResponse() (response *ModifyDatabaseCDCResponse) {
    response = &ModifyDatabaseCDCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseCDC
// This API is used to enable or disable the change data capture (CDC) feature.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseCDC(request *ModifyDatabaseCDCRequest) (response *ModifyDatabaseCDCResponse, err error) {
    return c.ModifyDatabaseCDCWithContext(context.Background(), request)
}

// ModifyDatabaseCDC
// This API is used to enable or disable the change data capture (CDC) feature.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseCDCWithContext(ctx context.Context, request *ModifyDatabaseCDCRequest) (response *ModifyDatabaseCDCResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseCDCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseCDC require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseCDCResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseCTRequest() (request *ModifyDatabaseCTRequest) {
    request = &ModifyDatabaseCTRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseCT")
    
    
    return
}

func NewModifyDatabaseCTResponse() (response *ModifyDatabaseCTResponse) {
    response = &ModifyDatabaseCTResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseCT
// This API is used to enable or disable the change tracking (CT) feature.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDatabaseCT(request *ModifyDatabaseCTRequest) (response *ModifyDatabaseCTResponse, err error) {
    return c.ModifyDatabaseCTWithContext(context.Background(), request)
}

// ModifyDatabaseCT
// This API is used to enable or disable the change tracking (CT) feature.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDatabaseCTWithContext(ctx context.Context, request *ModifyDatabaseCTRequest) (response *ModifyDatabaseCTResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseCTRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseCT require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseCTResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseMdfRequest() (request *ModifyDatabaseMdfRequest) {
    request = &ModifyDatabaseMdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseMdf")
    
    
    return
}

func NewModifyDatabaseMdfResponse() (response *ModifyDatabaseMdfResponse) {
    response = &ModifyDatabaseMdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseMdf
// This API is used to shrink database MDF files.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseMdf(request *ModifyDatabaseMdfRequest) (response *ModifyDatabaseMdfResponse, err error) {
    return c.ModifyDatabaseMdfWithContext(context.Background(), request)
}

// ModifyDatabaseMdf
// This API is used to shrink database MDF files.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseMdfWithContext(ctx context.Context, request *ModifyDatabaseMdfRequest) (response *ModifyDatabaseMdfResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseMdfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseMdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseMdfResponse()
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
    return c.ModifyIncrementalMigrationWithContext(context.Background(), request)
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
func (c *Client) ModifyIncrementalMigrationWithContext(ctx context.Context, request *ModifyIncrementalMigrationRequest) (response *ModifyIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewModifyIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceEncryptAttributesRequest() (request *ModifyInstanceEncryptAttributesRequest) {
    request = &ModifyInstanceEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceEncryptAttributes")
    
    
    return
}

func NewModifyInstanceEncryptAttributesResponse() (response *ModifyInstanceEncryptAttributesResponse) {
    response = &ModifyInstanceEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceEncryptAttributes
// This API is used to enable TDE of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyInstanceEncryptAttributes(request *ModifyInstanceEncryptAttributesRequest) (response *ModifyInstanceEncryptAttributesResponse, err error) {
    return c.ModifyInstanceEncryptAttributesWithContext(context.Background(), request)
}

// ModifyInstanceEncryptAttributes
// This API is used to enable TDE of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyInstanceEncryptAttributesWithContext(ctx context.Context, request *ModifyInstanceEncryptAttributesRequest) (response *ModifyInstanceEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewModifyInstanceEncryptAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceEncryptAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceEncryptAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceParam")
    
    
    return
}

func NewModifyInstanceParamResponse() (response *ModifyInstanceParamResponse) {
    response = &ModifyInstanceParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceParam
// This API is used to modify instance parameters.
//
// <b>Note</b>: if <b>the instance needs to be restarted</b> for the modified parameter to take effect, <b>it will be restarted</b> immediately or during the maintenance time according to the `WaitSwitch` parameter.
//
// Before you modify a parameter, you can use the `DescribeInstanceParams` API to query whether the instance needs to be restarted.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    return c.ModifyInstanceParamWithContext(context.Background(), request)
}

// ModifyInstanceParam
// This API is used to modify instance parameters.
//
// <b>Note</b>: if <b>the instance needs to be restarted</b> for the modified parameter to take effect, <b>it will be restarted</b> immediately or during the maintenance time according to the `WaitSwitch` parameter.
//
// Before you modify a parameter, you can use the `DescribeInstanceParams` API to query whether the instance needs to be restarted.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyInstanceParamWithContext(ctx context.Context, request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamResponse()
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
    return c.ModifyMigrationWithContext(context.Background(), request)
}

// ModifyMigration
// This API is used to modify an existing migration task.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
func (c *Client) ModifyMigrationWithContext(ctx context.Context, request *ModifyMigrationRequest) (response *ModifyMigrationResponse, err error) {
    if request == nil {
        request = NewModifyMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewOpenInterCommunicationRequest() (request *OpenInterCommunicationRequest) {
    request = &OpenInterCommunicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "OpenInterCommunication")
    
    
    return
}

func NewOpenInterCommunicationResponse() (response *OpenInterCommunicationResponse) {
    response = &OpenInterCommunicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenInterCommunication
// This API is used to enable instance interconnection, which can interconnect business intelligence services.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenInterCommunication(request *OpenInterCommunicationRequest) (response *OpenInterCommunicationResponse, err error) {
    return c.OpenInterCommunicationWithContext(context.Background(), request)
}

// OpenInterCommunication
// This API is used to enable instance interconnection, which can interconnect business intelligence services.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenInterCommunicationWithContext(ctx context.Context, request *OpenInterCommunicationRequest) (response *OpenInterCommunicationResponse, err error) {
    if request == nil {
        request = NewOpenInterCommunicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenInterCommunication require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenInterCommunicationResponse()
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
    return c.RecycleDBInstanceWithContext(context.Background(), request)
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
func (c *Client) RecycleDBInstanceWithContext(ctx context.Context, request *RecycleDBInstanceRequest) (response *RecycleDBInstanceResponse, err error) {
    if request == nil {
        request = NewRecycleDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecycleDBInstance require credential")
    }

    request.SetContext(ctx)
    
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RestartDBInstanceWithContext(context.Background(), request)
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
func (c *Client) RestartDBInstanceWithContext(ctx context.Context, request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestartDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDBInstance require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to roll back the database by backup set.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    return c.RestoreInstanceWithContext(context.Background(), request)
}

// RestoreInstance
// This API is used to roll back the database by backup set.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestoreInstanceWithContext(ctx context.Context, request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    if request == nil {
        request = NewRestoreInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreInstance require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to roll back the instance by time point.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollbackInstance(request *RollbackInstanceRequest) (response *RollbackInstanceResponse, err error) {
    return c.RollbackInstanceWithContext(context.Background(), request)
}

// RollbackInstance
// This API is used to roll back the instance by time point.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollbackInstanceWithContext(ctx context.Context, request *RollbackInstanceRequest) (response *RollbackInstanceResponse, err error) {
    if request == nil {
        request = NewRollbackInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackInstance require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RunMigration(request *RunMigrationRequest) (response *RunMigrationResponse, err error) {
    return c.RunMigrationWithContext(context.Background(), request)
}

// RunMigration
// This API is used to start running a migration task.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RunMigrationWithContext(ctx context.Context, request *RunMigrationRequest) (response *RunMigrationResponse, err error) {
    if request == nil {
        request = NewRunMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunMigration require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
    return c.StartBackupMigrationWithContext(context.Background(), request)
}

// StartBackupMigration
// This API is used to start a backup import task.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
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
func (c *Client) StartBackupMigrationWithContext(ctx context.Context, request *StartBackupMigrationRequest) (response *StartBackupMigrationResponse, err error) {
    if request == nil {
        request = NewStartBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBackupMigration require credential")
    }

    request.SetContext(ctx)
    
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
    return c.StartIncrementalMigrationWithContext(context.Background(), request)
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
func (c *Client) StartIncrementalMigrationWithContext(ctx context.Context, request *StartIncrementalMigrationRequest) (response *StartIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewStartIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstanceXEventRequest() (request *StartInstanceXEventRequest) {
    request = &StartInstanceXEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartInstanceXEvent")
    
    
    return
}

func NewStartInstanceXEventResponse() (response *StartInstanceXEventResponse) {
    response = &StartInstanceXEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartInstanceXEvent
// This API is used to start and stop an extended event.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
func (c *Client) StartInstanceXEvent(request *StartInstanceXEventRequest) (response *StartInstanceXEventResponse, err error) {
    return c.StartInstanceXEventWithContext(context.Background(), request)
}

// StartInstanceXEvent
// This API is used to start and stop an extended event.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
func (c *Client) StartInstanceXEventWithContext(ctx context.Context, request *StartInstanceXEventRequest) (response *StartInstanceXEventResponse, err error) {
    if request == nil {
        request = NewStartInstanceXEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstanceXEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartInstanceXEventResponse()
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
    return c.TerminateDBInstanceWithContext(context.Background(), request)
}

// TerminateDBInstance
// This API is used to isolate an instance to move it into a recycle bin.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDBInstanceWithContext(ctx context.Context, request *TerminateDBInstanceRequest) (response *TerminateDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDBInstance require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    return c.UpgradeDBInstanceWithContext(context.Background(), request)
}

// UpgradeDBInstance
// This API is used to upgrade an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDBInstanceWithContext(ctx context.Context, request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
