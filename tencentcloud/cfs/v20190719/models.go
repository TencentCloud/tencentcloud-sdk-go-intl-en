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

package v20190719

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AvailableProtoStatus struct {
	// Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable)
	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`

	// Protocol type. Valid values: NFS, CIFS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type AvailableRegion struct {
	// Region name, such as "ap-beijing"
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region name, such as "bj"
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region availability. If a region has at least one AZ where resources are purchasable, this value will be `AVAILABLE`; otherwise, it will be `UNAVAILABLE`
	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`

	// Array of AZs
	Zones []*AvailableZone `json:"Zones,omitempty" name:"Zones"`

	// Region name, such as "Guangzhou"
	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type AvailableType struct {
	// Protocol and sale details
	Protocols []*AvailableProtoStatus `json:"Protocols,omitempty" name:"Protocols"`

	// Storage class. Valid values: `SD` (standard storage) and `HP` (high-performance storage)
	Type *string `json:"Type,omitempty" name:"Type"`

	// Indicates whether prepaid is supported. `true`: yes; `false`: no
	Prepayment *bool `json:"Prepayment,omitempty" name:"Prepayment"`
}

type AvailableZone struct {
	// AZ name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Chinese name of an AZ
	ZoneCnName *string `json:"ZoneCnName,omitempty" name:"ZoneCnName"`

	// Array of classes
	Types []*AvailableType `json:"Types,omitempty" name:"Types"`

	// Chinese and English names of an AZ
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

// Predefined struct for user
type CreateCfsFileSystemRequestParams struct {
	// AZ name, such as "ap-beijing-1". For the list of regions and AZs, please see [Overview](https://intl.cloud.tencent.com/document/product/582/13225?from_cn_redirect=1)
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Network type. Valid values: `VPC` (private network), `BASIC` (classic network), `CCN` (Cloud Connect Network). You must set this parameter to `CCN` if you use the Turbo series. Classic network will be phased out and is not recommended.
	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`

	// Permission group ID (required for Standard and High-Performance). For the Turbo series, set it to `pgroupbasic`.
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Storage class of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), `TP` (High-Performance Turbo)
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// VPC ID. This field is required if network type is VPC.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID. This field is required if network type is VPC.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// IP address (this parameter supports only the VPC network type, and the Turbo series is not supported). If this parameter is left empty, a random IP in the subnet will be assigned.
	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// File system tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. This string is valid for 2 hours.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// CCN instance ID (required if the network type is CCN)
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// AZ name, such as "ap-beijing-1". For the list of regions and AZs, please see [Overview](https://intl.cloud.tencent.com/document/product/582/13225?from_cn_redirect=1)
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Network type. Valid values: `VPC` (private network), `BASIC` (classic network), `CCN` (Cloud Connect Network). You must set this parameter to `CCN` if you use the Turbo series. Classic network will be phased out and is not recommended.
	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`

	// Permission group ID (required for Standard and High-Performance). For the Turbo series, set it to `pgroupbasic`.
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system protocol. Valid values: `NFS`, `CIFS`, `TURBO`. If this parameter is left empty, `NFS` is used by default. For the Turbo series, you must set this parameter to `TURBO`.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Storage class of the file system. Valid values: `SD` (Standard), `HP` (High-Performance), `TB` (Standard Turbo), `TP` (High-Performance Turbo)
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// VPC ID. This field is required if network type is VPC.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID. This field is required if network type is VPC.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// IP address (this parameter supports only the VPC network type, and the Turbo series is not supported). If this parameter is left empty, a random IP in the subnet will be assigned.
	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// File system tag
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. This string is valid for 2 hours.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// CCN instance ID (required if the network type is CCN)
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// CCN IP range used by the CFS (required if the network type is CCN), which cannot conflict with other IP ranges bound in CCN
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// File system capacity, in GiB (required for the Turbo series). For Standard Turbo, the minimum purchase required is 40,960 GiB (40 TiB) and the expansion increment is 20,480 GiB (20 TiB). For High-Performance Turbo, the minimum purchase required is 20,480 GiB (20 TiB) and the expansion increment is 10,240 GiB (10 TiB).
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
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
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Custom file system name
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// File system status. Valid values: `creating`, `create_failed`, `available`, `unserviced`, `upgrading`, `deleting`
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// Storage used by the file system, in bytes
	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// AZ ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// Whether a file system is encrypted
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

type CreateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// The number of file systems bound to this permission group
	BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`

	// Permission group creation time
	CDate *string `json:"CDate,omitempty" name:"CDate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
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
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Client IP
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read & write permissions
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteCfsFileSystemRequestParams struct {
	// File system ID. Note: please call the `DeleteMountTarget` API to delete the mount target first before deleting a file system; otherwise, the delete operation will fail.
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system ID. Note: please call the `DeleteMountTarget` API to delete the mount target first before deleting a file system; otherwise, the delete operation will fail.
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
}

type DeleteCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// User ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
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
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteMountTargetRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
}

type DeleteMountTargetRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RegionZones []*AvailableRegion `json:"RegionZones,omitempty" name:"RegionZones"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCfsFileSystemClientsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeCfsFileSystemClientsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	ClientList []*FileSystemClient `json:"ClientList,omitempty" name:"ClientList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemsResponseParams struct {
	// File system information
	FileSystems []*FileSystemInfo `json:"FileSystems,omitempty" name:"FileSystems"`

	// Total number of file systems
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupList []*PGroupInfo `json:"PGroupList,omitempty" name:"PGroupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
}

type DescribeCfsRulesRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
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
	RuleList []*PGroupRuleInfo `json:"RuleList,omitempty" name:"RuleList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeMountTargetsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	MountTargets []*MountInfo `json:"MountTargets,omitempty" name:"MountTargets"`

	// The number of mount targets
	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type FileSystemClient struct {
	// IP address of the file system
	CfsVip *string `json:"CfsVip,omitempty" name:"CfsVip"`

	// Client IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// File system VPCID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Name of the availability zone, e.g. ap-beijing-1. For more information, see regions and availability zones in the Overview document
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Path in which the file system is mounted to the client
	MountDirectory *string `json:"MountDirectory,omitempty" name:"MountDirectory"`
}

type FileSystemInfo struct {
	// Creation time
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Custom name
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// File system status
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// Used file system capacity
	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// Maximum storage limit of a file system
	SizeLimit *uint64 `json:"SizeLimit,omitempty" name:"SizeLimit"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// Region name
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// File system protocol type
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// File system storage class
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// Prepaid storage pack bound with the file system
	StorageResourcePkg *string `json:"StorageResourcePkg,omitempty" name:"StorageResourcePkg"`

	// Prepaid bandwidth pack bound to a file system (not supported currently)
	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitempty" name:"BandwidthResourcePkg"`

	// Information of permission groups bound to a file system
	PGroup *PGroup `json:"PGroup,omitempty" name:"PGroup"`

	// Custom name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// Whether a file system is encrypted
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`

	// Key used for encryption, which can be the key ID or ARN
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// Application ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// The upper limit on the file system’s throughput, which is determined based on its current usage, and bound resource packs for both storage and throughput
	BandwidthLimit *float64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`

	// Total capacity of the file system
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`

	// File system tag list
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
}

type MountInfo struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount target ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`

	// Mount target IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// Mount root-directory
	FSID *string `json:"FSID,omitempty" name:"FSID"`

	// Mount target status
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// Network type
	NetworkInterface *string `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Subnet name
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// CCN instance ID used by CFS Turbo
	CcnID *string `json:"CcnID,omitempty" name:"CcnID"`

	// CCN IP range used by CFS Turbo
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
}

type PGroup struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PGroupInfo struct {
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// Creation time
	CDate *string `json:"CDate,omitempty" name:"CDate"`

	// The number of bound file system
	BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`
}

type PGroupRuleInfo struct {
	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Client IP allowed for access
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read/write permission. ro: read-only; rw: read & write
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
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
	// Current status of the CFS service for this user. Valid values: none (not activated), creating (activating), created (activated)
	CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpdateCfsFileSystemNameRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`
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
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Custom file system name
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FsLimit *uint64 `json:"FsLimit,omitempty" name:"FsLimit"`

	// File system ID. Currently, only Standard file systems are supported.
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemSizeLimitRequest struct {
	*tchttp.BaseRequest
	
	// File system capacity limit in GB. Value range: 0-1,073,741,824. If 0 is entered, no limit will be imposed on the file system capacity.
	FsLimit *uint64 `json:"FsLimit,omitempty" name:"FsLimit"`

	// File system ID. Currently, only Standard file systems are supported.
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

type UpdateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name, which can contain 1-64 Chinese characters, letters, numbers, underscores, or dashes
	Name *string `json:"Name,omitempty" name:"Name"`

	// Permission group description, which can contain 1-255 characters
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Permission group name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Description
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type UpdateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// You can enter a single IP or IP range, such as 10.1.10.11 or 10.10.1.0/24. The default visiting address is `*`, indicating that all IPs are allowed. Please note that you need to enter the CVM instance's private IP here.
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read/write permission. Valid values: RO (read-only), RW (read & write). Default value: RO
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission. Valid values: all_squash, no_all_squash, root_squash, no_root_squash. Specifically, all_squash: any visiting user will be mapped to an anonymous user or user group; no_all_squash: a visiting user will be first matched with a local user, and if the match fails, it will be mapped to an anonymous user or user group; root_squash: a visiting root user will be mapped to an anonymous user or user group; no_root_squash: a visiting root user will be allowed to maintain root account permissions. Default value: root_squash.
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Rule priority. Value range: 1-100. 1 represents the highest priority, while 100 the lowest
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// Rule ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Client IP or IP range allowed for access
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// Read & write permission
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// User permission
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// Priority
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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