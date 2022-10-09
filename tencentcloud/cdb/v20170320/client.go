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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// AddTimeWindow
// This API (AddTimeWindow) is used to add a maintenance time window for a TencentDB instance, so as to specify when the instance can automatically perform access switch operations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) AddTimeWindow(request *AddTimeWindowRequest) (response *AddTimeWindowResponse, err error) {
    return c.AddTimeWindowWithContext(context.Background(), request)
}

// AddTimeWindow
// This API (AddTimeWindow) is used to add a maintenance time window for a TencentDB instance, so as to specify when the instance can automatically perform access switch operations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) AddTimeWindowWithContext(ctx context.Context, request *AddTimeWindowRequest) (response *AddTimeWindowResponse, err error) {
    if request == nil {
        request = NewAddTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTimeWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewApplyCDBProxyRequest() (request *ApplyCDBProxyRequest) {
    request = &ApplyCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ApplyCDBProxy")
    
    
    return
}

func NewApplyCDBProxyResponse() (response *ApplyCDBProxyResponse) {
    response = &ApplyCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyCDBProxy
// This API is used to create a database proxy group for a source instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REPEATCREATEPROXYERROR = "FailedOperation.RepeatCreateProxyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ApplyCDBProxy(request *ApplyCDBProxyRequest) (response *ApplyCDBProxyResponse, err error) {
    return c.ApplyCDBProxyWithContext(context.Background(), request)
}

// ApplyCDBProxy
// This API is used to create a database proxy group for a source instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REPEATCREATEPROXYERROR = "FailedOperation.RepeatCreateProxyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ApplyCDBProxyWithContext(ctx context.Context, request *ApplyCDBProxyRequest) (response *ApplyCDBProxyResponse, err error) {
    if request == nil {
        request = NewApplyCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyCDBProxyResponse()
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

// AssociateSecurityGroups
// This API (AssociateSecurityGroups) is used to bind security groups to instances in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// This API (AssociateSecurityGroups) is used to bind security groups to instances in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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

// BalanceRoGroupLoad
// This API is used to rebalance the loads of instances in an RO group. Please note that the database connections to those instances will be interrupted transiently; therefore, you should ensure that your application can reconnect to the databases. This operation should be performed with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) BalanceRoGroupLoad(request *BalanceRoGroupLoadRequest) (response *BalanceRoGroupLoadResponse, err error) {
    return c.BalanceRoGroupLoadWithContext(context.Background(), request)
}

// BalanceRoGroupLoad
// This API is used to rebalance the loads of instances in an RO group. Please note that the database connections to those instances will be interrupted transiently; therefore, you should ensure that your application can reconnect to the databases. This operation should be performed with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) BalanceRoGroupLoadWithContext(ctx context.Context, request *BalanceRoGroupLoadRequest) (response *BalanceRoGroupLoadResponse, err error) {
    if request == nil {
        request = NewBalanceRoGroupLoadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BalanceRoGroupLoad require credential")
    }

    request.SetContext(ctx)
    
    response = NewBalanceRoGroupLoadResponse()
    err = c.Send(request, response)
    return
}

func NewCloseCDBProxyRequest() (request *CloseCDBProxyRequest) {
    request = &CloseCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CloseCDBProxy")
    
    
    return
}

func NewCloseCDBProxyResponse() (response *CloseCDBProxyResponse) {
    response = &CloseCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseCDBProxy
// This API is used to disable database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CloseCDBProxy(request *CloseCDBProxyRequest) (response *CloseCDBProxyResponse, err error) {
    return c.CloseCDBProxyWithContext(context.Background(), request)
}

// CloseCDBProxy
// This API is used to disable database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CloseCDBProxyWithContext(ctx context.Context, request *CloseCDBProxyRequest) (response *CloseCDBProxyResponse, err error) {
    if request == nil {
        request = NewCloseCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseCDBProxyResponse()
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

// CloseWanService
// This API (CloseWanService) is used to disable public network access for TencentDB instance, which will make public IP addresses inaccessible.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CloseWanService(request *CloseWanServiceRequest) (response *CloseWanServiceResponse, err error) {
    return c.CloseWanServiceWithContext(context.Background(), request)
}

// CloseWanService
// This API (CloseWanService) is used to disable public network access for TencentDB instance, which will make public IP addresses inaccessible.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CloseWanServiceWithContext(ctx context.Context, request *CloseWanServiceRequest) (response *CloseWanServiceResponse, err error) {
    if request == nil {
        request = NewCloseWanServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseWanService require credential")
    }

    request.SetContext(ctx)
    
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

// CreateAccounts
// This API is used to create a TencentDB account. The account name, host address, and password are required. Account remarks and maximum connections can also be configured.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) CreateAccounts(request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    return c.CreateAccountsWithContext(context.Background(), request)
}

// CreateAccounts
// This API is used to create a TencentDB account. The account name, host address, and password are required. Account remarks and maximum connections can also be configured.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) CreateAccountsWithContext(ctx context.Context, request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    if request == nil {
        request = NewCreateAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditPolicyRequest() (request *CreateAuditPolicyRequest) {
    request = &CreateAuditPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAuditPolicy")
    
    
    return
}

func NewCreateAuditPolicyResponse() (response *CreateAuditPolicyResponse) {
    response = &CreateAuditPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAuditPolicy
// This API is used to create an audit policy for a TencentDB instance by associating an audit rule with the TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  FAILEDOPERATION_TYPEINCONFLICT = "FailedOperation.TypeInConflict"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITOSSLOGICERROR = "InternalError.AuditOssLogicError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITRULEHASBIND = "OperationDenied.AuditRuleHasBind"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_DBBRAINPOLICYCONFLICT = "OperationDenied.DBBrainPolicyConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAuditPolicy(request *CreateAuditPolicyRequest) (response *CreateAuditPolicyResponse, err error) {
    return c.CreateAuditPolicyWithContext(context.Background(), request)
}

// CreateAuditPolicy
// This API is used to create an audit policy for a TencentDB instance by associating an audit rule with the TencentDB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  FAILEDOPERATION_TYPEINCONFLICT = "FailedOperation.TypeInConflict"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITOSSLOGICERROR = "InternalError.AuditOssLogicError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITRULEHASBIND = "OperationDenied.AuditRuleHasBind"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_DBBRAINPOLICYCONFLICT = "OperationDenied.DBBrainPolicyConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAuditPolicyWithContext(ctx context.Context, request *CreateAuditPolicyRequest) (response *CreateAuditPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAuditPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditPolicyResponse()
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

// CreateBackup
// This API (CreateBackup) is used to create a TencentDB instance backup.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OVERQUOTA = "OverQuota"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// This API (CreateBackup) is used to create a TencentDB instance backup.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OVERQUOTA = "OverQuota"
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

// CreateCloneInstance
// This API is used to create a clone of a specific instance, and roll back the clone by using a physical backup file of the instance or roll back the clone to a point in time.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateCloneInstance(request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    return c.CreateCloneInstanceWithContext(context.Background(), request)
}

// CreateCloneInstance
// This API is used to create a clone of a specific instance, and roll back the clone by using a physical backup file of the instance or roll back the clone to a point in time.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateCloneInstanceWithContext(ctx context.Context, request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    if request == nil {
        request = NewCreateCloneInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloneInstance require credential")
    }

    request.SetContext(ctx)
    
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

// CreateDBImportJob
// This API (CreateDBImportJob) is used to create a data import task for a TencentDB instance.
//
// 
//
// Note that the files for a data import task must be uploaded to Tencent Cloud in advance. You need to do so in the console.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_EXECUTESQLERROR = "InternalError.ExecuteSQLError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CreateDBImportJob(request *CreateDBImportJobRequest) (response *CreateDBImportJobResponse, err error) {
    return c.CreateDBImportJobWithContext(context.Background(), request)
}

// CreateDBImportJob
// This API (CreateDBImportJob) is used to create a data import task for a TencentDB instance.
//
// 
//
// Note that the files for a data import task must be uploaded to Tencent Cloud in advance. You need to do so in the console.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_EXECUTESQLERROR = "InternalError.ExecuteSQLError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CreateDBImportJobWithContext(ctx context.Context, request *CreateDBImportJobRequest) (response *CreateDBImportJobResponse, err error) {
    if request == nil {
        request = NewCreateDBImportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBImportJob require credential")
    }

    request.SetContext(ctx)
    
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

// CreateDBInstanceHour
// This API is used to create pay-as-you-go TencentDB instances (which can be source instances, disaster recovery instances, or read-only replicas) by passing in information such as instance specifications, MySQL version number, and instance quantity.
//
// 
//
// This is an asynchronous API. You can also use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query instance details. If the output parameter `Status` is `1` and the output parameter `TaskStatus` is `0`, the instances have been successfully delivered.
//
// 
//
// 1. Use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the purchasable instance specifications, and then use the [DescribeDBPrice](https://intl.cloud.tencent.com/document/api/236/18566?from_cn_redirect=1) API to query the prices of the purchasable instances;
//
// 2. You can create up to 100 instances at a time, with an instance validity period of up to 36 months;
//
// 3. MySQL v5.5, v5.6, v5.7, and v8.0 are supported;
//
// 4. Source instances, disaster recovery instances, and read-only replicas can be created;
//
// 5. If `Port`, `ParamList`, or `Password` is specified in the input parameters, the instance (excluding basic instances) will be initialized.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    return c.CreateDBInstanceHourWithContext(context.Background(), request)
}

// CreateDBInstanceHour
// This API is used to create pay-as-you-go TencentDB instances (which can be source instances, disaster recovery instances, or read-only replicas) by passing in information such as instance specifications, MySQL version number, and instance quantity.
//
// 
//
// This is an asynchronous API. You can also use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query instance details. If the output parameter `Status` is `1` and the output parameter `TaskStatus` is `0`, the instances have been successfully delivered.
//
// 
//
// 1. Use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the purchasable instance specifications, and then use the [DescribeDBPrice](https://intl.cloud.tencent.com/document/api/236/18566?from_cn_redirect=1) API to query the prices of the purchasable instances;
//
// 2. You can create up to 100 instances at a time, with an instance validity period of up to 36 months;
//
// 3. MySQL v5.5, v5.6, v5.7, and v8.0 are supported;
//
// 4. Source instances, disaster recovery instances, and read-only replicas can be created;
//
// 5. If `Port`, `ParamList`, or `Password` is specified in the input parameters, the instance (excluding basic instances) will be initialized.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceHourWithContext(ctx context.Context, request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceHour require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceHourResponse()
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

// CreateParamTemplate
// This API is used to create a parameter template. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    return c.CreateParamTemplateWithContext(context.Background(), request)
}

// CreateParamTemplate
// This API is used to create a parameter template. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) CreateParamTemplateWithContext(ctx context.Context, request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    if request == nil {
        request = NewCreateParamTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateParamTemplate require credential")
    }

    request.SetContext(ctx)
    
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

// CreateRoInstanceIp
// This API is used to create a VIP exclusive to a TencentDB read-only instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEROVIPERROR = "FailedOperation.CreateRoVipError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) CreateRoInstanceIp(request *CreateRoInstanceIpRequest) (response *CreateRoInstanceIpResponse, err error) {
    return c.CreateRoInstanceIpWithContext(context.Background(), request)
}

// CreateRoInstanceIp
// This API is used to create a VIP exclusive to a TencentDB read-only instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEROVIPERROR = "FailedOperation.CreateRoVipError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) CreateRoInstanceIpWithContext(ctx context.Context, request *CreateRoInstanceIpRequest) (response *CreateRoInstanceIpResponse, err error) {
    if request == nil {
        request = NewCreateRoInstanceIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoInstanceIp require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteAccounts
// This API (DeleteAccounts) is used to delete TencentDB accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DeleteAccounts(request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    return c.DeleteAccountsWithContext(context.Background(), request)
}

// DeleteAccounts
// This API (DeleteAccounts) is used to delete TencentDB accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DeleteAccountsWithContext(ctx context.Context, request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccounts require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteBackup
// This API is used to delete a database backup. It can only delete manually initiated backups.
//
// error code that may be returned:
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    return c.DeleteBackupWithContext(context.Background(), request)
}

// DeleteBackup
// This API is used to delete a database backup. It can only delete manually initiated backups.
//
// error code that may be returned:
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DeleteBackupWithContext(ctx context.Context, request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    if request == nil {
        request = NewDeleteBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupResponse()
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

// DeleteParamTemplate
// This API is used to delete a parameter template. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    return c.DeleteParamTemplateWithContext(context.Background(), request)
}

// DeleteParamTemplate
// This API is used to delete a parameter template. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteParamTemplateWithContext(ctx context.Context, request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteParamTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteParamTemplate require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteTimeWindow
// This API (DeleteTimeWindow) is used to delete a maintenance time window for a TencentDB instance. After it is deleted, the default maintenance time window will be 03:00-04:00, i.e., switch to a new instance will be performed during 03:00-04:00 by default.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DeleteTimeWindow(request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    return c.DeleteTimeWindowWithContext(context.Background(), request)
}

// DeleteTimeWindow
// This API (DeleteTimeWindow) is used to delete a maintenance time window for a TencentDB instance. After it is deleted, the default maintenance time window will be 03:00-04:00, i.e., switch to a new instance will be performed during 03:00-04:00 by default.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DeleteTimeWindowWithContext(ctx context.Context, request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    if request == nil {
        request = NewDeleteTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTimeWindow require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeAccountPrivileges
// This API (DescribeAccountPrivileges) is used to query the information of TencentDB account permissions.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// This API (DescribeAccountPrivileges) is used to query the information of TencentDB account permissions.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// This API is used to query information of all TencentDB accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// This API is used to query information of all TencentDB accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
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

// DescribeAsyncRequestInfo
// This API (DescribeAsyncRequestInfo) is used to query the async task execution result of a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// This API (DescribeAsyncRequestInfo) is used to query the async task execution result of a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditPoliciesRequest() (request *DescribeAuditPoliciesRequest) {
    request = &DescribeAuditPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditPolicies")
    
    
    return
}

func NewDescribeAuditPoliciesResponse() (response *DescribeAuditPoliciesResponse) {
    response = &DescribeAuditPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditPolicies
// This API is used to query the audit policies of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeAuditPolicies(request *DescribeAuditPoliciesRequest) (response *DescribeAuditPoliciesResponse, err error) {
    return c.DescribeAuditPoliciesWithContext(context.Background(), request)
}

// DescribeAuditPolicies
// This API is used to query the audit policies of a TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeAuditPoliciesWithContext(ctx context.Context, request *DescribeAuditPoliciesRequest) (response *DescribeAuditPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRulesRequest() (request *DescribeAuditRulesRequest) {
    request = &DescribeAuditRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditRules")
    
    
    return
}

func NewDescribeAuditRulesResponse() (response *DescribeAuditRulesResponse) {
    response = &DescribeAuditRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditRules
// This API is used to query the audit rules in the current region.
//
// error code that may be returned:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
func (c *Client) DescribeAuditRules(request *DescribeAuditRulesRequest) (response *DescribeAuditRulesResponse, err error) {
    return c.DescribeAuditRulesWithContext(context.Background(), request)
}

// DescribeAuditRules
// This API is used to query the audit rules in the current region.
//
// error code that may be returned:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
func (c *Client) DescribeAuditRulesWithContext(ctx context.Context, request *DescribeAuditRulesRequest) (response *DescribeAuditRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRulesResponse()
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

// DescribeBackupConfig
// This API (DescribeBackupConfig) is used to query the configuration information of a TencentDB instance backup.
//
// error code that may be returned:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    return c.DescribeBackupConfigWithContext(context.Background(), request)
}

// DescribeBackupConfig
// This API (DescribeBackupConfig) is used to query the configuration information of a TencentDB instance backup.
//
// error code that may be returned:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupConfigWithContext(ctx context.Context, request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupDownloadRestriction
// This API is used to query the restrictions of downloading backups in a region.
//
// error code that may be returned:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// This API is used to query the restrictions of downloading backups in a region.
//
// error code that may be returned:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBackupDownloadRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadRestrictionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadRestrictionResponse()
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

// DescribeBackupOverview
// This API is used to query the backup overview of a user. It will return the user's current total number of backups, total capacity used by backups, capacity in the free tier, and paid capacity (all capacity values are in bytes).
//
// error code that may be returned:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeBackupOverview(request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    return c.DescribeBackupOverviewWithContext(context.Background(), request)
}

// DescribeBackupOverview
// This API is used to query the backup overview of a user. It will return the user's current total number of backups, total capacity used by backups, capacity in the free tier, and paid capacity (all capacity values are in bytes).
//
// error code that may be returned:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeBackupOverviewWithContext(ctx context.Context, request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupOverview require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBackupSummaries
// This API is used to query the statistics of backups. It will return the capacity used by backups at the instance level and the number and used capacity of data backups and log backups of each instance (all capacity values are in bytes).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBackupSummaries(request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    return c.DescribeBackupSummariesWithContext(context.Background(), request)
}

// DescribeBackupSummaries
// This API is used to query the statistics of backups. It will return the capacity used by backups at the instance level and the number and used capacity of data backups and log backups of each instance (all capacity values are in bytes).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBackupSummariesWithContext(ctx context.Context, request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupSummariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupSummaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupSummariesResponse()
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

// DescribeBackups
// This API (DescribeBackups) is used to query the backups of a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    return c.DescribeBackupsWithContext(context.Background(), request)
}

// DescribeBackups
// This API (DescribeBackups) is used to query the backups of a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

// DescribeBinlogBackupOverview
// This API is used to query the log backup overview of a user in the current region.
//
// error code that may be returned:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBinlogBackupOverview(request *DescribeBinlogBackupOverviewRequest) (response *DescribeBinlogBackupOverviewResponse, err error) {
    return c.DescribeBinlogBackupOverviewWithContext(context.Background(), request)
}

// DescribeBinlogBackupOverview
// This API is used to query the log backup overview of a user in the current region.
//
// error code that may be returned:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBinlogBackupOverviewWithContext(ctx context.Context, request *DescribeBinlogBackupOverviewRequest) (response *DescribeBinlogBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogBackupOverview require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeBinlogs
// This API is used to query the list of binlog files of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    return c.DescribeBinlogsWithContext(context.Background(), request)
}

// DescribeBinlogs
// This API is used to query the list of binlog files of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeBinlogsWithContext(ctx context.Context, request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCDBProxyRequest() (request *DescribeCDBProxyRequest) {
    request = &DescribeCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCDBProxy")
    
    
    return
}

func NewDescribeCDBProxyResponse() (response *DescribeCDBProxyResponse) {
    response = &DescribeCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCDBProxy
// This API is used to query database proxy. It will be deprecated and replaced by the `QueryCDBProxy` API.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCDBProxy(request *DescribeCDBProxyRequest) (response *DescribeCDBProxyResponse, err error) {
    return c.DescribeCDBProxyWithContext(context.Background(), request)
}

// DescribeCDBProxy
// This API is used to query database proxy. It will be deprecated and replaced by the `QueryCDBProxy` API.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCDBProxyWithContext(ctx context.Context, request *DescribeCDBProxyRequest) (response *DescribeCDBProxyResponse, err error) {
    if request == nil {
        request = NewDescribeCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdbZoneConfigRequest() (request *DescribeCdbZoneConfigRequest) {
    request = &DescribeCdbZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCdbZoneConfig")
    
    
    return
}

func NewDescribeCdbZoneConfigResponse() (response *DescribeCdbZoneConfigResponse) {
    response = &DescribeCdbZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdbZoneConfig
// This API is used to query the purchasable specifications of TencentDB instances in a region.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCdbZoneConfig(request *DescribeCdbZoneConfigRequest) (response *DescribeCdbZoneConfigResponse, err error) {
    return c.DescribeCdbZoneConfigWithContext(context.Background(), request)
}

// DescribeCdbZoneConfig
// This API is used to query the purchasable specifications of TencentDB instances in a region.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCdbZoneConfigWithContext(ctx context.Context, request *DescribeCdbZoneConfigRequest) (response *DescribeCdbZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribeCdbZoneConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdbZoneConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdbZoneConfigResponse()
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

// DescribeCloneList
// This API is used to query the clone task list of an instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCloneList(request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    return c.DescribeCloneListWithContext(context.Background(), request)
}

// DescribeCloneList
// This API is used to query the clone task list of an instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCloneListWithContext(ctx context.Context, request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    if request == nil {
        request = NewDescribeCloneListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloneList require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBImportRecords
// This API (DescribeDBImportRecords) is used to query the records of import tasks in a TencentDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBImportRecords(request *DescribeDBImportRecordsRequest) (response *DescribeDBImportRecordsResponse, err error) {
    return c.DescribeDBImportRecordsWithContext(context.Background(), request)
}

// DescribeDBImportRecords
// This API (DescribeDBImportRecords) is used to query the records of import tasks in a TencentDB instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBImportRecordsWithContext(ctx context.Context, request *DescribeDBImportRecordsRequest) (response *DescribeDBImportRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBImportRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBImportRecords require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBInstanceCharset
// This API (DescribeDBInstanceCharset) is used to query the character set and its name of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceCharset(request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    return c.DescribeDBInstanceCharsetWithContext(context.Background(), request)
}

// DescribeDBInstanceCharset
// This API (DescribeDBInstanceCharset) is used to query the character set and its name of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceCharsetWithContext(ctx context.Context, request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceCharsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceCharset require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBInstanceConfig
// This API (DescribeDBInstanceConfig) is used to query the configuration information of a TencentDB instance, such as its synchronization mode and deployment mode.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceConfig(request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    return c.DescribeDBInstanceConfigWithContext(context.Background(), request)
}

// DescribeDBInstanceConfig
// This API (DescribeDBInstanceConfig) is used to query the configuration information of a TencentDB instance, such as its synchronization mode and deployment mode.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceConfigWithContext(ctx context.Context, request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceConfig require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBInstanceGTID
// This API (DescribeDBInstanceGTID) is used to query whether GTID is activated for a TencentDB instance. Instances on or below version 5.5 are not supported.
//
// error code that may be returned:
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceGTID(request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    return c.DescribeDBInstanceGTIDWithContext(context.Background(), request)
}

// DescribeDBInstanceGTID
// This API (DescribeDBInstanceGTID) is used to query whether GTID is activated for a TencentDB instance. Instances on or below version 5.5 are not supported.
//
// error code that may be returned:
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceGTIDWithContext(ctx context.Context, request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceGTIDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceGTID require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBInstanceInfo
// This API is used to query the basic information of an instance (instance ID, instance name, and whether encryption is enabled).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) DescribeDBInstanceInfo(request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    return c.DescribeDBInstanceInfoWithContext(context.Background(), request)
}

// DescribeDBInstanceInfo
// This API is used to query the basic information of an instance (instance ID, instance name, and whether encryption is enabled).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) DescribeDBInstanceInfoWithContext(ctx context.Context, request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceInfo require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBInstanceRebootTime
// This API (DescribeDBInstanceRebootTime) is used to query the estimated time needed for a TencentDB instance to restart.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceRebootTime(request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    return c.DescribeDBInstanceRebootTimeWithContext(context.Background(), request)
}

// DescribeDBInstanceRebootTime
// This API (DescribeDBInstanceRebootTime) is used to query the estimated time needed for a TencentDB instance to restart.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceRebootTimeWithContext(ctx context.Context, request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceRebootTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceRebootTime require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBInstances
// This API (DescribeDBInstances) is used to query the list of TencentDB instances (which can be primary, disaster recovery, or read-only instances). It supports filtering instances by project ID, instance ID, access address, and instance status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// This API (DescribeDBInstances) is used to query the list of TencentDB instances (which can be primary, disaster recovery, or read-only instances). It supports filtering instances by project ID, instance ID, access address, and instance status.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeDBPriceRequest() (request *DescribeDBPriceRequest) {
    request = &DescribeDBPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBPrice")
    
    
    return
}

func NewDescribeDBPriceResponse() (response *DescribeDBPriceResponse) {
    response = &DescribeDBPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBPrice
// This API is used to query the purchase or renewal price of a pay-as-you-go or monthly subscribed TencentDB instance by passing in information such as instance type, purchase duration, number of instances to purchase, memory size, disk size, and AZ. For the price of instance renewal, you can pass in instance name to query.
//
// 
//
// Note: To query prices in a specific region, you need to use the access point of the region. For more information on access points, see <a href="https://www.tencentcloud.com/document/product/236/15832">Service Address</a>. For example, to query prices in Guangzhou, send a request to: cdb.ap-guangzhou.tencentcloudapi.com. Likewise, to query prices in Shanghai, send a request to: cdb.ap-shanghai.tencentcloudapi.com.
//
// error code that may be returned:
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDBPrice(request *DescribeDBPriceRequest) (response *DescribeDBPriceResponse, err error) {
    return c.DescribeDBPriceWithContext(context.Background(), request)
}

// DescribeDBPrice
// This API is used to query the purchase or renewal price of a pay-as-you-go or monthly subscribed TencentDB instance by passing in information such as instance type, purchase duration, number of instances to purchase, memory size, disk size, and AZ. For the price of instance renewal, you can pass in instance name to query.
//
// 
//
// Note: To query prices in a specific region, you need to use the access point of the region. For more information on access points, see <a href="https://www.tencentcloud.com/document/product/236/15832">Service Address</a>. For example, to query prices in Guangzhou, send a request to: cdb.ap-guangzhou.tencentcloudapi.com. Likewise, to query prices in Shanghai, send a request to: cdb.ap-shanghai.tencentcloudapi.com.
//
// error code that may be returned:
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDBPriceWithContext(ctx context.Context, request *DescribeDBPriceRequest) (response *DescribeDBPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDBPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBPriceResponse()
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

// DescribeDBSecurityGroups
// This API (DescribeDBSecurityGroups) is used to query the security group details of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// This API (DescribeDBSecurityGroups) is used to query the security group details of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDBSwitchRecords
// This API (DescribeDBSwitchRecords) is used to query the instance switch records.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBSwitchRecords(request *DescribeDBSwitchRecordsRequest) (response *DescribeDBSwitchRecordsResponse, err error) {
    return c.DescribeDBSwitchRecordsWithContext(context.Background(), request)
}

// DescribeDBSwitchRecords
// This API (DescribeDBSwitchRecords) is used to query the instance switch records.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBSwitchRecordsWithContext(ctx context.Context, request *DescribeDBSwitchRecordsRequest) (response *DescribeDBSwitchRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSwitchRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSwitchRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSwitchRecordsResponse()
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

// DescribeDataBackupOverview
// This API is used to query the data backup overview of a user in the current region.
//
// error code that may be returned:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDataBackupOverview(request *DescribeDataBackupOverviewRequest) (response *DescribeDataBackupOverviewResponse, err error) {
    return c.DescribeDataBackupOverviewWithContext(context.Background(), request)
}

// DescribeDataBackupOverview
// This API is used to query the data backup overview of a user in the current region.
//
// error code that may be returned:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDataBackupOverviewWithContext(ctx context.Context, request *DescribeDataBackupOverviewRequest) (response *DescribeDataBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDataBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataBackupOverview require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDatabases
// This API is used to query the information of databases in a TencentDB instance which must be a source or disaster recovery instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// This API is used to query the information of databases in a TencentDB instance which must be a source or disaster recovery instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDefaultParams
// This API (DescribeDefaultParams) is used to query the list of default configurable parameters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDefaultParams(request *DescribeDefaultParamsRequest) (response *DescribeDefaultParamsResponse, err error) {
    return c.DescribeDefaultParamsWithContext(context.Background(), request)
}

// DescribeDefaultParams
// This API (DescribeDefaultParams) is used to query the list of default configurable parameters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDefaultParamsWithContext(ctx context.Context, request *DescribeDefaultParamsRequest) (response *DescribeDefaultParamsResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultParamsResponse()
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

// DescribeDeviceMonitorInfo
// This API (DescribeDeviceMonitorInfo) is used to query the monitoring information of a TencentDB physical machine on the day. Currently, it only supports instances with 488 GB memory and 6 TB disk.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_FUNCTIONDENIED = "OperationDenied.FunctionDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeDeviceMonitorInfo(request *DescribeDeviceMonitorInfoRequest) (response *DescribeDeviceMonitorInfoResponse, err error) {
    return c.DescribeDeviceMonitorInfoWithContext(context.Background(), request)
}

// DescribeDeviceMonitorInfo
// This API (DescribeDeviceMonitorInfo) is used to query the monitoring information of a TencentDB physical machine on the day. Currently, it only supports instances with 488 GB memory and 6 TB disk.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_FUNCTIONDENIED = "OperationDenied.FunctionDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeDeviceMonitorInfoWithContext(ctx context.Context, request *DescribeDeviceMonitorInfoRequest) (response *DescribeDeviceMonitorInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceMonitorInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceMonitorInfo require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeErrorLogData
// This API is used to query the error logs of an instance over the past month by search criteria.
//
// Note: the HTTP response packet will be very large if it contain a single large error log, which causes the API call to time out. If this happens, we recommend you lower the value of the input parameter `Limit` to reduce the packet size so that the API can respond timely.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeErrorLogData(request *DescribeErrorLogDataRequest) (response *DescribeErrorLogDataResponse, err error) {
    return c.DescribeErrorLogDataWithContext(context.Background(), request)
}

// DescribeErrorLogData
// This API is used to query the error logs of an instance over the past month by search criteria.
//
// Note: the HTTP response packet will be very large if it contain a single large error log, which causes the API call to time out. If this happens, we recommend you lower the value of the input parameter `Limit` to reduce the packet size so that the API can respond timely.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeErrorLogDataWithContext(ctx context.Context, request *DescribeErrorLogDataRequest) (response *DescribeErrorLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeErrorLogDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeErrorLogData require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeInstanceParamRecords
// This API (DescribeInstanceParamRecords) is used to query the parameter modification records of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// This API (DescribeInstanceParamRecords) is used to query the parameter modification records of an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParams
// This API (DescribeInstanceParams) is used to query the list of parameters for an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// This API (DescribeInstanceParams) is used to query the list of parameters for an instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
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

func NewDescribeLocalBinlogConfigRequest() (request *DescribeLocalBinlogConfigRequest) {
    request = &DescribeLocalBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeLocalBinlogConfig")
    
    
    return
}

func NewDescribeLocalBinlogConfigResponse() (response *DescribeLocalBinlogConfigResponse) {
    response = &DescribeLocalBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLocalBinlogConfig
// This API is used to query the retention policy of local binlog of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeLocalBinlogConfig(request *DescribeLocalBinlogConfigRequest) (response *DescribeLocalBinlogConfigResponse, err error) {
    return c.DescribeLocalBinlogConfigWithContext(context.Background(), request)
}

// DescribeLocalBinlogConfig
// This API is used to query the retention policy of local binlog of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeLocalBinlogConfigWithContext(ctx context.Context, request *DescribeLocalBinlogConfigRequest) (response *DescribeLocalBinlogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeLocalBinlogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLocalBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLocalBinlogConfigResponse()
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

// DescribeParamTemplateInfo
// This API is used to query parameter template details. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplateInfo(request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    return c.DescribeParamTemplateInfoWithContext(context.Background(), request)
}

// DescribeParamTemplateInfo
// This API is used to query parameter template details. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplateInfoWithContext(ctx context.Context, request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplateInfo require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeParamTemplates
// This API is used to query the parameter template list. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// This API is used to query the parameter template list. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplatesWithContext(ctx context.Context, request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplates require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeProjectSecurityGroups
// This API (DescribeProjectSecurityGroups) is used to query the security group details of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// This API (DescribeProjectSecurityGroups) is used to query the security group details of a project.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProjectSecurityGroupsWithContext(ctx context.Context, request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyConnectionPoolConfRequest() (request *DescribeProxyConnectionPoolConfRequest) {
    request = &DescribeProxyConnectionPoolConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProxyConnectionPoolConf")
    
    
    return
}

func NewDescribeProxyConnectionPoolConfResponse() (response *DescribeProxyConnectionPoolConfResponse) {
    response = &DescribeProxyConnectionPoolConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyConnectionPoolConf
// This API is used to query the connection pool configuration of database proxy.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeProxyConnectionPoolConf(request *DescribeProxyConnectionPoolConfRequest) (response *DescribeProxyConnectionPoolConfResponse, err error) {
    return c.DescribeProxyConnectionPoolConfWithContext(context.Background(), request)
}

// DescribeProxyConnectionPoolConf
// This API is used to query the connection pool configuration of database proxy.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeProxyConnectionPoolConfWithContext(ctx context.Context, request *DescribeProxyConnectionPoolConfRequest) (response *DescribeProxyConnectionPoolConfResponse, err error) {
    if request == nil {
        request = NewDescribeProxyConnectionPoolConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyConnectionPoolConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyConnectionPoolConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyCustomConfRequest() (request *DescribeProxyCustomConfRequest) {
    request = &DescribeProxyCustomConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProxyCustomConf")
    
    
    return
}

func NewDescribeProxyCustomConfResponse() (response *DescribeProxyCustomConfResponse) {
    response = &DescribeProxyCustomConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyCustomConf
// This API is used to query the proxy configuration.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeProxyCustomConf(request *DescribeProxyCustomConfRequest) (response *DescribeProxyCustomConfResponse, err error) {
    return c.DescribeProxyCustomConfWithContext(context.Background(), request)
}

// DescribeProxyCustomConf
// This API is used to query the proxy configuration.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeProxyCustomConfWithContext(ctx context.Context, request *DescribeProxyCustomConfRequest) (response *DescribeProxyCustomConfResponse, err error) {
    if request == nil {
        request = NewDescribeProxyCustomConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyCustomConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyCustomConfResponse()
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

// DescribeRoGroups
// This API is used to query the information of all RO groups of a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeRoGroups(request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    return c.DescribeRoGroupsWithContext(context.Background(), request)
}

// DescribeRoGroups
// This API is used to query the information of all RO groups of a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeRoGroupsWithContext(ctx context.Context, request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRoGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoGroups require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeRoMinScale
// This API is used to query the minimum specification of a read-only instance that can be purchased or upgraded to.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeRoMinScale(request *DescribeRoMinScaleRequest) (response *DescribeRoMinScaleResponse, err error) {
    return c.DescribeRoMinScaleWithContext(context.Background(), request)
}

// DescribeRoMinScale
// This API is used to query the minimum specification of a read-only instance that can be purchased or upgraded to.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeRoMinScaleWithContext(ctx context.Context, request *DescribeRoMinScaleRequest) (response *DescribeRoMinScaleResponse, err error) {
    if request == nil {
        request = NewDescribeRoMinScaleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoMinScale require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeRollbackRangeTime
// This API (DescribeRollbackRangeTime) is used to query the time range available for instance rollback.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRollbackRangeTime(request *DescribeRollbackRangeTimeRequest) (response *DescribeRollbackRangeTimeResponse, err error) {
    return c.DescribeRollbackRangeTimeWithContext(context.Background(), request)
}

// DescribeRollbackRangeTime
// This API (DescribeRollbackRangeTime) is used to query the time range available for instance rollback.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRollbackRangeTimeWithContext(ctx context.Context, request *DescribeRollbackRangeTimeRequest) (response *DescribeRollbackRangeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackRangeTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackRangeTime require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeRollbackTaskDetail
// This API is used to query the details of a TencentDB instance rollback task.
//
// error code that may be returned:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeRollbackTaskDetail(request *DescribeRollbackTaskDetailRequest) (response *DescribeRollbackTaskDetailResponse, err error) {
    return c.DescribeRollbackTaskDetailWithContext(context.Background(), request)
}

// DescribeRollbackTaskDetail
// This API is used to query the details of a TencentDB instance rollback task.
//
// error code that may be returned:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeRollbackTaskDetailWithContext(ctx context.Context, request *DescribeRollbackTaskDetailRequest) (response *DescribeRollbackTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTaskDetail require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeSlowLogData
// This API is used to query the slow logs of an instance over the past month by search criteria.
//
// Note: the HTTP response packet will be very large if it contain a single large slow log, which causes the API call to time out. If this happens, we recommend you lower the value of the input parameter `Limit` to reduce the packet size so that the API can respond timely.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeSlowLogData(request *DescribeSlowLogDataRequest) (response *DescribeSlowLogDataResponse, err error) {
    return c.DescribeSlowLogDataWithContext(context.Background(), request)
}

// DescribeSlowLogData
// This API is used to query the slow logs of an instance over the past month by search criteria.
//
// Note: the HTTP response packet will be very large if it contain a single large slow log, which causes the API call to time out. If this happens, we recommend you lower the value of the input parameter `Limit` to reduce the packet size so that the API can respond timely.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeSlowLogDataWithContext(ctx context.Context, request *DescribeSlowLogDataRequest) (response *DescribeSlowLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogData require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeSlowLogs
// This API (DescribeSlowLogs) is used to query the slow logs of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    return c.DescribeSlowLogsWithContext(context.Background(), request)
}

// DescribeSlowLogs
// This API (DescribeSlowLogs) is used to query the slow logs of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogs require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeSupportedPrivileges
// This API (DescribeSupportedPrivileges) is used to query the information of TencentDB account permissions, including global permissions, database permissions, table permissions, and column permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeSupportedPrivileges(request *DescribeSupportedPrivilegesRequest) (response *DescribeSupportedPrivilegesResponse, err error) {
    return c.DescribeSupportedPrivilegesWithContext(context.Background(), request)
}

// DescribeSupportedPrivileges
// This API (DescribeSupportedPrivileges) is used to query the information of TencentDB account permissions, including global permissions, database permissions, table permissions, and column permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeSupportedPrivilegesWithContext(ctx context.Context, request *DescribeSupportedPrivilegesRequest) (response *DescribeSupportedPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeSupportedPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportedPrivileges require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTables
// This API is used to query the information of database tables in a TencentDB instance. It only supports source or disaster recovery instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// This API is used to query the information of database tables in a TencentDB instance. It only supports source or disaster recovery instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTagsOfInstanceIds
// This API (DescribeTagsOfInstanceIds) is used to query the tag information of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_TASKERROR = "InternalError.TaskError"
//  INTERNALERROR_TIMEWINDOWERROR = "InternalError.TimeWindowError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTagsOfInstanceIds(request *DescribeTagsOfInstanceIdsRequest) (response *DescribeTagsOfInstanceIdsResponse, err error) {
    return c.DescribeTagsOfInstanceIdsWithContext(context.Background(), request)
}

// DescribeTagsOfInstanceIds
// This API (DescribeTagsOfInstanceIds) is used to query the tag information of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_TASKERROR = "InternalError.TaskError"
//  INTERNALERROR_TIMEWINDOWERROR = "InternalError.TimeWindowError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTagsOfInstanceIdsWithContext(ctx context.Context, request *DescribeTagsOfInstanceIdsRequest) (response *DescribeTagsOfInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsOfInstanceIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagsOfInstanceIds require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTasks
// This API (DescribeTasks) is used to query the list of tasks for a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// This API (DescribeTasks) is used to query the list of tasks for a TencentDB instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTimeWindow
// This API (DescribeTimeWindow) is used to query the maintenance time window of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeTimeWindow(request *DescribeTimeWindowRequest) (response *DescribeTimeWindowResponse, err error) {
    return c.DescribeTimeWindowWithContext(context.Background(), request)
}

// DescribeTimeWindow
// This API (DescribeTimeWindow) is used to query the maintenance time window of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeTimeWindowWithContext(ctx context.Context, request *DescribeTimeWindowRequest) (response *DescribeTimeWindowResponse, err error) {
    if request == nil {
        request = NewDescribeTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimeWindow require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeUploadedFiles
// This API is used to query the list of SQL files imported by users. The common request parameter `Region` must be `ap-shanghai`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUploadedFiles(request *DescribeUploadedFilesRequest) (response *DescribeUploadedFilesResponse, err error) {
    return c.DescribeUploadedFilesWithContext(context.Background(), request)
}

// DescribeUploadedFiles
// This API is used to query the list of SQL files imported by users. The common request parameter `Region` must be `ap-shanghai`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUploadedFilesWithContext(ctx context.Context, request *DescribeUploadedFilesRequest) (response *DescribeUploadedFilesResponse, err error) {
    if request == nil {
        request = NewDescribeUploadedFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadedFiles require credential")
    }

    request.SetContext(ctx)
    
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

// DisassociateSecurityGroups
// This API (DisassociateSecurityGroups) is used to unbind security groups from instances in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// This API (DisassociateSecurityGroups) is used to unbind security groups from instances in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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

// InitDBInstances
// This API (InitDBInstances) is used to initialize instances, including their password, default character set, and instance port number.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    return c.InitDBInstancesWithContext(context.Background(), request)
}

// InitDBInstances
// This API (InitDBInstances) is used to initialize instances, including their password, default character set, and instance port number.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) InitDBInstancesWithContext(ctx context.Context, request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitDBInstances require credential")
    }

    request.SetContext(ctx)
    
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

// IsolateDBInstance
// This API is used to isolate a TencentDB instance, which will no longer be accessible via IP and port. The isolated instance can be started up in the recycle bin. If it is isolated due to arrears, please top up your account as soon as possible.
//
// error code that may be returned:
//  FAILEDOPERATION_CDBINSTANCELOCKFAILERROR = "FailedOperation.CdbInstanceLockFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTREFUNDERROR = "OperationDenied.UnSupportRefundError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// This API is used to isolate a TencentDB instance, which will no longer be accessible via IP and port. The isolated instance can be started up in the recycle bin. If it is isolated due to arrears, please top up your account as soon as possible.
//
// error code that may be returned:
//  FAILEDOPERATION_CDBINSTANCELOCKFAILERROR = "FailedOperation.CdbInstanceLockFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTREFUNDERROR = "OperationDenied.UnSupportRefundError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstance require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyAccountDescription
// This API (ModifyAccountDescription) is used to modify the remarks of a TencentDB instance account.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    return c.ModifyAccountDescriptionWithContext(context.Background(), request)
}

// ModifyAccountDescription
// This API (ModifyAccountDescription) is used to modify the remarks of a TencentDB instance account.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountMaxUserConnectionsRequest() (request *ModifyAccountMaxUserConnectionsRequest) {
    request = &ModifyAccountMaxUserConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountMaxUserConnections")
    
    
    return
}

func NewModifyAccountMaxUserConnectionsResponse() (response *ModifyAccountMaxUserConnectionsResponse) {
    response = &ModifyAccountMaxUserConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountMaxUserConnections
// This API is used to modify the maximum connections of one or more TencentDB instance accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_ASYNCTASKSTATUSERROR = "FailedOperation.AsyncTaskStatusError"
//  FAILEDOPERATION_JSONUNMARSHALERROR = "FailedOperation.JsonUnmarshalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
func (c *Client) ModifyAccountMaxUserConnections(request *ModifyAccountMaxUserConnectionsRequest) (response *ModifyAccountMaxUserConnectionsResponse, err error) {
    return c.ModifyAccountMaxUserConnectionsWithContext(context.Background(), request)
}

// ModifyAccountMaxUserConnections
// This API is used to modify the maximum connections of one or more TencentDB instance accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_ASYNCTASKSTATUSERROR = "FailedOperation.AsyncTaskStatusError"
//  FAILEDOPERATION_JSONUNMARSHALERROR = "FailedOperation.JsonUnmarshalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
func (c *Client) ModifyAccountMaxUserConnectionsWithContext(ctx context.Context, request *ModifyAccountMaxUserConnectionsRequest) (response *ModifyAccountMaxUserConnectionsResponse, err error) {
    if request == nil {
        request = NewModifyAccountMaxUserConnectionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountMaxUserConnections require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountMaxUserConnectionsResponse()
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

// ModifyAccountPassword
// This API (ModifyAccountPassword) is used to modify the password of a TencentDB instance account.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (response *ModifyAccountPasswordResponse, err error) {
    return c.ModifyAccountPasswordWithContext(context.Background(), request)
}

// ModifyAccountPassword
// This API (ModifyAccountPassword) is used to modify the password of a TencentDB instance account.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPasswordWithContext(ctx context.Context, request *ModifyAccountPasswordRequest) (response *ModifyAccountPasswordResponse, err error) {
    if request == nil {
        request = NewModifyAccountPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPassword require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyAccountPrivileges
// This API is used to modify the permissions of a TencentDB instance account.
//
// 
//
// Note that when modifying account permissions, you need to pass in the full permission information of the account. You can [query the account permission information
//
// ](https://intl.cloud.tencent.com/document/api/236/17500?from_cn_redirect=1) first before modifying permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// This API is used to modify the permissions of a TencentDB instance account.
//
// 
//
// Note that when modifying account permissions, you need to pass in the full permission information of the account. You can [query the account permission information
//
// ](https://intl.cloud.tencent.com/document/api/236/17500?from_cn_redirect=1) first before modifying permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyAutoRenewFlag
// This API is used to modify the auto-renewal flag of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// This API is used to modify the auto-renewal flag of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyBackupConfig
// This API (ModifyBackupConfig) is used to modify the database backup configuration.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    return c.ModifyBackupConfigWithContext(context.Background(), request)
}

// ModifyBackupConfig
// This API (ModifyBackupConfig) is used to modify the database backup configuration.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyBackupConfigWithContext(ctx context.Context, request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyBackupDownloadRestriction")
    
    
    return
}

func NewModifyBackupDownloadRestrictionResponse() (response *ModifyBackupDownloadRestrictionResponse) {
    response = &ModifyBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupDownloadRestriction
// This API is used to modify the restrictions of downloading backups in a region. You can specify which types of networks (private, or both private and public), VPCs, and IPs to download backups.
//
// error code that may be returned:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// This API is used to modify the restrictions of downloading backups in a region. You can specify which types of networks (private, or both private and public), VPCs, and IPs to download backups.
//
// error code that may be returned:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyBackupDownloadRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadRestrictionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyRequest() (request *ModifyCDBProxyRequest) {
    request = &ModifyCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxy")
    
    
    return
}

func NewModifyCDBProxyResponse() (response *ModifyCDBProxyResponse) {
    response = &ModifyCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxy
// This API is used to configure read/write separation of database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCDBProxy(request *ModifyCDBProxyRequest) (response *ModifyCDBProxyResponse, err error) {
    return c.ModifyCDBProxyWithContext(context.Background(), request)
}

// ModifyCDBProxy
// This API is used to configure read/write separation of database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCDBProxyWithContext(ctx context.Context, request *ModifyCDBProxyRequest) (response *ModifyCDBProxyResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyConnectionPoolRequest() (request *ModifyCDBProxyConnectionPoolRequest) {
    request = &ModifyCDBProxyConnectionPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxyConnectionPool")
    
    
    return
}

func NewModifyCDBProxyConnectionPoolResponse() (response *ModifyCDBProxyConnectionPoolResponse) {
    response = &ModifyCDBProxyConnectionPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxyConnectionPool
// This API is used to configure the connection pool of database proxy. You can use the `DescribeProxyConnectionPoolConf` API to query the supported connection pool configurations.
//
// error code that may be returned:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyConnectionPool(request *ModifyCDBProxyConnectionPoolRequest) (response *ModifyCDBProxyConnectionPoolResponse, err error) {
    return c.ModifyCDBProxyConnectionPoolWithContext(context.Background(), request)
}

// ModifyCDBProxyConnectionPool
// This API is used to configure the connection pool of database proxy. You can use the `DescribeProxyConnectionPoolConf` API to query the supported connection pool configurations.
//
// error code that may be returned:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyConnectionPoolWithContext(ctx context.Context, request *ModifyCDBProxyConnectionPoolRequest) (response *ModifyCDBProxyConnectionPoolResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyConnectionPoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxyConnectionPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyConnectionPoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyDescRequest() (request *ModifyCDBProxyDescRequest) {
    request = &ModifyCDBProxyDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxyDesc")
    
    
    return
}

func NewModifyCDBProxyDescResponse() (response *ModifyCDBProxyDescResponse) {
    response = &ModifyCDBProxyDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxyDesc
// This API is used to modify the description of database proxy.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyDesc(request *ModifyCDBProxyDescRequest) (response *ModifyCDBProxyDescResponse, err error) {
    return c.ModifyCDBProxyDescWithContext(context.Background(), request)
}

// ModifyCDBProxyDesc
// This API is used to modify the description of database proxy.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyDescWithContext(ctx context.Context, request *ModifyCDBProxyDescRequest) (response *ModifyCDBProxyDescResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyDescRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxyDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyDescResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyVipVPortRequest() (request *ModifyCDBProxyVipVPortRequest) {
    request = &ModifyCDBProxyVipVPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxyVipVPort")
    
    
    return
}

func NewModifyCDBProxyVipVPortResponse() (response *ModifyCDBProxyVipVPortResponse) {
    response = &ModifyCDBProxyVipVPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxyVipVPort
// This API is used to modify the VIP or port of database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCDBProxyVipVPort(request *ModifyCDBProxyVipVPortRequest) (response *ModifyCDBProxyVipVPortResponse, err error) {
    return c.ModifyCDBProxyVipVPortWithContext(context.Background(), request)
}

// ModifyCDBProxyVipVPort
// This API is used to modify the VIP or port of database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCDBProxyVipVPortWithContext(ctx context.Context, request *ModifyCDBProxyVipVPortRequest) (response *ModifyCDBProxyVipVPortResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyVipVPortRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxyVipVPort require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyVipVPortResponse()
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

// ModifyDBInstanceName
// This API (ModifyDBInstanceName) is used to rename a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    return c.ModifyDBInstanceNameWithContext(context.Background(), request)
}

// ModifyDBInstanceName
// This API (ModifyDBInstanceName) is used to rename a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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

// ModifyDBInstanceProject
// This API (ModifyDBInstanceProject) is used to modify the project to which a TencentDB instance belongs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    return c.ModifyDBInstanceProjectWithContext(context.Background(), request)
}

// ModifyDBInstanceProject
// This API (ModifyDBInstanceProject) is used to modify the project to which a TencentDB instance belongs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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

// ModifyDBInstanceSecurityGroups
// This API (ModifyDBInstanceSecurityGroups) is used to modify the security groups bound to a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// This API (ModifyDBInstanceSecurityGroups) is used to modify the security groups bound to a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyDBInstanceVipVport
// This API is used to modify the IP and port number of a TencentDB instance, switch from classic network to VPC, or change VPC subnets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceVipVport(request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    return c.ModifyDBInstanceVipVportWithContext(context.Background(), request)
}

// ModifyDBInstanceVipVport
// This API is used to modify the IP and port number of a TencentDB instance, switch from classic network to VPC, or change VPC subnets.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceVipVportWithContext(ctx context.Context, request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceVipVportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceVipVport require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyInstanceParam
// This API (ModifyInstanceParam) is used to modify instance parameters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    return c.ModifyInstanceParamWithContext(context.Background(), request)
}

// ModifyInstanceParam
// This API (ModifyInstanceParam) is used to modify instance parameters.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
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

func NewModifyInstancePasswordComplexityRequest() (request *ModifyInstancePasswordComplexityRequest) {
    request = &ModifyInstancePasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstancePasswordComplexity")
    
    
    return
}

func NewModifyInstancePasswordComplexityResponse() (response *ModifyInstancePasswordComplexityResponse) {
    response = &ModifyInstancePasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancePasswordComplexity
// This API is used to modify the password complexity of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstancePasswordComplexity(request *ModifyInstancePasswordComplexityRequest) (response *ModifyInstancePasswordComplexityResponse, err error) {
    return c.ModifyInstancePasswordComplexityWithContext(context.Background(), request)
}

// ModifyInstancePasswordComplexity
// This API is used to modify the password complexity of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstancePasswordComplexityWithContext(ctx context.Context, request *ModifyInstancePasswordComplexityRequest) (response *ModifyInstancePasswordComplexityResponse, err error) {
    if request == nil {
        request = NewModifyInstancePasswordComplexityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancePasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancePasswordComplexityResponse()
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

// ModifyInstanceTag
// This API (ModifyInstanceTag) is used to add, modify, or delete an instance tag.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyInstanceTag(request *ModifyInstanceTagRequest) (response *ModifyInstanceTagResponse, err error) {
    return c.ModifyInstanceTagWithContext(context.Background(), request)
}

// ModifyInstanceTag
// This API (ModifyInstanceTag) is used to add, modify, or delete an instance tag.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyInstanceTagWithContext(ctx context.Context, request *ModifyInstanceTagRequest) (response *ModifyInstanceTagResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceTagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLocalBinlogConfigRequest() (request *ModifyLocalBinlogConfigRequest) {
    request = &ModifyLocalBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyLocalBinlogConfig")
    
    
    return
}

func NewModifyLocalBinlogConfigResponse() (response *ModifyLocalBinlogConfigResponse) {
    response = &ModifyLocalBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLocalBinlogConfig
// This API is used to modify the retention policy of local binlog of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
func (c *Client) ModifyLocalBinlogConfig(request *ModifyLocalBinlogConfigRequest) (response *ModifyLocalBinlogConfigResponse, err error) {
    return c.ModifyLocalBinlogConfigWithContext(context.Background(), request)
}

// ModifyLocalBinlogConfig
// This API is used to modify the retention policy of local binlog of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
func (c *Client) ModifyLocalBinlogConfigWithContext(ctx context.Context, request *ModifyLocalBinlogConfigRequest) (response *ModifyLocalBinlogConfigResponse, err error) {
    if request == nil {
        request = NewModifyLocalBinlogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLocalBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLocalBinlogConfigResponse()
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

// ModifyNameOrDescByDpId
// This API is used to modify the name or description of a placement group.
//
// error code that may be returned:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) ModifyNameOrDescByDpId(request *ModifyNameOrDescByDpIdRequest) (response *ModifyNameOrDescByDpIdResponse, err error) {
    return c.ModifyNameOrDescByDpIdWithContext(context.Background(), request)
}

// ModifyNameOrDescByDpId
// This API is used to modify the name or description of a placement group.
//
// error code that may be returned:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) ModifyNameOrDescByDpIdWithContext(ctx context.Context, request *ModifyNameOrDescByDpIdRequest) (response *ModifyNameOrDescByDpIdResponse, err error) {
    if request == nil {
        request = NewModifyNameOrDescByDpIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNameOrDescByDpId require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyParamTemplate
// This API is used to modify a parameter template. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    return c.ModifyParamTemplateWithContext(context.Background(), request)
}

// ModifyParamTemplate
// This API is used to modify a parameter template. The common request parameter `Region` can only be set to `ap-guangzhou`.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) ModifyParamTemplateWithContext(ctx context.Context, request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    if request == nil {
        request = NewModifyParamTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyParamTemplate require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyRoGroupInfo
// This API is used to update the information of a TencentDB RO group, such as configuring a read-only instance removal policy in case of excessive delay, setting read weights of read-only instances, and setting the replication delay.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_CONFLICTROSTATUS = "OperationDenied.ConflictRoStatus"
//  OPERATIONDENIED_CONFLICTSTATUS = "OperationDenied.ConflictStatus"
//  OPERATIONDENIED_DELAYREPLICATIONRUNNING = "OperationDenied.DelayReplicationRunning"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) ModifyRoGroupInfo(request *ModifyRoGroupInfoRequest) (response *ModifyRoGroupInfoResponse, err error) {
    return c.ModifyRoGroupInfoWithContext(context.Background(), request)
}

// ModifyRoGroupInfo
// This API is used to update the information of a TencentDB RO group, such as configuring a read-only instance removal policy in case of excessive delay, setting read weights of read-only instances, and setting the replication delay.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_CONFLICTROSTATUS = "OperationDenied.ConflictRoStatus"
//  OPERATIONDENIED_CONFLICTSTATUS = "OperationDenied.ConflictStatus"
//  OPERATIONDENIED_DELAYREPLICATIONRUNNING = "OperationDenied.DelayReplicationRunning"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) ModifyRoGroupInfoWithContext(ctx context.Context, request *ModifyRoGroupInfoRequest) (response *ModifyRoGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyRoGroupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoGroupInfo require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyTimeWindow
// This API (ModifyTimeWindow) is used to update the maintenance time window of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyTimeWindow(request *ModifyTimeWindowRequest) (response *ModifyTimeWindowResponse, err error) {
    return c.ModifyTimeWindowWithContext(context.Background(), request)
}

// ModifyTimeWindow
// This API (ModifyTimeWindow) is used to update the maintenance time window of a TencentDB instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyTimeWindowWithContext(ctx context.Context, request *ModifyTimeWindowRequest) (response *ModifyTimeWindowResponse, err error) {
    if request == nil {
        request = NewModifyTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTimeWindow require credential")
    }

    request.SetContext(ctx)
    
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

// OfflineIsolatedInstances
// This API (OfflineIsolatedInstances) is used to deactivate isolated TencentDB instances immediately. The instances must be in isolated status, i.e., their `Status` value is 5 in the return of the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1).
//
// 
//
// This is an asynchronous API. There may be a delay in repossessing some resources. You can query the details by using the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) and specifying the InstanceId and the `Status` value as [5, 6, 7]. If the returned instance is empty, then all its resources have been released.
//
// 
//
// Note that once an instance is deactivated, its resources and data will not be recoverable. Please do so with caution.
//
// error code that may be returned:
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OfflineIsolatedInstances(request *OfflineIsolatedInstancesRequest) (response *OfflineIsolatedInstancesResponse, err error) {
    return c.OfflineIsolatedInstancesWithContext(context.Background(), request)
}

// OfflineIsolatedInstances
// This API (OfflineIsolatedInstances) is used to deactivate isolated TencentDB instances immediately. The instances must be in isolated status, i.e., their `Status` value is 5 in the return of the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1).
//
// 
//
// This is an asynchronous API. There may be a delay in repossessing some resources. You can query the details by using the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) and specifying the InstanceId and the `Status` value as [5, 6, 7]. If the returned instance is empty, then all its resources have been released.
//
// 
//
// Note that once an instance is deactivated, its resources and data will not be recoverable. Please do so with caution.
//
// error code that may be returned:
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OfflineIsolatedInstancesWithContext(ctx context.Context, request *OfflineIsolatedInstancesRequest) (response *OfflineIsolatedInstancesResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAuditServiceRequest() (request *OpenAuditServiceRequest) {
    request = &OpenAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OpenAuditService")
    
    
    return
}

func NewOpenAuditServiceResponse() (response *OpenAuditServiceResponse) {
    response = &OpenAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenAuditService
// This API is used to enable the audit service.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// This API is used to enable the audit service.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAuditServiceResponse()
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

// OpenDBInstanceGTID
// This API (OpenDBInstanceGTID) is used to enable GTID for a TencentDB instance. Only instances on or above version 5.6 are supported.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OpenDBInstanceGTID(request *OpenDBInstanceGTIDRequest) (response *OpenDBInstanceGTIDResponse, err error) {
    return c.OpenDBInstanceGTIDWithContext(context.Background(), request)
}

// OpenDBInstanceGTID
// This API (OpenDBInstanceGTID) is used to enable GTID for a TencentDB instance. Only instances on or above version 5.6 are supported.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OpenDBInstanceGTIDWithContext(ctx context.Context, request *OpenDBInstanceGTIDRequest) (response *OpenDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewOpenDBInstanceGTIDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenDBInstanceGTID require credential")
    }

    request.SetContext(ctx)
    
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

// OpenWanService
// This API (OpenWanService) is used to enable public network access for an instance.
//
// 
//
// Note that before enabling public network access, you need to first [initialize the instance](https://intl.cloud.tencent.com/document/api/236/15873?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) OpenWanService(request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    return c.OpenWanServiceWithContext(context.Background(), request)
}

// OpenWanService
// This API (OpenWanService) is used to enable public network access for an instance.
//
// 
//
// Note that before enabling public network access, you need to first [initialize the instance](https://intl.cloud.tencent.com/document/api/236/15873?from_cn_redirect=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) OpenWanServiceWithContext(ctx context.Context, request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    if request == nil {
        request = NewOpenWanServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWanService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCDBProxyRequest() (request *QueryCDBProxyRequest) {
    request = &QueryCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "QueryCDBProxy")
    
    
    return
}

func NewQueryCDBProxyResponse() (response *QueryCDBProxyResponse) {
    response = &QueryCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCDBProxy
// This API is used to query the proxy details.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) QueryCDBProxy(request *QueryCDBProxyRequest) (response *QueryCDBProxyResponse, err error) {
    return c.QueryCDBProxyWithContext(context.Background(), request)
}

// QueryCDBProxy
// This API is used to query the proxy details.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) QueryCDBProxyWithContext(ctx context.Context, request *QueryCDBProxyRequest) (response *QueryCDBProxyResponse, err error) {
    if request == nil {
        request = NewQueryCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCDBProxyResponse()
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

// ReleaseIsolatedDBInstances
// This API is used to deisolate an isolated TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) ReleaseIsolatedDBInstances(request *ReleaseIsolatedDBInstancesRequest) (response *ReleaseIsolatedDBInstancesResponse, err error) {
    return c.ReleaseIsolatedDBInstancesWithContext(context.Background(), request)
}

// ReleaseIsolatedDBInstances
// This API is used to deisolate an isolated TencentDB instance.
//
// error code that may be returned:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) ReleaseIsolatedDBInstancesWithContext(ctx context.Context, request *ReleaseIsolatedDBInstancesRequest) (response *ReleaseIsolatedDBInstancesResponse, err error) {
    if request == nil {
        request = NewReleaseIsolatedDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseIsolatedDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseIsolatedDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewReloadBalanceProxyNodeRequest() (request *ReloadBalanceProxyNodeRequest) {
    request = &ReloadBalanceProxyNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ReloadBalanceProxyNode")
    
    
    return
}

func NewReloadBalanceProxyNodeResponse() (response *ReloadBalanceProxyNodeResponse) {
    response = &ReloadBalanceProxyNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReloadBalanceProxyNode
// This API is used to rebalance the load on database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
func (c *Client) ReloadBalanceProxyNode(request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    return c.ReloadBalanceProxyNodeWithContext(context.Background(), request)
}

// ReloadBalanceProxyNode
// This API is used to rebalance the load on database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
func (c *Client) ReloadBalanceProxyNodeWithContext(ctx context.Context, request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    if request == nil {
        request = NewReloadBalanceProxyNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReloadBalanceProxyNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewReloadBalanceProxyNodeResponse()
    err = c.Send(request, response)
    return
}

func NewResetRootAccountRequest() (request *ResetRootAccountRequest) {
    request = &ResetRootAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ResetRootAccount")
    
    
    return
}

func NewResetRootAccountResponse() (response *ResetRootAccountResponse) {
    response = &ResetRootAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetRootAccount
// This API is used to reset the root account and initialize the account permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ResetRootAccount(request *ResetRootAccountRequest) (response *ResetRootAccountResponse, err error) {
    return c.ResetRootAccountWithContext(context.Background(), request)
}

// ResetRootAccount
// This API is used to reset the root account and initialize the account permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ResetRootAccountWithContext(ctx context.Context, request *ResetRootAccountRequest) (response *ResetRootAccountResponse, err error) {
    if request == nil {
        request = NewResetRootAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRootAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetRootAccountResponse()
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

// RestartDBInstances
// This API (RestartDBInstances) is used to restart TencentDB instances.
//
// 
//
// Note:
//
// 1. This API only supports restarting primary instances.
//
// 2. The instance status must be normal, and no other async tasks are in progress.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) RestartDBInstances(request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    return c.RestartDBInstancesWithContext(context.Background(), request)
}

// RestartDBInstances
// This API (RestartDBInstances) is used to restart TencentDB instances.
//
// 
//
// Note:
//
// 1. This API only supports restarting primary instances.
//
// 2. The instance status must be normal, and no other async tasks are in progress.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) RestartDBInstancesWithContext(ctx context.Context, request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    if request == nil {
        request = NewRestartDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDBInstances require credential")
    }

    request.SetContext(ctx)
    
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

// StartBatchRollback
// This API (StartBatchRollback) is used to roll back the tables of a TencentDB instance in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  OVERQUOTA = "OverQuota"
func (c *Client) StartBatchRollback(request *StartBatchRollbackRequest) (response *StartBatchRollbackResponse, err error) {
    return c.StartBatchRollbackWithContext(context.Background(), request)
}

// StartBatchRollback
// This API (StartBatchRollback) is used to roll back the tables of a TencentDB instance in batches.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  OVERQUOTA = "OverQuota"
func (c *Client) StartBatchRollbackWithContext(ctx context.Context, request *StartBatchRollbackRequest) (response *StartBatchRollbackResponse, err error) {
    if request == nil {
        request = NewStartBatchRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBatchRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBatchRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewStartReplicationRequest() (request *StartReplicationRequest) {
    request = &StartReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StartReplication")
    
    
    return
}

func NewStartReplicationResponse() (response *StartReplicationResponse) {
    response = &StartReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartReplication
// This API is used to start the data replication from the source instance to the read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StartReplication(request *StartReplicationRequest) (response *StartReplicationResponse, err error) {
    return c.StartReplicationWithContext(context.Background(), request)
}

// StartReplication
// This API is used to start the data replication from the source instance to the read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StartReplicationWithContext(ctx context.Context, request *StartReplicationRequest) (response *StartReplicationResponse, err error) {
    if request == nil {
        request = NewStartReplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartReplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartReplicationResponse()
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

// StopDBImportJob
// This API (StopDBImportJob) is used to stop a data import task.
//
// error code that may be returned:
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) StopDBImportJob(request *StopDBImportJobRequest) (response *StopDBImportJobResponse, err error) {
    return c.StopDBImportJobWithContext(context.Background(), request)
}

// StopDBImportJob
// This API (StopDBImportJob) is used to stop a data import task.
//
// error code that may be returned:
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) StopDBImportJobWithContext(ctx context.Context, request *StopDBImportJobRequest) (response *StopDBImportJobResponse, err error) {
    if request == nil {
        request = NewStopDBImportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDBImportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopReplicationRequest() (request *StopReplicationRequest) {
    request = &StopReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StopReplication")
    
    
    return
}

func NewStopReplicationResponse() (response *StopReplicationResponse) {
    response = &StopReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopReplication
// This API is used to stop the data replication from the source instance to the read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDELAYRO = "FailedOperation.NotDelayRo"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StopReplication(request *StopReplicationRequest) (response *StopReplicationResponse, err error) {
    return c.StopReplicationWithContext(context.Background(), request)
}

// StopReplication
// This API is used to stop the data replication from the source instance to the read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDELAYRO = "FailedOperation.NotDelayRo"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StopReplicationWithContext(ctx context.Context, request *StopReplicationRequest) (response *StopReplicationResponse, err error) {
    if request == nil {
        request = NewStopReplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopReplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopReplicationResponse()
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

// StopRollback
// This API is used to cancel a rollback task in progress, and returns an async task ID. You can use the `DescribeAsyncRequestInfo` API to query the result of cancellation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_ASYNCREQUESTERROR = "InternalError.AsyncRequestError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StopRollback(request *StopRollbackRequest) (response *StopRollbackResponse, err error) {
    return c.StopRollbackWithContext(context.Background(), request)
}

// StopRollback
// This API is used to cancel a rollback task in progress, and returns an async task ID. You can use the `DescribeAsyncRequestInfo` API to query the result of cancellation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_ASYNCREQUESTERROR = "InternalError.AsyncRequestError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StopRollbackWithContext(ctx context.Context, request *StopRollbackRequest) (response *StopRollbackResponse, err error) {
    if request == nil {
        request = NewStopRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchCDBProxyRequest() (request *SwitchCDBProxyRequest) {
    request = &SwitchCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchCDBProxy")
    
    
    return
}

func NewSwitchCDBProxyResponse() (response *SwitchCDBProxyResponse) {
    response = &SwitchCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchCDBProxy
// This API is used to switch database proxy after the proxy configuration is modified or the proxy version is upgraded.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) SwitchCDBProxy(request *SwitchCDBProxyRequest) (response *SwitchCDBProxyResponse, err error) {
    return c.SwitchCDBProxyWithContext(context.Background(), request)
}

// SwitchCDBProxy
// This API is used to switch database proxy after the proxy configuration is modified or the proxy version is upgraded.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) SwitchCDBProxyWithContext(ctx context.Context, request *SwitchCDBProxyRequest) (response *SwitchCDBProxyResponse, err error) {
    if request == nil {
        request = NewSwitchCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDBInstanceMasterSlaveRequest() (request *SwitchDBInstanceMasterSlaveRequest) {
    request = &SwitchDBInstanceMasterSlaveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchDBInstanceMasterSlave")
    
    
    return
}

func NewSwitchDBInstanceMasterSlaveResponse() (response *SwitchDBInstanceMasterSlaveResponse) {
    response = &SwitchDBInstanceMasterSlaveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDBInstanceMasterSlave
// This API is used for source-to-replica switch.
//
// error code that may be returned:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) SwitchDBInstanceMasterSlave(request *SwitchDBInstanceMasterSlaveRequest) (response *SwitchDBInstanceMasterSlaveResponse, err error) {
    return c.SwitchDBInstanceMasterSlaveWithContext(context.Background(), request)
}

// SwitchDBInstanceMasterSlave
// This API is used for source-to-replica switch.
//
// error code that may be returned:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) SwitchDBInstanceMasterSlaveWithContext(ctx context.Context, request *SwitchDBInstanceMasterSlaveRequest) (response *SwitchDBInstanceMasterSlaveResponse, err error) {
    if request == nil {
        request = NewSwitchDBInstanceMasterSlaveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDBInstanceMasterSlave require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDBInstanceMasterSlaveResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDrInstanceToMasterRequest() (request *SwitchDrInstanceToMasterRequest) {
    request = &SwitchDrInstanceToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchDrInstanceToMaster")
    
    
    return
}

func NewSwitchDrInstanceToMasterResponse() (response *SwitchDrInstanceToMasterResponse) {
    response = &SwitchDrInstanceToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDrInstanceToMaster
// This API is used to promote a disaster recovery instance to source instance. The request parameter `Region` must be the region of the disaster recovery instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) SwitchDrInstanceToMaster(request *SwitchDrInstanceToMasterRequest) (response *SwitchDrInstanceToMasterResponse, err error) {
    return c.SwitchDrInstanceToMasterWithContext(context.Background(), request)
}

// SwitchDrInstanceToMaster
// This API is used to promote a disaster recovery instance to source instance. The request parameter `Region` must be the region of the disaster recovery instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) SwitchDrInstanceToMasterWithContext(ctx context.Context, request *SwitchDrInstanceToMasterRequest) (response *SwitchDrInstanceToMasterResponse, err error) {
    if request == nil {
        request = NewSwitchDrInstanceToMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDrInstanceToMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDrInstanceToMasterResponse()
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

// SwitchForUpgrade
// This API (SwitchForUpgrade) is used to switch to a new instance. You can initiate this process when the primary instance being upgraded is pending switch.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) SwitchForUpgrade(request *SwitchForUpgradeRequest) (response *SwitchForUpgradeResponse, err error) {
    return c.SwitchForUpgradeWithContext(context.Background(), request)
}

// SwitchForUpgrade
// This API (SwitchForUpgrade) is used to switch to a new instance. You can initiate this process when the primary instance being upgraded is pending switch.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) SwitchForUpgradeWithContext(ctx context.Context, request *SwitchForUpgradeRequest) (response *SwitchForUpgradeResponse, err error) {
    if request == nil {
        request = NewSwitchForUpgradeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchForUpgrade require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchForUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeCDBProxyRequest() (request *UpgradeCDBProxyRequest) {
    request = &UpgradeCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeCDBProxy")
    
    
    return
}

func NewUpgradeCDBProxyResponse() (response *UpgradeCDBProxyResponse) {
    response = &UpgradeCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeCDBProxy
// This API is used to upgrade the configuration of database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) UpgradeCDBProxy(request *UpgradeCDBProxyRequest) (response *UpgradeCDBProxyResponse, err error) {
    return c.UpgradeCDBProxyWithContext(context.Background(), request)
}

// UpgradeCDBProxy
// This API is used to upgrade the configuration of database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) UpgradeCDBProxyWithContext(ctx context.Context, request *UpgradeCDBProxyRequest) (response *UpgradeCDBProxyResponse, err error) {
    if request == nil {
        request = NewUpgradeCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeCDBProxyVersionRequest() (request *UpgradeCDBProxyVersionRequest) {
    request = &UpgradeCDBProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeCDBProxyVersion")
    
    
    return
}

func NewUpgradeCDBProxyVersionResponse() (response *UpgradeCDBProxyVersionResponse) {
    response = &UpgradeCDBProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeCDBProxyVersion
// This API is used to upgrade the version of database proxy.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) UpgradeCDBProxyVersion(request *UpgradeCDBProxyVersionRequest) (response *UpgradeCDBProxyVersionResponse, err error) {
    return c.UpgradeCDBProxyVersionWithContext(context.Background(), request)
}

// UpgradeCDBProxyVersion
// This API is used to upgrade the version of database proxy.
//
// error code that may be returned:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) UpgradeCDBProxyVersionWithContext(ctx context.Context, request *UpgradeCDBProxyVersionRequest) (response *UpgradeCDBProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeCDBProxyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeCDBProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeCDBProxyVersionResponse()
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

// UpgradeDBInstance
// This API is used to upgrade or downgrade a TencentDB instance, which can be a primary instance, disaster recovery instance, or read-only instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    return c.UpgradeDBInstanceWithContext(context.Background(), request)
}

// UpgradeDBInstance
// This API is used to upgrade or downgrade a TencentDB instance, which can be a primary instance, disaster recovery instance, or read-only instance.
//
// error code that may be returned:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
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

// UpgradeDBInstanceEngineVersion
// This API (UpgradeDBInstanceEngineVersion) is used to upgrade the version of a TencentDB instance, which can be a primary instance, disaster recovery instance, or read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    return c.UpgradeDBInstanceEngineVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceEngineVersion
// This API (UpgradeDBInstanceEngineVersion) is used to upgrade the version of a TencentDB instance, which can be a primary instance, disaster recovery instance, or read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstanceEngineVersionWithContext(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceEngineVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstanceEngineVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceEngineVersionResponse()
    err = c.Send(request, response)
    return
}
