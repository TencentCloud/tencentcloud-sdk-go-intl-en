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

package v20170312

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

	// Description of target account
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

		// Async task flow ID.
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

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

	// ID of instance for which to disable public network access. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitempty" name:"SrcUserName"`

	// Access host allowed for source user
	SrcHost *string `json:"SrcHost,omitempty" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitempty" name:"DstUserName"`

	// Access host allowed for target user
	DstHost *string `json:"DstHost,omitempty" name:"DstHost"`

	// `ReadOnly` attribute of source account
	SrcReadOnly *string `json:"SrcReadOnly,omitempty" name:"SrcReadOnly"`

	// `ReadOnly` attribute of target account
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

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username, which can contain 1–32 letters, digits, underscores, and hyphens.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Account password, which can contain 6–32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the slave will be used first, and if it is unavailable, the master will be used; 2: the slave will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Account remarks, which can contain 0–256 letters, digits, and common symbols.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Determines whether the slave is unavailable based on the passed-in time
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

	// Read-only flag. 0: no; 1: for the account's SQL requests, the slave will be used first, and if it is unavailable, the master will be used; 2: the slave will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// This field is meaningful for read-only accounts, indicating to select a slave where the master/slave delay is below this value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DelayThresh *int64 `json:"DelayThresh,omitempty" name:"DelayThresh"`
}

type DBBackupTimeConfig struct {

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of daily backup window in the format of `mm:ss`, such as 22:00
	StartBackupTime *string `json:"StartBackupTime,omitempty" name:"StartBackupTime"`

	// End time of daily backup window in the format of `mm:ss`, such as 23:00
	EndBackupTime *string `json:"EndBackupTime,omitempty" name:"EndBackupTime"`
}

type DBInstance struct {

	// Instance ID, which uniquely identifies a TDSQL instance
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Customizable instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Application ID of instance
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Project ID of instance
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance region name, such as ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC ID, which is 0 if the basic network is used
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID, which is 0 if the basic network is used
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance status. 0: creating, 1: processing, 2: running, 3: instance not initialized, -1: instance isolated, -2: instance deleted
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Private IP address
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// Private network port
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// Domain name for public network access, which can be resolved by the public network
	WanDomain *string `json:"WanDomain,omitempty" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network
	WanVip *string `json:"WanVip,omitempty" name:"WanVip"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitempty" name:"WanPort"`

	// Instance creation time in the format of `2006-01-02 15:04:05`
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Last updated time of instance in the format of `2006-01-02 15:04:05`
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Auto-renewal flag. 0: no, 1: yes
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Instance expiration time in the format of `2006-01-02 15:04:05`
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" name:"PeriodEndTime"`

	// Instance account
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// TDSQL version information
	TdsqlVersion *string `json:"TdsqlVersion,omitempty" name:"TdsqlVersion"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// VPC ID in string type
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// VPC subnet ID in string type
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// Original ID of instance (this field is obsolete and should not be depended on)
	OriginSerialId *string `json:"OriginSerialId,omitempty" name:"OriginSerialId"`

	// Number of nodes. 2: one master and one slave, 3: one master and two slaves
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// Whether it is a temp instance. 0: no, non-zero value: yes
	IsTmp *uint64 `json:"IsTmp,omitempty" name:"IsTmp"`

	// Dedicated cluster ID. If this parameter is empty, the instance is a general instance
	ExclusterId *string `json:"ExclusterId,omitempty" name:"ExclusterId"`

	// Numeric ID of instance (this field is obsolete and should not be depended on)
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Product type ID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// Maximum QPS value
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// Async task flow ID when an async task is in progress on an instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	Locker *int64 `json:"Locker,omitempty" name:"Locker"`

	// Current instance running status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// Public network access status. 0: not enabled, 1: enabled, 2: disabled, 3: enabling
	WanStatus *int64 `json:"WanStatus,omitempty" name:"WanStatus"`

	// Whether the instance supports audit. 1: yes, 0: no
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitempty" name:"IsAuditSupported"`

	// Model
	Machine *string `json:"Machine,omitempty" name:"Machine"`

	// Whether data encryption is supported. 1: yes, 0: no
	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitempty" name:"IsEncryptSupported"`

	// Number of CPU cores of instance
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

type DBParamValue struct {

	// Parameter name
	Param *string `json:"Param,omitempty" name:"Param"`

	// Parameter value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Database struct {

	// Database name
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user
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

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be queried (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitempty" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions of the table will be queried; if `ColName` is a specific field name, the permissions of the corresponding field will be queried
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

		// Permission list.
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

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
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

		// Instance user list.
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

type DescribeBackupTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeBackupTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned configurations
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance backup time configuration information
	// Note: this field may return null, indicating that no valid values can be obtained.
		Items []*DBBackupTimeConfig `json:"Items,omitempty" name:"Items" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Queries by instance ID or IDs. Instance ID is in the format of `tdsql-ow728lmc`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Search field name. Valid values: instancename (search by instance name), vip (search by private IP), all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Queries by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitempty" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Sort by field. Valid values: projectId, createtime, instancename
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Queries by `OriginSerialId`
	OriginSerialIds []*string `json:"OriginSerialIds,omitempty" name:"OriginSerialIds" list`

	// Identifies whether to use the `ExclusterType` field. false: no, true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitempty" name:"IsFilterExcluster"`

	// Instance cluster type. 1: non-dedicated cluster, 2: dedicated cluster, 0: all
	ExclusterType *int64 `json:"ExclusterType,omitempty" name:"ExclusterType"`

	// Filters instances by dedicated cluster ID in the format of `dbdc-4ih6uct9`
	ExclusterIds []*string `json:"ExclusterIds,omitempty" name:"ExclusterIds" list`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance details list
		Instances []*DBInstance `json:"Instances,omitempty" name:"Instances" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type *uint64 `json:"Type,omitempty" name:"Type"`
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

		// Instance ID in the format of `tdsql-ow728lmc`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// Total number of requested logs
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// Information such as `uri`, `length`, and `mtime` (modification time)
		Files []*LogFileInfo `json:"Files,omitempty" name:"Files" list`

		// For an instance in a VPC, this prefix plus URI can be used as the download address
		VpcPrefix *string `json:"VpcPrefix,omitempty" name:"VpcPrefix"`

		// For an instance in a common network, this prefix plus URI can be used as the download address
		NormalPrefix *string `json:"NormalPrefix,omitempty" name:"NormalPrefix"`

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

	// Instance ID in the format of `tdsql-ow728lmc`.
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

		// Instance ID in the format of `tdsql-ow728lmc`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Requests the current parameter values of database
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

type DescribeDBPerformanceDetailsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start date in the format of `yyyy-mm-dd`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of `yyyy-mm-dd`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Name of pulled metric. Valid values: long_query, select_total, update_total, insert_total, delete_total, mem_hit_rate, disk_iops, conn_active, is_master_switched, slave_delay
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDBPerformanceDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBPerformanceDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPerformanceDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Master node performance monitoring data
		Master *PerformanceMonitorSet `json:"Master,omitempty" name:"Master"`

		// Slave 1 performance monitoring data
	// Note: this field may return null, indicating that no valid values can be obtained.
		Slave1 *PerformanceMonitorSet `json:"Slave1,omitempty" name:"Slave1"`

		// Slave 2 performance monitoring data. If the instance is one-master-one-slave, it does not have this field
	// Note: this field may return null, indicating that no valid values can be obtained.
		Slave2 *PerformanceMonitorSet `json:"Slave2,omitempty" name:"Slave2"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBPerformanceDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBPerformanceDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPerformanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start date in the format of `yyyy-mm-dd`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of `yyyy-mm-dd`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Name of pulled metric. Valid values: long_query, select_total, update_total, insert_total, delete_total, mem_hit_rate, disk_iops, conn_active, is_master_switched, slave_delay
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDBPerformanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBPerformanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of slow queries
		LongQuery *MonitorData `json:"LongQuery,omitempty" name:"LongQuery"`

		// Number of SELECT operations
		SelectTotal *MonitorData `json:"SelectTotal,omitempty" name:"SelectTotal"`

		// Number of UPDATE operations
		UpdateTotal *MonitorData `json:"UpdateTotal,omitempty" name:"UpdateTotal"`

		// Number of INSERT operations
		InsertTotal *MonitorData `json:"InsertTotal,omitempty" name:"InsertTotal"`

		// Number of DELETE operations
		DeleteTotal *MonitorData `json:"DeleteTotal,omitempty" name:"DeleteTotal"`

		// Cache hit rate
		MemHitRate *MonitorData `json:"MemHitRate,omitempty" name:"MemHitRate"`

		// Number of disk IOs per second
		DiskIops *MonitorData `json:"DiskIops,omitempty" name:"DiskIops"`

		// Number of active connections
		ConnActive *MonitorData `json:"ConnActive,omitempty" name:"ConnActive"`

		// Whether master/slave switch occurred. 1: yes, 0: no
		IsMasterSwitched *MonitorData `json:"IsMasterSwitched,omitempty" name:"IsMasterSwitched"`

		// Master/slave delay
		SlaveDelay *MonitorData `json:"SlaveDelay,omitempty" name:"SlaveDelay"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBPerformanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBResourceUsageDetailsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start date in the format of `yyyy-mm-dd`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of `yyyy-mm-dd`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Name of pulled metric. Valid values: data_disk_available, binlog_disk_available, mem_available, cpu_usage_rate
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDBResourceUsageDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBResourceUsageDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBResourceUsageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Master node resource usage monitoring node
		Master *ResourceUsageMonitorSet `json:"Master,omitempty" name:"Master"`

		// Slave 1 resource usage monitoring node
	// Note: this field may return null, indicating that no valid values can be obtained.
		Slave1 *ResourceUsageMonitorSet `json:"Slave1,omitempty" name:"Slave1"`

		// Slave 2 resource usage monitoring node
	// Note: this field may return null, indicating that no valid values can be obtained.
		Slave2 *ResourceUsageMonitorSet `json:"Slave2,omitempty" name:"Slave2"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBResourceUsageDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBResourceUsageDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBResourceUsageRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start date in the format of `yyyy-mm-dd`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End date in the format of `yyyy-mm-dd`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Name of pulled metric. Valid values: data_disk_available, binlog_disk_available, mem_available, cpu_usage_rate
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeDBResourceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBResourceUsageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBResourceUsageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Available capacity of binlog disk in GB
		BinlogDiskAvailable *MonitorData `json:"BinlogDiskAvailable,omitempty" name:"BinlogDiskAvailable"`

		// Available disk capacity in GB
		DataDiskAvailable *MonitorData `json:"DataDiskAvailable,omitempty" name:"DataDiskAvailable"`

		// CPU utilization
		CpuUsageRate *MonitorData `json:"CpuUsageRate,omitempty" name:"CpuUsageRate"`

		// Available memory size in GB
		MemAvailable *MonitorData `json:"MemAvailable,omitempty" name:"MemAvailable"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBResourceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBResourceUsageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Data entry number starting from which to return results
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Query start time in the format of `2016-07-23 14:55:20`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of `2016-08-22 14:55:20`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Specific name of database to be queried
	Db *string `json:"Db,omitempty" name:"Db"`

	// Sort by metric. Valid values: query_time_sum, query_count
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Whether to query slow queries of the slave. 0: master, 1: slave
	Slave *int64 `json:"Slave,omitempty" name:"Slave"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Slow query log data
		Data []*SlowLogData `json:"Data,omitempty" name:"Data" list`

		// Sum of all statement lock durations
		LockTimeSum *float64 `json:"LockTimeSum,omitempty" name:"LockTimeSum"`

		// Total number of statement queries
		QueryCount *int64 `json:"QueryCount,omitempty" name:"QueryCount"`

		// Total number of results
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// Sum of all statement query durations
		QueryTimeSum *float64 `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `dcdbt-ow7t8lmc`.
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

		// List of databases on instance.
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

type DescribeFlowRequest struct {
	*tchttp.BaseRequest

	// Flow ID returned by async request API.
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Flow status. 0: succeeded, 1: failed, 2: running
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFlowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeLogFileRetentionPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogFileRetentionPeriodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID in the format of `tdsql-ow728lmc`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Backup log retention days
		Days *uint64 `json:"Days,omitempty" name:"Days"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogFileRetentionPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogFileRetentionPeriodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// Database name. `\*` indicates that global permissions will be set (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored. If `DbName` is not `\*`, the input parameter `Type` is required.
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Global permission. Valid values: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER, SHOW DATABASES 
	// Database permission. Valid values: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER 
	// Table/view permission. Valid values: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE VIEW, SHOW VIEW, TRIGGER 
	// Stored procedure/function permission. Valid values: ALTER ROUTINE, EXECUTE 
	// Field permission. Valid values: INSERT, REFERENCES, SELECT, UPDATE
	Privileges []*string `json:"Privileges,omitempty" name:"Privileges" list`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitempty" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitempty" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
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

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// List of IDs of instances to be initialized. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Parameter list. Valid values: character_set_server (character set; required); lower_case_table_names (table name case sensitivity; required; 0: case-sensitive, 1: case-insensitive); innodb_page_size (InnoDB data page; default size: 16 KB); sync_mode (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: strong sync).
	Params []*DBParamValue `json:"Params,omitempty" name:"Params" list`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async task ID. The task status can be queried through the `DescribeFlow` API.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// Passed through from the input parameters.
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InitDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {

	// Last modified time of log
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

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// New account remarks, which can contain 0–256 characters.
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

type ModifyBackupTimeRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Start time of daily backup window in the format of `mm:ss`, such as 22:00
	StartBackupTime *string `json:"StartBackupTime,omitempty" name:"StartBackupTime"`

	// End time of daily backup window in the format of `mm:ss`, such as 23:59
	EndBackupTime *string `json:"EndBackupTime,omitempty" name:"EndBackupTime"`
}

func (r *ModifyBackupTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBackupTimeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBackupTimeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Setting status. 0: success
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBackupTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBackupTimeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// ID of instance to be modified, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// New name of instance, which can contain letters, digits, underscores, and hyphens.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// List of IDs of instances to be modified. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
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

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Parameter list. Each element is a combination of `Param` and `Value`.
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

		// Instance ID in the format of `tdsql-ow728lmc`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// Parameter modification result
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

type ModifyLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Retention days up to 30
	Days *uint64 `json:"Days,omitempty" name:"Days"`
}

func (r *ModifyLogFileRetentionPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLogFileRetentionPeriodRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID in the format of `tdsql-ow728lmc`.
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogFileRetentionPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLogFileRetentionPeriodResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MonitorData struct {

	// Start time in the format of `2018-03-24 23:59:59`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time in the format of `2018-03-24 23:59:59`
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Monitoring data
	Data []*float64 `json:"Data,omitempty" name:"Data" list`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// ID of instance for which to enable public network access. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

	// Constraint type, such as `enum` and `section`
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

	// Previously set value, which is the same as `value` after the parameter takes effect.
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

	// Result of parameter modification. 0: success, -1: failure, -2: invalid parameter value
	Code *int64 `json:"Code,omitempty" name:"Code"`
}

type PerformanceMonitorSet struct {

	// Number of UPDATE operations
	UpdateTotal *MonitorData `json:"UpdateTotal,omitempty" name:"UpdateTotal"`

	// Number of disk IOs per second
	DiskIops *MonitorData `json:"DiskIops,omitempty" name:"DiskIops"`

	// Number of active connections
	ConnActive *MonitorData `json:"ConnActive,omitempty" name:"ConnActive"`

	// Cache hit rate
	MemHitRate *MonitorData `json:"MemHitRate,omitempty" name:"MemHitRate"`

	// Master/slave delay
	SlaveDelay *MonitorData `json:"SlaveDelay,omitempty" name:"SlaveDelay"`

	// Number of SELECT operations
	SelectTotal *MonitorData `json:"SelectTotal,omitempty" name:"SelectTotal"`

	// Number of slow queries
	LongQuery *MonitorData `json:"LongQuery,omitempty" name:"LongQuery"`

	// Number of DELETE operations
	DeleteTotal *MonitorData `json:"DeleteTotal,omitempty" name:"DeleteTotal"`

	// Number of INSERT operations
	InsertTotal *MonitorData `json:"InsertTotal,omitempty" name:"InsertTotal"`

	// Whether master/slave switch occurred. 1: yes, 0: no
	IsMasterSwitched *MonitorData `json:"IsMasterSwitched,omitempty" name:"IsMasterSwitched"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitempty" name:"Host"`

	// New password, which can contain 6–32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
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

type ResourceUsageMonitorSet struct {

	// Available capacity of binlog disk in GB
	BinlogDiskAvailable *MonitorData `json:"BinlogDiskAvailable,omitempty" name:"BinlogDiskAvailable"`

	// CPU utilization
	CpuUsageRate *MonitorData `json:"CpuUsageRate,omitempty" name:"CpuUsageRate"`

	// Available memory size in GB
	MemAvailable *MonitorData `json:"MemAvailable,omitempty" name:"MemAvailable"`

	// Available disk capacity in GB
	DataDiskAvailable *MonitorData `json:"DataDiskAvailable,omitempty" name:"DataDiskAvailable"`
}

type SlowLogData struct {

	// Statement checksum for querying details
	CheckSum *string `json:"CheckSum,omitempty" name:"CheckSum"`

	// Database name
	Db *string `json:"Db,omitempty" name:"Db"`

	// Abstracted SQL statement
	FingerPrint *string `json:"FingerPrint,omitempty" name:"FingerPrint"`

	// Average lock duration
	LockTimeAvg *string `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`

	// Maximum lock duration
	LockTimeMax *string `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// Minimum lock duration
	LockTimeMin *string `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// Sum of lock durations
	LockTimeSum *string `json:"LockTimeSum,omitempty" name:"LockTimeSum"`

	// Number of queries
	QueryCount *string `json:"QueryCount,omitempty" name:"QueryCount"`

	// Average query duration
	QueryTimeAvg *string `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`

	// Maximum query duration
	QueryTimeMax *string `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// Minimum query duration
	QueryTimeMin *string `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`

	// Sum of query durations
	QueryTimeSum *string `json:"QueryTimeSum,omitempty" name:"QueryTimeSum"`

	// Number of scanned rows
	RowsExaminedSum *string `json:"RowsExaminedSum,omitempty" name:"RowsExaminedSum"`

	// Number of sent rows
	RowsSentSum *string `json:"RowsSentSum,omitempty" name:"RowsSentSum"`

	// Last execution time
	TsMax *string `json:"TsMax,omitempty" name:"TsMax"`

	// First execution time
	TsMin *string `json:"TsMin,omitempty" name:"TsMin"`

	// Account
	User *string `json:"User,omitempty" name:"User"`

	// Sample SQL
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExampleSql *string `json:"ExampleSql,omitempty" name:"ExampleSql"`
}
