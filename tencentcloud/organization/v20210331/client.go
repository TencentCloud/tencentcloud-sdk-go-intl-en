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

package v20210331

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-03-31"

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


func NewAddExternalSAMLIdPCertificateRequest() (request *AddExternalSAMLIdPCertificateRequest) {
    request = &AddExternalSAMLIdPCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddExternalSAMLIdPCertificate")
    
    
    return
}

func NewAddExternalSAMLIdPCertificateResponse() (response *AddExternalSAMLIdPCertificateResponse) {
    response = &AddExternalSAMLIdPCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddExternalSAMLIdPCertificate
// This API is used to add SAML signing certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_X509CERTIFICATEALREADYEXIST = "FailedOperation.X509CertificateAlreadyExist"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) AddExternalSAMLIdPCertificate(request *AddExternalSAMLIdPCertificateRequest) (response *AddExternalSAMLIdPCertificateResponse, err error) {
    return c.AddExternalSAMLIdPCertificateWithContext(context.Background(), request)
}

// AddExternalSAMLIdPCertificate
// This API is used to add SAML signing certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_X509CERTIFICATEALREADYEXIST = "FailedOperation.X509CertificateAlreadyExist"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) AddExternalSAMLIdPCertificateWithContext(ctx context.Context, request *AddExternalSAMLIdPCertificateRequest) (response *AddExternalSAMLIdPCertificateResponse, err error) {
    if request == nil {
        request = NewAddExternalSAMLIdPCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddExternalSAMLIdPCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddExternalSAMLIdPCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewAddOrganizationNodeRequest() (request *AddOrganizationNodeRequest) {
    request = &AddOrganizationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationNode")
    
    
    return
}

func NewAddOrganizationNodeResponse() (response *AddOrganizationNodeResponse) {
    response = &AddOrganizationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddOrganizationNode
// This API is used to add an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNode(request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    return c.AddOrganizationNodeWithContext(context.Background(), request)
}

// AddOrganizationNode
// This API is used to add an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNodeWithContext(ctx context.Context, request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewAddOrganizationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewAddPermissionPolicyToRoleConfigurationRequest() (request *AddPermissionPolicyToRoleConfigurationRequest) {
    request = &AddPermissionPolicyToRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddPermissionPolicyToRoleConfiguration")
    
    
    return
}

func NewAddPermissionPolicyToRoleConfigurationResponse() (response *AddPermissionPolicyToRoleConfigurationResponse) {
    response = &AddPermissionPolicyToRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddPermissionPolicyToRoleConfiguration
// This API is used to add policies to permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_CUSTOMPOLICYOVERUPPERLIMIT = "FailedOperation.CustomPolicyOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYSTEMPOLICYOVERUPPERLIMIT = "FailedOperation.SystemPolicyOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_BINDPOLICYNAMENOTALLOWED = "InvalidParameter.BindPolicyNameNotAllowed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYNAMEALREADYEXISTS = "InvalidParameter.PolicyNameAlreadyExists"
//  INVALIDPARAMETER_POLICYNAMESIZEOVERUPPERLIMIT = "InvalidParameter.PolicyNameSizeOverUpperLimit"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) AddPermissionPolicyToRoleConfiguration(request *AddPermissionPolicyToRoleConfigurationRequest) (response *AddPermissionPolicyToRoleConfigurationResponse, err error) {
    return c.AddPermissionPolicyToRoleConfigurationWithContext(context.Background(), request)
}

// AddPermissionPolicyToRoleConfiguration
// This API is used to add policies to permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_CUSTOMPOLICYOVERUPPERLIMIT = "FailedOperation.CustomPolicyOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYSTEMPOLICYOVERUPPERLIMIT = "FailedOperation.SystemPolicyOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_BINDPOLICYNAMENOTALLOWED = "InvalidParameter.BindPolicyNameNotAllowed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYNAMEALREADYEXISTS = "InvalidParameter.PolicyNameAlreadyExists"
//  INVALIDPARAMETER_POLICYNAMESIZEOVERUPPERLIMIT = "InvalidParameter.PolicyNameSizeOverUpperLimit"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) AddPermissionPolicyToRoleConfigurationWithContext(ctx context.Context, request *AddPermissionPolicyToRoleConfigurationRequest) (response *AddPermissionPolicyToRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewAddPermissionPolicyToRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddPermissionPolicyToRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddPermissionPolicyToRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
    request = &AddUserToGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddUserToGroup")
    
    
    return
}

func NewAddUserToGroupResponse() (response *AddUserToGroupResponse) {
    response = &AddUserToGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserToGroup
// This API is used to add users to a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//  FAILEDOPERATION_USERADDGROUPCOUNTOVERUPPERLIMIT = "FailedOperation.UserAddGroupCountOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERALREADYEXISTS = "InvalidParameter.GroupUserAlreadyExists"
//  LIMITEXCEEDED_ADDUSERTOGROUPLIMITEXCEEDED = "LimitExceeded.AddUserToGroupLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    return c.AddUserToGroupWithContext(context.Background(), request)
}

// AddUserToGroup
// This API is used to add users to a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//  FAILEDOPERATION_USERADDGROUPCOUNTOVERUPPERLIMIT = "FailedOperation.UserAddGroupCountOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERALREADYEXISTS = "InvalidParameter.GroupUserAlreadyExists"
//  LIMITEXCEEDED_ADDUSERTOGROUPLIMITEXCEEDED = "LimitExceeded.AddUserToGroupLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddUserToGroupWithContext(ctx context.Context, request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserToGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserToGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindOrganizationMemberAuthAccountRequest() (request *BindOrganizationMemberAuthAccountRequest) {
    request = &BindOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberAuthAccount")
    
    
    return
}

func NewBindOrganizationMemberAuthAccountResponse() (response *BindOrganizationMemberAuthAccountResponse) {
    response = &BindOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindOrganizationMemberAuthAccount
// This API is used to bind an organization member to a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    return c.BindOrganizationMemberAuthAccountWithContext(context.Background(), request)
}

// BindOrganizationMemberAuthAccount
// This API is used to bind an organization member to a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationMemberAuthAccountWithContext(ctx context.Context, request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationMemberAuthAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCancelOrganizationMemberAuthAccountRequest() (request *CancelOrganizationMemberAuthAccountRequest) {
    request = &CancelOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccount")
    
    
    return
}

func NewCancelOrganizationMemberAuthAccountResponse() (response *CancelOrganizationMemberAuthAccountResponse) {
    response = &CancelOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelOrganizationMemberAuthAccount
// This API is used to unbind an organization member from a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CancelOrganizationMemberAuthAccount(request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
    return c.CancelOrganizationMemberAuthAccountWithContext(context.Background(), request)
}

// CancelOrganizationMemberAuthAccount
// This API is used to unbind an organization member from a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CancelOrganizationMemberAuthAccountWithContext(ctx context.Context, request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewCancelOrganizationMemberAuthAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewClearExternalSAMLIdentityProviderRequest() (request *ClearExternalSAMLIdentityProviderRequest) {
    request = &ClearExternalSAMLIdentityProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ClearExternalSAMLIdentityProvider")
    
    
    return
}

func NewClearExternalSAMLIdentityProviderResponse() (response *ClearExternalSAMLIdentityProviderResponse) {
    response = &ClearExternalSAMLIdentityProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearExternalSAMLIdentityProvider
// This API is used to clear the SAML identity provider configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SSOSTATUSENABLENOTCLEARIDENTITYPROVIDER = "FailedOperation.SSoStatusEnableNotClearIdentityProvider"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) ClearExternalSAMLIdentityProvider(request *ClearExternalSAMLIdentityProviderRequest) (response *ClearExternalSAMLIdentityProviderResponse, err error) {
    return c.ClearExternalSAMLIdentityProviderWithContext(context.Background(), request)
}

// ClearExternalSAMLIdentityProvider
// This API is used to clear the SAML identity provider configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SSOSTATUSENABLENOTCLEARIDENTITYPROVIDER = "FailedOperation.SSoStatusEnableNotClearIdentityProvider"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) ClearExternalSAMLIdentityProviderWithContext(ctx context.Context, request *ClearExternalSAMLIdentityProviderRequest) (response *ClearExternalSAMLIdentityProviderResponse, err error) {
    if request == nil {
        request = NewClearExternalSAMLIdentityProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateGroup")
    
    
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroup
// This API is used to create user groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPOVERUPPERLIMIT = "FailedOperation.GroupOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNAMEFORMATERROR = "InvalidParameter.GroupNameFormatError"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

// CreateGroup
// This API is used to create user groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPOVERUPPERLIMIT = "FailedOperation.GroupOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNAMEFORMATERROR = "InvalidParameter.GroupNameFormatError"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrgServiceAssignRequest() (request *CreateOrgServiceAssignRequest) {
    request = &CreateOrgServiceAssignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrgServiceAssign")
    
    
    return
}

func NewCreateOrgServiceAssignResponse() (response *CreateOrgServiceAssignResponse) {
    response = &CreateOrgServiceAssignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrgServiceAssign
// This API is used to add a delegated admin of the organization service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrgServiceAssign(request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
    return c.CreateOrgServiceAssignWithContext(context.Background(), request)
}

// CreateOrgServiceAssign
// This API is used to add a delegated admin of the organization service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrgServiceAssignWithContext(ctx context.Context, request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
    if request == nil {
        request = NewCreateOrgServiceAssignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrgServiceAssignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberRequest() (request *CreateOrganizationMemberRequest) {
    request = &CreateOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMember")
    
    
    return
}

func NewCreateOrganizationMemberResponse() (response *CreateOrganizationMemberResponse) {
    response = &CreateOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMember
// This API is used to create an organization member.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) CreateOrganizationMember(request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    return c.CreateOrganizationMemberWithContext(context.Background(), request)
}

// CreateOrganizationMember
// This API is used to create an organization member.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) CreateOrganizationMemberWithContext(ctx context.Context, request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberPolicyRequest() (request *CreateOrganizationMemberPolicyRequest) {
    request = &CreateOrganizationMemberPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberPolicy")
    
    
    return
}

func NewCreateOrganizationMemberPolicyResponse() (response *CreateOrganizationMemberPolicyResponse) {
    response = &CreateOrganizationMemberPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMemberPolicy
// This API is used to create an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberPolicy(request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
    return c.CreateOrganizationMemberPolicyWithContext(context.Background(), request)
}

// CreateOrganizationMemberPolicy
// This API is used to create an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberPolicyWithContext(ctx context.Context, request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMemberPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleAssignmentRequest() (request *CreateRoleAssignmentRequest) {
    request = &CreateRoleAssignmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateRoleAssignment")
    
    
    return
}

func NewCreateRoleAssignmentResponse() (response *CreateRoleAssignmentResponse) {
    response = &CreateRoleAssignmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoleAssignment
// This API is used to grant authorizations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONALREADYEXIST = "FailedOperation.RoleConfigurationAuthorizationAlreadyExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONOVERLIMIT = "FailedOperation.RoleConfigurationAuthorizationOverLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  LIMITEXCEEDED_CREATEROLEASSIGNMENTLIMITEXCEEDED = "LimitExceeded.CreateRoleAssignmentLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateRoleAssignment(request *CreateRoleAssignmentRequest) (response *CreateRoleAssignmentResponse, err error) {
    return c.CreateRoleAssignmentWithContext(context.Background(), request)
}

// CreateRoleAssignment
// This API is used to grant authorizations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONALREADYEXIST = "FailedOperation.RoleConfigurationAuthorizationAlreadyExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONOVERLIMIT = "FailedOperation.RoleConfigurationAuthorizationOverLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  LIMITEXCEEDED_CREATEROLEASSIGNMENTLIMITEXCEEDED = "LimitExceeded.CreateRoleAssignmentLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateRoleAssignmentWithContext(ctx context.Context, request *CreateRoleAssignmentRequest) (response *CreateRoleAssignmentResponse, err error) {
    if request == nil {
        request = NewCreateRoleAssignmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoleAssignment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleAssignmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleConfigurationRequest() (request *CreateRoleConfigurationRequest) {
    request = &CreateRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateRoleConfiguration")
    
    
    return
}

func NewCreateRoleConfigurationResponse() (response *CreateRoleConfigurationResponse) {
    response = &CreateRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoleConfiguration
// This API is used to create permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFIGURATIONOVERUPPERLIMIT = "FailedOperation.ConfigurationOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_CONFIGURATIONNAMEALREADYEXISTS = "InvalidParameter.ConfigurationNameAlreadyExists"
//  INVALIDPARAMETER_CONFIGURATIONNAMEFORMATERROR = "InvalidParameter.ConfigurationNameFormatError"
func (c *Client) CreateRoleConfiguration(request *CreateRoleConfigurationRequest) (response *CreateRoleConfigurationResponse, err error) {
    return c.CreateRoleConfigurationWithContext(context.Background(), request)
}

// CreateRoleConfiguration
// This API is used to create permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFIGURATIONOVERUPPERLIMIT = "FailedOperation.ConfigurationOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_CONFIGURATIONNAMEALREADYEXISTS = "InvalidParameter.ConfigurationNameAlreadyExists"
//  INVALIDPARAMETER_CONFIGURATIONNAMEFORMATERROR = "InvalidParameter.ConfigurationNameFormatError"
func (c *Client) CreateRoleConfigurationWithContext(ctx context.Context, request *CreateRoleConfigurationRequest) (response *CreateRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// This API is used to create a user.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
//  LIMITEXCEEDED_CREATEUSERLIMITEXCEEDED = "LimitExceeded.CreateUserLimitExceeded"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to create a user.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
//  LIMITEXCEEDED_CREATEUSERLIMITEXCEEDED = "LimitExceeded.CreateUserLimitExceeded"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserSyncProvisioningRequest() (request *CreateUserSyncProvisioningRequest) {
    request = &CreateUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateUserSyncProvisioning")
    
    
    return
}

func NewCreateUserSyncProvisioningResponse() (response *CreateUserSyncProvisioningResponse) {
    response = &CreateUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserSyncProvisioning
// This API is used to create sub-user synchronization tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USERPROVISIONINGALREADYEXISTS = "FailedOperation.UserProvisioningAlreadyExists"
//  FAILEDOPERATION_USERPROVISIONINGOVERLIMIT = "FailedOperation.UserProvisioningOverLimit"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  LIMITEXCEEDED_CREATEUSERSYNCPROVISIONINGLIMITEXCEEDED = "LimitExceeded.CreateUserSyncProvisioningLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateUserSyncProvisioning(request *CreateUserSyncProvisioningRequest) (response *CreateUserSyncProvisioningResponse, err error) {
    return c.CreateUserSyncProvisioningWithContext(context.Background(), request)
}

// CreateUserSyncProvisioning
// This API is used to create sub-user synchronization tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USERPROVISIONINGALREADYEXISTS = "FailedOperation.UserProvisioningAlreadyExists"
//  FAILEDOPERATION_USERPROVISIONINGOVERLIMIT = "FailedOperation.UserProvisioningOverLimit"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  LIMITEXCEEDED_CREATEUSERSYNCPROVISIONINGLIMITEXCEEDED = "LimitExceeded.CreateUserSyncProvisioningLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateUserSyncProvisioningWithContext(ctx context.Context, request *CreateUserSyncProvisioningRequest) (response *CreateUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewCreateUserSyncProvisioningRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserSyncProvisioningResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// This API is used to delete user groups.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEGROUPNOTALLOWEXISTUSER = "FailedOperation.DeleteGroupNotAllowExistUser"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// This API is used to delete user groups.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEGROUPNOTALLOWEXISTUSER = "FailedOperation.DeleteGroupNotAllowExistUser"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrgServiceAssignRequest() (request *DeleteOrgServiceAssignRequest) {
    request = &DeleteOrgServiceAssignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrgServiceAssign")
    
    
    return
}

func NewDeleteOrgServiceAssignResponse() (response *DeleteOrgServiceAssignResponse) {
    response = &DeleteOrgServiceAssignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrgServiceAssign
// This API is used to delete a delegated admin of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DeleteOrgServiceAssign(request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
    return c.DeleteOrgServiceAssignWithContext(context.Background(), request)
}

// DeleteOrgServiceAssign
// This API is used to delete a delegated admin of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DeleteOrgServiceAssignWithContext(ctx context.Context, request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
    if request == nil {
        request = NewDeleteOrgServiceAssignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrgServiceAssignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMembersRequest() (request *DeleteOrganizationMembersRequest) {
    request = &DeleteOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembers")
    
    
    return
}

func NewDeleteOrganizationMembersResponse() (response *DeleteOrganizationMembersResponse) {
    response = &DeleteOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMembers
// This API is used to remove a member account from the organization, rather than delete the account.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    return c.DeleteOrganizationMembersWithContext(context.Background(), request)
}

// DeleteOrganizationMembers
// This API is used to remove a member account from the organization, rather than delete the account.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembersWithContext(ctx context.Context, request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationNodesRequest() (request *DeleteOrganizationNodesRequest) {
    request = &DeleteOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodes")
    
    
    return
}

func NewDeleteOrganizationNodesResponse() (response *DeleteOrganizationNodesResponse) {
    response = &DeleteOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationNodes
// This API is used to batch delete organization nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodes(request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    return c.DeleteOrganizationNodesWithContext(context.Background(), request)
}

// DeleteOrganizationNodes
// This API is used to batch delete organization nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodesWithContext(ctx context.Context, request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleAssignmentRequest() (request *DeleteRoleAssignmentRequest) {
    request = &DeleteRoleAssignmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteRoleAssignment")
    
    
    return
}

func NewDeleteRoleAssignmentResponse() (response *DeleteRoleAssignmentResponse) {
    response = &DeleteRoleAssignmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoleAssignment
// This API is used to remove authorizations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONAUTHORIZATIONNOTFOUND = "ResourceNotFound.RoleConfigurationAuthorizationNotFound"
func (c *Client) DeleteRoleAssignment(request *DeleteRoleAssignmentRequest) (response *DeleteRoleAssignmentResponse, err error) {
    return c.DeleteRoleAssignmentWithContext(context.Background(), request)
}

// DeleteRoleAssignment
// This API is used to remove authorizations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONAUTHORIZATIONNOTFOUND = "ResourceNotFound.RoleConfigurationAuthorizationNotFound"
func (c *Client) DeleteRoleAssignmentWithContext(ctx context.Context, request *DeleteRoleAssignmentRequest) (response *DeleteRoleAssignmentResponse, err error) {
    if request == nil {
        request = NewDeleteRoleAssignmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoleAssignment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleAssignmentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleConfigurationRequest() (request *DeleteRoleConfigurationRequest) {
    request = &DeleteRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteRoleConfiguration")
    
    
    return
}

func NewDeleteRoleConfigurationResponse() (response *DeleteRoleConfigurationResponse) {
    response = &DeleteRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoleConfiguration
// This API is used to delete the permission configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGALREADYDEPLOYED = "FailedOperation.RoleConfigurationProvisioningAlreadyDeployed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYALREADYEXIST = "InvalidParameter.RolePolicyAlreadyExist"
func (c *Client) DeleteRoleConfiguration(request *DeleteRoleConfigurationRequest) (response *DeleteRoleConfigurationResponse, err error) {
    return c.DeleteRoleConfigurationWithContext(context.Background(), request)
}

// DeleteRoleConfiguration
// This API is used to delete the permission configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGALREADYDEPLOYED = "FailedOperation.RoleConfigurationProvisioningAlreadyDeployed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYALREADYEXIST = "InvalidParameter.RolePolicyAlreadyExist"
func (c *Client) DeleteRoleConfigurationWithContext(ctx context.Context, request *DeleteRoleConfigurationRequest) (response *DeleteRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// This API is used to delete a user.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTDELETE = "FailedOperation.SynchronizedUserNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_USERALREADYEXISTSGROUP = "InvalidParameter.UserAlreadyExistsGroup"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete a user.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTDELETE = "FailedOperation.SynchronizedUserNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_USERALREADYEXISTSGROUP = "InvalidParameter.UserAlreadyExistsGroup"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserSyncProvisioningRequest() (request *DeleteUserSyncProvisioningRequest) {
    request = &DeleteUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteUserSyncProvisioning")
    
    
    return
}

func NewDeleteUserSyncProvisioningResponse() (response *DeleteUserSyncProvisioningResponse) {
    response = &DeleteUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserSyncProvisioning
// This API is used to delete sub-user synchronization tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USERPROVISIONINGFAILED = "FailedOperation.UserProvisioningFailed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) DeleteUserSyncProvisioning(request *DeleteUserSyncProvisioningRequest) (response *DeleteUserSyncProvisioningResponse, err error) {
    return c.DeleteUserSyncProvisioningWithContext(context.Background(), request)
}

// DeleteUserSyncProvisioning
// This API is used to delete sub-user synchronization tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USERPROVISIONINGFAILED = "FailedOperation.UserProvisioningFailed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) DeleteUserSyncProvisioningWithContext(ctx context.Context, request *DeleteUserSyncProvisioningRequest) (response *DeleteUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewDeleteUserSyncProvisioningRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserSyncProvisioningResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentityCenterRequest() (request *DescribeIdentityCenterRequest) {
    request = &DescribeIdentityCenterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeIdentityCenter")
    
    
    return
}

func NewDescribeIdentityCenterResponse() (response *DescribeIdentityCenterResponse) {
    response = &DescribeIdentityCenterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdentityCenter
// This API is used to obtain the CAM Identity Center service information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
func (c *Client) DescribeIdentityCenter(request *DescribeIdentityCenterRequest) (response *DescribeIdentityCenterResponse, err error) {
    return c.DescribeIdentityCenterWithContext(context.Background(), request)
}

// DescribeIdentityCenter
// This API is used to obtain the CAM Identity Center service information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
func (c *Client) DescribeIdentityCenterWithContext(ctx context.Context, request *DescribeIdentityCenterRequest) (response *DescribeIdentityCenterResponse, err error) {
    if request == nil {
        request = NewDescribeIdentityCenterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentityCenter require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentityCenterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
    request = &DescribeOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganization")
    
    
    return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
    response = &DescribeOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganization
// This API is used to get the organization information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    return c.DescribeOrganizationWithContext(context.Background(), request)
}

// DescribeOrganization
// This API is used to get the organization information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationWithContext(ctx context.Context, request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberAuthAccountsRequest() (request *DescribeOrganizationMemberAuthAccountsRequest) {
    request = &DescribeOrganizationMemberAuthAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
    
    
    return
}

func NewDescribeOrganizationMemberAuthAccountsResponse() (response *DescribeOrganizationMemberAuthAccountsResponse) {
    response = &DescribeOrganizationMemberAuthAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberAuthAccounts
// This API is used to get the list of sub-accounts bound to an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthAccounts(request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
    return c.DescribeOrganizationMemberAuthAccountsWithContext(context.Background(), request)
}

// DescribeOrganizationMemberAuthAccounts
// This API is used to get the list of sub-accounts bound to an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthAccountsWithContext(ctx context.Context, request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberAuthAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberAuthIdentitiesRequest() (request *DescribeOrganizationMemberAuthIdentitiesRequest) {
    request = &DescribeOrganizationMemberAuthIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthIdentities")
    
    
    return
}

func NewDescribeOrganizationMemberAuthIdentitiesResponse() (response *DescribeOrganizationMemberAuthIdentitiesResponse) {
    response = &DescribeOrganizationMemberAuthIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberAuthIdentities
// This API is used to obtain the list of organization member access authorization.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentities(request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
    return c.DescribeOrganizationMemberAuthIdentitiesWithContext(context.Background(), request)
}

// DescribeOrganizationMemberAuthIdentities
// This API is used to obtain the list of organization member access authorization.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentitiesWithContext(ctx context.Context, request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberAuthIdentitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthIdentities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberPoliciesRequest() (request *DescribeOrganizationMemberPoliciesRequest) {
    request = &DescribeOrganizationMemberPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberPolicies")
    
    
    return
}

func NewDescribeOrganizationMemberPoliciesResponse() (response *DescribeOrganizationMemberPoliciesResponse) {
    response = &DescribeOrganizationMemberPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberPolicies
// This API is used to get the list of authorization policies of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberPolicies(request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
    return c.DescribeOrganizationMemberPoliciesWithContext(context.Background(), request)
}

// DescribeOrganizationMemberPolicies
// This API is used to get the list of authorization policies of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberPoliciesWithContext(ctx context.Context, request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
    request = &DescribeOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembers")
    
    
    return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
    response = &DescribeOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMembers
// This API is used to get the list of organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    return c.DescribeOrganizationMembersWithContext(context.Background(), request)
}

// DescribeOrganizationMembers
// This API is used to get the list of organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DescribeOrganizationMembersWithContext(ctx context.Context, request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationNodesRequest() (request *DescribeOrganizationNodesRequest) {
    request = &DescribeOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodes")
    
    
    return
}

func NewDescribeOrganizationNodesResponse() (response *DescribeOrganizationNodesResponse) {
    response = &DescribeOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationNodes
// This API is used to get the list of organization nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationNodes(request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
    return c.DescribeOrganizationNodesWithContext(context.Background(), request)
}

// DescribeOrganizationNodes
// This API is used to get the list of organization nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationNodesWithContext(ctx context.Context, request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDismantleRoleConfigurationRequest() (request *DismantleRoleConfigurationRequest) {
    request = &DismantleRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DismantleRoleConfiguration")
    
    
    return
}

func NewDismantleRoleConfigurationResponse() (response *DismantleRoleConfigurationResponse) {
    response = &DismantleRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DismantleRoleConfiguration
// This API is used to undeploy permission configurations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGSTATUSERROR = "FailedOperation.RoleConfigurationProvisioningStatusError"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONPROVISIONINGNOTFOUND = "ResourceNotFound.RoleConfigurationProvisioningNotFound"
func (c *Client) DismantleRoleConfiguration(request *DismantleRoleConfigurationRequest) (response *DismantleRoleConfigurationResponse, err error) {
    return c.DismantleRoleConfigurationWithContext(context.Background(), request)
}

// DismantleRoleConfiguration
// This API is used to undeploy permission configurations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGSTATUSERROR = "FailedOperation.RoleConfigurationProvisioningStatusError"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONPROVISIONINGNOTFOUND = "ResourceNotFound.RoleConfigurationProvisioningNotFound"
func (c *Client) DismantleRoleConfigurationWithContext(ctx context.Context, request *DismantleRoleConfigurationRequest) (response *DismantleRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewDismantleRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismantleRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismantleRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewGetExternalSAMLIdentityProviderRequest() (request *GetExternalSAMLIdentityProviderRequest) {
    request = &GetExternalSAMLIdentityProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetExternalSAMLIdentityProvider")
    
    
    return
}

func NewGetExternalSAMLIdentityProviderResponse() (response *GetExternalSAMLIdentityProviderResponse) {
    response = &GetExternalSAMLIdentityProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetExternalSAMLIdentityProvider
// This API is used to query the SAML identity provider configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) GetExternalSAMLIdentityProvider(request *GetExternalSAMLIdentityProviderRequest) (response *GetExternalSAMLIdentityProviderResponse, err error) {
    return c.GetExternalSAMLIdentityProviderWithContext(context.Background(), request)
}

// GetExternalSAMLIdentityProvider
// This API is used to query the SAML identity provider configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) GetExternalSAMLIdentityProviderWithContext(ctx context.Context, request *GetExternalSAMLIdentityProviderRequest) (response *GetExternalSAMLIdentityProviderResponse, err error) {
    if request == nil {
        request = NewGetExternalSAMLIdentityProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupRequest() (request *GetGroupRequest) {
    request = &GetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetGroup")
    
    
    return
}

func NewGetGroupResponse() (response *GetGroupResponse) {
    response = &GetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroup
// This API is used to query the user group information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
    return c.GetGroupWithContext(context.Background(), request)
}

// GetGroup
// This API is used to query the user group information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest) (response *GetGroupResponse, err error) {
    if request == nil {
        request = NewGetGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewGetProvisioningTaskStatusRequest() (request *GetProvisioningTaskStatusRequest) {
    request = &GetProvisioningTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetProvisioningTaskStatus")
    
    
    return
}

func NewGetProvisioningTaskStatusResponse() (response *GetProvisioningTaskStatusResponse) {
    response = &GetProvisioningTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetProvisioningTaskStatus
// This API is used to query the status of async tasks of user synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGTASKNOTFOUND = "ResourceNotFound.UserProvisioningTaskNotFound"
func (c *Client) GetProvisioningTaskStatus(request *GetProvisioningTaskStatusRequest) (response *GetProvisioningTaskStatusResponse, err error) {
    return c.GetProvisioningTaskStatusWithContext(context.Background(), request)
}

// GetProvisioningTaskStatus
// This API is used to query the status of async tasks of user synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGTASKNOTFOUND = "ResourceNotFound.UserProvisioningTaskNotFound"
func (c *Client) GetProvisioningTaskStatusWithContext(ctx context.Context, request *GetProvisioningTaskStatusRequest) (response *GetProvisioningTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetProvisioningTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProvisioningTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProvisioningTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoleConfigurationRequest() (request *GetRoleConfigurationRequest) {
    request = &GetRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetRoleConfiguration")
    
    
    return
}

func NewGetRoleConfigurationResponse() (response *GetRoleConfigurationResponse) {
    response = &GetRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRoleConfiguration
// This API is used to query the permission configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) GetRoleConfiguration(request *GetRoleConfigurationRequest) (response *GetRoleConfigurationResponse, err error) {
    return c.GetRoleConfigurationWithContext(context.Background(), request)
}

// GetRoleConfiguration
// This API is used to query the permission configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) GetRoleConfigurationWithContext(ctx context.Context, request *GetRoleConfigurationRequest) (response *GetRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewGetRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskStatusRequest() (request *GetTaskStatusRequest) {
    request = &GetTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetTaskStatus")
    
    
    return
}

func NewGetTaskStatusResponse() (response *GetTaskStatusResponse) {
    response = &GetTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskStatus
// This API is used to query the status of async tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONTASKNOTFOUND = "ResourceNotFound.RoleConfigurationTaskNotFound"
func (c *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    return c.GetTaskStatusWithContext(context.Background(), request)
}

// GetTaskStatus
// This API is used to query the status of async tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONTASKNOTFOUND = "ResourceNotFound.RoleConfigurationTaskNotFound"
func (c *Client) GetTaskStatusWithContext(ctx context.Context, request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserRequest() (request *GetUserRequest) {
    request = &GetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetUser")
    
    
    return
}

func NewGetUserResponse() (response *GetUserResponse) {
    response = &GetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUser
// This API is used to query the user information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
    return c.GetUserWithContext(context.Background(), request)
}

// GetUser
// This API is used to query the user information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest) (response *GetUserResponse, err error) {
    if request == nil {
        request = NewGetUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserSyncProvisioningRequest() (request *GetUserSyncProvisioningRequest) {
    request = &GetUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetUserSyncProvisioning")
    
    
    return
}

func NewGetUserSyncProvisioningResponse() (response *GetUserSyncProvisioningResponse) {
    response = &GetUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUserSyncProvisioning
// This API is used to query the CAM user synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) GetUserSyncProvisioning(request *GetUserSyncProvisioningRequest) (response *GetUserSyncProvisioningResponse, err error) {
    return c.GetUserSyncProvisioningWithContext(context.Background(), request)
}

// GetUserSyncProvisioning
// This API is used to query the CAM user synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) GetUserSyncProvisioningWithContext(ctx context.Context, request *GetUserSyncProvisioningRequest) (response *GetUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewGetUserSyncProvisioningRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserSyncProvisioningResponse()
    err = c.Send(request, response)
    return
}

func NewGetZoneSAMLServiceProviderInfoRequest() (request *GetZoneSAMLServiceProviderInfoRequest) {
    request = &GetZoneSAMLServiceProviderInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetZoneSAMLServiceProviderInfo")
    
    
    return
}

func NewGetZoneSAMLServiceProviderInfoResponse() (response *GetZoneSAMLServiceProviderInfoResponse) {
    response = &GetZoneSAMLServiceProviderInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetZoneSAMLServiceProviderInfo
// This API is used to query the SAML service provider configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SAMLMETADATADOCUMENTGETFAILED = "FailedOperation.SAMLMetadataDocumentGetFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLSERVICEPROVIDERNOTFOUND = "ResourceNotFound.SAMLServiceProviderNotFound"
func (c *Client) GetZoneSAMLServiceProviderInfo(request *GetZoneSAMLServiceProviderInfoRequest) (response *GetZoneSAMLServiceProviderInfoResponse, err error) {
    return c.GetZoneSAMLServiceProviderInfoWithContext(context.Background(), request)
}

// GetZoneSAMLServiceProviderInfo
// This API is used to query the SAML service provider configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SAMLMETADATADOCUMENTGETFAILED = "FailedOperation.SAMLMetadataDocumentGetFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLSERVICEPROVIDERNOTFOUND = "ResourceNotFound.SAMLServiceProviderNotFound"
func (c *Client) GetZoneSAMLServiceProviderInfoWithContext(ctx context.Context, request *GetZoneSAMLServiceProviderInfoRequest) (response *GetZoneSAMLServiceProviderInfoResponse, err error) {
    if request == nil {
        request = NewGetZoneSAMLServiceProviderInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetZoneSAMLServiceProviderInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetZoneSAMLServiceProviderInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetZoneStatisticsRequest() (request *GetZoneStatisticsRequest) {
    request = &GetZoneStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetZoneStatistics")
    
    
    return
}

func NewGetZoneStatisticsResponse() (response *GetZoneStatisticsResponse) {
    response = &GetZoneStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetZoneStatistics
// This API is used to query space statistics.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetZoneStatistics(request *GetZoneStatisticsRequest) (response *GetZoneStatisticsResponse, err error) {
    return c.GetZoneStatisticsWithContext(context.Background(), request)
}

// GetZoneStatistics
// This API is used to query space statistics.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetZoneStatisticsWithContext(ctx context.Context, request *GetZoneStatisticsRequest) (response *GetZoneStatisticsResponse, err error) {
    if request == nil {
        request = NewGetZoneStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetZoneStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetZoneStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewListExternalSAMLIdPCertificatesRequest() (request *ListExternalSAMLIdPCertificatesRequest) {
    request = &ListExternalSAMLIdPCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListExternalSAMLIdPCertificates")
    
    
    return
}

func NewListExternalSAMLIdPCertificatesResponse() (response *ListExternalSAMLIdPCertificatesResponse) {
    response = &ListExternalSAMLIdPCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListExternalSAMLIdPCertificates
// This API is used to query the SAML signing certificate list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
func (c *Client) ListExternalSAMLIdPCertificates(request *ListExternalSAMLIdPCertificatesRequest) (response *ListExternalSAMLIdPCertificatesResponse, err error) {
    return c.ListExternalSAMLIdPCertificatesWithContext(context.Background(), request)
}

// ListExternalSAMLIdPCertificates
// This API is used to query the SAML signing certificate list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
func (c *Client) ListExternalSAMLIdPCertificatesWithContext(ctx context.Context, request *ListExternalSAMLIdPCertificatesRequest) (response *ListExternalSAMLIdPCertificatesResponse, err error) {
    if request == nil {
        request = NewListExternalSAMLIdPCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListExternalSAMLIdPCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListExternalSAMLIdPCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupMembersRequest() (request *ListGroupMembersRequest) {
    request = &ListGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListGroupMembers")
    
    
    return
}

func NewListGroupMembersResponse() (response *ListGroupMembersResponse) {
    response = &ListGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListGroupMembers
// This API is used to query the user list of the user group.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) ListGroupMembers(request *ListGroupMembersRequest) (response *ListGroupMembersResponse, err error) {
    return c.ListGroupMembersWithContext(context.Background(), request)
}

// ListGroupMembers
// This API is used to query the user list of the user group.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) ListGroupMembersWithContext(ctx context.Context, request *ListGroupMembersRequest) (response *ListGroupMembersResponse, err error) {
    if request == nil {
        request = NewListGroupMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupsRequest() (request *ListGroupsRequest) {
    request = &ListGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListGroups")
    
    
    return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
    response = &ListGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListGroups
// This API is used to query the user group list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    return c.ListGroupsWithContext(context.Background(), request)
}

// ListGroups
// This API is used to query the user group list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    if request == nil {
        request = NewListGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListJoinedGroupsForUserRequest() (request *ListJoinedGroupsForUserRequest) {
    request = &ListJoinedGroupsForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListJoinedGroupsForUser")
    
    
    return
}

func NewListJoinedGroupsForUserResponse() (response *ListJoinedGroupsForUserResponse) {
    response = &ListJoinedGroupsForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListJoinedGroupsForUser
// This API is used to query the user group joined by users.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListJoinedGroupsForUser(request *ListJoinedGroupsForUserRequest) (response *ListJoinedGroupsForUserResponse, err error) {
    return c.ListJoinedGroupsForUserWithContext(context.Background(), request)
}

// ListJoinedGroupsForUser
// This API is used to query the user group joined by users.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListJoinedGroupsForUserWithContext(ctx context.Context, request *ListJoinedGroupsForUserRequest) (response *ListJoinedGroupsForUserResponse, err error) {
    if request == nil {
        request = NewListJoinedGroupsForUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListJoinedGroupsForUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListJoinedGroupsForUserResponse()
    err = c.Send(request, response)
    return
}

func NewListOrgServiceAssignMemberRequest() (request *ListOrgServiceAssignMemberRequest) {
    request = &ListOrgServiceAssignMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrgServiceAssignMember")
    
    
    return
}

func NewListOrgServiceAssignMemberResponse() (response *ListOrgServiceAssignMemberResponse) {
    response = &ListOrgServiceAssignMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrgServiceAssignMember
// This API is used to obtain the list of delegated admins of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrgServiceAssignMember(request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
    return c.ListOrgServiceAssignMemberWithContext(context.Background(), request)
}

// ListOrgServiceAssignMember
// This API is used to obtain the list of delegated admins of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrgServiceAssignMemberWithContext(ctx context.Context, request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
    if request == nil {
        request = NewListOrgServiceAssignMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrgServiceAssignMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrgServiceAssignMemberResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationIdentityRequest() (request *ListOrganizationIdentityRequest) {
    request = &ListOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationIdentity")
    
    
    return
}

func NewListOrganizationIdentityResponse() (response *ListOrganizationIdentityResponse) {
    response = &ListOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationIdentity
// This API is used to get the list of access identities of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentity(request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
    return c.ListOrganizationIdentityWithContext(context.Background(), request)
}

// ListOrganizationIdentity
// This API is used to get the list of access identities of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentityWithContext(ctx context.Context, request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewListOrganizationIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationServiceRequest() (request *ListOrganizationServiceRequest) {
    request = &ListOrganizationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationService")
    
    
    return
}

func NewListOrganizationServiceResponse() (response *ListOrganizationServiceResponse) {
    response = &ListOrganizationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationService
// This API is used to obtain the list of organization service settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationService(request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
    return c.ListOrganizationServiceWithContext(context.Background(), request)
}

// ListOrganizationService
// This API is used to obtain the list of organization service settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationServiceWithContext(ctx context.Context, request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
    if request == nil {
        request = NewListOrganizationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewListPermissionPoliciesInRoleConfigurationRequest() (request *ListPermissionPoliciesInRoleConfigurationRequest) {
    request = &ListPermissionPoliciesInRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListPermissionPoliciesInRoleConfiguration")
    
    
    return
}

func NewListPermissionPoliciesInRoleConfigurationResponse() (response *ListPermissionPoliciesInRoleConfigurationResponse) {
    response = &ListPermissionPoliciesInRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListPermissionPoliciesInRoleConfiguration
// This API is used to obtain the policy list in permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) ListPermissionPoliciesInRoleConfiguration(request *ListPermissionPoliciesInRoleConfigurationRequest) (response *ListPermissionPoliciesInRoleConfigurationResponse, err error) {
    return c.ListPermissionPoliciesInRoleConfigurationWithContext(context.Background(), request)
}

// ListPermissionPoliciesInRoleConfiguration
// This API is used to obtain the policy list in permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) ListPermissionPoliciesInRoleConfigurationWithContext(ctx context.Context, request *ListPermissionPoliciesInRoleConfigurationRequest) (response *ListPermissionPoliciesInRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewListPermissionPoliciesInRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPermissionPoliciesInRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPermissionPoliciesInRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewListRoleAssignmentsRequest() (request *ListRoleAssignmentsRequest) {
    request = &ListRoleAssignmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListRoleAssignments")
    
    
    return
}

func NewListRoleAssignmentsResponse() (response *ListRoleAssignmentsResponse) {
    response = &ListRoleAssignmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRoleAssignments
// This API is used to query the authorization list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleAssignments(request *ListRoleAssignmentsRequest) (response *ListRoleAssignmentsResponse, err error) {
    return c.ListRoleAssignmentsWithContext(context.Background(), request)
}

// ListRoleAssignments
// This API is used to query the authorization list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleAssignmentsWithContext(ctx context.Context, request *ListRoleAssignmentsRequest) (response *ListRoleAssignmentsResponse, err error) {
    if request == nil {
        request = NewListRoleAssignmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleAssignments require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleAssignmentsResponse()
    err = c.Send(request, response)
    return
}

func NewListRoleConfigurationProvisioningsRequest() (request *ListRoleConfigurationProvisioningsRequest) {
    request = &ListRoleConfigurationProvisioningsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListRoleConfigurationProvisionings")
    
    
    return
}

func NewListRoleConfigurationProvisioningsResponse() (response *ListRoleConfigurationProvisioningsResponse) {
    response = &ListRoleConfigurationProvisioningsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRoleConfigurationProvisionings
// This API is used to query the permission configuration deployment list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurationProvisionings(request *ListRoleConfigurationProvisioningsRequest) (response *ListRoleConfigurationProvisioningsResponse, err error) {
    return c.ListRoleConfigurationProvisioningsWithContext(context.Background(), request)
}

// ListRoleConfigurationProvisionings
// This API is used to query the permission configuration deployment list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurationProvisioningsWithContext(ctx context.Context, request *ListRoleConfigurationProvisioningsRequest) (response *ListRoleConfigurationProvisioningsResponse, err error) {
    if request == nil {
        request = NewListRoleConfigurationProvisioningsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleConfigurationProvisionings require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleConfigurationProvisioningsResponse()
    err = c.Send(request, response)
    return
}

func NewListRoleConfigurationsRequest() (request *ListRoleConfigurationsRequest) {
    request = &ListRoleConfigurationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListRoleConfigurations")
    
    
    return
}

func NewListRoleConfigurationsResponse() (response *ListRoleConfigurationsResponse) {
    response = &ListRoleConfigurationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRoleConfigurations
// This API is used to query the permission configuration list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurations(request *ListRoleConfigurationsRequest) (response *ListRoleConfigurationsResponse, err error) {
    return c.ListRoleConfigurationsWithContext(context.Background(), request)
}

// ListRoleConfigurations
// This API is used to query the permission configuration list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurationsWithContext(ctx context.Context, request *ListRoleConfigurationsRequest) (response *ListRoleConfigurationsResponse, err error) {
    if request == nil {
        request = NewListRoleConfigurationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleConfigurations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewListTasksRequest() (request *ListTasksRequest) {
    request = &ListTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListTasks")
    
    
    return
}

func NewListTasksResponse() (response *ListTasksResponse) {
    response = &ListTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTasks
// This API is used to query the async task list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
    return c.ListTasksWithContext(context.Background(), request)
}

// ListTasks
// This API is used to query the async task list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) (response *ListTasksResponse, err error) {
    if request == nil {
        request = NewListTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListUserSyncProvisioningsRequest() (request *ListUserSyncProvisioningsRequest) {
    request = &ListUserSyncProvisioningsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListUserSyncProvisionings")
    
    
    return
}

func NewListUserSyncProvisioningsResponse() (response *ListUserSyncProvisioningsResponse) {
    response = &ListUserSyncProvisioningsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUserSyncProvisionings
// This API is used to query the CAM user synchronization list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListUserSyncProvisionings(request *ListUserSyncProvisioningsRequest) (response *ListUserSyncProvisioningsResponse, err error) {
    return c.ListUserSyncProvisioningsWithContext(context.Background(), request)
}

// ListUserSyncProvisionings
// This API is used to query the CAM user synchronization list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListUserSyncProvisioningsWithContext(ctx context.Context, request *ListUserSyncProvisioningsRequest) (response *ListUserSyncProvisioningsResponse, err error) {
    if request == nil {
        request = NewListUserSyncProvisioningsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserSyncProvisionings require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserSyncProvisioningsResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListUsers")
    
    
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsers
// This API is used to query the user list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    return c.ListUsersWithContext(context.Background(), request)
}

// ListUsers
// This API is used to query the user list.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest) (response *ListUsersResponse, err error) {
    if request == nil {
        request = NewListUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersResponse()
    err = c.Send(request, response)
    return
}

func NewMoveOrganizationNodeMembersRequest() (request *MoveOrganizationNodeMembersRequest) {
    request = &MoveOrganizationNodeMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNodeMembers")
    
    
    return
}

func NewMoveOrganizationNodeMembersResponse() (response *MoveOrganizationNodeMembersResponse) {
    response = &MoveOrganizationNodeMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MoveOrganizationNodeMembers
// This API is used to move a member to the specified organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationNodeMembers(request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
    return c.MoveOrganizationNodeMembersWithContext(context.Background(), request)
}

// MoveOrganizationNodeMembers
// This API is used to move a member to the specified organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationNodeMembersWithContext(ctx context.Context, request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
    if request == nil {
        request = NewMoveOrganizationNodeMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveOrganizationNodeMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveOrganizationNodeMembersResponse()
    err = c.Send(request, response)
    return
}

func NewOpenIdentityCenterRequest() (request *OpenIdentityCenterRequest) {
    request = &OpenIdentityCenterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "OpenIdentityCenter")
    
    
    return
}

func NewOpenIdentityCenterResponse() (response *OpenIdentityCenterResponse) {
    response = &OpenIdentityCenterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenIdentityCenter
// This API is used to activate the CIC service.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERALREADYOPEN = "FailedOperation.IdentityCenterAlreadyOpen"
//  FAILEDOPERATION_IDENTITYCENTERNOTAUTH = "FailedOperation.IdentityCenterNotAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTENTERPRISEAUTH = "FailedOperation.IdentityCenterNotEnterpriseAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTORGANIZATIONMANAGER = "FailedOperation.IdentityCenterNotOrganizationManager"
//  FAILEDOPERATION_IDENTITYCENTERORGANIZATIONNOTOPEN = "FailedOperation.IdentityCenterOrganizationNotOpen"
//  INVALIDPARAMETERVALUE_IDENTITYCENTERZONENAMEALREADYEXIST = "InvalidParameterValue.IdentityCenterZoneNameAlreadyExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
func (c *Client) OpenIdentityCenter(request *OpenIdentityCenterRequest) (response *OpenIdentityCenterResponse, err error) {
    return c.OpenIdentityCenterWithContext(context.Background(), request)
}

// OpenIdentityCenter
// This API is used to activate the CIC service.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERALREADYOPEN = "FailedOperation.IdentityCenterAlreadyOpen"
//  FAILEDOPERATION_IDENTITYCENTERNOTAUTH = "FailedOperation.IdentityCenterNotAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTENTERPRISEAUTH = "FailedOperation.IdentityCenterNotEnterpriseAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTORGANIZATIONMANAGER = "FailedOperation.IdentityCenterNotOrganizationManager"
//  FAILEDOPERATION_IDENTITYCENTERORGANIZATIONNOTOPEN = "FailedOperation.IdentityCenterOrganizationNotOpen"
//  INVALIDPARAMETERVALUE_IDENTITYCENTERZONENAMEALREADYEXIST = "InvalidParameterValue.IdentityCenterZoneNameAlreadyExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
func (c *Client) OpenIdentityCenterWithContext(ctx context.Context, request *OpenIdentityCenterRequest) (response *OpenIdentityCenterResponse, err error) {
    if request == nil {
        request = NewOpenIdentityCenterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenIdentityCenter require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenIdentityCenterResponse()
    err = c.Send(request, response)
    return
}

func NewProvisionRoleConfigurationRequest() (request *ProvisionRoleConfigurationRequest) {
    request = &ProvisionRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ProvisionRoleConfiguration")
    
    
    return
}

func NewProvisionRoleConfigurationResponse() (response *ProvisionRoleConfigurationResponse) {
    response = &ProvisionRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProvisionRoleConfiguration
// This API is used to deploy permission configurations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
func (c *Client) ProvisionRoleConfiguration(request *ProvisionRoleConfigurationRequest) (response *ProvisionRoleConfigurationResponse, err error) {
    return c.ProvisionRoleConfigurationWithContext(context.Background(), request)
}

// ProvisionRoleConfiguration
// This API is used to deploy permission configurations on member accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
func (c *Client) ProvisionRoleConfigurationWithContext(ctx context.Context, request *ProvisionRoleConfigurationRequest) (response *ProvisionRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewProvisionRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProvisionRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewProvisionRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveExternalSAMLIdPCertificateRequest() (request *RemoveExternalSAMLIdPCertificateRequest) {
    request = &RemoveExternalSAMLIdPCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RemoveExternalSAMLIdPCertificate")
    
    
    return
}

func NewRemoveExternalSAMLIdPCertificateResponse() (response *RemoveExternalSAMLIdPCertificateResponse) {
    response = &RemoveExternalSAMLIdPCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveExternalSAMLIdPCertificate
// This API is used to remove SAML signing certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_X509CERTIFICATENOTFOUND = "ResourceNotFound.X509CertificateNotFound"
func (c *Client) RemoveExternalSAMLIdPCertificate(request *RemoveExternalSAMLIdPCertificateRequest) (response *RemoveExternalSAMLIdPCertificateResponse, err error) {
    return c.RemoveExternalSAMLIdPCertificateWithContext(context.Background(), request)
}

// RemoveExternalSAMLIdPCertificate
// This API is used to remove SAML signing certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_X509CERTIFICATENOTFOUND = "ResourceNotFound.X509CertificateNotFound"
func (c *Client) RemoveExternalSAMLIdPCertificateWithContext(ctx context.Context, request *RemoveExternalSAMLIdPCertificateRequest) (response *RemoveExternalSAMLIdPCertificateResponse, err error) {
    if request == nil {
        request = NewRemoveExternalSAMLIdPCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveExternalSAMLIdPCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveExternalSAMLIdPCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePermissionPolicyFromRoleConfigurationRequest() (request *RemovePermissionPolicyFromRoleConfigurationRequest) {
    request = &RemovePermissionPolicyFromRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RemovePermissionPolicyFromRoleConfiguration")
    
    
    return
}

func NewRemovePermissionPolicyFromRoleConfigurationResponse() (response *RemovePermissionPolicyFromRoleConfigurationResponse) {
    response = &RemovePermissionPolicyFromRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemovePermissionPolicyFromRoleConfiguration
// This API is used to remove policies from permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) RemovePermissionPolicyFromRoleConfiguration(request *RemovePermissionPolicyFromRoleConfigurationRequest) (response *RemovePermissionPolicyFromRoleConfigurationResponse, err error) {
    return c.RemovePermissionPolicyFromRoleConfigurationWithContext(context.Background(), request)
}

// RemovePermissionPolicyFromRoleConfiguration
// This API is used to remove policies from permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) RemovePermissionPolicyFromRoleConfigurationWithContext(ctx context.Context, request *RemovePermissionPolicyFromRoleConfigurationRequest) (response *RemovePermissionPolicyFromRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewRemovePermissionPolicyFromRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemovePermissionPolicyFromRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemovePermissionPolicyFromRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
    request = &RemoveUserFromGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RemoveUserFromGroup")
    
    
    return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
    response = &RemoveUserFromGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveUserFromGroup
// This API is used to removes users from a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"
//  LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"
func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    return c.RemoveUserFromGroupWithContext(context.Background(), request)
}

// RemoveUserFromGroup
// This API is used to removes users from a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"
//  LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"
func (c *Client) RemoveUserFromGroupWithContext(ctx context.Context, request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserFromGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserFromGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSetExternalSAMLIdentityProviderRequest() (request *SetExternalSAMLIdentityProviderRequest) {
    request = &SetExternalSAMLIdentityProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "SetExternalSAMLIdentityProvider")
    
    
    return
}

func NewSetExternalSAMLIdentityProviderResponse() (response *SetExternalSAMLIdentityProviderResponse) {
    response = &SetExternalSAMLIdentityProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetExternalSAMLIdentityProvider
// This API is used to configure the SAML identity provider information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DECODEMETADATADOCUMENTFAILED = "FailedOperation.DecodeMetadataDocumentFailed"
//  FAILEDOPERATION_METADATACOSPARSINGFAILED = "FailedOperation.MetadataCosParsingFailed"
//  FAILEDOPERATION_SAMLSERVICEPROVIDERCREATEFAILED = "FailedOperation.SAMLServiceProviderCreateFailed"
//  FAILEDOPERATION_UPLOADMETADATAFAILED = "FailedOperation.UploadMetadataFailed"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  FAILEDOPERATION_XMLDATAUNMARSHALFAILED = "FailedOperation.XMLDataUnmarshalFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SSOSTATUSINVALID = "InvalidParameterValue.SSoStatusInvalid"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) SetExternalSAMLIdentityProvider(request *SetExternalSAMLIdentityProviderRequest) (response *SetExternalSAMLIdentityProviderResponse, err error) {
    return c.SetExternalSAMLIdentityProviderWithContext(context.Background(), request)
}

// SetExternalSAMLIdentityProvider
// This API is used to configure the SAML identity provider information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DECODEMETADATADOCUMENTFAILED = "FailedOperation.DecodeMetadataDocumentFailed"
//  FAILEDOPERATION_METADATACOSPARSINGFAILED = "FailedOperation.MetadataCosParsingFailed"
//  FAILEDOPERATION_SAMLSERVICEPROVIDERCREATEFAILED = "FailedOperation.SAMLServiceProviderCreateFailed"
//  FAILEDOPERATION_UPLOADMETADATAFAILED = "FailedOperation.UploadMetadataFailed"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  FAILEDOPERATION_XMLDATAUNMARSHALFAILED = "FailedOperation.XMLDataUnmarshalFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SSOSTATUSINVALID = "InvalidParameterValue.SSoStatusInvalid"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) SetExternalSAMLIdentityProviderWithContext(ctx context.Context, request *SetExternalSAMLIdentityProviderRequest) (response *SetExternalSAMLIdentityProviderResponse, err error) {
    if request == nil {
        request = NewSetExternalSAMLIdentityProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
    request = &UpdateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateGroup")
    
    
    return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
    response = &UpdateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGroup
// This API is used to modify user group information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    return c.UpdateGroupWithContext(context.Background(), request)
}

// UpdateGroup
// This API is used to modify user group information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    if request == nil {
        request = NewUpdateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationNodeRequest() (request *UpdateOrganizationNodeRequest) {
    request = &UpdateOrganizationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationNode")
    
    
    return
}

func NewUpdateOrganizationNodeResponse() (response *UpdateOrganizationNodeResponse) {
    response = &UpdateOrganizationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationNode
// This API is used to update an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNode(request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    return c.UpdateOrganizationNodeWithContext(context.Background(), request)
}

// UpdateOrganizationNode
// This API is used to update an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNodeWithContext(ctx context.Context, request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleConfigurationRequest() (request *UpdateRoleConfigurationRequest) {
    request = &UpdateRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateRoleConfiguration")
    
    
    return
}

func NewUpdateRoleConfigurationResponse() (response *UpdateRoleConfigurationResponse) {
    response = &UpdateRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRoleConfiguration
// This API is used to modify the permission configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) UpdateRoleConfiguration(request *UpdateRoleConfigurationRequest) (response *UpdateRoleConfigurationResponse, err error) {
    return c.UpdateRoleConfigurationWithContext(context.Background(), request)
}

// UpdateRoleConfiguration
// This API is used to modify the permission configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) UpdateRoleConfigurationWithContext(ctx context.Context, request *UpdateRoleConfigurationRequest) (response *UpdateRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateRoleConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateUser")
    
    
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUser
// This API is used to modify user information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    return c.UpdateUserWithContext(context.Background(), request)
}

// UpdateUser
// This API is used to modify user information.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserStatusRequest() (request *UpdateUserStatusRequest) {
    request = &UpdateUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateUserStatus")
    
    
    return
}

func NewUpdateUserStatusResponse() (response *UpdateUserStatusResponse) {
    response = &UpdateUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserStatus
// This API is used to modify the user status.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    return c.UpdateUserStatusWithContext(context.Background(), request)
}

// UpdateUserStatus
// This API is used to modify the user status.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserStatusWithContext(ctx context.Context, request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    if request == nil {
        request = NewUpdateUserStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserSyncProvisioningRequest() (request *UpdateUserSyncProvisioningRequest) {
    request = &UpdateUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateUserSyncProvisioning")
    
    
    return
}

func NewUpdateUserSyncProvisioningResponse() (response *UpdateUserSyncProvisioningResponse) {
    response = &UpdateUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserSyncProvisioning
// This API is used to create sub-user synchronization tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) UpdateUserSyncProvisioning(request *UpdateUserSyncProvisioningRequest) (response *UpdateUserSyncProvisioningResponse, err error) {
    return c.UpdateUserSyncProvisioningWithContext(context.Background(), request)
}

// UpdateUserSyncProvisioning
// This API is used to create sub-user synchronization tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) UpdateUserSyncProvisioningWithContext(ctx context.Context, request *UpdateUserSyncProvisioningRequest) (response *UpdateUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewUpdateUserSyncProvisioningRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserSyncProvisioningResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateZoneRequest() (request *UpdateZoneRequest) {
    request = &UpdateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateZone")
    
    
    return
}

func NewUpdateZoneResponse() (response *UpdateZoneResponse) {
    response = &UpdateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateZone
// This API is used to update the user's space name.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETERVALUE_ZONENAMEFORMATERROR = "InvalidParameterValue.ZoneNameFormatError"
func (c *Client) UpdateZone(request *UpdateZoneRequest) (response *UpdateZoneResponse, err error) {
    return c.UpdateZoneWithContext(context.Background(), request)
}

// UpdateZone
// This API is used to update the user's space name.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETERVALUE_ZONENAMEFORMATERROR = "InvalidParameterValue.ZoneNameFormatError"
func (c *Client) UpdateZoneWithContext(ctx context.Context, request *UpdateZoneRequest) (response *UpdateZoneResponse, err error) {
    if request == nil {
        request = NewUpdateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateZoneResponse()
    err = c.Send(request, response)
    return
}
