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

package v20191112

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CcnInfo struct {

	// Account of the CCN instance owner
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// CCN ID
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

type CopyFleetRequest struct {
	*tchttp.BaseRequest

	// Server fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Replica number. It should a value between 1 to the number of the remaining quota. It can be obtained through [Obtaining User Quota](https://intl.cloud.tencent.com/document/product/1165/48732?from_cn_redirect=1).
	CopyNumber *int64 `json:"CopyNumber,omitempty" name:"CopyNumber"`

	// Asset package ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// Description. The length is 0-100 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Network configuration
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions"`

	// Server type. It can be obtained through [Obtaining Server Instance Type List](https://intl.cloud.tencent.com/document/product/1165/48732?from_cn_redirect=1).
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Server fleet type, which only supports “ON_DEMAND” type now.
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// Server fleet name. The length is 1-50 characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Protection policy. Valid values: NoProtection·(no protection), FullProtection (full protection), TimeLimitProtection (time-limited protection)
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// Limit policy of resource creation
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// Progress configuration
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// Timeout period of time-limited protection. Value range: 5-1440 minutes. Default value: 60 minutes. This parameter is valid only when NewGameSessionProtectionPolicy is set as TimeLimitProtection.
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// Whether to select scaling. Valid values: SCALING_SELECTED, SCALING_UNSELECTED. Default value: SCALING_UNSELECTED.
	SelectedScalingType *string `json:"SelectedScalingType,omitempty" name:"SelectedScalingType"`

	// Whether to associate the fleet with a CCN instance: CCN_SELECTED_BEFORE_CREATE (associate before creation), CCN_SELECTED_AFTER_CREATE (associated after creation), or CCN_UNSELECTED (do not associate); CCN_UNSELECTED by default
	SelectedCcnType *string `json:"SelectedCcnType,omitempty" name:"SelectedCcnType"`

	// Tag list. Up to 50 tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// System disk. It can be a SSD (CLOUD_SSD) with 100-500 GB capacity or a Premium Cloud Storage disk (CLOUD_PREMIUM) with 50-500 GB capacity. The increment is 1.
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// Data disk. It can be SSD disks (CLOUD_SSD) with 100-32000 GB capacity or Premium Cloud Storage disks (CLOUD_PREMIUM) with 10-32000 GB capacity. The increment is 10. 
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// Whether to select to replicate the timer policy: TIMER_SELECTED or TIMER_UNSELECTED. The default value is TIMER_UNSELECTED.
	SelectedTimerType *string `json:"SelectedTimerType,omitempty" name:"SelectedTimerType"`

	// Information of the CCN instance, including the owner account and the instance ID.
	CcnInfos []*CcnInfo `json:"CcnInfos,omitempty" name:"CcnInfos"`

	// Maximum outbound public network bandwidth of the server fleet. Value range: 1 - 200 Mbps. Default value: 100 Mbps.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *CopyFleetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyFleetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "CopyNumber")
	delete(f, "AssetId")
	delete(f, "Description")
	delete(f, "InboundPermissions")
	delete(f, "InstanceType")
	delete(f, "FleetType")
	delete(f, "Name")
	delete(f, "NewGameServerSessionProtectionPolicy")
	delete(f, "ResourceCreationLimitPolicy")
	delete(f, "RuntimeConfiguration")
	delete(f, "GameServerSessionProtectionTimeLimit")
	delete(f, "SelectedScalingType")
	delete(f, "SelectedCcnType")
	delete(f, "Tags")
	delete(f, "SystemDiskInfo")
	delete(f, "DataDiskInfo")
	delete(f, "SelectedTimerType")
	delete(f, "CcnInfos")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyFleetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CopyFleetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Server fleet attributes
	// Note: this field may return `null`, indicating that no valid value is obtained.
		FleetAttributes []*FleetAttributes `json:"FleetAttributes,omitempty" name:"FleetAttributes"`

		// The number of server fleets
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyFleetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyFleetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// The maximum number of players, which cannot be less than 0.
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Alias ID. You need to specify an alias ID or fleet ID for each request. If both of them are specified, the fleet ID shall prevail.
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Creator ID. Up to 1024 ASCII characters are allowed.
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// Fleet ID. You need to specify an alias ID or fleet ID for each request. If both of them are specified, the fleet ID shall prevail.
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game attributes. Up to 16 groups of attributes are allowed.
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// The attribute details of game server session. Up to 4096 ASCII characters are allowed.
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// The custom ID of game server session. Up to 4096 ASCII characters are allowed.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Idempotency token. Up to 48 ASCII characters are allowed.
	IdempotencyToken *string `json:"IdempotencyToken,omitempty" name:"IdempotencyToken"`

	// The name of game server session. Up to 1024 ASCII characters are allowed.
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGameServerSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaximumPlayerSessionCount")
	delete(f, "AliasId")
	delete(f, "CreatorId")
	delete(f, "FleetId")
	delete(f, "GameProperties")
	delete(f, "GameServerSessionData")
	delete(f, "GameServerSessionId")
	delete(f, "IdempotencyToken")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGameServerSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGameServerSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {

	// SSH private key
	Secret *string `json:"Secret,omitempty" name:"Secret"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DeleteTimerScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// Unique ID of the policy
	TimerId *string `json:"TimerId,omitempty" name:"TimerId"`

	// ID of the fleet to be bound with the policy
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Scheduled scaling policy name
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`
}

func (r *DeleteTimerScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimerScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimerId")
	delete(f, "FleetId")
	delete(f, "TimerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTimerScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTimerScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTimerScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimerScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsRequest struct {
	*tchttp.BaseRequest

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game server session ID. It should contain 1 to 48 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Game server session status. Valid values: ACTIVE, ACTIVATING, TERMINATED, TERMINATING, ERROR
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "FleetId")
	delete(f, "GameServerSessionId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "StatusFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of game server session details
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSessionDetails []*GameServerSessionDetail `json:"GameServerSessionDetails,omitempty" name:"GameServerSessionDetails"`

		// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// Unique ID of game server session placement
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *DescribeGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionPlacementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlacementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionPlacementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session placement
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game server session ID. It should contain 1 to 48 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Game server session status. Valid values: ACTIVE, ACTIVATING, TERMINATED, TERMINATING, ERROR
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "FleetId")
	delete(f, "GameServerSessionId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "StatusFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGameServerSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session list
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions"`

		// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGameServerSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of server types
		InstanceTypeList []*InstanceTypeInfo `json:"InstanceTypeList,omitempty" name:"InstanceTypeList"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsRequest struct {
	*tchttp.BaseRequest

	// Game server session ID. It should contain 1 to 48 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Player ID. It should contain 1 to 1024 ASCII characters.
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player session ID. It should contain 1 to 1024 ASCII characters.
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// Player session status. Valid values: RESERVED, ACTIVE, COMPLETED, TIMEDOUT
	PlayerSessionStatusFilter *string `json:"PlayerSessionStatusFilter,omitempty" name:"PlayerSessionStatusFilter"`
}

func (r *DescribePlayerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayerSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "PlayerId")
	delete(f, "PlayerSessionId")
	delete(f, "PlayerSessionStatusFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePlayerSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Player session list
	// Note: this field may return null, indicating that no valid values can be obtained.
		PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions"`

		// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePlayerSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTimerScalingPoliciesRequest struct {
	*tchttp.BaseRequest

	// ID of the fleet to be bound with the policy
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Scheduled scaling policy name
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`

	// Start time of the scheduled scaling policy
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the scheduled scaling policy
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTimerScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimerScalingPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "TimerName")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimerScalingPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTimerScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration of the scheduled scaling policy
	// Note: this field may return `null`, indicating that no valid value is obtained.
		TimerScalingPolicies []*TimerScalingPolicy `json:"TimerScalingPolicies,omitempty" name:"TimerScalingPolicies"`

		// Total number of scheduled scaling policies
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTimerScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimerScalingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DesiredPlayerSession struct {

	// Unique player ID associated with player session
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Developer-defined player data
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

type DiskInfo struct {

	// Disk type: Premium Cloud Storage (CLOUD_PREMIUM) or SSD (CLOUD_SSD)
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk: the available disk capacity is 50-500 GB. Data disk: the available disk capacity is 100-32000 GB, and the value is a multiple of 10. When the disk type is SSD (CLOUD_SSD), the minimum capacity is 100 GB.
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type FleetAttributes struct {

	// Asset package ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// Server fleet creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Description
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Description of server fleet resource
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FleetArn *string `json:"FleetArn,omitempty" name:"FleetArn"`

	// Server fleet ID
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Server fleet type, which only supports ON_DEMAND now.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// Server type, such as S5.LARGE8
	// Note: this field may return `null`, indicating that no valid value is obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Server fleet name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Game session protection policy
	// Note: this field may return `null`, indicating that no valid value is obtained.
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// Operating system type
	// Note: this field may return `null`, indicating that no valid value is obtained.
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// Limit policy of resource creation
	// Note: this field may return `null`, indicating that no valid value is obtained.
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// Statuses: “Create”, “Downloading”, “Verifying”, “Generating”, “Activating”, “Active”, “Exception”, “Deleting”, and “End”.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The status of server fleet when it stopped. If this field is left empty, it means automatic scaling.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	StoppedActions []*string `json:"StoppedActions,omitempty" name:"StoppedActions"`

	// Server fleet termination time
	// Note: this field may return `null`, indicating that no valid value is obtained.
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// Timeout period of time-limited protection. Value range: 5-1440 minutes. Default value: 60 minutes.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	GameServerSessionProtectionTimeLimit *uint64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// Billing status: Unactivated, Activated, Exception, Isolated due to arrears, Terminated, and Unfrozen.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	BillingStatus *string `json:"BillingStatus,omitempty" name:"BillingStatus"`

	// Tag list. Up to 50 tags.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Data disk. It can be SSD disks (CLOUD_SSD) with 100-32000 GB capacity or Premium Cloud Storage disks (CLOUD_PREMIUM) with 10-32000 GB capacity. The increment is 10. 
	// Note: this field may return `null`, indicating that no valid value is obtained.
	DataDiskInfo []*DiskInfo `json:"DataDiskInfo,omitempty" name:"DataDiskInfo"`

	// System disk. It can be a SSD (CLOUD_SSD) with 100-500 GB capacity or a Premium Cloud Storage disk (CLOUD_PREMIUM) with 50-500 GB capacity. The increment is 1.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	SystemDiskInfo *DiskInfo `json:"SystemDiskInfo,omitempty" name:"SystemDiskInfo"`

	// CCN instance information
	// Note: this field may return `null`, indicating that no valid value is obtained.
	RelatedCcnInfos []*RelatedCcnInfo `json:"RelatedCcnInfos,omitempty" name:"RelatedCcnInfos"`

	// Maximum outbound public network bandwidth of the server fleet. Value range: 1 - 200 Mbps. Default value: 100 Mbps.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type GameProperty struct {

	// Attribute name. Up to 32 ASCII characters are allowed.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Attribute value. Up to 96 ASCII characters are allowed.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GameServerSession struct {

	// Game server session creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Creator ID. Up to 1024 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// The current number of players, which cannot be less than 0.
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// CVM DNS ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game attributes. Up to 16 groups of attributes are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// The attribute details of game server session. Up to 4096 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// Game server session ID. It should contain 1 to 48 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM IP address
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Battle progress details. Up to 400,000 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// The maximum number of players, which cannot be less than 0.
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// The name of game server session. Up to 1024 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player session creation policy. Valid values: ACCEPT_ALL, DENY_ALL
	// Note: this field may return `null`, indicating that no valid value is obtained.
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// Port number. It should be a value between 1 to 60000.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Game server session status. Valid values: ACTIVE, ACTIVATING, TERMINATED, TERMINATING, ERROR
	Status *string `json:"Status,omitempty" name:"Status"`

	// Additional information of game server session status
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// Termination time
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// Instance type. Up to 128 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Current custom count
	// Note: this field may return null, indicating that no valid values can be obtained.
	CurrentCustomCount *int64 `json:"CurrentCustomCount,omitempty" name:"CurrentCustomCount"`

	// Maximum custom count
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxCustomCount *int64 `json:"MaxCustomCount,omitempty" name:"MaxCustomCount"`

	// Weight
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Session availability status, i.e., whether it is blocked. Valid value: Enable, Disable
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AvailabilityStatus *string `json:"AvailabilityStatus,omitempty" name:"AvailabilityStatus"`
}

type GameServerSessionDetail struct {

	// Game server session
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// Protection policy. Valid values: NoProtection, FullProtection
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

type GameServerSessionPlacement struct {

	// Deployment ID
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// Service deployment group name
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// Player latency
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies"`

	// Service deployment status
	Status *string `json:"Status,omitempty" name:"Status"`

	// DNS ID assigned to the instance where the game session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// Game session ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Game session name
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// Service deployment region
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionRegion *string `json:"GameServerSessionRegion,omitempty" name:"GameServerSessionRegion"`

	// Game attributes
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// The maximum number of players that can be connected simultaneously to the game session. It should a value between 1 to the maximum number of player sessions.
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Game session data
	// Note: this field may return null, indicating that no valid values can be obtained.
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// IP address of the instance where the game session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Port number of the instance where the game session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Game match data
	// Note: this field may return null, indicating that no valid values can be obtained.
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// Deployed player game data
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlacedPlayerSessions []*PlacedPlayerSession `json:"PlacedPlayerSessions,omitempty" name:"PlacedPlayerSessions"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type GetGameServerSessionLogUrlRequest struct {
	*tchttp.BaseRequest

	// Game server session ID. It should contain 1 to 48 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`
}

func (r *GetGameServerSessionLogUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGameServerSessionLogUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGameServerSessionLogUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetGameServerSessionLogUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Log download URL. It should contain 1 to 1024 ASCII characters.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		PreSignedUrl *string `json:"PreSignedUrl,omitempty" name:"PreSignedUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGameServerSessionLogUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGameServerSessionLogUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessRequest struct {
	*tchttp.BaseRequest

	// Server fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetInstanceAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInstanceAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetInstanceAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Credentials required for instance login
		InstanceAccess *InstanceAccess `json:"InstanceAccess,omitempty" name:"InstanceAccess"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInstanceAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InboundPermission struct {

	// Start port number. Minimum value: 1025.
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP range. Valid range of the input IPv4 addresses in CIDR format; for example, 0.0.0.0.0/0.
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// Protocol type: TCP or UDP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// End port number. Maximum value: 60000.
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type InstanceAccess struct {

	// Credentials required for instance access
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// Service deployment ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Public IP of instance
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// OS
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
}

type InstanceTypeInfo struct {

	// Name of the server type, such as `Standard SA1`
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// Specification of the server type, such as `SA1.SMALL1`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU, in core
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory, in GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// The packet sending and receiving capability, in 10k PPS. 
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
}

type JoinGameServerSessionBatchRequest struct {
	*tchttp.BaseRequest

	// Game server session ID. It should contain 1 to 256 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Player ID list. At least 1 ID and up to 25 IDs.
	PlayerIds []*string `json:"PlayerIds,omitempty" name:"PlayerIds"`

	// Player custom data
	PlayerDataMap *PlayerDataMap `json:"PlayerDataMap,omitempty" name:"PlayerDataMap"`
}

func (r *JoinGameServerSessionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "PlayerIds")
	delete(f, "PlayerDataMap")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "JoinGameServerSessionBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Player session list. Up to 25 sessions.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinGameServerSessionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// Game server session ID. It should contain 1 to 256 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Player ID. Up to 1024 ASCII characters are allowed.
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player custom data. Up to 2048 ASCII characters are allowed.
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

func (r *JoinGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "PlayerId")
	delete(f, "PlayerData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "JoinGameServerSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Player session
	// Note: this field may return null, indicating that no valid values can be obtained.
		PlayerSession *PlayerSession `json:"PlayerSession,omitempty" name:"PlayerSession"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JoinGameServerSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlacedPlayerSession struct {

	// Player ID
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player session ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`
}

type PlayerDataMap struct {

	// The key of player custom data. It should contain 1 to 1024 ASCII characters.
	Key *string `json:"Key,omitempty" name:"Key"`

	// The value of player custom data. It should contain 1 to 2048 ASCII characters.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PlayerLatency struct {

	// Player ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Name of region corresponding to latency
	// Note: this field may return null, indicating that no valid values can be obtained.
	RegionIdentifier *string `json:"RegionIdentifier,omitempty" name:"RegionIdentifier"`

	// Latency in milliseconds
	LatencyInMilliseconds *uint64 `json:"LatencyInMilliseconds,omitempty" name:"LatencyInMilliseconds"`
}

type PlayerSession struct {

	// Player session creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// ID of the DNS where the game server session is running
	// Note: this field may return null, indicating that no valid values can be obtained.
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Game server session ID. It should contain 1 to 256 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// Address of the CVM instance where the game server session is running
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Player custom data. Up to 2048 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`

	// Player ID. Up to 1024 ASCII characters are allowed.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// Player session ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// Port number. It should be a value between 1 to 60000.
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Player session status. Valid values: RESERVED = 1, ACTIVE = 2, COMPLETED =3, TIMEDOUT = 4
	Status *string `json:"Status,omitempty" name:"Status"`

	// Player session termination time
	// Note: this field may return null, indicating that no valid values can be obtained.
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`
}

type PutTimerScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// Configuration of the scheduled scaling policy
	TimerScalingPolicy *TimerScalingPolicy `json:"TimerScalingPolicy,omitempty" name:"TimerScalingPolicy"`
}

func (r *PutTimerScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutTimerScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimerScalingPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutTimerScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PutTimerScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutTimerScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutTimerScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelatedCcnInfo struct {

	// Account of the CCN instance owner
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// CCN instance ID
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// Status of associated CCN instance
	AttachType *string `json:"AttachType,omitempty" name:"AttachType"`
}

type ResourceCreationLimitPolicy struct {

	// Creation quantity. Minimum value: 1. Default value: 2.
	NewGameServerSessionsPerCreator *uint64 `json:"NewGameServerSessionsPerCreator,omitempty" name:"NewGameServerSessionsPerCreator"`

	// Unit time. Minimum value: 1. Default value: 3. Unit: minute.
	PolicyPeriodInMinutes *uint64 `json:"PolicyPeriodInMinutes,omitempty" name:"PolicyPeriodInMinutes"`
}

type RuntimeConfiguration struct {

	// Game session timeout. Value range: 1-600. Unit: second.
	GameServerSessionActivationTimeoutSeconds *uint64 `json:"GameServerSessionActivationTimeoutSeconds,omitempty" name:"GameServerSessionActivationTimeoutSeconds"`

	// Maximum number of game sessions. Value range: 1-2,147,483,647.
	MaxConcurrentGameServerSessionActivations *uint64 `json:"MaxConcurrentGameServerSessionActivations,omitempty" name:"MaxConcurrentGameServerSessionActivations"`

	// Service process configuration. There must be at least one service configuration.
	ServerProcesses []*ServerProcesse `json:"ServerProcesses,omitempty" name:"ServerProcesses"`
}

type SearchGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// Alias ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// Fleet ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Maximum number of entries in a single query
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// Search filter expression. Valid values:
	// gameServerSessionName: game session name in `String` type
	// gameServerSessionId: game session ID in `String` type
	// maximumSessions: maximum number of player sessions in `Number` type
	// creationTimeMillis: creation time in milliseconds in `Number` type
	// playerSessionCount: current number of player sessions in `Number` type
	// hasAvailablePlayerSessions: whether there is available player session in `String` type. Valid values: true, false
	// gameServerSessionProperties: game session attributes in `String` type
	// 
	// Expressions in `String` type support = and <> for judgment
	// Expressions in `Number` type support =, <>, >, >=, <, and <= for judgment
	// 
	// Example:
	// If FilterExpression takes the value:
	// playerSessionCount>=2 AND hasAvailablePlayerSessions=true"
	// It means searching for game sessions that have at least two players and have player sessions available.
	// If FilterExpression takes the value:
	// gameServerSessionProperties.K1 = 'V1' AND gameServerSessionProperties.K2 = 'V2' OR gameServerSessionProperties.K3 = 'V3'
	// 
	// it means
	// searching for game sessions that meets the following game server session attributes
	// {
	//     "GameProperties":[
	//         {
	//             "Key":"K1",
	//             "Value":"V1"
	//         },
	//         {
	//             "Key":"K2",
	//             "Value":"V2"
	//         },
	//         {
	//             "Key":"K3",
	//             "Value":"V3"
	//         }
	//     ]
	// }
	FilterExpression *string `json:"FilterExpression,omitempty" name:"FilterExpression"`

	// Sorting keyword
	// Valid values:
	// gameServerSessionName: game session name in `String` type
	// gameServerSessionId: game session ID in `String` type
	// maximumSessions: maximum number of player sessions in `Number` type
	// creationTimeMillis: creation time in milliseconds in `Number` type
	// playerSessionCount: current number of player sessions in `Number` type
	SortExpression *string `json:"SortExpression,omitempty" name:"SortExpression"`
}

func (r *SearchGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchGameServerSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AliasId")
	delete(f, "FleetId")
	delete(f, "Limit")
	delete(f, "NextToken")
	delete(f, "FilterExpression")
	delete(f, "SortExpression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchGameServerSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session list
	// Note: this field may return null, indicating that no valid values can be obtained.
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions"`

		// Pagination offset, which is used for querying the next page. It should contain 1 to 1024 ASCII characters.
	// Note: this field may return `null`, indicating that no valid value is obtained.
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchGameServerSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerProcesse struct {

	// Number of concurrent processes. Value range of total concurrent processes: 1-50.
	ConcurrentExecutions *uint64 `json:"ConcurrentExecutions,omitempty" name:"ConcurrentExecutions"`

	// Launch Path. Linux: /local/game/ or Windows: C:\game\. The path length is 1-1024.
	LaunchPath *string `json:"LaunchPath,omitempty" name:"LaunchPath"`

	// Launch parameter. The length is 0-1024.
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
}

type SetServerReservedRequest struct {
	*tchttp.BaseRequest

	// ID of the fleet to be bound with the policy
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether the instance is retained. Valid values: 1 (retained), 0 (not retained). Default value: 0.
	ReserveValue *int64 `json:"ReserveValue,omitempty" name:"ReserveValue"`
}

func (r *SetServerReservedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetServerReservedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FleetId")
	delete(f, "InstanceId")
	delete(f, "ReserveValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetServerReservedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetServerReservedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetServerReservedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetServerReservedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// The unique ID of the game server session placement. It should contain up to 48 ASCII characters, supporting [a-zA-Z0-9-]+.
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// Game server session queue name
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// The maximum number of players that can be connected simultaneously to the game session. It should a value between 1 to the maximum number of player sessions.
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Player game session information
	DesiredPlayerSessions []*DesiredPlayerSession `json:"DesiredPlayerSessions,omitempty" name:"DesiredPlayerSessions"`

	// Player game session attributes
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties"`

	// Data of game server sessions. Up to 4096 ASCII characters are allowed.
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// Name of game server sessions. Up to 4096 ASCII characters are allowed.
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// Player latency
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies"`
}

func (r *StartGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartGameServerSessionPlacementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlacementId")
	delete(f, "GameServerSessionQueueName")
	delete(f, "MaximumPlayerSessionCount")
	delete(f, "DesiredPlayerSessions")
	delete(f, "GameProperties")
	delete(f, "GameServerSessionData")
	delete(f, "GameServerSessionName")
	delete(f, "PlayerLatencies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartGameServerSessionPlacementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session placement
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartGameServerSessionPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// Unique ID of game server session placement
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *StopGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameServerSessionPlacementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlacementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopGameServerSessionPlacementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Game server session placement
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGameServerSessionPlacementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// Tag key. Up to 127 bytes are allowed.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value. Up to 255 bytes are allowed.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TargetConfiguration struct {

	// Ratio of reserved server session resource 
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TargetValue *uint64 `json:"TargetValue,omitempty" name:"TargetValue"`
}

type TimerConfiguration struct {

	// The recurrence pattern of auto-scaling. `0`: undefined, `1`: once, `2`: daily, `3`: monthly, `4`: weekly
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerType *int64 `json:"TimerType,omitempty" name:"TimerType"`

	// Details of the recurrence pattern of auto-scaling
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerValue *TimerValue `json:"TimerValue,omitempty" name:"TimerValue"`

	// Start time of the scheduled scaling policy
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time of the scheduled scaling policy
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type TimerFleetCapacity struct {

	// ID of the fleet to be bound with the policy
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// Desired number of instances
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DesiredInstances *int64 `json:"DesiredInstances,omitempty" name:"DesiredInstances"`

	// Minimum number of instances
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum number of instances
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Scaling cooldown period, in minutes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ScalingInterval *int64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`

	// Scaling type. `1`: manual, `2`: automatic, `0`: undefined
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ScalingType *int64 `json:"ScalingType,omitempty" name:"ScalingType"`

	// Configuration of target tracking scaling
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

type TimerScalingPolicy struct {

	// Unique ID of the policy. When it’s filled in, the policy will be updated.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerId *string `json:"TimerId,omitempty" name:"TimerId"`

	// Scheduled scaling policy name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerName *string `json:"TimerName,omitempty" name:"TimerName"`

	// Scheduled scaling policy status. `0`: Undefined, `1`: Not started, 2: Activated, `3`: Stopped, `4`: Expired
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerStatus *int64 `json:"TimerStatus,omitempty" name:"TimerStatus"`

	// The capacity configurations of the scheduled scaling policy
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerFleetCapacity *TimerFleetCapacity `json:"TimerFleetCapacity,omitempty" name:"TimerFleetCapacity"`

	// The recurrence pattern of auto-scaling
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TimerConfiguration *TimerConfiguration `json:"TimerConfiguration,omitempty" name:"TimerConfiguration"`
}

type TimerValue struct {

	// Execute once every X day(s)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Day *int64 `json:"Day,omitempty" name:"Day"`

	// Specify the first day to execute the scaling action in a month (execute once per day)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FromDay *int64 `json:"FromDay,omitempty" name:"FromDay"`

	// Specify the last day to execute the scaling action in a month
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ToDay *int64 `json:"ToDay,omitempty" name:"ToDay"`

	// Specify the week days to repeat the scaling action. Multiple values are supported. Valid values: `1` (Monday), `2` (Tuesday), `3` (Wednesday), `4` (Thursday), `5` (Friday), `6` (Saturday), `7` (Sunday). 
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	WeekDays []*int64 `json:"WeekDays,omitempty" name:"WeekDays"`
}

type UpdateBucketAccelerateOptRequest struct {
	*tchttp.BaseRequest

	// `true`: enable global acceleration; `false`: disable global acceleration
	Allowed *bool `json:"Allowed,omitempty" name:"Allowed"`
}

func (r *UpdateBucketAccelerateOptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketAccelerateOptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Allowed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBucketAccelerateOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBucketAccelerateOptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBucketAccelerateOptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketAccelerateOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBucketCORSOptRequest struct {
	*tchttp.BaseRequest

	// Allowed access source. For details, see [COS Documentation](https://intl.cloud.tencent.com/document/product/436/8279?from_cn_redirect=1).
	AllowedOrigins []*string `json:"AllowedOrigins,omitempty" name:"AllowedOrigins"`

	// Allowed HTTP method(s). Multiple methods are allowed, including PUT, GET, POST, and HEAD. For details, see [COS Documentation](https://intl.cloud.tencent.com/document/product/436/8279?from_cn_redirect=1).
	AllowedMethods []*string `json:"AllowedMethods,omitempty" name:"AllowedMethods"`

	// Specifies the custom HTTP request headers that the browser is allowed to include in a CORS request. Wildcard (*) is supported, indicating allowing all headers (recommended). For details, see [COS Documentation](https://intl.cloud.tencent.com/document/product/436/8279?from_cn_redirect=1).
	AllowedHeaders []*string `json:"AllowedHeaders,omitempty" name:"AllowedHeaders"`

	// Sets the validity duration for the CORS configuration (in second). For details, see [COS Documentation](https://intl.cloud.tencent.com/document/product/436/8279?from_cn_redirect=1).
	MaxAgeSeconds *int64 `json:"MaxAgeSeconds,omitempty" name:"MaxAgeSeconds"`

	// CORS response header(s) that can be exposed to the browser, case-insensitive. If this parameter is not specified, the browser can access only simple response headers Cache-Control, Content-Type, Expires, and Last-Modified by default. For details, see [COS Documentation](https://intl.cloud.tencent.com/document/product/436/8279?from_cn_redirect=1).
	ExposeHeaders []*string `json:"ExposeHeaders,omitempty" name:"ExposeHeaders"`
}

func (r *UpdateBucketCORSOptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketCORSOptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllowedOrigins")
	delete(f, "AllowedMethods")
	delete(f, "AllowedHeaders")
	delete(f, "MaxAgeSeconds")
	delete(f, "ExposeHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBucketCORSOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBucketCORSOptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBucketCORSOptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBucketCORSOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// Game server session ID. It should contain 1 to 256 ASCII characters.
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// The maximum number of players, which cannot be less than 0.
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// Name of the game server session. It should contain 1 to 1024 ASCII characters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Player session creation policy, which includes `ACCEPT_ALL` (allow all players) and `DENY_ALL` (reject all players).
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// Protection policy, which includes `NoProtection`·(no protection), `TimeLimitProtection` (time-limited protection) and `FullProtection` (full protection)
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

func (r *UpdateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGameServerSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GameServerSessionId")
	delete(f, "MaximumPlayerSessionCount")
	delete(f, "Name")
	delete(f, "PlayerSessionCreationPolicy")
	delete(f, "ProtectionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGameServerSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Updated game session
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGameServerSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
