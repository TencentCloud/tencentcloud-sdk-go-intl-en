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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AddUserRequest struct {
	*tchttp.BaseRequest

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Sub-user remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether or not the sub-user is allowed to log in to the console. 0: No; 1: Yes.
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// Whether or not to generate keys for sub-users. 0: No; 1: Yes.
	UseApi *uint64 `json:"UseApi,omitempty" name:"UseApi"`

	// Sub-user's console login password. If no password rules have been set, the password must have a minimum of 8 characters containing uppercase letters, lowercase letters, digits, and special characters by default. This parameter will be valid only when the sub-user is allowed to log in to the console. If it is not specified and console login is allowed, the system will automatically generate a random 32-character password that contains uppercase letters, lowercase letters, digits, and special characters.
	Password *string `json:"Password,omitempty" name:"Password"`

	// If the sub-user needs to reset their password when they next log in to the console. 0: No; 1: Yes.
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`

	// Mobile number
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Country/Area Code
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Email
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *AddUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Sub-user UIN
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// Sub-user username
		Name *string `json:"Name,omitempty" name:"Name"`

		// If the combination of input parameters indicates that a random password should be generated, the generated password is returned
		Password *string `json:"Password,omitempty" name:"Password"`

		// Sub-user’s key ID
		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

		// Sub-user’s secret key
		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

		// Sub-user UID
		Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest

	// How sub-user UIDs are associated with the ID of the user group they are added to.
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info" list`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserToGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserToGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachEntityOfPolicy struct {

	// Entity ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Entity Name
	// Note: This field may return null, indicating that no valid value was found.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Entity UIN
	// Note: This field may return null, indicating that no valid value was found.
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Type of entity association. 1: Associate by users; 2: Associate by User Groups
	RelatedType *uint64 `json:"RelatedType,omitempty" name:"RelatedType"`
}

type AttachGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// User Group ID
	AttachGroupId *uint64 `json:"AttachGroupId,omitempty" name:"AttachGroupId"`
}

func (r *AttachGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachPolicyInfo struct {

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	// Note: This field may return null, indicating that no valid value was found.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Time created
	// Note: This field may return null, indicating that no valid value was found.
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// How the policy was created: 1: Via console; 2: Via syntax
	// Note: This field may return null, indicating that no valid value was found.
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`

	// Valid values: `user` and `QCS`
	// Note: This field may return null, indicating that no valid value was found.
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 
	OperateOwnerUin *string `json:"OperateOwnerUin,omitempty" name:"OperateOwnerUin"`

	// 
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 
	OperateUinType *uint64 `json:"OperateUinType,omitempty" name:"OperateUinType"`

	// Queries if the policy has been deactivated
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// List of deprecated products
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail" list`
}

type AttachRolePolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID, used to specify a role. Input either `AttachRoleId` or `AttachRoleName`
	AttachRoleId *string `json:"AttachRoleId,omitempty" name:"AttachRoleId"`

	// Role name, used to specify a role. Input either `AttachRoleId` or `AttachRoleName`
	AttachRoleName *string `json:"AttachRoleName,omitempty" name:"AttachRoleName"`

	// Policy name. Either `PolicyId` or `PolicyName` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *AttachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Sub-account UIN
	AttachUin *uint64 `json:"AttachUin,omitempty" name:"AttachUin"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachedPolicyOfRole struct {

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Time of association
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Policy type. `User` indicates custom policy; `QCS` indicates preset policy
	// Note: This field may return null, indicating that no valid value was found.
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// Policy creation method. 1: indicates the policy was created based on product function or item permission; other values indicate the policy was created based on the policy syntax
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`

	// Whether the product has been deprecated (0: no; 1: yes)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// List of deprecated products
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail" list`

	// Policy description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ConsumeCustomMFATokenRequest struct {
	*tchttp.BaseRequest

	// Custom multi-factor verification Token
	MFAToken *string `json:"MFAToken,omitempty" name:"MFAToken"`
}

func (r *ConsumeCustomMFATokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConsumeCustomMFATokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConsumeCustomMFATokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConsumeCustomMFATokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConsumeCustomMFATokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// User Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// User Group description
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// User Group ID
		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest

	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy document, such as `{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}`, where `principal` is used to specify the resources that the role is authorized to access. For more information on this parameter, please see the `RoleInfo` output parameter of the [GetRole](https://cloud.tencent.com/document/product/598/36221) API
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Policy description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of newly added policy
		PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// Role name
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Policy document
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Role description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Whether login is allowed. 1: yes, 0: no
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// The maximum validity period of the temporary key for creating a role (range: 0-43200)
	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Role ID
	// Note: This field may return null, indicating that no valid value was found.
		RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML identity provider description
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML identity provider metadata document (Base64)
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *CreateSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SAML identity provider resource descriptor
		ProviderArn *string `json:"ProviderArn,omitempty" name:"ProviderArn"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest

	// Array. Array elements are policy IDs. Policies can be deleted in a batch
	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId" list`
}

func (r *DeletePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest

	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to forcibly delete the sub-user. The default input parameter is `0`. `0`: do not delete the user if the user has undeleted API keys; `1`: first delete the API keys then delete the user if the user has undeleted API keys. To delete API keys, you need to have cam:DeleteApiKey permission, which enables you to delete both enabled and disabled API keys. If you do not have this permission, you will not be able to delete API keys and the user.
	Force *uint64 `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest

	// Page number, beginning from 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of lines per page, no greater than 200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Role details list
	// Note: This field may return null, indicating that no valid value was found.
		List []*RoleInfo `json:"List,omitempty" name:"List" list`

		// Total number of roles
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// User Group ID
	DetachGroupId *uint64 `json:"DetachGroupId,omitempty" name:"DetachGroupId"`
}

func (r *DetachGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachRolePolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID, used to specify a role. Input either `AttachRoleId` or `AttachRoleName`
	DetachRoleId *string `json:"DetachRoleId,omitempty" name:"DetachRoleId"`

	// Role name, used to specify a role. Input either `AttachRoleId` or `AttachRoleName`
	DetachRoleName *string `json:"DetachRoleName,omitempty" name:"DetachRoleName"`

	// Policy name. Either `PolicyId` or `PolicyName` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *DetachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Sub-account UIN
	DetachUin *uint64 `json:"DetachUin,omitempty" name:"DetachUin"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomMFATokenInfoRequest struct {
	*tchttp.BaseRequest

	// Custom multi-factor verification Token
	MFAToken *string `json:"MFAToken,omitempty" name:"MFAToken"`
}

func (r *GetCustomMFATokenInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomMFATokenInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomMFATokenInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Account ID corresponding to the custom multi-factor verification Token
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomMFATokenInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomMFATokenInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupRequest struct {
	*tchttp.BaseRequest

	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// User Group ID
		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

		// User Group name
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// Number of members in the User Group
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// User Group description
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// Time User Group created
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// User Group member information
		UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *GetPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Policy name
	// Note: This field may return null, indicating that no valid value was found.
		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

		// Policy description
	// Note: This field may return null, indicating that no valid value was found.
		Description *string `json:"Description,omitempty" name:"Description"`

		// 1: Custom policy; 2: Preset policy
	// Note: This field may return null, indicating that no valid value was found.
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// Time created
	// Note: This field may return null, indicating that no valid value was found.
		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

		// Time of latest update
	// Note: This field may return null, indicating that no valid value was found.
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// Policy document
	// Note: This field may return null, indicating that no valid value was found.
		PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRoleRequest struct {
	*tchttp.BaseRequest

	// Role ID, used to specify role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *GetRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Role details
		RoleInfo *RoleInfo `json:"RoleInfo,omitempty" name:"RoleInfo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SAML identity provider name
		Name *string `json:"Name,omitempty" name:"Name"`

		// SAML identity provider description
		Description *string `json:"Description,omitempty" name:"Description"`

		// Time SAML identity provider created
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// Time SAML identity provider last modified
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// SAML identity provider metadata document
		SAMLMetadata *string `json:"SAMLMetadata,omitempty" name:"SAMLMetadata"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserRequest struct {
	*tchttp.BaseRequest

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Sub-user UIN
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// Sub-user username
		Name *string `json:"Name,omitempty" name:"Name"`

		// Sub-user UID
		Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

		// Sub-user remarks
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// If sub-user can log in to the Console
		ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

		// Mobile number
		PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

		// Country/Area code
		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

		// Email
		Email *string `json:"Email,omitempty" name:"Email"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupIdOfUidInfo struct {

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

type GroupInfo struct {

	// User group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// User Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Time User Group created
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// User Group description
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type GroupMemberInfo struct {

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Sub-user UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Sub-user name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Mobile number
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Mobile number country/area code
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// If mobile number has been verified
	PhoneFlag *uint64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// If email has been verified
	EmailFlag *uint64 `json:"EmailFlag,omitempty" name:"EmailFlag"`

	// User type
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`

	// Time policy created
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// If the user is the main message recipient
	IsReceiverOwner *uint64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
}

type ListAttachedGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// User group ID
	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListAttachedGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of policies
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Policy list
		List []*AttachPolicyInfo `json:"List,omitempty" name:"List" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// Page number, beginning from 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of lines per page, no more than 200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Filter according to policy type. `User` indicates querying custom policies only. `QCS` indicates querying preset policies only
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of policies associated with the role
		List []*AttachedPolicyOfRole `json:"List,omitempty" name:"List" list`

		// Total number of policies associated with the role
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserPoliciesRequest struct {
	*tchttp.BaseRequest

	// Sub-account UIN
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListAttachedUserPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedUserPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of policies
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Policy list
		List []*AttachPolicyInfo `json:"List,omitempty" name:"List" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedUserPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedUserPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyRequest struct {
	*tchttp.BaseRequest

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Valid values: `All`, `User`, `Group`, and `Role`. `All` means all entity types will be returned; `User` means only sub-accounts will be returned; `Group` means only User Groups will be returned; `Role` means only roles will be returned. `All` is the default value.
	EntityFilter *string `json:"EntityFilter,omitempty" name:"EntityFilter"`
}

func (r *ListEntitiesForPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of entities
	// Note: This field may return null, indicating that no valid value was found.
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Entity list
	// Note: This field may return null, indicating that no valid value was found.
		List []*AttachEntityOfPolicy `json:"List,omitempty" name:"List" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsForUserRequest struct {
	*tchttp.BaseRequest

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Number of entries per page; default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Page number; default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Sub-account UIN
	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
}

func (r *ListGroupsForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsForUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of User Groups to which the sub-user has been added
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// User Group information
		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupsForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsForUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsRequest struct {
	*tchttp.BaseRequest

	// Page number; default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Filter by User Group name
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of User Groups
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// User group information array
		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest

	// Number of entries per page: must be greater than 0 and no greater than 200. Default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Page number. Starts from 1 and cannot be greater than 200. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Valid values: `All`, `QCS`, and `Local`. `All` means all policies will be returned; `QCS` means only preset policies will be returned; `Local` means only custom policies will be returned. `All` is the default value
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// Filter by policy name
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of policies
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Policy array. Each array contains fields including `policyId`, `policyName`, `addTime`, `type`, `description`, and `createMode`. 
	// policyId: policy ID 
	// policyName: policy name
	// addTime: policy creation time
	// type: 1: custom policy, 2: preset policy 
	// description: policy description 
	// createMode: 1 indicates a policy created based on business permissions, while other values indicate that the policy syntax can be viewed and the policy can be updated using the policy syntax
	// Attachments: number of associated users
	// ServiceType: the product the policy is associated with
	// IsAttached: this value should not be null when querying if a marked entity has been associated with a policy. 0 indicates that no policy has been associated, and 1 indicates that a policy has been associated
		List []*StrategyInfo `json:"List,omitempty" name:"List" list`

		// Reserved field
	// Note: This field may return null, indicating that no valid value was found.
		ServiceTypeList []*string `json:"ServiceTypeList,omitempty" name:"ServiceTypeList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSAMLProvidersRequest struct {
	*tchttp.BaseRequest
}

func (r *ListSAMLProvidersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSAMLProvidersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSAMLProvidersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of SAML identity providers
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of SAML identity providers
		SAMLProviderSet []*SAMLProviderInfo `json:"SAMLProviderSet,omitempty" name:"SAMLProviderSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSAMLProvidersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSAMLProvidersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupRequest struct {
	*tchttp.BaseRequest

	// User group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Page number; default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListUsersForGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersForGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of users associated with the User Group
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Sub-user information
		UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersForGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersForGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersRequest struct {
	*tchttp.BaseRequest
}

func (r *ListUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Sub-user information
		Data []*SubAccountInfo `json:"Data,omitempty" name:"Data" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginActionFlag struct {

	// Mobile phone
	Phone *uint64 `json:"Phone,omitempty" name:"Phone"`

	// Hard token
	Token *uint64 `json:"Token,omitempty" name:"Token"`

	// Soft token
	Stoken *uint64 `json:"Stoken,omitempty" name:"Stoken"`

	// WeChat
	Wechat *uint64 `json:"Wechat,omitempty" name:"Wechat"`

	// Custom
	Custom *uint64 `json:"Custom,omitempty" name:"Custom"`
}

type OffsiteFlag struct {

	// Verification identifier
	VerifyFlag *uint64 `json:"VerifyFlag,omitempty" name:"VerifyFlag"`

	// Phone notification
	NotifyPhone *uint64 `json:"NotifyPhone,omitempty" name:"NotifyPhone"`

	// Email notification
	NotifyEmail *int64 `json:"NotifyEmail,omitempty" name:"NotifyEmail"`

	// WeChat notification
	NotifyWechat *uint64 `json:"NotifyWechat,omitempty" name:"NotifyWechat"`

	// Alert
	Tips *uint64 `json:"Tips,omitempty" name:"Tips"`
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest

	// The UID of the user to be deleted and an array corresponding to the User Group IDs
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info" list`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserFromGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoleInfo struct {

	// Role ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Role policy document
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Role description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Time role created
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Time role last updated
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// If login is allowed for the role
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// User role. Valid values: `user`, `system`, `service_linked`
	// Note: this field may return null, indicating that no valid values can be obtained.
	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`

	// Valid period
	// Note: this field may return null, indicating that no valid values can be obtained.
	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`

	// Task identifier for deleting a service-linked role 
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`
}

type SAMLProviderInfo struct {

	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML identity provider description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Time SAML identity provider created
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Time SAML identity provider last modified
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type SetFlagRequest struct {
	*tchttp.BaseRequest

	// Set user UIN
	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`

	// Login settings
	LoginFlag *LoginActionFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// Sensitive operation settings
	ActionFlag *LoginActionFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`

	// Remote login settings
	OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`

	// Whether or not to reset MFA
	NeedResetMfa *uint64 `json:"NeedResetMfa,omitempty" name:"NeedResetMfa"`
}

func (r *SetFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StrategyInfo struct {

	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Time policy created
	// Note: This field may return null, indicating that no valid value was found.
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Policy type. 1: Custom policy; 2: Preset policy
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Policy description
	// Note: This field may return null, indicating that no valid value was found.
	Description *string `json:"Description,omitempty" name:"Description"`

	// How the policy was created: 1: Via console; 2: Via syntax
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`

	// Number of associated users
	Attachments *uint64 `json:"Attachments,omitempty" name:"Attachments"`

	// Product associated with the policy
	// Note: This field may return null, indicating that no valid value was found.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 
	IsAttached *uint64 `json:"IsAttached,omitempty" name:"IsAttached"`

	// Queries if the policy has been deactivated
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// List of deprecated products
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail" list`

	// The deletion task identifier used to check the deletion status of the service-linked role
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsServiceLinkedPolicy *uint64 `json:"IsServiceLinkedPolicy,omitempty" name:"IsServiceLinkedPolicy"`
}

type SubAccountInfo struct {

	// Sub-user user ID
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Sub-user remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// If sub-user can log in to the console
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// Mobile number
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Country/Area code
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Email
	Email *string `json:"Email,omitempty" name:"Email"`
}

type UpdateAssumeRolePolicyRequest struct {
	*tchttp.BaseRequest

	// Policy document
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateAssumeRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssumeRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAssumeRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssumeRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssumeRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest

	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// User Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// User Group description
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionRequest struct {
	*tchttp.BaseRequest

	// Role description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateRoleDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML identity provider description
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML identity provider metadata document (Base64)
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *UpdateSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Sub-user remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether or not the sub-user is allowed to log in to the console. 0: No; 1: Yes.
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// Sub-user's console login password. If no password rules have been set, the password must have a minimum of 8 characters containing uppercase letters, lowercase letters, digits, and special characters by default. This parameter will be valid only when the sub-user is allowed to log in to the console. If it is not specified and console login is allowed, the system will automatically generate a random 32-character password that contains uppercase letters, lowercase letters, digits, and special characters.
	Password *string `json:"Password,omitempty" name:"Password"`

	// If the sub-user needs to reset their password when they next log in to the console. 0: No; 1: Yes.
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`

	// Mobile number
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Country/Area Code
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Email
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *UpdateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
