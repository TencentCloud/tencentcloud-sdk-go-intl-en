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

package v20190924

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CheckInstanceRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be verified.
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *CheckInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Verification result. true: valid, false: invalid
		IsValidated *bool `json:"IsValidated,omitempty" name:"IsValidated"`

		// ID of the region where the instance is located.
		RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Rule
	Rule *ImmutableTagRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *CreateImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Security group policy
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateMultipleSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultipleSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMultipleSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Security group policy
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteMultipleSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMultipleSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID
		RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMultipleSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Rule list
	// Note: this field may return `null`, indicating that no valid value can be obtained.
		Rules []*ImmutableTagRule `json:"Rules,omitempty" name:"Rules"`

		// Namespace with no rules created
	// Note: this field may return `null`, indicating that no valid value can be obtained.
		EmptyNs []*string `json:"EmptyNs,omitempty" name:"EmptyNs"`

		// Total rules
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImmutableTagRule struct {

	// Repository matching rule
	RepositoryPattern *string `json:"RepositoryPattern,omitempty" name:"RepositoryPattern"`

	// Tag matching rule
	TagPattern *string `json:"TagPattern,omitempty" name:"TagPattern"`

	// repoMatches or repoExcludes
	RepositoryDecoration *string `json:"RepositoryDecoration,omitempty" name:"RepositoryDecoration"`

	// matches or excludes
	TagDecoration *string `json:"TagDecoration,omitempty" name:"TagDecoration"`

	// Disabling rule
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Namespace
	NsName *string `json:"NsName,omitempty" name:"NsName"`
}

type ModifyImmutableTagRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// Rule
	Rule *ImmutableTagRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RuleId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityPolicy struct {

	// Policy index
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// Remarks
	Description *string `json:"Description,omitempty" name:"Description"`

	// The public network IP address of the access source
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// The version of the security policy
	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}
