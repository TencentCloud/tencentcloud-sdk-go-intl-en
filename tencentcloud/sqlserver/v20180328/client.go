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


func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateSecurityGroups
// This API is used to bind security groups to instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// This API is used to bind security groups to instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "AssociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewBalanceReadOnlyGroupRequest() (request *BalanceReadOnlyGroupRequest) {
    request = &BalanceReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "BalanceReadOnlyGroup")
    
    
    return
}

func NewBalanceReadOnlyGroupResponse() (response *BalanceReadOnlyGroupResponse) {
    response = &BalanceReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BalanceReadOnlyGroup
// This API is used to balance the routing weight of each read-only instance according to the predefined weights. The DescribeReadOnlyGroupAutoWeight API is used to query the predefined weights.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) BalanceReadOnlyGroup(request *BalanceReadOnlyGroupRequest) (response *BalanceReadOnlyGroupResponse, err error) {
    return c.BalanceReadOnlyGroupWithContext(context.Background(), request)
}

// BalanceReadOnlyGroup
// This API is used to balance the routing weight of each read-only instance according to the predefined weights. The DescribeReadOnlyGroupAutoWeight API is used to query the predefined weights.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) BalanceReadOnlyGroupWithContext(ctx context.Context, request *BalanceReadOnlyGroupRequest) (response *BalanceReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewBalanceReadOnlyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "BalanceReadOnlyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BalanceReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewBalanceReadOnlyGroupResponse()
    err = c.Send(request, response)
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CloneDB")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CloseInterCommunication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseInterCommunication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseInterCommunicationResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteExpansionRequest() (request *CompleteExpansionRequest) {
    request = &CompleteExpansionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CompleteExpansion")
    
    
    return
}

func NewCompleteExpansionResponse() (response *CompleteExpansionResponse) {
    response = &CompleteExpansionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CompleteExpansion
// This API is used to complete the instance upgrade and switch immediately without waiting for the maintenance window when the instance status is "upgrade pending switch" after scale-out is initiated. This API needs to be called during off-peak hours of the instance. Some databases cannot be accessed before the switch is completed.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CompleteExpansion(request *CompleteExpansionRequest) (response *CompleteExpansionResponse, err error) {
    return c.CompleteExpansionWithContext(context.Background(), request)
}

// CompleteExpansion
// This API is used to complete the instance upgrade and switch immediately without waiting for the maintenance window when the instance status is "upgrade pending switch" after scale-out is initiated. This API needs to be called during off-peak hours of the instance. Some databases cannot be accessed before the switch is completed.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CompleteExpansionWithContext(ctx context.Context, request *CompleteExpansionRequest) (response *CompleteExpansionResponse, err error) {
    if request == nil {
        request = NewCompleteExpansionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CompleteExpansion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteExpansion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteExpansionResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateAccount")
    
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
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
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
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateBackup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateBackupMigration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBasicDBInstancesRequest() (request *CreateBasicDBInstancesRequest) {
    request = &CreateBasicDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBasicDBInstances")
    
    
    return
}

func NewCreateBasicDBInstancesResponse() (response *CreateBasicDBInstancesResponse) {
    response = &CreateBasicDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBasicDBInstances
// This API is used to create basic edition instances (cloud disk).
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBasicDBInstances(request *CreateBasicDBInstancesRequest) (response *CreateBasicDBInstancesResponse, err error) {
    return c.CreateBasicDBInstancesWithContext(context.Background(), request)
}

// CreateBasicDBInstances
// This API is used to create basic edition instances (cloud disk).
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBasicDBInstancesWithContext(ctx context.Context, request *CreateBasicDBInstancesRequest) (response *CreateBasicDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateBasicDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateBasicDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBasicDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBasicDBInstancesResponse()
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
// This API is used to create business intelligence service instances (cloud disk).
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
// This API is used to create business intelligence service instances (cloud disk).
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateBusinessDBInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateBusinessIntelligenceFile")
    
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
// This API is used to create high-availability instances (cloud disk).
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
// This API is used to create high-availability instances (cloud disk).
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateCloudDBInstances")
    
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
// This API is used to create read-only instances (cloud disk).
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
// This API is used to create read-only instances (cloud disk).
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateCloudReadOnlyDBInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateDB")
    
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
// This API is used to create high-availability instances (local disk).
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
// This API is used to create high-availability instances (local disk).
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateDBInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateIncrementalMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateMigration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublishSubscribeRequest() (request *CreatePublishSubscribeRequest) {
    request = &CreatePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreatePublishSubscribe")
    
    
    return
}

func NewCreatePublishSubscribeResponse() (response *CreatePublishSubscribeResponse) {
    response = &CreatePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePublishSubscribe
// This API is used to create a publish/subscribe relationship between two databases. A subscriber cannot act as a publisher, and a publisher can have multiple subscriber instances.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreatePublishSubscribe(request *CreatePublishSubscribeRequest) (response *CreatePublishSubscribeResponse, err error) {
    return c.CreatePublishSubscribeWithContext(context.Background(), request)
}

// CreatePublishSubscribe
// This API is used to create a publish/subscribe relationship between two databases. A subscriber cannot act as a publisher, and a publisher can have multiple subscriber instances.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreatePublishSubscribeWithContext(ctx context.Context, request *CreatePublishSubscribeRequest) (response *CreatePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewCreatePublishSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreatePublishSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyDBInstancesRequest() (request *CreateReadOnlyDBInstancesRequest) {
    request = &CreateReadOnlyDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateReadOnlyDBInstances")
    
    
    return
}

func NewCreateReadOnlyDBInstancesResponse() (response *CreateReadOnlyDBInstancesResponse) {
    response = &CreateReadOnlyDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReadOnlyDBInstances
// This API is used to create read-only instances (local disk).
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateReadOnlyDBInstances(request *CreateReadOnlyDBInstancesRequest) (response *CreateReadOnlyDBInstancesResponse, err error) {
    return c.CreateReadOnlyDBInstancesWithContext(context.Background(), request)
}

// CreateReadOnlyDBInstances
// This API is used to create read-only instances (local disk).
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateReadOnlyDBInstancesWithContext(ctx context.Context, request *CreateReadOnlyDBInstancesRequest) (response *CreateReadOnlyDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CreateReadOnlyDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReadOnlyDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReadOnlyDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCutXEventsRequest() (request *CutXEventsRequest) {
    request = &CutXEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CutXEvents")
    
    
    return
}

func NewCutXEventsResponse() (response *CutXEventsResponse) {
    response = &CutXEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CutXEvents
// This API is used to manually cut block logs and deadlock logs.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CutXEvents(request *CutXEventsRequest) (response *CutXEventsResponse, err error) {
    return c.CutXEventsWithContext(context.Background(), request)
}

// CutXEvents
// This API is used to manually cut block logs and deadlock logs.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CutXEventsWithContext(ctx context.Context, request *CutXEventsRequest) (response *CutXEventsResponse, err error) {
    if request == nil {
        request = NewCutXEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "CutXEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CutXEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewCutXEventsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteAccount")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteBackupMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteBusinessIntelligenceFile")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteDB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
    request = &DeleteDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDBInstance")
    
    
    return
}

func NewDeleteDBInstanceResponse() (response *DeleteDBInstanceResponse) {
    response = &DeleteDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBInstance
// This API is used to release SQL server instances (eliminated immediately) in the recycle bin. The released instances will be physically terminated after being retained for a period of time. Their publish-subscribe relationships will be automatically disassociated, and their RO replicas will be automatically released.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (response *DeleteDBInstanceResponse, err error) {
    return c.DeleteDBInstanceWithContext(context.Background(), request)
}

// DeleteDBInstance
// This API is used to release SQL server instances (eliminated immediately) in the recycle bin. The released instances will be physically terminated after being retained for a period of time. Their publish-subscribe relationships will be automatically disassociated, and their RO replicas will be automatically released.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDBInstanceWithContext(ctx context.Context, request *DeleteDBInstanceRequest) (response *DeleteDBInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBInstanceResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteIncrementalMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeleteMigration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublishSubscribeRequest() (request *DeletePublishSubscribeRequest) {
    request = &DeletePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeletePublishSubscribe")
    
    
    return
}

func NewDeletePublishSubscribeResponse() (response *DeletePublishSubscribeResponse) {
    response = &DeletePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePublishSubscribe
// This API is used to delete the publish/subscribe relationship between two databases.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeletePublishSubscribe(request *DeletePublishSubscribeRequest) (response *DeletePublishSubscribeResponse, err error) {
    return c.DeletePublishSubscribeWithContext(context.Background(), request)
}

// DeletePublishSubscribe
// This API is used to delete the publish/subscribe relationship between two databases.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeletePublishSubscribeWithContext(ctx context.Context, request *DeletePublishSubscribeRequest) (response *DeletePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewDeletePublishSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DeletePublishSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegeByDBRequest() (request *DescribeAccountPrivilegeByDBRequest) {
    request = &DescribeAccountPrivilegeByDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeAccountPrivilegeByDB")
    
    
    return
}

func NewDescribeAccountPrivilegeByDBResponse() (response *DescribeAccountPrivilegeByDBResponse) {
    response = &DescribeAccountPrivilegeByDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountPrivilegeByDB
// This API is used to query information on the account and permissions associated with the database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivilegeByDB(request *DescribeAccountPrivilegeByDBRequest) (response *DescribeAccountPrivilegeByDBResponse, err error) {
    return c.DescribeAccountPrivilegeByDBWithContext(context.Background(), request)
}

// DescribeAccountPrivilegeByDB
// This API is used to query information on the account and permissions associated with the database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivilegeByDBWithContext(ctx context.Context, request *DescribeAccountPrivilegeByDBRequest) (response *DescribeAccountPrivilegeByDBResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegeByDBRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeAccountPrivilegeByDB")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountPrivilegeByDB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountPrivilegeByDBResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupByFlowIdRequest() (request *DescribeBackupByFlowIdRequest) {
    request = &DescribeBackupByFlowIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupByFlowId")
    
    
    return
}

func NewDescribeBackupByFlowIdResponse() (response *DescribeBackupByFlowIdResponse) {
    response = &DescribeBackupByFlowIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupByFlowId
// This API is used to query the created backup details through the backup creation process ID. The process ID can be obtained through the CreateBackup API.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupByFlowId(request *DescribeBackupByFlowIdRequest) (response *DescribeBackupByFlowIdResponse, err error) {
    return c.DescribeBackupByFlowIdWithContext(context.Background(), request)
}

// DescribeBackupByFlowId
// This API is used to query the created backup details through the backup creation process ID. The process ID can be obtained through the CreateBackup API.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupByFlowIdWithContext(ctx context.Context, request *DescribeBackupByFlowIdRequest) (response *DescribeBackupByFlowIdResponse, err error) {
    if request == nil {
        request = NewDescribeBackupByFlowIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupByFlowId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupByFlowId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupByFlowIdResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupCommand")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupFiles")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupMigration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupMonitorRequest() (request *DescribeBackupMonitorRequest) {
    request = &DescribeBackupMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupMonitor")
    
    
    return
}

func NewDescribeBackupMonitorResponse() (response *DescribeBackupMonitorResponse) {
    response = &DescribeBackupMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupMonitor
// This API is used to query backup space usage details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupMonitor(request *DescribeBackupMonitorRequest) (response *DescribeBackupMonitorResponse, err error) {
    return c.DescribeBackupMonitorWithContext(context.Background(), request)
}

// DescribeBackupMonitor
// This API is used to query backup space usage details.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupMonitorWithContext(ctx context.Context, request *DescribeBackupMonitorRequest) (response *DescribeBackupMonitorResponse, err error) {
    if request == nil {
        request = NewDescribeBackupMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupStatisticalRequest() (request *DescribeBackupStatisticalRequest) {
    request = &DescribeBackupStatisticalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupStatistical")
    
    
    return
}

func NewDescribeBackupStatisticalResponse() (response *DescribeBackupStatisticalResponse) {
    response = &DescribeBackupStatisticalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupStatistical
// This API is used to query the real-time statistics list of backups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupStatistical(request *DescribeBackupStatisticalRequest) (response *DescribeBackupStatisticalResponse, err error) {
    return c.DescribeBackupStatisticalWithContext(context.Background(), request)
}

// DescribeBackupStatistical
// This API is used to query the real-time statistics list of backups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupStatisticalWithContext(ctx context.Context, request *DescribeBackupStatisticalRequest) (response *DescribeBackupStatisticalResponse, err error) {
    if request == nil {
        request = NewDescribeBackupStatisticalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupStatistical")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupStatistical require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupStatisticalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupSummaryRequest() (request *DescribeBackupSummaryRequest) {
    request = &DescribeBackupSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupSummary")
    
    
    return
}

func NewDescribeBackupSummaryResponse() (response *DescribeBackupSummaryResponse) {
    response = &DescribeBackupSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupSummary
// This API is used to query the backup overview information of databases.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupSummary(request *DescribeBackupSummaryRequest) (response *DescribeBackupSummaryResponse, err error) {
    return c.DescribeBackupSummaryWithContext(context.Background(), request)
}

// DescribeBackupSummary
// This API is used to query the backup overview information of databases.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupSummaryWithContext(ctx context.Context, request *DescribeBackupSummaryRequest) (response *DescribeBackupSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBackupSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupSummaryResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackupUploadSize")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBackups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeBusinessIntelligenceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBusinessIntelligenceFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCollationTimeZoneRequest() (request *DescribeCollationTimeZoneRequest) {
    request = &DescribeCollationTimeZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCollationTimeZone")
    
    
    return
}

func NewDescribeCollationTimeZoneResponse() (response *DescribeCollationTimeZoneResponse) {
    response = &DescribeCollationTimeZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCollationTimeZone
// This API is used to query the character set and time zone supported by the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCollationTimeZone(request *DescribeCollationTimeZoneRequest) (response *DescribeCollationTimeZoneResponse, err error) {
    return c.DescribeCollationTimeZoneWithContext(context.Background(), request)
}

// DescribeCollationTimeZone
// This API is used to query the character set and time zone supported by the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCollationTimeZoneWithContext(ctx context.Context, request *DescribeCollationTimeZoneRequest) (response *DescribeCollationTimeZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCollationTimeZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeCollationTimeZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCollationTimeZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCollationTimeZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossBackupStatisticalRequest() (request *DescribeCrossBackupStatisticalRequest) {
    request = &DescribeCrossBackupStatisticalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCrossBackupStatistical")
    
    
    return
}

func NewDescribeCrossBackupStatisticalResponse() (response *DescribeCrossBackupStatisticalResponse) {
    response = &DescribeCrossBackupStatisticalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossBackupStatistical
// This API is used to query the real-time statistics list of cross-region backups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossBackupStatistical(request *DescribeCrossBackupStatisticalRequest) (response *DescribeCrossBackupStatisticalResponse, err error) {
    return c.DescribeCrossBackupStatisticalWithContext(context.Background(), request)
}

// DescribeCrossBackupStatistical
// This API is used to query the real-time statistics list of cross-region backups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossBackupStatisticalWithContext(ctx context.Context, request *DescribeCrossBackupStatisticalRequest) (response *DescribeCrossBackupStatisticalResponse, err error) {
    if request == nil {
        request = NewDescribeCrossBackupStatisticalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeCrossBackupStatistical")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossBackupStatistical require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossBackupStatisticalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossRegionZoneRequest() (request *DescribeCrossRegionZoneRequest) {
    request = &DescribeCrossRegionZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCrossRegionZone")
    
    
    return
}

func NewDescribeCrossRegionZoneResponse() (response *DescribeCrossRegionZoneResponse) {
    response = &DescribeCrossRegionZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossRegionZone
// This API is used to query the disaster recovery region and AZ of the secondary node based on the primary instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossRegionZone(request *DescribeCrossRegionZoneRequest) (response *DescribeCrossRegionZoneResponse, err error) {
    return c.DescribeCrossRegionZoneWithContext(context.Background(), request)
}

// DescribeCrossRegionZone
// This API is used to query the disaster recovery region and AZ of the secondary node based on the primary instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossRegionZoneWithContext(ctx context.Context, request *DescribeCrossRegionZoneRequest) (response *DescribeCrossRegionZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCrossRegionZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeCrossRegionZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossRegionZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossRegionZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossRegionsRequest() (request *DescribeCrossRegionsRequest) {
    request = &DescribeCrossRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCrossRegions")
    
    
    return
}

func NewDescribeCrossRegionsResponse() (response *DescribeCrossRegionsResponse) {
    response = &DescribeCrossRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossRegions
// This API is used to query the target region for cross-region backups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossRegions(request *DescribeCrossRegionsRequest) (response *DescribeCrossRegionsResponse, err error) {
    return c.DescribeCrossRegionsWithContext(context.Background(), request)
}

// DescribeCrossRegions
// This API is used to query the target region for cross-region backups.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossRegionsWithContext(ctx context.Context, request *DescribeCrossRegionsRequest) (response *DescribeCrossRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeCrossRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeCrossRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossRegionsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBCharsets")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBInstanceInter")
    
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
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
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
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBInstancesAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstancesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBPrivilegeByAccountRequest() (request *DescribeDBPrivilegeByAccountRequest) {
    request = &DescribeDBPrivilegeByAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBPrivilegeByAccount")
    
    
    return
}

func NewDescribeDBPrivilegeByAccountResponse() (response *DescribeDBPrivilegeByAccountResponse) {
    response = &DescribeDBPrivilegeByAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBPrivilegeByAccount
// This API is used to query information on the databases and permissions associated with the account.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
func (c *Client) DescribeDBPrivilegeByAccount(request *DescribeDBPrivilegeByAccountRequest) (response *DescribeDBPrivilegeByAccountResponse, err error) {
    return c.DescribeDBPrivilegeByAccountWithContext(context.Background(), request)
}

// DescribeDBPrivilegeByAccount
// This API is used to query information on the databases and permissions associated with the account.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
func (c *Client) DescribeDBPrivilegeByAccountWithContext(ctx context.Context, request *DescribeDBPrivilegeByAccountRequest) (response *DescribeDBPrivilegeByAccountResponse, err error) {
    if request == nil {
        request = NewDescribeDBPrivilegeByAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBPrivilegeByAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBPrivilegeByAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBPrivilegeByAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBRestoreTimeRequest() (request *DescribeDBRestoreTimeRequest) {
    request = &DescribeDBRestoreTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBRestoreTime")
    
    
    return
}

func NewDescribeDBRestoreTimeResponse() (response *DescribeDBRestoreTimeResponse) {
    response = &DescribeDBRestoreTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBRestoreTime
// This API is used to query databases that can be rolled back.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBRestoreTime(request *DescribeDBRestoreTimeRequest) (response *DescribeDBRestoreTimeResponse, err error) {
    return c.DescribeDBRestoreTimeWithContext(context.Background(), request)
}

// DescribeDBRestoreTime
// This API is used to query databases that can be rolled back.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBRestoreTimeWithContext(ctx context.Context, request *DescribeDBRestoreTimeRequest) (response *DescribeDBRestoreTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBRestoreTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBRestoreTime")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBRestoreTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBRestoreTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSecurityGroups
// This API is used to query the security group details of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// This API is used to query the security group details of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSecurityGroupsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBs")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDBsNormal")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBsNormal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBsNormalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseNamesRequest() (request *DescribeDatabaseNamesRequest) {
    request = &DescribeDatabaseNamesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDatabaseNames")
    
    
    return
}

func NewDescribeDatabaseNamesResponse() (response *DescribeDatabaseNamesResponse) {
    response = &DescribeDatabaseNamesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseNames
// This API is used to query the database name associated with the account.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseNames(request *DescribeDatabaseNamesRequest) (response *DescribeDatabaseNamesResponse, err error) {
    return c.DescribeDatabaseNamesWithContext(context.Background(), request)
}

// DescribeDatabaseNames
// This API is used to query the database name associated with the account.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseNamesWithContext(ctx context.Context, request *DescribeDatabaseNamesRequest) (response *DescribeDatabaseNamesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseNamesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDatabaseNames")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseNames require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseNamesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabases
// This API is used to query the database list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// This API is used to query the database list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesNormalRequest() (request *DescribeDatabasesNormalRequest) {
    request = &DescribeDatabasesNormalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDatabasesNormal")
    
    
    return
}

func NewDescribeDatabasesNormalResponse() (response *DescribeDatabasesNormalResponse) {
    response = &DescribeDatabasesNormalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabasesNormal
// This API is used to query database configuration information. This API does not contain accounts associated with databases.
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
func (c *Client) DescribeDatabasesNormal(request *DescribeDatabasesNormalRequest) (response *DescribeDatabasesNormalResponse, err error) {
    return c.DescribeDatabasesNormalWithContext(context.Background(), request)
}

// DescribeDatabasesNormal
// This API is used to query database configuration information. This API does not contain accounts associated with databases.
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
func (c *Client) DescribeDatabasesNormalWithContext(ctx context.Context, request *DescribeDatabasesNormalRequest) (response *DescribeDatabasesNormalResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesNormalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeDatabasesNormal")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabasesNormal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabasesNormalResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeFlowStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHASwitchLogRequest() (request *DescribeHASwitchLogRequest) {
    request = &DescribeHASwitchLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeHASwitchLog")
    
    
    return
}

func NewDescribeHASwitchLogResponse() (response *DescribeHASwitchLogResponse) {
    response = &DescribeHASwitchLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHASwitchLog
// This API is used to perform the manual primary-secondary switch.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeHASwitchLog(request *DescribeHASwitchLogRequest) (response *DescribeHASwitchLogResponse, err error) {
    return c.DescribeHASwitchLogWithContext(context.Background(), request)
}

// DescribeHASwitchLog
// This API is used to perform the manual primary-secondary switch.
//
// error code that may be returned:
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeHASwitchLogWithContext(ctx context.Context, request *DescribeHASwitchLogRequest) (response *DescribeHASwitchLogResponse, err error) {
    if request == nil {
        request = NewDescribeHASwitchLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeHASwitchLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHASwitchLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHASwitchLogResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeIncrementalMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeInstanceByOrders")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeInstanceParamRecords")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTasksRequest() (request *DescribeInstanceTasksRequest) {
    request = &DescribeInstanceTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceTasks")
    
    
    return
}

func NewDescribeInstanceTasksResponse() (response *DescribeInstanceTasksResponse) {
    response = &DescribeInstanceTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceTasks
// This API is used to query the list of asynchronous tasks related to an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceTasks(request *DescribeInstanceTasksRequest) (response *DescribeInstanceTasksResponse, err error) {
    return c.DescribeInstanceTasksWithContext(context.Background(), request)
}

// DescribeInstanceTasks
// This API is used to query the list of asynchronous tasks related to an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceTasksWithContext(ctx context.Context, request *DescribeInstanceTasksRequest) (response *DescribeInstanceTasksResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeInstanceTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTradeParameterRequest() (request *DescribeInstanceTradeParameterRequest) {
    request = &DescribeInstanceTradeParameterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceTradeParameter")
    
    
    return
}

func NewDescribeInstanceTradeParameterResponse() (response *DescribeInstanceTradeParameterResponse) {
    response = &DescribeInstanceTradeParameterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceTradeParameter
// This API is used to query the instance billing parameters.
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
func (c *Client) DescribeInstanceTradeParameter(request *DescribeInstanceTradeParameterRequest) (response *DescribeInstanceTradeParameterResponse, err error) {
    return c.DescribeInstanceTradeParameterWithContext(context.Background(), request)
}

// DescribeInstanceTradeParameter
// This API is used to query the instance billing parameters.
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
func (c *Client) DescribeInstanceTradeParameterWithContext(ctx context.Context, request *DescribeInstanceTradeParameterRequest) (response *DescribeInstanceTradeParameterResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTradeParameterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeInstanceTradeParameter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceTradeParameter require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceTradeParameterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintenanceSpanRequest() (request *DescribeMaintenanceSpanRequest) {
    request = &DescribeMaintenanceSpanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMaintenanceSpan")
    
    
    return
}

func NewDescribeMaintenanceSpanResponse() (response *DescribeMaintenanceSpanResponse) {
    response = &DescribeMaintenanceSpanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMaintenanceSpan
// This API is used to query the maintenance time window of an instance based on its instance ID.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintenanceSpan(request *DescribeMaintenanceSpanRequest) (response *DescribeMaintenanceSpanResponse, err error) {
    return c.DescribeMaintenanceSpanWithContext(context.Background(), request)
}

// DescribeMaintenanceSpan
// This API is used to query the maintenance time window of an instance based on its instance ID.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintenanceSpanWithContext(ctx context.Context, request *DescribeMaintenanceSpanRequest) (response *DescribeMaintenanceSpanResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceSpanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeMaintenanceSpan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintenanceSpan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintenanceSpanResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeMigrationDetail")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeMigrations")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeOrders")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeProductConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeProjectSecurityGroups")
    
    
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectSecurityGroups
// This API is used to query security group details of a project.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// This API is used to query security group details of a project.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroupsWithContext(ctx context.Context, request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeProjectSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublishSubscribeRequest() (request *DescribePublishSubscribeRequest) {
    request = &DescribePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribePublishSubscribe")
    
    
    return
}

func NewDescribePublishSubscribeResponse() (response *DescribePublishSubscribeResponse) {
    response = &DescribePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublishSubscribe
// This API is used to query the publish/subscribe relationship list.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribePublishSubscribe(request *DescribePublishSubscribeRequest) (response *DescribePublishSubscribeResponse, err error) {
    return c.DescribePublishSubscribeWithContext(context.Background(), request)
}

// DescribePublishSubscribe
// This API is used to query the publish/subscribe relationship list.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribePublishSubscribeWithContext(ctx context.Context, request *DescribePublishSubscribeRequest) (response *DescribePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewDescribePublishSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribePublishSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupAutoWeightRequest() (request *DescribeReadOnlyGroupAutoWeightRequest) {
    request = &DescribeReadOnlyGroupAutoWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupAutoWeight")
    
    
    return
}

func NewDescribeReadOnlyGroupAutoWeightResponse() (response *DescribeReadOnlyGroupAutoWeightResponse) {
    response = &DescribeReadOnlyGroupAutoWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReadOnlyGroupAutoWeight
// This API is used to query the automatic weight assignment result of the read-only group. The BalanceReadOnlyGroup API is used to perform routing weight assignment according to the automatic weight assignment result.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupAutoWeight(request *DescribeReadOnlyGroupAutoWeightRequest) (response *DescribeReadOnlyGroupAutoWeightResponse, err error) {
    return c.DescribeReadOnlyGroupAutoWeightWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupAutoWeight
// This API is used to query the automatic weight assignment result of the read-only group. The BalanceReadOnlyGroup API is used to perform routing weight assignment according to the automatic weight assignment result.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupAutoWeightWithContext(ctx context.Context, request *DescribeReadOnlyGroupAutoWeightRequest) (response *DescribeReadOnlyGroupAutoWeightResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupAutoWeightRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeReadOnlyGroupAutoWeight")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupAutoWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupAutoWeightResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupByReadOnlyInstanceRequest() (request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) {
    request = &DescribeReadOnlyGroupByReadOnlyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupByReadOnlyInstance")
    
    
    return
}

func NewDescribeReadOnlyGroupByReadOnlyInstanceResponse() (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse) {
    response = &DescribeReadOnlyGroupByReadOnlyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReadOnlyGroupByReadOnlyInstance
// This API is used to query the read-only group where the read-only instance is located by its ID.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupByReadOnlyInstance(request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse, err error) {
    return c.DescribeReadOnlyGroupByReadOnlyInstanceWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupByReadOnlyInstance
// This API is used to query the read-only group where the read-only instance is located by its ID.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupByReadOnlyInstanceWithContext(ctx context.Context, request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupByReadOnlyInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeReadOnlyGroupByReadOnlyInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupByReadOnlyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupByReadOnlyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupDetailsRequest() (request *DescribeReadOnlyGroupDetailsRequest) {
    request = &DescribeReadOnlyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupDetails")
    
    
    return
}

func NewDescribeReadOnlyGroupDetailsResponse() (response *DescribeReadOnlyGroupDetailsResponse) {
    response = &DescribeReadOnlyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReadOnlyGroupDetails
// This API is used to query read-only group details.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupDetails(request *DescribeReadOnlyGroupDetailsRequest) (response *DescribeReadOnlyGroupDetailsResponse, err error) {
    return c.DescribeReadOnlyGroupDetailsWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupDetails
// This API is used to query read-only group details.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupDetailsWithContext(ctx context.Context, request *DescribeReadOnlyGroupDetailsRequest) (response *DescribeReadOnlyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeReadOnlyGroupDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupListRequest() (request *DescribeReadOnlyGroupListRequest) {
    request = &DescribeReadOnlyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupList")
    
    
    return
}

func NewDescribeReadOnlyGroupListResponse() (response *DescribeReadOnlyGroupListResponse) {
    response = &DescribeReadOnlyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReadOnlyGroupList
// This API is used to query the read-only group list.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupList(request *DescribeReadOnlyGroupListRequest) (response *DescribeReadOnlyGroupListResponse, err error) {
    return c.DescribeReadOnlyGroupListWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupList
// This API is used to query the read-only group list.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupListWithContext(ctx context.Context, request *DescribeReadOnlyGroupListRequest) (response *DescribeReadOnlyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeReadOnlyGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegularBackupPlanRequest() (request *DescribeRegularBackupPlanRequest) {
    request = &DescribeRegularBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRegularBackupPlan")
    
    
    return
}

func NewDescribeRegularBackupPlanResponse() (response *DescribeRegularBackupPlanResponse) {
    response = &DescribeRegularBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegularBackupPlan
// This API is used to query the scheduled backup retention plans of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
func (c *Client) DescribeRegularBackupPlan(request *DescribeRegularBackupPlanRequest) (response *DescribeRegularBackupPlanResponse, err error) {
    return c.DescribeRegularBackupPlanWithContext(context.Background(), request)
}

// DescribeRegularBackupPlan
// This API is used to query the scheduled backup retention plans of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
func (c *Client) DescribeRegularBackupPlanWithContext(ctx context.Context, request *DescribeRegularBackupPlanRequest) (response *DescribeRegularBackupPlanResponse, err error) {
    if request == nil {
        request = NewDescribeRegularBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeRegularBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegularBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegularBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRestoreTaskRequest() (request *DescribeRestoreTaskRequest) {
    request = &DescribeRestoreTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRestoreTask")
    
    
    return
}

func NewDescribeRestoreTaskResponse() (response *DescribeRestoreTaskResponse) {
    response = &DescribeRestoreTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRestoreTask
// This API is used to query the list of rollback tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRestoreTask(request *DescribeRestoreTaskRequest) (response *DescribeRestoreTaskResponse, err error) {
    return c.DescribeRestoreTaskWithContext(context.Background(), request)
}

// DescribeRestoreTask
// This API is used to query the list of rollback tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRestoreTaskWithContext(ctx context.Context, request *DescribeRestoreTaskRequest) (response *DescribeRestoreTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRestoreTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeRestoreTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRestoreTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRestoreTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRestoreTimeRangeRequest() (request *DescribeRestoreTimeRangeRequest) {
    request = &DescribeRestoreTimeRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRestoreTimeRange")
    
    
    return
}

func NewDescribeRestoreTimeRangeResponse() (response *DescribeRestoreTimeRangeResponse) {
    response = &DescribeRestoreTimeRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRestoreTimeRange
// This API is used to query the time range available for rollback by time point.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRestoreTimeRange(request *DescribeRestoreTimeRangeRequest) (response *DescribeRestoreTimeRangeResponse, err error) {
    return c.DescribeRestoreTimeRangeWithContext(context.Background(), request)
}

// DescribeRestoreTimeRange
// This API is used to query the time range available for rollback by time point.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRestoreTimeRangeWithContext(ctx context.Context, request *DescribeRestoreTimeRangeRequest) (response *DescribeRestoreTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeRestoreTimeRangeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeRestoreTimeRange")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRestoreTimeRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRestoreTimeRangeResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeRollbackTime")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeSlowlogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecSellStatusRequest() (request *DescribeSpecSellStatusRequest) {
    request = &DescribeSpecSellStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeSpecSellStatus")
    
    
    return
}

func NewDescribeSpecSellStatusResponse() (response *DescribeSpecSellStatusResponse) {
    response = &DescribeSpecSellStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpecSellStatus
// This API is used to query the status information on specifications, including the sales status and reference price. (The actual price is subject to the result returned by price querying APIs.)
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSpecSellStatus(request *DescribeSpecSellStatusRequest) (response *DescribeSpecSellStatusResponse, err error) {
    return c.DescribeSpecSellStatusWithContext(context.Background(), request)
}

// DescribeSpecSellStatus
// This API is used to query the status information on specifications, including the sales status and reference price. (The actual price is subject to the result returned by price querying APIs.)
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSpecSellStatusWithContext(ctx context.Context, request *DescribeSpecSellStatusRequest) (response *DescribeSpecSellStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSpecSellStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeSpecSellStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpecSellStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpecSellStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpgradeInstanceCheckRequest() (request *DescribeUpgradeInstanceCheckRequest) {
    request = &DescribeUpgradeInstanceCheckRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeUpgradeInstanceCheck")
    
    
    return
}

func NewDescribeUpgradeInstanceCheckResponse() (response *DescribeUpgradeInstanceCheckResponse) {
    response = &DescribeUpgradeInstanceCheckResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpgradeInstanceCheck
// This API is used to pre-check the impact of the instance configuration adjustment before the adjustment.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_NOTSUPPORTPUBSUBINSTANCE = "InvalidParameter.NotSupportPubSubInstance"
//  INVALIDPARAMETER_NOTSUPPORTREADONLYMASTERINSTANCE = "InvalidParameter.NotSupportReadOnlyMasterInstance"
//  INVALIDPARAMETER_NOTSUPPORTSINGLEINSTANCE = "InvalidParameter.NotSupportSingleInstance"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeUpgradeInstanceCheck(request *DescribeUpgradeInstanceCheckRequest) (response *DescribeUpgradeInstanceCheckResponse, err error) {
    return c.DescribeUpgradeInstanceCheckWithContext(context.Background(), request)
}

// DescribeUpgradeInstanceCheck
// This API is used to pre-check the impact of the instance configuration adjustment before the adjustment.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_NOTSUPPORTPUBSUBINSTANCE = "InvalidParameter.NotSupportPubSubInstance"
//  INVALIDPARAMETER_NOTSUPPORTREADONLYMASTERINSTANCE = "InvalidParameter.NotSupportReadOnlyMasterInstance"
//  INVALIDPARAMETER_NOTSUPPORTSINGLEINSTANCE = "InvalidParameter.NotSupportSingleInstance"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeUpgradeInstanceCheckWithContext(ctx context.Context, request *DescribeUpgradeInstanceCheckRequest) (response *DescribeUpgradeInstanceCheckResponse, err error) {
    if request == nil {
        request = NewDescribeUpgradeInstanceCheckRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeUpgradeInstanceCheck")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpgradeInstanceCheck require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpgradeInstanceCheckResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeUploadBackupInfo")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeXEvents")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateSecurityGroups
// This API is used to unbind security groups from instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// This API is used to unbind security groups from instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "DisassociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "InquiryPriceCreateDBInstances")
    
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
// This API is used to query the upgrade prices of a monthly subscribed instance.
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
// This API is used to query the upgrade prices of a monthly subscribed instance.
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "InquiryPriceUpgradeDBInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyAccountPrivilege")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyAccountRemark")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyBackupMigration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupNameRequest() (request *ModifyBackupNameRequest) {
    request = &ModifyBackupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupName")
    
    
    return
}

func NewModifyBackupNameResponse() (response *ModifyBackupNameResponse) {
    response = &ModifyBackupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupName
// This API is used to modify the name of a backup task.
//
// error code that may be returned:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupName(request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    return c.ModifyBackupNameWithContext(context.Background(), request)
}

// ModifyBackupName
// This API is used to modify the name of a backup task.
//
// error code that may be returned:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupNameWithContext(ctx context.Context, request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    if request == nil {
        request = NewModifyBackupNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyBackupName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupNameResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyBackupStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloseWanIpRequest() (request *ModifyCloseWanIpRequest) {
    request = &ModifyCloseWanIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyCloseWanIp")
    
    
    return
}

func NewModifyCloseWanIpResponse() (response *ModifyCloseWanIpResponse) {
    response = &ModifyCloseWanIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloseWanIp
// This API is used to disable the public network for the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_DOMAINOPERATIONFAILED = "FailedOperation.DomainOperationFailed"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyCloseWanIp(request *ModifyCloseWanIpRequest) (response *ModifyCloseWanIpResponse, err error) {
    return c.ModifyCloseWanIpWithContext(context.Background(), request)
}

// ModifyCloseWanIp
// This API is used to disable the public network for the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_DOMAINOPERATIONFAILED = "FailedOperation.DomainOperationFailed"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyCloseWanIpWithContext(ctx context.Context, request *ModifyCloseWanIpRequest) (response *ModifyCloseWanIpResponse, err error) {
    if request == nil {
        request = NewModifyCloseWanIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyCloseWanIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloseWanIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloseWanIpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCrossBackupStrategyRequest() (request *ModifyCrossBackupStrategyRequest) {
    request = &ModifyCrossBackupStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyCrossBackupStrategy")
    
    
    return
}

func NewModifyCrossBackupStrategyResponse() (response *ModifyCrossBackupStrategyResponse) {
    response = &ModifyCrossBackupStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCrossBackupStrategy
// This API is used to enable or disable cross-region backup policies.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyCrossBackupStrategy(request *ModifyCrossBackupStrategyRequest) (response *ModifyCrossBackupStrategyResponse, err error) {
    return c.ModifyCrossBackupStrategyWithContext(context.Background(), request)
}

// ModifyCrossBackupStrategy
// This API is used to enable or disable cross-region backup policies.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyCrossBackupStrategyWithContext(ctx context.Context, request *ModifyCrossBackupStrategyRequest) (response *ModifyCrossBackupStrategyResponse, err error) {
    if request == nil {
        request = NewModifyCrossBackupStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyCrossBackupStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCrossBackupStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCrossBackupStrategyResponse()
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
// This API is used to enable or disable TDE of a database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBEncryptAttributes(request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    return c.ModifyDBEncryptAttributesWithContext(context.Background(), request)
}

// ModifyDBEncryptAttributes
// This API is used to enable or disable TDE of a database.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBEncryptAttributesWithContext(ctx context.Context, request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBEncryptAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBEncryptAttributes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBInstanceName")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBInstanceNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNoteRequest() (request *ModifyDBInstanceNoteRequest) {
    request = &ModifyDBInstanceNoteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceNote")
    
    
    return
}

func NewModifyDBInstanceNoteResponse() (response *ModifyDBInstanceNoteResponse) {
    response = &ModifyDBInstanceNoteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceNote
// This API is used to modify the instance remarks.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_INSTANCENAMEISILLEGAL = "InvalidParameterValue.InstanceNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceNote(request *ModifyDBInstanceNoteRequest) (response *ModifyDBInstanceNoteResponse, err error) {
    return c.ModifyDBInstanceNoteWithContext(context.Background(), request)
}

// ModifyDBInstanceNote
// This API is used to modify the instance remarks.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_INSTANCENAMEISILLEGAL = "InvalidParameterValue.InstanceNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceNoteWithContext(ctx context.Context, request *ModifyDBInstanceNoteRequest) (response *ModifyDBInstanceNoteResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNoteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBInstanceNote")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNote require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNoteResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBInstanceProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSSLRequest() (request *ModifyDBInstanceSSLRequest) {
    request = &ModifyDBInstanceSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceSSL")
    
    
    return
}

func NewModifyDBInstanceSSLResponse() (response *ModifyDBInstanceSSLResponse) {
    response = &ModifyDBInstanceSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSSL
// This API is used to enable/disable/update SSL encryption.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (response *ModifyDBInstanceSSLResponse, err error) {
    return c.ModifyDBInstanceSSLWithContext(context.Background(), request)
}

// ModifyDBInstanceSSL
// This API is used to enable/disable/update SSL encryption.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBInstanceSSLWithContext(ctx context.Context, request *ModifyDBInstanceSSLRequest) (response *ModifyDBInstanceSSLResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBInstanceSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSSLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSecurityGroups
// This API is used to modify security groups bound to an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// This API is used to modify security groups bound to an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBName")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDBRemark")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDReadableRequest() (request *ModifyDReadableRequest) {
    request = &ModifyDReadableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDReadable")
    
    
    return
}

func NewModifyDReadableResponse() (response *ModifyDReadableResponse) {
    response = &ModifyDReadableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDReadable
// This API is used to enable or disable the read-only feature of the replica server.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
func (c *Client) ModifyDReadable(request *ModifyDReadableRequest) (response *ModifyDReadableResponse, err error) {
    return c.ModifyDReadableWithContext(context.Background(), request)
}

// ModifyDReadable
// This API is used to enable or disable the read-only feature of the replica server.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
func (c *Client) ModifyDReadableWithContext(ctx context.Context, request *ModifyDReadableRequest) (response *ModifyDReadableResponse, err error) {
    if request == nil {
        request = NewModifyDReadableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDReadable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDReadable require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDReadableResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDatabaseCDC")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDatabaseCT")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDatabaseMdf")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseMdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseMdfResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabasePrivilegeRequest() (request *ModifyDatabasePrivilegeRequest) {
    request = &ModifyDatabasePrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabasePrivilege")
    
    
    return
}

func NewModifyDatabasePrivilegeResponse() (response *ModifyDatabasePrivilegeResponse) {
    response = &ModifyDatabasePrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatabasePrivilege
// This API is used to modify instance database permissions.
//
// error code that may be returned:
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabasePrivilege(request *ModifyDatabasePrivilegeRequest) (response *ModifyDatabasePrivilegeResponse, err error) {
    return c.ModifyDatabasePrivilegeWithContext(context.Background(), request)
}

// ModifyDatabasePrivilege
// This API is used to modify instance database permissions.
//
// error code that may be returned:
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabasePrivilegeWithContext(ctx context.Context, request *ModifyDatabasePrivilegeRequest) (response *ModifyDatabasePrivilegeResponse, err error) {
    if request == nil {
        request = NewModifyDatabasePrivilegeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDatabasePrivilege")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabasePrivilege require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabasePrivilegeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseShrinkMDFRequest() (request *ModifyDatabaseShrinkMDFRequest) {
    request = &ModifyDatabaseShrinkMDFRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseShrinkMDF")
    
    
    return
}

func NewModifyDatabaseShrinkMDFResponse() (response *ModifyDatabaseShrinkMDFResponse) {
    response = &ModifyDatabaseShrinkMDFResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatabaseShrinkMDF
// This API is used to shrink the database mdf (Shrink mdf).
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
func (c *Client) ModifyDatabaseShrinkMDF(request *ModifyDatabaseShrinkMDFRequest) (response *ModifyDatabaseShrinkMDFResponse, err error) {
    return c.ModifyDatabaseShrinkMDFWithContext(context.Background(), request)
}

// ModifyDatabaseShrinkMDF
// This API is used to shrink the database mdf (Shrink mdf).
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
func (c *Client) ModifyDatabaseShrinkMDFWithContext(ctx context.Context, request *ModifyDatabaseShrinkMDFRequest) (response *ModifyDatabaseShrinkMDFResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseShrinkMDFRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyDatabaseShrinkMDF")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseShrinkMDF require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseShrinkMDFResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyIncrementalMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyInstanceEncryptAttributes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyInstanceParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceSpanRequest() (request *ModifyMaintenanceSpanRequest) {
    request = &ModifyMaintenanceSpanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyMaintenanceSpan")
    
    
    return
}

func NewModifyMaintenanceSpanResponse() (response *ModifyMaintenanceSpanResponse) {
    response = &ModifyMaintenanceSpanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMaintenanceSpan
// This API is used to modify the maintenance window of the instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintenanceSpan(request *ModifyMaintenanceSpanRequest) (response *ModifyMaintenanceSpanResponse, err error) {
    return c.ModifyMaintenanceSpanWithContext(context.Background(), request)
}

// ModifyMaintenanceSpan
// This API is used to modify the maintenance window of the instance.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintenanceSpanWithContext(ctx context.Context, request *ModifyMaintenanceSpanRequest) (response *ModifyMaintenanceSpanResponse, err error) {
    if request == nil {
        request = NewModifyMaintenanceSpanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyMaintenanceSpan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintenanceSpan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintenanceSpanResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyMigration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOpenWanIpRequest() (request *ModifyOpenWanIpRequest) {
    request = &ModifyOpenWanIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyOpenWanIp")
    
    
    return
}

func NewModifyOpenWanIpResponse() (response *ModifyOpenWanIpResponse) {
    response = &ModifyOpenWanIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOpenWanIp
// This API is used to enable the public network for the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_DOMAINOPERATIONFAILED = "FailedOperation.DomainOperationFailed"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyOpenWanIp(request *ModifyOpenWanIpRequest) (response *ModifyOpenWanIpResponse, err error) {
    return c.ModifyOpenWanIpWithContext(context.Background(), request)
}

// ModifyOpenWanIp
// This API is used to enable the public network for the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_DOMAINOPERATIONFAILED = "FailedOperation.DomainOperationFailed"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyOpenWanIpWithContext(ctx context.Context, request *ModifyOpenWanIpRequest) (response *ModifyOpenWanIpResponse, err error) {
    if request == nil {
        request = NewModifyOpenWanIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyOpenWanIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOpenWanIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOpenWanIpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublishSubscribeRequest() (request *ModifyPublishSubscribeRequest) {
    request = &ModifyPublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyPublishSubscribe")
    
    
    return
}

func NewModifyPublishSubscribeResponse() (response *ModifyPublishSubscribeResponse) {
    response = &ModifyPublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPublishSubscribe
// This API is used to modify the publish/subscribe relationship of the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
func (c *Client) ModifyPublishSubscribe(request *ModifyPublishSubscribeRequest) (response *ModifyPublishSubscribeResponse, err error) {
    return c.ModifyPublishSubscribeWithContext(context.Background(), request)
}

// ModifyPublishSubscribe
// This API is used to modify the publish/subscribe relationship of the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
func (c *Client) ModifyPublishSubscribeWithContext(ctx context.Context, request *ModifyPublishSubscribeRequest) (response *ModifyPublishSubscribeResponse, err error) {
    if request == nil {
        request = NewModifyPublishSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyPublishSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublishSubscribeNameRequest() (request *ModifyPublishSubscribeNameRequest) {
    request = &ModifyPublishSubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyPublishSubscribeName")
    
    
    return
}

func NewModifyPublishSubscribeNameResponse() (response *ModifyPublishSubscribeNameResponse) {
    response = &ModifyPublishSubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPublishSubscribeName
// This API is used to modify the publish/subscribe names.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyPublishSubscribeName(request *ModifyPublishSubscribeNameRequest) (response *ModifyPublishSubscribeNameResponse, err error) {
    return c.ModifyPublishSubscribeNameWithContext(context.Background(), request)
}

// ModifyPublishSubscribeName
// This API is used to modify the publish/subscribe names.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyPublishSubscribeNameWithContext(ctx context.Context, request *ModifyPublishSubscribeNameRequest) (response *ModifyPublishSubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifyPublishSubscribeNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyPublishSubscribeName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublishSubscribeName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublishSubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReadOnlyGroupDetailsRequest() (request *ModifyReadOnlyGroupDetailsRequest) {
    request = &ModifyReadOnlyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyReadOnlyGroupDetails")
    
    
    return
}

func NewModifyReadOnlyGroupDetailsResponse() (response *ModifyReadOnlyGroupDetailsResponse) {
    response = &ModifyReadOnlyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReadOnlyGroupDetails
// This API is used to modify read-only group details.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyReadOnlyGroupDetails(request *ModifyReadOnlyGroupDetailsRequest) (response *ModifyReadOnlyGroupDetailsResponse, err error) {
    return c.ModifyReadOnlyGroupDetailsWithContext(context.Background(), request)
}

// ModifyReadOnlyGroupDetails
// This API is used to modify read-only group details.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyReadOnlyGroupDetailsWithContext(ctx context.Context, request *ModifyReadOnlyGroupDetailsRequest) (response *ModifyReadOnlyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewModifyReadOnlyGroupDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ModifyReadOnlyGroupDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReadOnlyGroupDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReadOnlyGroupDetailsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "OpenInterCommunication")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RecycleDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecycleDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecycleDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleReadOnlyGroupRequest() (request *RecycleReadOnlyGroupRequest) {
    request = &RecycleReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RecycleReadOnlyGroup")
    
    
    return
}

func NewRecycleReadOnlyGroupResponse() (response *RecycleReadOnlyGroupResponse) {
    response = &RecycleReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecycleReadOnlyGroup
// This API is used to reclaim resources of read-only groups immediately. The resources, such as VIP, occupied by the read-only group will be released immediately and cannot be recovered.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ROGROUPSTATUSISILLEGAL = "InvalidParameterValue.RoGroupStatusIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecycleReadOnlyGroup(request *RecycleReadOnlyGroupRequest) (response *RecycleReadOnlyGroupResponse, err error) {
    return c.RecycleReadOnlyGroupWithContext(context.Background(), request)
}

// RecycleReadOnlyGroup
// This API is used to reclaim resources of read-only groups immediately. The resources, such as VIP, occupied by the read-only group will be released immediately and cannot be recovered.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ROGROUPSTATUSISILLEGAL = "InvalidParameterValue.RoGroupStatusIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecycleReadOnlyGroupWithContext(ctx context.Context, request *RecycleReadOnlyGroupRequest) (response *RecycleReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRecycleReadOnlyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RecycleReadOnlyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecycleReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecycleReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveBackupsRequest() (request *RemoveBackupsRequest) {
    request = &RemoveBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RemoveBackups")
    
    
    return
}

func NewRemoveBackupsResponse() (response *RemoveBackupsResponse) {
    response = &RemoveBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveBackups
// This API is used to delete backup files created by users manually. The backup policy to be deleted can be instance backup or multi-database backup.
//
// error code that may be returned:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveBackups(request *RemoveBackupsRequest) (response *RemoveBackupsResponse, err error) {
    return c.RemoveBackupsWithContext(context.Background(), request)
}

// RemoveBackups
// This API is used to delete backup files created by users manually. The backup policy to be deleted can be instance backup or multi-database backup.
//
// error code that may be returned:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveBackupsWithContext(ctx context.Context, request *RemoveBackupsRequest) (response *RemoveBackupsResponse, err error) {
    if request == nil {
        request = NewRemoveBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RemoveBackups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewRenewPostpaidDBInstanceRequest() (request *RenewPostpaidDBInstanceRequest) {
    request = &RenewPostpaidDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RenewPostpaidDBInstance")
    
    
    return
}

func NewRenewPostpaidDBInstanceResponse() (response *RenewPostpaidDBInstanceResponse) {
    response = &RenewPostpaidDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewPostpaidDBInstance
// This API is used to recover the pay-as-you-go instance that is manually isolated through the API TerminateDBInstance from the recycle bin.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewPostpaidDBInstance(request *RenewPostpaidDBInstanceRequest) (response *RenewPostpaidDBInstanceResponse, err error) {
    return c.RenewPostpaidDBInstanceWithContext(context.Background(), request)
}

// RenewPostpaidDBInstance
// This API is used to recover the pay-as-you-go instance that is manually isolated through the API TerminateDBInstance from the recycle bin.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewPostpaidDBInstanceWithContext(ctx context.Context, request *RenewPostpaidDBInstanceRequest) (response *RenewPostpaidDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewPostpaidDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RenewPostpaidDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewPostpaidDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewPostpaidDBInstanceResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "ResetAccountPassword")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RestartDBInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RestoreInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RollbackInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "RunMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "StartBackupMigration")
    
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "StartIncrementalMigration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "StartInstanceXEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstanceXEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartInstanceXEventResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchCloudInstanceHARequest() (request *SwitchCloudInstanceHARequest) {
    request = &SwitchCloudInstanceHARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "SwitchCloudInstanceHA")
    
    
    return
}

func NewSwitchCloudInstanceHAResponse() (response *SwitchCloudInstanceHAResponse) {
    response = &SwitchCloudInstanceHAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchCloudInstanceHA
// This API is used to manually switch between primary and secondary.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchCloudInstanceHA(request *SwitchCloudInstanceHARequest) (response *SwitchCloudInstanceHAResponse, err error) {
    return c.SwitchCloudInstanceHAWithContext(context.Background(), request)
}

// SwitchCloudInstanceHA
// This API is used to manually switch between primary and secondary.
//
// error code that may be returned:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_NOTSUPPORT = "FailedOperation.NotSupport"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchCloudInstanceHAWithContext(ctx context.Context, request *SwitchCloudInstanceHARequest) (response *SwitchCloudInstanceHAResponse, err error) {
    if request == nil {
        request = NewSwitchCloudInstanceHARequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "SwitchCloudInstanceHA")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchCloudInstanceHA require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchCloudInstanceHAResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "TerminateDBInstance")
    
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
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
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDBInstanceWithContext(ctx context.Context, request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "sqlserver", APIVersion, "UpgradeDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
