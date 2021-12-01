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

package v20200324

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ApplyInstanceSnapshotRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *ApplyInstanceSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInstanceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyInstanceSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyInstanceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInstanceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Key pair ID list. Each request can contain up to 100 key pairs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachCcnRequest struct {
	*tchttp.BaseRequest

	// CCN instance ID.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *AttachCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Blueprint struct {

	// Image ID, which is the unique identifier of `Blueprint`.
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// Image title to be displayed.
	DisplayTitle *string `json:"DisplayTitle,omitempty" name:"DisplayTitle"`

	// Image version to be displayed.
	DisplayVersion *string `json:"DisplayVersion,omitempty" name:"DisplayVersion"`

	// Image description information.
	Description *string `json:"Description,omitempty" name:"Description"`

	// OS name.
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// OS type.
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// OS type, such as LINUX_UNIX and WINDOWS.
	PlatformType *string `json:"PlatformType,omitempty" name:"PlatformType"`

	// Image type, such as APP_OS, PURE_OS, and PRIVATE.
	BlueprintType *string `json:"BlueprintType,omitempty" name:"BlueprintType"`

	// Image picture URL.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// System disk size required by image in GB.
	RequiredSystemDiskSize *int64 `json:"RequiredSystemDiskSize,omitempty" name:"RequiredSystemDiskSize"`

	// Image status.
	BlueprintState *string `json:"BlueprintState,omitempty" name:"BlueprintState"`

	// Creation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Image name.
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// Whether the image supports automation tools.
	SupportAutomationTools *bool `json:"SupportAutomationTools,omitempty" name:"SupportAutomationTools"`

	// Memory size required by image in GB.
	RequiredMemorySize *int64 `json:"RequiredMemorySize,omitempty" name:"RequiredMemorySize"`

	// ID of the Lighthouse image shared from a CVM image
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type BlueprintInstance struct {

	// Image information.
	Blueprint *Blueprint `json:"Blueprint,omitempty" name:"Blueprint"`

	// Software list.
	SoftwareSet []*Software `json:"SoftwareSet,omitempty" name:"SoftwareSet"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type BlueprintPrice struct {

	// Original image unit price in USD.
	OriginalBlueprintPrice *float64 `json:"OriginalBlueprintPrice,omitempty" name:"OriginalBlueprintPrice"`

	// Original image total price in USD.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discount.
	Discount *int64 `json:"Discount,omitempty" name:"Discount"`

	// Discounted image total price in USD.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type Bundle struct {

	// Package ID.
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// Memory size in GB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// System disk type.
	// Valid values: 
	// <li> LOCAL_BASIC: local disk</li><li> LOCAL_SSD: local SSD disk</li><li> CLOUD_BASIC: HDD cloud disk</li><li> CLOUD_SSD: SSD cloud disk</li><li> CLOUD_PREMIUM: Premium Cloud Storage</li>
	SystemDiskType *string `json:"SystemDiskType,omitempty" name:"SystemDiskType"`

	// System disk size.
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// Monthly network traffic in Gb.
	MonthlyTraffic *int64 `json:"MonthlyTraffic,omitempty" name:"MonthlyTraffic"`

	// Whether Linux/Unix is supported.
	SupportLinuxUnixPlatform *bool `json:"SupportLinuxUnixPlatform,omitempty" name:"SupportLinuxUnixPlatform"`

	// Whether Windows is supported.
	SupportWindowsPlatform *bool `json:"SupportWindowsPlatform,omitempty" name:"SupportWindowsPlatform"`

	// Current package unit price information.
	Price *Price `json:"Price,omitempty" name:"Price"`

	// Number of CPU cores.
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Peak bandwidth in Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Network billing mode.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// Package sale status. Valid values: AVAILABLE, SOLD_OUT
	BundleSalesState *string `json:"BundleSalesState,omitempty" name:"BundleSalesState"`

	// Package type.
	// Valid values:
	// <li> GENERAL_BUNDLE: general</li><li> STORAGE_BUNDLE: Storage</li>
	BundleType *string `json:"BundleType,omitempty" name:"BundleType"`

	// Package tag.
	// Valid values:
	// "ACTIVITY": promotional package
	// "NORMAL": regular package
	// "CAREFREE": carefree package
	BundleDisplayLabel *string `json:"BundleDisplayLabel,omitempty" name:"BundleDisplayLabel"`
}

type CcnAttachedInstance struct {

	// CCN instance ID.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// CIDR block of associated instance.
	CidrBlock []*string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Associated instance status:
	// 
	// •  PENDING: applying
	// •  ACTIVE: connected
	// •  EXPIRED: expired
	// •  REJECTED: rejected
	// •  DELETED: deleted
	// •  FAILED: failed (it will be asynchronously unassociated after 2 hours)
	// •  ATTACHING: associating
	// •  DETACHING: unassociating
	// •  DETACHFAILED: failed to unassociate (it will be asynchronously unassociated after 2 hours)
	State *string `json:"State,omitempty" name:"State"`

	// Association time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`

	// Remarks
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateBlueprintRequest struct {
	*tchttp.BaseRequest

	// Image name, which can contain up to 60 characters.
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// Image description, which can contain up to 60 characters.
	Description *string `json:"Description,omitempty" name:"Description"`

	// ID of the instance for which to make an image.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateBlueprintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlueprintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintName")
	delete(f, "Description")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Custom image ID.
		BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBlueprintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlueprintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *CreateFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRules")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSnapshotRequest struct {
	*tchttp.BaseRequest

	// ID of the instance for which to create a snapshot.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Snapshot name, which can contain up to 60 characters.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
}

func (r *CreateInstanceSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Snapshot ID.
		SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest

	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Key pair information.
		KeyPair *KeyPair `json:"KeyPair,omitempty" name:"KeyPair"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlueprintsRequest struct {
	*tchttp.BaseRequest

	// Image ID list, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds"`
}

func (r *DeleteBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlueprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *DeleteFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRules")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Key pair ID list. Each request can contain up to 10 key pairs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest

	// List of IDs of snapshots to be deleted, which can be queried through `DescribeSnapshots`.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotsRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeniedAction struct {

	// Restricted operation name.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Restricted operation message code.
	Code *string `json:"Code,omitempty" name:"Code"`

	// Restricted operation message.
	Message *string `json:"Message,omitempty" name:"Message"`
}

type DescribeBlueprintInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list, which currently can contain only one instance.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeBlueprintInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlueprintInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlueprintInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible image instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Image instance list information.
		BlueprintInstanceSet []*BlueprintInstance `json:"BlueprintInstanceSet,omitempty" name:"BlueprintInstanceSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlueprintInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlueprintsRequest struct {
	*tchttp.BaseRequest

	// Image ID list.
	BlueprintIds []*string `json:"BlueprintIds,omitempty" name:"BlueprintIds"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter list
	// <li>blueprint-id</li>Filter by **image ID**.
	// Type: String
	// Required: no
	// <li>blueprint-type</li>Filter by **image type**.
	// Valid values: `APP_OS`: application image; `PURE_OS`: system image; `PRIVATE`: custom image; `SHARED`: shared image
	// Type: String
	// Required: no
	// <li>platform-type</li>Filter by **image platform type**.
	// Valid values: `LINUX_UNIX`: Linux or Unix; `WINDOWS`: Windows
	// Type: String
	// Required: no
	// <li>blueprint-name</li>Filter by **image name**.
	// Type: String
	// Required: no
	// <li>blueprint-state</li>Filter by **image status**.
	// Type: String
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. `BlueprintIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlueprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible images.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Image details list.
		BlueprintSet []*Blueprint `json:"BlueprintSet,omitempty" name:"BlueprintSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBundlesRequest struct {
	*tchttp.BaseRequest

	// Package ID list.
	BundleIds []*string `json:"BundleIds,omitempty" name:"BundleIds"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter list.
	// <li>bundle-id</li>Filter by **package ID**.
	// Type: String
	// Required: no
	// <li>support-platform-type</li>Filter by **system type**.
	// Valid values: LINUX_UNIX (Linux/Unix), WINDOWS ( Windows)
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. You cannot specify both `BundleIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// AZ list, which contains all AZs by default.
	Zones []*string `json:"Zones,omitempty" name:"Zones"`
}

func (r *DescribeBundlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundlesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Zones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBundlesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBundlesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of package details.
		BundleSet []*Bundle `json:"BundleSet,omitempty" name:"BundleSet"`

		// Total number of eligible packages, which is used for pagination.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBundlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundlesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCcnAttachedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAttachedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnAttachedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instances associated with the CCN instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
		CcnAttachedInstanceSet []*CcnAttachedInstance `json:"CcnAttachedInstanceSet,omitempty" name:"CcnAttachedInstanceSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnAttachedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible firewall rules.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Firewall rule details list.
		FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitempty" name:"FirewallRuleSet"`

		// Firewall version number.
		FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirewallRulesTemplateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFirewallRulesTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirewallRulesTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFirewallRulesTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible firewall rules.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Firewall rule details list.
		FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitempty" name:"FirewallRuleSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFirewallRulesTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirewallRulesTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralResourceQuotasRequest struct {
	*tchttp.BaseRequest

	// List of resource names. Valid values: USER_KEY_PAIR, INSTANCE, SNAPSHOT.
	ResourceNames []*string `json:"ResourceNames,omitempty" name:"ResourceNames"`
}

func (r *DescribeGeneralResourceQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralResourceQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralResourceQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralResourceQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of general resource quota details.
		GeneralResourceQuotaSet []*GeneralResourceQuota `json:"GeneralResourceQuotaSet,omitempty" name:"GeneralResourceQuotaSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralResourceQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralResourceQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLoginKeyPairAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceLoginKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLoginKeyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLoginKeyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLoginKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Whether to allow login with the default key pair. Valid values: YES, NO.
		PermitLogin *string `json:"PermitLogin,omitempty" name:"PermitLogin"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceLoginKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLoginKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVncUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceVncUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Instance VNC URL.
		InstanceVncUrl *string `json:"InstanceVncUrl,omitempty" name:"InstanceVncUrl"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceVncUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance operation limit details.
		InstanceDeniedActionSet []*InstanceDeniedActions `json:"InstanceDeniedActionSet,omitempty" name:"InstanceDeniedActionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Filter list.
	// <li>instance-name</li>Filter by **instance name**.
	// Type: String
	// Required: no
	// <li>private-ip-address</li>Filter by **private IP of instance primary ENI**.
	// Type: String
	// Required: no
	// <li>public-ip-address</li>Filter by **public IP of instance primary ENI**.
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. You cannot specify both `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance details.
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesReturnableRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesReturnableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesReturnableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesReturnableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesReturnableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of returnable instance details.
		InstanceReturnableSet []*InstanceReturnable `json:"InstanceReturnableSet,omitempty" name:"InstanceReturnableSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesReturnableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesReturnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesTrafficPackagesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible instance traffic package details.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of instance traffic package details.
		InstanceTrafficPackageSet []*InstanceTrafficPackage `json:"InstanceTrafficPackageSet,omitempty" name:"InstanceTrafficPackageSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Key pair ID list.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter list
	// <li>key-id</li>Filter by **key pair ID**.
	// Type: String
	// Required: no
	// <li>key-name</li>Filter by **key pair name**.
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. `KeyIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible key pairs.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of key pair details.
		KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModifyInstanceBundlesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Filter list.
	// <li>bundle-id</li>Filter by **package ID**.
	// Type: String
	// Required: no
	// <li>support-platform-type</li>Filter by **system type**.
	// Valid values: LINUX_UNIX (Linux/Unix), WINDOWS ( Windows)
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeModifyInstanceBundlesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyInstanceBundlesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModifyInstanceBundlesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModifyInstanceBundlesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of matched instances.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// New package details
		ModifyBundleSet []*ModifyBundle `json:"ModifyBundleSet,omitempty" name:"ModifyBundleSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModifyInstanceBundlesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyInstanceBundlesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of regions.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Region information list.
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResetInstanceBlueprintsRequest struct {
	*tchttp.BaseRequest

	// Instance ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter list
	// <li>blueprint-id</li>Filter by **image ID**.
	// Type: String
	// Required: no
	// <li>blueprint-type</li>Filter by **image type**.
	// Valid values: `APP_OS`: application image; `PURE_OS`: system image; `PRIVATE`: custom image
	// Type: String
	// Required: no
	// <li>platform-type</li>Filter by **image platform type**.
	// Valid values: `LINUX_UNIX`: Linux or Unix; `WINDOWS`: Windows
	// Type: String
	// Required: no
	// <li>blueprint-name</li>Filter by **image name**.
	// Type: String
	// Required: no
	// <li>blueprint-state</li>Filter by **image status**.
	// Type: String
	// Required: no
	// 
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. `BlueprintIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeResetInstanceBlueprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResetInstanceBlueprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResetInstanceBlueprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResetInstanceBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible images.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Image reset information list
		ResetInstanceBlueprintSet []*ResetInstanceBlueprint `json:"ResetInstanceBlueprintSet,omitempty" name:"ResetInstanceBlueprintSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResetInstanceBlueprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResetInstanceBlueprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// Snapshot ID list, which can be queried through `DescribeSnapshots`.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

func (r *DescribeSnapshotsDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of snapshot operation limit details.
		SnapshotDeniedActionSet []*SnapshotDeniedActions `json:"SnapshotDeniedActionSet,omitempty" name:"SnapshotDeniedActionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSnapshotsDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest

	// List of IDs of snapshots to be queried.
	// You cannot specify `SnapshotIds` and `Filters` at the same time.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Filter list.
	// <li>snapshot-id</li>Filter by **snapshot ID**.
	// Type: String
	// Required: no
	// <li>disk-id</li>Filter by **disk ID**.
	// Type: String
	// Required: no
	// <li>snapshot-name</li>Filter by **snapshot name**.
	// Type: String
	// Required: no
	// <li>instance-id</li>Filter by **instance ID**.
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. You cannot specify both `SnapshotIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of snapshots.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of AZs
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of AZ details
		ZoneInfoSet []*ZoneInfo `json:"ZoneInfoSet,omitempty" name:"ZoneInfoSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachCcnRequest struct {
	*tchttp.BaseRequest

	// CCN instance ID.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *DetachCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Key pair ID list. Each request can contain up to 100 key pairs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Field to be filtered.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FirewallRule struct {

	// Protocol. Valid values: TCP, UDP, ICMP, ALL.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port. Valid values: ALL, one single port, multiple ports separated by commas, or port range indicated by a minus sign
	Port *string `json:"Port,omitempty" name:"Port"`

	// IP range or IP (mutually exclusive). Default value: 0.0.0.0/0, which indicates all sources.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Valid values: ACCEPT, DROP. Default value: ACCEPT.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Firewall rule description.
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitempty" name:"FirewallRuleDescription"`
}

type FirewallRuleInfo struct {

	// Application type. Valid values: custom, HTTP (80), HTTPS (443), Linux login (22), Windows login (3389), MySQL (3306), SQL Server (1433), all TCP ports, all UDP ports, Ping-ICMP, ALL.
	AppType *string `json:"AppType,omitempty" name:"AppType"`

	// Protocol. Valid values: TCP, UDP, ICMP, ALL.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Port. Valid values: ALL, one single port, multiple ports separated by commas, or port range indicated by a minus sign
	Port *string `json:"Port,omitempty" name:"Port"`

	// IP range or IP (mutually exclusive). Default value: 0.0.0.0/0, which indicates all sources.
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// Valid values: ACCEPT, DROP. Default value: ACCEPT.
	Action *string `json:"Action,omitempty" name:"Action"`

	// Firewall rule description.
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitempty" name:"FirewallRuleDescription"`
}

type GeneralResourceQuota struct {

	// Resource name.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Number of available resources.
	ResourceQuotaAvailable *int64 `json:"ResourceQuotaAvailable,omitempty" name:"ResourceQuotaAvailable"`

	// Total number of resources.
	ResourceQuotaTotal *int64 `json:"ResourceQuotaTotal,omitempty" name:"ResourceQuotaTotal"`
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest

	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// Public key content of the key pair, which is in the OpenSSH RSA format.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	delete(f, "PublicKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Key pair ID.
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateBlueprintRequest struct {
	*tchttp.BaseRequest

	// Number of custom images. Default value: 1.
	BlueprintCount *int64 `json:"BlueprintCount,omitempty" name:"BlueprintCount"`
}

func (r *InquirePriceCreateBlueprintRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateBlueprintRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Custom image price.
		BlueprintPrice *BlueprintPrice `json:"BlueprintPrice,omitempty" name:"BlueprintPrice"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateBlueprintResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateBlueprintResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance package ID.
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// Number of instances to be created. Default value: 1.
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Prepaid mode, i.e., monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. It is required for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// Application image ID, which is required if a paid application image is used and can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`
}

func (r *InquirePriceCreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleId")
	delete(f, "InstanceCount")
	delete(f, "InstanceChargePrepaid")
	delete(f, "BlueprintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price query information.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance to be renewed.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Prepaid mode, i.e., monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. It is required for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price query information.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Package ID.
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// Image ID.
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// Number of instance CPU cores.
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Instance memory capacity in GB.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance billing mode. Valid values: 
	// PREPAID: prepaid (i.e., monthly subscription).
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Instance system disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Private IP of instance primary ENI. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateAddresses []*string `json:"PrivateAddresses,omitempty" name:"PrivateAddresses"`

	// Public IP of instance primary ENI. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicAddresses []*string `json:"PublicAddresses,omitempty" name:"PublicAddresses"`

	// Instance bandwidth information.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Auto-Renewal flag. Valid values: 
	// NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically  
	// NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Instance login settings.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Instance status. Valid values: 
	// <li>PENDING: creating</li><li>LAUNCH_FAILED: creation failed</li><li>RUNNING: running</li><li>STOPPED: shut down</li><li>STARTING: starting</li><li>STOPPING: shutting down</li><li>REBOOTING: rebooting</li><li>SHUTDOWN: shut down and to be terminated</li><li>TERMINATING: terminating</li>
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// Globally unique ID of instance.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Last instance operation, such as `StopInstances` and `ResetInstance`. Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// Last instance operation status. Valid values: 
	// SUCCESS: operation succeeded 
	// OPERATING: the operation is being executed 
	// FAILED: operation failed 
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// Unique request ID for the last operation of the instance. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// Isolation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// Creation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Expiration time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// OS type, such as LINUX_UNIX and WINDOWS.
	PlatformType *string `json:"PlatformType,omitempty" name:"PlatformType"`

	// OS type.
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// OS name.
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// AZ.
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type InstanceChargePrepaid struct {

	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Auto-Renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew <br><li>DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically<br><br>Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed monthly if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceDeniedActions struct {

	// Instance ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of operation limits.
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type InstancePrice struct {

	// Original package unit price.
	OriginalBundlePrice *float64 `json:"OriginalBundlePrice,omitempty" name:"OriginalBundlePrice"`

	// Original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discount.
	Discount *int64 `json:"Discount,omitempty" name:"Discount"`

	// Discounted price.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type InstanceReturnable struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Whether the instance can be returned.
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// Error code of instance return failure.
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// Error message of instance return failure.
	ReturnFailMessage *string `json:"ReturnFailMessage,omitempty" name:"ReturnFailMessage"`
}

type InstanceTrafficPackage struct {

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// List of traffic package details.
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitempty" name:"TrafficPackageSet"`
}

type InternetAccessible struct {

	// Network billing mode. Valid values:
	// <li>Bill by traffic package: TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>Bill by bandwidth: BANDWIDTH_POSTPAID_BY_HOUR</li>
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// Public network outbound bandwidth cap in Mbps.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type KeyPair struct {

	// Key pair ID, which is the unique identifier of a key pair.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Key pair name.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// Public key (in plain text) of key pair.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// List of IDs of instances associated with the key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Private key of key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

type LoginSettings struct {

	// Key ID list. After a key is associated, you can use it to access the instance. Note: this field may return [], indicating that no valid values can be obtained.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type ModifyBlueprintAttributeRequest struct {
	*tchttp.BaseRequest

	// Image ID, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`

	// New image name, which can contain up to 60 characters.
	BlueprintName *string `json:"BlueprintName,omitempty" name:"BlueprintName"`

	// New image description, which can contain up to 60 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyBlueprintAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlueprintAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BlueprintId")
	delete(f, "BlueprintName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlueprintAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBlueprintAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBlueprintAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlueprintAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBundle struct {

	// Price difference that you need to pay for the new instance package after change.
	ModifyPrice *Price `json:"ModifyPrice,omitempty" name:"ModifyPrice"`

	// Package change status. Valid values:
	// <li>SOLD_OUT: packages are sold out</li>
	// <li>AVAILABLE: packages can be changed</li>
	// <li>UNAVAILABLE: packages cannot be changed currently</li>
	ModifyBundleState *string `json:"ModifyBundleState,omitempty" name:"ModifyBundleState"`

	// Package information.
	Bundle *Bundle `json:"Bundle,omitempty" name:"Bundle"`
}

type ModifyFirewallRuleDescriptionRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Firewall rule.
	FirewallRule *FirewallRule `json:"FirewallRule,omitempty" name:"FirewallRule"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *ModifyFirewallRuleDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRuleDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRule")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFirewallRuleDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFirewallRuleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFirewallRuleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRuleDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFirewallRulesRequest struct {
	*tchttp.BaseRequest

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitempty" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitempty" name:"FirewallVersion"`
}

func (r *ModifyFirewallRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FirewallRules")
	delete(f, "FirewallVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFirewallRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFirewallRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFirewallRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance name, which is customizable and can contain up to 60 characters.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesLoginKeyPairAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Whether to allow login with the default key pair. Valid values: YES: yes; NO: no.
	PermitLogin *string `json:"PermitLogin,omitempty" name:"PermitLogin"`
}

func (r *ModifyInstancesLoginKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesLoginKeyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "PermitLogin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesLoginKeyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesLoginKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesLoginKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesLoginKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Auto-Renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed monthly if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest

	// Snapshot ID, which can be queried through `DescribeSnapshots`.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// New snapshot name, which can contain up to 60 characters.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotAttributeRequest has unknown keys!", "")
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {

	// Instance price.
	InstancePrice *InstancePrice `json:"InstancePrice,omitempty" name:"InstancePrice"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region name, such as `ap-guangzhou`.
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region description, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Region availability status. Its value can only be `AVAILABLE`.
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`

	// Whether the region is in the Chinese mainland
	IsChinaMainland *bool `json:"IsChinaMainland,omitempty" name:"IsChinaMainland"`
}

type ResetAttachCcnRequest struct {
	*tchttp.BaseRequest

	// CCN instance ID.
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *ResetAttachCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAttachCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAttachCcnResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAttachCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceBlueprint struct {

	// Image details
	BlueprintInfo *Blueprint `json:"BlueprintInfo,omitempty" name:"BlueprintInfo"`

	// Whether the image can be reset as the target image
	IsResettable *bool `json:"IsResettable,omitempty" name:"IsResettable"`

	// Non-Resettable flag. If the image is resettable, it will be ""
	NonResettableMessage *string `json:"NonResettableMessage,omitempty" name:"NonResettableMessage"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image ID, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitempty" name:"BlueprintId"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BlueprintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Login password of the instance(s). The password requirements vary among different operating systems:
	// The password of a `LINUX_UNIX` instance must contain 8–30 characters (above 12 characters preferably) in at least three of the following types and cannot begin with "/": <br><li>Lowercase letters: [a–z]<br><li>Uppercase letters: [A–Z]<br><li>Digits: 0–9<br><li>Special symbols: ()\`~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	// The password of a `WINDOWS` instance must contain 12–30 characters in at least three of the following types and cannot begin with "/" or include the username: <br><li>Lowercase letters: [a–z]<br><li>Uppercase letters: [A–Z]<br><li>Digits: 0–9<br><li>Special symbols: ()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>If both `LINUX_UNIX` and `WINDOWS` instances exist, the requirements for password complexity of `WINDOWS` instances shall prevail.
	Password *string `json:"Password,omitempty" name:"Password"`

	// OS username of the instance for which you want to reset the password, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Password")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Snapshot struct {

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Type of the disk for which the snapshot is created. Valid values: <li>SYSTEM_DISK: system disk</li>
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// ID of the disk for which the snapshot is created.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Size of the disk in GB for which the snapshot is created.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Snapshot name, which is a custom snapshot alias.
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// Snapshot status. Valid values:
	// <li>NORMAL: normal </li>
	// <li>CREATING: creating</li>
	// <li>ROLLBACKING: rolling back</li>
	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`

	// Snapshot creation or rollback progress in percentage. After success, the value of this field will become 100.
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// Last snapshot operation. It is recorded only during snapshot creation and rollback.
	// Example values: CreateInstanceSnapshot, RollbackInstanceSnapshot.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// Last snapshot operation status. It is recorded only during snapshot creation and rollback.
	// Valid values:
	// <li>SUCCESS: operation succeeded</li>
	// <li>OPERATING: the operation is being executed</li>
	// <li>FAILED: operation failed</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// Unique request ID for the last snapshot operation. It is recorded only during snapshot creation and rollback.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// Snapshot creation time.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type SnapshotDeniedActions struct {

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// List of operation limits.
	DeniedActions []*DeniedAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type Software struct {

	// Software name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Software version.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Software picture URL.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Software installation directory.
	InstallDir *string `json:"InstallDir,omitempty" name:"InstallDir"`

	// List of software details.
	DetailSet []*SoftwareDetail `json:"DetailSet,omitempty" name:"DetailSet"`
}

type SoftwareDetail struct {

	// Unique detail key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Detail title.
	Title *string `json:"Title,omitempty" name:"Title"`

	// Detail value.
	Value *string `json:"Value,omitempty" name:"Value"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// System disk type.
	// Valid values: 
	// <li> LOCAL_BASIC: local disk</li><li> LOCAL_SSD: local SSD disk</li><li> CLOUD_BASIC: HDD cloud disk</li><li> CLOUD_SSD: SSD cloud disk</li><li> CLOUD_PREMIUM: Premium Cloud Storage</li>
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// System disk ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID list, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrafficPackage struct {

	// Traffic package ID.
	TrafficPackageId *string `json:"TrafficPackageId,omitempty" name:"TrafficPackageId"`

	// Used traffic in bytes during traffic package validity period.
	TrafficUsed *int64 `json:"TrafficUsed,omitempty" name:"TrafficUsed"`

	// Total traffic in bytes during traffic package validity period.
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitempty" name:"TrafficPackageTotal"`

	// Remaining traffic in bytes during traffic package validity period.
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitempty" name:"TrafficPackageRemaining"`

	// Traffic exceeding package amount in bytes during traffic package validity period.
	TrafficOverflow *int64 `json:"TrafficOverflow,omitempty" name:"TrafficOverflow"`

	// Start time of traffic package validity period according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time of traffic package validity period according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Traffic package expiration time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// Traffic package status:
	// <li>NETWORK_NORMAL: normal</li>
	// <li>OVERDUE_NETWORK_DISABLED: the network is disconnected due to overdue payments</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ZoneInfo struct {

	// AZ
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Chinese name of AZ
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// AZ tags on instance purchase page
	InstanceDisplayLabel *string `json:"InstanceDisplayLabel,omitempty" name:"InstanceDisplayLabel"`
}
