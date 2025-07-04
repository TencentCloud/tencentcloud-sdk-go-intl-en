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

package v20181225

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-12-25"

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


func NewAcceptOrganizationInvitationRequest() (request *AcceptOrganizationInvitationRequest) {
    request = &AcceptOrganizationInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AcceptOrganizationInvitation")
    
    
    return
}

func NewAcceptOrganizationInvitationResponse() (response *AcceptOrganizationInvitationResponse) {
    response = &AcceptOrganizationInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcceptOrganizationInvitation
// This API is used to accept an invitation to an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_INORGANIZATIONALREADY = "FailedOperation.InOrganizationAlready"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
func (c *Client) AcceptOrganizationInvitation(request *AcceptOrganizationInvitationRequest) (response *AcceptOrganizationInvitationResponse, err error) {
    return c.AcceptOrganizationInvitationWithContext(context.Background(), request)
}

// AcceptOrganizationInvitation
// This API is used to accept an invitation to an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_INORGANIZATIONALREADY = "FailedOperation.InOrganizationAlready"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
func (c *Client) AcceptOrganizationInvitationWithContext(ctx context.Context, request *AcceptOrganizationInvitationRequest) (response *AcceptOrganizationInvitationResponse, err error) {
    if request == nil {
        request = NewAcceptOrganizationInvitationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptOrganizationInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcceptOrganizationInvitationResponse()
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
// This API is used to add an organizational unit.
//
// error code that may be returned:
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCEINUSE_NODENAME = "ResourceInUse.NodeName"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNode(request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    return c.AddOrganizationNodeWithContext(context.Background(), request)
}

// AddOrganizationNode
// This API is used to add an organizational unit.
//
// error code that may be returned:
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCEINUSE_NODENAME = "ResourceInUse.NodeName"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
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

func NewCancelOrganizationInvitationRequest() (request *CancelOrganizationInvitationRequest) {
    request = &CancelOrganizationInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationInvitation")
    
    
    return
}

func NewCancelOrganizationInvitationResponse() (response *CancelOrganizationInvitationResponse) {
    response = &CancelOrganizationInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelOrganizationInvitation
// This API is used to cancel an invitation to an organization.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
func (c *Client) CancelOrganizationInvitation(request *CancelOrganizationInvitationRequest) (response *CancelOrganizationInvitationResponse, err error) {
    return c.CancelOrganizationInvitationWithContext(context.Background(), request)
}

// CancelOrganizationInvitation
// This API is used to cancel an invitation to an organization.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
func (c *Client) CancelOrganizationInvitationWithContext(ctx context.Context, request *CancelOrganizationInvitationRequest) (response *CancelOrganizationInvitationResponse, err error) {
    if request == nil {
        request = NewCancelOrganizationInvitationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelOrganizationInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelOrganizationInvitationResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateOrganization(request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
    return c.CreateOrganizationWithContext(context.Background(), request)
}

// CreateOrganization
// This API is used to create an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateOrganizationWithContext(ctx context.Context, request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationResponse()
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
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    return c.DeleteOrganizationWithContext(context.Background(), request)
}

// DeleteOrganization
// This API is used to delete an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationWithContext(ctx context.Context, request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMemberFromNodeRequest() (request *DeleteOrganizationMemberFromNodeRequest) {
    request = &DeleteOrganizationMemberFromNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMemberFromNode")
    
    
    return
}

func NewDeleteOrganizationMemberFromNodeResponse() (response *DeleteOrganizationMemberFromNodeResponse) {
    response = &DeleteOrganizationMemberFromNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMemberFromNode
// This API is used to delete an organization member.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEDELETEMEMBERFROMROOTNODE = "FailedOperation.DisableDeleteMemberFromRootNode"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberFromNode(request *DeleteOrganizationMemberFromNodeRequest) (response *DeleteOrganizationMemberFromNodeResponse, err error) {
    return c.DeleteOrganizationMemberFromNodeWithContext(context.Background(), request)
}

// DeleteOrganizationMemberFromNode
// This API is used to delete an organization member.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEDELETEMEMBERFROMROOTNODE = "FailedOperation.DisableDeleteMemberFromRootNode"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberFromNodeWithContext(ctx context.Context, request *DeleteOrganizationMemberFromNodeRequest) (response *DeleteOrganizationMemberFromNodeResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMemberFromNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMemberFromNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMemberFromNodeResponse()
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
// This API is used to delete multiple organization members in a single request.
//
// error code that may be returned:
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    return c.DeleteOrganizationMembersWithContext(context.Background(), request)
}

// DeleteOrganizationMembers
// This API is used to delete multiple organization members in a single request.
//
// error code that may be returned:
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
// This API is used to delete multiple organizational units in a single request.
//
// error code that may be returned:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodes(request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    return c.DeleteOrganizationNodesWithContext(context.Background(), request)
}

// DeleteOrganizationNodes
// This API is used to delete multiple organizational units in a single request.
//
// error code that may be returned:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
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

func NewDenyOrganizationInvitationRequest() (request *DenyOrganizationInvitationRequest) {
    request = &DenyOrganizationInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DenyOrganizationInvitation")
    
    
    return
}

func NewDenyOrganizationInvitationResponse() (response *DenyOrganizationInvitationResponse) {
    response = &DenyOrganizationInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DenyOrganizationInvitation
// This API is used to decline an invitation to an organization.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
func (c *Client) DenyOrganizationInvitation(request *DenyOrganizationInvitationRequest) (response *DenyOrganizationInvitationResponse, err error) {
    return c.DenyOrganizationInvitationWithContext(context.Background(), request)
}

// DenyOrganizationInvitation
// This API is used to decline an invitation to an organization.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
func (c *Client) DenyOrganizationInvitationWithContext(ctx context.Context, request *DenyOrganizationInvitationRequest) (response *DenyOrganizationInvitationResponse, err error) {
    if request == nil {
        request = NewDenyOrganizationInvitationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DenyOrganizationInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDenyOrganizationInvitationResponse()
    err = c.Send(request, response)
    return
}

func NewGetOrganizationRequest() (request *GetOrganizationRequest) {
    request = &GetOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetOrganization")
    
    
    return
}

func NewGetOrganizationResponse() (response *GetOrganizationResponse) {
    response = &GetOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOrganization
// This API is used to obtain information on organizations.
//
// error code that may be returned:
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetOrganization(request *GetOrganizationRequest) (response *GetOrganizationResponse, err error) {
    return c.GetOrganizationWithContext(context.Background(), request)
}

// GetOrganization
// This API is used to obtain information on organizations.
//
// error code that may be returned:
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetOrganizationWithContext(ctx context.Context, request *GetOrganizationRequest) (response *GetOrganizationResponse, err error) {
    if request == nil {
        request = NewGetOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewGetOrganizationMemberRequest() (request *GetOrganizationMemberRequest) {
    request = &GetOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetOrganizationMember")
    
    
    return
}

func NewGetOrganizationMemberResponse() (response *GetOrganizationMemberResponse) {
    response = &GetOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOrganizationMember
// This API is used to obtain information on organization members.
//
// error code that may be returned:
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) GetOrganizationMember(request *GetOrganizationMemberRequest) (response *GetOrganizationMemberResponse, err error) {
    return c.GetOrganizationMemberWithContext(context.Background(), request)
}

// GetOrganizationMember
// This API is used to obtain information on organization members.
//
// error code that may be returned:
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) GetOrganizationMemberWithContext(ctx context.Context, request *GetOrganizationMemberRequest) (response *GetOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewGetOrganizationMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationInvitationsRequest() (request *ListOrganizationInvitationsRequest) {
    request = &ListOrganizationInvitationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationInvitations")
    
    
    return
}

func NewListOrganizationInvitationsResponse() (response *ListOrganizationInvitationsResponse) {
    response = &ListOrganizationInvitationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationInvitations
// This API is used to obtain an invitation list.
//
// error code that may be returned:
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListOrganizationInvitations(request *ListOrganizationInvitationsRequest) (response *ListOrganizationInvitationsResponse, err error) {
    return c.ListOrganizationInvitationsWithContext(context.Background(), request)
}

// ListOrganizationInvitations
// This API is used to obtain an invitation list.
//
// error code that may be returned:
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListOrganizationInvitationsWithContext(ctx context.Context, request *ListOrganizationInvitationsRequest) (response *ListOrganizationInvitationsResponse, err error) {
    if request == nil {
        request = NewListOrganizationInvitationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationInvitations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationInvitationsResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationMembersRequest() (request *ListOrganizationMembersRequest) {
    request = &ListOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationMembers")
    
    
    return
}

func NewListOrganizationMembersResponse() (response *ListOrganizationMembersResponse) {
    response = &ListOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationMembers
// This API is used to obtain a list of organization members.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationMembers(request *ListOrganizationMembersRequest) (response *ListOrganizationMembersResponse, err error) {
    return c.ListOrganizationMembersWithContext(context.Background(), request)
}

// ListOrganizationMembers
// This API is used to obtain a list of organization members.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationMembersWithContext(ctx context.Context, request *ListOrganizationMembersRequest) (response *ListOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewListOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationNodeMembersRequest() (request *ListOrganizationNodeMembersRequest) {
    request = &ListOrganizationNodeMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationNodeMembers")
    
    
    return
}

func NewListOrganizationNodeMembersResponse() (response *ListOrganizationNodeMembersResponse) {
    response = &ListOrganizationNodeMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationNodeMembers
// This API is used to obtain a list of organizational unit members.
//
// error code that may be returned:
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationNodeMembers(request *ListOrganizationNodeMembersRequest) (response *ListOrganizationNodeMembersResponse, err error) {
    return c.ListOrganizationNodeMembersWithContext(context.Background(), request)
}

// ListOrganizationNodeMembers
// This API is used to obtain a list of organizational unit members.
//
// error code that may be returned:
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationNodeMembersWithContext(ctx context.Context, request *ListOrganizationNodeMembersRequest) (response *ListOrganizationNodeMembersResponse, err error) {
    if request == nil {
        request = NewListOrganizationNodeMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationNodeMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationNodeMembersResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationNodesRequest() (request *ListOrganizationNodesRequest) {
    request = &ListOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationNodes")
    
    
    return
}

func NewListOrganizationNodesResponse() (response *ListOrganizationNodesResponse) {
    response = &ListOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationNodes
// This API is used to obtain a list of organizational units.
//
// error code that may be returned:
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationNodes(request *ListOrganizationNodesRequest) (response *ListOrganizationNodesResponse, err error) {
    return c.ListOrganizationNodesWithContext(context.Background(), request)
}

// ListOrganizationNodes
// This API is used to obtain a list of organizational units.
//
// error code that may be returned:
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationNodesWithContext(ctx context.Context, request *ListOrganizationNodesRequest) (response *ListOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewListOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewMoveOrganizationMembersToNodeRequest() (request *MoveOrganizationMembersToNodeRequest) {
    request = &MoveOrganizationMembersToNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationMembersToNode")
    
    
    return
}

func NewMoveOrganizationMembersToNodeResponse() (response *MoveOrganizationMembersToNodeResponse) {
    response = &MoveOrganizationMembersToNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MoveOrganizationMembersToNode
// This API is used to move members to a specified organizational unit.
//
// error code that may be returned:
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationMembersToNode(request *MoveOrganizationMembersToNodeRequest) (response *MoveOrganizationMembersToNodeResponse, err error) {
    return c.MoveOrganizationMembersToNodeWithContext(context.Background(), request)
}

// MoveOrganizationMembersToNode
// This API is used to move members to a specified organizational unit.
//
// error code that may be returned:
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationMembersToNodeWithContext(ctx context.Context, request *MoveOrganizationMembersToNodeRequest) (response *MoveOrganizationMembersToNodeResponse, err error) {
    if request == nil {
        request = NewMoveOrganizationMembersToNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveOrganizationMembersToNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveOrganizationMembersToNodeResponse()
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
// This API is used to quit an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) QuitOrganization(request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
    return c.QuitOrganizationWithContext(context.Background(), request)
}

// QuitOrganization
// This API is used to quit an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) QuitOrganizationWithContext(ctx context.Context, request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
    if request == nil {
        request = NewQuitOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuitOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewQuitOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewSendOrganizationInvitationRequest() (request *SendOrganizationInvitationRequest) {
    request = &SendOrganizationInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "SendOrganizationInvitation")
    
    
    return
}

func NewSendOrganizationInvitationResponse() (response *SendOrganizationInvitationResponse) {
    response = &SendOrganizationInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendOrganizationInvitation
// This API is used to send an invitation to join an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"
//  FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"
//  FAILEDOPERATION_USERINORGANIZATION = "FailedOperation.UserInOrganization"
//  FAILEDOPERATION_USERNOTREGISTER = "FailedOperation.UserNotRegister"
//  LIMITEXCEEDED_MEMBERS = "LimitExceeded.Members"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SendOrganizationInvitation(request *SendOrganizationInvitationRequest) (response *SendOrganizationInvitationResponse, err error) {
    return c.SendOrganizationInvitationWithContext(context.Background(), request)
}

// SendOrganizationInvitation
// This API is used to send an invitation to join an organization.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"
//  FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"
//  FAILEDOPERATION_USERINORGANIZATION = "FailedOperation.UserInOrganization"
//  FAILEDOPERATION_USERNOTREGISTER = "FailedOperation.UserNotRegister"
//  LIMITEXCEEDED_MEMBERS = "LimitExceeded.Members"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) SendOrganizationInvitationWithContext(ctx context.Context, request *SendOrganizationInvitationRequest) (response *SendOrganizationInvitationResponse, err error) {
    if request == nil {
        request = NewSendOrganizationInvitationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendOrganizationInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendOrganizationInvitationResponse()
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
// This API is used to update information on organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationMember(request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
    return c.UpdateOrganizationMemberWithContext(context.Background(), request)
}

// UpdateOrganizationMember
// This API is used to update information on organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationMemberWithContext(ctx context.Context, request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMemberResponse()
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
// This API is used to update organizational units.
//
// error code that may be returned:
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  RESOURCEINUSE_NODENAMEUSED = "ResourceInUse.NodeNameUsed"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNode(request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    return c.UpdateOrganizationNodeWithContext(context.Background(), request)
}

// UpdateOrganizationNode
// This API is used to update organizational units.
//
// error code that may be returned:
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  RESOURCEINUSE_NODENAMEUSED = "ResourceInUse.NodeNameUsed"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
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
