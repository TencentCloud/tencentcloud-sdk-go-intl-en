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

package v20180411

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-11"

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
    request.Init().WithApiInfo("dcdb", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// This API is used to associate security groups with Tencent Cloud resources in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_DBROWSAFFECTEDERROR = "InternalError.DBRowsAffectedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSGQUOTAERROR = "InternalError.GetUsgQuotaError"
//  INTERNALERROR_INSERTFAIL = "InternalError.InsertFail"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INTERNALERROR_UPDATEDATABASEFAILED = "InternalError.UpdateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_SGCHECKFAIL = "InvalidParameter.SGCheckFail"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// AssociateSecurityGroups
// This API is used to associate security groups with Tencent Cloud resources in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_DBROWSAFFECTEDERROR = "InternalError.DBRowsAffectedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSGQUOTAERROR = "InternalError.GetUsgQuotaError"
//  INTERNALERROR_INSERTFAIL = "InternalError.InsertFail"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INTERNALERROR_UPDATEDATABASEFAILED = "InternalError.UpdateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_SGCHECKFAIL = "InvalidParameter.SGCheckFail"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCancelDcnJobRequest() (request *CancelDcnJobRequest) {
    request = &CancelDcnJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CancelDcnJob")
    
    
    return
}

func NewCancelDcnJobResponse() (response *CancelDcnJobResponse) {
    response = &CancelDcnJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelDcnJob
// This API is used to cancel DCN synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CancelDcnJob(request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
    if request == nil {
        request = NewCancelDcnJobRequest()
    }
    
    response = NewCancelDcnJobResponse()
    err = c.Send(request, response)
    return
}

// CancelDcnJob
// This API is used to cancel DCN synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CancelDcnJobWithContext(ctx context.Context, request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
    if request == nil {
        request = NewCancelDcnJobRequest()
    }
    request.SetContext(ctx)
    
    response = NewCancelDcnJobResponse()
    err = c.Send(request, response)
    return
}

func NewCloneAccountRequest() (request *CloneAccountRequest) {
    request = &CloneAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CloneAccount")
    
    
    return
}

func NewCloneAccountResponse() (response *CloneAccountResponse) {
    response = &CloneAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloneAccount
// This API is used to clone an instance account.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) CloneAccount(request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    if request == nil {
        request = NewCloneAccountRequest()
    }
    
    response = NewCloneAccountResponse()
    err = c.Send(request, response)
    return
}

// CloneAccount
// This API is used to clone an instance account.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) CloneAccountWithContext(ctx context.Context, request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    if request == nil {
        request = NewCloneAccountRequest()
    }
    request.SetContext(ctx)
    
    response = NewCloneAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseDBExtranetAccessRequest() (request *CloseDBExtranetAccessRequest) {
    request = &CloseDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CloseDBExtranetAccess")
    
    
    return
}

func NewCloseDBExtranetAccessResponse() (response *CloseDBExtranetAccessResponse) {
    response = &CloseDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseDBExtranetAccess
// This API is used to disable public network access for a TencentDB instance, which will make the public IP address inaccessible. The `DescribeDCDBInstances` API will not return the public domain name and port information of the corresponding instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

// CloseDBExtranetAccess
// This API is used to disable public network access for a TencentDB instance, which will make the public IP address inaccessible. The `DescribeDCDBInstances` API will not return the public domain name and port information of the corresponding instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseDBExtranetAccessWithContext(ctx context.Context, request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    request.SetContext(ctx)
    
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAccountPrivilegesRequest() (request *CopyAccountPrivilegesRequest) {
    request = &CopyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CopyAccountPrivileges")
    
    
    return
}

func NewCopyAccountPrivilegesResponse() (response *CopyAccountPrivilegesResponse) {
    response = &CopyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyAccountPrivileges
// This API is used to copy the permissions of a TencentDB account.
//
// Note: Accounts with the same username but different hosts are different accounts. Permissions can only be copied between accounts with the same `Readonly` attribute.
//
// error code that may be returned:
//  FAILEDOPERATION_COPYRIGHTERROR = "FailedOperation.CopyRightError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERTYPE = "InvalidParameterValue.BadUserType"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyAccountPrivileges(request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewCopyAccountPrivilegesRequest()
    }
    
    response = NewCopyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

// CopyAccountPrivileges
// This API is used to copy the permissions of a TencentDB account.
//
// Note: Accounts with the same username but different hosts are different accounts. Permissions can only be copied between accounts with the same `Readonly` attribute.
//
// error code that may be returned:
//  FAILEDOPERATION_COPYRIGHTERROR = "FailedOperation.CopyRightError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERTYPE = "InvalidParameterValue.BadUserType"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyAccountPrivilegesWithContext(ctx context.Context, request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewCopyAccountPrivilegesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCopyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// This API is used to create a TencentDB account. Multiple accounts can be created for one instance. Accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEUSERFAILED = "FailedOperation.CreateUserFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_OLDPROXYVERSION = "UnsupportedOperation.OldProxyVersion"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

// CreateAccount
// This API is used to create a TencentDB account. Multiple accounts can be created for one instance. Accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEUSERFAILED = "FailedOperation.CreateUserFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_OLDPROXYVERSION = "UnsupportedOperation.OldProxyVersion"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DeleteAccount")
    
    
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccount
// This API is used to delete a TencentDB account, which is uniquely identified by username and host.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSERFAILED = "FailedOperation.DeleteUserFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

// DeleteAccount
// This API is used to delete a TencentDB account, which is uniquely identified by username and host.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSERFAILED = "FailedOperation.DeleteUserFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccountPrivileges")
    
    
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountPrivileges
// This API is used to query the permissions of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccountPrivileges
// This API is used to query the permissions of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// This API is used to query the list of accounts of a specified TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccounts
// This API is used to query the list of accounts of a specified TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBLogFiles")
    
    
    return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
    response = &DescribeDBLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBLogFiles
// This API is used to get the list of various logs of a database, including cold backups, binlogs, errlogs, and slowlogs.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_COSCONFIGURATION = "InternalError.CosConfiguration"
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_COSAPIFAILED = "ResourceUnavailable.CosApiFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    
    response = NewDescribeDBLogFilesResponse()
    err = c.Send(request, response)
    return
}

// DescribeDBLogFiles
// This API is used to get the list of various logs of a database, including cold backups, binlogs, errlogs, and slowlogs.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_COSCONFIGURATION = "InternalError.CosConfiguration"
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_COSAPIFAILED = "ResourceUnavailable.CosApiFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBLogFilesWithContext(ctx context.Context, request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDBLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
    request = &DescribeDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBParameters")
    
    
    return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
    response = &DescribeDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBParameters
// This API is used to get the current parameter settings of a database.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBParametersRequest()
    }
    
    response = NewDescribeDBParametersResponse()
    err = c.Send(request, response)
    return
}

// DescribeDBParameters
// This API is used to get the current parameter settings of a database.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBParametersWithContext(ctx context.Context, request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBParametersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSecurityGroups
// This API is used to query the security group details of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDBSecurityGroups
// This API is used to query the security group details of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSyncModeRequest() (request *DescribeDBSyncModeRequest) {
    request = &DescribeDBSyncModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSyncMode")
    
    
    return
}

func NewDescribeDBSyncModeResponse() (response *DescribeDBSyncModeResponse) {
    response = &DescribeDBSyncModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSyncMode
// This API is used to query the sync mode of a TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSyncMode(request *DescribeDBSyncModeRequest) (response *DescribeDBSyncModeResponse, err error) {
    if request == nil {
        request = NewDescribeDBSyncModeRequest()
    }
    
    response = NewDescribeDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

// DescribeDBSyncMode
// This API is used to query the sync mode of a TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSyncModeWithContext(ctx context.Context, request *DescribeDBSyncModeRequest) (response *DescribeDBSyncModeResponse, err error) {
    if request == nil {
        request = NewDescribeDBSyncModeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstanceNodeInfoRequest() (request *DescribeDCDBInstanceNodeInfoRequest) {
    request = &DescribeDCDBInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstanceNodeInfo")
    
    
    return
}

func NewDescribeDCDBInstanceNodeInfoResponse() (response *DescribeDCDBInstanceNodeInfoResponse) {
    response = &DescribeDCDBInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBInstanceNodeInfo
// This API is used to query the information of instance nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDCDBInstanceNodeInfo(request *DescribeDCDBInstanceNodeInfoRequest) (response *DescribeDCDBInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstanceNodeInfoRequest()
    }
    
    response = NewDescribeDCDBInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDCDBInstanceNodeInfo
// This API is used to query the information of instance nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDCDBInstanceNodeInfoWithContext(ctx context.Context, request *DescribeDCDBInstanceNodeInfoRequest) (response *DescribeDCDBInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstanceNodeInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDCDBInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstancesRequest() (request *DescribeDCDBInstancesRequest) {
    request = &DescribeDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstances")
    
    
    return
}

func NewDescribeDCDBInstancesResponse() (response *DescribeDCDBInstancesResponse) {
    response = &DescribeDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBInstances
// This API is used to query the list of TencentDB instances. It supports filtering instances by project ID, instance ID, private network address, and instance name.
//
// If no filter is specified, 10 instances will be returned by default. Up to 100 instances can be returned for a single request.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBInstances(request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstancesRequest()
    }
    
    response = NewDescribeDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeDCDBInstances
// This API is used to query the list of TencentDB instances. It supports filtering instances by project ID, instance ID, private network address, and instance name.
//
// If no filter is specified, 10 instances will be returned by default. Up to 100 instances can be returned for a single request.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBInstancesWithContext(ctx context.Context, request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBShardsRequest() (request *DescribeDCDBShardsRequest) {
    request = &DescribeDCDBShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBShards")
    
    
    return
}

func NewDescribeDCDBShardsResponse() (response *DescribeDCDBShardsResponse) {
    response = &DescribeDCDBShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBShards
// This API is used to query the information of shards of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBShards(request *DescribeDCDBShardsRequest) (response *DescribeDCDBShardsResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBShardsRequest()
    }
    
    response = NewDescribeDCDBShardsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDCDBShards
// This API is used to query the information of shards of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBShardsWithContext(ctx context.Context, request *DescribeDCDBShardsRequest) (response *DescribeDCDBShardsResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBShardsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDCDBShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseObjects")
    
    
    return
}

func NewDescribeDatabaseObjectsResponse() (response *DescribeDatabaseObjectsResponse) {
    response = &DescribeDatabaseObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabaseObjects
// This API is used to query the list of database objects in a TencentDB instance, including tables, stored procedures, views, and functions.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBOBJECTFAILED = "InternalError.GetDbObjectFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    
    response = NewDescribeDatabaseObjectsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDatabaseObjects
// This API is used to query the list of database objects in a TencentDB instance, including tables, stored procedures, views, and functions.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBOBJECTFAILED = "InternalError.GetDbObjectFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseObjectsWithContext(ctx context.Context, request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDatabaseObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseTableRequest() (request *DescribeDatabaseTableRequest) {
    request = &DescribeDatabaseTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseTable")
    
    
    return
}

func NewDescribeDatabaseTableResponse() (response *DescribeDatabaseTableResponse) {
    response = &DescribeDatabaseTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabaseTable
// This API is used to query the table information of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETTABLEINFOFAILED = "InternalError.GetTableInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseTable(request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseTableRequest()
    }
    
    response = NewDescribeDatabaseTableResponse()
    err = c.Send(request, response)
    return
}

// DescribeDatabaseTable
// This API is used to query the table information of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETTABLEINFOFAILED = "InternalError.GetTableInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseTableWithContext(ctx context.Context, request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseTableRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDatabaseTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabases
// This API is used to query the list of databases of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBLISTFAILED = "InternalError.GetDbListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

// DescribeDatabases
// This API is used to query the list of databases of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBLISTFAILED = "InternalError.GetDbListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDcnDetailRequest() (request *DescribeDcnDetailRequest) {
    request = &DescribeDcnDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDcnDetail")
    
    
    return
}

func NewDescribeDcnDetailResponse() (response *DescribeDcnDetailResponse) {
    response = &DescribeDcnDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDcnDetail
// This API is used to query the disaster recovery details of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDcnDetail(request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDcnDetailRequest()
    }
    
    response = NewDescribeDcnDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeDcnDetail
// This API is used to query the disaster recovery details of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDcnDetailWithContext(ctx context.Context, request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDcnDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDcnDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeFlow")
    
    
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlow
// This API is used to query task status.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_FLOWNOTFOUND = "InvalidParameter.FlowNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

// DescribeFlow
// This API is used to query task status.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_FLOWNOTFOUND = "InvalidParameter.FlowNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlowWithContext(ctx context.Context, request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjectSecurityGroups")
    
    
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectSecurityGroups
// This API is used to query the security group details of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeProjectSecurityGroups
// This API is used to query the security group details of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
func (c *Client) DescribeProjectSecurityGroupsWithContext(ctx context.Context, request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjects
// This API is used to query the project list.
//
// error code that may be returned:
//  INTERNALERROR_LISTPROJECTFAILED = "InternalError.ListProjectFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

// DescribeProjects
// This API is used to query the project list.
//
// error code that may be returned:
//  INTERNALERROR_LISTPROJECTFAILED = "InternalError.ListProjectFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDCDBInstanceRequest() (request *DestroyDCDBInstanceRequest) {
    request = &DestroyDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DestroyDCDBInstance")
    
    
    return
}

func NewDestroyDCDBInstanceResponse() (response *DestroyDCDBInstanceResponse) {
    response = &DestroyDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyDCDBInstance
// This API is used to terminate an isolated monthly-subscribed instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
func (c *Client) DestroyDCDBInstance(request *DestroyDCDBInstanceRequest) (response *DestroyDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyDCDBInstanceRequest()
    }
    
    response = NewDestroyDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

// DestroyDCDBInstance
// This API is used to terminate an isolated monthly-subscribed instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
func (c *Client) DestroyDCDBInstanceWithContext(ctx context.Context, request *DestroyDCDBInstanceRequest) (response *DestroyDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyDCDBInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDestroyDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyHourDCDBInstanceRequest() (request *DestroyHourDCDBInstanceRequest) {
    request = &DestroyHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DestroyHourDCDBInstance")
    
    
    return
}

func NewDestroyHourDCDBInstanceResponse() (response *DestroyHourDCDBInstanceResponse) {
    response = &DestroyHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyHourDCDBInstance
// This API is used to terminate a pay-as-you-go instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
func (c *Client) DestroyHourDCDBInstance(request *DestroyHourDCDBInstanceRequest) (response *DestroyHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyHourDCDBInstanceRequest()
    }
    
    response = NewDestroyHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

// DestroyHourDCDBInstance
// This API is used to terminate a pay-as-you-go instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
func (c *Client) DestroyHourDCDBInstanceWithContext(ctx context.Context, request *DestroyHourDCDBInstanceRequest) (response *DestroyHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyHourDCDBInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDestroyHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateSecurityGroups
// This API is used to unassociate security groups from instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// DisassociateSecurityGroups
// This API is used to unassociate security groups from instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
    request = &GrantAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "GrantAccountPrivileges")
    
    
    return
}

func NewGrantAccountPrivilegesResponse() (response *GrantAccountPrivilegesResponse) {
    response = &GrantAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GrantAccountPrivileges
// This API is used to grant permissions to a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_MODIFYRIGHTFAILED = "FailedOperation.ModifyRightFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  INVALIDPARAMETERVALUE_ILLEGALRIGHTPARAM = "InvalidParameterValue.IllegalRightParam"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewGrantAccountPrivilegesRequest()
    }
    
    response = NewGrantAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

// GrantAccountPrivileges
// This API is used to grant permissions to a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_MODIFYRIGHTFAILED = "FailedOperation.ModifyRightFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  INVALIDPARAMETERVALUE_ILLEGALRIGHTPARAM = "InvalidParameterValue.IllegalRightParam"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivilegesWithContext(ctx context.Context, request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewGrantAccountPrivilegesRequest()
    }
    request.SetContext(ctx)
    
    response = NewGrantAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewInitDCDBInstancesRequest() (request *InitDCDBInstancesRequest) {
    request = &InitDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "InitDCDBInstances")
    
    
    return
}

func NewInitDCDBInstancesResponse() (response *InitDCDBInstancesResponse) {
    response = &InitDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitDCDBInstances
// This API is used to initialize instances, including setting the default character set and table name case sensitivity.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ILLEGALINITPARAM = "InvalidParameterValue.IllegalInitParam"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InitDCDBInstances(request *InitDCDBInstancesRequest) (response *InitDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDCDBInstancesRequest()
    }
    
    response = NewInitDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

// InitDCDBInstances
// This API is used to initialize instances, including setting the default character set and table name case sensitivity.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ILLEGALINITPARAM = "InvalidParameterValue.IllegalInitParam"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InitDCDBInstancesWithContext(ctx context.Context, request *InitDCDBInstancesRequest) (response *InitDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDCDBInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewInitDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAccountDescription")
    
    
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountDescription
// This API is used to modify the remarks of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

// ModifyAccountDescription
// This API is used to modify the remarks of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceSecurityGroups
// This API is used to modify the security groups associated with TencentDB.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// ModifyDBInstanceSecurityGroups
// This API is used to modify the security groups associated with TencentDB.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstancesProjectRequest() (request *ModifyDBInstancesProjectRequest) {
    request = &ModifyDBInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstancesProject")
    
    
    return
}

func NewModifyDBInstancesProjectResponse() (response *ModifyDBInstancesProjectResponse) {
    response = &ModifyDBInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstancesProject
// This API is used to modify the project to which TencentDB instances belong.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstancesProject(request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

// ModifyDBInstancesProject
// This API is used to modify the project to which TencentDB instances belong.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstancesProjectWithContext(ctx context.Context, request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
    request = &ModifyDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBParameters")
    
    
    return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
    response = &ModifyDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBParameters
// This API is used to modify database parameters.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

// ModifyDBParameters
// This API is used to modify database parameters.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBParametersWithContext(ctx context.Context, request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBSyncModeRequest() (request *ModifyDBSyncModeRequest) {
    request = &ModifyDBSyncModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBSyncMode")
    
    
    return
}

func NewModifyDBSyncModeResponse() (response *ModifyDBSyncModeResponse) {
    response = &ModifyDBSyncModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBSyncMode
// This API is used to modify the sync mode of a TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_BADSYNCMODE = "InvalidParameterValue.BadSyncMode"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBSyncMode(request *ModifyDBSyncModeRequest) (response *ModifyDBSyncModeResponse, err error) {
    if request == nil {
        request = NewModifyDBSyncModeRequest()
    }
    
    response = NewModifyDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

// ModifyDBSyncMode
// This API is used to modify the sync mode of a TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_BADSYNCMODE = "InvalidParameterValue.BadSyncMode"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBSyncModeWithContext(ctx context.Context, request *ModifyDBSyncModeRequest) (response *ModifyDBSyncModeResponse, err error) {
    if request == nil {
        request = NewModifyDBSyncModeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBExtranetAccessRequest() (request *OpenDBExtranetAccessRequest) {
    request = &OpenDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "OpenDBExtranetAccess")
    
    
    return
}

func NewOpenDBExtranetAccessResponse() (response *OpenDBExtranetAccessResponse) {
    response = &OpenDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenDBExtranetAccess
// This API is used to enable public network access for a TencentDB instance. After that, you can access the instance with the public domain name and port obtained through the `DescribeDCDBInstances` API.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenDBExtranetAccess(request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

// OpenDBExtranetAccess
// This API is used to enable public network access for a TencentDB instance. After that, you can access the instance with the public domain name and port obtained through the `DescribeDCDBInstances` API.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenDBExtranetAccessWithContext(ctx context.Context, request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    request.SetContext(ctx)
    
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dcdb", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAccountPassword
// This API is used to reset the password of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_RESETPASSWORDFAILED = "FailedOperation.ResetPasswordFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

// ResetAccountPassword
// This API is used to reset the password of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_RESETPASSWORDFAILED = "FailedOperation.ResetPasswordFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}
