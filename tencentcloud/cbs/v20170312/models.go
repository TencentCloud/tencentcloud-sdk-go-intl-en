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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AdvancedRetentionPolicy struct {
	// Retains one latest snapshot each day within `Days` days. Value range: [0, 100].
	// Note: This field may return null, indicating that no valid values can be obtained.
	Days *uint64 `json:"Days,omitempty" name:"Days"`

	// Retains one latest snapshot each week within `Weeks` weeks. Value range: [0, 100].
	// Note: This field may return null, indicating that no valid values can be obtained.
	Weeks *uint64 `json:"Weeks,omitempty" name:"Weeks"`

	// Retains one latest snapshot each month within `Months` months. Value range: [0, 100].
	// Note: This field may return null, indicating that no valid values can be obtained.
	Months *uint64 `json:"Months,omitempty" name:"Months"`

	// Retains one latest snapshot each year within `Years` years. Value range: [0, 100].
	// Note: This field may return null, indicating that no valid values can be obtained.
	Years *uint64 `json:"Years,omitempty" name:"Years"`
}

// Predefined struct for user
type ApplyDiskBackupRequestParams struct {
	// ID of the cloud disk backup point, which can be queried through the `DescribeDiskBackups` API.
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// ID of the original cloud disk of the backup point, which can be queried through the `DescribeDisks` API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type ApplyDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cloud disk backup point, which can be queried through the `DescribeDiskBackups` API.
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// ID of the original cloud disk of the backup point, which can be queried through the `DescribeDisks` API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *ApplyDiskBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyDiskBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupId")
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyDiskBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyDiskBackupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyDiskBackupResponse struct {
	*tchttp.BaseResponse
	Response *ApplyDiskBackupResponseParams `json:"Response"`
}

func (r *ApplyDiskBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyDiskBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySnapshotRequestParams struct {
	// Snapshot ID, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// ID of the original cloud disk corresponding to the snapshot, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Specifies whether to shut down a CVM automatically before a rollback
	AutoStopInstance *bool `json:"AutoStopInstance,omitempty" name:"AutoStopInstance"`

	// Specifies whether to start up a CVM automatically after a rollback
	AutoStartInstance *bool `json:"AutoStartInstance,omitempty" name:"AutoStartInstance"`
}

type ApplySnapshotRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot ID, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// ID of the original cloud disk corresponding to the snapshot, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Specifies whether to shut down a CVM automatically before a rollback
	AutoStopInstance *bool `json:"AutoStopInstance,omitempty" name:"AutoStopInstance"`

	// Specifies whether to start up a CVM automatically after a rollback
	AutoStartInstance *bool `json:"AutoStartInstance,omitempty" name:"AutoStartInstance"`
}

func (r *ApplySnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "DiskId")
	delete(f, "AutoStopInstance")
	delete(f, "AutoStartInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplySnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplySnapshotResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplySnapshotResponse struct {
	*tchttp.BaseResponse
	Response *ApplySnapshotResponseParams `json:"Response"`
}

func (r *ApplySnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The number of instances mounted to data disk.
	AttachedDiskCount *uint64 `json:"AttachedDiskCount,omitempty" name:"AttachedDiskCount"`

	// The maximum number of instances that can be mounted to data disk.
	MaxAttachCount *uint64 `json:"MaxAttachCount,omitempty" name:"MaxAttachCount"`
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// ID of the CVM instance on which the cloud disk will be mounted. It can be queried via the API [DescribeInstances](https://intl.cloud.tencent.com/document/product/213/15728?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the elastic cloud disk to be mounted, which can be queried through the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1). A maximum of 10 elastic cloud disks can be mounted in a single request.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Optional parameter. If this is not passed only the mount operation is executed. If `True` is passed, the cloud disk will be configured to be terminated when the server it is mounted to is terminated. This is only valid for pay-as-you-go cloud disks.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// (Optional) Specifies the cloud disk mounting method. It’s only valid for BM models. Valid values: <br><li>PF<br><li>VF
	AttachMode *string `json:"AttachMode,omitempty" name:"AttachMode"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// ID of the CVM instance on which the cloud disk will be mounted. It can be queried via the API [DescribeInstances](https://intl.cloud.tencent.com/document/product/213/15728?from_cn_redirect=1).
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ID of the elastic cloud disk to be mounted, which can be queried through the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1). A maximum of 10 elastic cloud disks can be mounted in a single request.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Optional parameter. If this is not passed only the mount operation is executed. If `True` is passed, the cloud disk will be configured to be terminated when the server it is mounted to is terminated. This is only valid for pay-as-you-go cloud disks.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// (Optional) Specifies the cloud disk mounting method. It’s only valid for BM models. Valid values: <br><li>PF<br><li>VF
	AttachMode *string `json:"AttachMode,omitempty" name:"AttachMode"`
}

func (r *AttachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskIds")
	delete(f, "DeleteWithInstance")
	delete(f, "AttachMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse
	Response *AttachDisksResponseParams `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoMountConfiguration struct {
	// ID of the instance to which the cloud disk is attached.
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Mount point in the instance.
	MountPoint []*string `json:"MountPoint,omitempty" name:"MountPoint"`

	// File system type. Valid values: `ext4`, `xfs`.
	FileSystemType *string `json:"FileSystemType,omitempty" name:"FileSystemType"`
}

type AutoSnapshotPolicy struct {
	// The list of cloud disk IDs that the current scheduled snapshot policy is bound to.
	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`

	// Whether scheduled snapshot policy is activated.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Scheduled snapshot policy state. Value range:<br><li>NORMAL: Normal<br><li>ISOLATED: Isolated.
	AutoSnapshotPolicyState *string `json:"AutoSnapshotPolicyState,omitempty" name:"AutoSnapshotPolicyState"`

	// Whether it is to replicate a snapshot across accounts. `1`: yes, `0`: no.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCopyToRemote *uint64 `json:"IsCopyToRemote,omitempty" name:"IsCopyToRemote"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The time the scheduled snapshot will be triggered again.
	NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`

	// Scheduled snapshot policy name.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// Scheduled snapshot policy ID.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// The time the scheduled snapshot policy was created.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Number of days the snapshot created by this scheduled snapshot policy is retained.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`

	// ID of the replication target account
	// Note: This field may return null, indicating that no valid values can be obtained.
	CopyToAccountUin *string `json:"CopyToAccountUin,omitempty" name:"CopyToAccountUin"`

	// List of IDs of the instances associated with the scheduled snapshot policy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// The number of months for which the snapshots created by this scheduled snapshot policy can be retained.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionMonths *uint64 `json:"RetentionMonths,omitempty" name:"RetentionMonths"`

	// The maximum number of snapshots created by this scheduled snapshot policy that can be retained.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionAmount *uint64 `json:"RetentionAmount,omitempty" name:"RetentionAmount"`

	// Retention policy for scheduled snapshots.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdvancedRetentionPolicy *AdvancedRetentionPolicy `json:"AdvancedRetentionPolicy,omitempty" name:"AdvancedRetentionPolicy"`


	CopyFromAccountUin *string `json:"CopyFromAccountUin,omitempty" name:"CopyFromAccountUin"`


	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

// Predefined struct for user
type BindAutoSnapshotPolicyRequestParams struct {
	// ID of scheduled snapshot policy to be bound.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// List of cloud disk IDs to be bound. Maximum of 80 cloud disks can be bound per request.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// ID of scheduled snapshot policy to be bound.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// List of cloud disk IDs to be bound. Maximum of 80 cloud disks can be bound per request.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
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
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoSnapshotPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CopySnapshotCrossRegionsRequestParams struct {
	// Destination regions of the replication task. You can query the value of regions by calling [DescribeRegions](https://intl.cloud.tencent.com/document/product/213/9456?from_cn_redirect=1) API. Note that you can only specify regions that support snapshots.
	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`

	// Snapshot ID, which can be queried via the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Name of the snapshot replica. If it’s not specified, it defaults to “Copied [source snapshot ID from [region name]”
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

type CopySnapshotCrossRegionsRequest struct {
	*tchttp.BaseRequest
	
	// Destination regions of the replication task. You can query the value of regions by calling [DescribeRegions](https://intl.cloud.tencent.com/document/product/213/9456?from_cn_redirect=1) API. Note that you can only specify regions that support snapshots.
	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`

	// Snapshot ID, which can be queried via the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Name of the snapshot replica. If it’s not specified, it defaults to “Copied [source snapshot ID from [region name]”
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *CopySnapshotCrossRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopySnapshotCrossRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DestinationRegions")
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopySnapshotCrossRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopySnapshotCrossRegionsResponseParams struct {
	// Result of the cross-region replication task. The ID of the new snapshot replica is returned if the request succeeds. Otherwise `Error` is returned.
	SnapshotCopyResultSet []*SnapshotCopyResult `json:"SnapshotCopyResultSet,omitempty" name:"SnapshotCopyResultSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopySnapshotCrossRegionsResponse struct {
	*tchttp.BaseResponse
	Response *CopySnapshotCrossRegionsResponseParams `json:"Response"`
}

func (r *CopySnapshotCrossRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopySnapshotCrossRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyRequestParams struct {
	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// The name of the scheduled snapshot policy to be created. If it is left empty, the default is 'Not named'. The maximum length cannot exceed 60 bytes.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// Whether or not the scheduled snapshot policy is activated. FALSE: Not activated. TRUE: Activated. The default value is TRUE.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently. FALSE: Not retained permanently. TRUE: Retained permanently. The default value is FALSE.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The number of days that a snapshot created by this scheduled snapshot policy is retained. The default value is 7. If this parameter is specified, the IsPermanent input parameter can not be TRUE, otherwise a conflict will occur.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`

	// Whether to create an execution policy for the scheduled snapshot. TRUE: Only the time of the initial backup needs to be obtained, and no scheduled snapshot policy is actually created. FALSE: Create. The default value is FALSE.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// The name of the scheduled snapshot policy to be created. If it is left empty, the default is 'Not named'. The maximum length cannot exceed 60 bytes.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// Whether or not the scheduled snapshot policy is activated. FALSE: Not activated. TRUE: Activated. The default value is TRUE.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently. FALSE: Not retained permanently. TRUE: Retained permanently. The default value is FALSE.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The number of days that a snapshot created by this scheduled snapshot policy is retained. The default value is 7. If this parameter is specified, the IsPermanent input parameter can not be TRUE, otherwise a conflict will occur.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`

	// Whether to create an execution policy for the scheduled snapshot. TRUE: Only the time of the initial backup needs to be obtained, and no scheduled snapshot policy is actually created. FALSE: Create. The default value is FALSE.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
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
	delete(f, "Policy")
	delete(f, "AutoSnapshotPolicyName")
	delete(f, "IsActivated")
	delete(f, "IsPermanent")
	delete(f, "RetentionDays")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyResponseParams struct {
	// The ID of the newly created scheduled snapshot policy.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// The time that initial backup will start.
	NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateDiskBackupRequestParams struct {
	// Name of the cloud disk for which to create a backup point.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Name of the cloud disk backup point, which can contain up to 100 characters.
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`
}

type CreateDiskBackupRequest struct {
	*tchttp.BaseRequest
	
	// Name of the cloud disk for which to create a backup point.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Name of the cloud disk backup point, which can contain up to 100 characters.
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`
}

func (r *CreateDiskBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDiskBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDiskBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiskBackupResponseParams struct {
	// ID of the cloud disk backup point.
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDiskBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDiskBackupResponseParams `json:"Response"`
}

func (r *CreateDiskBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDiskBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisksRequestParams struct {
	// Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone and project. If no project is specified, the default project will be used.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Cloud disk billing mode. POSTPAID_BY_HOUR: Pay-as-you-go by hour<br><li>CDCPAID: Billed together with the bound dedicated cluster<br>For more information on the pricing in each mode, see [Pricing Overview](https://intl.cloud.tencent.com/document/product/362/2413?from_cn_redirect=1).
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD Cloud Storage<br><li>CLOUD_PREMIUM: Premium Cloud Disk<br><li>CLOUD_BSSD: Balanced SSD<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: ulTra SSD.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk name. If it is not specified, "Unnamed" will be used by default. The maximum length is 60 bytes.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// Tags bound to the cloud disk.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Snapshot ID. If this parameter is specified, the cloud disk will be created based on the snapshot. The snapshot must be a data disk snapshot. To query the type of a snapshot, call the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API and see the `DiskUsage` field in the response.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Number of cloud disks to be created. If it is not specified, `1` will be used by default. There is an upper limit on the maximum number of cloud disks that can be created in a single request. For more information, see [Use Limits](https://intl.cloud.tencent.com/doc/product/362/5145?from_cn_redirect=1).
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// Extra performance purchased for a cloud disk.<br>This optional parameter is only valid for ulTra SSD (CLOUD_TSSD) and Enhanced SSD (CLOUD_HSSD).
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// Cloud disk size in GB. <br><li>`DiskSize` is not required if `SnapshotId` is specified. In this case, the size of the cloud disk will be equal to that of the snapshot. <br><li>If you specify both `SnapshotId` and `DiskSize`, the specified disk size cannot be smaller than the snapshot size. <br><li>For the value range of cloud disk size, see [Cloud Disk Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Optional parameter. Default value: `False`. If `True` is specified, the new cloud disk will be shared.
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// This parameter is used to create encrypted cloud disks. It is fixed at `ENCRYPT`.
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). The monthly subscription cloud disk purchase attributes such as usage period and whether or not auto-renewal is set up can be specified using this parameter. <br>This parameter is required when creating a prepaid cloud disk. This parameter is not required when creating an hourly postpaid cloud disk. 
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// Whether to delete the associated non-permanently reserved snapshots upon deletion of the source cloud disk. `0`: No (default value). `1`: Yes. To check whether a snapshot is permanently reserved, see the `IsPermanent` field returned by the `DescribeSnapshots` API.
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitempty" name:"DeleteSnapshot"`

	// Specifies whether to automatically attach and initialize the newly created data disk.
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitempty" name:"AutoMountConfiguration"`

	// Specifies the cloud disk backup point quota.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone and project. If no project is specified, the default project will be used.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Cloud disk billing mode. POSTPAID_BY_HOUR: Pay-as-you-go by hour<br><li>CDCPAID: Billed together with the bound dedicated cluster<br>For more information on the pricing in each mode, see [Pricing Overview](https://intl.cloud.tencent.com/document/product/362/2413?from_cn_redirect=1).
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD Cloud Storage<br><li>CLOUD_PREMIUM: Premium Cloud Disk<br><li>CLOUD_BSSD: Balanced SSD<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: ulTra SSD.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk name. If it is not specified, "Unnamed" will be used by default. The maximum length is 60 bytes.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// Tags bound to the cloud disk.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Snapshot ID. If this parameter is specified, the cloud disk will be created based on the snapshot. The snapshot must be a data disk snapshot. To query the type of a snapshot, call the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API and see the `DiskUsage` field in the response.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Number of cloud disks to be created. If it is not specified, `1` will be used by default. There is an upper limit on the maximum number of cloud disks that can be created in a single request. For more information, see [Use Limits](https://intl.cloud.tencent.com/doc/product/362/5145?from_cn_redirect=1).
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// Extra performance purchased for a cloud disk.<br>This optional parameter is only valid for ulTra SSD (CLOUD_TSSD) and Enhanced SSD (CLOUD_HSSD).
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// Cloud disk size in GB. <br><li>`DiskSize` is not required if `SnapshotId` is specified. In this case, the size of the cloud disk will be equal to that of the snapshot. <br><li>If you specify both `SnapshotId` and `DiskSize`, the specified disk size cannot be smaller than the snapshot size. <br><li>For the value range of cloud disk size, see [Cloud Disk Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Optional parameter. Default value: `False`. If `True` is specified, the new cloud disk will be shared.
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// This parameter is used to create encrypted cloud disks. It is fixed at `ENCRYPT`.
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). The monthly subscription cloud disk purchase attributes such as usage period and whether or not auto-renewal is set up can be specified using this parameter. <br>This parameter is required when creating a prepaid cloud disk. This parameter is not required when creating an hourly postpaid cloud disk. 
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// Whether to delete the associated non-permanently reserved snapshots upon deletion of the source cloud disk. `0`: No (default value). `1`: Yes. To check whether a snapshot is permanently reserved, see the `IsPermanent` field returned by the `DescribeSnapshots` API.
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitempty" name:"DeleteSnapshot"`

	// Specifies whether to automatically attach and initialize the newly created data disk.
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitempty" name:"AutoMountConfiguration"`

	// Specifies the cloud disk backup point quota.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

func (r *CreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "DiskChargeType")
	delete(f, "DiskType")
	delete(f, "DiskName")
	delete(f, "Tags")
	delete(f, "SnapshotId")
	delete(f, "DiskCount")
	delete(f, "ThroughputPerformance")
	delete(f, "DiskSize")
	delete(f, "Shareable")
	delete(f, "ClientToken")
	delete(f, "Encrypt")
	delete(f, "DiskChargePrepaid")
	delete(f, "DeleteSnapshot")
	delete(f, "AutoMountConfiguration")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisksResponseParams struct {
	// List of IDs of the created cloud disks.
	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisksResponseParams `json:"Response"`
}

func (r *CreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotRequestParams struct {
	// ID of the cloud disk for which to create a snapshot, which can be queried through the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Snapshot name. If it is not specified, "Unnamed" will be used by default.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Expiration time of the snapshot. It must be in UTC ISO-8601 format, such as 2022-01-08T09:47:55+00:00. The snapshot will be automatically deleted when it expires.
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// ID of the cloud disk backup point. When this parameter is specified, the snapshot will be created from the backup point.
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// Tags bound to the snapshot.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cloud disk for which to create a snapshot, which can be queried through the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Snapshot name. If it is not specified, "Unnamed" will be used by default.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Expiration time of the snapshot. It must be in UTC ISO-8601 format, such as 2022-01-08T09:47:55+00:00. The snapshot will be automatically deleted when it expires.
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// ID of the cloud disk backup point. When this parameter is specified, the snapshot will be created from the backup point.
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// Tags bound to the snapshot.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "SnapshotName")
	delete(f, "Deadline")
	delete(f, "DiskBackupId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotResponseParams struct {
	// ID of the new snapshot.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotResponseParams `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPoliciesRequestParams struct {
	// List of scheduled snapshot policy IDs to be deleted.
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
}

type DeleteAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// List of scheduled snapshot policy IDs to be deleted.
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`
}

func (r *DeleteAutoSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPoliciesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DeleteAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDiskBackupsRequestParams struct {
	// ID of the cloud disk backup point to be deleted.
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`
}

type DeleteDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cloud disk backup point to be deleted.
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`
}

func (r *DeleteDiskBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDiskBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDiskBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDiskBackupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDiskBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDiskBackupsResponseParams `json:"Response"`
}

func (r *DeleteDiskBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDiskBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// List of IDs of snapshots to be deleted, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Whether to forcibly delete the image associated with the snapshot
	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of snapshots to be deleted, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Whether to forcibly delete the image associated with the snapshot
	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	delete(f, "DeleteBindImages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotsResponseParams `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesRequestParams struct {
	// List of scheduled snapshot policy IDs to be queried. The parameter does not support specifying both `SnapshotIds` and `Filters`.
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`

	// Filter conditions. Specification of both the `AutoSnapshotPolicyIds` and `Filters` parameters is not supported.<br><li>auto-snapshot-policy-id - Array of String - Required or not: No - (Filter condition) Filters according to the scheduled snapshot policy ID. The format of the scheduled snapshot policy ID is as follows: `asp-11112222`. <br><li>auto-snapshot-policy-state - Array of String - Required or not: No - (Filter condition) Filters according to the status of the scheduled snapshot policy. The format of the scheduled snapshot policy ID is as follows: `asp-11112222`. (NORMAL: normal | ISOLATED: isolated)<br><li>auto-snapshot-policy-name - Array of String - Required or not: No - (Filter condition) Filters according to the name of the scheduled snapshot policy.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of results to be returned. Default is 20. Maximum is 100. For more information on `Limit`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default is 0. For more information on `Offset`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Outputs the ordering of the scheduled snapshot lists. Value range: <br><li>ASC: Ascending order <br><li>DESC: Descending order.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The sorting filter applied to the scheduled snapshot list. Value range: <Sort by creation time of scheduled snapshot. By default, this is sorted by creation time.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// List of scheduled snapshot policy IDs to be queried. The parameter does not support specifying both `SnapshotIds` and `Filters`.
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`

	// Filter conditions. Specification of both the `AutoSnapshotPolicyIds` and `Filters` parameters is not supported.<br><li>auto-snapshot-policy-id - Array of String - Required or not: No - (Filter condition) Filters according to the scheduled snapshot policy ID. The format of the scheduled snapshot policy ID is as follows: `asp-11112222`. <br><li>auto-snapshot-policy-state - Array of String - Required or not: No - (Filter condition) Filters according to the status of the scheduled snapshot policy. The format of the scheduled snapshot policy ID is as follows: `asp-11112222`. (NORMAL: normal | ISOLATED: isolated)<br><li>auto-snapshot-policy-name - Array of String - Required or not: No - (Filter condition) Filters according to the name of the scheduled snapshot policy.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Number of results to be returned. Default is 20. Maximum is 100. For more information on `Limit`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset. Default is 0. For more information on `Offset`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Outputs the ordering of the scheduled snapshot lists. Value range: <br><li>ASC: Ascending order <br><li>DESC: Descending order.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The sorting filter applied to the scheduled snapshot list. Value range: <Sort by creation time of scheduled snapshot. By default, this is sorted by creation time.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
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
	delete(f, "AutoSnapshotPolicyIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesResponseParams struct {
	// The quantity of valid scheduled snapshot policies.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of scheduled snapshot policies.
	AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDiskAssociatedAutoSnapshotPolicyRequestParams struct {
	// The ID of the queried cloud disk.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type DescribeDiskAssociatedAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the queried cloud disk.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskAssociatedAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskAssociatedAutoSnapshotPolicyResponseParams struct {
	// The quantity of scheduled snapshots binded to cloud disk.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of scheduled snapshots bound to cloud disk.
	AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskAssociatedAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskAssociatedAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsRequestParams struct {
	// List of IDs of the backup points to be queried. `DiskBackupIds` and `Filters` cannot be specified at the same time.
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`

	// Filter. `DiskBackupIds` and `Filters` cannot be specified at the same time. Valid values: <br><li>disk-backup-id - Array of String - Required: No - (Filter) Filter by backup point ID in the format of `dbp-11112222`.
	// <br><li>disk-id - Array of String - Required: No - (Filter) Filter by ID of the cloud disk for which backup points are created.
	// <br><li>disk-usage - Array of String - Required: No - (Filter) Filter by type of the cloud disk for which backup points are created. (SYSTEM_DISK: System disk | DATA_DISK: Data disk)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting order of cloud disk backup points. Valid values:<br><li>ASC: Ascending<br><li>DESC: Descending
	Order *string `json:"Order,omitempty" name:"Order"`

	// The field by which cloud disk backup points are sorted. Valid values:<br><li>CREATE_TIME: Sort by creation time<br>Backup points are sorted by creation time by default.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeDiskBackupsRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the backup points to be queried. `DiskBackupIds` and `Filters` cannot be specified at the same time.
	DiskBackupIds []*string `json:"DiskBackupIds,omitempty" name:"DiskBackupIds"`

	// Filter. `DiskBackupIds` and `Filters` cannot be specified at the same time. Valid values: <br><li>disk-backup-id - Array of String - Required: No - (Filter) Filter by backup point ID in the format of `dbp-11112222`.
	// <br><li>disk-id - Array of String - Required: No - (Filter) Filter by ID of the cloud disk for which backup points are created.
	// <br><li>disk-usage - Array of String - Required: No - (Filter) Filter by type of the cloud disk for which backup points are created. (SYSTEM_DISK: System disk | DATA_DISK: Data disk)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant section of the API [Overview](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting order of cloud disk backup points. Valid values:<br><li>ASC: Ascending<br><li>DESC: Descending
	Order *string `json:"Order,omitempty" name:"Order"`

	// The field by which cloud disk backup points are sorted. Valid values:<br><li>CREATE_TIME: Sort by creation time<br>Backup points are sorted by creation time by default.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeDiskBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskBackupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskBackupsResponseParams struct {
	// Number of eligible cloud disk backup points.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of details of cloud disk backup points.
	DiskBackupSet []*DiskBackup `json:"DiskBackupSet,omitempty" name:"DiskBackupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskBackupsResponseParams `json:"Response"`
}

func (r *DescribeDiskBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskConfigQuotaRequestParams struct {
	// Inquiry type. Value range: INQUIRY_CBS_CONFIG: query the configuration list of cloud disks <br><li>INQUIRY_CVM_CONFIG: query the configuration list of cloud disks and instances.
	InquiryType *string `json:"InquiryType,omitempty" name:"InquiryType"`

	// Query configuration under one or more [availability zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo).
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Billing mode. Value range: <br><li>POSTPAID_BY_HOUR: postpaid.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD
	DiskTypes []*string `json:"DiskTypes,omitempty" name:"DiskTypes"`

	// The system disk or data disk. Value range: <br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Filter by the instance model series, such as S1, I1 and M1. For more information, please see [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1)
	InstanceFamilies []*string `json:"InstanceFamilies,omitempty" name:"InstanceFamilies"`

	// Instance CPU cores.
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// Instance memory size.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
}

type DescribeDiskConfigQuotaRequest struct {
	*tchttp.BaseRequest
	
	// Inquiry type. Value range: INQUIRY_CBS_CONFIG: query the configuration list of cloud disks <br><li>INQUIRY_CVM_CONFIG: query the configuration list of cloud disks and instances.
	InquiryType *string `json:"InquiryType,omitempty" name:"InquiryType"`

	// Query configuration under one or more [availability zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo).
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// Billing mode. Value range: <br><li>POSTPAID_BY_HOUR: postpaid.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD
	DiskTypes []*string `json:"DiskTypes,omitempty" name:"DiskTypes"`

	// The system disk or data disk. Value range: <br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Filter by the instance model series, such as S1, I1 and M1. For more information, please see [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1)
	InstanceFamilies []*string `json:"InstanceFamilies,omitempty" name:"InstanceFamilies"`

	// Instance CPU cores.
	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`

	// Instance memory size.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
}

func (r *DescribeDiskConfigQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InquiryType")
	delete(f, "Zones")
	delete(f, "DiskChargeType")
	delete(f, "DiskTypes")
	delete(f, "DiskUsage")
	delete(f, "InstanceFamilies")
	delete(f, "CPU")
	delete(f, "Memory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskConfigQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskConfigQuotaResponseParams struct {
	// List of cloud disk configurations.
	DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitempty" name:"DiskConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskConfigQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskConfigQuotaResponseParams `json:"Response"`
}

func (r *DescribeDiskConfigQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskOperationLogsRequestParams struct {
	// Filter conditions. The following conditions are supported:
	// <li>disk-id - Array of String - Required or not: Yes - Filter by cloud disk ID, with maximum of 10 cloud disk IDs able to be specified per request.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The start time of the operation logs to be queried, for example: '2019-11-22 00:00:00"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the operation logs to be queried, for example: '2019-11-22 23:59:59"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeDiskOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions. The following conditions are supported:
	// <li>disk-id - Array of String - Required or not: Yes - Filter by cloud disk ID, with maximum of 10 cloud disk IDs able to be specified per request.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The start time of the operation logs to be queried, for example: '2019-11-22 00:00:00"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the operation logs to be queried, for example: '2019-11-22 23:59:59"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDiskOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskOperationLogsResponseParams struct {
	// List of cloud disk operation logs.
	DiskOperationLogSet []*DiskOperationLog `json:"DiskOperationLogSet,omitempty" name:"DiskOperationLogSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiskOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeDiskOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksRequestParams struct {
	// Query by one or more cloud disk IDs, such as `disk-11112222`. For the format of this parameter, please see the ids.N section of the API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1). This parameter does not support specifying both `DiskIds` and `Filters`.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Filters. You cannot specify `DiskIds` and `Filters` at the same time. <br><li>disk-usage - Array of String - Optional - Filters by cloud disk type. (SYSTEM_DISK: system disk | DATA_DISK: data disk) <br><li>disk-charge-type - Array of String - Optional - Filters by cloud disk billing method. (POSTPAID_BY_HOUR: pay-as-you-go) <br><li>portable - Array of String- Optional - Filters by whether the cloud disk is elastic or not. (TRUE: elastic | FALSE: non-elastic) <br><li>project-id - Array of Integer - Optional - Filters by the ID of the project to which a cloud disk belongs. <br><li>disk-id - Array of String - Optional - Filters by cloud disk ID, such as `disk-11112222`. <br><li>disk-name - Array of String - Optional - Filters by cloud disk name. <br><li>disk-type - Array of String - Optional - Filters by cloud disk media type (CLOUD_BASIC: HDD cloud disk | CLOUD_PREMIUM: Premium Cloud Storage | CLOUD_SSD: SSD cloud disk.) <br><li>disk-state - Array of String - Optional - Filters by cloud disk state. (UNATTACHED: not mounted | ATTACHING: being mounted | ATTACHED: mounted | DETACHING: being unmounted | EXPANDING: being expanded | ROLLBACKING: being rolled back | TORECYCLE: to be repossessed.) <br><li>instance-id - Array of String - Optional - Filters by the ID of the CVM instance on which a cloud disk is mounted. You can use this parameter to query the cloud disks mounted on specific CVMs. <br><li>zone - Array of String - Optional - Filters by [availability zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) <br><li>instance-ip-address - Array of String - Optional - Filters by the private or public IP of the CVM on which a cloud disk is mounted. <br><li>instance-name - Array of String - Optional - Filters by the name of the instance on which a cloud disk is mounted. <br><li>tag-key - Array of String - Optional - Filters by tag key. <br><li>tag-value - Array of String - Optional - Filters by tag value. <br><li>tag:tag-key - Array of String - Optional - Filters by tag key-value pair. Please replace `tag-key` with a specific tag key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default is 0. For more information on `Offset`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default is 20. Maximum is 100. For more information on `Limit`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Outputs the ordering of the cloud disk list. Value range: <br><li>ASC: Ascending order <br><li>DESC: Descending order.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The field by which the cloud disk list is sorted. Value range: <br><li>CREATE_TIME: sorted by the creation time of cloud disks <br><li>DEADLINE: sorted by the expiration time of cloud disks<br>By default, the cloud disk list is sorted by the creation time of cloud disks.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Whether the ID of the periodic snapshot policy bound to the cloud disk needs to be returned in the cloud disk details. TRUE: return; FALSE: do not return.
	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// Query by one or more cloud disk IDs, such as `disk-11112222`. For the format of this parameter, please see the ids.N section of the API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1). This parameter does not support specifying both `DiskIds` and `Filters`.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Filters. You cannot specify `DiskIds` and `Filters` at the same time. <br><li>disk-usage - Array of String - Optional - Filters by cloud disk type. (SYSTEM_DISK: system disk | DATA_DISK: data disk) <br><li>disk-charge-type - Array of String - Optional - Filters by cloud disk billing method. (POSTPAID_BY_HOUR: pay-as-you-go) <br><li>portable - Array of String- Optional - Filters by whether the cloud disk is elastic or not. (TRUE: elastic | FALSE: non-elastic) <br><li>project-id - Array of Integer - Optional - Filters by the ID of the project to which a cloud disk belongs. <br><li>disk-id - Array of String - Optional - Filters by cloud disk ID, such as `disk-11112222`. <br><li>disk-name - Array of String - Optional - Filters by cloud disk name. <br><li>disk-type - Array of String - Optional - Filters by cloud disk media type (CLOUD_BASIC: HDD cloud disk | CLOUD_PREMIUM: Premium Cloud Storage | CLOUD_SSD: SSD cloud disk.) <br><li>disk-state - Array of String - Optional - Filters by cloud disk state. (UNATTACHED: not mounted | ATTACHING: being mounted | ATTACHED: mounted | DETACHING: being unmounted | EXPANDING: being expanded | ROLLBACKING: being rolled back | TORECYCLE: to be repossessed.) <br><li>instance-id - Array of String - Optional - Filters by the ID of the CVM instance on which a cloud disk is mounted. You can use this parameter to query the cloud disks mounted on specific CVMs. <br><li>zone - Array of String - Optional - Filters by [availability zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) <br><li>instance-ip-address - Array of String - Optional - Filters by the private or public IP of the CVM on which a cloud disk is mounted. <br><li>instance-name - Array of String - Optional - Filters by the name of the instance on which a cloud disk is mounted. <br><li>tag-key - Array of String - Optional - Filters by tag key. <br><li>tag-value - Array of String - Optional - Filters by tag value. <br><li>tag:tag-key - Array of String - Optional - Filters by tag key-value pair. Please replace `tag-key` with a specific tag key.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default is 0. For more information on `Offset`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default is 20. Maximum is 100. For more information on `Limit`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Outputs the ordering of the cloud disk list. Value range: <br><li>ASC: Ascending order <br><li>DESC: Descending order.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The field by which the cloud disk list is sorted. Value range: <br><li>CREATE_TIME: sorted by the creation time of cloud disks <br><li>DEADLINE: sorted by the expiration time of cloud disks<br>By default, the cloud disk list is sorted by the creation time of cloud disks.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// Whether the ID of the periodic snapshot policy bound to the cloud disk needs to be returned in the cloud disk details. TRUE: return; FALSE: do not return.
	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "ReturnBindAutoSnapshotPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksResponseParams struct {
	// The quantity of cloud disks meeting the conditions.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of cloud disk details.
	DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksResponseParams `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDiskNumRequestParams struct {
	// ID of the CVM instance can be queried via the API [DescribeInstances](https://intl.cloud.tencent.com/document/product/213/15728?from_cn_redirect=1).
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeInstancesDiskNumRequest struct {
	*tchttp.BaseRequest
	
	// ID of the CVM instance can be queried via the API [DescribeInstances](https://intl.cloud.tencent.com/document/product/213/15728?from_cn_redirect=1).
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesDiskNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDiskNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDiskNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDiskNumResponseParams struct {
	// The quantity of mounted and mountable elastic cloud disks for each cloud virtual machine
	AttachDetail []*AttachDetail `json:"AttachDetail,omitempty" name:"AttachDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesDiskNumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDiskNumResponseParams `json:"Response"`
}

func (r *DescribeInstancesDiskNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDiskNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOperationLogsRequestParams struct {
	// Filter conditions. The following conditions are supported:
	// <li>snapshot-id - Array of String - Required or not: Yes - Filter by snapshot ID, with maximum of 10 snapshot IDs able to be specified per request.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The start time of the operation logs to be queried, for example: '2019-11-22 00:00:00"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the operation logs to be queried, for example: '2019-11-22 23:59:59"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeSnapshotOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions. The following conditions are supported:
	// <li>snapshot-id - Array of String - Required or not: Yes - Filter by snapshot ID, with maximum of 10 snapshot IDs able to be specified per request.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The start time of the operation logs to be queried, for example: '2019-11-22 00:00:00"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the operation logs to be queried, for example: '2019-11-22 23:59:59"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	delete(f, "Filters")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOperationLogsResponseParams struct {
	// List of snapshot operation logs.
	SnapshotOperationLogSet []*SnapshotOperationLog `json:"SnapshotOperationLogSet,omitempty" name:"SnapshotOperationLogSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeSnapshotSharePermissionRequestParams struct {
	// The ID of the snapshot to be queried. You can obtain this by using [DescribeSnapshots](https://intl.cloud.tencent.com/document/api/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type DescribeSnapshotSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the snapshot to be queried. You can obtain this by using [DescribeSnapshots](https://intl.cloud.tencent.com/document/api/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *DescribeSnapshotSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotSharePermissionResponseParams struct {
	// The set of snapshot sharing information
	SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotSharePermissionResponseParams `json:"Response"`
}

func (r *DescribeSnapshotSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// List of snapshot IDs to be queried. The parameter does not support specifying both `SnapshotIds` and `Filters`.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Filters. It cannot be specified together with `SnapshotIds`.<br><li>snapshot-id - Array of String - Optional - Filters by snapshot ID, such as `snap-11112222`.<br><li>snapshot-name - Array of String - Optional - Filters by snapshot name. <br><li>snapshot-state - Array of String - Optional - Filters by snapshot state (NORMAL: normal | CREATING: creating | ROLLBACKING: rolling back). <br><li>disk-usage - Array of String - Optional - Filters by the type of the cloud disk from which a snapshot is created (SYSTEM_DISK: system disk | DATA_DISK: data disk).<br><li>project-id - Array of String - Optional - Filters by the ID of the project to which a cloud disk belongs. <br><li>disk-id - Array of String - Optional - Filters by the ID of the cloud disk from which a snapshot is created.<br><li>zone - Array of String - Optional - Filters by [availability zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo).<br><li>encrypt - Array of String - Optional - Filters by whether a snapshot is created from an encrypted cloud disk. (TRUE: a snapshot of an encrypted cloud disk | FALSE: not a snapshot of an encrypted cloud disk.)
	// <li>snapshot-type- Array of String - Optional - Filters by the snapshot type specified in `snapshot-type`.
	// (SHARED_SNAPSHOT: a shared snapshot | PRIVATE_SNAPSHOT: a private snapshot.)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default is 0. For more information on `Offset`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default is 20. Maximum is 100. For more information on `Limit`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Outputs the ordering of the cloud disk list. Value range: <br><li>ASC: Ascending order <br><li>DESC: Descending order.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The field by which the snapshot list is sorted. Value range: <br><li>CREATE_TIME: sorted by the creation time of the snapshots <br>By default, the snapshot list is sorted by the creation time of snapshots.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// List of snapshot IDs to be queried. The parameter does not support specifying both `SnapshotIds` and `Filters`.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Filters. It cannot be specified together with `SnapshotIds`.<br><li>snapshot-id - Array of String - Optional - Filters by snapshot ID, such as `snap-11112222`.<br><li>snapshot-name - Array of String - Optional - Filters by snapshot name. <br><li>snapshot-state - Array of String - Optional - Filters by snapshot state (NORMAL: normal | CREATING: creating | ROLLBACKING: rolling back). <br><li>disk-usage - Array of String - Optional - Filters by the type of the cloud disk from which a snapshot is created (SYSTEM_DISK: system disk | DATA_DISK: data disk).<br><li>project-id - Array of String - Optional - Filters by the ID of the project to which a cloud disk belongs. <br><li>disk-id - Array of String - Optional - Filters by the ID of the cloud disk from which a snapshot is created.<br><li>zone - Array of String - Optional - Filters by [availability zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo).<br><li>encrypt - Array of String - Optional - Filters by whether a snapshot is created from an encrypted cloud disk. (TRUE: a snapshot of an encrypted cloud disk | FALSE: not a snapshot of an encrypted cloud disk.)
	// <li>snapshot-type- Array of String - Optional - Filters by the snapshot type specified in `snapshot-type`.
	// (SHARED_SNAPSHOT: a shared snapshot | PRIVATE_SNAPSHOT: a private snapshot.)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default is 0. For more information on `Offset`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. Default is 20. Maximum is 100. For more information on `Limit`, please see relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/product/362/15633?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Outputs the ordering of the cloud disk list. Value range: <br><li>ASC: Ascending order <br><li>DESC: Descending order.
	Order *string `json:"Order,omitempty" name:"Order"`

	// The field by which the snapshot list is sorted. Value range: <br><li>CREATE_TIME: sorted by the creation time of the snapshots <br>By default, the snapshot list is sorted by the creation time of snapshots.
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// Number of snapshots.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of snapshot details.
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachDisksRequestParams struct {
	// IDs of the cloud disks to be unmounted, which can be queried via the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API. Up to 10 elastic cloud disks can be unmounted in a single request.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Indicates the CVM from which you want to unmount the disks. This parameter is only available for shared cloud disks.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest
	
	// IDs of the cloud disks to be unmounted, which can be queried via the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API. Up to 10 elastic cloud disks can be unmounted in a single request.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Indicates the CVM from which you want to unmount the disks. This parameter is only available for shared cloud disks.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachDisksResponse struct {
	*tchttp.BaseResponse
	Response *DetachDisksResponseParams `json:"Response"`
}

func (r *DetachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {
	// Whether the cloud disk terminates along with the instance mounted to it. <br><li>true: Cloud disk will also be terminated when instance terminates, so only hourly postpaid cloud disk are supported.<br><li>false: Cloud disk does not terminate when instance terminates.
	// Note: This field may return null, indicating that no valid value was found.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// Auto renewal flag. Supported values:<br><li>NOTIFY_AND_AUTO_RENEW: Notify expiry and renew automatically<br><li>NOTIFY_AND_MANUAL_RENEW: Notify expiry but not renew automatically<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Neither notify expiry nor renew automatically.
	// Note: This field may return null, indicating that no valid value was found.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Cloud disk types. Valid values: <br><li>CLOUD_BASIC: HDD cloud disk <br><li>CLOUD_PREMIUM: Premium Cloud Disk <br><li>CLOUD_BSSD: General Purpose SSD <br><li>CLOUD_SSD: SSD <br><li>CLOUD_HSSD: Enhanced SSD <br><li>CLOUD_TSSD: Tremendous SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// The state of the cloud disk. Value range: <br><li>UNATTACHED: Not mounted <br><li>ATTACHING: Mounting <br><li>ATTACHED: Mounted <br><li>DETACHING: Un-mounting <br><li>EXPANDING: Expanding <br><li>ROLLBACKING: Rolling back <br><li>TORECYCE: Pending recycling. <br><li>DUMPING: Copying the hard drive.
	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`

	// The total number of snapshots of the cloud disk.
	SnapshotCount *int64 `json:"SnapshotCount,omitempty" name:"SnapshotCount"`

	// Cloud disk already mounted to CVM, and both CVM and cloud disk use monthly subscription.<br><li>true: CVM has auto-renewal flag set up, but cloud disk does not.<br><li>false: Cloud disk auto-renewal flag set up normally.
	// Note: This field may return null, indicating that no valid value was found.
	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`

	// Whether the cloud disk is in the status of snapshot rollback. Value range: <br><li>false: No <br><li>true: Yes
	Rollbacking *bool `json:"Rollbacking,omitempty" name:"Rollbacking"`

	// For non-shareable cloud disks, this parameter is null. For shareable cloud disks, this parameters indicates this cloud disk's Instance IDs currently mounted to the CVM.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// Whether the cloud disk is encrypted. Value range: <br><li>false: Not encrypted <br><li>true: Encrypted.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// Cloud disk name.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// Specifies whether to create a snapshot when the cloud disk is terminated due to overdue payment or expiration. `true`: create snapshot; `false`: do not create snapshot.
	BackupDisk *bool `json:"BackupDisk,omitempty" name:"BackupDisk"`

	// The tag bound to the cloud disk. The value Null is used when no tag is bound to the cloud disk.
	// Note: This field may return null, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// ID of the CVM to which the cloud disk is mounted.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Cloud disk mount method. Valid values: <br><li>PF: mount as a PF (Physical Function)<br><li>VF: mount as a VF (Virtual Function)
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AttachMode *string `json:"AttachMode,omitempty" name:"AttachMode"`

	// ID of the periodic snapshot associated to the cloud disk. This parameter is returned only if the value of parameter ReturnBindAutoSnapshotPolicy is TRUE when the API DescribeDisks is called.
	// Note: This field may return null, indicating that no valid value was found.
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`

	// Extra performance for a cloud disk, in MB/sec.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// Whether cloud disk is in process of type change. Value range: <br><li>false: Cloud disk not in process of type change. <br><li>true: Cloud disk type change has been launched, and migration is in process.
	// Note: This field may return null, indicating that no valid value was found.
	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`

	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// The total capacity of the snapshots of the cloud disk. Unit: MB.
	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`

	// Location of the cloud disk.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Determines whether or not prepaid cloud disk supports active return. <br><li>true: Active return supported.<br><li>false: Active return not supported.
	// Note: This field may return null, indicating that no valid value was found.
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// Expiration time of the cloud disk.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Whether the cloud disk is mounted to the CVM. Value range: <br><li>false: Unmounted <br><li>true: Mounted.
	Attached *bool `json:"Attached,omitempty" name:"Attached"`

	// Cloud disk size (in GB).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Migration progress of cloud disk type change, from 0 to 100.
	// Note: This field may return null, indicating that no valid value was found.
	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`

	// Cloud disk type. Value range:<br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Billing method. Value range: <br><li>PREPAID: Prepaid, that is, monthly subscription<br><li>POSTPAID_BY_HOUR: Postpaid, that is, pay as you go.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Whether it is an elastic cloud disk. false: Non-elastic cloud disk; true: Elastic cloud disk.
	Portable *bool `json:"Portable,omitempty" name:"Portable"`

	// Whether the cloud disk has the capability to create snapshots. Value range: <br><li>false: Cannot create snapshots. true: Can create snapshots.
	SnapshotAbility *bool `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`

	// This field is only applicable when the instance is already mounted to the cloud disk, and both the instance and the cloud disk use monthly subscription. <br><li>true: Expiration time of cloud disk is earlier than that of the instance.<br><li>false:Expiration time of cloud disk is later than that of the instance.
	// Note: This field may return null, indicating that no valid value was found.
	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`

	// Rollback progress of a cloud disk snapshot.
	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`

	// Number of days from current time until disk expiration (only applicable for prepaid disks).
	// Note: This field may return null, indicating that no valid value was found.
	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`

	// In circumstances where the prepaid cloud disk does not support active return, this parameter indicates the reason that return is not supported. Value range: <br><li>1: The cloud disk has already been returned. <br><li>2: The cloud disk has already expired. <br><li>3: The cloud disk does not support return. <br><li> 8: The limit on the number of returns is exceeded.
	// Note: This field may return null, indicating that no valid value was found.
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// Whether or not cloud disk is shareable cloud disk.
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// Creation time of the cloud disk.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Delete the associated non-permanently reserved snapshots upon deletion of the source cloud disk. `0`: No (default). `1`: Yes. To check whether a snapshot is permanently reserved, refer to the `IsPermanent` field returned by the `DescribeSnapshots` API. 
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitempty" name:"DeleteSnapshot"`

	// Quota of cloud disk backup points, i.e., the maximum number of backup points that a cloud disk can have.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`

	// Number of used cloud disk backups.
	DiskBackupCount *uint64 `json:"DiskBackupCount,omitempty" name:"DiskBackupCount"`

	// Type of the instance mounted to the cloud disk. Valid values: <br><li>CVM<br><li>EKS
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`


	LastAttachInsId *string `json:"LastAttachInsId,omitempty" name:"LastAttachInsId"`


	ErrorPrompt *string `json:"ErrorPrompt,omitempty" name:"ErrorPrompt"`
}

type DiskBackup struct {
	// Cloud disk backup point ID.
	DiskBackupId *string `json:"DiskBackupId,omitempty" name:"DiskBackupId"`

	// ID of the cloud disk with which the backup point is associated.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk size in GB.
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Cloud disk type. Valid values:<br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Backup point name.
	DiskBackupName *string `json:"DiskBackupName,omitempty" name:"DiskBackupName"`

	// Cloud disk backup point status. Valid values:<br><li>NORMAL: Normal<br><li>CREATING: Creating<br><li>ROLLBACKING: Rolling back
	DiskBackupState *string `json:"DiskBackupState,omitempty" name:"DiskBackupState"`

	// Cloud disk creation progress in percentage.
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// Creation time of the cloud disk backup point.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether the cloud disk is encrypted. Valid values: <br><li>false: Not encrypted <br><li>true: Encrypted
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
}

type DiskChargePrepaid struct {
	// Subscription period of the cloud disk in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: Notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: Notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Neither notify upon expiration nor renew automatically <br><br>Default value: NOTIFY_AND_MANUAL_RENEW.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// You can specify this parameter when you need to ensure that a cloud disk and the CVM instance to which it is attached have the same expiration time. This parameter represents the current expiration time of the instance. In this case, if you specify `Period`, `Period` will represent how long you want to renew the instance, and the cloud disk will be renewed based on the new expiration time of the instance. For example, the value of this parameter can be `2018-03-30 20:15:03`.
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitempty" name:"CurInstanceDeadline"`
}

type DiskConfig struct {
	// Whether the configuration is available.
	Available *bool `json:"Available,omitempty" name:"Available"`

	// Billing method. Value range: <br><li>PREPAID: Prepaid, that is, monthly subscription<br><li>POSTPAID_BY_HOUR: Postpaid, that is, pay as you go.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// The [Availability Region](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) of the cloud drive.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model series. For more information, please see [Instance Models](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1)
	// Note: This field may return null, indicating that no valid value was found.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Type of cloud disk medium. Value range: <br><li>CLOUD_BASIC: Ordinary cloud disk <br><li>CLOUD_PREMIUM: Premium cloud storage <br><li>CLOUD_SSD: SSD cloud disk.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Minimum increment of cloud disk size adjustment in GB.
	// Note: This field might return null, indicating that no valid values can be obtained.
	StepSize *uint64 `json:"StepSize,omitempty" name:"StepSize"`

	// Additional performance range.
	// Note: This field might return null, indicating that no valid values can be obtained.
	ExtraPerformanceRange []*int64 `json:"ExtraPerformanceRange,omitempty" name:"ExtraPerformanceRange"`

	// Instance model.
	// Note: This field may return null, indicating that no valid value was found.
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// Cloud disk type. Value range: <br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// The minimum configurable cloud disk size (in GB).
	MinDiskSize *uint64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`

	// The maximum configurable cloud disk size (in GB).
	MaxDiskSize *uint64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`
}

type DiskOperationLog struct {
	// UIN of operator.
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Operation type. Value range:
	// CBS_OPERATION_ATTACH: Mount cloud disk
	// CBS_OPERATION_DETACH: Unmount cloud disk
	// CBS_OPERATION_RENEW: Renew
	// CBS_OPERATION_EXPAND: Expand
	// CBS_OPERATION_CREATE: Create
	// CBS_OPERATION_ISOLATE: Isolate
	// CBS_OPERATION_MODIFY: Modify cloud disk attributes
	// ASP_OPERATION_BIND: Associate scheduled snapshot policy
	// ASP_OPERATION_UNBIND: Cancel associated scheduled snapshot policy
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Cloud disk ID of operation.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Status of operation. Value range:
	// SUCCESS: Operation successful 
	// FAILED: Operation failed 
	// PROCESSING: Operation in process
	OperationState *string `json:"OperationState,omitempty" name:"OperationState"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Filter struct {
	// Name of filter key.
	Name *string `json:"Name,omitempty" name:"Name"`

	// One or more filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GetSnapOverviewRequestParams struct {

}

type GetSnapOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetSnapOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSnapOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSnapOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSnapOverviewResponseParams struct {
	// The total snapshot size of the user
	TotalSize *float64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// The total billed snapshot size of the user
	RealTradeSize *float64 `json:"RealTradeSize,omitempty" name:"RealTradeSize"`

	// Free tier of snapshot
	FreeQuota *float64 `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// Total number of snapshots
	TotalNums *int64 `json:"TotalNums,omitempty" name:"TotalNums"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetSnapOverviewResponse struct {
	*tchttp.BaseResponse
	Response *GetSnapOverviewResponseParams `json:"Response"`
}

func (r *GetSnapOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSnapOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Image struct {
	// Image instance ID.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Image name.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`
}

// Predefined struct for user
type InitializeDisksRequestParams struct {
	// ID list of the cloud disks to be reinitialized. Up to 20 disks can be reinitialized at a time.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

type InitializeDisksRequest struct {
	*tchttp.BaseRequest
	
	// ID list of the cloud disks to be reinitialized. Up to 20 disks can be reinitialized at a time.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *InitializeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitializeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitializeDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InitializeDisksResponse struct {
	*tchttp.BaseResponse
	Response *InitializeDisksResponseParams `json:"Response"`
}

func (r *InitializeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskBackupQuotaRequestParams struct {
	// Cloud disk ID, which can be queried through the `DescribeDisks` API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk backup point quota after the modification, i.e., the number of backup points that a cloud disk can have.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type InquirePriceModifyDiskBackupQuotaRequest struct {
	*tchttp.BaseRequest
	
	// Cloud disk ID, which can be queried through the `DescribeDisks` API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk backup point quota after the modification, i.e., the number of backup points that a cloud disk can have.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

func (r *InquirePriceModifyDiskBackupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskBackupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDiskBackupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskBackupQuotaResponseParams struct {
	// Price of the cloud disk after its backup point quota is modified.
	DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceModifyDiskBackupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDiskBackupQuotaResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDiskBackupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskBackupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskExtraPerformanceRequestParams struct {
	// Cloud disk ID, which can be queried via the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// The extra throughput to purchase, in MB/s
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

type InquirePriceModifyDiskExtraPerformanceRequest struct {
	*tchttp.BaseRequest
	
	// Cloud disk ID, which can be queried via the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// The extra throughput to purchase, in MB/s
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *InquirePriceModifyDiskExtraPerformanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "ThroughputPerformance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDiskExtraPerformanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDiskExtraPerformanceResponseParams struct {
	// Price for purchasing the extra performance
	DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePriceModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDiskExtraPerformanceResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDiskExtraPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDisksRequestParams struct {
	// Cloud disk billing mode. <br><li>POSTPAID_BY_HOUR: Hourly pay-as-you-go.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD Cloud Storage<br><li>CLOUD_PREMIUM: Premium Cloud Disk<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: ulTra SSD.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk size in GB. For the value range, see [Cloud Disk Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of the project to which the cloud disk belongs.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of cloud disks to be purchased. If it is not specified, `1` will be used by default.
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// Extra performance in MB/s purchased for a cloud disk.<br>This parameter is only valid for Enhanced SSD (CLOUD_HSSD) and ulTra SSD (CLOUD_TSSD).
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). The monthly subscription cloud disk purchase attributes such as usage period and whether or not auto-renewal is set up can be specified using this parameter. <br>This parameter is required when creating a prepaid cloud disk. This parameter is not required when creating an hourly postpaid cloud disk.
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// Specifies the cloud disk backup point quota.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type InquiryPriceCreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// Cloud disk billing mode. <br><li>POSTPAID_BY_HOUR: Hourly pay-as-you-go.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD Cloud Storage<br><li>CLOUD_PREMIUM: Premium Cloud Disk<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: ulTra SSD.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk size in GB. For the value range, see [Cloud Disk Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of the project to which the cloud disk belongs.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Number of cloud disks to be purchased. If it is not specified, `1` will be used by default.
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// Extra performance in MB/s purchased for a cloud disk.<br>This parameter is only valid for Enhanced SSD (CLOUD_HSSD) and ulTra SSD (CLOUD_TSSD).
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). The monthly subscription cloud disk purchase attributes such as usage period and whether or not auto-renewal is set up can be specified using this parameter. <br>This parameter is required when creating a prepaid cloud disk. This parameter is not required when creating an hourly postpaid cloud disk.
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// Specifies the cloud disk backup point quota.
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

func (r *InquiryPriceCreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskChargeType")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "ProjectId")
	delete(f, "DiskCount")
	delete(f, "ThroughputPerformance")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDisksResponseParams struct {
	// Describes the price of newly purchased cloud disks.
	DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceCreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateDisksResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeDiskRequestParams struct {
	// ID of the cloud disk, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk size after scale out (in GB). This cannot be smaller than the current size of the cloud disk. For the value range of the cloud disk sizes, see cloud disk [Product Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of project the cloud disk belongs to. If selected, it can only be used for authentication.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type InquiryPriceResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cloud disk, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk size after scale out (in GB). This cannot be smaller than the current size of the cloud disk. For the value range of the cloud disk sizes, see cloud disk [Product Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of project the cloud disk belongs to. If selected, it can only be used for authentication.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *InquiryPriceResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskSize")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeDiskResponseParams struct {
	// Describes the price of expanding the cloud disk.
	DiskPrice *PrepayPrice `json:"DiskPrice,omitempty" name:"DiskPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResizeDiskResponseParams `json:"Response"`
}

func (r *InquiryPriceResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoSnapshotPolicyAttributeRequestParams struct {
	// Scheduled snapshot policy ID.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// Whether or not the scheduled snapshot policy is activated. FALSE: Not activated. TRUE: Activated. The default value is TRUE.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently. FALSE: Not retained permanently. TRUE: Retained permanently. The default value is FALSE.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The name of the scheduled snapshot policy to be created. If it is left empty, the default is 'Not named'. The maximum length cannot exceed 60 bytes.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// Number of days to retain the snapshots created according to this scheduled snapshot policy. If this parameter is specified, `IsPermanent` cannot be specified as `TRUE`; otherwise, they will conflict with each other.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

type ModifyAutoSnapshotPolicyAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Scheduled snapshot policy ID.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// Whether or not the scheduled snapshot policy is activated. FALSE: Not activated. TRUE: Activated. The default value is TRUE.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently. FALSE: Not retained permanently. TRUE: Retained permanently. The default value is FALSE.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The name of the scheduled snapshot policy to be created. If it is left empty, the default is 'Not named'. The maximum length cannot exceed 60 bytes.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// Number of days to retain the snapshots created according to this scheduled snapshot policy. If this parameter is specified, `IsPermanent` cannot be specified as `TRUE`; otherwise, they will conflict with each other.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoSnapshotPolicyAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "IsActivated")
	delete(f, "IsPermanent")
	delete(f, "AutoSnapshotPolicyName")
	delete(f, "Policy")
	delete(f, "RetentionDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoSnapshotPolicyAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoSnapshotPolicyAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAutoSnapshotPolicyAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoSnapshotPolicyAttributeResponseParams `json:"Response"`
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoSnapshotPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskAttributesRequestParams struct {
	// IDs of one or more cloud disks to be operated. If multiple cloud disk IDs are selected, it only supports modifying all cloud disks with the same attributes.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// The new project ID of the cloud disk. Only the project ID of elastic cloud disk can be modified. The available projects and their IDs can be queried via the API [DescribeProject](https://intl.cloud.tencent.com/document/api/378/4400?from_cn_redirect=1).
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Name of new cloud disk.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// Whether it is an elastic cloud disk. FALSE: non-elastic cloud disk; TRUE: elastic cloud disk. You can only modify non-elastic cloud disks to elastic cloud disks.
	Portable *bool `json:"Portable,omitempty" name:"Portable"`

	// Whether the cloud disk is terminated with the CVM after it has been successfully mounted. `TRUE` indicates that it is terminated with the CVM. `FALSE` indicates that it is not terminated with the CVM. This is only supported for cloud disks and data disks that are pay-as-you-go.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// When changing the type of a cloud disk, this parameter can be passed to indicate the desired cloud disk type. Value range: <br><li>CLOUD_PREMIUM: Premium cloud storage.  <br><li>CLOUD_SSD: SSD cloud disk. <br>Currently, batch operations are not supported for changing type. That is, when `DiskType` is passed, only one cloud disk can be passed through `DiskIds`. <br>When the cloud disk type is changed, the changing of other attributes is not supported concurrently.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

type ModifyDiskAttributesRequest struct {
	*tchttp.BaseRequest
	
	// IDs of one or more cloud disks to be operated. If multiple cloud disk IDs are selected, it only supports modifying all cloud disks with the same attributes.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// The new project ID of the cloud disk. Only the project ID of elastic cloud disk can be modified. The available projects and their IDs can be queried via the API [DescribeProject](https://intl.cloud.tencent.com/document/api/378/4400?from_cn_redirect=1).
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Name of new cloud disk.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// Whether it is an elastic cloud disk. FALSE: non-elastic cloud disk; TRUE: elastic cloud disk. You can only modify non-elastic cloud disks to elastic cloud disks.
	Portable *bool `json:"Portable,omitempty" name:"Portable"`

	// Whether the cloud disk is terminated with the CVM after it has been successfully mounted. `TRUE` indicates that it is terminated with the CVM. `FALSE` indicates that it is not terminated with the CVM. This is only supported for cloud disks and data disks that are pay-as-you-go.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// When changing the type of a cloud disk, this parameter can be passed to indicate the desired cloud disk type. Value range: <br><li>CLOUD_PREMIUM: Premium cloud storage.  <br><li>CLOUD_SSD: SSD cloud disk. <br>Currently, batch operations are not supported for changing type. That is, when `DiskType` is passed, only one cloud disk can be passed through `DiskIds`. <br>When the cloud disk type is changed, the changing of other attributes is not supported concurrently.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

func (r *ModifyDiskAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "ProjectId")
	delete(f, "DiskName")
	delete(f, "Portable")
	delete(f, "DeleteWithInstance")
	delete(f, "DiskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskAttributesResponseParams `json:"Response"`
}

func (r *ModifyDiskAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskBackupQuotaRequestParams struct {
	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk backup point quota after the adjustment
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

type ModifyDiskBackupQuotaRequest struct {
	*tchttp.BaseRequest
	
	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk backup point quota after the adjustment
	DiskBackupQuota *uint64 `json:"DiskBackupQuota,omitempty" name:"DiskBackupQuota"`
}

func (r *ModifyDiskBackupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskBackupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskBackupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskBackupQuotaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDiskBackupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskBackupQuotaResponseParams `json:"Response"`
}

func (r *ModifyDiskBackupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskBackupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskExtraPerformanceRequestParams struct {
	// ID of the cloud disk to create a snapshot, which can be obtained via the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// The extra throughput to purchase, in MB/s
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

type ModifyDiskExtraPerformanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cloud disk to create a snapshot, which can be obtained via the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// The extra throughput to purchase, in MB/s
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *ModifyDiskExtraPerformanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "ThroughputPerformance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiskExtraPerformanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiskExtraPerformanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiskExtraPerformanceResponseParams `json:"Response"`
}

func (r *ModifyDiskExtraPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotAttributeRequestParams struct {
	// Snapshot ID, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Name of new snapshot. Maximum length is 60 bytes.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Snapshot retention mode. Valid values: `FALSE`: non-permanent retention; `TRUE`: permanent retention.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// Expiration time of the snapshot. Setting this parameter will set the snapshot retention mode to `FALSE` (non-permanent retention) and the snapshot will be automatically deleted upon expiration.
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot ID, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Name of new snapshot. Maximum length is 60 bytes.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Snapshot retention mode. Valid values: `FALSE`: non-permanent retention; `TRUE`: permanent retention.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// Expiration time of the snapshot. Setting this parameter will set the snapshot retention mode to `FALSE` (non-permanent retention) and the snapshot will be automatically deleted upon expiration.
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
}

func (r *ModifySnapshotAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	delete(f, "IsPermanent")
	delete(f, "Deadline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySnapshotAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotAttributeResponseParams `json:"Response"`
}

func (r *ModifySnapshotAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsSharePermissionRequestParams struct {
	// List of account IDs with which a snapshot is shared. For the format of array-type parameters, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1). You can find the account ID in [Account Information](https://console.cloud.tencent.com/developer).
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`

	// Operations. Valid values: `SHARE`, sharing an image; `CANCEL`, cancelling the sharing of an image.
	Permission *string `json:"Permission,omitempty" name:"Permission"`

	// The ID of the snapshot. You can obtain this by using [DescribeSnapshots](https://intl.cloud.tencent.com/document/api/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

type ModifySnapshotsSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// List of account IDs with which a snapshot is shared. For the format of array-type parameters, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1). You can find the account ID in [Account Information](https://console.cloud.tencent.com/developer).
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`

	// Operations. Valid values: `SHARE`, sharing an image; `CANCEL`, cancelling the sharing of an image.
	Permission *string `json:"Permission,omitempty" name:"Permission"`

	// The ID of the snapshot. You can obtain this by using [DescribeSnapshots](https://intl.cloud.tencent.com/document/api/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *ModifySnapshotsSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountIds")
	delete(f, "Permission")
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotsSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsSharePermissionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySnapshotsSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotsSharePermissionResponseParams `json:"Response"`
}

func (r *ModifySnapshotsSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {
	// The ID of the [Availability Zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) to which the cloud disk belongs. This parameter can be obtained from the Zone field in the returned values of [DescribeZones](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1).
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Cage ID. When it is an input parameter, the specified CageID resource is operated, and it can be left blank. When it is an output parameter, it is the ID of the cage the resource belongs to, and it can be left blank.
	// Note: This field may return null, indicating that no valid value was found.
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the projectId field in the returned values of [DescribeProject](https://intl.cloud.tencent.com/document/api/378/4400?from_cn_redirect=1). If this is left empty, default project is used.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Dedicated cluster name. When it is an input parameter, it is ignored.  When it is an output parameter, it is the name of the dedicated cluster the cloud disk belongs to, and it can be left blank.
	// Note: This field may return null, indicating that no valid value was found.
	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`

	// ID of dedicated cluster which the instance belongs to. When it is an input parameter, the specified CdcId dedicated cluster resource is operated, and it can be left blank. When it is an output parameter, it is the ID of the dedicated cluster which the resource belongs to, and it can be left blank.
	// Note: This field may return null, indicating that no valid value was found.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`

	// Dedicated cluster ID
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

type Policy struct {
	// Specifies the time that that the scheduled snapshot policy will be triggered. The unit is hour. The value range is [0-23]. 00:00-23:00 is a total of 24 time points that can be selected. 1 indicates 01:00, and so on.
	Hour []*uint64 `json:"Hour,omitempty" name:"Hour"`

	// Specifies the days of the week, from Monday to Sunday, on which a scheduled snapshot will be triggered. Value range: [0, 6]. 0 indicates triggering on Sunday, 1-6 indicate triggering on Monday-Saturday.
	DayOfWeek []*uint64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// Specifies the dates of the month on which a scheduled snapshot will be triggered. Value range: [1, 31]. `1` to `31` indicate the specific dates of the month; for example, `5` indicates the 5th day of the month. Note: If you set a date that does not exist in some months such as 29, 30, and 31, these months will be skipped for scheduled snapshot creation.
	DayOfMonth []*uint64 `json:"DayOfMonth,omitempty" name:"DayOfMonth"`

	// Specifies the interval for creating scheduled snapshots in days. Value range: [1, 365]. For example, if it is set to `5`, scheduled snapshots will be created every 5 days. Note: If you choose to back up by day, the time for the first backup is theoretically the day when the backup policy is created. If the backup policy creation time on the current day is later than the set backup time, the first backup will be performed in the second backup cycle.
	IntervalDays *uint64 `json:"IntervalDays,omitempty" name:"IntervalDays"`
}

type PrepayPrice struct {
	// Original payment of a monthly-subscribed cloud disk or a snapshot, in USD.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discounted price of a monthly-subscribed cloud disk or a snapshot, in USD.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Original payment of a monthly-subscribed cloud disk or a snapshot, in USD, with six decimal places.
	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitempty" name:"OriginalPriceHigh"`

	// Discounted price of a monthly-subscribed cloud disk or a snapshot, in USD, with six decimal places.
	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitempty" name:"DiscountPriceHigh"`

	// Original unit price of a pay-as-you-go cloud disk, in USD.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Billing unit for pay-as-you-go cloud disks. Valid value: <br><li>HOUR: billed hourly.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// Discount unit price of a pay-as-you-go cloud disk, in USD.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// Original unit price of a pay-as-you-go cloud disk, in USD, with six decimal places.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPriceHigh *string `json:"UnitPriceHigh,omitempty" name:"UnitPriceHigh"`

	// Discounted unit price of a pay-as-you-go cloud disk, in USD, with six decimal places.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitempty" name:"UnitPriceDiscountHigh"`
}

type Price struct {
	// Original price of a monthly-subscribed cloud disk, in USD.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discounted price of a monthly-subscribed cloud disk, in USD.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Original unit price of a pay-as-you-go cloud disk, in USD.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Billing unit of a postpaid cloud disk. Value range: <br><li>HOUR: Billed by hour.
	// Note: This field may return null, indicating that no valid value was found.
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// Discount unit price of a pay-as-you-go cloud disk, in USD.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// Original payment of a monthly-subscribed cloud disk, in USD, with six decimal places.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OriginalPriceHigh *string `json:"OriginalPriceHigh,omitempty" name:"OriginalPriceHigh"`

	// Discounted payment price of a monthly-subscribed cloud disk, in USD, with six decimal places.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DiscountPriceHigh *string `json:"DiscountPriceHigh,omitempty" name:"DiscountPriceHigh"`

	// Original unit price of a pay-as-you-go cloud disk, in USD, with six decimal places.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPriceHigh *string `json:"UnitPriceHigh,omitempty" name:"UnitPriceHigh"`

	// Discounted unit price of a pay-as-you-go cloud disk, in USD, with six decimal places.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnitPriceDiscountHigh *string `json:"UnitPriceDiscountHigh,omitempty" name:"UnitPriceDiscountHigh"`
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// ID of the cloud disk, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk size after scale out (in GB). This must be larger than the current size of the cloud disk. For the value range of the cloud disk sizes, see cloud disk [Product Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cloud disk, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk size after scale out (in GB). This must be larger than the current size of the cloud disk. For the value range of the cloud disk sizes, see cloud disk [Product Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDiskResponseParams `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SharePermission struct {
	// Snapshot sharing time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// ID of the shared account
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
}

type Snapshot struct {
	// Location of the snapshot.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Whether the snapshot is replicated across regions. Value range: <br><li>true: Indicates that the snapshot is replicated across regions. <br><li>false: Indicates that the snapshot belongs to the local region.
	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`

	// Snapshot status. Valid values: <br><li>NORMAL: normal <br><li>CREATING: creating<br><li>ROLLBACKING: rolling back<br><li>COPYING_FROM_REMOTE: cross-region replicating<li>CHECKING_COPIED: verifying the cross-region replicated data<br><li>TORECYCLE: to be repossessed.
	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`

	// Whether it is a permanent snapshot. Value range: <br><li>true: Permanent snapshot <br><li>false: Non-permanent snapshot.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// Snapshot name, the user-defined snapshot alias. Call [ModifySnapshotAttribute](https://intl.cloud.tencent.com/document/product/362/15650?from_cn_redirect=1) to modify this field.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// The expiration time of the snapshot. If the snapshot is permanently retained, this field is blank.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// The progress percentage for snapshot creation. This field is always 100 after the snapshot is created successfully.
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// List of images associated with snapshot.
	Images []*Image `json:"Images,omitempty" name:"Images"`

	// Number of snapshots currently shared
	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`

	// Snapshot type. This value can currently be either PRIVATE_SNAPSHOT or SHARED_SNAPSHOT.
	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`

	// Size of the cloud disk used to create this snapshot (in GB).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of the cloud disk used to create this snapshot.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// The destination region to which the snapshot is being replicated. Default value is [ ].
	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`

	// Whether the snapshot is created from an encrypted disk. Value range: <br><li>true: Yes <br><li>false: No.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// Creation time of the snapshot.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Number of images associated with snapshot.
	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`

	// The type of the cloud disk used to create the snapshot. Value range: <br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// The time when the snapshot sharing starts
	TimeStartShare *string `json:"TimeStartShare,omitempty" name:"TimeStartShare"`

	// List of tags associated with the snapshot.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type SnapshotCopyResult struct {
	// ID of the snapshot replica
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Error message. It’s null if the request succeeds.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Error code. It’s `Success` if the request succeeds.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Destination region of the replication task
	DestinationRegion *string `json:"DestinationRegion,omitempty" name:"DestinationRegion"`
}

type SnapshotOperationLog struct {
	// Status of operation. Value range:
	// SUCCESS: Operation successful 
	// FAILED: Operation failed 
	// PROCESSING: Operation in process
	OperationState *string `json:"OperationState,omitempty" name:"OperationState"`

	// Start time
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// UIN of operator.
	// Note: This field may return null, indicating that no valid value was found.
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// ID of snapshot being operated.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Operation type. Value range:
	// SNAP_OPERATION_DELETE: Delete snapshot
	// SNAP_OPERATION_ROLLBACK: Roll back snapshot
	// SNAP_OPERATION_MODIFY: Modify snapshot attributes
	// SNAP_OPERATION_CREATE: Create snapshot
	// SNAP_OPERATION_COPY: Cross-region replication of snapshot
	// ASP_OPERATION_CREATE_SNAP: Create snapshot with scheduled snapshot policy
	// ASP_OPERATION_DELETE_SNAP: Delete snapshot from scheduled snapshot policy
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// End time
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Tag struct {
	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateDisksRequestParams struct {
	// List of cloud disk IDs required to be returned.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Delete the associated non-permanently reserved snapshots upon deletion of the source cloud disk. `0`: No (default). `1`: Yes. To check whether a snapshot is permanently reserved, refer to the `IsPermanent` field returned by the `DescribeSnapshots` API. 
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitempty" name:"DeleteSnapshot"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs required to be returned.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// Delete the associated non-permanently reserved snapshots upon deletion of the source cloud disk. `0`: No (default). `1`: Yes. To check whether a snapshot is permanently reserved, refer to the `IsPermanent` field returned by the `DescribeSnapshots` API. 
	DeleteSnapshot *int64 `json:"DeleteSnapshot,omitempty" name:"DeleteSnapshot"`
}

func (r *TerminateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DeleteSnapshot")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateDisksResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDisksResponseParams `json:"Response"`
}

func (r *TerminateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyRequestParams struct {
	// List of cloud disk IDs scheduled snapshot policy to be unbound from.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// ID of scheduled snapshot policy to be unbound.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs scheduled snapshot policy to be unbound from.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// ID of scheduled snapshot policy to be unbound.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
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
	delete(f, "DiskIds")
	delete(f, "AutoSnapshotPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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