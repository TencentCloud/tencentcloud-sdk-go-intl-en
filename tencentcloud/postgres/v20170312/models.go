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

type AccountInfo struct {

	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Account
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Account remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Account status. 1: creating, 2: normal, 3: modifying, 4: resetting password, -1: deleting
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Account creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Account last modified time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Whether to disable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
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

		// Async task flow ID
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

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Purchasable specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeProductConfig` API.
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL kernel version. Currently, only two versions are supported: 9.3.5 and 9.5.4.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Instance capacity size in GB.
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// Number of instances purchased at a time. Value range: 1–100.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Length of purchase in months. Currently, only 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36 are supported.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance billing type.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: no.
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// VPC ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Renewal flag. 0: normal renewal (default), 1: auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no)
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order number list. Each instance corresponds to an order number.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DBBackup struct {

	// Unique backup file ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// File size in KB
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// Type (0: scheduled)
	Way *int64 `json:"Way,omitempty" name:"Way"`

	// Backup mode (1: full)
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Status (1: creating, 2: success, 3: failure)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// DB list
	DbList []*string `json:"DbList,omitempty" name:"DbList" list`

	// Download address on private network
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// Download address on public network
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`
}

type DBInstance struct {

	// Instance region such as ap-guangzhou, which corresponds to the `Region` field of `RegionSet`
	Region *string `json:"Region,omitempty" name:"Region"`

	// Instance AZ such as ap-guangzhou-3, which corresponds to the `Zone` field of `ZoneSet`
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// Instance status. Valid values: applying, init (to be initialized), initing (initializing), running, limited run, isolated, recycling, recycled, job running, offline, migrating, expanding, readonly, restarting
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// Assigned instance memory size in GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory,omitempty" name:"DBInstanceMemory"`

	// Assigned instance storage capacity in GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage,omitempty" name:"DBInstanceStorage"`

	// Number of assigned CPUs
	DBInstanceCpu *uint64 `json:"DBInstanceCpu,omitempty" name:"DBInstanceCpu"`

	// Purchasable specification ID
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" name:"DBInstanceClass"`

	// Instance type. 1: primary (main instance), 2: readonly (read-only instance), 3: guard (disaster recovery instance), 4: temp (temp instance)
	DBInstanceType *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`

	// Instance edition. Currently, only `standard` edition (dual-server high-availability one-main-one-subordinate edition) is supported
	DBInstanceVersion *string `json:"DBInstanceVersion,omitempty" name:"DBInstanceVersion"`

	// Instance database character set
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// PostgreSQL kernel version
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance last modified time
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Instance expiration time
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// Instance isolation time
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// Billing mode. postpaid: pay-as-you-go
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// Whether to renew automatically. 1: yes, 0: no
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// Instance network connection information
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`

	// Machine type
	Type *string `json:"Type,omitempty" name:"Type"`

	// 
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Whether the instance supports IPv6 address access. Valid values: 1 (yes), 0 (no)
	SupportIpv6 *uint64 `json:"SupportIpv6,omitempty" name:"SupportIpv6"`
}

type DBInstanceNetInfo struct {

	// DNS domain name
	Address *string `json:"Address,omitempty" name:"Address"`

	// IP address
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Connection port address
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Network type. Valid values: inner (private address of classic network), private (private address of VPC), public (public address of classic network/VPC)
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Network connection status
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Number of entries returned per page. Default value: 20. Value range: 1–100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to sort by creation time or username. Valid values: `createTime` (sort by creation time), `name` (sort by username)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Whether returns are sorted in ascending or descending order. Valid values: `desc` (descending), `asc` (ascending)
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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

		// Number of date entries returned for this API call.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Account list details.
		Details []*AccountInfo `json:"Details,omitempty" name:"Details" list`

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

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-4wdeb0zv.
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Backup mode (1: full). Currently, only full backup is supported. The value is 1.
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Number of entries to be returned per page for backup list. Default value: 20. Minimum value: 1. Maximum value: 100. (If this parameter is left empty or 0, the default value will be used)
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of backup files in the returned backup list
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Backup list
		BackupList []*DBBackup `json:"BackupList,omitempty" name:"BackupList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBErrlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-5bq3wfjd
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-01-01 00:00:00, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Search keyword
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys" list`

	// Number of entries returned per page. Value range: 1–100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBErrlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBErrlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBErrlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of date entries returned for this call
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Error log list
		Details []*ErrLogDetail `json:"Details,omitempty" name:"Details" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBErrlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBErrlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance details.
		DBInstance *DBInstance `json:"DBInstance,omitempty" name:"DBInstance"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Filter. Valid values: db-instance-id, db-instance-name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of entries returned per page. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Page number, starting from 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
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

		// Number of instances found.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Instance details set.
		DBInstanceSet []*DBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet" list`

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

type DescribeDBSlowlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Metric for sorting. Valid values: `sum_calls` (total number of calls), `sum_cost_time` (total time consumed)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. desc: descending, asc: ascending
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// Number of entries returned per page. Value range: 1–100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of date entries returned this time
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Slow query log details
		Detail *SlowlogDetail `json:"Detail,omitempty" name:"Detail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSlowlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBXlogsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-4wdeb0zv.
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries returned per page in paged query. Value range: 1–100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBXlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBXlogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBXlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of date entries returned this time.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Xlog list
		XlogList []*Xlog `json:"XlogList,omitempty" name:"XlogList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBXlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBXlogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
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

		// Database information
		Items []*string `json:"Items,omitempty" name:"Items" list`

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

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// Order name set
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of orders
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Order array
		Deals []*PgDeal `json:"Deals,omitempty" name:"Deals" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOrdersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest

	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Purchasable specification list.
		SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned results.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Region information set.
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of returned results.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// AZ information set.
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest

	// ID of the instance to be deleted
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DestroyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ErrLogDetail struct {

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Database name
	Database *string `json:"Database,omitempty" name:"Database"`

	// Error occurrence time
	ErrTime *string `json:"ErrTime,omitempty" name:"ErrTime"`

	// Error message
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type Filter struct {

	// Filter name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// One or more filter values.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID set.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// Instance admin account username.
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// Password corresponding to instance root account username.
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// Instance character set. Valid values: UTF8, LATIN1.
	Charset *string `json:"Charset,omitempty" name:"Charset"`
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

		// Instance ID set.
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

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

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeProductConfig` API.
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// Storage capacity size in GB.
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// Number of instances. Maximum value: 100. If you need to create more instances at a time, please contact customer service.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Length of purchase in months. Currently, only 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36 are supported.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Billing ID, which can be obtained through the `Pid` field in the returned value of the `DescribeProductConfig` API.
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// Instance billing type. Valid value: POSTPAID_BY_HOUR (pay-as-you-go)
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Original price in 0.01 CNY.
		OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// Discounted price in 0.01 CNY.
		Price *uint64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Renewal duration in months. Maximum value: 48
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *InquiryPriceRenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total cost before discount; for example, 24650 indicates 246.5 CNY
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// Actual amount payable; for example, 24650 indicates 246.5 CNY
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance ID in the format of postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance billing type. Valid value: `POSTPAID_BY_HOUR` (pay-as-you-go hourly)
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total cost before discount.
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// Actual amount payable
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// New remarks corresponding to user `UserName`
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// Database instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// New name of database instance
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

	// TencentDB for PostgreSQL instance ID array
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// New project ID of TencentDB for PostgreSQL instance
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
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

		// Number of successfully transferred instances
		Count *int64 `json:"Count,omitempty" name:"Count"`

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

type NormalQueryItem struct {

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Number of calls
	Calls *int64 `json:"Calls,omitempty" name:"Calls"`

	// Granularity
	CallsGrids []*int64 `json:"CallsGrids,omitempty" name:"CallsGrids" list`

	// Total time consumed
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`

	// Number of affected rows
	Rows *int64 `json:"Rows,omitempty" name:"Rows"`

	// Minimum time consumed
	MinCostTime *float64 `json:"MinCostTime,omitempty" name:"MinCostTime"`

	// Maximum time consumed
	MaxCostTime *float64 `json:"MaxCostTime,omitempty" name:"MaxCostTime"`

	// Time of the earliest slow SQL statement
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// Time of the latest slow SQL statement
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// Shared memory blocks for reads
	SharedReadBlks *int64 `json:"SharedReadBlks,omitempty" name:"SharedReadBlks"`

	// Shared memory blocks for writes
	SharedWriteBlks *int64 `json:"SharedWriteBlks,omitempty" name:"SharedWriteBlks"`

	// Total IO read time
	ReadCostTime *int64 `json:"ReadCostTime,omitempty" name:"ReadCostTime"`

	// Total IO write time
	WriteCostTime *int64 `json:"WriteCostTime,omitempty" name:"WriteCostTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// Slow SQL statement after desensitization
	NormalQuery *string `json:"NormalQuery,omitempty" name:"NormalQuery"`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Whether to enable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
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

		// Async task flow ID
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

type PgDeal struct {

	// Order name
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// User
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Number of instances involved in order
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// Billing mode. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// Instance ID array
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`
}

type RegionInfo struct {

	// Region abbreviation
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region number
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// Availability status. UNAVAILABLE: unavailable, AVAILABLE: available
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of `postgres-6fego161`
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Renewal duration in months
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order name
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance account name
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// New password corresponding to `UserName` account
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

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Async flow ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Instance ID array
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// Renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal upon expiration
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of successfully set instances
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SlowlogDetail struct {

	// Total time consumed
	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// Total number of calls
	TotalCalls *int64 `json:"TotalCalls,omitempty" name:"TotalCalls"`

	// List of slow SQL statements after desensitization
	NormalQueries []*NormalQueryItem `json:"NormalQueries,omitempty" name:"NormalQueries" list`
}

type SpecInfo struct {

	// Region abbreviation, which corresponds to the `Region` field of `RegionSet`
	Region *string `json:"Region,omitempty" name:"Region"`

	// AZ abbreviate, which corresponds to the `Zone` field of `ZoneSet`
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Specification details list
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList,omitempty" name:"SpecItemInfoList" list`
}

type SpecItemInfo struct {

	// Specification ID
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL kernel version number
	Version *string `json:"Version,omitempty" name:"Version"`

	// Full version name corresponding to kernel number
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size in MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Maximum storage capacity in GB supported by this specification
	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// Minimum storage capacity in GB supported by this specification
	MinStorage *uint64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// Estimated QPS for this specification
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Billing ID for this specification
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// Machine type
	Type *string `json:"Type,omitempty" name:"Type"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance memory size in GB after upgrade
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance disk size in GB after upgrade
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently)
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Transaction name.
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Xlog struct {

	// Unique backup file ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Download address on private network
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// Download address on public network
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`
}

type ZoneInfo struct {

	// AZ abbreviation
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// AZ number
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Availability status. UNAVAILABLE: unavailable, AVAILABLE: available
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`

	// Whether the availability zone supports IPv6 address access
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitempty" name:"ZoneSupportIpv6"`
}
