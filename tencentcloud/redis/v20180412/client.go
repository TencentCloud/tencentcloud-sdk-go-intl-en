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

package v20180412

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-12"

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


func NewAllocateWanAddressRequest() (request *AllocateWanAddressRequest) {
    request = &AllocateWanAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "AllocateWanAddress")
    
    
    return
}

func NewAllocateWanAddressResponse() (response *AllocateWanAddressResponse) {
    response = &AllocateWanAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateWanAddress
// This API is used to enable public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_ISNOTVPCINSTANCE = "InvalidParameter.IsNotVpcInstance"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
func (c *Client) AllocateWanAddress(request *AllocateWanAddressRequest) (response *AllocateWanAddressResponse, err error) {
    return c.AllocateWanAddressWithContext(context.Background(), request)
}

// AllocateWanAddress
// This API is used to enable public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_ISNOTVPCINSTANCE = "InvalidParameter.IsNotVpcInstance"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
func (c *Client) AllocateWanAddressWithContext(ctx context.Context, request *AllocateWanAddressRequest) (response *AllocateWanAddressResponse, err error) {
    if request == nil {
        request = NewAllocateWanAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateWanAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateWanAddressResponse()
    err = c.Send(request, response)
    return
}

func NewApplyParamsTemplateRequest() (request *ApplyParamsTemplateRequest) {
    request = &ApplyParamsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ApplyParamsTemplate")
    
    
    return
}

func NewApplyParamsTemplateResponse() (response *ApplyParamsTemplateResponse) {
    response = &ApplyParamsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyParamsTemplate
// This API is used to apply a parameter template to instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ApplyParamsTemplate(request *ApplyParamsTemplateRequest) (response *ApplyParamsTemplateResponse, err error) {
    return c.ApplyParamsTemplateWithContext(context.Background(), request)
}

// ApplyParamsTemplate
// This API is used to apply a parameter template to instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ApplyParamsTemplateWithContext(ctx context.Context, request *ApplyParamsTemplateRequest) (response *ApplyParamsTemplateResponse, err error) {
    if request == nil {
        request = NewApplyParamsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyParamsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyParamsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// This API is used to bind a security group to instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// This API is used to bind a security group to instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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

func NewChangeInstanceRoleRequest() (request *ChangeInstanceRoleRequest) {
    request = &ChangeInstanceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ChangeInstanceRole")
    
    
    return
}

func NewChangeInstanceRoleResponse() (response *ChangeInstanceRoleResponse) {
    response = &ChangeInstanceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeInstanceRole
// This API is used to modify the role of an instance in a replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeInstanceRole(request *ChangeInstanceRoleRequest) (response *ChangeInstanceRoleResponse, err error) {
    return c.ChangeInstanceRoleWithContext(context.Background(), request)
}

// ChangeInstanceRole
// This API is used to modify the role of an instance in a replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeInstanceRoleWithContext(ctx context.Context, request *ChangeInstanceRoleRequest) (response *ChangeInstanceRoleResponse, err error) {
    if request == nil {
        request = NewChangeInstanceRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeInstanceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeInstanceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewChangeMasterInstanceRequest() (request *ChangeMasterInstanceRequest) {
    request = &ChangeMasterInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ChangeMasterInstance")
    
    
    return
}

func NewChangeMasterInstanceResponse() (response *ChangeMasterInstanceResponse) {
    response = &ChangeMasterInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeMasterInstance
// This API is used to switch with master instance in a replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeMasterInstance(request *ChangeMasterInstanceRequest) (response *ChangeMasterInstanceResponse, err error) {
    return c.ChangeMasterInstanceWithContext(context.Background(), request)
}

// ChangeMasterInstance
// This API is used to switch with master instance in a replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeMasterInstanceWithContext(ctx context.Context, request *ChangeMasterInstanceRequest) (response *ChangeMasterInstanceResponse, err error) {
    if request == nil {
        request = NewChangeMasterInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeMasterInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeMasterInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewChangeReplicaToMasterRequest() (request *ChangeReplicaToMasterRequest) {
    request = &ChangeReplicaToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ChangeReplicaToMaster")
    
    
    return
}

func NewChangeReplicaToMasterResponse() (response *ChangeReplicaToMasterResponse) {
    response = &ChangeReplicaToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeReplicaToMaster
// This API is used to promote a replica node group of a multi-AZ deployed instance to master node group or a replica node of a single-AZ deployed instance to master node.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) ChangeReplicaToMaster(request *ChangeReplicaToMasterRequest) (response *ChangeReplicaToMasterResponse, err error) {
    return c.ChangeReplicaToMasterWithContext(context.Background(), request)
}

// ChangeReplicaToMaster
// This API is used to promote a replica node group of a multi-AZ deployed instance to master node group or a replica node of a single-AZ deployed instance to master node.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) ChangeReplicaToMasterWithContext(ctx context.Context, request *ChangeReplicaToMasterRequest) (response *ChangeReplicaToMasterResponse, err error) {
    if request == nil {
        request = NewChangeReplicaToMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeReplicaToMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeReplicaToMasterResponse()
    err = c.Send(request, response)
    return
}

func NewCleanUpInstanceRequest() (request *CleanUpInstanceRequest) {
    request = &CleanUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CleanUpInstance")
    
    
    return
}

func NewCleanUpInstanceResponse() (response *CleanUpInstanceResponse) {
    response = &CleanUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CleanUpInstance
// This API is used to eliminate an instance in the recycle bin immediately.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CleanUpInstance(request *CleanUpInstanceRequest) (response *CleanUpInstanceResponse, err error) {
    return c.CleanUpInstanceWithContext(context.Background(), request)
}

// CleanUpInstance
// This API is used to eliminate an instance in the recycle bin immediately.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CleanUpInstanceWithContext(ctx context.Context, request *CleanUpInstanceRequest) (response *CleanUpInstanceResponse, err error) {
    if request == nil {
        request = NewCleanUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CleanUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCleanUpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewClearInstanceRequest() (request *ClearInstanceRequest) {
    request = &ClearInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ClearInstance")
    
    
    return
}

func NewClearInstanceResponse() (response *ClearInstanceResponse) {
    response = &ClearInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearInstance
// This API is used to clear the data of a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ClearInstance(request *ClearInstanceRequest) (response *ClearInstanceResponse, err error) {
    return c.ClearInstanceWithContext(context.Background(), request)
}

// ClearInstance
// This API is used to clear the data of a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ClearInstanceWithContext(ctx context.Context, request *ClearInstanceRequest) (response *ClearInstanceResponse, err error) {
    if request == nil {
        request = NewClearInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSSLRequest() (request *CloseSSLRequest) {
    request = &CloseSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CloseSSL")
    
    
    return
}

func NewCloseSSLResponse() (response *CloseSSLResponse) {
    response = &CloseSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseSSL
// This API is used to disable SSL.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) CloseSSL(request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    return c.CloseSSLWithContext(context.Background(), request)
}

// CloseSSL
// This API is used to disable SSL.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) CloseSSLWithContext(ctx context.Context, request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    if request == nil {
        request = NewCloseSSLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseSSLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceAccountRequest() (request *CreateInstanceAccountRequest) {
    request = &CreateInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateInstanceAccount")
    
    
    return
}

func NewCreateInstanceAccountResponse() (response *CreateInstanceAccountResponse) {
    response = &CreateInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceAccount
// This API is used to create an instance sub-account.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) CreateInstanceAccount(request *CreateInstanceAccountRequest) (response *CreateInstanceAccountResponse, err error) {
    return c.CreateInstanceAccountWithContext(context.Background(), request)
}

// CreateInstanceAccount
// This API is used to create an instance sub-account.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) CreateInstanceAccountWithContext(ctx context.Context, request *CreateInstanceAccountRequest) (response *CreateInstanceAccountResponse, err error) {
    if request == nil {
        request = NewCreateInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstances
// This API is used to create a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"
//  INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"
//  RESOURCEUNAVAILABLE_NOTYPEIDREDISSERVICE = "ResourceUnavailable.NoTypeIdRedisService"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// This API is used to create a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"
//  INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"
//  RESOURCEUNAVAILABLE_NOTYPEIDREDISSERVICE = "ResourceUnavailable.NoTypeIdRedisService"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateParamTemplateRequest() (request *CreateParamTemplateRequest) {
    request = &CreateParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateParamTemplate")
    
    
    return
}

func NewCreateParamTemplateResponse() (response *CreateParamTemplateResponse) {
    response = &CreateParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateParamTemplate
// This API is used to create a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    return c.CreateParamTemplateWithContext(context.Background(), request)
}

// CreateParamTemplate
// This API is used to create a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewDeleteInstanceAccountRequest() (request *DeleteInstanceAccountRequest) {
    request = &DeleteInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DeleteInstanceAccount")
    
    
    return
}

func NewDeleteInstanceAccountResponse() (response *DeleteInstanceAccountResponse) {
    response = &DeleteInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstanceAccount
// This API is used to delete an instance sub-account.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DeleteInstanceAccount(request *DeleteInstanceAccountRequest) (response *DeleteInstanceAccountResponse, err error) {
    return c.DeleteInstanceAccountWithContext(context.Background(), request)
}

// DeleteInstanceAccount
// This API is used to delete an instance sub-account.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DeleteInstanceAccountWithContext(ctx context.Context, request *DeleteInstanceAccountRequest) (response *DeleteInstanceAccountResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParamTemplateRequest() (request *DeleteParamTemplateRequest) {
    request = &DeleteParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DeleteParamTemplate")
    
    
    return
}

func NewDeleteParamTemplateResponse() (response *DeleteParamTemplateResponse) {
    response = &DeleteParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteParamTemplate
// This API is used to delete a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    return c.DeleteParamTemplateWithContext(context.Background(), request)
}

// DeleteParamTemplate
// This API is used to delete a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewDescribeAutoBackupConfigRequest() (request *DescribeAutoBackupConfigRequest) {
    request = &DescribeAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeAutoBackupConfig")
    
    
    return
}

func NewDescribeAutoBackupConfigResponse() (response *DescribeAutoBackupConfigResponse) {
    response = &DescribeAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoBackupConfig
// This API is used to get the backup configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeAutoBackupConfig(request *DescribeAutoBackupConfigRequest) (response *DescribeAutoBackupConfigResponse, err error) {
    return c.DescribeAutoBackupConfigWithContext(context.Background(), request)
}

// DescribeAutoBackupConfig
// This API is used to get the backup configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeAutoBackupConfigWithContext(ctx context.Context, request *DescribeAutoBackupConfigRequest) (response *DescribeAutoBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAutoBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupDownloadRestriction
// This API is used to query the download address for a database backup file in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// This API is used to query the download address for a database backup file in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewDescribeBackupUrlRequest() (request *DescribeBackupUrlRequest) {
    request = &DescribeBackupUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBackupUrl")
    
    
    return
}

func NewDescribeBackupUrlResponse() (response *DescribeBackupUrlResponse) {
    response = &DescribeBackupUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupUrl
// This API is used to query the download address of a backup RDB file.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSINVALID = "ResourceUnavailable.BackupStatusInvalid"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) DescribeBackupUrl(request *DescribeBackupUrlRequest) (response *DescribeBackupUrlResponse, err error) {
    return c.DescribeBackupUrlWithContext(context.Background(), request)
}

// DescribeBackupUrl
// This API is used to query the download address of a backup RDB file.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSINVALID = "ResourceUnavailable.BackupStatusInvalid"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) DescribeBackupUrlWithContext(ctx context.Context, request *DescribeBackupUrlRequest) (response *DescribeBackupUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBackupUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommonDBInstancesRequest() (request *DescribeCommonDBInstancesRequest) {
    request = &DescribeCommonDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeCommonDBInstances")
    
    
    return
}

func NewDescribeCommonDBInstancesResponse() (response *DescribeCommonDBInstancesResponse) {
    response = &DescribeCommonDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCommonDBInstances
// This API has been disused. It was used to query the list of Redis instance information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeCommonDBInstances(request *DescribeCommonDBInstancesRequest) (response *DescribeCommonDBInstancesResponse, err error) {
    return c.DescribeCommonDBInstancesWithContext(context.Background(), request)
}

// DescribeCommonDBInstances
// This API has been disused. It was used to query the list of Redis instance information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeCommonDBInstancesWithContext(ctx context.Context, request *DescribeCommonDBInstancesRequest) (response *DescribeCommonDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCommonDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCommonDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCommonDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeDBSecurityGroups")
    
    
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
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// This API is used to query the security group details of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewDescribeInstanceAccountRequest() (request *DescribeInstanceAccountRequest) {
    request = &DescribeInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceAccount")
    
    
    return
}

func NewDescribeInstanceAccountResponse() (response *DescribeInstanceAccountResponse) {
    response = &DescribeInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceAccount
// This API is used to view instance sub-account information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeInstanceAccount(request *DescribeInstanceAccountRequest) (response *DescribeInstanceAccountResponse, err error) {
    return c.DescribeInstanceAccountWithContext(context.Background(), request)
}

// DescribeInstanceAccount
// This API is used to view instance sub-account information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeInstanceAccountWithContext(ctx context.Context, request *DescribeInstanceAccountRequest) (response *DescribeInstanceAccountResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceBackupsRequest() (request *DescribeInstanceBackupsRequest) {
    request = &DescribeInstanceBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceBackups")
    
    
    return
}

func NewDescribeInstanceBackupsResponse() (response *DescribeInstanceBackupsResponse) {
    response = &DescribeInstanceBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceBackups
// This API is used to query the backup list of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceBackups(request *DescribeInstanceBackupsRequest) (response *DescribeInstanceBackupsResponse, err error) {
    return c.DescribeInstanceBackupsWithContext(context.Background(), request)
}

// DescribeInstanceBackups
// This API is used to query the backup list of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceBackupsWithContext(ctx context.Context, request *DescribeInstanceBackupsRequest) (response *DescribeInstanceBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDTSInfoRequest() (request *DescribeInstanceDTSInfoRequest) {
    request = &DescribeInstanceDTSInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceDTSInfo")
    
    
    return
}

func NewDescribeInstanceDTSInfoResponse() (response *DescribeInstanceDTSInfoResponse) {
    response = &DescribeInstanceDTSInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceDTSInfo
// This API is used to query the DTS task details of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceDTSInfo(request *DescribeInstanceDTSInfoRequest) (response *DescribeInstanceDTSInfoResponse, err error) {
    return c.DescribeInstanceDTSInfoWithContext(context.Background(), request)
}

// DescribeInstanceDTSInfo
// This API is used to query the DTS task details of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceDTSInfoWithContext(ctx context.Context, request *DescribeInstanceDTSInfoRequest) (response *DescribeInstanceDTSInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDTSInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDTSInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDTSInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDealDetailRequest() (request *DescribeInstanceDealDetailRequest) {
    request = &DescribeInstanceDealDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceDealDetail")
    
    
    return
}

func NewDescribeInstanceDealDetailResponse() (response *DescribeInstanceDealDetailResponse) {
    response = &DescribeInstanceDealDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceDealDetail
// This API is used to query the order information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) DescribeInstanceDealDetail(request *DescribeInstanceDealDetailRequest) (response *DescribeInstanceDealDetailResponse, err error) {
    return c.DescribeInstanceDealDetailWithContext(context.Background(), request)
}

// DescribeInstanceDealDetail
// This API is used to query the order information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) DescribeInstanceDealDetailWithContext(ctx context.Context, request *DescribeInstanceDealDetailRequest) (response *DescribeInstanceDealDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDealDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDealDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDealDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeyRequest() (request *DescribeInstanceMonitorBigKeyRequest) {
    request = &DescribeInstanceMonitorBigKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKey")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeyResponse() (response *DescribeInstanceMonitorBigKeyResponse) {
    response = &DescribeInstanceMonitorBigKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorBigKey
// The API for querying big keys of a TencentDB for Redis instance was disused on October 31, 2022. For more information, see [API for Querying Instance Big Key Will Be Disused](https://intl.cloud.tencent.com/document/product/239/81005?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKey(request *DescribeInstanceMonitorBigKeyRequest) (response *DescribeInstanceMonitorBigKeyResponse, err error) {
    return c.DescribeInstanceMonitorBigKeyWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKey
// The API for querying big keys of a TencentDB for Redis instance was disused on October 31, 2022. For more information, see [API for Querying Instance Big Key Will Be Disused](https://intl.cloud.tencent.com/document/product/239/81005?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeyRequest) (response *DescribeInstanceMonitorBigKeyResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeySizeDistRequest() (request *DescribeInstanceMonitorBigKeySizeDistRequest) {
    request = &DescribeInstanceMonitorBigKeySizeDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKeySizeDist")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeySizeDistResponse() (response *DescribeInstanceMonitorBigKeySizeDistResponse) {
    response = &DescribeInstanceMonitorBigKeySizeDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorBigKeySizeDist
// The API for querying big keys of a TencentDB for Redis instance was disused on October 31, 2022. For more information, see [API for Querying Instance Big Key Will Be Disused](https://intl.cloud.tencent.com/document/product/239/81005?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeySizeDist(request *DescribeInstanceMonitorBigKeySizeDistRequest) (response *DescribeInstanceMonitorBigKeySizeDistResponse, err error) {
    return c.DescribeInstanceMonitorBigKeySizeDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKeySizeDist
// The API for querying big keys of a TencentDB for Redis instance was disused on October 31, 2022. For more information, see [API for Querying Instance Big Key Will Be Disused](https://intl.cloud.tencent.com/document/product/239/81005?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeySizeDistWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeySizeDistRequest) (response *DescribeInstanceMonitorBigKeySizeDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeySizeDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKeySizeDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeySizeDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeyTypeDistRequest() (request *DescribeInstanceMonitorBigKeyTypeDistRequest) {
    request = &DescribeInstanceMonitorBigKeyTypeDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKeyTypeDist")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeyTypeDistResponse() (response *DescribeInstanceMonitorBigKeyTypeDistResponse) {
    response = &DescribeInstanceMonitorBigKeyTypeDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorBigKeyTypeDist
// The API for querying big keys of a TencentDB for Redis instance was disused on October 31, 2022. For more information, see [API for Querying Instance Big Key Will Be Disused](https://intl.cloud.tencent.com/document/product/239/81005?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyTypeDist(request *DescribeInstanceMonitorBigKeyTypeDistRequest) (response *DescribeInstanceMonitorBigKeyTypeDistResponse, err error) {
    return c.DescribeInstanceMonitorBigKeyTypeDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKeyTypeDist
// The API for querying big keys of a TencentDB for Redis instance was disused on October 31, 2022. For more information, see [API for Querying Instance Big Key Will Be Disused](https://intl.cloud.tencent.com/document/product/239/81005?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyTypeDistWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeyTypeDistRequest) (response *DescribeInstanceMonitorBigKeyTypeDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeyTypeDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKeyTypeDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeyTypeDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorHotKeyRequest() (request *DescribeInstanceMonitorHotKeyRequest) {
    request = &DescribeInstanceMonitorHotKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorHotKey")
    
    
    return
}

func NewDescribeInstanceMonitorHotKeyResponse() (response *DescribeInstanceMonitorHotKeyResponse) {
    response = &DescribeInstanceMonitorHotKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorHotKey
// This API is used to query the hot key of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorHotKey(request *DescribeInstanceMonitorHotKeyRequest) (response *DescribeInstanceMonitorHotKeyResponse, err error) {
    return c.DescribeInstanceMonitorHotKeyWithContext(context.Background(), request)
}

// DescribeInstanceMonitorHotKey
// This API is used to query the hot key of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorHotKeyWithContext(ctx context.Context, request *DescribeInstanceMonitorHotKeyRequest) (response *DescribeInstanceMonitorHotKeyResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorHotKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorHotKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorHotKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorSIPRequest() (request *DescribeInstanceMonitorSIPRequest) {
    request = &DescribeInstanceMonitorSIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorSIP")
    
    
    return
}

func NewDescribeInstanceMonitorSIPResponse() (response *DescribeInstanceMonitorSIPResponse) {
    response = &DescribeInstanceMonitorSIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorSIP
// This API is used to query the access source information of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorSIP(request *DescribeInstanceMonitorSIPRequest) (response *DescribeInstanceMonitorSIPResponse, err error) {
    return c.DescribeInstanceMonitorSIPWithContext(context.Background(), request)
}

// DescribeInstanceMonitorSIP
// This API is used to query the access source information of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorSIPWithContext(ctx context.Context, request *DescribeInstanceMonitorSIPRequest) (response *DescribeInstanceMonitorSIPResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorSIPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorSIP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorSIPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTookDistRequest() (request *DescribeInstanceMonitorTookDistRequest) {
    request = &DescribeInstanceMonitorTookDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTookDist")
    
    
    return
}

func NewDescribeInstanceMonitorTookDistResponse() (response *DescribeInstanceMonitorTookDistResponse) {
    response = &DescribeInstanceMonitorTookDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorTookDist
// This API is used to query the distribution of instance access duration.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTookDist(request *DescribeInstanceMonitorTookDistRequest) (response *DescribeInstanceMonitorTookDistResponse, err error) {
    return c.DescribeInstanceMonitorTookDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTookDist
// This API is used to query the distribution of instance access duration.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTookDistWithContext(ctx context.Context, request *DescribeInstanceMonitorTookDistRequest) (response *DescribeInstanceMonitorTookDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTookDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTookDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTookDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTopNCmdRequest() (request *DescribeInstanceMonitorTopNCmdRequest) {
    request = &DescribeInstanceMonitorTopNCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTopNCmd")
    
    
    return
}

func NewDescribeInstanceMonitorTopNCmdResponse() (response *DescribeInstanceMonitorTopNCmdResponse) {
    response = &DescribeInstanceMonitorTopNCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorTopNCmd
// This API is used to query an instance access command.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmd(request *DescribeInstanceMonitorTopNCmdRequest) (response *DescribeInstanceMonitorTopNCmdResponse, err error) {
    return c.DescribeInstanceMonitorTopNCmdWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTopNCmd
// This API is used to query an instance access command.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdWithContext(ctx context.Context, request *DescribeInstanceMonitorTopNCmdRequest) (response *DescribeInstanceMonitorTopNCmdResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTopNCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTopNCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTopNCmdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTopNCmdTookRequest() (request *DescribeInstanceMonitorTopNCmdTookRequest) {
    request = &DescribeInstanceMonitorTopNCmdTookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTopNCmdTook")
    
    
    return
}

func NewDescribeInstanceMonitorTopNCmdTookResponse() (response *DescribeInstanceMonitorTopNCmdTookResponse) {
    response = &DescribeInstanceMonitorTopNCmdTookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceMonitorTopNCmdTook
// This API is used to query the instance CPU time.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdTook(request *DescribeInstanceMonitorTopNCmdTookRequest) (response *DescribeInstanceMonitorTopNCmdTookResponse, err error) {
    return c.DescribeInstanceMonitorTopNCmdTookWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTopNCmdTook
// This API is used to query the instance CPU time.
//
// error code that may be returned:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdTookWithContext(ctx context.Context, request *DescribeInstanceMonitorTopNCmdTookRequest) (response *DescribeInstanceMonitorTopNCmdTookResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTopNCmdTookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTopNCmdTook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTopNCmdTookResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodeInfoRequest() (request *DescribeInstanceNodeInfoRequest) {
    request = &DescribeInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceNodeInfo")
    
    
    return
}

func NewDescribeInstanceNodeInfoResponse() (response *DescribeInstanceNodeInfoResponse) {
    response = &DescribeInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceNodeInfo
// This API is used to query instance node information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfo(request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    return c.DescribeInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodeInfo
// This API is used to query instance node information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfoWithContext(ctx context.Context, request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceParamRecords")
    
    
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParamRecords
// This API is used to query the list of parameter modifications.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// This API is used to query the list of parameter modifications.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParams
// This API is used to query the list of instance parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// This API is used to query the list of instance parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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

func NewDescribeInstanceSecurityGroupRequest() (request *DescribeInstanceSecurityGroupRequest) {
    request = &DescribeInstanceSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceSecurityGroup")
    
    
    return
}

func NewDescribeInstanceSecurityGroupResponse() (response *DescribeInstanceSecurityGroupResponse) {
    response = &DescribeInstanceSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceSecurityGroup
// This API is used to query the security group information of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceSecurityGroup(request *DescribeInstanceSecurityGroupRequest) (response *DescribeInstanceSecurityGroupResponse, err error) {
    return c.DescribeInstanceSecurityGroupWithContext(context.Background(), request)
}

// DescribeInstanceSecurityGroup
// This API is used to query the security group information of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceSecurityGroupWithContext(ctx context.Context, request *DescribeInstanceSecurityGroupRequest) (response *DescribeInstanceSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceShardsRequest() (request *DescribeInstanceShardsRequest) {
    request = &DescribeInstanceShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceShards")
    
    
    return
}

func NewDescribeInstanceShardsResponse() (response *DescribeInstanceShardsResponse) {
    response = &DescribeInstanceShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceShards
// This API is used to get the information of Cluster Edition instance shards.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
func (c *Client) DescribeInstanceShards(request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    return c.DescribeInstanceShardsWithContext(context.Background(), request)
}

// DescribeInstanceShards
// This API is used to get the information of Cluster Edition instance shards.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
func (c *Client) DescribeInstanceShardsWithContext(ctx context.Context, request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceShardsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceShards require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceZoneInfoRequest() (request *DescribeInstanceZoneInfoRequest) {
    request = &DescribeInstanceZoneInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceZoneInfo")
    
    
    return
}

func NewDescribeInstanceZoneInfoResponse() (response *DescribeInstanceZoneInfoResponse) {
    response = &DescribeInstanceZoneInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceZoneInfo
// This API is used to query the Redis node details.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceZoneInfo(request *DescribeInstanceZoneInfoRequest) (response *DescribeInstanceZoneInfoResponse, err error) {
    return c.DescribeInstanceZoneInfoWithContext(context.Background(), request)
}

// DescribeInstanceZoneInfo
// This API is used to query the Redis node details.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceZoneInfoWithContext(ctx context.Context, request *DescribeInstanceZoneInfoRequest) (response *DescribeInstanceZoneInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceZoneInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceZoneInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceZoneInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// This API is used to query the list of Redis instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query the list of Redis instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintenanceWindowRequest() (request *DescribeMaintenanceWindowRequest) {
    request = &DescribeMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeMaintenanceWindow")
    
    
    return
}

func NewDescribeMaintenanceWindowResponse() (response *DescribeMaintenanceWindowResponse) {
    response = &DescribeMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaintenanceWindow
// This API is used to query instance maintenance window. The maintenance window specifies a time period during which compatible version upgrade, architecture upgrade, backend maintenance, and other operations can be performed to avoid affecting business.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeMaintenanceWindow(request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    return c.DescribeMaintenanceWindowWithContext(context.Background(), request)
}

// DescribeMaintenanceWindow
// This API is used to query instance maintenance window. The maintenance window specifies a time period during which compatible version upgrade, architecture upgrade, backend maintenance, and other operations can be performed to avoid affecting business.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeMaintenanceWindowWithContext(ctx context.Context, request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplateInfoRequest() (request *DescribeParamTemplateInfoRequest) {
    request = &DescribeParamTemplateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeParamTemplateInfo")
    
    
    return
}

func NewDescribeParamTemplateInfoResponse() (response *DescribeParamTemplateInfoResponse) {
    response = &DescribeParamTemplateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeParamTemplateInfo
// This API is used to query the parameter template details.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeParamTemplateInfo(request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    return c.DescribeParamTemplateInfoWithContext(context.Background(), request)
}

// DescribeParamTemplateInfo
// This API is used to query the parameter template details.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeParamTemplates")
    
    
    return
}

func NewDescribeParamTemplatesResponse() (response *DescribeParamTemplatesResponse) {
    response = &DescribeParamTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeParamTemplates
// This API is used to query the list of parameter templates.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// This API is used to query the list of parameter templates.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
    request = &DescribeProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProductInfo")
    
    
    return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
    response = &DescribeProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductInfo
// This API is used to query purchasable TencentDB for Redis specifications in all regions.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProductInfo(request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    return c.DescribeProductInfoWithContext(context.Background(), request)
}

// DescribeProductInfo
// This API is used to query purchasable TencentDB for Redis specifications in all regions.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProductInfoWithContext(ctx context.Context, request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProductInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupRequest() (request *DescribeProjectSecurityGroupRequest) {
    request = &DescribeProjectSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProjectSecurityGroup")
    
    
    return
}

func NewDescribeProjectSecurityGroupResponse() (response *DescribeProjectSecurityGroupResponse) {
    response = &DescribeProjectSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectSecurityGroup
// This API is used to query the security group information of a project.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  RESOURCEUNAVAILABLE_SECURITYGROUPNOTSUPPORTED = "ResourceUnavailable.SecurityGroupNotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeProjectSecurityGroup(request *DescribeProjectSecurityGroupRequest) (response *DescribeProjectSecurityGroupResponse, err error) {
    return c.DescribeProjectSecurityGroupWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroup
// This API is used to query the security group information of a project.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  RESOURCEUNAVAILABLE_SECURITYGROUPNOTSUPPORTED = "ResourceUnavailable.SecurityGroupNotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeProjectSecurityGroupWithContext(ctx context.Context, request *DescribeProjectSecurityGroupRequest) (response *DescribeProjectSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProjectSecurityGroups")
    
    
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
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_FLOWNOTEXISTS = "FailedOperation.FlowNotExists"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// This API is used to query the security group details of a project.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_FLOWNOTEXISTS = "FailedOperation.FlowNotExists"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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

func NewDescribeProxySlowLogRequest() (request *DescribeProxySlowLogRequest) {
    request = &DescribeProxySlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProxySlowLog")
    
    
    return
}

func NewDescribeProxySlowLogResponse() (response *DescribeProxySlowLogResponse) {
    response = &DescribeProxySlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxySlowLog
// This API is used to query the slow queries of the proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeProxySlowLog(request *DescribeProxySlowLogRequest) (response *DescribeProxySlowLogResponse, err error) {
    return c.DescribeProxySlowLogWithContext(context.Background(), request)
}

// DescribeProxySlowLog
// This API is used to query the slow queries of the proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeProxySlowLogWithContext(ctx context.Context, request *DescribeProxySlowLogRequest) (response *DescribeProxySlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeProxySlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationGroupRequest() (request *DescribeReplicationGroupRequest) {
    request = &DescribeReplicationGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeReplicationGroup")
    
    
    return
}

func NewDescribeReplicationGroupResponse() (response *DescribeReplicationGroupResponse) {
    response = &DescribeReplicationGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReplicationGroup
// This API is used to query a replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
func (c *Client) DescribeReplicationGroup(request *DescribeReplicationGroupRequest) (response *DescribeReplicationGroupResponse, err error) {
    return c.DescribeReplicationGroupWithContext(context.Background(), request)
}

// DescribeReplicationGroup
// This API is used to query a replication group.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
func (c *Client) DescribeReplicationGroupWithContext(ctx context.Context, request *DescribeReplicationGroupRequest) (response *DescribeReplicationGroupResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogRequest() (request *DescribeSlowLogRequest) {
    request = &DescribeSlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeSlowLog")
    
    
    return
}

func NewDescribeSlowLogResponse() (response *DescribeSlowLogResponse) {
    response = &DescribeSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowLog
// This API is used to query the records of slow query.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSlowLog(request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    return c.DescribeSlowLogWithContext(context.Background(), request)
}

// DescribeSlowLog
// This API is used to query the records of slow query.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSlowLogWithContext(ctx context.Context, request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTaskInfo")
    
    
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskInfo
// This API is used to query the task result.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    return c.DescribeTaskInfoWithContext(context.Background(), request)
}

// DescribeTaskInfo
// This API is used to query the task result.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskInfoWithContext(ctx context.Context, request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskListRequest() (request *DescribeTaskListRequest) {
    request = &DescribeTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTaskList")
    
    
    return
}

func NewDescribeTaskListResponse() (response *DescribeTaskListResponse) {
    response = &DescribeTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskList
// This API is used to query the list of tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskList(request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    return c.DescribeTaskListWithContext(context.Background(), request)
}

// DescribeTaskList
// This API is used to query the list of tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskListWithContext(ctx context.Context, request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTendisSlowLogRequest() (request *DescribeTendisSlowLogRequest) {
    request = &DescribeTendisSlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTendisSlowLog")
    
    
    return
}

func NewDescribeTendisSlowLogResponse() (response *DescribeTendisSlowLogResponse) {
    response = &DescribeTendisSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTendisSlowLog
// This API is used to query slow queries of a Tendis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeTendisSlowLog(request *DescribeTendisSlowLogRequest) (response *DescribeTendisSlowLogResponse, err error) {
    return c.DescribeTendisSlowLogWithContext(context.Background(), request)
}

// DescribeTendisSlowLog
// This API is used to query slow queries of a Tendis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeTendisSlowLogWithContext(ctx context.Context, request *DescribeTendisSlowLogRequest) (response *DescribeTendisSlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeTendisSlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTendisSlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTendisSlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPostpaidInstanceRequest() (request *DestroyPostpaidInstanceRequest) {
    request = &DestroyPostpaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DestroyPostpaidInstance")
    
    
    return
}

func NewDestroyPostpaidInstanceResponse() (response *DestroyPostpaidInstanceResponse) {
    response = &DestroyPostpaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyPostpaidInstance
// This API is used to terminate a pay-as-you-go instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DestroyPostpaidInstance(request *DestroyPostpaidInstanceRequest) (response *DestroyPostpaidInstanceResponse, err error) {
    return c.DestroyPostpaidInstanceWithContext(context.Background(), request)
}

// DestroyPostpaidInstance
// This API is used to terminate a pay-as-you-go instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DestroyPostpaidInstanceWithContext(ctx context.Context, request *DestroyPostpaidInstanceRequest) (response *DestroyPostpaidInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPostpaidInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPostpaidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPostpaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPrepaidInstanceRequest() (request *DestroyPrepaidInstanceRequest) {
    request = &DestroyPrepaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DestroyPrepaidInstance")
    
    
    return
}

func NewDestroyPrepaidInstanceResponse() (response *DestroyPrepaidInstanceResponse) {
    response = &DestroyPrepaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyPrepaidInstance
// This API is used to return a monthly subscribed instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEISOLATED = "ResourceUnavailable.InstanceIsolated"
//  RESOURCEUNAVAILABLE_INSTANCENODEAL = "ResourceUnavailable.InstanceNoDeal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DestroyPrepaidInstance(request *DestroyPrepaidInstanceRequest) (response *DestroyPrepaidInstanceResponse, err error) {
    return c.DestroyPrepaidInstanceWithContext(context.Background(), request)
}

// DestroyPrepaidInstance
// This API is used to return a monthly subscribed instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEISOLATED = "ResourceUnavailable.InstanceIsolated"
//  RESOURCEUNAVAILABLE_INSTANCENODEAL = "ResourceUnavailable.InstanceNoDeal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DestroyPrepaidInstanceWithContext(ctx context.Context, request *DestroyPrepaidInstanceRequest) (response *DestroyPrepaidInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPrepaidInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPrepaidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPrepaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisableReplicaReadonlyRequest() (request *DisableReplicaReadonlyRequest) {
    request = &DisableReplicaReadonlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DisableReplicaReadonly")
    
    
    return
}

func NewDisableReplicaReadonlyResponse() (response *DisableReplicaReadonlyResponse) {
    response = &DisableReplicaReadonlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableReplicaReadonly
// This API is used to disable read/write separation.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
func (c *Client) DisableReplicaReadonly(request *DisableReplicaReadonlyRequest) (response *DisableReplicaReadonlyResponse, err error) {
    return c.DisableReplicaReadonlyWithContext(context.Background(), request)
}

// DisableReplicaReadonly
// This API is used to disable read/write separation.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
func (c *Client) DisableReplicaReadonlyWithContext(ctx context.Context, request *DisableReplicaReadonlyRequest) (response *DisableReplicaReadonlyResponse, err error) {
    if request == nil {
        request = NewDisableReplicaReadonlyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableReplicaReadonly require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableReplicaReadonlyResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateSecurityGroups
// This API is used to unbind a security group from instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// This API is used to unbind a security group from instances in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewEnableReplicaReadonlyRequest() (request *EnableReplicaReadonlyRequest) {
    request = &EnableReplicaReadonlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "EnableReplicaReadonly")
    
    
    return
}

func NewEnableReplicaReadonlyResponse() (response *EnableReplicaReadonlyResponse) {
    response = &EnableReplicaReadonlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableReplicaReadonly
// This API is used to enable read/write separation.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) EnableReplicaReadonly(request *EnableReplicaReadonlyRequest) (response *EnableReplicaReadonlyResponse, err error) {
    return c.EnableReplicaReadonlyWithContext(context.Background(), request)
}

// EnableReplicaReadonly
// This API is used to enable read/write separation.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) EnableReplicaReadonlyWithContext(ctx context.Context, request *EnableReplicaReadonlyRequest) (response *EnableReplicaReadonlyResponse, err error) {
    if request == nil {
        request = NewEnableReplicaReadonlyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableReplicaReadonly require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableReplicaReadonlyResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
    request = &InquiryPriceCreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceCreateInstance")
    
    
    return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
    response = &InquiryPriceCreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateInstance
// This API is used to query the price for purchasing an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    return c.InquiryPriceCreateInstanceWithContext(context.Background(), request)
}

// InquiryPriceCreateInstance
// This API is used to query the price for purchasing an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) InquiryPriceCreateInstanceWithContext(ctx context.Context, request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeInstanceRequest() (request *InquiryPriceUpgradeInstanceRequest) {
    request = &InquiryPriceUpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceUpgradeInstance")
    
    
    return
}

func NewInquiryPriceUpgradeInstanceResponse() (response *InquiryPriceUpgradeInstanceResponse) {
    response = &InquiryPriceUpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceUpgradeInstance
// This API is used to query the price for scaling an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) InquiryPriceUpgradeInstance(request *InquiryPriceUpgradeInstanceRequest) (response *InquiryPriceUpgradeInstanceResponse, err error) {
    return c.InquiryPriceUpgradeInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpgradeInstance
// This API is used to query the price for scaling an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) InquiryPriceUpgradeInstanceWithContext(ctx context.Context, request *InquiryPriceUpgradeInstanceRequest) (response *InquiryPriceUpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewKillMasterGroupRequest() (request *KillMasterGroupRequest) {
    request = &KillMasterGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "KillMasterGroup")
    
    
    return
}

func NewKillMasterGroupResponse() (response *KillMasterGroupResponse) {
    response = &KillMasterGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// KillMasterGroup
// This API is used to perform a failure simulation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTOPERATION = "ResourceUnavailable.InstanceNotSupportOperation"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) KillMasterGroup(request *KillMasterGroupRequest) (response *KillMasterGroupResponse, err error) {
    return c.KillMasterGroupWithContext(context.Background(), request)
}

// KillMasterGroup
// This API is used to perform a failure simulation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTOPERATION = "ResourceUnavailable.InstanceNotSupportOperation"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) KillMasterGroupWithContext(ctx context.Context, request *KillMasterGroupRequest) (response *KillMasterGroupResponse, err error) {
    if request == nil {
        request = NewKillMasterGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillMasterGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillMasterGroupResponse()
    err = c.Send(request, response)
    return
}

func NewManualBackupInstanceRequest() (request *ManualBackupInstanceRequest) {
    request = &ManualBackupInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ManualBackupInstance")
    
    
    return
}

func NewManualBackupInstanceResponse() (response *ManualBackupInstanceResponse) {
    response = &ManualBackupInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManualBackupInstance
// This API is used to manually back up a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMMITFLOWERROR = "FailedOperation.CommitFlowError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ManualBackupInstance(request *ManualBackupInstanceRequest) (response *ManualBackupInstanceResponse, err error) {
    return c.ManualBackupInstanceWithContext(context.Background(), request)
}

// ManualBackupInstance
// This API is used to manually back up a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMMITFLOWERROR = "FailedOperation.CommitFlowError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ManualBackupInstanceWithContext(ctx context.Context, request *ManualBackupInstanceRequest) (response *ManualBackupInstanceResponse, err error) {
    if request == nil {
        request = NewManualBackupInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManualBackupInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewManualBackupInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModfiyInstancePasswordRequest() (request *ModfiyInstancePasswordRequest) {
    request = &ModfiyInstancePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModfiyInstancePassword")
    
    
    return
}

func NewModfiyInstancePasswordResponse() (response *ModfiyInstancePasswordResponse) {
    response = &ModfiyInstancePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModfiyInstancePassword
// This API is used to change the Redis password.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModfiyInstancePassword(request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    return c.ModfiyInstancePasswordWithContext(context.Background(), request)
}

// ModfiyInstancePassword
// This API is used to change the Redis password.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModfiyInstancePasswordWithContext(ctx context.Context, request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    if request == nil {
        request = NewModfiyInstancePasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModfiyInstancePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModfiyInstancePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoBackupConfigRequest() (request *ModifyAutoBackupConfigRequest) {
    request = &ModifyAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyAutoBackupConfig")
    
    
    return
}

func NewModifyAutoBackupConfigResponse() (response *ModifyAutoBackupConfigResponse) {
    response = &ModifyAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoBackupConfig
// This API is used to configure automatic backup.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_WEEKDAYSISINVALID = "InvalidParameterValue.WeekDaysIsInvalid"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyAutoBackupConfig(request *ModifyAutoBackupConfigRequest) (response *ModifyAutoBackupConfigResponse, err error) {
    return c.ModifyAutoBackupConfigWithContext(context.Background(), request)
}

// ModifyAutoBackupConfig
// This API is used to configure automatic backup.
//
// error code that may be returned:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_WEEKDAYSISINVALID = "InvalidParameterValue.WeekDaysIsInvalid"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyAutoBackupConfigWithContext(ctx context.Context, request *ModifyAutoBackupConfigRequest) (response *ModifyAutoBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyBackupDownloadRestriction")
    
    
    return
}

func NewModifyBackupDownloadRestrictionResponse() (response *ModifyBackupDownloadRestrictionResponse) {
    response = &ModifyBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupDownloadRestriction
// This API is used to modify the network information and address for downloading a backup file.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// This API is used to modify the network information and address for downloading a backup file.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceSecurityGroups
// This API is used to modify the security groups bound to an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_INSTANCESGOVERLIMITERROR = "InvalidParameter.InstanceSGOverLimitError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// This API is used to modify the security groups bound to an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_INSTANCESGOVERLIMITERROR = "InvalidParameter.InstanceSGOverLimitError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstance
// This API is used to modify instance information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// This API is used to modify instance information.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAccountRequest() (request *ModifyInstanceAccountRequest) {
    request = &ModifyInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceAccount")
    
    
    return
}

func NewModifyInstanceAccountResponse() (response *ModifyInstanceAccountResponse) {
    response = &ModifyInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceAccount
// This API is used to modify an instance sub-account.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceAccount(request *ModifyInstanceAccountRequest) (response *ModifyInstanceAccountResponse, err error) {
    return c.ModifyInstanceAccountWithContext(context.Background(), request)
}

// ModifyInstanceAccount
// This API is used to modify an instance sub-account.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceAccountWithContext(ctx context.Context, request *ModifyInstanceAccountRequest) (response *ModifyInstanceAccountResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamsRequest() (request *ModifyInstanceParamsRequest) {
    request = &ModifyInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceParams")
    
    
    return
}

func NewModifyInstanceParamsResponse() (response *ModifyInstanceParamsResponse) {
    response = &ModifyInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceParams
// This API is used to modify the parameters of TencentDB for Redis instances
//
// error code that may be returned:
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceParams(request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    return c.ModifyInstanceParamsWithContext(context.Background(), request)
}

// ModifyInstanceParams
// This API is used to modify the parameters of TencentDB for Redis instances
//
// error code that may be returned:
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceParamsWithContext(ctx context.Context, request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceReadOnlyRequest() (request *ModifyInstanceReadOnlyRequest) {
    request = &ModifyInstanceReadOnlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceReadOnly")
    
    
    return
}

func NewModifyInstanceReadOnlyResponse() (response *ModifyInstanceReadOnlyResponse) {
    response = &ModifyInstanceReadOnlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceReadOnly
// This API is used to set instance input mode.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyInstanceReadOnly(request *ModifyInstanceReadOnlyRequest) (response *ModifyInstanceReadOnlyResponse, err error) {
    return c.ModifyInstanceReadOnlyWithContext(context.Background(), request)
}

// ModifyInstanceReadOnly
// This API is used to set instance input mode.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyInstanceReadOnlyWithContext(ctx context.Context, request *ModifyInstanceReadOnlyRequest) (response *ModifyInstanceReadOnlyResponse, err error) {
    if request == nil {
        request = NewModifyInstanceReadOnlyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceReadOnly require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceReadOnlyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceWindowRequest() (request *ModifyMaintenanceWindowRequest) {
    request = &ModifyMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyMaintenanceWindow")
    
    
    return
}

func NewModifyMaintenanceWindowResponse() (response *ModifyMaintenanceWindowResponse) {
    response = &ModifyMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMaintenanceWindow
// This API is used to modify the instance maintenance time. The maintenance time specifies a time period during which compatible version upgrade, architecture upgrade, backend maintenance, and other operations can be performed to avoid affecting business. Note: if the compatible version upgrade or architecture upgrade task has been initiated for an instance, its maintenance time cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
func (c *Client) ModifyMaintenanceWindow(request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    return c.ModifyMaintenanceWindowWithContext(context.Background(), request)
}

// ModifyMaintenanceWindow
// This API is used to modify the instance maintenance time. The maintenance time specifies a time period during which compatible version upgrade, architecture upgrade, backend maintenance, and other operations can be performed to avoid affecting business. Note: if the compatible version upgrade or architecture upgrade task has been initiated for an instance, its maintenance time cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
func (c *Client) ModifyMaintenanceWindowWithContext(ctx context.Context, request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewModifyMaintenanceWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkConfigRequest() (request *ModifyNetworkConfigRequest) {
    request = &ModifyNetworkConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyNetworkConfig")
    
    
    return
}

func NewModifyNetworkConfigResponse() (response *ModifyNetworkConfigResponse) {
    response = &ModifyNetworkConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetworkConfig
// This API is used to modify the network configuration of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyNetworkConfig(request *ModifyNetworkConfigRequest) (response *ModifyNetworkConfigResponse, err error) {
    return c.ModifyNetworkConfigWithContext(context.Background(), request)
}

// ModifyNetworkConfig
// This API is used to modify the network configuration of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyNetworkConfigWithContext(ctx context.Context, request *ModifyNetworkConfigRequest) (response *ModifyNetworkConfigResponse, err error) {
    if request == nil {
        request = NewModifyNetworkConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyParamTemplateRequest() (request *ModifyParamTemplateRequest) {
    request = &ModifyParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyParamTemplate")
    
    
    return
}

func NewModifyParamTemplateResponse() (response *ModifyParamTemplateResponse) {
    response = &ModifyParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyParamTemplate
// This API is used to modify a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    return c.ModifyParamTemplateWithContext(context.Background(), request)
}

// ModifyParamTemplate
// This API is used to modify a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewOpenSSLRequest() (request *OpenSSLRequest) {
    request = &OpenSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "OpenSSL")
    
    
    return
}

func NewOpenSSLResponse() (response *OpenSSLResponse) {
    response = &OpenSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenSSL
// This API is used to enable SSL.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) OpenSSL(request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    return c.OpenSSLWithContext(context.Background(), request)
}

// OpenSSL
// This API is used to enable SSL.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) OpenSSLWithContext(ctx context.Context, request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    if request == nil {
        request = NewOpenSSLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenSSLResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseWanAddressRequest() (request *ReleaseWanAddressRequest) {
    request = &ReleaseWanAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ReleaseWanAddress")
    
    
    return
}

func NewReleaseWanAddressResponse() (response *ReleaseWanAddressResponse) {
    response = &ReleaseWanAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseWanAddress
// This API is used to disable public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ReleaseWanAddress(request *ReleaseWanAddressRequest) (response *ReleaseWanAddressResponse, err error) {
    return c.ReleaseWanAddressWithContext(context.Background(), request)
}

// ReleaseWanAddress
// This API is used to disable public network access.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ReleaseWanAddressWithContext(ctx context.Context, request *ReleaseWanAddressRequest) (response *ReleaseWanAddressResponse, err error) {
    if request == nil {
        request = NewReleaseWanAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseWanAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseWanAddressResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "RenewInstance")
    
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewInstance
// This API is used to renew an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    return c.RenewInstanceWithContext(context.Background(), request)
}

// RenewInstance
// This API is used to renew an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
func (c *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ResetPassword")
    
    
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetPassword
// This API is used to reset a password.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

// ResetPassword
// This API is used to reset a password.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreInstanceRequest() (request *RestoreInstanceRequest) {
    request = &RestoreInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "RestoreInstance")
    
    
    return
}

func NewRestoreInstanceResponse() (response *RestoreInstanceResponse) {
    response = &RestoreInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestoreInstance
// This API is used to restore a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPLOCKEDERROR = "ResourceUnavailable.BackupLockedError"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSABNORMAL = "ResourceUnavailable.BackupStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    return c.RestoreInstanceWithContext(context.Background(), request)
}

// RestoreInstance
// This API is used to restore a Redis instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPLOCKEDERROR = "ResourceUnavailable.BackupLockedError"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSABNORMAL = "ResourceUnavailable.BackupStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewStartupInstanceRequest() (request *StartupInstanceRequest) {
    request = &StartupInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "StartupInstance")
    
    
    return
}

func NewStartupInstanceResponse() (response *StartupInstanceResponse) {
    response = &StartupInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartupInstance
// This API is used to deisolate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) StartupInstance(request *StartupInstanceRequest) (response *StartupInstanceResponse, err error) {
    return c.StartupInstanceWithContext(context.Background(), request)
}

// StartupInstance
// This API is used to deisolate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) StartupInstanceWithContext(ctx context.Context, request *StartupInstanceRequest) (response *StartupInstanceResponse, err error) {
    if request == nil {
        request = NewStartupInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartupInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartupInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchInstanceVipRequest() (request *SwitchInstanceVipRequest) {
    request = &SwitchInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "SwitchInstanceVip")
    
    
    return
}

func NewSwitchInstanceVipResponse() (response *SwitchInstanceVipResponse) {
    response = &SwitchInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchInstanceVip
// This API is used to swap the VIPs of instances for instance disaster recovery switch in scenarios where cross-AZ disaster recovery is supported through DTS. After the VIPs of the source and target instances are swapped, the target instance can be written into and the DTS sync task between them will be disconnected.
//
// error code that may be returned:
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) SwitchInstanceVip(request *SwitchInstanceVipRequest) (response *SwitchInstanceVipResponse, err error) {
    return c.SwitchInstanceVipWithContext(context.Background(), request)
}

// SwitchInstanceVip
// This API is used to swap the VIPs of instances for instance disaster recovery switch in scenarios where cross-AZ disaster recovery is supported through DTS. After the VIPs of the source and target instances are swapped, the target instance can be written into and the DTS sync task between them will be disconnected.
//
// error code that may be returned:
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) SwitchInstanceVipWithContext(ctx context.Context, request *SwitchInstanceVipRequest) (response *SwitchInstanceVipResponse, err error) {
    if request == nil {
        request = NewSwitchInstanceVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchInstanceVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchInstanceVipResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchProxyRequest() (request *SwitchProxyRequest) {
    request = &SwitchProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "SwitchProxy")
    
    
    return
}

func NewSwitchProxyResponse() (response *SwitchProxyResponse) {
    response = &SwitchProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchProxy
// This API is used to simulate the failure of a proxy node.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) SwitchProxy(request *SwitchProxyRequest) (response *SwitchProxyResponse, err error) {
    return c.SwitchProxyWithContext(context.Background(), request)
}

// SwitchProxy
// This API is used to simulate the failure of a proxy node.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) SwitchProxyWithContext(ctx context.Context, request *SwitchProxyRequest) (response *SwitchProxyResponse, err error) {
    if request == nil {
        request = NewSwitchProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchProxyResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeInstance
// This API is used to modify the instance configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MEMSIZENOTINRANGE = "InvalidParameterValue.MemSizeNotInRange"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// This API is used to modify the instance configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MEMSIZENOTINRANGE = "InvalidParameterValue.MemSizeNotInRange"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceVersionRequest() (request *UpgradeInstanceVersionRequest) {
    request = &UpgradeInstanceVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeInstanceVersion")
    
    
    return
}

func NewUpgradeInstanceVersionResponse() (response *UpgradeInstanceVersionResponse) {
    response = &UpgradeInstanceVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeInstanceVersion
// This API is used to upgrade compatible instance version (for example, from Redis 2.8 to 4.0), or upgrade instance architecture (for example, from standard architecture to cluster architecture).
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) UpgradeInstanceVersion(request *UpgradeInstanceVersionRequest) (response *UpgradeInstanceVersionResponse, err error) {
    return c.UpgradeInstanceVersionWithContext(context.Background(), request)
}

// UpgradeInstanceVersion
// This API is used to upgrade compatible instance version (for example, from Redis 2.8 to 4.0), or upgrade instance architecture (for example, from standard architecture to cluster architecture).
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) UpgradeInstanceVersionWithContext(ctx context.Context, request *UpgradeInstanceVersionRequest) (response *UpgradeInstanceVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstanceVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeProxyVersionRequest() (request *UpgradeProxyVersionRequest) {
    request = &UpgradeProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeProxyVersion")
    
    
    return
}

func NewUpgradeProxyVersionResponse() (response *UpgradeProxyVersionResponse) {
    response = &UpgradeProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeProxyVersion
// This API is used to upgrade instance proxy version.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) UpgradeProxyVersion(request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    return c.UpgradeProxyVersionWithContext(context.Background(), request)
}

// UpgradeProxyVersion
// This API is used to upgrade instance proxy version.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) UpgradeProxyVersionWithContext(ctx context.Context, request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeProxyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeProxyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeSmallVersionRequest() (request *UpgradeSmallVersionRequest) {
    request = &UpgradeSmallVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeSmallVersion")
    
    
    return
}

func NewUpgradeSmallVersionResponse() (response *UpgradeSmallVersionResponse) {
    response = &UpgradeSmallVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeSmallVersion
// This API is used to upgrade instance minor version.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) UpgradeSmallVersion(request *UpgradeSmallVersionRequest) (response *UpgradeSmallVersionResponse, err error) {
    return c.UpgradeSmallVersionWithContext(context.Background(), request)
}

// UpgradeSmallVersion
// This API is used to upgrade instance minor version.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) UpgradeSmallVersionWithContext(ctx context.Context, request *UpgradeSmallVersionRequest) (response *UpgradeSmallVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeSmallVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeSmallVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeSmallVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeVersionToMultiAvailabilityZonesRequest() (request *UpgradeVersionToMultiAvailabilityZonesRequest) {
    request = &UpgradeVersionToMultiAvailabilityZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeVersionToMultiAvailabilityZones")
    
    
    return
}

func NewUpgradeVersionToMultiAvailabilityZonesResponse() (response *UpgradeVersionToMultiAvailabilityZonesResponse) {
    response = &UpgradeVersionToMultiAvailabilityZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeVersionToMultiAvailabilityZones
// This API is used to upgrade an instance to support multi-AZ deployment.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) UpgradeVersionToMultiAvailabilityZones(request *UpgradeVersionToMultiAvailabilityZonesRequest) (response *UpgradeVersionToMultiAvailabilityZonesResponse, err error) {
    return c.UpgradeVersionToMultiAvailabilityZonesWithContext(context.Background(), request)
}

// UpgradeVersionToMultiAvailabilityZones
// This API is used to upgrade an instance to support multi-AZ deployment.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) UpgradeVersionToMultiAvailabilityZonesWithContext(ctx context.Context, request *UpgradeVersionToMultiAvailabilityZonesRequest) (response *UpgradeVersionToMultiAvailabilityZonesResponse, err error) {
    if request == nil {
        request = NewUpgradeVersionToMultiAvailabilityZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeVersionToMultiAvailabilityZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeVersionToMultiAvailabilityZonesResponse()
    err = c.Send(request, response)
    return
}
