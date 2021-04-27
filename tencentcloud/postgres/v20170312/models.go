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

type AddDBInstanceToReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("AddDBInstanceToReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddDBInstanceToReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Whether to disable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return errors.New("CloseDBExtranetAccessRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an instance
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return errors.New("CloseServerlessDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
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

	// Number of instances purchased at a time. Value range: 1-100.
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

	// Activity ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no)
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// The information of tags to be associated with instances. This parameter is left empty by default.
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecCode")
	delete(f, "DBVersion")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AutoRenewFlag")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order number list. Each instance corresponds to an order number.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// Bill ID of frozen fees
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// ID set of instances which have been created successfully. The parameter value will be returned only when the billing mode is postpaid.
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Purchasable specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeProductConfig` API.
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL kernel version, which must be the same as that of the primary instance
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Instance storage capacity in GB
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// Number of instances purchased at a time. Value range: 1–100.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Valid period in months of purchased instances. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This parameter is set to `1` when the pay-as-you-go billing mode is used.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// ID of the primary instance to which the read-only replica belongs
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// Availability zone ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Instance billing mode. Valid values: `PREPAID` (monthly subscription), `POSTPAID_BY_HOUR` (pay-as-you-go).
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Whether to automatically use vouchers. Valid values: `1` (yes), `0` (no). Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// Renewal flag. Valid values: `0` (manual renewal), `1` (auto-renewal). Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Special offer ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitempty" name:"Name"`

	// Whether to support IPv6 address access. Valid values: `1` (yes), `0` (no).
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// The information of tags to be associated with instances. This parameter is left empty by default.
	TagList *Tag `json:"TagList,omitempty" name:"TagList"`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecCode")
	delete(f, "DBVersion")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "MasterDBInstanceId")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "AutoRenewFlag")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "ReadOnlyGroupId")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateReadOnlyDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Order number list. Each instance corresponds to an order number.
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// Bill ID of frozen fees
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// ID set of instances which have been created successfully. The parameter value will be returned only when the pay-as-you-go billing mode is used.
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// Primary instance ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// RO group name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// Delay threshold in ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// Delayed log size threshold in MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// The minimum number of read-only replicas that must be retained in an RO group
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MasterDBInstanceId")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLag")
	delete(f, "MaxReplayLatency")
	delete(f, "MinDelayEliminateReserve")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RO group ID
		ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

		// Task ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Availability zone ID. Only ap-shanghai-2, ap-beijing-1, and ap-guangzhou-2 are supported during the beta test.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance name. The value must be unique for the same account.
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// Kernel version of a PostgreSQL instance. Currently, only 10.4 is supported.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Database character set of a PostgreSQL instance. Currently, only UTF-8 is supported.
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Array of tags to be bound with the instance
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBInstanceName")
	delete(f, "DBVersion")
	delete(f, "DBCharset")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "TagList")
	if len(f) > 0 {
		return errors.New("CreateServerlessDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance ID, such as "postgres-xxxxx". The value must be globally unique.
		DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceResponse) FromJsonString(s string) error {
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

	// SubnetId
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// Instance status
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// Assigned instance memory size in GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory,omitempty" name:"DBInstanceMemory"`

	// Assigned instance storage capacity in GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage,omitempty" name:"DBInstanceStorage"`

	// Number of assigned CPUs
	DBInstanceCpu *uint64 `json:"DBInstanceCpu,omitempty" name:"DBInstanceCpu"`

	// Purchasable specification ID
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" name:"DBInstanceClass"`

	// Instance type. 1: primary (master instance), 2: readonly (read-only instance), 3: guard (disaster recovery instance), 4: temp (temp instance)
	DBInstanceType *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`

	// Instance edition. Currently, only `standard` edition (dual-server high-availability one-master-one-slave edition) is supported
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

	// User `AppId`
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// Instance `Uid`
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// Whether the instance supports IPv6 address access. Valid values: 1 (yes), 0 (no)
	SupportIpv6 *uint64 `json:"SupportIpv6,omitempty" name:"SupportIpv6"`

	// The information of tags associated with instances.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`

	// Primary instance information, which is returned only when the instance is read-only
	// Note: this field may return null, indicating that no valid values can be obtained.
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// Number of read-only instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReadOnlyInstanceNum *int64 `json:"ReadOnlyInstanceNum,omitempty" name:"ReadOnlyInstanceNum"`

	// The status of a instance in a read-only group
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusInReadonlyGroup *string `json:"StatusInReadonlyGroup,omitempty" name:"StatusInReadonlyGroup"`

	// Elimination time
	// Note: this field may return null, indicating that no valid values can be obtained.
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
}

type DBInstanceNetInfo struct {

	// DNS domain name
	Address *string `json:"Address,omitempty" name:"Address"`

	// Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Connection port address
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Network type. 1: inner (private network address), 2: public (public network address)
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// Network connection status
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// ID of the RO group to be deleted
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("DeleteReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance name. Either instance name or instance ID (or both) must be passed in. If both are passed in, the instance ID will prevail.
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// Instance ID. Either instance name or instance ID (or both) must be passed in. If both are passed in, the instance ID will prevail.
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceName")
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DeleteServerlessDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Number of entries returned per page. Default value: 20. Value range: 1-100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Whether to sort by creation time or username. Valid values: `createTime` (sort by creation time), `name` (sort by username)
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Whether returns are sorted in ascending or descending order. Valid values: `desc` (descending), `asc` (ascending)
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeAccountsRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDBBackupsRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Number of entries returned per page. Value range: 1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "SearchKeys")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDBErrlogsRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDBInstanceAttributeRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Filter condition. Valid values: db-instance-id, db-instance-name, db-project-id, db-pay-mode, db-tag-key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Number of entries returned per page. Default value: 10.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Data offset which starts from 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting metric, such as instance name or creation time. Valid values: DBInstanceId, CreateTime, Name, EndTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// In ascending or descending order
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeDBInstancesRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Number of entries returned per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDBSlowlogsRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Number of entries returned per page in paged query. Value range: 1-100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeDBXlogsRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDatabasesRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// Order name set
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return errors.New("DescribeOrdersRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest

	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return errors.New("DescribeProductConfigRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReadOnlyGroupsRequest struct {
	*tchttp.BaseRequest

	// Filter condition. The primary ID must be specified in the format of `db-master-instance-id` to filter results, or else `null` will be returned.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// The number of results per page. Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Specify which page is displayed. Default value: 1 (the first page).
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// Sorting criterion. Valid values: `ROGroupId`, `CreateTime`, `Name`.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc`.
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeReadOnlyGroupsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReadOnlyGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// RO group list
		ReadOnlyGroupList []*ReadOnlyGroup `json:"ReadOnlyGroupList,omitempty" name:"ReadOnlyGroupList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeRegionsRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Query conditions
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`

	// The number of queries
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The offset value
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Sorting metric. Currently, only "CreateTime" (instance creation time) is supported.
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// Sorting order. Ascending and descending are supported.
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeServerlessDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The number of query results
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Query results
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		DBInstanceSet []*ServerlessDBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeZonesRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest

	// The ID of the instance to be eliminated
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DestroyDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Resource ID list
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// Specify the valid period (in months) of the monthly-subscribed instance when removing it from isolation
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Whether to use vouchers
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// Voucher ID list
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return errors.New("DisIsolateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisIsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesResponse) FromJsonString(s string) error {
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "Charset")
	if len(f) > 0 {
		return errors.New("InitDBInstancesRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Pid")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return errors.New("InquiryPriceCreateDBInstancesRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	if len(f) > 0 {
		return errors.New("InquiryPriceRenewDBInstanceRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Storage")
	delete(f, "Memory")
	delete(f, "DBInstanceId")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return errors.New("InquiryPriceUpgradeDBInstanceRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID set
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	if len(f) > 0 {
		return errors.New("IsolateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesResponse) FromJsonString(s string) error {
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Remark")
	if len(f) > 0 {
		return errors.New("ModifyAccountRemarkRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return errors.New("ModifyDBInstanceNameRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// ID of the RO group to which the read-only replica belongs
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// ID of the new RO group into which the read-only replica will move
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitempty" name:"NewReadOnlyGroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "NewReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("ModifyDBInstanceReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// TencentDB for PostgreSQL instance ID array
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// New project ID of TencentDB for PostgreSQL instance
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("ModifyDBInstancesProjectRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReadOnlyGroupConfigRequest struct {
	*tchttp.BaseRequest

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// RO group name
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// Delayed log size threshold in MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// Delay threshold in ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// Whether to enable automatic load balancing. Valid values: `0` (disable), `1` (enable).
	Rebalance *uint64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// The minimum number of read-only replicas that must be retained in an RO group
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLatency")
	delete(f, "MaxReplayLag")
	delete(f, "Rebalance")
	delete(f, "MinDelayEliminateReserve")
	if len(f) > 0 {
		return errors.New("ModifyReadOnlyGroupConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReadOnlyGroupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigResponse) FromJsonString(s string) error {
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return errors.New("OpenDBExtranetAccessRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// Unique ID of an instance
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return errors.New("OpenServerlessDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
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

type ReadOnlyGroup struct {

	// RO group identifier
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// RO group name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// Project ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Primary instance ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// The minimum number of read-only replicas that must be retained in an RO group
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MinDelayEliminateReserve *int64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// Delayed log size threshold
	MaxReplayLatency *int64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *int64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// Delay threshold
	MaxReplayLag *float64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *int64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Region ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// Availability zone ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// Instance details
	ReadOnlyDBInstanceList []*DBInstance `json:"ReadOnlyDBInstanceList,omitempty" name:"ReadOnlyDBInstanceList" list`

	// Whether to enable automatic load balancing
	Rebalance *int64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// Network information
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`
}

type RebalanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("RebalanceReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebalanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Whether the resource can be purchased in this region. Valid values: `0` (no), `1` (yes).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SupportInternational *uint64 `json:"SupportInternational,omitempty" name:"SupportInternational"`
}

type RemoveDBInstanceFromReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("RemoveDBInstanceFromReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveDBInstanceFromReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Task ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return errors.New("RenewInstanceRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return errors.New("ResetAccountPasswordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("RestartDBInstanceRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerlessDBAccount struct {

	// Username
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBUser *string `json:"DBUser,omitempty" name:"DBUser"`

	// Password
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBPassword *string `json:"DBPassword,omitempty" name:"DBPassword"`

	// The maximum number of connections
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBConnLimit *int64 `json:"DBConnLimit,omitempty" name:"DBConnLimit"`
}

type ServerlessDBInstance struct {

	// Instance ID, which is the unique identifier
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// Instance name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// Instance status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// Region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Availability zone
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Project ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// VPC ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Character set
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// Database version
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// Creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Instance network information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceNetInfo []*ServerlessDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`

	// Instance account information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBAccountSet []*ServerlessDBAccount `json:"DBAccountSet,omitempty" name:"DBAccountSet" list`

	// Information of the databases in an instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBDatabaseList []*string `json:"DBDatabaseList,omitempty" name:"DBDatabaseList" list`

	// The array of tags bound to an instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`
}

type ServerlessDBInstanceNetInfo struct {

	// Address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitempty" name:"Address"`

	// IP address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Port number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// Status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Network type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Instance ID array
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// Renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal upon expiration
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return errors.New("SetAutoRenewFlagRequest has unknown keys!")
	}
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

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

type Tag struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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

	// Activity ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "DBInstanceId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return errors.New("UpgradeDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Transaction name.
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// Bill ID of frozen fees
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// Backup file size
	Size *int64 `json:"Size,omitempty" name:"Size"`
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

	// Whether the AZ supports IPv6 address access
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitempty" name:"ZoneSupportIpv6"`
}
