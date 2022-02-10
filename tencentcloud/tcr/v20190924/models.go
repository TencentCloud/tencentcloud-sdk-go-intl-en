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

type CreateInstanceTokenRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Access credential type. Values: `longterm` and `temp` (default, valid for one hour)
	TokenType *string `json:"TokenType,omitempty" name:"TokenType"`

	// Description of the long-term access credential
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "TokenType")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Username
	// Note: this field may return `null`, indicating that no valid value can be found.
		Username *string `json:"Username,omitempty" name:"Username"`

		// Access credential
		Token *string `json:"Token,omitempty" name:"Token"`

		// Expiration timestamp of access credential. It is a string of numbers without unit.
		ExpTime *int64 `json:"ExpTime,omitempty" name:"ExpTime"`

		// Token ID of long-term access credential. It is not available to temporary access credential.
	// Note: this field may return `null`, indicating that no valid value can be found.
		TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceTokenResponse) FromJsonString(s string) error {
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

type CreateReplicationInstanceRequest struct {
	*tchttp.BaseRequest

	// Master instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// Region name of the replication instance
	ReplicationRegionName *string `json:"ReplicationRegionName,omitempty" name:"ReplicationRegionName"`
}

func (r *CreateReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegionId")
	delete(f, "ReplicationRegionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Enterprise Registry Instance ID
		ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceResponse) FromJsonString(s string) error {
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

type DescribeReplicationInstanceCreateTasksRequest struct {
	*tchttp.BaseRequest

	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`
}

func (r *DescribeReplicationInstanceCreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceCreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstanceCreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationInstanceCreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task details
		TaskDetail []*TaskDetail `json:"TaskDetail,omitempty" name:"TaskDetail"`

		// Overall task status
		Status *string `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplicationInstanceCreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceCreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationInstanceSyncStatusRequest struct {
	*tchttp.BaseRequest

	// Master instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// Whether to show the synchronization log
	ShowReplicationLog *bool `json:"ShowReplicationLog,omitempty" name:"ShowReplicationLog"`

	// Page offset for log display. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 5, maximum value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstanceSyncStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceSyncStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	delete(f, "ShowReplicationLog")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstanceSyncStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationInstanceSyncStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Synchronization status
		ReplicationStatus *string `json:"ReplicationStatus,omitempty" name:"ReplicationStatus"`

		// Synchronization completion time
		ReplicationTime *string `json:"ReplicationTime,omitempty" name:"ReplicationTime"`

		// Synchronization log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ReplicationLog *ReplicationLog `json:"ReplicationLog,omitempty" name:"ReplicationLog"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplicationInstanceSyncStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20, maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of instances
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Replication instance list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		ReplicationRegistries []*ReplicationRegistry `json:"ReplicationRegistries,omitempty" name:"ReplicationRegistries"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplicationInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesResponse) FromJsonString(s string) error {
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

type ManageReplicationRequest struct {
	*tchttp.BaseRequest

	// Source instance ID
	SourceRegistryId *string `json:"SourceRegistryId,omitempty" name:"SourceRegistryId"`

	// Destination instance ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitempty" name:"DestinationRegistryId"`

	// Synchronization rule
	Rule *ReplicationRule `json:"Rule,omitempty" name:"Rule"`

	// Rule description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Region ID of the destination instance. For example, `1` represents Guangzhou
	DestinationRegionId *uint64 `json:"DestinationRegionId,omitempty" name:"DestinationRegionId"`

	// Configuration of the synchronization rule
	PeerReplicationOption *PeerReplicationOption `json:"PeerReplicationOption,omitempty" name:"PeerReplicationOption"`
}

func (r *ManageReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceRegistryId")
	delete(f, "DestinationRegistryId")
	delete(f, "Rule")
	delete(f, "Description")
	delete(f, "DestinationRegionId")
	delete(f, "PeerReplicationOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ManageReplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManageReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Instance specification
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RegistryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeerReplicationOption struct {

	// UIN of the destination instance
	PeerRegistryUin *string `json:"PeerRegistryUin,omitempty" name:"PeerRegistryUin"`

	// Permanent access Token for the destination instance
	PeerRegistryToken *string `json:"PeerRegistryToken,omitempty" name:"PeerRegistryToken"`

	// Whether to enable cross-account synchronization
	EnablePeerReplication *bool `json:"EnablePeerReplication,omitempty" name:"EnablePeerReplication"`
}

type ReplicationFilter struct {

	// Type (`name`, `tag` and `resource`)
	Type *string `json:"Type,omitempty" name:"Type"`

	// It is left blank by default
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ReplicationLog struct {

	// Resource type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// Path of the source resource
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitempty" name:"Source"`

	// Path of the destination resource
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Destination *string `json:"Destination,omitempty" name:"Destination"`

	// Synchronization status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Start time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ReplicationRegistry struct {

	// Master instance ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// Region name of the replication instance
	ReplicationRegionName *string `json:"ReplicationRegionName,omitempty" name:"ReplicationRegionName"`

	// Status of the replication instance
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type ReplicationRule struct {

	// Name of synchronization rule
	Name *string `json:"Name,omitempty" name:"Name"`

	// Destination namespace
	DestNamespace *string `json:"DestNamespace,omitempty" name:"DestNamespace"`

	// Whether to override
	Override *bool `json:"Override,omitempty" name:"Override"`

	// Synchronization filters
	Filters []*ReplicationFilter `json:"Filters,omitempty" name:"Filters"`
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

type TaskDetail struct {

	// Task
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// Task UUID
	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`

	// Task status
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// Task details
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

	// Start time of the task
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// End time of the task
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`
}
