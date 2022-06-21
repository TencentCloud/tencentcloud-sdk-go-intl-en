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

package v20210420

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccountGroupInfo struct {
	// Account group ID.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// Account group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Remarks.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`
}

type AccountGroupSearchCriteria struct {
	// Keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

// Predefined struct for user
type AddAccountToAccountGroupRequestParams struct {
	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// List of IDs of the accounts to be added to the account group.
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`
}

type AddAccountToAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// List of IDs of the accounts to be added to the account group.
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`
}

func (r *AddAccountToAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAccountToAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	delete(f, "AccountIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAccountToAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAccountToAccountGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddAccountToAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddAccountToAccountGroupResponseParams `json:"Response"`
}

func (r *AddAccountToAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAccountToAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToUserGroupRequestParams struct {
	// List of IDs of the users to be added to the user group.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

type AddUserToUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the users to be added to the user group.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *AddUserToUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserToUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToUserGroupResponseParams struct {
	// List of IDs of the users failed to be added to the user group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailedItems []*string `json:"FailedItems,omitempty" name:"FailedItems"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddUserToUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddUserToUserGroupResponseParams `json:"Response"`
}

func (r *AddUserToUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppAccountInfo struct {
	// Account ID.
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// Account name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// User information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserList []*LinkUserInfo `json:"UserList,omitempty" name:"UserList"`

	// Description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`
}

type AppAccountSearchCriteria struct {
	// Keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type ApplicationAuthorizationInfo struct {
	// List of the user's accounts under authorized applications
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationAccounts []*string `json:"ApplicationAccounts,omitempty" name:"ApplicationAccounts"`

	// Application ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// List of IDs of the user's user groups and organization nodes that have access to the application.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InheritedForm *InheritedForm `json:"InheritedForm,omitempty" name:"InheritedForm"`

	// Application name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Application creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`
}

type ApplicationInfoSearchCriteria struct {
	// Application search keyword, which can be application name or ID.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// Application type. Valid values: OAUTH2, JWT, CAS, SAML2, FORM, OIDC, APIGW
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`
}

type ApplicationInformation struct {
	// Application ID, which is globally unique.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Displayed application name, which can contain up to 64 characters and is the same as the application name by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Application creation time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// Last update time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// Application status.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

	// Application icon.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// Application type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Client ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
}

type AuthorizationInfo struct {
	// Unique application ID.
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// Application name.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// Type name.
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`

	// Unique type ID.
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// Last update time in ISO 8601 format.
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// Unique authorization type ID.
	AuthorizationId *string `json:"AuthorizationId,omitempty" name:"AuthorizationId"`
}

type AuthorizationInfoSearchCriteria struct {
	// Search by name. When the query type is user, the match criteria include username and application name. When the query type is user group, the match criteria include user group name and application name. When the query type is organization, the match criteria include organization name and application name.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type AuthorizationUserResouceInfo struct {
	// Resource ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Resource type
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Authorized resource
	// Note: this field may return null, indicating that no valid values can be obtained.
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Inheritance relationship
	// Note: this field may return null, indicating that no valid values can be obtained.
	InheritedForm *InheritedForm `json:"InheritedForm,omitempty" name:"InheritedForm"`

	// Application account
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationAccounts []*string `json:"ApplicationAccounts,omitempty" name:"ApplicationAccounts"`

	// Resource name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

// Predefined struct for user
type CreateAccountGroupRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Account group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Account group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "GroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountGroupResponseParams struct {
	// Account group ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountGroupResponseParams `json:"Response"`
}

func (r *CreateAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppAccountRequestParams struct {
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Account password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAppAccountRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Account name
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Account password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAppAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "AccountName")
	delete(f, "Password")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppAccountResponseParams struct {
	// Account ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAppAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppAccountResponseParams `json:"Response"`
}

func (r *CreateAppAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrgNodeRequestParams struct {
	// Organization node name, which can contain up to 64 characters.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Parent organization node ID. If this parameter is left empty, the organization will be created under the root organization node by default.
	ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

	// Organization node description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// External ID of the organization node, which is optional and customizable. If this parameter is specified, its uniqueness will be verified.
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`
}

type CreateOrgNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organization node name, which can contain up to 64 characters.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Parent organization node ID. If this parameter is left empty, the organization will be created under the root organization node by default.
	ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

	// Organization node description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// External ID of the organization node, which is optional and customizable. If this parameter is specified, its uniqueness will be verified.
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`
}

func (r *CreateOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayName")
	delete(f, "ParentOrgNodeId")
	delete(f, "Description")
	delete(f, "CustomizedOrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrgNodeResponseParams struct {
	// Organization node ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrgNodeResponseParams `json:"Response"`
}

func (r *CreateOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupRequestParams struct {
	// User group nickname, which is unique and can contain up to 64 characters.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User group remarks, which can contain up to 512 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// User group nickname, which is unique and can contain up to 64 characters.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User group remarks, which can contain up to 512 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupResponseParams struct {
	// User group ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserGroupResponseParams `json:"Response"`
}

func (r *CreateUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// Username, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User password, which needs to be configured according to the password policy.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User remarks, which can contain up to 512 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of IDs of the user's user groups.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// User's mobile number, such as `+86-1xxxxxxxxxx`.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// Unique ID of the user's primary organization. If this parameter is left empty, the user will be created under the root node by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// User expiration time in ISO 8601 format.
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// User's email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Whether the password needs to be reset. Default value: false (no).
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`

	// List of IDs of the user's secondary organizations.
	SecondaryOrgNodeIdList []*string `json:"SecondaryOrgNodeIdList,omitempty" name:"SecondaryOrgNodeIdList"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// Username, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User password, which needs to be configured according to the password policy.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User remarks, which can contain up to 512 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of IDs of the user's user groups.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// User's mobile number, such as `+86-1xxxxxxxxxx`.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// Unique ID of the user's primary organization. If this parameter is left empty, the user will be created under the root node by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// User expiration time in ISO 8601 format.
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// User's email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Whether the password needs to be reset. Default value: false (no).
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`

	// List of IDs of the user's secondary organizations.
	SecondaryOrgNodeIdList []*string `json:"SecondaryOrgNodeIdList,omitempty" name:"SecondaryOrgNodeIdList"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "UserGroupIds")
	delete(f, "Phone")
	delete(f, "OrgNodeId")
	delete(f, "ExpirationTime")
	delete(f, "Email")
	delete(f, "PwdNeedReset")
	delete(f, "SecondaryOrgNodeIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// Returned ID of the newly created user, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountGroupRequestParams struct {
	// Array of account group IDs.
	AccountGroupIdList []*string `json:"AccountGroupIdList,omitempty" name:"AccountGroupIdList"`
}

type DeleteAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Array of account group IDs.
	AccountGroupIdList []*string `json:"AccountGroupIdList,omitempty" name:"AccountGroupIdList"`
}

func (r *DeleteAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountGroupResponseParams `json:"Response"`
}

func (r *DeleteAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppAccountRequestParams struct {
	// Array of account IDs .
	AccountIdList []*string `json:"AccountIdList,omitempty" name:"AccountIdList"`
}

type DeleteAppAccountRequest struct {
	*tchttp.BaseRequest
	
	// Array of account IDs .
	AccountIdList []*string `json:"AccountIdList,omitempty" name:"AccountIdList"`
}

func (r *DeleteAppAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAppAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppAccountResponseParams `json:"Response"`
}

func (r *DeleteAppAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrgNodeRequestParams struct {
	// Organization node ID, which is globally unique.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`
}

type DeleteOrgNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID, which is globally unique.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`
}

func (r *DeleteOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrgNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrgNodeResponseParams `json:"Response"`
}

func (r *DeleteOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupRequestParams struct {
	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

type DeleteUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *DeleteUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserGroupResponseParams `json:"Response"`
}

func (r *DeleteUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// Username, which can contain up to 32 characters. You need to select either `UserName` or `UserId` as the search criterion; if both are selected, `UserName` will be used by default.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User ID. You need to select either `UserName` or `UserId` as the search criterion. If both are selected, `UserName` will be used by default.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// Username, which can contain up to 32 characters. You need to select either `UserName` or `UserId` as the search criterion; if both are selected, `UserName` will be used by default.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User ID. You need to select either `UserName` or `UserId` as the search criterion. If both are selected, `UserName` will be used by default.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// List of IDs of the users to be deleted. You need to specify at least `DeleteIdList` or `DeleteNameList`. If both are specified, `DeleteNameList` will be used first.
	DeleteIdList []*string `json:"DeleteIdList,omitempty" name:"DeleteIdList"`

	// List of usernames of the users to be deleted. You need to specify at least `DeleteIdList` or `DeleteNameList`. If both are specified, `DeleteNameList` will be used first.
	DeleteNameList []*string `json:"DeleteNameList,omitempty" name:"DeleteNameList"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the users to be deleted. You need to specify at least `DeleteIdList` or `DeleteNameList`. If both are specified, `DeleteNameList` will be used first.
	DeleteIdList []*string `json:"DeleteIdList,omitempty" name:"DeleteIdList"`

	// List of usernames of the users to be deleted. You need to specify at least `DeleteIdList` or `DeleteNameList`. If both are specified, `DeleteNameList` will be used first.
	DeleteNameList []*string `json:"DeleteNameList,omitempty" name:"DeleteNameList"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeleteIdList")
	delete(f, "DeleteNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersResponseParams struct {
	// Information of the users failed to be deleted. When the business parameter is `DeleteIdList`, this field will return the list of IDs of the users failed to be deleted. When the business parameter is `DeleteNameList`, this field will return the list of usernames of the users failed to be deleted.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailedItems []*string `json:"FailedItems,omitempty" name:"FailedItems"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersResponseParams `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountGroupRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *AccountGroupSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *AccountGroupSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SearchCondition")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountGroupResponseParams struct {
	// Total number of records returned for the query.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Application ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Returned list of eligible data.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountGroupList []*AccountGroupInfo `json:"AccountGroupList,omitempty" name:"AccountGroupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountGroupResponseParams `json:"Response"`
}

func (r *DescribeAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppAccountRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *AppAccountSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAppAccountRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *AppAccountSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SearchCondition")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppAccountResponseParams struct {
	// Total number of records returned for the query.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Application ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Returned list of eligible data.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppAccountList []*AppAccountInfo `json:"AppAccountList,omitempty" name:"AppAccountList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAppAccountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppAccountResponseParams `json:"Response"`
}

func (r *DescribeAppAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationRequestParams struct {
	// Application ID, which is globally unique. You must specify at least this parameter or `ClientId`.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Client ID. You must specify at least this parameter or `ApplicationId`.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
}

type DescribeApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is globally unique. You must specify at least this parameter or `ClientId`.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Client ID. You must specify at least this parameter or `ApplicationId`.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
}

func (r *DescribeApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ClientId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationResponseParams struct {
	// Key ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Displayed application name, which can contain up to 64 characters and is the same as the application name by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Last modification time of the application in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// Client ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// Application type, i.e., the application template type selected during application creation.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Application creation time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// Application ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Token validity period in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TokenExpired *int64 `json:"TokenExpired,omitempty" name:"TokenExpired"`

	// Client secret.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`

	// Public key information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Authorization address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthorizeUrl *string `json:"AuthorizeUrl,omitempty" name:"AuthorizeUrl"`

	// Access address of the application icon image.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// Security level.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

	// Application status.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

	// Description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationResponseParams `json:"Response"`
}

func (r *DescribeApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrgNodeRequestParams struct {
	// Organization node ID, which is globally unique and can contain up to 64 characters. If this parameter is left empty, the information of the root organization node will be read by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Whether to read the information of its sub-nodes. When this parameter is left empty or specified as `false`, only the information of the current organization node will be read by default. When it is specified as `true`, the information of the current organization node and its level-1 sub-nodes will be read.
	IncludeOrgNodeChildInfo *bool `json:"IncludeOrgNodeChildInfo,omitempty" name:"IncludeOrgNodeChildInfo"`
}

type DescribeOrgNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID, which is globally unique and can contain up to 64 characters. If this parameter is left empty, the information of the root organization node will be read by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Whether to read the information of its sub-nodes. When this parameter is left empty or specified as `false`, only the information of the current organization node will be read by default. When it is specified as `true`, the information of the current organization node and its level-1 sub-nodes will be read.
	IncludeOrgNodeChildInfo *bool `json:"IncludeOrgNodeChildInfo,omitempty" name:"IncludeOrgNodeChildInfo"`
}

func (r *DescribeOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	delete(f, "IncludeOrgNodeChildInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrgNodeResponseParams struct {
	// Displayed organization node name, which can contain up to 64 characters and is the same as the organization name by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Last modification time of the organization node in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// External ID of the organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`

	// Parent node ID of the current organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

	// Organization node ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Data source.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Organization node creation time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// List of sub-nodes under the current organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeChildInfo []*OrgNodeChildInfo `json:"OrgNodeChildInfo,omitempty" name:"OrgNodeChildInfo"`

	// Organization node description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrgNodeResponseParams `json:"Response"`
}

func (r *DescribeOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicKeyRequestParams struct {
	// Application ID, which is globally unique.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribePublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is globally unique.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribePublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicKeyResponseParams struct {
	// Public key information used for JWT signature verification.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// JWT key ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Application ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicKeyResponseParams `json:"Response"`
}

func (r *DescribePublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupRequestParams struct {
	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

type DescribeUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *DescribeUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupResponseParams struct {
	// User group nickname, which is not unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User group remarks, which can contain up to 512 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// User group ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupResponseParams `json:"Response"`
}

func (r *DescribeUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoRequestParams struct {
	// Username, which can contain up to 64 characters. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User ID, which can contain up to 64 characters. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// Username, which can contain up to 64 characters. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User ID, which can contain up to 64 characters. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoResponseParams struct {
	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User status. Valid values: NORMAL: normal; FREEZE: frozen; LOCKED: locked; NOT_ENABLED: disabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Nickname
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User remarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of IDs of the user's user groups.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// User ID, which can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User's email address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitempty" name:"Email"`

	// User's mobile number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// Unique ID of the user's primary organization.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Data source
	// Note: this field may return null, indicating that no valid values can be obtained.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// User expiration time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// User activation time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ActivationTime *string `json:"ActivationTime,omitempty" name:"ActivationTime"`

	// Whether the password of the current user needs to be reset. `false` indicates no.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`

	// List of IDs of the user's secondary organizations.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SecondaryOrgNodeIdList []*string `json:"SecondaryOrgNodeIdList,omitempty" name:"SecondaryOrgNodeIdList"`

	// Whether the user is an admin. Valid values: 0: no; 1: yes.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminFlag *int64 `json:"AdminFlag,omitempty" name:"AdminFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInfoResponseParams `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResourcesAuthorizationRequestParams struct {
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// User ID. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whether the query scope includes the application access of the user groups and organizations associated with the user. Valid values: false: no; true: yes. Default value: false.
	IncludeInheritedAuthorizations *bool `json:"IncludeInheritedAuthorizations,omitempty" name:"IncludeInheritedAuthorizations"`
}

type DescribeUserResourcesAuthorizationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// User ID. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username. You need to specify at least `UserName` or `UserId`. If both are specified, `UserName` will be used first.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whether the query scope includes the application access of the user groups and organizations associated with the user. Valid values: false: no; true: yes. Default value: false.
	IncludeInheritedAuthorizations *bool `json:"IncludeInheritedAuthorizations,omitempty" name:"IncludeInheritedAuthorizations"`
}

func (r *DescribeUserResourcesAuthorizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResourcesAuthorizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "UserId")
	delete(f, "UserName")
	delete(f, "IncludeInheritedAuthorizations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserResourcesAuthorizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResourcesAuthorizationResponseParams struct {
	// Unique application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Application account.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationAccounts []*string `json:"ApplicationAccounts,omitempty" name:"ApplicationAccounts"`

	// Unique ID of the authorized user.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username of the authorized user.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Returned resource list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthorizationUserResourceList []*AuthorizationUserResouceInfo `json:"AuthorizationUserResourceList,omitempty" name:"AuthorizationUserResourceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserResourcesAuthorizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResourcesAuthorizationResponseParams `json:"Response"`
}

func (r *DescribeUserResourcesAuthorizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResourcesAuthorizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserThirdPartyAccountInfoRequestParams struct {
	// Username. You need to specify at least `Username` or `UserId`. If both are specified, `Username` will be used first.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User ID. You need to specify at least `Username` or `UserId`. If both are specified, `Username` will be used first.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeUserThirdPartyAccountInfoRequest struct {
	*tchttp.BaseRequest
	
	// Username. You need to specify at least `Username` or `UserId`. If both are specified, `Username` will be used first.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User ID. You need to specify at least `Username` or `UserId`. If both are specified, `Username` will be used first.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserThirdPartyAccountInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserThirdPartyAccountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserThirdPartyAccountInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserThirdPartyAccountInfoResponseParams struct {
	// User ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Third-Party account binding information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ThirdPartyAccounts []*ThirdPartyAccountInfo `json:"ThirdPartyAccounts,omitempty" name:"ThirdPartyAccounts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserThirdPartyAccountInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserThirdPartyAccountInfoResponseParams `json:"Response"`
}

func (r *DescribeUserThirdPartyAccountInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserThirdPartyAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InheritedForm struct {
	// List of IDs of the user's user groups.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// List of IDs of the user's organization nodes.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeIds []*string `json:"OrgNodeIds,omitempty" name:"OrgNodeIds"`
}

type LinkUserInfo struct {
	// User ID, which is globally unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

// Predefined struct for user
type ListAccountInAccountGroupRequestParams struct {
	// Account group ID.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges.
	SearchCondition *AccountGroupSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type ListAccountInAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Account group ID.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges.
	SearchCondition *AccountGroupSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListAccountInAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountInAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	delete(f, "SearchCondition")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAccountInAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountInAccountGroupResponseParams struct {
	// List of accounts returned for the query.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountList []*AppAccountInfo `json:"AccountList,omitempty" name:"AccountList"`

	// Total number of accounts returned for the query.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Account group ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAccountInAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *ListAccountInAccountGroupResponseParams `json:"Response"`
}

func (r *ListAccountInAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountInAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListApplicationAuthorizationsRequestParams struct {
	// Query type. Valid values: User: user; UserGroup: user group; OrgNode: organization.
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *AuthorizationInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. You can sort the results by last modification time (lastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by application name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListApplicationAuthorizationsRequest struct {
	*tchttp.BaseRequest
	
	// Query type. Valid values: User: user; UserGroup: user group; OrgNode: organization.
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *AuthorizationInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. You can sort the results by last modification time (lastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by application name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListApplicationAuthorizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationAuthorizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EntityType")
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListApplicationAuthorizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListApplicationAuthorizationsResponseParams struct {
	// Returned list of application authorization information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AuthorizationInfoList []*AuthorizationInfo `json:"AuthorizationInfoList,omitempty" name:"AuthorizationInfoList"`

	// Total number of returned application information items.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListApplicationAuthorizationsResponse struct {
	*tchttp.BaseResponse
	Response *ListApplicationAuthorizationsResponseParams `json:"Response"`
}

func (r *ListApplicationAuthorizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationAuthorizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListApplicationsRequestParams struct {
	// Fuzzy match search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, and an asterisk (*) at the end of the field indicates partial match. The fuzzy match search feature and the exact match query feature will not take effect at the same time. If both `SearchCondition` and `ApplicationIdList` are specified, `ApplicationIdList` will take effect by default for exact match query; otherwise, the information of all applications will be returned by default.
	SearchCondition *ApplicationInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. Valid values: DisplayName: application name; CreatedDate: creation time; LastModifiedDate: last modification time. If this field is left empty, the results will be sorted in alphabetical order by application name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Set of sort criteria. Valid values: DisplayName: application name; CreatedDate: creation time; LastModifiedDate: last modification time. If this field is left empty, the results will be sorted in alphabetical order by application name.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Application ID list, through which the corresponding application information will be matched exactly. The fuzzy match search feature and the exact match query feature will not take effect at the same time. If both `SearchCondition` and `ApplicationIdList` are specified, `ApplicationIdList` will take effect by default for exact match query; otherwise, the information of all applications will be returned by default.
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList"`
}

type ListApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// Fuzzy match search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, and an asterisk (*) at the end of the field indicates partial match. The fuzzy match search feature and the exact match query feature will not take effect at the same time. If both `SearchCondition` and `ApplicationIdList` are specified, `ApplicationIdList` will take effect by default for exact match query; otherwise, the information of all applications will be returned by default.
	SearchCondition *ApplicationInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. Valid values: DisplayName: application name; CreatedDate: creation time; LastModifiedDate: last modification time. If this field is left empty, the results will be sorted in alphabetical order by application name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Set of sort criteria. Valid values: DisplayName: application name; CreatedDate: creation time; LastModifiedDate: last modification time. If this field is left empty, the results will be sorted in alphabetical order by application name.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Application ID list, through which the corresponding application information will be matched exactly. The fuzzy match search feature and the exact match query feature will not take effect at the same time. If both `SearchCondition` and `ApplicationIdList` are specified, `ApplicationIdList` will take effect by default for exact match query; otherwise, the information of all applications will be returned by default.
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList"`
}

func (r *ListApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListApplicationsResponseParams struct {
	// Total number of returned application information items.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Returned application information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationInfoList []*ApplicationInformation `json:"ApplicationInfoList,omitempty" name:"ApplicationInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *ListApplicationsResponseParams `json:"Response"`
}

func (r *ListApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuthorizedApplicationsToOrgNodeRequestParams struct {
	// Organization node ID.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`
}

type ListAuthorizedApplicationsToOrgNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`
}

func (r *ListAuthorizedApplicationsToOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuthorizedApplicationsToOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuthorizedApplicationsToOrgNodeResponseParams struct {
	// List of IDs of the applications accessible to the organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAuthorizedApplicationsToOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *ListAuthorizedApplicationsToOrgNodeResponseParams `json:"Response"`
}

func (r *ListAuthorizedApplicationsToOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuthorizedApplicationsToUserGroupRequestParams struct {
	// User group ID.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

type ListAuthorizedApplicationsToUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// User group ID.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *ListAuthorizedApplicationsToUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuthorizedApplicationsToUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuthorizedApplicationsToUserGroupResponseParams struct {
	// List of IDs of the applications accessible to the user group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAuthorizedApplicationsToUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *ListAuthorizedApplicationsToUserGroupResponseParams `json:"Response"`
}

func (r *ListAuthorizedApplicationsToUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuthorizedApplicationsToUserRequestParams struct {
	// User ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Whether the query scope includes the application access of the user groups and organizations associated with the user. Valid values: false: no; true: yes. Default value: false.
	IncludeInheritedAuthorizations *bool `json:"IncludeInheritedAuthorizations,omitempty" name:"IncludeInheritedAuthorizations"`
}

type ListAuthorizedApplicationsToUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Whether the query scope includes the application access of the user groups and organizations associated with the user. Valid values: false: no; true: yes. Default value: false.
	IncludeInheritedAuthorizations *bool `json:"IncludeInheritedAuthorizations,omitempty" name:"IncludeInheritedAuthorizations"`
}

func (r *ListAuthorizedApplicationsToUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "IncludeInheritedAuthorizations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuthorizedApplicationsToUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuthorizedApplicationsToUserResponseParams struct {
	// List of information of the applications accessible to the user.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ApplicationAuthorizationInfo []*ApplicationAuthorizationInfo `json:"ApplicationAuthorizationInfo,omitempty" name:"ApplicationAuthorizationInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAuthorizedApplicationsToUserResponse struct {
	*tchttp.BaseResponse
	Response *ListAuthorizedApplicationsToUserResponseParams `json:"Response"`
}

func (r *ListAuthorizedApplicationsToUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsOfUserRequestParams struct {
	// User ID, which is globally unique.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Fuzzy search criterion. You can search by user group name (DisplayName). If this field is left empty, all of the user's user groups will be displayed by default.
	SearchCondition *UserGroupInformationSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. Valid values: DisplayName: user group name; UserGroupId: user group ID; CreatedDate: creation time. If this field is left empty, the results will be sorted in alphabetical order by user group name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 user groups will be returned.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 user groups will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListUserGroupsOfUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID, which is globally unique.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Fuzzy search criterion. You can search by user group name (DisplayName). If this field is left empty, all of the user's user groups will be displayed by default.
	SearchCondition *UserGroupInformationSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. Valid values: DisplayName: user group name; UserGroupId: user group ID; CreatedDate: creation time. If this field is left empty, the results will be sorted in alphabetical order by user group name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 user groups will be returned.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 user groups will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListUserGroupsOfUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsOfUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserGroupsOfUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsOfUserResponseParams struct {
	// List of IDs of the user's user groups.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// User ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// List of information of the user's user groups.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupInfoList []*UserGroupInfo `json:"UserGroupInfoList,omitempty" name:"UserGroupInfoList"`

	// Total number of returned user group information items.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUserGroupsOfUserResponse struct {
	*tchttp.BaseResponse
	Response *ListUserGroupsOfUserResponseParams `json:"Response"`
}

func (r *ListUserGroupsOfUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsOfUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsRequestParams struct {
	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *UserGroupInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. The supported attributes for sorting include user group name (DisplayName), user group ID (UserGroupId), and last modification time (LastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by user group name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Search criterion. You can combine multiple search criteria and search in multiple data ranges. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, and an empty field indicates to query the full table by default.
	SearchCondition *UserGroupInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. The supported attributes for sorting include user group name (DisplayName), user group ID (UserGroupId), and last modification time (LastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by user group name.
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsResponseParams struct {
	// Returned user group list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupList []*UserGroupInformation `json:"UserGroupList,omitempty" name:"UserGroupList"`

	// Total number of returned user group information items.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListUserGroupsResponseParams `json:"Response"`
}

func (r *ListUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersInOrgNodeRequestParams struct {
	// Organization node ID, which is globally unique and can contain up to 64 characters. If this parameter is left empty, the user information under the root organization node will be read by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Whether to read the information of its sub-nodes. When this parameter is left empty or specified as `false`, only the information of the current organization node will be read by default. When it is specified as `true`, the information of the current organization node and its level-1 sub-nodes will be read.
	IncludeOrgNodeChildInfo *bool `json:"IncludeOrgNodeChildInfo,omitempty" name:"IncludeOrgNodeChildInfo"`

	// User attribute search criterion. The supported search criteria include username, mobile number, email address, user locking status, user freezing status, creation time, and last modification time, which can also be combined. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, brackets separated by a comma ([Min,Max]) indicate query within a closed interval, braces separated by a comma ({Min,Max}) indicate query within an open interval, and a bracket and a brace can be used together (for example, {Min,Max] indicates that the minimum value is excluded and the maximum value is included in the query). Range query supports using an asterisk (for example, {20,*] indicates an interval including all data greater than 20) and querying by time period. The supported attributes include creation time (CreationTime) and last modification time (LastUpdateTime) in ISO 8601 format, such as `2021-01-13T09:44:07.182+0000`.
	SearchCondition *ListUsersInOrgNodeSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. The supported attributes for sorting include username (UserName), mobile number (Phone), email address (Email), user status (Status), creation time (CreatedDate), and last modification time (LastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by nickname (DisplayName).
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListUsersInOrgNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID, which is globally unique and can contain up to 64 characters. If this parameter is left empty, the user information under the root organization node will be read by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Whether to read the information of its sub-nodes. When this parameter is left empty or specified as `false`, only the information of the current organization node will be read by default. When it is specified as `true`, the information of the current organization node and its level-1 sub-nodes will be read.
	IncludeOrgNodeChildInfo *bool `json:"IncludeOrgNodeChildInfo,omitempty" name:"IncludeOrgNodeChildInfo"`

	// User attribute search criterion. The supported search criteria include username, mobile number, email address, user locking status, user freezing status, creation time, and last modification time, which can also be combined. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, brackets separated by a comma ([Min,Max]) indicate query within a closed interval, braces separated by a comma ({Min,Max}) indicate query within an open interval, and a bracket and a brace can be used together (for example, {Min,Max] indicates that the minimum value is excluded and the maximum value is included in the query). Range query supports using an asterisk (for example, {20,*] indicates an interval including all data greater than 20) and querying by time period. The supported attributes include creation time (CreationTime) and last modification time (LastUpdateTime) in ISO 8601 format, such as `2021-01-13T09:44:07.182+0000`.
	SearchCondition *ListUsersInOrgNodeSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. The supported attributes for sorting include username (UserName), mobile number (Phone), email address (Email), user status (Status), creation time (CreatedDate), and last modification time (LastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by nickname (DisplayName).
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListUsersInOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	delete(f, "IncludeOrgNodeChildInfo")
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersInOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersInOrgNodeResponseParams struct {
	// User information list under the organization sub-node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeChildUserInfo []*OrgNodeChildUserInfo `json:"OrgNodeChildUserInfo,omitempty" name:"OrgNodeChildUserInfo"`

	// Organization node ID, which is globally unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// User information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// Total number of users under the current organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalUserNum *int64 `json:"TotalUserNum,omitempty" name:"TotalUserNum"`

	// Organization ID path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeIdPath *string `json:"OrgNodeIdPath,omitempty" name:"OrgNodeIdPath"`

	// Organization name path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeNamePath *string `json:"OrgNodeNamePath,omitempty" name:"OrgNodeNamePath"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUsersInOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *ListUsersInOrgNodeResponseParams `json:"Response"`
}

func (r *ListUsersInOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersInOrgNodeSearchCriteria struct {
	// Username, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User's mobile number.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// User's email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// User status. Valid values: NORMAL: normal; FREEZE: frozen; LOCKED: locked; NOT_ENABLED: disabled.
	Status *string `json:"Status,omitempty" name:"Status"`

	// User creation time in ISO 8601 format.
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Last update time of the user.
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// Search by name. The match criteria include username and user's mobile number.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

// Predefined struct for user
type ListUsersInUserGroupRequestParams struct {
	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// User attribute search criterion. The supported search criteria include username, mobile number, email address, user locking status, user freezing status, creation time, and last modification time, which can also be combined. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, brackets separated by a comma ([Min,Max]) indicate query within a closed interval, braces separated by a comma ({Min,Max}) indicate query within an open interval, and a bracket and a brace can be used together (for example, {Min,Max] indicates that the minimum value is excluded and the maximum value is included in the query). Range query supports using an asterisk (for example, {20,*] indicates an interval including all data greater than 20) and querying by time period. The supported attributes include creation time (CreationTime) and last modification time (LastUpdateTime) in ISO 8601 format, such as `2021-01-13T09:44:07.182+0000`.
	SearchCondition *UserSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. The supported attributes for sorting include username (UserName), nickname (DisplayName), mobile number (Phone), email address (Email), user status (Status), creation time (CreatedDate), and last modification time (LastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by nickname (DisplayName).
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type ListUsersInUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// User attribute search criterion. The supported search criteria include username, mobile number, email address, user locking status, user freezing status, creation time, and last modification time, which can also be combined. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, brackets separated by a comma ([Min,Max]) indicate query within a closed interval, braces separated by a comma ({Min,Max}) indicate query within an open interval, and a bracket and a brace can be used together (for example, {Min,Max] indicates that the minimum value is excluded and the maximum value is included in the query). Range query supports using an asterisk (for example, {20,*] indicates an interval including all data greater than 20) and querying by time period. The supported attributes include creation time (CreationTime) and last modification time (LastUpdateTime) in ISO 8601 format, such as `2021-01-13T09:44:07.182+0000`.
	SearchCondition *UserSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// Set of sort criteria. The supported attributes for sorting include username (UserName), nickname (DisplayName), mobile number (Phone), email address (Email), user status (Status), creation time (CreatedDate), and last modification time (LastModifiedDate). If this field is left empty, the results will be sorted in alphabetical order by nickname (DisplayName).
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 50 users will be returned.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListUsersInUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersInUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersInUserGroupResponseParams struct {
	// User group ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// Returned user information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// Total number of returned user information items.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUsersInUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *ListUsersInUserGroupResponseParams `json:"Response"`
}

func (r *ListUsersInUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersRequestParams struct {
	// User attribute search criterion. The supported search criteria include username, mobile number, email address, user locking status, user freezing status, creation time, and last modification time, which can also be combined. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, brackets separated by a comma ([Min,Max]) indicate query within a closed interval, braces separated by a comma ({Min,Max}) indicate query within an open interval, and a bracket and a brace can be used together (for example, {Min,Max] indicates that the minimum value is excluded and the maximum value is included in the query). Range query supports using an asterisk (for example, {20,*] indicates an interval including all data greater than 20) and querying by time period. The supported attributes include creation time (CreationTime) and last modification time (LastUpdateTime) in ISO 8601 format, such as `2021-01-13T09:44:07.182+0000`.
	SearchCondition *UserSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// User attributes expected to be returned. All built-in user attributes will be returned by default, including user UUID (UserId), nickname (DisplayName), username (UserName), mobile number (Phone), email address (Email), status (Status), user group (SubjectGroups), organization path (OrgPath), remarks (Description), creation time (CreationTime), last modification time (LastUpdateTime), and last login time (LastLoginTime).
	ExpectedFields []*string `json:"ExpectedFields,omitempty" name:"ExpectedFields"`

	// Set of sort criteria. The supported attributes for sorting include username (UserName), nickname (DisplayName), mobile number (Phone), email address (Email), user status (Status), creation time (CreatedDate), last modification time (LastUpdateTime), and last login time (LastLoginTime). If this field is left empty, the results will be sorted in alphabetical order by nickname (DisplayName).
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 1,000 users will be returned.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 1,000 users will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to view the total number of search results. Default value: false (no).
	IncludeTotal *bool `json:"IncludeTotal,omitempty" name:"IncludeTotal"`
}

type ListUsersRequest struct {
	*tchttp.BaseRequest
	
	// User attribute search criterion. The supported search criteria include username, mobile number, email address, user locking status, user freezing status, creation time, and last modification time, which can also be combined. In addition, multiple query methods such as full match, partial match, and range match are supported. Specifically, double quotation marks ("") indicate full match, an asterisk (*) at the end of the field indicates partial match, brackets separated by a comma ([Min,Max]) indicate query within a closed interval, braces separated by a comma ({Min,Max}) indicate query within an open interval, and a bracket and a brace can be used together (for example, {Min,Max] indicates that the minimum value is excluded and the maximum value is included in the query). Range query supports using an asterisk (for example, {20,*] indicates an interval including all data greater than 20) and querying by time period. The supported attributes include creation time (CreationTime) and last modification time (LastUpdateTime) in ISO 8601 format, such as `2021-01-13T09:44:07.182+0000`.
	SearchCondition *UserSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// User attributes expected to be returned. All built-in user attributes will be returned by default, including user UUID (UserId), nickname (DisplayName), username (UserName), mobile number (Phone), email address (Email), status (Status), user group (SubjectGroups), organization path (OrgPath), remarks (Description), creation time (CreationTime), last modification time (LastUpdateTime), and last login time (LastLoginTime).
	ExpectedFields []*string `json:"ExpectedFields,omitempty" name:"ExpectedFields"`

	// Set of sort criteria. The supported attributes for sorting include username (UserName), nickname (DisplayName), mobile number (Phone), email address (Email), user status (Status), creation time (CreatedDate), last modification time (LastUpdateTime), and last login time (LastLoginTime). If this field is left empty, the results will be sorted in alphabetical order by nickname (DisplayName).
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// Pagination offset. Default value: 0. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 1,000 users will be returned.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results read per page. Default value: 50. Maximum value: 100. The `Offset` and `Limit` fields need to be used together; otherwise, the query results will not be paginated, and up to 1,000 users will be returned.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Whether to view the total number of search results. Default value: false (no).
	IncludeTotal *bool `json:"IncludeTotal,omitempty" name:"IncludeTotal"`
}

func (r *ListUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchCondition")
	delete(f, "ExpectedFields")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IncludeTotal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersResponseParams struct {
	// List of users returned for the query.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserList []*UserInformation `json:"UserList,omitempty" name:"UserList"`

	// Total number of users returned for the query, which will be returned only when the `IncludeTotal` input parameter is set to `true`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUsersResponse struct {
	*tchttp.BaseResponse
	Response *ListUsersResponseParams `json:"Response"`
}

func (r *ListUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountGroupRequestParams struct {
	// Account group ID.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// Account group name. When this parameter is not specified, the name will not be modified.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Description. When this parameter is not specified, the description will not be modified.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Account group ID.
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// Account group name. When this parameter is not specified, the name will not be modified.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Description. When this parameter is not specified, the description will not be modified.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	delete(f, "GroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountGroupResponseParams `json:"Response"`
}

func (r *ModifyAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppAccountRequestParams struct {
	// Account ID.
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// Account name. When this parameter is not specified, the name will not be modified.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Account password. When this parameter is not specified, the password will not be changed.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Description. When this parameter is not specified, the description will not be modified.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyAppAccountRequest struct {
	*tchttp.BaseRequest
	
	// Account ID.
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// Account name. When this parameter is not specified, the name will not be modified.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Account password. When this parameter is not specified, the password will not be changed.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Description. When this parameter is not specified, the description will not be modified.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAppAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountId")
	delete(f, "AccountName")
	delete(f, "Password")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAppAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppAccountResponseParams `json:"Response"`
}

func (r *ModifyAppAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationRequestParams struct {
	// Application ID, which is globally unique.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Security level.
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

	// Displayed application name, which can contain up to 32 characters and is the same as the application name by default.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Application status. Valid values: true: enabled; false: disabled.
	AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

	// Access address of the application icon image.
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// Description, which can contain up to 128 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest
	
	// Application ID, which is globally unique.
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// Security level.
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

	// Displayed application name, which can contain up to 32 characters and is the same as the application name by default.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Application status. Valid values: true: enabled; false: disabled.
	AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

	// Access address of the application icon image.
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// Description, which can contain up to 128 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SecureLevel")
	delete(f, "DisplayName")
	delete(f, "AppStatus")
	delete(f, "IconUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationResponseParams `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserInfoRequestParams struct {
	// Username, which can contain up to 32 characters. You need to select either `Username` or `UserId` as the search criterion; if both are selected, `Username` will be used by default.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User remarks, which can contain up to 512 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of IDs of the user's user groups.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// User ID. You need to select either `UserName` or `UserId` as the search criterion. If both are selected, `UserName` will be used by default.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User's mobile number.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// User expiration time in ISO 8601 format.
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// User password, which needs to be configured according to the password policy.
	Password *string `json:"Password,omitempty" name:"Password"`

	// User's email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Whether the password needs to be reset. Default value: false (no).
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`

	// Unique ID of the user's primary organization. If this parameter is left empty, the user will be created under the root node by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// List of IDs of the user's secondary organizations.
	SecondaryOrgNodeIdList []*string `json:"SecondaryOrgNodeIdList,omitempty" name:"SecondaryOrgNodeIdList"`
}

type ModifyUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// Username, which can contain up to 32 characters. You need to select either `Username` or `UserId` as the search criterion; if both are selected, `Username` will be used by default.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User remarks, which can contain up to 512 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of IDs of the user's user groups.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// User ID. You need to select either `UserName` or `UserId` as the search criterion. If both are selected, `UserName` will be used by default.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User's mobile number.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// User expiration time in ISO 8601 format.
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// User password, which needs to be configured according to the password policy.
	Password *string `json:"Password,omitempty" name:"Password"`

	// User's email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Whether the password needs to be reset. Default value: false (no).
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`

	// Unique ID of the user's primary organization. If this parameter is left empty, the user will be created under the root node by default.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// List of IDs of the user's secondary organizations.
	SecondaryOrgNodeIdList []*string `json:"SecondaryOrgNodeIdList,omitempty" name:"SecondaryOrgNodeIdList"`
}

func (r *ModifyUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "UserGroupIds")
	delete(f, "UserId")
	delete(f, "Phone")
	delete(f, "ExpirationTime")
	delete(f, "Password")
	delete(f, "Email")
	delete(f, "PwdNeedReset")
	delete(f, "OrgNodeId")
	delete(f, "SecondaryOrgNodeIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserInfoResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserInfoResponseParams `json:"Response"`
}

func (r *ModifyUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgNodeChildInfo struct {
	// Displayed organization node name, which can contain up to 64 characters and is the same as the organization name by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Last modification time of the organization node in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// External ID of the organization node, which is optional and customizable.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`

	// Parent node ID of the current organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

	// Organization node ID, which is globally unique.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Data source.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// Organization node creation time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// Organization node description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type OrgNodeChildUserInfo struct {
	// Organization node ID, which is globally unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// User information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// Total number of users under the current organization node.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalUserNum *int64 `json:"TotalUserNum,omitempty" name:"TotalUserNum"`

	// Organization ID path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeIdPath *string `json:"OrgNodeIdPath,omitempty" name:"OrgNodeIdPath"`

	// Organization name path.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgNodeNamePath *string `json:"OrgNodeNamePath,omitempty" name:"OrgNodeNamePath"`
}

// Predefined struct for user
type RemoveAccountFromAccountGroupRequestParams struct {
	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// List of IDs of the accounts to be removed.
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`
}

type RemoveAccountFromAccountGroupRequest struct {
	*tchttp.BaseRequest
	
	// Account group ID
	AccountGroupId *string `json:"AccountGroupId,omitempty" name:"AccountGroupId"`

	// List of IDs of the accounts to be removed.
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`
}

func (r *RemoveAccountFromAccountGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAccountFromAccountGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	delete(f, "AccountIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveAccountFromAccountGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveAccountFromAccountGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveAccountFromAccountGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveAccountFromAccountGroupResponseParams `json:"Response"`
}

func (r *RemoveAccountFromAccountGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAccountFromAccountGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromUserGroupRequestParams struct {
	// List of IDs of the users to be added to the user group.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

type RemoveUserFromUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the users to be added to the user group.
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// User group ID, which is globally unique.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *RemoveUserFromUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserFromUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromUserGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveUserFromUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserFromUserGroupResponseParams `json:"Response"`
}

func (r *RemoveUserFromUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SortCondition struct {
	// Sorting attribute.
	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`

	// Sorting order. Valid values: ASC: ascending order; DESC: descending order.
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type ThirdPartyAccountInfo struct {
	// Third-Party account code. `2` indicates WeCom account.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountCode *string `json:"AccountCode,omitempty" name:"AccountCode"`

	// Username of the account.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

// Predefined struct for user
type UpdateOrgNodeRequestParams struct {
	// Organization node ID, which is globally unique.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Organization node name, which can contain up to 64 characters.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Organization node description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// External ID of the organization node, which is optional and customizable. If this parameter is specified, its uniqueness will be verified.
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`
}

type UpdateOrgNodeRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID, which is globally unique.
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// Organization node name, which can contain up to 64 characters.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Organization node description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// External ID of the organization node, which is optional and customizable. If this parameter is specified, its uniqueness will be verified.
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`
}

func (r *UpdateOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "CustomizedOrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrgNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrgNodeResponseParams `json:"Response"`
}

func (r *UpdateOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserGroupInfo struct {
	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User group ID, which is globally unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// User group remarks.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`
}

type UserGroupInfoSearchCriteria struct {
	// Search by name. The match criteria include user group name and user group ID.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type UserGroupInformation struct {
	// User group ID.
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// User group name.
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`

	// Last update time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`
}

type UserGroupInformationSearchCriteria struct {
	// Search by name. The match criteria include user group name.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type UserInfo struct {
	// User ID, which is globally unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User's mobile number.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// Email address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitempty" name:"Email"`

	// User status.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Data source.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
}

type UserInformation struct {
	// Username, which can contain up to 32 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User status.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Nickname, which can contain up to 64 characters and is the same as the username by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// User remarks, which can contain up to 512 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Last update time of the user in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// User creation time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Path ID of the user's primary organization.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrgPath *string `json:"OrgPath,omitempty" name:"OrgPath"`

	// User's mobile number with country code, such as `+86-00000000000`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// List of IDs of the user's user groups.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubjectGroups []*string `json:"SubjectGroups,omitempty" name:"SubjectGroups"`

	// User's email address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Last login time of the user in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`

	// User ID, which is globally unique and can contain up to 64 characters.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type UserSearchCriteria struct {
	// Username, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// User's mobile number.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// User's email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// User status. Valid values: NORMAL: normal; FREEZE: frozen; LOCKED: locked; NOT_ENABLED: disabled.
	Status *string `json:"Status,omitempty" name:"Status"`

	// User creation time in ISO 8601 format.
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// The user's last update time.
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// Search by name. The match criteria include username and user ID.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}