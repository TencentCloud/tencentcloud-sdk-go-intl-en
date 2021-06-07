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
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ApplySnapshotRequest struct {
	*tchttp.BaseRequest

	// Snapshot ID, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// ID of the original cloud disk corresponding to the snapshot, which can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (r *ApplySnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplySnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "DiskId")
	if len(f) > 0 {
		return errors.New("ApplySnapshotRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplySnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplySnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("AttachDisksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoSnapshotPolicy struct {

	// Scheduled snapshot policy ID.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// Scheduled snapshot policy name.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// Scheduled snapshot policy state. Value range:<br><li>NORMAL: Normal<br><li>ISOLATED: Isolated.
	AutoSnapshotPolicyState *string `json:"AutoSnapshotPolicyState,omitempty" name:"AutoSnapshotPolicyState"`

	// Whether scheduled snapshot policy is activated.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// Number of days the snapshot created by this scheduled snapshot policy is retained.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`

	// The time the scheduled snapshot policy was created.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The time the scheduled snapshot will be triggered again.
	NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`

	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// The list of cloud disk IDs that the current scheduled snapshot policy is bound to.
	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "DiskIds")
	if len(f) > 0 {
		return errors.New("BindAutoSnapshotPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateAutoSnapshotPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The ID of the newly created scheduled snapshot policy.
		AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

		// The time that initial backup will start.
		NextTriggerTime *string `json:"NextTriggerTime,omitempty" name:"NextTriggerTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: Tremendous SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk billing method. POSTPAID_BY_HOUR: pay as you go by hour<br><li>CDCPAID: Billed together with the bound dedicated cluster<br>For information about the pricing of each method, see the cloud disk [Pricing Overview](https://intl.cloud.tencent.com/document/product/362/2413?from_cn_redirect=1).
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// The location of the instance. The availability zone and the project that the instance belongs to can be specified using this parameter. If the project is not specified, it will be created under the default project.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The displayed name of the cloud disk. If it is left empty, the default is 'Not named'. The maximum length cannot exceed 60 bytes.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// If the number of cloud disks to be created is left empty, the default is 1. There is a limit to the maximum number of cloud disks that can be created for a single request. For more information, please see [CBS Use Limits](https://intl.cloud.tencent.com/doc/product/362/5145?from_cn_redirect=1).
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). The monthly subscription cloud disk purchase attributes such as usage period and whether or not auto-renewal is set up can be specified using this parameter. <br>This parameter is required when creating a prepaid cloud disk. This parameter is not required when creating an hourly postpaid cloud disk. 
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// Cloud hard disk size (in GB). <br><li> If `SnapshotId` is passed, `DiskSize` cannot be passed. In this case, the size of the cloud disk is the size of the snapshot. <br><li>To pass `SnapshotId` and `DiskSize` at the same time, the size of the disk must be larger than or equal to the size of the snapshot. <br><li>For information about the size range of cloud disks, see cloud disk [Product Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Snapshot ID. If this parameter is specified, the cloud disk is created based on the snapshot. The snapshot type must be a data disk snapshot. The snapshot can be queried in the DiskUsage field in the output parameter through the API [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// A string to ensure the idempotency of the request, which is generated by the client. Each request shall have a unique string with a maximum of 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be ensured.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// This parameter is used to create an encrypted cloud disk. Its value is always ENCRYPT.
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// Cloud disk binding tag.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// The default of optional parameter is False. When True is selected, the cloud disk will be created as a shareable cloud disk.
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// Extra performance purchased for a cloud disk.<br>This optional parameter is only valid for Tremendous SSD (CLOUD_TSSD) and Enhanced SSD (CLOUD_HSSD).
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *CreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskType")
	delete(f, "DiskChargeType")
	delete(f, "Placement")
	delete(f, "DiskName")
	delete(f, "DiskCount")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskSize")
	delete(f, "SnapshotId")
	delete(f, "ClientToken")
	delete(f, "Encrypt")
	delete(f, "Tags")
	delete(f, "Shareable")
	delete(f, "ThroughputPerformance")
	if len(f) > 0 {
		return errors.New("CreateDisksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of created cloud disk IDs. 
		DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest

	// ID of the cloud disk, for which a snapshot needs to be created. It can be queried via the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Snapshot name. If it is left empty, the new snapshot name is "Not named" by default.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Expiration time of the snapshot. The snapshot will be automatically deleted upon expiration.
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "SnapshotName")
	delete(f, "Deadline")
	if len(f) > 0 {
		return errors.New("CreateSnapshotRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the new snapshot.
		SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyIds")
	if len(f) > 0 {
		return errors.New("DeleteAutoSnapshotPoliciesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest

	// List of IDs of snapshots to be deleted, which can be queried via [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return errors.New("DeleteSnapshotsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeAutoSnapshotPoliciesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The quantity of valid scheduled snapshot policies.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of scheduled snapshot policies.
		AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskAssociatedAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	if len(f) > 0 {
		return errors.New("DescribeDiskAssociatedAutoSnapshotPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskAssociatedAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The quantity of scheduled snapshots binded to cloud disk.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of scheduled snapshots bound to cloud disk.
		AutoSnapshotPolicySet []*AutoSnapshotPolicy `json:"AutoSnapshotPolicySet,omitempty" name:"AutoSnapshotPolicySet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskAssociatedAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeDiskConfigQuotaRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskConfigQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of cloud disk configurations.
		DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitempty" name:"DiskConfigSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskConfigQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeDiskOperationLogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDiskOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of cloud disk operation logs.
		DiskOperationLogSet []*DiskOperationLog `json:"DiskOperationLogSet,omitempty" name:"DiskOperationLogSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDiskOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeDisksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The quantity of cloud disks meeting the conditions.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of cloud disk details.
		DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDiskNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return errors.New("DescribeInstancesDiskNumRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDiskNumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The quantity of mounted and mountable elastic cloud disks for each cloud virtual machine
		AttachDetail []*AttachDetail `json:"AttachDetail,omitempty" name:"AttachDetail"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDiskNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDiskNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeSnapshotOperationLogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of snapshot operation logs.
		SnapshotOperationLogSet []*SnapshotOperationLog `json:"SnapshotOperationLogSet,omitempty" name:"SnapshotOperationLogSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return errors.New("DescribeSnapshotSharePermissionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The set of snapshot sharing information
		SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeSnapshotsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of snapshots.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of snapshot details.
		SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return errors.New("DetachDisksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {

	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Cloud disk type. Value range:<br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Billing method. Value range: <br><li>PREPAID: Prepaid, that is, monthly subscription<br><li>POSTPAID_BY_HOUR: Postpaid, that is, pay as you go.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Whether it is an elastic cloud disk. false: Non-elastic cloud disk; true: Elastic cloud disk.
	Portable *bool `json:"Portable,omitempty" name:"Portable"`

	// Location of the cloud disk.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Whether the cloud disk has the capability to create snapshots. Value range: <br><li>false: Cannot create snapshots. true: Can create snapshots.
	SnapshotAbility *bool `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`

	// Cloud disk name.
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// Cloud disk size (in GB).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// The state of the cloud disk. Value range: <br><li>UNATTACHED: Not mounted <br><li>ATTACHING: Mounting <br><li>ATTACHED: Mounted <br><li>DETACHING: Un-mounting <br><li>EXPANDING: Expanding <br><li>ROLLBACKING: Rolling back <br><li>TORECYCE: Pending recycling. <br><li>DUMPING: Copying the hard drive.
	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: Tremendous SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Whether the cloud disk is mounted to the CVM. Value range: <br><li>false: Unmounted <br><li>true: Mounted.
	Attached *bool `json:"Attached,omitempty" name:"Attached"`

	// ID of the CVM to which the cloud disk is mounted.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Creation time of the cloud disk.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Expiration time of the cloud disk.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Whether the cloud disk is in the status of snapshot rollback. Value range: <br><li>false: No <br><li>true: Yes
	Rollbacking *bool `json:"Rollbacking,omitempty" name:"Rollbacking"`

	// Rollback progress of a cloud disk snapshot.
	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`

	// Whether the cloud disk is encrypted. Value range: <br><li>false: Not encrypted <br><li>true: Encrypted.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// Cloud disk already mounted to CVM, and both CVM and cloud disk use monthly subscription.<br><li>true: CVM has auto-renewal flag set up, but cloud disk does not.<br><li>false: Cloud disk auto-renewal flag set up normally.
	// Note: This field may return null, indicating that no valid value was found.
	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`

	// Auto renewal flag. Supported values:<br><li>NOTIFY_AND_AUTO_RENEW: Notify expiry and renew automatically<br><li>NOTIFY_AND_MANUAL_RENEW: Notify expiry but not renew automatically<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Neither notify expiry nor renew automatically.
	// Note: This field may return null, indicating that no valid value was found.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// This field is only applicable when the instance is already mounted to the cloud disk, and both the instance and the cloud disk use monthly subscription. <br><li>true: Expiration time of cloud disk is earlier than that of the instance.<br><li>false:Expiration time of cloud disk is later than that of the instance.
	// Note: This field may return null, indicating that no valid value was found.
	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`

	// Determines whether or not prepaid cloud disk supports active return. <br><li>true: Active return supported.<br><li>false: Active return not supported.
	// Note: This field may return null, indicating that no valid value was found.
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// In circumstances where the prepaid cloud disk does not support active return, this parameter indicates the reason that return is not supported. Value range: <br><li>1: The cloud disk has already been returned. <br><li>2: The cloud disk has already expired. <br><li>3: The cloud disk does not support return. <br><li> 8: The limit on the number of returns is exceeded.
	// Note: This field may return null, indicating that no valid value was found.
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// ID of the periodic snapshot associated to the cloud disk. This parameter is returned only if the value of parameter ReturnBindAutoSnapshotPolicy is TRUE when the API DescribeDisks is called.
	// Note: This field may return null, indicating that no valid value was found.
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`

	// The tag bound to the cloud disk. The value Null is used when no tag is bound to the cloud disk.
	// Note: This field may return null, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether the cloud disk terminates along with the instance mounted to it. <br><li>true: Cloud disk will also be terminated when instance terminates, so only hourly postpaid cloud disk are supported.<br><li>false: Cloud disk does not terminate when instance terminates.
	// Note: This field may return null, indicating that no valid value was found.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// Number of days from current time until disk expiration (only applicable for prepaid disks).
	// Note: This field may return null, indicating that no valid value was found.
	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`

	// Whether cloud disk is in process of type change. Value range: <br><li>false: Cloud disk not in process of type change. <br><li>true: Cloud disk type change has been launched, and migration is in process.
	// Note: This field may return null, indicating that no valid value was found.
	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`

	// Migration progress of cloud disk type change, from 0 to 100.
	// Note: This field may return null, indicating that no valid value was found.
	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`

	// Whether or not cloud disk is shareable cloud disk.
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// For non-shareable cloud disks, this parameter is null. For shareable cloud disks, this parameters indicates this cloud disk's Instance IDs currently mounted to the CVM.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// The total number of snapshots of the cloud disk.
	SnapshotCount *int64 `json:"SnapshotCount,omitempty" name:"SnapshotCount"`

	// The total capacity of the snapshots of the cloud disk. Unit: MB.
	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`

	// Specifies whether to create a snapshot when the cloud disk is terminated due to overdue payment or expiration. `true`: create snapshot; `false`: do not create snapshot.
	BackupDisk *bool `json:"BackupDisk,omitempty" name:"BackupDisk"`

	// Extra performance for a cloud disk, in MB/sec.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

type DiskChargePrepaid struct {

	// The purchased usage period of a cloud disk (in months). Value range: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Auto Renewal flag. Value range: <br><li>NOTIFY_AND_AUTO_RENEW: Notify expiry and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: Notify expiry but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: Neither notify expiry nor renew automatically <br><br>Default value range: NOTIFY_AND_MANUAL_RENEW: Notify expiry but do not renew automatically.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// This parameter is used when you align the expiration time of the cloud disk with that of the mounted server. It is the current expiration time of the server. In this case, the Period passed represents the renewal period of the server, and the cloud disk will be automatically renewed in alignment with the expiration time of the renewed server. Example value: 2018-03-30 20:15:03.
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitempty" name:"CurInstanceDeadline"`
}

type DiskConfig struct {

	// Whether the configuration is available.
	Available *bool `json:"Available,omitempty" name:"Available"`

	// Type of cloud disk medium. Value range: <br><li>CLOUD_BASIC: Ordinary cloud disk <br><li>CLOUD_PREMIUM: Premium cloud storage <br><li>CLOUD_SSD: SSD cloud disk.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk type. Value range: <br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Billing method. Value range: <br><li>PREPAID: Prepaid, that is, monthly subscription<br><li>POSTPAID_BY_HOUR: Postpaid, that is, pay as you go.
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// The maximum configurable cloud disk size (in GB).
	MaxDiskSize *uint64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// The minimum configurable cloud disk size (in GB).
	MinDiskSize *uint64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`

	// The [Availability Region](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) of the cloud drive.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model.
	// Note: This field may return null, indicating that no valid value was found.
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// Instance model series. For more information, please see [Instance Models](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1)
	// Note: This field may return null, indicating that no valid value was found.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
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

type GetSnapOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *GetSnapOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSnapOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("GetSnapOverviewRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSnapOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *GetSnapOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "ThroughputPerformance")
	if len(f) > 0 {
		return errors.New("InquirePriceModifyDiskExtraPerformanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price for purchasing the extra performance
		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceModifyDiskExtraPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDisksRequest struct {
	*tchttp.BaseRequest

	// Cloud disk media type. Valid values: <br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: Tremendous SSD.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Cloud disk size (in GB). For the value range of the cloud disk sizes, see cloud disk [Product Types](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Cloud disk billing method. <br><li>POSTPAID_BY_HOUR: Pay-as-you-go on an hourly basis
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// Relevant parameter settings for the prepaid mode (i.e., monthly subscription). The monthly subscription cloud disk purchase attributes such as usage period and whether or not auto-renewal is set up can be specified using this parameter. <br>This parameter is required when creating a prepaid cloud disk. This parameter is not required when creating an hourly postpaid cloud disk.
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// Quantity of cloud disks purchased. If left empty, default is 1.
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// ID of project the cloud disk belongs to.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Extra performance (in MB/sec) purchased for a cloud disk.<br>This parameter is only valid for Enhanced SSD (CLOUD_HSSD) and Tremendous SSD (CLOUD_TSSD).
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
}

func (r *InquiryPriceCreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "DiskChargeType")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskCount")
	delete(f, "ProjectId")
	delete(f, "ThroughputPerformance")
	if len(f) > 0 {
		return errors.New("InquiryPriceCreateDisksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Describes the price of purchasing new cloud disk.
		DiskPrice *Price `json:"DiskPrice,omitempty" name:"DiskPrice"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("InquiryPriceResizeDiskRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Describes the price of expanding the cloud disk.
		DiskPrice *PrepayPrice `json:"DiskPrice,omitempty" name:"DiskPrice"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoSnapshotPolicyAttributeRequest struct {
	*tchttp.BaseRequest

	// Scheduled snapshot policy ID.
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// The policy for executing the scheduled snapshot.
	Policy []*Policy `json:"Policy,omitempty" name:"Policy"`

	// The name of the scheduled snapshot policy to be created. If it is left empty, the default is 'Not named'. The maximum length cannot exceed 60 bytes.
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" name:"AutoSnapshotPolicyName"`

	// Whether or not the scheduled snapshot policy is activated. FALSE: Not activated. TRUE: Activated. The default value is TRUE.
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// Whether the snapshot created by this scheduled snapshot policy is retained permanently. FALSE: Not retained permanently. TRUE: Retained permanently. The default value is FALSE.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The number of days for which snapshots created by this policy are retained. This parameter cannot clash with `IsPermanent`, which is, if the scheduled snapshot policy is configured to retain permanently, `RetentionDays` must be 0.
	RetentionDays *uint64 `json:"RetentionDays,omitempty" name:"RetentionDays"`
}

func (r *ModifyAutoSnapshotPolicyAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoSnapshotPolicyAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "Policy")
	delete(f, "AutoSnapshotPolicyName")
	delete(f, "IsActivated")
	delete(f, "IsPermanent")
	delete(f, "RetentionDays")
	if len(f) > 0 {
		return errors.New("ModifyAutoSnapshotPolicyAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoSnapshotPolicyAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoSnapshotPolicyAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoSnapshotPolicyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyDiskAttributesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskExtraPerformanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "ThroughputPerformance")
	if len(f) > 0 {
		return errors.New("ModifyDiskExtraPerformanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDiskExtraPerformanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDiskExtraPerformanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiskExtraPerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifySnapshotAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifySnapshotsSharePermissionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotsSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySnapshotsSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Placement struct {

	// The ID of the [Availability Zone](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#ZoneInfo) to which the cloud disk belongs. This parameter can be obtained from the Zone field in the returned values of [DescribeZones](https://intl.cloud.tencent.com/document/product/213/15707?from_cn_redirect=1).
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ID of the project to which the instance belongs. This parameter can be obtained from the projectId field in the returned values of [DescribeProject](https://intl.cloud.tencent.com/document/api/378/4400?from_cn_redirect=1). If this is left empty, default project is used.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID of dedicated cluster which the instance belongs to. When it is an input parameter, the specified CdcId dedicated cluster resource is operated, and it can be left blank. When it is an output parameter, it is the ID of the dedicated cluster which the resource belongs to, and it can be left blank.
	// Note: This field may return null, indicating that no valid value was found.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`

	// Cage ID. When it is an input parameter, the specified CageID resource is operated, and it can be left blank. When it is an output parameter, it is the ID of the cage the resource belongs to, and it can be left blank.
	// Note: This field may return null, indicating that no valid value was found.
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// Dedicated cluster name. When it is an input parameter, it is ignored.  When it is an output parameter, it is the name of the dedicated cluster the cloud disk belongs to, and it can be left blank.
	// Note: This field may return null, indicating that no valid value was found.
	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`
}

type Policy struct {

	// Specifies the days of the week, from Monday to Sunday, on which a scheduled snapshot will be triggered. Value range: [0, 6]. 0 indicates triggering on Sunday, 1-6 indicate triggering on Monday-Saturday.
	DayOfWeek []*uint64 `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// Specifies the time that that the scheduled snapshot policy will be triggered. The unit is hour. The value range is [0-23]. 00:00-23:00 is a total of 24 time points that can be selected. 1 indicates 01:00, and so on.
	Hour []*uint64 `json:"Hour,omitempty" name:"Hour"`
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskId")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return errors.New("ResizeDiskRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Location of the snapshot.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The type of the cloud disk used to create the snapshot. Value range: <br><li>SYSTEM_DISK: System disk <br><li>DATA_DISK: Data disk.
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// ID of the cloud disk used to create this snapshot.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Size of the cloud disk used to create this snapshot (in GB).
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Snapshot status. Valid values: <br><li>NORMAL: normal <br><li>CREATING: creating<br><li>ROLLBACKING: rolling back<br><li>COPYING_FROM_REMOTE: cross-region replicating<li>CHECKING_COPIED: verifying the cross-region replicated data<br><li>TORECYCLE: to be repossessed.
	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`

	// Snapshot name, the user-defined snapshot alias. Call [ModifySnapshotAttribute](https://intl.cloud.tencent.com/document/product/362/15650?from_cn_redirect=1) to modify this field.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// The progress percentage for snapshot creation. This field is always 100 after the snapshot is created successfully.
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// Creation time of the snapshot.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// The expiration time of the snapshot. If the snapshot is permanently retained, this field is blank.
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// Whether the snapshot is created from an encrypted disk. Value range: <br><li>true: Yes <br><li>false: No.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// Whether it is a permanent snapshot. Value range: <br><li>true: Permanent snapshot <br><li>false: Non-permanent snapshot.
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// The destination region to which the snapshot is being replicated. Default value is [ ].
	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`

	// Whether the snapshot is replicated across regions. Value range: <br><li>true: Indicates that the snapshot is replicated across regions. <br><li>false: Indicates that the snapshot belongs to the local region.
	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`

	// List of images associated with snapshot.
	Images []*Image `json:"Images,omitempty" name:"Images"`

	// Number of images associated with snapshot.
	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`

	// Snapshot type. This value can currently be either PRIVATE_SNAPSHOT or SHARED_SNAPSHOT.
	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`

	// Number of snapshots currently shared
	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`

	// The time when the snapshot sharing starts
	TimeStartShare *string `json:"TimeStartShare,omitempty" name:"TimeStartShare"`
}

type SnapshotOperationLog struct {

	// UIN of operator.
	// Note: This field may return null, indicating that no valid value was found.
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Operation type. Value range:
	// SNAP_OPERATION_DELETE: Delete snapshot
	// SNAP_OPERATION_ROLLBACK: Roll back snapshot
	// SNAP_OPERATION_MODIFY: Modify snapshot attributes
	// SNAP_OPERATION_CREATE: Create snapshot
	// SNAP_OPERATION_COPY: Cross-region replication of snapshot
	// ASP_OPERATION_CREATE_SNAP: Create snapshot with scheduled snapshot policy
	// ASP_OPERATION_DELETE_SNAP: Delete snapshot from scheduled snapshot policy
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// ID of snapshot being operated.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

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

type Tag struct {

	// Tag key.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest

	// List of cloud disk IDs required to be returned.
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *TerminateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return errors.New("TerminateDisksRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TerminateDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "AutoSnapshotPolicyId")
	if len(f) > 0 {
		return errors.New("UnbindAutoSnapshotPolicyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
