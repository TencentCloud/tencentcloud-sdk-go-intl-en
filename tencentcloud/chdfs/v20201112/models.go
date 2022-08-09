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

package v20201112

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessGroup struct {
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// Permission group name
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// Permission group description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// VPC type (1: CVM; 2: BM 1.0)
	VpcType *uint64 `json:"VpcType,omitempty" name:"VpcType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type AccessRule struct {
	// Permission rule ID
	AccessRuleId *uint64 `json:"AccessRuleId,omitempty" name:"AccessRuleId"`

	// Permission rule address (IP range or IP)
	Address *string `json:"Address,omitempty" name:"Address"`

	// Permission rule access mode (1: read-only; 2: read-write)
	AccessMode *uint64 `json:"AccessMode,omitempty" name:"AccessMode"`

	// Priority (value range: 1–100. The smaller the value, the higher the priority)
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type AssociateAccessGroupsRequestParams struct {
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// List of permission group IDs
	AccessGroupIds []*string `json:"AccessGroupIds,omitempty" name:"AccessGroupIds"`
}

type AssociateAccessGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// List of permission group IDs
	AccessGroupIds []*string `json:"AccessGroupIds,omitempty" name:"AccessGroupIds"`
}

func (r *AssociateAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAccessGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	delete(f, "AccessGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateAccessGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateAccessGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateAccessGroupsResponseParams `json:"Response"`
}

func (r *AssociateAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAccessGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessGroupRequestParams struct {
	// Permission group name
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// VPC type (1: CVM; 2: BM 1.0)
	VpcType *uint64 `json:"VpcType,omitempty" name:"VpcType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Permission group description, which is an empty string by default
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group name
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// VPC type (1: CVM; 2: BM 1.0)
	VpcType *uint64 `json:"VpcType,omitempty" name:"VpcType"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Permission group description, which is an empty string by default
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupName")
	delete(f, "VpcType")
	delete(f, "VpcId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessGroupResponseParams struct {
	// Permission group
	AccessGroup *AccessGroup `json:"AccessGroup,omitempty" name:"AccessGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessGroupResponseParams `json:"Response"`
}

func (r *CreateAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessRulesRequestParams struct {
	// Multiple permission rules (up to 10)
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules"`

	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

type CreateAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// Multiple permission rules (up to 10)
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules"`

	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *CreateAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRules")
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessRulesResponseParams `json:"Response"`
}

func (r *CreateAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemRequestParams struct {
	// File system name
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// File system capacity (in bytes), which can range from 1 GB to 1 PB and must be an integer multiple of 1 GB
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// Whether to verify POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitempty" name:"PosixAcl"`

	// File system description, which is an empty string by default
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of superuser names, which is an empty array by default
	SuperUsers []*string `json:"SuperUsers,omitempty" name:"SuperUsers"`

	// Username of the root directory Inode, which is `hadoop` by default
	RootInodeUser *string `json:"RootInodeUser,omitempty" name:"RootInodeUser"`

	// Group name of the root directory Inode, which is `supergroup` by default
	RootInodeGroup *string `json:"RootInodeGroup,omitempty" name:"RootInodeGroup"`

	// Whether to enable verification of Ranger service addresses
	EnableRanger *bool `json:"EnableRanger,omitempty" name:"EnableRanger"`

	// List of Ranger service addresses (empty array by default)
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitempty" name:"RangerServiceAddresses"`

	// Multiple resource tags, which can be an empty array
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system name
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// File system capacity (in bytes), which can range from 1 GB to 1 PB and must be an integer multiple of 1 GB
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// Whether to verify POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitempty" name:"PosixAcl"`

	// File system description, which is an empty string by default
	Description *string `json:"Description,omitempty" name:"Description"`

	// List of superuser names, which is an empty array by default
	SuperUsers []*string `json:"SuperUsers,omitempty" name:"SuperUsers"`

	// Username of the root directory Inode, which is `hadoop` by default
	RootInodeUser *string `json:"RootInodeUser,omitempty" name:"RootInodeUser"`

	// Group name of the root directory Inode, which is `supergroup` by default
	RootInodeGroup *string `json:"RootInodeGroup,omitempty" name:"RootInodeGroup"`

	// Whether to enable verification of Ranger service addresses
	EnableRanger *bool `json:"EnableRanger,omitempty" name:"EnableRanger"`

	// List of Ranger service addresses (empty array by default)
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitempty" name:"RangerServiceAddresses"`

	// Multiple resource tags, which can be an empty array
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemName")
	delete(f, "CapacityQuota")
	delete(f, "PosixAcl")
	delete(f, "Description")
	delete(f, "SuperUsers")
	delete(f, "RootInodeUser")
	delete(f, "RootInodeGroup")
	delete(f, "EnableRanger")
	delete(f, "RangerServiceAddresses")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemResponseParams struct {
	// File system
	FileSystem *FileSystem `json:"FileSystem,omitempty" name:"FileSystem"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileSystemResponseParams `json:"Response"`
}

func (r *CreateFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifeCycleRulesRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Multiple lifecycle rules (up to 10)
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitempty" name:"LifeCycleRules"`
}

type CreateLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Multiple lifecycle rules (up to 10)
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitempty" name:"LifeCycleRules"`
}

func (r *CreateLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "LifeCycleRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifeCycleRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifeCycleRulesResponseParams `json:"Response"`
}

func (r *CreateLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMountPointRequestParams struct {
	// Mount point name
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount point status (1: enabled; 2: disabled)
	MountPointStatus *uint64 `json:"MountPointStatus,omitempty" name:"MountPointStatus"`
}

type CreateMountPointRequest struct {
	*tchttp.BaseRequest
	
	// Mount point name
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount point status (1: enabled; 2: disabled)
	MountPointStatus *uint64 `json:"MountPointStatus,omitempty" name:"MountPointStatus"`
}

func (r *CreateMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointName")
	delete(f, "FileSystemId")
	delete(f, "MountPointStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMountPointResponseParams struct {
	// Mount point
	MountPoint *MountPoint `json:"MountPoint,omitempty" name:"MountPoint"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMountPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateMountPointResponseParams `json:"Response"`
}

func (r *CreateMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRestoreTasksRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Multiple restoration tasks (up to 10)
	RestoreTasks []*RestoreTask `json:"RestoreTasks,omitempty" name:"RestoreTasks"`
}

type CreateRestoreTasksRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Multiple restoration tasks (up to 10)
	RestoreTasks []*RestoreTask `json:"RestoreTasks,omitempty" name:"RestoreTasks"`
}

func (r *CreateRestoreTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRestoreTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "RestoreTasks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRestoreTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRestoreTasksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRestoreTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateRestoreTasksResponseParams `json:"Response"`
}

func (r *CreateRestoreTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRestoreTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessGroupRequestParams struct {
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

type DeleteAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *DeleteAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessGroupResponseParams `json:"Response"`
}

func (r *DeleteAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessRulesRequestParams struct {
	// Multiple permission rule IDs (up to 10)
	AccessRuleIds []*uint64 `json:"AccessRuleIds,omitempty" name:"AccessRuleIds"`
}

type DeleteAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// Multiple permission rule IDs (up to 10)
	AccessRuleIds []*uint64 `json:"AccessRuleIds,omitempty" name:"AccessRuleIds"`
}

func (r *DeleteAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessRulesResponseParams `json:"Response"`
}

func (r *DeleteAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSystemRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DeleteFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSystemResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileSystemResponseParams `json:"Response"`
}

func (r *DeleteFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifeCycleRulesRequestParams struct {
	// Multiple lifecycle rule IDs (up to 10)
	LifeCycleRuleIds []*uint64 `json:"LifeCycleRuleIds,omitempty" name:"LifeCycleRuleIds"`
}

type DeleteLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// Multiple lifecycle rule IDs (up to 10)
	LifeCycleRuleIds []*uint64 `json:"LifeCycleRuleIds,omitempty" name:"LifeCycleRuleIds"`
}

func (r *DeleteLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifeCycleRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifeCycleRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLifeCycleRulesResponseParams `json:"Response"`
}

func (r *DeleteLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountPointRequestParams struct {
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`
}

type DeleteMountPointRequest struct {
	*tchttp.BaseRequest
	
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`
}

func (r *DeleteMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountPointResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMountPointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMountPointResponseParams `json:"Response"`
}

func (r *DeleteMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupRequestParams struct {
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

type DescribeAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *DescribeAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupResponseParams struct {
	// Permission group
	AccessGroup *AccessGroup `json:"AccessGroup,omitempty" name:"AccessGroup"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessGroupResponseParams `json:"Response"`
}

func (r *DescribeAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupsRequestParams struct {
	// VPC ID
	// Note: either `VpcId` or `OwnerUin` can be specified as the input parameter
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Resource owner `Uin`
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type DescribeAccessGroupsRequest struct {
	*tchttp.BaseRequest
	
	// VPC ID
	// Note: either `VpcId` or `OwnerUin` can be specified as the input parameter
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Resource owner `Uin`
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *DescribeAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupsResponseParams struct {
	// List of permission groups
	AccessGroups []*AccessGroup `json:"AccessGroups,omitempty" name:"AccessGroups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessGroupsResponseParams `json:"Response"`
}

func (r *DescribeAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessRulesRequestParams struct {
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

type DescribeAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *DescribeAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessRulesResponseParams struct {
	// List of permission rules
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessRulesResponseParams `json:"Response"`
}

func (r *DescribeAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemResponseParams struct {
	// File system
	FileSystem *FileSystem `json:"FileSystem,omitempty" name:"FileSystem"`

	// Used capacity of the file system, in bytes
	// Note: this field may return `null`, indicating that no valid value was found.
	CapacityUsed *uint64 `json:"CapacityUsed,omitempty" name:"CapacityUsed"`

	// Used ARCHIVE capacity of COS, in bytes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ArchiveCapacityUsed *uint64 `json:"ArchiveCapacityUsed,omitempty" name:"ArchiveCapacityUsed"`

	// Used STANDARD capacity of COS, in bytes
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StandardCapacityUsed *uint64 `json:"StandardCapacityUsed,omitempty" name:"StandardCapacityUsed"`

	// Used STANDARD_IA capacity of COS, in bytes
	// Note: this field may return `null`, indicating that no valid value was found.
	DegradeCapacityUsed *uint64 `json:"DegradeCapacityUsed,omitempty" name:"DegradeCapacityUsed"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSystemResponseParams `json:"Response"`
}

func (r *DescribeFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemsRequestParams struct {

}

type DescribeFileSystemsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemsResponseParams struct {
	// List of file systems
	FileSystems []*FileSystem `json:"FileSystems,omitempty" name:"FileSystems"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSystemsResponseParams `json:"Response"`
}

func (r *DescribeFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifeCycleRulesRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifeCycleRulesResponseParams struct {
	// List of lifecycle rules
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitempty" name:"LifeCycleRules"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLifeCycleRulesResponseParams `json:"Response"`
}

func (r *DescribeLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointRequestParams struct {
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`
}

type DescribeMountPointRequest struct {
	*tchttp.BaseRequest
	
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`
}

func (r *DescribeMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointResponseParams struct {
	// Mount point
	MountPoint *MountPoint `json:"MountPoint,omitempty" name:"MountPoint"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMountPointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountPointResponseParams `json:"Response"`
}

func (r *DescribeMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointsRequestParams struct {
	// File system ID
	// Note: only one of `AccessGroupId`, `FileSystemId`, and `OwnerUin` can be specified as the input parameter
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// Resource owner `Uin`
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type DescribeMountPointsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	// Note: only one of `AccessGroupId`, `FileSystemId`, and `OwnerUin` can be specified as the input parameter
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// Resource owner `Uin`
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

func (r *DescribeMountPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "AccessGroupId")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointsResponseParams struct {
	// List of mount points
	MountPoints []*MountPoint `json:"MountPoints,omitempty" name:"MountPoints"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMountPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountPointsResponseParams `json:"Response"`
}

func (r *DescribeMountPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsResponseParams struct {
	// List of resource tags
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTasksRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeRestoreTasksRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeRestoreTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRestoreTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTasksResponseParams struct {
	// List of restoration tasks
	RestoreTasks []*RestoreTask `json:"RestoreTasks,omitempty" name:"RestoreTasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRestoreTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRestoreTasksResponseParams `json:"Response"`
}

func (r *DescribeRestoreTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAccessGroupsRequestParams struct {
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// List of permission group IDs
	AccessGroupIds []*string `json:"AccessGroupIds,omitempty" name:"AccessGroupIds"`
}

type DisassociateAccessGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// List of permission group IDs
	AccessGroupIds []*string `json:"AccessGroupIds,omitempty" name:"AccessGroupIds"`
}

func (r *DisassociateAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAccessGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	delete(f, "AccessGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateAccessGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAccessGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateAccessGroupsResponseParams `json:"Response"`
}

func (r *DisassociateAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAccessGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystem struct {
	// Resource owner `AppId`
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// File system name
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// File system description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// File system block size (in bytes)
	BlockSize *uint64 `json:"BlockSize,omitempty" name:"BlockSize"`

	// File system capacity (in bytes)
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// File system status (1: creating; 2: created successfully; 3: failed to create)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// List of superuser names
	SuperUsers []*string `json:"SuperUsers,omitempty" name:"SuperUsers"`

	// POSIX permission control
	PosixAcl *bool `json:"PosixAcl,omitempty" name:"PosixAcl"`

	// Whether to enable verification of Ranger service addresses
	// Note: this field may return `null`, indicating that no valid value was found.
	EnableRanger *bool `json:"EnableRanger,omitempty" name:"EnableRanger"`

	// List of Ranger service addresses
	// Note: this field may return `null`, indicating that no valid value was found.
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitempty" name:"RangerServiceAddresses"`
}

type LifeCycleRule struct {
	// Lifecycle rule ID
	LifeCycleRuleId *uint64 `json:"LifeCycleRuleId,omitempty" name:"LifeCycleRuleId"`

	// Lifecycle rule name
	LifeCycleRuleName *string `json:"LifeCycleRuleName,omitempty" name:"LifeCycleRuleName"`

	// Lifecycle rule path (directory or file)
	Path *string `json:"Path,omitempty" name:"Path"`

	// List of lifecycle rule transitions
	Transitions []*Transition `json:"Transitions,omitempty" name:"Transitions"`

	// Lifecycle rule status (1: enabled; 2: disabled)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type ModifyAccessGroupRequestParams struct {
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// Permission group name
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// Permission group description
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// Permission group ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// Permission group name
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// Permission group description
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	delete(f, "AccessGroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessGroupResponseParams `json:"Response"`
}

func (r *ModifyAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessRulesRequestParams struct {
	// Multiple permission rules (up to 10)
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules"`
}

type ModifyAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// Multiple permission rules (up to 10)
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules"`
}

func (r *ModifyAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessRulesResponseParams `json:"Response"`
}

func (r *ModifyAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// File system name
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// File system description
	Description *string `json:"Description,omitempty" name:"Description"`

	// File system capacity (in bytes), which can range from 1 GB to 1 PB and must be an integer multiple of 1 GB
	// Note: the file system capacity after change cannot be smaller than the currently used capacity
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// List of superuser names, which can be an empty array
	SuperUsers []*string `json:"SuperUsers,omitempty" name:"SuperUsers"`

	// Whether to verify POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitempty" name:"PosixAcl"`

	// Whether to enable verification of Ranger service addresses
	EnableRanger *bool `json:"EnableRanger,omitempty" name:"EnableRanger"`

	// List of Ranger service addresses, which can be an empty array
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitempty" name:"RangerServiceAddresses"`
}

type ModifyFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// File system name
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// File system description
	Description *string `json:"Description,omitempty" name:"Description"`

	// File system capacity (in bytes), which can range from 1 GB to 1 PB and must be an integer multiple of 1 GB
	// Note: the file system capacity after change cannot be smaller than the currently used capacity
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// List of superuser names, which can be an empty array
	SuperUsers []*string `json:"SuperUsers,omitempty" name:"SuperUsers"`

	// Whether to verify POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitempty" name:"PosixAcl"`

	// Whether to enable verification of Ranger service addresses
	EnableRanger *bool `json:"EnableRanger,omitempty" name:"EnableRanger"`

	// List of Ranger service addresses, which can be an empty array
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitempty" name:"RangerServiceAddresses"`
}

func (r *ModifyFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FileSystemName")
	delete(f, "Description")
	delete(f, "CapacityQuota")
	delete(f, "SuperUsers")
	delete(f, "PosixAcl")
	delete(f, "EnableRanger")
	delete(f, "RangerServiceAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFileSystemResponseParams `json:"Response"`
}

func (r *ModifyFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifeCycleRulesRequestParams struct {
	// Multiple lifecycle rules (up to 10)
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitempty" name:"LifeCycleRules"`
}

type ModifyLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// Multiple lifecycle rules (up to 10)
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitempty" name:"LifeCycleRules"`
}

func (r *ModifyLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifeCycleRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifeCycleRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLifeCycleRulesResponseParams `json:"Response"`
}

func (r *ModifyLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMountPointRequestParams struct {
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// Mount point name
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// Mount point status
	MountPointStatus *uint64 `json:"MountPointStatus,omitempty" name:"MountPointStatus"`
}

type ModifyMountPointRequest struct {
	*tchttp.BaseRequest
	
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// Mount point name
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// Mount point status
	MountPointStatus *uint64 `json:"MountPointStatus,omitempty" name:"MountPointStatus"`
}

func (r *ModifyMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	delete(f, "MountPointName")
	delete(f, "MountPointStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMountPointResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMountPointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMountPointResponseParams `json:"Response"`
}

func (r *ModifyMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsRequestParams struct {
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Multiple resource tags, which can be an empty array
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ModifyResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Multiple resource tags, which can be an empty array
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceTagsResponseParams `json:"Response"`
}

func (r *ModifyResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountPoint struct {
	// Mount point ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// Mount point name
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// File system ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// Mount point status (1: enabled; 2: disabled)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// List of IDs of the bound permission groups
	AccessGroupIds []*string `json:"AccessGroupIds,omitempty" name:"AccessGroupIds"`
}

type RestoreTask struct {
	// Restoration task ID
	RestoreTaskId *uint64 `json:"RestoreTaskId,omitempty" name:"RestoreTaskId"`

	// Restoration task file path
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// Restoration task type (`1`: standard; `2`: expedited; `3`: bulk, with only the expedited type available currently)
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Validity period (in days) of the temporary copy generated during restoration
	Days *uint64 `json:"Days,omitempty" name:"Days"`

	// Restoration task status (1: binding file; 2: file binding completed; 3: restoring file; 4: file restoration completed)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Transition struct {
	// Trigger time (in days)
	Days *uint64 `json:"Days,omitempty" name:"Days"`

	// Transition type (`1`: transition to ARCHIVE; `2`: delete; `3`: transition to STANDARD_IA)
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}