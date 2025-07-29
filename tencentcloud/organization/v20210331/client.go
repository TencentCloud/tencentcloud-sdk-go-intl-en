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


func NewAcceptJoinShareUnitInvitationRequest() (request *AcceptJoinShareUnitInvitationRequest) {
    request = &AcceptJoinShareUnitInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AcceptJoinShareUnitInvitation")
    
    
    return
}

func NewAcceptJoinShareUnitInvitationResponse() (response *AcceptJoinShareUnitInvitationResponse) {
    response = &AcceptJoinShareUnitInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcceptJoinShareUnitInvitation
// This API is used to accept an invitation to join a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AcceptJoinShareUnitInvitation(request *AcceptJoinShareUnitInvitationRequest) (response *AcceptJoinShareUnitInvitationResponse, err error) {
    return c.AcceptJoinShareUnitInvitationWithContext(context.Background(), request)
}

// AcceptJoinShareUnitInvitation
// This API is used to accept an invitation to join a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AcceptJoinShareUnitInvitationWithContext(ctx context.Context, request *AcceptJoinShareUnitInvitationRequest) (response *AcceptJoinShareUnitInvitationResponse, err error) {
    if request == nil {
        request = NewAcceptJoinShareUnitInvitationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AcceptJoinShareUnitInvitation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptJoinShareUnitInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcceptJoinShareUnitInvitationResponse()
    err = c.Send(request, response)
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddExternalSAMLIdPCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddExternalSAMLIdPCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddExternalSAMLIdPCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewAddOrganizationMemberEmailRequest() (request *AddOrganizationMemberEmailRequest) {
    request = &AddOrganizationMemberEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationMemberEmail")
    
    
    return
}

func NewAddOrganizationMemberEmailResponse() (response *AddOrganizationMemberEmailResponse) {
    response = &AddOrganizationMemberEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddOrganizationMemberEmail
// This API is used to add an organization member's mailbox.
//
// error code that may be returned:
//  FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"
//  LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddOrganizationMemberEmail(request *AddOrganizationMemberEmailRequest) (response *AddOrganizationMemberEmailResponse, err error) {
    return c.AddOrganizationMemberEmailWithContext(context.Background(), request)
}

// AddOrganizationMemberEmail
// This API is used to add an organization member's mailbox.
//
// error code that may be returned:
//  FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"
//  LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddOrganizationMemberEmailWithContext(ctx context.Context, request *AddOrganizationMemberEmailRequest) (response *AddOrganizationMemberEmailResponse, err error) {
    if request == nil {
        request = NewAddOrganizationMemberEmailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddOrganizationMemberEmail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganizationMemberEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationMemberEmailResponse()
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNodeWithContext(ctx context.Context, request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewAddOrganizationNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddOrganizationNode")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddPermissionPolicyToRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddPermissionPolicyToRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddPermissionPolicyToRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewAddShareUnitRequest() (request *AddShareUnitRequest) {
    request = &AddShareUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddShareUnit")
    
    
    return
}

func NewAddShareUnitResponse() (response *AddShareUnitResponse) {
    response = &AddShareUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddShareUnit
// This API is used to create a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_RESOURCEOVERLIMIT = "FailedOperation.ResourceOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddShareUnit(request *AddShareUnitRequest) (response *AddShareUnitResponse, err error) {
    return c.AddShareUnitWithContext(context.Background(), request)
}

// AddShareUnit
// This API is used to create a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_RESOURCEOVERLIMIT = "FailedOperation.ResourceOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddShareUnitWithContext(ctx context.Context, request *AddShareUnitRequest) (response *AddShareUnitResponse, err error) {
    if request == nil {
        request = NewAddShareUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddShareUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitResponse()
    err = c.Send(request, response)
    return
}

func NewAddShareUnitMembersRequest() (request *AddShareUnitMembersRequest) {
    request = &AddShareUnitMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddShareUnitMembers")
    
    
    return
}

func NewAddShareUnitMembersResponse() (response *AddShareUnitMembersResponse) {
    response = &AddShareUnitMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddShareUnitMembers
// This API is used to add a shared unit member.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SHARINGTOOTHERORGANIZATIONMEMBER = "UnsupportedOperation.SharingToOtherOrganizationMember"
func (c *Client) AddShareUnitMembers(request *AddShareUnitMembersRequest) (response *AddShareUnitMembersResponse, err error) {
    return c.AddShareUnitMembersWithContext(context.Background(), request)
}

// AddShareUnitMembers
// This API is used to add a shared unit member.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SHARINGTOOTHERORGANIZATIONMEMBER = "UnsupportedOperation.SharingToOtherOrganizationMember"
func (c *Client) AddShareUnitMembersWithContext(ctx context.Context, request *AddShareUnitMembersRequest) (response *AddShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewAddShareUnitMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddShareUnitMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnitMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitMembersResponse()
    err = c.Send(request, response)
    return
}

func NewAddShareUnitResourcesRequest() (request *AddShareUnitResourcesRequest) {
    request = &AddShareUnitResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddShareUnitResources")
    
    
    return
}

func NewAddShareUnitResourcesResponse() (response *AddShareUnitResourcesResponse) {
    response = &AddShareUnitResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddShareUnitResources
// This API is used to add resources to a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASDIFFERENTRESOURCETYPE = "FailedOperation.HasDifferentResourceType"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MEMBERUNSUPPORTEDOPERATION = "UnsupportedOperation.MemberUnsupportedOperation"
func (c *Client) AddShareUnitResources(request *AddShareUnitResourcesRequest) (response *AddShareUnitResourcesResponse, err error) {
    return c.AddShareUnitResourcesWithContext(context.Background(), request)
}

// AddShareUnitResources
// This API is used to add resources to a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASDIFFERENTRESOURCETYPE = "FailedOperation.HasDifferentResourceType"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MEMBERUNSUPPORTEDOPERATION = "UnsupportedOperation.MemberUnsupportedOperation"
func (c *Client) AddShareUnitResourcesWithContext(ctx context.Context, request *AddShareUnitResourcesRequest) (response *AddShareUnitResourcesResponse, err error) {
    if request == nil {
        request = NewAddShareUnitResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddShareUnitResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitResourcesResponse()
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
//  FAILEDOPERATION_GROUPTYPEUSERTYPENOTMATCH = "FailedOperation.GroupTypeUserTypeNotMatch"
//  FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
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
//  FAILEDOPERATION_GROUPTYPEUSERTYPENOTMATCH = "FailedOperation.GroupTypeUserTypeNotMatch"
//  FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddUserToGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "BindOrganizationMemberAuthAccount")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CancelOrganizationMemberAuthAccount")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ClearExternalSAMLIdentityProvider")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrgServiceAssign")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrgServiceAssignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationRequest() (request *CreateOrganizationRequest) {
    request = &CreateOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganization")
    
    
    return
}

func NewCreateOrganizationResponse() (response *CreateOrganizationResponse) {
    response = &CreateOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganization
// This API is used to create an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"
func (c *Client) CreateOrganization(request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
    return c.CreateOrganizationWithContext(context.Background(), request)
}

// CreateOrganization
// This API is used to create an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"
func (c *Client) CreateOrganizationWithContext(ctx context.Context, request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationIdentityRequest() (request *CreateOrganizationIdentityRequest) {
    request = &CreateOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationIdentity")
    
    
    return
}

func NewCreateOrganizationIdentityResponse() (response *CreateOrganizationIdentityResponse) {
    response = &CreateOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationIdentity
// This API is used to add an organization identity.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) CreateOrganizationIdentity(request *CreateOrganizationIdentityRequest) (response *CreateOrganizationIdentityResponse, err error) {
    return c.CreateOrganizationIdentityWithContext(context.Background(), request)
}

// CreateOrganizationIdentity
// This API is used to add an organization identity.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) CreateOrganizationIdentityWithContext(ctx context.Context, request *CreateOrganizationIdentityRequest) (response *CreateOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationIdentityResponse()
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberAuthIdentityRequest() (request *CreateOrganizationMemberAuthIdentityRequest) {
    request = &CreateOrganizationMemberAuthIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberAuthIdentity")
    
    
    return
}

func NewCreateOrganizationMemberAuthIdentityResponse() (response *CreateOrganizationMemberAuthIdentityResponse) {
    response = &CreateOrganizationMemberAuthIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMemberAuthIdentity
// This API is used to add organization member access authorization.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMemberAuthIdentity(request *CreateOrganizationMemberAuthIdentityRequest) (response *CreateOrganizationMemberAuthIdentityResponse, err error) {
    return c.CreateOrganizationMemberAuthIdentityWithContext(context.Background(), request)
}

// CreateOrganizationMemberAuthIdentity
// This API is used to add organization member access authorization.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMemberAuthIdentityWithContext(ctx context.Context, request *CreateOrganizationMemberAuthIdentityRequest) (response *CreateOrganizationMemberAuthIdentityResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberAuthIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMemberAuthIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMemberAuthIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberAuthIdentityResponse()
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
// This API is used to create an organization member access authorization policy.
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
// This API is used to create an organization member access authorization policy.
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMemberPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMemberPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMembersPolicyRequest() (request *CreateOrganizationMembersPolicyRequest) {
    request = &CreateOrganizationMembersPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMembersPolicy")
    
    
    return
}

func NewCreateOrganizationMembersPolicyResponse() (response *CreateOrganizationMembersPolicyResponse) {
    response = &CreateOrganizationMembersPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMembersPolicy
// This API is used to create an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMembersPolicy(request *CreateOrganizationMembersPolicyRequest) (response *CreateOrganizationMembersPolicyResponse, err error) {
    return c.CreateOrganizationMembersPolicyWithContext(context.Background(), request)
}

// CreateOrganizationMembersPolicy
// This API is used to create an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMembersPolicyWithContext(ctx context.Context, request *CreateOrganizationMembersPolicyRequest) (response *CreateOrganizationMembersPolicyResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMembersPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMembersPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMembersPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMembersPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateRoleAssignment")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSCIMCredentialRequest() (request *CreateSCIMCredentialRequest) {
    request = &CreateSCIMCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateSCIMCredential")
    
    
    return
}

func NewCreateSCIMCredentialResponse() (response *CreateSCIMCredentialResponse) {
    response = &CreateSCIMCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSCIMCredential
// This API is used to create a SCIM key.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SCIMCREDENTIALGENERATEERROR = "FailedOperation.ScimCredentialGenerateError"
//  LIMITEXCEEDED_SCIMCREDENTIALLIMITEXCEEDED = "LimitExceeded.ScimCredentialLimitExceeded"
func (c *Client) CreateSCIMCredential(request *CreateSCIMCredentialRequest) (response *CreateSCIMCredentialResponse, err error) {
    return c.CreateSCIMCredentialWithContext(context.Background(), request)
}

// CreateSCIMCredential
// This API is used to create a SCIM key.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SCIMCREDENTIALGENERATEERROR = "FailedOperation.ScimCredentialGenerateError"
//  LIMITEXCEEDED_SCIMCREDENTIALLIMITEXCEEDED = "LimitExceeded.ScimCredentialLimitExceeded"
func (c *Client) CreateSCIMCredentialWithContext(ctx context.Context, request *CreateSCIMCredentialRequest) (response *CreateSCIMCredentialResponse, err error) {
    if request == nil {
        request = NewCreateSCIMCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateSCIMCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSCIMCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSCIMCredentialResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateUser")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateUserSyncProvisioning")
    
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
//  FAILEDOPERATION_MANUALGROUPNOTDELETE = "FailedOperation.ManualGroupNotDelete"
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
//  FAILEDOPERATION_MANUALGROUPNOTDELETE = "FailedOperation.ManualGroupNotDelete"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrgServiceAssign")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrgServiceAssignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
    request = &DeleteOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganization")
    
    
    return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
    response = &DeleteOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganization
// This API is used to delete an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    return c.DeleteOrganizationWithContext(context.Background(), request)
}

// DeleteOrganization
// This API is used to delete an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationWithContext(ctx context.Context, request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationIdentityRequest() (request *DeleteOrganizationIdentityRequest) {
    request = &DeleteOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationIdentity")
    
    
    return
}

func NewDeleteOrganizationIdentityResponse() (response *DeleteOrganizationIdentityResponse) {
    response = &DeleteOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationIdentity
// This API is used to delete an organization identity.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationIdentity(request *DeleteOrganizationIdentityRequest) (response *DeleteOrganizationIdentityResponse, err error) {
    return c.DeleteOrganizationIdentityWithContext(context.Background(), request)
}

// DeleteOrganizationIdentity
// This API is used to delete an organization identity.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationIdentityWithContext(ctx context.Context, request *DeleteOrganizationIdentityRequest) (response *DeleteOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMemberAuthIdentityRequest() (request *DeleteOrganizationMemberAuthIdentityRequest) {
    request = &DeleteOrganizationMemberAuthIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMemberAuthIdentity")
    
    
    return
}

func NewDeleteOrganizationMemberAuthIdentityResponse() (response *DeleteOrganizationMemberAuthIdentityResponse) {
    response = &DeleteOrganizationMemberAuthIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMemberAuthIdentity
// This API is used to delete organization member access authorization.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberAuthIdentity(request *DeleteOrganizationMemberAuthIdentityRequest) (response *DeleteOrganizationMemberAuthIdentityResponse, err error) {
    return c.DeleteOrganizationMemberAuthIdentityWithContext(context.Background(), request)
}

// DeleteOrganizationMemberAuthIdentity
// This API is used to delete organization member access authorization.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberAuthIdentityWithContext(ctx context.Context, request *DeleteOrganizationMemberAuthIdentityRequest) (response *DeleteOrganizationMemberAuthIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMemberAuthIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationMemberAuthIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMemberAuthIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMemberAuthIdentityResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMembersPolicyRequest() (request *DeleteOrganizationMembersPolicyRequest) {
    request = &DeleteOrganizationMembersPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembersPolicy")
    
    
    return
}

func NewDeleteOrganizationMembersPolicyResponse() (response *DeleteOrganizationMembersPolicyResponse) {
    response = &DeleteOrganizationMembersPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMembersPolicy
// This API is used to delete an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMembersPolicy(request *DeleteOrganizationMembersPolicyRequest) (response *DeleteOrganizationMembersPolicyResponse, err error) {
    return c.DeleteOrganizationMembersPolicyWithContext(context.Background(), request)
}

// DeleteOrganizationMembersPolicy
// This API is used to delete an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMembersPolicyWithContext(ctx context.Context, request *DeleteOrganizationMembersPolicyRequest) (response *DeleteOrganizationMembersPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMembersPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationMembersPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMembersPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMembersPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationNodes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteRoleAssignment")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSCIMCredentialRequest() (request *DeleteSCIMCredentialRequest) {
    request = &DeleteSCIMCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteSCIMCredential")
    
    
    return
}

func NewDeleteSCIMCredentialResponse() (response *DeleteSCIMCredentialResponse) {
    response = &DeleteSCIMCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSCIMCredential
// This API is used to delete a SCIM key.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
func (c *Client) DeleteSCIMCredential(request *DeleteSCIMCredentialRequest) (response *DeleteSCIMCredentialResponse, err error) {
    return c.DeleteSCIMCredentialWithContext(context.Background(), request)
}

// DeleteSCIMCredential
// This API is used to delete a SCIM key.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
func (c *Client) DeleteSCIMCredentialWithContext(ctx context.Context, request *DeleteSCIMCredentialRequest) (response *DeleteSCIMCredentialResponse, err error) {
    if request == nil {
        request = NewDeleteSCIMCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteSCIMCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSCIMCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSCIMCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareUnitRequest() (request *DeleteShareUnitRequest) {
    request = &DeleteShareUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteShareUnit")
    
    
    return
}

func NewDeleteShareUnitResponse() (response *DeleteShareUnitResponse) {
    response = &DeleteShareUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShareUnit
// This API is used to delete a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnit(request *DeleteShareUnitRequest) (response *DeleteShareUnitResponse, err error) {
    return c.DeleteShareUnitWithContext(context.Background(), request)
}

// DeleteShareUnit
// This API is used to delete a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitWithContext(ctx context.Context, request *DeleteShareUnitRequest) (response *DeleteShareUnitResponse, err error) {
    if request == nil {
        request = NewDeleteShareUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteShareUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareUnitMembersRequest() (request *DeleteShareUnitMembersRequest) {
    request = &DeleteShareUnitMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteShareUnitMembers")
    
    
    return
}

func NewDeleteShareUnitMembersResponse() (response *DeleteShareUnitMembersResponse) {
    response = &DeleteShareUnitMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShareUnitMembers
// This API is used to delete a shared unit member.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitMembers(request *DeleteShareUnitMembersRequest) (response *DeleteShareUnitMembersResponse, err error) {
    return c.DeleteShareUnitMembersWithContext(context.Background(), request)
}

// DeleteShareUnitMembers
// This API is used to delete a shared unit member.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitMembersWithContext(ctx context.Context, request *DeleteShareUnitMembersRequest) (response *DeleteShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewDeleteShareUnitMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteShareUnitMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnitMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareUnitResourcesRequest() (request *DeleteShareUnitResourcesRequest) {
    request = &DeleteShareUnitResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteShareUnitResources")
    
    
    return
}

func NewDeleteShareUnitResourcesResponse() (response *DeleteShareUnitResourcesResponse) {
    response = &DeleteShareUnitResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShareUnitResources
// This API is used to delete shared unit resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitResources(request *DeleteShareUnitResourcesRequest) (response *DeleteShareUnitResourcesResponse, err error) {
    return c.DeleteShareUnitResourcesWithContext(context.Background(), request)
}

// DeleteShareUnitResources
// This API is used to delete shared unit resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitResourcesWithContext(ctx context.Context, request *DeleteShareUnitResourcesRequest) (response *DeleteShareUnitResourcesResponse, err error) {
    if request == nil {
        request = NewDeleteShareUnitResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteShareUnitResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitResourcesResponse()
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
//  FAILEDOPERATION_MANUALUSERNOTDELETE = "FailedOperation.ManualUserNotDelete"
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
//  FAILEDOPERATION_MANUALUSERNOTDELETE = "FailedOperation.ManualUserNotDelete"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteUser")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteUserSyncProvisioning")
    
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
// This API is used to obtain TCO Identity Center service information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
func (c *Client) DescribeIdentityCenter(request *DescribeIdentityCenterRequest) (response *DescribeIdentityCenterResponse, err error) {
    return c.DescribeIdentityCenterWithContext(context.Background(), request)
}

// DescribeIdentityCenter
// This API is used to obtain TCO Identity Center service information.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
func (c *Client) DescribeIdentityCenterWithContext(ctx context.Context, request *DescribeIdentityCenterRequest) (response *DescribeIdentityCenterResponse, err error) {
    if request == nil {
        request = NewDescribeIdentityCenterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeIdentityCenter")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganization")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberAuthIdentities")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthIdentities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberEmailBindRequest() (request *DescribeOrganizationMemberEmailBindRequest) {
    request = &DescribeOrganizationMemberEmailBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberEmailBind")
    
    
    return
}

func NewDescribeOrganizationMemberEmailBindResponse() (response *DescribeOrganizationMemberEmailBindResponse) {
    response = &DescribeOrganizationMemberEmailBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberEmailBind
// This API is used to query detailed information about member mailbox binding.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"
//  FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"
//  FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"
//  FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"
//  INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"
//  INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"
//  INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeOrganizationMemberEmailBind(request *DescribeOrganizationMemberEmailBindRequest) (response *DescribeOrganizationMemberEmailBindResponse, err error) {
    return c.DescribeOrganizationMemberEmailBindWithContext(context.Background(), request)
}

// DescribeOrganizationMemberEmailBind
// This API is used to query detailed information about member mailbox binding.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"
//  FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"
//  FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"
//  FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"
//  INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"
//  INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"
//  INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeOrganizationMemberEmailBindWithContext(ctx context.Context, request *DescribeOrganizationMemberEmailBindRequest) (response *DescribeOrganizationMemberEmailBindResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberEmailBindRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberEmailBind")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberEmailBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberEmailBindResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberPolicies")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareAreasRequest() (request *DescribeShareAreasRequest) {
    request = &DescribeShareAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareAreas")
    
    
    return
}

func NewDescribeShareAreasResponse() (response *DescribeShareAreasResponse) {
    response = &DescribeShareAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareAreas
// This API is used to obtain a list of shareable regions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareAreas(request *DescribeShareAreasRequest) (response *DescribeShareAreasResponse, err error) {
    return c.DescribeShareAreasWithContext(context.Background(), request)
}

// DescribeShareAreas
// This API is used to obtain a list of shareable regions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareAreasWithContext(ctx context.Context, request *DescribeShareAreasRequest) (response *DescribeShareAreasResponse, err error) {
    if request == nil {
        request = NewDescribeShareAreasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareAreas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareUnitMembersRequest() (request *DescribeShareUnitMembersRequest) {
    request = &DescribeShareUnitMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareUnitMembers")
    
    
    return
}

func NewDescribeShareUnitMembersResponse() (response *DescribeShareUnitMembersResponse) {
    response = &DescribeShareUnitMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareUnitMembers
// This API is used to obtain the member list of a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitMembers(request *DescribeShareUnitMembersRequest) (response *DescribeShareUnitMembersResponse, err error) {
    return c.DescribeShareUnitMembersWithContext(context.Background(), request)
}

// DescribeShareUnitMembers
// This API is used to obtain the member list of a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitMembersWithContext(ctx context.Context, request *DescribeShareUnitMembersRequest) (response *DescribeShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewDescribeShareUnitMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareUnitMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareUnitMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareUnitMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareUnitResourcesRequest() (request *DescribeShareUnitResourcesRequest) {
    request = &DescribeShareUnitResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareUnitResources")
    
    
    return
}

func NewDescribeShareUnitResourcesResponse() (response *DescribeShareUnitResourcesResponse) {
    response = &DescribeShareUnitResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareUnitResources
// This API is used to obtain the resource list of a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitResources(request *DescribeShareUnitResourcesRequest) (response *DescribeShareUnitResourcesResponse, err error) {
    return c.DescribeShareUnitResourcesWithContext(context.Background(), request)
}

// DescribeShareUnitResources
// This API is used to obtain the resource list of a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitResourcesWithContext(ctx context.Context, request *DescribeShareUnitResourcesRequest) (response *DescribeShareUnitResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeShareUnitResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareUnitResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareUnitResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareUnitsRequest() (request *DescribeShareUnitsRequest) {
    request = &DescribeShareUnitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareUnits")
    
    
    return
}

func NewDescribeShareUnitsResponse() (response *DescribeShareUnitsResponse) {
    response = &DescribeShareUnitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareUnits
// This API is used to obtain a list of shared units.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnits(request *DescribeShareUnitsRequest) (response *DescribeShareUnitsResponse, err error) {
    return c.DescribeShareUnitsWithContext(context.Background(), request)
}

// DescribeShareUnits
// This API is used to obtain a list of shared units.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitsWithContext(ctx context.Context, request *DescribeShareUnitsRequest) (response *DescribeShareUnitsResponse, err error) {
    if request == nil {
        request = NewDescribeShareUnitsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareUnits")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareUnits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareUnitsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DismantleRoleConfiguration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetExternalSAMLIdentityProvider")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetProvisioningTaskStatus")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewGetSCIMSynchronizationStatusRequest() (request *GetSCIMSynchronizationStatusRequest) {
    request = &GetSCIMSynchronizationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetSCIMSynchronizationStatus")
    
    
    return
}

func NewGetSCIMSynchronizationStatusResponse() (response *GetSCIMSynchronizationStatusResponse) {
    response = &GetSCIMSynchronizationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSCIMSynchronizationStatus
// This API is used to query SCIM synchronization status.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetSCIMSynchronizationStatus(request *GetSCIMSynchronizationStatusRequest) (response *GetSCIMSynchronizationStatusResponse, err error) {
    return c.GetSCIMSynchronizationStatusWithContext(context.Background(), request)
}

// GetSCIMSynchronizationStatus
// This API is used to query SCIM synchronization status.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetSCIMSynchronizationStatusWithContext(ctx context.Context, request *GetSCIMSynchronizationStatusRequest) (response *GetSCIMSynchronizationStatusResponse, err error) {
    if request == nil {
        request = NewGetSCIMSynchronizationStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetSCIMSynchronizationStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSCIMSynchronizationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSCIMSynchronizationStatusResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetTaskStatus")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetUser")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetUserSyncProvisioning")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetZoneSAMLServiceProviderInfo")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetZoneStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetZoneStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetZoneStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewInviteOrganizationMemberRequest() (request *InviteOrganizationMemberRequest) {
    request = &InviteOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "InviteOrganizationMember")
    
    
    return
}

func NewInviteOrganizationMemberResponse() (response *InviteOrganizationMemberResponse) {
    response = &InviteOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InviteOrganizationMember
// This API is used to invite a member.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLYEXIST = "FailedOperation.ApplyExist"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_EXISTOTHERORGANIZATIONMEMBERSHARED = "FailedOperation.ExistOtherOrganizationMemberShared"
//  FAILEDOPERATION_GETACCOUNTREGION = "FailedOperation.GetAccountRegion"
//  FAILEDOPERATION_IMPORTFILEILLEGAL = "FailedOperation.ImportFileIllegal"
//  FAILEDOPERATION_INVITATIONEXIST = "FailedOperation.InvitationExist"
//  FAILEDOPERATION_MEMBEREXISTINOTHERORGANIZATION = "FailedOperation.MemberExistInOtherOrganization"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHRELATIONEXIST = "FailedOperation.OrganizationAuthRelationExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBEREXIST = "FailedOperation.OrganizationMemberExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_INVITATIONOVERLIMIT = "LimitExceeded.InvitationOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFMEMBER = "UnsupportedOperation.AbnormalFinancialStatusOfMember"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_AGENTNOTSAME = "UnsupportedOperation.AgentNotSame"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERHASVOUCHER = "UnsupportedOperation.MemberHasVoucher"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERISNOTCLIENT = "UnsupportedOperation.MemberIsNotClient"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) InviteOrganizationMember(request *InviteOrganizationMemberRequest) (response *InviteOrganizationMemberResponse, err error) {
    return c.InviteOrganizationMemberWithContext(context.Background(), request)
}

// InviteOrganizationMember
// This API is used to invite a member.
//
// error code that may be returned:
//  FAILEDOPERATION_APPLYEXIST = "FailedOperation.ApplyExist"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_EXISTOTHERORGANIZATIONMEMBERSHARED = "FailedOperation.ExistOtherOrganizationMemberShared"
//  FAILEDOPERATION_GETACCOUNTREGION = "FailedOperation.GetAccountRegion"
//  FAILEDOPERATION_IMPORTFILEILLEGAL = "FailedOperation.ImportFileIllegal"
//  FAILEDOPERATION_INVITATIONEXIST = "FailedOperation.InvitationExist"
//  FAILEDOPERATION_MEMBEREXISTINOTHERORGANIZATION = "FailedOperation.MemberExistInOtherOrganization"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHRELATIONEXIST = "FailedOperation.OrganizationAuthRelationExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBEREXIST = "FailedOperation.OrganizationMemberExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_INVITATIONOVERLIMIT = "LimitExceeded.InvitationOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFMEMBER = "UnsupportedOperation.AbnormalFinancialStatusOfMember"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_AGENTNOTSAME = "UnsupportedOperation.AgentNotSame"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERHASVOUCHER = "UnsupportedOperation.MemberHasVoucher"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERISNOTCLIENT = "UnsupportedOperation.MemberIsNotClient"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) InviteOrganizationMemberWithContext(ctx context.Context, request *InviteOrganizationMemberRequest) (response *InviteOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewInviteOrganizationMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "InviteOrganizationMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InviteOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewInviteOrganizationMemberResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListExternalSAMLIdPCertificates")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListGroupMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListGroups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListJoinedGroupsForUser")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListOrgServiceAssignMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListOrganizationIdentity")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListOrganizationService")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListPermissionPoliciesInRoleConfiguration")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListRoleAssignments")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListRoleConfigurationProvisionings")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListRoleConfigurations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleConfigurations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewListSCIMCredentialsRequest() (request *ListSCIMCredentialsRequest) {
    request = &ListSCIMCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListSCIMCredentials")
    
    
    return
}

func NewListSCIMCredentialsResponse() (response *ListSCIMCredentialsResponse) {
    response = &ListSCIMCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSCIMCredentials
// This API is used to query the user SCIM key list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) ListSCIMCredentials(request *ListSCIMCredentialsRequest) (response *ListSCIMCredentialsResponse, err error) {
    return c.ListSCIMCredentialsWithContext(context.Background(), request)
}

// ListSCIMCredentials
// This API is used to query the user SCIM key list.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) ListSCIMCredentialsWithContext(ctx context.Context, request *ListSCIMCredentialsRequest) (response *ListSCIMCredentialsResponse, err error) {
    if request == nil {
        request = NewListSCIMCredentialsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListSCIMCredentials")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSCIMCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSCIMCredentialsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListTasks")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListUserSyncProvisionings")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListUsers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "MoveOrganizationNodeMembers")
    
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
// This API is used to activate Identity Center service (CIC).
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
// This API is used to activate Identity Center service (CIC).
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "OpenIdentityCenter")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ProvisionRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProvisionRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewProvisionRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewQuitOrganizationRequest() (request *QuitOrganizationRequest) {
    request = &QuitOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "QuitOrganization")
    
    
    return
}

func NewQuitOrganizationResponse() (response *QuitOrganizationResponse) {
    response = &QuitOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QuitOrganization
// This API is used to exit an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
func (c *Client) QuitOrganization(request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
    return c.QuitOrganizationWithContext(context.Background(), request)
}

// QuitOrganization
// This API is used to exit an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
func (c *Client) QuitOrganizationWithContext(ctx context.Context, request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
    if request == nil {
        request = NewQuitOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "QuitOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuitOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewQuitOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewRejectJoinShareUnitInvitationRequest() (request *RejectJoinShareUnitInvitationRequest) {
    request = &RejectJoinShareUnitInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RejectJoinShareUnitInvitation")
    
    
    return
}

func NewRejectJoinShareUnitInvitationResponse() (response *RejectJoinShareUnitInvitationResponse) {
    response = &RejectJoinShareUnitInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RejectJoinShareUnitInvitation
// This API is used to reject an invitation to join a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectJoinShareUnitInvitation(request *RejectJoinShareUnitInvitationRequest) (response *RejectJoinShareUnitInvitationResponse, err error) {
    return c.RejectJoinShareUnitInvitationWithContext(context.Background(), request)
}

// RejectJoinShareUnitInvitation
// This API is used to reject an invitation to join a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectJoinShareUnitInvitationWithContext(ctx context.Context, request *RejectJoinShareUnitInvitationRequest) (response *RejectJoinShareUnitInvitationResponse, err error) {
    if request == nil {
        request = NewRejectJoinShareUnitInvitationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RejectJoinShareUnitInvitation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RejectJoinShareUnitInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewRejectJoinShareUnitInvitationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RemoveExternalSAMLIdPCertificate")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RemovePermissionPolicyFromRoleConfiguration")
    
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
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
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
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"
//  LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"
func (c *Client) RemoveUserFromGroupWithContext(ctx context.Context, request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RemoveUserFromGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserFromGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserFromGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSendOrgMemberAccountBindEmailRequest() (request *SendOrgMemberAccountBindEmailRequest) {
    request = &SendOrgMemberAccountBindEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "SendOrgMemberAccountBindEmail")
    
    
    return
}

func NewSendOrgMemberAccountBindEmailResponse() (response *SendOrgMemberAccountBindEmailResponse) {
    response = &SendOrgMemberAccountBindEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendOrgMemberAccountBindEmail
// This API is used to resend an email for activating the member's bound mailbox.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SENDEMAILLIMIT = "LimitExceeded.SendEmailLimit"
//  LIMITEXCEEDED_SENDEMAILWITHINONEHOURLIMIT = "LimitExceeded.SendEmailWithinOneHourLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) SendOrgMemberAccountBindEmail(request *SendOrgMemberAccountBindEmailRequest) (response *SendOrgMemberAccountBindEmailResponse, err error) {
    return c.SendOrgMemberAccountBindEmailWithContext(context.Background(), request)
}

// SendOrgMemberAccountBindEmail
// This API is used to resend an email for activating the member's bound mailbox.
//
// error code that may be returned:
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SENDEMAILLIMIT = "LimitExceeded.SendEmailLimit"
//  LIMITEXCEEDED_SENDEMAILWITHINONEHOURLIMIT = "LimitExceeded.SendEmailWithinOneHourLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) SendOrgMemberAccountBindEmailWithContext(ctx context.Context, request *SendOrgMemberAccountBindEmailRequest) (response *SendOrgMemberAccountBindEmailResponse, err error) {
    if request == nil {
        request = NewSendOrgMemberAccountBindEmailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "SendOrgMemberAccountBindEmail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendOrgMemberAccountBindEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendOrgMemberAccountBindEmailResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "SetExternalSAMLIdentityProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCustomPolicyForRoleConfigurationRequest() (request *UpdateCustomPolicyForRoleConfigurationRequest) {
    request = &UpdateCustomPolicyForRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateCustomPolicyForRoleConfiguration")
    
    
    return
}

func NewUpdateCustomPolicyForRoleConfigurationResponse() (response *UpdateCustomPolicyForRoleConfigurationResponse) {
    response = &UpdateCustomPolicyForRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCustomPolicyForRoleConfiguration
// This API is used to modify a custom policy for permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYTYPEERROR = "InvalidParameter.PolicyTypeError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) UpdateCustomPolicyForRoleConfiguration(request *UpdateCustomPolicyForRoleConfigurationRequest) (response *UpdateCustomPolicyForRoleConfigurationResponse, err error) {
    return c.UpdateCustomPolicyForRoleConfigurationWithContext(context.Background(), request)
}

// UpdateCustomPolicyForRoleConfiguration
// This API is used to modify a custom policy for permission configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYTYPEERROR = "InvalidParameter.PolicyTypeError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) UpdateCustomPolicyForRoleConfigurationWithContext(ctx context.Context, request *UpdateCustomPolicyForRoleConfigurationRequest) (response *UpdateCustomPolicyForRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateCustomPolicyForRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateCustomPolicyForRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCustomPolicyForRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCustomPolicyForRoleConfigurationResponse()
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
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
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
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    if request == nil {
        request = NewUpdateGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationIdentityRequest() (request *UpdateOrganizationIdentityRequest) {
    request = &UpdateOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationIdentity")
    
    
    return
}

func NewUpdateOrganizationIdentityResponse() (response *UpdateOrganizationIdentityResponse) {
    response = &UpdateOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationIdentity
// This API is used to update an organization identity.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) UpdateOrganizationIdentity(request *UpdateOrganizationIdentityRequest) (response *UpdateOrganizationIdentityResponse, err error) {
    return c.UpdateOrganizationIdentityWithContext(context.Background(), request)
}

// UpdateOrganizationIdentity
// This API is used to update an organization identity.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) UpdateOrganizationIdentityWithContext(ctx context.Context, request *UpdateOrganizationIdentityRequest) (response *UpdateOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationMemberRequest() (request *UpdateOrganizationMemberRequest) {
    request = &UpdateOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMember")
    
    
    return
}

func NewUpdateOrganizationMemberResponse() (response *UpdateOrganizationMemberResponse) {
    response = &UpdateOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationMember
// This API is used to update organization member information.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CHANGEPERMISSIONRECORDEXIST = "FailedOperation.ChangePermissionRecordExist"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INVALIDPARAMETER_ALLOWQUITILLEGAL = "InvalidParameter.AllowQuitIllegal"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_DELETEDELEGATEPAYERNOTALLOW = "UnsupportedOperation.DeleteDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
func (c *Client) UpdateOrganizationMember(request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
    return c.UpdateOrganizationMemberWithContext(context.Background(), request)
}

// UpdateOrganizationMember
// This API is used to update organization member information.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CHANGEPERMISSIONRECORDEXIST = "FailedOperation.ChangePermissionRecordExist"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INVALIDPARAMETER_ALLOWQUITILLEGAL = "InvalidParameter.AllowQuitIllegal"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_DELETEDELEGATEPAYERNOTALLOW = "UnsupportedOperation.DeleteDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
func (c *Client) UpdateOrganizationMemberWithContext(ctx context.Context, request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationMemberEmailBindRequest() (request *UpdateOrganizationMemberEmailBindRequest) {
    request = &UpdateOrganizationMemberEmailBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMemberEmailBind")
    
    
    return
}

func NewUpdateOrganizationMemberEmailBindResponse() (response *UpdateOrganizationMemberEmailBindResponse) {
    response = &UpdateOrganizationMemberEmailBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationMemberEmailBind
// This API is used to modify the mailbox of a bound member.
//
// error code that may be returned:
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMemberEmailBind(request *UpdateOrganizationMemberEmailBindRequest) (response *UpdateOrganizationMemberEmailBindResponse, err error) {
    return c.UpdateOrganizationMemberEmailBindWithContext(context.Background(), request)
}

// UpdateOrganizationMemberEmailBind
// This API is used to modify the mailbox of a bound member.
//
// error code that may be returned:
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMemberEmailBindWithContext(ctx context.Context, request *UpdateOrganizationMemberEmailBindRequest) (response *UpdateOrganizationMemberEmailBindResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationMemberEmailBindRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationMemberEmailBind")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMemberEmailBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMemberEmailBindResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationNode")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSCIMCredentialStatusRequest() (request *UpdateSCIMCredentialStatusRequest) {
    request = &UpdateSCIMCredentialStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateSCIMCredentialStatus")
    
    
    return
}

func NewUpdateSCIMCredentialStatusResponse() (response *UpdateSCIMCredentialStatusResponse) {
    response = &UpdateSCIMCredentialStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSCIMCredentialStatus
// This API is used to enable or disable a SCIM key.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
//  INVALIDPARAMETER_USERSCIMCREDENTIALSTATUSERROR = "InvalidParameter.UserScimCredentialStatusError"
func (c *Client) UpdateSCIMCredentialStatus(request *UpdateSCIMCredentialStatusRequest) (response *UpdateSCIMCredentialStatusResponse, err error) {
    return c.UpdateSCIMCredentialStatusWithContext(context.Background(), request)
}

// UpdateSCIMCredentialStatus
// This API is used to enable or disable a SCIM key.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
//  INVALIDPARAMETER_USERSCIMCREDENTIALSTATUSERROR = "InvalidParameter.UserScimCredentialStatusError"
func (c *Client) UpdateSCIMCredentialStatusWithContext(ctx context.Context, request *UpdateSCIMCredentialStatusRequest) (response *UpdateSCIMCredentialStatusResponse, err error) {
    if request == nil {
        request = NewUpdateSCIMCredentialStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateSCIMCredentialStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSCIMCredentialStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSCIMCredentialStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSCIMSynchronizationStatusRequest() (request *UpdateSCIMSynchronizationStatusRequest) {
    request = &UpdateSCIMSynchronizationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateSCIMSynchronizationStatus")
    
    
    return
}

func NewUpdateSCIMSynchronizationStatusResponse() (response *UpdateSCIMSynchronizationStatusResponse) {
    response = &UpdateSCIMSynchronizationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSCIMSynchronizationStatus
// This API is used to enable or disable user SCIM synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMSYNCSTATUSERROR = "InvalidParameter.ScimSyncStatusError"
func (c *Client) UpdateSCIMSynchronizationStatus(request *UpdateSCIMSynchronizationStatusRequest) (response *UpdateSCIMSynchronizationStatusResponse, err error) {
    return c.UpdateSCIMSynchronizationStatusWithContext(context.Background(), request)
}

// UpdateSCIMSynchronizationStatus
// This API is used to enable or disable user SCIM synchronization.
//
// error code that may be returned:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMSYNCSTATUSERROR = "InvalidParameter.ScimSyncStatusError"
func (c *Client) UpdateSCIMSynchronizationStatusWithContext(ctx context.Context, request *UpdateSCIMSynchronizationStatusRequest) (response *UpdateSCIMSynchronizationStatusResponse, err error) {
    if request == nil {
        request = NewUpdateSCIMSynchronizationStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateSCIMSynchronizationStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSCIMSynchronizationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSCIMSynchronizationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateShareUnitRequest() (request *UpdateShareUnitRequest) {
    request = &UpdateShareUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateShareUnit")
    
    
    return
}

func NewUpdateShareUnitResponse() (response *UpdateShareUnitResponse) {
    response = &UpdateShareUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateShareUnit
// This API is used to update a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTSHAREMEMBERNOTINORGANIZATION = "FailedOperation.ExistShareMemberNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateShareUnit(request *UpdateShareUnitRequest) (response *UpdateShareUnitResponse, err error) {
    return c.UpdateShareUnitWithContext(context.Background(), request)
}

// UpdateShareUnit
// This API is used to update a shared unit.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTSHAREMEMBERNOTINORGANIZATION = "FailedOperation.ExistShareMemberNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateShareUnitWithContext(ctx context.Context, request *UpdateShareUnitRequest) (response *UpdateShareUnitResponse, err error) {
    if request == nil {
        request = NewUpdateShareUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateShareUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateShareUnitResponse()
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
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
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
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateUser")
    
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
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
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
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserStatusWithContext(ctx context.Context, request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    if request == nil {
        request = NewUpdateUserStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateUserStatus")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateUserSyncProvisioning")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateZoneResponse()
    err = c.Send(request, response)
    return
}
