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

package v20190116

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-01-16"

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


func NewAddUserRequest() (request *AddUserRequest) {
    request = &AddUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "AddUser")
    
    
    return
}

func NewAddUserResponse() (response *AddUserResponse) {
    response = &AddUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddUser
// This API is used to add sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PASSWORDVIOLATEDRULES = "InvalidParameter.PasswordViolatedRules"
//  INVALIDPARAMETER_SUBUSERFULL = "InvalidParameter.SubUserFull"
//  INVALIDPARAMETER_SUBUSERNAMEINUSE = "InvalidParameter.SubUserNameInUse"
//  INVALIDPARAMETER_USERNAMEILLEGAL = "InvalidParameter.UserNameIllegal"
//  REQUESTLIMITEXCEEDED_CREATEUSER = "RequestLimitExceeded.CreateUser"
func (c *Client) AddUser(request *AddUserRequest) (response *AddUserResponse, err error) {
    return c.AddUserWithContext(context.Background(), request)
}

// AddUser
// This API is used to add sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PASSWORDVIOLATEDRULES = "InvalidParameter.PasswordViolatedRules"
//  INVALIDPARAMETER_SUBUSERFULL = "InvalidParameter.SubUserFull"
//  INVALIDPARAMETER_SUBUSERNAMEINUSE = "InvalidParameter.SubUserNameInUse"
//  INVALIDPARAMETER_USERNAMEILLEGAL = "InvalidParameter.UserNameIllegal"
//  REQUESTLIMITEXCEEDED_CREATEUSER = "RequestLimitExceeded.CreateUser"
func (c *Client) AddUserWithContext(ctx context.Context, request *AddUserRequest) (response *AddUserResponse, err error) {
    if request == nil {
        request = NewAddUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
    request = &AddUserToGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "AddUserToGroup")
    
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERFULL = "InvalidParameter.GroupUserFull"
//  INVALIDPARAMETER_USERGROUPFULL = "InvalidParameter.UserGroupFull"
//  INVALIDPARAMETER_USERUINANDUINNOTALLNULL = "InvalidParameter.UserUinAndUinNotAllNull"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    return c.AddUserToGroupWithContext(context.Background(), request)
}

// AddUserToGroup
// This API is used to add users to a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERFULL = "InvalidParameter.GroupUserFull"
//  INVALIDPARAMETER_USERGROUPFULL = "InvalidParameter.UserGroupFull"
//  INVALIDPARAMETER_USERUINANDUINNOTALLNULL = "InvalidParameter.UserUinAndUinNotAllNull"
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

func NewAttachGroupPolicyRequest() (request *AttachGroupPolicyRequest) {
    request = &AttachGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "AttachGroupPolicy")
    
    
    return
}

func NewAttachGroupPolicyResponse() (response *AttachGroupPolicyResponse) {
    response = &AttachGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachGroupPolicy
// This API (AttachGroupPolicy) is used to associate a policy with a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AttachGroupPolicy(request *AttachGroupPolicyRequest) (response *AttachGroupPolicyResponse, err error) {
    return c.AttachGroupPolicyWithContext(context.Background(), request)
}

// AttachGroupPolicy
// This API (AttachGroupPolicy) is used to associate a policy with a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AttachGroupPolicyWithContext(ctx context.Context, request *AttachGroupPolicyRequest) (response *AttachGroupPolicyResponse, err error) {
    if request == nil {
        request = NewAttachGroupPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAttachRolePolicyRequest() (request *AttachRolePolicyRequest) {
    request = &AttachRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "AttachRolePolicy")
    
    
    return
}

func NewAttachRolePolicyResponse() (response *AttachRolePolicyResponse) {
    response = &AttachRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachRolePolicy
// This API (AttachRolePolicy) is used to associate a policy with a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
func (c *Client) AttachRolePolicy(request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
    return c.AttachRolePolicyWithContext(context.Background(), request)
}

// AttachRolePolicy
// This API (AttachRolePolicy) is used to associate a policy with a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
func (c *Client) AttachRolePolicyWithContext(ctx context.Context, request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
    if request == nil {
        request = NewAttachRolePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachRolePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAttachUserPolicyRequest() (request *AttachUserPolicyRequest) {
    request = &AttachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "AttachUserPolicy")
    
    
    return
}

func NewAttachUserPolicyResponse() (response *AttachUserPolicyResponse) {
    response = &AttachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachUserPolicy
// This API (AttachUserPolicy) is used to associates a policy with a user.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AttachUserPolicy(request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    return c.AttachUserPolicyWithContext(context.Background(), request)
}

// AttachUserPolicy
// This API (AttachUserPolicy) is used to associates a policy with a user.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AttachUserPolicyWithContext(ctx context.Context, request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    if request == nil {
        request = NewAttachUserPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachUserPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewConsumeCustomMFATokenRequest() (request *ConsumeCustomMFATokenRequest) {
    request = &ConsumeCustomMFATokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ConsumeCustomMFAToken")
    
    
    return
}

func NewConsumeCustomMFATokenResponse() (response *ConsumeCustomMFATokenResponse) {
    response = &ConsumeCustomMFATokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ConsumeCustomMFAToken
// This API is used to verify a custom multi-factor Token.
//
// error code that may be returned:
//  INVALIDPARAMETER_MFATOKENERROR = "InvalidParameter.MFATokenError"
func (c *Client) ConsumeCustomMFAToken(request *ConsumeCustomMFATokenRequest) (response *ConsumeCustomMFATokenResponse, err error) {
    return c.ConsumeCustomMFATokenWithContext(context.Background(), request)
}

// ConsumeCustomMFAToken
// This API is used to verify a custom multi-factor Token.
//
// error code that may be returned:
//  INVALIDPARAMETER_MFATOKENERROR = "InvalidParameter.MFATokenError"
func (c *Client) ConsumeCustomMFATokenWithContext(ctx context.Context, request *ConsumeCustomMFATokenRequest) (response *ConsumeCustomMFATokenResponse, err error) {
    if request == nil {
        request = NewConsumeCustomMFATokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConsumeCustomMFAToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewConsumeCustomMFATokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateGroup")
    
    
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroup
// This API is used to create a user group.
//
// error code that may be returned:
//  INVALIDPARAMETER_GROUPFULL = "InvalidParameter.GroupFull"
//  INVALIDPARAMETER_GROUPNAMEINUSE = "InvalidParameter.GroupNameInUse"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

// CreateGroup
// This API is used to create a user group.
//
// error code that may be returned:
//  INVALIDPARAMETER_GROUPFULL = "InvalidParameter.GroupFull"
//  INVALIDPARAMETER_GROUPNAMEINUSE = "InvalidParameter.GroupNameInUse"
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

func NewCreateOIDCConfigRequest() (request *CreateOIDCConfigRequest) {
    request = &CreateOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateOIDCConfig")
    
    
    return
}

func NewCreateOIDCConfigResponse() (response *CreateOIDCConfigResponse) {
    response = &CreateOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOIDCConfig
// This API is used to create role OIDC configurations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateOIDCConfig(request *CreateOIDCConfigRequest) (response *CreateOIDCConfigResponse, err error) {
    return c.CreateOIDCConfigWithContext(context.Background(), request)
}

// CreateOIDCConfig
// This API is used to create role OIDC configurations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateOIDCConfigWithContext(ctx context.Context, request *CreateOIDCConfigRequest) (response *CreateOIDCConfigResponse, err error) {
    if request == nil {
        request = NewCreateOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
    request = &CreatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreatePolicy")
    
    
    return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
    response = &CreatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePolicy
// This API (CreatePolicy) is used to create a policy.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    return c.CreatePolicyWithContext(context.Background(), request)
}

// CreatePolicy
// This API (CreatePolicy) is used to create a policy.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyVersionRequest() (request *CreatePolicyVersionRequest) {
    request = &CreatePolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreatePolicyVersion")
    
    
    return
}

func NewCreatePolicyVersionResponse() (response *CreatePolicyVersionResponse) {
    response = &CreatePolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePolicyVersion
// This API is used to add a policy version. After creating a policy version, you can easily change the policy by changing the policy version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  FAILEDOPERATION_POLICYVERSIONFULL = "FailedOperation.PolicyVersionFull"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (response *CreatePolicyVersionResponse, err error) {
    return c.CreatePolicyVersionWithContext(context.Background(), request)
}

// CreatePolicyVersion
// This API is used to add a policy version. After creating a policy version, you can easily change the policy by changing the policy version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  FAILEDOPERATION_POLICYVERSIONFULL = "FailedOperation.PolicyVersionFull"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePolicyVersionWithContext(ctx context.Context, request *CreatePolicyVersionRequest) (response *CreatePolicyVersionResponse, err error) {
    if request == nil {
        request = NewCreatePolicyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateRole")
    
    
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRole
// This API (CreateRole) is used to create a role.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGRESOURCEFAILED = "FailedOperation.TagResourceFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSCROSSERROR = "InvalidParameter.PrincipalQcsCrossError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_ROLEFULL = "InvalidParameter.RoleFull"
//  INVALIDPARAMETER_ROLENAMEERROR = "InvalidParameter.RoleNameError"
//  INVALIDPARAMETER_ROLENAMEINUSE = "InvalidParameter.RoleNameInUse"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_TAGLIMITEXCEEDED = "InvalidParameter.TagLimitExceeded"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    return c.CreateRoleWithContext(context.Background(), request)
}

// CreateRole
// This API (CreateRole) is used to create a role.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGRESOURCEFAILED = "FailedOperation.TagResourceFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSCROSSERROR = "InvalidParameter.PrincipalQcsCrossError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_ROLEFULL = "InvalidParameter.RoleFull"
//  INVALIDPARAMETER_ROLENAMEERROR = "InvalidParameter.RoleNameError"
//  INVALIDPARAMETER_ROLENAMEINUSE = "InvalidParameter.RoleNameInUse"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_TAGLIMITEXCEEDED = "InvalidParameter.TagLimitExceeded"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
func (c *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSAMLProviderRequest() (request *CreateSAMLProviderRequest) {
    request = &CreateSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateSAMLProvider")
    
    
    return
}

func NewCreateSAMLProviderResponse() (response *CreateSAMLProviderResponse) {
    response = &CreateSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSAMLProvider
// This API is used to create a SAML identity provider.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateSAMLProvider(request *CreateSAMLProviderRequest) (response *CreateSAMLProviderResponse, err error) {
    return c.CreateSAMLProviderWithContext(context.Background(), request)
}

// CreateSAMLProvider
// This API is used to create a SAML identity provider.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateSAMLProviderWithContext(ctx context.Context, request *CreateSAMLProviderRequest) (response *CreateSAMLProviderResponse, err error) {
    if request == nil {
        request = NewCreateSAMLProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSAMLProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceLinkedRoleRequest() (request *CreateServiceLinkedRoleRequest) {
    request = &CreateServiceLinkedRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateServiceLinkedRole")
    
    
    return
}

func NewCreateServiceLinkedRoleResponse() (response *CreateServiceLinkedRoleResponse) {
    response = &CreateServiceLinkedRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServiceLinkedRole
// This API is used to create a service-linked role.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGRESOURCEFAILED = "FailedOperation.TagResourceFailed"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_ROLENAMEERROR = "InvalidParameter.RoleNameError"
//  INVALIDPARAMETER_ROLENAMEINUSE = "InvalidParameter.RoleNameInUse"
//  INVALIDPARAMETER_TAGLIMITEXCEEDED = "InvalidParameter.TagLimitExceeded"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
func (c *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (response *CreateServiceLinkedRoleResponse, err error) {
    return c.CreateServiceLinkedRoleWithContext(context.Background(), request)
}

// CreateServiceLinkedRole
// This API is used to create a service-linked role.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGRESOURCEFAILED = "FailedOperation.TagResourceFailed"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_ROLENAMEERROR = "InvalidParameter.RoleNameError"
//  INVALIDPARAMETER_ROLENAMEINUSE = "InvalidParameter.RoleNameInUse"
//  INVALIDPARAMETER_TAGLIMITEXCEEDED = "InvalidParameter.TagLimitExceeded"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
func (c *Client) CreateServiceLinkedRoleWithContext(ctx context.Context, request *CreateServiceLinkedRoleRequest) (response *CreateServiceLinkedRoleResponse, err error) {
    if request == nil {
        request = NewCreateServiceLinkedRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceLinkedRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceLinkedRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserOIDCConfigRequest() (request *CreateUserOIDCConfigRequest) {
    request = &CreateUserOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateUserOIDCConfig")
    
    
    return
}

func NewCreateUserOIDCConfigResponse() (response *CreateUserOIDCConfigResponse) {
    response = &CreateUserOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserOIDCConfig
// This API is used to create a user OIDC configuration. Only one user OIDC IdP can be created, and the user SAML SSO IdP will be automatically disabled after it is created.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateUserOIDCConfig(request *CreateUserOIDCConfigRequest) (response *CreateUserOIDCConfigResponse, err error) {
    return c.CreateUserOIDCConfigWithContext(context.Background(), request)
}

// CreateUserOIDCConfig
// This API is used to create a user OIDC configuration. Only one user OIDC IdP can be created, and the user SAML SSO IdP will be automatically disabled after it is created.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
func (c *Client) CreateUserOIDCConfigWithContext(ctx context.Context, request *CreateUserOIDCConfigRequest) (response *CreateUserOIDCConfigResponse, err error) {
    if request == nil {
        request = NewCreateUserOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserSAMLConfigRequest() (request *CreateUserSAMLConfigRequest) {
    request = &CreateUserSAMLConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "CreateUserSAMLConfig")
    
    
    return
}

func NewCreateUserSAMLConfigResponse() (response *CreateUserSAMLConfigResponse) {
    response = &CreateUserSAMLConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserSAMLConfig
// This API is used to create user SAML configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
func (c *Client) CreateUserSAMLConfig(request *CreateUserSAMLConfigRequest) (response *CreateUserSAMLConfigResponse, err error) {
    return c.CreateUserSAMLConfigWithContext(context.Background(), request)
}

// CreateUserSAMLConfig
// This API is used to create user SAML configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
func (c *Client) CreateUserSAMLConfigWithContext(ctx context.Context, request *CreateUserSAMLConfigRequest) (response *CreateUserSAMLConfigResponse, err error) {
    if request == nil {
        request = NewCreateUserSAMLConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserSAMLConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserSAMLConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroup
// This API is used to delete a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// This API is used to delete a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

func NewDeleteOIDCConfigRequest() (request *DeleteOIDCConfigRequest) {
    request = &DeleteOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteOIDCConfig")
    
    
    return
}

func NewDeleteOIDCConfigResponse() (response *DeleteOIDCConfigResponse) {
    response = &DeleteOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOIDCConfig
// This API is used to delete OIDC IdPs.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DeleteOIDCConfig(request *DeleteOIDCConfigRequest) (response *DeleteOIDCConfigResponse, err error) {
    return c.DeleteOIDCConfigWithContext(context.Background(), request)
}

// DeleteOIDCConfig
// This API is used to delete OIDC IdPs.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DeleteOIDCConfigWithContext(ctx context.Context, request *DeleteOIDCConfigRequest) (response *DeleteOIDCConfigResponse, err error) {
    if request == nil {
        request = NewDeleteOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
    request = &DeletePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeletePolicy")
    
    
    return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
    response = &DeletePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePolicy
// This API (DeletePolicy) is used to delete a policy.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    return c.DeletePolicyWithContext(context.Background(), request)
}

// DeletePolicy
// This API (DeletePolicy) is used to delete a policy.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    if request == nil {
        request = NewDeletePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyVersionRequest() (request *DeletePolicyVersionRequest) {
    request = &DeletePolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeletePolicyVersion")
    
    
    return
}

func NewDeletePolicyVersionResponse() (response *DeletePolicyVersionResponse) {
    response = &DeletePolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePolicyVersion
// This API is used to delete a policy version of a policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  FAILEDOPERATION_POLICYVERSIONALREADYDEFAULT = "FailedOperation.PolicyVersionAlreadyDefault"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYVERSIONNOTEXISTS = "InvalidParameter.PolicyVersionNotExists"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (response *DeletePolicyVersionResponse, err error) {
    return c.DeletePolicyVersionWithContext(context.Background(), request)
}

// DeletePolicyVersion
// This API is used to delete a policy version of a policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  FAILEDOPERATION_POLICYVERSIONALREADYDEFAULT = "FailedOperation.PolicyVersionAlreadyDefault"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYVERSIONNOTEXISTS = "InvalidParameter.PolicyVersionNotExists"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeletePolicyVersionWithContext(ctx context.Context, request *DeletePolicyVersionRequest) (response *DeletePolicyVersionResponse, err error) {
    if request == nil {
        request = NewDeletePolicyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteRole")
    
    
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRole
// This API (DeleteRole) is used to delete a specified role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    return c.DeleteRoleWithContext(context.Background(), request)
}

// DeleteRole
// This API (DeleteRole) is used to delete a specified role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRolePermissionsBoundaryRequest() (request *DeleteRolePermissionsBoundaryRequest) {
    request = &DeleteRolePermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteRolePermissionsBoundary")
    
    
    return
}

func NewDeleteRolePermissionsBoundaryResponse() (response *DeleteRolePermissionsBoundaryResponse) {
    response = &DeleteRolePermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRolePermissionsBoundary
// This API is used to delete a role permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_SERVICELINKEDROLECANTUSEPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedRoleCantUsePermissionBoundary"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRolePermissionsBoundary(request *DeleteRolePermissionsBoundaryRequest) (response *DeleteRolePermissionsBoundaryResponse, err error) {
    return c.DeleteRolePermissionsBoundaryWithContext(context.Background(), request)
}

// DeleteRolePermissionsBoundary
// This API is used to delete a role permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_SERVICELINKEDROLECANTUSEPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedRoleCantUsePermissionBoundary"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRolePermissionsBoundaryWithContext(ctx context.Context, request *DeleteRolePermissionsBoundaryRequest) (response *DeleteRolePermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewDeleteRolePermissionsBoundaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRolePermissionsBoundary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRolePermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSAMLProviderRequest() (request *DeleteSAMLProviderRequest) {
    request = &DeleteSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteSAMLProvider")
    
    
    return
}

func NewDeleteSAMLProviderResponse() (response *DeleteSAMLProviderResponse) {
    response = &DeleteSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSAMLProvider
// This API is used to delete a SAML identity provider.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DeleteSAMLProvider(request *DeleteSAMLProviderRequest) (response *DeleteSAMLProviderResponse, err error) {
    return c.DeleteSAMLProviderWithContext(context.Background(), request)
}

// DeleteSAMLProvider
// This API is used to delete a SAML identity provider.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DeleteSAMLProviderWithContext(ctx context.Context, request *DeleteSAMLProviderRequest) (response *DeleteSAMLProviderResponse, err error) {
    if request == nil {
        request = NewDeleteSAMLProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSAMLProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceLinkedRoleRequest() (request *DeleteServiceLinkedRoleRequest) {
    request = &DeleteServiceLinkedRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteServiceLinkedRole")
    
    
    return
}

func NewDeleteServiceLinkedRoleResponse() (response *DeleteServiceLinkedRoleResponse) {
    response = &DeleteServiceLinkedRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServiceLinkedRole
// This API is used to delete a service-linked role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) DeleteServiceLinkedRole(request *DeleteServiceLinkedRoleRequest) (response *DeleteServiceLinkedRoleResponse, err error) {
    return c.DeleteServiceLinkedRoleWithContext(context.Background(), request)
}

// DeleteServiceLinkedRole
// This API is used to delete a service-linked role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) DeleteServiceLinkedRoleWithContext(ctx context.Context, request *DeleteServiceLinkedRoleRequest) (response *DeleteServiceLinkedRoleResponse, err error) {
    if request == nil {
        request = NewDeleteServiceLinkedRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceLinkedRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceLinkedRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUser
// This API is used to delete a sub-user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_HAVEKEYS = "OperationDenied.HaveKeys"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION_DELETEAPIKEY = "UnauthorizedOperation.DeleteApiKey"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete a sub-user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_HAVEKEYS = "OperationDenied.HaveKeys"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION_DELETEAPIKEY = "UnauthorizedOperation.DeleteApiKey"
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

func NewDeleteUserPermissionsBoundaryRequest() (request *DeleteUserPermissionsBoundaryRequest) {
    request = &DeleteUserPermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DeleteUserPermissionsBoundary")
    
    
    return
}

func NewDeleteUserPermissionsBoundaryResponse() (response *DeleteUserPermissionsBoundaryResponse) {
    response = &DeleteUserPermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserPermissionsBoundary
// This API is used to delete a user permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteUserPermissionsBoundary(request *DeleteUserPermissionsBoundaryRequest) (response *DeleteUserPermissionsBoundaryResponse, err error) {
    return c.DeleteUserPermissionsBoundaryWithContext(context.Background(), request)
}

// DeleteUserPermissionsBoundary
// This API is used to delete a user permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteUserPermissionsBoundaryWithContext(ctx context.Context, request *DeleteUserPermissionsBoundaryRequest) (response *DeleteUserPermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewDeleteUserPermissionsBoundaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserPermissionsBoundary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserPermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOIDCConfigRequest() (request *DescribeOIDCConfigRequest) {
    request = &DescribeOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeOIDCConfig")
    
    
    return
}

func NewDescribeOIDCConfigResponse() (response *DescribeOIDCConfigResponse) {
    response = &DescribeOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOIDCConfig
// This API is used to query role OIDC configurations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeOIDCConfig(request *DescribeOIDCConfigRequest) (response *DescribeOIDCConfigResponse, err error) {
    return c.DescribeOIDCConfigWithContext(context.Background(), request)
}

// DescribeOIDCConfig
// This API is used to query role OIDC configurations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeOIDCConfigWithContext(ctx context.Context, request *DescribeOIDCConfigRequest) (response *DescribeOIDCConfigResponse, err error) {
    if request == nil {
        request = NewDescribeOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
    request = &DescribeRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeRoleList")
    
    
    return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
    response = &DescribeRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoleList
// This API (DescribeRoleList) is used to get the role list under the account.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    return c.DescribeRoleListWithContext(context.Background(), request)
}

// DescribeRoleList
// This API (DescribeRoleList) is used to get the role list under the account.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeRoleListWithContext(ctx context.Context, request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeRoleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeAuthFlagCollRequest() (request *DescribeSafeAuthFlagCollRequest) {
    request = &DescribeSafeAuthFlagCollRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeSafeAuthFlagColl")
    
    
    return
}

func NewDescribeSafeAuthFlagCollResponse() (response *DescribeSafeAuthFlagCollResponse) {
    response = &DescribeSafeAuthFlagCollResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSafeAuthFlagColl
// This API is used to query security settings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeSafeAuthFlagColl(request *DescribeSafeAuthFlagCollRequest) (response *DescribeSafeAuthFlagCollResponse, err error) {
    return c.DescribeSafeAuthFlagCollWithContext(context.Background(), request)
}

// DescribeSafeAuthFlagColl
// This API is used to query security settings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeSafeAuthFlagCollWithContext(ctx context.Context, request *DescribeSafeAuthFlagCollRequest) (response *DescribeSafeAuthFlagCollResponse, err error) {
    if request == nil {
        request = NewDescribeSafeAuthFlagCollRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSafeAuthFlagColl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSafeAuthFlagCollResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeAuthFlagIntlRequest() (request *DescribeSafeAuthFlagIntlRequest) {
    request = &DescribeSafeAuthFlagIntlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeSafeAuthFlagIntl")
    
    
    return
}

func NewDescribeSafeAuthFlagIntlResponse() (response *DescribeSafeAuthFlagIntlResponse) {
    response = &DescribeSafeAuthFlagIntlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSafeAuthFlagIntl
// This API is used to query security settings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeSafeAuthFlagIntl(request *DescribeSafeAuthFlagIntlRequest) (response *DescribeSafeAuthFlagIntlResponse, err error) {
    return c.DescribeSafeAuthFlagIntlWithContext(context.Background(), request)
}

// DescribeSafeAuthFlagIntl
// This API is used to query security settings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeSafeAuthFlagIntlWithContext(ctx context.Context, request *DescribeSafeAuthFlagIntlRequest) (response *DescribeSafeAuthFlagIntlResponse, err error) {
    if request == nil {
        request = NewDescribeSafeAuthFlagIntlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSafeAuthFlagIntl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSafeAuthFlagIntlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubAccountsRequest() (request *DescribeSubAccountsRequest) {
    request = &DescribeSubAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeSubAccounts")
    
    
    return
}

func NewDescribeSubAccountsResponse() (response *DescribeSubAccountsResponse) {
    response = &DescribeSubAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubAccounts
// This API is used to query sub-users through the sub-user UIN list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubAccounts(request *DescribeSubAccountsRequest) (response *DescribeSubAccountsResponse, err error) {
    return c.DescribeSubAccountsWithContext(context.Background(), request)
}

// DescribeSubAccounts
// This API is used to query sub-users through the sub-user UIN list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubAccountsWithContext(ctx context.Context, request *DescribeSubAccountsRequest) (response *DescribeSubAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeSubAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserOIDCConfigRequest() (request *DescribeUserOIDCConfigRequest) {
    request = &DescribeUserOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeUserOIDCConfig")
    
    
    return
}

func NewDescribeUserOIDCConfigResponse() (response *DescribeUserOIDCConfigResponse) {
    response = &DescribeUserOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserOIDCConfig
// This API is used to query the user OIDC configuration.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeUserOIDCConfig(request *DescribeUserOIDCConfigRequest) (response *DescribeUserOIDCConfigResponse, err error) {
    return c.DescribeUserOIDCConfigWithContext(context.Background(), request)
}

// DescribeUserOIDCConfig
// This API is used to query the user OIDC configuration.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeUserOIDCConfigWithContext(ctx context.Context, request *DescribeUserOIDCConfigRequest) (response *DescribeUserOIDCConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUserOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserSAMLConfigRequest() (request *DescribeUserSAMLConfigRequest) {
    request = &DescribeUserSAMLConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DescribeUserSAMLConfig")
    
    
    return
}

func NewDescribeUserSAMLConfigResponse() (response *DescribeUserSAMLConfigResponse) {
    response = &DescribeUserSAMLConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserSAMLConfig
// This API is used to query user SAML configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeUserSAMLConfig(request *DescribeUserSAMLConfigRequest) (response *DescribeUserSAMLConfigResponse, err error) {
    return c.DescribeUserSAMLConfigWithContext(context.Background(), request)
}

// DescribeUserSAMLConfig
// This API is used to query user SAML configurations.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) DescribeUserSAMLConfigWithContext(ctx context.Context, request *DescribeUserSAMLConfigRequest) (response *DescribeUserSAMLConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUserSAMLConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserSAMLConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserSAMLConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDetachGroupPolicyRequest() (request *DetachGroupPolicyRequest) {
    request = &DetachGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DetachGroupPolicy")
    
    
    return
}

func NewDetachGroupPolicyResponse() (response *DetachGroupPolicyResponse) {
    response = &DetachGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachGroupPolicy
// This API (DetachGroupPolicy) is used to unassociate a policy and a user group.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DetachGroupPolicy(request *DetachGroupPolicyRequest) (response *DetachGroupPolicyResponse, err error) {
    return c.DetachGroupPolicyWithContext(context.Background(), request)
}

// DetachGroupPolicy
// This API (DetachGroupPolicy) is used to unassociate a policy and a user group.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DetachGroupPolicyWithContext(ctx context.Context, request *DetachGroupPolicyRequest) (response *DetachGroupPolicyResponse, err error) {
    if request == nil {
        request = NewDetachGroupPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachRolePolicyRequest() (request *DetachRolePolicyRequest) {
    request = &DetachRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DetachRolePolicy")
    
    
    return
}

func NewDetachRolePolicyResponse() (response *DetachRolePolicyResponse) {
    response = &DetachRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachRolePolicy
// This API (DetachRolePolicy) is used to unassociate a policy and a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) DetachRolePolicy(request *DetachRolePolicyRequest) (response *DetachRolePolicyResponse, err error) {
    return c.DetachRolePolicyWithContext(context.Background(), request)
}

// DetachRolePolicy
// This API (DetachRolePolicy) is used to unassociate a policy and a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) DetachRolePolicyWithContext(ctx context.Context, request *DetachRolePolicyRequest) (response *DetachRolePolicyResponse, err error) {
    if request == nil {
        request = NewDetachRolePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachRolePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachUserPolicyRequest() (request *DetachUserPolicyRequest) {
    request = &DetachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DetachUserPolicy")
    
    
    return
}

func NewDetachUserPolicyResponse() (response *DetachUserPolicyResponse) {
    response = &DetachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachUserPolicy
// This API (DetachUserPolicy) is used to unassociate a policy and a user.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DetachUserPolicy(request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    return c.DetachUserPolicyWithContext(context.Background(), request)
}

// DetachUserPolicy
// This API (DetachUserPolicy) is used to unassociate a policy and a user.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DetachUserPolicyWithContext(ctx context.Context, request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    if request == nil {
        request = NewDetachUserPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachUserPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUserSSORequest() (request *DisableUserSSORequest) {
    request = &DisableUserSSORequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "DisableUserSSO")
    
    
    return
}

func NewDisableUserSSOResponse() (response *DisableUserSSOResponse) {
    response = &DisableUserSSOResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableUserSSO
// This API is used to disable user SSO.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
func (c *Client) DisableUserSSO(request *DisableUserSSORequest) (response *DisableUserSSOResponse, err error) {
    return c.DisableUserSSOWithContext(context.Background(), request)
}

// DisableUserSSO
// This API is used to disable user SSO.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
func (c *Client) DisableUserSSOWithContext(ctx context.Context, request *DisableUserSSORequest) (response *DisableUserSSOResponse, err error) {
    if request == nil {
        request = NewDisableUserSSORequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableUserSSO require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableUserSSOResponse()
    err = c.Send(request, response)
    return
}

func NewGetAccountSummaryRequest() (request *GetAccountSummaryRequest) {
    request = &GetAccountSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetAccountSummary")
    
    
    return
}

func NewGetAccountSummaryResponse() (response *GetAccountSummaryResponse) {
    response = &GetAccountSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAccountSummary
// This API is used to query account summary. 
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) GetAccountSummary(request *GetAccountSummaryRequest) (response *GetAccountSummaryResponse, err error) {
    return c.GetAccountSummaryWithContext(context.Background(), request)
}

// GetAccountSummary
// This API is used to query account summary. 
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) GetAccountSummaryWithContext(ctx context.Context, request *GetAccountSummaryRequest) (response *GetAccountSummaryResponse, err error) {
    if request == nil {
        request = NewGetAccountSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAccountSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAccountSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetCustomMFATokenInfoRequest() (request *GetCustomMFATokenInfoRequest) {
    request = &GetCustomMFATokenInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetCustomMFATokenInfo")
    
    
    return
}

func NewGetCustomMFATokenInfoResponse() (response *GetCustomMFATokenInfoResponse) {
    response = &GetCustomMFATokenInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCustomMFATokenInfo
// This API is used to get information associated with a custom multi-factor Token
//
// error code that may be returned:
//  INVALIDPARAMETER_MFATOKENERROR = "InvalidParameter.MFATokenError"
func (c *Client) GetCustomMFATokenInfo(request *GetCustomMFATokenInfoRequest) (response *GetCustomMFATokenInfoResponse, err error) {
    return c.GetCustomMFATokenInfoWithContext(context.Background(), request)
}

// GetCustomMFATokenInfo
// This API is used to get information associated with a custom multi-factor Token
//
// error code that may be returned:
//  INVALIDPARAMETER_MFATOKENERROR = "InvalidParameter.MFATokenError"
func (c *Client) GetCustomMFATokenInfoWithContext(ctx context.Context, request *GetCustomMFATokenInfoRequest) (response *GetCustomMFATokenInfoResponse, err error) {
    if request == nil {
        request = NewGetCustomMFATokenInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCustomMFATokenInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCustomMFATokenInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupRequest() (request *GetGroupRequest) {
    request = &GetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetGroup")
    
    
    return
}

func NewGetGroupResponse() (response *GetGroupResponse) {
    response = &GetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGroup
// This API is used to query user group details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
    return c.GetGroupWithContext(context.Background(), request)
}

// GetGroup
// This API is used to query user group details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

func NewGetPolicyRequest() (request *GetPolicyRequest) {
    request = &GetPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetPolicy")
    
    
    return
}

func NewGetPolicyResponse() (response *GetPolicyResponse) {
    response = &GetPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPolicy
// This API (GetPolicy) is used to query and view policy details.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) GetPolicy(request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
    return c.GetPolicyWithContext(context.Background(), request)
}

// GetPolicy
// This API (GetPolicy) is used to query and view policy details.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) GetPolicyWithContext(ctx context.Context, request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
    if request == nil {
        request = NewGetPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewGetPolicyVersionRequest() (request *GetPolicyVersionRequest) {
    request = &GetPolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetPolicyVersion")
    
    
    return
}

func NewGetPolicyVersionResponse() (response *GetPolicyVersionResponse) {
    response = &GetPolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPolicyVersion
// This API is used to query policy version details.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYVERSIONNOTEXISTS = "InvalidParameter.PolicyVersionNotExists"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) GetPolicyVersion(request *GetPolicyVersionRequest) (response *GetPolicyVersionResponse, err error) {
    return c.GetPolicyVersionWithContext(context.Background(), request)
}

// GetPolicyVersion
// This API is used to query policy version details.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYVERSIONNOTEXISTS = "InvalidParameter.PolicyVersionNotExists"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) GetPolicyVersionWithContext(ctx context.Context, request *GetPolicyVersionRequest) (response *GetPolicyVersionResponse, err error) {
    if request == nil {
        request = NewGetPolicyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPolicyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoleRequest() (request *GetRoleRequest) {
    request = &GetRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetRole")
    
    
    return
}

func NewGetRoleResponse() (response *GetRoleResponse) {
    response = &GetRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRole
// This API (GetRole) is used to get the details of a specified role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) GetRole(request *GetRoleRequest) (response *GetRoleResponse, err error) {
    return c.GetRoleWithContext(context.Background(), request)
}

// GetRole
// This API (GetRole) is used to get the details of a specified role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) GetRoleWithContext(ctx context.Context, request *GetRoleRequest) (response *GetRoleResponse, err error) {
    if request == nil {
        request = NewGetRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoleResponse()
    err = c.Send(request, response)
    return
}

func NewGetSAMLProviderRequest() (request *GetSAMLProviderRequest) {
    request = &GetSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetSAMLProvider")
    
    
    return
}

func NewGetSAMLProviderResponse() (response *GetSAMLProviderResponse) {
    response = &GetSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSAMLProvider
// This API is used to query SAML identity provider details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) GetSAMLProvider(request *GetSAMLProviderRequest) (response *GetSAMLProviderResponse, err error) {
    return c.GetSAMLProviderWithContext(context.Background(), request)
}

// GetSAMLProvider
// This API is used to query SAML identity provider details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) GetSAMLProviderWithContext(ctx context.Context, request *GetSAMLProviderRequest) (response *GetSAMLProviderResponse, err error) {
    if request == nil {
        request = NewGetSAMLProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSAMLProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewGetSecurityLastUsedRequest() (request *GetSecurityLastUsedRequest) {
    request = &GetSecurityLastUsedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetSecurityLastUsed")
    
    
    return
}

func NewGetSecurityLastUsedResponse() (response *GetSecurityLastUsedResponse) {
    response = &GetSecurityLastUsedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSecurityLastUsed
// This API is used to get a key’s recent usage details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetSecurityLastUsed(request *GetSecurityLastUsedRequest) (response *GetSecurityLastUsedResponse, err error) {
    return c.GetSecurityLastUsedWithContext(context.Background(), request)
}

// GetSecurityLastUsed
// This API is used to get a key’s recent usage details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetSecurityLastUsedWithContext(ctx context.Context, request *GetSecurityLastUsedRequest) (response *GetSecurityLastUsedResponse, err error) {
    if request == nil {
        request = NewGetSecurityLastUsedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSecurityLastUsed require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSecurityLastUsedResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceLinkedRoleDeletionStatusRequest() (request *GetServiceLinkedRoleDeletionStatusRequest) {
    request = &GetServiceLinkedRoleDeletionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetServiceLinkedRoleDeletionStatus")
    
    
    return
}

func NewGetServiceLinkedRoleDeletionStatusResponse() (response *GetServiceLinkedRoleDeletionStatusResponse) {
    response = &GetServiceLinkedRoleDeletionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetServiceLinkedRoleDeletionStatus
// This API is used to get the status of the service-linked role deletion based on the `TaskId`
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DELETIONTASKNOTEXIST = "InvalidParameter.DeletionTaskNotExist"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) GetServiceLinkedRoleDeletionStatus(request *GetServiceLinkedRoleDeletionStatusRequest) (response *GetServiceLinkedRoleDeletionStatusResponse, err error) {
    return c.GetServiceLinkedRoleDeletionStatusWithContext(context.Background(), request)
}

// GetServiceLinkedRoleDeletionStatus
// This API is used to get the status of the service-linked role deletion based on the `TaskId`
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DELETIONTASKNOTEXIST = "InvalidParameter.DeletionTaskNotExist"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) GetServiceLinkedRoleDeletionStatusWithContext(ctx context.Context, request *GetServiceLinkedRoleDeletionStatusRequest) (response *GetServiceLinkedRoleDeletionStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceLinkedRoleDeletionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetServiceLinkedRoleDeletionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetServiceLinkedRoleDeletionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserRequest() (request *GetUserRequest) {
    request = &GetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetUser")
    
    
    return
}

func NewGetUserResponse() (response *GetUserResponse) {
    response = &GetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUser
// This API is used to query sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
    return c.GetUserWithContext(context.Background(), request)
}

// GetUser
// This API is used to query sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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

func NewGetUserAppIdRequest() (request *GetUserAppIdRequest) {
    request = &GetUserAppIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "GetUserAppId")
    
    
    return
}

func NewGetUserAppIdResponse() (response *GetUserAppIdResponse) {
    response = &GetUserAppIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUserAppId
// This API is used to get the user AppId.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) GetUserAppId(request *GetUserAppIdRequest) (response *GetUserAppIdResponse, err error) {
    return c.GetUserAppIdWithContext(context.Background(), request)
}

// GetUserAppId
// This API is used to get the user AppId.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) GetUserAppIdWithContext(ctx context.Context, request *GetUserAppIdRequest) (response *GetUserAppIdResponse, err error) {
    if request == nil {
        request = NewGetUserAppIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUserAppId require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserAppIdResponse()
    err = c.Send(request, response)
    return
}

func NewListAccessKeysRequest() (request *ListAccessKeysRequest) {
    request = &ListAccessKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListAccessKeys")
    
    
    return
}

func NewListAccessKeysResponse() (response *ListAccessKeysResponse) {
    response = &ListAccessKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAccessKeys
// This API is used to list the access keys associated with a specified CAM user.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ACCESSKEY = "FailedOperation.Accesskey"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  OPERATIONDENIED_ACCESSKEYOVERLIMIT = "OperationDenied.AccessKeyOverLimit"
//  OPERATIONDENIED_SUBUIN = "OperationDenied.SubUin"
//  OPERATIONDENIED_UINNOTMATCH = "OperationDenied.UinNotMatch"
func (c *Client) ListAccessKeys(request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
    return c.ListAccessKeysWithContext(context.Background(), request)
}

// ListAccessKeys
// This API is used to list the access keys associated with a specified CAM user.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ACCESSKEY = "FailedOperation.Accesskey"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  OPERATIONDENIED_ACCESSKEYOVERLIMIT = "OperationDenied.AccessKeyOverLimit"
//  OPERATIONDENIED_SUBUIN = "OperationDenied.SubUin"
//  OPERATIONDENIED_UINNOTMATCH = "OperationDenied.UinNotMatch"
func (c *Client) ListAccessKeysWithContext(ctx context.Context, request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
    if request == nil {
        request = NewListAccessKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAccessKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAccessKeysResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedGroupPoliciesRequest() (request *ListAttachedGroupPoliciesRequest) {
    request = &ListAttachedGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedGroupPolicies")
    
    
    return
}

func NewListAttachedGroupPoliciesResponse() (response *ListAttachedGroupPoliciesResponse) {
    response = &ListAttachedGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAttachedGroupPolicies
// This API (ListAttachedGroupPolicies) is used to query the list of policies associated with a user group.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListAttachedGroupPolicies(request *ListAttachedGroupPoliciesRequest) (response *ListAttachedGroupPoliciesResponse, err error) {
    return c.ListAttachedGroupPoliciesWithContext(context.Background(), request)
}

// ListAttachedGroupPolicies
// This API (ListAttachedGroupPolicies) is used to query the list of policies associated with a user group.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListAttachedGroupPoliciesWithContext(ctx context.Context, request *ListAttachedGroupPoliciesRequest) (response *ListAttachedGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedGroupPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttachedGroupPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttachedGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedRolePoliciesRequest() (request *ListAttachedRolePoliciesRequest) {
    request = &ListAttachedRolePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedRolePolicies")
    
    
    return
}

func NewListAttachedRolePoliciesResponse() (response *ListAttachedRolePoliciesResponse) {
    response = &ListAttachedRolePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAttachedRolePolicies
// This API (ListAttachedRolePolicies) is used to obtain the list of the policies associated with a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListAttachedRolePolicies(request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
    return c.ListAttachedRolePoliciesWithContext(context.Background(), request)
}

// ListAttachedRolePolicies
// This API (ListAttachedRolePolicies) is used to obtain the list of the policies associated with a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListAttachedRolePoliciesWithContext(ctx context.Context, request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedRolePoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttachedRolePolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttachedRolePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedUserAllPoliciesRequest() (request *ListAttachedUserAllPoliciesRequest) {
    request = &ListAttachedUserAllPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedUserAllPolicies")
    
    
    return
}

func NewListAttachedUserAllPoliciesResponse() (response *ListAttachedUserAllPoliciesResponse) {
    response = &ListAttachedUserAllPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAttachedUserAllPolicies
// This API is used to list policies associated with the user (including those inherited from the user group).
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAttachedUserAllPolicies(request *ListAttachedUserAllPoliciesRequest) (response *ListAttachedUserAllPoliciesResponse, err error) {
    return c.ListAttachedUserAllPoliciesWithContext(context.Background(), request)
}

// ListAttachedUserAllPolicies
// This API is used to list policies associated with the user (including those inherited from the user group).
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAttachedUserAllPoliciesWithContext(ctx context.Context, request *ListAttachedUserAllPoliciesRequest) (response *ListAttachedUserAllPoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedUserAllPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttachedUserAllPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttachedUserAllPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedUserPoliciesRequest() (request *ListAttachedUserPoliciesRequest) {
    request = &ListAttachedUserPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedUserPolicies")
    
    
    return
}

func NewListAttachedUserPoliciesResponse() (response *ListAttachedUserPoliciesResponse) {
    response = &ListAttachedUserPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAttachedUserPolicies
// This API (ListAttachedUserPolicies) is used to query the list of policies associated with a sub-account.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListAttachedUserPolicies(request *ListAttachedUserPoliciesRequest) (response *ListAttachedUserPoliciesResponse, err error) {
    return c.ListAttachedUserPoliciesWithContext(context.Background(), request)
}

// ListAttachedUserPolicies
// This API (ListAttachedUserPolicies) is used to query the list of policies associated with a sub-account.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListAttachedUserPoliciesWithContext(ctx context.Context, request *ListAttachedUserPoliciesRequest) (response *ListAttachedUserPoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedUserPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttachedUserPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttachedUserPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListCollaboratorsRequest() (request *ListCollaboratorsRequest) {
    request = &ListCollaboratorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListCollaborators")
    
    
    return
}

func NewListCollaboratorsResponse() (response *ListCollaboratorsResponse) {
    response = &ListCollaboratorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListCollaborators
// This API is used to get the collaborator list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCollaborators(request *ListCollaboratorsRequest) (response *ListCollaboratorsResponse, err error) {
    return c.ListCollaboratorsWithContext(context.Background(), request)
}

// ListCollaborators
// This API is used to get the collaborator list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCollaboratorsWithContext(ctx context.Context, request *ListCollaboratorsRequest) (response *ListCollaboratorsResponse, err error) {
    if request == nil {
        request = NewListCollaboratorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCollaborators require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCollaboratorsResponse()
    err = c.Send(request, response)
    return
}

func NewListEntitiesForPolicyRequest() (request *ListEntitiesForPolicyRequest) {
    request = &ListEntitiesForPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListEntitiesForPolicy")
    
    
    return
}

func NewListEntitiesForPolicyResponse() (response *ListEntitiesForPolicyResponse) {
    response = &ListEntitiesForPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListEntitiesForPolicy
// This API (ListEntitiesForPolicy) is used to query the list of entities associated with a policy.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ENTITYFILTERERROR = "InvalidParameter.EntityFilterError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
func (c *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
    return c.ListEntitiesForPolicyWithContext(context.Background(), request)
}

// ListEntitiesForPolicy
// This API (ListEntitiesForPolicy) is used to query the list of entities associated with a policy.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ENTITYFILTERERROR = "InvalidParameter.EntityFilterError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
func (c *Client) ListEntitiesForPolicyWithContext(ctx context.Context, request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
    if request == nil {
        request = NewListEntitiesForPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListEntitiesForPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewListEntitiesForPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupsRequest() (request *ListGroupsRequest) {
    request = &ListGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListGroups")
    
    
    return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
    response = &ListGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListGroups
// This API is used to query the list of user groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    return c.ListGroupsWithContext(context.Background(), request)
}

// ListGroups
// This API is used to query the list of user groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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

func NewListGroupsForUserRequest() (request *ListGroupsForUserRequest) {
    request = &ListGroupsForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListGroupsForUser")
    
    
    return
}

func NewListGroupsForUserResponse() (response *ListGroupsForUserResponse) {
    response = &ListGroupsForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListGroupsForUser
// This API is used to list user groups associated with a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (response *ListGroupsForUserResponse, err error) {
    return c.ListGroupsForUserWithContext(context.Background(), request)
}

// ListGroupsForUser
// This API is used to list user groups associated with a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListGroupsForUserWithContext(ctx context.Context, request *ListGroupsForUserRequest) (response *ListGroupsForUserResponse, err error) {
    if request == nil {
        request = NewListGroupsForUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGroupsForUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGroupsForUserResponse()
    err = c.Send(request, response)
    return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
    request = &ListPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListPolicies")
    
    
    return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
    response = &ListPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListPolicies
// This API is used to query the policy list.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_GROUPIDERROR = "InvalidParameter.GroupIdError"
//  INVALIDPARAMETER_KEYWORDERROR = "InvalidParameter.KeywordError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SCOPEERROR = "InvalidParameter.ScopeError"
//  INVALIDPARAMETER_SERVICETYPEERROR = "InvalidParameter.ServiceTypeError"
//  INVALIDPARAMETER_UINERROR = "InvalidParameter.UinError"
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    return c.ListPoliciesWithContext(context.Background(), request)
}

// ListPolicies
// This API is used to query the policy list.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_GROUPIDERROR = "InvalidParameter.GroupIdError"
//  INVALIDPARAMETER_KEYWORDERROR = "InvalidParameter.KeywordError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SCOPEERROR = "InvalidParameter.ScopeError"
//  INVALIDPARAMETER_SERVICETYPEERROR = "InvalidParameter.ServiceTypeError"
//  INVALIDPARAMETER_UINERROR = "InvalidParameter.UinError"
func (c *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    if request == nil {
        request = NewListPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListPolicyVersionsRequest() (request *ListPolicyVersionsRequest) {
    request = &ListPolicyVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListPolicyVersions")
    
    
    return
}

func NewListPolicyVersionsResponse() (response *ListPolicyVersionsResponse) {
    response = &ListPolicyVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListPolicyVersions
// This API is used to get the list of policy versions.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (response *ListPolicyVersionsResponse, err error) {
    return c.ListPolicyVersionsWithContext(context.Background(), request)
}

// ListPolicyVersions
// This API is used to get the list of policy versions.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) ListPolicyVersionsWithContext(ctx context.Context, request *ListPolicyVersionsRequest) (response *ListPolicyVersionsResponse, err error) {
    if request == nil {
        request = NewListPolicyVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPolicyVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPolicyVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListSAMLProvidersRequest() (request *ListSAMLProvidersRequest) {
    request = &ListSAMLProvidersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListSAMLProviders")
    
    
    return
}

func NewListSAMLProvidersResponse() (response *ListSAMLProvidersResponse) {
    response = &ListSAMLProvidersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSAMLProviders
// This API is used to query the list of SAML identity providers.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) ListSAMLProviders(request *ListSAMLProvidersRequest) (response *ListSAMLProvidersResponse, err error) {
    return c.ListSAMLProvidersWithContext(context.Background(), request)
}

// ListSAMLProviders
// This API is used to query the list of SAML identity providers.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
func (c *Client) ListSAMLProvidersWithContext(ctx context.Context, request *ListSAMLProvidersRequest) (response *ListSAMLProvidersResponse, err error) {
    if request == nil {
        request = NewListSAMLProvidersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSAMLProviders require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSAMLProvidersResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListUsers")
    
    
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUsers
// This API is used to pull sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    return c.ListUsersWithContext(context.Background(), request)
}

// ListUsers
// This API is used to pull sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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

func NewListUsersForGroupRequest() (request *ListUsersForGroupRequest) {
    request = &ListUsersForGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "ListUsersForGroup")
    
    
    return
}

func NewListUsersForGroupResponse() (response *ListUsersForGroupResponse) {
    response = &ListUsersForGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUsersForGroup
// This API is used to query the list of users associated with a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (response *ListUsersForGroupResponse, err error) {
    return c.ListUsersForGroupWithContext(context.Background(), request)
}

// ListUsersForGroup
// This API is used to query the list of users associated with a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ListUsersForGroupWithContext(ctx context.Context, request *ListUsersForGroupRequest) (response *ListUsersForGroupResponse, err error) {
    if request == nil {
        request = NewListUsersForGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsersForGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersForGroupResponse()
    err = c.Send(request, response)
    return
}

func NewPutRolePermissionsBoundaryRequest() (request *PutRolePermissionsBoundaryRequest) {
    request = &PutRolePermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "PutRolePermissionsBoundary")
    
    
    return
}

func NewPutRolePermissionsBoundaryResponse() (response *PutRolePermissionsBoundaryResponse) {
    response = &PutRolePermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutRolePermissionsBoundary
// This API is used to set a role permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_SERVICELINKEDPOLICYCANTINPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedPolicyCantInPermissionBoundary"
//  INVALIDPARAMETER_SERVICELINKEDROLECANTUSEPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedRoleCantUsePermissionBoundary"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutRolePermissionsBoundary(request *PutRolePermissionsBoundaryRequest) (response *PutRolePermissionsBoundaryResponse, err error) {
    return c.PutRolePermissionsBoundaryWithContext(context.Background(), request)
}

// PutRolePermissionsBoundary
// This API is used to set a role permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_SERVICELINKEDPOLICYCANTINPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedPolicyCantInPermissionBoundary"
//  INVALIDPARAMETER_SERVICELINKEDROLECANTUSEPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedRoleCantUsePermissionBoundary"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutRolePermissionsBoundaryWithContext(ctx context.Context, request *PutRolePermissionsBoundaryRequest) (response *PutRolePermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewPutRolePermissionsBoundaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutRolePermissionsBoundary require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutRolePermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewPutUserPermissionsBoundaryRequest() (request *PutUserPermissionsBoundaryRequest) {
    request = &PutUserPermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "PutUserPermissionsBoundary")
    
    
    return
}

func NewPutUserPermissionsBoundaryResponse() (response *PutUserPermissionsBoundaryResponse) {
    response = &PutUserPermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutUserPermissionsBoundary
// This API is used to set a user permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_SERVICELINKEDPOLICYCANTINPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedPolicyCantInPermissionBoundary"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutUserPermissionsBoundary(request *PutUserPermissionsBoundaryRequest) (response *PutUserPermissionsBoundaryResponse, err error) {
    return c.PutUserPermissionsBoundaryWithContext(context.Background(), request)
}

// PutUserPermissionsBoundary
// This API is used to set a user permission boundary.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_OPERATEENTITIESOVERLIMIT = "InvalidParameter.OperateEntitiesOverLimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_SERVICELINKEDPOLICYCANTINPERMISSIONBOUNDARY = "InvalidParameter.ServiceLinkedPolicyCantInPermissionBoundary"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) PutUserPermissionsBoundaryWithContext(ctx context.Context, request *PutUserPermissionsBoundaryRequest) (response *PutUserPermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewPutUserPermissionsBoundaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutUserPermissionsBoundary require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutUserPermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
    request = &RemoveUserFromGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "RemoveUserFromGroup")
    
    
    return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
    response = &RemoveUserFromGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveUserFromGroup
// This API is used to delete users from a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_USERUINANDUINNOTALLNULL = "InvalidParameter.UserUinAndUinNotAllNull"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    return c.RemoveUserFromGroupWithContext(context.Background(), request)
}

// RemoveUserFromGroup
// This API is used to delete users from a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_USERUINANDUINNOTALLNULL = "InvalidParameter.UserUinAndUinNotAllNull"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

func NewSetDefaultPolicyVersionRequest() (request *SetDefaultPolicyVersionRequest) {
    request = &SetDefaultPolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "SetDefaultPolicyVersion")
    
    
    return
}

func NewSetDefaultPolicyVersionResponse() (response *SetDefaultPolicyVersionResponse) {
    response = &SetDefaultPolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetDefaultPolicyVersion
// This API is used to set the operative policy version.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  FAILEDOPERATION_POLICYVERSIONALREADYDEFAULT = "FailedOperation.PolicyVersionAlreadyDefault"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYVERSIONNOTEXISTS = "InvalidParameter.PolicyVersionNotExists"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SetDefaultPolicyVersion(request *SetDefaultPolicyVersionRequest) (response *SetDefaultPolicyVersionResponse, err error) {
    return c.SetDefaultPolicyVersionWithContext(context.Background(), request)
}

// SetDefaultPolicyVersion
// This API is used to set the operative policy version.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  FAILEDOPERATION_POLICYVERSIONALREADYDEFAULT = "FailedOperation.PolicyVersionAlreadyDefault"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYVERSIONNOTEXISTS = "InvalidParameter.PolicyVersionNotExists"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SetDefaultPolicyVersionWithContext(ctx context.Context, request *SetDefaultPolicyVersionRequest) (response *SetDefaultPolicyVersionResponse, err error) {
    if request == nil {
        request = NewSetDefaultPolicyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDefaultPolicyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDefaultPolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewSetMfaFlagRequest() (request *SetMfaFlagRequest) {
    request = &SetMfaFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "SetMfaFlag")
    
    
    return
}

func NewSetMfaFlagResponse() (response *SetMfaFlagResponse) {
    response = &SetMfaFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetMfaFlag
// This API is used to set account verification for login and sensitive operations for sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION_USERNOTBINDPHONE = "FailedOperation.UserNotBindPhone"
//  FAILEDOPERATION_USERNOTBINDWECHAT = "FailedOperation.UserNotBindWechat"
//  FAILEDOPERATION_USERUNBINDNOPERMISSION = "FailedOperation.UserUnbindNoPermission"
//  INVALIDPARAMETER_MFATOKENERROR = "InvalidParameter.MFATokenError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SetMfaFlag(request *SetMfaFlagRequest) (response *SetMfaFlagResponse, err error) {
    return c.SetMfaFlagWithContext(context.Background(), request)
}

// SetMfaFlag
// This API is used to set account verification for login and sensitive operations for sub-users.
//
// error code that may be returned:
//  FAILEDOPERATION_USERNOTBINDPHONE = "FailedOperation.UserNotBindPhone"
//  FAILEDOPERATION_USERNOTBINDWECHAT = "FailedOperation.UserNotBindWechat"
//  FAILEDOPERATION_USERUNBINDNOPERMISSION = "FailedOperation.UserUnbindNoPermission"
//  INVALIDPARAMETER_MFATOKENERROR = "InvalidParameter.MFATokenError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SetMfaFlagWithContext(ctx context.Context, request *SetMfaFlagRequest) (response *SetMfaFlagResponse, err error) {
    if request == nil {
        request = NewSetMfaFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetMfaFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetMfaFlagResponse()
    err = c.Send(request, response)
    return
}

func NewTagRoleRequest() (request *TagRoleRequest) {
    request = &TagRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "TagRole")
    
    
    return
}

func NewTagRoleResponse() (response *TagRoleResponse) {
    response = &TagRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TagRole
// This API is used to bind tags to a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_TAGRESOURCEFAILED = "FailedOperation.TagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_TAGLIMITEXCEEDED = "InvalidParameter.TagLimitExceeded"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
func (c *Client) TagRole(request *TagRoleRequest) (response *TagRoleResponse, err error) {
    return c.TagRoleWithContext(context.Background(), request)
}

// TagRole
// This API is used to bind tags to a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_TAGRESOURCEFAILED = "FailedOperation.TagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_TAGLIMITEXCEEDED = "InvalidParameter.TagLimitExceeded"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
func (c *Client) TagRoleWithContext(ctx context.Context, request *TagRoleRequest) (response *TagRoleResponse, err error) {
    if request == nil {
        request = NewTagRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TagRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewTagRoleResponse()
    err = c.Send(request, response)
    return
}

func NewUntagRoleRequest() (request *UntagRoleRequest) {
    request = &UntagRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UntagRole")
    
    
    return
}

func NewUntagRoleResponse() (response *UntagRoleResponse) {
    response = &UntagRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UntagRole
// This API is used to unbind tags from a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UnTagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENAMEERROR = "InvalidParameter.RoleNameError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
func (c *Client) UntagRole(request *UntagRoleRequest) (response *UntagRoleResponse, err error) {
    return c.UntagRoleWithContext(context.Background(), request)
}

// UntagRole
// This API is used to unbind tags from a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UnTagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENAMEERROR = "InvalidParameter.RoleNameError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_TAGPARAMERROR = "InvalidParameter.TagParamError"
func (c *Client) UntagRoleWithContext(ctx context.Context, request *UntagRoleRequest) (response *UntagRoleResponse, err error) {
    if request == nil {
        request = NewUntagRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UntagRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewUntagRoleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssumeRolePolicyRequest() (request *UpdateAssumeRolePolicyRequest) {
    request = &UpdateAssumeRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateAssumeRolePolicy")
    
    
    return
}

func NewUpdateAssumeRolePolicyResponse() (response *UpdateAssumeRolePolicyResponse) {
    response = &UpdateAssumeRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAssumeRolePolicy
// This API (UpdateAssumeRolePolicy) is used to modify the trust policy of a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSCROSSERROR = "InvalidParameter.PrincipalQcsCrossError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
func (c *Client) UpdateAssumeRolePolicy(request *UpdateAssumeRolePolicyRequest) (response *UpdateAssumeRolePolicyResponse, err error) {
    return c.UpdateAssumeRolePolicyWithContext(context.Background(), request)
}

// UpdateAssumeRolePolicy
// This API (UpdateAssumeRolePolicy) is used to modify the trust policy of a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSCROSSERROR = "InvalidParameter.PrincipalQcsCrossError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
func (c *Client) UpdateAssumeRolePolicyWithContext(ctx context.Context, request *UpdateAssumeRolePolicyRequest) (response *UpdateAssumeRolePolicyResponse, err error) {
    if request == nil {
        request = NewUpdateAssumeRolePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAssumeRolePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAssumeRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
    request = &UpdateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateGroup")
    
    
    return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
    response = &UpdateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateGroup
// This API is used to update a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_GROUPNAMEINUSE = "InvalidParameter.GroupNameInUse"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    return c.UpdateGroupWithContext(context.Background(), request)
}

// UpdateGroup
// This API is used to update a user group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_GROUPNAMEINUSE = "InvalidParameter.GroupNameInUse"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
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

func NewUpdateOIDCConfigRequest() (request *UpdateOIDCConfigRequest) {
    request = &UpdateOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateOIDCConfig")
    
    
    return
}

func NewUpdateOIDCConfigResponse() (response *UpdateOIDCConfigResponse) {
    response = &UpdateOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateOIDCConfig
// This API is used to modify role OIDC configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateOIDCConfig(request *UpdateOIDCConfigRequest) (response *UpdateOIDCConfigResponse, err error) {
    return c.UpdateOIDCConfigWithContext(context.Background(), request)
}

// UpdateOIDCConfig
// This API is used to modify role OIDC configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  INVALIDPARAMETERVALUE_NAMEERROR = "InvalidParameterValue.NameError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateOIDCConfigWithContext(ctx context.Context, request *UpdateOIDCConfigRequest) (response *UpdateOIDCConfigResponse, err error) {
    if request == nil {
        request = NewUpdateOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdatePolicy")
    
    
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePolicy
// This API is used to update a policy.
//
// This API will update the default version of an existing policy instead of creating a new one. If no policy exists, a default version will be created.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    return c.UpdatePolicyWithContext(context.Background(), request)
}

// UpdatePolicy
// This API is used to update a policy.
//
// This API will update the default version of an existing policy instead of creating a new one. If no policy exists, a default version will be created.
//
// error code that may be returned:
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDERROR = "InvalidParameter.PolicyIdError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdatePolicyWithContext(ctx context.Context, request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleConsoleLoginRequest() (request *UpdateRoleConsoleLoginRequest) {
    request = &UpdateRoleConsoleLoginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleConsoleLogin")
    
    
    return
}

func NewUpdateRoleConsoleLoginResponse() (response *UpdateRoleConsoleLoginResponse) {
    response = &UpdateRoleConsoleLoginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRoleConsoleLogin
// This API is used to modify a role's login permissions.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) UpdateRoleConsoleLogin(request *UpdateRoleConsoleLoginRequest) (response *UpdateRoleConsoleLoginResponse, err error) {
    return c.UpdateRoleConsoleLoginWithContext(context.Background(), request)
}

// UpdateRoleConsoleLogin
// This API is used to modify a role's login permissions.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) UpdateRoleConsoleLoginWithContext(ctx context.Context, request *UpdateRoleConsoleLoginRequest) (response *UpdateRoleConsoleLoginResponse, err error) {
    if request == nil {
        request = NewUpdateRoleConsoleLoginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRoleConsoleLogin require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRoleConsoleLoginResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleDescriptionRequest() (request *UpdateRoleDescriptionRequest) {
    request = &UpdateRoleDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleDescription")
    
    
    return
}

func NewUpdateRoleDescriptionResponse() (response *UpdateRoleDescriptionResponse) {
    response = &UpdateRoleDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRoleDescription
// This API (UpdateRoleDescription) is used to modify the description of a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) UpdateRoleDescription(request *UpdateRoleDescriptionRequest) (response *UpdateRoleDescriptionResponse, err error) {
    return c.UpdateRoleDescriptionWithContext(context.Background(), request)
}

// UpdateRoleDescription
// This API (UpdateRoleDescription) is used to modify the description of a role.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLENOTEXIST = "InvalidParameter.RoleNotExist"
func (c *Client) UpdateRoleDescriptionWithContext(ctx context.Context, request *UpdateRoleDescriptionRequest) (response *UpdateRoleDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateRoleDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRoleDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRoleDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSAMLProviderRequest() (request *UpdateSAMLProviderRequest) {
    request = &UpdateSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateSAMLProvider")
    
    
    return
}

func NewUpdateSAMLProviderResponse() (response *UpdateSAMLProviderResponse) {
    response = &UpdateSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateSAMLProvider
// This API is used to update SAML identity provider information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateSAMLProvider(request *UpdateSAMLProviderRequest) (response *UpdateSAMLProviderResponse, err error) {
    return c.UpdateSAMLProviderWithContext(context.Background(), request)
}

// UpdateSAMLProvider
// This API is used to update SAML identity provider information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateSAMLProviderWithContext(ctx context.Context, request *UpdateSAMLProviderRequest) (response *UpdateSAMLProviderResponse, err error) {
    if request == nil {
        request = NewUpdateSAMLProviderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSAMLProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateUser")
    
    
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUser
// This API is used to update a sub-user.
//
// error code that may be returned:
//  INVALIDPARAMETER_PASSWORDVIOLATEDRULES = "InvalidParameter.PasswordViolatedRules"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    return c.UpdateUserWithContext(context.Background(), request)
}

// UpdateUser
// This API is used to update a sub-user.
//
// error code that may be returned:
//  INVALIDPARAMETER_PASSWORDVIOLATEDRULES = "InvalidParameter.PasswordViolatedRules"
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

func NewUpdateUserOIDCConfigRequest() (request *UpdateUserOIDCConfigRequest) {
    request = &UpdateUserOIDCConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateUserOIDCConfig")
    
    
    return
}

func NewUpdateUserOIDCConfigResponse() (response *UpdateUserOIDCConfigResponse) {
    response = &UpdateUserOIDCConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUserOIDCConfig
// This API is used to modify the user OIDC configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateUserOIDCConfig(request *UpdateUserOIDCConfigRequest) (response *UpdateUserOIDCConfigResponse, err error) {
    return c.UpdateUserOIDCConfigWithContext(context.Background(), request)
}

// UpdateUserOIDCConfig
// This API is used to modify the user OIDC configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IDENTITYNAMEINUSE = "InvalidParameter.IdentityNameInUse"
//  INVALIDPARAMETERVALUE_IDENTITYKEYERROR = "InvalidParameterValue.IdentityKeyError"
//  INVALIDPARAMETERVALUE_IDENTITYURLERROR = "InvalidParameterValue.IdentityUrlError"
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  LIMITEXCEEDED_IDENTITYFULL = "LimitExceeded.IdentityFull"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateUserOIDCConfigWithContext(ctx context.Context, request *UpdateUserOIDCConfigRequest) (response *UpdateUserOIDCConfigResponse, err error) {
    if request == nil {
        request = NewUpdateUserOIDCConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserOIDCConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserOIDCConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserSAMLConfigRequest() (request *UpdateUserSAMLConfigRequest) {
    request = &UpdateUserSAMLConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cam", APIVersion, "UpdateUserSAMLConfig")
    
    
    return
}

func NewUpdateUserSAMLConfigResponse() (response *UpdateUserSAMLConfigResponse) {
    response = &UpdateUserSAMLConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUserSAMLConfig
// This API is used to modify user SAML configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateUserSAMLConfig(request *UpdateUserSAMLConfigRequest) (response *UpdateUserSAMLConfigResponse, err error) {
    return c.UpdateUserSAMLConfigWithContext(context.Background(), request)
}

// UpdateUserSAMLConfig
// This API is used to modify user SAML configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_METADATAERROR = "InvalidParameterValue.MetadataError"
//  RESOURCENOTFOUND_IDENTITYNOTEXIST = "ResourceNotFound.IdentityNotExist"
func (c *Client) UpdateUserSAMLConfigWithContext(ctx context.Context, request *UpdateUserSAMLConfigRequest) (response *UpdateUserSAMLConfigResponse, err error) {
    if request == nil {
        request = NewUpdateUserSAMLConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserSAMLConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserSAMLConfigResponse()
    err = c.Send(request, response)
    return
}
