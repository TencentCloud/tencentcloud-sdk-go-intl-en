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

package v20170312

import (
    "context"
    "errors"
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "AssociateSecurityGroups")
    
    
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
//  FAILEDOPERATION_SETSVCLOCATIONFAILED = "FailedOperation.SetSvcLocationFailed"
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
//  RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
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
//  FAILEDOPERATION_SETSVCLOCATIONFAILED = "FailedOperation.SetSvcLocationFailed"
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
//  RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "AssociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroups require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CancelDcnJob")
    
    
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
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CancelDcnJob(request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
    return c.CancelDcnJobWithContext(context.Background(), request)
}

// CancelDcnJob
// This API is used to cancel DCN synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CancelDcnJobWithContext(ctx context.Context, request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
    if request == nil {
        request = NewCancelDcnJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CancelDcnJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelDcnJob require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CloneAccount")
    
    
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
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) CloneAccount(request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    return c.CloneAccountWithContext(context.Background(), request)
}

// CloneAccount
// This API is used to clone an instance account.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CloneAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneAccount require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CloseDBExtranetAccess")
    
    
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
//  INTERNALERROR_WANSERVICEFAILED = "InternalError.WanServiceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    return c.CloseDBExtranetAccessWithContext(context.Background(), request)
}

// CloseDBExtranetAccess
// This API is used to disable public network access for a TencentDB instance, which will make the public IP address inaccessible. The `DescribeDCDBInstances` API will not return the public domain name and port information of the corresponding instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INTERNALERROR_WANSERVICEFAILED = "InternalError.WanServiceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseDBExtranetAccessWithContext(ctx context.Context, request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CloseDBExtranetAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseDBExtranetAccess require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CopyAccountPrivileges")
    
    
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
// Note: accounts with the same username but different hosts are different accounts. Permissions can only be copied between accounts with the same `Readonly` attribute.
//
// error code that may be returned:
//  FAILEDOPERATION_COPYRIGHTERROR = "FailedOperation.CopyRightError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERTYPE = "InvalidParameterValue.BadUserType"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyAccountPrivileges(request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    return c.CopyAccountPrivilegesWithContext(context.Background(), request)
}

// CopyAccountPrivileges
// This API is used to copy the permissions of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts. Permissions can only be copied between accounts with the same `Readonly` attribute.
//
// error code that may be returned:
//  FAILEDOPERATION_COPYRIGHTERROR = "FailedOperation.CopyRightError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERTYPE = "InvalidParameterValue.BadUserType"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyAccountPrivilegesWithContext(ctx context.Context, request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewCopyAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CopyAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyAccountPrivileges require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateAccount")
    
    
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
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
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
    return c.CreateAccountWithContext(context.Background(), request)
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
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CreateAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateDBInstance")
    
    
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstance
// This API is used to create a monthly subscribed TencentDB for MariaDB instance by passing in information such as instance specifications, database version number, validity period, and quantity.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALCOUNT = "InvalidParameterValue.IllegalCount"
//  INVALIDPARAMETERVALUE_ILLEGALQUANTITY = "InvalidParameterValue.IllegalQuantity"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_DBVERSIONNOTSUPPORTED = "UnsupportedOperation.DbVersionNotSupported"
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    return c.CreateDBInstanceWithContext(context.Background(), request)
}

// CreateDBInstance
// This API is used to create a monthly subscribed TencentDB for MariaDB instance by passing in information such as instance specifications, database version number, validity period, and quantity.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALCOUNT = "InvalidParameterValue.IllegalCount"
//  INVALIDPARAMETERVALUE_ILLEGALQUANTITY = "InvalidParameterValue.IllegalQuantity"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_DBVERSIONNOTSUPPORTED = "UnsupportedOperation.DbVersionNotSupported"
func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CreateDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHourDBInstanceRequest() (request *CreateHourDBInstanceRequest) {
    request = &CreateHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "CreateHourDBInstance")
    
    
    return
}

func NewCreateHourDBInstanceResponse() (response *CreateHourDBInstanceResponse) {
    response = &CreateHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHourDBInstance
// This API is used to create a pay-as-you-go TencentDB for MariaDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALQUANTITY = "InvalidParameterValue.IllegalQuantity"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_DBVERSIONNOTSUPPORTED = "UnsupportedOperation.DbVersionNotSupported"
func (c *Client) CreateHourDBInstance(request *CreateHourDBInstanceRequest) (response *CreateHourDBInstanceResponse, err error) {
    return c.CreateHourDBInstanceWithContext(context.Background(), request)
}

// CreateHourDBInstance
// This API is used to create a pay-as-you-go TencentDB for MariaDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALQUANTITY = "InvalidParameterValue.IllegalQuantity"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_DBVERSIONNOTSUPPORTED = "UnsupportedOperation.DbVersionNotSupported"
func (c *Client) CreateHourDBInstanceWithContext(ctx context.Context, request *CreateHourDBInstanceRequest) (response *CreateHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateHourDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "CreateHourDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHourDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHourDBInstanceResponse()
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

// DeleteAccount
// This API is used to delete a TencentDB account, which is uniquely identified by username and host.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSERFAILED = "FailedOperation.DeleteUserFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    return c.DeleteAccountWithContext(context.Background(), request)
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
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DeleteAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccount require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAccountPrivileges")
    
    
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
// Note: Accounts with the same username but different hosts are different accounts.
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
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// This API is used to query the permissions of a TencentDB account.
//
// Note: Accounts with the same username but different hosts are different accounts.
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountPrivileges require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeAccounts")
    
    
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
    return c.DescribeAccountsWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupFilesRequest() (request *DescribeBackupFilesRequest) {
    request = &DescribeBackupFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeBackupFiles")
    
    
    return
}

func NewDescribeBackupFilesResponse() (response *DescribeBackupFilesResponse) {
    response = &DescribeBackupFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupFiles
// This API is used to query the list of backup files.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupFiles(request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
    return c.DescribeBackupFilesWithContext(context.Background(), request)
}

// DescribeBackupFiles
// This API is used to query the list of backup files.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupFilesWithContext(ctx context.Context, request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeBackupFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBEncryptAttributesRequest() (request *DescribeDBEncryptAttributesRequest) {
    request = &DescribeDBEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBEncryptAttributes")
    
    
    return
}

func NewDescribeDBEncryptAttributesResponse() (response *DescribeDBEncryptAttributesResponse) {
    response = &DescribeDBEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBEncryptAttributes
// This API is used to query the encryption status of the instance data.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCIPHERTEXTFAILED = "InternalError.GetCipherTextFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) DescribeDBEncryptAttributes(request *DescribeDBEncryptAttributesRequest) (response *DescribeDBEncryptAttributesResponse, err error) {
    return c.DescribeDBEncryptAttributesWithContext(context.Background(), request)
}

// DescribeDBEncryptAttributes
// This API is used to query the encryption status of the instance data.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCIPHERTEXTFAILED = "InternalError.GetCipherTextFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) DescribeDBEncryptAttributesWithContext(ctx context.Context, request *DescribeDBEncryptAttributesRequest) (response *DescribeDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeDBEncryptAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBEncryptAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBEncryptAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBEncryptAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceDetailRequest() (request *DescribeDBInstanceDetailRequest) {
    request = &DescribeDBInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBInstanceDetail")
    
    
    return
}

func NewDescribeDBInstanceDetailResponse() (response *DescribeDBInstanceDetailResponse) {
    response = &DescribeDBInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceDetail
// This API is used to query the details of a specified instance.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstanceDetail(request *DescribeDBInstanceDetailRequest) (response *DescribeDBInstanceDetailResponse, err error) {
    return c.DescribeDBInstanceDetailWithContext(context.Background(), request)
}

// DescribeDBInstanceDetail
// This API is used to query the details of a specified instance.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstanceDetailWithContext(ctx context.Context, request *DescribeDBInstanceDetailRequest) (response *DescribeDBInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceDetailResponse()
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

// DescribeDBInstances
// This API is used to query the TencentDB instance list. It supports filtering instances by project ID, instance ID, private address, and instance name.
//
// If no filter is specified, 20 instances will be returned by default. Up to 100 instances can be returned for a single request.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// This API is used to query the TencentDB instance list. It supports filtering instances by project ID, instance ID, private address, and instance name.
//
// If no filter is specified, 20 instances will be returned by default. Up to 100 instances can be returned for a single request.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBLogFiles
// This API is used to get the list of various logs of a database, including cold backups, binlogs, errlogs, and slowlogs.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_COSCONFIGURATION = "InternalError.CosConfiguration"
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_HDFSERROR = "InternalError.HDFSError"
//  INTERNALERROR_INNERCONFIGURATIONMISSING = "InternalError.InnerConfigurationMissing"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SHARDRESOURCEIDNOTFOUND = "InvalidParameter.ShardResourceIdNotFound"
//  RESOURCEUNAVAILABLE_COSAPIFAILED = "ResourceUnavailable.CosApiFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    return c.DescribeDBLogFilesWithContext(context.Background(), request)
}

// DescribeDBLogFiles
// This API is used to get the list of various logs of a database, including cold backups, binlogs, errlogs, and slowlogs.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_COSCONFIGURATION = "InternalError.CosConfiguration"
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_HDFSERROR = "InternalError.HDFSError"
//  INTERNALERROR_INNERCONFIGURATIONMISSING = "InternalError.InnerConfigurationMissing"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SHARDRESOURCEIDNOTFOUND = "InvalidParameter.ShardResourceIdNotFound"
//  RESOURCEUNAVAILABLE_COSAPIFAILED = "ResourceUnavailable.CosApiFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBLogFilesWithContext(ctx context.Context, request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBLogFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBLogFiles require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBParameters")
    
    
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
    return c.DescribeDBParametersWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBParameters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBParameters require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBSecurityGroups")
    
    
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
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// This API is used to query the security group details of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBSlowLogs
// This API is used to query the list of slow query logs.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETSLOWLOGFAILED = "InternalError.GetSlowLogFailed"
//  INTERNALERROR_LOGDBFAILED = "InternalError.LogDBFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_ILLEGALTIME = "InvalidParameter.IllegalTime"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSlowLogs(request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
    return c.DescribeDBSlowLogsWithContext(context.Background(), request)
}

// DescribeDBSlowLogs
// This API is used to query the list of slow query logs.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETSLOWLOGFAILED = "InternalError.GetSlowLogFailed"
//  INTERNALERROR_LOGDBFAILED = "InternalError.LogDBFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_ILLEGALTIME = "InvalidParameter.IllegalTime"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSlowLogsWithContext(ctx context.Context, request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBTmpInstancesRequest() (request *DescribeDBTmpInstancesRequest) {
    request = &DescribeDBTmpInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDBTmpInstances")
    
    
    return
}

func NewDescribeDBTmpInstancesResponse() (response *DescribeDBTmpInstancesResponse) {
    response = &DescribeDBTmpInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBTmpInstances
// This API is used to obtain a temp rollback instance.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBTmpInstances(request *DescribeDBTmpInstancesRequest) (response *DescribeDBTmpInstancesResponse, err error) {
    return c.DescribeDBTmpInstancesWithContext(context.Background(), request)
}

// DescribeDBTmpInstances
// This API is used to obtain a temp rollback instance.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBTmpInstancesWithContext(ctx context.Context, request *DescribeDBTmpInstancesRequest) (response *DescribeDBTmpInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBTmpInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDBTmpInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBTmpInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBTmpInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabaseObjects")
    
    
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
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    return c.DescribeDatabaseObjectsWithContext(context.Background(), request)
}

// DescribeDatabaseObjects
// This API is used to query the list of database objects in a TencentDB instance, including tables, stored procedures, views, and functions.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBOBJECTFAILED = "InternalError.GetDbObjectFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseObjectsWithContext(ctx context.Context, request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDatabaseObjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseObjects require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabaseTable")
    
    
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
    return c.DescribeDatabaseTableWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDatabaseTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseTable require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabases
// This API is used to query the database list of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBLISTFAILED = "InternalError.GetDbListFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// This API is used to query the database list of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBLISTFAILED = "InternalError.GetDbListFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeDcnDetail")
    
    
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
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDcnDetail(request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
    return c.DescribeDcnDetailWithContext(context.Background(), request)
}

// DescribeDcnDetail
// This API is used to query the disaster recovery details of an instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDcnDetailWithContext(ctx context.Context, request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDcnDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeDcnDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDcnDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDcnDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileDownloadUrlRequest() (request *DescribeFileDownloadUrlRequest) {
    request = &DescribeFileDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeFileDownloadUrl")
    
    
    return
}

func NewDescribeFileDownloadUrlResponse() (response *DescribeFileDownloadUrlResponse) {
    response = &DescribeFileDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileDownloadUrl
// This API is used to get the download URL of a specific backup or log file of a database.
//
// error code that may be returned:
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) DescribeFileDownloadUrl(request *DescribeFileDownloadUrlRequest) (response *DescribeFileDownloadUrlResponse, err error) {
    return c.DescribeFileDownloadUrlWithContext(context.Background(), request)
}

// DescribeFileDownloadUrl
// This API is used to get the download URL of a specific backup or log file of a database.
//
// error code that may be returned:
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) DescribeFileDownloadUrlWithContext(ctx context.Context, request *DescribeFileDownloadUrlRequest) (response *DescribeFileDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeFileDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeFileDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodeInfoRequest() (request *DescribeInstanceNodeInfoRequest) {
    request = &DescribeInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeInstanceNodeInfo")
    
    
    return
}

func NewDescribeInstanceNodeInfoResponse() (response *DescribeInstanceNodeInfoResponse) {
    response = &DescribeInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodeInfo
// This API is used to query the information of primary and replica nodes of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeInstanceNodeInfo(request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    return c.DescribeInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodeInfo
// This API is used to query the information of primary and replica nodes of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeInstanceNodeInfoWithContext(ctx context.Context, request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodeInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeInstanceNodeInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodeInfoResponse()
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

// DescribeLogFileRetentionPeriod
// This API is used to view the backup log retention days.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeLogFileRetentionPeriod(request *DescribeLogFileRetentionPeriodRequest) (response *DescribeLogFileRetentionPeriodResponse, err error) {
    return c.DescribeLogFileRetentionPeriodWithContext(context.Background(), request)
}

// DescribeLogFileRetentionPeriod
// This API is used to view the backup log retention days.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeLogFileRetentionPeriodWithContext(ctx context.Context, request *DescribeLogFileRetentionPeriodRequest) (response *DescribeLogFileRetentionPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeLogFileRetentionPeriodRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeLogFileRetentionPeriod")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogFileRetentionPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogFileRetentionPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribeOrders")
    
    
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrders
// This API is used to query TencentDB order information. You can pass in an order ID to query the TencentDB instance associated with the order and the corresponding task process ID.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYORDERFAILED = "InternalError.QueryOrderFailed"
//  INVALIDPARAMETER_DEALNAMENOTGIVEN = "InvalidParameter.DealNameNotGiven"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    return c.DescribeOrdersWithContext(context.Background(), request)
}

// DescribeOrders
// This API is used to query TencentDB order information. You can pass in an order ID to query the TencentDB instance associated with the order and the corresponding task process ID.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYORDERFAILED = "InternalError.QueryOrderFailed"
//  INVALIDPARAMETER_DEALNAMENOTGIVEN = "InvalidParameter.DealNameNotGiven"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeOrdersWithContext(ctx context.Context, request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeOrders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrders require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
    request = &DescribePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DescribePrice")
    
    
    return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
    response = &DescribePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrice
// This API is used to query the instance price before you purchase it.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALCOUNT = "InvalidParameterValue.IllegalCount"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
    return c.DescribePriceWithContext(context.Background(), request)
}

// DescribePrice
// This API is used to query the instance price before you purchase it.
//
// error code that may be returned:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALCOUNT = "InvalidParameterValue.IllegalCount"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
    if request == nil {
        request = NewDescribePriceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribePrice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePriceResponse()
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

// DescribeProjectSecurityGroups
// This API is used to query the security group details of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSERSGCOUNTFAILED = "InternalError.GetUserSGCountFailed"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// This API is used to query the security group details of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSERSGCOUNTFAILED = "InternalError.GetUserSGCountFailed"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
func (c *Client) DescribeProjectSecurityGroupsWithContext(ctx context.Context, request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DescribeProjectSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDBInstanceRequest() (request *DestroyDBInstanceRequest) {
    request = &DestroyDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DestroyDBInstance")
    
    
    return
}

func NewDestroyDBInstanceResponse() (response *DestroyDBInstanceResponse) {
    response = &DestroyDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyDBInstance
// This API is used to terminate an isolated monthly subscribed instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyDBInstance(request *DestroyDBInstanceRequest) (response *DestroyDBInstanceResponse, err error) {
    return c.DestroyDBInstanceWithContext(context.Background(), request)
}

// DestroyDBInstance
// This API is used to terminate an isolated monthly subscribed instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyDBInstanceWithContext(ctx context.Context, request *DestroyDBInstanceRequest) (response *DestroyDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DestroyDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyHourDBInstanceRequest() (request *DestroyHourDBInstanceRequest) {
    request = &DestroyHourDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "DestroyHourDBInstance")
    
    
    return
}

func NewDestroyHourDBInstanceResponse() (response *DestroyHourDBInstanceResponse) {
    response = &DestroyHourDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyHourDBInstance
// This API is used to terminate a pay-as-you-go TencentDB for MariaDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  FAILEDOPERATION_INSTANCERETURNFAILED = "FailedOperation.InstanceReturnFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyHourDBInstance(request *DestroyHourDBInstanceRequest) (response *DestroyHourDBInstanceResponse, err error) {
    return c.DestroyHourDBInstanceWithContext(context.Background(), request)
}

// DestroyHourDBInstance
// This API is used to terminate a pay-as-you-go TencentDB for MariaDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  FAILEDOPERATION_INSTANCERETURNFAILED = "FailedOperation.InstanceReturnFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyHourDBInstanceWithContext(ctx context.Context, request *DestroyHourDBInstanceRequest) (response *DestroyHourDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyHourDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DestroyHourDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyHourDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyHourDBInstanceResponse()
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
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
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
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
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
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "DisassociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "GrantAccountPrivileges")
    
    
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
    return c.GrantAccountPrivilegesWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "GrantAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GrantAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewGrantAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "IsolateDBInstance")
    
    
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDBInstance
// This API is used to isolate a monthly subscribed TencentDB for MariaDB instance, which will no longer be accessible via IP and port.  The isolated instance can be started up in the recycle bin.  If it is isolated due to overdue payments, top up your account as soon as possible.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_INSTANCECANNOTRETURN = "FailedOperation.InstanceCanNotReturn"
//  FAILEDOPERATION_INSTANCERETURNFAILED = "FailedOperation.InstanceReturnFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// This API is used to isolate a monthly subscribed TencentDB for MariaDB instance, which will no longer be accessible via IP and port.  The isolated instance can be started up in the recycle bin.  If it is isolated due to overdue payments, top up your account as soon as possible.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_INSTANCECANNOTRETURN = "FailedOperation.InstanceCanNotReturn"
//  FAILEDOPERATION_INSTANCERETURNFAILED = "FailedOperation.InstanceReturnFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "IsolateDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDedicatedDBInstanceRequest() (request *IsolateDedicatedDBInstanceRequest) {
    request = &IsolateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "IsolateDedicatedDBInstance")
    
    
    return
}

func NewIsolateDedicatedDBInstanceResponse() (response *IsolateDedicatedDBInstanceResponse) {
    response = &IsolateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDedicatedDBInstance
// This API is used to isolate a dedicated TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSISOLATEINSTANCEFAILED = "FailedOperation.OssIsolateInstanceFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) IsolateDedicatedDBInstance(request *IsolateDedicatedDBInstanceRequest) (response *IsolateDedicatedDBInstanceResponse, err error) {
    return c.IsolateDedicatedDBInstanceWithContext(context.Background(), request)
}

// IsolateDedicatedDBInstance
// This API is used to isolate a dedicated TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSISOLATEINSTANCEFAILED = "FailedOperation.OssIsolateInstanceFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) IsolateDedicatedDBInstanceWithContext(ctx context.Context, request *IsolateDedicatedDBInstanceRequest) (response *IsolateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDedicatedDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "IsolateDedicatedDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDedicatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewKillSessionRequest() (request *KillSessionRequest) {
    request = &KillSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "KillSession")
    
    
    return
}

func NewKillSessionResponse() (response *KillSessionResponse) {
    response = &KillSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillSession
// This API is used to kill the specified session.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) KillSession(request *KillSessionRequest) (response *KillSessionResponse, err error) {
    return c.KillSessionWithContext(context.Background(), request)
}

// KillSession
// This API is used to kill the specified session.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) KillSessionWithContext(ctx context.Context, request *KillSessionRequest) (response *KillSessionResponse, err error) {
    if request == nil {
        request = NewKillSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "KillSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillSessionResponse()
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
    return c.ModifyAccountDescriptionWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyAccountDescription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyAccountPrivileges")
    
    
    return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
    response = &ModifyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountPrivileges
// This API is used to modify the permissions of a TencentDB instance account. \n\n**Note**\n-Only the SELECT permission (that is, set the permission parameter to `["SELECT"]`) of the system database `mysql` can be granted.An error will be reported if read-write permissions are granted to a read-only account. If the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted view permissions, set `Privileges` to an empty array.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// This API is used to modify the permissions of a TencentDB instance account. \n\n**Note**\n-Only the SELECT permission (that is, set the permission parameter to `["SELECT"]`) of the system database `mysql` can be granted.An error will be reported if read-write permissions are granted to a read-only account. If the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted view permissions, set `Privileges` to an empty array.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBEncryptAttributesRequest() (request *ModifyDBEncryptAttributesRequest) {
    request = &ModifyDBEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBEncryptAttributes")
    
    
    return
}

func NewModifyDBEncryptAttributesResponse() (response *ModifyDBEncryptAttributesResponse) {
    response = &ModifyDBEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBEncryptAttributes
// This API is used to modify the instance data encryption.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBEncryptAttributes(request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    return c.ModifyDBEncryptAttributesWithContext(context.Background(), request)
}

// ModifyDBEncryptAttributes
// This API is used to modify the instance data encryption.
//
// error code that may be returned:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBEncryptAttributesWithContext(ctx context.Context, request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBEncryptAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyDBEncryptAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBEncryptAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBEncryptAttributesResponse()
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
    return c.ModifyDBInstancesProjectWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyDBInstancesProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstancesProject require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBParameters")
    
    
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
    return c.ModifyDBParametersWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyDBParameters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBParameters require credential")
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
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyDBSyncMode")
    
    
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
    return c.ModifyDBSyncModeWithContext(context.Background(), request)
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyDBSyncMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBSyncMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNetworkRequest() (request *ModifyInstanceNetworkRequest) {
    request = &ModifyInstanceNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceNetwork")
    
    
    return
}

func NewModifyInstanceNetworkResponse() (response *ModifyInstanceNetworkResponse) {
    response = &ModifyInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceNetwork
// This API is used to modify instance network.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_WANSTATUSABNORMAL = "FailedOperation.WanStatusAbnormal"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_VPCOPERATIONFAILED = "InternalError.VpcOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNetwork(request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    return c.ModifyInstanceNetworkWithContext(context.Background(), request)
}

// ModifyInstanceNetwork
// This API is used to modify instance network.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_WANSTATUSABNORMAL = "FailedOperation.WanStatusAbnormal"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_VPCOPERATIONFAILED = "InternalError.VpcOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNetworkWithContext(ctx context.Context, request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNetworkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyInstanceNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceVipRequest() (request *ModifyInstanceVipRequest) {
    request = &ModifyInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceVip")
    
    
    return
}

func NewModifyInstanceVipResponse() (response *ModifyInstanceVipResponse) {
    response = &ModifyInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceVip
// This API is used to modify instance VIP.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_VIPNOTCHANGE = "FailedOperation.VipNotChange"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVip(request *ModifyInstanceVipRequest) (response *ModifyInstanceVipResponse, err error) {
    return c.ModifyInstanceVipWithContext(context.Background(), request)
}

// ModifyInstanceVip
// This API is used to modify instance VIP.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_VIPNOTCHANGE = "FailedOperation.VipNotChange"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVipWithContext(ctx context.Context, request *ModifyInstanceVipRequest) (response *ModifyInstanceVipResponse, err error) {
    if request == nil {
        request = NewModifyInstanceVipRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyInstanceVip")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceVipResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceVportRequest() (request *ModifyInstanceVportRequest) {
    request = &ModifyInstanceVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifyInstanceVport")
    
    
    return
}

func NewModifyInstanceVportResponse() (response *ModifyInstanceVportResponse) {
    response = &ModifyInstanceVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceVport
// This API is used to modify instance Vport.
//
// error code that may be returned:
//  FAILEDOPERATION_SGCHANGEVIP = "FailedOperation.SGChangeVip"
//  FAILEDOPERATION_VPCADDSERVICEFAILED = "FailedOperation.VpcAddServiceFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  INVALIDPARAMETER_VPORTUSED = "InvalidParameter.VportUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVport(request *ModifyInstanceVportRequest) (response *ModifyInstanceVportResponse, err error) {
    return c.ModifyInstanceVportWithContext(context.Background(), request)
}

// ModifyInstanceVport
// This API is used to modify instance Vport.
//
// error code that may be returned:
//  FAILEDOPERATION_SGCHANGEVIP = "FailedOperation.SGChangeVip"
//  FAILEDOPERATION_VPCADDSERVICEFAILED = "FailedOperation.VpcAddServiceFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  INVALIDPARAMETER_VPORTUSED = "InvalidParameter.VportUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVportWithContext(ctx context.Context, request *ModifyInstanceVportRequest) (response *ModifyInstanceVportResponse, err error) {
    if request == nil {
        request = NewModifyInstanceVportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifyInstanceVport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifySyncTaskAttributeRequest() (request *ModifySyncTaskAttributeRequest) {
    request = &ModifySyncTaskAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "ModifySyncTaskAttribute")
    
    
    return
}

func NewModifySyncTaskAttributeResponse() (response *ModifySyncTaskAttributeResponse) {
    response = &ModifySyncTaskAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySyncTaskAttribute
// This API is used to modify sync task attributes (currently, only the task name can be modified).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  RESOURCENOTFOUND_SYNCTASKDELETED = "ResourceNotFound.SyncTaskDeleted"
func (c *Client) ModifySyncTaskAttribute(request *ModifySyncTaskAttributeRequest) (response *ModifySyncTaskAttributeResponse, err error) {
    return c.ModifySyncTaskAttributeWithContext(context.Background(), request)
}

// ModifySyncTaskAttribute
// This API is used to modify sync task attributes (currently, only the task name can be modified).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  RESOURCENOTFOUND_SYNCTASKDELETED = "ResourceNotFound.SyncTaskDeleted"
func (c *Client) ModifySyncTaskAttributeWithContext(ctx context.Context, request *ModifySyncTaskAttributeRequest) (response *ModifySyncTaskAttributeResponse, err error) {
    if request == nil {
        request = NewModifySyncTaskAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ModifySyncTaskAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySyncTaskAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySyncTaskAttributeResponse()
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

// ResetAccountPassword
// This API is used to reset the password of a TencentDB account.
//
// Note: accounts with the same username but different hosts are different accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_RESETPASSWORDFAILED = "FailedOperation.ResetPasswordFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
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
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
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
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "ResetAccountPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDedicatedDBInstanceRequest() (request *TerminateDedicatedDBInstanceRequest) {
    request = &TerminateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "TerminateDedicatedDBInstance")
    
    
    return
}

func NewTerminateDedicatedDBInstanceResponse() (response *TerminateDedicatedDBInstanceResponse) {
    response = &TerminateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateDedicatedDBInstance
// This API is used to terminate the isolated dedicated TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDedicatedDBInstance(request *TerminateDedicatedDBInstanceRequest) (response *TerminateDedicatedDBInstanceResponse, err error) {
    return c.TerminateDedicatedDBInstanceWithContext(context.Background(), request)
}

// TerminateDedicatedDBInstance
// This API is used to terminate the isolated dedicated TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDedicatedDBInstanceWithContext(ctx context.Context, request *TerminateDedicatedDBInstanceRequest) (response *TerminateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDedicatedDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "TerminateDedicatedDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDedicatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDedicatedDBInstanceRequest() (request *UpgradeDedicatedDBInstanceRequest) {
    request = &UpgradeDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mariadb", APIVersion, "UpgradeDedicatedDBInstance")
    
    
    return
}

func NewUpgradeDedicatedDBInstanceResponse() (response *UpgradeDedicatedDBInstanceResponse) {
    response = &UpgradeDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDedicatedDBInstance
// This API is used to expand the dedicated TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) UpgradeDedicatedDBInstance(request *UpgradeDedicatedDBInstanceRequest) (response *UpgradeDedicatedDBInstanceResponse, err error) {
    return c.UpgradeDedicatedDBInstanceWithContext(context.Background(), request)
}

// UpgradeDedicatedDBInstance
// This API is used to expand the dedicated TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) UpgradeDedicatedDBInstanceWithContext(ctx context.Context, request *UpgradeDedicatedDBInstanceRequest) (response *UpgradeDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDedicatedDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mariadb", APIVersion, "UpgradeDedicatedDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDedicatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}
