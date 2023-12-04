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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyInstanceSnapshotRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`
}

type ApplyInstanceSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`
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

// Predefined struct for user
type ApplyInstanceSnapshotResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ApplyInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *ApplyInstanceSnapshotResponseParams `json:"Response"`
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

// Predefined struct for user
type AssociateInstancesKeyPairsRequestParams struct {
	// Key pair ID list. Each request can contain up to 100 key pairs.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Key pair ID list. Each request can contain up to 100 key pairs.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type AssociateInstancesKeyPairsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateInstancesKeyPairsResponseParams `json:"Response"`
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

// Predefined struct for user
type AttachCcnRequestParams struct {
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

type AttachCcnRequest struct {
	*tchttp.BaseRequest
	
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
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

// Predefined struct for user
type AttachCcnResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachCcnResponse struct {
	*tchttp.BaseResponse
	Response *AttachCcnResponseParams `json:"Response"`
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

type AttachDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of elastic cloud disks attached to the instance
	AttachedDiskCount *int64 `json:"AttachedDiskCount,omitnil" name:"AttachedDiskCount"`

	// Upper limit of attached elastic cloud disks
	MaxAttachCount *int64 `json:"MaxAttachCount,omitnil" name:"MaxAttachCount"`
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Specify whether to renew an instance automatically when it expires. Values: 
	// 
	// `NOTIFY_AND_AUTO_RENEW`: Trigger expiration notification and renew automatically; `NOTIFY_AND_MANUAL_RENEW`: Trigger expiration notification but do not renew; `DISABLE_NOTIFY_AND_MANUAL_RENEW`: Do not trigger the notification and do not renew.
	// 
	// Default: `NOTIFY_AND_MANUAL_RENEW`. If `NOTIFY_AND_AUTO_RENEW` is specified, the instance is automatically renewed on a monthly basis when the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Specify whether to renew an instance automatically when it expires. Values: 
	// 
	// `NOTIFY_AND_AUTO_RENEW`: Trigger expiration notification and renew automatically; `NOTIFY_AND_MANUAL_RENEW`: Trigger expiration notification but do not renew; `DISABLE_NOTIFY_AND_MANUAL_RENEW`: Do not trigger the notification and do not renew.
	// 
	// Default: `NOTIFY_AND_MANUAL_RENEW`. If `NOTIFY_AND_AUTO_RENEW` is specified, the instance is automatically renewed on a monthly basis when the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
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
	delete(f, "DiskIds")
	delete(f, "InstanceId")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// ID of the instance to be mounted to. The instance must be **Running**.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The mount point within the instance. Only Linux instances are supported. If it's not specified, the default mount point is "/data/disk".
	MountPoint *string `json:"MountPoint,omitnil" name:"MountPoint"`

	// The file system type. Values: `ext4` (default) and `xfs`. Only Linux instances are supported. 
	FileSystemType *string `json:"FileSystemType,omitnil" name:"FileSystemType"`
}

type Blueprint struct {
	// Image ID, which is the unique identifier of `Blueprint`.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// Image title to be displayed.
	DisplayTitle *string `json:"DisplayTitle,omitnil" name:"DisplayTitle"`

	// Image version to be displayed.
	DisplayVersion *string `json:"DisplayVersion,omitnil" name:"DisplayVersion"`

	// Image description information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil" name:"Description"`

	// OS name.
	OsName *string `json:"OsName,omitnil" name:"OsName"`

	// OS type.
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// OS type, such as LINUX_UNIX and WINDOWS.
	PlatformType *string `json:"PlatformType,omitnil" name:"PlatformType"`

	// Image type, such as APP_OS, PURE_OS, and PRIVATE.
	BlueprintType *string `json:"BlueprintType,omitnil" name:"BlueprintType"`

	// Image picture URL.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// System disk size required by image in GB.
	RequiredSystemDiskSize *int64 `json:"RequiredSystemDiskSize,omitnil" name:"RequiredSystemDiskSize"`

	// Image status.
	BlueprintState *string `json:"BlueprintState,omitnil" name:"BlueprintState"`

	// Creation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Image name.
	BlueprintName *string `json:"BlueprintName,omitnil" name:"BlueprintName"`

	// Whether the image supports automation tools.
	SupportAutomationTools *bool `json:"SupportAutomationTools,omitnil" name:"SupportAutomationTools"`

	// Memory size required by image in GB.
	RequiredMemorySize *int64 `json:"RequiredMemorySize,omitnil" name:"RequiredMemorySize"`

	// ID of the Lighthouse image shared from a CVM image
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// URL of official website of the open-source project
	CommunityUrl *string `json:"CommunityUrl,omitnil" name:"CommunityUrl"`

	// Guide documentation URL
	GuideUrl *string `json:"GuideUrl,omitnil" name:"GuideUrl"`

	// Array of IDs of scenes associated with an image
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SceneIdSet []*string `json:"SceneIdSet,omitnil" name:"SceneIdSet"`

	// Docker version.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DockerVersion *string `json:"DockerVersion,omitnil" name:"DockerVersion"`
}

type BlueprintInstance struct {
	// Image information.
	Blueprint *Blueprint `json:"Blueprint,omitnil" name:"Blueprint"`

	// Software list.
	SoftwareSet []*Software `json:"SoftwareSet,omitnil" name:"SoftwareSet"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type BlueprintPrice struct {
	// Original image unit price in USD.
	OriginalBlueprintPrice *float64 `json:"OriginalBlueprintPrice,omitnil" name:"OriginalBlueprintPrice"`

	// Original image total price in USD.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// Discount.
	Discount *int64 `json:"Discount,omitnil" name:"Discount"`

	// Discounted image total price in USD.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`
}

type Bundle struct {
	// Package ID.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Memory size in GB.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// System disk type.
	// Values: 
	// <li>`CLOUD_SSD`: SSD cloud disks</li><li>`CLOUD_PREMIUM`: Premium cloud disks</li>
	SystemDiskType *string `json:"SystemDiskType,omitnil" name:"SystemDiskType"`

	// System disk size in GB.
	SystemDiskSize *int64 `json:"SystemDiskSize,omitnil" name:"SystemDiskSize"`

	// Monthly network traffic in GB.
	MonthlyTraffic *int64 `json:"MonthlyTraffic,omitnil" name:"MonthlyTraffic"`

	// Whether Linux/Unix is supported.
	SupportLinuxUnixPlatform *bool `json:"SupportLinuxUnixPlatform,omitnil" name:"SupportLinuxUnixPlatform"`

	// Whether Windows is supported.
	SupportWindowsPlatform *bool `json:"SupportWindowsPlatform,omitnil" name:"SupportWindowsPlatform"`

	// Current package unit price information.
	Price *Price `json:"Price,omitnil" name:"Price"`

	// Number of CPU cores.
	CPU *int64 `json:"CPU,omitnil" name:"CPU"`

	// Peak bandwidth in Mbps.
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// Network billing mode.
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// Package sale status. Valid values: AVAILABLE, SOLD_OUT
	BundleSalesState *string `json:"BundleSalesState,omitnil" name:"BundleSalesState"`

	// Bundle type. 
	// Valid values: 
	// <li>STARTER_BUNDLE: Starter bundle</li>
	// <li>GENERAL_BUNDLE: General bundle</li>
	// <li>ENTERPRISE_BUNDLE: Enterprise bundle</li>
	// <li>STORAGE_BUNDLE: Storage-optimized bundle</li>
	// <li>EXCLUSIVE_BUNDLE: Dedicated bundle</li>
	// <li>HK_EXCLUSIVE_BUNDLE: Hong Kong-dedicated bundle </li>
	// <li>CAREFREE_BUNDLE: Lighthouse Care bundle</li>
	// <li>BEFAST_BUNDLE: BeFast bundle </li>
	BundleType *string `json:"BundleType,omitnil" name:"BundleType"`

	// Bundle type description 
	// Note: This parameter may return null, indicating that no valid values can be obtained.
	BundleTypeDescription *string `json:"BundleTypeDescription,omitnil" name:"BundleTypeDescription"`

	// Package tag.
	// Valid values:
	// "ACTIVITY": promotional package
	// "NORMAL": regular package
	// "CAREFREE": carefree package
	BundleDisplayLabel *string `json:"BundleDisplayLabel,omitnil" name:"BundleDisplayLabel"`
}

type CcnAttachedInstance struct {
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CIDR block of associated instance.
	CidrBlock []*string `json:"CidrBlock,omitnil" name:"CidrBlock"`

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
	State *string `json:"State,omitnil" name:"State"`

	// Association time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AttachedTime *string `json:"AttachedTime,omitnil" name:"AttachedTime"`

	// Remarks
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ContainerEnv struct {
	// Environment variable key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Environment variable value
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type CreateBlueprintRequestParams struct {
	// Image name, which can contain up to 60 characters.
	BlueprintName *string `json:"BlueprintName,omitnil" name:"BlueprintName"`

	// Image description, which can contain up to 60 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// ID of the instance for which to make an image.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to forcibly shut down the instance before creating the image 
	// Valid values: 
	// `True`: Shut down and instance first 
	// `False`: Create the image when the instance is running 
	// Default: `True` 
	// Note that if you create an image when the instance is running, there might be data loss.
	ForcePowerOff *bool `json:"ForcePowerOff,omitnil" name:"ForcePowerOff"`
}

type CreateBlueprintRequest struct {
	*tchttp.BaseRequest
	
	// Image name, which can contain up to 60 characters.
	BlueprintName *string `json:"BlueprintName,omitnil" name:"BlueprintName"`

	// Image description, which can contain up to 60 characters.
	Description *string `json:"Description,omitnil" name:"Description"`

	// ID of the instance for which to make an image.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to forcibly shut down the instance before creating the image 
	// Valid values: 
	// `True`: Shut down and instance first 
	// `False`: Create the image when the instance is running 
	// Default: `True` 
	// Note that if you create an image when the instance is running, there might be data loss.
	ForcePowerOff *bool `json:"ForcePowerOff,omitnil" name:"ForcePowerOff"`
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
	delete(f, "ForcePowerOff")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlueprintRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlueprintResponseParams struct {
	// Custom image ID.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlueprintResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDisksRequestParams struct {
	// Availability zone. You can call [DescribeZones](https://intl.cloud.tencent.com/document/product/1207/57513?from_cn_redirect=1) and get the information in the `Zone` parameter re 
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Cloud disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Cloud disk media type. Valid values: "CLOUD_PREMIUM" (premium cloud disk), "CLOUD_SSD" (SSD cloud disk).
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Parameters of monthly subscribed cloud disks
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil" name:"DiskChargePrepaid"`

	// Image name, which can contain up to 60 characters.
	DiskName *string `json:"DiskName,omitnil" name:"DiskName"`

	// Number of cloud disks. Range: [1, 30]. Default value: 1.
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// Specify the quota of disk backups. No quota if it’s left empty. Only one quota is allowed.
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`

	// Whether to use the vouchers automatically. It defaults to No.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// Automatically mount and initialize the data disk.
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitnil" name:"AutoMountConfiguration"`
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone. You can call [DescribeZones](https://intl.cloud.tencent.com/document/product/1207/57513?from_cn_redirect=1) and get the information in the `Zone` parameter re 
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Cloud disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Cloud disk media type. Valid values: "CLOUD_PREMIUM" (premium cloud disk), "CLOUD_SSD" (SSD cloud disk).
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Parameters of monthly subscribed cloud disks
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil" name:"DiskChargePrepaid"`

	// Image name, which can contain up to 60 characters.
	DiskName *string `json:"DiskName,omitnil" name:"DiskName"`

	// Number of cloud disks. Range: [1, 30]. Default value: 1.
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// Specify the quota of disk backups. No quota if it’s left empty. Only one quota is allowed.
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`

	// Whether to use the vouchers automatically. It defaults to No.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// Automatically mount and initialize the data disk.
	AutoMountConfiguration *AutoMountConfiguration `json:"AutoMountConfiguration,omitnil" name:"AutoMountConfiguration"`
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
	delete(f, "Zone")
	delete(f, "DiskSize")
	delete(f, "DiskType")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskName")
	delete(f, "DiskCount")
	delete(f, "DiskBackupQuota")
	delete(f, "AutoVoucher")
	delete(f, "AutoMountConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisksResponseParams struct {
	// List of IDs created by using this API. The returning of IDs does not mean that the instances are created successfully.
	// 
	// You can call `DescribeDisks` API, and find the disk ID in the `DiskSet` returned to check its status. If the status changes from `PENDING` to `UNATTACHED` or `ATTACHED`, the instance is created successfully.
	DiskIdSet []*string `json:"DiskIdSet,omitnil" name:"DiskIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateFirewallRulesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
}

type CreateFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
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

// Predefined struct for user
type CreateFirewallRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFirewallRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateInstanceSnapshotRequestParams struct {
	// ID of the instance for which to create a snapshot.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Snapshot name, which can contain up to 60 characters.
	SnapshotName *string `json:"SnapshotName,omitnil" name:"SnapshotName"`
}

type CreateInstanceSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance for which to create a snapshot.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Snapshot name, which can contain up to 60 characters.
	SnapshotName *string `json:"SnapshotName,omitnil" name:"SnapshotName"`
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

// Predefined struct for user
type CreateInstanceSnapshotResponseParams struct {
	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstanceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceSnapshotResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// Bundle ID. You can get it via the [DescribeBundles](https://intl.cloud.tencent.com/document/api/1207/47575?from_cn_redirect=1) API.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Image ID. You can get it via the [DescribeBlueprints](https://intl.cloud.tencent.com/document/api/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// Monthly subscription information for the instance, including the purchase period, setting of auto-renewal, etc.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Instance display name.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Number of the instances to purchase. For monthly subscribed instances, the value can be 1 to 30. The default value is `1`. Note that this number can not exceed the remaining quota under the current account.
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// List of availability zones. A random AZ is selected by default.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Whether the request is a dry run only.
	// `true`: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// `false` (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Login password of the instance. It’s only available for Windows instances. If it’s not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitnil" name:"LoginConfiguration"`

	// Configuration of the containers to create
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil" name:"Containers"`

	// Whether to use the vouchers automatically. It defaults to No.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Bundle ID. You can get it via the [DescribeBundles](https://intl.cloud.tencent.com/document/api/1207/47575?from_cn_redirect=1) API.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Image ID. You can get it via the [DescribeBlueprints](https://intl.cloud.tencent.com/document/api/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// Monthly subscription information for the instance, including the purchase period, setting of auto-renewal, etc.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Instance display name.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Number of the instances to purchase. For monthly subscribed instances, the value can be 1 to 30. The default value is `1`. Note that this number can not exceed the remaining quota under the current account.
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// List of availability zones. A random AZ is selected by default.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Whether the request is a dry run only.
	// `true`: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// `false` (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// Login password of the instance. It’s only available for Windows instances. If it’s not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration *LoginConfiguration `json:"LoginConfiguration,omitnil" name:"LoginConfiguration"`

	// Configuration of the containers to create
	Containers []*DockerContainerConfiguration `json:"Containers,omitnil" name:"Containers"`

	// Whether to use the vouchers automatically. It defaults to No.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleId")
	delete(f, "BlueprintId")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceName")
	delete(f, "InstanceCount")
	delete(f, "Zones")
	delete(f, "DryRun")
	delete(f, "ClientToken")
	delete(f, "LoginConfiguration")
	delete(f, "Containers")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// List of IDs created by using this API. The returning of IDs does not mean that the instances are created successfully.
	// 
	// You can call `DescribeInstances` API, and find the instance ID in the `InstancesSet` returned to check its status. If the `status` is `running`, the instance is created successfully.
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil" name:"InstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairRequestParams struct {
	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitnil" name:"KeyName"`
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitnil" name:"KeyName"`
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

// Predefined struct for user
type CreateKeyPairResponseParams struct {
	// Key pair information.
	KeyPair *KeyPair `json:"KeyPair,omitnil" name:"KeyPair"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyPairResponseParams `json:"Response"`
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

type DataDiskPrice struct {
	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitnil" name:"DiskId"`

	// Cloud disk unit price.
	OriginalDiskPrice *float64 `json:"OriginalDiskPrice,omitnil" name:"OriginalDiskPrice"`

	// Total price of cloud disk
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// Discount.
	Discount *float64 `json:"Discount,omitnil" name:"Discount"`

	// Discounted total price.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`

	// ID of the instance to which the data disk is mounted.
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

// Predefined struct for user
type DeleteBlueprintsRequestParams struct {
	// Image ID list, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintIds []*string `json:"BlueprintIds,omitnil" name:"BlueprintIds"`
}

type DeleteBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// Image ID list, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintIds []*string `json:"BlueprintIds,omitnil" name:"BlueprintIds"`
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

// Predefined struct for user
type DeleteBlueprintsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBlueprintsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteFirewallRulesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
}

type DeleteFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
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

// Predefined struct for user
type DeleteFirewallRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFirewallRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteKeyPairsRequestParams struct {
	// Key pair ID list. Each request can contain up to 10 key pairs.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Key pair ID list. Each request can contain up to 10 key pairs.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`
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

// Predefined struct for user
type DeleteKeyPairsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKeyPairsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// List of IDs of snapshots to be deleted, which can be queried through `DescribeSnapshots`.
	SnapshotIds []*string `json:"SnapshotIds,omitnil" name:"SnapshotIds"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of snapshots to be deleted, which can be queried through `DescribeSnapshots`.
	SnapshotIds []*string `json:"SnapshotIds,omitnil" name:"SnapshotIds"`
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

// Predefined struct for user
type DeleteSnapshotsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type DeniedAction struct {
	// Restricted operation name.
	Action *string `json:"Action,omitnil" name:"Action"`

	// Restricted operation message code.
	Code *string `json:"Code,omitnil" name:"Code"`

	// Restricted operation message.
	Message *string `json:"Message,omitnil" name:"Message"`
}

// Predefined struct for user
type DescribeAllScenesRequestParams struct {
	// List of scene IDs
	SceneIds []*string `json:"SceneIds,omitnil" name:"SceneIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAllScenesRequest struct {
	*tchttp.BaseRequest
	
	// List of scene IDs
	SceneIds []*string `json:"SceneIds,omitnil" name:"SceneIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAllScenesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllScenesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllScenesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllScenesResponseParams struct {
	// List of scenes
	SceneInfoSet []*SceneInfo `json:"SceneInfoSet,omitnil" name:"SceneInfoSet"`

	// Total count of scenes
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAllScenesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllScenesResponseParams `json:"Response"`
}

func (r *DescribeAllScenesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllScenesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlueprintInstancesRequestParams struct {
	// Instance ID list, which currently can contain only one instance.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DescribeBlueprintInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list, which currently can contain only one instance.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type DescribeBlueprintInstancesResponseParams struct {
	// Number of eligible image instances.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Image instance list information.
	BlueprintInstanceSet []*BlueprintInstance `json:"BlueprintInstanceSet,omitnil" name:"BlueprintInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBlueprintInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlueprintInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBlueprintsRequestParams struct {
	// Image ID list.
	BlueprintIds []*string `json:"BlueprintIds,omitnil" name:"BlueprintIds"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter list.
	// <li>blueprint-id</li>Filter by the **image ID**.
	// Type: String
	// Required: no
	// <li>blueprint-type</li>Filter by the **image type**.
	// Valid values: `APP_OS` (application image), `PURE_OS` (system image), `DOCKER` (Docker container image), `PRIVATE` (custom image), `SHARED` (shared image)
	// Type: String
	// Required: no
	// <li>platform-type</li>Filter by the **image operating system**.
	// Valid values: `LINUX_UNIX` (Linux or Unix), `WINDOWS` (Windows)
	// Type: String
	// Required: no
	// <li>blueprint-name</li>Filter by the **image name**.
	// Type: String
	// Required: no
	// <li>blueprint-state</li>Filter by the **image status**.
	// Type: String
	// Required: no
	// <li>scene-id</li>Filter by the **scene ID**.
	// Type: String
	// Required: no
	// 
	// Each request can contain up to 10 `Filters`, each of which can contain up to 100 `Filter.Values`. `BlueprintIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// Image ID list.
	BlueprintIds []*string `json:"BlueprintIds,omitnil" name:"BlueprintIds"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter list.
	// <li>blueprint-id</li>Filter by the **image ID**.
	// Type: String
	// Required: no
	// <li>blueprint-type</li>Filter by the **image type**.
	// Valid values: `APP_OS` (application image), `PURE_OS` (system image), `DOCKER` (Docker container image), `PRIVATE` (custom image), `SHARED` (shared image)
	// Type: String
	// Required: no
	// <li>platform-type</li>Filter by the **image operating system**.
	// Valid values: `LINUX_UNIX` (Linux or Unix), `WINDOWS` (Windows)
	// Type: String
	// Required: no
	// <li>blueprint-name</li>Filter by the **image name**.
	// Type: String
	// Required: no
	// <li>blueprint-state</li>Filter by the **image status**.
	// Type: String
	// Required: no
	// <li>scene-id</li>Filter by the **scene ID**.
	// Type: String
	// Required: no
	// 
	// Each request can contain up to 10 `Filters`, each of which can contain up to 100 `Filter.Values`. `BlueprintIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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

// Predefined struct for user
type DescribeBlueprintsResponseParams struct {
	// Number of eligible images.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Image details list.
	BlueprintSet []*Blueprint `json:"BlueprintSet,omitnil" name:"BlueprintSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlueprintsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBundleDiscountRequestParams struct {
	// Package ID.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`
}

type DescribeBundleDiscountRequest struct {
	*tchttp.BaseRequest
	
	// Package ID.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`
}

func (r *DescribeBundleDiscountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundleDiscountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BundleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBundleDiscountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBundleDiscountResponseParams struct {
	// Currency: CNY, USD.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Discount tier details. The information of each tier includes the duration, discounted quantity, total price, discounted price, and discount details (user discount, official website discount, or final discount).
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitnil" name:"DiscountDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBundleDiscountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBundleDiscountResponseParams `json:"Response"`
}

func (r *DescribeBundleDiscountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBundleDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBundlesRequestParams struct {
	// Package ID list.
	BundleIds []*string `json:"BundleIds,omitnil" name:"BundleIds"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter list.
	// <li>bundle-id</li>Filter by the **bundle ID**.
	// Type: String
	// Required: No
	// <li>`support-platform-type`<li>Filter by the **system type**.
	// Values: `LINUX_UNIX` (Linux/Unix), `WINDOWS` (Windows).
	// Type: String
	// Required: No
	// <li>bundle-type</li>Filter by the **bundle type**.
	// Values: `GENERAL_BUNDLE` (General bundle), `STORAGE_BUNDLE` (Storage bundle), `ENTERPRISE_BUNDLE` (Enterprise bundle), `EXCLUSIVE_BUNDLE` (Dedicated bundle), `BEFAST_BUNDLE` (BeFast bundle), `STARTER_BUNDLE` (Beginner bundle); `CAREFREE_BUNDLE` (Carefree bundle);
	// Type: String
	// Required: No
	// <li>bundle-state</li>Filter by the **bundle status**.
	// Values: `ONLINE`, `OFFLINE`
	// Type: String
	// Required: No
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. You cannot specify both `BundleIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// AZ list, which contains all AZs by default.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`
}

type DescribeBundlesRequest struct {
	*tchttp.BaseRequest
	
	// Package ID list.
	BundleIds []*string `json:"BundleIds,omitnil" name:"BundleIds"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter list.
	// <li>bundle-id</li>Filter by the **bundle ID**.
	// Type: String
	// Required: No
	// <li>`support-platform-type`<li>Filter by the **system type**.
	// Values: `LINUX_UNIX` (Linux/Unix), `WINDOWS` (Windows).
	// Type: String
	// Required: No
	// <li>bundle-type</li>Filter by the **bundle type**.
	// Values: `GENERAL_BUNDLE` (General bundle), `STORAGE_BUNDLE` (Storage bundle), `ENTERPRISE_BUNDLE` (Enterprise bundle), `EXCLUSIVE_BUNDLE` (Dedicated bundle), `BEFAST_BUNDLE` (BeFast bundle), `STARTER_BUNDLE` (Beginner bundle); `CAREFREE_BUNDLE` (Carefree bundle);
	// Type: String
	// Required: No
	// <li>bundle-state</li>Filter by the **bundle status**.
	// Values: `ONLINE`, `OFFLINE`
	// Type: String
	// Required: No
	// Each request can contain up to 10 `Filters` and 5 `Filter.Values`. You cannot specify both `BundleIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// AZ list, which contains all AZs by default.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`
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

// Predefined struct for user
type DescribeBundlesResponseParams struct {
	// List of package details.
	BundleSet []*Bundle `json:"BundleSet,omitnil" name:"BundleSet"`

	// Total number of eligible packages, which is used for pagination.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBundlesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBundlesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCcnAttachedInstancesRequestParams struct {

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

// Predefined struct for user
type DescribeCcnAttachedInstancesResponseParams struct {
	// List of instances associated with the CCN instance.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CcnAttachedInstanceSet []*CcnAttachedInstance `json:"CcnAttachedInstanceSet,omitnil" name:"CcnAttachedInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnAttachedInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDiskConfigsRequestParams struct {
	// Filter list.
	// <li>zone</li>Filter by availability zone.
	// Type: String
	// Required: no
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeDiskConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Filter list.
	// <li>zone</li>Filter by availability zone.
	// Type: String
	// Required: no
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeDiskConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskConfigsResponseParams struct {
	// List of cloud disk configurations.
	DiskConfigSet []*DiskConfig `json:"DiskConfigSet,omitnil" name:"DiskConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDiskConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskConfigsResponseParams `json:"Response"`
}

func (r *DescribeDiskConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskDiscountRequestParams struct {
	// Cloud disk type. Valid values: "CLOUD_PREMIUM".
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Cloud disk size.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Specify the quota of disk backups. No quota if it’s left empty. Only one quota is allowed.
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`
}

type DescribeDiskDiscountRequest struct {
	*tchttp.BaseRequest
	
	// Cloud disk type. Valid values: "CLOUD_PREMIUM".
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Cloud disk size.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Specify the quota of disk backups. No quota if it’s left empty. Only one quota is allowed.
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`
}

func (r *DescribeDiskDiscountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskDiscountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiskDiscountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiskDiscountResponseParams struct {
	// Currency: CNY, USD.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Discount tier details. The information of each tier includes the duration, discounted quantity, total price, discounted price, and discount details (user discount, official website discount, or final discount).
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitnil" name:"DiscountDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDiskDiscountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiskDiscountResponseParams `json:"Response"`
}

func (r *DescribeDiskDiscountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiskDiscountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksDeniedActionsRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
}

type DescribeDisksDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
}

func (r *DescribeDisksDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksDeniedActionsResponseParams struct {
	// List of operation limits of cloud disks.
	DiskDeniedActionSet []*DiskDeniedActions `json:"DiskDeniedActionSet,omitnil" name:"DiskDeniedActionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDisksDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeDisksDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Filter list
	// disk-id
	// Filter by **cloud disk ID**.
	// Type: String
	// Required: No
	// instance-id
	// Filter by **instance ID**.
	// Type: String
	// Required: No
	// disk-name
	// Filter by **cloud disk name**.
	// Type: String
	// Required: No
	// zone
	// Filter by **availability zone**.
	// Type: String
	// Required: No
	// disk-usage
	// Filter by **cloud disk type**.
	// Type: String
	// Required: No
	// Values: `SYSTEM_DISK` and `DATA_DISK`
	// disk-state
	// Filter by **cloud disk status**.
	// Type: String
	// Required: No
	// Values: See `DiskState` in [Disk](https://intl.cloud.tencent.com/document/api/1207/47576?from_cn_redirect=1#Disk)
	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`. `DiskIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The field by which the cloud disks are sorted. Valid values: "CREATED_TIME" (creation time), "EXPIRED_TIME" (expiration time), "DISK_SIZE" (size of cloud disks). Default value: "CREATED_TIME".
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Sorting order of the output cloud disks. Valid values: "ASC" (ascending order), "DESC" (descending order). Default value: "DESC".
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Filter list
	// disk-id
	// Filter by **cloud disk ID**.
	// Type: String
	// Required: No
	// instance-id
	// Filter by **instance ID**.
	// Type: String
	// Required: No
	// disk-name
	// Filter by **cloud disk name**.
	// Type: String
	// Required: No
	// zone
	// Filter by **availability zone**.
	// Type: String
	// Required: No
	// disk-usage
	// Filter by **cloud disk type**.
	// Type: String
	// Required: No
	// Values: `SYSTEM_DISK` and `DATA_DISK`
	// disk-state
	// Filter by **cloud disk status**.
	// Type: String
	// Required: No
	// Values: See `DiskState` in [Disk](https://intl.cloud.tencent.com/document/api/1207/47576?from_cn_redirect=1#Disk)
	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`. `DiskIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The field by which the cloud disks are sorted. Valid values: "CREATED_TIME" (creation time), "EXPIRED_TIME" (expiration time), "DISK_SIZE" (size of cloud disks). Default value: "CREATED_TIME".
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Sorting order of the output cloud disks. Valid values: "ASC" (ascending order), "DESC" (descending order). Default value: "DESC".
	Order *string `json:"Order,omitnil" name:"Order"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksResponseParams struct {
	// List of cloud disk information.
	DiskSet []*Disk `json:"DiskSet,omitnil" name:"DiskSet"`

	// Number of eligible cloud disks.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeDisksReturnableRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeDisksReturnableRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeDisksReturnableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksReturnableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksReturnableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksReturnableResponseParams struct {
	// List of returnable cloud disks.
	DiskReturnableSet []*DiskReturnable `json:"DiskReturnableSet,omitnil" name:"DiskReturnableSet"`

	// Number of eligible cloud disks.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDisksReturnableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksReturnableResponseParams `json:"Response"`
}

func (r *DescribeDisksReturnableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksReturnableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirewallRulesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeFirewallRulesResponseParams struct {
	// Number of eligible firewall rules.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Firewall rule details list.
	FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitnil" name:"FirewallRuleSet"`

	// Firewall version number.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeFirewallRulesTemplateRequestParams struct {

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

// Predefined struct for user
type DescribeFirewallRulesTemplateResponseParams struct {
	// Number of eligible firewall rules.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Firewall rule details list.
	FirewallRuleSet []*FirewallRuleInfo `json:"FirewallRuleSet,omitnil" name:"FirewallRuleSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFirewallRulesTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirewallRulesTemplateResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeGeneralResourceQuotasRequestParams struct {
	// Resource name list. Values:
	// - `GENERAL_BUNDLE_INSTANCE`: General bundle
	// - `STORAGE_BUNDLE_INSTANCE`:  Storage bundle 
	// - `ENTERPRISE_BUNDLE_INSTANCE`: Enterprise bundle 
	// - `EXCLUSIVE_BUNDLE_INSTANCE`： Dedicated bundle
	// - `BEFAST_BUNDLE_INSTANCE`: BeFast bundle
	// - `USER_KEY_PAIR`: Key pair
	// - `SNAPSHOT`: Snapshot
	// - `BLUEPRINT`: Custom image
	// - `FREE_BLUEPRINT`: Free custom image
	// - `DATA_DISK`: Data disk
	// - `FIREWALL_RULE`: Firewall rules
	ResourceNames []*string `json:"ResourceNames,omitnil" name:"ResourceNames"`
}

type DescribeGeneralResourceQuotasRequest struct {
	*tchttp.BaseRequest
	
	// Resource name list. Values:
	// - `GENERAL_BUNDLE_INSTANCE`: General bundle
	// - `STORAGE_BUNDLE_INSTANCE`:  Storage bundle 
	// - `ENTERPRISE_BUNDLE_INSTANCE`: Enterprise bundle 
	// - `EXCLUSIVE_BUNDLE_INSTANCE`： Dedicated bundle
	// - `BEFAST_BUNDLE_INSTANCE`: BeFast bundle
	// - `USER_KEY_PAIR`: Key pair
	// - `SNAPSHOT`: Snapshot
	// - `BLUEPRINT`: Custom image
	// - `FREE_BLUEPRINT`: Free custom image
	// - `DATA_DISK`: Data disk
	// - `FIREWALL_RULE`: Firewall rules
	ResourceNames []*string `json:"ResourceNames,omitnil" name:"ResourceNames"`
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

// Predefined struct for user
type DescribeGeneralResourceQuotasResponseParams struct {
	// List of general resource quota details.
	GeneralResourceQuotaSet []*GeneralResourceQuota `json:"GeneralResourceQuotaSet,omitnil" name:"GeneralResourceQuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGeneralResourceQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralResourceQuotasResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceLoginKeyPairAttributeRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceLoginKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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

// Predefined struct for user
type DescribeInstanceLoginKeyPairAttributeResponseParams struct {
	// Whether to allow login with the default key pair. Valid values: YES, NO.
	PermitLogin *string `json:"PermitLogin,omitnil" name:"PermitLogin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceLoginKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLoginKeyPairAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstanceVncUrlRequestParams struct {
	// Instance ID, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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

// Predefined struct for user
type DescribeInstanceVncUrlResponseParams struct {
	// Instance VNC URL.
	InstanceVncUrl *string `json:"InstanceVncUrl,omitnil" name:"InstanceVncUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceVncUrlResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesDeniedActionsRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type DescribeInstancesDeniedActionsResponseParams struct {
	// List of instance operation limit details.
	InstanceDeniedActionSet []*InstanceDeniedActions `json:"InstanceDeniedActionSet,omitnil" name:"InstanceDeniedActionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDeniedActionsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesDiskNumRequestParams struct {
	// List of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DescribeInstancesDiskNumRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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
	// Information of all attached disks
	AttachDetailSet []*AttachDetail `json:"AttachDetailSet,omitnil" name:"AttachDetailSet"`

	// Number of attached cloud disks
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeInstancesRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Filter list. 
	// <li>instance-name</li>Filter by the **instance name**. 
	// Type: String 
	// Required: No 
	// <li>private-ip-address</li>Filter by the **private IP of instance primary ENI**. 
	// Type: String 
	// Required: No 
	// <li>public-ip-address</li>Filter by the **public IP of instance primary ENI**. 
	// Type: String 
	// Required: No 
	// <li>zone</li>Filter by the availability zone. 
	// Type: String 
	// Required: No 
	// <li>instance-state</li>Filter by the **instance status**. 
	// Type: String 
	// Required: No 
	// <li>tag-key</li>Filter by the **tag key**. 
	// Type: String 
	// Required: No 
	// <li>tag-value</li>Filter by the **tag value**. 
	// Type: String 
	// Required: No 
	// <li> tag:tag-key</li>Filter by tag key-value pair. The `tag-key` should be replaced with a specific tag key. 
	// Type: String 
	// Required: No 
	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`. You cannot specify both `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Filter list. 
	// <li>instance-name</li>Filter by the **instance name**. 
	// Type: String 
	// Required: No 
	// <li>private-ip-address</li>Filter by the **private IP of instance primary ENI**. 
	// Type: String 
	// Required: No 
	// <li>public-ip-address</li>Filter by the **public IP of instance primary ENI**. 
	// Type: String 
	// Required: No 
	// <li>zone</li>Filter by the availability zone. 
	// Type: String 
	// Required: No 
	// <li>instance-state</li>Filter by the **instance status**. 
	// Type: String 
	// Required: No 
	// <li>tag-key</li>Filter by the **tag key**. 
	// Type: String 
	// Required: No 
	// <li>tag-value</li>Filter by the **tag value**. 
	// Type: String 
	// Required: No 
	// <li> tag:tag-key</li>Filter by tag key-value pair. The `tag-key` should be replaced with a specific tag key. 
	// Type: String 
	// Required: No 
	// Each request can contain up to 10 `Filters` and 100 `Filter.Values`. You cannot specify both `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of instance details
	InstanceSet []*Instance `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesReturnableRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeInstancesReturnableRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeInstancesReturnableResponseParams struct {
	// Number of eligible instances.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of returnable instance details.
	InstanceReturnableSet []*InstanceReturnable `json:"InstanceReturnableSet,omitnil" name:"InstanceReturnableSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesReturnableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesReturnableResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeInstancesTrafficPackagesRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeInstancesTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeInstancesTrafficPackagesResponseParams struct {
	// Number of eligible instance traffic package details.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of instance traffic package details.
	InstanceTrafficPackageSet []*InstanceTrafficPackage `json:"InstanceTrafficPackageSet,omitnil" name:"InstanceTrafficPackageSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesTrafficPackagesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeKeyPairsRequestParams struct {
	// Key pair ID list.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter list.
	// <li>key-id</li>Filter by **key pair ID**.
	// Type: String
	// Required: no
	// <li>key-name</li>Filter by the **key pair name**. Fuzzy match is supported.
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and up to 5 `Filter.Values` for each filter. `KeyIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Key pair ID list.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Filter list.
	// <li>key-id</li>Filter by **key pair ID**.
	// Type: String
	// Required: no
	// <li>key-name</li>Filter by the **key pair name**. Fuzzy match is supported.
	// Type: String
	// Required: no
	// Each request can contain up to 10 `Filters` and up to 5 `Filter.Values` for each filter. `KeyIds` and `Filters` cannot be specified at the same time.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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

// Predefined struct for user
type DescribeKeyPairsResponseParams struct {
	// Number of eligible key pairs.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of key pair details.
	KeyPairSet []*KeyPair `json:"KeyPairSet,omitnil" name:"KeyPairSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeyPairsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeModifyInstanceBundlesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Filter list
	// <li>bundle-id</li>Filter by the **bundle ID**.
	// Type: String
	// Required: No
	// <li>support-platform-type</li>Filter by the **OS type**.
	// Valid values: `LINUX_UNIX` (Linux or Unix), `WINDOWS` (Windows)
	// Type: String
	// Required: No
	// <li>bundle-type</li>Filter by the **bundle type**.
	// Valid values: `GENERAL_BUNDLE` (General bundle), `STORAGE_BUNDLE` (Storage bundle), `ENTERPRISE_BUNDLE` (Enterprise bundle), `EXCLUSIVE_BUNDLE` (Dedicated bundle), `BEFAST_BUNDLE` (BeFast bundle)
	// Type: String
	// Required: No
	// <li>bundle-state</li>Filter by the **bundle status**.
	// Valid values: `ONLINE`, `OFFLINE`
	// Type: String
	// Required: No
	// Each request can contain up to 10 `Filters`, and each filter can have up to 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeModifyInstanceBundlesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Filter list
	// <li>bundle-id</li>Filter by the **bundle ID**.
	// Type: String
	// Required: No
	// <li>support-platform-type</li>Filter by the **OS type**.
	// Valid values: `LINUX_UNIX` (Linux or Unix), `WINDOWS` (Windows)
	// Type: String
	// Required: No
	// <li>bundle-type</li>Filter by the **bundle type**.
	// Valid values: `GENERAL_BUNDLE` (General bundle), `STORAGE_BUNDLE` (Storage bundle), `ENTERPRISE_BUNDLE` (Enterprise bundle), `EXCLUSIVE_BUNDLE` (Dedicated bundle), `BEFAST_BUNDLE` (BeFast bundle)
	// Type: String
	// Required: No
	// <li>bundle-state</li>Filter by the **bundle status**.
	// Valid values: `ONLINE`, `OFFLINE`
	// Type: String
	// Required: No
	// Each request can contain up to 10 `Filters`, and each filter can have up to 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeModifyInstanceBundlesResponseParams struct {
	// Number of matched instances.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// New package details
	ModifyBundleSet []*ModifyBundle `json:"ModifyBundleSet,omitnil" name:"ModifyBundleSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeModifyInstanceBundlesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModifyInstanceBundlesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRegionsRequestParams struct {

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

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// Number of regions.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Region information list.
	RegionSet []*RegionInfo `json:"RegionSet,omitnil" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeResetInstanceBlueprintsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

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
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeResetInstanceBlueprintsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset. Default value: 0. For more information on `Offset`, please see the relevant section in [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, please see the relevant section in the API [Overview](https://intl.cloud.tencent.com/document/product/1207/47578?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

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
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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

// Predefined struct for user
type DescribeResetInstanceBlueprintsResponseParams struct {
	// Number of eligible images.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Image reset information list
	ResetInstanceBlueprintSet []*ResetInstanceBlueprint `json:"ResetInstanceBlueprintSet,omitnil" name:"ResetInstanceBlueprintSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResetInstanceBlueprintsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResetInstanceBlueprintsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeScenesRequestParams struct {
	// List of scene IDs
	SceneIds []*string `json:"SceneIds,omitnil" name:"SceneIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeScenesRequest struct {
	*tchttp.BaseRequest
	
	// List of scene IDs
	SceneIds []*string `json:"SceneIds,omitnil" name:"SceneIds"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeScenesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScenesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScenesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScenesResponseParams struct {
	// List of scenes
	SceneSet []*Scene `json:"SceneSet,omitnil" name:"SceneSet"`

	// Total number of scenes
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScenesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScenesResponseParams `json:"Response"`
}

func (r *DescribeScenesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScenesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsDeniedActionsRequestParams struct {
	// Snapshot ID list, which can be queried through `DescribeSnapshots`.
	SnapshotIds []*string `json:"SnapshotIds,omitnil" name:"SnapshotIds"`
}

type DescribeSnapshotsDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot ID list, which can be queried through `DescribeSnapshots`.
	SnapshotIds []*string `json:"SnapshotIds,omitnil" name:"SnapshotIds"`
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

// Predefined struct for user
type DescribeSnapshotsDeniedActionsResponseParams struct {
	// List of snapshot operation limit details.
	SnapshotDeniedActionSet []*SnapshotDeniedActions `json:"SnapshotDeniedActionSet,omitnil" name:"SnapshotDeniedActionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSnapshotsDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsDeniedActionsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// List of IDs of snapshots to be queried.
	// You cannot specify `SnapshotIds` and `Filters` at the same time.
	SnapshotIds []*string `json:"SnapshotIds,omitnil" name:"SnapshotIds"`

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
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of snapshots to be queried.
	// You cannot specify `SnapshotIds` and `Filters` at the same time.
	SnapshotIds []*string `json:"SnapshotIds,omitnil" name:"SnapshotIds"`

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
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// Number of snapshots.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of snapshot details.
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitnil" name:"SnapshotSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeZonesRequestParams struct {
	// Sorting field. Valid values:
	// <li>`ZONE`: Sort by the availability zone.
	// <li>`INSTANCE_DISPLAY_LABEL`: Sort by visibility labels (`HIDDEN`, `NORMAL` and `SELECTED`). Default: ['HIDDEN', 'NORMAL', 'SELECTED'].
	// The default value is `ZONE`.
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Specifies how availability zones are listed. Valid values:
	// <li>ASC: Ascending sort. 
	// <li>DESC: Descending sort.
	// The default value is `ASC`.
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// Sorting field. Valid values:
	// <li>`ZONE`: Sort by the availability zone.
	// <li>`INSTANCE_DISPLAY_LABEL`: Sort by visibility labels (`HIDDEN`, `NORMAL` and `SELECTED`). Default: ['HIDDEN', 'NORMAL', 'SELECTED'].
	// The default value is `ZONE`.
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// Specifies how availability zones are listed. Valid values:
	// <li>ASC: Ascending sort. 
	// <li>DESC: Descending sort.
	// The default value is `ASC`.
	Order *string `json:"Order,omitnil" name:"Order"`
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
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of AZs
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of AZ details
	ZoneInfoSet []*ZoneInfo `json:"ZoneInfoSet,omitnil" name:"ZoneInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
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

// Predefined struct for user
type DetachCcnRequestParams struct {
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

type DetachCcnRequest struct {
	*tchttp.BaseRequest
	
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
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

// Predefined struct for user
type DetachCcnResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachCcnResponse struct {
	*tchttp.BaseResponse
	Response *DetachCcnResponseParams `json:"Response"`
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

// Predefined struct for user
type DetachDisksRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type DetailPrice struct {
	// Values: 
	// <li>"DiskSpace": Cloud disk space</li>
	// <li>"DiskBackupQuota": Cloud disk backups</li>
	PriceName *string `json:"PriceName,omitnil" name:"PriceName"`

	// Official unit price of the billable item
	OriginUnitPrice *float64 `json:"OriginUnitPrice,omitnil" name:"OriginUnitPrice"`

	// Official total price of the billable item
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// Discount of the billable item
	Discount *float64 `json:"Discount,omitnil" name:"Discount"`

	// Discounted total price of the billable item
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`
}

// Predefined struct for user
type DisassociateInstancesKeyPairsRequestParams struct {
	// Key pair ID list. Each request can contain up to 100 key pairs.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Key pair ID list. Each request can contain up to 100 key pairs.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`

	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type DisassociateInstancesKeyPairsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateInstancesKeyPairsResponseParams `json:"Response"`
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

type DiscountDetail struct {
	// Billing duration.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// Billing unit.
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Total price.
	TotalCost *float64 `json:"TotalCost,omitnil" name:"TotalCost"`

	// Discounted total price.
	RealTotalCost *float64 `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// Discount.
	Discount *int64 `json:"Discount,omitnil" name:"Discount"`

	// Discount details.
	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitnil" name:"PolicyDetail"`
}

type Disk struct {
	// Disk ID
	DiskId *string `json:"DiskId,omitnil" name:"DiskId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Availability zone
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Disk name
	DiskName *string `json:"DiskName,omitnil" name:"DiskName"`

	// Disk type
	DiskUsage *string `json:"DiskUsage,omitnil" name:"DiskUsage"`

	// Disk media type
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Disk payment type
	DiskChargeType *string `json:"DiskChargeType,omitnil" name:"DiskChargeType"`

	// Disk size
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Renewal flag
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// Disk status. Values: 
	// <li>`PENDING`: Creating</li>
	// <li>`UNATTACHED`: Not attached</li>
	// <li>`ATTACHING`: Attaching</li>
	// <li>`ATTACHED`: Attached</li>
	// <li>`DETACHING`: Detaching</li>
	// <li>`SHUTDOWN`: Isolated</li>
	// <li>`CREATED_FAILED`: Failed to create</li>
	// <li>`TERMINATING`: Terminating</li>
	// <li>`DELETING`: Deleting</li>
	// <li>`FREEZING`: Freezing</li>
	DiskState *string `json:"DiskState,omitnil" name:"DiskState"`

	// Whether the disk is attached to an instance
	Attached *bool `json:"Attached,omitnil" name:"Attached"`

	// Whether to release the disk along with the instance
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil" name:"DeleteWithInstance"`

	// Last operation
	LatestOperation *string `json:"LatestOperation,omitnil" name:"LatestOperation"`

	// Last operation status
	LatestOperationState *string `json:"LatestOperationState,omitnil" name:"LatestOperationState"`

	// Last request ID
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil" name:"LatestOperationRequestId"`

	// Creation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Expiration time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// Isolation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitnil" name:"IsolatedTime"`

	// Total disk backups
	DiskBackupCount *int64 `json:"DiskBackupCount,omitnil" name:"DiskBackupCount"`

	// Disk backup quota
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`
}

type DiskChargePrepaid struct {
	// Purchase duration.
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Auto-Renewal flag. Valid values:
	// 
	// `NOTIFY_AND_AUTO_RENEW`: Trigger expiration notification and renew automatically
	// `NOTIFY_AND_MANUAL_RENEW`: Trigger expiration notification but do not renew
	// `u200cDISABLE_NOTIFY_AND_AUTO_RENEW`: Neither trigger expiration notification nor renew
	// 
	// Default: `NOTIFY_AND_MANUAL_RENEW`. If `NOTIFY_AND_AUTO_RENEW` is specified, the instance is automatically renewed on a monthly basis when the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// Purchase duration unit. Default value: "m" (month)
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`
}

type DiskConfig struct {
	// Availability zone.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Cloud disk type.
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Cloud disk sale status.
	DiskSalesState *string `json:"DiskSalesState,omitnil" name:"DiskSalesState"`

	// Maximum cloud disk size.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil" name:"MaxDiskSize"`

	// Minimum cloud disk size.
	MinDiskSize *int64 `json:"MinDiskSize,omitnil" name:"MinDiskSize"`

	// Cloud disk increment.
	DiskStepSize *int64 `json:"DiskStepSize,omitnil" name:"DiskStepSize"`
}

type DiskDeniedActions struct {
	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitnil" name:"DiskId"`

	// List of operation limits.
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil" name:"DeniedActions"`
}

type DiskPrice struct {
	// Cloud disk unit price.
	OriginalDiskPrice *float64 `json:"OriginalDiskPrice,omitnil" name:"OriginalDiskPrice"`

	// Total cloud disk price.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// Discount.
	Discount *float64 `json:"Discount,omitnil" name:"Discount"`

	// Discounted total price.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`

	// Detailed billing items
	DetailPrices []*DetailPrice `json:"DetailPrices,omitnil" name:"DetailPrices"`
}

type DiskReturnable struct {
	// Cloud disk ID.
	DiskId *string `json:"DiskId,omitnil" name:"DiskId"`

	// Whether the cloud disk can be returned.
	IsReturnable *bool `json:"IsReturnable,omitnil" name:"IsReturnable"`

	// Error code of cloud disk return failure.
	ReturnFailCode *int64 `json:"ReturnFailCode,omitnil" name:"ReturnFailCode"`

	// Error message of cloud disk return failure.
	ReturnFailMessage *string `json:"ReturnFailMessage,omitnil" name:"ReturnFailMessage"`
}

type DockerContainerConfiguration struct {
	// Container image address
	ContainerImage *string `json:"ContainerImage,omitnil" name:"ContainerImage"`

	// Container name
	ContainerName *string `json:"ContainerName,omitnil" name:"ContainerName"`

	// List of environment variables
	Envs []*ContainerEnv `json:"Envs,omitnil" name:"Envs"`

	// List of mappings of container ports and host ports
	PublishPorts []*DockerContainerPublishPort `json:"PublishPorts,omitnil" name:"PublishPorts"`

	// List of container mount volumes
	Volumes []*DockerContainerVolume `json:"Volumes,omitnil" name:"Volumes"`

	// The command to run
	Command *string `json:"Command,omitnil" name:"Command"`
}

type DockerContainerPublishPort struct {
	// Host port
	HostPort *int64 `json:"HostPort,omitnil" name:"HostPort"`

	// Container port
	ContainerPort *int64 `json:"ContainerPort,omitnil" name:"ContainerPort"`

	// External IP. It defaults to 0.0.0.0.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// The protocol defaults to `tcp`. Valid values: `tcp`, `udp` and `sctp`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type DockerContainerVolume struct {
	// Container path
	ContainerPath *string `json:"ContainerPath,omitnil" name:"ContainerPath"`

	// Host path
	HostPath *string `json:"HostPath,omitnil" name:"HostPath"`
}

type Filter struct {
	// Field to be filtered.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Filter value of field.
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FirewallRule struct {
	// Protocol. Valid values: TCP, UDP, ICMP, ALL.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Port. Valid values: ALL, one single port, multiple ports separated by commas, or port range indicated by a minus sign
	Port *string `json:"Port,omitnil" name:"Port"`

	// IP range or IP (mutually exclusive). Default value: 0.0.0.0/0, which indicates all sources.
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// Valid values: ACCEPT, DROP. Default value: ACCEPT.
	Action *string `json:"Action,omitnil" name:"Action"`

	// Firewall rule description.
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitnil" name:"FirewallRuleDescription"`
}

type FirewallRuleInfo struct {
	// Application type. Valid values: custom, HTTP (80), HTTPS (443), Linux login (22), Windows login (3389), MySQL (3306), SQL Server (1433), all TCP ports, all UDP ports, Ping-ICMP, ALL.
	AppType *string `json:"AppType,omitnil" name:"AppType"`

	// Protocol. Valid values: TCP, UDP, ICMP, ALL.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Port. Valid values: ALL, one single port, multiple ports separated by commas, or port range indicated by a minus sign
	Port *string `json:"Port,omitnil" name:"Port"`

	// IP range or IP (mutually exclusive). Default value: 0.0.0.0/0, which indicates all sources.
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// Valid values: ACCEPT, DROP. Default value: ACCEPT.
	Action *string `json:"Action,omitnil" name:"Action"`

	// Firewall rule description.
	FirewallRuleDescription *string `json:"FirewallRuleDescription,omitnil" name:"FirewallRuleDescription"`
}

type GeneralResourceQuota struct {
	// Resource name.
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// Number of available resources.
	ResourceQuotaAvailable *int64 `json:"ResourceQuotaAvailable,omitnil" name:"ResourceQuotaAvailable"`

	// Total number of resources.
	ResourceQuotaTotal *int64 `json:"ResourceQuotaTotal,omitnil" name:"ResourceQuotaTotal"`
}

// Predefined struct for user
type ImportKeyPairRequestParams struct {
	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitnil" name:"KeyName"`

	// Public key content of the key pair, which is in the OpenSSH RSA format.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// Key pair name, which can contain up to 25 digits, letters, and underscores.
	KeyName *string `json:"KeyName,omitnil" name:"KeyName"`

	// Public key content of the key pair, which is in the OpenSSH RSA format.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`
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

// Predefined struct for user
type ImportKeyPairResponseParams struct {
	// Key pair ID.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *ImportKeyPairResponseParams `json:"Response"`
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

// Predefined struct for user
type InquirePriceCreateBlueprintRequestParams struct {
	// Number of custom images. Default value: 1.
	BlueprintCount *int64 `json:"BlueprintCount,omitnil" name:"BlueprintCount"`
}

type InquirePriceCreateBlueprintRequest struct {
	*tchttp.BaseRequest
	
	// Number of custom images. Default value: 1.
	BlueprintCount *int64 `json:"BlueprintCount,omitnil" name:"BlueprintCount"`
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

// Predefined struct for user
type InquirePriceCreateBlueprintResponseParams struct {
	// Custom image price.
	BlueprintPrice *BlueprintPrice `json:"BlueprintPrice,omitnil" name:"BlueprintPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceCreateBlueprintResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateBlueprintResponseParams `json:"Response"`
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

// Predefined struct for user
type InquirePriceCreateDisksRequestParams struct {
	// Cloud disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Cloud disk media type. Valid values: "CLOUD_PREMIUM" (premium cloud storage), "CLOUD_SSD" (SSD cloud disk).
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Parameter settings for purchasing the monthly subscribed cloud disk.
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil" name:"DiskChargePrepaid"`

	// Number of cloud disks. Default value: 1.
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// Specify the quota of disk backups. No quota if it’s left empty. Only one quota is allowed.
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`
}

type InquirePriceCreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// Cloud disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Cloud disk media type. Valid values: "CLOUD_PREMIUM" (premium cloud storage), "CLOUD_SSD" (SSD cloud disk).
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// Parameter settings for purchasing the monthly subscribed cloud disk.
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitnil" name:"DiskChargePrepaid"`

	// Number of cloud disks. Default value: 1.
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// Specify the quota of disk backups. No quota if it’s left empty. Only one quota is allowed.
	DiskBackupQuota *int64 `json:"DiskBackupQuota,omitnil" name:"DiskBackupQuota"`
}

func (r *InquirePriceCreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskSize")
	delete(f, "DiskType")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskCount")
	delete(f, "DiskBackupQuota")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDisksResponseParams struct {
	// Cloud disk price.
	DiskPrice *DiskPrice `json:"DiskPrice,omitnil" name:"DiskPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceCreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDisksResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateInstancesRequestParams struct {
	// Instance package ID.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Parameter setting for prepaid mode. This parameter can specify the purchase period, whether to enable auto-renewal, and other attributes of the monthly subscribed instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Number of instances to be created. Default value: 1.
	InstanceCount *int64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// Application image ID, which is required if a paid application image is used and can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`
}

type InquirePriceCreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance package ID.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Parameter setting for prepaid mode. This parameter can specify the purchase period, whether to enable auto-renewal, and other attributes of the monthly subscribed instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Number of instances to be created. Default value: 1.
	InstanceCount *int64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// Application image ID, which is required if a paid application image is used and can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`
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
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceCount")
	delete(f, "BlueprintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateInstancesResponseParams struct {
	// Price query information.
	Price *Price `json:"Price,omitnil" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceCreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type InquirePriceRenewDisksRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Parameter settings for renewing the monthly subscribed cloud disk.
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil" name:"RenewDiskChargePrepaid"`
}

type InquirePriceRenewDisksRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Parameter settings for renewing the monthly subscribed cloud disk.
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil" name:"RenewDiskChargePrepaid"`
}

func (r *InquirePriceRenewDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "RenewDiskChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDisksResponseParams struct {
	// Cloud disk price.
	DiskPrice *DiskPrice `json:"DiskPrice,omitnil" name:"DiskPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceRenewDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewDisksResponseParams `json:"Response"`
}

func (r *InquirePriceRenewDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewInstancesRequestParams struct {
	// IDs of the instances to be renewed. Each request can contain up to 50 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Parameter setting for prepaid mode. This parameter can specify the renewal period, whether to enable auto-renewal, and other attributes of the monthly subscribed instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Whether to renew the data disk. Default: `false`.
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil" name:"RenewDataDisk"`

	// Whether to align the data disk expiration with the instance expiration time. Default: `false`.
	AlignInstanceExpiredTime *bool `json:"AlignInstanceExpiredTime,omitnil" name:"AlignInstanceExpiredTime"`
}

type InquirePriceRenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// IDs of the instances to be renewed. Each request can contain up to 50 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Parameter setting for prepaid mode. This parameter can specify the renewal period, whether to enable auto-renewal, and other attributes of the monthly subscribed instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Whether to renew the data disk. Default: `false`.
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil" name:"RenewDataDisk"`

	// Whether to align the data disk expiration with the instance expiration time. Default: `false`.
	AlignInstanceExpiredTime *bool `json:"AlignInstanceExpiredTime,omitnil" name:"AlignInstanceExpiredTime"`
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
	delete(f, "RenewDataDisk")
	delete(f, "AlignInstanceExpiredTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewInstancesResponseParams struct {
	// Price information. It defaults to the price information of the first instance in the list.
	Price *Price `json:"Price,omitnil" name:"Price"`

	// List of data disk price information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataDiskPriceSet []*DataDiskPrice `json:"DataDiskPriceSet,omitnil" name:"DataDiskPriceSet"`

	// Price list of the instances to be renewed.
	// Note: This field may return `null`, indicating that no valid value was found.
	InstancePriceDetailSet []*InstancePriceDetail `json:"InstancePriceDetailSet,omitnil" name:"InstancePriceDetailSet"`

	// Total price
	TotalPrice *TotalPrice `json:"TotalPrice,omitnil" name:"TotalPrice"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewInstancesResponseParams `json:"Response"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Package ID.
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Image ID.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// Number of instance CPU cores.
	CPU *int64 `json:"CPU,omitnil" name:"CPU"`

	// Instance memory capacity in GB.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance billing mode. Valid values: 
	// PREPAID: prepaid (i.e., monthly subscription).
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// Instance system disk information.
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// Private IP of instance primary ENI. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateAddresses []*string `json:"PrivateAddresses,omitnil" name:"PrivateAddresses"`

	// Public IP of instance primary ENI. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicAddresses []*string `json:"PublicAddresses,omitnil" name:"PublicAddresses"`

	// Instance bandwidth information.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// Auto-Renewal flag. Valid values: 
	// NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically  
	// NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// Instance login settings.
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// Instance status. Valid values: 
	// <li>PENDING: Creating</li><li>LAUNCH_FAILED: Failed to create</li><li>RUNNING: Running</li><li>STOPPED: Shut down</li><li>STARTING: Starting up</li><li>STOPPING: Shutting down</li><li>REBOOTING: Restarting</li><li>SHUTDOWN: Shutdown and to be terminated</li><li>TERMINATING: Terminating</li><li>DELETING: Deleting</li><li>FREEZING: Frozen</li><li>ENTER_RESCUE_MODE: Entering the rescue mode</li><li>RESCUE_MODE: Rescue mode</li><li>EXIT_RESCUE_MODE: Exiting from the rescue mode</li>
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// Globally unique ID of instance.
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// Last instance operation, such as `StopInstances` and `ResetInstance`. Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperation *string `json:"LatestOperation,omitnil" name:"LatestOperation"`

	// Last instance operation status. Valid values: 
	// SUCCESS: operation succeeded 
	// OPERATING: the operation is being executed 
	// FAILED: operation failed 
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationState *string `json:"LatestOperationState,omitnil" name:"LatestOperationState"`

	// Unique request ID for the last operation of the instance. 
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil" name:"LatestOperationRequestId"`

	// Isolation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsolatedTime *string `json:"IsolatedTime,omitnil" name:"IsolatedTime"`

	// Creation time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Expiration time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// OS type, such as LINUX_UNIX and WINDOWS.
	PlatformType *string `json:"PlatformType,omitnil" name:"PlatformType"`

	// OS type.
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// OS name.
	OsName *string `json:"OsName,omitnil" name:"OsName"`

	// AZ.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// The list of tags associated with the instance
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Obtain instance status
	// <li>NORMAL: The instance is normal</li><li>NETWORK_RESTRICT: The instance is blocked from the network.</li>
	InstanceRestrictState *string `json:"InstanceRestrictState,omitnil" name:"InstanceRestrictState"`
}

type InstanceChargePrepaid struct {
	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Auto-Renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew <br><li>DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically<br><br>Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed monthly if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type InstanceDeniedActions struct {
	// Instance ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of operation limits.
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil" name:"DeniedActions"`
}

type InstancePrice struct {
	// Original package unit price.
	OriginalBundlePrice *float64 `json:"OriginalBundlePrice,omitnil" name:"OriginalBundlePrice"`

	// Original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// Discount.
	Discount *int64 `json:"Discount,omitnil" name:"Discount"`

	// Discounted price.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`

	// Currency unit. Valid values: `CNY` and `USD`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil" name:"Currency"`
}

type InstancePriceDetail struct {
	// Instance ID.
	// Note: This field may return `null`, indicating that no valid value was found.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Price query information.
	// Note: This field may return `null`, indicating that no valid value was found.
	InstancePrice *InstancePrice `json:"InstancePrice,omitnil" name:"InstancePrice"`

	// Tiered-pricing details. The information of each tier includes the billable period, discount percentage, total price, discounted price, and discount details (`UserDiscount`, `CommonDiscount` and `FinalDiscount`).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiscountDetail []*DiscountDetail `json:"DiscountDetail,omitnil" name:"DiscountDetail"`
}

type InstanceReturnable struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether the instance can be returned.
	IsReturnable *bool `json:"IsReturnable,omitnil" name:"IsReturnable"`

	// Error code of instance return failure.
	ReturnFailCode *int64 `json:"ReturnFailCode,omitnil" name:"ReturnFailCode"`

	// Error message of instance return failure.
	ReturnFailMessage *string `json:"ReturnFailMessage,omitnil" name:"ReturnFailMessage"`
}

type InstanceTrafficPackage struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of traffic package details.
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitnil" name:"TrafficPackageSet"`
}

type InternetAccessible struct {
	// Network billing mode. Valid values:
	// <li>Bill by traffic package: TRAFFIC_POSTPAID_BY_HOUR</li>
	// <li>Bill by bandwidth: BANDWIDTH_POSTPAID_BY_HOUR</li>
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// Public network outbound bandwidth cap in Mbps.
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil" name:"PublicIpAssigned"`
}

// Predefined struct for user
type IsolateDisksRequestParams struct {
	// IDs of cloud disks. The value can be obtained from the `InstanceId` parameter returned by the [DescribeDisks](https://intl.cloud.tencent.com/document/product/1207/66093?from_cn_redirect=1) API. Up to 20 disks can be processed at a time.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
}

type IsolateDisksRequest struct {
	*tchttp.BaseRequest
	
	// IDs of cloud disks. The value can be obtained from the `InstanceId` parameter returned by the [DescribeDisks](https://intl.cloud.tencent.com/document/product/1207/66093?from_cn_redirect=1) API. Up to 20 disks can be processed at a time.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
}

func (r *IsolateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateDisksResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDisksResponseParams `json:"Response"`
}

func (r *IsolateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstancesRequestParams struct {
	// IDs of target instances. You can get the IDs from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 20 instances can be specified at the same time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Whether to return data disks mounted on the instance together with the instance. Valid values: 
	// `TRUE`: Return the mounted data disks at the same time 
	// `FALSE`: Do not return the mounted data disks at the same time 
	// Default value: `TRUE`.
	IsolateDataDisk *bool `json:"IsolateDataDisk,omitnil" name:"IsolateDataDisk"`
}

type IsolateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// IDs of target instances. You can get the IDs from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 20 instances can be specified at the same time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Whether to return data disks mounted on the instance together with the instance. Valid values: 
	// `TRUE`: Return the mounted data disks at the same time 
	// `FALSE`: Do not return the mounted data disks at the same time 
	// Default value: `TRUE`.
	IsolateDataDisk *bool `json:"IsolateDataDisk,omitnil" name:"IsolateDataDisk"`
}

func (r *IsolateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "IsolateDataDisk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *IsolateInstancesResponseParams `json:"Response"`
}

func (r *IsolateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyPair struct {
	// Key pair ID, which is the unique identifier of a key pair.
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// Key pair name.
	KeyName *string `json:"KeyName,omitnil" name:"KeyName"`

	// Public key (in plain text) of key pair.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// List of IDs of instances associated with the key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitnil" name:"AssociatedInstanceIds"`

	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Private key of key pair.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

type LoginConfiguration struct {
	// <li>`YES`: Random password. In this case, `Password` cannot be specified. </li>
	// <li>`No`: Custom. `Password` must be specified. </li>
	AutoGeneratePassword *string `json:"AutoGeneratePassword,omitnil" name:"AutoGeneratePassword"`

	// Instance login password. 
	// For Windows instances, the password must contain 12 to 30 characters of the following types. It cannot start with “/” and cannot include the username. 
	// <li>Lowercase letters: [a–z]</li>
	// <li>Uppercase letters: [A–Z]</li>
	// <li>Digits: 0-9</li>
	// <li>Symbols: ()`~!@#$%^&*-+=_|{}[]:;'<>,.?/</li>
	Password *string `json:"Password,omitnil" name:"Password"`


	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`
}

type LoginSettings struct {
	// Key ID list. After a key is associated, you can use it to access the instance. Note: this field may return [], indicating that no valid values can be obtained.
	KeyIds []*string `json:"KeyIds,omitnil" name:"KeyIds"`
}

// Predefined struct for user
type ModifyBlueprintAttributeRequestParams struct {
	// Image ID, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// New image name, which can contain up to 60 characters.
	BlueprintName *string `json:"BlueprintName,omitnil" name:"BlueprintName"`

	// New image description, which can contain up to 60 characters.
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyBlueprintAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Image ID, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`

	// New image name, which can contain up to 60 characters.
	BlueprintName *string `json:"BlueprintName,omitnil" name:"BlueprintName"`

	// New image description, which can contain up to 60 characters.
	Description *string `json:"Description,omitnil" name:"Description"`
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

// Predefined struct for user
type ModifyBlueprintAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBlueprintAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlueprintAttributeResponseParams `json:"Response"`
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
	ModifyPrice *Price `json:"ModifyPrice,omitnil" name:"ModifyPrice"`

	// Package change status. Valid values:
	// <li>SOLD_OUT: packages are sold out</li>
	// <li>AVAILABLE: packages can be changed</li>
	// <li>UNAVAILABLE: packages cannot be changed currently</li>
	ModifyBundleState *string `json:"ModifyBundleState,omitnil" name:"ModifyBundleState"`

	// Package information.
	Bundle *Bundle `json:"Bundle,omitnil" name:"Bundle"`

	// The reason of package changing failure. It’s empty if the package change status is `AVAILABLE`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	NotSupportModifyMessage *string `json:"NotSupportModifyMessage,omitnil" name:"NotSupportModifyMessage"`
}

// Predefined struct for user
type ModifyDisksAttributeRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Cloud disk name.
	DiskName *string `json:"DiskName,omitnil" name:"DiskName"`
}

type ModifyDisksAttributeRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Cloud disk name.
	DiskName *string `json:"DiskName,omitnil" name:"DiskName"`
}

func (r *ModifyDisksAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "DiskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisksAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDisksAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisksAttributeResponseParams `json:"Response"`
}

func (r *ModifyDisksAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksRenewFlagRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Whether Auto-Renewal is enabled 
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type ModifyDisksRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Whether Auto-Renewal is enabled 
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

func (r *ModifyDisksRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisksRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisksRenewFlagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDisksRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisksRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyDisksRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisksRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFirewallRuleDescriptionRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule.
	FirewallRule *FirewallRule `json:"FirewallRule,omitnil" name:"FirewallRule"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
}

type ModifyFirewallRuleDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule.
	FirewallRule *FirewallRule `json:"FirewallRule,omitnil" name:"FirewallRule"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
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

// Predefined struct for user
type ModifyFirewallRuleDescriptionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyFirewallRuleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFirewallRuleDescriptionResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyFirewallRulesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
}

type ModifyFirewallRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Firewall rule list.
	FirewallRules []*FirewallRule `json:"FirewallRules,omitnil" name:"FirewallRules"`

	// Current firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the rule from expiring. If it is left empty, conflicts will not be considered.
	FirewallVersion *uint64 `json:"FirewallVersion,omitnil" name:"FirewallVersion"`
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

// Predefined struct for user
type ModifyFirewallRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyFirewallRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFirewallRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyInstancesAttributeRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Instance name, which is customizable and can contain up to 60 characters.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Instance name, which is customizable and can contain up to 60 characters.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
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

// Predefined struct for user
type ModifyInstancesAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyInstancesBundleRequestParams struct {
	// IDs of target instances. You can get the IDs from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 15 instances can be specified at the same time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of bundles to change. You can get the IDs from the `BundleId` returned by the [DescribeBundles](https://intl.cloud.tencent.com/document/api/1207/47575?from_cn_redirect=1).
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Whether to use existing vouchers under the current account automatically. Valid values: 
	// `true`: Deduct from existing vouchers automatically 
	// `false`: Do not deduct from existing vouchers automatically 
	// Default value: `false`.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

type ModifyInstancesBundleRequest struct {
	*tchttp.BaseRequest
	
	// IDs of target instances. You can get the IDs from the `InstanceId` parameter returned by the `DescribeInstances` API. Up to 15 instances can be specified at the same time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of bundles to change. You can get the IDs from the `BundleId` returned by the [DescribeBundles](https://intl.cloud.tencent.com/document/api/1207/47575?from_cn_redirect=1).
	BundleId *string `json:"BundleId,omitnil" name:"BundleId"`

	// Whether to use existing vouchers under the current account automatically. Valid values: 
	// `true`: Deduct from existing vouchers automatically 
	// `false`: Do not deduct from existing vouchers automatically 
	// Default value: `false`.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

func (r *ModifyInstancesBundleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesBundleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "BundleId")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesBundleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesBundleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstancesBundleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesBundleResponseParams `json:"Response"`
}

func (r *ModifyInstancesBundleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesBundleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesLoginKeyPairAttributeRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Whether to allow login with the default key pair. Valid values: YES: yes; NO: no.
	PermitLogin *string `json:"PermitLogin,omitnil" name:"PermitLogin"`
}

type ModifyInstancesLoginKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Whether to allow login with the default key pair. Valid values: YES: yes; NO: no.
	PermitLogin *string `json:"PermitLogin,omitnil" name:"PermitLogin"`
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

// Predefined struct for user
type ModifyInstancesLoginKeyPairAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstancesLoginKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesLoginKeyPairAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyInstancesRenewFlagRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Auto-Renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed monthly if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Auto-Renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed monthly if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
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

// Predefined struct for user
type ModifyInstancesRenewFlagResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesRenewFlagResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifySnapshotAttributeRequestParams struct {
	// Snapshot ID, which can be queried through `DescribeSnapshots`.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`

	// New snapshot name, which can contain up to 60 characters.
	SnapshotName *string `json:"SnapshotName,omitnil" name:"SnapshotName"`
}

type ModifySnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Snapshot ID, which can be queried through `DescribeSnapshots`.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`

	// New snapshot name, which can contain up to 60 characters.
	SnapshotName *string `json:"SnapshotName,omitnil" name:"SnapshotName"`
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

// Predefined struct for user
type ModifySnapshotAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type PolicyDetail struct {
	// User discount.
	UserDiscount *int64 `json:"UserDiscount,omitnil" name:"UserDiscount"`

	// Public discount.
	CommonDiscount *int64 `json:"CommonDiscount,omitnil" name:"CommonDiscount"`

	// Final discount.
	FinalDiscount *int64 `json:"FinalDiscount,omitnil" name:"FinalDiscount"`

	// Activity discount. The value `null` indicates no discount.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ActivityDiscount *float64 `json:"ActivityDiscount,omitnil" name:"ActivityDiscount"`

	// Discount type.
	// Valid values: `user` (user discount), `common` (discount displayed on the official website), `activity` (activity discount), `null` (no discount).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiscountType *string `json:"DiscountType,omitnil" name:"DiscountType"`
}

type Price struct {
	// Instance price.
	InstancePrice *InstancePrice `json:"InstancePrice,omitnil" name:"InstancePrice"`
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RebootInstancesResponseParams `json:"Response"`
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
	Region *string `json:"Region,omitnil" name:"Region"`

	// Region description, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// Region availability status. Its value can only be `AVAILABLE`.
	RegionState *string `json:"RegionState,omitnil" name:"RegionState"`

	// Whether the region is in the Chinese mainland
	IsChinaMainland *bool `json:"IsChinaMainland,omitnil" name:"IsChinaMainland"`
}

type RenewDiskChargePrepaid struct {
	// Renewal period
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Whether to renew the disk automatically. Values:
	// 
	// `NOTIFY_AND_AUTO_RENEW`: Trigger expiration notification and renew automatically; `NOTIFY_AND_MANUAL_RENEW`: Trigger expiration notification but do not renew; `DISABLE_NOTIFY_AND_MANUAL_RENEW`: Do not trigger the notification and do not renew.
	// 
	// Default: `NOTIFY_AND_MANUAL_RENEW`. If `NOTIFY_AND_AUTO_RENEW` is specified, the instance is automatically renewed on a monthly basis when the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// Unit of the period. Values: `m` (month).
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// Expiration time of the current instance, such as "2018-01-01 00:00:00". Specify this parameter to align the expiration time of the instance and attached cloud disks. `CurInstanceDeadline` and `Period` cannot be both specified.
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitnil" name:"CurInstanceDeadline"`
}

// Predefined struct for user
type RenewDisksRequestParams struct {
	// IDs of cloud disks. The value can be obtained from the `DiskId` parameter returned by the [DescribeDisks](https://intl.cloud.tencent.com/document/product/1207/66093?from_cn_redirect=1) API. Up to 50 disks can be requested at a time.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Parameter settings for renewing the monthly subscribed cloud disk.
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil" name:"RenewDiskChargePrepaid"`

	// Whether to use the vouchers automatically. It defaults to No.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

type RenewDisksRequest struct {
	*tchttp.BaseRequest
	
	// IDs of cloud disks. The value can be obtained from the `DiskId` parameter returned by the [DescribeDisks](https://intl.cloud.tencent.com/document/product/1207/66093?from_cn_redirect=1) API. Up to 50 disks can be requested at a time.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`

	// Parameter settings for renewing the monthly subscribed cloud disk.
	RenewDiskChargePrepaid *RenewDiskChargePrepaid `json:"RenewDiskChargePrepaid,omitnil" name:"RenewDiskChargePrepaid"`

	// Whether to use the vouchers automatically. It defaults to No.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

func (r *RenewDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "RenewDiskChargePrepaid")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewDisksResponse struct {
	*tchttp.BaseResponse
	Response *RenewDisksResponseParams `json:"Response"`
}

func (r *RenewDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesRequestParams struct {
	// IDs of one or more instances to be operated. The value can be obtained from the `InstanceId` parameter returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API. Up to 100 instances can be requested at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Prepaid mode, i.e., monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. It is required for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Whether to renew elastic data disks. Values: 
	// `TRUE`: Renew the elastic data disks attached to the instance as well when the related instance is renewed.
	// `FALSE`: Do not renew the elastic data disks attached to the instance as well when the related instance is renewed.
	// Default: `TRUE`
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil" name:"RenewDataDisk"`

	// Whether to automatically use vouchers. Values:
	// `TRUE`: Use vouchers for payment automatically.
	// `FALSE`: Do not use vouchers for payment automatically.
	// Default: `FALSE`.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// IDs of one or more instances to be operated. The value can be obtained from the `InstanceId` parameter returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API. Up to 100 instances can be requested at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Prepaid mode, i.e., monthly subscription. This parameter can specify the purchase period and other attributes such as auto-renewal. It is required for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// Whether to renew elastic data disks. Values: 
	// `TRUE`: Renew the elastic data disks attached to the instance as well when the related instance is renewed.
	// `FALSE`: Do not renew the elastic data disks attached to the instance as well when the related instance is renewed.
	// Default: `TRUE`
	RenewDataDisk *bool `json:"RenewDataDisk,omitnil" name:"RenewDataDisk"`

	// Whether to automatically use vouchers. Values:
	// `TRUE`: Use vouchers for payment automatically.
	// `FALSE`: Do not use vouchers for payment automatically.
	// Default: `FALSE`.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	delete(f, "RenewDataDisk")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstancesResponseParams `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnRequestParams struct {
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

type ResetAttachCcnRequest struct {
	*tchttp.BaseRequest
	
	// CCN instance ID.
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
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

// Predefined struct for user
type ResetAttachCcnResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetAttachCcnResponse struct {
	*tchttp.BaseResponse
	Response *ResetAttachCcnResponseParams `json:"Response"`
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
	BlueprintInfo *Blueprint `json:"BlueprintInfo,omitnil" name:"BlueprintInfo"`

	// Whether the image can be reset as the target image
	IsResettable *bool `json:"IsResettable,omitnil" name:"IsResettable"`

	// Non-Resettable flag. If the image is resettable, it will be ""
	NonResettableMessage *string `json:"NonResettableMessage,omitnil" name:"NonResettableMessage"`
}

// Predefined struct for user
type ResetInstanceRequestParams struct {
	// Instance ID, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Image ID, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Image ID, which can be obtained from the `BlueprintId` returned by the [DescribeBlueprints](https://intl.cloud.tencent.com/document/product/1207/47689?from_cn_redirect=1) API.
	BlueprintId *string `json:"BlueprintId,omitnil" name:"BlueprintId"`
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

// Predefined struct for user
type ResetInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstanceResponseParams `json:"Response"`
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

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Login password of the instance(s). The password requirements vary among different operating systems:
	// The password of a `LINUX_UNIX` instance must contain 8–30 characters (above 12 characters preferably) in at least three of the following types and cannot begin with "/": <br><li>Lowercase letters: [a–z]<br><li>Uppercase letters: [A–Z]<br><li>Digits: 0–9<br><li>Special symbols: ()\`~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	// The password of a `WINDOWS` instance must contain 12–30 characters in at least three of the following types and cannot begin with "/" or include the username: <br><li>Lowercase letters: [a–z]<br><li>Uppercase letters: [A–Z]<br><li>Digits: 0–9<br><li>Special symbols: ()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>If both `LINUX_UNIX` and `WINDOWS` instances exist, the requirements for password complexity of `WINDOWS` instances shall prevail.
	Password *string `json:"Password,omitnil" name:"Password"`

	// OS username of the instance for which you want to reset the password, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Login password of the instance(s). The password requirements vary among different operating systems:
	// The password of a `LINUX_UNIX` instance must contain 8–30 characters (above 12 characters preferably) in at least three of the following types and cannot begin with "/": <br><li>Lowercase letters: [a–z]<br><li>Uppercase letters: [A–Z]<br><li>Digits: 0–9<br><li>Special symbols: ()\`~!@#$%^&\*-+=\_|{}[]:;'<>,.?/</li>
	// The password of a `WINDOWS` instance must contain 12–30 characters in at least three of the following types and cannot begin with "/" or include the username: <br><li>Lowercase letters: [a–z]<br><li>Uppercase letters: [A–Z]<br><li>Digits: 0–9<br><li>Special symbols: ()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>If both `LINUX_UNIX` and `WINDOWS` instances exist, the requirements for password complexity of `WINDOWS` instances shall prevail.
	Password *string `json:"Password,omitnil" name:"Password"`

	// OS username of the instance for which you want to reset the password, which can contain up to 64 characters.
	UserName *string `json:"UserName,omitnil" name:"UserName"`
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

// Predefined struct for user
type ResetInstancesPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesPasswordResponseParams `json:"Response"`
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

type Scene struct {
	// Scene ID
	SceneId *string `json:"SceneId,omitnil" name:"SceneId"`

	// Display name of the scene
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// Scene description
	Description *string `json:"Description,omitnil" name:"Description"`
}

type SceneInfo struct {
	// Scene ID
	SceneId *string `json:"SceneId,omitnil" name:"SceneId"`

	// Display name of the scene
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// Scene description
	Description *string `json:"Description,omitnil" name:"Description"`
}

type Snapshot struct {
	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`

	// Type of the disk for which the snapshot is created. Valid values: <li>SYSTEM_DISK: system disk</li>
	DiskUsage *string `json:"DiskUsage,omitnil" name:"DiskUsage"`

	// ID of the disk for which the snapshot is created.
	DiskId *string `json:"DiskId,omitnil" name:"DiskId"`

	// Size of the disk in GB for which the snapshot is created.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Snapshot name, which is a custom snapshot alias.
	SnapshotName *string `json:"SnapshotName,omitnil" name:"SnapshotName"`

	// Snapshot status. Valid values:
	// <li>NORMAL: normal </li>
	// <li>CREATING: creating</li>
	// <li>ROLLBACKING: rolling back</li>
	SnapshotState *string `json:"SnapshotState,omitnil" name:"SnapshotState"`

	// Snapshot creation or rollback progress in percentage. After success, the value of this field will become 100.
	Percent *int64 `json:"Percent,omitnil" name:"Percent"`

	// Last snapshot operation. It is recorded only during snapshot creation and rollback.
	// Example values: CreateInstanceSnapshot, RollbackInstanceSnapshot.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperation *string `json:"LatestOperation,omitnil" name:"LatestOperation"`

	// Last snapshot operation status. It is recorded only during snapshot creation and rollback.
	// Valid values:
	// <li>SUCCESS: operation succeeded</li>
	// <li>OPERATING: the operation is being executed</li>
	// <li>FAILED: operation failed</li>
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationState *string `json:"LatestOperationState,omitnil" name:"LatestOperationState"`

	// Unique request ID for the last snapshot operation. It is recorded only during snapshot creation and rollback.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitnil" name:"LatestOperationRequestId"`

	// Snapshot creation time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type SnapshotDeniedActions struct {
	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitnil" name:"SnapshotId"`

	// List of operation limits.
	DeniedActions []*DeniedAction `json:"DeniedActions,omitnil" name:"DeniedActions"`
}

type Software struct {
	// Software name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Software version.
	Version *string `json:"Version,omitnil" name:"Version"`

	// Software picture URL.
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// Software installation directory.
	InstallDir *string `json:"InstallDir,omitnil" name:"InstallDir"`

	// List of software details.
	DetailSet []*SoftwareDetail `json:"DetailSet,omitnil" name:"DetailSet"`
}

type SoftwareDetail struct {
	// Unique detail key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Detail title.
	Title *string `json:"Title,omitnil" name:"Title"`

	// Detail value.
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type StartInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartInstancesResponseParams `json:"Response"`
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

// Predefined struct for user
type StopInstancesRequestParams struct {
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Each request can contain up to 100 instances at a time. You can get an instance ID from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type StopInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopInstancesResponseParams `json:"Response"`
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
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// System disk size in GB.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// System disk ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DiskId *string `json:"DiskId,omitnil" name:"DiskId"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type TerminateDisksRequestParams struct {
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest
	
	// List of cloud disk IDs.
	DiskIds []*string `json:"DiskIds,omitnil" name:"DiskIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type TerminateInstancesRequestParams struct {
	// Instance ID list, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list, which can be obtained from the `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/1207/47573?from_cn_redirect=1) API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
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

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstancesResponseParams `json:"Response"`
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

type TotalPrice struct {
	// Total original price
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// Total discounted price
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`
}

type TrafficPackage struct {
	// Traffic package ID.
	TrafficPackageId *string `json:"TrafficPackageId,omitnil" name:"TrafficPackageId"`

	// Used traffic in bytes during traffic package validity period.
	TrafficUsed *int64 `json:"TrafficUsed,omitnil" name:"TrafficUsed"`

	// Total traffic in bytes during traffic package validity period.
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitnil" name:"TrafficPackageTotal"`

	// Remaining traffic in bytes during traffic package validity period.
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitnil" name:"TrafficPackageRemaining"`

	// Traffic exceeding package amount in bytes during traffic package validity period.
	TrafficOverflow *int64 `json:"TrafficOverflow,omitnil" name:"TrafficOverflow"`

	// Start time of traffic package validity period according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of traffic package validity period according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Traffic package expiration time according to ISO 8601 standard. UTC time is used. 
	// Format: YYYY-MM-DDThh:mm:ssZ.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deadline *string `json:"Deadline,omitnil" name:"Deadline"`

	// Traffic package status:
	// <li>NETWORK_NORMAL: normal</li>
	// <li>OVERDUE_NETWORK_DISABLED: the network is disconnected due to overdue payments</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ZoneInfo struct {
	// AZ
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Chinese name of AZ
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// AZ tags on instance purchase page
	InstanceDisplayLabel *string `json:"InstanceDisplayLabel,omitnil" name:"InstanceDisplayLabel"`
}