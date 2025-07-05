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

package v20210420

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-04-20"

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


func NewAddAccountToAccountGroupRequest() (request *AddAccountToAccountGroupRequest) {
    request = &AddAccountToAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "AddAccountToAccountGroup")
    
    
    return
}

func NewAddAccountToAccountGroupResponse() (response *AddAccountToAccountGroupResponse) {
    response = &AddAccountToAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAccountToAccountGroup
// This API is used to add an account to an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTALREADYEXISTEDINACCOUNTGROUP = "FailedOperation.AccountAlreadyExistedInAccountGroup"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTIDSISNULL = "FailedOperation.AccountIdsIsNull"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) AddAccountToAccountGroup(request *AddAccountToAccountGroupRequest) (response *AddAccountToAccountGroupResponse, err error) {
    return c.AddAccountToAccountGroupWithContext(context.Background(), request)
}

// AddAccountToAccountGroup
// This API is used to add an account to an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTALREADYEXISTEDINACCOUNTGROUP = "FailedOperation.AccountAlreadyExistedInAccountGroup"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTIDSISNULL = "FailedOperation.AccountIdsIsNull"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) AddAccountToAccountGroupWithContext(ctx context.Context, request *AddAccountToAccountGroupRequest) (response *AddAccountToAccountGroupResponse, err error) {
    if request == nil {
        request = NewAddAccountToAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAccountToAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAccountToAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserToUserGroupRequest() (request *AddUserToUserGroupRequest) {
    request = &AddUserToUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "AddUserToUserGroup")
    
    
    return
}

func NewAddUserToUserGroupResponse() (response *AddUserToUserGroupResponse) {
    response = &AddUserToUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserToUserGroup
// This API is used to add a user to a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDUSERSTOUSERGROUP = "FailedOperation.AddUsersToUserGroup"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERALREADYEXISTEDINUSERGROUP = "FailedOperation.UserAlreadyExistedInUserGroup"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) AddUserToUserGroup(request *AddUserToUserGroupRequest) (response *AddUserToUserGroupResponse, err error) {
    return c.AddUserToUserGroupWithContext(context.Background(), request)
}

// AddUserToUserGroup
// This API is used to add a user to a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDUSERSTOUSERGROUP = "FailedOperation.AddUsersToUserGroup"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERALREADYEXISTEDINUSERGROUP = "FailedOperation.UserAlreadyExistedInUserGroup"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) AddUserToUserGroupWithContext(ctx context.Context, request *AddUserToUserGroupRequest) (response *AddUserToUserGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserToUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserToUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountGroupRequest() (request *CreateAccountGroupRequest) {
    request = &CreateAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateAccountGroup")
    
    
    return
}

func NewCreateAccountGroupResponse() (response *CreateAccountGroupResponse) {
    response = &CreateAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccountGroup
// This API is used to create an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEEXISTED = "FailedOperation.AccountGroupNameExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAccountGroup(request *CreateAccountGroupRequest) (response *CreateAccountGroupResponse, err error) {
    return c.CreateAccountGroupWithContext(context.Background(), request)
}

// CreateAccountGroup
// This API is used to create an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEEXISTED = "FailedOperation.AccountGroupNameExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAccountGroupWithContext(ctx context.Context, request *CreateAccountGroupRequest) (response *CreateAccountGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppAccountRequest() (request *CreateAppAccountRequest) {
    request = &CreateAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateAppAccount")
    
    
    return
}

func NewCreateAppAccountResponse() (response *CreateAppAccountResponse) {
    response = &CreateAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAppAccount
// This API is used to create an application account.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNAMEEXISTED = "FailedOperation.AccountNameExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAppAccount(request *CreateAppAccountRequest) (response *CreateAppAccountResponse, err error) {
    return c.CreateAppAccountWithContext(context.Background(), request)
}

// CreateAppAccount
// This API is used to create an application account.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNAMEEXISTED = "FailedOperation.AccountNameExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAppAccountWithContext(ctx context.Context, request *CreateAppAccountRequest) (response *CreateAppAccountResponse, err error) {
    if request == nil {
        request = NewCreateAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrgNodeRequest() (request *CreateOrgNodeRequest) {
    request = &CreateOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateOrgNode")
    
    
    return
}

func NewCreateOrgNodeResponse() (response *CreateOrgNodeResponse) {
    response = &CreateOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrgNode
// This API is used to create an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CREATEORGNODEERROR = "FailedOperation.CreateOrgNodeError"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) CreateOrgNode(request *CreateOrgNodeRequest) (response *CreateOrgNodeResponse, err error) {
    return c.CreateOrgNodeWithContext(context.Background(), request)
}

// CreateOrgNode
// This API is used to create an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CREATEORGNODEERROR = "FailedOperation.CreateOrgNodeError"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) CreateOrgNodeWithContext(ctx context.Context, request *CreateOrgNodeRequest) (response *CreateOrgNodeResponse, err error) {
    if request == nil {
        request = NewCreateOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateUser")
    
    
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
//  FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LIMITQUOTANOTENOUGH = "FailedOperation.LimitQuotaNotEnough"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_USEREMAILEXISTED = "FailedOperation.UserEmailExisted"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to create a user.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LIMITQUOTANOTENOUGH = "FailedOperation.LimitQuotaNotEnough"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_USEREMAILEXISTED = "FailedOperation.UserEmailExisted"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
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

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
    request = &CreateUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateUserGroup")
    
    
    return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
    response = &CreateUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserGroup
// This API is used to create a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEUSERGROUPFAILURE = "FailedOperation.CreateUserGroupFailure"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    return c.CreateUserGroupWithContext(context.Background(), request)
}

// CreateUserGroup
// This API is used to create a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEUSERGROUPFAILURE = "FailedOperation.CreateUserGroupFailure"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateUserGroupWithContext(ctx context.Context, request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    if request == nil {
        request = NewCreateUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountGroupRequest() (request *DeleteAccountGroupRequest) {
    request = &DeleteAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteAccountGroup")
    
    
    return
}

func NewDeleteAccountGroupResponse() (response *DeleteAccountGroupResponse) {
    response = &DeleteAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccountGroup
// This API is used to delete an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteAccountGroup(request *DeleteAccountGroupRequest) (response *DeleteAccountGroupResponse, err error) {
    return c.DeleteAccountGroupWithContext(context.Background(), request)
}

// DeleteAccountGroup
// This API is used to delete an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteAccountGroupWithContext(ctx context.Context, request *DeleteAccountGroupRequest) (response *DeleteAccountGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppAccountRequest() (request *DeleteAppAccountRequest) {
    request = &DeleteAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteAppAccount")
    
    
    return
}

func NewDeleteAppAccountResponse() (response *DeleteAppAccountResponse) {
    response = &DeleteAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAppAccount
// This API is used to delete an application account.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTIDSISNULL = "FailedOperation.AccountIdsIsNull"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteAppAccount(request *DeleteAppAccountRequest) (response *DeleteAppAccountResponse, err error) {
    return c.DeleteAppAccountWithContext(context.Background(), request)
}

// DeleteAppAccount
// This API is used to delete an application account.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTIDSISNULL = "FailedOperation.AccountIdsIsNull"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteAppAccountWithContext(ctx context.Context, request *DeleteAppAccountRequest) (response *DeleteAppAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrgNodeRequest() (request *DeleteOrgNodeRequest) {
    request = &DeleteOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteOrgNode")
    
    
    return
}

func NewDeleteOrgNodeResponse() (response *DeleteOrgNodeResponse) {
    response = &DeleteOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrgNode
// This API is used to delete an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCanNotBeDeleted"
//  FAILEDOPERATION_DEFAULTORGNODECANNOTBEDELETED = "FailedOperation.DefaultOrgNodeCanNotBeDeleted"
//  FAILEDOPERATION_DELETEORGNODEERROR = "FailedOperation.DeleteOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCanNotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCanNotBeDeleted"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteOrgNode(request *DeleteOrgNodeRequest) (response *DeleteOrgNodeResponse, err error) {
    return c.DeleteOrgNodeWithContext(context.Background(), request)
}

// DeleteOrgNode
// This API is used to delete an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCanNotBeDeleted"
//  FAILEDOPERATION_DEFAULTORGNODECANNOTBEDELETED = "FailedOperation.DefaultOrgNodeCanNotBeDeleted"
//  FAILEDOPERATION_DELETEORGNODEERROR = "FailedOperation.DeleteOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCanNotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCanNotBeDeleted"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteOrgNodeWithContext(ctx context.Context, request *DeleteOrgNodeRequest) (response *DeleteOrgNodeResponse, err error) {
    if request == nil {
        request = NewDeleteOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// This API is used to delete a user by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_DELETEUSERFAILURE = "FailedOperation.DeleteUserFailure"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete a user by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_DELETEUSERFAILURE = "FailedOperation.DeleteUserFailure"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
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

func NewDeleteUserGroupRequest() (request *DeleteUserGroupRequest) {
    request = &DeleteUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteUserGroup")
    
    
    return
}

func NewDeleteUserGroupResponse() (response *DeleteUserGroupResponse) {
    response = &DeleteUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserGroup
// This API is used to delete a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSERGROUPFAILURE = "FailedOperation.DeleteUserGroupFailure"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    return c.DeleteUserGroupWithContext(context.Background(), request)
}

// DeleteUserGroup
// This API is used to delete a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSERGROUPFAILURE = "FailedOperation.DeleteUserGroupFailure"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUserGroupWithContext(ctx context.Context, request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsers
// This API is used to batch delete the users under the current node. If an error occurs when a user is deleted, the deletion of other selected users will not be affected, and the username/ID of the user who fails to be deleted will be prompted.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// This API is used to batch delete the users under the current node. If an error occurs when a user is deleted, the deletion of other selected users will not be affected, and the username/ID of the user who fails to be deleted will be prompted.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountGroupRequest() (request *DescribeAccountGroupRequest) {
    request = &DescribeAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeAccountGroup")
    
    
    return
}

func NewDescribeAccountGroupResponse() (response *DescribeAccountGroupResponse) {
    response = &DescribeAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountGroup
// This API is used to query the list of account groups.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAccountGroup(request *DescribeAccountGroupRequest) (response *DescribeAccountGroupResponse, err error) {
    return c.DescribeAccountGroupWithContext(context.Background(), request)
}

// DescribeAccountGroup
// This API is used to query the list of account groups.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAccountGroupWithContext(ctx context.Context, request *DescribeAccountGroupRequest) (response *DescribeAccountGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppAccountRequest() (request *DescribeAppAccountRequest) {
    request = &DescribeAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeAppAccount")
    
    
    return
}

func NewDescribeAppAccountResponse() (response *DescribeAppAccountResponse) {
    response = &DescribeAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppAccount
// This API is used to query the list of application accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAppAccount(request *DescribeAppAccountRequest) (response *DescribeAppAccountResponse, err error) {
    return c.DescribeAppAccountWithContext(context.Background(), request)
}

// DescribeAppAccount
// This API is used to query the list of application accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAppAccountWithContext(ctx context.Context, request *DescribeAppAccountRequest) (response *DescribeAppAccountResponse, err error) {
    if request == nil {
        request = NewDescribeAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeApplication")
    
    
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplication
// This API is used to get the information of an application.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    return c.DescribeApplicationWithContext(context.Background(), request)
}

// DescribeApplication
// This API is used to get the information of an application.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeApplicationWithContext(ctx context.Context, request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrgNodeRequest() (request *DescribeOrgNodeRequest) {
    request = &DescribeOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeOrgNode")
    
    
    return
}

func NewDescribeOrgNodeResponse() (response *DescribeOrgNodeResponse) {
    response = &DescribeOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrgNode
// This API is used to read the information of an organization node by organization node ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeOrgNode(request *DescribeOrgNodeRequest) (response *DescribeOrgNodeResponse, err error) {
    return c.DescribeOrgNodeWithContext(context.Background(), request)
}

// DescribeOrgNode
// This API is used to read the information of an organization node by organization node ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeOrgNodeWithContext(ctx context.Context, request *DescribeOrgNodeRequest) (response *DescribeOrgNodeResponse, err error) {
    if request == nil {
        request = NewDescribeOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicKeyRequest() (request *DescribePublicKeyRequest) {
    request = &DescribePublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribePublicKey")
    
    
    return
}

func NewDescribePublicKeyResponse() (response *DescribePublicKeyResponse) {
    response = &DescribePublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicKey
// This API is used to get the information of a JWT public key.
//
// error code that may be returned:
//  FAILEDOPERATION_APPIDISNULL = "FailedOperation.AppIdIsNull"
//  FAILEDOPERATION_APPIDNOTFOUND = "FailedOperation.AppIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribePublicKey(request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    return c.DescribePublicKeyWithContext(context.Background(), request)
}

// DescribePublicKey
// This API is used to get the information of a JWT public key.
//
// error code that may be returned:
//  FAILEDOPERATION_APPIDISNULL = "FailedOperation.AppIdIsNull"
//  FAILEDOPERATION_APPIDNOTFOUND = "FailedOperation.AppIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribePublicKeyWithContext(ctx context.Context, request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    if request == nil {
        request = NewDescribePublicKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupRequest() (request *DescribeUserGroupRequest) {
    request = &DescribeUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserGroup")
    
    
    return
}

func NewDescribeUserGroupResponse() (response *DescribeUserGroupResponse) {
    response = &DescribeUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserGroup
// This API is used to get the information of a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserGroup(request *DescribeUserGroupRequest) (response *DescribeUserGroupResponse, err error) {
    return c.DescribeUserGroupWithContext(context.Background(), request)
}

// DescribeUserGroup
// This API is used to get the information of a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserGroupWithContext(ctx context.Context, request *DescribeUserGroupRequest) (response *DescribeUserGroupResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserInfo")
    
    
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserInfo
// This API is used to search for a user by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    return c.DescribeUserInfoWithContext(context.Background(), request)
}

// DescribeUserInfo
// This API is used to search for a user by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"
func (c *Client) DescribeUserInfoWithContext(ctx context.Context, request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserResourcesAuthorizationRequest() (request *DescribeUserResourcesAuthorizationRequest) {
    request = &DescribeUserResourcesAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserResourcesAuthorization")
    
    
    return
}

func NewDescribeUserResourcesAuthorizationResponse() (response *DescribeUserResourcesAuthorizationResponse) {
    response = &DescribeUserResourcesAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserResourcesAuthorization
// This API is used to query the list of resource authorizations under the specified user.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserResourcesAuthorization(request *DescribeUserResourcesAuthorizationRequest) (response *DescribeUserResourcesAuthorizationResponse, err error) {
    return c.DescribeUserResourcesAuthorizationWithContext(context.Background(), request)
}

// DescribeUserResourcesAuthorization
// This API is used to query the list of resource authorizations under the specified user.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserResourcesAuthorizationWithContext(ctx context.Context, request *DescribeUserResourcesAuthorizationRequest) (response *DescribeUserResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeUserResourcesAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserResourcesAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResourcesAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserThirdPartyAccountInfoRequest() (request *DescribeUserThirdPartyAccountInfoRequest) {
    request = &DescribeUserThirdPartyAccountInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserThirdPartyAccountInfo")
    
    
    return
}

func NewDescribeUserThirdPartyAccountInfoResponse() (response *DescribeUserThirdPartyAccountInfoResponse) {
    response = &DescribeUserThirdPartyAccountInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserThirdPartyAccountInfo
// This API is used to get the third-party account binding information by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserThirdPartyAccountInfo(request *DescribeUserThirdPartyAccountInfoRequest) (response *DescribeUserThirdPartyAccountInfoResponse, err error) {
    return c.DescribeUserThirdPartyAccountInfoWithContext(context.Background(), request)
}

// DescribeUserThirdPartyAccountInfo
// This API is used to get the third-party account binding information by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserThirdPartyAccountInfoWithContext(ctx context.Context, request *DescribeUserThirdPartyAccountInfoRequest) (response *DescribeUserThirdPartyAccountInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserThirdPartyAccountInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserThirdPartyAccountInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserThirdPartyAccountInfoResponse()
    err = c.Send(request, response)
    return
}

func NewListAccountInAccountGroupRequest() (request *ListAccountInAccountGroupRequest) {
    request = &ListAccountInAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAccountInAccountGroup")
    
    
    return
}

func NewListAccountInAccountGroupResponse() (response *ListAccountInAccountGroupResponse) {
    response = &ListAccountInAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAccountInAccountGroup
//  This API is used to get the list of accounts in an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAccountInAccountGroup(request *ListAccountInAccountGroupRequest) (response *ListAccountInAccountGroupResponse, err error) {
    return c.ListAccountInAccountGroupWithContext(context.Background(), request)
}

// ListAccountInAccountGroup
//  This API is used to get the list of accounts in an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAccountInAccountGroupWithContext(ctx context.Context, request *ListAccountInAccountGroupRequest) (response *ListAccountInAccountGroupResponse, err error) {
    if request == nil {
        request = NewListAccountInAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAccountInAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAccountInAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewListApplicationAuthorizationsRequest() (request *ListApplicationAuthorizationsRequest) {
    request = &ListApplicationAuthorizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListApplicationAuthorizations")
    
    
    return
}

func NewListApplicationAuthorizationsResponse() (response *ListApplicationAuthorizationsResponse) {
    response = &ListApplicationAuthorizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListApplicationAuthorizations
// This API is used to get the list of authorization relationships of an application (including search criteria-based match).
//
// error code that may be returned:
//  FAILEDOPERATION_ENTITYTYPENOTEXISTED = "FailedOperation.EntityTypeNotExisted"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListApplicationAuthorizations(request *ListApplicationAuthorizationsRequest) (response *ListApplicationAuthorizationsResponse, err error) {
    return c.ListApplicationAuthorizationsWithContext(context.Background(), request)
}

// ListApplicationAuthorizations
// This API is used to get the list of authorization relationships of an application (including search criteria-based match).
//
// error code that may be returned:
//  FAILEDOPERATION_ENTITYTYPENOTEXISTED = "FailedOperation.EntityTypeNotExisted"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListApplicationAuthorizationsWithContext(ctx context.Context, request *ListApplicationAuthorizationsRequest) (response *ListApplicationAuthorizationsResponse, err error) {
    if request == nil {
        request = NewListApplicationAuthorizationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApplicationAuthorizations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListApplicationAuthorizationsResponse()
    err = c.Send(request, response)
    return
}

func NewListApplicationsRequest() (request *ListApplicationsRequest) {
    request = &ListApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListApplications")
    
    
    return
}

func NewListApplicationsResponse() (response *ListApplicationsResponse) {
    response = &ListApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListApplications
// This API is used to get the list of applications.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListApplications(request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
    return c.ListApplicationsWithContext(context.Background(), request)
}

// ListApplications
// This API is used to get the list of applications.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListApplicationsWithContext(ctx context.Context, request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
    if request == nil {
        request = NewListApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewListApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewListAuthorizedApplicationsToOrgNodeRequest() (request *ListAuthorizedApplicationsToOrgNodeRequest) {
    request = &ListAuthorizedApplicationsToOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAuthorizedApplicationsToOrgNode")
    
    
    return
}

func NewListAuthorizedApplicationsToOrgNodeResponse() (response *ListAuthorizedApplicationsToOrgNodeResponse) {
    response = &ListAuthorizedApplicationsToOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAuthorizedApplicationsToOrgNode
// This API is used to get the list of accessible applications by organization node ID.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToOrgNode(request *ListAuthorizedApplicationsToOrgNodeRequest) (response *ListAuthorizedApplicationsToOrgNodeResponse, err error) {
    return c.ListAuthorizedApplicationsToOrgNodeWithContext(context.Background(), request)
}

// ListAuthorizedApplicationsToOrgNode
// This API is used to get the list of accessible applications by organization node ID.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToOrgNodeWithContext(ctx context.Context, request *ListAuthorizedApplicationsToOrgNodeRequest) (response *ListAuthorizedApplicationsToOrgNodeResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAuthorizedApplicationsToOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuthorizedApplicationsToOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewListAuthorizedApplicationsToUserRequest() (request *ListAuthorizedApplicationsToUserRequest) {
    request = &ListAuthorizedApplicationsToUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAuthorizedApplicationsToUser")
    
    
    return
}

func NewListAuthorizedApplicationsToUserResponse() (response *ListAuthorizedApplicationsToUserResponse) {
    response = &ListAuthorizedApplicationsToUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAuthorizedApplicationsToUser
// This API is used to get the list of accessible applications by user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_IDTOCODEERROR = "FailedOperation.IdToCodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERAUTHLISTERROR = "FailedOperation.UserAuthListError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUser(request *ListAuthorizedApplicationsToUserRequest) (response *ListAuthorizedApplicationsToUserResponse, err error) {
    return c.ListAuthorizedApplicationsToUserWithContext(context.Background(), request)
}

// ListAuthorizedApplicationsToUser
// This API is used to get the list of accessible applications by user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_IDTOCODEERROR = "FailedOperation.IdToCodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERAUTHLISTERROR = "FailedOperation.UserAuthListError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUserWithContext(ctx context.Context, request *ListAuthorizedApplicationsToUserRequest) (response *ListAuthorizedApplicationsToUserResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAuthorizedApplicationsToUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuthorizedApplicationsToUserResponse()
    err = c.Send(request, response)
    return
}

func NewListAuthorizedApplicationsToUserGroupRequest() (request *ListAuthorizedApplicationsToUserGroupRequest) {
    request = &ListAuthorizedApplicationsToUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAuthorizedApplicationsToUserGroup")
    
    
    return
}

func NewListAuthorizedApplicationsToUserGroupResponse() (response *ListAuthorizedApplicationsToUserGroupResponse) {
    response = &ListAuthorizedApplicationsToUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAuthorizedApplicationsToUserGroup
// This API is used to get the list of accessible applications by user group ID.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERGROUPNOTEXIST = "FailedOperation.UserGroupNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUserGroup(request *ListAuthorizedApplicationsToUserGroupRequest) (response *ListAuthorizedApplicationsToUserGroupResponse, err error) {
    return c.ListAuthorizedApplicationsToUserGroupWithContext(context.Background(), request)
}

// ListAuthorizedApplicationsToUserGroup
// This API is used to get the list of accessible applications by user group ID.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERGROUPNOTEXIST = "FailedOperation.UserGroupNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUserGroupWithContext(ctx context.Context, request *ListAuthorizedApplicationsToUserGroupRequest) (response *ListAuthorizedApplicationsToUserGroupResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAuthorizedApplicationsToUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuthorizedApplicationsToUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewListUserGroupsRequest() (request *ListUserGroupsRequest) {
    request = &ListUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUserGroups")
    
    
    return
}

func NewListUserGroupsResponse() (response *ListUserGroupsResponse) {
    response = &ListUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUserGroups
// This API is used to get the information of the user group list (including query conditions).
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUserGroups(request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    return c.ListUserGroupsWithContext(context.Background(), request)
}

// ListUserGroups
// This API is used to get the information of the user group list (including query conditions).
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUserGroupsWithContext(ctx context.Context, request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    if request == nil {
        request = NewListUserGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListUserGroupsOfUserRequest() (request *ListUserGroupsOfUserRequest) {
    request = &ListUserGroupsOfUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUserGroupsOfUser")
    
    
    return
}

func NewListUserGroupsOfUserResponse() (response *ListUserGroupsOfUserResponse) {
    response = &ListUserGroupsOfUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUserGroupsOfUser
// This API is used to get the list of a user's user groups.
//
// error code that may be returned:
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERERROR = "FailedOperation.ListUserGroupsOfUserError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_USERINFOSORTKEYISILLEGAL = "InvalidParameter.UserInfoSortKeyIsIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUserGroupsOfUser(request *ListUserGroupsOfUserRequest) (response *ListUserGroupsOfUserResponse, err error) {
    return c.ListUserGroupsOfUserWithContext(context.Background(), request)
}

// ListUserGroupsOfUser
// This API is used to get the list of a user's user groups.
//
// error code that may be returned:
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERERROR = "FailedOperation.ListUserGroupsOfUserError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_USERINFOSORTKEYISILLEGAL = "InvalidParameter.UserInfoSortKeyIsIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUserGroupsOfUserWithContext(ctx context.Context, request *ListUserGroupsOfUserRequest) (response *ListUserGroupsOfUserResponse, err error) {
    if request == nil {
        request = NewListUserGroupsOfUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserGroupsOfUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserGroupsOfUserResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUsers")
    
    
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsers
// This API is used to get the information of the user list.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETER_USERINFOSORTKEYISILLEGAL = "InvalidParameter.UserInfoSortKeyIsIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    return c.ListUsersWithContext(context.Background(), request)
}

// ListUsers
// This API is used to get the information of the user list.
//
// error code that may be returned:
//  FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETER_USERINFOSORTKEYISILLEGAL = "InvalidParameter.UserInfoSortKeyIsIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
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

func NewListUsersInOrgNodeRequest() (request *ListUsersInOrgNodeRequest) {
    request = &ListUsersInOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUsersInOrgNode")
    
    
    return
}

func NewListUsersInOrgNodeResponse() (response *ListUsersInOrgNodeResponse) {
    response = &ListUsersInOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsersInOrgNode
// This API is used to read the users under an organization node by organization node ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LISTUSERSINORGNODEERROR = "FailedOperation.ListUsersInOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInOrgNode(request *ListUsersInOrgNodeRequest) (response *ListUsersInOrgNodeResponse, err error) {
    return c.ListUsersInOrgNodeWithContext(context.Background(), request)
}

// ListUsersInOrgNode
// This API is used to read the users under an organization node by organization node ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LISTUSERSINORGNODEERROR = "FailedOperation.ListUsersInOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInOrgNodeWithContext(ctx context.Context, request *ListUsersInOrgNodeRequest) (response *ListUsersInOrgNodeResponse, err error) {
    if request == nil {
        request = NewListUsersInOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsersInOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersInOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersInUserGroupRequest() (request *ListUsersInUserGroupRequest) {
    request = &ListUsersInUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUsersInUserGroup")
    
    
    return
}

func NewListUsersInUserGroupResponse() (response *ListUsersInUserGroupResponse) {
    response = &ListUsersInUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsersInUserGroup
// This API is used to get the list of the users in a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_LISTUSERSINUSERGROUPERROR = "FailedOperation.ListUsersInUserGroupError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInUserGroup(request *ListUsersInUserGroupRequest) (response *ListUsersInUserGroupResponse, err error) {
    return c.ListUsersInUserGroupWithContext(context.Background(), request)
}

// ListUsersInUserGroup
// This API is used to get the list of the users in a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_LISTUSERSINUSERGROUPERROR = "FailedOperation.ListUsersInUserGroupError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInUserGroupWithContext(ctx context.Context, request *ListUsersInUserGroupRequest) (response *ListUsersInUserGroupResponse, err error) {
    if request == nil {
        request = NewListUsersInUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsersInUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersInUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountGroupRequest() (request *ModifyAccountGroupRequest) {
    request = &ModifyAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyAccountGroup")
    
    
    return
}

func NewModifyAccountGroupResponse() (response *ModifyAccountGroupResponse) {
    response = &ModifyAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountGroup
// This API is used to modify an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEEXISTED = "FailedOperation.AccountGroupNameExisted"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAccountGroup(request *ModifyAccountGroupRequest) (response *ModifyAccountGroupResponse, err error) {
    return c.ModifyAccountGroupWithContext(context.Background(), request)
}

// ModifyAccountGroup
// This API is used to modify an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEEXISTED = "FailedOperation.AccountGroupNameExisted"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAccountGroupWithContext(ctx context.Context, request *ModifyAccountGroupRequest) (response *ModifyAccountGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppAccountRequest() (request *ModifyAppAccountRequest) {
    request = &ModifyAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyAppAccount")
    
    
    return
}

func NewModifyAppAccountResponse() (response *ModifyAppAccountResponse) {
    response = &ModifyAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAppAccount
// This API is used to modify an application account.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTNAMEEXISTED = "FailedOperation.AccountNameExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAppAccount(request *ModifyAppAccountRequest) (response *ModifyAppAccountResponse, err error) {
    return c.ModifyAppAccountWithContext(context.Background(), request)
}

// ModifyAppAccount
// This API is used to modify an application account.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTNAMEEXISTED = "FailedOperation.AccountNameExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAppAccountWithContext(ctx context.Context, request *ModifyAppAccountRequest) (response *ModifyAppAccountResponse, err error) {
    if request == nil {
        request = NewModifyAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// This API is used to update the information of an application.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_APPDISPLAYNAMEEXISTED = "InvalidParameter.AppDisplayNameExisted"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// This API is used to update the information of an application.
//
// error code that may be returned:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_APPDISPLAYNAMEEXISTED = "InvalidParameter.AppDisplayNameExisted"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserInfoRequest() (request *ModifyUserInfoRequest) {
    request = &ModifyUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyUserInfo")
    
    
    return
}

func NewModifyUserInfoResponse() (response *ModifyUserInfoResponse) {
    response = &ModifyUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserInfo
// This API is used to modify the information of a user by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_NEWPASSWORDMUSTNOTBLANK = "FailedOperation.NewPasswordMustNotBlank"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_UPDATELDAPUSERORGEXCEEDSRANGE = "FailedOperation.UpdateLDAPUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATELDAPUSERORGNOTINSAMECROP = "FailedOperation.UpdateLDAPUserOrgNotInSameCrop"
//  FAILEDOPERATION_UPDATEUSEREXCEEDSRANGE = "FailedOperation.UpdateUserExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGEXCEEDSRANGE = "FailedOperation.UpdateWeComUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGNOTINSAMECROP = "FailedOperation.UpdateWeComUserOrgNotInSameCrop"
//  FAILEDOPERATION_USEREMAILEXISTED = "FailedOperation.UserEmailExisted"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  FAILEDOPERATION_USERPHONEISEMPTY = "FailedOperation.UserPhoneIsEmpty"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) ModifyUserInfo(request *ModifyUserInfoRequest) (response *ModifyUserInfoResponse, err error) {
    return c.ModifyUserInfoWithContext(context.Background(), request)
}

// ModifyUserInfo
// This API is used to modify the information of a user by username or user ID.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_NEWPASSWORDMUSTNOTBLANK = "FailedOperation.NewPasswordMustNotBlank"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_UPDATELDAPUSERORGEXCEEDSRANGE = "FailedOperation.UpdateLDAPUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATELDAPUSERORGNOTINSAMECROP = "FailedOperation.UpdateLDAPUserOrgNotInSameCrop"
//  FAILEDOPERATION_UPDATEUSEREXCEEDSRANGE = "FailedOperation.UpdateUserExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGEXCEEDSRANGE = "FailedOperation.UpdateWeComUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGNOTINSAMECROP = "FailedOperation.UpdateWeComUserOrgNotInSameCrop"
//  FAILEDOPERATION_USEREMAILEXISTED = "FailedOperation.UserEmailExisted"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  FAILEDOPERATION_USERPHONEISEMPTY = "FailedOperation.UserPhoneIsEmpty"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) ModifyUserInfoWithContext(ctx context.Context, request *ModifyUserInfoRequest) (response *ModifyUserInfoResponse, err error) {
    if request == nil {
        request = NewModifyUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAccountFromAccountGroupRequest() (request *RemoveAccountFromAccountGroupRequest) {
    request = &RemoveAccountFromAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "RemoveAccountFromAccountGroup")
    
    
    return
}

func NewRemoveAccountFromAccountGroupResponse() (response *RemoveAccountFromAccountGroupResponse) {
    response = &RemoveAccountFromAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveAccountFromAccountGroup
// This API is used to remove an account from an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTIDSISNULL = "FailedOperation.AccountIdsIsNull"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) RemoveAccountFromAccountGroup(request *RemoveAccountFromAccountGroupRequest) (response *RemoveAccountFromAccountGroupResponse, err error) {
    return c.RemoveAccountFromAccountGroupWithContext(context.Background(), request)
}

// RemoveAccountFromAccountGroup
// This API is used to remove an account from an account group.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTIDSISNULL = "FailedOperation.AccountIdsIsNull"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) RemoveAccountFromAccountGroupWithContext(ctx context.Context, request *RemoveAccountFromAccountGroupRequest) (response *RemoveAccountFromAccountGroupResponse, err error) {
    if request == nil {
        request = NewRemoveAccountFromAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAccountFromAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAccountFromAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserFromUserGroupRequest() (request *RemoveUserFromUserGroupRequest) {
    request = &RemoveUserFromUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "RemoveUserFromUserGroup")
    
    
    return
}

func NewRemoveUserFromUserGroupResponse() (response *RemoveUserFromUserGroupResponse) {
    response = &RemoveUserFromUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveUserFromUserGroup
// This API is used to remove a user from a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPERROR = "FailedOperation.RemoveUsersFromUserGroupError"
//  FAILEDOPERATION_USERNOTEXISTINUSERGROUP = "FailedOperation.UserNotExistInUserGroup"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) RemoveUserFromUserGroup(request *RemoveUserFromUserGroupRequest) (response *RemoveUserFromUserGroupResponse, err error) {
    return c.RemoveUserFromUserGroupWithContext(context.Background(), request)
}

// RemoveUserFromUserGroup
// This API is used to remove a user from a user group.
//
// error code that may be returned:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPERROR = "FailedOperation.RemoveUsersFromUserGroupError"
//  FAILEDOPERATION_USERNOTEXISTINUSERGROUP = "FailedOperation.UserNotExistInUserGroup"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) RemoveUserFromUserGroupWithContext(ctx context.Context, request *RemoveUserFromUserGroupRequest) (response *RemoveUserFromUserGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserFromUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserFromUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrgNodeRequest() (request *UpdateOrgNodeRequest) {
    request = &UpdateOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "UpdateOrgNode")
    
    
    return
}

func NewUpdateOrgNodeResponse() (response *UpdateOrgNodeResponse) {
    response = &UpdateOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrgNode
// This API is used to create an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) UpdateOrgNode(request *UpdateOrgNodeRequest) (response *UpdateOrgNodeResponse, err error) {
    return c.UpdateOrgNodeWithContext(context.Background(), request)
}

// UpdateOrgNode
// This API is used to create an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) UpdateOrgNodeWithContext(ctx context.Context, request *UpdateOrgNodeRequest) (response *UpdateOrgNodeResponse, err error) {
    if request == nil {
        request = NewUpdateOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrgNodeResponse()
    err = c.Send(request, response)
    return
}
