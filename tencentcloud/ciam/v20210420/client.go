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


func NewListUserGroupsRequest() (request *ListUserGroupsRequest) {
    request = &ListUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ciam", APIVersion, "ListUserGroups")
    
    
    return
}

func NewListUserGroupsResponse() (response *ListUserGroupsResponse) {
    response = &ListUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUserGroups
// This API is used to list user groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTNOTFOUND = "FailedOperation.AccountNotFound"
//  FAILEDOPERATION_ADDUSERSTOUSERGROUP = "FailedOperation.AddUsersToUserGroup"
//  FAILEDOPERATION_APPIDISNULL = "FailedOperation.AppIdIsNull"
//  FAILEDOPERATION_APPIDNOTEXITED = "FailedOperation.AppIdNotExited"
//  FAILEDOPERATION_APPIDNOTFOUND = "FailedOperation.AppIdNotFound"
//  FAILEDOPERATION_APPNOTEXITED = "FailedOperation.AppNotExited"
//  FAILEDOPERATION_AUTHLISTFAILED = "FailedOperation.AuthListFailed"
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCannotBeDeleted"
//  FAILEDOPERATION_CREATEORGNODEFAILURE = "FailedOperation.CreateOrgNodeFailure"
//  FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"
//  FAILEDOPERATION_CREATEUSERGROUPFAILURE = "FailedOperation.CreateUserGroupFailure"
//  FAILEDOPERATION_CUSTOMIZEDPARENTORGNODEIDEXISTED = "FailedOperation.CustomizedParentOrgNodeIdExisted"
//  FAILEDOPERATION_DELETEORGNODEFAILURE = "FailedOperation.DeleteOrgNodeFailure"
//  FAILEDOPERATION_DELETEUSEREXCEPTION = "FailedOperation.DeleteUserException"
//  FAILEDOPERATION_DELETEUSERFAILURE = "FailedOperation.DeleteUserFailure"
//  FAILEDOPERATION_DELETEUSERGROUPFAILURE = "FailedOperation.DeleteUserGroupFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEFAILURE = "FailedOperation.DescribeOrgNodeFailure"
//  FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_LISTALLUSERGROUPSFAILURE = "FailedOperation.ListAllUserGroupsFailure"
//  FAILEDOPERATION_LISTMATCHEDUSERINFOFAILURE = "FailedOperation.ListMatchedUserInfoFailure"
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERFAILURE = "FailedOperation.ListUserGroupsOfUserFailure"
//  FAILEDOPERATION_LISTUSERSINORGNODEFAILURE = "FailedOperation.ListUsersInOrgNodeFailure"
//  FAILEDOPERATION_LISTUSERSINUSERGROUPFAILURE = "FailedOperation.ListUsersInUserGroupFailure"
//  FAILEDOPERATION_MODIFYAPPLICATIONDISPLAYNAMEISNULL = "FailedOperation.ModifyApplicationDisplayNameIsNull"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCannotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCannotBeDeleted"
//  FAILEDOPERATION_PAGEPARAMETERISNOTPAIRED = "FailedOperation.PageParameterIsNotPaired"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  FAILEDOPERATION_PARENTORGNODEISEMPTY = "FailedOperation.ParentOrgNodeIsEmpty"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPFAILURE = "FailedOperation.RemoveUsersFromUserGroupFailure"
//  FAILEDOPERATION_REPEATEDAUTHORIZATIONEXCEPTION = "FailedOperation.RepeatedAuthorizationException"
//  FAILEDOPERATION_REPEATEDUNAUTHORIZATIONEXCEPTION = "FailedOperation.RepeatedUnauthorizationException"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  FAILEDOPERATION_USERALREADYEXISTEDINUSERGROUP = "FailedOperation.UserAlreadyExistedInUserGroup"
//  FAILEDOPERATION_USEREXPRIATIONTIMEISILLEGAL = "FailedOperation.UserExpriationTimeIsIllegal"
//  FAILEDOPERATION_USERGROUPNOTEXISTED = "FailedOperation.UserGroupNotExisted"
//  FAILEDOPERATION_USERIDNOTFOUND = "FailedOperation.UserIdNotFound"
//  FAILEDOPERATION_USERINFOSORTATTRIBUTECODEISILLEGAL = "FailedOperation.UserInfoSortAttributeCodeIsIllegal"
//  FAILEDOPERATION_USERNAMEEXISTED = "FailedOperation.UserNameExisted"
//  FAILEDOPERATION_USERNOTEXISTEDINUSERGROUP = "FailedOperation.UserNotExistedInUserGroup"
//  FAILEDOPERATION_USERPHONEEXISTED = "FailedOperation.UserPhoneExisted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDFAILURE = "InvalidParameter.AttributeValueValidFailure"
//  INVALIDPARAMETER_DATATYPENOTMATCH = "InvalidParameter.DataTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERISILLEGAL = "InvalidParameter.ParameterIsIllegal"
//  INVALIDPARAMETER_PARAMETERTYPEISILLEGAL = "InvalidParameter.ParameterTypeIsIllegal"
//  INVALIDPARAMETER_PASSWORDISILLEGAL = "InvalidParameter.PasswordIsIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DATATYPEISILLEGAL = "OperationDenied.DataTypeIsIllegal"
//  OPERATIONDENIED_UINNOTEXISTED = "OperationDenied.UinNotExisted"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListUserGroups(request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    return c.ListUserGroupsWithContext(context.Background(), request)
}

// ListUserGroups
// This API is used to list user groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTNOTFOUND = "FailedOperation.AccountNotFound"
//  FAILEDOPERATION_ADDUSERSTOUSERGROUP = "FailedOperation.AddUsersToUserGroup"
//  FAILEDOPERATION_APPIDISNULL = "FailedOperation.AppIdIsNull"
//  FAILEDOPERATION_APPIDNOTEXITED = "FailedOperation.AppIdNotExited"
//  FAILEDOPERATION_APPIDNOTFOUND = "FailedOperation.AppIdNotFound"
//  FAILEDOPERATION_APPNOTEXITED = "FailedOperation.AppNotExited"
//  FAILEDOPERATION_AUTHLISTFAILED = "FailedOperation.AuthListFailed"
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCannotBeDeleted"
//  FAILEDOPERATION_CREATEORGNODEFAILURE = "FailedOperation.CreateOrgNodeFailure"
//  FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"
//  FAILEDOPERATION_CREATEUSERGROUPFAILURE = "FailedOperation.CreateUserGroupFailure"
//  FAILEDOPERATION_CUSTOMIZEDPARENTORGNODEIDEXISTED = "FailedOperation.CustomizedParentOrgNodeIdExisted"
//  FAILEDOPERATION_DELETEORGNODEFAILURE = "FailedOperation.DeleteOrgNodeFailure"
//  FAILEDOPERATION_DELETEUSEREXCEPTION = "FailedOperation.DeleteUserException"
//  FAILEDOPERATION_DELETEUSERFAILURE = "FailedOperation.DeleteUserFailure"
//  FAILEDOPERATION_DELETEUSERGROUPFAILURE = "FailedOperation.DeleteUserGroupFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEFAILURE = "FailedOperation.DescribeOrgNodeFailure"
//  FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_INVALIDTENANT = "FailedOperation.InvalidTenant"
//  FAILEDOPERATION_LISTALLUSERGROUPSFAILURE = "FailedOperation.ListAllUserGroupsFailure"
//  FAILEDOPERATION_LISTMATCHEDUSERINFOFAILURE = "FailedOperation.ListMatchedUserInfoFailure"
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERFAILURE = "FailedOperation.ListUserGroupsOfUserFailure"
//  FAILEDOPERATION_LISTUSERSINORGNODEFAILURE = "FailedOperation.ListUsersInOrgNodeFailure"
//  FAILEDOPERATION_LISTUSERSINUSERGROUPFAILURE = "FailedOperation.ListUsersInUserGroupFailure"
//  FAILEDOPERATION_MODIFYAPPLICATIONDISPLAYNAMEISNULL = "FailedOperation.ModifyApplicationDisplayNameIsNull"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCannotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCannotBeDeleted"
//  FAILEDOPERATION_PAGEPARAMETERISNOTPAIRED = "FailedOperation.PageParameterIsNotPaired"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  FAILEDOPERATION_PARENTORGNODEISEMPTY = "FailedOperation.ParentOrgNodeIsEmpty"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPFAILURE = "FailedOperation.RemoveUsersFromUserGroupFailure"
//  FAILEDOPERATION_REPEATEDAUTHORIZATIONEXCEPTION = "FailedOperation.RepeatedAuthorizationException"
//  FAILEDOPERATION_REPEATEDUNAUTHORIZATIONEXCEPTION = "FailedOperation.RepeatedUnauthorizationException"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  FAILEDOPERATION_USERALREADYEXISTEDINUSERGROUP = "FailedOperation.UserAlreadyExistedInUserGroup"
//  FAILEDOPERATION_USEREXPRIATIONTIMEISILLEGAL = "FailedOperation.UserExpriationTimeIsIllegal"
//  FAILEDOPERATION_USERGROUPNOTEXISTED = "FailedOperation.UserGroupNotExisted"
//  FAILEDOPERATION_USERIDNOTFOUND = "FailedOperation.UserIdNotFound"
//  FAILEDOPERATION_USERINFOSORTATTRIBUTECODEISILLEGAL = "FailedOperation.UserInfoSortAttributeCodeIsIllegal"
//  FAILEDOPERATION_USERNAMEEXISTED = "FailedOperation.UserNameExisted"
//  FAILEDOPERATION_USERNOTEXISTEDINUSERGROUP = "FailedOperation.UserNotExistedInUserGroup"
//  FAILEDOPERATION_USERPHONEEXISTED = "FailedOperation.UserPhoneExisted"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDFAILURE = "InvalidParameter.AttributeValueValidFailure"
//  INVALIDPARAMETER_DATATYPENOTMATCH = "InvalidParameter.DataTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERISILLEGAL = "InvalidParameter.ParameterIsIllegal"
//  INVALIDPARAMETER_PARAMETERTYPEISILLEGAL = "InvalidParameter.ParameterTypeIsIllegal"
//  INVALIDPARAMETER_PASSWORDISILLEGAL = "InvalidParameter.PasswordIsIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DATATYPEISILLEGAL = "OperationDenied.DataTypeIsIllegal"
//  OPERATIONDENIED_UINNOTEXISTED = "OperationDenied.UinNotExisted"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
