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

package v20170320

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Account struct {
	// Account name, enter 1-32 characters.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Account's host.
	// Note:
	// 1. IP format. You can specify a percent sign (%).
	// 2. Multiple hosts are separated by a separator, which supports ;, |, line break, and space.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type AccountInfo struct {
	// Account remarks
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// Account domain name
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Account name
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Account information modification time
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Password modification time
	ModifyPasswordTime *string `json:"ModifyPasswordTime,omitnil,omitempty" name:"ModifyPasswordTime"`

	// This parameter is deprecated.
	//
	// Deprecated: CreateTime is deprecated.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The maximum number of instance connections supported by an account
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

	// Is password rotation enabled for the user account?
	OpenCam *bool `json:"OpenCam,omitnil,omitempty" name:"OpenCam"`
}

// Predefined struct for user
type AddTimeWindowRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maintenance window on Monday. The format should be 10:00-12:00. You can set multiple time windows on a day. Each time window lasts from half an hour to three hours, and must start and end on the hour or half hour. At least one time window is required in a week. The same rule applies to the following parameters.
	Monday []*string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// Maintenance window on Tuesday. At least one time window is required in a week.
	Tuesday []*string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// Maintenance window on Wednesday. At least one time window is required in a week.
	Wednesday []*string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// Maintenance window on Thursday. At least one time window is required in a week.
	Thursday []*string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// Maintenance window on Friday. At least one time window is required in a week.
	Friday []*string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// Maintenance window on Saturday. At least one time window is required in a week.
	Saturday []*string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// Maintenance window on Sunday. At least one time window is required in a week.
	Sunday []*string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// Maximum delay threshold (seconds), only applicable to primary instance and disaster recovery instance. Default value: 10. Value ranges from 1 to 10 integers.
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maintenance window on Monday. The format should be 10:00-12:00. You can set multiple time windows on a day. Each time window lasts from half an hour to three hours, and must start and end on the hour or half hour. At least one time window is required in a week. The same rule applies to the following parameters.
	Monday []*string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// Maintenance window on Tuesday. At least one time window is required in a week.
	Tuesday []*string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// Maintenance window on Wednesday. At least one time window is required in a week.
	Wednesday []*string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// Maintenance window on Thursday. At least one time window is required in a week.
	Thursday []*string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// Maintenance window on Friday. At least one time window is required in a week.
	Friday []*string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// Maintenance window on Saturday. At least one time window is required in a week.
	Saturday []*string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// Maintenance window on Sunday. At least one time window is required in a week.
	Sunday []*string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// Maximum delay threshold (seconds), only applicable to primary instance and disaster recovery instance. Default value: 10. Value ranges from 1 to 10 integers.
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

func (r *AddTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Monday")
	delete(f, "Tuesday")
	delete(f, "Wednesday")
	delete(f, "Thursday")
	delete(f, "Friday")
	delete(f, "Saturday")
	delete(f, "Sunday")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTimeWindowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *AddTimeWindowResponseParams `json:"Response"`
}

func (r *AddTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddressInfo struct {
	// Resource ID identifier of the address.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// The vpc where the address resides.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// The subnet where the address resides.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// vip address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// The port of the address.
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Public network address domain name.
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public network address port.
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`
}

// Predefined struct for user
type AdjustCdbProxyAddressRequestParams struct {
	// <p>Proxy group ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1">DescribeCdbProxyInfo</a> API.</p>
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// <p>Weight allocation mode,<br>system Auto-Assignment: "system", custom: "custom"</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Whether delay removal is enabled. Value: "true" | "false"</p>
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// <p>Minimum retention quantity, minimum value: 0.<br>Description: Valid only when IsKickOut is true.</p>
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// <p>Delay removal threshold, minimum value: 1, value ranges from 1 to 10000, integer.</p>
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// <p>Whether fault migration is enabled, value: "true" | "false"</p>
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Automatically add RO. Parameter: "true" | "false"</p>
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Whether it is read-only. Value: "true" | "false".</p>
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// <p>Proxy group address ID. Obtain through the <a href="https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1">DescribeCdbProxyInfo</a> API.</p>
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// <p>Whether transaction splitting is enabled. Value: "true" | "false". Default value: false.</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Whether the connection pool is enabled. Off by default.<br>Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be at least MySQL 8.0 20230630.</p>
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// <p>Read-write weight allocation. If WeightMode is passed in as system, the passed-in weight does not take effect and the default weight is assigned by the system.</p>
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// <p>Whether self-adaptive load balancing is enabled. Off by default.</p>
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// <p>Access mode: nearby - proximity access, balance - balanced allocation. Default is proximity access.</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Whether to treat the libra node as an ordinary RO node</p>
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// <p>Whether to forward to other nodes in case of a libra node fault</p>
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

type AdjustCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// <p>Proxy group ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1">DescribeCdbProxyInfo</a> API.</p>
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// <p>Weight allocation mode,<br>system Auto-Assignment: "system", custom: "custom"</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>Whether delay removal is enabled. Value: "true" | "false"</p>
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// <p>Minimum retention quantity, minimum value: 0.<br>Description: Valid only when IsKickOut is true.</p>
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// <p>Delay removal threshold, minimum value: 1, value ranges from 1 to 10000, integer.</p>
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// <p>Whether fault migration is enabled, value: "true" | "false"</p>
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// <p>Automatically add RO. Parameter: "true" | "false"</p>
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// <p>Whether it is read-only. Value: "true" | "false".</p>
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// <p>Proxy group address ID. Obtain through the <a href="https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1">DescribeCdbProxyInfo</a> API.</p>
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// <p>Whether transaction splitting is enabled. Value: "true" | "false". Default value: false.</p>
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// <p>Whether the connection pool is enabled. Off by default.<br>Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be at least MySQL 8.0 20230630.</p>
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// <p>Read-write weight allocation. If WeightMode is passed in as system, the passed-in weight does not take effect and the default weight is assigned by the system.</p>
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// <p>Whether self-adaptive load balancing is enabled. Off by default.</p>
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// <p>Access mode: nearby - proximity access, balance - balanced allocation. Default is proximity access.</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Whether to treat the libra node as an ordinary RO node</p>
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// <p>Whether to forward to other nodes in case of a libra node fault</p>
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

func (r *AdjustCdbProxyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "WeightMode")
	delete(f, "IsKickOut")
	delete(f, "MinCount")
	delete(f, "MaxDelay")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "ReadOnly")
	delete(f, "ProxyAddressId")
	delete(f, "TransSplit")
	delete(f, "ConnectionPool")
	delete(f, "ProxyAllocation")
	delete(f, "AutoLoadBalance")
	delete(f, "AccessMode")
	delete(f, "ApNodeAsRoNode")
	delete(f, "ApQueryToOtherNode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyAddressResponseParams struct {
	// <p>Asynchronous Task ID</p>
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AdjustCdbProxyAddressResponse struct {
	*tchttp.BaseResponse
	Response *AdjustCdbProxyAddressResponseParams `json:"Response"`
}

func (r *AdjustCdbProxyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Node specification configuration
	// Remark: Database proxy supported node specifications are 2C4000MB, 4C8000MB, 8C16000MB.
	// Parameter description in the example.
	// NodeCount: Number of nodes
	// Region: Node region
	// Zone: Node availability zone
	// Cpu: Number of node cores for one agent (Unit: core)
	// Mem: Memory size of each proxy node (unit: MB)
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Rebalance. Valid values:  `auto` (automatic), `manual` (manual).
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// The upgrade switch time. Valid values:  `nowTime` (upgrade immediately), `timeWindow` (upgrade during instance maintenance time).
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

type AdjustCdbProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Node specification configuration
	// Remark: Database proxy supported node specifications are 2C4000MB, 4C8000MB, 8C16000MB.
	// Parameter description in the example.
	// NodeCount: Number of nodes
	// Region: Node region
	// Zone: Node availability zone
	// Cpu: Number of node cores for one agent (Unit: core)
	// Mem: Memory size of each proxy node (unit: MB)
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Rebalance. Valid values:  `auto` (automatic), `manual` (manual).
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// The upgrade switch time. Valid values:  `nowTime` (upgrade immediately), `timeWindow` (upgrade during instance maintenance time).
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

func (r *AdjustCdbProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "ProxyNodeCustom")
	delete(f, "ReloadBalance")
	delete(f, "UpgradeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustCdbProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyResponseParams struct {
	// Asynchronous Task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AdjustCdbProxyResponse struct {
	*tchttp.BaseResponse
	Response *AdjustCdbProxyResponseParams `json:"Response"`
}

func (r *AdjustCdbProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustCdbProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggregationCondition struct {
	// Aggregation field. Valid values: `host` (source IP), `user` （username), `dbName` (database name), `sqlType` (SQL type).
	AggregationField *string `json:"AggregationField,omitnil,omitempty" name:"AggregationField"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of buckets returned under this field. Maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type AnalysisNodeInfo struct {
	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Data loading status.
	DataStatus *string `json:"DataStatus,omitnil,omitempty" name:"DataStatus"`

	// Number of CPU cores, in cores.
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size, in MB.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size, in GB.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Node AZ.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Data synchronization error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type AnalyzeAuditLogsRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting conditions for aggregation dimension
	AggregationConditions []*AggregationCondition `json:"AggregationConditions,omitnil,omitempty" name:"AggregationConditions"`

	// Deprecated.
	//
	// Deprecated: AuditLogFilter is deprecated.
	AuditLogFilter *AuditLogFilter `json:"AuditLogFilter,omitnil,omitempty" name:"AuditLogFilter"`

	// The result set of the audit log filtered by this condition is set as the analysis Log.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type AnalyzeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting conditions for aggregation dimension
	AggregationConditions []*AggregationCondition `json:"AggregationConditions,omitnil,omitempty" name:"AggregationConditions"`

	// Deprecated.
	AuditLogFilter *AuditLogFilter `json:"AuditLogFilter,omitnil,omitempty" name:"AuditLogFilter"`

	// The result set of the audit log filtered by this condition is set as the analysis Log.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

func (r *AnalyzeAuditLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeAuditLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AggregationConditions")
	delete(f, "AuditLogFilter")
	delete(f, "LogFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AnalyzeAuditLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeAuditLogsResponseParams struct {
	// Returned aggregation bucket information set
	Items []*AuditLogAggregationResult `json:"Items,omitnil,omitempty" name:"Items"`

	// Number of logs scanned
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AnalyzeAuditLogsResponse struct {
	*tchttp.BaseResponse
	Response *AnalyzeAuditLogsResponseParams `json:"Response"`
}

func (r *AnalyzeAuditLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeAuditLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// Security group ID, which can be obtained through the [DescribeDBSecurityGroups](https://www.tencentcloud.com/document/api/236/15854?from_cn_redirect=1) API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance ID list, an array consisting of one or more instance IDs. You can obtain it through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be bound to the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be bound to the read-only replicas themselves.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID, which can be obtained through the [DescribeDBSecurityGroups](https://www.tencentcloud.com/document/api/236/15854?from_cn_redirect=1) API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance ID list, an array consisting of one or more instance IDs. You can obtain it through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be bound to the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be bound to the read-only replicas themselves.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditFilter struct {
	// Filter parameter names. Valid values:
	// SrcIp: Client IP;
	// User: Database account;
	// DB: Database name.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter match type. Valid value:
	// `INC`: Include;
	// `EXC`: Exclude;
	// `EQ`: Equal to;
	// `NEQ`: Not equal to.
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// Filter match value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AuditInstanceFilters struct {
	// Filter condition name. Valid values: InstanceId - Instance ID, InstanceName - Instance name, ProjectId - Project ID, TagKey - Tag key, Tag - Tag (using a vertical bar as separator, for example: TagKey|Tagvalue).
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// true indicates exact matching; false indicates fuzzy matching.
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`

	// Filter value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AuditInstanceInfo struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Tag information.
	TagList []*TagInfoUnit `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Database kernel type.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database kernel version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type AuditLog struct {
	// Number of affected rows
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// The error code
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`


	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// Audit policy name, which will be unavailable soon.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`


	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`


	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Client address
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Username
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Execution time (μs)
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// Time
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Number of returned rows
	SentRows *int64 `json:"SentRows,omitnil,omitempty" name:"SentRows"`

	// Thread ID
	ThreadId *int64 `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// Number of scanned rows.
	CheckRows *int64 `json:"CheckRows,omitnil,omitempty" name:"CheckRows"`

	// cpu execution time, µs.
	CpuTime *float64 `json:"CpuTime,omitnil,omitempty" name:"CpuTime"`

	// IO wait time, µs.
	IoWaitTime *uint64 `json:"IoWaitTime,omitnil,omitempty" name:"IoWaitTime"`

	// Lock waiting time (unit: microsecond).
	LockWaitTime *uint64 `json:"LockWaitTime,omitnil,omitempty" name:"LockWaitTime"`

	// Start time, which combines with timestamp to form a time accurate to nanoseconds.
	NsTime *uint64 `json:"NsTime,omitnil,omitempty" name:"NsTime"`

	// Transaction duration, µs.
	TrxLivingTime *uint64 `json:"TrxLivingTime,omitnil,omitempty" name:"TrxLivingTime"`

	// Basic info of the log hit rule template
	TemplateInfo []*LogRuleTemplateInfo `json:"TemplateInfo,omitnil,omitempty" name:"TemplateInfo"`

	// Transaction ID
	TrxId *int64 `json:"TrxId,omitnil,omitempty" name:"TrxId"`

	// Port.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientPort *int64 `json:"ClientPort,omitnil,omitempty" name:"ClientPort"`
}

type AuditLogAggregationResult struct {
	// Aggregation dimension
	AggregationField *string `json:"AggregationField,omitnil,omitempty" name:"AggregationField"`

	// Aggregate bucket result set
	Buckets []*Bucket `json:"Buckets,omitnil,omitempty" name:"Buckets"`
}

type AuditLogFile struct {
	// Audit log file name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Creation time of the audit log file, in the format: "2019-03-20 17:09:13".
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// File status. Possible return values:"creating" - Generating;"failed" - Creation failed;"success" - Generated.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// File size in KB.
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Download URL for the audit log.
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// Error message.
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type AuditLogFilter struct {
	// Client address
	Host []*string `json:"Host,omitnil,omitempty" name:"Host"`

	// Username
	User []*string `json:"User,omitnil,omitempty" name:"User"`


	DBName []*string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Table name
	TableName []*string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Audit policy name
	PolicyName []*string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`


	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`


	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// Execution time in ms, which is used to filter the audit log with execution time greater than this value.
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// Number of affected rows, which is used to filter the audit log with affected rows greater than this value.
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// SQL type (Multiple types can be queried at same time). Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `ALTER`, `SET`, `REPLACE`, `EXECUTE`.
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// SQL statement. Multiple SQL statements can be passed in.
	Sqls []*string `json:"Sqls,omitnil,omitempty" name:"Sqls"`

	// Number of rows affected in the format of M-N, such as 10-200.
	AffectRowsSection *string `json:"AffectRowsSection,omitnil,omitempty" name:"AffectRowsSection"`

	// Number of rows returned in the format of M-N, such as 10-200.
	SentRowsSection *string `json:"SentRowsSection,omitnil,omitempty" name:"SentRowsSection"`

	// Execution time in the format of M-N, such as 10-200.
	ExecTimeSection *string `json:"ExecTimeSection,omitnil,omitempty" name:"ExecTimeSection"`

	// Lock wait time in the format of M-N, such as 10-200.
	LockWaitTimeSection *string `json:"LockWaitTimeSection,omitnil,omitempty" name:"LockWaitTimeSection"`

	// IO wait time in the format of M-N, such as 10-200.
	IoWaitTimeSection *string `json:"IoWaitTimeSection,omitnil,omitempty" name:"IoWaitTimeSection"`

	// Transaction duration in the format of M-N, such as 10-200.
	TransactionLivingTimeSection *string `json:"TransactionLivingTimeSection,omitnil,omitempty" name:"TransactionLivingTimeSection"`

	// Thread ID
	ThreadId []*string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// Number of returned rows,  which is used to filter the audit log with affected rows greater than this value.
	SentRows *int64 `json:"SentRows,omitnil,omitempty" name:"SentRows"`

	// MySQL error codes
	ErrCode []*int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`
}

type AuditPolicy struct {
	// Audit policy ID.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Audit policy status. Valid values:
	// `creating`;
	// `running`,
	// `paused`;
	// `failed`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Database instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Creation time of audit policy in the format of 2019-03-20 17:09:13
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of audit policy in the format of 2019-03-20 17:09:13
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Audit policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Audit rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Audit rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Database instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type AuditRule struct {
	// Audit rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Creation time of audit rule in the format of 2019-03-20 17:09:13
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last modified time of audit rule in the format of 2019-03-20 17:09:13
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Audit rule name
	// Note: This field may return `null`, indicating that no valid value was found.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Audit rule description
	// Note: This field may return `null`, indicating that no valid value was found.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Audit rule filters
	// Note: This field may return `null`, indicating that no valid value was found.
	RuleFilters []*AuditFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Whether to enable full audit
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type AuditRuleFilters struct {
	// A single audit rule.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

type AuditRuleTemplateInfo struct {
	// Rule template ID.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Rule template name.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Filter conditions of the rule template.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Rule template creation time.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// Instances to which this rule template is applied.
	AffectedInstances []*string `json:"AffectedInstances,omitnil,omitempty" name:"AffectedInstances"`

	// Template status. Valid values: 0 - No task, 1 - modifying.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Template update time.
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type AutoStrategy struct {
	// Auto scaling threshold. Available values: 40, 50, 60, 70, 80, 90. Represents the CPU utilization reaches 40%, 50%, 60%, 70%, 80%, or 90% to trigger auto scaling in the background.
	ExpandThreshold *int64 `json:"ExpandThreshold,omitnil,omitempty" name:"ExpandThreshold"`

	// CPU utilization threshold (percent value). Valid values: 10, 20, and 30. Automatic scale-in will be triggered when CPU utilization reaches the set threshold.
	ShrinkThreshold *int64 `json:"ShrinkThreshold,omitnil,omitempty" name:"ShrinkThreshold"`

	// Auto-scaling observation period, in minutes, available values 1, 3, 5, 10, 15, 30. The backend will judge scaling out according to the configured period.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ExpandPeriod is deprecated.
	ExpandPeriod *int64 `json:"ExpandPeriod,omitnil,omitempty" name:"ExpandPeriod"`

	// Automatic scaling down observation period, in minutes, available values 5, 10, 15, 30. The backend performs scale-in judgment according to the configured period.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: ShrinkPeriod is deprecated.
	ShrinkPeriod *int64 `json:"ShrinkPeriod,omitnil,omitempty" name:"ShrinkPeriod"`

	// Elastic scaling observation period (in seconds). Value is 15, 30, 45, 60, 180, 300, 600, 900, or 1800.
	ExpandSecondPeriod *int64 `json:"ExpandSecondPeriod,omitnil,omitempty" name:"ExpandSecondPeriod"`

	// Scale-down observation period (in seconds). Valid values: 300, 600, 900, 1800.
	ShrinkSecondPeriod *int64 `json:"ShrinkSecondPeriod,omitnil,omitempty" name:"ShrinkSecondPeriod"`
}

type BackupConfig struct {
	// Replication mode of secondary database 2. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitnil,omitempty" name:"ReplicationMode"`

	// The canonical name of the second secondary availability zone, for example ap-shanghai-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Private IP address of secondary database 2
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Access port of secondary database 2
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type BackupInfo struct {
	// <p>Backup file name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Backup file size, unit: Byte</p>
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// <p>Backup snapshot time. Time format: 2016-03-17 02:10:37</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// <p>Download link</p>
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// <p>Download link</p>
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// <p>Specific type of logs. Possible values: "logical": logical cold backup, "physical": physical cold backup.</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>ID of the backup subtask, used when deleting backup files</p>
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// <p>Backup task status. Possible values: "SUCCESS": backup successful, "FAILED": backup FAILED, "RUNNING": backup in progress.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Backup task completion time</p>
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// <p>(This value will be deprecated and is not recommended for use) Creator of the backup. Valid values: SYSTEM - generated by the system, Uin - Uin value of the initiator.</p>
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>Start time of the backup task</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>Backup method. Possible values are "full": full backup, "partial": partial backup.</p>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>Backup method. Possible values are "manual": manual backup, "automatic": automatic backup.</p>
	Way *string `json:"Way,omitnil,omitempty" name:"Way"`

	// <p>Manual backup alias</p>
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// <p>Backup retention type, save_mode_regular - Regular backup saving, save_mode_period - Periodic backup</p>
	SaveMode *string `json:"SaveMode,omitnil,omitempty" name:"SaveMode"`

	// <p>Local backup region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Detailed information of offsite backup</p>
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitnil,omitempty" name:"RemoteInfo"`

	// <p>Storage method: 0 - regular storage, 1 - archive storage, 2 - standard storage, defaults to 0.</p>
	CosStorageType *int64 `json:"CosStorageType,omitnil,omitempty" name:"CosStorageType"`

	// <p>Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup completion progress</p>
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// <p>Whether the backup file is encrypted, on-encrypted, off-unencrypted</p>
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`

	// <p>Backup GTID position</p>
	ExecutedGTIDSet *string `json:"ExecutedGTIDSet,omitnil,omitempty" name:"ExecutedGTIDSet"`

	// <p>MD5 value of the backup file</p>
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`
}

type BackupItem struct {
	// Name of the database to be backed up
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Name of the table to be backed up. If this parameter is passed in, the specified table in the database will be backed up; otherwise, the database will be backed up.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type BackupLimitVpcItem struct {
	// The region where the backup download restrictions take effect. It must be the same as the common request parameter `Region` of the API.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of VPCs used to restrict backup download
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`
}

type BackupSummaryItem struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Number of automatic data backups of an instance.
	AutoBackupCount *int64 `json:"AutoBackupCount,omitnil,omitempty" name:"AutoBackupCount"`

	// Capacity of automatic data backups of an instance.
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitnil,omitempty" name:"AutoBackupVolume"`

	// Number of manual data backups of an instance.
	ManualBackupCount *int64 `json:"ManualBackupCount,omitnil,omitempty" name:"ManualBackupCount"`

	// Capacity of manual data backups of an instance.
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitnil,omitempty" name:"ManualBackupVolume"`

	// Total number of data backups of an instance (including automatic backups and manual backups).
	DataBackupCount *int64 `json:"DataBackupCount,omitnil,omitempty" name:"DataBackupCount"`

	// Total capacity of data backups of an instance.
	DataBackupVolume *int64 `json:"DataBackupVolume,omitnil,omitempty" name:"DataBackupVolume"`

	// Number of log backups of an instance.
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitnil,omitempty" name:"BinlogBackupCount"`

	// Capacity of log backups of an instance.
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitnil,omitempty" name:"BinlogBackupVolume"`

	// Total capacity of backups of an instance (including data backups and log backups).
	BackupVolume *int64 `json:"BackupVolume,omitnil,omitempty" name:"BackupVolume"`
}

// Predefined struct for user
type BalanceRoGroupLoadRequestParams struct {
	// ID of the RO group, in the format of cdbrg-c1nl9rpv. You can obtain it via [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1).
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type BalanceRoGroupLoadRequest struct {
	*tchttp.BaseRequest
	
	// ID of the RO group, in the format of cdbrg-c1nl9rpv. You can obtain it via [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1).
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

func (r *BalanceRoGroupLoadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BalanceRoGroupLoadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BalanceRoGroupLoadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BalanceRoGroupLoadResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BalanceRoGroupLoadResponse struct {
	*tchttp.BaseResponse
	Response *BalanceRoGroupLoadResponseParams `json:"Response"`
}

func (r *BalanceRoGroupLoadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BalanceRoGroupLoadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BinlogInfo struct {
	// <p>binlog backup file name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Backup file size, unit: Byte</p>
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// <p>File storage time. Time format: 2016-03-17 02:10:37</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// <p>Download link<br>Description: This download link is the same as the download address of the parameter InternetUrl.</p>
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// <p>Download address<br>Description: This download address is the same as the IntranetUrl download address.</p>
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// <p>Log specific type. Possible values: binlog - binary log</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>binlog file start time</p>
	BinlogStartTime *string `json:"BinlogStartTime,omitnil,omitempty" name:"BinlogStartTime"`

	// <p>binlog file expiration time</p>
	BinlogFinishTime *string `json:"BinlogFinishTime,omitnil,omitempty" name:"BinlogFinishTime"`

	// <p>Region where local binlog files are located</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Backup task status. Possible values: "SUCCESS": backup successful, "FAILED": backup FAILED, "RUNNING": backup in progress.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Detailed information of binlog offsite backup</p>
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitnil,omitempty" name:"RemoteInfo"`

	// <p>Storage method: 0 - regular storage, 1 - archive storage, 2 - standard storage, defaults to 0.</p>
	CosStorageType *int64 `json:"CosStorageType,omitnil,omitempty" name:"CosStorageType"`

	// <p>Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.</p>
	//
	// Deprecated: InstanceId is deprecated.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Backup completion progress</p>
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`
}

type Bucket struct {
	// None
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Number of occurrences of the key value
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type CdbRegionSellConf struct {
	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Area
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Whether it is a default region
	IsDefaultRegion *int64 `json:"IsDefaultRegion,omitnil,omitempty" name:"IsDefaultRegion"`

	// Region name
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The purchasable configuration in an AZ in a region
	RegionConfig []*CdbZoneSellConf `json:"RegionConfig,omitnil,omitempty" name:"RegionConfig"`
}

type CdbSellConfig struct {
	// Memory size in MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CPU core count
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Minimum disk size in GB
	VolumeMin *int64 `json:"VolumeMin,omitnil,omitempty" name:"VolumeMin"`

	// Maximum disk size in GB
	VolumeMax *int64 `json:"VolumeMax,omitnil,omitempty" name:"VolumeMax"`

	// Disk capacity increment in GB
	VolumeStep *int64 `json:"VolumeStep,omitnil,omitempty" name:"VolumeStep"`

	// IO operations per second
	Iops *int64 `json:"Iops,omitnil,omitempty" name:"Iops"`

	// Application scenario description
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// Status. The value `0` indicates that this specification is available.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance type, possible value ranges from UNIVERSAL (universal type), EXCLUSIVE (exclusive), BASIC (basic), to BASIC_V2 (basic v2).
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Engine type description. Valid values: `Innodb`, `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// Purchasable specifications ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type CdbSellType struct {
	// Purchasable instance name.
	// Z3: High-availability, corresponds to the specified specification DeviceType, including UNIVERSAL and EXCLUSIVE.
	// CVM: It is a basic edition type, and the DeviceType in the corresponding specifications is BASIC (Offline).
	// TKE: It is the basic version v2 type, and the DeviceType in the corresponding specifications is BASIC_V2.
	// CLOUD_NATIVE_CLUSTER: Represents the standard type of cloud disk edition.
	// CLOUD_NATIVE_CLUSTER_EXCLUSIVE: Indicates the enhanced cloud disk edition.
	// ECONOMICAL: Means economical.
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// Engine version number
	EngineVersion []*string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Purchasable specifications ID
	ConfigIds []*int64 `json:"ConfigIds,omitnil,omitempty" name:"ConfigIds"`
}

type CdbZoneDataResult struct {
	// List of purchasable specifications
	Configs []*CdbSellConfig `json:"Configs,omitnil,omitempty" name:"Configs"`

	// List of AZs in purchasable regions
	Regions []*CdbRegionSellConf `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type CdbZoneSellConf struct {
	// AZ status, which is used to indicate whether instances are purchasable. Valid values: `1` (purchasable), `3` (not purchasable), `4` (AZ not displayed)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Whether it is a custom instance type
	IsCustom *bool `json:"IsCustom,omitnil,omitempty" name:"IsCustom"`

	// Whether disaster recovery is supported
	IsSupportDr *bool `json:"IsSupportDr,omitnil,omitempty" name:"IsSupportDr"`

	// Whether VPC is supported
	IsSupportVpc *bool `json:"IsSupportVpc,omitnil,omitempty" name:"IsSupportVpc"`

	// Maximum purchasable quantity of hourly billed instances
	HourInstanceSaleMaxNum *int64 `json:"HourInstanceSaleMaxNum,omitnil,omitempty" name:"HourInstanceSaleMaxNum"`

	// Whether it is a default AZ
	IsDefaultZone *bool `json:"IsDefaultZone,omitnil,omitempty" name:"IsDefaultZone"`

	// Whether it is a BM zone
	IsBm *bool `json:"IsBm,omitnil,omitempty" name:"IsBm"`

	// Supported billing method. Valid values: `0` (yearly/monthly subscribed), `1` (hourly billed), `2` (pay-as-you-go)
	PayType []*string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// Data replication type. Valid values: `0` (async), `1` (semi-sync), `2` (strong sync)
	ProtectMode []*string `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// AZ name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Multi-AZ information
	ZoneConf *ZoneConf `json:"ZoneConf,omitnil,omitempty" name:"ZoneConf"`

	// Information of supported disaster recovery AZs
	DrZone []*string `json:"DrZone,omitnil,omitempty" name:"DrZone"`

	// Whether cross-AZ read-only access is supported
	IsSupportRemoteRo *bool `json:"IsSupportRemoteRo,omitnil,omitempty" name:"IsSupportRemoteRo"`

	// Information of supported cross-AZ read-only zone
	RemoteRoZone []*string `json:"RemoteRoZone,omitnil,omitempty" name:"RemoteRoZone"`

	// AZ status, which is used to indicate whether dedicated instances are purchasable. Valid values: `1 (purchasable), `3` (not purchasable), `4` (AZ not displayed)
	ExClusterStatus *int64 `json:"ExClusterStatus,omitnil,omitempty" name:"ExClusterStatus"`

	// Information of cross-AZ read-only zones supported by a dedicated instance
	ExClusterRemoteRoZone []*string `json:"ExClusterRemoteRoZone,omitnil,omitempty" name:"ExClusterRemoteRoZone"`

	// AZ information of a multi-AZ deployed dedicated instance.
	ExClusterZoneConf *ZoneConf `json:"ExClusterZoneConf,omitnil,omitempty" name:"ExClusterZoneConf"`

	// Array of purchasable instance types. The value of `configIds` and `configs` have a one-to-one correspondence.
	SellType []*CdbSellType `json:"SellType,omitnil,omitempty" name:"SellType"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Whether IPv6 is supported
	IsSupportIpv6 *bool `json:"IsSupportIpv6,omitnil,omitempty" name:"IsSupportIpv6"`

	// Supported engine types for purchasable database
	EngineType []*string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// Sales status of the cloud disk edition instance in the current availability zone. Possible returned values: 1-launched; 3-not available for sale; 4-not displayed.
	CloudNativeClusterStatus *int64 `json:"CloudNativeClusterStatus,omitnil,omitempty" name:"CloudNativeClusterStatus"`

	// Cloud disk edition or single-node basic edition supported disk type.
	DiskTypeConf []*DiskTypeConfigItem `json:"DiskTypeConf,omitnil,omitempty" name:"DiskTypeConf"`
}

type CloneItem struct {
	// ID of the original instance in a clone task
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// ID of the cloned instance in a clone task
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// Clone task ID
	CloneJobId *int64 `json:"CloneJobId,omitnil,omitempty" name:"CloneJobId"`

	// The policy used in a clone task. Valid values: `timepoint` (roll back to a specific point in time), `backupset` (roll back by using a specific backup file).
	RollbackStrategy *string `json:"RollbackStrategy,omitnil,omitempty" name:"RollbackStrategy"`

	// The point in time to which the cloned instance will be rolled back
	RollbackTargetTime *string `json:"RollbackTargetTime,omitnil,omitempty" name:"RollbackTargetTime"`

	// Task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task status. Valid values: `initial`, `running`, `wait_complete`, `success`, `failed`.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Clone instance region ID
	NewRegionId *int64 `json:"NewRegionId,omitnil,omitempty" name:"NewRegionId"`

	// Source instance region ID
	SrcRegionId *int64 `json:"SrcRegionId,omitnil,omitempty" name:"SrcRegionId"`
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseAuditServiceResponseParams `json:"Response"`
}

func (r *CloseAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCDBProxyRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`. Default value: `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

type CloseCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`. Default value: `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

func (r *CloseCDBProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCDBProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "OnlyCloseRW")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseCDBProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCDBProxyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseCDBProxyResponse struct {
	*tchttp.BaseResponse
	Response *CloseCDBProxyResponseParams `json:"Response"`
}

func (r *CloseCDBProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCDBProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCdbProxyAddressRequestParams struct {
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. You can obtain it through the API [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1).
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

type CloseCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. You can obtain it through the API [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1).
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

func (r *CloseCdbProxyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCdbProxyAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCdbProxyAddressResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseCdbProxyAddressResponse struct {
	*tchttp.BaseResponse
	Response *CloseCdbProxyAddressResponseParams `json:"Response"`
}

func (r *CloseCdbProxyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCdbProxyAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLRequestParams struct {
	// Instance ID. Required when the read-only group ID is empty. Can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only group ID. Required when the instance ID is empty. Can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Required when the read-only group ID is empty. Can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only group ID. Required when the instance ID is empty. Can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

func (r *CloseSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// Asynchronous request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseSSLResponse struct {
	*tchttp.BaseResponse
	Response *CloseSSLResponseParams `json:"Response"`
}

func (r *CloseSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanServiceRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. Use the query instance list API (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1) to obtain it, with its value being the InstanceId field in the output parameter. Input the read-only group ID to disable public network access for the read-only group.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// When updating the read-only group of a cloud disk edition instance, specify the instance ID in InstanceId and this parameter to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type CloseWanServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. Use the query instance list API (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1) to obtain it, with its value being the InstanceId field in the output parameter. Input the read-only group ID to disable public network access for the read-only group.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// When updating the read-only group of a cloud disk edition instance, specify the instance ID in InstanceId and this parameter to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

func (r *CloseWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWanServiceResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseWanServiceResponseParams `json:"Response"`
}

func (r *CloseWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWanServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {
	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node type: primary node and secondary node.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Region.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ClusterNodeInfo struct {
	// node id.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node role.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Node AZ.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Weight of the node
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Node status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ClusterTopology struct {
	// RW node topology.
	// Description: NodeId can be obtained through [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1).
	ReadWriteNode *ReadWriteNode `json:"ReadWriteNode,omitnil,omitempty" name:"ReadWriteNode"`

	// RO node topology.
	// Description: NodeId can be obtained through [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1).
	ReadOnlyNodes []*ReadonlyNode `json:"ReadOnlyNodes,omitnil,omitempty" name:"ReadOnlyNodes"`
}

type ColumnPrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Column name
	Column *string `json:"Column,omitnil,omitempty" name:"Column"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type CommonTimeWindow struct {
	// Time window on Monday in the format of 02:00-06:00
	Monday *string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// Time window on Tuesday in the format of 02:00-06:00
	Tuesday *string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// Time window on Wednesday in the format of 02:00-06:00
	Wednesday *string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// Time window on Thursday in the format of 02:00-06:00
	Thursday *string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// Time window on Friday in the format of 02:00-06:00
	Friday *string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// Time window on Saturday in the format of 02:00-06:00
	Saturday *string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// Time window on Sunday in the format of 02:00-06:00
	Sunday *string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// Non-archive backup retention policy. Valid values: `weekly` (back up by week), monthly (back up by month), default value: `weekly`.
	BackupPeriodStrategy *string `json:"BackupPeriodStrategy,omitnil,omitempty" name:"BackupPeriodStrategy"`

	// If `BackupPeriodStrategy` is `monthly`, you need to pass in the specific backup dates. The time interval between any two adjacent dates cannot exceed 2 days, for example, [1,4,7,9,11,14,17,19,22,25,28,30,31].
	Days []*int64 `json:"Days,omitnil,omitempty" name:"Days"`

	// Backup time by month in the format of 02:00–06:00, which is required when `BackupPeriodStrategy` is `monthly`.
	BackupPeriodTime *string `json:"BackupPeriodTime,omitnil,omitempty" name:"BackupPeriodTime"`
}

// Predefined struct for user
type CreateAccountsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of TencentDB accounts
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Password of the new account.
	// Note:
	// 1. Within 8–64 characters (recommend not exceeding 12).
	// 2. At least two of the following items: lowercase letter a – z or uppercase letter A – Z, digit 0 – 9, _+-,&=!@#$%^*().|.
	// 3. Cannot contain invalid characters.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Remark information. Input limit: 255 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Maximum connections of the new account. Default value: `10240`. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type CreateAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of TencentDB accounts
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Password of the new account.
	// Note:
	// 1. Within 8–64 characters (recommend not exceeding 12).
	// 2. At least two of the following items: lowercase letter a – z or uppercase letter A – Z, digit 0 – 9, _+-,&=!@#$%^*().|.
	// 3. Cannot contain invalid characters.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Remark information. Input limit: 255 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Maximum connections of the new account. Default value: `10240`. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

func (r *CreateAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountsResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountsResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountsResponseParams `json:"Response"`
}

func (r *CreateAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time. We recommend that the interval between start and end time does not exceed 7 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. We recommend that the interval between start and end time does not exceed 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sort order. Valid values: "ASC" - Ascending order, "DESC" - Descending order. Default value: "DESC".
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field to sort by. Valid values: "timestamp" - Timestamp; "affectRows" - Number of affected rows; "execTime" - Execution time. Default value: "timestamp".
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Deprecated.
	//
	// Deprecated: Filter is deprecated.
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Filter conditions. You can filter logs based on these conditions.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// Columns to include in the download.
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
}

type CreateAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time. We recommend that the interval between start and end time does not exceed 7 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. We recommend that the interval between start and end time does not exceed 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sort order. Valid values: "ASC" - Ascending order, "DESC" - Descending order. Default value: "DESC".
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field to sort by. Valid values: "timestamp" - Timestamp; "affectRows" - Number of affected rows; "execTime" - Execution time. Default value: "timestamp".
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Deprecated.
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Filter conditions. You can filter logs based on these conditions.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`

	// Columns to include in the download.
	ColumnFilter []*string `json:"ColumnFilter,omitnil,omitempty" name:"ColumnFilter"`
}

func (r *CreateAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "Filter")
	delete(f, "LogFilter")
	delete(f, "ColumnFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileResponseParams struct {
	// Audit log file name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditLogFileResponseParams `json:"Response"`
}

func (r *CreateAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditPolicyRequestParams struct {
	// Audit policy name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Audit rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period of audit logs. Valid values:
	// 7: seven days (a week);
	// 30: 30 days (a month);
	// 180: 180 days (six months);
	// 365: 365 days (a year);
	// 1095: 1095 days (three years);
	// 1825: 1825 days (five years).
	// This parameter specifies the retention period (30 days by default) of audit logs, which is valid when you create the first audit policy for an instance. If the instance already has audit policies, this parameter is invalid, but you can use the `ModifyAuditConfig` API to modify the retention period.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`
}

type CreateAuditPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Audit policy name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Audit rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period of audit logs. Valid values:
	// 7: seven days (a week);
	// 30: 30 days (a month);
	// 180: 180 days (six months);
	// 365: 365 days (a year);
	// 1095: 1095 days (three years);
	// 1825: 1825 days (five years).
	// This parameter specifies the retention period (30 days by default) of audit logs, which is valid when you create the first audit policy for an instance. If the instance already has audit policies, this parameter is invalid, but you can use the `ModifyAuditConfig` API to modify the retention period.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`
}

func (r *CreateAuditPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "RuleId")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditPolicyResponseParams struct {
	// Audit policy ID.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditPolicyResponseParams `json:"Response"`
}

func (r *CreateAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateRequestParams struct {
	// Audit rule.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template name. Up to 30 characters are allowed.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule template description. Up to 200 characters are allowed.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk. Default value: 1.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled. Default value: 0.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type CreateAuditRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Rule template name. Up to 30 characters are allowed.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule template description. Up to 200 characters are allowed.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk. Default value: 1.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled. Default value: 0.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

func (r *CreateAuditRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditRuleTemplateResponseParams struct {
	// Generated rule template ID.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuditRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditRuleTemplateResponseParams `json:"Response"`
}

func (r *CreateAuditRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target backup method. Valid values: `logical` (logical cold backup), `physical` (physical cold backup), `snapshot` (snapshot backup). Basic Edition instances only support snapshot backups.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Database table information to be backed up. If this parameter is not set, the whole instance is backed up by default. This parameter can only be set when BackupMethod=logical. The specified database and tables must exist. Otherwise, backup may fail.
	// If necessary to back up tables tb1 and tb2 in database db1 and database db2, configure the parameter as [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"}].
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitnil,omitempty" name:"BackupDBTableList"`

	// Manually back up the alias. Keep the input length within 60 characters.
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// Whether the physical backup needs encryption, optional values: on - yes, off - no. This value is meaningful only when BackupMethod is physical. If not specified, use the default encryption policy of instance backup. Here, the default encryption policy refers to the current instance encryption policy queried via the api for the query [DescribeBackupEncryptionStatus](https://www.tencentcloud.com/document/product/236/86508?from_cn_redirect=1).
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target backup method. Valid values: `logical` (logical cold backup), `physical` (physical cold backup), `snapshot` (snapshot backup). Basic Edition instances only support snapshot backups.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Database table information to be backed up. If this parameter is not set, the whole instance is backed up by default. This parameter can only be set when BackupMethod=logical. The specified database and tables must exist. Otherwise, backup may fail.
	// If necessary to back up tables tb1 and tb2 in database db1 and database db2, configure the parameter as [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"}].
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitnil,omitempty" name:"BackupDBTableList"`

	// Manually back up the alias. Keep the input length within 60 characters.
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// Whether the physical backup needs encryption, optional values: on - yes, off - no. This value is meaningful only when BackupMethod is physical. If not specified, use the default encryption policy of instance backup. Here, the default encryption policy refers to the current instance encryption policy queried via the api for the query [DescribeBackupEncryptionStatus](https://www.tencentcloud.com/document/product/236/86508?from_cn_redirect=1).
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupDBTableList")
	delete(f, "ManualBackupName")
	delete(f, "EncryptionFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// Backup task ID
	BackupId *uint64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupResponseParams `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyAddressRequestParams struct {
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group. Valid values: `true`, `false`.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// Delay removal threshold, minimum value: 1, range: 1–10000. The value is an integer.
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// Whether to enable failover. Valid values: `true`, `false`.
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add newly created read-only instances. Valid values: `true`, `false`.
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether it is read-only. Valid values: `true`, `false`.
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// Whether to enable transaction splitting. Valid values: `true`, `false`.
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Assignment of read/write weights
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// VPC ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Private subnet ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Whether to enable connection pool. Off by default.
	// Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be equal to or greater than MySQL 8.0 20230630.
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// IP. Leave it blank to default to a random supported IP in the selected VPC.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port. Default value 3306.
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Connection pool type. Available values: transaction (transaction-level connection pool), connection (session-level connection pool). This parameter is valid only when ConnectionPool is true. Default value: connection.
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether adaptive load balancing is enabled. Off by default.
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// Access mode. nearBy - proximity access, balance - balanced allocation. Default value: nearBy.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

type CreateCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group. Valid values: `true`, `false`.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// Delay removal threshold, minimum value: 1, range: 1–10000. The value is an integer.
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// Whether to enable failover. Valid values: `true`, `false`.
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add newly created read-only instances. Valid values: `true`, `false`.
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether it is read-only. Valid values: `true`, `false`.
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// Whether to enable transaction splitting. Valid values: `true`, `false`.
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Assignment of read/write weights
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// VPC ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Private subnet ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Whether to enable connection pool. Off by default.
	// Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be equal to or greater than MySQL 8.0 20230630.
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// IP. Leave it blank to default to a random supported IP in the selected VPC.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port. Default value 3306.
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Connection pool type. Available values: transaction (transaction-level connection pool), connection (session-level connection pool). This parameter is valid only when ConnectionPool is true. Default value: connection.
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	// Whether adaptive load balancing is enabled. Off by default.
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// Access mode. nearBy - proximity access, balance - balanced allocation. Default value: nearBy.
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

func (r *CreateCdbProxyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "WeightMode")
	delete(f, "IsKickOut")
	delete(f, "MinCount")
	delete(f, "MaxDelay")
	delete(f, "FailOver")
	delete(f, "AutoAddRo")
	delete(f, "ReadOnly")
	delete(f, "TransSplit")
	delete(f, "ProxyAllocation")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ConnectionPool")
	delete(f, "Desc")
	delete(f, "Vip")
	delete(f, "VPort")
	delete(f, "SecurityGroup")
	delete(f, "ConnectionPoolType")
	delete(f, "AutoLoadBalance")
	delete(f, "AccessMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyAddressResponseParams struct {
	// Asynchronous Task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCdbProxyAddressResponse struct {
	*tchttp.BaseResponse
	Response *CreateCdbProxyAddressResponseParams `json:"Response"`
}

func (r *CreateCdbProxyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VPC ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Private subnet ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Node specification configuration.
	// Parameter description in the example.
	// NodeCount: Number of nodes.
	// Region: Node region.
	// Zone: Node availability zone.
	// Cpu: Number of cores per proxy node (unit: core).
	// Mem: Memory size of each proxy node (unit: MB).
	// Remarks:
	// 1. Database proxy supported node specifications are: 2C4000MB, 4C8000MB, 8C16000MB.
	// 2. The above parameters (such as number of nodes, availability zone) are required. When calling the API, if incomplete, creation may fail.
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Connection pool threshold
	// Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be equal to or greater than MySQL 8.0 20230630.
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// Specify the Linux kernel version of the purchased proxy. Leave it blank to ship the latest version by default.
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`
}

type CreateCdbProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VPC ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Private subnet ID. Obtain through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Node specification configuration.
	// Parameter description in the example.
	// NodeCount: Number of nodes.
	// Region: Node region.
	// Zone: Node availability zone.
	// Cpu: Number of cores per proxy node (unit: core).
	// Mem: Memory size of each proxy node (unit: MB).
	// Remarks:
	// 1. Database proxy supported node specifications are: 2C4000MB, 4C8000MB, 8C16000MB.
	// 2. The above parameters (such as number of nodes, availability zone) are required. When calling the API, if incomplete, creation may fail.
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Connection pool threshold
	// Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be equal to or greater than MySQL 8.0 20230630.
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// Specify the Linux kernel version of the purchased proxy. Leave it blank to ship the latest version by default.
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`
}

func (r *CreateCdbProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProxyNodeCustom")
	delete(f, "SecurityGroup")
	delete(f, "Desc")
	delete(f, "ConnectionPoolLimit")
	delete(f, "ProxyVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCdbProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyResponseParams struct {
	// Asynchronous Task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCdbProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCdbProxyResponseParams `json:"Response"`
}

func (r *CreateCdbProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCdbProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceRequestParams struct {
	// <p>Clone source instance ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">DescribeDBInstances</a> API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>If necessary, specify this value when cloning an instance and rolling back to a specified time. The time format is yyyy-mm-dd hh:mm:ss.<br>Note: This parameter and the SpecifiedBackupId parameter require a choice between the two for configuration.</p>
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitnil,omitempty" name:"SpecifiedRollbackTime"`

	// <p>If necessary to clone an instance and roll back to a designated backup set, specify this value as the Id of the backup file. Please use <a href="/document/api/236/15842">query data backup file list</a>.</p><p>If it is a clone of a two-node, three-node, or four-node instance, the backup file is a physical backup. If it is a clone of a single-node or cloud disk edition instance, the backup file is a snapshot backup.</p>
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitnil,omitempty" name:"SpecifiedBackupId"`

	// <p>VPC ID. Please use <a href="/document/api/215/15778">Querying VPC List</a>.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID in the private network. If UniqVpcId is set up, UniqSubnetId is required. Please use <a href="/document/api/215/15784">query subnet list</a>.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Instance memory size, unit: MB, must not be less than the clone source instance. Default is same as the source instance.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB, must not be less than the clone source instance. Default is same as the source instance.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Name of the newly generated clone instance. Support input of up to 60 characters.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Security group parameters. Use the API <a href="https://www.tencentcloud.com/document/api/236/15850?from_cn_redirect=1">Query Project Security Group Information</a> to query security group details of a certain project.</p>
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// <p>Tag information of the instance.</p>
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Instance Cpu cores, must not be less than the clone source instance. Default is same as the source instance.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Multiple Availability Zones, defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>The AZ information of the newly generated clone instance standby 1 is the same as the source instance Zone by default.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>AZ information of standby 2, empty by default. Specify this parameter when you clone a strong sync primary instance.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Clone instance type. Supported values include: "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "CLOUD_NATIVE_CLUSTER" - standard type for CLOUD disk, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - enhanced type for CLOUD disk. If not specified, it defaults to general-purpose instance.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Number of nodes in the new clone instance.</p><p>To clone a three-node instance, set this value to 3 or specify the BackupZone parameter. To clone a dual-node instance, set this value to 2. By default, a dual-node instance is cloned. To clone a four-node instance, set this value to 4 or specify the FourthZone parameter.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Placement group ID.</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>Whether to only pre-check this request. true: Send a check request without creating an instance. Check items include required parameters, request format, and service limits. If the check fails, return the corresponding error code; if the check passes, return RequestId. Default false: Send a normal request and create the instance directly after passing the check.</p>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Financial Enclosure ID.</p>
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// <p>Project ID. Default project ID 0.</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Payment type. Valid values: PRE_PAID (prepaid, also known as yearly/monthly subscription) and USED_PAID (pay-as-you-go). Default billing mode is pay-as-you-go.</p>
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// <p>Instance duration, required when PayType is PRE_PAID, measurement unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Topology configuration for cloud disk edition nodes.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Original instance region. Required when importing a remote backup, for example: ap-guangzhou</p>
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// <p>Offsite data backup id</p>
	SpecifiedSubBackupId *int64 `json:"SpecifiedSubBackupId,omitnil,omitempty" name:"SpecifiedSubBackupId"`

	// <p>The AZ information of the newly generated clone instance primary database is the same as the source instance Zone by default.</p>
	//
	// Deprecated: MasterZone is deprecated.
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// <p>The AZ information of the newly generated clone instance's primary database defaults to the same as the source instance's Zone.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>AZ information of standby 3, empty by default. Specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Clone source instance ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">DescribeDBInstances</a> API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>If necessary, specify this value when cloning an instance and rolling back to a specified time. The time format is yyyy-mm-dd hh:mm:ss.<br>Note: This parameter and the SpecifiedBackupId parameter require a choice between the two for configuration.</p>
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitnil,omitempty" name:"SpecifiedRollbackTime"`

	// <p>If necessary to clone an instance and roll back to a designated backup set, specify this value as the Id of the backup file. Please use <a href="/document/api/236/15842">query data backup file list</a>.</p><p>If it is a clone of a two-node, three-node, or four-node instance, the backup file is a physical backup. If it is a clone of a single-node or cloud disk edition instance, the backup file is a snapshot backup.</p>
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitnil,omitempty" name:"SpecifiedBackupId"`

	// <p>VPC ID. Please use <a href="/document/api/215/15778">Querying VPC List</a>.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID in the private network. If UniqVpcId is set up, UniqSubnetId is required. Please use <a href="/document/api/215/15784">query subnet list</a>.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Instance memory size, unit: MB, must not be less than the clone source instance. Default is same as the source instance.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB, must not be less than the clone source instance. Default is same as the source instance.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Name of the newly generated clone instance. Support input of up to 60 characters.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Security group parameters. Use the API <a href="https://www.tencentcloud.com/document/api/236/15850?from_cn_redirect=1">Query Project Security Group Information</a> to query security group details of a certain project.</p>
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// <p>Tag information of the instance.</p>
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Instance Cpu cores, must not be less than the clone source instance. Default is same as the source instance.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Multiple Availability Zones, defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>The AZ information of the newly generated clone instance standby 1 is the same as the source instance Zone by default.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>AZ information of standby 2, empty by default. Specify this parameter when you clone a strong sync primary instance.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Clone instance type. Supported values include: "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "CLOUD_NATIVE_CLUSTER" - standard type for CLOUD disk, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - enhanced type for CLOUD disk. If not specified, it defaults to general-purpose instance.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Number of nodes in the new clone instance.</p><p>To clone a three-node instance, set this value to 3 or specify the BackupZone parameter. To clone a dual-node instance, set this value to 2. By default, a dual-node instance is cloned. To clone a four-node instance, set this value to 4 or specify the FourthZone parameter.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Placement group ID.</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>Whether to only pre-check this request. true: Send a check request without creating an instance. Check items include required parameters, request format, and service limits. If the check fails, return the corresponding error code; if the check passes, return RequestId. Default false: Send a normal request and create the instance directly after passing the check.</p>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Financial Enclosure ID.</p>
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// <p>Project ID. Default project ID 0.</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Payment type. Valid values: PRE_PAID (prepaid, also known as yearly/monthly subscription) and USED_PAID (pay-as-you-go). Default billing mode is pay-as-you-go.</p>
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// <p>Instance duration, required when PayType is PRE_PAID, measurement unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Topology configuration for cloud disk edition nodes.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Original instance region. Required when importing a remote backup, for example: ap-guangzhou</p>
	SrcRegion *string `json:"SrcRegion,omitnil,omitempty" name:"SrcRegion"`

	// <p>Offsite data backup id</p>
	SpecifiedSubBackupId *int64 `json:"SpecifiedSubBackupId,omitnil,omitempty" name:"SpecifiedSubBackupId"`

	// <p>The AZ information of the newly generated clone instance primary database is the same as the source instance Zone by default.</p>
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// <p>The AZ information of the newly generated clone instance's primary database defaults to the same as the source instance's Zone.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>AZ information of standby 3, empty by default. Specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

func (r *CreateCloneInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloneInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecifiedRollbackTime")
	delete(f, "SpecifiedBackupId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "InstanceName")
	delete(f, "SecurityGroup")
	delete(f, "ResourceTags")
	delete(f, "Cpu")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "BackupZone")
	delete(f, "DeviceType")
	delete(f, "InstanceNodes")
	delete(f, "DeployGroupId")
	delete(f, "DryRun")
	delete(f, "CageId")
	delete(f, "ProjectId")
	delete(f, "PayType")
	delete(f, "Period")
	delete(f, "ClusterTopology")
	delete(f, "SrcRegion")
	delete(f, "SpecifiedSubBackupId")
	delete(f, "MasterZone")
	delete(f, "Zone")
	delete(f, "FourthZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceResponseParams struct {
	// <p>Request ID of the asynchronous task. Use this ID to query the outcome of the async task.</p>
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloneInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloneInstanceResponseParams `json:"Response"`
}

func (r *CreateCloneInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloneInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBImportJobRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TencentDB username
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Filename. The file must be a .sql file uploaded to Tencent Cloud.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Password of a TencentDB instance user account
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Name of the target database. If this parameter is not passed in, no database is specified.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// URL of a .sql file stored in COS. Either `FileName` or `CosUrl` must be specified.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`
}

type CreateDBImportJobRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TencentDB username
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Filename. The file must be a .sql file uploaded to Tencent Cloud.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Password of a TencentDB instance user account
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Name of the target database. If this parameter is not passed in, no database is specified.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// URL of a .sql file stored in COS. Either `FileName` or `CosUrl` must be specified.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`
}

func (r *CreateDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBImportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "FileName")
	delete(f, "Password")
	delete(f, "DbName")
	delete(f, "CosUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBImportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBImportJobResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBImportJobResponseParams `json:"Response"`
}

func (r *CreateDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBImportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourRequestParams struct {
	// <p>Instance count. Default value is 1, minimum value 1, maximum value 100.</p>
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Instance memory size. Unit: MB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to obtain creatable memory specifications.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the creatable disk range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>MySQL version, including 5.5, 5.6, 5.7, and 8.0. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the instance version created.<br>Note: When creating a non-cloud disk edition instance, specify the instance version as needed (recommend 5.7 or 8.0). If this parameter is left blank, the default value is 8.0. If creating a cloud disk edition instance, this parameter can only be set to 5.7 or 8.0.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>VPC ID. Please use <a href="/document/api/215/15778">Querying VPC List</a>.<br>Description: If you create a cloud disk edition instance, this item is required and must be VPC type. If this item is left blank, the system will select the default VPC by default.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID in the private network. If UniqVpcId is set up, UniqSubnetId is required. Please use <a href="/document/api/215/15784">query subnet list</a>.<br>Description: If this item is not filled, the system will select the default subnet in the Default VPC.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Project ID. If this is left empty, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>For availability zone information, please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the availability zones where instances can be created.</p><p>If you create a single-node, two-node, three-node, or four-node instance, this parameter is required. Please specify an availability zone. If you do not specify one, the system will automatically select an availability zone (which may not be the one you want to deploy in). If you create a cloud disk edition instance, leave this parameter blank and configure the availability zones for read-write nodes and read-only nodes with parameter ClusterTopology.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance ID, required when you purchase a read-only instance or disaster recovery instance. This field represents the primary instance ID of the read-only instance or disaster recovery instance. Please use the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API to query the cloud database instance ID.</p>
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// <p>Instance type. Supported values include: master - primary instance, dr - disaster recovery instance, ro - read-only instance.<br>Description: Select instance type. master is selected by default if left blank.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Region of the primary instance. This field is required when you purchase a disaster recovery or RO instance.</p>
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// <p>Custom port. Supported range: [1024-65535].<br>Description: If left blank, it defaults to 3306.</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Set the root account password. The password must contain 8 to 64 characters and at least two of the following: letters, digits, or characters (supported characters: _+-&amp;=!@#$%^*()). You can specify this parameter when purchasing a primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Parameter list. The parameter format is ParamList.0.Name=auto_increment&amp;ParamList.0.Value=1.</p><p>Query the configurable parameters by referring to <a href="https://www.tencentcloud.com/document/api/236/32662?from_cn_redirect=1">querying the default configurable parameter list</a>.<br>Note: table Name case sensitivity can be enabled or disabled with the parameter lower_case_table_names. a parameter Value of 0 means enabling, and a Value of 1 means disabling. If not set, the default Value is 0. If you create a MySQL 8.0 edition instance, you need to set the lower_case_table_names parameter when creating the instance to enable or disable table Name case sensitivity. After the instance is created, the parameter cannot be modified, meaning table Name case sensitivity cannot be changed once created. Instances of other database versions support modifying the lower_case_table_names parameter after creation. For the function invocation method to set table Name case sensitivity when creating an instance, please see examples in this document.</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication. You can specify this parameter when purchasing a primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Multiple Availability Zones, defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone. Specify this parameter when purchasing the primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>AZ information of standby database 1.</p><p>For two-node, three-node, or four-node instances, specify this parameter. If not specified, it defaults to the Zone value. For cloud disk edition instances, this parameter is optional. Configure the AZ for read-write and read-only nodes with parameter ClusterTopology. Single-node instances are in a single availability zone, so no need to specify this parameter.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>AZ information of standby 2, empty by default.</p><p>Specify this parameter when you proceed to purchase a three-node primary instance or a four-node primary instance.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Security group parameters. Use the API <a href="https://www.tencentcloud.com/document/api/236/15850?from_cn_redirect=1">Query Project Security Group Information</a> to query security group details of a certain project.</p>
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// <p>Read-only instance information. This parameter is required when you purchase a read-only instance.</p>
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// <p>This field is meaningless for pay-as-you-go instances.</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Instance name. When you purchase multiple instances only once, suffix numbers are used for case-sensitive instance naming. For example, instanceName=db and goodsNum=3, the instance names are db1, db2, and db3 respectively.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag information of the instance.</p>
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Placement group ID.</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>String used to ensure request idempotency. This string is generated by the customer and must be unique between different requests within 48 hours, with a maximum value of 64 ASCII characters. If not specified, request idempotency cannot be guaranteed.</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>Instance isolation type. Supported values include: "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "BASIC_V2" - ONTKE single-node instance, "CLOUD_NATIVE_CLUSTER" - CLOUD disk edition standard type, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - CLOUD disk edition enhanced. If not specified, it defaults to general-purpose instance.<br>Description: If a CLOUD disk edition instance is created, this parameter is required.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Parameter template ID.<br>Remark: If you use a custom parameter template ID, you can input the custom parameter template ID. If you plan to use the default parameter template, the input ID is invalid and you need to set ParamTemplateType.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Array of alarm policy IDs. OriginId returned by the Tencent Cloud observability platform DescribeAlarmPolicy API.</p>
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Number of instance nodes.</p><p>For RO and basic edition instances, the value defaults to 1. To purchase a three-node instance, set this value to 3 or specify the BackupZone parameter. When purchasing a primary instance without specifying this parameter or the BackupZone parameter, the default value is 2, which means purchasing a dual-node instance. To purchase a four-node instance, set this value to 4 or specify the FourthZone parameter.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Number of Cpu cores of the instance.</p><p>When multiple Cpu configurations exist for the Memory specification (for example, 64000MB Memory corresponds to 8-core/16-core/32-core), the Cpu parameter must be provided.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Whether to automatically initiate disaster recovery sync. This parameter only takes effect when purchasing a disaster recovery instance. Available values are: 0 - Do not automatically initiate disaster recovery sync; 1 - Automatically initiate disaster recovery sync. The default is 0.</p>
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// <p>Financial Enclosure ID.</p>
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// <p>Default parameter template type. Supported values include "HIGH_STABILITY" - HIGH-STABILITY template, "HIGH_PERFORMANCE" - HIGH-PERFORMANCE template. Default value is "HIGH_STABILITY".<br>Remark: If you need to use the cloud database MySQL default parameter template, set up ParamTemplateType.</p>
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// <p>Alarm policy name array, such as ["policy-uyoee9wg"]. This parameter is invalid when AlarmPolicyList is not empty.</p>
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// <p>Whether to only pre-check this request. true: Send a check request without creating an instance. Check items include required parameters, request format, and service limits. If the check fails, return the corresponding error code; if the check passes, return RequestId. Default false: Send a normal request and create the instance directly after passing the check.</p>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Instance engine type, defaults to "InnoDB". Supported values include "InnoDB" and "RocksDB".</p>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// <p>Specify the IP list of the instance. Only the primary instance is supported. Process by instance sequence. Handle as unspecified if insufficient.</p>
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// <p>The data protection space size of the cloud disk edition instance, in GB, has a setting range of 1 - 10.</p>
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// <p>Topology configuration for cloud disk edition nodes.<br>Description: If a cloud disk edition instance is purchased, this parameter is required. Set the topology for RW and RO nodes of the cloud disk edition instance. The node scope for RO nodes is 1-5. Set at least 1 RO node.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Disk type. This parameter can be specified for single-node (cloud disk) or cloud disk edition instances. CLOUD_SSD means SSD Cloud Block Storage, CLOUD_HSSD means enhanced SSD cloud disk, and CLOUD_PREMIUM means high-performance cloud block storage.<br>Note: The supported regions for disk types of single-node (cloud disk) and cloud disk edition instances vary slightly. For specific support situation, refer to <a href="https://www.tencentcloud.com/document/product/236/8458?from_cn_redirect=1">Regions and Availability Zones</a>.</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>ClusterType: cage—Financial Enclosure, cdc—CDB ON CDC; dedicate—dedicated cluster</p>
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// <p>Turn on or off instance destruction protection. on - turn on, off - turn off.</p>
	DestroyProtect *string `json:"DestroyProtect,omitnil,omitempty" name:"DestroyProtect"`

	// <p>AZ information of standby 3, empty by default. Specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance count. Default value is 1, minimum value 1, maximum value 100.</p>
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Instance memory size. Unit: MB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to obtain creatable memory specifications.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the creatable disk range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>MySQL version, including 5.5, 5.6, 5.7, and 8.0. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the instance version created.<br>Note: When creating a non-cloud disk edition instance, specify the instance version as needed (recommend 5.7 or 8.0). If this parameter is left blank, the default value is 8.0. If creating a cloud disk edition instance, this parameter can only be set to 5.7 or 8.0.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>VPC ID. Please use <a href="/document/api/215/15778">Querying VPC List</a>.<br>Description: If you create a cloud disk edition instance, this item is required and must be VPC type. If this item is left blank, the system will select the default VPC by default.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID in the private network. If UniqVpcId is set up, UniqSubnetId is required. Please use <a href="/document/api/215/15784">query subnet list</a>.<br>Description: If this item is not filled, the system will select the default subnet in the Default VPC.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Project ID. If this is left empty, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>For availability zone information, please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the availability zones where instances can be created.</p><p>If you create a single-node, two-node, three-node, or four-node instance, this parameter is required. Please specify an availability zone. If you do not specify one, the system will automatically select an availability zone (which may not be the one you want to deploy in). If you create a cloud disk edition instance, leave this parameter blank and configure the availability zones for read-write nodes and read-only nodes with parameter ClusterTopology.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance ID, required when you purchase a read-only instance or disaster recovery instance. This field represents the primary instance ID of the read-only instance or disaster recovery instance. Please use the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API to query the cloud database instance ID.</p>
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// <p>Instance type. Supported values include: master - primary instance, dr - disaster recovery instance, ro - read-only instance.<br>Description: Select instance type. master is selected by default if left blank.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Region of the primary instance. This field is required when you purchase a disaster recovery or RO instance.</p>
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// <p>Custom port. Supported range: [1024-65535].<br>Description: If left blank, it defaults to 3306.</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Set the root account password. The password must contain 8 to 64 characters and at least two of the following: letters, digits, or characters (supported characters: _+-&amp;=!@#$%^*()). You can specify this parameter when purchasing a primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Parameter list. The parameter format is ParamList.0.Name=auto_increment&amp;ParamList.0.Value=1.</p><p>Query the configurable parameters by referring to <a href="https://www.tencentcloud.com/document/api/236/32662?from_cn_redirect=1">querying the default configurable parameter list</a>.<br>Note: table Name case sensitivity can be enabled or disabled with the parameter lower_case_table_names. a parameter Value of 0 means enabling, and a Value of 1 means disabling. If not set, the default Value is 0. If you create a MySQL 8.0 edition instance, you need to set the lower_case_table_names parameter when creating the instance to enable or disable table Name case sensitivity. After the instance is created, the parameter cannot be modified, meaning table Name case sensitivity cannot be changed once created. Instances of other database versions support modifying the lower_case_table_names parameter after creation. For the function invocation method to set table Name case sensitivity when creating an instance, please see examples in this document.</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication. You can specify this parameter when purchasing a primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Multiple Availability Zones, defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone. Specify this parameter when purchasing the primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>AZ information of standby database 1.</p><p>For two-node, three-node, or four-node instances, specify this parameter. If not specified, it defaults to the Zone value. For cloud disk edition instances, this parameter is optional. Configure the AZ for read-write and read-only nodes with parameter ClusterTopology. Single-node instances are in a single availability zone, so no need to specify this parameter.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>AZ information of standby 2, empty by default.</p><p>Specify this parameter when you proceed to purchase a three-node primary instance or a four-node primary instance.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Security group parameters. Use the API <a href="https://www.tencentcloud.com/document/api/236/15850?from_cn_redirect=1">Query Project Security Group Information</a> to query security group details of a certain project.</p>
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// <p>Read-only instance information. This parameter is required when you purchase a read-only instance.</p>
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// <p>This field is meaningless for pay-as-you-go instances.</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Instance name. When you purchase multiple instances only once, suffix numbers are used for case-sensitive instance naming. For example, instanceName=db and goodsNum=3, the instance names are db1, db2, and db3 respectively.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag information of the instance.</p>
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Placement group ID.</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>String used to ensure request idempotency. This string is generated by the customer and must be unique between different requests within 48 hours, with a maximum value of 64 ASCII characters. If not specified, request idempotency cannot be guaranteed.</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>Instance isolation type. Supported values include: "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "BASIC_V2" - ONTKE single-node instance, "CLOUD_NATIVE_CLUSTER" - CLOUD disk edition standard type, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - CLOUD disk edition enhanced. If not specified, it defaults to general-purpose instance.<br>Description: If a CLOUD disk edition instance is created, this parameter is required.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Parameter template ID.<br>Remark: If you use a custom parameter template ID, you can input the custom parameter template ID. If you plan to use the default parameter template, the input ID is invalid and you need to set ParamTemplateType.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Array of alarm policy IDs. OriginId returned by the Tencent Cloud observability platform DescribeAlarmPolicy API.</p>
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Number of instance nodes.</p><p>For RO and basic edition instances, the value defaults to 1. To purchase a three-node instance, set this value to 3 or specify the BackupZone parameter. When purchasing a primary instance without specifying this parameter or the BackupZone parameter, the default value is 2, which means purchasing a dual-node instance. To purchase a four-node instance, set this value to 4 or specify the FourthZone parameter.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Number of Cpu cores of the instance.</p><p>When multiple Cpu configurations exist for the Memory specification (for example, 64000MB Memory corresponds to 8-core/16-core/32-core), the Cpu parameter must be provided.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Whether to automatically initiate disaster recovery sync. This parameter only takes effect when purchasing a disaster recovery instance. Available values are: 0 - Do not automatically initiate disaster recovery sync; 1 - Automatically initiate disaster recovery sync. The default is 0.</p>
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// <p>Financial Enclosure ID.</p>
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// <p>Default parameter template type. Supported values include "HIGH_STABILITY" - HIGH-STABILITY template, "HIGH_PERFORMANCE" - HIGH-PERFORMANCE template. Default value is "HIGH_STABILITY".<br>Remark: If you need to use the cloud database MySQL default parameter template, set up ParamTemplateType.</p>
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// <p>Alarm policy name array, such as ["policy-uyoee9wg"]. This parameter is invalid when AlarmPolicyList is not empty.</p>
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// <p>Whether to only pre-check this request. true: Send a check request without creating an instance. Check items include required parameters, request format, and service limits. If the check fails, return the corresponding error code; if the check passes, return RequestId. Default false: Send a normal request and create the instance directly after passing the check.</p>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Instance engine type, defaults to "InnoDB". Supported values include "InnoDB" and "RocksDB".</p>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// <p>Specify the IP list of the instance. Only the primary instance is supported. Process by instance sequence. Handle as unspecified if insufficient.</p>
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// <p>The data protection space size of the cloud disk edition instance, in GB, has a setting range of 1 - 10.</p>
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// <p>Topology configuration for cloud disk edition nodes.<br>Description: If a cloud disk edition instance is purchased, this parameter is required. Set the topology for RW and RO nodes of the cloud disk edition instance. The node scope for RO nodes is 1-5. Set at least 1 RO node.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Disk type. This parameter can be specified for single-node (cloud disk) or cloud disk edition instances. CLOUD_SSD means SSD Cloud Block Storage, CLOUD_HSSD means enhanced SSD cloud disk, and CLOUD_PREMIUM means high-performance cloud block storage.<br>Note: The supported regions for disk types of single-node (cloud disk) and cloud disk edition instances vary slightly. For specific support situation, refer to <a href="https://www.tencentcloud.com/document/product/236/8458?from_cn_redirect=1">Regions and Availability Zones</a>.</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>ClusterType: cage—Financial Enclosure, cdc—CDB ON CDC; dedicate—dedicated cluster</p>
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// <p>Turn on or off instance destruction protection. on - turn on, off - turn off.</p>
	DestroyProtect *string `json:"DestroyProtect,omitnil,omitempty" name:"DestroyProtect"`

	// <p>AZ information of standby 3, empty by default. Specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "EngineVersion")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProjectId")
	delete(f, "Zone")
	delete(f, "MasterInstanceId")
	delete(f, "InstanceRole")
	delete(f, "MasterRegion")
	delete(f, "Port")
	delete(f, "Password")
	delete(f, "ParamList")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "BackupZone")
	delete(f, "SecurityGroup")
	delete(f, "RoGroup")
	delete(f, "AutoRenewFlag")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "DeployGroupId")
	delete(f, "ClientToken")
	delete(f, "DeviceType")
	delete(f, "ParamTemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "AutoSyncFlag")
	delete(f, "CageId")
	delete(f, "ParamTemplateType")
	delete(f, "AlarmPolicyIdList")
	delete(f, "DryRun")
	delete(f, "EngineType")
	delete(f, "Vips")
	delete(f, "DataProtectVolume")
	delete(f, "ClusterTopology")
	delete(f, "DiskType")
	delete(f, "ClusterType")
	delete(f, "DestroyProtect")
	delete(f, "FourthZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// <p>Short order ID.</p>
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// <p>Instance ID list.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceHourResponseParams `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
	// <p>Instance memory size. Unit: MB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to obtain creatable memory specifications.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the creatable disk range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Instance duration, measurement unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Instance count. Default value is 1, minimum value 1, maximum value 100.</p>
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>For availability zone information, please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the availability zones where instances can be created.</p><p>If you create a single-node, two-node, three-node, or four-node instance, this parameter is required. Please specify an availability zone. If you do not specify one, the system will automatically select an availability zone (which may not be the one you want to deploy in). If you create a cloud disk edition instance, leave this parameter blank and configure the availability zones for read-write nodes and read-only nodes with parameter ClusterTopology.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>VPC ID. Please use <a href="/document/api/215/15778">Querying VPC List</a>.<br>Description: If you create a cloud disk edition instance, this item is required and must be VPC type. If this item is left blank, the system will select the default VPC by default.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID in the private network. If UniqVpcId is set up, UniqSubnetId is required. Please use <a href="/document/api/215/15784">query subnet list</a>.<br>Description: If this item is not filled, the system will select the default subnet in the Default VPC.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Project ID. The default project is used if left empty. When you purchase a read-only instance or disaster recovery instance, the project ID is consistent with the primary instance by default.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Custom port. Supported range: [1024-65535].<br>Description: If this item is left blank, it defaults to 3306.</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Instance type. Supported values include: master - primary instance, dr - disaster recovery instance, ro - read-only instance.<br>Description: Select instance type. master is selected by default if left blank.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Instance ID, required when purchasing a read-only instance or disaster recovery instance. This field represents the primary instance ID of the read-only instance or disaster recovery instance. Please use the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API to query the cloud database instance ID.</p>
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// <p>MySQL version, including 5.5, 5.6, 5.7, and 8.0. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the instance version created.<br>Note: When creating a non-cloud disk edition instance, specify the instance version as needed (recommend 5.7 or 8.0). If this parameter is left blank, the default value is 8.0. If creating a cloud disk edition instance, this parameter can only be set to 5.7 or 8.0.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Set the root account password. The password must contain 8 to 64 characters and at least two of the following: letters, digits, or characters (supported characters: _+-&amp;=!@#$%^*()). You can specify this parameter when purchasing a primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Multiple Availability Zones, defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>AZ information of standby database 1.</p><p>For two-node, three-node, or four-node instances, specify this parameter. If not specified, it defaults to the Zone value. For cloud disk edition instances, this parameter is optional. Configure the AZ for read-write and read-only nodes with parameter ClusterTopology. Single-node instances are in a single availability zone, so no need to specify this parameter.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>Parameter list. The parameter format is ParamList.0.Name=auto_increment&amp;ParamList.0.Value=1. You can query the configurable parameters by referring to <a href="https://www.tencentcloud.com/document/api/236/32662?from_cn_redirect=1">querying the default configurable parameter list</a>.<br>Description: table Name case sensitivity can be turned on or off by setting the parameter lower_case_table_names. a parameter Value of 0 means enabling, and a Value of 1 means disabling. If not set, the default Value is 0. If you create a MySQL 8.0 edition instance, you need to set the lower_case_table_names parameter when creating the instance to turn on or off table Name case sensitivity. After the instance is created, the parameter cannot be modified, meaning table Name case sensitivity cannot be changed once created. Instances of other database versions support modifying the lower_case_table_names parameter after creation. For the function invocation method to set table Name case sensitivity when creating an instance, please see example 3 in this document.</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>AZ information of standby 2, empty by default.</p><p>Specify this parameter when you proceed to purchase a three-node primary instance or a four-node primary instance.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Auto-renewal flag. Available values are: 0 - no auto-renewal; 1 - auto-renewal. Default is 0.</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Region of the primary instance. This field is required when you purchase a disaster recovery or RO instance.</p>
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// <p>Security group parameters. Use the API <a href="https://www.tencentcloud.com/document/api/236/15850?from_cn_redirect=1">Query Project Security Group Information</a> to query security group details of a certain project.</p>
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// <p>Read-only instance parameter. This parameter is required when you purchase a read-only instance.</p>
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// <p>Instance name. When you purchase multiple instances only once, suffix numbers are used for case-sensitive instance naming. For example, instnaceName=db and goodsNum=3, the instance names are db1, db2, and db3 respectively.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag information of the instance.</p>
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Placement group ID.</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>String used to ensure request idempotency. This string is generated by the customer and must be unique between different requests within 48 hours, with a maximum value of 64 ASCII characters. If not specified, request idempotency cannot be guaranteed.</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>Instance isolation type. Supported values include: "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "BASIC_V2" - ONTKE single-node instance, "CLOUD_NATIVE_CLUSTER" - CLOUD disk edition standard type, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - CLOUD disk edition enhanced. If not specified, it defaults to general-purpose instance.<br>Description: If a CLOUD disk edition instance is created, this parameter is required.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Parameter template ID.<br>Remark: If you use a custom parameter template ID, you can input the custom parameter template ID. If you plan to use the default parameter template, the input ID is invalid and you need to set ParamTemplateType.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Array of alarm policy IDs. OriginId returned by the Tencent Cloud observability platform DescribeAlarmPolicy API.</p>
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Number of instance nodes.</p><p>For RO and basic edition instances, the default value is 1. To purchase a three-node instance, set this value to 3 or specify the BackupZone parameter. When purchasing a primary instance without specifying this parameter or the BackupZone parameter, the default is 2, meaning a dual-node instance will be purchased. To purchase a four-node instance, set this value to 4 or specify the FourthZone parameter.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Number of Cpu cores of the instance.</p><p>When multiple Cpu configurations exist for the Memory specification (for example, 64000MB Memory corresponds to 8-core/16-core/32-core), the Cpu parameter must be provided.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Whether to automatically initiate disaster recovery sync. This parameter only takes effect when purchasing a disaster recovery instance. Available values are: 0 - Do not automatically initiate disaster recovery sync; 1 - Automatically initiate disaster recovery sync. The default is 0.</p>
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// <p>Financial Enclosure ID.</p>
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// <p>Default parameter template type. Supported values include "HIGH_STABILITY" - HIGH-STABILITY template, "HIGH_PERFORMANCE" - HIGH-PERFORMANCE template.<br>Remark: If you need to use TencentDB for MySQL default parameter template, set up ParamTemplateType.</p>
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// <p>Alarm policy name array, such as ["policy-uyoee9wg"]. This parameter is invalid when AlarmPolicyList is not empty.</p>
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// <p>Whether to perform a pre-check only for this request. true: Send a check request without creating an instance. Check items include whether required parameters are filled, request format, and service limit. If the check fails, return the corresponding error code; if the check passes, return RequestId. false: Send a normal request and create an instance directly after the check passes.<br>Default to false.</p>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Instance engine type, defaults to "InnoDB". Supported values include "InnoDB" and "RocksDB".</p>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// <p>Specify the IP list of the instance. Only the primary instance is supported. Process by instance sequence. Handle as unspecified if insufficient.</p>
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// <p>The data protection space size of the cloud disk edition instance, in GB, has a setting range of 1 - 10.</p>
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// <p>Topology configuration for cloud disk edition nodes.<br>Description: If a cloud disk edition instance is purchased, this parameter is required. Set the topology for RW and RO nodes of the cloud disk edition instance. The node scope for RO nodes is 1-5. Set at least 1 RO node.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Disk type. This parameter can be specified for single-node (cloud disk edition) or cloud disk edition instances. CLOUD_SSD means SSD Cloud Block Storage, CLOUD_HSSD means enhanced SSD cloud disk, and CLOUD_PREMIUM means high-performance cloud block storage.<br>Note: The supported regions for hard disk types of single-node (cloud disk edition) and cloud disk edition instances vary slightly. For specific support situation, refer to <a href="https://www.tencentcloud.com/document/product/236/8458?from_cn_redirect=1">Regions and Availability Zones</a>.</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Turn on or off instance destruction protection. on - turn on, off - turn off.</p>
	DestroyProtect *string `json:"DestroyProtect,omitnil,omitempty" name:"DestroyProtect"`

	// <p>AZ information of standby 3, empty by default. Specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance memory size. Unit: MB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to obtain creatable memory specifications.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the creatable disk range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Instance duration, measurement unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Instance count. Default value is 1, minimum value 1, maximum value 100.</p>
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>For availability zone information, please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the availability zones where instances can be created.</p><p>If you create a single-node, two-node, three-node, or four-node instance, this parameter is required. Please specify an availability zone. If you do not specify one, the system will automatically select an availability zone (which may not be the one you want to deploy in). If you create a cloud disk edition instance, leave this parameter blank and configure the availability zones for read-write nodes and read-only nodes with parameter ClusterTopology.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>VPC ID. Please use <a href="/document/api/215/15778">Querying VPC List</a>.<br>Description: If you create a cloud disk edition instance, this item is required and must be VPC type. If this item is left blank, the system will select the default VPC by default.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID in the private network. If UniqVpcId is set up, UniqSubnetId is required. Please use <a href="/document/api/215/15784">query subnet list</a>.<br>Description: If this item is not filled, the system will select the default subnet in the Default VPC.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Project ID. The default project is used if left empty. When you purchase a read-only instance or disaster recovery instance, the project ID is consistent with the primary instance by default.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Custom port. Supported range: [1024-65535].<br>Description: If this item is left blank, it defaults to 3306.</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Instance type. Supported values include: master - primary instance, dr - disaster recovery instance, ro - read-only instance.<br>Description: Select instance type. master is selected by default if left blank.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Instance ID, required when purchasing a read-only instance or disaster recovery instance. This field represents the primary instance ID of the read-only instance or disaster recovery instance. Please use the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API to query the cloud database instance ID.</p>
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// <p>MySQL version, including 5.5, 5.6, 5.7, and 8.0. Please use the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the instance version created.<br>Note: When creating a non-cloud disk edition instance, specify the instance version as needed (recommend 5.7 or 8.0). If this parameter is left blank, the default value is 8.0. If creating a cloud disk edition instance, this parameter can only be set to 5.7 or 8.0.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Set the root account password. The password must contain 8 to 64 characters and at least two of the following: letters, digits, or characters (supported characters: _+-&amp;=!@#$%^*()). You can specify this parameter when purchasing a primary instance. This parameter is invalid when purchasing a read-only instance or disaster recovery instance.</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Multiple Availability Zones, defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>AZ information of standby database 1.</p><p>For two-node, three-node, or four-node instances, specify this parameter. If not specified, it defaults to the Zone value. For cloud disk edition instances, this parameter is optional. Configure the AZ for read-write and read-only nodes with parameter ClusterTopology. Single-node instances are in a single availability zone, so no need to specify this parameter.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>Parameter list. The parameter format is ParamList.0.Name=auto_increment&amp;ParamList.0.Value=1. You can query the configurable parameters by referring to <a href="https://www.tencentcloud.com/document/api/236/32662?from_cn_redirect=1">querying the default configurable parameter list</a>.<br>Description: table Name case sensitivity can be turned on or off by setting the parameter lower_case_table_names. a parameter Value of 0 means enabling, and a Value of 1 means disabling. If not set, the default Value is 0. If you create a MySQL 8.0 edition instance, you need to set the lower_case_table_names parameter when creating the instance to turn on or off table Name case sensitivity. After the instance is created, the parameter cannot be modified, meaning table Name case sensitivity cannot be changed once created. Instances of other database versions support modifying the lower_case_table_names parameter after creation. For the function invocation method to set table Name case sensitivity when creating an instance, please see example 3 in this document.</p>
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>AZ information of standby 2, empty by default.</p><p>Specify this parameter when you proceed to purchase a three-node primary instance or a four-node primary instance.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Auto-renewal flag. Available values are: 0 - no auto-renewal; 1 - auto-renewal. Default is 0.</p>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// <p>Region of the primary instance. This field is required when you purchase a disaster recovery or RO instance.</p>
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// <p>Security group parameters. Use the API <a href="https://www.tencentcloud.com/document/api/236/15850?from_cn_redirect=1">Query Project Security Group Information</a> to query security group details of a certain project.</p>
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// <p>Read-only instance parameter. This parameter is required when you purchase a read-only instance.</p>
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// <p>Instance name. When you purchase multiple instances only once, suffix numbers are used for case-sensitive instance naming. For example, instnaceName=db and goodsNum=3, the instance names are db1, db2, and db3 respectively.</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Tag information of the instance.</p>
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// <p>Placement group ID.</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>String used to ensure request idempotency. This string is generated by the customer and must be unique between different requests within 48 hours, with a maximum value of 64 ASCII characters. If not specified, request idempotency cannot be guaranteed.</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>Instance isolation type. Supported values include: "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "BASIC_V2" - ONTKE single-node instance, "CLOUD_NATIVE_CLUSTER" - CLOUD disk edition standard type, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - CLOUD disk edition enhanced. If not specified, it defaults to general-purpose instance.<br>Description: If a CLOUD disk edition instance is created, this parameter is required.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Parameter template ID.<br>Remark: If you use a custom parameter template ID, you can input the custom parameter template ID. If you plan to use the default parameter template, the input ID is invalid and you need to set ParamTemplateType.</p>
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>Array of alarm policy IDs. OriginId returned by the Tencent Cloud observability platform DescribeAlarmPolicy API.</p>
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// <p>Number of instance nodes.</p><p>For RO and basic edition instances, the default value is 1. To purchase a three-node instance, set this value to 3 or specify the BackupZone parameter. When purchasing a primary instance without specifying this parameter or the BackupZone parameter, the default is 2, meaning a dual-node instance will be purchased. To purchase a four-node instance, set this value to 4 or specify the FourthZone parameter.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Number of Cpu cores of the instance.</p><p>When multiple Cpu configurations exist for the Memory specification (for example, 64000MB Memory corresponds to 8-core/16-core/32-core), the Cpu parameter must be provided.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Whether to automatically initiate disaster recovery sync. This parameter only takes effect when purchasing a disaster recovery instance. Available values are: 0 - Do not automatically initiate disaster recovery sync; 1 - Automatically initiate disaster recovery sync. The default is 0.</p>
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// <p>Financial Enclosure ID.</p>
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// <p>Default parameter template type. Supported values include "HIGH_STABILITY" - HIGH-STABILITY template, "HIGH_PERFORMANCE" - HIGH-PERFORMANCE template.<br>Remark: If you need to use TencentDB for MySQL default parameter template, set up ParamTemplateType.</p>
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// <p>Alarm policy name array, such as ["policy-uyoee9wg"]. This parameter is invalid when AlarmPolicyList is not empty.</p>
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// <p>Whether to perform a pre-check only for this request. true: Send a check request without creating an instance. Check items include whether required parameters are filled, request format, and service limit. If the check fails, return the corresponding error code; if the check passes, return RequestId. false: Send a normal request and create an instance directly after the check passes.<br>Default to false.</p>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>Instance engine type, defaults to "InnoDB". Supported values include "InnoDB" and "RocksDB".</p>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// <p>Specify the IP list of the instance. Only the primary instance is supported. Process by instance sequence. Handle as unspecified if insufficient.</p>
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// <p>The data protection space size of the cloud disk edition instance, in GB, has a setting range of 1 - 10.</p>
	DataProtectVolume *int64 `json:"DataProtectVolume,omitnil,omitempty" name:"DataProtectVolume"`

	// <p>Topology configuration for cloud disk edition nodes.<br>Description: If a cloud disk edition instance is purchased, this parameter is required. Set the topology for RW and RO nodes of the cloud disk edition instance. The node scope for RO nodes is 1-5. Set at least 1 RO node.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Disk type. This parameter can be specified for single-node (cloud disk edition) or cloud disk edition instances. CLOUD_SSD means SSD Cloud Block Storage, CLOUD_HSSD means enhanced SSD cloud disk, and CLOUD_PREMIUM means high-performance cloud block storage.<br>Note: The supported regions for hard disk types of single-node (cloud disk edition) and cloud disk edition instances vary slightly. For specific support situation, refer to <a href="https://www.tencentcloud.com/document/product/236/8458?from_cn_redirect=1">Regions and Availability Zones</a>.</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Turn on or off instance destruction protection. on - turn on, off - turn off.</p>
	DestroyProtect *string `json:"DestroyProtect,omitnil,omitempty" name:"DestroyProtect"`

	// <p>AZ information of standby 3, empty by default. Specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "Period")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ProjectId")
	delete(f, "Port")
	delete(f, "InstanceRole")
	delete(f, "MasterInstanceId")
	delete(f, "EngineVersion")
	delete(f, "Password")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "ParamList")
	delete(f, "BackupZone")
	delete(f, "AutoRenewFlag")
	delete(f, "MasterRegion")
	delete(f, "SecurityGroup")
	delete(f, "RoGroup")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "DeployGroupId")
	delete(f, "ClientToken")
	delete(f, "DeviceType")
	delete(f, "ParamTemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "AutoSyncFlag")
	delete(f, "CageId")
	delete(f, "ParamTemplateType")
	delete(f, "AlarmPolicyIdList")
	delete(f, "DryRun")
	delete(f, "EngineType")
	delete(f, "Vips")
	delete(f, "DataProtectVolume")
	delete(f, "ClusterTopology")
	delete(f, "DiskType")
	delete(f, "DestroyProtect")
	delete(f, "FourthZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// <p>Billing sub-order ID.</p>
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// <p>Instance ID list.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseRequestParams struct {
	// Instance ID in the format of `cdb-c1nl9rpv`,  which is the same as the one displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name, length not exceeding 64.
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Character set. Valid values:  `utf8`, `gbk`, `latin1`, `utf8mb4`.
	CharacterSetName *string `json:"CharacterSetName,omitnil,omitempty" name:"CharacterSetName"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `cdb-c1nl9rpv`,  which is the same as the one displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name, length not exceeding 64.
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Character set. Valid values:  `utf8`, `gbk`, `latin1`, `utf8mb4`.
	CharacterSetName *string `json:"CharacterSetName,omitnil,omitempty" name:"CharacterSetName"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBName")
	delete(f, "CharacterSetName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatabaseResponseParams `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// Parameter template name. Up to 60 characters are allowed.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MySQL version number. Available values: 5.6, 5.7, and 8.0.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Source parameter template ID, which can be obtained through the [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1) API.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Instance engine type, defaults to "InnoDB". Supported values include "InnoDB" and "RocksDB".
	// Description: RocksDB is only supported in database versions MySQL 5.7 and MySQL 8.0.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template name. Up to 60 characters are allowed.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MySQL version number. Available values: 5.6, 5.7, and 8.0.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Source parameter template ID, which can be obtained through the [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1) API.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Instance engine type, defaults to "InnoDB". Supported values include "InnoDB" and "RocksDB".
	// Description: RocksDB is only supported in database versions MySQL 5.7 and MySQL 8.0.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "EngineVersion")
	delete(f, "TemplateId")
	delete(f, "ParamList")
	delete(f, "TemplateType")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParamTemplateResponseParams `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoInstanceIpRequestParams struct {
	// Read-only instance ID in the format of "cdbro-3i70uj0k". Its value is the same as the read-only instance ID in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Subnet descriptor, such as "subnet-1typ0s7d".
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// VPC descriptor, such as "vpc-a23yt67j". If this field is passed in, `UniqSubnetId` will be required.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`
}

type CreateRoInstanceIpRequest struct {
	*tchttp.BaseRequest
	
	// Read-only instance ID in the format of "cdbro-3i70uj0k". Its value is the same as the read-only instance ID in the TencentDB Console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Subnet descriptor, such as "subnet-1typ0s7d".
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// VPC descriptor, such as "vpc-a23yt67j". If this field is passed in, `UniqSubnetId` will be required.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`
}

func (r *CreateRoInstanceIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoInstanceIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UniqSubnetId")
	delete(f, "UniqVpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoInstanceIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoInstanceIpResponseParams struct {
	// VPC ID of the read-only instance.
	RoVpcId *int64 `json:"RoVpcId,omitnil,omitempty" name:"RoVpcId"`

	// Subnet ID of the read-only instance.
	RoSubnetId *int64 `json:"RoSubnetId,omitnil,omitempty" name:"RoSubnetId"`

	// Private IP address of the read-only instance.
	RoVip *string `json:"RoVip,omitnil,omitempty" name:"RoVip"`

	// Private port number of the read-only instance.
	RoVport *int64 `json:"RoVport,omitnil,omitempty" name:"RoVport"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoInstanceIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoInstanceIpResponseParams `json:"Response"`
}

func (r *CreateRoInstanceIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoInstanceIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRotationPasswordRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Currently, enable password rotation for account information, including account name and host name.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateRotationPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Currently, enable password rotation for account information, including account name and host name.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *CreateRotationPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRotationPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRotationPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRotationPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRotationPasswordResponse struct {
	*tchttp.BaseResponse
	Response *CreateRotationPasswordResponseParams `json:"Response"`
}

func (r *CreateRotationPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRotationPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomConfig struct {
	// device
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// Type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Device type
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Memory, measured in MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type DBSwitchInfo struct {
	// Switch time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-09-03 01:34:31
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`

	// Switch type. Value range: TRANSFER (data migration), MASTER2SLAVE (primary/secondary switch), RECOVERY (primary/secondary recovery)
	SwitchType *string `json:"SwitchType,omitnil,omitempty" name:"SwitchType"`
}

type DatabasePrivilege struct {
	// Permission information
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`
}

type DatabasesWithCharacterLists struct {
	// Database name
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Character set
	CharacterSet *string `json:"CharacterSet,omitnil,omitempty" name:"CharacterSet"`
}

// Predefined struct for user
type DeleteAccountsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TencentDB account.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type DeleteAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TencentDB account.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

func (r *DeleteAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountsResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountsResponseParams `json:"Response"`
}

func (r *DeleteAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileRequestParams struct {
	// Audit log file name, which can be obtained through the [DescribeAuditLogFiles](https://www.tencentcloud.com/document/api/236/45454?from_cn_redirect=1) API.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// Audit log file name, which can be obtained through the [DescribeAuditLogFiles](https://www.tencentcloud.com/document/api/236/45454?from_cn_redirect=1) API.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditLogFileResponseParams `json:"Response"`
}

func (r *DeleteAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditPolicyRequestParams struct {
	// Audit policy ID.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteAuditPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Audit policy ID.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteAuditPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditPolicyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuditPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditPolicyResponseParams `json:"Response"`
}

func (r *DeleteAuditPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API. A maximum of 5 rule templates can be deleted per request.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API. A maximum of 5 rule templates can be deleted per request.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

func (r *DeleteAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRuleTemplatesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DeleteAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup task ID. Confirm it by querying the data backup file list (https://www.tencentcloud.com/document/api/236/15842?from_cn_redirect=1) to get the target backup task ID.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup task ID. Confirm it by querying the data backup file list (https://www.tencentcloud.com/document/api/236/15842?from_cn_redirect=1) to get the target backup task ID.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

func (r *DeleteBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupResponseParams `json:"Response"`
}

func (r *DeleteBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatabaseRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name, length not exceeding 64.
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`
}

type DeleteDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name, length not exceeding 64.
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`
}

func (r *DeleteDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatabaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatabaseResponseParams `json:"Response"`
}

func (r *DeleteDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// Parameter template ID, which can be obtained through the API [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1).
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID, which can be obtained through the API [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1).
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParamTemplateResponseParams `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRotationPasswordRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance account name with password rotation disabled, such as root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Disable the domain name of the instance account with password rotation, such as%
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The latest password of the instance account after disabling password rotation
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// If the input is not null, the password is encrypted.
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

type DeleteRotationPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance account name with password rotation disabled, such as root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Disable the domain name of the instance account with password rotation, such as%
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The latest password of the instance account after disabling password rotation
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// If the input is not null, the password is encrypted.
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

func (r *DeleteRotationPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRotationPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "EncryptMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRotationPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRotationPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRotationPasswordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRotationPasswordResponseParams `json:"Response"`
}

func (r *DeleteRotationPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRotationPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimeWindowRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTimeWindowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTimeWindowResponseParams `json:"Response"`
}

func (r *DeleteTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Account name of the database. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Domain name of the database account. Obtain through the API [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1).
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Account name of the database. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Domain name of the database account. Obtain through the API [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1).
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// Array of global permissions.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Array of database permissions.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Array of table permissions in the database.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// Array of column permissions in the table.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Value range: 1-100. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Regex for matching account names, which complies with the rules at MySQL's official website
	AccountRegexp *string `json:"AccountRegexp,omitnil,omitempty" name:"AccountRegexp"`

	// Default none, support: ASC, DESC, asc, desc
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Time field for sorting. Options: CreateTime (account creation time), ModifyTime (update time), ModifyPasswordTime (password modification time).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Regular expression to match the account host address (Host). The rule is the same as that on the MySQL official website.
	HostRegexp *string `json:"HostRegexp,omitnil,omitempty" name:"HostRegexp"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Record offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Value range: 1-100. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Regex for matching account names, which complies with the rules at MySQL's official website
	AccountRegexp *string `json:"AccountRegexp,omitnil,omitempty" name:"AccountRegexp"`

	// Default none, support: ASC, DESC, asc, desc
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Time field for sorting. Options: CreateTime (account creation time), ModifyTime (update time), ModifyPasswordTime (password modification time).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Regular expression to match the account host address (Host). The rule is the same as that on the MySQL official website.
	HostRegexp *string `json:"HostRegexp,omitnil,omitempty" name:"HostRegexp"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AccountRegexp")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "HostRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Number of eligible accounts
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of eligible accounts
	Items []*AccountInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The maximum number of instance connections
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// Task execution result. Possible values: INITIAL - Initialization, RUNNING - Running, SUCCESS - Execution successful, FAILED - Execution failed, KILLED - Terminated, REMOVED - Deleted, PAUSED - Terminating.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task execution information description.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditConfigRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAuditConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAuditConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditConfigResponseParams struct {
	// Audit log retention period. Valid values: [0, 7, 30, 180, 365, 1095, 1825].
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Audit log storage type. Valid value: "storage" - Storage type.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Whether the audit service is being disabled. Valid values: "false" - No, "true" - Yes.
	IsClosing *string `json:"IsClosing,omitnil,omitempty" name:"IsClosing"`

	// Whether the audit service is being enabled. Valid values: "false" - No, "true" - Yes.
	IsOpening *string `json:"IsOpening,omitnil,omitempty" name:"IsOpening"`

	// Time when the audit service was activated.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditConfigResponseParams `json:"Response"`
}

func (r *DescribeAuditConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListRequestParams struct {
	// Whether audit is enabled for the instance. Valid values: 1 - Enabled; 0 - Disabled.
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// Filter conditions for querying the instance list.
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Audit rule mode for the instance. Valid values: 1 - Rule-based audit; 0 - Full audit.
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// Number of entries to return per request. Default value: 30. Maximum value: 20000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// Whether audit is enabled for the instance. Valid values: 1 - Enabled; 0 - Disabled.
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// Filter conditions for querying the instance list.
	Filters []*AuditInstanceFilters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Audit rule mode for the instance. Valid values: 1 - Rule-based audit; 0 - Full audit.
	AuditMode *int64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// Number of entries to return per request. Default value: 30. Maximum value: 20000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAuditInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditSwitch")
	delete(f, "Filters")
	delete(f, "AuditMode")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListResponseParams struct {
	// Total number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of audit instance details.Note: This field may return null, indicating that no valid values can be obtained.
	Items []*InstanceDbAuditStatus `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAuditInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page size. Default value: 20; minimum value: 1; maximum value: 300.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Audit log file name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type DescribeAuditLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page size. Default value: 20; minimum value: 1; maximum value: 300.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Audit log file name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *DescribeAuditLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesResponseParams struct {
	// Number of eligible audit log files.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Audit log file details.
	Items []*AuditLogFile `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogFilesResponseParams `json:"Response"`
}

func (r *DescribeAuditLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogsRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time. We recommend that the interval between start and end time does not exceed 7 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. We recommend that the interval between start and end time does not exceed 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The pagination parameter, which specifies the number of entries per page. Maximum value: 100 (default).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log offset, supports up to 65535 log entries for offset querying. Fill in the range: 0 - 65535.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort order. Valid values: "ASC" - Ascending order, "DESC" - Descending order. Default value: "DESC".
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field to sort by. Valid values:
	// "timestamp" - timestamp;
	// "affectRows" - Number of affected rows.
	// "execTime" - Execution time.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter. Multiple values are in `AND` relationship.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type DescribeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time. We recommend that the interval between start and end time does not exceed 7 days.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. We recommend that the interval between start and end time does not exceed 7 days.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The pagination parameter, which specifies the number of entries per page. Maximum value: 100 (default).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log offset, supports up to 65535 log entries for offset querying. Fill in the range: 0 - 65535.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort order. Valid values: "ASC" - Ascending order, "DESC" - Descending order. Default value: "DESC".
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Field to sort by. Valid values:
	// "timestamp" - timestamp;
	// "affectRows" - Number of affected rows.
	// "execTime" - Execution time.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter. Multiple values are in `AND` relationship.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

func (r *DescribeAuditLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "LogFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogsResponseParams struct {
	// Number of eligible audit logs
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Audit log details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*AuditLog `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogsResponseParams `json:"Response"`
}

func (r *DescribeAuditLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditPoliciesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit policy ID.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Audit policy name. Fuzzy match query is supported.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Audit rule ID, which can be used to query its associated audit policies.
	// Note: At least one of the parameters (“RuleId”, “PolicyId”, PolicyId”, or “PolicyName”) must be passed in.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type DescribeAuditPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit policy ID.
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// Audit policy name. Fuzzy match query is supported.
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Audit rule ID, which can be used to query its associated audit policies.
	// Note: At least one of the parameters (“RuleId”, “PolicyId”, PolicyId”, or “PolicyName”) must be passed in.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *DescribeAuditPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	delete(f, "PolicyName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditPoliciesResponseParams struct {
	// Number of eligible audit policies
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Audit policy details.
	Items []*AuditPolicy `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAuditPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplateModifyHistoryRequestParams struct {
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Start time of the query range.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query range.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of entries to return. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort order. DESC - Sorted by modification time in descending order, ASC - Ascending order. Default value: DESC.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeAuditRuleTemplateModifyHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Start time of the query range.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query range.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of entries to return. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sort order. DESC - Sorted by modification time in descending order, ASC - Ascending order. Default value: DESC.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeAuditRuleTemplateModifyHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplateModifyHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplateModifyHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplateModifyHistoryResponseParams struct {
	// Total number of entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Change details.
	Items []*RuleTemplateRecordInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditRuleTemplateModifyHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleTemplateModifyHistoryResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleTemplateModifyHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplateModifyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesRequestParams struct {
	// Rule template ID.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Rule template name.
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// Number of entries to return per request. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type DescribeAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Rule template ID.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Rule template name.
	RuleTemplateNames []*string `json:"RuleTemplateNames,omitnil,omitempty" name:"RuleTemplateNames"`

	// Number of entries to return per request. Default value: 20. Maximum value: 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

func (r *DescribeAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleTemplateNames")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRuleTemplatesResponseParams struct {
	// Total number of eligible instances.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of rule template details.
	Items []*AuditRuleTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRulesRequestParams struct {
	// Audit rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Audit rule name. Fuzzy match query is supported.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditRulesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule ID.
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Audit rule name. Fuzzy match query is supported.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAuditRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRulesResponseParams struct {
	// Number of eligible audit rules
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Audit rule details
	// Note: This field may return `null`, indicating that no valid value was found.
	Items []*AuditRule `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditRulesResponseParams `json:"Response"`
}

func (r *DescribeAuditRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigResponseParams struct {
	// Earliest start time point of automatic backup, such as 2 (for 2:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
	//
	// Deprecated: StartTimeMin is deprecated.
	StartTimeMin *int64 `json:"StartTimeMin,omitnil,omitempty" name:"StartTimeMin"`

	// Latest start time point of automatic backup, such as 6 (for 6:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
	//
	// Deprecated: StartTimeMax is deprecated.
	StartTimeMax *int64 `json:"StartTimeMax,omitnil,omitempty" name:"StartTimeMax"`

	// Backup file retention period in days.
	BackupExpireDays *int64 `json:"BackupExpireDays,omitnil,omitempty" name:"BackupExpireDays"`

	// Backup mode. Value range: physical, logical
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Binlog file retention period in days.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// Time window for automatic instance backup.
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// Switch for archive backup retention. Valid values: `off` (disable), `on` (enable). Default value:`off`.
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitnil,omitempty" name:"EnableBackupPeriodSave"`

	// Maximum days of archive backup retention. Valid range: 90-3650. Default value: 1080.
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitnil,omitempty" name:"BackupPeriodSaveDays"`

	// Archive backup retention period. Valid values: `weekly` (a week), `monthly` (a month), `quarterly` (a quarter), `yearly` (a year). Default value: `monthly`.
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitnil,omitempty" name:"BackupPeriodSaveInterval"`

	// Number of archive backups. Minimum value: `1`, Maximum value: Number of non-archive backups in archive backup retention period. Default value: `1`.
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitnil,omitempty" name:"BackupPeriodSaveCount"`

	// The start time in the format: yyyy-mm-dd HH:MM:SS, which is used to enable archive backup retention policy.
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitnil,omitempty" name:"StartBackupPeriodSaveDate"`

	// Whether to enable the archive backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// The period (in days) of how long a data backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// Whether to enable the archive backup of logs. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// The period (in days) of how long a log backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// Whether to enable the standard storage policy for data backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// The period (in days) of how long a data backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// Whether to enable the standard storage policy for log backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// The period (in days) of how long a log backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDecryptionKeyRequestParams struct {
	// Instance ID, in the format such as cdb-fybaegd8. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup ID of the instance, which can be obtained through the [DescribeBackups](https://www.tencentcloud.com/document/api/236/15842?from_cn_redirect=1) API.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup type. data - data backup, binlog - log backup. The default value is data.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`
}

type DescribeBackupDecryptionKeyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-fybaegd8. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup ID of the instance, which can be obtained through the [DescribeBackups](https://www.tencentcloud.com/document/api/236/15842?from_cn_redirect=1) API.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup type. data - data backup, binlog - log backup. The default value is data.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`
}

func (r *DescribeBackupDecryptionKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDecryptionKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "BackupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDecryptionKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDecryptionKeyResponseParams struct {
	// The decryption key of a backup file
	DecryptionKey *string `json:"DecryptionKey,omitnil,omitempty" name:"DecryptionKey"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDecryptionKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDecryptionKeyResponseParams `json:"Response"`
}

func (r *DescribeBackupDecryptionKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDecryptionKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionRequestParams struct {

}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// Valid values: `NoLimit` (backups can be downloaded over both private and public networks with any IPs), `LimitOnlyIntranet` (backups can be downloaded over the private network with any private IPs), `Customize` (backups can be downloaded over specified VPCs with specified IPs). The `LimitVpc` and `LimitIp` parameters are valid only when this parameter is set to `Customize`.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Valid value: `In` (backups can only be downloaded over the VPCs specified in `LimitVpc`).
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Valid values: `In` (backups can only be downloaded with the IPs specified in `LimitIp`), `NotIn` (backups cannot be downloaded with the IPs specified in `LimitIp`).
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPCs used to restrict backup download.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// IPs used to restrict backup download.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupEncryptionStatusRequestParams struct {
	// Instance ID in the format of cdb-XXXX, which is the same as that displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-XXXX, which is the same as that displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupEncryptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupEncryptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupEncryptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupEncryptionStatusResponseParams struct {
	// Whether the physical cold backup is enabled for the instance. Valid values: `on`, `off`.
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupEncryptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupEncryptionStatusResponseParams `json:"Response"`
}

func (r *DescribeBackupEncryptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupEncryptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewRequestParams struct {
	// The cloud database product type for which a backup overview needs to be queried. The value can be mysql (two-node/three-node high-availability instances), mysql-basic (single-node cloud disk edition instance), or mysql-cluster (cloud disk edition instance).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// The cloud database product type for which a backup overview needs to be queried. The value can be mysql (two-node/three-node high-availability instances), mysql-basic (single-node cloud disk edition instance), or mysql-cluster (cloud disk edition instance).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewResponseParams struct {
	// Total number of backups of a user in the current region (including data backups and log backups).
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`

	// Total capacity of backups of a user in the current region.
	BackupVolume *int64 `json:"BackupVolume,omitnil,omitempty" name:"BackupVolume"`

	// Paid capacity of backups of a user in the current region, i.e., capacity that exceeds the free tier.
	BillingVolume *int64 `json:"BillingVolume,omitnil,omitempty" name:"BillingVolume"`

	// Backup capacity in the free tier of a user in the current region.
	FreeVolume *int64 `json:"FreeVolume,omitnil,omitempty" name:"FreeVolume"`

	// Total offsite backup capacity of the user in current region.
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitnil,omitempty" name:"RemoteBackupVolume"`

	// Archive backup capacity, including data backup and log backup.
	BackupArchiveVolume *int64 `json:"BackupArchiveVolume,omitnil,omitempty" name:"BackupArchiveVolume"`

	// Standard storage backup capacity includes data backup and log backup.
	BackupStandbyVolume *int64 `json:"BackupStandbyVolume,omitnil,omitempty" name:"BackupStandbyVolume"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesRequestParams struct {
	// The cloud database product type for which real-time backup statistics need to be queried. The value can be mysql (two-node/three-node high-availability instances), mysql-basic (single-node cloud disk edition instance), or mysql-cluster (cloud disk edition instance).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries returned per page, which ranges from 1 to 100. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting criterion. Valid values: `BackupVolume` (backup capacity), `DataBackupVolume` (data backup capacity), `BinlogBackupVolume` (log backup capacity), `AutoBackupVolume` (automatic backup capacity), `ManualBackupVolume` (manual backup capacity). Default value: `BackupVolume`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending). Default value: `ASC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest
	
	// The cloud database product type for which real-time backup statistics need to be queried. The value can be mysql (two-node/three-node high-availability instances), mysql-basic (single-node cloud disk edition instance), or mysql-cluster (cloud disk edition instance).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries returned per page, which ranges from 1 to 100. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting criterion. Valid values: `BackupVolume` (backup capacity), `DataBackupVolume` (data backup capacity), `BinlogBackupVolume` (log backup capacity), `AutoBackupVolume` (automatic backup capacity), `ManualBackupVolume` (manual backup capacity). Default value: `BackupVolume`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `ASC` (ascending), `DESC` (descending). Default value: `ASC`.
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeBackupSummariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupSummariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesResponseParams struct {
	// Statistical items of instance backup.
	Items []*BackupSummaryItem `json:"Items,omitnil,omitempty" name:"Items"`

	// Total number of instance backups.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupSummariesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupSummariesResponseParams `json:"Response"`
}

func (r *DescribeBackupSummariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of eligible backups.
	Items []*BackupInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupsResponseParams `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogBackupOverviewRequestParams struct {
	// The cloud database product type for which a log backup overview needs to be queried. The value can be mysql (two-node/three-node high-availability instances), mysql-basic (single-node cloud disk edition instance), or mysql-cluster (cloud disk edition instance).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeBinlogBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// The cloud database product type for which a log backup overview needs to be queried. The value can be mysql (two-node/three-node high-availability instances), mysql-basic (single-node cloud disk edition instance), or mysql-cluster (cloud disk edition instance).
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeBinlogBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogBackupOverviewResponseParams struct {
	// Total capacity of log backups in bytes (including remote log backups)
	BinlogBackupVolume *int64 `json:"BinlogBackupVolume,omitnil,omitempty" name:"BinlogBackupVolume"`

	// Total number of log backups (include remote log backups)
	BinlogBackupCount *int64 `json:"BinlogBackupCount,omitnil,omitempty" name:"BinlogBackupCount"`

	// Capacity of remote log backups in bytes
	RemoteBinlogVolume *int64 `json:"RemoteBinlogVolume,omitnil,omitempty" name:"RemoteBinlogVolume"`

	// Number of remote backups
	RemoteBinlogCount *int64 `json:"RemoteBinlogCount,omitnil,omitempty" name:"RemoteBinlogCount"`

	// Capacity of archive log backups in bytes
	BinlogArchiveVolume *int64 `json:"BinlogArchiveVolume,omitnil,omitempty" name:"BinlogArchiveVolume"`

	// Number of archived log backups
	BinlogArchiveCount *int64 `json:"BinlogArchiveCount,omitnil,omitempty" name:"BinlogArchiveCount"`

	// Log backup capacity of standard storage in bytes
	BinlogStandbyVolume *int64 `json:"BinlogStandbyVolume,omitnil,omitempty" name:"BinlogStandbyVolume"`

	// Number of log backups of standard storage
	BinlogStandbyCount *int64 `json:"BinlogStandbyCount,omitnil,omitempty" name:"BinlogStandbyCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeBinlogBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The earliest start time of binlog  in the format of 2016-03-17 02:10:37.
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// The latest start time of binlog  in the format of 2016-03-17 02:10:37.
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`

	// Whether the binlog list contains the starting node MinStartTime, no by default
	ContainsMinStartTime *bool `json:"ContainsMinStartTime,omitnil,omitempty" name:"ContainsMinStartTime"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The earliest start time of binlog  in the format of 2016-03-17 02:10:37.
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// The latest start time of binlog  in the format of 2016-03-17 02:10:37.
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`

	// Whether the binlog list contains the starting node MinStartTime, no by default
	ContainsMinStartTime *bool `json:"ContainsMinStartTime,omitnil,omitempty" name:"ContainsMinStartTime"`
}

func (r *DescribeBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MinStartTime")
	delete(f, "MaxStartTime")
	delete(f, "ContainsMinStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBinlogsResponseParams struct {
	// Number of eligible log files.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Number of eligible binlog files.
	Items []*BinlogInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBinlogsResponseParams `json:"Response"`
}

func (r *DescribeBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBinlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCPUExpandStrategyInfoRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCPUExpandStrategyInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCPUExpandStrategyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCPUExpandStrategyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCPUExpandStrategyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCPUExpandStrategyInfoResponseParams struct {
	// Policy type. Output value: auto, manual, timeInterval, period.
	// Description: 1. auto means auto-scaling. 2. manual means custom scaling with immediate effect. 3. timeInterval means custom scaling by time. 4. period means custom scaling by cycle. 5. If the return is NULL, the elastic expansion strategy is not yet opened.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Custom expansion with CPU that takes effect immediately. Valid when Type is manual.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Auto scale-out policy. Valid when Type is auto.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// Scaling policy by cycle. Valid when Type is period.
	PeriodStrategy *PeriodStrategy `json:"PeriodStrategy,omitnil,omitempty" name:"PeriodStrategy"`

	// Scaling policy by time period. Valid when Type is timeInterval.
	TimeIntervalStrategy *TimeIntervalStrategy `json:"TimeIntervalStrategy,omitnil,omitempty" name:"TimeIntervalStrategy"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCPUExpandStrategyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCPUExpandStrategyInfoResponseParams `json:"Response"`
}

func (r *DescribeCPUExpandStrategyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCPUExpandStrategyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbProxyInfoRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type DescribeCdbProxyInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *DescribeCdbProxyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbProxyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdbProxyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbProxyInfoResponseParams struct {
	// Number of proxy groups
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Proxy group information
	ProxyInfos []*ProxyGroupInfo `json:"ProxyInfos,omitnil,omitempty" name:"ProxyInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCdbProxyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdbProxyInfoResponseParams `json:"Response"`
}

func (r *DescribeCdbProxyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbProxyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbZoneConfigRequestParams struct {

}

type DescribeCdbZoneConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCdbZoneConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbZoneConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdbZoneConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdbZoneConfigResponseParams struct {
	// List of purchasable specification and region information
	DataResult *CdbZoneDataResult `json:"DataResult,omitnil,omitempty" name:"DataResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCdbZoneConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdbZoneConfigResponseParams `json:"Response"`
}

func (r *DescribeCdbZoneConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdbZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneListRequestParams struct {
	// Query the cloning task list of the specified source instance. Obtain the instance ID through the [DescribeDBInstances](https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page for paging query. Default value: 20. Maximum value: 100 recommended.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCloneListRequest struct {
	*tchttp.BaseRequest
	
	// Query the cloning task list of the specified source instance. Obtain the instance ID through the [DescribeDBInstances](https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page for paging query. Default value: 20. Maximum value: 100 recommended.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCloneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneListResponseParams struct {
	// The number of results which meet the conditions
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Clone task list
	Items []*CloneItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloneListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloneListResponseParams `json:"Response"`
}

func (r *DescribeCloneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInfoRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	// Description: Only able to input the instance ID of instances with cloud disk architecture, corresponding to console instance configurations displayed as "Cloud Disk Edition (Cloud Disk)".
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeClusterInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	// Description: Only able to input the instance ID of instances with cloud disk architecture, corresponding to console instance configurations displayed as "Cloud Disk Edition (Cloud Disk)".
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInfoResponseParams struct {
	// Instance name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Address information for reading and writing of the cloud disk edition instance.
	ReadWriteAddress *AddressInfo `json:"ReadWriteAddress,omitnil,omitempty" name:"ReadWriteAddress"`

	// Read-only address information of the cloud disk edition instance.
	ReadOnlyAddress []*AddressInfo `json:"ReadOnlyAddress,omitnil,omitempty" name:"ReadOnlyAddress"`

	// Node list information of the Cloud Disk Edition instance.
	NodeList []*ClusterNodeInfo `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// Read-only space protection threshold, GB
	ReadonlyLimit *int64 `json:"ReadonlyLimit,omitnil,omitempty" name:"ReadonlyLimit"`

	// Number of instance nodes.
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInfoResponseParams `json:"Response"`
}

func (r *DescribeClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBFeaturesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBFeaturesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBFeaturesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBFeaturesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBFeaturesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBFeaturesResponseParams struct {
	// Whether database audit is supported
	IsSupportAudit *bool `json:"IsSupportAudit,omitnil,omitempty" name:"IsSupportAudit"`

	// Whether enabling audit requires a kernel version upgrade
	AuditNeedUpgrade *bool `json:"AuditNeedUpgrade,omitnil,omitempty" name:"AuditNeedUpgrade"`

	// Whether database encryption is supported
	IsSupportEncryption *bool `json:"IsSupportEncryption,omitnil,omitempty" name:"IsSupportEncryption"`

	// Whether enabling encryption requires a kernel version upgrade
	EncryptionNeedUpgrade *bool `json:"EncryptionNeedUpgrade,omitnil,omitempty" name:"EncryptionNeedUpgrade"`

	// Whether the instance is a remote read-only instance
	IsRemoteRo *bool `json:"IsRemoteRo,omitnil,omitempty" name:"IsRemoteRo"`

	// Primary instance region.
	// Description: This parameter may return null. You can ignore this return value. If needed, you can call the [Query Instance List](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API to obtain the instance region details.
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// Whether minor version upgrade is supported
	IsSupportUpdateSubVersion *bool `json:"IsSupportUpdateSubVersion,omitnil,omitempty" name:"IsSupportUpdateSubVersion"`

	// The current kernel version
	CurrentSubVersion *string `json:"CurrentSubVersion,omitnil,omitempty" name:"CurrentSubVersion"`

	// Available kernel version for upgrade
	TargetSubVersion *string `json:"TargetSubVersion,omitnil,omitempty" name:"TargetSubVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBFeaturesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBFeaturesResponseParams `json:"Response"`
}

func (r *DescribeDBFeaturesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBFeaturesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBImportRecordsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 00:00:01.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 23:59:59.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination parameter indicating the offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter indicating the number of results to be returned for a single request. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBImportRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 00:00:01.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-01-01 23:59:59.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pagination parameter indicating the offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameter indicating the number of results to be returned for a single request. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBImportRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBImportRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBImportRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBImportRecordsResponseParams struct {
	// Number of eligible import task operation logs.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of import operation records.
	Items []*ImportRecord `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBImportRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBImportRecordsResponseParams `json:"Response"`
}

func (r *DescribeDBImportRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBImportRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceCharsetRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceCharsetRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceCharsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceCharsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceCharsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceCharsetResponseParams struct {
	// Default character set of the instance, such as "latin1" and "utf8".
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceCharsetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceCharsetResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceCharsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceCharsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceConfigRequestParams struct {
	// <p>Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceConfigResponseParams struct {
	// <p>Data protection method of the primary instance, possible returned values: 0 - asynchronous replication mode, 1 - semi-sync replication mode, 2 - strong sync replication mode.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Primary instance deployment mode. Possible returned values: 0 - single AZ deployment, 1 - multi-AZ deployment.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>Primary AZ information of the instance, in the format of "ap-shanghai-2".</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Configuration message of the first standby for two-node, three-node, and four-node instances.</p><p>When querying a two-node instance, this parameter returns the standby information of the two-node instance. When querying a three-node or four-node instance, this parameter returns the first standby information of the instance.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	SlaveConfig *SlaveConfig `json:"SlaveConfig,omitnil,omitempty" name:"SlaveConfig"`

	// <p>Configuration message of the second standby database for three-node and 4-node instances.</p><p>When querying three-node and 4-node instances, this parameter returns the information of the second standby database.</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupConfig *BackupConfig `json:"BackupConfig,omitnil,omitempty" name:"BackupConfig"`

	// <p>Whether to switch over to the standby database.</p>
	Switched *bool `json:"Switched,omitnil,omitempty" name:"Switched"`

	// <p>Configuration message of the third standby database in a 4-node instance.</p><p>When querying a 4-node instance, this parameter returns the info of the third standby database.</p>
	FourthConfig *BackupConfig `json:"FourthConfig,omitnil,omitempty" name:"FourthConfig"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceConfigResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceGTIDRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceGTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceGTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceGTIDResponseParams struct {
	// GTID enablement flag. Value range: 0 (not enabled), 1 (enabled).
	IsGTIDOpen *int64 `json:"IsGTIDOpen,omitnil,omitempty" name:"IsGTIDOpen"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceGTIDResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceGTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInfoRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	// Description: Only the primary instance supports querying. This item only supports input of the primary instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	// Description: Only the primary instance supports querying. This item only supports input of the primary instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInfoResponseParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Whether encryption is enabled. YES: enabled, NO: not enabled.
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// Key ID used for encryption.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key region.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// The default region of the KMS service used by the current CDB backend service.
	DefaultKmsRegion *string `json:"DefaultKmsRegion,omitnil,omitempty" name:"DefaultKmsRegion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceInfoResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceLogToCLSRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region of the CLS service
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

type DescribeDBInstanceLogToCLSRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region of the CLS service
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

func (r *DescribeDBInstanceLogToCLSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceLogToCLSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClsRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceLogToCLSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceLogToCLSResponseParams struct {
	// Error log delivery CLS configuration
	ErrorLog *LogToCLSConfig `json:"ErrorLog,omitnil,omitempty" name:"ErrorLog"`

	// Slow log delivery CLS configuration
	SlowLog *LogToCLSConfig `json:"SlowLog,omitnil,omitempty" name:"SlowLog"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceLogToCLSResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceLogToCLSResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceLogToCLSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceLogToCLSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceRebootTimeRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	// Description: Multiple instance IDs allowed for query. json format as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	// Description: Multiple instance IDs allowed for query. json format as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeDBInstanceRebootTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceRebootTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceRebootTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceRebootTimeResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Returned parameter information.
	Items []*InstanceRebootTime `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceRebootTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceRebootTimeResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceRebootTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceRebootTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Instance type. Valid values: 1 - Primary instance, 2 - Disaster recovery instance, 3 - Read-only instance.</p>
	InstanceTypes []*uint64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// <p>Private IP address of the instance.</p>
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// <p>Instance status. Valid values:<br>0 - Creating<br>1 - Running<br>4 - Isolation operation in progress<br>5 - Isolated (can be restored from the Recycle Bin)</p>
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Offset. Default value is 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items returned per request. Default value: 20. Maximum value: 2000.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Security group ID. When using security group ID as the filter condition, the WithSecurityGroup parameter needs to be specified as 1.</p>
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// <p>Payment type. Valid values: 0 - yearly/monthly subscription; 1 - bill by hour.</p>
	PayTypes []*uint64 `json:"PayTypes,omitnil,omitempty" name:"PayTypes"`

	// <p>Instance name.</p>
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// <p>Instance task status, possible values:<br>0 - No tasks<br>1 - Upgrading<br>2 - Data import in progress<br>3 - Enabling Slave<br>4 - Enabling public network access<br>5 - Batch operation in progress<br>6 - Rolling back<br>7 - Disabling public network access<br>8 - Password change in progress<br>9 - Renaming instance<br>10 - Restarting<br>12 - Self-built migration in progress<br>13 - Deleting database table<br>14 - Disaster recovery instance creation sync in progress<br>15 - Upgrade pending switch<br>16 - Upgrade and switch in progress<br>17 - Switch completed<br>19 - Parameter setting pending execution<br>34 - Node in-place upgrade to be executed</p>
	TaskStatus []*uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// <p>Database engine version of the instance. Possible values: 5.1, 5.5, 5.6, and 5.7.</p>
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// <p>VPC ID.</p>
	VpcIds []*uint64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// <p>Availability zone ID.</p>
	ZoneIds []*uint64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>Subnet ID.</p>
	SubnetIds []*uint64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// <p>Whether to set the lock flag. Available values: 0 - not lock, 1 - lock. Default is 0.</p>
	CdbErrors []*int64 `json:"CdbErrors,omitnil,omitempty" name:"CdbErrors"`

	// <p>Sorting field of the returned result set. Currently supports: "instanceId", "instanceName", "createTime", and "deadlineTime".</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting method of the returned result set. Valid values: "ASC" - ascending order; "DESC" - descending order. The default is "DESC".</p>
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// <p>Whether to use security group ID as the filter condition.<br>Description: 0 indicates no, 1 indicates yes.</p>
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitnil,omitempty" name:"WithSecurityGroup"`

	// <p>Whether the exclusive cluster detail is included. Value range: 0 - not contained, 1 - contained.</p>
	WithExCluster *int64 `json:"WithExCluster,omitnil,omitempty" name:"WithExCluster"`

	// <p>Dedicated cluster ID.</p>
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// <p>Instance ID.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Initialization flag. Valid values: 0 - uninitialized, 1 - initialized.</p>
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// <p>Whether the corresponding instance in the disaster recovery relationship is included. Valid values: 0 - excluding, 1 - included. Default value: 1. If pulling the primary instance, the data of the disaster recovery relationship is in the DrInfo field. If pulling the disaster recovery instance, the data of the disaster recovery relationship is in the MasterInfo field. The disaster recovery relationship only contains partial basic data. Detailed data must be pulled manually via the interface.</p>
	WithDr *int64 `json:"WithDr,omitnil,omitempty" name:"WithDr"`

	// <p>Whether it contains read-only instances. Valid values: 0 - does not include, 1 - includes. Default value is 1.</p>
	WithRo *int64 `json:"WithRo,omitnil,omitempty" name:"WithRo"`

	// <p>Whether the primary instance is included. Valid values: 0 - does not include, 1 - includes. Default value is 1.</p>
	WithMaster *int64 `json:"WithMaster,omitnil,omitempty" name:"WithMaster"`

	// <p>Placement group ID list.</p>
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`

	// <p>Filter by tag key.</p>
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitnil,omitempty" name:"TagKeysForSearch"`

	// <p>Financial Enclosure ID.</p>
	CageIds []*string `json:"CageIds,omitnil,omitempty" name:"CageIds"`

	// <p>Tag value</p>
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`

	// <p>Character type VPC ID</p>
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitnil,omitempty" name:"UniqueVpcIds"`

	// <p>VPC character type subnetId</p>
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// <p>Tag key value<br>Please note, tags of the instance being created are unable to query.</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Database proxy IP.</p>
	ProxyVips []*string `json:"ProxyVips,omitnil,omitempty" name:"ProxyVips"`

	// <p>Database proxy ID.</p>
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// <p>Database engine type. Valid values: InnoDB, RocksDB.</p>
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// <p>Whether to obtain the Cluster Edition instance node information. Valid values: true or false. The default value is false.</p>
	QueryClusterInfo *bool `json:"QueryClusterInfo,omitnil,omitempty" name:"QueryClusterInfo"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Instance type. Valid values: 1 - Primary instance, 2 - Disaster recovery instance, 3 - Read-only instance.</p>
	InstanceTypes []*uint64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// <p>Private IP address of the instance.</p>
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// <p>Instance status. Valid values:<br>0 - Creating<br>1 - Running<br>4 - Isolation operation in progress<br>5 - Isolated (can be restored from the Recycle Bin)</p>
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Offset. Default value is 0.</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items returned per request. Default value: 20. Maximum value: 2000.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Security group ID. When using security group ID as the filter condition, the WithSecurityGroup parameter needs to be specified as 1.</p>
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// <p>Payment type. Valid values: 0 - yearly/monthly subscription; 1 - bill by hour.</p>
	PayTypes []*uint64 `json:"PayTypes,omitnil,omitempty" name:"PayTypes"`

	// <p>Instance name.</p>
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// <p>Instance task status, possible values:<br>0 - No tasks<br>1 - Upgrading<br>2 - Data import in progress<br>3 - Enabling Slave<br>4 - Enabling public network access<br>5 - Batch operation in progress<br>6 - Rolling back<br>7 - Disabling public network access<br>8 - Password change in progress<br>9 - Renaming instance<br>10 - Restarting<br>12 - Self-built migration in progress<br>13 - Deleting database table<br>14 - Disaster recovery instance creation sync in progress<br>15 - Upgrade pending switch<br>16 - Upgrade and switch in progress<br>17 - Switch completed<br>19 - Parameter setting pending execution<br>34 - Node in-place upgrade to be executed</p>
	TaskStatus []*uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// <p>Database engine version of the instance. Possible values: 5.1, 5.5, 5.6, and 5.7.</p>
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// <p>VPC ID.</p>
	VpcIds []*uint64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// <p>Availability zone ID.</p>
	ZoneIds []*uint64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// <p>Subnet ID.</p>
	SubnetIds []*uint64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// <p>Whether to set the lock flag. Available values: 0 - not lock, 1 - lock. Default is 0.</p>
	CdbErrors []*int64 `json:"CdbErrors,omitnil,omitempty" name:"CdbErrors"`

	// <p>Sorting field of the returned result set. Currently supports: "instanceId", "instanceName", "createTime", and "deadlineTime".</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting method of the returned result set. Valid values: "ASC" - ascending order; "DESC" - descending order. The default is "DESC".</p>
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// <p>Whether to use security group ID as the filter condition.<br>Description: 0 indicates no, 1 indicates yes.</p>
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitnil,omitempty" name:"WithSecurityGroup"`

	// <p>Whether the exclusive cluster detail is included. Value range: 0 - not contained, 1 - contained.</p>
	WithExCluster *int64 `json:"WithExCluster,omitnil,omitempty" name:"WithExCluster"`

	// <p>Dedicated cluster ID.</p>
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// <p>Instance ID.</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>Initialization flag. Valid values: 0 - uninitialized, 1 - initialized.</p>
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// <p>Whether the corresponding instance in the disaster recovery relationship is included. Valid values: 0 - excluding, 1 - included. Default value: 1. If pulling the primary instance, the data of the disaster recovery relationship is in the DrInfo field. If pulling the disaster recovery instance, the data of the disaster recovery relationship is in the MasterInfo field. The disaster recovery relationship only contains partial basic data. Detailed data must be pulled manually via the interface.</p>
	WithDr *int64 `json:"WithDr,omitnil,omitempty" name:"WithDr"`

	// <p>Whether it contains read-only instances. Valid values: 0 - does not include, 1 - includes. Default value is 1.</p>
	WithRo *int64 `json:"WithRo,omitnil,omitempty" name:"WithRo"`

	// <p>Whether the primary instance is included. Valid values: 0 - does not include, 1 - includes. Default value is 1.</p>
	WithMaster *int64 `json:"WithMaster,omitnil,omitempty" name:"WithMaster"`

	// <p>Placement group ID list.</p>
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`

	// <p>Filter by tag key.</p>
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitnil,omitempty" name:"TagKeysForSearch"`

	// <p>Financial Enclosure ID.</p>
	CageIds []*string `json:"CageIds,omitnil,omitempty" name:"CageIds"`

	// <p>Tag value</p>
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`

	// <p>Character type VPC ID</p>
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitnil,omitempty" name:"UniqueVpcIds"`

	// <p>VPC character type subnetId</p>
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// <p>Tag key value<br>Please note, tags of the instance being created are unable to query.</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Database proxy IP.</p>
	ProxyVips []*string `json:"ProxyVips,omitnil,omitempty" name:"ProxyVips"`

	// <p>Database proxy ID.</p>
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// <p>Database engine type. Valid values: InnoDB, RocksDB.</p>
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// <p>Whether to obtain the Cluster Edition instance node information. Valid values: true or false. The default value is false.</p>
	QueryClusterInfo *bool `json:"QueryClusterInfo,omitnil,omitempty" name:"QueryClusterInfo"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceTypes")
	delete(f, "Vips")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SecurityGroupId")
	delete(f, "PayTypes")
	delete(f, "InstanceNames")
	delete(f, "TaskStatus")
	delete(f, "EngineVersions")
	delete(f, "VpcIds")
	delete(f, "ZoneIds")
	delete(f, "SubnetIds")
	delete(f, "CdbErrors")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "WithSecurityGroup")
	delete(f, "WithExCluster")
	delete(f, "ExClusterId")
	delete(f, "InstanceIds")
	delete(f, "InitFlag")
	delete(f, "WithDr")
	delete(f, "WithRo")
	delete(f, "WithMaster")
	delete(f, "DeployGroupIds")
	delete(f, "TagKeysForSearch")
	delete(f, "CageIds")
	delete(f, "TagValues")
	delete(f, "UniqueVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "Tags")
	delete(f, "ProxyVips")
	delete(f, "ProxyIds")
	delete(f, "EngineTypes")
	delete(f, "QueryClusterInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// <p>Total number of eligible instances.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Instance detail list.</p>
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBPriceRequestParams struct {
	// <p>Instance duration, in months, minimum value 1, maximum value 36. This field is invalid when querying the pay-as-you-go rate.</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>AZ information, in the format of "ap-guangzhou-2". For available values, query the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">DescribeDBZoneConfig</a> api. This parameter is required when InstanceId is empty.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance count. Default value is 1, minimum value 1, maximum value 100. This parameter is required when InstanceId is empty.</p>
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Instance memory size, measurement unit: MB. Required when InstanceId is empty. To ensure the input value is valid, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to obtain the saleable instance memory size range.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB. This parameter is required when InstanceId is empty. To ensure the input value is valid, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the saleable disk size range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Instance type, defaults to master. Supported values include: master - primary instance, ro - read-only instance, dr - disaster recovery instance. Required when InstanceId is empty.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Payment type. Supported values: PRE_PAID (yearly/monthly subscription) and HOUR_PAID (pay-as-you-go). This parameter is required when InstanceId is empty.</p>
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Instance isolation type.</p><p>Enumeration value:</p><ul><li>UNIVERSAL: General-purpose instance</li><li>EXCLUSIVE: Dedicated instance</li><li>CLOUD_NATIVE_CLUSTER: Standard type of cloud disk edition</li><li>CLOUD_NATIVE_CLUSTER_EXCLUSIVE: Enhanced type of cloud disk edition</li><li>CLOUD_NATIVE_CLUSTER_ULTRA: Flagship type of cloud disk edition</li></ul><p>Default value: UNIVERSAL</p><p>If needed to query the price of a single-node instance of cloud disk edition, set up this parameter as CLOUD_NATIVE_CLUSTER and specify parameter InstanceNodes as 1.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Number of instance nodes.<br>1. When querying the price of a read-only instance or a single-node instance, the value of this field is 1.<br>2. When querying the price of a dual-node instance, the value of this field is 2.<br>3. When querying the price of a three-node instance, the value of this field is 3.<br>4. When querying the price of a cloud disk edition instance, the value range of this field can be 2 - 6. A value of 2 means the cloud disk edition instance has 1 read-write node + 1 read-only node. A value of 6 means the cloud disk edition instance has 1 read-write node + 5 read-only nodes. For other values (3 - 5), the rule is 1 read-write node + (value - 1) read-only nodes.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>The number of CPU cores of the price-query instance, measurement unit: core. To ensure the validity of the input CPU value, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API acquisition to get the saleable number of cores. When this value is not specified, a default value will be padded based on the Memory size.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Query the ID of the instance to be renewed. If needed, fill in InstanceId and Period to query instance renewal price.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Usage-based billing tier. Valid only when PayType=HOUR_PAID. Supported values include: 1, 2, 3. For step duration, see https://www.tencentcloud.com/document/product/236/18335.?from_cn_redirect=1</p>
	Ladder *uint64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

	// <p>Disk Type. Specify this parameter when querying the price of CLOUD disk edition or single-node instance of CLOUD disk edition. Default value: SSD CLOUD Block Storage.<br>Supported values include:<br>"CLOUD_SSD" - SSD CLOUD Block Storage.<br>"CLOUD_HSSD" - Enhanced SSD CLOUD Disk.<br>"CLOUD_PREMIUM" - High-performance CLOUD Block Storage.</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

type DescribeDBPriceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance duration, in months, minimum value 1, maximum value 36. This field is invalid when querying the pay-as-you-go rate.</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>AZ information, in the format of "ap-guangzhou-2". For available values, query the <a href="https://www.tencentcloud.com/document/api/236/17229?from_cn_redirect=1">DescribeDBZoneConfig</a> api. This parameter is required when InstanceId is empty.</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance count. Default value is 1, minimum value 1, maximum value 100. This parameter is required when InstanceId is empty.</p>
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// <p>Instance memory size, measurement unit: MB. Required when InstanceId is empty. To ensure the input value is valid, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to obtain the saleable instance memory size range.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance disk size, unit: GB. This parameter is required when InstanceId is empty. To ensure the input value is valid, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the saleable disk size range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Instance type, defaults to master. Supported values include: master - primary instance, ro - read-only instance, dr - disaster recovery instance. Required when InstanceId is empty.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Payment type. Supported values: PRE_PAID (yearly/monthly subscription) and HOUR_PAID (pay-as-you-go). This parameter is required when InstanceId is empty.</p>
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// <p>Data replication method, defaults to 0. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Instance isolation type.</p><p>Enumeration value:</p><ul><li>UNIVERSAL: General-purpose instance</li><li>EXCLUSIVE: Dedicated instance</li><li>CLOUD_NATIVE_CLUSTER: Standard type of cloud disk edition</li><li>CLOUD_NATIVE_CLUSTER_EXCLUSIVE: Enhanced type of cloud disk edition</li><li>CLOUD_NATIVE_CLUSTER_ULTRA: Flagship type of cloud disk edition</li></ul><p>Default value: UNIVERSAL</p><p>If needed to query the price of a single-node instance of cloud disk edition, set up this parameter as CLOUD_NATIVE_CLUSTER and specify parameter InstanceNodes as 1.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Number of instance nodes.<br>1. When querying the price of a read-only instance or a single-node instance, the value of this field is 1.<br>2. When querying the price of a dual-node instance, the value of this field is 2.<br>3. When querying the price of a three-node instance, the value of this field is 3.<br>4. When querying the price of a cloud disk edition instance, the value range of this field can be 2 - 6. A value of 2 means the cloud disk edition instance has 1 read-write node + 1 read-only node. A value of 6 means the cloud disk edition instance has 1 read-write node + 5 read-only nodes. For other values (3 - 5), the rule is 1 read-write node + (value - 1) read-only nodes.</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>The number of CPU cores of the price-query instance, measurement unit: core. To ensure the validity of the input CPU value, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API acquisition to get the saleable number of cores. When this value is not specified, a default value will be padded based on the Memory size.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Query the ID of the instance to be renewed. If needed, fill in InstanceId and Period to query instance renewal price.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Usage-based billing tier. Valid only when PayType=HOUR_PAID. Supported values include: 1, 2, 3. For step duration, see https://www.tencentcloud.com/document/product/236/18335.?from_cn_redirect=1</p>
	Ladder *uint64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

	// <p>Disk Type. Specify this parameter when querying the price of CLOUD disk edition or single-node instance of CLOUD disk edition. Default value: SSD CLOUD Block Storage.<br>Supported values include:<br>"CLOUD_SSD" - SSD CLOUD Block Storage.<br>"CLOUD_HSSD" - Enhanced SSD CLOUD Disk.<br>"CLOUD_PREMIUM" - High-performance CLOUD Block Storage.</p>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

func (r *DescribeDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "GoodsNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "InstanceRole")
	delete(f, "PayType")
	delete(f, "ProtectMode")
	delete(f, "DeviceType")
	delete(f, "InstanceNodes")
	delete(f, "Cpu")
	delete(f, "InstanceId")
	delete(f, "Ladder")
	delete(f, "DiskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBPriceResponseParams struct {
	// <p>Instance price, unit: cent.</p>
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// <p>Original price of instance. Measurement unit: cent.</p>
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// <p>Currency unit. CNY - RMB, USD - USD.</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBPriceResponseParams `json:"Response"`
}

func (r *DescribeDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// This parameter takes effect only when the ID of a read-only instance is passed in. If the parameter is set to `False` or left empty, the security groups bound with the RO groups of the read-only instance can only be queried. If it is set to `True`, the security groups can be modified.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// When updating the read-only group of a cluster edition instance, specify the instance id in InstanceId and this parameter to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// This parameter takes effect only when the ID of a read-only instance is passed in. If the parameter is set to `False` or left empty, the security groups bound with the RO groups of the read-only instance can only be queried. If it is set to `True`, the security groups can be modified.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// When updating the read-only group of a cluster edition instance, specify the instance id in InstanceId and this parameter to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ForReadonlyInstance")
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group details.
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSwitchRecordsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 50. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBSwitchRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 50. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBSwitchRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSwitchRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSwitchRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSwitchRecordsResponseParams struct {
	// Number of instance switches.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of instance switches.
	Items []*DBSwitchInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSwitchRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSwitchRecordsResponseParams `json:"Response"`
}

func (r *DescribeDBSwitchRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSwitchRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataBackupOverviewRequestParams struct {
	// The cloud database product type for which you need to query the data backup overview. Value is: mysql refers to two-node/three-node high-availability instances, mysql-basic refers to single-node (cloud disk) instances, mysql-cluster refers to cloud disk edition instances.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDataBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// The cloud database product type for which you need to query the data backup overview. Value is: mysql refers to two-node/three-node high-availability instances, mysql-basic refers to single-node (cloud disk) instances, mysql-cluster refers to cloud disk edition instances.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDataBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataBackupOverviewResponseParams struct {
	// Total capacity of data backups in bytes in the current region (including automatic backups and manual backups).
	DataBackupVolume *int64 `json:"DataBackupVolume,omitnil,omitempty" name:"DataBackupVolume"`

	// Total number of data backups in the current region.
	DataBackupCount *int64 `json:"DataBackupCount,omitnil,omitempty" name:"DataBackupCount"`

	// Total capacity of automatic backups in the current region.
	AutoBackupVolume *int64 `json:"AutoBackupVolume,omitnil,omitempty" name:"AutoBackupVolume"`

	// Total number of automatic backups in the current region.
	AutoBackupCount *int64 `json:"AutoBackupCount,omitnil,omitempty" name:"AutoBackupCount"`

	// Total capacity of manual backups in the current region.
	ManualBackupVolume *int64 `json:"ManualBackupVolume,omitnil,omitempty" name:"ManualBackupVolume"`

	// Total number of manual backups in the current region.
	ManualBackupCount *int64 `json:"ManualBackupCount,omitnil,omitempty" name:"ManualBackupCount"`

	// Total capacity of remote backups
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitnil,omitempty" name:"RemoteBackupVolume"`

	// Total number of remote backups
	RemoteBackupCount *int64 `json:"RemoteBackupCount,omitnil,omitempty" name:"RemoteBackupCount"`

	// Total capacity of archive backups in the current region
	DataBackupArchiveVolume *int64 `json:"DataBackupArchiveVolume,omitnil,omitempty" name:"DataBackupArchiveVolume"`

	// Total number of archive backups in the current region
	DataBackupArchiveCount *int64 `json:"DataBackupArchiveCount,omitnil,omitempty" name:"DataBackupArchiveCount"`

	// Total backup capacity of standard storage in current region
	DataBackupStandbyVolume *int64 `json:"DataBackupStandbyVolume,omitnil,omitempty" name:"DataBackupStandbyVolume"`

	// Total number of standard storage backups in current region
	DataBackupStandbyCount *int64 `json:"DataBackupStandbyCount,omitnil,omitempty" name:"DataBackupStandbyCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeDataBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per request. Default value: 20. Minimum value: 1. Maximum value: 5000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Regular expression for matching database names.
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per request. Default value: 20. Minimum value: 1. Maximum value: 5000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Regular expression for matching database names.
	DatabaseRegexp *string `json:"DatabaseRegexp,omitnil,omitempty" name:"DatabaseRegexp"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DatabaseRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of an instance.
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// Database name and character set
	DatabaseList []*DatabasesWithCharacterLists `json:"DatabaseList,omitnil,omitempty" name:"DatabaseList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParamsRequestParams struct {
	// Engine version. Currently supports ["5.1", "5.5", "5.6", "5.7", "8.0"].
	// Description: Engine version is required.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Default parameter template type. Supported values include "HIGH_STABILITY" - high-stability template, "HIGH_PERFORMANCE" - high-performance template. Default value: HIGH_STABILITY.
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine, default value: InnoDB, valid values: InnoDB, RocksDB.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest
	
	// Engine version. Currently supports ["5.1", "5.5", "5.6", "5.7", "8.0"].
	// Description: Engine version is required.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Default parameter template type. Supported values include "HIGH_STABILITY" - high-stability template, "HIGH_PERFORMANCE" - high-performance template. Default value: HIGH_STABILITY.
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine, default value: InnoDB, valid values: InnoDB, RocksDB.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

func (r *DescribeDefaultParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersion")
	delete(f, "TemplateType")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParamsResponseParams struct {
	// Number of parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter details.
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDefaultParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultParamsResponseParams `json:"Response"`
}

func (r *DescribeDefaultParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceMonitorInfoRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// This parameter is used to return the monitoring data of Count 5-minute time periods on the day. Value range: 1-288. If this parameter is not passed in, all monitoring data in a 5-minute granularity on the day will be returned by default.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DescribeDeviceMonitorInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// This parameter is used to return the monitoring data of Count 5-minute time periods on the day. Value range: 1-288. If this parameter is not passed in, all monitoring data in a 5-minute granularity on the day will be returned by default.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

func (r *DescribeDeviceMonitorInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceMonitorInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceMonitorInfoResponseParams struct {
	// CPU monitoring data of the instance
	Cpu *DeviceCpuInfo `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory monitoring data of the instance
	Mem *DeviceMemInfo `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Network monitoring data of the instance
	Net *DeviceNetInfo `json:"Net,omitnil,omitempty" name:"Net"`

	// Disk monitoring data of the instance
	Disk *DeviceDiskInfo `json:"Disk,omitnil,omitempty" name:"Disk"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceMonitorInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceMonitorInfoResponseParams `json:"Response"`
}

func (r *DescribeDeviceMonitorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceMonitorInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorLogDataRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start timestamp. For example, 1585142640, in seconds.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp. For example, 1585142640, in seconds.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Keyword list to match, supports up to 15 keywords with fuzzy matching support.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// The number of results per page in paginated queries. Default value: 100. Maximum value: 400.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`
}

type DescribeErrorLogDataRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start timestamp. For example, 1585142640, in seconds.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp. For example, 1585142640, in seconds.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Keyword list to match, supports up to 15 keywords with fuzzy matching support.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// The number of results per page in paginated queries. Default value: 100. Maximum value: 400.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`
}

func (r *DescribeErrorLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "KeyWords")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorLogDataResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Returned records.
	Items []*ErrlogItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeErrorLogDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeErrorLogDataResponseParams `json:"Response"`
}

func (r *DescribeErrorLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAlarmEventsRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event query range start time, closed interval.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the event query range, closed interval.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Event name. Outofmemory - Memory OOM (status event); Switch - primary-secondary switch (status event); Roremove - read-only instance removal (status event); MemoryUsedHigh - high memory utilization (status event); CPUExpansion - CPU performance scale-out (stateless event); CPUExpansionFailed - CPU performance scale-out failed (stateless event); CPUContraction - CPU performance scale-in (stateless event); Restart - instance restart (status event); ServerFailureNodeMigration - ServerFailureNodeMigration (status event); PlannedSwitch - planned active/standby switch (stateless event); OverusedReadonlySet - instance will be locked (stateless event); OverusedReadWriteSet - instance unlock (stateless event).
	EventName []*string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Event status. "1" - Event; "0" - Recovery event; "-" - Stateless event.
	EventStatus *string `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// Sorting method. Sort by event occurrence. "DESC" - inverted; "ASC" - in order. Default is inverted.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Number of displayed events. Default is 100, maximum is 200.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DescribeInstanceAlarmEventsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Event query range start time, closed interval.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the event query range, closed interval.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Event name. Outofmemory - Memory OOM (status event); Switch - primary-secondary switch (status event); Roremove - read-only instance removal (status event); MemoryUsedHigh - high memory utilization (status event); CPUExpansion - CPU performance scale-out (stateless event); CPUExpansionFailed - CPU performance scale-out failed (stateless event); CPUContraction - CPU performance scale-in (stateless event); Restart - instance restart (status event); ServerFailureNodeMigration - ServerFailureNodeMigration (status event); PlannedSwitch - planned active/standby switch (stateless event); OverusedReadonlySet - instance will be locked (stateless event); OverusedReadWriteSet - instance unlock (stateless event).
	EventName []*string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Event status. "1" - Event; "0" - Recovery event; "-" - Stateless event.
	EventStatus *string `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// Sorting method. Sort by event occurrence. "DESC" - inverted; "ASC" - in order. Default is inverted.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Number of displayed events. Default is 100, maximum is 200.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DescribeInstanceAlarmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAlarmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EventName")
	delete(f, "EventStatus")
	delete(f, "Order")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAlarmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAlarmEventsResponseParams struct {
	// Number of events.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Event information. Items is null when info is not found.
	Items []*InstEventInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAlarmEventsResponseParams `json:"Response"`
}

func (r *DescribeInstanceAlarmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// Number of eligible records.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter modification records.
	Items []*ParamRecord `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// Number of instance parameters.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter details.
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancePasswordComplexityRequestParams struct {
	// Instance ID. 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstancePasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstancePasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancePasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancePasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancePasswordComplexityResponseParams struct {
	// Total number of password complexity related parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Password complexity parameter details. The policy value ranges from "" to "LOW" or "MEDIUM". Empty or LOW means password complexity is off. MEDIUM means password complexity is on. When the policy parameter value is MEDIUM, the following parameters take effect. length: value ranges from 8 to 64, means the minimum number of characters. mixed_case_count: value ranges from 1 to 16, means the minimum count of uppercase and lowercase letters. number_count: value ranges from 1 to 16, means the minimum count of numeric characters. special_char_count: value ranges from 1 to 16, means the minimum count of special characters.
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancePasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancePasswordComplexityResponseParams `json:"Response"`
}

func (r *DescribeInstancePasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancePasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeCheckJobRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target database version.
	// Description: Available values 5.6, 5.7, 8.0. Cross-version upgrade is not supported. Downgrade is not supported after upgrade.
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

type DescribeInstanceUpgradeCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target database version.
	// Description: Available values 5.6, 5.7, 8.0. Cross-version upgrade is not supported. Downgrade is not supported after upgrade.
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

func (r *DescribeInstanceUpgradeCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstMysqlVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceUpgradeCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeCheckJobResponseParams struct {
	// Existence of historic upgrade validation task within 24 hours
	ExistUpgradeCheckJob *bool `json:"ExistUpgradeCheckJob,omitnil,omitempty" name:"ExistUpgradeCheckJob"`

	// Task ID.
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceUpgradeCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceUpgradeCheckJobResponseParams `json:"Response"`
}

func (r *DescribeInstanceUpgradeCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeTypeRequestParams struct {
	// <p>Instance ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1">DescribeDBInstances</a> API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The number of CPU cores of the target instance. To ensure the input value is valid, please use <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> to get the saleable CPU range of the instance.</p>
	DstCpu *float64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// <p>Target instance memory size, measurement unit: MB. To ensure the input value is valid, please use <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> to get the saleable memory size range of the instance.</p>
	DstMemory *uint64 `json:"DstMemory,omitnil,omitempty" name:"DstMemory"`

	// <p>Target instance disk size, unit: GB. To ensure the input value is valid, please use <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> to get the saleable disk size range of the instance.</p>
	DstDisk *uint64 `json:"DstDisk,omitnil,omitempty" name:"DstDisk"`

	// <p>Target instance database version. Available values: 5.6, 5.7, 8.0.</p>
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// <p>Deployment model of the target instance. Defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone.</p>
	DstDeployMode *int64 `json:"DstDeployMode,omitnil,omitempty" name:"DstDeployMode"`

	// <p>Replication type of the target instance. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	DstProtectMode *int64 `json:"DstProtectMode,omitnil,omitempty" name:"DstProtectMode"`

	// <p>AZ ID of the standby instance 1 of the target instance. You can use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the availability zone ID.</p>
	DstSlaveZone *int64 `json:"DstSlaveZone,omitnil,omitempty" name:"DstSlaveZone"`

	// <p>AZ ID of the standby instance 2. You can use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the AZ ID.</p>
	DstBackupZone *int64 `json:"DstBackupZone,omitnil,omitempty" name:"DstBackupZone"`

	// <p>Target instance type. Supported values include: "CUSTOM" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "ONTKE" - ONTKE single-node instance, "CLOUD_NATIVE_CLUSTER" - standard type for CLOUD disk, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - enhanced type for CLOUD disk.</p>
	DstCdbType *string `json:"DstCdbType,omitnil,omitempty" name:"DstCdbType"`

	// <p>Primary availability zone ID of the target instance. You can use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the AZ ID.</p>
	DstZoneId *int64 `json:"DstZoneId,omitnil,omitempty" name:"DstZoneId"`

	// <p>Node distribution of CDB instances in the dedicated cluster.</p>
	NodeDistribution *NodeDistribution `json:"NodeDistribution,omitnil,omitempty" name:"NodeDistribution"`

	// <p>Topology configuration for cloud disk edition nodes. Nodeld information can be obtained through the <a href="https://www.tencentcloud.com/document/api/236/105116?from_cn_redirect=1">DescribeClusterInfo</a> API.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>AZ ID of the standby instance 3 in the target instance. Use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the availability zone ID.</p>
	DstFourthZone *int64 `json:"DstFourthZone,omitnil,omitempty" name:"DstFourthZone"`
}

type DescribeInstanceUpgradeTypeRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1">DescribeDBInstances</a> API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The number of CPU cores of the target instance. To ensure the input value is valid, please use <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> to get the saleable CPU range of the instance.</p>
	DstCpu *float64 `json:"DstCpu,omitnil,omitempty" name:"DstCpu"`

	// <p>Target instance memory size, measurement unit: MB. To ensure the input value is valid, please use <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> to get the saleable memory size range of the instance.</p>
	DstMemory *uint64 `json:"DstMemory,omitnil,omitempty" name:"DstMemory"`

	// <p>Target instance disk size, unit: GB. To ensure the input value is valid, please use <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> to get the saleable disk size range of the instance.</p>
	DstDisk *uint64 `json:"DstDisk,omitnil,omitempty" name:"DstDisk"`

	// <p>Target instance database version. Available values: 5.6, 5.7, 8.0.</p>
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// <p>Deployment model of the target instance. Defaults to 0. Supported values include: 0 - means single availability zone, 1 - means multi-availability zone.</p>
	DstDeployMode *int64 `json:"DstDeployMode,omitnil,omitempty" name:"DstDeployMode"`

	// <p>Replication type of the target instance. Supported values include: 0 - means async replication, 1 - means semi-sync replication, 2 - means strong sync replication.</p>
	DstProtectMode *int64 `json:"DstProtectMode,omitnil,omitempty" name:"DstProtectMode"`

	// <p>AZ ID of the standby instance 1 of the target instance. You can use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the availability zone ID.</p>
	DstSlaveZone *int64 `json:"DstSlaveZone,omitnil,omitempty" name:"DstSlaveZone"`

	// <p>AZ ID of the standby instance 2. You can use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the AZ ID.</p>
	DstBackupZone *int64 `json:"DstBackupZone,omitnil,omitempty" name:"DstBackupZone"`

	// <p>Target instance type. Supported values include: "CUSTOM" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "ONTKE" - ONTKE single-node instance, "CLOUD_NATIVE_CLUSTER" - standard type for CLOUD disk, "CLOUD_NATIVE_CLUSTER_EXCLUSIVE" - enhanced type for CLOUD disk.</p>
	DstCdbType *string `json:"DstCdbType,omitnil,omitempty" name:"DstCdbType"`

	// <p>Primary availability zone ID of the target instance. You can use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the AZ ID.</p>
	DstZoneId *int64 `json:"DstZoneId,omitnil,omitempty" name:"DstZoneId"`

	// <p>Node distribution of CDB instances in the dedicated cluster.</p>
	NodeDistribution *NodeDistribution `json:"NodeDistribution,omitnil,omitempty" name:"NodeDistribution"`

	// <p>Topology configuration for cloud disk edition nodes. Nodeld information can be obtained through the <a href="https://www.tencentcloud.com/document/api/236/105116?from_cn_redirect=1">DescribeClusterInfo</a> API.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>AZ ID of the standby instance 3 in the target instance. Use the <a href="https://www.tencentcloud.com/document/product/236/80281?from_cn_redirect=1">DescribeCdbZoneConfig</a> API to obtain the availability zone ID.</p>
	DstFourthZone *int64 `json:"DstFourthZone,omitnil,omitempty" name:"DstFourthZone"`
}

func (r *DescribeInstanceUpgradeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstCpu")
	delete(f, "DstMemory")
	delete(f, "DstDisk")
	delete(f, "DstVersion")
	delete(f, "DstDeployMode")
	delete(f, "DstProtectMode")
	delete(f, "DstSlaveZone")
	delete(f, "DstBackupZone")
	delete(f, "DstCdbType")
	delete(f, "DstZoneId")
	delete(f, "NodeDistribution")
	delete(f, "ClusterTopology")
	delete(f, "DstFourthZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceUpgradeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUpgradeTypeResponseParams struct {
	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance upgrade type. Trsf - Migration upgrade, InPlace - In-place upgrade, Topology - Architecture upgrade.</p>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceUpgradeTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceUpgradeTypeResponseParams `json:"Response"`
}

func (r *DescribeInstanceUpgradeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUpgradeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalBinlogConfigRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLocalBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLocalBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalBinlogConfigResponseParams struct {
	// Binlog retention policy of the instance.
	LocalBinlogConfig *LocalBinlogConfig `json:"LocalBinlogConfig,omitnil,omitempty" name:"LocalBinlogConfig"`

	// Default binlog retention policy in the region.
	LocalBinlogConfigDefault *LocalBinlogConfigDefault `json:"LocalBinlogConfigDefault,omitnil,omitempty" name:"LocalBinlogConfigDefault"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLocalBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLocalBinlogConfigResponseParams `json:"Response"`
}

func (r *DescribeLocalBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoRequestParams struct {
	// Parameter template ID, which can be obtained through the [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1) API.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID, which can be obtained through the [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1) API.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoResponseParams struct {
	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The parameter template corresponds to the instance version. Valid values: 5.5, 5.6, 5.7, 8.0.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Number of parameters in the parameter template
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter details
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// Parameter template description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Type of the parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine. Supported values include "InnoDB", "RocksDB".
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateInfoResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// Engine version. Query all if default. Valid values: 5.5, 5.6, 5.7, 8.0.
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Engine type. Query all if default. Valid values: InnoDB, RocksDB. Case-insensitive.
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// Template name. Query all if default. Support fuzzy matching.
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID. Query all if default.
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Engine version. Query all if default. Valid values: 5.5, 5.6, 5.7, 8.0.
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Engine type. Query all if default. Valid values: InnoDB, RocksDB. Case-insensitive.
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// Template name. Query all if default. Support fuzzy matching.
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID. Query all if default.
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineVersions")
	delete(f, "EngineTypes")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// Number of parameter templates of the user.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter template details.
	Items []*ParamTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Project ID. You can obtain it through the [DescribeProjects](https://www.tencentcloud.com/document/api/651/78725?from_cn_redirect=1) API.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID. You can obtain it through the [DescribeProjects](https://www.tencentcloud.com/document/api/651/78725?from_cn_redirect=1) API.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group details.
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// Number of security group rules
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyCustomConfRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries returned per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeProxyCustomConfRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries returned per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeProxyCustomConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyCustomConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyCustomConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyCustomConfResponseParams struct {
	// Number of proxy configurations
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// proxy configuration
	//
	// Deprecated: CustomConf is deprecated.
	CustomConf *CustomConfig `json:"CustomConf,omitnil,omitempty" name:"CustomConf"`

	// Weight limit
	WeightRule *Rule `json:"WeightRule,omitnil,omitempty" name:"WeightRule"`

	// proxy configuration
	CustomConfInfo []*CustomConfig `json:"CustomConfInfo,omitnil,omitempty" name:"CustomConfInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxyCustomConfResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyCustomConfResponseParams `json:"Response"`
}

func (r *DescribeProxyCustomConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyCustomConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySupportParamRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeProxySupportParamRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeProxySupportParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySupportParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySupportParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySupportParamResponseParams struct {
	// Proxy supports the maximum version
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// Whether connection pool is supported
	SupportPool *bool `json:"SupportPool,omitnil,omitempty" name:"SupportPool"`

	// Minimum value of the connection pool
	PoolMin *uint64 `json:"PoolMin,omitnil,omitempty" name:"PoolMin"`

	// Maximum value of connection pool
	PoolMax *uint64 `json:"PoolMax,omitnil,omitempty" name:"PoolMax"`

	// Whether transaction split is supported
	SupportTransSplit *bool `json:"SupportTransSplit,omitnil,omitempty" name:"SupportTransSplit"`

	// Minimum proxy version that supports connection pool
	SupportPoolMinVersion *string `json:"SupportPoolMinVersion,omitnil,omitempty" name:"SupportPoolMinVersion"`

	// Minimum proxy version supporting transaction split
	SupportTransSplitMinVersion *string `json:"SupportTransSplitMinVersion,omitnil,omitempty" name:"SupportTransSplitMinVersion"`

	// Whether setting as read-only is supported.
	SupportReadOnly *bool `json:"SupportReadOnly,omitnil,omitempty" name:"SupportReadOnly"`

	// Whether to automatically balance the load
	SupportAutoLoadBalance *bool `json:"SupportAutoLoadBalance,omitnil,omitempty" name:"SupportAutoLoadBalance"`

	// Whether support access mode
	SupportAccessMode *bool `json:"SupportAccessMode,omitnil,omitempty" name:"SupportAccessMode"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySupportParamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySupportParamResponseParams `json:"Response"`
}

func (r *DescribeProxySupportParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySupportParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRemoteBackupConfigRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRemoteBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeRemoteBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRemoteBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRemoteBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRemoteBackupConfigResponseParams struct {
	// Remote backup retention period in days
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// Remote data backup. Valid values:`off` (disable), `on` (enable).
	RemoteBackupSave *string `json:"RemoteBackupSave,omitnil,omitempty" name:"RemoteBackupSave"`

	// Remote log backup. Valid values: `off` (disable), `on` (enable). Only when the parameter `RemoteBackupSave` is `on`, the `RemoteBinlogSave` parameter can be set to `on`.
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitnil,omitempty" name:"RemoteBinlogSave"`

	// List of configured remote backup regions
	RemoteRegion []*string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`

	// List of remote backup regions
	RegionList []*string `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRemoteBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRemoteBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeRemoteBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRemoteBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoGroupsRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRoGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeRoGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoGroupsResponseParams struct {
	// RO group information array. An instance can associate with multiple RO groups.
	RoGroups []*RoGroup `json:"RoGroups,omitnil,omitempty" name:"RoGroups"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoGroupsResponseParams `json:"Response"`
}

func (r *DescribeRoGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoMinScaleRequestParams struct {
	// Read-only instance ID in the format of "cdbro-c1nl9rpv". Its value is the same as the instance ID in the TencentDB Console. This parameter and the `MasterInstanceId` parameter cannot both be empty.
	RoInstanceId *string `json:"RoInstanceId,omitnil,omitempty" name:"RoInstanceId"`

	// Primary instance ID in the format of "cdbro-c1nl9rpv". Its value is the same as the instance ID in the TencentDB Console. This parameter and the `RoInstanceId` parameter cannot both be empty. Note: when the parameters are passed in with `RoInstanceId`, the return value refers to the minimum specification to which a read-only instance can be upgraded; when the parameters are passed in with `MasterInstanceId` but without `RoInstanceId`, the return value refers to the minimum purchasable specification for a read-only instance.
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`
}

type DescribeRoMinScaleRequest struct {
	*tchttp.BaseRequest
	
	// Read-only instance ID in the format of "cdbro-c1nl9rpv". Its value is the same as the instance ID in the TencentDB Console. This parameter and the `MasterInstanceId` parameter cannot both be empty.
	RoInstanceId *string `json:"RoInstanceId,omitnil,omitempty" name:"RoInstanceId"`

	// Primary instance ID in the format of "cdbro-c1nl9rpv". Its value is the same as the instance ID in the TencentDB Console. This parameter and the `RoInstanceId` parameter cannot both be empty. Note: when the parameters are passed in with `RoInstanceId`, the return value refers to the minimum specification to which a read-only instance can be upgraded; when the parameters are passed in with `MasterInstanceId` but without `RoInstanceId`, the return value refers to the minimum purchasable specification for a read-only instance.
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`
}

func (r *DescribeRoMinScaleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoMinScaleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoInstanceId")
	delete(f, "MasterInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoMinScaleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoMinScaleResponseParams struct {
	// Memory size in MB.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size in GB.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoMinScaleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoMinScaleResponseParams `json:"Response"`
}

func (r *DescribeRoMinScaleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoMinScaleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackRangeTimeRequestParams struct {
	// Instance ID list. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether the clone instance and the source instance are in the same AZ. Valid values: `true` (yes), `false` (no).
	IsRemoteZone *string `json:"IsRemoteZone,omitnil,omitempty" name:"IsRemoteZone"`

	// The region of the clone instance, such as `ap-guangzhou`.
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`
}

type DescribeRollbackRangeTimeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Whether the clone instance and the source instance are in the same AZ. Valid values: `true` (yes), `false` (no).
	IsRemoteZone *string `json:"IsRemoteZone,omitnil,omitempty" name:"IsRemoteZone"`

	// The region of the clone instance, such as `ap-guangzhou`.
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`
}

func (r *DescribeRollbackRangeTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackRangeTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "IsRemoteZone")
	delete(f, "BackupRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackRangeTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackRangeTimeResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Returned parameter information.
	Items []*InstanceRollbackRangeTime `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRollbackRangeTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackRangeTimeResponseParams `json:"Response"`
}

func (r *DescribeRollbackRangeTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackRangeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTaskDetailRequestParams struct {
	// Instance ID, which is the same as the instance ID displayed in the TencentDB Console. You can use the [DescribeDBInstances API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Async task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Pagination parameter. Number of records returned per request. Default value: 20. Maximum value: 100 is recommended.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRollbackTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is the same as the instance ID displayed in the TencentDB Console. You can use the [DescribeDBInstances API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Async task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Pagination parameter. Number of records returned per request. Default value: 20. Maximum value: 100 is recommended.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRollbackTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTaskDetailResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Rollback task details.
	Items []*RollbackTask `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRollbackTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeRollbackTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	// Description: Fill in either the instance ID or the read-only group ID. To query the SSL activation status of two-node or three-node instances, enter the instance ID parameter. Single-node (cloud disk) and cloud disk edition instances do not support enabling SSL, so queries are not supported.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only group ID. Obtain through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	// Description: Fill in either the instance ID or the read-only group ID. To query the SSL activation status of a read-only instance or read-only group, fill in the RoGroupId parameter. Note that you should always enter the read-only group ID. Single-node (cloud disk) and cloud disk edition instances do not support enabling SSL, so they do not support querying.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	// Description: Fill in either the instance ID or the read-only group ID. To query the SSL activation status of two-node or three-node instances, enter the instance ID parameter. Single-node (cloud disk) and cloud disk edition instances do not support enabling SSL, so queries are not supported.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only group ID. Obtain through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	// Description: Fill in either the instance ID or the read-only group ID. To query the SSL activation status of a read-only instance or read-only group, fill in the RoGroupId parameter. Note that you should always enter the read-only group ID. Single-node (cloud disk) and cloud disk edition instances do not support enabling SSL, so they do not support querying.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

func (r *DescribeSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusResponseParams struct {
	// Whether to enable SSL. ON represents enabled, OFF represents not enabled.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Certificate download URL.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSLStatusResponseParams `json:"Response"`
}

func (r *DescribeSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogDataRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Session start timestamp. For example, 1585142640.
	// Description: This parameter is a timestamp in seconds.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp. Example: 1585142640.
	// Description: This parameter is a timestamp in seconds.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Client `Host` list.
	UserHosts []*string `json:"UserHosts,omitnil,omitempty" name:"UserHosts"`

	// Client username list.
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// Accessed database list.
	DataBases []*string `json:"DataBases,omitnil,omitempty" name:"DataBases"`

	// Sorting field. Currently supported fields and their meanings are as follows. Default value is Timestamp.
	// 1. Timestamp: SQL execution time
	// 2. QueryTime: SQL execution duration (seconds)
	// 3. LockTime: Lock duration (seconds)
	// 4. RowsExamined: Number of scanned rows
	// 5. RowsSent: Result set row count
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Ascending or descending order. Valid values: "ASC" - Ascending order, "DESC" - Descending order. Default value: "ASC".
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Offset. The default is 0, and the maximum is 9999.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of records returned in a single use, default is 100, maximum is 800.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`

	// Node ID.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Session start timestamp. For example, 1585142640.
	// Description: This parameter is a timestamp in seconds.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp. Example: 1585142640.
	// Description: This parameter is a timestamp in seconds.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Client `Host` list.
	UserHosts []*string `json:"UserHosts,omitnil,omitempty" name:"UserHosts"`

	// Client username list.
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// Accessed database list.
	DataBases []*string `json:"DataBases,omitnil,omitempty" name:"DataBases"`

	// Sorting field. Currently supported fields and their meanings are as follows. Default value is Timestamp.
	// 1. Timestamp: SQL execution time
	// 2. QueryTime: SQL execution duration (seconds)
	// 3. LockTime: Lock duration (seconds)
	// 4. RowsExamined: Number of scanned rows
	// 5. RowsSent: Result set row count
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Ascending or descending order. Valid values: "ASC" - Ascending order, "DESC" - Descending order. Default value: "ASC".
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Offset. The default is 0, and the maximum is 9999.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of records returned in a single use, default is 100, maximum is 800.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`

	// Node ID.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

func (r *DescribeSlowLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserHosts")
	delete(f, "UserNames")
	delete(f, "DataBases")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstType")
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogDataResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Queried records.
	Items []*SlowLogItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogDataResponseParams `json:"Response"`
}

func (r *DescribeSlowLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset, starting from `0`. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset, starting from `0`. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Default value: 20. Minimum value: 1. Maximum value: 1000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// Number of eligible slow logs.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of eligible slow logs.
	Items []*SlowLogInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedPrivilegesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSupportedPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeSupportedPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedPrivilegesResponseParams struct {
	// Global permissions supported by the instance
	GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitnil,omitempty" name:"GlobalSupportedPrivileges"`

	// Database permissions supported by the instance.
	DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitnil,omitempty" name:"DatabaseSupportedPrivileges"`

	// Table permissions supported by the instance.
	TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitnil,omitempty" name:"TableSupportedPrivileges"`

	// Column permissions supported by the instance.
	ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitnil,omitempty" name:"ColumnSupportedPrivileges"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSupportedPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportedPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeSupportedPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableColumnsRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the [Query Instance List](https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1) API. Its value is the InstanceId field in the output parameter.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name. Obtain through the [Query Database](https://www.tencentcloud.com/document/api/236/17493?from_cn_redirect=1) API.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Name of the table in the database.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type DescribeTableColumnsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the [Query Instance List](https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1) API. Its value is the InstanceId field in the output parameter.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name. Obtain through the [Query Database](https://www.tencentcloud.com/document/api/236/17493?from_cn_redirect=1) API.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Name of the table in the database.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

func (r *DescribeTableColumnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableColumnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Database")
	delete(f, "Table")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableColumnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableColumnsResponseParams struct {
	// Total number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Returned database column information.
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableColumnsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableColumnsResponseParams `json:"Response"`
}

func (r *DescribeTableColumnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableColumnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items returned per request. Default value: 20. Maximum value: 5000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Regular expression for matching table names, which complies with the rules at MySQL's official website
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items returned per request. Default value: 20. Maximum value: 5000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Regular expression for matching table names, which complies with the rules at MySQL's official website
	TableRegexp *string `json:"TableRegexp,omitnil,omitempty" name:"TableRegexp"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Database")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TableRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// Number of eligible tables.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of a table.
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsOfInstanceIdsRequestParams struct {
	// Instance list. Instance ID can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API. The length of the array passed in is not limited.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Defaults to 15.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// Instance list. Instance ID can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API. The length of the array passed in is not limited.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page size. Defaults to 15.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTagsOfInstanceIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsOfInstanceIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsOfInstanceIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsOfInstanceIdsResponseParams struct {
	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Instance tag information.
	Rows []*TagsInfoOfInstance `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagsOfInstanceIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsOfInstanceIdsResponseParams `json:"Response"`
}

func (r *DescribeTagsOfInstanceIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsOfInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ID of an async task request, i.e., `AsyncRequestId` returned by relevant TencentDB operations.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Task type. If no value is passed in, all task types will be queried. Valid values:
	// 1 - rolling back a database;
	// 2 - performing an SQL operation;
	// 3 - importing data;
	// 5 - setting a parameter;
	// 6 - initializing a TencentDB instance;
	// 7 - restarting a TencentDB instance;
	// 8 - enabling GTID of a TencentDB instance;
	// 9 - upgrading a read-only instance;
	// 10 - rolling back databases in batches;
	// 11 - upgrading a primary instance;
	// 12 - deleting a TencentDB table;
	// 13 - promoting a disaster recovery instance.
	TaskTypes []*int64 `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// Task status. If no value is passed in, all task statuses will be queried. Valid values:
	// -1 - undefined;
	// 0 - initializing;
	// 1 - running;
	// 2 - succeeded;
	// 3 - failed;
	// 4 - terminated;
	// 5 - deleted;
	// 6 - paused.
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// ID of an async task request, i.e., `AsyncRequestId` returned by relevant TencentDB operations.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Task type. If no value is passed in, all task types will be queried. Valid values:
	// 1 - rolling back a database;
	// 2 - performing an SQL operation;
	// 3 - importing data;
	// 5 - setting a parameter;
	// 6 - initializing a TencentDB instance;
	// 7 - restarting a TencentDB instance;
	// 8 - enabling GTID of a TencentDB instance;
	// 9 - upgrading a read-only instance;
	// 10 - rolling back databases in batches;
	// 11 - upgrading a primary instance;
	// 12 - deleting a TencentDB table;
	// 13 - promoting a disaster recovery instance.
	TaskTypes []*int64 `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// Task status. If no value is passed in, all task statuses will be queried. Valid values:
	// -1 - undefined;
	// 0 - initializing;
	// 1 - running;
	// 2 - succeeded;
	// 3 - failed;
	// 4 - terminated;
	// 5 - deleted;
	// 6 - paused.
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeBegin *string `json:"StartTimeBegin,omitnil,omitempty" name:"StartTimeBegin"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2017-12-31 10:40:01. It is used for queries by time range.
	StartTimeEnd *string `json:"StartTimeEnd,omitnil,omitempty" name:"StartTimeEnd"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "TaskTypes")
	delete(f, "TaskStatus")
	delete(f, "StartTimeBegin")
	delete(f, "StartTimeEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Information of an instance task.
	Items []*TaskDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeWindowRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimeWindowResponseParams struct {
	// List of maintenance time windows on Monday.
	Monday []*string `json:"Monday,omitnil,omitempty" name:"Monday"`

	// List of maintenance time windows on Tuesday.
	Tuesday []*string `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// List of maintenance time windows on Wednesday.
	Wednesday []*string `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// List of maintenance time windows on Thursday.
	Thursday []*string `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// List of maintenance time windows on Friday.
	Friday []*string `json:"Friday,omitnil,omitempty" name:"Friday"`

	// List of maintenance time windows on Saturday.
	Saturday []*string `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// List of maintenance time windows on Sunday.
	Sunday []*string `json:"Sunday,omitnil,omitempty" name:"Sunday"`

	// Maximum data delay threshold
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimeWindowResponseParams `json:"Response"`
}

func (r *DescribeTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadedFilesRequestParams struct {
	// File path. `OwnerUin` information of the root account should be entered in this field.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUploadedFilesRequest struct {
	*tchttp.BaseRequest
	
	// File path. `OwnerUin` information of the root account should be entered in this field.
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUploadedFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadedFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Path")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadedFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadedFilesResponseParams struct {
	// Number of eligible SQL files.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of returned SQL files.
	Items []*SqlFileInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUploadedFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadedFilesResponseParams `json:"Response"`
}

func (r *DescribeUploadedFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadedFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCpuInfo struct {
	// Average instance CPU utilization
	Rate []*DeviceCpuRateInfo `json:"Rate,omitnil,omitempty" name:"Rate"`

	// CPU monitoring data of the instance
	Load []*int64 `json:"Load,omitnil,omitempty" name:"Load"`
}

type DeviceCpuRateInfo struct {
	// CPU core number
	CpuCore *int64 `json:"CpuCore,omitnil,omitempty" name:"CpuCore"`

	// CPU utilization
	Rate []*int64 `json:"Rate,omitnil,omitempty" name:"Rate"`
}

type DeviceDiskInfo struct {
	// Time percentage of IO operations per second
	IoRatioPerSec []*int64 `json:"IoRatioPerSec,omitnil,omitempty" name:"IoRatioPerSec"`

	// Average wait time of device I/O operations * 100 in milliseconds. For example, if the value is 201, the average wait time of I/O operations is 201/100 = 2.1 milliseconds.
	IoWaitTime []*int64 `json:"IoWaitTime,omitnil,omitempty" name:"IoWaitTime"`

	// Average number of read operations completed by the disk per second * 100. For example, if the value is 2,002, the average number of read operations completed by the disk per second is 2,002/100=20.2.
	Read []*int64 `json:"Read,omitnil,omitempty" name:"Read"`

	// Average number of write operations completed by the disk per second * 100. For example, if the value is 30,001, the average number of write operations completed by the disk per second is 30,001/100=300.01.
	Write []*int64 `json:"Write,omitnil,omitempty" name:"Write"`

	// Disk capacity. Each value is comprised of two data, with the first data representing the used capacity and the second one representing the total disk capacity.
	CapacityRatio []*int64 `json:"CapacityRatio,omitnil,omitempty" name:"CapacityRatio"`
}

type DeviceMemInfo struct {
	// Total memory size in KB, which is the value of `total` in the `Mem:` in the `free` command
	Total []*int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Used memory size in KB, which is the value of `used` in the `Mem:` row in the `free` command
	Used []*int64 `json:"Used,omitnil,omitempty" name:"Used"`
}

type DeviceNetInfo struct {
	// Number of TCP connections
	Conn []*int64 `json:"Conn,omitnil,omitempty" name:"Conn"`

	// ENI inbound packets per second
	PackageIn []*int64 `json:"PackageIn,omitnil,omitempty" name:"PackageIn"`

	// ENI outbound packets per second
	PackageOut []*int64 `json:"PackageOut,omitnil,omitempty" name:"PackageOut"`

	// Inbound traffic in Kbps
	FlowIn []*int64 `json:"FlowIn,omitnil,omitempty" name:"FlowIn"`

	// Outbound traffic in Kbps
	FlowOut []*int64 `json:"FlowOut,omitnil,omitempty" name:"FlowOut"`
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// Security group ID. Obtain through the [DescribeDBSecurityGroups](https://www.tencentcloud.com/document/api/236/15854?from_cn_redirect=1) API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance ID list, an array consisting of one or more instance IDs, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// When importing a read-only instance ID, the default operation is performed on the corresponding security group of the read-only group. If necessary to operate the security group of the read-only instance ID, set this input parameter to True. Default False.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID. Obtain through the [DescribeDBSecurityGroups](https://www.tencentcloud.com/document/api/236/15854?from_cn_redirect=1) API.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Instance ID list, an array consisting of one or more instance IDs, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// When importing a read-only instance ID, the default operation is performed on the corresponding security group of the read-only group. If necessary to operate the security group of the read-only instance ID, set this input parameter to True. Default False.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	delete(f, "ForReadonlyInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskTypeConfigItem struct {
	// Type of instance corresponding to the disk. Only support single node (cloud disk) and cloud disk edition.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// List of disk types to choose.
	DiskType []*string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

type DrInfo struct {
	// Disaster recovery instance status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// AZ information
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region information
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance sync status. Possible returned values include:
	// 0 - disaster recovery not synced;
	// 1 - disaster recovery syncing;
	// 2 - disaster recovery synced successfully;
	// 3 - disaster recovery sync failed;
	// 4 - repairing disaster recovery sync;
	SyncStatus *int64 `json:"SyncStatus,omitnil,omitempty" name:"SyncStatus"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type ErrlogItem struct {
	// Error occurrence time. Timestamp in seconds.
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Error details
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ImportRecord struct {
	// Status value. 0 - Initializing, 1 - Running, 2 - Operation successful, 3 - Operation failure.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status value. Task exception when the value is negative.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Execution time, unit: seconds.
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backend task ID
	WorkId *string `json:"WorkId,omitnil,omitempty" name:"WorkId"`

	// Name of the file to be imported
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Execution progress, measurement unit: percentage.
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// File size, unit: byte.
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Task execution information
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Task ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Name of the table to be imported
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Async task request ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type Inbound struct {
	// Policy, which can be ACCEPT or DROP
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Source IP or IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// The direction of the rule, which is INPUT for inbound rules
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`

	// Address module
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// Rule ID, rule ID of the nested security group
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Rule description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type InstEventInfo struct {
	// Event name.
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// Event status.
	EventStatus *string `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// Event occurrence time.
	OccurTime *string `json:"OccurTime,omitnil,omitempty" name:"OccurTime"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type InstanceAuditLogFilters struct {
	// Filter condition. The search conditions are supported as follows:
	// 
	// Include/Exclude, and Include/Exclude (segment dimension) can be used to search for:
	// `sql` - SQL details.
	// 
	// `Equal to`, `Not equal to`, `Include`, and `Exclude` can be used to search for:
	// `host` - Client IP,
	// `user` - Username,
	// `DBName` - Database name.
	// 
	// `Equal to` and `Not equal to` can be used to search for:
	// `sqlType` - SQL type,
	// `errCode` - Error code,
	// `threadId` - Thread ID.
	// 
	// Range search is supported for:
	// `execTime`- Execution time (μs),
	// `lockWaitTime` - Lock wait time (μs),
	// `ioWaitTime` - IO wait time (μs),
	// `trxLivingTime` - Transaction duration (μs),
	// `cpuTime` - CPU time (μs),
	// `checkRows` - Number of scanned rows,
	// `affectRows` - Number of affected rows,
	// `sentRows` - Number of returned rows.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter. Valid values:
	// `WINC` - Include (segment dimension)
	// `WEXC` - Exclude (segment dimension)
	// `INC` - Include,
	// `EXC` - Exclude,
	// `EQS` - Equal to,
	// `NEQ` - Not equal to.
	// `RA` - Range
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// The filter value. In a reverse query, multiple values are in an "AND" relationship; while in a forward query, multiple values are in an "OR" relationship.
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InstanceDbAuditStatus struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit status. ON - Audit is enabled, OFF - Audit is disabled.
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// Task status. Valid values: 0 - No task; 1 - Enabling audit; 2 - Disabling audit.
	AuditTask *uint64 `json:"AuditTask,omitnil,omitempty" name:"AuditTask"`

	// Log retention period. Valid values:7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High-frequency storage period. Valid values:3 - 3 days;7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Low-frequency storage period (in days). This equals the log retention period minus the high-frequency storage period.
	LowLogExpireDay *uint64 `json:"LowLogExpireDay,omitnil,omitempty" name:"LowLogExpireDay"`

	// Log storage volume (in GB).
	BillingAmount *float64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// High-frequency storage volume (in GB).
	HighRealStorage *float64 `json:"HighRealStorage,omitnil,omitempty" name:"HighRealStorage"`

	// Low-frequency storage volume (in GB).
	LowRealStorage *float64 `json:"LowRealStorage,omitnil,omitempty" name:"LowRealStorage"`

	// Whether full audit is enabled. true - Full audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Time when the audit service was activated.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Related information about the instance.
	InstanceInfo *AuditInstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// Total storage volume (in GB).
	RealStorage *float64 `json:"RealStorage,omitnil,omitempty" name:"RealStorage"`

	// Whether an audit policy is configured.
	OldRule *bool `json:"OldRule,omitnil,omitempty" name:"OldRule"`

	// Rule template applied to the instance.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Trial status.
	TrialStatus *string `json:"TrialStatus,omitnil,omitempty" name:"TrialStatus"`

	// Trial start time.
	TrialStartTime *int64 `json:"TrialStartTime,omitnil,omitempty" name:"TrialStartTime"`

	// Trial duration.
	TrialDuration *int64 `json:"TrialDuration,omitnil,omitempty" name:"TrialDuration"`

	// Trial end time.
	TrialCloseTime *int64 `json:"TrialCloseTime,omitnil,omitempty" name:"TrialCloseTime"`

	// Log query time limit during the trial period.
	TrialDescribeLogHours *int64 `json:"TrialDescribeLogHours,omitnil,omitempty" name:"TrialDescribeLogHours"`
}

type InstanceInfo struct {
	// <p>Public network status. Possible returned values: 0 - External network not enabled; 1 - Public network enabled; 2 - Public network disabled</p>
	WanStatus *int64 `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// <p>AZ information</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Initialization flag. Possible returned values: 0 - uninitialized; 1 - initialized.</p>
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// <p>Read-only vip information. This field is available only for read-only instances with separate instance access enabled.</p>
	RoVipInfo *RoVipInfo `json:"RoVipInfo,omitnil,omitempty" name:"RoVipInfo"`

	// <p>Memory capacity, in MB.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Instance status. Valid values: 0: creating; 1: running; 4: isolation operation in progress; 5: isolated.</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>VPC ID, for example: 51102</p>
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>Secondary server information</p>
	SlaveInfo *SlaveInfo `json:"SlaveInfo,omitnil,omitempty" name:"SlaveInfo"`

	// <p>Instance ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Disk capacity, in GB.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Auto-renewal flag. Possible returned values: 0 - auto-renewal is not enabled; 1 - auto-renewal is enabled; 2 - automatic renewal is disabled.</p>
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>Data replication mode. 0 - async replication; 1 - semi-sync replication; 2 - strong sync replication</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Read-only group detailed information</p>
	RoGroups []*RoGroup `json:"RoGroups,omitnil,omitempty" name:"RoGroups"`

	// <p>Subnet ID, for example: 2333</p>
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Instance type. Possible returned values: 1 - Primary instance; 2 - Disaster recovery instance; 3 - Read-only instance.</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Project ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Regional information</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Instance expiration time</p>
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// <p>Availability Zone Deployment method. Valid values: 0 - single availability zone; 1 - multi-availability zone.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>Instance task status. 0 - No tasks, 1 - Upgrading, 2 - Data import, 3 - Opening Slave, 4 - Public network access enabling, 5 - Batch operation executing, 6 - Rolling back, 7 - Public network access disabling, 8 - Password modification, 9 - Renaming instance, 10 - Restarting, 12 - Self-built migration, 13 - Database deletion, 14 - Disaster recovery instance creation sync, 15 - Upgrade pending switch, 16 - Upgrade and switch, 17 - Upgrade and switch completed</p>
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// <p>Detailed information about the primary instance.</p>
	MasterInfo *MasterInfo `json:"MasterInfo,omitnil,omitempty" name:"MasterInfo"`

	// <p>Instance type</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Kernel version</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Disaster recovery instance details</p>
	DrInfo []*DrInfo `json:"DrInfo,omitnil,omitempty" name:"DrInfo"`

	// <p>public network domain name</p>
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// <p>Public network port number</p>
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Billing type
	PayType *int64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// <p>Instance creation time</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Instance IP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Port number</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Whether disk write is locked (data write volume of the instance exceeds disk quota). 0 - Unlocked 1 - Locked</p>
	CdbError *int64 `json:"CdbError,omitnil,omitempty" name:"CdbError"`

	// <p>Private network descriptor, for example: "vpc-5v8wn9mg"</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet descriptor, such as "subnet-1typ0s7d"</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Physical ID</p>
	PhysicalId *string `json:"PhysicalId,omitnil,omitempty" name:"PhysicalId"`

	// <p>Core count</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Queries per second.</p>
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// <p>Chinese Name of Availability Zone</p>
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// <p>Physical machine model</p>
	DeviceClass *string `json:"DeviceClass,omitnil,omitempty" name:"DeviceClass"`

	// <p>Placement group ID</p>
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// <p>Availability zone ID</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Number of nodes</p>
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// <p>Tag list</p>
	TagList []*TagInfoItem `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>Engine type</p>
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// <p>Maximum delay threshold</p>
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// <p>Instance disk type. Only CLOUD disk edition and single-node (CLOUD disk) instances will return a valid value.<br>Description:</p><ol><li>If "DiskType": "CLOUD_HSSD" is returned, it indicates that the instance disk type is enhanced SSD CLOUD disk.</li><li>If "DiskType": "CLOUD_SSD" is returned, it indicates that the instance disk type is SSD CLOUD Block Storage.</li><li>If "DiskType": "" is returned and the DeviceType parameter value is UNIVERSAL or EXCLUSIVE, it means that the instance uses local SSD.</li></ol>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>Current number of CPU cores for scale-out.</p>
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// <p>Cloud Disk Edition instance node information</p>
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <p>Analysis engine node list</p>
	AnalysisNodeInfos []*AnalysisNodeInfo `json:"AnalysisNodeInfos,omitnil,omitempty" name:"AnalysisNodeInfos"`

	// <p>Device bandwidth, in G. This parameter is valid only when DeviceClass is not empty. For example, 25 means the current device bandwidth is 25G; 10 means the current device bandwidth is 10G.</p>
	DeviceBandwidth *uint64 `json:"DeviceBandwidth,omitnil,omitempty" name:"DeviceBandwidth"`

	// <p>Instance termination protection status. on indicates enabled; otherwise, the protection is disabled.</p>
	DestroyProtect *string `json:"DestroyProtect,omitnil,omitempty" name:"DestroyProtect"`

	// <p>TDSQL engine parameters</p>
	CpuModel *string `json:"CpuModel,omitnil,omitempty" name:"CpuModel"`

	// <p>Analysis engine instance version upgrade information</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	AnalysisUpgradeVersionInfo *UpgradeAnalysisInstanceVersionInfo `json:"AnalysisUpgradeVersionInfo,omitnil,omitempty" name:"AnalysisUpgradeVersionInfo"`
}

type InstanceRebootTime struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Expected restart time, unit: second.
	TimeInSeconds *int64 `json:"TimeInSeconds,omitnil,omitempty" name:"TimeInSeconds"`
}

type InstanceRollbackRangeTime struct {
	// Query database error codes. 0 - Normal, 1600001 - Internal error, 1600003 - Input parameter exception, 1600009 - Instance does not exist, 1624001 - DB access exception.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Queries database error message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// List of instance IDs. An instance ID is in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time range available for rollback
	Times []*RollbackTimeRange `json:"Times,omitnil,omitempty" name:"Times"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// Request ID of the async task. Use this ID to query the outcome of the async task. (This returned field is currently abandoned. The quarantined state of instances can be queried through the API to query instances.)
	//
	// Deprecated: AsyncRequestId is deprecated.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocalBinlogConfig struct {
	// Local binlog retention duration. Valid values: [6,168].
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Local binlog space utilization. Valid values: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type LocalBinlogConfigDefault struct {
	// Local binlog retention duration. Valid values: [6,168].
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Local binlog space utilization. Valid values: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type LogRuleTemplateInfo struct {
	// Template ID.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Rule template name
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *string `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Rule template change status. Valid values: 0 - Not changed, 1 - changed, 2 - deleted.
	RuleTemplateStatus *int64 `json:"RuleTemplateStatus,omitnil,omitempty" name:"RuleTemplateStatus"`
}

type LogToCLSConfig struct {
	// Delivery status on or turn off
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CLS Logset ID
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// Log topic ID
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// Region of the CLS service
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`
}

type MasterInfo struct {
	// <p>Regional information</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Region ID</p>
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>Availability zone ID.</p>
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>AZ information</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>Instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Instance long ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>Instance status</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Instance name</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>Instance type</p>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>Task status.</p>
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// <p>Memory capacity</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Disk capacity</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Instance model</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Queries per second.</p>
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// <p>VPC ID</p>
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>subnet ID</p>
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>Dedicated cluster ID</p>
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// <p>Dedicated cluster name</p>
	ExClusterName *string `json:"ExClusterName,omitnil,omitempty" name:"ExClusterName"`
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TDSQL for MySQL accounts. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Remark information of the database account. Input limit: 255 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TDSQL for MySQL accounts. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Remark information of the database account. Input limit: 255 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountDescriptionResponseParams `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountMaxUserConnectionsRequestParams struct {
	// Cloud Database account. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum connections of the account. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type ModifyAccountMaxUserConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// Cloud Database account. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum connections of the account. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

func (r *ModifyAccountMaxUserConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountMaxUserConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Accounts")
	delete(f, "InstanceId")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountMaxUserConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountMaxUserConnectionsResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountMaxUserConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountMaxUserConnectionsResponseParams `json:"Response"`
}

func (r *ModifyAccountMaxUserConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountMaxUserConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPasswordRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New password of the database account. It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (_+-&=!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`

	// TDSQL for MySQL accounts. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Deprecated.
	//
	// Deprecated: SkipValidatePassword is deprecated.
	SkipValidatePassword *bool `json:"SkipValidatePassword,omitnil,omitempty" name:"SkipValidatePassword"`
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New password of the database account. It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (_+-&=!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`

	// TDSQL for MySQL accounts. Obtain through the [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1) API.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Deprecated.
	SkipValidatePassword *bool `json:"SkipValidatePassword,omitnil,omitempty" name:"SkipValidatePassword"`
}

func (r *ModifyAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewPassword")
	delete(f, "Accounts")
	delete(f, "SkipValidatePassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPasswordResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPasswordResponseParams `json:"Response"`
}

func (r *ModifyAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database account, includes users and domain name. Obtain through the API [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1).
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Global permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "PROCESS", "DROP", "REFERENCES", "INDEX", "ALTER", "SHOW DATABASES", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER", "CREATE USER", "RELOAD", "REPLICATION CLIENT", "REPLICATION SLAVE".
	// Note: When “ModifyAction” is empty, if `GlobalPrivileges` is not passed in, it indicates the global permission will become ineffective.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: When “ModifyAction” is empty, if `DatabasePrivileges` is not passed in, it indicates the permission of the database will become ineffective.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table permission in the database. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE VIEW", "SHOW VIEW", "TRIGGER".
	// Note: When “ModifyAction” is empty, if `TablePrivileges` is not passed in, it indicates the permission of the table will become ineffective.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// Column permission in the table. Valid values: "SELECT", "INSERT", "UPDATE", "REFERENCES".
	// Note: When “ModifyAction” is empty, if `ColumnPrivileges` is not passed in, it indicates the permission of the column will become ineffective.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// When this parameter is not empty, it indicates that the permission will be modified. Valid values: `grant` (grant permission), `revoke` (revoke permission)
	ModifyAction *string `json:"ModifyAction,omitnil,omitempty" name:"ModifyAction"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database account, includes users and domain name. Obtain through the API [DescribeAccounts](https://www.tencentcloud.com/document/api/236/17499?from_cn_redirect=1).
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Global permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "PROCESS", "DROP", "REFERENCES", "INDEX", "ALTER", "SHOW DATABASES", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER", "CREATE USER", "RELOAD", "REPLICATION CLIENT", "REPLICATION SLAVE".
	// Note: When “ModifyAction” is empty, if `GlobalPrivileges` is not passed in, it indicates the global permission will become ineffective.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// Database permission. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES", "EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", "TRIGGER".
	// Note: When “ModifyAction” is empty, if `DatabasePrivileges` is not passed in, it indicates the permission of the database will become ineffective.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// Table permission in the database. Valid values: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE VIEW", "SHOW VIEW", "TRIGGER".
	// Note: When “ModifyAction” is empty, if `TablePrivileges` is not passed in, it indicates the permission of the table will become ineffective.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// Column permission in the table. Valid values: "SELECT", "INSERT", "UPDATE", "REFERENCES".
	// Note: When “ModifyAction” is empty, if `ColumnPrivileges` is not passed in, it indicates the permission of the column will become ineffective.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// When this parameter is not empty, it indicates that the permission will be modified. Valid values: `grant` (grant permission), `revoke` (revoke permission)
	ModifyAction *string `json:"ModifyAction,omitnil,omitempty" name:"ModifyAction"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ModifyAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditConfigRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit log retention period. Valid values:7 - One week;30 - One month;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Whether to disable the audit service. Valid values: true - Disable; false - Do not disable. Default value: false.Notes:1. When the audit service is disabled, your audit logs and files will be deleted, and all audit policies for this instance will be removed.2. At least one of CloseAudit and LogExpireDay must be provided. If both are provided, CloseAudit takes priority.3. You can use this parameter to disable the audit service. Once disabled, the audit service cannot be re-enabled via this API.
	CloseAudit *bool `json:"CloseAudit,omitnil,omitempty" name:"CloseAudit"`

	// High-frequency audit log retention period. Valid values:7 - One week;30 - One month;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`
}

type ModifyAuditConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit log retention period. Valid values:7 - One week;30 - One month;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Whether to disable the audit service. Valid values: true - Disable; false - Do not disable. Default value: false.Notes:1. When the audit service is disabled, your audit logs and files will be deleted, and all audit policies for this instance will be removed.2. At least one of CloseAudit and LogExpireDay must be provided. If both are provided, CloseAudit takes priority.3. You can use this parameter to disable the audit service. Once disabled, the audit service cannot be re-enabled via this API.
	CloseAudit *bool `json:"CloseAudit,omitnil,omitempty" name:"CloseAudit"`

	// High-frequency audit log retention period. Valid values:7 - One week;30 - One month;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`
}

func (r *ModifyAuditConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "CloseAudit")
	delete(f, "HighLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditConfigResponseParams `json:"Response"`
}

func (r *ModifyAuditConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesRequestParams struct {
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Modified audit rule.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Modified rule template name.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Modified rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

type ModifyAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Modified audit rule.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Modified rule template name.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Modified rule template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled.
	AlarmPolicy *uint64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`
}

func (r *ModifyAuditRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleTemplateIds")
	delete(f, "RuleFilters")
	delete(f, "RuleTemplateName")
	delete(f, "Description")
	delete(f, "AlarmLevel")
	delete(f, "AlarmPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditRuleTemplatesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditRuleTemplatesResponseParams `json:"Response"`
}

func (r *ModifyAuditRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period. Valid values:7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High-frequency log retention period. Default value: 7. This value must be less than or equal to LogExpireDay. Valid values include:3 - 3 days;7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Modifies the instance audit rule to Full audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Deprecated.
	//
	// Deprecated: AuditRuleFilters is deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period. Valid values:7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High-frequency log retention period. Default value: 7. This value must be less than or equal to LogExpireDay. Valid values include:3 - 3 days;7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Modifies the instance audit rule to Full audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.com/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditAll")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagRequestParams struct {
	// Instance ID, in the format of cdb-c1nl9rpv. This is identical to the instance ID displayed on the TencentDB console.
	// Description: Multiple instance IDs can be entered for modification. The json format is as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled).
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format of cdb-c1nl9rpv. This is identical to the instance ID displayed on the TencentDB console.
	// Description: Multiple instance IDs can be entered for modification. The json format is as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled).
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

func (r *ModifyAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoRenewFlagResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention time of the data backup file, in days.
	// 1. MySQL two-node, three-node, and cloud disk edition data backup files can be retained for 7-1830 days.
	// 2. MySQL single-node (cloud disk) data backup files can be retained for 7-30 days.
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// (This parameter will be disused. The `BackupTimeWindow` parameter is recommended.) Backup time range in the format of 02:00-06:00, with the start time and end time on the hour. Valid values: 00:00-12:00, 02:00-06:00, 06:00-10:00, 10:00-14:00, 14:00-18:00, 18:00-22:00, 22:00-02:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Automatic backup mode. Only `physical` (physical cold backup) is supported
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// binlog retention time in days.
	// 1. MySQL two-node, three-node, and cloud disk log backup files can be retained for 7 to 3650 days.
	// 2. MySQL single-node (cloud disk) log backup files can be retained for 7-30 days.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// Backup time window; for example, to set up backup between 10:00 and 14:00 on every Tuesday and Sunday, you should set this parameter as follows: {"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"} (Note: You can set up backup on different days, but the backup time windows need to be the same. If this field is set, the `StartTime` field will be ignored)
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// Periodic backup retention switch. off - periodic backup retention policy is not enabled, on - periodic backup retention policy is enabled. Default is off.
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitnil,omitempty" name:"EnableBackupPeriodSave"`

	// Switch for long-term backup retention (This field can be ignored, for its feature hasn’t been launched). Valid values: `off` (disable), `on` (enable). Default value: `off`. Once enabled, the parameters (BackupPeriodSaveDays, BackupPeriodSaveInterval, and BackupPeriodSaveCount) will be invalid.
	EnableBackupPeriodLongTermSave *string `json:"EnableBackupPeriodLongTermSave,omitnil,omitempty" name:"EnableBackupPeriodLongTermSave"`

	// Maximum days of archive backup retention. Valid range: 90-3650. Default value: 1080.
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitnil,omitempty" name:"BackupPeriodSaveDays"`

	// Archive backup retention period. Valid values: `weekly` (a week), `monthly` (a month), `quarterly` (a quarter), `yearly` (a year). Default value: `monthly`.
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitnil,omitempty" name:"BackupPeriodSaveInterval"`

	// Number of archive backups. Minimum value: `1`, Maximum value: Number of non-archive backups in archive backup retention period. Default value: `1`.
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitnil,omitempty" name:"BackupPeriodSaveCount"`

	// The start time in the format of yyyy-mm-dd HH:MM:SS, which is used to enable archive backup retention policy.
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitnil,omitempty" name:"StartBackupPeriodSaveDate"`

	// Whether the data backup/archive policy is enabled. off - disabled, on - enabled. If not specified, remain unchanged.
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// The period (in days) of how long a data backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// The period (in days) of how long a log backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// Whether to enable log backup archive strategy. off - off, on - on. If not specified, remain unchanged.
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// Whether to enable the standard storage policy for data backup. off - disabled, on - enabled. If not specified, it remains unchanged.
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// The period (in days) of how long a data backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// Whether to enable log backup standard storage policy. off - off, on - on. If not specified, remain unchanged.
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// The period (in days) of how long a log backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention time of the data backup file, in days.
	// 1. MySQL two-node, three-node, and cloud disk edition data backup files can be retained for 7-1830 days.
	// 2. MySQL single-node (cloud disk) data backup files can be retained for 7-30 days.
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// (This parameter will be disused. The `BackupTimeWindow` parameter is recommended.) Backup time range in the format of 02:00-06:00, with the start time and end time on the hour. Valid values: 00:00-12:00, 02:00-06:00, 06:00-10:00, 10:00-14:00, 14:00-18:00, 18:00-22:00, 22:00-02:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Automatic backup mode. Only `physical` (physical cold backup) is supported
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// binlog retention time in days.
	// 1. MySQL two-node, three-node, and cloud disk log backup files can be retained for 7 to 3650 days.
	// 2. MySQL single-node (cloud disk) log backup files can be retained for 7-30 days.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// Backup time window; for example, to set up backup between 10:00 and 14:00 on every Tuesday and Sunday, you should set this parameter as follows: {"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"} (Note: You can set up backup on different days, but the backup time windows need to be the same. If this field is set, the `StartTime` field will be ignored)
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// Periodic backup retention switch. off - periodic backup retention policy is not enabled, on - periodic backup retention policy is enabled. Default is off.
	EnableBackupPeriodSave *string `json:"EnableBackupPeriodSave,omitnil,omitempty" name:"EnableBackupPeriodSave"`

	// Switch for long-term backup retention (This field can be ignored, for its feature hasn’t been launched). Valid values: `off` (disable), `on` (enable). Default value: `off`. Once enabled, the parameters (BackupPeriodSaveDays, BackupPeriodSaveInterval, and BackupPeriodSaveCount) will be invalid.
	EnableBackupPeriodLongTermSave *string `json:"EnableBackupPeriodLongTermSave,omitnil,omitempty" name:"EnableBackupPeriodLongTermSave"`

	// Maximum days of archive backup retention. Valid range: 90-3650. Default value: 1080.
	BackupPeriodSaveDays *int64 `json:"BackupPeriodSaveDays,omitnil,omitempty" name:"BackupPeriodSaveDays"`

	// Archive backup retention period. Valid values: `weekly` (a week), `monthly` (a month), `quarterly` (a quarter), `yearly` (a year). Default value: `monthly`.
	BackupPeriodSaveInterval *string `json:"BackupPeriodSaveInterval,omitnil,omitempty" name:"BackupPeriodSaveInterval"`

	// Number of archive backups. Minimum value: `1`, Maximum value: Number of non-archive backups in archive backup retention period. Default value: `1`.
	BackupPeriodSaveCount *int64 `json:"BackupPeriodSaveCount,omitnil,omitempty" name:"BackupPeriodSaveCount"`

	// The start time in the format of yyyy-mm-dd HH:MM:SS, which is used to enable archive backup retention policy.
	StartBackupPeriodSaveDate *string `json:"StartBackupPeriodSaveDate,omitnil,omitempty" name:"StartBackupPeriodSaveDate"`

	// Whether the data backup/archive policy is enabled. off - disabled, on - enabled. If not specified, remain unchanged.
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// The period (in days) of how long a data backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// The period (in days) of how long a log backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// Whether to enable log backup archive strategy. off - off, on - on. If not specified, remain unchanged.
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// Whether to enable the standard storage policy for data backup. off - disabled, on - enabled. If not specified, it remains unchanged.
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// The period (in days) of how long a data backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// Whether to enable log backup standard storage policy. off - off, on - on. If not specified, remain unchanged.
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// The period (in days) of how long a log backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`
}

func (r *ModifyBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ExpireDays")
	delete(f, "StartTime")
	delete(f, "BackupMethod")
	delete(f, "BinlogExpireDays")
	delete(f, "BackupTimeWindow")
	delete(f, "EnableBackupPeriodSave")
	delete(f, "EnableBackupPeriodLongTermSave")
	delete(f, "BackupPeriodSaveDays")
	delete(f, "BackupPeriodSaveInterval")
	delete(f, "BackupPeriodSaveCount")
	delete(f, "StartBackupPeriodSaveDate")
	delete(f, "EnableBackupArchive")
	delete(f, "BackupArchiveDays")
	delete(f, "BinlogArchiveDays")
	delete(f, "EnableBinlogArchive")
	delete(f, "EnableBackupStandby")
	delete(f, "BackupStandbyDays")
	delete(f, "EnableBinlogStandby")
	delete(f, "BinlogStandbyDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// Valid values: `NoLimit` (backups can be downloaded over both private and public networks with any IPs), `LimitOnlyIntranet` (backups can be downloaded over the private network with any private IPs), `Customize` (backups can be downloaded over specified VPCs with specified IPs). The `LimitVpc` and `LimitIp` parameters are valid only when this parameter is set to `Customize`.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Valid value: `In` (backups can only be downloaded over the VPCs specified in `LimitVpc`). Default value: `In`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Valid values: `In` (backups can only be downloaded with the IPs specified in `LimitIp`), `NotIn` (backups cannot be downloaded with the IPs specified in `LimitIp`). Default value: `In`.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPCs used to restrict backup download.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// IPs used to restrict backup download.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Valid values: `NoLimit` (backups can be downloaded over both private and public networks with any IPs), `LimitOnlyIntranet` (backups can be downloaded over the private network with any private IPs), `Customize` (backups can be downloaded over specified VPCs with specified IPs). The `LimitVpc` and `LimitIp` parameters are valid only when this parameter is set to `Customize`.
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// Valid value: `In` (backups can only be downloaded over the VPCs specified in `LimitVpc`). Default value: `In`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// Valid values: `In` (backups can only be downloaded with the IPs specified in `LimitIp`), `NotIn` (backups cannot be downloaded with the IPs specified in `LimitIp`). Default value: `In`.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// VPCs used to restrict backup download.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// IPs used to restrict backup download.
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupEncryptionStatusRequestParams struct {
	// Instance ID in the format of cdb-XXXX, which is the same as that displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Default encryption status for the new auto-generated physical backup files. Valid values: `on`, `off`.
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`

	// Set the default encryption status of the newly-added automated log backup file for the instance. Available values are on or off.
	BinlogEncryptionStatus *string `json:"BinlogEncryptionStatus,omitnil,omitempty" name:"BinlogEncryptionStatus"`
}

type ModifyBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-XXXX, which is the same as that displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Default encryption status for the new auto-generated physical backup files. Valid values: `on`, `off`.
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`

	// Set the default encryption status of the newly-added automated log backup file for the instance. Available values are on or off.
	BinlogEncryptionStatus *string `json:"BinlogEncryptionStatus,omitnil,omitempty" name:"BinlogEncryptionStatus"`
}

func (r *ModifyBackupEncryptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupEncryptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EncryptionStatus")
	delete(f, "BinlogEncryptionStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupEncryptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupEncryptionStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupEncryptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupEncryptionStatusResponseParams `json:"Response"`
}

func (r *ModifyBackupEncryptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupEncryptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressDescRequestParams struct {
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. You can obtain it through the API [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1).
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ModifyCdbProxyAddressDescRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. You can obtain it through the API [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1).
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *ModifyCdbProxyAddressDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCdbProxyAddressDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressDescResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCdbProxyAddressDescResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCdbProxyAddressDescResponseParams `json:"Response"`
}

func (r *ModifyCdbProxyAddressDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressVipAndVPortRequestParams struct {
	// Proxy group ID. Obtain through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. Obtain through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// VPC ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Private subnet ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// IP. If not specified, the system will assign an available IP under subnet.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port. Default value 3306, value ranges from 1024 to 65535.
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Old IP valid hours. Measurement unit: hr, default value: 24, value ranges from 0 to 168.
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
}

type ModifyCdbProxyAddressVipAndVPortRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID. Obtain through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. Obtain through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// VPC ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Private subnet ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// IP. If not specified, the system will assign an available IP under subnet.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port. Default value 3306, value ranges from 1024 to 65535.
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Old IP valid hours. Measurement unit: hr, default value: 24, value ranges from 0 to 168.
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
}

func (r *ModifyCdbProxyAddressVipAndVPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressVipAndVPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "Vip")
	delete(f, "VPort")
	delete(f, "ReleaseDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCdbProxyAddressVipAndVPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyAddressVipAndVPortResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCdbProxyAddressVipAndVPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCdbProxyAddressVipAndVPortResponseParams `json:"Response"`
}

func (r *ModifyCdbProxyAddressVipAndVPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyAddressVipAndVPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyParamRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Connection pool threshold. Value ranges from above 0 to less than or equal to 300.
	// Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be equal to or greater than MySQL 8.0 20230630.
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
}

type ModifyCdbProxyParamRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Connection pool threshold. Value ranges from above 0 to less than or equal to 300.
	// Note: If you need to use the database proxy connection pool capability, the kernel minor version of the MySQL 8.0 primary instance must be equal to or greater than MySQL 8.0 20230630.
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
}

func (r *ModifyCdbProxyParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "ConnectionPoolLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCdbProxyParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCdbProxyParamResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCdbProxyParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCdbProxyParamResponseParams `json:"Response"`
}

func (r *ModifyCdbProxyParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCdbProxyParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceLogToCLSRequestParams struct {
	// <p>Instance ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1">DescribeDBInstances</a> API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type. Error: error log, slowlog: slow log.</p>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Delivery status. ON: enabled, OFF: disabled.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Whether required to create logset. Default to false.</p>
	CreateLogset *bool `json:"CreateLogset,omitnil,omitempty" name:"CreateLogset"`

	// <p>Logset name when creating a logset; logset ID when selecting an existing logset. Empty by default.<br>Description: When the Status parameter is ON, either the Logset or LogTopic parameter must be filled.</p>
	Logset *string `json:"Logset,omitnil,omitempty" name:"Logset"`

	// <p>Whether required to create log topic. Default to false.</p>
	CreateLogTopic *bool `json:"CreateLogTopic,omitnil,omitempty" name:"CreateLogTopic"`

	// <p>Enter a log topic name when creating a log topic, or enter a log topic ID when selecting an existing log topic. Empty by default.<br>Description: When the Status parameter is set to ON, either the Logset or LogTopic parameter must be specified.</p>
	LogTopic *string `json:"LogTopic,omitnil,omitempty" name:"LogTopic"`

	// <p>Valid period of the log topic. Default value: 30 days if left empty. Maximum value: 3600 days.</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Whether to create an index when creating a log topic. Defaults to false.</p>
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`

	// <p>CLS region. If left empty, it defaults to the Region parameter value.</p>
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// <p>Selectable when creating a log topic. Cannot exceed 10 tags</p>
	ResourceTags []*TagInfoItem `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type ModifyDBInstanceLogToCLSRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, which can be obtained through the <a href="https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1">DescribeDBInstances</a> API.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Log type. Error: error log, slowlog: slow log.</p>
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// <p>Delivery status. ON: enabled, OFF: disabled.</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Whether required to create logset. Default to false.</p>
	CreateLogset *bool `json:"CreateLogset,omitnil,omitempty" name:"CreateLogset"`

	// <p>Logset name when creating a logset; logset ID when selecting an existing logset. Empty by default.<br>Description: When the Status parameter is ON, either the Logset or LogTopic parameter must be filled.</p>
	Logset *string `json:"Logset,omitnil,omitempty" name:"Logset"`

	// <p>Whether required to create log topic. Default to false.</p>
	CreateLogTopic *bool `json:"CreateLogTopic,omitnil,omitempty" name:"CreateLogTopic"`

	// <p>Enter a log topic name when creating a log topic, or enter a log topic ID when selecting an existing log topic. Empty by default.<br>Description: When the Status parameter is set to ON, either the Logset or LogTopic parameter must be specified.</p>
	LogTopic *string `json:"LogTopic,omitnil,omitempty" name:"LogTopic"`

	// <p>Valid period of the log topic. Default value: 30 days if left empty. Maximum value: 3600 days.</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>Whether to create an index when creating a log topic. Defaults to false.</p>
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`

	// <p>CLS region. If left empty, it defaults to the Region parameter value.</p>
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// <p>Selectable when creating a log topic. Cannot exceed 10 tags</p>
	ResourceTags []*TagInfoItem `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

func (r *ModifyDBInstanceLogToCLSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceLogToCLSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "Status")
	delete(f, "CreateLogset")
	delete(f, "Logset")
	delete(f, "CreateLogTopic")
	delete(f, "LogTopic")
	delete(f, "Period")
	delete(f, "CreateIndex")
	delete(f, "ClsRegion")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceLogToCLSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceLogToCLSResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceLogToCLSResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceLogToCLSResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceLogToCLSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceLogToCLSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceModesRequestParams struct {
	// <p>Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The mode of cloud databases currently only supports input "protectMode" to modify the Primary-standby sync mode.</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>Data synchronization mode, available values: 0 - async replication, 1 - semi-sync replication, 2 - strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`
}

type ModifyDBInstanceModesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>The mode of cloud databases currently only supports input "protectMode" to modify the Primary-standby sync mode.</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>Data synchronization mode, available values: 0 - async replication, 1 - semi-sync replication, 2 - strong sync replication.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`
}

func (r *ModifyDBInstanceModesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceModesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Mode")
	delete(f, "ProtectMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceModesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceModesResponseParams struct {
	// <p>Request ID of the asynchronous task. Use this ID to query the outcome of the async task.</p>
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceModesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceModesResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceModesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceModesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modified instance name, which can only contain digits, English uppercase and lowercase letters, Chinese, and special characters -_./()[]（）+=:：@. Its length cannot exceed 60.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modified instance name, which can only contain digits, English uppercase and lowercase letters, Chinese, and special characters -_./()[]（）+=:：@. Its length cannot exceed 60.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceProjectRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the query instance list API (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1). The value is the InstanceId field in the output parameter.
	// Description: Multiple instance IDs can be entered for modification. The json format is as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// ID of the project to which instance belongs can be queried on the Projects page in the account center.
	// Description: This item is required.
	NewProjectId *int64 `json:"NewProjectId,omitnil,omitempty" name:"NewProjectId"`
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the query instance list API (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1). The value is the InstanceId field in the output parameter.
	// Description: Multiple instance IDs can be entered for modification. The json format is as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// ID of the project to which instance belongs can be queried on the Projects page in the account center.
	// Description: This item is required.
	NewProjectId *int64 `json:"NewProjectId,omitnil,omitempty" name:"NewProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "NewProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceProjectResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of security group IDs to modify, an array of security group IDs. It can be obtained through the API [DescribeDBSecurityGroups](https://www.tencentcloud.com/document/product/236/15854?from_cn_redirect=1). The input security group ID array has no length limit.
	// **Note**: This input parameter performs a full replacement on all existing collections but not an incremental update. To modify it, import the expected full collections.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// When importing a read-only instance ID, the default operation is performed on the corresponding security group of the read-only group. If necessary to operate the security group of the read-only instance ID, set this input parameter to True. Default False.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// When updating the read-only group of a cloud disk edition instance, specify the instance ID in InstanceId and this parameter to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of security group IDs to modify, an array of security group IDs. It can be obtained through the API [DescribeDBSecurityGroups](https://www.tencentcloud.com/document/product/236/15854?from_cn_redirect=1). The input security group ID array has no length limit.
	// **Note**: This input parameter performs a full replacement on all existing collections but not an incremental update. To modify it, import the expected full collections.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// When importing a read-only instance ID, the default operation is performed on the corresponding security group of the read-only group. If necessary to operate the security group of the read-only instance ID, set this input parameter to True. Default False.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`

	// When updating the read-only group of a cloud disk edition instance, specify the instance ID in InstanceId and this parameter to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	delete(f, "ForReadonlyInstance")
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceVipVportRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv, cdbro-c2nl9rpv, or cdbrg-c3nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872) API to query the ID, which is the value of the `InstanceId` output parameter.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target IP address.
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// Destination port. Support scope: [1024-65535].
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Unified subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Repossession duration in hours for old IP in the original network when changing from classic network to VPC or changing the VPC subnet. Value range: 0–168. Default value: `24`.
	ReleaseDuration *int64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`

	// When updating the read-only group of a cluster edition instance, specify the instance id in InstanceId and this parameter is required to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv, cdbro-c2nl9rpv, or cdbrg-c3nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872) API to query the ID, which is the value of the `InstanceId` output parameter.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target IP address.
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// Destination port. Support scope: [1024-65535].
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Unified subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Repossession duration in hours for old IP in the original network when changing from classic network to VPC or changing the VPC subnet. Value range: 0–168. Default value: `24`.
	ReleaseDuration *int64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`

	// When updating the read-only group of a cluster edition instance, specify the instance id in InstanceId and this parameter is required to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

func (r *ModifyDBInstanceVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "ReleaseDuration")
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceVipVportResponseParams struct {
	// Asynchronous Task ID. (This returned field is currently abandoned)
	//
	// Deprecated: AsyncRequestId is deprecated.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceVipVportResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamRequestParams struct {
	// Instance ID list, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value).
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Template ID. At least one of ParamList and TemplateId must be provided. It can be obtained through the API [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1).
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// When to perform the parameter adjustment task. Default value: 0. Valid values: 0 - execute immediately, 1 - execute during window. When its value is 1, only one instance ID can be passed in (i.e., only one `InstanceIds` can be passed in).
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// Whether to sync the parameters to read-only instance of the source instance. Valid values: `true` (not sync), `false` (sync). Default value: `false`.
	NotSyncRo *bool `json:"NotSyncRo,omitnil,omitempty" name:"NotSyncRo"`

	// Whether to sync the parameters to disaster recovery instance of the source instance. Valid values: `true` (not sync), `false` (sync). Default value: `false`.
	NotSyncDr *bool `json:"NotSyncDr,omitnil,omitempty" name:"NotSyncDr"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value).
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Template ID. At least one of ParamList and TemplateId must be provided. It can be obtained through the API [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1).
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// When to perform the parameter adjustment task. Default value: 0. Valid values: 0 - execute immediately, 1 - execute during window. When its value is 1, only one instance ID can be passed in (i.e., only one `InstanceIds` can be passed in).
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// Whether to sync the parameters to read-only instance of the source instance. Valid values: `true` (not sync), `false` (sync). Default value: `false`.
	NotSyncRo *bool `json:"NotSyncRo,omitnil,omitempty" name:"NotSyncRo"`

	// Whether to sync the parameters to disaster recovery instance of the source instance. Valid values: `true` (not sync), `false` (sync). Default value: `false`.
	NotSyncDr *bool `json:"NotSyncDr,omitnil,omitempty" name:"NotSyncDr"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ParamList")
	delete(f, "TemplateId")
	delete(f, "WaitSwitch")
	delete(f, "NotSyncRo")
	delete(f, "NotSyncDr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// Async task ID, which can be used to query task progress.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePasswordComplexityRequestParams struct {
	// Instance ID of the instance for which the password complexity needs to be modified. The [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API can be called to obtain it.
	// Description: Support multiple instance IDs for modification.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Options to modify password complexity. Each option is written in metric combinations. A group includes Name and CurrentValue. Among them, Name refers to the parameter name of the corresponding option, and CurrentValue represents the parameter value. For example: [{"Name": "validate_password.length", "CurrentValue": "10"}] means changing the minimum number of characters in a password to 10.
	// Description: The options for modifying password complexity vary by database version of instances as follows.
	// 1. MySQL 8.0:
	// The option validate_password.policy means the switch for password complexity. A value of LOW means disabling; a value of MEDIUM means enabling.
	// The option validate_password.length indicates the minimum number of characters for the total code length.
	// The option validate_password.mixed_case_count indicates the minimum number of lowercase and uppercase letters.
	// Option validate_password.number_count indicates the minimum number of digits.
	// The option validate_password.special_char_count indicates the minimum number of special characters.
	// 2. MySQL 5.6,MySQL 5.7:
	// The option validate_password_policy means the password complexity switch. A value of LOW means disabling; a value of MEDIUM means enabling.
	// The option validate_password_length indicates the minimum number of characters for the total code length.
	// The option validate_password_mixed_case_count means the minimum number of uppercase and lowercase letters.
	// Option validate_password_number_count means the minimum number of digits.
	// Option validate_password_special_char_count indicates the minimum number of special characters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyInstancePasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID of the instance for which the password complexity needs to be modified. The [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API can be called to obtain it.
	// Description: Support multiple instance IDs for modification.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Options to modify password complexity. Each option is written in metric combinations. A group includes Name and CurrentValue. Among them, Name refers to the parameter name of the corresponding option, and CurrentValue represents the parameter value. For example: [{"Name": "validate_password.length", "CurrentValue": "10"}] means changing the minimum number of characters in a password to 10.
	// Description: The options for modifying password complexity vary by database version of instances as follows.
	// 1. MySQL 8.0:
	// The option validate_password.policy means the switch for password complexity. A value of LOW means disabling; a value of MEDIUM means enabling.
	// The option validate_password.length indicates the minimum number of characters for the total code length.
	// The option validate_password.mixed_case_count indicates the minimum number of lowercase and uppercase letters.
	// Option validate_password.number_count indicates the minimum number of digits.
	// The option validate_password.special_char_count indicates the minimum number of special characters.
	// 2. MySQL 5.6,MySQL 5.7:
	// The option validate_password_policy means the password complexity switch. A value of LOW means disabling; a value of MEDIUM means enabling.
	// The option validate_password_length indicates the minimum number of characters for the total code length.
	// The option validate_password_mixed_case_count means the minimum number of uppercase and lowercase letters.
	// Option validate_password_number_count means the minimum number of digits.
	// Option validate_password_special_char_count indicates the minimum number of special characters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyInstancePasswordComplexityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePasswordComplexityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancePasswordComplexityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancePasswordComplexityResponseParams struct {
	// Async task ID, which can be used to query task progress.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstancePasswordComplexityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancePasswordComplexityResponseParams `json:"Response"`
}

func (r *ModifyInstancePasswordComplexityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancePasswordComplexityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTagRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tags to add or modify. ReplaceTags or DeleteTags is mandatory to fill in one.
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitnil,omitempty" name:"ReplaceTags"`

	// Tag to delete. ReplaceTags or DeleteTags is mandatory to fill in one.
	DeleteTags []*TagInfo `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tags to add or modify. ReplaceTags or DeleteTags is mandatory to fill in one.
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitnil,omitempty" name:"ReplaceTags"`

	// Tag to delete. ReplaceTags or DeleteTags is mandatory to fill in one.
	DeleteTags []*TagInfo `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`
}

func (r *ModifyInstanceTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTagResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceTagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceTagResponseParams `json:"Response"`
}

func (r *ModifyInstanceTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalBinlogConfigRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Local binlog retention duration. Values for different instances are as follows:
	// 1. The local binlog retention duration (hr) for cloud disk edition instances, dual-node instances, and three-node instances defaults to 120, with a range of 6 - 168.
	// 2. The retention duration of local binlog for disaster recovery instance defaults to 120 hr, with a range of 120 - 168.
	// 3. The retention duration (hr) of local binlog for a single-node cloud disk instance defaults to 120, with a range of 0 - 168.
	// 4. If a dual-node instance or three-node instance has no disaster recovery instance, the retention duration (hr) of local binlog for the primary instance ranges from 6 to 168. If a dual-node instance or three-node instance has a disaster recovery instance, or you want to add a disaster recovery instance to a dual-node instance or three-node instance, to avoid synchronization exception, the retention duration (hr) of local binlog for the primary instance cannot be set to less than 120 hr, ranging from 120 to 168.
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Local binlog space utilization. Valid values: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type ModifyLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Local binlog retention duration. Values for different instances are as follows:
	// 1. The local binlog retention duration (hr) for cloud disk edition instances, dual-node instances, and three-node instances defaults to 120, with a range of 6 - 168.
	// 2. The retention duration of local binlog for disaster recovery instance defaults to 120 hr, with a range of 120 - 168.
	// 3. The retention duration (hr) of local binlog for a single-node cloud disk instance defaults to 120, with a range of 0 - 168.
	// 4. If a dual-node instance or three-node instance has no disaster recovery instance, the retention duration (hr) of local binlog for the primary instance ranges from 6 to 168. If a dual-node instance or three-node instance has a disaster recovery instance, or you want to add a disaster recovery instance to a dual-node instance or three-node instance, to avoid synchronization exception, the retention duration (hr) of local binlog for the primary instance cannot be set to less than 120 hr, ranging from 120 to 168.
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Local binlog space utilization. Valid values: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

func (r *ModifyLocalBinlogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalBinlogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SaveHours")
	delete(f, "MaxUsage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLocalBinlogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalBinlogConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLocalBinlogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLocalBinlogConfigResponseParams `json:"Response"`
}

func (r *ModifyLocalBinlogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalBinlogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameOrDescByDpIdRequestParams struct {
	// Placement group ID
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// Name of a placement group, which can contain up to 60 characters. The placement group name and description can’t be empty.
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// Description of a placement group, which can contain up to 200 characters. The placement group name and description can’t be empty.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyNameOrDescByDpIdRequest struct {
	*tchttp.BaseRequest
	
	// Placement group ID
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// Name of a placement group, which can contain up to 60 characters. The placement group name and description can’t be empty.
	DeployGroupName *string `json:"DeployGroupName,omitnil,omitempty" name:"DeployGroupName"`

	// Description of a placement group, which can contain up to 200 characters. The placement group name and description can’t be empty.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyNameOrDescByDpIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameOrDescByDpIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployGroupId")
	delete(f, "DeployGroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNameOrDescByDpIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameOrDescByDpIdResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNameOrDescByDpIdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNameOrDescByDpIdResponseParams `json:"Response"`
}

func (r *ModifyNameOrDescByDpIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameOrDescByDpIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// Template ID, which can be obtained through the [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1) API.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name, supports numbers, English uppercase and lowercase letters, Chinese, and special characters _-./()[]+=:@, and the length cannot exceed 60.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description (up to 255 characters)
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID, which can be obtained through the [DescribeParamTemplates](https://www.tencentcloud.com/document/api/236/32659?from_cn_redirect=1) API.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name, supports numbers, English uppercase and lowercase letters, Chinese, and special characters _-./()[]+=:@, and the length cannot exceed 60.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description (up to 255 characters)
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParamTemplateResponseParams `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectModeRequestParams struct {
	// Data replication method, defaults to 0. Supported values include: 0 - asynchronous replication, 1 - semi-sync replication, 2 - strong sync replication.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyProtectModeRequest struct {
	*tchttp.BaseRequest
	
	// Data replication method, defaults to 0. Supported values include: 0 - asynchronous replication, 1 - semi-sync replication, 2 - strong sync replication.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyProtectModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectMode")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProtectModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectModeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProtectModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProtectModeResponseParams `json:"Response"`
}

func (r *ModifyProtectModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRemoteBackupConfigRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Remote data backup. Valid values:`off` (disable), `on` (enable).
	RemoteBackupSave *string `json:"RemoteBackupSave,omitnil,omitempty" name:"RemoteBackupSave"`

	// Remote log backup. Valid values: `off` (disable), `on` (enable). Only when the parameter `RemoteBackupSave` is `on`, the `RemoteBinlogSave` parameter can be set to `on`.
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitnil,omitempty" name:"RemoteBinlogSave"`

	// The custom backup region list
	RemoteRegion []*string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`

	// Remote backup retention period in days
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`
}

type ModifyRemoteBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Remote data backup. Valid values:`off` (disable), `on` (enable).
	RemoteBackupSave *string `json:"RemoteBackupSave,omitnil,omitempty" name:"RemoteBackupSave"`

	// Remote log backup. Valid values: `off` (disable), `on` (enable). Only when the parameter `RemoteBackupSave` is `on`, the `RemoteBinlogSave` parameter can be set to `on`.
	RemoteBinlogSave *string `json:"RemoteBinlogSave,omitnil,omitempty" name:"RemoteBinlogSave"`

	// The custom backup region list
	RemoteRegion []*string `json:"RemoteRegion,omitnil,omitempty" name:"RemoteRegion"`

	// Remote backup retention period in days
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`
}

func (r *ModifyRemoteBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRemoteBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RemoteBackupSave")
	delete(f, "RemoteBinlogSave")
	delete(f, "RemoteRegion")
	delete(f, "ExpireDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRemoteBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRemoteBackupConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRemoteBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRemoteBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyRemoteBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRemoteBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupInfoRequestParams struct {
	// ID of the RO group, which can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// RO group details.
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitnil,omitempty" name:"RoGroupInfo"`

	// Weight of instances in the RO group. If modification is needed to set the weight mode of the RO group to user-defined mode (custom), this parameter must be set, and the weight value of each read-only instance needs to be set. The RO instance ID can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitnil,omitempty" name:"RoWeightValues"`

	// Whether to rebalance the loads of read-only replicas in the RO group. Valid values: `1` (yes), `0` (no). Default value: `0`. If this parameter is set to `1`, connections to the read-only replicas in the RO group will be interrupted transiently. Please ensure that your application has a reconnection mechanism.
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitnil,omitempty" name:"IsBalanceRoLoad"`

	// This field has been deprecated.
	//
	// Deprecated: ReplicationDelayTime is deprecated.
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of the RO group, which can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// RO group details.
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitnil,omitempty" name:"RoGroupInfo"`

	// Weight of instances in the RO group. If modification is needed to set the weight mode of the RO group to user-defined mode (custom), this parameter must be set, and the weight value of each read-only instance needs to be set. The RO instance ID can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitnil,omitempty" name:"RoWeightValues"`

	// Whether to rebalance the loads of read-only replicas in the RO group. Valid values: `1` (yes), `0` (no). Default value: `0`. If this parameter is set to `1`, connections to the read-only replicas in the RO group will be interrupted transiently. Please ensure that your application has a reconnection mechanism.
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitnil,omitempty" name:"IsBalanceRoLoad"`

	// This field has been deprecated.
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
}

func (r *ModifyRoGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoGroupId")
	delete(f, "RoGroupInfo")
	delete(f, "RoWeightValues")
	delete(f, "IsBalanceRoLoad")
	delete(f, "ReplicationDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupInfoResponseParams struct {
	// Asynchronous Task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoGroupInfoResponseParams `json:"Response"`
}

func (r *ModifyRoGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupVipVportRequestParams struct {
	// ID of the RO group.
	UGroupId *string `json:"UGroupId,omitnil,omitempty" name:"UGroupId"`

	// Target IP address.
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// Target Port.
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`
}

type ModifyRoGroupVipVportRequest struct {
	*tchttp.BaseRequest
	
	// ID of the RO group.
	UGroupId *string `json:"UGroupId,omitnil,omitempty" name:"UGroupId"`

	// Target IP address.
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// Target Port.
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`
}

func (r *ModifyRoGroupVipVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupVipVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UGroupId")
	delete(f, "DstIp")
	delete(f, "DstPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoGroupVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoGroupVipVportResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoGroupVipVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoGroupVipVportResponseParams `json:"Response"`
}

func (r *ModifyRoGroupVipVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoGroupVipVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTimeWindowRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The modified maintenance time slot. Among them, each time period is in the format of 10:00-12:00. The start and end time is aligned by half hour. The shortest is half hour and the longest is three hours. Up to two time periods can be set. The start and end time ranges from [00:00, 24:00].
	// Description: The following is an example of setting two time periods in json.
	// [
	//     "01:00-01:30",
	//     "02:00-02:30"
	//   ]
	TimeRanges []*string `json:"TimeRanges,omitnil,omitempty" name:"TimeRanges"`

	// Specify which day to modify the maintenance time slot. Possible values are: monday, tuesday, wednesday, thursday, friday, saturday, sunday. If not specified or empty, modify all seven days of the week by default.
	// Description: The json example for modifying more than one day is as follows.
	// [
	//     "monday",
	//     "tuesday"
	//   ]
	Weekdays []*string `json:"Weekdays,omitnil,omitempty" name:"Weekdays"`

	// Data latency threshold (seconds), only applicable to primary instance and disaster recovery instance. No modification by default to keep the original threshold. Value ranges from 1 to 10 integers.
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The modified maintenance time slot. Among them, each time period is in the format of 10:00-12:00. The start and end time is aligned by half hour. The shortest is half hour and the longest is three hours. Up to two time periods can be set. The start and end time ranges from [00:00, 24:00].
	// Description: The following is an example of setting two time periods in json.
	// [
	//     "01:00-01:30",
	//     "02:00-02:30"
	//   ]
	TimeRanges []*string `json:"TimeRanges,omitnil,omitempty" name:"TimeRanges"`

	// Specify which day to modify the maintenance time slot. Possible values are: monday, tuesday, wednesday, thursday, friday, saturday, sunday. If not specified or empty, modify all seven days of the week by default.
	// Description: The json example for modifying more than one day is as follows.
	// [
	//     "monday",
	//     "tuesday"
	//   ]
	Weekdays []*string `json:"Weekdays,omitnil,omitempty" name:"Weekdays"`

	// Data latency threshold (seconds), only applicable to primary instance and disaster recovery instance. No modification by default to keep the original threshold. Value ranges from 1 to 10 integers.
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

func (r *ModifyTimeWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTimeWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TimeRanges")
	delete(f, "Weekdays")
	delete(f, "MaxDelayTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTimeWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTimeWindowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTimeWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTimeWindowResponseParams `json:"Response"`
}

func (r *ModifyTimeWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTimeWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeDistribution struct {
	// Host ID of the Master node of the primary instance or host ID of the read-only instance
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// Host ID where the first Slave node of the primary instance resides
	SlaveNodeOne *string `json:"SlaveNodeOne,omitnil,omitempty" name:"SlaveNodeOne"`

	// Host ID where the second Slave node of the primary instance resides
	SlaveNodeTwo *string `json:"SlaveNodeTwo,omitnil,omitempty" name:"SlaveNodeTwo"`
}

// Predefined struct for user
type OfflineIsolatedInstancesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type OfflineIsolatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *OfflineIsolatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedInstancesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OfflineIsolatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedInstancesResponseParams `json:"Response"`
}

func (r *OfflineIsolatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// CDB instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit log retention period. Supported values include:
	// 7 - A week;
	// 30 - one month
	// 90 - three months;
	// 180 - 6 months;
	// 365 - One year;
	// 1095 - Three years;
	// 1825 - Five years.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High frequency audit log retention period. Default value is 7. This item must take value less than or equal to LogExpireDay. Supported values include:
	// 3 - 3 days;
	// 7 - A week;
	// 30 - one month
	// 90 - three months;
	// 180 - 6 months;
	// 365 - One year;
	// 1095 - Three years;
	// 1825 - Five years.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule (deprecated, no longer effective).
	//
	// Deprecated: AuditRuleFilters is deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit type. true - full audit; default false - rule audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// CDB instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Audit log retention period. Supported values include:
	// 7 - A week;
	// 30 - one month
	// 90 - three months;
	// 180 - 6 months;
	// 365 - One year;
	// 1095 - Three years;
	// 1825 - Five years.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High frequency audit log retention period. Default value is 7. This item must take value less than or equal to LogExpireDay. Supported values include:
	// 3 - 3 days;
	// 7 - A week;
	// 30 - one month
	// 90 - three months;
	// 180 - 6 months;
	// 365 - One year;
	// 1095 - Three years;
	// 1825 - Five years.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule (deprecated, no longer effective).
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit type. true - full audit; default false - rule audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HighLogExpireDay")
	delete(f, "AuditRuleFilters")
	delete(f, "RuleTemplateIds")
	delete(f, "AuditAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceEncryptionRequestParams struct {
	// Cloud database instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Custom key ID, the unique identifier of CMK. If left empty, the automatically generated key KMS-CDB by using Tencent Cloud will be used.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Storage region of the custom key. For example: ap-guangzhou. This parameter is required when KeyId is not empty.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`
}

type OpenDBInstanceEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// Cloud database instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Custom key ID, the unique identifier of CMK. If left empty, the automatically generated key KMS-CDB by using Tencent Cloud will be used.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Storage region of the custom key. For example: ap-guangzhou. This parameter is required when KeyId is not empty.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`
}

func (r *OpenDBInstanceEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceEncryptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KeyId")
	delete(f, "KeyRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBInstanceEncryptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceEncryptionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenDBInstanceEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBInstanceEncryptionResponseParams `json:"Response"`
}

func (r *OpenDBInstanceEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceEncryptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceGTIDRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenDBInstanceGTIDRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *OpenDBInstanceGTIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceGTIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBInstanceGTIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBInstanceGTIDResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenDBInstanceGTIDResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBInstanceGTIDResponseParams `json:"Response"`
}

func (r *OpenDBInstanceGTIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBInstanceGTIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLRequestParams struct {
	// Instance ID. Required when the read-only group ID is empty. Can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only group ID. Required when the instance ID is empty. Can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Required when the read-only group ID is empty. Can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Read-only group ID. Required when the instance ID is empty. Can be obtained through the [DescribeRoGroups](https://www.tencentcloud.com/document/api/236/40939?from_cn_redirect=1) API.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

func (r *OpenSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RoGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// Asynchronous request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenSSLResponse struct {
	*tchttp.BaseResponse
	Response *OpenSSLResponseParams `json:"Response"`
}

func (r *OpenSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanServiceRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the query instance list API (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1). The value is the InstanceId field in the output parameter. The read-only group ID can be passed in.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// When updating the read-only group of a cluster edition instance, specify the instance id in InstanceId and this parameter is required to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

type OpenWanServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the query instance list API (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1). The value is the InstanceId field in the output parameter. The read-only group ID can be passed in.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// When updating the read-only group of a cluster edition instance, specify the instance id in InstanceId and this parameter is required to indicate the operation is for the read-only group. If you perform the operation on the read-write node, this parameter is not required.
	OpResourceId *string `json:"OpResourceId,omitnil,omitempty" name:"OpResourceId"`
}

func (r *OpenWanServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWanServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWanServiceResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenWanServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenWanServiceResponseParams `json:"Response"`
}

func (r *OpenWanServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWanServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {
	// Policy, which can be ACCEPT or DROP
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Destination IP or IP range, such as 172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Port or port range
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// The direction of the rule, which is OUTPUT for inbound rules
	Dir *string `json:"Dir,omitnil,omitempty" name:"Dir"`

	// Address module
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// Rule description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ParamInfo struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ParamRecord struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Parameter name
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Parameter value before modification
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// Parameter value after modification
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// Whether the parameter is modified successfully
	//
	// Deprecated: IsSucess is deprecated.
	IsSucess *bool `json:"IsSucess,omitnil,omitempty" name:"IsSucess"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Indicates whether the parameter is modified successfully.
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`
}

type ParamTemplateInfo struct {
	// parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Instance engine version. Values: 5.5, 5.6, 5.7, and 8.0.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Parameter template type. Valid values: HIGH_STABILITY, HIGH_PERFORMANCE.
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine, values: InnoDB, RocksDB.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type Parameter struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter type. Valid values: `integer`, `enum`, `float`, `string`, `func`
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Whether the database needs to be restarted for the modified parameter to take effect. Value range: 0 (no); 1 (yes)
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Maximum value of the parameter
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`

	// Minimum value of the parameter
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Enumerated values of the parameter. It is null if the parameter is non-enumerated
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// Maximum parameter value, which is valid only when `ParamType` is set to `func`
	MaxFunc *string `json:"MaxFunc,omitnil,omitempty" name:"MaxFunc"`

	// Minimum parameter value, which is valid only when `ParamType` is set to `func`
	MinFunc *string `json:"MinFunc,omitnil,omitempty" name:"MinFunc"`

	// Whether the parameter cannot be modified
	IsNotSupportEdit *bool `json:"IsNotSupportEdit,omitnil,omitempty" name:"IsNotSupportEdit"`
}

type PeriodStrategy struct {
	// Scale-out period
	TimeCycle *TImeCycle `json:"TimeCycle,omitnil,omitempty" name:"TimeCycle"`

	// Time interval
	TimeInterval *TimeInterval `json:"TimeInterval,omitnil,omitempty" name:"TimeInterval"`
}

type ProxyAddress struct {
	// Address ID of the proxy group
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// IP address
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Weight allocation mode.
	// System Auto-Assignment: "system", Custom: "custom"
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to enable delay removal. Parameter value: "true" | "false"
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Minimum retention quantity, minimum value: 0.
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// Delay removal threshold, minimum value: 0
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// Automatically add RO. Value: "true" | "false"
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether it is read-only. Value: "true" | "false".
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// Whether transaction splitting is enabled
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Whether fault migration is enabled
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to enable connection pool
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Read weight distribution of an instance
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	// Access mode
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// Whether automatic CLB is enabled
	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	// Whether to treat libra as a read-only node
	ApNodeAsRoNode *bool `json:"ApNodeAsRoNode,omitnil,omitempty" name:"ApNodeAsRoNode"`

	// libra node fault, whether to forward to other nodes
	ApQueryToOtherNode *bool `json:"ApQueryToOtherNode,omitnil,omitempty" name:"ApQueryToOtherNode"`
}

type ProxyAllocation struct {
	// Proxy node region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ of proxy node region
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Proxy instance allocation
	ProxyInstance []*ProxyInst `json:"ProxyInstance,omitnil,omitempty" name:"ProxyInstance"`
}

type ProxyGroupInfo struct {
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// proxy version
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// Proxy supports edition upgrade
	SupportUpgradeProxyVersion *string `json:"SupportUpgradeProxyVersion,omitnil,omitempty" name:"SupportUpgradeProxyVersion"`

	// Agent status. 0 - Initializing, 1 - Online, 2 - Online - Read-write separation, 3 - Offline, 4 - Terminated.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Agent task status, Upgrading - upgrading, UpgradeTo - upgrade pending switch, UpgradeSwitching - upgrade and switch in progress, ProxyCreateAddress - configuring address, ProxyModifyAddress - changing address, ProxyCloseAddress - closing address.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Proxy group node information
	ProxyNode []*ProxyNode `json:"ProxyNode,omitnil,omitempty" name:"ProxyNode"`

	// Proxy group address information
	ProxyAddress []*ProxyAddress `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`

	// Connection pool threshold
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// Support creating an address
	SupportCreateProxyAddress *bool `json:"SupportCreateProxyAddress,omitnil,omitempty" name:"SupportCreateProxyAddress"`

	// cdb version required for proxy version upgrade
	SupportUpgradeProxyMysqlVersion *string `json:"SupportUpgradeProxyMysqlVersion,omitnil,omitempty" name:"SupportUpgradeProxyMysqlVersion"`
}

type ProxyInst struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type: 1 master primary instance; 2 read-only instance; 3 dr disaster recovery instance; 4 sdr (small disaster recovery) instance
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance status. Valid values: 0: creating; 1: running; 4: isolation; 5: isolated.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Read-only weight. If the weight is automatically assigned by the system, this value does not take effect and only indicates whether the instance is enabled or not.
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Instance region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone to which the instance belongs
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance Node ID
	InstNodeId *string `json:"InstNodeId,omitnil,omitempty" name:"InstNodeId"`

	// Node role
	InstNodeRole *string `json:"InstNodeRole,omitnil,omitempty" name:"InstNodeRole"`
}

type ProxyNode struct {
	// Proxy node ID
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Number of CPU cores.
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size, measured in MB.
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Node status: 0 - Initializing, 1 - Online, 2 - Offline, 3 - Being destroyed, 4 - Recovering, 5 - Node fault, 6 - Switching.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Proxy node availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Proxy Node Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Number of connections
	Connection *uint64 `json:"Connection,omitnil,omitempty" name:"Connection"`
}

type ProxyNodeCustom struct {
	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ReadWriteNode struct {
	// Availability zone where the RW node is located.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// When upgrading a cloud disk edition instance, if you need to adjust the Availability Zone of Read-Only Nodes, you must specify the node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type ReadonlyNode struct {
	// Whether distributed in a random availability Zone. Import YES means random availability Zone. Otherwise used specified availability Zone.
	IsRandomZone *string `json:"IsRandomZone,omitnil,omitempty" name:"IsRandomZone"`

	// Specify the availability zone for node distribution.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// When upgrading a cloud disk edition instance, if you need to adjust the Availability Zone of Read-Only Nodes, you must specify the node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

// Predefined struct for user
type ReleaseIsolatedDBInstancesRequestParams struct {
	// Instance ID. The instance ID is in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the API for querying the instance list (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1). The value is the InstanceId field in the output parameter.
	// Description: Multiple instance IDs can be entered for operations. The json format is as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. The instance ID is in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the API for querying the instance list (https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1). The value is the InstanceId field in the output parameter.
	// Description: Multiple instance IDs can be entered for operations. The json format is as follows.
	// [
	//     "cdb-30z11v8s",
	//     "cdb-93h11efg"
	//   ]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ReleaseIsolatedDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIsolatedDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseIsolatedDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIsolatedDBInstancesResponseParams struct {
	// Deisolation result set.
	Items []*ReleaseResult `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseIsolatedDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseIsolatedDBInstancesResponseParams `json:"Response"`
}

func (r *ReleaseIsolatedDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIsolatedDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseResult struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Result value of instance deisolation. A returned value of 0 indicates success.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message for instance deisolation.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type ReloadBalanceProxyNodeRequestParams struct {
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. You can obtain it through the API [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1).
	// Note:
	// 1. For dual-node instances, this parameter is optional. If not provided, load balancing will be performed for ALL proxy group addresses.
	// 2. For cloud disk edition instances, this parameter is required.
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Proxy group address ID. You can obtain it through the API [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1).
	// Note:
	// 1. For dual-node instances, this parameter is optional. If not provided, load balancing will be performed for ALL proxy group addresses.
	// 2. For cloud disk edition instances, this parameter is required.
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

func (r *ReloadBalanceProxyNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProxyGroupId")
	delete(f, "ProxyAddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReloadBalanceProxyNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReloadBalanceProxyNodeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReloadBalanceProxyNodeResponse struct {
	*tchttp.BaseResponse
	Response *ReloadBalanceProxyNodeResponseParams `json:"Response"`
}

func (r *ReloadBalanceProxyNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReloadBalanceProxyNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteBackupInfo struct {
	// ID of the remote backup subtask
	SubBackupId *int64 `json:"SubBackupId,omitnil,omitempty" name:"SubBackupId"`

	// The region where the remote backup resides
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Backup task status. Valid values: `SUCCESS` (backup succeeded), `FAILED` (backup failed), `RUNNING` (backup is in progress).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The start time of remote backup
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// The end time of remote backup
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// The download address
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type RenewDBInstanceRequestParams struct {
	// ID of the instance to be renewed in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed in the TencentDB console. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Renewal period in months. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// To renew a pay-as-you-go instance to a yearly/monthly subscribed one, you need to set this parameter to `PREPAID`.
	ModifyPayType *string `json:"ModifyPayType,omitnil,omitempty" name:"ModifyPayType"`

	// Auto-renewal flag. 0 means no auto-renewal, 1 means auto-renewal.
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type RenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be renewed in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed in the TencentDB console. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Renewal period in months. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// To renew a pay-as-you-go instance to a yearly/monthly subscribed one, you need to set this parameter to `PREPAID`.
	ModifyPayType *string `json:"ModifyPayType,omitnil,omitempty" name:"ModifyPayType"`

	// Auto-renewal flag. 0 means no auto-renewal, 1 means auto-renewal.
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

func (r *RenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TimeSpan")
	delete(f, "ModifyPayType")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstanceResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBInstanceResponseParams `json:"Response"`
}

func (r *RenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance account name with manually refreshed rotation password, such as root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Manually refresh the domain name of the instance account with a rotated password, such as %
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance account name with manually refreshed rotation password, such as root
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Manually refresh the domain name of the instance account with a rotated password, such as %
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRootAccountRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ResetRootAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ResetRootAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRootAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRootAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRootAccountResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetRootAccountResponse struct {
	*tchttp.BaseResponse
	Response *ResetRootAccountResponseParams `json:"Response"`
}

func (r *ResetRootAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRootAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstancesRequestParams struct {
	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type RestartDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *RestartDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstancesResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstancesResponseParams `json:"Response"`
}

func (r *RestartDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoGroup struct {
	// <p>Read-only group mode. Available values are: alone-automatic allocation by the system; allinone-create a read-only group; join-use an existing read-only group.</p>
	RoGroupMode *string `json:"RoGroupMode,omitnil,omitempty" name:"RoGroupMode"`

	// <p>Read-only group ID.<br>Note: If the data structure is used during instance purchase, this item is required only when the read-only group mode is set to join.</p>
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// <p>Read-only group name.</p>
	RoGroupName *string `json:"RoGroupName,omitnil,omitempty" name:"RoGroupName"`

	// <p>Whether to enable the feature to isolate an instance that exceeds the latency threshold. After enabling this feature, if the delay between a read-only instance and the primary instance exceeds the delay threshold, the read-only instance will be isolated. Available values: 1-enable; 0-disable.</p>
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitnil,omitempty" name:"RoOfflineDelay"`

	// <p>Delay threshold, in seconds. Value range: 1–10000. The value is an integer.</p>
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitnil,omitempty" name:"RoMaxDelayTime"`

	// <p>Minimum number of instances to retain. If the number of read-only instances purchased is less than the set number, removal will not occur.</p>
	MinRoInGroup *int64 `json:"MinRoInGroup,omitnil,omitempty" name:"MinRoInGroup"`

	// <p>Read-write weight allocation mode. Available values: system-automatic allocation by the system; custom-customization.</p>
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// <p>This field is deprecated and meaningless. To view the weight of a read-only instance, check the Weight value in the RoInstances field.</p>
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>Details of read-only instances in the read-only group.</p>
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitnil,omitempty" name:"RoInstances"`

	// <p>Private IP address of the read-only group.</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Private network port number of the read-only group.</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>VPC ID.</p>
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// <p>Subnet ID.</p>
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// <p>Region of the read-only group.</p>
	RoGroupRegion *string `json:"RoGroupRegion,omitnil,omitempty" name:"RoGroupRegion"`

	// <p>AZ of the read-only group.</p>
	RoGroupZone *string `json:"RoGroupZone,omitnil,omitempty" name:"RoGroupZone"`

	// <p>Replication delay time, in seconds. Value range: 1–259200. The value is an integer.</p>
	DelayReplicationTime *int64 `json:"DelayReplicationTime,omitnil,omitempty" name:"DelayReplicationTime"`
}

type RoGroupAttr struct {
	// RO group name.
	RoGroupName *string `json:"RoGroupName,omitnil,omitempty" name:"RoGroupName"`

	// Maximum delay threshold for the RO instance. Unit: seconds, minimum value is 1. Range: [1,10000], integer.
	// Note: The RO group must have enabled the instance latency removal policy for this value to be valid.
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitnil,omitempty" name:"RoMaxDelayTime"`

	// Whether to enable instance removal. Valid values: 1 (enabled), 0 (not enabled). Please note that if instance removal is enabled, the delay threshold parameter (`RoMaxDelayTime`) must be set.
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitnil,omitempty" name:"RoOfflineDelay"`

	// Minimum reserved instances. Can be set to any value ≤ the number of instances in the RO group. Default value: 1.
	// Note: If the set value is larger than the RO instance count, do not remove. If set to 0, all instances with delay above the limit will be excluded.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitnil,omitempty" name:"MinRoInGroup"`

	// Weighting mode. Supported values include `system` (automatically assigned by the system) and `custom` (defined by user). Please note that if the `custom` mode is selected, the RO instance weight configuration parameter (RoWeightValues) must be set.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Delayed replication time. Unit: second, range: 1 - 259200 seconds, not required to enable delayed replication for the instance.
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
}

type RoInstanceInfo struct {
	// Master instance ID corresponding to the RO group
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// RO instance status in the RO group. Value range: online, offline
	RoStatus *string `json:"RoStatus,omitnil,omitempty" name:"RoStatus"`

	// Last deactivation time of a RO instance in the RO group
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// RO instance weight in the RO group
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// RO instance region name, such as ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Canonical name of the RO Availability Zone, for example ap-shanghai-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// RO instance ID in the format of cdbro-c1nl9rpv
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// RO instance status. Valid values: `0` (creating), `1` (running), `3` (remote RO), `4` (deleting). When the `DescribeDBInstances` API is used to query the information of the source instance, if the source instance is associated with a remote read-only instance, the returned status value of the remote read-only instance always shows 3.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only)
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// RO instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Pay-as-you-go status. Valid values: 1 - normal; 2 - in arrears.
	HourFeeStatus *int64 `json:"HourFeeStatus,omitnil,omitempty" name:"HourFeeStatus"`

	// RO instance task status. Value range: <br>0 - no task <br>1 - upgrading <br>2 - importing data <br>3 - activating secondary <br>4 - public network access enabled <br>5 - batch operation in progress <br>6 - rolling back <br>7 - public network access not enabled <br>8 - modifying password <br>9 - renaming instance <br>10 - restarting <br>12 - migrating self-built instance <br>13 - dropping table <br>14 - creating and syncing disaster recovery instance
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// RO instance memory size in MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// RO instance disk size in GB
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Queries per second
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// Private IP address of the RO instance
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Access port of the RO instance
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// VPC ID of the RO instance
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID of the RO instance
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// RO instance specification description. Value range: CUSTOM
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Database engine version of the read-only replica. Valid values: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// RO instance expiration time in the format of yyyy-mm-dd hh:mm:ss. If it is a pay-as-you-go instance, the value of this field is 0000-00-00 00:00:00
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// Billing type of the RO instance. Valid values: 0 - yearly/monthly subscription; 1 - pay-as-you-go; 2-postpaid by month.
	PayType *int64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// RO replication delay status.
	ReplicationStatus *string `json:"ReplicationStatus,omitnil,omitempty" name:"ReplicationStatus"`
}

type RoVipInfo struct {
	// VIP status of the read-only instance
	RoVipStatus *int64 `json:"RoVipStatus,omitnil,omitempty" name:"RoVipStatus"`

	// VPC subnet of the read-only instance
	RoSubnetId *int64 `json:"RoSubnetId,omitnil,omitempty" name:"RoSubnetId"`

	// VPC of the read-only instance
	RoVpcId *int64 `json:"RoVpcId,omitnil,omitempty" name:"RoVpcId"`

	// VIP port number of the read-only instance
	RoVport *int64 `json:"RoVport,omitnil,omitempty" name:"RoVport"`

	// VIP of the read-only instance
	RoVip *string `json:"RoVip,omitnil,omitempty" name:"RoVip"`
}

type RoWeightValue struct {
	// RO instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Weight value. Value range: [0, 100].
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type RollbackDBName struct {
	// Original database name before rollback
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Name of the rolled back database
	NewDatabaseName *string `json:"NewDatabaseName,omitnil,omitempty" name:"NewDatabaseName"`
}

type RollbackInstancesInfo struct {
	// Cloud database instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rollback strategy. Optional values: table, db, full. table - Ultra-fast rollback mode, only imports selected table-level backups and binlog. If there are cross-table operations and the associated table hasn't been selected, it will cause rollback failure. In this mode, parameter Databases must be empty. db - Quick mode, only imports selected database-level backups and binlog. If there are cross-database operations and the associated database hasn't been selected, it will cause rollback failure. full - Standard rollback mode, imports backups and binlog of the entire instance, speed is not as fast.
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Database rollback time in the format of yyyy-mm-dd hh:mm:ss.
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// Database information to be rolled back, which means database rollback.
	Databases []*RollbackDBName `json:"Databases,omitnil,omitempty" name:"Databases"`

	// Database table information to be rolled back, which means rollback by table.
	Tables []*RollbackTables `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type RollbackTableName struct {
	// Original database table name before rollback
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Rolled back database table name
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`
}

type RollbackTables struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Database table details
	Table []*RollbackTableName `json:"Table,omitnil,omitempty" name:"Table"`
}

type RollbackTask struct {
	// Task execution information.
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// Task execution result. Valid values: INITIAL: initializing, RUNNING: running, SUCCESS: succeeded, FAILED: failed, KILLED: terminated, REMOVED: deleted, PAUSED: paused.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task execution progress. Value range: [0,100].
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Rollback task detail.
	Detail []*RollbackInstancesInfo `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type RollbackTimeRange struct {
	// Start time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-10-29 01:06:04
	Begin *string `json:"Begin,omitnil,omitempty" name:"Begin"`

	// End time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-11-02 11:44:47
	End *string `json:"End,omitnil,omitempty" name:"End"`
}

type Rule struct {
	// Division ceiling
	LessThan *uint64 `json:"LessThan,omitnil,omitempty" name:"LessThan"`

	// Weight
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type RuleFilters struct {
	// Parameter name of the audit rule filter.  Valid values:  `host` (client IP), `user` (database account), `dbName` (database name), `sqlType` (SQL type), `sql` (SQL statement), `affectRows` (affected rows), `sentRows` (returned rows), `checkRows` (scanned rows), `execTime` (execution rows).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter match value of the audit rule Valid values:  `INC` (including), `EXC` (excluding), `EQS` (equal to), `NEQ` (not equal to), `REG` (regex), `GT` (greater than), `LT` (less than).
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// Filter match value of the audit rule Valid values for `sqlType`: `alter`, `changeuser`, `create`, `delete`, `drop`, `execute`, `insert`, `login`, `logout`, `other`, `replace`, `select`, `set, `update`.
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RuleTemplateInfo struct {
	// Rule template ID.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Rule template name.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Rule content.
	RuleFilters []*RuleFilters `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`

	// Alarm level. Valid values: 1 - Low risk, 2 - Medium risk, 3 - High risk.
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Alarm policy. Valid values: 0 - Alarm disabled, 1 - Alarm enabled.
	AlarmPolicy *int64 `json:"AlarmPolicy,omitnil,omitempty" name:"AlarmPolicy"`

	// Rule description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type RuleTemplateRecordInfo struct {
	// Task ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Details of the original rule template.
	ModifyBeforeInfo *RuleTemplateInfo `json:"ModifyBeforeInfo,omitnil,omitempty" name:"ModifyBeforeInfo"`

	// Details of the modified rule template.
	ModifyAfterInfo *RuleTemplateInfo `json:"ModifyAfterInfo,omitnil,omitempty" name:"ModifyAfterInfo"`

	// Affected instances.
	AffectedInstances []*string `json:"AffectedInstances,omitnil,omitempty" name:"AffectedInstances"`

	// Operator (account UIN).
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Time of the change.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SecurityGroup struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*Inbound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*Outbound `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`
}

type SlaveConfig struct {
	// Replication mode of the secondary database. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitnil,omitempty" name:"ReplicationMode"`

	// Canonical name of the read-only availability zone, for example ap-shanghai-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type SlaveInfo struct {
	// <p>Secondary server information of the top spot</p>
	First *SlaveInstanceInfo `json:"First,omitnil,omitempty" name:"First"`

	// <p>Second standby machine information</p>
	Second *SlaveInstanceInfo `json:"Second,omitnil,omitempty" name:"Second"`
}

type SlaveInstanceInfo struct {
	// Port number
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Region information
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Virtual IP information
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// AZ information
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type SlowLogInfo struct {
	// Backup filename
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Backup snapshot time. Time format: 2016-03-17.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Download address on the private network
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// Download address on the public network
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// Log type. Value range: slowlog (slow log)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SlowLogItem struct {
	// Sql execution time. Unix second-level timestamp.
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Execution duration of Sql (seconds).
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// Sql statement.
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Client IP address.
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Lock duration (unit: second).
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// Number of scanned rows.
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// Result set row count.
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// Sql Template.
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// md5 of the Sql statement.
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type SqlFileInfo struct {
	// Upload time
	UploadTime *string `json:"UploadTime,omitnil,omitempty" name:"UploadTime"`

	// Upload progress
	UploadInfo *UploadInfo `json:"UploadInfo,omitnil,omitempty" name:"UploadInfo"`

	// Filename
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size in bytes
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Whether upload is finished. Valid values: 0 (not completed), 1 (completed)
	IsUploadFinished *int64 `json:"IsUploadFinished,omitnil,omitempty" name:"IsUploadFinished"`

	// File ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`
}

// Predefined struct for user
type StartBatchRollbackRequestParams struct {
	// Details of the instance for rollback
	Instances []*RollbackInstancesInfo `json:"Instances,omitnil,omitempty" name:"Instances"`
}

type StartBatchRollbackRequest struct {
	*tchttp.BaseRequest
	
	// Details of the instance for rollback
	Instances []*RollbackInstancesInfo `json:"Instances,omitnil,omitempty" name:"Instances"`
}

func (r *StartBatchRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBatchRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBatchRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBatchRollbackResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartBatchRollbackResponse struct {
	*tchttp.BaseResponse
	Response *StartBatchRollbackResponseParams `json:"Response"`
}

func (r *StartBatchRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBatchRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCpuExpandRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Scale-out type supports auto-scaling and custom scaling.
	// Description: 1. auto means automatic scaling. 2. manual means custom scaling with immediate effect. 3. timeInterval means custom scaling by time. 4. period means custom scaling by cycle.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of CPU cores for scale-out during customization.
	// Description: 1. This parameter is required when Type is set to manual, timeInterval, or period. 2. The maximum number of CPU cores to increase is the current instance's CPU core number. For example, an 8-core 16GB instance can scale out up to 8 CPU cores, with a range of 1 - 8.
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Automatic scale-out policy. This parameter is required when Type is set to auto.
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// Scaling policy by time period.
	// Description: When Type is timeInterval, TimeIntervalStrategy is required.
	TimeIntervalStrategy *TimeIntervalStrategy `json:"TimeIntervalStrategy,omitnil,omitempty" name:"TimeIntervalStrategy"`

	// Scale by cycle.
	// Description: When Type is period, PeriodStrategy is required.
	PeriodStrategy *PeriodStrategy `json:"PeriodStrategy,omitnil,omitempty" name:"PeriodStrategy"`
}

type StartCpuExpandRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Scale-out type supports auto-scaling and custom scaling.
	// Description: 1. auto means automatic scaling. 2. manual means custom scaling with immediate effect. 3. timeInterval means custom scaling by time. 4. period means custom scaling by cycle.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of CPU cores for scale-out during customization.
	// Description: 1. This parameter is required when Type is set to manual, timeInterval, or period. 2. The maximum number of CPU cores to increase is the current instance's CPU core number. For example, an 8-core 16GB instance can scale out up to 8 CPU cores, with a range of 1 - 8.
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Automatic scale-out policy. This parameter is required when Type is set to auto.
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// Scaling policy by time period.
	// Description: When Type is timeInterval, TimeIntervalStrategy is required.
	TimeIntervalStrategy *TimeIntervalStrategy `json:"TimeIntervalStrategy,omitnil,omitempty" name:"TimeIntervalStrategy"`

	// Scale by cycle.
	// Description: When Type is period, PeriodStrategy is required.
	PeriodStrategy *PeriodStrategy `json:"PeriodStrategy,omitnil,omitempty" name:"PeriodStrategy"`
}

func (r *StartCpuExpandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCpuExpandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "ExpandCpu")
	delete(f, "AutoStrategy")
	delete(f, "TimeIntervalStrategy")
	delete(f, "PeriodStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCpuExpandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCpuExpandResponseParams struct {
	// Asynchronous Task ID. Call the API DescribeAsyncRequest and input the ID to query the task execution progress.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartCpuExpandResponse struct {
	*tchttp.BaseResponse
	Response *StartCpuExpandResponseParams `json:"Response"`
}

func (r *StartCpuExpandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCpuExpandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartReplicationRequestParams struct {
	// Instance ID. It only supports read-only instances. It can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StartReplicationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. It only supports read-only instances. It can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StartReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartReplicationResponseParams struct {
	// Asynchronous Task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartReplicationResponse struct {
	*tchttp.BaseResponse
	Response *StartReplicationResponseParams `json:"Response"`
}

func (r *StartReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCpuExpandRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopCpuExpandRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopCpuExpandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCpuExpandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCpuExpandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCpuExpandResponseParams struct {
	// Asynchronous Task ID. When calling [DescribeAsyncRequestInfo](https://www.tencentcloud.com/document/api/236/20410?from_cn_redirect=1) to query the task execution progress, you can pass in this ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCpuExpandResponse struct {
	*tchttp.BaseResponse
	Response *StopCpuExpandResponseParams `json:"Response"`
}

func (r *StopCpuExpandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCpuExpandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDBImportJobRequestParams struct {
	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type StopDBImportJobRequest struct {
	*tchttp.BaseRequest
	
	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

func (r *StopDBImportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDBImportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDBImportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDBImportJobResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopDBImportJobResponse struct {
	*tchttp.BaseResponse
	Response *StopDBImportJobResponseParams `json:"Response"`
}

func (r *StopDBImportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDBImportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopReplicationRequestParams struct {
	// Instance ID. It only supports read-only instances. It can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopReplicationRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. It only supports read-only instances. It can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopReplicationResponseParams struct {
	// Asynchronous Task ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopReplicationResponse struct {
	*tchttp.BaseResponse
	Response *StopReplicationResponseParams `json:"Response"`
}

func (r *StopReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRollbackRequestParams struct {
	// Revoke the corresponding instance ID of the rollback task. Obtain through the API [DescribeDBInstances](https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopRollbackRequest struct {
	*tchttp.BaseRequest
	
	// Revoke the corresponding instance ID of the rollback task. Obtain through the API [DescribeDBInstances](https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StopRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRollbackResponseParams struct {
	// Asynchronous Task ID of the execution request.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRollbackResponse struct {
	*tchttp.BaseResponse
	Response *StopRollbackResponseParams `json:"Response"`
}

func (r *StopRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitInstanceUpgradeCheckJobRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target database version.
	// Description: Available values: 5.6, 5.7, 8.0. Cross-version upgrade is not supported. Version downgrade is unsupported after upgrade.
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

type SubmitInstanceUpgradeCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target database version.
	// Description: Available values: 5.6, 5.7, 8.0. Cross-version upgrade is not supported. Version downgrade is unsupported after upgrade.
	DstMysqlVersion *string `json:"DstMysqlVersion,omitnil,omitempty" name:"DstMysqlVersion"`
}

func (r *SubmitInstanceUpgradeCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitInstanceUpgradeCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstMysqlVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitInstanceUpgradeCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitInstanceUpgradeCheckJobResponseParams struct {
	// Task ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitInstanceUpgradeCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitInstanceUpgradeCheckJobResponseParams `json:"Response"`
}

func (r *SubmitInstanceUpgradeCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitInstanceUpgradeCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchCDBProxyRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type SwitchCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

func (r *SwitchCDBProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchCDBProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchCDBProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchCDBProxyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchCDBProxyResponse struct {
	*tchttp.BaseResponse
	Response *SwitchCDBProxyResponseParams `json:"Response"`
}

func (r *SwitchCDBProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchCDBProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceMasterSlaveRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the replica server to switched to. Valid values: `first` (the first replica server), `second` (the second replica server). Default value: `first`. `second` is valid only for a multi-AZ instance.
	DstSlave *string `json:"DstSlave,omitnil,omitempty" name:"DstSlave"`

	// Whether to force the switch. Valid values: `True`, `False` (default). If this parameter is set to `True`, instance data may be lost during the switch.
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// Whether to perform the switch during a time window. Valid values: `True`, `False` (default). If `ForceSwitch` is set to `True`, this parameter is invalid.
	WaitSwitch *bool `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// Trigger primary-secondary switch for the designated node ID of the cluster edition instance.
	DstNodeId *string `json:"DstNodeId,omitnil,omitempty" name:"DstNodeId"`
}

type SwitchDBInstanceMasterSlaveRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the replica server to switched to. Valid values: `first` (the first replica server), `second` (the second replica server). Default value: `first`. `second` is valid only for a multi-AZ instance.
	DstSlave *string `json:"DstSlave,omitnil,omitempty" name:"DstSlave"`

	// Whether to force the switch. Valid values: `True`, `False` (default). If this parameter is set to `True`, instance data may be lost during the switch.
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// Whether to perform the switch during a time window. Valid values: `True`, `False` (default). If `ForceSwitch` is set to `True`, this parameter is invalid.
	WaitSwitch *bool `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// Trigger primary-secondary switch for the designated node ID of the cluster edition instance.
	DstNodeId *string `json:"DstNodeId,omitnil,omitempty" name:"DstNodeId"`
}

func (r *SwitchDBInstanceMasterSlaveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceMasterSlaveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DstSlave")
	delete(f, "ForceSwitch")
	delete(f, "WaitSwitch")
	delete(f, "DstNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstanceMasterSlaveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceMasterSlaveResponseParams struct {
	// Async task ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDBInstanceMasterSlaveResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDBInstanceMasterSlaveResponseParams `json:"Response"`
}

func (r *SwitchDBInstanceMasterSlaveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceMasterSlaveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDrInstanceToMasterRequestParams struct {
	// Disaster recovery instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type SwitchDrInstanceToMasterRequest struct {
	*tchttp.BaseRequest
	
	// Disaster recovery instance ID, in the format such as cdb-c1nl9rpv. This matches the instance ID displayed on the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *SwitchDrInstanceToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrInstanceToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDrInstanceToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDrInstanceToMasterResponseParams struct {
	// Async task request ID, which can be used to query the execution result of an async task
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDrInstanceToMasterResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDrInstanceToMasterResponseParams `json:"Response"`
}

func (r *SwitchDrInstanceToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDrInstanceToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchForUpgradeRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to enable association switchover. Enable for true and shutdown for false. Default false.
	IsRelatedSwitch *bool `json:"IsRelatedSwitch,omitnil,omitempty" name:"IsRelatedSwitch"`
}

type SwitchForUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to enable association switchover. Enable for true and shutdown for false. Default false.
	IsRelatedSwitch *bool `json:"IsRelatedSwitch,omitnil,omitempty" name:"IsRelatedSwitch"`
}

func (r *SwitchForUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchForUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IsRelatedSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchForUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchForUpgradeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchForUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *SwitchForUpgradeResponseParams `json:"Response"`
}

func (r *SwitchForUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchForUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TImeCycle struct {
	// Whether to choose Monday for scaling during the period.
	// Description: Value "true" means select, value "false" means not select.
	Monday *bool `json:"Monday,omitnil,omitempty" name:"Monday"`

	// During scaling, whether to choose Tuesday for expansion.
	// Description: Value "true" means select, value "false" means not select.
	Tuesday *bool `json:"Tuesday,omitnil,omitempty" name:"Tuesday"`

	// Whether to choose Wednesday for scaling during the period.
	// Description: Value "true" means select, value "false" means not select.
	Wednesday *bool `json:"Wednesday,omitnil,omitempty" name:"Wednesday"`

	// Whether to choose Thursday for scaling during the period.
	// Description: Value "true" means select, value "false" means not select.
	Thursday *bool `json:"Thursday,omitnil,omitempty" name:"Thursday"`

	// Whether to choose Friday for scaling during the period.
	// Description: Value "true" means select, value "false" means not select.
	Friday *bool `json:"Friday,omitnil,omitempty" name:"Friday"`

	// Whether to choose Saturday for scaling during the period.
	// Description: Value "true" means select, value "false" means not select.
	Saturday *bool `json:"Saturday,omitnil,omitempty" name:"Saturday"`

	// Whether to choose Sunday for scaling during the period.
	// Description: Value "true" means select, value "false" means not select.
	Sunday *bool `json:"Sunday,omitnil,omitempty" name:"Sunday"`
}

type TablePrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagInfoItem struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagInfoUnit struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagsInfoOfInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tag information
	Tags []*TagInfoUnit `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TaskAttachInfo struct {
	// Upgrade task
	// "FastUpgradeStatus": Indicates the upgrade type. 1 - in-place upgrade; 0 - normal upgrade.
	AttachKey *string `json:"AttachKey,omitnil,omitempty" name:"AttachKey"`

	// Upgrade task
	// "FastUpgradeStatus": Indicates the upgrade type. 1 - In-place upgrade; 0 - Normal upgrade.
	AttachValue *string `json:"AttachValue,omitnil,omitempty" name:"AttachValue"`
}

type TaskDetail struct {
	// Error code. `0` indicates success. Other values correspond to different error scenarios.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Error message.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// ID of an instance task.
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Instance task progress.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Instance task status. Valid values:
	// "UNDEFINED" - undefined;
	// "INITIAL" - initializing;
	// "RUNNING" - running;
	// "SUCCEED" - succeeded;
	// "FAILED" - failed;
	// "KILLED" - terminated;
	// "REMOVED" - deleted;
	// "PAUSED" - paused.
	// "WAITING" - waiting (which can be canceled)
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Instance task type. Valid values:
	// "ROLLBACK" - rolling back a database;
	// "SQL OPERATION" - performing an SQL operation;
	// "IMPORT DATA" - importing data;
	// "MODIFY PARAM" - setting a parameter;
	// "INITIAL" - initializing a TencentDB instance;
	// "REBOOT" - restarting a TencentDB instance;
	// "OPEN GTID" - enabling GTID of a TencentDB instance;
	// "UPGRADE RO" - upgrading a read-only instance;
	// "BATCH ROLLBACK" - rolling back databases in batches;
	// "UPGRADE MASTER" - upgrading a primary instance;
	// "DROP TABLES" - dropping a TencentDB table;
	// "SWITCH DR TO MASTER" - promoting a disaster recovery instance.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Instance task start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Instance task end time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ID of the associated instance.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// Additional information of the task.
	TaskAttachInfo []*TaskAttachInfo `json:"TaskAttachInfo,omitnil,omitempty" name:"TaskAttachInfo"`
}

type TimeInterval struct {
	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type TimeIntervalStrategy struct {
	// Expansion time started.
	// Description: This value is a timestamp in seconds in Integer format.
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Expansion time ended.
	// Description: This value is a timestamp in seconds in Integer format.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type UpgradeAnalysisInstanceVersionInfo struct {
	// <p>Grayscale ip for version upgrade</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>Grayscale port for version upgrade</p>
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// <p>Upgrade to a later version</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Grayscale event for instance upgrade</p><p>Unit: day</p>
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type UpgradeCDBProxyVersionRequestParams struct {
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Current version of database proxy
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// Target version of database proxy
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// Upgrade time. Valid values: `nowTime` (upgrade immediately), `timeWindow` (upgrade during instance maintenance time)
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

type UpgradeCDBProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID, which can be obtained through the [DescribeCdbProxyInfo](https://www.tencentcloud.com/document/api/236/90585?from_cn_redirect=1) API.
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Current version of database proxy
	SrcProxyVersion *string `json:"SrcProxyVersion,omitnil,omitempty" name:"SrcProxyVersion"`

	// Target version of database proxy
	DstProxyVersion *string `json:"DstProxyVersion,omitnil,omitempty" name:"DstProxyVersion"`

	// Upgrade time. Valid values: `nowTime` (upgrade immediately), `timeWindow` (upgrade during instance maintenance time)
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

func (r *UpgradeCDBProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeCDBProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyGroupId")
	delete(f, "SrcProxyVersion")
	delete(f, "DstProxyVersion")
	delete(f, "UpgradeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeCDBProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeCDBProxyVersionResponseParams struct {
	// Async Processing ID
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeCDBProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeCDBProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeCDBProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeCDBProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceEngineVersionRequestParams struct {
	// <p>Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API, with its value being the InstanceId field in the output parameter.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Database engine version of the primary instance. Supported values include 5.6, 5.7, 8.0.<br>Description: Cross-version upgrade is not supported. Downgrade is not supported after upgrade.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>The way to switch to a new instance defaults to 0. Supported values include: 0 - switch immediately, 1 - switch in a time window. When the value is 1, during the upgrade process, the switchover to a new instance will be performed in the time window, or the user can proactively call the API <a href="https://www.tencentcloud.com/document/product/236/15864?from_cn_redirect=1">switch to a new instance</a> to trigger the process.</p>
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// <p>Whether to upgrade the kernel subversion. Supported values: 1 - upgrade kernel subversion; 0 - upgrade database engine version. No default value. Specify the version type to upgrade.</p>
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitnil,omitempty" name:"UpgradeSubversion"`

	// <p>Delay threshold. Value ranges from 1 to 10. No default value. When not specified, the delay threshold is 0, which means the delay threshold is not set.</p>
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// <p>Whether to ignore keyword errors when upgrading from 5.7 to 8.0. The value ranges from 0 to 1. 1 means ignored, 0 means not ignored. No default value. Not specified means no action taken.</p>
	IgnoreErrKeyword *int64 `json:"IgnoreErrKeyword,omitnil,omitempty" name:"IgnoreErrKeyword"`

	// <p>Upgrade support for specified parameters</p>
	ParamList []*UpgradeEngineVersionParams `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API, with its value being the InstanceId field in the output parameter.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Database engine version of the primary instance. Supported values include 5.6, 5.7, 8.0.<br>Description: Cross-version upgrade is not supported. Downgrade is not supported after upgrade.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>The way to switch to a new instance defaults to 0. Supported values include: 0 - switch immediately, 1 - switch in a time window. When the value is 1, during the upgrade process, the switchover to a new instance will be performed in the time window, or the user can proactively call the API <a href="https://www.tencentcloud.com/document/product/236/15864?from_cn_redirect=1">switch to a new instance</a> to trigger the process.</p>
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// <p>Whether to upgrade the kernel subversion. Supported values: 1 - upgrade kernel subversion; 0 - upgrade database engine version. No default value. Specify the version type to upgrade.</p>
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitnil,omitempty" name:"UpgradeSubversion"`

	// <p>Delay threshold. Value ranges from 1 to 10. No default value. When not specified, the delay threshold is 0, which means the delay threshold is not set.</p>
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// <p>Whether to ignore keyword errors when upgrading from 5.7 to 8.0. The value ranges from 0 to 1. 1 means ignored, 0 means not ignored. No default value. Not specified means no action taken.</p>
	IgnoreErrKeyword *int64 `json:"IgnoreErrKeyword,omitnil,omitempty" name:"IgnoreErrKeyword"`

	// <p>Upgrade support for specified parameters</p>
	ParamList []*UpgradeEngineVersionParams `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *UpgradeDBInstanceEngineVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceEngineVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EngineVersion")
	delete(f, "WaitSwitch")
	delete(f, "UpgradeSubversion")
	delete(f, "MaxDelayTime")
	delete(f, "IgnoreErrKeyword")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceEngineVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceEngineVersionResponseParams struct {
	// <p>Asynchronous Task ID. Use <a href="https://www.tencentcloud.com/document/api/236/20410?from_cn_redirect=1">Query Asynchronous Task</a> to get its execution situation.</p>
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceEngineVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceEngineVersionResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceEngineVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceEngineVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
	// <p>Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API, with its value being the InstanceId field in the output parameter.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Memory size after upgrade, unit: MB. To ensure the validity of the imported Memory value, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the upgradeable memory specifications.<br>Note: If you perform business migration, fill in the instance specification (CPU, memory), otherwise the system will use the minimum allowed specification by default.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Disk size after upgrade, unit: GB. To ensure the validity of the imported Volume value, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the upgradeable disk range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Data replication method. Supported values include: 0 - async replication, 1 - semi-sync replication, 2 - strong sync replication. Specify this parameter when upgrading the primary instance. This parameter is invalid when upgrading a read-only instance or disaster recovery instance.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Deployment mode, defaults to 0. Supported values include: 0 - single-AZ deployment, 1 - multi-AZ deployment. You can specify this parameter when upgrading the primary instance. This parameter is invalid when upgrading a read-only instance or disaster recovery instance.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>The availability zone information of standby database 1 matches the Zone parameter of the instance by default. You can specify this parameter when upgrading the primary instance to multi-AZ deployment. This parameter is invalid when upgrading a read-only instance or disaster recovery instance. You can query the supported availability zones via the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>Database engine version of the primary instance. Supported values include 5.5, 5.6, 5.7, and 8.0.<br>Note: Please use the <a href="https://www.tencentcloud.com/document/api/236/15870?from_cn_redirect=1">UpgradeDBInstanceEngineVersion</a> API to upgrade the database version.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>The way to switch to a new instance defaults to 0. Supported values include: 0 - switch immediately, 1 - switch within a time window. When the value is 1, during the upgrade, the process to switch to a new instance will be performed within the time window, or the user can actively call the API <a href="https://www.tencentcloud.com/document/product/236/15864?from_cn_redirect=1">Switch to a New Instance</a> to trigger the process.</p>
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// <p>The availability zone information of standby 2 is empty by default. You can specify this parameter when upgrading the primary instance, but it is invalid when upgrading a read-only instance or disaster recovery instance. Query the supported AZs via the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API.<br>Remark: To downgrade a three-node instance to a two-node one, set this parameter to empty to achieve it.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Instance type, defaults to master. Supported values include: master - refers to the primary instance, dr - refers to the disaster recovery instance, ro - refers to the read-only instance.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Instance isolation type. Supported values include "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "BASIC" - BASIC edition instance.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Number of cpu cores of the instance after upgrade. If not provided, the system will automatically fill in the minimum allowed specification based on the Memory size specified by Memory.<br>Description: If you need to migrate business, make sure to fill in the instance specification (cpu, Memory). Otherwise, the system will use the minimum allowed specification by default.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Whether to perform Rapid Configuration Change. 0-Normal upgrade, 1-Rapid Configuration Change, 2-Precedence to rapid change. Selecting Rapid Configuration Change will validate whether it is possible to perform ultra-fast reconfiguration based on resource status. If conditions are met, ultra-fast reconfiguration will be performed; otherwise, error information will be returned.</p>
	FastUpgrade *int64 `json:"FastUpgrade,omitnil,omitempty" name:"FastUpgrade"`

	// <p>Delay threshold. Value ranges from 1 to 10, default value is 10.</p>
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// <p>Whether to perform cross-region migration. 0 - ordinary migration, 1 - cross-region migration, default value is 0. When set to 1, it supports changes to the primary node availability zone of the instance.</p>
	CrossCluster *int64 `json:"CrossCluster,omitnil,omitempty" name:"CrossCluster"`

	// <p>Primary node availability zone. This parameter is valid only when cross-AZ migration. You can only migrate in the same region.</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>For cross-cluster migration scenarios, select the processing logic for intra-AZ read-only instances. together-intra-AZ read-only instances migrate with the primary instance to the target availability zone (default option), severally-intra-AZ read-only instances maintain the original deployment mode and do not move to the target availability zone.</p>
	RoTransType *string `json:"RoTransType,omitnil,omitempty" name:"RoTransType"`

	// <p>Topology configuration of cloud disk edition nodes.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Check whether in-place upgrade requires restart. 1-Check, 0-Do not check. If the value is 1 and the check shows that in-place upgrade must be restarted, the upgrade will be stopped and a notification will be returned. If in-place upgrade does not require restart, the upgrade process will proceed normally.</p>
	CheckFastUpgradeReboot *int64 `json:"CheckFastUpgradeReboot,omitnil,omitempty" name:"CheckFastUpgradeReboot"`

	// <p>Data validation sensitivity. This parameter is used for non-Rapid Configuration Change scenarios. Sensitivity is calculated based on current instance specifications to determine cpu resource usage for data comparison during the migration process. Corresponding options are: "high", "normal", "low", empty by default. Parameter explanation: "high": Corresponds to high in the console, not recommended when database load is too high. "normal": Corresponds to standard in the console. "low": Corresponds to low in the console.</p>
	DataCheckSensitive *string `json:"DataCheckSensitive,omitnil,omitempty" name:"DataCheckSensitive"`

	// <p>AZ information of standby database 3 is empty by default. You can specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Instance ID, in the format such as cdb-c1nl9rpv or cdbro-c1nl9rpv. This matches the instance ID displayed on the TencentDB console. You can obtain it through the <a href="https://www.tencentcloud.com/document/api/236/15872?from_cn_redirect=1">Query Instance List</a> API, with its value being the InstanceId field in the output parameter.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Memory size after upgrade, unit: MB. To ensure the validity of the imported Memory value, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the upgradeable memory specifications.<br>Note: If you perform business migration, fill in the instance specification (CPU, memory), otherwise the system will use the minimum allowed specification by default.</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>Disk size after upgrade, unit: GB. To ensure the validity of the imported Volume value, please use the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API to get the upgradeable disk range.</p>
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// <p>Data replication method. Supported values include: 0 - async replication, 1 - semi-sync replication, 2 - strong sync replication. Specify this parameter when upgrading the primary instance. This parameter is invalid when upgrading a read-only instance or disaster recovery instance.</p>
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// <p>Deployment mode, defaults to 0. Supported values include: 0 - single-AZ deployment, 1 - multi-AZ deployment. You can specify this parameter when upgrading the primary instance. This parameter is invalid when upgrading a read-only instance or disaster recovery instance.</p>
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// <p>The availability zone information of standby database 1 matches the Zone parameter of the instance by default. You can specify this parameter when upgrading the primary instance to multi-AZ deployment. This parameter is invalid when upgrading a read-only instance or disaster recovery instance. You can query the supported availability zones via the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API.</p>
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// <p>Database engine version of the primary instance. Supported values include 5.5, 5.6, 5.7, and 8.0.<br>Note: Please use the <a href="https://www.tencentcloud.com/document/api/236/15870?from_cn_redirect=1">UpgradeDBInstanceEngineVersion</a> API to upgrade the database version.</p>
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>The way to switch to a new instance defaults to 0. Supported values include: 0 - switch immediately, 1 - switch within a time window. When the value is 1, during the upgrade, the process to switch to a new instance will be performed within the time window, or the user can actively call the API <a href="https://www.tencentcloud.com/document/product/236/15864?from_cn_redirect=1">Switch to a New Instance</a> to trigger the process.</p>
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// <p>The availability zone information of standby 2 is empty by default. You can specify this parameter when upgrading the primary instance, but it is invalid when upgrading a read-only instance or disaster recovery instance. Query the supported AZs via the <a href="https://www.tencentcloud.com/document/product/236/17229?from_cn_redirect=1">obtain the purchasable specifications of cloud databases</a> API.<br>Remark: To downgrade a three-node instance to a two-node one, set this parameter to empty to achieve it.</p>
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// <p>Instance type, defaults to master. Supported values include: master - refers to the primary instance, dr - refers to the disaster recovery instance, ro - refers to the read-only instance.</p>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// <p>Instance isolation type. Supported values include "UNIVERSAL" - general-purpose instance, "EXCLUSIVE" - dedicated instance, "BASIC" - BASIC edition instance.</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>Number of cpu cores of the instance after upgrade. If not provided, the system will automatically fill in the minimum allowed specification based on the Memory size specified by Memory.<br>Description: If you need to migrate business, make sure to fill in the instance specification (cpu, Memory). Otherwise, the system will use the minimum allowed specification by default.</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>Whether to perform Rapid Configuration Change. 0-Normal upgrade, 1-Rapid Configuration Change, 2-Precedence to rapid change. Selecting Rapid Configuration Change will validate whether it is possible to perform ultra-fast reconfiguration based on resource status. If conditions are met, ultra-fast reconfiguration will be performed; otherwise, error information will be returned.</p>
	FastUpgrade *int64 `json:"FastUpgrade,omitnil,omitempty" name:"FastUpgrade"`

	// <p>Delay threshold. Value ranges from 1 to 10, default value is 10.</p>
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// <p>Whether to perform cross-region migration. 0 - ordinary migration, 1 - cross-region migration, default value is 0. When set to 1, it supports changes to the primary node availability zone of the instance.</p>
	CrossCluster *int64 `json:"CrossCluster,omitnil,omitempty" name:"CrossCluster"`

	// <p>Primary node availability zone. This parameter is valid only when cross-AZ migration. You can only migrate in the same region.</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>For cross-cluster migration scenarios, select the processing logic for intra-AZ read-only instances. together-intra-AZ read-only instances migrate with the primary instance to the target availability zone (default option), severally-intra-AZ read-only instances maintain the original deployment mode and do not move to the target availability zone.</p>
	RoTransType *string `json:"RoTransType,omitnil,omitempty" name:"RoTransType"`

	// <p>Topology configuration of cloud disk edition nodes.</p>
	ClusterTopology *ClusterTopology `json:"ClusterTopology,omitnil,omitempty" name:"ClusterTopology"`

	// <p>Check whether in-place upgrade requires restart. 1-Check, 0-Do not check. If the value is 1 and the check shows that in-place upgrade must be restarted, the upgrade will be stopped and a notification will be returned. If in-place upgrade does not require restart, the upgrade process will proceed normally.</p>
	CheckFastUpgradeReboot *int64 `json:"CheckFastUpgradeReboot,omitnil,omitempty" name:"CheckFastUpgradeReboot"`

	// <p>Data validation sensitivity. This parameter is used for non-Rapid Configuration Change scenarios. Sensitivity is calculated based on current instance specifications to determine cpu resource usage for data comparison during the migration process. Corresponding options are: "high", "normal", "low", empty by default. Parameter explanation: "high": Corresponds to high in the console, not recommended when database load is too high. "normal": Corresponds to standard in the console. "low": Corresponds to low in the console.</p>
	DataCheckSensitive *string `json:"DataCheckSensitive,omitnil,omitempty" name:"DataCheckSensitive"`

	// <p>AZ information of standby database 3 is empty by default. You can specify this parameter when you proceed to purchase a four-node primary instance.</p>
	FourthZone *string `json:"FourthZone,omitnil,omitempty" name:"FourthZone"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ProtectMode")
	delete(f, "DeployMode")
	delete(f, "SlaveZone")
	delete(f, "EngineVersion")
	delete(f, "WaitSwitch")
	delete(f, "BackupZone")
	delete(f, "InstanceRole")
	delete(f, "DeviceType")
	delete(f, "Cpu")
	delete(f, "FastUpgrade")
	delete(f, "MaxDelayTime")
	delete(f, "CrossCluster")
	delete(f, "ZoneId")
	delete(f, "RoTransType")
	delete(f, "ClusterTopology")
	delete(f, "CheckFastUpgradeReboot")
	delete(f, "DataCheckSensitive")
	delete(f, "FourthZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// <p>Order ID.</p>
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// <p>Request ID of the async task. Use this ID to <a href="https://www.tencentcloud.com/document/product/236/20410?from_cn_redirect=1">query the outcome of the async task</a>.</p>
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeEngineVersionParams struct {
	// Parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type UploadInfo struct {
	// Number of parts of file
	AllSliceNum *int64 `json:"AllSliceNum,omitnil,omitempty" name:"AllSliceNum"`

	// Number of completed parts
	CompleteNum *int64 `json:"CompleteNum,omitnil,omitempty" name:"CompleteNum"`
}

type ZoneConf struct {
	// AZ deployment mode. Value range: 0 (single-AZ), 1 (multi-AZ)
	DeployMode []*int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// AZ where the primary instance is located
	MasterZone []*string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// AZ where salve database 1 is located when the instance is deployed in multi-AZ mode
	SlaveZone []*string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// AZ where salve database 2 is located when the instance is deployed in multi-AZ mode
	BackupZone []*string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`
}