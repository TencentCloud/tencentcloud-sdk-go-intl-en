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
	// New account name
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// New account domain name
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
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The maximum number of instance connections supported by an account
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

// Predefined struct for user
type AddTimeWindowRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
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

	// Maximum delay threshold, which takes effect only for source instances and disaster recovery instances.
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type AddTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
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

	// Maximum delay threshold, which takes effect only for source instances and disaster recovery instances.
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

// Predefined struct for user
type AdjustCdbProxyAddressRequestParams struct {
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group Valid values: `true`, `false`.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// The delay threshold. Minimum value:  `0`
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// Whether to enable failover. Valid values: `true`, `false`.
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add newly created read-only instances. Valid values: `true`, `false`.
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether it is read-only. Valid values: `true`, `false`.
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// Address ID of the proxy group
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// Whether to enable transaction splitting. Valid values: `true`, `false`.
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Whether to enable the connection pool
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Assignment of read/write weights If `system` is passed in for `WeightMode`, only the default weight assigned by the system will take effect.
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`


	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`


	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

type AdjustCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group Valid values: `true`, `false`.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// The delay threshold. Minimum value:  `0`
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// Whether to enable failover. Valid values: `true`, `false`.
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to automatically add newly created read-only instances. Valid values: `true`, `false`.
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether it is read-only. Valid values: `true`, `false`.
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// Address ID of the proxy group
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// Whether to enable transaction splitting. Valid values: `true`, `false`.
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Whether to enable the connection pool
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Assignment of read/write weights If `system` is passed in for `WeightMode`, only the default weight assigned by the system will take effect.
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`

	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustCdbProxyAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustCdbProxyAddressResponseParams struct {
	// Async task ID Note: This field may return null, indicating that no valid values can be obtained.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// The specification configuration of a node
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Rebalance. Valid values:  `auto` (automatic), `manual` (manual).
	ReloadBalance *string `json:"ReloadBalance,omitnil,omitempty" name:"ReloadBalance"`

	// The upgrade switch time. Valid values:  `nowTime` (upgrade immediately), `timeWindow` (upgrade during instance maintenance time).
	UpgradeTime *string `json:"UpgradeTime,omitnil,omitempty" name:"UpgradeTime"`
}

type AdjustCdbProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// The specification configuration of a node
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
	// Async task ID Note: This field may return null, indicating that no valid values can be obtained.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting conditions for aggregation dimension
	AggregationConditions []*AggregationCondition `json:"AggregationConditions,omitnil,omitempty" name:"AggregationConditions"`

	// This parameter is disused. The result set of the audit log filtered by this condition is set as the analysis log.
	AuditLogFilter *AuditLogFilter `json:"AuditLogFilter,omitnil,omitempty" name:"AuditLogFilter"`

	// The result set of the audit log filtered by this condition is set as the analysis Log.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type AnalyzeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the log to be analyzed in the format of `2023-02-16 00:00:20`.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sorting conditions for aggregation dimension
	AggregationConditions []*AggregationCondition `json:"AggregationConditions,omitnil,omitempty" name:"AggregationConditions"`

	// This parameter is disused. The result set of the audit log filtered by this condition is set as the analysis log.
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
	// Information set of the aggregation bucket returned
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*AuditLogAggregationResult `json:"Items,omitnil,omitempty" name:"Items"`

	// Number of scanned logs
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be bound to the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be bound to the read-only replicas themselves.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
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

	// Database engine type.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database engine version.
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

	// Number of scanned rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	CheckRows *int64 `json:"CheckRows,omitnil,omitempty" name:"CheckRows"`

	// CPU execution time (μs)
	// Note: This field may return null, indicating that no valid values can be obtained.
	CpuTime *float64 `json:"CpuTime,omitnil,omitempty" name:"CpuTime"`

	// IO wait time (μs)
	// Note: This field may return null, indicating that no valid values can be obtained.
	IoWaitTime *uint64 `json:"IoWaitTime,omitnil,omitempty" name:"IoWaitTime"`

	// Lock wait time (μs)
	// Note: This field may return null, indicating that no valid values can be obtained.
	LockWaitTime *uint64 `json:"LockWaitTime,omitnil,omitempty" name:"LockWaitTime"`

	// Start time, which forms a time accurate to nanoseconds with·`timestamp`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NsTime *uint64 `json:"NsTime,omitnil,omitempty" name:"NsTime"`

	// Transaction duration (μs)
	// Note: This field may return null, indicating that no valid values can be obtained.
	TrxLivingTime *uint64 `json:"TrxLivingTime,omitnil,omitempty" name:"TrxLivingTime"`

	// Basic information on the rule template hit by the log.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	TemplateInfo []*LogRuleTemplateInfo `json:"TemplateInfo,omitnil,omitempty" name:"TemplateInfo"`
}

type AuditLogAggregationResult struct {
	// Aggregation dimension
	// Note: This field may return null, indicating that no valid values can be obtained.
	AggregationField *string `json:"AggregationField,omitnil,omitempty" name:"AggregationField"`

	// Result set of an aggregation bucket
	// Note: This field may return null, indicating that no valid values can be obtained.
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

	// Audit rule name
	// Note: This field may return `null`, indicating that no valid value was found.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Database instance name
	// Note: This field may return `null`, indicating that no valid value was found.
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
	// CPU utilization threshold (percent value). Valid values: 70, 80, and 90. Automatic scale-out will be triggered when CPU utilization reaches the set threshold.
	ExpandThreshold *int64 `json:"ExpandThreshold,omitnil,omitempty" name:"ExpandThreshold"`

	// Interval, in seconds. Valid values: 1, 3, 5, 10, 15, and 30. The system backend determines whether automatic scale-out is required at the set interval.
	ExpandPeriod *int64 `json:"ExpandPeriod,omitnil,omitempty" name:"ExpandPeriod"`

	// CPU utilization threshold (percent value). Valid values: 10, 20, and 30. Automatic scale-in will be triggered when CPU utilization reaches the set threshold.
	ShrinkThreshold *int64 `json:"ShrinkThreshold,omitnil,omitempty" name:"ShrinkThreshold"`

	// Interval, in seconds. Valid values: 5, 10, 15, and 30. The system backend determines whether automatic scale-in is required at the set interval.
	ShrinkPeriod *int64 `json:"ShrinkPeriod,omitnil,omitempty" name:"ShrinkPeriod"`
}

type BackupConfig struct {
	// Replication mode of secondary database 2. Value range: async, semi-sync
	ReplicationMode *string `json:"ReplicationMode,omitnil,omitempty" name:"ReplicationMode"`

	// Name of the AZ of secondary database 2, such as ap-shanghai-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Private IP address of secondary database 2
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Access port of secondary database 2
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type BackupInfo struct {
	// Backup filename
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Backup snapshot time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-03-17 02:10:37
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Download address
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// Download address
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// Log type. Valid values: `logical` (logical cold backup), `physical` (physical cold backup).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Backup subtask ID, which is used when backup files are deleted
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Backup task status. Valid values: `SUCCESS` (backup succeeded), `FAILED` (backup failed), `RUNNING` (backup is in progress).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Backup task completion time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// (This field will be disused and is thus not recommended) backup creator. Valid values: `SYSTEM` (created by system), `Uin` (initiator's `Uin` value).
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Backup task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup method. Valid values: `full` (full backup), `partial` (partial backup).
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Backup mode. Valid values: `manual` (manual backup), `automatic` (automatic backup).
	Way *string `json:"Way,omitnil,omitempty" name:"Way"`

	// Manual backup alias
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`

	// Backup retention type. Valid values: `save_mode_regular` (non-archive backup), save_mode_period`(archive backup).
	SaveMode *string `json:"SaveMode,omitnil,omitempty" name:"SaveMode"`

	// The region where local backup resides
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Detailed information of remote backups
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitnil,omitempty" name:"RemoteInfo"`

	// Storage method. Valid values: `0` (regular storage), `1`(archive storage). Default value: `0`.
	CosStorageType *int64 `json:"CosStorageType,omitnil,omitempty" name:"CosStorageType"`

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether the backup file is encrypted. Valid values: `on` (encrypted), `off` (unencrypted).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptionFlag *string `json:"EncryptionFlag,omitnil,omitempty" name:"EncryptionFlag"`
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
	// RO group ID in the format of `cdbrg-c1nl9rpv`.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`
}

type BalanceRoGroupLoadRequest struct {
	*tchttp.BaseRequest
	
	// RO group ID in the format of `cdbrg-c1nl9rpv`.
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
	// Binlog backup filename
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Backup file size in bytes
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// File stored time in the format of 2016-03-17 02:10:37
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Download address
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// Download address
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// Log type. Value range: binlog
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Binlog file start file
	BinlogStartTime *string `json:"BinlogStartTime,omitnil,omitempty" name:"BinlogStartTime"`

	// Binlog file end time
	BinlogFinishTime *string `json:"BinlogFinishTime,omitnil,omitempty" name:"BinlogFinishTime"`

	// The region where the binlog file resides
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Backup task status. Valid values: `SUCCESS` (backup succeeded), `FAILED` (backup failed), `RUNNING` (backup is in progress).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The detailed information of remote binlog backups
	RemoteInfo []*RemoteBackupInfo `json:"RemoteInfo,omitnil,omitempty" name:"RemoteInfo"`

	// Storage method. Valid values: `0` (regular storage), `1`(archive storage). Default value: `0`.
	CosStorageType *int64 `json:"CosStorageType,omitnil,omitempty" name:"CosStorageType"`

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type Bucket struct {
	// None
	// Note: This field may return null, indicating that no valid values can be obtained.
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

	// Instance type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance), `BASIC_V2` (basic v2 instance).
	// Note: This field may return null, indicating that no valid values can be obtained.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Engine type description. Valid values: `Innodb`, `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// Purchasable specifications ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type CdbSellType struct {
	// Name of the purchasable instance. Valid values: `Z3` (High-availability instance. `DeviceType`:`UNIVERSAL`, `EXCLUSIVE`; `CVM` (basic instance. `DeviceType`: `BASIC`); `TKE` (basic v2 instance. `DeviceType`: `BASIC_V2`).
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

	// Supported billing method. Valid values: `0` (monthly subscribed), `1` (hourly billed), `2` (pay-as-you-go)
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
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Whether only to disable read/write separation. Valid values: `true`, `false`. Default value: `false`.
	OnlyCloseRW *bool `json:"OnlyCloseRW,omitnil,omitempty" name:"OnlyCloseRW"`
}

type CloseCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
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
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Address ID of the proxy group
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

type CloseCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Address ID of the proxy group
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
type CloseWanServiceRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseWanServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

	// Password of the new account
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Remarks
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

	// Password of the new account
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Remarks
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
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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

	// Information of the table to be backed up. If this parameter is not set, the entire instance will be backed up by default. It can be set only in logical backup (i.e., BackupMethod = logical). The specified table must exist; otherwise, backup may fail.
	// For example, if you want to backup tb1 and tb2 in db1 and the entire db2, you should set the parameter as [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ].
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitnil,omitempty" name:"BackupDBTableList"`

	// Manual backup alias
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target backup method. Valid values: `logical` (logical cold backup), `physical` (physical cold backup), `snapshot` (snapshot backup). Basic Edition instances only support snapshot backups.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Information of the table to be backed up. If this parameter is not set, the entire instance will be backed up by default. It can be set only in logical backup (i.e., BackupMethod = logical). The specified table must exist; otherwise, backup may fail.
	// For example, if you want to backup tb1 and tb2 in db1 and the entire db2, you should set the parameter as [{"Db": "db1", "Table": "tb1"}, {"Db": "db1", "Table": "tb2"}, {"Db": "db2"} ].
	BackupDBTableList []*BackupItem `json:"BackupDBTableList,omitnil,omitempty" name:"BackupDBTableList"`

	// Manual backup alias
	ManualBackupName *string `json:"ManualBackupName,omitnil,omitempty" name:"ManualBackupName"`
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
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group. Valid values: `true`, `false`.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// The delay threshold. Minimum value:  `0`
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

	// VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Whether to enable the connection pool. Valid values: 
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// IP address
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Connection pool type, which will take effect only when `ConnectionPool` is `true`. Valid values:  `transaction` (transaction-level), `connection` (session-level).
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`


	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`


	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`
}

type CreateCdbProxyAddressRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group. Valid values: `true`, `false`.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// The delay threshold. Minimum value:  `0`
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

	// VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Whether to enable the connection pool. Valid values: 
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// IP address
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Connection pool type, which will take effect only when `ConnectionPool` is `true`. Valid values:  `transaction` (transaction-level), `connection` (session-level).
	ConnectionPoolType *string `json:"ConnectionPoolType,omitnil,omitempty" name:"ConnectionPoolType"`

	AutoLoadBalance *bool `json:"AutoLoadBalance,omitnil,omitempty" name:"AutoLoadBalance"`

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
	// Async task ID Note: This field may return null, indicating that no valid values can be obtained.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// The specification configuration of a node
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Connection pool threshold
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
}

type CreateCdbProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// The specification configuration of a node
	ProxyNodeCustom []*ProxyNodeCustom `json:"ProxyNodeCustom,omitnil,omitempty" name:"ProxyNodeCustom"`

	// Security group
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Connection pool threshold
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCdbProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCdbProxyResponseParams struct {
	// Async task ID Note: This field may return null, indicating that no valid values can be obtained.
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
	// ID of the instance to be cloned from
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// To roll back the cloned instance to a specific point in time, set this parameter to a value in the format of "yyyy-mm-dd hh:mm:ss".
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitnil,omitempty" name:"SpecifiedRollbackTime"`

	// To roll back the cloned instance to a specific physical backup file, set this parameter to the ID of the physical backup file. The ID can be obtained by the [DescribeBackups](https://intl.cloud.tencent.com/document/api/236/15842?from_cn_redirect=1) API.
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitnil,omitempty" name:"SpecifiedBackupId"`

	// VPC ID, which can be obtained by the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API. If this parameter is left empty, the classic network will be used by default.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID, which can be obtained by the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API. If `UniqVpcId` is set, `UniqSubnetId` will be required.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Memory of the cloned instance in MB, which should be equal to (by default) or larger than that of the original instance
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk capacity of the cloned instance in GB, which should be equal to (by default) or larger than that of the original instance
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Name of the cloned instance
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Security group parameter, which can be obtained by the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Information of the cloned instance tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// The number of CPU cores of the cloned instance. It should be equal to (by default) or larger than that of the original instance.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Multi-AZ or single-AZ. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Availability zone information of replica 1 of the cloned instance, which is the same as the value of `Zone` of the original instance by default
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Availability zone information of replica 2 of the cloned instance, 
	// which is left empty by default. Specify this parameter when cloning a strong sync source instance.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Resource isolation type of the clone. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// The number of nodes of the clone. If this parameter is set to `3` or the `BackupZone` parameter is specified, the clone will have three nodes. If this parameter is set to `2` or left empty, the clone will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// Project ID. Default value: 0.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type CreateCloneInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be cloned from
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// To roll back the cloned instance to a specific point in time, set this parameter to a value in the format of "yyyy-mm-dd hh:mm:ss".
	SpecifiedRollbackTime *string `json:"SpecifiedRollbackTime,omitnil,omitempty" name:"SpecifiedRollbackTime"`

	// To roll back the cloned instance to a specific physical backup file, set this parameter to the ID of the physical backup file. The ID can be obtained by the [DescribeBackups](https://intl.cloud.tencent.com/document/api/236/15842?from_cn_redirect=1) API.
	SpecifiedBackupId *int64 `json:"SpecifiedBackupId,omitnil,omitempty" name:"SpecifiedBackupId"`

	// VPC ID, which can be obtained by the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API. If this parameter is left empty, the classic network will be used by default.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID, which can be obtained by the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API. If `UniqVpcId` is set, `UniqSubnetId` will be required.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Memory of the cloned instance in MB, which should be equal to (by default) or larger than that of the original instance
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk capacity of the cloned instance in GB, which should be equal to (by default) or larger than that of the original instance
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Name of the cloned instance
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Security group parameter, which can be obtained by the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Information of the cloned instance tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// The number of CPU cores of the cloned instance. It should be equal to (by default) or larger than that of the original instance.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Multi-AZ or single-AZ. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Availability zone information of replica 1 of the cloned instance, which is the same as the value of `Zone` of the original instance by default
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Availability zone information of replica 2 of the cloned instance, 
	// which is left empty by default. Specify this parameter when cloning a strong sync source instance.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Resource isolation type of the clone. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// The number of nodes of the clone. If this parameter is set to `3` or the `BackupZone` parameter is specified, the clone will have three nodes. If this parameter is set to `2` or left empty, the clone will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// Project ID. Default value: 0.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloneInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloneInstanceResponseParams struct {
	// LimitAsync task request ID, which can be used to query the execution result of an async task
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
	// Number of instances. Value range: 1-100. Default value: 1.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance memory size in MB. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported memory specifications.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported disk specifications.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// MySQL version. Valid values: `5.5`, `5.6`, `5.7`, `8.0`. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported versions.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// VPC ID. If this parameter is not passed in, the basic network will be selected by default. Please use the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API to query the VPCs.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API to query the subnet lists.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Project ID. If this is left empty, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// AZ information. By default, the system will automatically select an AZ. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported AZs.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance ID, which is required and the same as the primary instance ID when purchasing read-only or disaster recovery instances. Please use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the instance IDs.
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// Instance type. Valid values: master (primary instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Region information of the source instance, which is required when purchasing a read-only or disaster recovery instance.
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// Custom port. Value range: [1024-65535].
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Sets the root account password. Rule: the password can contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special symbols (_+-&=!@#$%^*()). This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// List of parameters in the format of `ParamList.0.Name=auto_increment&ParamList.0.Value=1`. You can use the [DescribeDefaultParams](https://intl.cloud.tencent.com/document/api/236/32662?from_cn_redirect=1) API to query the configurable parameters.
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Multi-AZ. Valid value: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// AZ information of secondary database 1, which is the `Zone` value by default. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// The availability zone information of Replica 2, which is left empty by default. Specify this parameter when purchasing a source instance in the one-source-two-replica architecture.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Security group parameter. You can use the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API to query the security group details of a project.
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Read-only instance information. This parameter must be passed in when purchasing read-only instances.
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// This field is meaningless when purchasing pay-as-you-go instances.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Instance name For multiple instances purchased at one time, they will be distinguished by the name suffix number, such as instnaceName=db and goodsNum=3, and their instance names are db1, db2, and db3, respectively.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance tag information.
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Instance resource isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Parameter template ID.
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Array of alarm policy IDs,  which is `OriginId` obtained through the `DescribeAlarmPolicy` API.
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// The number of nodes of the instance. To purchase a read-only replica or a basic instance, set this parameter to `1` or leave it empty. To purchase a three-node instance, set this parameter to `3` or specify the `BackupZone` parameter. If the instance to be purchased is a source instance and both `BackupZone` and this parameter are left empty, the value `2` will be used, which indicates the source instance will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// The number of CPU cores of the instance. If this parameter is left empty, the number of CPU cores depends on the `Memory` value.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Whether to automatically start disaster recovery synchronization. This parameter takes effect only for disaster recovery instances. Valid values: `0` (no), `1` (yes). Default value: `0`.
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template). Default value: `HIGH_STABILITY`.
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// The array of alarm policy names, such as ["policy-uyoee9wg"]. If the `AlarmPolicyList` parameter is specified, this parameter is invalid.
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// Instance engine type. Valid values: `InnoDB` (default); `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// The list of IPs for sources instances. Only one IP address can be assigned to a single source instance. If all IPs are used up, the system will automatically assign IPs to the remaining source instances that do not have one.
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// Number of instances. Value range: 1-100. Default value: 1.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance memory size in MB. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported memory specifications.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported disk specifications.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// MySQL version. Valid values: `5.5`, `5.6`, `5.7`, `8.0`. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported versions.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// VPC ID. If this parameter is not passed in, the basic network will be selected by default. Please use the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API to query the VPCs.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. Please use the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API to query the subnet lists.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Project ID. If this is left empty, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// AZ information. By default, the system will automatically select an AZ. Please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported AZs.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance ID, which is required and the same as the primary instance ID when purchasing read-only or disaster recovery instances. Please use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the instance IDs.
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// Instance type. Valid values: master (primary instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Region information of the source instance, which is required when purchasing a read-only or disaster recovery instance.
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// Custom port. Value range: [1024-65535].
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Sets the root account password. Rule: the password can contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special symbols (_+-&=!@#$%^*()). This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// List of parameters in the format of `ParamList.0.Name=auto_increment&ParamList.0.Value=1`. You can use the [DescribeDefaultParams](https://intl.cloud.tencent.com/document/api/236/32662?from_cn_redirect=1) API to query the configurable parameters.
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). Default value: 0. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Multi-AZ. Valid value: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// AZ information of secondary database 1, which is the `Zone` value by default. This parameter can be specified when purchasing primary instances and is meaningless for read-only or disaster recovery instances.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// The availability zone information of Replica 2, which is left empty by default. Specify this parameter when purchasing a source instance in the one-source-two-replica architecture.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Security group parameter. You can use the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API to query the security group details of a project.
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Read-only instance information. This parameter must be passed in when purchasing read-only instances.
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// This field is meaningless when purchasing pay-as-you-go instances.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Instance name For multiple instances purchased at one time, they will be distinguished by the name suffix number, such as instnaceName=db and goodsNum=3, and their instance names are db1, db2, and db3, respectively.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance tag information.
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Instance resource isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Parameter template ID.
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Array of alarm policy IDs,  which is `OriginId` obtained through the `DescribeAlarmPolicy` API.
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// The number of nodes of the instance. To purchase a read-only replica or a basic instance, set this parameter to `1` or leave it empty. To purchase a three-node instance, set this parameter to `3` or specify the `BackupZone` parameter. If the instance to be purchased is a source instance and both `BackupZone` and this parameter are left empty, the value `2` will be used, which indicates the source instance will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// The number of CPU cores of the instance. If this parameter is left empty, the number of CPU cores depends on the `Memory` value.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Whether to automatically start disaster recovery synchronization. This parameter takes effect only for disaster recovery instances. Valid values: `0` (no), `1` (yes). Default value: `0`.
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template). Default value: `HIGH_STABILITY`.
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// The array of alarm policy names, such as ["policy-uyoee9wg"]. If the `AlarmPolicyList` parameter is specified, this parameter is invalid.
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// Instance engine type. Valid values: `InnoDB` (default); `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// The list of IPs for sources instances. Only one IP address can be assigned to a single source instance. If all IPs are used up, the system will automatically assign IPs to the remaining source instances that do not have one.
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// Short order ID.
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// Instance ID list
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
	// Instance memory size in MB. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported memory specifications.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported disk specifications.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Instance validity period in months. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of instances. Value range: 1-100. Default value: `1`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// AZ information. The system will automatically select an AZ by default. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported AZs.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID. If this parameter is not passed in, the basic network will be selected by default. You can use the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API to query the VPCs.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. You can use the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API to query the subnet lists.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Project ID. If this parameter is left empty, the default project will be used. When you purchase read-only instances and disaster recovery instances, the project ID is the same as that of the source instance by default.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Custom port. Value range: 1024-65535.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance typeA. Valid values: `master` (source instance), `dr` (disaster recovery instance), `ro` (read-only instance).
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Instance ID. It is required when purchasing a read-only instance, which is the same as the source instance ID. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the instance ID.
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// MySQL version. Valid values: `5.5`, `5.6`, `5.7`, and `8.0`. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported instance versions.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// The root account password, which can contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and symbols `_+-&=!@#$%^*()`. This parameter applies to source instances but not to read-only or disaster recovery instances.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Data replication mode. Valid values: `0` (async replication), `1` (semi-sync replication), `2` (strong sync replication). Default value: `0`.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Multi-AZ or single-AZ. Valid values: `0` (single-AZ), `1` (multi-AZ). Default value: `0`.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Information of replica AZ 1, which is the `Zone` value by default.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// List of parameters in the format of ParamList.0.Name=auto_increment&ParamList.0.Value=1. You can use the [DescribeDefaultParams](https://intl.cloud.tencent.com/document/api/236/32662?from_cn_redirect=1) API to query the configurable parameters.
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Information of replica AZ 2, which is left empty by default. Specify this parameter when purchasing a source instance in the one-source-two-replica architecture.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Auto-renewal flag. Valid values: `0` (auto-renewal not enabled), `1` (auto-renewal enabled).
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Region information of the source instance, which is required when purchasing a read-only or disaster recovery instance.
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// Security group parameter. You can use the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API to query the security group details of a project.
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Read-only instance parameter. This parameter must be passed in when purchasing read-only instances.
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// Instance name. For multiple instances purchased at one time, they will be distinguished by the name suffix number, such as instnaceName=db and goodsNum=3, and their instance names are db1, db2, and db3, respectively.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance tag information
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Placement group ID
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// A string unique in 48 hours, which is supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Instance isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Array of alarm policy IDs, which can be obtained through the `OriginId` field in the return value of the `DescribeAlarmPolicy` API of TCOP.
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// The number of nodes of the instance. To purchase a read-only instance or a basic instance, set this parameter to `1` or leave it empty. To purchase a three-node instance, set this parameter to `3` or specify the `BackupZone` parameter. If the instance to be purchased is a source instance and both `BackupZone` and this parameter are left empty, the value `2` will be used, which indicates the source instance will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// The number of the instance CPU cores. If this parameter is left empty, it will be subject to the `Memory` value.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Whether to automatically start disaster recovery synchronization. This parameter takes effect only for disaster recovery instances. Valid values: `0` (no), `1` (yes). Default value: `0`.
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// The array of alarm policy names, such as ["policy-uyoee9wg"]. If the `AlarmPolicyList` parameter is specified, this parameter is invalid.
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// Instance engine type. Valid values: `InnoDB` (default), `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// The list of IPs for sources instances. Only one IP address can be assigned to a single source instance. If all IPs are used up, the system will automatically assign IPs to the remaining source instances that do not have one.
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance memory size in MB. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported memory specifications.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported disk specifications.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Instance validity period in months. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Number of instances. Value range: 1-100. Default value: `1`.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// AZ information. The system will automatically select an AZ by default. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported AZs.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID. If this parameter is not passed in, the basic network will be selected by default. You can use the [DescribeVpcs](https://intl.cloud.tencent.com/document/api/215/15778?from_cn_redirect=1) API to query the VPCs.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC subnet ID. If `UniqVpcId` is set, then `UniqSubnetId` will be required. You can use the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API to query the subnet lists.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Project ID. If this parameter is left empty, the default project will be used. When you purchase read-only instances and disaster recovery instances, the project ID is the same as that of the source instance by default.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Custom port. Value range: 1024-65535.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Instance typeA. Valid values: `master` (source instance), `dr` (disaster recovery instance), `ro` (read-only instance).
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Instance ID. It is required when purchasing a read-only instance, which is the same as the source instance ID. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the instance ID.
	MasterInstanceId *string `json:"MasterInstanceId,omitnil,omitempty" name:"MasterInstanceId"`

	// MySQL version. Valid values: `5.5`, `5.6`, `5.7`, and `8.0`. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1) API to query the supported instance versions.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// The root account password, which can contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and symbols `_+-&=!@#$%^*()`. This parameter applies to source instances but not to read-only or disaster recovery instances.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Data replication mode. Valid values: `0` (async replication), `1` (semi-sync replication), `2` (strong sync replication). Default value: `0`.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Multi-AZ or single-AZ. Valid values: `0` (single-AZ), `1` (multi-AZ). Default value: `0`.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Information of replica AZ 1, which is the `Zone` value by default.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// List of parameters in the format of ParamList.0.Name=auto_increment&ParamList.0.Value=1. You can use the [DescribeDefaultParams](https://intl.cloud.tencent.com/document/api/236/32662?from_cn_redirect=1) API to query the configurable parameters.
	ParamList []*ParamInfo `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Information of replica AZ 2, which is left empty by default. Specify this parameter when purchasing a source instance in the one-source-two-replica architecture.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Auto-renewal flag. Valid values: `0` (auto-renewal not enabled), `1` (auto-renewal enabled).
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Region information of the source instance, which is required when purchasing a read-only or disaster recovery instance.
	MasterRegion *string `json:"MasterRegion,omitnil,omitempty" name:"MasterRegion"`

	// Security group parameter. You can use the [DescribeProjectSecurityGroups](https://intl.cloud.tencent.com/document/api/236/15850?from_cn_redirect=1) API to query the security group details of a project.
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// Read-only instance parameter. This parameter must be passed in when purchasing read-only instances.
	RoGroup *RoGroup `json:"RoGroup,omitnil,omitempty" name:"RoGroup"`

	// Instance name. For multiple instances purchased at one time, they will be distinguished by the name suffix number, such as instnaceName=db and goodsNum=3, and their instance names are db1, db2, and db3, respectively.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance tag information
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// Placement group ID
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// A string unique in 48 hours, which is supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Instance isolation type. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Parameter template ID
	ParamTemplateId *int64 `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// Array of alarm policy IDs, which can be obtained through the `OriginId` field in the return value of the `DescribeAlarmPolicy` API of TCOP.
	AlarmPolicyList []*int64 `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// The number of nodes of the instance. To purchase a read-only instance or a basic instance, set this parameter to `1` or leave it empty. To purchase a three-node instance, set this parameter to `3` or specify the `BackupZone` parameter. If the instance to be purchased is a source instance and both `BackupZone` and this parameter are left empty, the value `2` will be used, which indicates the source instance will have two nodes.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// The number of the instance CPU cores. If this parameter is left empty, it will be subject to the `Memory` value.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Whether to automatically start disaster recovery synchronization. This parameter takes effect only for disaster recovery instances. Valid values: `0` (no), `1` (yes). Default value: `0`.
	AutoSyncFlag *int64 `json:"AutoSyncFlag,omitnil,omitempty" name:"AutoSyncFlag"`

	// Financial cage ID.
	CageId *string `json:"CageId,omitnil,omitempty" name:"CageId"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	ParamTemplateType *string `json:"ParamTemplateType,omitnil,omitempty" name:"ParamTemplateType"`

	// The array of alarm policy names, such as ["policy-uyoee9wg"]. If the `AlarmPolicyList` parameter is specified, this parameter is invalid.
	AlarmPolicyIdList []*string `json:"AlarmPolicyIdList,omitnil,omitempty" name:"AlarmPolicyIdList"`

	// Whether to check the request without creating any instance. Valid values: `true`, `false` (default). After being submitted, the request will be checked to see if it is in correct format and has all required parameters with valid values. An error code is returned if the check failed, and `RequestId` is returned if the check succeeded. After a successful check, no instance will be created if this parameter is set to `true`, whereas an instance will be created and if it is set to `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// Instance engine type. Valid values: `InnoDB` (default), `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// The list of IPs for sources instances. Only one IP address can be assigned to a single source instance. If all IPs are used up, the system will automatically assign IPs to the remaining source instances that do not have one.
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// Billing sub-order ID
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// List of instance IDs
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


	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// Character set. Valid values:  `utf8`, `gbk`, `latin1`, `utf8mb4`.
	CharacterSetName *string `json:"CharacterSetName,omitnil,omitempty" name:"CharacterSetName"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `cdb-c1nl9rpv`,  which is the same as the one displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

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
	// Parameter template name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MySQL version number.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Source parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Instance engine type. Valid values: `InnoDB` (default), `RocksDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// MySQL version number.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Source parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Instance engine type. Valid values: `InnoDB` (default), `RocksDB`.
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
	// Instance ID, in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB for MySQL console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Information about the account for which password rotation needs to be enabled. The account and host names are included.
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type CreateRotationPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed on the TencentDB for MySQL console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Information about the account for which password rotation needs to be enabled. The account and host names are included.
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
	// Device
	// Note: this field may return `null`, indicating that no valid value can be found.
	Device *string `json:"Device,omitnil,omitempty" name:"Device"`

	// Type
	// Note: this field may return `null`, indicating that no valid value can be found.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Device type
	// Note: this field may return `null`, indicating that no valid value can be found.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Memory
	// Note: this field may return `null`, indicating that no valid value can be found.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Number of CPU cores
	// Note: this field may return `null`, indicating that no valid value can be found.
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
	// Audit log file name, which can be obtained through the [DescribeAuditLogFiles](https://www.tencentcloud.comom/document/api/236/45454?from_cn_redirect=1) API.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// Audit log file name, which can be obtained through the [DescribeAuditLogFiles](https://www.tencentcloud.comom/document/api/236/45454?from_cn_redirect=1) API.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API. A maximum of 5 rule templates can be deleted per request.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type DeleteAuditRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API. A maximum of 5 rule templates can be deleted per request.
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

	// Backup task ID, which is the task ID returned by the [TencentDB instance backup creating API](https://intl.cloud.tencent.com/document/api/236/15844?from_cn_redirect=1).
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DeleteBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup task ID, which is the task ID returned by the [TencentDB instance backup creating API](https://intl.cloud.tencent.com/document/api/236/15844?from_cn_redirect=1).
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
type DeleteParamTemplateRequestParams struct {
	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID.
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

	// Database user account.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database account domain name.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database user account.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database account domain name.
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
	// Task execution result. Valid values: INITIAL, RUNNING, SUCCESS, FAILED, KILLED, REMOVED, PAUSED.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task execution information.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The pagination parameter, which specifies the number of entries per page. Maximum value: 100 (default).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting order Valid values: `ASC (ascending), `DESC` (descending).
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field Valid values: 
	// `timestamp`: Timestamp,
	// `affectRows`: Number of affected rows,
	// `execTime`: Execution time.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Filter. Multiple values are in `AND` relationship.
	LogFilter []*InstanceAuditLogFilters `json:"LogFilter,omitnil,omitempty" name:"LogFilter"`
}

type DescribeAuditLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The pagination parameter, which specifies the number of entries per page. Maximum value: 100 (default).
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting order Valid values: `ASC (ascending), `DESC` (descending).
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field Valid values: 
	// `timestamp`: Timestamp,
	// `affectRows`: Number of affected rows,
	// `execTime`: Execution time.
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

	// Audit policy details
	// Note: This field may return `null`, indicating that no valid value was found.
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
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API.
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
	
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API.
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
	StartTimeMin *int64 `json:"StartTimeMin,omitnil,omitempty" name:"StartTimeMin"`

	// Latest start time point of automatic backup, such as 6 (for 6:00 AM). (This field has been disused. You are recommended to use the `BackupTimeWindow` field)
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
	// Instance ID in the format of  cdb-XXXX,  which is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance backup ID, which can be obtained by the `DescribeBackups` API.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DescribeBackupDecryptionKeyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of  cdb-XXXX,  which is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance backup ID, which can be obtained by the `DescribeBackups` API.
	BackupId *int64 `json:"BackupId,omitnil,omitempty" name:"BackupId"`
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
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
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

	// Total capacity of backups of a user in the current region
	// Note: This field may return null, indicating that no valid value can be obtained.
	RemoteBackupVolume *int64 `json:"RemoteBackupVolume,omitnil,omitempty" name:"RemoteBackupVolume"`

	// Archive backup capacity, which includes data backups and log backups.
	// Note: This field may return null, indicating that no valid value can be obtained.
	BackupArchiveVolume *int64 `json:"BackupArchiveVolume,omitnil,omitempty" name:"BackupArchiveVolume"`

	// Backup capacity of standard storage, which includes data backups and log backups.
	// Note: This field may return null, indicating that no valid value can be obtained.
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
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
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
	
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
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
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
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
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeBinlogBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
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

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The earliest start time of binlog  in the format of 2016-03-17 02:10:37.
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// The latest start time of binlog  in the format of 2016-03-17 02:10:37.
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`
}

type DescribeBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Minimum value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The earliest start time of binlog  in the format of 2016-03-17 02:10:37.
	MinStartTime *string `json:"MinStartTime,omitnil,omitempty" name:"MinStartTime"`

	// The latest start time of binlog  in the format of 2016-03-17 02:10:37.
	MaxStartTime *string `json:"MaxStartTime,omitnil,omitempty" name:"MaxStartTime"`
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
type DescribeCdbProxyInfoRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type DescribeCdbProxyInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
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
	// Number of proxy groups Note: This field may return null, indicating that no valid values can be obtained.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Proxy group information Note: This field may return null, indicating that no valid values can be obtained.
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
	// ID of the original instance. This parameter is used to query the clone task list of a specific original instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCloneListRequest struct {
	*tchttp.BaseRequest
	
	// ID of the original instance. This parameter is used to query the clone task list of a specific original instance.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results per page. Default value: `20`.
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
type DescribeCpuExpandStrategyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCpuExpandStrategyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCpuExpandStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCpuExpandStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCpuExpandStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCpuExpandStrategyResponseParams struct {
	// Policy type. Valid values: `auto`, `manual`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Manually expanded CPU, which is valid when `Type` is `manual`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpandCpu *string `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Automatic expansion policy, which is valid when `Type` is `auto`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoStrategy *string `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCpuExpandStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCpuExpandStrategyResponseParams `json:"Response"`
}

func (r *DescribeCpuExpandStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCpuExpandStrategyResponse) FromJsonString(s string) error {
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

	// Region of the source instance
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
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
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
	// Data protection mode of the primary instance. Value range: 0 (async replication), 1 (semi-sync replication), 2 (strong sync replication).
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Master instance deployment mode. Value range: 0 (single-AZ), 1 (multi-AZ)
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Instance AZ information in the format of "ap-shanghai-2".
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Configurations of the replica node
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	SlaveConfig *SlaveConfig `json:"SlaveConfig,omitnil,omitempty" name:"SlaveConfig"`

	// Configurations of the second replica node of a strong-sync instance
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	BackupConfig *BackupConfig `json:"BackupConfig,omitnil,omitempty" name:"BackupConfig"`

	// This parameter is only available for multi-AZ instances. It indicates whether the source AZ is the same as the one specified upon purchase. `true`: not the same, `false`: the same.
	Switched *bool `json:"Switched,omitnil,omitempty" name:"Switched"`

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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
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

	// Encryption key ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key region.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// The default region of the KMS service currently used by the TencentDB backend service.
	// Note: this field may return `null`, indicating that no valid value can be found.
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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceLogToCLSRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceLogToCLSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceLogToCLSResponseParams struct {
	// Configurations of sending error logs to CLS.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	ErrorLog *LogToCLSConfig `json:"ErrorLog,omitnil,omitempty" name:"ErrorLog"`

	// Configurations of sending slow logs to CLS.
	// Note: The return value may be null, indicating that no valid data can be obtained.
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
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeDBInstanceRebootTimeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
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

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only).
	InstanceTypes []*uint64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Private IP address of the instance.
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// Instance status. Valid values: <br>`0` (creating) <br>`1` (running) <br>`4` (isolating) <br>`5` (isolated; the instance can be restored and started in the recycle bin)
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Security group ID. When it is used as a filter, the `WithSecurityGroup` parameter should be set to 1.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Payment type. Valid values: 0 - yearly/monthly subscription; 1 - bill by hour.
	PayTypes []*uint64 `json:"PayTypes,omitnil,omitempty" name:"PayTypes"`

	// Instance name.
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Instance task status. Valid values:<br>0 - no task;<br>1 - upgrading;<br>2 - importing data;<br>3 - enabling secondary nodes;<br>4 - enabling public network access;<br>5 - executing batch operations;<br>6 - rolling back;<br>7 - disabling public network access;<br>8 - changing the password;<br>9 - renaming the instance;<br>10 - restarting;<br>12 - migrating self-built databases;<br>13 - deleting databases and tables;<br>14 - synchronizing the creation of disaster recovery instances;<br>15 - pending upgrade switch;<br>16 - under upgrade switch;<br>17 - upgrade switch completed;<br>19 - parameter settings pending execution;<br>34 - in-place upgrade pending execution.
	TaskStatus []*uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Version of the instance database engine. Value range: 5.1, 5.5, 5.6, 5.7.
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// VPC ID.
	VpcIds []*uint64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// AZ ID.
	ZoneIds []*uint64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Subnet ID.
	SubnetIds []*uint64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Whether to lock disk write. Valid values: `0`(unlock), `1`(lock). Default value: 0.
	CdbErrors []*int64 `json:"CdbErrors,omitnil,omitempty" name:"CdbErrors"`

	// Sorting field of the query results. Valid values: "instanceId", "instanceName", "createTime", and "deadlineTime".
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method of the returned result set. Valid values: "ASC" - ascending order; "DESC" - descending order. The default value is "DESC".
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// Whether to use the security group ID as the filter condition.
	// Note: 0 indicates no; 1 indicates yes.
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitnil,omitempty" name:"WithSecurityGroup"`

	// Whether dedicated cluster details are included. Value range: 0 (not included), 1 (included)
	WithExCluster *int64 `json:"WithExCluster,omitnil,omitempty" name:"WithExCluster"`

	// Exclusive cluster ID.
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// Whether instances corresponding to the disaster recovery relationship are included. Valid values: 0 (not included), 1 (included). Default value: 1. If a primary instance is pulled, the data of the disaster recovery relationship will be in the `DrInfo` field. If a disaster recovery instance is pulled, the data of the disaster recovery relationship will be in the `MasterInfo` field. The disaster recovery relationship contains only partial basic data. To get the detailed data, you need to call an API to pull it.
	WithDr *int64 `json:"WithDr,omitnil,omitempty" name:"WithDr"`

	// Whether read-only instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithRo *int64 `json:"WithRo,omitnil,omitempty" name:"WithRo"`

	// Whether primary instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithMaster *int64 `json:"WithMaster,omitnil,omitempty" name:"WithMaster"`

	// Placement group ID list.
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`

	// Whether to use the tag key as a filter condition
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitnil,omitempty" name:"TagKeysForSearch"`

	// Financial cage IDs.
	CageIds []*string `json:"CageIds,omitnil,omitempty" name:"CageIds"`

	// Tag value
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`

	// VPC character vpcId
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitnil,omitempty" name:"UniqueVpcIds"`

	// VPC character subnetId
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// Tag key value.
	// Note that tags cannot be queried for instances being created.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Database proxy IP
	ProxyVips []*string `json:"ProxyVips,omitnil,omitempty" name:"ProxyVips"`

	// Database proxy ID
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// Database engine type. Valid values: InnoDB; RocksDB.
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// Whether to obtain the Cluster Edition instance node information. Valid values: true or false. The default value is false.
	QueryClusterInfo *bool `json:"QueryClusterInfo,omitnil,omitempty" name:"QueryClusterInfo"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only).
	InstanceTypes []*uint64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Private IP address of the instance.
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// Instance status. Valid values: <br>`0` (creating) <br>`1` (running) <br>`4` (isolating) <br>`5` (isolated; the instance can be restored and started in the recycle bin)
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Offset. Default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Security group ID. When it is used as a filter, the `WithSecurityGroup` parameter should be set to 1.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Payment type. Valid values: 0 - yearly/monthly subscription; 1 - bill by hour.
	PayTypes []*uint64 `json:"PayTypes,omitnil,omitempty" name:"PayTypes"`

	// Instance name.
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// Instance task status. Valid values:<br>0 - no task;<br>1 - upgrading;<br>2 - importing data;<br>3 - enabling secondary nodes;<br>4 - enabling public network access;<br>5 - executing batch operations;<br>6 - rolling back;<br>7 - disabling public network access;<br>8 - changing the password;<br>9 - renaming the instance;<br>10 - restarting;<br>12 - migrating self-built databases;<br>13 - deleting databases and tables;<br>14 - synchronizing the creation of disaster recovery instances;<br>15 - pending upgrade switch;<br>16 - under upgrade switch;<br>17 - upgrade switch completed;<br>19 - parameter settings pending execution;<br>34 - in-place upgrade pending execution.
	TaskStatus []*uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Version of the instance database engine. Value range: 5.1, 5.5, 5.6, 5.7.
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// VPC ID.
	VpcIds []*uint64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// AZ ID.
	ZoneIds []*uint64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Subnet ID.
	SubnetIds []*uint64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// Whether to lock disk write. Valid values: `0`(unlock), `1`(lock). Default value: 0.
	CdbErrors []*int64 `json:"CdbErrors,omitnil,omitempty" name:"CdbErrors"`

	// Sorting field of the query results. Valid values: "instanceId", "instanceName", "createTime", and "deadlineTime".
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting method of the returned result set. Valid values: "ASC" - ascending order; "DESC" - descending order. The default value is "DESC".
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// Whether to use the security group ID as the filter condition.
	// Note: 0 indicates no; 1 indicates yes.
	WithSecurityGroup *int64 `json:"WithSecurityGroup,omitnil,omitempty" name:"WithSecurityGroup"`

	// Whether dedicated cluster details are included. Value range: 0 (not included), 1 (included)
	WithExCluster *int64 `json:"WithExCluster,omitnil,omitempty" name:"WithExCluster"`

	// Exclusive cluster ID.
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// Instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized).
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// Whether instances corresponding to the disaster recovery relationship are included. Valid values: 0 (not included), 1 (included). Default value: 1. If a primary instance is pulled, the data of the disaster recovery relationship will be in the `DrInfo` field. If a disaster recovery instance is pulled, the data of the disaster recovery relationship will be in the `MasterInfo` field. The disaster recovery relationship contains only partial basic data. To get the detailed data, you need to call an API to pull it.
	WithDr *int64 `json:"WithDr,omitnil,omitempty" name:"WithDr"`

	// Whether read-only instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithRo *int64 `json:"WithRo,omitnil,omitempty" name:"WithRo"`

	// Whether primary instances are included. Valid values: 0 (not included), 1 (included). Default value: 1.
	WithMaster *int64 `json:"WithMaster,omitnil,omitempty" name:"WithMaster"`

	// Placement group ID list.
	DeployGroupIds []*string `json:"DeployGroupIds,omitnil,omitempty" name:"DeployGroupIds"`

	// Whether to use the tag key as a filter condition
	TagKeysForSearch []*string `json:"TagKeysForSearch,omitnil,omitempty" name:"TagKeysForSearch"`

	// Financial cage IDs.
	CageIds []*string `json:"CageIds,omitnil,omitempty" name:"CageIds"`

	// Tag value
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`

	// VPC character vpcId
	UniqueVpcIds []*string `json:"UniqueVpcIds,omitnil,omitempty" name:"UniqueVpcIds"`

	// VPC character subnetId
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// Tag key value.
	// Note that tags cannot be queried for instances being created.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Database proxy IP
	ProxyVips []*string `json:"ProxyVips,omitnil,omitempty" name:"ProxyVips"`

	// Database proxy ID
	ProxyIds []*string `json:"ProxyIds,omitnil,omitempty" name:"ProxyIds"`

	// Database engine type. Valid values: InnoDB; RocksDB.
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// Whether to obtain the Cluster Edition instance node information. Valid values: true or false. The default value is false.
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
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instance details
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
	// Instance validity period in months. Value range: 1-36. This field is invalid when querying the prices of pay-as-you-go instances.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// AZ information in the format of "ap-guangzhou-3". You can use the <a href="https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1">DescribeDBZoneConfig</a> API to query the configurable values. This parameter is required when `InstanceId` is empty.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of instances. Value range: 1-100. Default value: 1. This parameter is required when `InstanceId` is empty.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance memory size in MB. This parameter is required when `InstanceId` is empty.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB. This parameter is required when `InstanceId` is empty.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Instance type. Valid values: `master` (source instance), `dr` (disaster recovery instance), `ro` (read-only instance). Default value: `master`. This parameter is required when `InstanceId` is empty.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Billing mode. Valid values: `PRE_PAID` (monthly subscribed), `HOUR_PAID` (pay-as-you-go). This parameter is required when `InstanceId` is empty.
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// Data replication mode. Valid values: `0` (async), 1 (semi-sync), `2` (strong sync). Default value: `0`.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Instance isolation types Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.  Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// The number of the instance. Valid values: `1` (for read-only and basic instances), `2` (for other source instances). To query the price of a three-node instance, set this value to `3`.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// CPU core count of the price-queried instance. To ensure that the CPU value to be passed in is valid, use the [DescribeDBZoneConfig](https://www.tencentcloud.com/document/product/236/17229) API to query the number of purchasable cores. If this value is not specified, a default value based on memory size will be set.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance ID for querying renewal price. To query the renewal price of the instance, pass in the values of `InstanceId` and `Period`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tiered pay-as-you-go pricing, which is valid only when `PayType` is set to `HOUR_PAID`. Valid values: `1`, `2`, `3`. For more information on tiered duration, visit https://intl.cloud.tencent.com/document/product/236/18335.?from_cn_redirect=1
	Ladder *uint64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`


	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`
}

type DescribeDBPriceRequest struct {
	*tchttp.BaseRequest
	
	// Instance validity period in months. Value range: 1-36. This field is invalid when querying the prices of pay-as-you-go instances.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// AZ information in the format of "ap-guangzhou-3". You can use the <a href="https://intl.cloud.tencent.com/document/api/236/17229?from_cn_redirect=1">DescribeDBZoneConfig</a> API to query the configurable values. This parameter is required when `InstanceId` is empty.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of instances. Value range: 1-100. Default value: 1. This parameter is required when `InstanceId` is empty.
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// Instance memory size in MB. This parameter is required when `InstanceId` is empty.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB. This parameter is required when `InstanceId` is empty.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Instance type. Valid values: `master` (source instance), `dr` (disaster recovery instance), `ro` (read-only instance). Default value: `master`. This parameter is required when `InstanceId` is empty.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// Billing mode. Valid values: `PRE_PAID` (monthly subscribed), `HOUR_PAID` (pay-as-you-go). This parameter is required when `InstanceId` is empty.
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// Data replication mode. Valid values: `0` (async), 1 (semi-sync), `2` (strong sync). Default value: `0`.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Instance isolation types Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). Default value: `UNIVERSAL`.  Default value: `UNIVERSAL`.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// The number of the instance. Valid values: `1` (for read-only and basic instances), `2` (for other source instances). To query the price of a three-node instance, set this value to `3`.
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// CPU core count of the price-queried instance. To ensure that the CPU value to be passed in is valid, use the [DescribeDBZoneConfig](https://www.tencentcloud.com/document/product/236/17229) API to query the number of purchasable cores. If this value is not specified, a default value based on memory size will be set.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Instance ID for querying renewal price. To query the renewal price of the instance, pass in the values of `InstanceId` and `Period`.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tiered pay-as-you-go pricing, which is valid only when `PayType` is set to `HOUR_PAID`. Valid values: `1`, `2`, `3`. For more information on tiered duration, visit https://intl.cloud.tencent.com/document/product/236/18335.?from_cn_redirect=1
	Ladder *uint64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

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
	// Instance price. If `Currency` is set to `CNY`, the unit will be 0.01 CNY. If `Currency` is set to `USD`, the unit will be US Cent.
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Original price of the instance. If `Currency` is set to `CNY`, the unit will be 0.01 CNY. If `Currency` is set to `USD`, the unit will be US Cent.
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Currency: `CNY`, `USD`.
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
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// This parameter takes effect only when the ID of a read-only instance is passed in. If the parameter is set to `False` or left empty, the security groups bound with the RO groups of the read-only instance can only be queried. If it is set to `True`, the security groups can be modified.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
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

	// Number of entries per page. Value range: 1-2,000. Default value: 50.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBSwitchRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-2,000. Default value: 50.
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
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDataBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
	// TencentDB product type to be queried. Currently, only `mysql` is supported.
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

	// Number of results to be returned for a single request. Value range: 1-100. Maximum value: 20.
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

	// Number of results to be returned for a single request. Value range: 1-100. Maximum value: 20.
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
	// Engine version. Currently, the supported versions are `5.1`, `5.5`, `5.6`, `5.7`, and `8.0`.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine. Default value: `InnoDB`.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type DescribeDefaultParamsRequest struct {
	*tchttp.BaseRequest
	
	// Engine version. Currently, the supported versions are `5.1`, `5.5`, `5.6`, `5.7`, and `8.0`.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Type of the default parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine. Default value: `InnoDB`.
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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start timestamp, such as 1585142640.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, such as 1585142640.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// List of keywords to match. Up to 15 keywords are supported.
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
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start timestamp, such as 1585142640.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, such as 1585142640.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// List of keywords to match. Up to 15 keywords are supported.
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

	// Returned result.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
type DescribeInstanceParamRecordsRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: 20.
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
	// Parameter template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID.
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

	// Database engine version specified in the parameter template
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Number of parameters in the parameter template
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter details
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// Parameter template description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Type of the parameter template. Valid values: `HIGH_STABILITY` (high-stability template), `HIGH_PERFORMANCE` (high-performance template).
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine.  Valid values: `InnoDB`, `RocksDB`. 
	// Note:  This field may return null, indicating that no valid values can be obtained.
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
	// Engine version. If it is left empty, all parameter templates will be queried.
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Engine type. If it is left empty, all engine types will be queried.
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// Template name. If it is left empty, all template names will be queried.
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID. If it is left empty, all template IDs will be queried.
	TemplateIds []*int64 `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Engine version. If it is left empty, all parameter templates will be queried.
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// Engine type. If it is left empty, all engine types will be queried.
	EngineTypes []*string `json:"EngineTypes,omitnil,omitempty" name:"EngineTypes"`

	// Template name. If it is left empty, all template names will be queried.
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// Template ID. If it is left empty, all template IDs will be queried.
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
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Paginated query offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum entries returned per page
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeProxyCustomConfRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
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
	// Number of queried proxy configurations
	// Note: this field may return `null`, indicating that no valid value can be found.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Proxy configuration details
	// Note: this field may return `null`, indicating that no valid value can be found.
	CustomConf *CustomConfig `json:"CustomConf,omitnil,omitempty" name:"CustomConf"`

	// Weight rule
	// Note: this field may return `null`, indicating that no valid value can be found.
	WeightRule *Rule `json:"WeightRule,omitnil,omitempty" name:"WeightRule"`

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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeProxySupportParamRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
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
	// The supported maximum proxy version Note: This field may return null, indicating that no valid values can be obtained.
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// Whether to support the connection pool Note: This field may return null, indicating that no valid values can be obtained.
	SupportPool *bool `json:"SupportPool,omitnil,omitempty" name:"SupportPool"`

	// Minimum connections in the connection pool Note: This field may return null, indicating that no valid values can be obtained.
	PoolMin *uint64 `json:"PoolMin,omitnil,omitempty" name:"PoolMin"`

	// Maximum connections in the connection pool Note: This field may return null, indicating that no valid values can be obtained.
	PoolMax *uint64 `json:"PoolMax,omitnil,omitempty" name:"PoolMax"`

	// Whether to support transaction splitting Note: This field may return null, indicating that no valid values can be obtained.
	SupportTransSplit *bool `json:"SupportTransSplit,omitnil,omitempty" name:"SupportTransSplit"`

	// Minimum proxy version supporting connection pool Note: This field may return null, indicating that no valid values can be obtained.
	SupportPoolMinVersion *string `json:"SupportPoolMinVersion,omitnil,omitempty" name:"SupportPoolMinVersion"`

	// Minimum proxy version supporting transaction splitting Note: This field may return null, indicating that no valid values can be obtained.
	SupportTransSplitMinVersion *string `json:"SupportTransSplitMinVersion,omitnil,omitempty" name:"SupportTransSplitMinVersion"`

	// Whether read-only mode is supported Note: This field may return null, indicating that no valid values can be obtained.
	SupportReadOnly *bool `json:"SupportReadOnly,omitnil,omitempty" name:"SupportReadOnly"`

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
	// Instance ID in the format of `cdb-c1nl9rpv` or `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRoGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `cdb-c1nl9rpv` or `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page.
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
	// RO group information array. An instance can be associated with multiple RO groups.
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

	// Pagination parameter, i.e., the number of entries to be returned for a single request. Default value: 20. Maximum value: 100.
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

	// Pagination parameter, i.e., the number of entries to be returned for a single request. Default value: 20. Maximum value: 100.
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
	// Note: this field may return null, indicating that no valid values can be obtained.
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
type DescribeSlowLogDataRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start timestamp, such as 1585142640.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, such as 1585142640.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Client `Host` list.
	UserHosts []*string `json:"UserHosts,omitnil,omitempty" name:"UserHosts"`

	// Client username list.
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// Accessed database list.
	DataBases []*string `json:"DataBases,omitnil,omitempty" name:"DataBases"`

	// Sort by field. Valid values: Timestamp, QueryTime, LockTime, RowsExamined, RowsSent.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of results per page in paginated queries. Default value: 100. Maximum value: 400.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`
}

type DescribeSlowLogDataRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start timestamp, such as 1585142640.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, such as 1585142640.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Client `Host` list.
	UserHosts []*string `json:"UserHosts,omitnil,omitempty" name:"UserHosts"`

	// Client username list.
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// Accessed database list.
	DataBases []*string `json:"DataBases,omitnil,omitempty" name:"DataBases"`

	// Sort by field. Valid values: Timestamp, QueryTime, LockTime, RowsExamined, RowsSent.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of results per page in paginated queries. Default value: 100. Maximum value: 400.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// This parameter is valid only for source or disaster recovery instances. Valid value: `slave`, which indicates pulling logs from the replica.
	InstType *string `json:"InstType,omitnil,omitempty" name:"InstType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogDataResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Queried results.
	// Note: this field may return null, indicating that no valid values can be obtained.
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

	// Number of entries per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination offset, starting from `0`. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Value range: 1-100. Default value: 20.
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
type DescribeTablesRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database name.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Record offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
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

	// Number of results to be returned for a single request. Default value: 20. Maximum value: 2,000.
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
	// List of instances.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTagsOfInstanceIdsRequest struct {
	*tchttp.BaseRequest
	
	// List of instances.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Pagination offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page.
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
	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be unbound from the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be unbound from the read-only replicas themselves.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// This parameter takes effect only when the IDs of read-only replicas are passed in. If this parameter is set to `False` or left empty, the security group will be unbound from the RO groups of these read-only replicas. If this parameter is set to `True`, the security group will be unbound from the read-only replicas themselves.
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
	// Error occurrence time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Error details
	// Note: this field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ImportRecord struct {
	// Status value
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status value
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Execution duration
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backend task ID
	WorkId *string `json:"WorkId,omitnil,omitempty" name:"WorkId"`

	// Name of the file to be imported
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Execution progress
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// File size
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

	// Rule description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
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
	// Public network access status. Value range: 0 (not enabled), 1 (enabled), 2 (disabled)
	WanStatus *int64 `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// AZ information
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Initialization flag. Value range: 0 (not initialized), 1 (initialized)
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// Read-only VIP information. This field is available only for read-only instances with dedicated access enabled.
	RoVipInfo *RoVipInfo `json:"RoVipInfo,omitnil,omitempty" name:"RoVipInfo"`

	// Memory capacity in MB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance status. Valid values: `0` (creating), `1` (running), `4` (isolating), `5` (isolated).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// VPC ID, such as 51102
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Secondary server information.
	SlaveInfo *SlaveInfo `json:"SlaveInfo,omitnil,omitempty" name:"SlaveInfo"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Disk capacity in GB
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled), 2 (auto-renewal disabled)
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync)
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Detailed information about the read-only group.
	RoGroups []*RoGroup `json:"RoGroups,omitnil,omitempty" name:"RoGroups"`

	// Subnet ID, such as 2333
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance type. Value range: 1 (primary), 2 (disaster recovery), 3 (read-only)
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Region information
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance expiration time
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// AZ deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ)
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// Instance task status. 0 - no task; 1 - upgrading; 2 - importing data; 3 - activating secondary; 4 - enabling public network access; 5 - batch operation in progress; 6 - rolling back; 7 - disabling public network access; 8 - changing password; 9 - renaming instance; 10 - restarting; 12 - migrating self-built instance; 13 - dropping table; 14 - creating and syncing disaster recovery instance; 15 - pending upgrade and switch; 16 - upgrade and switch in progress; 17 - upgrade and switch completed
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Detailed information about the primary instance.
	MasterInfo *MasterInfo `json:"MasterInfo,omitnil,omitempty" name:"MasterInfo"`

	// Instance type
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Kernel version
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Detailed information about the disaster recovery instance.
	DrInfo []*DrInfo `json:"DrInfo,omitnil,omitempty" name:"DrInfo"`

	// Public domain name
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// Public network port number
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// Billing type
	PayType *int64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Instance IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port number
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Whether the disk write is locked (It depends on whether the instance data in disk exceeds its quota). Valid values: `0` (unlocked), `1` (locked).
	CdbError *int64 `json:"CdbError,omitnil,omitempty" name:"CdbError"`

	// VPC descriptor, such as "vpc-5v8wn9mg"
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet descriptor, such as "subnet-1typ0s7d"
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Physical ID
	PhysicalId *string `json:"PhysicalId,omitnil,omitempty" name:"PhysicalId"`

	// Number of cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Queries per second
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Physical server model.
	DeviceClass *string `json:"DeviceClass,omitnil,omitempty" name:"DeviceClass"`

	// Placement group ID.
	DeployGroupId *string `json:"DeployGroupId,omitnil,omitempty" name:"DeployGroupId"`

	// AZ ID.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Number of nodes
	InstanceNodes *int64 `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// Tag list.
	TagList []*TagInfoItem `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Engine type.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// Maximum delay threshold.
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// Instance disk type. Valid values are returned only for Cluster Edition and single-node (cloud disk) instances.
	// Note:
	// 1. If "DiskType": "CLOUD_HSSD" is returned, it indicates that the instance disk type is Enhanced SSD.
	// 2. If "DiskType": "CLOUD_SSD" is returned, it indicates that the instance disk type is Cloud SSD.
	// 3. If "DiskType": "" is returned and the DeviceType parameter value is UNIVERSAL or EXCLUSIVE, it indicates that the instance uses a local SSD.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Current number of CPU cores for scale-out.
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Cluster Edition instance node information.
	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// Analysis engine node list.
	AnalysisNodeInfos []*AnalysisNodeInfo `json:"AnalysisNodeInfos,omitnil,omitempty" name:"AnalysisNodeInfos"`

	// Device bandwidth, in GB. This parameter is valid when DeviceClass is specified. For example, 25 means the current device bandwidth is 25 GB; 10 means the current device bandwidth is 10 GB.
	DeviceBandwidth *uint64 `json:"DeviceBandwidth,omitnil,omitempty" name:"DeviceBandwidth"`

	// Instance termination protection status. on indicates enabled; otherwise, the protection is disabled.
	DestroyProtect *string `json:"DestroyProtect,omitnil,omitempty" name:"DestroyProtect"`
}

type InstanceRebootTime struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Estimated restart time
	TimeInSeconds *int64 `json:"TimeInSeconds,omitnil,omitempty" name:"TimeInSeconds"`
}

type InstanceRollbackRangeTime struct {
	// Queries database error code
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
	// Async task request ID, which can be used to query the execution result of an async task. (This returned field has been disused. You can query the isolation status of an instance through the `DescribeDBInstances` API.)
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Retention period of local binlog. Value range: [72,168].
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type LocalBinlogConfigDefault struct {
	// Retention period of local binlog. Value range: [72,168].
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type LogRuleTemplateInfo struct {
	// Template ID. 
	// Note: The return value may be null, indicating that no valid data can be obtained.
	RuleTemplateId *string `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// Template name.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	RuleTemplateName *string `json:"RuleTemplateName,omitnil,omitempty" name:"RuleTemplateName"`

	// Alarm level. Valid values: 1: Low risk; 2: Medium risk; 3: High risk. 
	// Note: The return value may be null, indicating that no valid data can be obtained.
	AlarmLevel *string `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// Template change status. Valid values: 0: Unchanged; 1: Changed; 2: Deleted.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	RuleTemplateStatus *int64 `json:"RuleTemplateStatus,omitnil,omitempty" name:"RuleTemplateStatus"`
}

type LogToCLSConfig struct {
	// Enabling status of the feature.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CLS log set ID.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// Log topic ID.
	// Note: The return value may be null, indicating that no valid data can be obtained.
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`
}

type MasterInfo struct {
	// Region information
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// AZ information
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Long instance ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Task status
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Memory capacity
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk capacity
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Instance model
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// Queries per second
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// VPC ID
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Dedicated cluster ID
	ExClusterId *string `json:"ExClusterId,omitnil,omitempty" name:"ExClusterId"`

	// Dedicated cluster name
	ExClusterName *string `json:"ExClusterName,omitnil,omitempty" name:"ExClusterName"`
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Database account remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Database account remarks
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
	// List of TencentDB accounts
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Maximum connections of the account. Maximum value: `10240`.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type ModifyAccountMaxUserConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// List of TencentDB accounts
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

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
}

type ModifyAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// New password of the database account. It can only contain 8-64 characters and must contain at least two of the following types of characters: letters, digits, and special characters (_+-&=!@#$%^*()).
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`

	// TencentDB account
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`
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

	// Database account, including username and domain name.
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

	// Database account, including username and domain name.
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
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API.
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
	
	// Audit rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API.
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
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
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

	// Rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained through the [DescribeDBInstances](https://www.tencentcloud.comom/document/product/236/15872?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log retention period. Valid values:7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// High-frequency log retention period. Default value: 7. This value must be less than or equal to LogExpireDay. Valid values include:3 - 3 days;7 - One week;30 - One month;90 - Three months;180 - Six months;365 - One year;1095 - Three years;1825 - Five years.
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Modifies the instance audit rule to Full audit.
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// Deprecated.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID, which can be obtained through the [DescribeAuditRuleTemplates](https://www.tencentcloud.comom/document/api/236/101811?from_cn_redirect=1) API.
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
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Auto-renewal flag. Value range: 0 (auto-renewal not enabled), 1 (auto-renewal enabled).
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type ModifyAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
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
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup file retention period in days. Value range: 7-1830.
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// (This parameter will be disused. The `BackupTimeWindow` parameter is recommended.) Backup time range in the format of 02:00-06:00, with the start time and end time on the hour. Valid values: 00:00-12:00, 02:00-06:00, 06:00-10:00, 10:00-14:00, 14:00-18:00, 18:00-22:00, 22:00-02:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Automatic backup mode. Only `physical` (physical cold backup) is supported
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Binlog retention period in days. Value range: 7-1830. It can’t be greater than the retention period of backup files.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// Backup time window; for example, to set up backup between 10:00 and 14:00 on every Tuesday and Sunday, you should set this parameter as follows: {"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"} (Note: You can set up backup on different days, but the backup time windows need to be the same. If this field is set, the `StartTime` field will be ignored)
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// Switch for periodic archive. Valid values: `off` (disable), `on` (enable). Default value:`off`. When you enable the periodic archive policy for the first time, you need to enter the `BackupPeriodSaveDays`, `BackupPeriodSaveInterval`, `BackupPeriodSaveCount`, and `StartBackupPeriodSaveDate` parameters; otherwise, the policy will not take effect.
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

	// Whether to enable the archive backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// The period (in days) of how long a data backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// The period (in days) of how long a log backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// Whether to enable the archive backup of the log. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// Whether to enable the standard storage policy for data backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// The period (in days) of how long a data backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// Whether to enable the standard storage policy for log backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBinlogStandby *string `json:"EnableBinlogStandby,omitnil,omitempty" name:"EnableBinlogStandby"`

	// The period (in days) of how long a log backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BinlogStandbyDays *int64 `json:"BinlogStandbyDays,omitnil,omitempty" name:"BinlogStandbyDays"`
}

type ModifyBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup file retention period in days. Value range: 7-1830.
	ExpireDays *int64 `json:"ExpireDays,omitnil,omitempty" name:"ExpireDays"`

	// (This parameter will be disused. The `BackupTimeWindow` parameter is recommended.) Backup time range in the format of 02:00-06:00, with the start time and end time on the hour. Valid values: 00:00-12:00, 02:00-06:00, 06:00-10:00, 10:00-14:00, 14:00-18:00, 18:00-22:00, 22:00-02:00.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Automatic backup mode. Only `physical` (physical cold backup) is supported
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Binlog retention period in days. Value range: 7-1830. It can’t be greater than the retention period of backup files.
	BinlogExpireDays *int64 `json:"BinlogExpireDays,omitnil,omitempty" name:"BinlogExpireDays"`

	// Backup time window; for example, to set up backup between 10:00 and 14:00 on every Tuesday and Sunday, you should set this parameter as follows: {"Monday": "", "Tuesday": "10:00-14:00", "Wednesday": "", "Thursday": "", "Friday": "", "Saturday": "", "Sunday": "10:00-14:00"} (Note: You can set up backup on different days, but the backup time windows need to be the same. If this field is set, the `StartTime` field will be ignored)
	BackupTimeWindow *CommonTimeWindow `json:"BackupTimeWindow,omitnil,omitempty" name:"BackupTimeWindow"`

	// Switch for periodic archive. Valid values: `off` (disable), `on` (enable). Default value:`off`. When you enable the periodic archive policy for the first time, you need to enter the `BackupPeriodSaveDays`, `BackupPeriodSaveInterval`, `BackupPeriodSaveCount`, and `StartBackupPeriodSaveDate` parameters; otherwise, the policy will not take effect.
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

	// Whether to enable the archive backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBackupArchive *string `json:"EnableBackupArchive,omitnil,omitempty" name:"EnableBackupArchive"`

	// The period (in days) of how long a data backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BackupArchiveDays *int64 `json:"BackupArchiveDays,omitnil,omitempty" name:"BackupArchiveDays"`

	// The period (in days) of how long a log backup is retained before being archived, which falls between 180 days and the number of days from the time it is created until it expires.
	BinlogArchiveDays *int64 `json:"BinlogArchiveDays,omitnil,omitempty" name:"BinlogArchiveDays"`

	// Whether to enable the archive backup of the log. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBinlogArchive *string `json:"EnableBinlogArchive,omitnil,omitempty" name:"EnableBinlogArchive"`

	// Whether to enable the standard storage policy for data backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
	EnableBackupStandby *string `json:"EnableBackupStandby,omitnil,omitempty" name:"EnableBackupStandby"`

	// The period (in days) of how long a data backup is retained before switching to standard storage, which falls between 30 days and the number of days from the time it is created until it expires. If the archive backup is enabled, this period cannot be greater than archive backup period.
	BackupStandbyDays *int64 `json:"BackupStandbyDays,omitnil,omitempty" name:"BackupStandbyDays"`

	// Whether to enable the standard storage policy for log backup. Valid values: `off` (disable), `on` (enable). Default value: `off`.
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
}

type ModifyBackupEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-XXXX, which is the same as that displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Default encryption status for the new auto-generated physical backup files. Valid values: `on`, `off`.
	EncryptionStatus *string `json:"EncryptionStatus,omitnil,omitempty" name:"EncryptionStatus"`
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
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Address ID of the proxy group
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`

	// Description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type ModifyCdbProxyAddressDescRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Address ID of the proxy group
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
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

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

	// Valid Hours of Old IP
	ReleaseDuration *uint64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
}

type ModifyCdbProxyAddressVipAndVPortRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

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

	// Valid Hours of Old IP
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Connection pool threshold
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`
}

type ModifyCdbProxyParamRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Connection pool threshold
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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Valid values: error and slowLog.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Enabling status. Valid values: ON and OFF.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Indicates whether a log set needs to be created.
	CreateLogset *bool `json:"CreateLogset,omitnil,omitempty" name:"CreateLogset"`

	// Log set name if the log set is to be created or ID of the selected existing log set.
	Logset *string `json:"Logset,omitnil,omitempty" name:"Logset"`

	// Indicates whether a log topic needs to be created.
	CreateLogTopic *bool `json:"CreateLogTopic,omitnil,omitempty" name:"CreateLogTopic"`

	// Log topic name if the topic is to be created or ID of the selected existing topic.
	LogTopic *string `json:"LogTopic,omitnil,omitempty" name:"LogTopic"`

	// Log topic validity period, which is 30 days by default if not specified.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Indicates whether to create an index when creating the log topic.
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`
}

type ModifyDBInstanceLogToCLSRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Log type. Valid values: error and slowLog.
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// Enabling status. Valid values: ON and OFF.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Indicates whether a log set needs to be created.
	CreateLogset *bool `json:"CreateLogset,omitnil,omitempty" name:"CreateLogset"`

	// Log set name if the log set is to be created or ID of the selected existing log set.
	Logset *string `json:"Logset,omitnil,omitempty" name:"Logset"`

	// Indicates whether a log topic needs to be created.
	CreateLogTopic *bool `json:"CreateLogTopic,omitnil,omitempty" name:"CreateLogTopic"`

	// Log topic name if the topic is to be created or ID of the selected existing topic.
	LogTopic *string `json:"LogTopic,omitnil,omitempty" name:"LogTopic"`

	// Log topic validity period, which is 30 days by default if not specified.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Indicates whether to create an index when creating the log topic.
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`
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
type ModifyDBInstanceNameRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The modified instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The modified instance name.
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
	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Project ID.
	NewProjectId *int64 `json:"NewProjectId,omitnil,omitempty" name:"NewProjectId"`
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance IDs in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Project ID.
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

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// This parameter takes effect only when the ID of read-only replica is passed in. If this parameter is set to `False` or left empty, the security groups bound with the RO group of the read-only replicas will be modified. If this parameter is set to `True`, the security groups bound with the read-only replica itself will be modified.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// This parameter takes effect only when the ID of read-only replica is passed in. If this parameter is set to `False` or left empty, the security groups bound with the RO group of the read-only replicas will be modified. If this parameter is set to `True`, the security groups bound with the read-only replica itself will be modified.
	ForReadonlyInstance *bool `json:"ForReadonlyInstance,omitnil,omitempty" name:"ForReadonlyInstance"`
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

	// Target IP. Either this parameter or `DstPort` must be passed in.
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// Target port number. Value range: 1024-65535. Either this parameter or `DstIp` must be passed in.
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Unified subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Repossession duration in hours for old IP in the original network when changing from classic network to VPC or changing the VPC subnet. Value range: 0–168. Default value: `24`.
	ReleaseDuration *int64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
}

type ModifyDBInstanceVipVportRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv, cdbro-c2nl9rpv, or cdbrg-c3nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://www.tencentcloud.com/document/product/236/15872) API to query the ID, which is the value of the `InstanceId` output parameter.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Target IP. Either this parameter or `DstPort` must be passed in.
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// Target port number. Value range: 1024-65535. Either this parameter or `DstIp` must be passed in.
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Unified subnet ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Repossession duration in hours for old IP in the original network when changing from classic network to VPC or changing the VPC subnet. Value range: 0–168. Default value: `24`.
	ReleaseDuration *int64 `json:"ReleaseDuration,omitnil,omitempty" name:"ReleaseDuration"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceVipVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceVipVportResponseParams struct {
	// Async task ID. This parameter is deprecated.
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	// List of short instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value).
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Template ID. At least one of `ParamList` and `TemplateId` must be passed in.
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
	
	// List of short instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value).
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// Template ID. At least one of `ParamList` and `TemplateId` must be passed in.
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
	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyInstancePasswordComplexityRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tag to be added or modified.
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitnil,omitempty" name:"ReplaceTags"`

	// Tag to be deleted.
	DeleteTags []*TagInfo `json:"DeleteTags,omitnil,omitempty" name:"DeleteTags"`
}

type ModifyInstanceTagRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Tag to be added or modified.
	ReplaceTags []*TagInfo `json:"ReplaceTags,omitnil,omitempty" name:"ReplaceTags"`

	// Tag to be deleted.
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
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Space utilization of local binlog. Value range: [30,50].
	MaxUsage *int64 `json:"MaxUsage,omitnil,omitempty" name:"MaxUsage"`
}

type ModifyLocalBinlogConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
	SaveHours *int64 `json:"SaveHours,omitnil,omitempty" name:"SaveHours"`

	// Space utilization of local binlog. Value range: [30,50].
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
	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name (up to 64 characters)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Template description (up to 255 characters)
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// List of parameters.
	ParamList []*Parameter `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name (up to 64 characters)
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
	// RO group ID.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// RO group details.
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitnil,omitempty" name:"RoGroupInfo"`

	// Weights of instances in RO group. If the weighting mode of an RO group is changed to custom mode, this parameter must be set, and a weight value needs to be set for each RO instance.
	RoWeightValues []*RoWeightValue `json:"RoWeightValues,omitnil,omitempty" name:"RoWeightValues"`

	// Whether to rebalance the loads of read-only replicas in the RO group. Valid values: `1` (yes), `0` (no). Default value: `0`. If this parameter is set to `1`, connections to the read-only replicas in the RO group will be interrupted transiently. Please ensure that your application has a reconnection mechanism.
	IsBalanceRoLoad *int64 `json:"IsBalanceRoLoad,omitnil,omitempty" name:"IsBalanceRoLoad"`

	// This field has been deprecated.
	ReplicationDelayTime *int64 `json:"ReplicationDelayTime,omitnil,omitempty" name:"ReplicationDelayTime"`
}

type ModifyRoGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// RO group ID.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// RO group details.
	RoGroupInfo *RoGroupAttr `json:"RoGroupInfo,omitnil,omitempty" name:"RoGroupInfo"`

	// Weights of instances in RO group. If the weighting mode of an RO group is changed to custom mode, this parameter must be set, and a weight value needs to be set for each RO instance.
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
	// Async task ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
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
type ModifyTimeWindowRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
	TimeRanges []*string `json:"TimeRanges,omitnil,omitempty" name:"TimeRanges"`

	// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
	Weekdays []*string `json:"Weekdays,omitnil,omitempty" name:"Weekdays"`

	// Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
	MaxDelayTime *uint64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type ModifyTimeWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
	TimeRanges []*string `json:"TimeRanges,omitnil,omitempty" name:"TimeRanges"`

	// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
	Weekdays []*string `json:"Weekdays,omitnil,omitempty" name:"Weekdays"`

	// Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
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
	// TencentDB for MySQL instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period of the audit log. Valid values:  `7` (one week), `30` (one month), `90` (three months), `180` (six months), `365` (one year), `1095` (three years), `1825` (five years).
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Retention period of high-frequency audit logs. Valid values:  `7` (one week), `30` (one month).
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule If both this parameter and `RuleTemplateIds` are left empty, full audit will be applied.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and AuditRuleFilters are not specified, all SQL statements will be recorded.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit type. Valid values: true: Record all; false: Record by rules (default value).
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// TencentDB for MySQL instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Retention period of the audit log. Valid values:  `7` (one week), `30` (one month), `90` (three months), `180` (six months), `365` (one year), `1095` (three years), `1825` (five years).
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// Retention period of high-frequency audit logs. Valid values:  `7` (one week), `30` (one month).
	HighLogExpireDay *uint64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// Audit rule If both this parameter and `RuleTemplateIds` are left empty, full audit will be applied.
	AuditRuleFilters []*AuditRuleFilters `json:"AuditRuleFilters,omitnil,omitempty" name:"AuditRuleFilters"`

	// Rule template ID. If both this parameter and AuditRuleFilters are not specified, all SQL statements will be recorded.
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// Audit type. Valid values: true: Record all; false: Record by rules (default value).
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
	// TencentDB instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`
}

type OpenDBInstanceEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// TencentDB instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
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
type OpenWanServiceRequestParams struct {
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenWanServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// Parameter template ID
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Instance engine version
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Parameter template type
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// Parameter template engine Note: This field may return null, indicating that no valid values can be obtained.
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

	// Whether the parameter can be modified Note: This field may return null, indicating that no valid values can be obtained.
	IsNotSupportEdit *bool `json:"IsNotSupportEdit,omitnil,omitempty" name:"IsNotSupportEdit"`
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

	// Assignment mode of weights. Valid values: `system` (auto-assigned), `custom`. Note: This field may return null, indicating that no valid values can be obtained.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Whether to remove delayed read-only instances from the proxy group Valid values: `true`, `false`. Note: This field may return null, indicating that no valid values can be obtained.
	IsKickOut *bool `json:"IsKickOut,omitnil,omitempty" name:"IsKickOut"`

	// Least read-only instances. Minimum value:  `0`. Note: This field may return null, indicating that no valid values can be obtained.
	MinCount *uint64 `json:"MinCount,omitnil,omitempty" name:"MinCount"`

	// The delay threshold. Minimum value:  `0`. Note: This field may return null, indicating that no valid values can be obtained.
	MaxDelay *uint64 `json:"MaxDelay,omitnil,omitempty" name:"MaxDelay"`

	// Whether to automatically add newly created read-only instances. Valid values: `true`, `false`. Note: This field may return null, indicating that no valid values can be obtained.
	AutoAddRo *bool `json:"AutoAddRo,omitnil,omitempty" name:"AutoAddRo"`

	// Whether it is read-only. Valid values: `true`, `false`. Note: This field may return null, indicating that no valid values can be obtained.
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// Whether to enable transaction splitting Note: This field may return null, indicating that no valid values can be obtained.
	TransSplit *bool `json:"TransSplit,omitnil,omitempty" name:"TransSplit"`

	// Whether to enable failover Note: This field may return null, indicating that no valid values can be obtained.
	FailOver *bool `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Whether to enable the connection pool Note: This field may return null, indicating that no valid values can be obtained.
	ConnectionPool *bool `json:"ConnectionPool,omitnil,omitempty" name:"ConnectionPool"`

	// Note:  This field may return null, indicating that no valid values can be obtained.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Read weight assignment for an instance Note: This field may return null, indicating that no valid values can be obtained.
	ProxyAllocation []*ProxyAllocation `json:"ProxyAllocation,omitnil,omitempty" name:"ProxyAllocation"`
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

	// Proxy version Note: This field may return null, indicating that no valid values can be obtained.
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// Supported proxy upgrade version Note: This field may return null, indicating that no valid values can be obtained.
	SupportUpgradeProxyVersion *string `json:"SupportUpgradeProxyVersion,omitnil,omitempty" name:"SupportUpgradeProxyVersion"`

	// Proxy status Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Proxy task status Note: This field may return null, indicating that no valid values can be obtained.
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Node information of the proxy group Note: This field may return null, indicating that no valid values can be obtained.
	ProxyNode []*ProxyNode `json:"ProxyNode,omitnil,omitempty" name:"ProxyNode"`

	// Address information of the proxy group Note: This field may return null, indicating that no valid values can be obtained.
	ProxyAddress []*ProxyAddress `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`

	// Connection pool threshold Note: This field may return null, indicating that no valid values can be obtained.
	ConnectionPoolLimit *uint64 `json:"ConnectionPoolLimit,omitnil,omitempty" name:"ConnectionPoolLimit"`

	// Whether to support address creation Note: This field may return null, indicating that no valid values can be obtained.
	SupportCreateProxyAddress *bool `json:"SupportCreateProxyAddress,omitnil,omitempty" name:"SupportCreateProxyAddress"`

	// TencentDB versions supporting proxy versions upgrade Note: This field may return null, indicating that no valid values can be obtained.
	SupportUpgradeProxyMysqlVersion *string `json:"SupportUpgradeProxyMysqlVersion,omitnil,omitempty" name:"SupportUpgradeProxyMysqlVersion"`
}

type ProxyInst struct {
	// Instance ID Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type. Valid values:  `master` (source instance), `ro` (read-only instance), `dr` (disaster recovery instance), `sdr` (disaster recovery instance of small specifications). Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance status. Valid values:  `0` (creating), `1` (running), `4` (isolating), `5` (isolated). Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Read weight. If it is assigned by the system automatically, the modification will not take effect but represents whether the instance is enabled. Note: This field may return null, indicating that no valid values can be obtained.
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Instance region Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance AZ Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ProxyNode struct {
	// Proxy node ID Note: This field may return null, indicating that no valid values can be obtained.
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`

	// Number of CPU cores Note: This field may return null, indicating that no valid values can be obtained.
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size Note: This field may return null, indicating that no valid values can be obtained.
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Node status Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Proxy node AZ Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Proxy node region Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Connections Note: This field may return null, indicating that no valid values can be obtained.
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

// Predefined struct for user
type ReleaseIsolatedDBInstancesRequestParams struct {
	// Array of instance IDs in the format of `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID, whose value is the `InstanceId` value in the output parameters.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ReleaseIsolatedDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance IDs in the format of `cdb-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID, whose value is the `InstanceId` value in the output parameters.
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
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Address ID of the proxy group
	ProxyAddressId *string `json:"ProxyAddressId,omitnil,omitempty" name:"ProxyAddressId"`
}

type ReloadBalanceProxyNodeRequest struct {
	*tchttp.BaseRequest
	
	// Proxy group ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`

	// Address ID of the proxy group
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

	// To renew a pay-as-you-go instance to a monthly subscribed one, you need to set this parameter to `PREPAID`.
	ModifyPayType *string `json:"ModifyPayType,omitnil,omitempty" name:"ModifyPayType"`
}

type RenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be renewed in the format of cdb-c1nl9rpv, which is the same as the instance ID displayed in the TencentDB console. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Renewal period in months. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// To renew a pay-as-you-go instance to a monthly subscribed one, you need to set this parameter to `PREPAID`.
	ModifyPayType *string `json:"ModifyPayType,omitnil,omitempty" name:"ModifyPayType"`
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
type ResetRootAccountRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ResetRootAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
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
	// Read-only group mode. Valid values: `alone` (the system assigns a read-only group automatically), `allinone` (a new read-only group will be created), `join` (an existing read-only group will be used).
	RoGroupMode *string `json:"RoGroupMode,omitnil,omitempty" name:"RoGroupMode"`

	// Read-only group ID.
	// Note: If the data structure is used during instance purchase, this item is required only when the read-only group mode is set to join.
	RoGroupId *string `json:"RoGroupId,omitnil,omitempty" name:"RoGroupId"`

	// Read-only group name.
	RoGroupName *string `json:"RoGroupName,omitnil,omitempty" name:"RoGroupName"`

	// Whether to enable the function of isolating an instance that exceeds the latency threshold. If it is enabled, when the latency between the read-only instance and the primary instance exceeds the latency threshold, the read-only instance will be isolated. Valid values: 1 (enabled), 0 (not enabled)
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitnil,omitempty" name:"RoOfflineDelay"`

	// Delay threshold, in seconds. Value range: 1–10000. The value is an integer.
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitnil,omitempty" name:"RoMaxDelayTime"`

	// Minimum number of instances to be retained. If the number of the purchased read-only instances is smaller than the set value, they will not be removed.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitnil,omitempty" name:"MinRoInGroup"`

	// Read/write weight distribution mode. Valid values: `system` (weights are assigned by the system automatically), `custom` (weights are customized)
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// This field has been disused. To view the weight of a read-only instance, check the `Weight` value in the `RoInstances` field.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Details of read-only instances in read-only group
	RoInstances []*RoInstanceInfo `json:"RoInstances,omitnil,omitempty" name:"RoInstances"`

	// Private IP of read-only group.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port number of read-only group.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Virtual Private Cloud (VPC) ID.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Subnet ID.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// Region of the read-only group.
	RoGroupRegion *string `json:"RoGroupRegion,omitnil,omitempty" name:"RoGroupRegion"`

	// AZ of the read-only group.
	RoGroupZone *string `json:"RoGroupZone,omitnil,omitempty" name:"RoGroupZone"`

	// Replication delay time, in seconds. Value range: 1–259200. The value is an integer.
	DelayReplicationTime *int64 `json:"DelayReplicationTime,omitnil,omitempty" name:"DelayReplicationTime"`
}

type RoGroupAttr struct {
	// RO group name.
	RoGroupName *string `json:"RoGroupName,omitnil,omitempty" name:"RoGroupName"`

	// Maximum delay threshold for RO instances in seconds. Minimum value: 1. Please note that this value will take effect only if an instance removal policy is enabled in the RO group.
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitnil,omitempty" name:"RoMaxDelayTime"`

	// Whether to enable instance removal. Valid values: 1 (enabled), 0 (not enabled). Please note that if instance removal is enabled, the delay threshold parameter (`RoMaxDelayTime`) must be set.
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitnil,omitempty" name:"RoOfflineDelay"`

	// Minimum number of instances to be retained, which can be set to any value less than or equal to the number of RO instances in the RO group. Please note that if this value is set to be greater than the number of RO instances, no removal will be performed, and if it is set to 0, all instances with an excessive delay will be removed.
	MinRoInGroup *int64 `json:"MinRoInGroup,omitnil,omitempty" name:"MinRoInGroup"`

	// Weighting mode. Supported values include `system` (automatically assigned by the system) and `custom` (defined by user). Please note that if the `custom` mode is selected, the RO instance weight configuration parameter (RoWeightValues) must be set.
	WeightMode *string `json:"WeightMode,omitnil,omitempty" name:"WeightMode"`

	// Replication delay.
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

	// Name of RO AZ, such as ap-shanghai-2
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
	// Note: this field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// New database name after rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	NewDatabaseName *string `json:"NewDatabaseName,omitnil,omitempty" name:"NewDatabaseName"`
}

type RollbackInstancesInfo struct {
	// TencentDB instance ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Rollback policy. Valid values: `table` (ultrafast mode), `db` (faster mode), and `full` (fast mode). Default value: `full`. In the ultrafast mode, only backups and binlogs of the tables specified by the `Tables` parameter are imported; if `Tables` does not include all of the tables involved in cross-table operations, the rollback may fail; and the `Database` parameter must be left empty. In the faster mode, only backups and binlogs of the databases specified by the `Databases` parameter are imported, and if `Databases` does not include all of the databases involved in cross-database operations, the rollback may fail. In the fast mode, backups and binlogs of the entire instance will be imported in a speed slower than the other modes.
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Database rollback time in the format of yyyy-mm-dd hh:mm:ss
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// Information of the databases to be rolled back, which means rollback at the database level
	// Note: this field may return null, indicating that no valid values can be obtained.
	Databases []*RollbackDBName `json:"Databases,omitnil,omitempty" name:"Databases"`

	// Information of the tables to be rolled back, which means rollback at the table level
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tables []*RollbackTables `json:"Tables,omitnil,omitempty" name:"Tables"`
}

type RollbackTableName struct {
	// Original table name before rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// New table name after rollback
	// Note: this field may return null, indicating that no valid values can be obtained.
	NewTableName *string `json:"NewTableName,omitnil,omitempty" name:"NewTableName"`
}

type RollbackTables struct {
	// Database name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table details
	// Note: this field may return null, indicating that no valid values can be obtained.
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

	// Rollback task details.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Detail []*RollbackInstancesInfo `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type RollbackTimeRange struct {
	// Start time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-10-29 01:06:04
	Begin *string `json:"Begin,omitnil,omitempty" name:"Begin"`

	// End time available for rollback in the format of yyyy-MM-dd HH:mm:ss, such as 2016-11-02 11:44:47
	End *string `json:"End,omitnil,omitempty" name:"End"`
}

type Rule struct {
	// The maximum weight
	// Note: this field may return `null`, indicating that no valid value can be found.
	LessThan *uint64 `json:"LessThan,omitnil,omitempty" name:"LessThan"`

	// Weight
	// Note: this field may return `null`, indicating that no valid value can be found.
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

	// AZ name of the secondary database, such as ap-shanghai-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type SlaveInfo struct {
	// Information of secondary server 1
	First *SlaveInstanceInfo `json:"First,omitnil,omitempty" name:"First"`

	// Second secondary server information.
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

	// Backup snapshot time in the format of yyyy-MM-dd HH:mm:ss, such as 2016-03-17 02:10:37
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// Download address on the private network
	IntranetUrl *string `json:"IntranetUrl,omitnil,omitempty" name:"IntranetUrl"`

	// Download address on the public network
	InternetUrl *string `json:"InternetUrl,omitnil,omitempty" name:"InternetUrl"`

	// Log type. Value range: slowlog (slow log)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SlowLogItem struct {
	// SQL execution time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// SQL execution duration in seconds.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// SQL statement.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// Client address.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// Username.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Lock duration in seconds.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// Number of scanned rows.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// Number of rows in result set.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// SQL template.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// SQL statement MD5.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Scale-out mode. Valid values: auto and
	// manual.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of CPU cores to increase during manual scale-out. This parameter is required when Type is set to manual.
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Automatic scale-out policy. This parameter is required when Type is set to auto.
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`
}

type StartCpuExpandRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Scale-out mode. Valid values: auto and
	// manual.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of CPU cores to increase during manual scale-out. This parameter is required when Type is set to manual.
	ExpandCpu *int64 `json:"ExpandCpu,omitnil,omitempty" name:"ExpandCpu"`

	// Automatic scale-out policy. This parameter is required when Type is set to auto.
	AutoStrategy *AutoStrategy `json:"AutoStrategy,omitnil,omitempty" name:"AutoStrategy"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCpuExpandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCpuExpandResponseParams struct {
	// Async task ID, which can be passed in by calling the `DescribeAsyncRequest` API for task progress query.
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
	// Read-Only instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StartReplicationRequest struct {
	*tchttp.BaseRequest
	
	// Read-Only instance ID.
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
	// Async task ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
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
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopCpuExpandRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
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
	// Async task ID, which can be passed in by calling the `DescribeAsyncRequest` API for task progress query.
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
	// Read-Only instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopReplicationRequest struct {
	*tchttp.BaseRequest
	
	// Read-Only instance ID.
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
	// Async task ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
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
	// ID of the instance whose rollback task is canceled
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StopRollbackRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance whose rollback task is canceled
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
	// Async task request ID
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
type SwitchCDBProxyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID
	ProxyGroupId *string `json:"ProxyGroupId,omitnil,omitempty" name:"ProxyGroupId"`
}

type SwitchCDBProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID
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
	// Disaster recovery instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type SwitchDrInstanceToMasterRequest struct {
	*tchttp.BaseRequest
	
	// Disaster recovery instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
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
}

type SwitchForUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

type TaskDetail struct {
	// Error code.
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

	// ID of an instance associated with a task.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// Async task request ID.
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

// Predefined struct for user
type UpgradeCDBProxyVersionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID
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
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database proxy ID
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
	// Async request ID
	// Note: this field may return `null`, indicating that no valid value can be found.
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
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Version of primary instance database engine. Value range: 5.6, 5.7
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Switch mode for accessing the new instance.  Valid values:  `0` (switch immediately), `1` (switch within a time window). Default value: `0`. If the value is `1`, the switch process will be performed within a time window. Or, you can call the [SwitchForUpgrade](https://intl.cloud.tencent.com/document/product/236/15864?from_cn_redirect=1) API to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// Whether to upgrade kernel minor version. Valid values: 1 (upgrade kernel minor version), 0 (upgrade database engine).
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitnil,omitempty" name:"UpgradeSubversion"`

	// Delay threshold. Value range: 1-10
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
}

type UpgradeDBInstanceEngineVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [instance list querying API](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Version of primary instance database engine. Value range: 5.6, 5.7
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Switch mode for accessing the new instance.  Valid values:  `0` (switch immediately), `1` (switch within a time window). Default value: `0`. If the value is `1`, the switch process will be performed within a time window. Or, you can call the [SwitchForUpgrade](https://intl.cloud.tencent.com/document/product/236/15864?from_cn_redirect=1) API to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// Whether to upgrade kernel minor version. Valid values: 1 (upgrade kernel minor version), 0 (upgrade database engine).
	UpgradeSubversion *int64 `json:"UpgradeSubversion,omitnil,omitempty" name:"UpgradeSubversion"`

	// Delay threshold. Value range: 1-10
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceEngineVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceEngineVersionResponseParams struct {
	// Async task ID. The task execution result can be queried using the [async task execution result querying API](https://intl.cloud.tencent.com/document/api/236/20410?from_cn_redirect=1).
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
	// Instance ID in the format of `cdb-c1nl9rpv` or `cdbro-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Memory size in MB after upgrade. To ensure that the `Memory` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the specifications of the memory that can be upgraded to.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size in GB after upgrade. To ensure that the `Volume` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the specifications of the disk that can be upgraded to.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// AZ information of secondary database 1, which is the `Zone` value of the instance by default. This parameter can be specified when upgrading primary instances in multi-AZ mode and is meaningless for read-only or disaster recovery instances. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the supported AZs.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Version of primary instance database engine. Valid values: 5.5, 5.6, 5.7.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Switch mode for accessing the new instance.  Valid values:  `0` (switch immediately), `1` (switch within a time window). Default value: `0`. If the value is `1`, the switch process will be performed within a time window. Or, you can call the [SwitchForUpgrade](https://intl.cloud.tencent.com/document/product/236/15864?from_cn_redirect=1) API to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// AZ information of secondary database 2, which is empty by default. This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Instance type. Valid values: master (primary instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// The resource isolation type after the instance is upgraded. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). If this parameter is left empty, the resource isolation type will be the same as the original one.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// The number of CPU cores after the instance is upgraded. If this parameter is left empty, the minimum value will be automatically filled based on the value specified by `Memory`.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// QuickChange options. Valid values: `0` (common upgrade), `1` (QuickChange), `2` (QuickChange first). After QuickChange is enabled, the required resources will be checked. QuickChange will be performed only when the required resources support the feature; otherwise, an error message will be returned.
	FastUpgrade *int64 `json:"FastUpgrade,omitnil,omitempty" name:"FastUpgrade"`

	// Delay threshold. Value range: 1-10. Default value: `10`.
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// Whether to migrate the source node across AZs. Valid values: `0` (no), `1`(yes). Default value: `0`. If it is `1`, you can modify the source node AZ.
	CrossCluster *int64 `json:"CrossCluster,omitnil,omitempty" name:"CrossCluster"`

	// New AZ of the source node. This field is only valid when `CrossCluster` is `1`. Only migration across AZs in the same region is supported.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Processing logic of the intra-AZ read-only instance for cross-cluster migration. Valid values: `together` (intra-AZ read-only instances will be migrated to the target AZ with the source instance by default.), `severally` (intra-AZ read-only instances will maintain the original deployment mode and will not be migrated to the target AZ.).
	RoTransType *string `json:"RoTransType,omitnil,omitempty" name:"RoTransType"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `cdb-c1nl9rpv` or `cdbro-c1nl9rpv`. It is the same as the instance ID displayed on the TencentDB Console page. You can use the [DescribeDBInstances](https://intl.cloud.tencent.com/document/api/236/15872?from_cn_redirect=1) API to query the ID, whose value is the `InstanceId` value in output parameters.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Memory size in MB after upgrade. To ensure that the `Memory` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the specifications of the memory that can be upgraded to.
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size in GB after upgrade. To ensure that the `Volume` value to be passed in is valid, please use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the specifications of the disk that can be upgraded to.
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Data replication mode. Valid values: 0 (async), 1 (semi-sync), 2 (strong sync). This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	ProtectMode *int64 `json:"ProtectMode,omitnil,omitempty" name:"ProtectMode"`

	// Deployment mode. Valid values: 0 (single-AZ), 1 (multi-AZ). Default value: 0. This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// AZ information of secondary database 1, which is the `Zone` value of the instance by default. This parameter can be specified when upgrading primary instances in multi-AZ mode and is meaningless for read-only or disaster recovery instances. You can use the [DescribeDBZoneConfig](https://intl.cloud.tencent.com/document/product/236/17229?from_cn_redirect=1) API to query the supported AZs.
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// Version of primary instance database engine. Valid values: 5.5, 5.6, 5.7.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// Switch mode for accessing the new instance.  Valid values:  `0` (switch immediately), `1` (switch within a time window). Default value: `0`. If the value is `1`, the switch process will be performed within a time window. Or, you can call the [SwitchForUpgrade](https://intl.cloud.tencent.com/document/product/236/15864?from_cn_redirect=1) API to trigger the process.
	WaitSwitch *int64 `json:"WaitSwitch,omitnil,omitempty" name:"WaitSwitch"`

	// AZ information of secondary database 2, which is empty by default. This parameter can be specified when upgrading primary instances and is meaningless for read-only or disaster recovery instances.
	BackupZone *string `json:"BackupZone,omitnil,omitempty" name:"BackupZone"`

	// Instance type. Valid values: master (primary instance), dr (disaster recovery instance), ro (read-only instance). Default value: master.
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// The resource isolation type after the instance is upgraded. Valid values: `UNIVERSAL` (general instance), `EXCLUSIVE` (dedicated instance), `BASIC` (basic instance). If this parameter is left empty, the resource isolation type will be the same as the original one.
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// The number of CPU cores after the instance is upgraded. If this parameter is left empty, the minimum value will be automatically filled based on the value specified by `Memory`.
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// QuickChange options. Valid values: `0` (common upgrade), `1` (QuickChange), `2` (QuickChange first). After QuickChange is enabled, the required resources will be checked. QuickChange will be performed only when the required resources support the feature; otherwise, an error message will be returned.
	FastUpgrade *int64 `json:"FastUpgrade,omitnil,omitempty" name:"FastUpgrade"`

	// Delay threshold. Value range: 1-10. Default value: `10`.
	MaxDelayTime *int64 `json:"MaxDelayTime,omitnil,omitempty" name:"MaxDelayTime"`

	// Whether to migrate the source node across AZs. Valid values: `0` (no), `1`(yes). Default value: `0`. If it is `1`, you can modify the source node AZ.
	CrossCluster *int64 `json:"CrossCluster,omitnil,omitempty" name:"CrossCluster"`

	// New AZ of the source node. This field is only valid when `CrossCluster` is `1`. Only migration across AZs in the same region is supported.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Processing logic of the intra-AZ read-only instance for cross-cluster migration. Valid values: `together` (intra-AZ read-only instances will be migrated to the target AZ with the source instance by default.), `severally` (intra-AZ read-only instances will maintain the original deployment mode and will not be migrated to the target AZ.).
	RoTransType *string `json:"RoTransType,omitnil,omitempty" name:"RoTransType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// Order ID.
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// Async task request ID, which can be used to query the execution result of an async task.
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