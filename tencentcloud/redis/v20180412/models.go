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

package v20180412

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Account struct {

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Account name (`root` for a root account)
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Account description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Read/write policy. r: read-only; w: write-only; rw: read/write
	// Note: This field may return null, indicating that no valid values can be obtained.
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// Routing policy. master: master node; replication: secondary node
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy"`

	// Sub-account status. 1: account is being changed; 2: account is valid; -4: account has been deleted
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type AllocateWanAddressRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *AllocateWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AllocateWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Status of enabling public network access
		WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BigKeyInfo struct {

	// Database
	DB *int64 `json:"DB,omitempty" name:"DB"`

	// Big key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Size
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Data timestamp
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type BigKeyTypeInfo struct {

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Size
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Timestamp
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type ChangeReplicaToMasterRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Replica group ID, which is required for multi-AZ instances.
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ChangeReplicaToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeReplicaToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ChangeReplicaToMasterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeReplicaToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommandTake struct {

	// Command
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// Duration
	Took *int64 `json:"Took,omitempty" name:"Took"`
}

type DelayDistribution struct {

	// Delay distribution. The mapping between delay range and `Ladder` value is as follows:
	// [0ms,1ms]: 1;
	// [1ms,5ms]: 5;
	// [5ms,10ms]: 10;
	// [10ms,50ms]: 50;
	// [50ms,200ms]: 200;
	// [200ms,∞]: -1.
	Ladder *int64 `json:"Ladder,omitempty" name:"Ladder"`

	// The number of commands whose delay falls within the current delay range
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Modification time
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type DescribeInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Account details
	// Note: This field may return null, indicating that no valid values can be obtained.
		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

		// Number of accounts
	// Note: This field may return null, indicating that no valid values can be obtained.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitempty" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReqType")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Big key details
		Data []*BigKeyInfo `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeySizeDistRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeySizeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeySizeDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Big key size distribution details
		Data []*DelayDistribution `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyTypeDistRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyTypeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyTypeDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Big key type distribution details
		Data []*BigKeyTypeInfo `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorHotKeyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: past 30 minutes; 3: past 6 hours; 4: past 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorHotKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorHotKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorHotKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Hot key details
		Data []*HotKeyInfo `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorHotKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorSIPRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceMonitorSIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorSIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorSIPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access source information
		Data []*SourceInfo `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorSIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTookDistRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitempty" name:"Date"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTookDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTookDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTookDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Latency distribution information
		Data []*DelayDistribution `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTookDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Access command information
		Data []*SourceCommand `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdTookRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdTookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdTookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Duration details
		Data []*CommandTake `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List size
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The offset value
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of proxy nodes
		ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

		// Proxy node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Proxy []*ProxyNodes `json:"Proxy,omitempty" name:"Proxy"`

		// The number of redis nodes
		RedisCount *int64 `json:"RedisCount,omitempty" name:"RedisCount"`

		// Redis node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Redis []*RedisNodes `json:"Redis,omitempty" name:"Redis"`

		// The number of tendis nodes
		TendisCount *int64 `json:"TendisCount,omitempty" name:"TendisCount"`

		// Tendis node information
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Tendis []*TendisNodes `json:"Tendis,omitempty" name:"Tendis"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Start time of the maintenance window, such as 17:00.
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// End time of the maintenance window, such as 19:00.
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Sale information of a region
		RegionSet []*RegionConf `json:"RegionSet,omitempty" name:"RegionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationGroupRequest struct {
	*tchttp.BaseRequest

	// Instance list size. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Replication group ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Instance ID/name. Fuzzy query is supported.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupId")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of replication group
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Replication group info
		Groups []*Groups `json:"Groups,omitempty" name:"Groups"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Slow log threshold in microseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of slow logs
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Slow log details
		InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitempty" name:"InstanceSlowlogDetail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Groups struct {

	// User App ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Region ID. 1: Guangzhou; 4: Shanghai; 5: Hong Kong (China); 6: Toronto; 7: Shanghai Finance; 8: Beijing; 9: Singapore; 11: Shenzhen Finance; 15: West US (Silicon Valley); 16: Chengdu; 17: Germany; 18: South Korea; 19: Chongqing; 21: India; 22: East US (Virginia); 23: Thailand; 24: Russia; 25: Japan
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Replication group info
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Replication group name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Replication group status. Valid values: 37 (Associating replication group), 38 (Reconnecting to replication group), 51 (Disassociating the replication group), 52 (Switching primary instance in the replication group), 53 (Modifying roles)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Number of replication groups
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instances in replication group
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Instances []*Instances `json:"Instances,omitempty" name:"Instances"`

	// Remarks
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type HotKeyInfo struct {

	// Hot key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance type. Valid values: `2` (Redis 2.8 memory edition in standard architecture), `3` (CKV 3.2 memory edition in standard architecture), `4` (CKV 3.2 memory edition in cluster architecture), `6` (Redis 4.0 memory edition in standard architecture), `7` (Redis 4.0 memory edition in cluster architecture), `8` (Redis 5.0 memory edition in standard architecture), `9` (Redis 5.0 memory edition in cluster architecture).
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` indicates the standard architecture, `MemSize` indicates the total memory capacity of an instance; if `TypeId` indicates the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly-subscribed instance. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing mode. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Instance shard quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, Redis 2.8 standalone edition, and Redis 4.0 standard architecture.
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Instance replica quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price. Unit: USD (accurate down to the cent)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
		Price *float64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard size in MB.
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// Number of shards. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Number of replicas. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`
}

func (r *InquiryPriceUpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price. Unit: USD
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Price *float64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceSlowlogDetail struct {

	// Slow log duration
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// Client address
	Client *string `json:"Client,omitempty" name:"Client"`

	// Command
	Command *string `json:"Command,omitempty" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// Execution duration
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`

	// Node ID
	Node *string `json:"Node,omitempty" name:"Node"`
}

type Instances struct {

	// User App ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Region ID. 1: Guangzhou; 4: Shanghai; 5: Hong Kong (China); 6: Toronto; 7: Shanghai Finance; 8: Beijing; 9: Singapore; 11: Shenzhen Finance; 15: West US (Silicon Valley)
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Number of replicas
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// Number of shards
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// Shard size
	RedisShardSize *int64 `json:"RedisShardSize,omitempty" name:"RedisShardSize"`

	// Instance disk size
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Engine: Redis community edition, Tencent Cloud CKV
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Instance role. Valid values: `rw` (read-write), `r`( read-only)
	Role *string `json:"Role,omitempty" name:"Role"`

	// Instance VIP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Internal parameter, which can be ignored.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// VPC ID, such as 75101
	VpcID *int64 `json:"VpcID,omitempty" name:"VpcID"`

	// Instance Port
	VPort *int64 `json:"VPort,omitempty" name:"VPort"`

	// Instance status. 0: to be initialized; 1: in process; 2: running; -2: isolated; -3: to be deleted
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Repository ID
	GrocerySysId *int64 `json:"GrocerySysId,omitempty" name:"GrocerySysId"`

	// Instance type. Valid values: `1` (Redis 2.8 memory edition in cluster architecture), `2` (Redis 2.8 memory edition in standard architecture), `3` (CKV 3.2 memory edition in standard architecture), `4` (CKV 3.2 memory edition in cluster architecture), `5` (Redis 2.8 memory edition in standalone architecture), `6` (Redis 4.0 memory edition in standard architecture), `7` (Redis 4.0 memory edition in cluster architecture), `8` (Redis 5.0 memory edition in standard architecture), `9` (Redis 5.0 memory edition in cluster architecture)
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ModfiyInstancePasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// New password of an instance
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModfiyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldPassword")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModfiyInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModfiyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModfiyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceReadOnlyRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance input mode. Valid values: `0` (read/write), `1` (read-only)
	InputMode *string `json:"InputMode,omitempty" name:"InputMode"`
}

func (r *ModifyInstanceReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InputMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductConf struct {

	// Product type. Valid values: `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture), `10` (Redis 4.0 Hybrid Storage Edition (Tendis)).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Product name: Redis Master-Replica Edition, CKV Master-Replica Edition, CKV Cluster Edition, Redis Standalone Edition, Redis Cluster Edition, Tendis Hybrid Storage Edition
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// Minimum purchasable quantity
	MinBuyNum *int64 `json:"MinBuyNum,omitempty" name:"MinBuyNum"`

	// Maximum purchasable quantity
	MaxBuyNum *int64 `json:"MaxBuyNum,omitempty" name:"MaxBuyNum"`

	// Whether a product is sold out
	Saleout *bool `json:"Saleout,omitempty" name:"Saleout"`

	// Product engine: Tencent Cloud CKV or Redis community edition
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// Compatible version: Redis 2.8, Redis 3.2, or Redis 4.0
	Version *string `json:"Version,omitempty" name:"Version"`

	// Total capacity in GB
	TotalSize []*string `json:"TotalSize,omitempty" name:"TotalSize"`

	// Shard size in GB
	ShardSize []*string `json:"ShardSize,omitempty" name:"ShardSize"`

	// Number of replicas
	ReplicaNum []*string `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// Number of shards
	ShardNum []*string `json:"ShardNum,omitempty" name:"ShardNum"`

	// Supported billing method. 1: monthly subscription; 0: pay-as-you-go
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to support read-only replicas
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitempty" name:"EnableRepicaReadOnly"`
}

type ProxyNodes struct {

	// Node ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

type RedisNodes struct {

	// Node ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// Shard ID
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type RegionConf struct {

	// Region ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region abbreviation
	RegionShortName *string `json:"RegionShortName,omitempty" name:"RegionShortName"`

	// Name of the area where a region is located
	Area *string `json:"Area,omitempty" name:"Area"`

	// AZ information
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitempty" name:"ZoneSet"`
}

type ReleaseWanAddressRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Status of disabling public network access
		WanStatus *string `json:"WanStatus,omitempty" name:"WanStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SourceCommand struct {

	// Command
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// Number of executions
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SourceInfo struct {

	// Source IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Number of connections
	Conn *int64 `json:"Conn,omitempty" name:"Conn"`

	// Command
	Cmd *int64 `json:"Cmd,omitempty" name:"Cmd"`
}

type TendisNodes struct {

	// Node ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
}

type ZoneCapacityConf struct {

	// AZ ID, such as ap-guangzhou-3
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Whether a product is sold out in an AZ
	IsSaleout *bool `json:"IsSaleout,omitempty" name:"IsSaleout"`

	// Whether it is a default AZ
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// Network type. basenet: basic network; vpcnet: VPC
	NetWorkType []*string `json:"NetWorkType,omitempty" name:"NetWorkType"`

	// Information of an AZ, such as product specifications in it
	ProductSet []*ProductConf `json:"ProductSet,omitempty" name:"ProductSet"`

	// AZ ID, such as 100003
	OldZoneId *int64 `json:"OldZoneId,omitempty" name:"OldZoneId"`
}
