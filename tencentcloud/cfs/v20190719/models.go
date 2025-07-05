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

package v20190719

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AutoScaleUpRule struct {

	Status *string `json:"Status,omitnil,omitempty" name:"Status"`


	ScaleThreshold *uint64 `json:"ScaleThreshold,omitnil,omitempty" name:"ScaleThreshold"`


	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`
}

type AutoSnapshotPolicyInfo struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// Snapshot policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Snapshot policy creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Number of bound file systems
	FileSystemNums *uint64 `json:"FileSystemNums,omitnil,omitempty" name:"FileSystemNums"`

	// The specific day of the week on which to create a snapshot. This parameter is mutually exclusive with `DayOfMonth` and `IntervalDays`.
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// The hour of a day at which to regularly back up the snapshot
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Whether to activate the scheduled snapshot feature
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// Next time to trigger snapshot
	NextActiveTime *string `json:"NextActiveTime,omitnil,omitempty" name:"NextActiveTime"`

	// Snapshot policy status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Account ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Retention period
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// Region
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// File system information
	FileSystems []*FileSystemByPolicy `json:"FileSystems,omitnil,omitempty" name:"FileSystems"`

	// The specific day of the month on which to create a snapshot. This parameter is mutually exclusive with `DayOfWeek` and `IntervalDays`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// The snapshot interval (1 to 365 days). This parameter is mutually exclusive with `DayOfWeek` and `DayOfMonth`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type AvailableProtoStatus struct {
	// Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable)
	SaleStatus *string `json:"SaleStatus,omitnil,omitempty" name:"SaleStatus"`

	// Protocol type. Valid values: NFS, CIFS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type AvailableRegion struct {
	// Region name, such as "ap-beijing"
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region name, such as "bj"
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Region availability. If a region has at least one AZ where resources are purchasable, this value will be `AVAILABLE`; otherwise, it will be `UNAVAILABLE`
	RegionStatus *string `json:"RegionStatus,omitnil,omitempty" name:"RegionStatus"`

	// Array of AZs
	Zones []*AvailableZone `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Region name, such as "Guangzhou"
	RegionCnName *string `json:"RegionCnName,omitnil,omitempty" name:"RegionCnName"`
}

type AvailableType struct {
	// Protocol and sale details
	Protocols []*AvailableProtoStatus `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// Storage class. Valid values: `SD` (standard storage) and `HP` (high-performance storage)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Indicates whether prepaid is supported. `true`: yes; `false`: no
	Prepayment *bool `json:"Prepayment,omitnil,omitempty" name:"Prepayment"`
}

type AvailableZone struct {
	// AZ name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Chinese name of an AZ
	ZoneCnName *string `json:"ZoneCnName,omitnil,omitempty" name:"ZoneCnName"`

	// Array of classes
	Types []*AvailableType `json:"Types,omitnil,omitempty" name:"Types"`

	// Chinese and English names of an AZ
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

// Predefined struct for user
type BindAutoSnapshotPolicyRequestParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// List of file systems
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// List of file systems
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`
}

func (r *BindAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "FileSystemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoSnapshotPolicyResponseParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *BindAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *BindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BucketInfo struct {
	// Bucket name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bucket region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type CreateAutoSnapshotPolicyRequestParams struct {
	// The time point when to repeat the snapshot operation
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// The day of the week on which to repeat the snapshot operation
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// Snapshot retention period
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// The specific day (day 1 to day 31) of the month on which to automatically create a snapshot.
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// The snapshot interval, in days.
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The time point when to repeat the snapshot operation
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// The day of the week on which to repeat the snapshot operation
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// Snapshot retention period
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// The specific day (day 1 to day 31) of the month on which to automatically create a snapshot.
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// The snapshot interval, in days.
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

func (r *CreateAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hour")
	delete(f, "PolicyName")
	delete(f, "DayOfWeek")
	delete(f, "AliveDays")
	delete(f, "DayOfMonth")
	delete(f, "IntervalDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyResponseParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *CreateAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsFileSystemRequestParams struct {
	// AZ name, such as "ap-beijing-1". For the list of regions and AZs, please see [Overview](https://intl.cloud.tencent.com/document/product/582/13225?from_cn_redirect=1)
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Network type. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system, and `CCN` for a Standard Turbo or High-Performance Turbo one.
	NetInterface *string `json:"NetInterface,omitnil,omitempty" name:"NetInterface"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP` (High-Performance Turbo). Default value: `SD`.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// VPC ID. This field is required if network type is VPC.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. This field is required if network type is VPC.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// IP address (this parameter supports only the VPC network type, and the Turbo series is not supported). If this parameter is left empty, a random IP in the subnet will be assigned.
	MountIP *string `json:"MountIP,omitnil,omitempty" name:"MountIP"`

	// Custom file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// File system tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. This string is valid for 2 hours.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// CCN instance ID (required if the network type is CCN)
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// AZ name, such as "ap-beijing-1". For the list of regions and AZs, please see [Overview](https://intl.cloud.tencent.com/document/product/582/13225?from_cn_redirect=1)
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Network type. Valid values: `VPC` and `CCN`. Select `VPC` for a Standard or High-Performance file system, and `CCN` for a Standard Turbo or High-Performance Turbo one.
	NetInterface *string `json:"NetInterface,omitnil,omitempty" name:"NetInterface"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Storage type of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), and `TP` (High-Performance Turbo). Default value: `SD`.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// VPC ID. This field is required if network type is VPC.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. This field is required if network type is VPC.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// IP address (this parameter supports only the VPC network type, and the Turbo series is not supported). If this parameter is left empty, a random IP in the subnet will be assigned.
	MountIP *string `json:"MountIP,omitnil,omitempty" name:"MountIP"`

	// Custom file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// File system tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. This string is valid for 2 hours.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// CCN instance ID (required if the network type is CCN)
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

func (r *CreateCfsFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "NetInterface")
	delete(f, "PGroupId")
	delete(f, "Protocol")
	delete(f, "StorageType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "MountIP")
	delete(f, "FsName")
	delete(f, "ResourceTags")
	delete(f, "ClientToken")
	delete(f, "CcnId")
	delete(f, "CidrBlock")
	delete(f, "Capacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsFileSystemResponseParams struct {
	// File system creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Custom file system name
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// File system status. Valid values: `creating`, `create_failed`, `available`, `unserviced`, `upgrading`, `deleting`
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// Storage used by the file system, in bytes
	SizeByte *uint64 `json:"SizeByte,omitnil,omitempty" name:"SizeByte"`

	// AZ ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Custom file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// Whether a file system is encrypted
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsFileSystemResponseParams `json:"Response"`
}

func (r *CreateCfsFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsPGroupRequestParams struct {
	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

type CreateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

func (r *CreateCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DescInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsPGroupResponseParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Permission group description
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`

	// The number of file systems bound to this permission group
	BindCfsNum *int64 `json:"BindCfsNum,omitnil,omitempty" name:"BindCfsNum"`

	// Permission group creation time
	CDate *string `json:"CDate,omitnil,omitempty" name:"CDate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsPGroupResponseParams `json:"Response"`
}

func (r *CreateCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsRuleRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`
}

func (r *CreateCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "AuthClientIp")
	delete(f, "Priority")
	delete(f, "RWPermission")
	delete(f, "UserPermission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Client IP
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Read & write permissions
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// Priority
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsRuleResponseParams `json:"Response"`
}

func (r *CreateCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsSnapshotRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// Snapshot tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type CreateCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// Snapshot tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

func (r *CreateCfsSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SnapshotName")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsSnapshotResponseParams struct {
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsSnapshotResponseParams `json:"Response"`
}

func (r *CreateCfsSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationTaskRequestParams struct {
	// Migration task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Migration type. Valid values: `0` (bucket) and `1` (list). Default value: `0`.
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// Migration mode. Default value: `0` (full migration).
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// SecretId of the data source account
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// SecretKey of the data source account
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`

	// File system instance ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// File system path
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// Overwrite policy for files with the same name. Valid values: `0` (retain the file with the latest modified time), `1` (overwrite); and `2` (not overwrite). Default value: `0`.
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// Data source service provider. Valid values: `COS` (Tencent Cloud COS), `OSS` (Alibaba Cloud OSS), and `OBS` (Huawei Cloud OBS).
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// Data source bucket name. Specify at least one of the bucket name or address.
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// Data source bucket region
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// Data source bucket address. Specify at least one of the bucket name or address.
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// List address. This parameter is required if `MigrationType` is set to `1` (list).
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// Target file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// Source bucket path, which defaults to `/`
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

type CreateMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// Migration task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Migration type. Valid values: `0` (bucket) and `1` (list). Default value: `0`.
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// Migration mode. Default value: `0` (full migration).
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// SecretId of the data source account
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// SecretKey of the data source account
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`

	// File system instance ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// File system path
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// Overwrite policy for files with the same name. Valid values: `0` (retain the file with the latest modified time), `1` (overwrite); and `2` (not overwrite). Default value: `0`.
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// Data source service provider. Valid values: `COS` (Tencent Cloud COS), `OSS` (Alibaba Cloud OSS), and `OBS` (Huawei Cloud OBS).
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// Data source bucket name. Specify at least one of the bucket name or address.
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// Data source bucket region
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// Data source bucket address. Specify at least one of the bucket name or address.
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// List address. This parameter is required if `MigrationType` is set to `1` (list).
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// Target file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// Source bucket path, which defaults to `/`
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

func (r *CreateMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "MigrationType")
	delete(f, "MigrationMode")
	delete(f, "SrcSecretId")
	delete(f, "SrcSecretKey")
	delete(f, "FileSystemId")
	delete(f, "FsPath")
	delete(f, "CoverType")
	delete(f, "SrcService")
	delete(f, "BucketName")
	delete(f, "BucketRegion")
	delete(f, "BucketAddress")
	delete(f, "ListAddress")
	delete(f, "FsName")
	delete(f, "BucketPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationTaskResponseParams struct {
	// Migration task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrationTaskResponseParams `json:"Response"`
}

func (r *CreateMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPolicyRequestParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type DeleteAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

func (r *DeleteAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPolicyResponseParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *DeleteAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsFileSystemRequestParams struct {
	// File system ID. Note: please call the `DeleteMountTarget` API to delete the mount target first before deleting a file system; otherwise, the delete operation will fail.
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system ID. Note: please call the `DeleteMountTarget` API to delete the mount target first before deleting a file system; otherwise, the delete operation will fail.
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DeleteCfsFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsFileSystemResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsFileSystemResponseParams `json:"Response"`
}

func (r *DeleteCfsFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsPGroupRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

type DeleteCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

func (r *DeleteCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsPGroupResponseParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// User ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsPGroupResponseParams `json:"Response"`
}

func (r *DeleteCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsRuleRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsRuleResponseParams struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsRuleResponseParams `json:"Response"`
}

func (r *DeleteCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsSnapshotRequestParams struct {
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The list of the IDs of the file system snapshots to be deleted. At least one of `SnapshotId` and `SnapshotIds` must be specified.
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

type DeleteCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The list of the IDs of the file system snapshots to be deleted. At least one of `SnapshotId` and `SnapshotIds` must be specified.
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

func (r *DeleteCfsSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsSnapshotResponseParams struct {
	// File system ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsSnapshotResponseParams `json:"Response"`
}

func (r *DeleteCfsSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrationTaskRequestParams struct {
	// Migration task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// Migration task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrationTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMigrationTaskResponseParams `json:"Response"`
}

func (r *DeleteMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountTargetRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitnil,omitempty" name:"MountTargetId"`
}

type DeleteMountTargetRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitnil,omitempty" name:"MountTargetId"`
}

func (r *DeleteMountTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "MountTargetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMountTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountTargetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMountTargetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMountTargetResponseParams `json:"Response"`
}

func (r *DeleteMountTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesRequestParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page length
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Ascending or descending order
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// Page offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page length
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Ascending or descending order
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Sorting field
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeAutoSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesResponseParams struct {
	// Total number of snapshot policies
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Snapshot policy information
	AutoSnapshotPolicies []*AutoSnapshotPolicyInfo `json:"AutoSnapshotPolicies,omitnil,omitempty" name:"AutoSnapshotPolicies"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableZoneInfoRequestParams struct {

}

type DescribeAvailableZoneInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailableZoneInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableZoneInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableZoneInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableZoneInfoResponseParams struct {
	// Information such as resource availability in each AZ and the supported storage classes and protocols
	RegionZones []*AvailableRegion `json:"RegionZones,omitnil,omitempty" name:"RegionZones"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableZoneInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableZoneInfoResponseParams `json:"Response"`
}

func (r *DescribeAvailableZoneInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableZoneInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBucketListRequestParams struct {
	// Data source service provider. Valid values: `COS` (Tencent Cloud COS), `OSS` (Alibaba Cloud OSS), and `OBS` (Huawei Cloud OBS).
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// SecretId of the data source account
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// SecretKey of the data source account
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`
}

type DescribeBucketListRequest struct {
	*tchttp.BaseRequest
	
	// Data source service provider. Valid values: `COS` (Tencent Cloud COS), `OSS` (Alibaba Cloud OSS), and `OBS` (Huawei Cloud OBS).
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// SecretId of the data source account
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// SecretKey of the data source account
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`
}

func (r *DescribeBucketListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBucketListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcService")
	delete(f, "SrcSecretId")
	delete(f, "SrcSecretKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBucketListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBucketListResponseParams struct {
	// Number of buckets
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Bucket list
	BucketList []*BucketInfo `json:"BucketList,omitnil,omitempty" name:"BucketList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBucketListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBucketListResponseParams `json:"Response"`
}

func (r *DescribeBucketListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBucketListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemClientsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeCfsFileSystemClientsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeCfsFileSystemClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemClientsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemClientsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemClientsResponseParams struct {
	// Client list
	ClientList []*FileSystemClient `json:"ClientList,omitnil,omitempty" name:"ClientList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsFileSystemClientsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsFileSystemClientsResponseParams `json:"Response"`
}

func (r *DescribeCfsFileSystemClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`


	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`


	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`


	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`
}

type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`
}

func (r *DescribeCfsFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemsResponseParams struct {
	// File system information
	FileSystems []*FileSystemInfo `json:"FileSystems,omitnil,omitempty" name:"FileSystems"`

	// Total number of file systems
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsFileSystemsResponseParams `json:"Response"`
}

func (r *DescribeCfsFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsPGroupsRequestParams struct {

}

type DescribeCfsPGroupsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfsPGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsPGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsPGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsPGroupsResponseParams struct {
	// Permission group information list
	PGroupList []*PGroupInfo `json:"PGroupList,omitnil,omitempty" name:"PGroupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsPGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsPGroupsResponseParams `json:"Response"`
}

func (r *DescribeCfsPGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsPGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsRulesRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

type DescribeCfsRulesRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

func (r *DescribeCfsRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsRulesResponseParams struct {
	// List of permission group rules
	RuleList []*PGroupRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsRulesResponseParams `json:"Response"`
}

func (r *DescribeCfsRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsServiceStatusRequestParams struct {

}

type DescribeCfsServiceStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfsServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsServiceStatusResponseParams struct {
	// Current status of the CFS service for this user. Valid values: none (not activated), creating (activating), created (activated)
	CfsServiceStatus *string `json:"CfsServiceStatus,omitnil,omitempty" name:"CfsServiceStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsServiceStatusResponseParams `json:"Response"`
}

func (r *DescribeCfsServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotOverviewRequestParams struct {

}

type DescribeCfsSnapshotOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfsSnapshotOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsSnapshotOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotOverviewResponseParams struct {
	// Statistics
	StatisticsList []*SnapshotStatistics `json:"StatisticsList,omitnil,omitempty" name:"StatisticsList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsSnapshotOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsSnapshotOverviewResponseParams `json:"Response"`
}

func (r *DescribeCfsSnapshotOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The starting position of paging
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page length
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Order field
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Sorting order (ascending or descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeCfsSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The starting position of paging
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Page length
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Order field
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// Sorting order (ascending or descending)
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeCfsSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SnapshotId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotsResponseParams struct {
	// Total number
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Snapshot information description
	Snapshots []*SnapshotInfo `json:"Snapshots,omitnil,omitempty" name:"Snapshots"`

	// Total size of snapshots
	TotalSize *uint64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeCfsSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTasksRequestParams struct {
	// The pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <br><li> taskId
	// 
	// Filter by **migration task ID**
	// Type: String
	// 
	// Required: No
	// 
	// <br><li> taskName
	// 
	// Fuzzy filter by **migration task name**
	// Type: String
	// 
	// Required: No
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMigrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// The pagination offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of entries per page. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <br><li> taskId
	// 
	// Filter by **migration task ID**
	// Type: String
	// 
	// Required: No
	// 
	// <br><li> taskName
	// 
	// Fuzzy filter by **migration task name**
	// Type: String
	// 
	// Required: No
	// 
	// Each request can have up to 10 `Filters` and 100 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMigrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTasksResponseParams struct {
	// Number of migration tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Migration task details
	MigrationTasks []*MigrationTaskInfo `json:"MigrationTasks,omitnil,omitempty" name:"MigrationTasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationTasksResponseParams `json:"Response"`
}

func (r *DescribeMigrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountTargetsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeMountTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountTargetsResponseParams struct {
	// Mount target details
	MountTargets []*MountInfo `json:"MountTargets,omitnil,omitempty" name:"MountTargets"`

	// The number of mount targets
	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitnil,omitempty" name:"NumberOfMountTargets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountTargetsResponseParams `json:"Response"`
}

func (r *DescribeMountTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOperationLogsRequestParams struct {
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeSnapshotOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeSnapshotOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOperationLogsResponseParams struct {
	// Snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Operation log
	SnapshotOperates []*SnapshotOperateLog `json:"SnapshotOperates,omitnil,omitempty" name:"SnapshotOperates"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystemByPolicy struct {
	// File system name
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// File system size
	SizeByte *uint64 `json:"SizeByte,omitnil,omitempty" name:"SizeByte"`

	// Storage class
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Total snapshot size
	TotalSnapshotSize *uint64 `json:"TotalSnapshotSize,omitnil,omitempty" name:"TotalSnapshotSize"`

	// File system creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Region ID of the file system
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type FileSystemClient struct {
	// IP address of the file system
	CfsVip *string `json:"CfsVip,omitnil,omitempty" name:"CfsVip"`

	// Client IP
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// File system VPCID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Name of the availability zone, e.g. ap-beijing-1. For more information, see regions and availability zones in the Overview document
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Path in which the file system is mounted to the client
	MountDirectory *string `json:"MountDirectory,omitnil,omitempty" name:"MountDirectory"`
}

type FileSystemInfo struct {
	// Creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Custom name
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// File system status. Valid values:
	// - creating
	// - mounting
	// - create_failed
	// - available
	// - unserviced
	// - upgrading
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// Used file system capacity
	SizeByte *uint64 `json:"SizeByte,omitnil,omitempty" name:"SizeByte"`

	// Maximum storage limit of a file system
	SizeLimit *uint64 `json:"SizeLimit,omitnil,omitempty" name:"SizeLimit"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Region name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// File system protocol type
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// File system storage class
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Prepaid storage pack bound with the file system
	StorageResourcePkg *string `json:"StorageResourcePkg,omitnil,omitempty" name:"StorageResourcePkg"`

	// Prepaid bandwidth pack bound to a file system (not supported currently)
	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitnil,omitempty" name:"BandwidthResourcePkg"`

	// Information of permission groups bound to a file system
	PGroup *PGroup `json:"PGroup,omitnil,omitempty" name:"PGroup"`

	// Custom name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// Whether a file system is encrypted
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// Key used for encryption, which can be the key ID or ARN
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// Application ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The upper limit on the file system’s throughput, which is determined based on its current usage, and bound resource packs for both storage and throughput
	BandwidthLimit *float64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`


	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`


	SnapStatus *string `json:"SnapStatus,omitnil,omitempty" name:"SnapStatus"`

	// Total capacity of the file system
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// File system tag list
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// The lifecycle management status of a file system.
	TieringState *string `json:"TieringState,omitnil,omitempty" name:"TieringState"`

	// The details about tiered storage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TieringDetail *TieringDetailInfo `json:"TieringDetail,omitnil,omitempty" name:"TieringDetail"`


	AutoScaleUpRule *AutoScaleUpRule `json:"AutoScaleUpRule,omitnil,omitempty" name:"AutoScaleUpRule"`
}

type Filter struct {
	// Value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type MigrationTaskInfo struct {
	// Migration task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Migration task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Migration type. Valid values: `0` (bucket) and `1` (list). Default value: `0`.
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// Migration mode. Default value: `0` (full migration).
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// Data source bucket name
	// Note: This field may return null, indicating that no valid values can be obtained.
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// Data source bucket region
	// Note: This field may return null, indicating that no valid values can be obtained.
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// Data source bucket address
	// Note: This field may return null, indicating that no valid values can be obtained.
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// List address
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// File system instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// File system instance ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// File system path
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// Overwrite policy for files with the same name. Valid values: `0` (retain the file with the latest modified time), `1` (overwrite); and `2` (not overwrite). Default value: `0`.
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// Creation time
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// End time
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Migration status. Valid values: `0` (completed), `1` (in progress), and `2` (stopped).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of files
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileTotalCount *uint64 `json:"FileTotalCount,omitnil,omitempty" name:"FileTotalCount"`

	// Number of migrated files
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileMigratedCount *uint64 `json:"FileMigratedCount,omitnil,omitempty" name:"FileMigratedCount"`

	// Number of files that failed to be migrated
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileFailedCount *uint64 `json:"FileFailedCount,omitnil,omitempty" name:"FileFailedCount"`

	// File size, in bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileTotalSize *int64 `json:"FileTotalSize,omitnil,omitempty" name:"FileTotalSize"`

	// Size of migrated files, in bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileMigratedSize *int64 `json:"FileMigratedSize,omitnil,omitempty" name:"FileMigratedSize"`

	// Size of files that failed to be migrated, in bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileFailedSize *int64 `json:"FileFailedSize,omitnil,omitempty" name:"FileFailedSize"`

	// List of all files
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileTotalList *string `json:"FileTotalList,omitnil,omitempty" name:"FileTotalList"`

	// List of migrated files
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileCompletedList *string `json:"FileCompletedList,omitnil,omitempty" name:"FileCompletedList"`

	// List of files that failed to be migrated
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileFailedList *string `json:"FileFailedList,omitnil,omitempty" name:"FileFailedList"`

	// Source bucket path
	// Note: This field may return null, indicating that no valid values can be obtained.
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

// Predefined struct for user
type ModifyFileSystemAutoScaleUpRuleRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Threshold for triggering scaling. Value range: 10-90
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// Target threshold after scaling. Value range: 10-90. The value of this parameter must be smaller than that of `ScaleUpThreshold`.
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// Rule status. Valid values: `0` (disabled) and `1` (enabled).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyFileSystemAutoScaleUpRuleRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Threshold for triggering scaling. Value range: 10-90
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// Target threshold after scaling. Value range: 10-90. The value of this parameter must be smaller than that of `ScaleUpThreshold`.
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// Rule status. Valid values: `0` (disabled) and `1` (enabled).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyFileSystemAutoScaleUpRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemAutoScaleUpRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "ScaleUpThreshold")
	delete(f, "TargetThreshold")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFileSystemAutoScaleUpRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemAutoScaleUpRuleResponseParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Rule status. Valid values: `0` (disabled) and `1` (enabled).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Threshold for triggering scaling. Value range: 10-90
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// Target threshold after scaling. Value range: 10-90
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFileSystemAutoScaleUpRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFileSystemAutoScaleUpRuleResponseParams `json:"Response"`
}

func (r *ModifyFileSystemAutoScaleUpRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemAutoScaleUpRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountInfo struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitnil,omitempty" name:"MountTargetId"`

	// Mount target IP
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// Mount root-directory
	FSID *string `json:"FSID,omitnil,omitempty" name:"FSID"`

	// Mount target status
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// Network type
	NetworkInterface *string `json:"NetworkInterface,omitnil,omitempty" name:"NetworkInterface"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Subnet name
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// CCN instance ID used by CFS Turbo
	CcnID *string `json:"CcnID,omitnil,omitempty" name:"CcnID"`

	// CCN IP range used by CFS Turbo
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`
}

type PGroup struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type PGroupInfo struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`

	// Creation time
	CDate *string `json:"CDate,omitnil,omitempty" name:"CDate"`

	// The number of bound file system
	BindCfsNum *int64 `json:"BindCfsNum,omitnil,omitempty" name:"BindCfsNum"`
}

type PGroupRuleInfo struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Client IP allowed for access
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Read/write permission. ro: read-only; rw: read & write
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission. all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions.
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

// Predefined struct for user
type ScaleUpFileSystemRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Target capacity after scaling
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`
}

type ScaleUpFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Target capacity after scaling
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`
}

func (r *ScaleUpFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "TargetCapacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpFileSystemResponseParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Target capacity after scaling
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpFileSystemResponseParams `json:"Response"`
}

func (r *ScaleUpFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignUpCfsServiceRequestParams struct {

}

type SignUpCfsServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SignUpCfsServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignUpCfsServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignUpCfsServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignUpCfsServiceResponseParams struct {
	// Current status of the CFS service for this user. Valid values: `creating` (activating); `created` (activated)
	CfsServiceStatus *string `json:"CfsServiceStatus,omitnil,omitempty" name:"CfsServiceStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SignUpCfsServiceResponse struct {
	*tchttp.BaseResponse
	Response *SignUpCfsServiceResponseParams `json:"Response"`
}

func (r *SignUpCfsServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignUpCfsServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotInfo struct {
	// Snapshot creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// Snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// Snapshot status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Snapshot size
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Retention period in days
	AliveDay *uint64 `json:"AliveDay,omitnil,omitempty" name:"AliveDay"`

	// Snapshot progress
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// Account ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Snapshot deletion time
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// File system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// Snapshot tag
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Snapshot type
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapshotType *string `json:"SnapshotType,omitnil,omitempty" name:"SnapshotType"`
}

type SnapshotOperateLog struct {
	// Operation type
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Operation time
	ActionTime *string `json:"ActionTime,omitnil,omitempty" name:"ActionTime"`

	// Operation name
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// Operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Result
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`
}

type SnapshotStatistics struct {
	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Total number of snapshots
	SnapshotNumber *uint64 `json:"SnapshotNumber,omitnil,omitempty" name:"SnapshotNumber"`

	// Total snapshot size
	SnapshotSize *uint64 `json:"SnapshotSize,omitnil,omitempty" name:"SnapshotSize"`
}

// Predefined struct for user
type StopMigrationTaskRequestParams struct {
	// Migration task name
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// Migration task name
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrationTaskResponseParams struct {
	// Migration task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Migration status. Valid values: `0` (completed), `1` (in progress), and `2` (stopped).
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrationTaskResponseParams `json:"Response"`
}

func (r *StopMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TieringDetailInfo struct {
	// STANDARD_IA storage usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	TieringSizeInBytes *int64 `json:"TieringSizeInBytes,omitnil,omitempty" name:"TieringSizeInBytes"`
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyRequestParams struct {
	// List of IDs of the file systems to be unbound, separated by comma
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`

	// ID of the snapshot to be unbound
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the file systems to be unbound, separated by comma
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`

	// ID of the snapshot to be unbound
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

func (r *UnbindAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemIds")
	delete(f, "AutoSnapshotPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyResponseParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UnbindAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *UnbindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAutoSnapshotPolicyRequestParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// Snapshot policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// The day of the week on which to regularly back up the snapshot
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// The hour of a day at which to regularly back up the snapshot
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Snapshot retention period
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// Whether to activate the scheduled snapshot feature
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// The specific day of the month on which to create a snapshot. This parameter is mutually exclusive with `DayOfWeek`.
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// The snapshot interval. This parameter is mutually exclusive with `DayOfWeek` and `DayOfMonth`.
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type UpdateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// Snapshot policy name
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// The day of the week on which to regularly back up the snapshot
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// The hour of a day at which to regularly back up the snapshot
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// Snapshot retention period
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// Whether to activate the scheduled snapshot feature
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// The specific day of the month on which to create a snapshot. This parameter is mutually exclusive with `DayOfWeek`.
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// The snapshot interval. This parameter is mutually exclusive with `DayOfWeek` and `DayOfMonth`.
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

func (r *UpdateAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "PolicyName")
	delete(f, "DayOfWeek")
	delete(f, "Hour")
	delete(f, "AliveDays")
	delete(f, "IsActivated")
	delete(f, "DayOfMonth")
	delete(f, "IntervalDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAutoSnapshotPolicyResponseParams struct {
	// Snapshot policy ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *UpdateAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemNameRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`
}

func (r *UpdateCfsFileSystemNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FsName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsFileSystemNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemNameResponseParams struct {
	// Custom file system name
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsFileSystemNameResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsFileSystemNameResponseParams `json:"Response"`
}

func (r *UpdateCfsFileSystemNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemPGroupRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsFileSystemPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemPGroupResponseParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsFileSystemPGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsFileSystemPGroupResponseParams `json:"Response"`
}

func (r *UpdateCfsFileSystemPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemSizeLimitRequestParams struct {
	// File system capacity limit in GB. Value range: 0-1,073,741,824. If 0 is entered, no limit will be imposed on the file system capacity.
	FsLimit *uint64 `json:"FsLimit,omitnil,omitempty" name:"FsLimit"`

	// File system ID. Currently, only Standard file systems are supported.
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemSizeLimitRequest struct {
	*tchttp.BaseRequest
	
	// File system capacity limit in GB. Value range: 0-1,073,741,824. If 0 is entered, no limit will be imposed on the file system capacity.
	FsLimit *uint64 `json:"FsLimit,omitnil,omitempty" name:"FsLimit"`

	// File system ID. Currently, only Standard file systems are supported.
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemSizeLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemSizeLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FsLimit")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsFileSystemSizeLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemSizeLimitResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsFileSystemSizeLimitResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsFileSystemSizeLimitResponseParams `json:"Response"`
}

func (r *UpdateCfsFileSystemSizeLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemSizeLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsPGroupRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

type UpdateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

func (r *UpdateCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "Name")
	delete(f, "DescInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsPGroupResponseParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsPGroupResponseParams `json:"Response"`
}

func (r *UpdateCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsRuleRequestParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type UpdateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

func (r *UpdateCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "RuleId")
	delete(f, "AuthClientIp")
	delete(f, "RWPermission")
	delete(f, "UserPermission")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsRuleResponseParams struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Client IP or IP range allowed for access
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// Read & write permission
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// User permission
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// Priority
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsRuleResponseParams `json:"Response"`
}

func (r *UpdateCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsSnapshotAttributeRequestParams struct {
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// File system snapshot name
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// File system snapshot retention period in days
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`
}

type UpdateCfsSnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// File system snapshot name
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// File system snapshot retention period in days
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`
}

func (r *UpdateCfsSnapshotAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsSnapshotAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	delete(f, "AliveDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsSnapshotAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsSnapshotAttributeResponseParams struct {
	// File system snapshot ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsSnapshotAttributeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsSnapshotAttributeResponseParams `json:"Response"`
}

func (r *UpdateCfsSnapshotAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsSnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}