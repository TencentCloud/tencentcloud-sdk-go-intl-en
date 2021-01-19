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

package v20180411

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CloneAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Source user account name
	SrcUser *string `json:"SrcUser,omitempty" name:"SrcUser"`

	// Source user host
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// Target user account name
	DstUser *string `json:"DstUser,omitempty" name:"DstUser"`

	// Target user host
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// Description of a target account
	DstDesc *string `json:"DstDesc,omitempty" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloneAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// ID of an instance for which to disable public network access. The ID is in the format of dcdbt-ow728lmc and can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task status can be queried through the `DescribeFlow` API.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConstraintRange struct {

	// Minimum value when constraint type is `section`
	Min *string `json:"Min,omitempty" name:"Min"`

	// Maximum value when constraint type is `section`
	Max *string `json:"Max,omitempty" name:"Max"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`

	// Access host allowed for a source user
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`

	// Access host allowed for a target user
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// `ReadOnly` attribute of a source account
	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`

	// `ReadOnly` attribute of a target account
	DstReadOnly *string `json:"DstReadOnly,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Account password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail; 3: only the secondary will be read from.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Account remarks, which can contain 0-256 letters, digits, and common symbols.
	Description *string `json:"Description,omitempty" name:"Description"`

	// If the secondary delay exceeds the set value of this parameter, the secondary will be deemed to have failed.
	// It is recommended that this parameter be set to a value greater than 10. This parameter takes effect when `ReadOnly` is 1 or 2.
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID, which is passed through from the input parameters.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Username, which is passed through from the input parameters.
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// Host allowed for access, which is passed through from the input parameters.
		Host *string `json:"Host,omitempty" name:"Host"`

		// Passed through from the input parameters.
		ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Host from which a user can log in (corresponding to the `host` field for a MySQL user; a user is uniquely identified by username and host; this parameter is in IP format and ends with % for IP range; % can be entered; if this parameter is left empty, % will be used by default)
	Host *string `json:"Host,omitempty" name:"Host"`

	// User remarks
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last updated time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Read-only flag. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// If the secondary delay exceeds the set value of this parameter, the secondary will be deemed to have failed.
	// It is recommended that this parameter be set to a value greater than 10. This parameter takes effect when `ReadOnly` is 1 or 2.
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

type DBParamValue struct {

	// Parameter name
	Param *string `json:"Param,omitempty" name:"Param"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DCDBInstanceInfo struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Application ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Numeric ID of a VPC
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet Digital ID
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Status
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Private IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Auto-renewal flag
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Number of shards
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// Expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Isolation time
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitempty" name:"IsolatedTimestamp"`

	// Account ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Shard details
	ShardDetail []*ShardInfo `json:"ShardDetail,omitempty" name:"ShardDetail" list`

	// Number of nodes. 2: one master and one slave; 3: one master and two slaves
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Temporary instance flag. 0: non-temporary instance
	IsTmp *int64 `json:"IsTmp,omitempty" name:"IsTmp"`

	// Dedicated cluster ID. If this parameter is empty, the instance is a non-dedicated cluster instance
	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`

	// VPC ID in string type
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID in string type
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// Numeric ID of instance (this field is obsolete and should not be depended on)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Domain name for public network access, which can be resolved by the public network
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network
	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Product type ID (this field is obsolete and should not be depended on)
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Last updated time of an instance in the format of 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Database engine
	DbEngine *string `json:"DbEngine,omitempty" name:"DbEngine"`

	// Database engine version
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// Billing mode
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// Async task flow ID when an async task is in progress on an instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	Locker *int64 `json:"Locker,omitempty" name:"Locker"`

	// Public network access status. 0: not enabled; 1: enabled; 2: disabled; 3: enabling
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// Whether the instance supports audit. 1: yes; 0: no
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Indicates whether the instance uses IPv6
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`

	// Private network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vipv6 *string `json:"Vipv6,omitempty" name:"Vipv6"`

	// Public network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanVipv6 *string `json:"WanVipv6,omitempty" name:"WanVipv6"`

	// Public network IPv6 port
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitempty" name:"WanPortIpv6"`

	// Public network IPv6 status
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitempty" name:"WanStatusIpv6"`

	// DCN type. Valid values: 0 (null), 1 (primary instance), 2 (disaster recovery instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`

	// DCN status. Valid values: 0 (null), 1 (creating), 2 (syncing), 3 (disconnected)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnStatus *int64 `json:"DcnStatus,omitempty" name:"DcnStatus"`

	// The number of DCN disaster recovery instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnDstNum *int64 `json:"DcnDstNum,omitempty" name:"DcnDstNum"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (standard primary instance), `3` (standard disaster recovery instance), `4` (dedicated disaster recovery instance)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DCDBShardInfo struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard SQL passthrough ID, which is used to pass through SQL statements to the specified shard for execution
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// Globally unique shard ID
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// Status. 0: creating; 1: processing; 2: running; 3: shard not initialized
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Status description
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// VPC ID in string format
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID in string format
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Expiration time
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Number of nodes. 2: one primary and one secondary; 3: one primary and two secondaries
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Storage utilization in %
	StorageUsage *float64 `json:"StorageUsage,omitempty" name:"StorageUsage"`

	// Memory utilization in %
	MemoryUsage *float64 `json:"MemoryUsage,omitempty" name:"MemoryUsage"`

	// Numeric ID of a shard (this field is obsolete and should not be depended on)
	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`

	// ProductID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Proxy version
	ProxyVersion *string `json:"ProxyVersion,omitempty" name:"ProxyVersion"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// Master AZ of a shard
	// Note: this field may return null, indicating that no valid values can be obtained.
	ShardMasterZone *string `json:"ShardMasterZone,omitempty" name:"ShardMasterZone"`

	// List of secondary AZs of a shard
	// Note: this field may return null, indicating that no valid values can be obtained.
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitempty" name:"ShardSlaveZones" list`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

type Database struct {

	// Database name
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DatabaseFunction struct {

	// Function name
	Func *string `json:"Func,omitempty" name:"Func"`
}

type DatabaseProcedure struct {

	// Stored procedure name
	Proc *string `json:"Proc,omitempty" name:"Proc"`
}

type DatabaseTable struct {

	// Table name
	Table *string `json:"Table,omitempty" name:"Table"`
}

type DatabaseView struct {

	// View name
	View *string `json:"View,omitempty" name:"View"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for a user
	Host *string `json:"Host,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Type. Valid values: table; view; proc; func; \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be queried (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type name. For example, if `Type` is table, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitempty" name:"Object"`

	// If `Type` = table and `ColName` is `\*`, the permissions of the table will be queried; if `ColName` is a specific field name, the permissions of the corresponding field will be queried
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// List of permissions.
		Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

		// Database account username
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// Database account host
		Host *string `json:"Host,omitempty" name:"Host"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID, which is passed through from the input parameters.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// List of instance users.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Users []*DBAccount `json:"Users,omitempty" name:"Users" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard ID in the format of shard-7noic7tv
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// Requested log type. Valid values: 1 (binlog); 2 (cold backup); 3 (errlog); 4 (slowlog).
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID in the format of dcdbt-ow728lmc.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Requested log type. Valid values: 1 (binlog); 2 (cold backup); 3 (errlog); 4 (slowlog).
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// Total number of requested logs
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of log files
		Files []*LogFileInfo `json:"Files,omitempty" name:"Files" list`

		// For an instance in a VPC, this prefix plus URI can be used as the download address
		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`

		// For an instance in a common network, this prefix plus URI can be used as the download address
		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`

		// Shard ID in the format of shard-7noic7tv
		ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID in the format of dcdbt-ow7t8lmc.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Requests the current parameter values of a DB
		Params []*ParamDesc `json:"Params,omitempty" name:"Params" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// ID of an instance for which to modify the sync mode. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Sync mode. 0: async; 1: strong sync; 2: downgradable strong sync
		SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`

		// Whether a modification is in progress. 1: yes; 0: no.
		IsModifying *int64 `json:"IsModifying,omitempty" name:"IsModifying"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Query by instance ID or IDs. Instance ID is in the format of dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Search field name. Valid values: instancename (search by instance name); vip (search by private IP); all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Query by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Sort by field. Valid values: projectId; createtime; instancename
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sort by type. Valid values: desc; asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 10. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 1: non-dedicated cluster; 2: dedicated cluster; 0: all
	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`

	// Identifies whether to use the `ExclusterType` field. false: no; true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`

	// Dedicated cluster ID
	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds" list`

	// Tag key used in queries
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys" list`

	// Instance types used in filtering. Valid values: 1 (dedicated instance), 2 (primary instance), 3 (disaster recovery instance). Multiple values should be separated by commas.
	FilterInstanceType *string `json:"FilterInstanceType,omitempty" name:"FilterInstanceType"`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance details
		Instances []*DCDBInstanceInfo `json:"Instances,omitempty" name:"Instances" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Shard ID list.
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitempty" name:"ShardInstanceIds" list`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Sort by field. Only `createtime` is supported currently
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sort by type. Valid values: desc; asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDCDBShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBShardsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDCDBShardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible shards
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Shard information list
		Shards []*DCDBShardInfo `json:"Shards,omitempty" name:"Shards" list`

		// DCN type. Valid values: 0 (null), 1 (primary instance), 2 (disaster recovery instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
		DcnFlag *int64 `json:"DcnFlag,omitempty" name:"DcnFlag"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDCDBShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDCDBShardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Passed through from the input parameters.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Database name.
		DbName *string `json:"DbName,omitempty" name:"DbName"`

		// List of tables.
		Tables []*DatabaseTable `json:"Tables,omitempty" name:"Tables" list`

		// List of views.
		Views []*DatabaseView `json:"Views,omitempty" name:"Views" list`

		// List of stored procedures.
		Procs []*DatabaseProcedure `json:"Procs,omitempty" name:"Procs" list`

		// List of functions.
		Funcs []*DatabaseFunction `json:"Funcs,omitempty" name:"Funcs" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Table name, which can be obtained through the `DescribeDatabaseObjects` API.
	Table *string `json:"Table,omitempty" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance name.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Database name.
		DbName *string `json:"DbName,omitempty" name:"DbName"`

		// Table name.
		Table *string `json:"Table,omitempty" name:"Table"`

		// Column information.
		Cols []*TableColumn `json:"Cols,omitempty" name:"Cols" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow7t8lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of databases on an instance.
		Databases []*Database `json:"Databases,omitempty" name:"Databases" list`

		// Passed through from the input parameters.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Global permission. Valid values: SELECT; INSERT; UPDATE; DELETE; CREATE; DROP; REFERENCES; INDEX; ALTER; CREATE TEMPORARY TABLES; LOCK TABLES; EXECUTE; CREATE VIEW; SHOW VIEW; CREATE ROUTINE; ALTER ROUTINE; EVENT; TRIGGER; SHOW DATABASES 
	// Database permission. Valid values: SELECT; INSERT; UPDATE; DELETE; CREATE; DROP; REFERENCES; INDEX; ALTER; CREATE TEMPORARY TABLES; LOCK TABLES; EXECUTE; CREATE VIEW; SHOW VIEW; CREATE ROUTINE; ALTER ROUTINE; EVENT; TRIGGER 
	// Table/view permission. Valid values: SELECT; INSERT; UPDATE; DELETE; CREATE; DROP; REFERENCES; INDEX; ALTER; CREATE VIEW; SHOW VIEW; TRIGGER 
	// Stored procedure/function permission. Valid values: ALTER ROUTINE; EXECUTE 
	// Field permission. Valid values: INSERT; REFERENCES; SELECT; UPDATE
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

	// Type. Valid values: table; view; proc; func; \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type name. For example, if `Type` is table, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitempty" name:"Object"`

	// If `Type` = table and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
	ColName *string `json:"ColName,omitempty" name:"ColName"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesRequest struct {
	*tchttp.BaseRequest

	// List of IDs of instances to be initialized. The ID is in the format of `dcdbt-ow728lmc` and can be obtained through the `DescribeDCDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// List of parameters. Valid values: character_set_server (character set; required); lower_case_table_names (table name case sensitivity; required; 0: case-sensitive; 1: case-insensitive); innodb_page_size (InnoDB data page; default size: 16 KB); sync_mode (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: strong sync).
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *InitDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDCDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task status can be queried through the `DescribeFlow` API.
		FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds" list`

		// Passed through from the input parameters.
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDCDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {

	// Last modified time of a log
	Mtime *uint64 `json:"Mtime,omitempty" name:"Mtime"`

	// File length
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// Uniform resource identifier (URI) used during log download
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// Filename
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// New account remarks, which can contain 0-256 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// List of IDs of instances to be modified. Instance ID is in the format of dcdbt-ow728lmc.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// ID of the project to be assigned, which can be obtained through the `DescribeProjects` API.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of parameters. Every element is a combination of `Param` and `Value`.
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID in the format of dcdbt-ow728lmc.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Parameter modification results
		Result []*ParamModifyResult `json:"Result,omitempty" name:"Result" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest

	// ID of an instance for which to modify the sync mode. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Sync mode. 0: async; 1: strong sync; 2: downgradable strong sync
	SyncMode *int64 `json:"SyncMode,omitempty" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task status can be queried through the `DescribeFlow` API.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// ID of an instance for which to enable public network access. The ID is in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitempty" name:"Ipv6Flag"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task status can be queried through the `DescribeFlow` API.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ParamConstraint struct {

	// Constraint type, such as enum and section
	Type *string `json:"Type,omitempty" name:"Type"`

	// List of valid values when constraint type is `enum`
	Enum *string `json:"Enum,omitempty" name:"Enum"`

	// Range when constraint type is `section`
	// Note: this field may return null, indicating that no valid values can be obtained.
	Range *ConstraintRange `json:"Range,omitempty" name:"Range"`

	// List of valid values when constraint type is `string`
	String *string `json:"String,omitempty" name:"String"`
}

type ParamDesc struct {

	// Parameter name
	Param *string `json:"Param,omitempty" name:"Param"`

	// Current parameter value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Previously set value, which is the same as `value` after the parameter takes effect. If no value has been set, this field will not be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SetValue *string `json:"SetValue,omitempty" name:"SetValue"`

	// Default value
	Default *string `json:"Default,omitempty" name:"Default"`

	// Parameter constraint
	Constraint *ParamConstraint `json:"Constraint,omitempty" name:"Constraint"`

	// Whether a value has been set. false: no, true: yes
	HaveSetValue *bool `json:"HaveSetValue,omitempty" name:"HaveSetValue"`
}

type ParamModifyResult struct {

	// Renames a parameter
	Param *string `json:"Param,omitempty" name:"Param"`

	// Result of parameter modification. 0: success; -1: failure; -2: invalid parameter value
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of dcdbt-ow728lmc.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// New password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {

	// Shard ID
	ShardInstanceId *string `json:"ShardInstanceId,omitempty" name:"ShardInstanceId"`

	// Shard set ID
	ShardSerialId *string `json:"ShardSerialId,omitempty" name:"ShardSerialId"`

	// Status. 0: creating; 1: processing; 2: running; 3: shard not initialized; -2: shard deleted
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// Memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Numeric ID of a shard
	ShardId *int64 `json:"ShardId,omitempty" name:"ShardId"`

	// Number of nodes. 2: one primary and one secondary; 3: one primary and two secondaries
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Product type ID (this field is obsolete and should not be depended on)
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
}

type TableColumn struct {

	// Column name
	Col *string `json:"Col,omitempty" name:"Col"`

	// Column type
	Type *string `json:"Type,omitempty" name:"Type"`
}
