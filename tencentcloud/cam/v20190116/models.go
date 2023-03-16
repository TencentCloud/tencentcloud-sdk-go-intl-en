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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccessKey struct {
	// Access key ID
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key status. Valid values: Active (activated), Inactive (not activated)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AccessKeyDetail struct {
	// Access key ID
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Access key, which is visible only when it is created. Keep it properly.
	SecretAccessKey *string `json:"SecretAccessKey,omitempty" name:"SecretAccessKey"`

	// Key status. Valid values: `Active` (activated), `Inactive` (not activated).
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type AddUserRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "ConsoleLogin")
	delete(f, "UseApi")
	delete(f, "Password")
	delete(f, "NeedResetPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserResponseParams struct {
	// Sub-user UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// If the combination of input parameters indicates that a random password should be generated, the generated password is returned
	Password *string `json:"Password,omitempty" name:"Password"`

	// Sub-user's key ID
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// Sub-user's secret key
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddUserResponse struct {
	*tchttp.BaseResponse
	Response *AddUserResponseParams `json:"Response"`
}

func (r *AddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupRequestParams struct {
	// The association between the user group ID and the sub-user UIN/UID.
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info"`
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest
	
	// The association between the user group ID and the sub-user UIN/UID.
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info"`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserToGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddUserToGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddUserToGroupResponseParams `json:"Response"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Policy association time
	// Note: this field may return `null`, indicating that no valid value was found.
	AttachmentTime *string `json:"AttachmentTime,omitempty" name:"AttachmentTime"`
}

// Predefined struct for user
type AttachGroupPolicyRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// User Group ID
	AttachGroupId *uint64 `json:"AttachGroupId,omitempty" name:"AttachGroupId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "AttachGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachGroupPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachGroupPolicyResponseParams `json:"Response"`
}

func (r *AttachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Policy remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Root account of the operator associating the policy
	// Note: this field may return null, indicating that no valid values can be obtained.
	OperateOwnerUin *string `json:"OperateOwnerUin,omitempty" name:"OperateOwnerUin"`

	// The ID of the account associating the policy. If `UinType` is 0, this indicates that this is a sub-account `UIN`. If `UinType` is 1, this indicates this is a role ID
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// If `UinType` is 0, `OperateUin` indicates that this is a sub-account `UIN`. If `UinType` is 1, `OperateUin` indicates that this is a role ID
	OperateUinType *uint64 `json:"OperateUinType,omitempty" name:"OperateUinType"`

	// Queries if the policy has been deactivated
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// List of deprecated products
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail"`
}

// Predefined struct for user
type AttachRolePolicyRequestParams struct {
	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID, used to specify a role. Input either `AttachRoleId` or `AttachRoleName`
	AttachRoleId *string `json:"AttachRoleId,omitempty" name:"AttachRoleId"`

	// Role name, used to specify a role. Input either `AttachRoleId` or `AttachRoleName`
	AttachRoleName *string `json:"AttachRoleName,omitempty" name:"AttachRoleName"`

	// Policy name. Either `PolicyId` or `PolicyName` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "AttachRoleId")
	delete(f, "AttachRoleName")
	delete(f, "PolicyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachRolePolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachRolePolicyResponseParams `json:"Response"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Sub-account UIN
	AttachUin *uint64 `json:"AttachUin,omitempty" name:"AttachUin"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "AttachUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachUserPolicyResponseParams `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail"`

	// Policy description
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type AttachedUserPolicy struct {
	// Policy ID.
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time.
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// Policy type (`1`: custom policy; `2`: preset policy).
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// Creation mode (`1`: create by product feature or project permission; other values: create by policy syntax).
	CreateMode *string `json:"CreateMode,omitempty" name:"CreateMode"`

	// Information on policies inherited from the user group.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Groups []*AttachedUserPolicyGroupInfo `json:"Groups,omitempty" name:"Groups"`

	// Whether the product has been deprecated (`0`: no; `1`: yes).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// List of deprecated products.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail"`
}

type AttachedUserPolicyGroupInfo struct {
	// Group ID.
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Group name.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

// Predefined struct for user
type ConsumeCustomMFATokenRequestParams struct {
	// Custom multi-factor verification Token
	MFAToken *string `json:"MFAToken,omitempty" name:"MFAToken"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConsumeCustomMFATokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MFAToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConsumeCustomMFATokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConsumeCustomMFATokenResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ConsumeCustomMFATokenResponse struct {
	*tchttp.BaseResponse
	Response *ConsumeCustomMFATokenResponseParams `json:"Response"`
}

func (r *ConsumeCustomMFATokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConsumeCustomMFATokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeyRequestParams struct {
	// UIN of the specified user. If this parameter is left empty, the access key will be created for the current user by default.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

type CreateAccessKeyRequest struct {
	*tchttp.BaseRequest
	
	// UIN of the specified user. If this parameter is left empty, the access key will be created for the current user by default.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *CreateAccessKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeyResponseParams struct {
	// Access key
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessKey *AccessKeyDetail `json:"AccessKey,omitempty" name:"AccessKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccessKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessKeyResponseParams `json:"Response"`
}

func (r *CreateAccessKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// User Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// User Group description
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOIDCConfigRequestParams struct {
	// IdP URL.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Public key for signature, which must be Base64-encoded.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID.
	ClientId []*string `json:"ClientId,omitempty" name:"ClientId"`

	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// IdP URL.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Public key for signature, which must be Base64-encoded.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID.
	ClientId []*string `json:"ClientId,omitempty" name:"ClientId"`

	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "IdentityKey")
	delete(f, "ClientId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOIDCConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateOIDCConfigResponseParams `json:"Response"`
}

func (r *CreateOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyRequestParams struct {
	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy document, such as `{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}`, where `principal` is used to specify the resources that the role is authorized to access. For more information on this parameter, please see the `RoleInfo` output parameter of the [GetRole](https://intl.cloud.tencent.com/document/product/598/36221?from_cn_redirect=1) API
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Policy description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy document, such as `{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}`, where `principal` is used to specify the resources that the role is authorized to access. For more information on this parameter, please see the `RoleInfo` output parameter of the [GetRole](https://intl.cloud.tencent.com/document/product/598/36221?from_cn_redirect=1) API
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Policy description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyName")
	delete(f, "PolicyDocument")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyResponseParams struct {
	// ID of newly added policy
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyResponseParams `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyVersionRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The policy document to use as the content for the new version
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Specifies whether to set this version as the default version
	SetAsDefault *bool `json:"SetAsDefault,omitempty" name:"SetAsDefault"`
}

type CreatePolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The policy document to use as the content for the new version
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Specifies whether to set this version as the default version
	SetAsDefault *bool `json:"SetAsDefault,omitempty" name:"SetAsDefault"`
}

func (r *CreatePolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "PolicyDocument")
	delete(f, "SetAsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyVersionResponseParams struct {
	// Policy version ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyVersionResponseParams `json:"Response"`
}

func (r *CreatePolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleRequestParams struct {
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

	// Tags bound to the role.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
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

	// Tags bound to the role.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "PolicyDocument")
	delete(f, "Description")
	delete(f, "ConsoleLogin")
	delete(f, "SessionDuration")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleResponseParams struct {
	// Role ID
	// Note: This field may return null, indicating that no valid value was found.
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleResponseParams `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSAMLProviderRequestParams struct {
	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML identity provider description
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML identity provider metadata document (Base64)
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "SAMLMetadataDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSAMLProviderResponseParams struct {
	// SAML identity provider resource descriptor
	ProviderArn *string `json:"ProviderArn,omitempty" name:"ProviderArn"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *CreateSAMLProviderResponseParams `json:"Response"`
}

func (r *CreateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceLinkedRoleRequestParams struct {
	// Authorized service, i.e., Tencent Cloud service entity with this role attached.
	QCSServiceName []*string `json:"QCSServiceName,omitempty" name:"QCSServiceName"`

	// Custom suffix. A string you provide, which is combined with the service-provided prefix to form the complete role name.
	CustomSuffix *string `json:"CustomSuffix,omitempty" name:"CustomSuffix"`

	// Role description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Tags bound to the role.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
}

type CreateServiceLinkedRoleRequest struct {
	*tchttp.BaseRequest
	
	// Authorized service, i.e., Tencent Cloud service entity with this role attached.
	QCSServiceName []*string `json:"QCSServiceName,omitempty" name:"QCSServiceName"`

	// Custom suffix. A string you provide, which is combined with the service-provided prefix to form the complete role name.
	CustomSuffix *string `json:"CustomSuffix,omitempty" name:"CustomSuffix"`

	// Role description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Tags bound to the role.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateServiceLinkedRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceLinkedRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QCSServiceName")
	delete(f, "CustomSuffix")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceLinkedRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceLinkedRoleResponseParams struct {
	// Role ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServiceLinkedRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceLinkedRoleResponseParams `json:"Response"`
}

func (r *CreateServiceLinkedRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceLinkedRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserOIDCConfigRequestParams struct {
	// OpenID Connect IdP URL.
	// It corresponds to the value of the `issuer` field in the `Openid-configuration` provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Signature public key, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the `authorization_endpoint` field in the `Openid-configuration` provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always `id_token`.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the `id_token` of the IdP is mapped to the username of a sub-user. It is usually the `sub` or `name` field
	MappingFiled *string `json:"MappingFiled,omitempty" name:"MappingFiled"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitempty" name:"Scope"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// OpenID Connect IdP URL.
	// It corresponds to the value of the `issuer` field in the `Openid-configuration` provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Signature public key, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the `authorization_endpoint` field in the `Openid-configuration` provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always `id_token`.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the `id_token` of the IdP is mapped to the username of a sub-user. It is usually the `sub` or `name` field
	MappingFiled *string `json:"MappingFiled,omitempty" name:"MappingFiled"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitempty" name:"Scope"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "IdentityKey")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserOIDCConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserOIDCConfigResponseParams `json:"Response"`
}

func (r *CreateUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSAMLConfigRequestParams struct {
	// SAML metadata document, which must be Base64 encoded.
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

type CreateUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
	
	// SAML metadata document, which must be Base64 encoded.
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *CreateUserSAMLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSAMLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SAMLMetadataDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserSAMLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSAMLConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserSAMLConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserSAMLConfigResponseParams `json:"Response"`
}

func (r *CreateUserSAMLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessKeyRequestParams struct {
	// ID of the specified access key that needs to be deleted
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// UIN of the specified user. If this parameter is left empty, the access key will be deleted for the current user by default.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

type DeleteAccessKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the specified access key that needs to be deleted
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// UIN of the specified user. If this parameter is left empty, the access key will be deleted for the current user by default.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DeleteAccessKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccessKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessKeyResponseParams `json:"Response"`
}

func (r *DeleteAccessKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOIDCConfigRequestParams struct {
	// OIDC IdP name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DeleteOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// OIDC IdP name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOIDCConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOIDCConfigResponseParams `json:"Response"`
}

func (r *DeleteOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyRequestParams struct {
	// Array. Array elements are policy IDs. Policies can be deleted in a batch
	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest
	
	// Array. Array elements are policy IDs. Policies can be deleted in a batch
	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyResponseParams `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyVersionRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy version ID
	VersionId []*uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

type DeletePolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy version ID
	VersionId []*uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DeletePolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyVersionResponseParams `json:"Response"`
}

func (r *DeletePolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRolePermissionsBoundaryRequestParams struct {
	// Role ID (either it or the role name must be entered)
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name (either it or the role ID must be entered)
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type DeleteRolePermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// Role ID (either it or the role name must be entered)
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name (either it or the role ID must be entered)
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRolePermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolePermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRolePermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRolePermissionsBoundaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRolePermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRolePermissionsBoundaryResponseParams `json:"Response"`
}

func (r *DeleteRolePermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolePermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleRequestParams struct {
	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleResponseParams `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSAMLProviderRequestParams struct {
	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSAMLProviderResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSAMLProviderResponseParams `json:"Response"`
}

func (r *DeleteSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceLinkedRoleRequestParams struct {
	// Name of the service-linked role to be deleted.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type DeleteServiceLinkedRoleRequest struct {
	*tchttp.BaseRequest
	
	// Name of the service-linked role to be deleted.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteServiceLinkedRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceLinkedRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceLinkedRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceLinkedRoleResponseParams struct {
	// Deletion task identifier, which can be used to check the status of a service-linked role deletion.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServiceLinkedRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceLinkedRoleResponseParams `json:"Response"`
}

func (r *DeleteServiceLinkedRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceLinkedRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserPermissionsBoundaryRequestParams struct {
	// Sub-account `Uin`
	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

type DeleteUserPermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// Sub-account `Uin`
	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DeleteUserPermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserPermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserPermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserPermissionsBoundaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUserPermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserPermissionsBoundaryResponseParams `json:"Response"`
}

func (r *DeleteUserPermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserPermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to forcibly delete the sub-user. The default input parameter is `0`. `0`: do not delete the user if the user has undeleted API keys; `1`: first delete the API keys then delete the user if the user has undeleted API keys. To delete API keys, you need to have cam:DeleteApiKey permission, which enables you to delete both enabled and disabled API keys. If you do not have this permission, you will not be able to delete API keys and the user.
	Force *uint64 `json:"Force,omitempty" name:"Force"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Force")
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
type DescribeOIDCConfigRequestParams struct {
	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOIDCConfigResponseParams struct {
	// IdP type. 11: Role IdP.
	ProviderType *uint64 `json:"ProviderType,omitempty" name:"ProviderType"`

	// IdP URL.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Public key for signature.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID.
	ClientId []*string `json:"ClientId,omitempty" name:"ClientId"`

	// Status. 0: Not set; 2: Disabled; 11: Enabled.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOIDCConfigResponseParams `json:"Response"`
}

func (r *DescribeOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListRequestParams struct {
	// Page number, beginning from 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of lines per page, no greater than 200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// A parameter used to filter the list of roles under a tag.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest
	
	// Page number, beginning from 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of lines per page, no greater than 200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// A parameter used to filter the list of roles under a tag.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListResponseParams struct {
	// Role details list
	// Note: This field may return null, indicating that no valid value was found.
	List []*RoleInfo `json:"List,omitempty" name:"List"`

	// Total number of roles
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoleListResponseParams `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagCollRequestParams struct {
	// Sub-account
	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
}

type DescribeSafeAuthFlagCollRequest struct {
	*tchttp.BaseRequest
	
	// Sub-account
	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
}

func (r *DescribeSafeAuthFlagCollRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagCollRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeAuthFlagCollRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagCollResponseParams struct {
	// Login protection settings
	LoginFlag *LoginActionFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// Sensitive operation protection settings
	ActionFlag *LoginActionFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`

	// Suspicious login location protection settings
	OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSafeAuthFlagCollResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeAuthFlagCollResponseParams `json:"Response"`
}

func (r *DescribeSafeAuthFlagCollResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagCollResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagIntlRequestParams struct {

}

type DescribeSafeAuthFlagIntlRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSafeAuthFlagIntlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagIntlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeAuthFlagIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagIntlResponseParams struct {
	// Login protection settings
	LoginFlag *LoginActionFlagIntl `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// Sensitive operation protection settings
	ActionFlag *LoginActionFlagIntl `json:"ActionFlag,omitempty" name:"ActionFlag"`

	// Suspicious login location protection settings
	OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSafeAuthFlagIntlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeAuthFlagIntlResponseParams `json:"Response"`
}

func (r *DescribeSafeAuthFlagIntlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagIntlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAccountsRequestParams struct {
	// List of sub-user UINs. Up to 50 UINs are supported.
	FilterSubAccountUin []*uint64 `json:"FilterSubAccountUin,omitempty" name:"FilterSubAccountUin"`
}

type DescribeSubAccountsRequest struct {
	*tchttp.BaseRequest
	
	// List of sub-user UINs. Up to 50 UINs are supported.
	FilterSubAccountUin []*uint64 `json:"FilterSubAccountUin,omitempty" name:"FilterSubAccountUin"`
}

func (r *DescribeSubAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAccountsResponseParams struct {
	// Sub-user list
	SubAccounts []*SubAccountUser `json:"SubAccounts,omitempty" name:"SubAccounts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubAccountsResponseParams `json:"Response"`
}

func (r *DescribeSubAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserOIDCConfigRequestParams struct {

}

type DescribeUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserOIDCConfigResponseParams struct {
	// IdP type. 12: user OIDC IdP
	ProviderType *uint64 `json:"ProviderType,omitempty" name:"ProviderType"`

	// IdP URL
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Signature public key
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// Status. 0: not set; 2: disabled; 11: enabled.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Authorization endpoint
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`

	// Authorization scope
	Scope []*string `json:"Scope,omitempty" name:"Scope"`

	// Authorization response type
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Authorization response mode
	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`

	// Mapping field name
	MappingFiled *string `json:"MappingFiled,omitempty" name:"MappingFiled"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserOIDCConfigResponseParams `json:"Response"`
}

func (r *DescribeUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSAMLConfigRequestParams struct {

}

type DescribeUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserSAMLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSAMLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSAMLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSAMLConfigResponseParams struct {
	// SAML metadata document.
	SAMLMetadata *string `json:"SAMLMetadata,omitempty" name:"SAMLMetadata"`

	// Status. `0`: not set, `1`: enabled, `2`: disabled.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserSAMLConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSAMLConfigResponseParams `json:"Response"`
}

func (r *DescribeUserSAMLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachGroupPolicyRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// User Group ID
	DetachGroupId *uint64 `json:"DetachGroupId,omitempty" name:"DetachGroupId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "DetachGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachGroupPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachGroupPolicyResponseParams `json:"Response"`
}

func (r *DetachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachRolePolicyRequestParams struct {
	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID, which is used to specify a role. The input parameter is either `DetachRoleId` or `DetachRoleName`.
	DetachRoleId *string `json:"DetachRoleId,omitempty" name:"DetachRoleId"`

	// Role name, which is used to specify a role. The input parameter is either `DetachRoleId` or `DetachRoleName`.
	DetachRoleName *string `json:"DetachRoleName,omitempty" name:"DetachRoleName"`

	// Policy name. Either `PolicyId` or `PolicyName` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type DetachRolePolicyRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID, which is used to specify a role. The input parameter is either `DetachRoleId` or `DetachRoleName`.
	DetachRoleId *string `json:"DetachRoleId,omitempty" name:"DetachRoleId"`

	// Role name, which is used to specify a role. The input parameter is either `DetachRoleId` or `DetachRoleName`.
	DetachRoleName *string `json:"DetachRoleName,omitempty" name:"DetachRoleName"`

	// Policy name. Either `PolicyId` or `PolicyName` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *DetachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "DetachRoleId")
	delete(f, "DetachRoleName")
	delete(f, "PolicyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachRolePolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachRolePolicyResponseParams `json:"Response"`
}

func (r *DetachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Sub-account UIN
	DetachUin *uint64 `json:"DetachUin,omitempty" name:"DetachUin"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "DetachUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachUserPolicyResponseParams `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserSSORequestParams struct {

}

type DisableUserSSORequest struct {
	*tchttp.BaseRequest
	
}

func (r *DisableUserSSORequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserSSORequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableUserSSORequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserSSOResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableUserSSOResponse struct {
	*tchttp.BaseResponse
	Response *DisableUserSSOResponseParams `json:"Response"`
}

func (r *DisableUserSSOResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserSSOResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountSummaryRequestParams struct {

}

type GetAccountSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAccountSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAccountSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountSummaryResponseParams struct {
	// Number of policies
	Policies *uint64 `json:"Policies,omitempty" name:"Policies"`

	// Number of roles
	Roles *uint64 `json:"Roles,omitempty" name:"Roles"`

	// Number of identity providers
	Idps *uint64 `json:"Idps,omitempty" name:"Idps"`

	// Number of sub-accounts
	User *uint64 `json:"User,omitempty" name:"User"`

	// Number of groups
	Group *uint64 `json:"Group,omitempty" name:"Group"`

	// Total number of grouped users
	Member *uint64 `json:"Member,omitempty" name:"Member"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAccountSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetAccountSummaryResponseParams `json:"Response"`
}

func (r *GetAccountSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomMFATokenInfoRequestParams struct {
	// Custom multi-factor verification Token
	MFAToken *string `json:"MFAToken,omitempty" name:"MFAToken"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomMFATokenInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MFAToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCustomMFATokenInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomMFATokenInfoResponseParams struct {
	// Account ID corresponding to the custom multi-factor verification Token
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCustomMFATokenInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetCustomMFATokenInfoResponseParams `json:"Response"`
}

func (r *GetCustomMFATokenInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomMFATokenInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupRequestParams struct {
	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupResponseParams struct {
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
	UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetGroupResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupResponseParams `json:"Response"`
}

func (r *GetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyResponseParams struct {
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

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	PresetAlias *string `json:"PresetAlias,omitempty" name:"PresetAlias"`

	// Whether it is a service-linked policy
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsServiceLinkedRolePolicy *uint64 `json:"IsServiceLinkedRolePolicy,omitempty" name:"IsServiceLinkedRolePolicy"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPolicyResponse struct {
	*tchttp.BaseResponse
	Response *GetPolicyResponseParams `json:"Response"`
}

func (r *GetPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyVersionRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy version, which can be obtained through `ListPolicyVersions`.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

type GetPolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy version, which can be obtained through `ListPolicyVersions`.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *GetPolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyVersionResponseParams struct {
	// Policy version details
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyVersion *PolicyVersionDetail `json:"PolicyVersion,omitempty" name:"PolicyVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *GetPolicyVersionResponseParams `json:"Response"`
}

func (r *GetPolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleRequestParams struct {
	// Role ID, used to specify role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleResponseParams struct {
	// Role details
	RoleInfo *RoleInfo `json:"RoleInfo,omitempty" name:"RoleInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRoleResponse struct {
	*tchttp.BaseResponse
	Response *GetRoleResponseParams `json:"Response"`
}

func (r *GetRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSAMLProviderRequestParams struct {
	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSAMLProviderResponseParams struct {
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
}

type GetSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *GetSAMLProviderResponseParams `json:"Response"`
}

func (r *GetSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSecurityLastUsedRequestParams struct {
	// Key ID list query. Up to 10 key IDs can be queried.
	SecretIdList []*string `json:"SecretIdList,omitempty" name:"SecretIdList"`
}

type GetSecurityLastUsedRequest struct {
	*tchttp.BaseRequest
	
	// Key ID list query. Up to 10 key IDs can be queried.
	SecretIdList []*string `json:"SecretIdList,omitempty" name:"SecretIdList"`
}

func (r *GetSecurityLastUsedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecurityLastUsedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSecurityLastUsedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSecurityLastUsedResponseParams struct {
	// List of key ID’s recent usage records.
	SecretIdLastUsedRows []*SecretIdLastUsed `json:"SecretIdLastUsedRows,omitempty" name:"SecretIdLastUsedRows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetSecurityLastUsedResponse struct {
	*tchttp.BaseResponse
	Response *GetSecurityLastUsedResponseParams `json:"Response"`
}

func (r *GetSecurityLastUsedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecurityLastUsedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceLinkedRoleDeletionStatusRequestParams struct {
	// Deletion task ID
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`
}

type GetServiceLinkedRoleDeletionStatusRequest struct {
	*tchttp.BaseRequest
	
	// Deletion task ID
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`
}

func (r *GetServiceLinkedRoleDeletionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceLinkedRoleDeletionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeletionTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetServiceLinkedRoleDeletionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceLinkedRoleDeletionStatusResponseParams struct {
	// Status: NOT_STARTED, IN_PROGRESS, SUCCEEDED, FAILED
	Status *string `json:"Status,omitempty" name:"Status"`

	// Reasons why the deletion failed.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// Service type
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// Service name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetServiceLinkedRoleDeletionStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetServiceLinkedRoleDeletionStatusResponseParams `json:"Response"`
}

func (r *GetServiceLinkedRoleDeletionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceLinkedRoleDeletionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserAppIdRequestParams struct {

}

type GetUserAppIdRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetUserAppIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserAppIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserAppIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserAppIdResponseParams struct {
	// UIN of the current account.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// OwnerUin of the current account.
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// AppId of the current account.
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUserAppIdResponse struct {
	*tchttp.BaseResponse
	Response *GetUserAppIdResponseParams `json:"Response"`
}

func (r *GetUserAppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserRequestParams struct {
	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserResponseParams struct {
	// Sub-user UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Sub-user username
	Name *string `json:"Name,omitempty" name:"Name"`

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Sub-user remarks
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether the sub-user can log in to the console. `0`: No; `1`: Yes.
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// Mobile number
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Country/Area code
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Email
	Email *string `json:"Email,omitempty" name:"Email"`

	// Last login IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecentlyLoginIP *string `json:"RecentlyLoginIP,omitempty" name:"RecentlyLoginIP"`

	// Last login time
	// Note: This field may return null, indicating that no valid values can be obtained.
	RecentlyLoginTime *string `json:"RecentlyLoginTime,omitempty" name:"RecentlyLoginTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUserResponse struct {
	*tchttp.BaseResponse
	Response *GetUserResponseParams `json:"Response"`
}

func (r *GetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupIdOfUidInfo struct {
	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Sub-user UIN. For UIN and UID, at least one of them is required.
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
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

	// Whether the mobile phone has been verified. `0`: No; `1`: Yes.
	PhoneFlag *uint64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// Whether the email has been verified. `0`: No; `1`: Yes.
	EmailFlag *uint64 `json:"EmailFlag,omitempty" name:"EmailFlag"`

	// User type. `1`: Global collaborator; `2`: Project collaborator; `3`: Message recipient.
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`

	// Time policy created
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether the user is the primary message recipient. `0`: No; `1`: Yes.
	IsReceiverOwner *uint64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
}

// Predefined struct for user
type ListAccessKeysRequestParams struct {
	// `UIN` of the specified user. If this parameter is left empty, access keys of the current user will be listed by default
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

type ListAccessKeysRequest struct {
	*tchttp.BaseRequest
	
	// `UIN` of the specified user. If this parameter is left empty, access keys of the current user will be listed by default
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *ListAccessKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccessKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAccessKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccessKeysResponseParams struct {
	// Access key list
	// Note: this field may return null, indicating that no valid values can be obtained.
	AccessKeys []*AccessKey `json:"AccessKeys,omitempty" name:"AccessKeys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAccessKeysResponse struct {
	*tchttp.BaseResponse
	Response *ListAccessKeysResponseParams `json:"Response"`
}

func (r *ListAccessKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccessKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedGroupPoliciesRequestParams struct {
	// User group ID
	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Search by keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type ListAttachedGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// User group ID
	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Search by keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListAttachedGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedGroupPoliciesResponseParams struct {
	// Total number of policies
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Policy list
	List []*AttachPolicyInfo `json:"List,omitempty" name:"List"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAttachedGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedGroupPoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedRolePoliciesRequestParams struct {
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

	// Search by keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
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

	// Search by keyword
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "RoleId")
	delete(f, "RoleName")
	delete(f, "PolicyType")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedRolePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedRolePoliciesResponseParams struct {
	// List of policies associated with the role
	List []*AttachedPolicyOfRole `json:"List,omitempty" name:"List"`

	// Total number of policies associated with the role
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAttachedRolePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedRolePoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserAllPoliciesRequestParams struct {
	// Target user ID.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// The number of policies displayed on each page. Value range: 1-200.
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Page number. Value range: 1-200.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// `0`: return policies that are directly associated and inherited from the user group; `1`: return policies that are directly associated; `2`: return policies inherited from the user group.
	AttachType *uint64 `json:"AttachType,omitempty" name:"AttachType"`

	// Policy type.
	StrategyType *uint64 `json:"StrategyType,omitempty" name:"StrategyType"`

	// Keyword for searching.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type ListAttachedUserAllPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Target user ID.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// The number of policies displayed on each page. Value range: 1-200.
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Page number. Value range: 1-200.
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// `0`: return policies that are directly associated and inherited from the user group; `1`: return policies that are directly associated; `2`: return policies inherited from the user group.
	AttachType *uint64 `json:"AttachType,omitempty" name:"AttachType"`

	// Policy type.
	StrategyType *uint64 `json:"StrategyType,omitempty" name:"StrategyType"`

	// Keyword for searching.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListAttachedUserAllPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserAllPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "AttachType")
	delete(f, "StrategyType")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedUserAllPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserAllPoliciesResponseParams struct {
	// Policy list.
	PolicyList []*AttachedUserPolicy `json:"PolicyList,omitempty" name:"PolicyList"`

	// Total number of policies.
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAttachedUserAllPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedUserAllPoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedUserAllPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserAllPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserPoliciesRequestParams struct {
	// Sub-account UIN
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "Page")
	delete(f, "Rp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedUserPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserPoliciesResponseParams struct {
	// Total number of policies
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Policy list
	List []*AttachPolicyInfo `json:"List,omitempty" name:"List"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAttachedUserPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedUserPoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedUserPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCollaboratorsRequestParams struct {
	// Number of entries per page. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type ListCollaboratorsRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page. Default value: 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination start value. Default value: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCollaboratorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCollaboratorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCollaboratorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCollaboratorsResponseParams struct {
	// Total number
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Collaborator information
	Data []*SubAccountInfo `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListCollaboratorsResponse struct {
	*tchttp.BaseResponse
	Response *ListCollaboratorsResponseParams `json:"Response"`
}

func (r *ListCollaboratorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCollaboratorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEntitiesForPolicyRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Page number, which starts from 1. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; 20 by default
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Valid values: `All`, `User`, `Group`, and `Role`. `All` means all entity types will be returned; `User` means only sub-accounts will be returned; `Group` means only User Groups will be returned; `Role` means only roles will be returned. `All` is the default value.
	EntityFilter *string `json:"EntityFilter,omitempty" name:"EntityFilter"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "EntityFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEntitiesForPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEntitiesForPolicyResponseParams struct {
	// Number of entities
	// Note: This field may return null, indicating that no valid value was found.
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Entity list
	// Note: This field may return null, indicating that no valid value was found.
	List []*AttachEntityOfPolicy `json:"List,omitempty" name:"List"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListEntitiesForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListEntitiesForPolicyResponseParams `json:"Response"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsForUserRequestParams struct {
	// Sub-user UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Number of entries per page; default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Page number; default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Sub-account UIN
	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uid")
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGroupsForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsForUserResponseParams struct {
	// Total number of User Groups to which the sub-user has been added
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// User Group information
	GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListGroupsForUserResponse struct {
	*tchttp.BaseResponse
	Response *ListGroupsForUserResponseParams `json:"Response"`
}

func (r *ListGroupsForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsRequestParams struct {
	// Page number; default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Filter by User Group name
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsResponseParams struct {
	// Total number of User Groups
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// User group information array
	GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListGroupsResponseParams `json:"Response"`
}

func (r *ListGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesRequestParams struct {
	// Number of entries per page: must be greater than 0 and no greater than 200. Default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// Page number. Starts from 1 and cannot be greater than 200. Default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Valid values: `All`, `QCS`, and `Local`. `All` means all policies will be returned; `QCS` means only preset policies will be returned; `Local` means only custom policies will be returned. `All` is the default value
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// Filter by policy name
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "Scope")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesResponseParams struct {
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
	List []*StrategyInfo `json:"List,omitempty" name:"List"`

	// Reserved field
	// Note: This field may return null, indicating that no valid value was found.
	ServiceTypeList []*string `json:"ServiceTypeList,omitempty" name:"ServiceTypeList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListPoliciesResponseParams `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPolicyVersionsRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type ListPolicyVersionsRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *ListPolicyVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPolicyVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPolicyVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPolicyVersionsResponseParams struct {
	// Policy version list
	// Note: this field may return null, indicating that no valid values can be obtained.
	Versions []*PolicyVersionItem `json:"Versions,omitempty" name:"Versions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListPolicyVersionsResponse struct {
	*tchttp.BaseResponse
	Response *ListPolicyVersionsResponseParams `json:"Response"`
}

func (r *ListPolicyVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPolicyVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSAMLProvidersRequestParams struct {

}

type ListSAMLProvidersRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListSAMLProvidersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSAMLProvidersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSAMLProvidersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSAMLProvidersResponseParams struct {
	// Total number of SAML identity providers
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of SAML identity providers
	SAMLProviderSet []*SAMLProviderInfo `json:"SAMLProviderSet,omitempty" name:"SAMLProviderSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListSAMLProvidersResponse struct {
	*tchttp.BaseResponse
	Response *ListSAMLProvidersResponseParams `json:"Response"`
}

func (r *ListSAMLProvidersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSAMLProvidersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersForGroupRequestParams struct {
	// User group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Page number; default is 1
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of entries per page; default is 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersForGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Page")
	delete(f, "Rp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersForGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersForGroupResponseParams struct {
	// Total number of users associated with the User Group
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Sub-user information
	UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUsersForGroupResponse struct {
	*tchttp.BaseResponse
	Response *ListUsersForGroupResponseParams `json:"Response"`
}

func (r *ListUsersForGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersForGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersRequestParams struct {

}

type ListUsersRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersResponseParams struct {
	// Sub-user information
	Data []*SubAccountInfo `json:"Data,omitempty" name:"Data"`

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

type LoginActionFlag struct {
	// Phone
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

type LoginActionFlagIntl struct {
	// Mobile number
	Phone *uint64 `json:"Phone,omitempty" name:"Phone"`

	// Hard token
	Token *uint64 `json:"Token,omitempty" name:"Token"`

	// Soft token
	Stoken *uint64 `json:"Stoken,omitempty" name:"Stoken"`

	// WeChat
	Wechat *uint64 `json:"Wechat,omitempty" name:"Wechat"`

	// Custom
	Custom *uint64 `json:"Custom,omitempty" name:"Custom"`

	// Email
	Mail *uint64 `json:"Mail,omitempty" name:"Mail"`
}

type LoginActionMfaFlag struct {
	// Mobile phone
	Phone *uint64 `json:"Phone,omitempty" name:"Phone"`

	// Soft token
	Stoken *uint64 `json:"Stoken,omitempty" name:"Stoken"`

	// WeChat
	Wechat *uint64 `json:"Wechat,omitempty" name:"Wechat"`
}

type OffsiteFlag struct {
	// Verification flag
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

type PolicyVersionDetail struct {
	// Policy version ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`

	// Policy version creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Whether it is the operative version. 0: no, 1: yes
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDefaultVersion *uint64 `json:"IsDefaultVersion,omitempty" name:"IsDefaultVersion"`

	// Policy syntax text
	// Note: this field may return null, indicating that no valid values can be obtained.
	Document *string `json:"Document,omitempty" name:"Document"`
}

type PolicyVersionItem struct {
	// Policy version ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`

	// Policy version creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// Whether it is the operative version. 0: no, 1: yes
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDefaultVersion *int64 `json:"IsDefaultVersion,omitempty" name:"IsDefaultVersion"`
}

// Predefined struct for user
type PutRolePermissionsBoundaryRequestParams struct {
	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID (either it or the role name must be entered)
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name (either it or the role ID must be entered)
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type PutRolePermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Role ID (either it or the role name must be entered)
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name (either it or the role ID must be entered)
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *PutRolePermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutRolePermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutRolePermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutRolePermissionsBoundaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PutRolePermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *PutRolePermissionsBoundaryResponseParams `json:"Response"`
}

func (r *PutRolePermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutRolePermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutUserPermissionsBoundaryRequestParams struct {
	// Sub-account `Uin`
	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type PutUserPermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// Sub-account `Uin`
	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// Policy ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *PutUserPermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutUserPermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutUserPermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutUserPermissionsBoundaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PutUserPermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *PutUserPermissionsBoundaryResponseParams `json:"Response"`
}

func (r *PutUserPermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutUserPermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromGroupRequestParams struct {
	// The user’s UIN/UID to be deleted and the array corresponding to the user group ID.
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info"`
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest
	
	// The user’s UIN/UID to be deleted and the array corresponding to the user group ID.
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info"`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserFromGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserFromGroupResponseParams `json:"Response"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Tags.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`
}

type RoleTags struct {
	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
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

type SecretIdLastUsed struct {
	// Key ID.
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// The date when the key ID was last used (the value is obtained one day later).
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	LastUsedDate *string `json:"LastUsedDate,omitempty" name:"LastUsedDate"`

	// The most recent date the key was accessed
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastSecretUsedDate *uint64 `json:"LastSecretUsedDate,omitempty" name:"LastSecretUsedDate"`
}

// Predefined struct for user
type SetDefaultPolicyVersionRequestParams struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy version, which can be obtained through `ListPolicyVersions`.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

type SetDefaultPolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy version, which can be obtained through `ListPolicyVersions`.
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *SetDefaultPolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultPolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDefaultPolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDefaultPolicyVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetDefaultPolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *SetDefaultPolicyVersionResponseParams `json:"Response"`
}

func (r *SetDefaultPolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultPolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMfaFlagRequestParams struct {
	// Sets user UIN
	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`

	// Sets login protection
	LoginFlag *LoginActionMfaFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// Sets operation protection
	ActionFlag *LoginActionMfaFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
}

type SetMfaFlagRequest struct {
	*tchttp.BaseRequest
	
	// Sets user UIN
	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`

	// Sets login protection
	LoginFlag *LoginActionMfaFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// Sets operation protection
	ActionFlag *LoginActionMfaFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
}

func (r *SetMfaFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMfaFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpUin")
	delete(f, "LoginFlag")
	delete(f, "ActionFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetMfaFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMfaFlagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetMfaFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetMfaFlagResponseParams `json:"Response"`
}

func (r *SetMfaFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMfaFlagResponse) FromJsonString(s string) error {
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

	// This value should not be null when querying whether a marked entity has been associated with a policy. 0 indicates that no policy has been associated, while 1 indicates that a policy has been associated
	IsAttached *uint64 `json:"IsAttached,omitempty" name:"IsAttached"`

	// Queries if the policy has been deactivated
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// List of deprecated products
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail"`

	// The deletion task identifier used to check the deletion status of the service-linked role
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsServiceLinkedPolicy *uint64 `json:"IsServiceLinkedPolicy,omitempty" name:"IsServiceLinkedPolicy"`

	// The number of entities associated with the policy.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AttachEntityCount *int64 `json:"AttachEntityCount,omitempty" name:"AttachEntityCount"`

	// The number of entities associated with the permission boundary.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AttachEntityBoundaryCount *int64 `json:"AttachEntityBoundaryCount,omitempty" name:"AttachEntityBoundaryCount"`

	// The last edited time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
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

	// Creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Nickname.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitempty" name:"NickName"`
}

type SubAccountUser struct {
	// Sub-user ID
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Sub-user name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Sub-user UID. UID is the unique identifier of a user who is a message recipient, while UIN is a unique identifier of a user.
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Sub-user remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// User type (1: root account; 2: sub-user; 3: WeCom sub-user; 4: collaborator; 5: message recipient)
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`


	LastLoginIp *string `json:"LastLoginIp,omitempty" name:"LastLoginIp"`


	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
}

// Predefined struct for user
type TagRoleRequestParams struct {
	// Tag.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`

	// Role name. Specify either the role name or role ID.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Role ID. Specify either the role ID or role name.
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

type TagRoleRequest struct {
	*tchttp.BaseRequest
	
	// Tag.
	Tags []*RoleTags `json:"Tags,omitempty" name:"Tags"`

	// Role name. Specify either the role name or role ID.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Role ID. Specify either the role ID or role name.
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

func (r *TagRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	delete(f, "RoleName")
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TagRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TagRoleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TagRoleResponse struct {
	*tchttp.BaseResponse
	Response *TagRoleResponseParams `json:"Response"`
}

func (r *TagRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntagRoleRequestParams struct {
	// Tag key.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Role name. Specify either the role name or role ID.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Role ID. Specify either the role ID or role name.
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

type UntagRoleRequest struct {
	*tchttp.BaseRequest
	
	// Tag key.
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// Role name. Specify either the role name or role ID.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Role ID. Specify either the role ID or role name.
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`
}

func (r *UntagRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntagRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "RoleName")
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntagRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntagRoleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UntagRoleResponse struct {
	*tchttp.BaseResponse
	Response *UntagRoleResponseParams `json:"Response"`
}

func (r *UntagRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntagRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyRequestParams struct {
	// ID of the specified access key that needs to be updated
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key status. Valid values: `Active` (activated), `Inactive` (not activated).
	Status *string `json:"Status,omitempty" name:"Status"`

	// UIN of the specified user. If this parameter is left empty, the access key will be updated for the current user by default.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

type UpdateAccessKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID of the specified access key that needs to be updated
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// Key status. Valid values: `Active` (activated), `Inactive` (not activated).
	Status *string `json:"Status,omitempty" name:"Status"`

	// UIN of the specified user. If this parameter is left empty, the access key will be updated for the current user by default.
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *UpdateAccessKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	delete(f, "Status")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAccessKeyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccessKeyResponseParams `json:"Response"`
}

func (r *UpdateAccessKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssumeRolePolicyRequestParams struct {
	// Policy document
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssumeRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyDocument")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAssumeRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssumeRolePolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAssumeRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAssumeRolePolicyResponseParams `json:"Response"`
}

func (r *UpdateAssumeRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssumeRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupRequestParams struct {
	// User Group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// User Group name
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// User Group description
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGroupResponseParams `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOIDCConfigRequestParams struct {
	// IdP URL.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Public key for signature, which must be Base64-encoded.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID.
	ClientId []*string `json:"ClientId,omitempty" name:"ClientId"`

	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// IdP URL.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Public key for signature, which must be Base64-encoded.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID.
	ClientId []*string `json:"ClientId,omitempty" name:"ClientId"`

	// Name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "IdentityKey")
	delete(f, "ClientId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOIDCConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOIDCConfigResponseParams `json:"Response"`
}

func (r *UpdateOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyRequestParams struct {
	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name. Either `PolicyName` or `PolicyId` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Policy documentation, for example: `{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}`, where `principal` is used to specify the service that is authorized to use the role. For more information about this parameter, see **RoleInfo** under **Output Parameters** in the [GetRole](https://intl.cloud.tencent.com/document/product/598/36221?from_cn_redirect=1).
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Preset policy remark
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// Policy ID. Either `PolicyId` or `PolicyName` must be entered
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name. Either `PolicyName` or `PolicyId` must be entered
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Policy description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Policy documentation, for example: `{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}`, where `principal` is used to specify the service that is authorized to use the role. For more information about this parameter, see **RoleInfo** under **Output Parameters** in the [GetRole](https://intl.cloud.tencent.com/document/product/598/36221?from_cn_redirect=1).
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// Preset policy remark
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "PolicyName")
	delete(f, "Description")
	delete(f, "PolicyDocument")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyResponseParams struct {
	// Policy ID, which will be returned only if the input parameter is `PolicyName`
	// Note: this field may return null, indicating that no valid values can be obtained.
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePolicyResponseParams `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConsoleLoginRequestParams struct {
	// Whether login is allowed. 1: yes, 0: no
	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// Role ID. Use either `RoleId` or `RoleName` as the input parameter.
	RoleId *int64 `json:"RoleId,omitempty" name:"RoleId"`

	// Role name. Use either `RoleId` or `RoleName` as the input parameter.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type UpdateRoleConsoleLoginRequest struct {
	*tchttp.BaseRequest
	
	// Whether login is allowed. 1: yes, 0: no
	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// Role ID. Use either `RoleId` or `RoleName` as the input parameter.
	RoleId *int64 `json:"RoleId,omitempty" name:"RoleId"`

	// Role name. Use either `RoleId` or `RoleName` as the input parameter.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateRoleConsoleLoginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConsoleLoginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConsoleLogin")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRoleConsoleLoginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConsoleLoginResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRoleConsoleLoginResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleConsoleLoginResponseParams `json:"Response"`
}

func (r *UpdateRoleConsoleLoginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConsoleLoginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleDescriptionRequestParams struct {
	// Role description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Role ID, used to specify a role. Input either `RoleId` or `RoleName`
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// Role name, used to specify a role. Input either `RoleId` or `RoleName`
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRoleDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleDescriptionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRoleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleDescriptionResponseParams `json:"Response"`
}

func (r *UpdateRoleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSAMLProviderRequestParams struct {
	// SAML identity provider name
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML identity provider description
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML identity provider metadata document (Base64)
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "SAMLMetadataDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSAMLProviderResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSAMLProviderResponseParams `json:"Response"`
}

func (r *UpdateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserOIDCConfigRequestParams struct {
	// OpenID Connect IdP URL.
	// It corresponds to the value of the `issuer` field in the `Openid-configuration` provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Signature public key, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the `authorization_endpoint` field in the `Openid-configuration` provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always `id_token`.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the `id_token` of the IdP is mapped to the username of a sub-user. It is usually the `sub` or `name` field
	MappingFiled *string `json:"MappingFiled,omitempty" name:"MappingFiled"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitempty" name:"Scope"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// OpenID Connect IdP URL.
	// It corresponds to the value of the `issuer` field in the `Openid-configuration` provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitempty" name:"IdentityUrl"`

	// Signature public key, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitempty" name:"IdentityKey"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the `authorization_endpoint` field in the `Openid-configuration` provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always `id_token`.
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the `id_token` of the IdP is mapped to the username of a sub-user. It is usually the `sub` or `name` field
	MappingFiled *string `json:"MappingFiled,omitempty" name:"MappingFiled"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitempty" name:"Scope"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "IdentityKey")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserOIDCConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserOIDCConfigResponseParams `json:"Response"`
}

func (r *UpdateUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "ConsoleLogin")
	delete(f, "Password")
	delete(f, "NeedResetPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserResponseParams `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserSAMLConfigRequestParams struct {
	// Type of the modification operation. `enable`: enable, `disable`: disable, `updateSAML`: modify metadata document.
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// Metadata document, which must be Base64 encoded. This parameter is required only when the value of `Operate` is `updateSAML`.
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

type UpdateUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
	
	// Type of the modification operation. `enable`: enable, `disable`: disable, `updateSAML`: modify metadata document.
	Operate *string `json:"Operate,omitempty" name:"Operate"`

	// Metadata document, which must be Base64 encoded. This parameter is required only when the value of `Operate` is `updateSAML`.
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *UpdateUserSAMLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSAMLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operate")
	delete(f, "SAMLMetadataDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserSAMLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserSAMLConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateUserSAMLConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserSAMLConfigResponseParams `json:"Response"`
}

func (r *UpdateUserSAMLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}