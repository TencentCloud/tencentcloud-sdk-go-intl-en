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

type ActionTimer struct {

	// Additional data
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// Timer name. Currently `TerminateInstances` is the only supported value.
	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`

	// Execution time, which must be at least 5 minutes later than the current time. For example, 2018-5-29 11:26:40.
	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
}

type AllocateHostsRequest struct {
	*tchttp.BaseRequest

	// Instance location. This parameter is used to specify the attributes of an instance, such as its availability zone and project.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// A string used to ensure the idempotency of the request.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Configuration of prepaid instances. You can use the parameter to specify the attributes of prepaid instances, such as the subscription period and the auto-renewal plan. This parameter is required for prepaid instances.
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// The billing method of an instance. Currently only `PREPAID` is supported.
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// CDH instance model. Default value: `HS1`.
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// The quantity of CDH instances you want to purchase.
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// Tag description. You can specify the parameter to associate a tag with an instance.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`
}

func (r *AllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The ID list of the CVM instances newly created on the CDH.
		HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). The maximum number of instances in each request is 100. <br>You can obtain the available instance IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the instance IDs. <br><li>Call [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Key ID(s). The maximum number of key pairs in each request is 100. The key pair ID is in the format of `skey-3glfot13`. <br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before associating a key pair with it. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
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

func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// ID of the security group to be associated, such as `sg-efil73jd`. Only one security group can be associated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// ID(s) of the instance(s) to be associated，such as `ins-lesecurk`. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {

	// 
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Auto renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>Default value: NOTIFY_AND_AUTO_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type CreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest

	// Name of the spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of the spread placement group. Valid values: <br><li>HOST: physical machine <br><li>SW: switch <br><li>RACK: rack
	Type *string `json:"Type,omitempty" name:"Type"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. <br>For more information, see “How to ensure idempotency”.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *CreateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of spread placement group IDs.
		DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

		// Type of the spread placement group. Valid values: <br><li>HOST: physical machine <br><li>SW: switch <br><li>RACK: rack.
		Type *string `json:"Type,omitempty" name:"Type"`

		// Name of the spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
		Name *string `json:"Name,omitempty" name:"Name"`

		// The maximum number of CVMs in a placement group.
		CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`

		// The current number of CVMs in a placement group.
		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

		// Creation time of the placement group.
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageRequest struct {
	*tchttp.BaseRequest

	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// The ID of the instance used to create an image
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Whether to force shut down an instance to create an image when a soft shutdown fails
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`

	// Whether to enable Sysprep when creating a Windows image
	Sysprep *string `json:"Sysprep,omitempty" name:"Sysprep"`

	// The ID of the data disk used to create an image
	DataDiskIds []*string `json:"DataDiskIds,omitempty" name:"DataDiskIds" list`

	// The ID of the snapshot used to create an image. A system disk snapshot must be included.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds" list`

	// Verifies the validity of the request without affecting the resources involved.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Image ID
	// Note: This field may return null, indicating that no valid value was found.
		ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest

	// Name of the key pair, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// The ID of the project to which the new key pair belongs.
	// You can query the project IDs in two ways:
	// <li>Query the project IDs in the project list.
	// <li>Call `DescribeProject` and look for `projectId` in the response.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyPairRequest) FromJsonString(s string) error {
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

func (r *CreateKeyPairResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// Data disk size (in GB). The minimum adjustment increment is 10 GB. The value range varies by data disk type. For more information on limits, see [Storage Overview](https://cloud.tencent.com/document/product/213/4952). The default value is 0, indicating that no data disk is purchased. For more information, see the product documentation.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// The type of the data disk. For more information regarding data disk types and limits, refer to [Storage Overview](https://cloud.tencent.com/document/product/213/4952). Valid values: <br><li>LOCAL_BASIC: local disk<br><li>LOCAL_SSD: local SSD disk<br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: premium cloud storage<br><li>CLOUD_SSD: SSD cloud disk<br><br>Default value: LOCAL_BASIC.<br><br>This parameter is invalid for `ResizeInstanceDisk`.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Data disk ID. Data disks of the type `LOCAL_BASIC` or `LOCAL_SSD` do not have IDs and do not support this parameter.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// Whether to terminate the data disk when its CVM is terminated. Valid values:
	// <li>TRUE: terminate the data disk when its CVM is terminated. This value only supports pay-as-you-go cloud disks billed on an hourly basis.
	// <li>FALSE: retain the data disk when its CVM is terminated.<br>
	// Default value: TRUE<br>
	// Currently this parameter is only used in the `RunInstances` API.
	// Note: This field may return null, indicating that no valid value is found.
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// Data disk snapshot ID. The size of the selected data disk snapshot must be smaller than that of the data disk.
	// Note: This field may return null, indicating that no valid value is found.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Specifies whether the data disk is encrypted. Valid values: 
	// <li>TRUE: encrypted
	// <li>FALSE: not encrypted<br>
	// Default value: FALSE<br>
	// This parameter is only used with `RunInstances`.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`
}

type DeleteDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// List of spread placement group IDs, which can be obtained by calling the [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810) API.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`
}

func (r *DeleteDisasterRecoverGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDisasterRecoverGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDisasterRecoverGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest

	// List of the IDs of the instances to be deleted.
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Key ID(s). The maximum number of key pairs in each request is 100. <br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
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

func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The maximum number of placement groups that can be created.
		GroupQuota *int64 `json:"GroupQuota,omitempty" name:"GroupQuota"`

		// The number of placement groups that have been created by the current user.
		CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

		// Quota on instances in a physical-machine-type disaster recovery group.
		CvmInHostGroupQuota *int64 `json:"CvmInHostGroupQuota,omitempty" name:"CvmInHostGroupQuota"`

		// Quota on instances in a switch-type disaster recovery group.
		CvmInSwGroupQuota *int64 `json:"CvmInSwGroupQuota,omitempty" name:"CvmInSwGroupQuota"`

		// Quota on instances in a rack-type disaster recovery group.
		CvmInRackGroupQuota *int64 `json:"CvmInRackGroupQuota,omitempty" name:"CvmInRackGroupQuota"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest

	// List of spread placement group IDs.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`

	// Name of a spread placement group. Fuzzy match is supported.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDisasterRecoverGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information on spread placement groups.
		DisasterRecoverGroupSet []*DisasterRecoverGroup `json:"DisasterRecoverGroupSet,omitempty" name:"DisasterRecoverGroupSet" list`

		// Total number of placement groups of the user.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDisasterRecoverGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest

	// Filters.
	// <li> `zone` - String - Optional - Filter results by availability zone.</li>
	// <li> `project-id` - Integer - Optional - Filter results by project ID. You can call `DescribeProject` or log in to the console to view the list of existing projects. You can also create a new project by calling `AddProject`.</li>
	// <li> `host-id` - String - Optional - Filter results by CDH ID. CDH IDs are in the format of `host-11112222`.</li>
	// <li> `host-name` - String - Optional - Filter results by CDH instance name.</li>
	// <li> `host-state` - String - Optional - Filter results by CDH instance state. (PENDING: creating | LAUNCH_FAILURE: creation failed | RUNNING: running | EXPIRED: expired)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset; default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of CDH instances meeting the query conditions
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Information on CDH instances
		HostSet []*HostItem `json:"HostSet,omitempty" name:"HostSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The image quota of an account
		ImageNumQuota *int64 `json:"ImageNumQuota,omitempty" name:"ImageNumQuota"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionRequest struct {
	*tchttp.BaseRequest

	// The ID of the image to be shared
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information on image sharing.
		SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// List of image IDs, such as `img-gvbnzy6f`. For the format of array-type parameters, see [API Introduction](https://cloud.tencent.com/document/api/213/15688). You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response. <br><li>View the image IDs in the [Image Console](https://console.cloud.tencent.com/cvm/image).
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`

	// Filters. Each request can have up to 10 `Filters` and 5 `Filters.Values`. You cannot specify `ImageIds` and `Filters` at the same time. Specific filters:
	// <li>`image-id` - String - Optional - Filter results by image ID</li>
	// <li>`image-type` - String - Optional - Filter results by image type. Valid values:
	//     PRIVATE_IMAGE: private image created by the current account 
	//     PUBLIC_IMAGE: public image created by Tencent Cloud
	//    SHARED_IMAGE: image shared with the current account by another account.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset; default value: 0. For more information on `Offset`, see [API Introduction](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see [API Introduction](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Instance type, e.g. `S1.SMALL1`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information on an image, including its state and attributes.
		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet" list`

		// Number of images meeting the filtering conditions.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Supported operating system types of imported images.
		ImportImageOsListSupported *ImageOsList `json:"ImportImageOsListSupported,omitempty" name:"ImportImageOsListSupported"`

		// Supported operating system versions of imported images. 
		ImportImageOsVersionSet []*OsVersion `json:"ImportImageOsVersionSet,omitempty" name:"ImportImageOsVersionSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance model families
		InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest

	// Filters.
	// <li> `zone` - String - Optional - Filter results by [availability zone](https://cloud.tencent.com/document/product/213/15753#ZoneInfo).</li>
	// <li> `instance-family` - String - Optional - Filter results by instance model family, such as `S1`, `I1`, and `M1`.</li>
	// Each request can have up to 10 `Filters` and 1 `Filters.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of instance model families
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest

	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVncUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
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

func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesOperationLimitRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The instance ID format is `ins-xxxxxxxx`. For more information on the format of this parameter, see the `id.N` section of [API Introduction](https://cloud.tencent.com/document/api/213/15688)). The maximum number of instance IDs in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Operation on the instance(s).
	// <li> INSTANCE_DEGRADE: downgrade the instance configurations</li>
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeInstancesOperationLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesOperationLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesOperationLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The maximum number of times you can modify the instance configurations (degrading the configurations)
		InstanceOperationLimitSet []*OperationCountLimit `json:"InstanceOperationLimitSet,omitempty" name:"InstanceOperationLimitSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesOperationLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesOperationLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s) in the format of `ins-xxxxxxxx`. For more information on the format of this parameter, see the `id.N` section of [API Introduction](https://cloud.tencent.com/document/api/213/15688). The maximum number of instances in each request is 100. You cannot specify `InstanceIds` and `Filters` at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Filters.
	// <li> `zone` - String - Optional - Filter results by availability zone.</li>
	// <li> `project-id` - Integer - Optional - Filter results by project ID. You can call [DescribeProject](https://cloud.tencent.com/document/api/378/4400) or log in to the [console](https://console.cloud.tencent.com/cvm/index) to view the list of existing projects. You can also create a new project by calling [AddProject](https://cloud.tencent.com/document/api/378/4398).</li>
	// <li> `host-id` - String - Optional - Filter results by [CDH](https://cloud.tencent.com/document/product/416) ID. [CDH](https://cloud.tencent.com/document/product/416) ID format: `host-xxxxxxxx`.</li>
	// </li>`vpc-id` - String - Optional - Filter results by VPC ID. VPC ID format: `vpc-xxxxxxxx`.</li>
	// <li> `subnet-id` - String - Optional - Filter results by subnet ID. Subnet ID format: `subnet-xxxxxxxx`.</li>
	// </li>`instance-id` - String - Optional - Filter results by instance ID. Instance ID format: `ins-xxxxxxxx`.</li>
	// </li>`security-group-id` - String - Optional - Filter results by security group ID. Security group ID format: `sg-8jlk3f3r`.</li>
	// </li>`instance-name` - String - Optional - Filter results by instance name.</li>
	// </li>`instance-charge-type` - String - Optional - Filter results by instance billing method. `POSTPAID_BY_HOUR`: pay-as-you-go | `CDHPAID`: you are only billed for [CDH](https://cloud.tencent.com/document/product/416) instances, not the CVMs running on the [CDH](https://cloud.tencent.com/document/product/416) instances.</li>
	// </li>`private-ip-address` - String - Optional - Filter results by the private IP address of the instance’s primary ENI.</li>
	// </li>`public-ip-address` - String - Optional - Filter results by the public IP address of the instance’s primary ENI, including the IP addresses automatically assigned during the instance creation and the EIPs manually associated after the instance creation.</li>
	// <li> `tag-key` - String - Optional - Filter results by tag key.</li>
	// </li>`tag-value` - String - Optional - Filter results by tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Filter results by tag key-value pair. Replace `tag-key` with specific tag keys, as shown in example 2.</li>
	// Each request can have up to 10 `Filters` and 5 `Filters.Values`. You cannot specify `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instances meeting the filtering conditions.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Detailed instance information.
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s) in the format of `ins-xxxxxxxx`. For more information on the format of this parameter, see the `id.N` section of [API Introduction](https://cloud.tencent.com/document/api/213/15688). The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of instance states meeting the filtering conditions.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of [instance states](https://cloud.tencent.com/document/api/213/15728).
		InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet,omitempty" name:"InstanceStatusSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInternetChargeTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternetChargeTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of network billing methods.
		InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet,omitempty" name:"InternetChargeTypeConfigSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternetChargeTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternetChargeTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Key pair ID(s) in the format of `skey-11112222`. This API supports using multiple IDs as filters at the same time. For more information on the format of this parameter, see the `id.N` section in [API Introduction](https://cloud.tencent.com/document/api/213/15688). You cannot specify `KeyIds` and `Filters` at the same time. You can log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the key pair IDs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// Filters.
	// <li> `project-id` - Integer - Optional - Filter results by project ID. To view the list of project IDs, you can go to [Project Management](https://console.cloud.tencent.com/project), or call [DescribeProject](https://cloud.tencent.com/document/api/378/4400) and look for `projectId` in the response. </li>
	// <li> `key-name` - String - Optional - Filter results by key pair name. </li> You cannot specify `KeyIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding sections in API [Introduction](https://intl.cloud.tencent.com/document/product/377). Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of key pairs meeting the filtering conditions.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Detailed information on key pairs.
		KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
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

		// Number of regions
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of regions
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

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// Filters.
	// 
	// <li> `zone` - String - Optional - Filter results by availability zone.</li>
	// 
	// <li>`instance-family` - String - Optional - Filter results by instance model family, such as `S1`, `I1`, and `M1`.</li>
	// 
	// <li>`instance-type` - String - Optional - Filter results by model. Different instance models have different configurations. You can call `DescribeInstanceTypeConfigs` to query the latest configuration list or refer to the documentation on instance types. If this parameter is not specified, `S1.SMALL1` will be used by default.</li>
	// 
	// <li>`instance-charge-type` - String - Optional - Filter results by instance billing method. `POSTPAID_BY_HOUR`: pay-as-you-go | `CDHPAID`: you are only billed for CDH instances, not the CVMs running on the CDH instances.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of model configurations for the availability zone.
		InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
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

		// Number of availability zones.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of availability zones.
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

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). The maximum number of instances in each request is 100. <br><br>You can obtain the available instance IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the instance IDs. <br><li>Call [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// List of key pair IDs. The maximum number of key pairs in each request is 100. The key pair ID is in the format of `skey-11112222`. <br><br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before disassociating a key pair from it. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
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

func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// ID of the security group to be disassociated, such as `sg-efil73jd`. Only one security group can be disassociated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// ID(s) of the instance(s) to be disassociated，such as `ins-lesecurk`. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisasterRecoverGroup struct {

	// ID of a spread placement group.
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// Name of a spread placement group, which must be 1-60 characters long.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of a spread placement group. Valid values: <br><li>HOST: physical machine <br><li>SW: switch <br><li>RACK: rack.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The maximum number of CVMs that can be hosted in a spread placement group.
	CvmQuotaTotal *int64 `json:"CvmQuotaTotal,omitempty" name:"CvmQuotaTotal"`

	// The current number of CVMs in a spread placement group.
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// The list of CVM IDs in a spread placement group.
	// Note: This field may return null, indicating that no valid value was found.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Creation time of a spread placement group.
	// Note: This field may return null, indicating that no valid value is found.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type EnhancedService struct {

	// Enables cloud security service. If this parameter is not specified, the cloud security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Enables cloud monitor service. If this parameter is not specified, the cloud monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type Externals struct {

	// Release address
	// Note: This field may return null, indicating that no valid value is found.
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// Unsupported network types
	// Note: This field may return null, indicating that no valid value is found.
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty" name:"UnsupportNetworks" list`

	// Attributes of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty" name:"StorageBlockAttr"`
}

type Filter struct {

	// Filters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values.
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type HostItem struct {

	// Location of the CDH instance. You can use this parameter to specify the attributes of the instance, such as its availability zone and project.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// CDH instance ID
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// CDH instance type
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// CDH instance name
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Billing method of the CDH instance
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// Auto renewal flag of the CDH instance
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Creation time of the CDH instance
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Expiration time of the CDH instance
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// List of IDs of CVM instances created on the CDH
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// CDH instance state
	HostState *string `json:"HostState,omitempty" name:"HostState"`

	// CDH instance IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// Resource information of the CDH instance
	HostResource *HostResource `json:"HostResource,omitempty" name:"HostResource"`

	// Cage ID of the CDH instance. This parameter is only valid for CDH instances in the cages of finance availability zones.
	// Note: This field may return null, indicating that no valid value is found.
	CageId *string `json:"CageId,omitempty" name:"CageId"`
}

type HostResource struct {

	// Total number of CPU cores in the CDH instance
	CpuTotal *uint64 `json:"CpuTotal,omitempty" name:"CpuTotal"`

	// Number of available CPU cores in the CDH instance
	CpuAvailable *uint64 `json:"CpuAvailable,omitempty" name:"CpuAvailable"`

	// Total memory of the CDH instance; unit: GiB
	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// Available memory of the CDH instance; unit: GiB
	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`

	// Total disk size of the CDH instance; unit: GiB
	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`

	// Avilable disk size of the CDH instance; unit: GiB
	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`
}

type Image struct {

	// Image ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Operating system of the image
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Image type
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// Creation time of the image
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Image size
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// Image architecture
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// Image state
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// Source platform of the image
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// Image creator
	ImageCreator *string `json:"ImageCreator,omitempty" name:"ImageCreator"`

	// Image source
	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`

	// Synchronization percentage
	// Note: This field may return null, indicating that no valid value is found.
	SyncPercent *int64 `json:"SyncPercent,omitempty" name:"SyncPercent"`

	// Whether the image supports cloud-init
	// Note: This field may return null, indicating that no valid value is found.
	IsSupportCloudinit *bool `json:"IsSupportCloudinit,omitempty" name:"IsSupportCloudinit"`

	// Information on the snapshots associated with the image
	// Note: This field may return null, indicating that no valid value is found.
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet" list`
}

type ImageOsList struct {

	// Supported Windows OS
	// Note: This field may return null, indicating that no valid value is found.
	Windows []*string `json:"Windows,omitempty" name:"Windows" list`

	// Supported Linux OS
	// Note: This field may return null, indicating that no valid value is found.
	Linux []*string `json:"Linux,omitempty" name:"Linux" list`
}

type ImportImageRequest struct {
	*tchttp.BaseRequest

	// OS architecture of the image to be imported, `x86_64` or `i386`.
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// OS type of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating systems.
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// OS version of the image to be imported. You can call `DescribeImportImageOs` to obtain the list of supported operating systems.
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// Address on COS where the image to be imported is stored.
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Dry run to check the parameters without performing the operation
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Whether to force import the image. For more information, see [Forcibly Import Image](https://intl.cloud.tencent.com/document/product/213/12849).
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest

	// Key pair name, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// The ID of the [project](https://cloud.tencent.com/document/product/378/10861) to which the created key pair belongs.<br><br>You can retrieve the project ID in two ways:<br><li>Query the project ID in [Project Management](https://console.cloud.tencent.com/project).<br><li>Call [DescribeProject](https://cloud.tencent.com/document/api/378/4400) and search for `projectId` in the response.
	// 
	// If you want to use the default project, specify 0 for the parameter.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Content of the public key in the key pair in the `OpenSSH RSA` format.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyPairRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Key pair ID
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyPairResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// [Image](/document/product/213/4940) ID in the format of `img-xxx`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For IDs of `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the information; for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Configuration of the system disk of the instance. For instances with a cloud disk as the system disk, you can expand the system disk by using this parameter to specify the new capacity after reinstallation. If the parameter is not specified, the system disk capacity remains unchanged by default. You can only expand the capacity of the system disk; reducing its capacity is not supported. When reinstalling the system, you can only modify the capacity of the system disk, not the type.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Monitor and Cloud Security. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *InquiryPriceResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price of reinstalling the instance with the specified configuration.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100. When changing the bandwidth of instances with `BANDWIDTH_PREPAID` or `BANDWIDTH_POSTPAID_BY_HOUR` as the network billing method, you can only specify one instance at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Configuration of public network egress bandwidth. The maximum bandwidth varies among different models. For more information, see the documentation on bandwidth limits. Currently only the `InternetMaxBandwidthOut` parameter is supported.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Date from which the new bandwidth takes effect. Format: `YYYY-MM-DD`, such as `2016-10-30`. The starting date cannot be earlier than the current date. If the starting date is the current date, the new bandwidth takes effect immediately. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Date until which the new bandwidth is effective. Format: `YYYY-MM-DD`, such as `2016-10-30`. The validity period of the new bandwidth covers the end date. The end date cannot be later than the expiration date of a prepaid instance. You can query the expiration time of an instance by calling [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728) and looking for `ExpiredTime` in the response. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price of the new bandwidth
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 1.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Instance model. Different resource specifications are specified for different models. For specific values, see the instance resource specification table. You can also obtain the latest specification list by calling the API for querying the list of instance resource specifications.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *InquiryPriceResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price of the instance using the specified model
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The configuration of data disks to be expanded. Currently, you can only use the API to expand non-elastic data disks whose [disk type](https://cloud.tencent.com/document/product/213/15753#DataDisk) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://cloud.tencent.com/document/api/362/16315) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic. Data disk capacity unit: GB; minimum increment: 10 GB. For more information about selecting a data disk type, see the product overview on cloud disks. Available data disk types are subject to the instance type (`InstanceType`). In addition, the maximum capacity allowed for expansion varies by data disk type.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *InquiryPriceResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price of the disks after being expanded to the specified configurations
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesRequest struct {
	*tchttp.BaseRequest

	// Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone and project.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// [Image](/document/product/213/4940) ID in the format of `img-xxx`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For IDs of `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the information; for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// The instance [billing method](https://cloud.tencent.com/document/product/213/2180).<br><li>POSTPAID_BY_HOUR: hourly, pay-as-you-go<br>Default value: POSTPAID_BY_HOUR.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Configuration of prepaid instances. You can use the parameter to specify the attributes of prepaid instances, such as the subscription period and the auto-renewal plan. This parameter is required for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// The instance model. Different resource specifications are specified for different models. For specific values, call [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) to retrieve the latest specification list or refer to [Instance Types](https://cloud.tencent.com/document/product/213/11518). If the parameter is not specified, `S1.SMALL1` will be used by default.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// The configuration information of the instance data disk. If this parameter is not specified, no data disk will be purchased by default. When purchasing, you can specify 21 data disks, which can contain at most 1 LOCAL_BASIC data disk or LOCAL_SSD data disk, and at most 20 CLOUD_BASIC data disks, CLOUD_PREMIUM data disks, or CLOUD_SSD data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// VPC configurations. You can use this parameter to specify the VPC ID, subnet ID, etc. If this parameter is not specified, the basic network will be used by default. If a VPC IP is specified in this parameter, the `InstanceCount` parameter can only be 1. 
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Number of instances to be purchased. Value range: [1, 100]; default value: 1. The specified number of instances to be purchased cannot exceed the remaining quota allowed for the user. For more information on quota, see [CVM instance purchase limit](https://intl.cloud.tencent.com/document/product/213/2664).
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instance name to be displayed. <br><li>If this parameter is not specified, "Unnamed" will be displayed by default. </li><li>If you purchase multiple instances at the same time and specify a pattern string `{R:x}`, numbers `[x, x+n-1]` will be generated, where `n` represents the number of instances purchased. For example, you specify a pattern string, `server_{R:3}`. If you only purchase 1 instance, the instance will be named `server_3`; if you purchase 2, they will be named `server_3` and `server_4`. You can specify multiple pattern strings in the format of `{R:x}`. </li><li>If you purchase multiple instances at the same time and do not specify a pattern string, the instance names will be suffixed by `1, 2...n`, where `n` represents the number of instances purchased. For example, if you purchase 2 instances and the instance name body is `server_`, the instance names will be `server_1` and `server_2`.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response. If this parameter is not specified, the instance will not be associated with any security group by default.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Monitor and Cloud Security. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. <br>For more information, see “How to ensure idempotency”.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Host name of the CVM. <br><li>Periods (.) or hyphens (-) cannot be the start or end of a host name or appear consecutively in a host name.<br><li>For Windows instances, the host name must be 2-15 characters long and can contain uppercase and lowercase letters, numbers, and hyphens (-). It cannot contain periods (.) or contain only numbers. <br><li>For other instances, such as Linux instances, the host name must be 2-30 characters long. It supports multiple periods (.) and allows uppercase and lowercase letters, numbers, and hyphens (-) between any two periods (.).
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// The tag description list. This parameter is used to bind a tag to a resource instance. A tag can only be bound to CVM instances.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// The market options of the instance.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`
}

func (r *InquiryPriceRunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Price of the instance with the specified configurations.
		Price *Price `json:"Price,omitempty" name:"Price"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// Location of the instance
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance `ID`
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Instance model
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of CPU cores of the instance; unit: core
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Memory capacity; unit: `GB`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance status. Valid values: <br><li>NORMAL: instance is normal. <br><li>EXPIRED: instance expired. <br><li>PROTECTIVELY_ISOLATED: instance is protectively isolated.
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance billing plan. Valid values:<br><li>`POSTPAID_BY_HOUR`: pay after use. You are billed by the hour, by traffic.<br><li>`CDHPAID`: `CDH` billing plan. Applicable to `CDH` only, not the instances on the host.<br>
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Information on the system disk of the instance
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Information on the data disks of the instance, which only covers the data disks purchased together with the instance. 
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// List of private IPs of the instance's primary ENI.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// List of public IPs of the instance's primary ENI.
	// Note: This field may return null, indicating that no valid value is found.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses" list`

	// Information on instance bandwidth.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Information on the VPC where the instance resides.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// `ID` of the image used to create the instance.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Auto renewal flag. Valid values: <br><li>`NOTIFY_AND_MANUAL_RENEW`: notify upon expiration, but do not renew automatically <br><li>`NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically <br><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`: do not notify upon expiration and do not renew automatically.
	// <br><li>Note: this parameter is `null` for postpaid instances.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Creation time following the `ISO8601` standard and using `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Expiration time in UTC format following the `ISO8601` standard: `YYYY-MM-DDThh:mm:ssZ`. Note: this parameter is `null` for postpaid instances.
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Operating system name.
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Login settings of the instance. Currently only the key associated with the instance is returned.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Instance state. Valid values: <br><li>PENDING: creating <br></li><li>LAUNCH_FAILED: creation failed <br></li><li>RUNNING: running <br></li><li>STOPPED: shut down <br></li><li>STARTING: starting <br></li><li>STOPPING: shutting down <br></li><li>REBOOTING: rebooting <br></li><li>SHUTDOWN: shut down and to be terminated <br></li><li>TERMINATING: terminating. <br></li>
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// List of tags associated with the instance.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// Instance billing method after shutdown.
	// Valid values: <br><li>KEEP_CHARGING: billing continues after shutdown <br><li>STOP_CHARGING: billing stops after shutdown <li>NOT_APPLICABLE: the instance is not shut down or stopping billing after shutdown is not applicable to the instance. <br>
	StopChargingMode *string `json:"StopChargingMode,omitempty" name:"StopChargingMode"`

	// 
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// ID of a spread placement group.
	// Note: this field may return null, indicating that no valid value is obtained.
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// IPv6 address of the instance.
	// Note: this field may return null, indicating that no valid value is obtained.
	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses" list`

	// CAM role name.
	// Note: this field may return null, indicating that no valid value is obtained.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`
}

type InstanceChargePrepaid struct {

	// Subscription period; unit: month; valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Auto renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceFamilyConfig struct {

	// Full name of the model family.
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// Acronym of the model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceMarketOptionsRequest struct {

	// Options related to bidding
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitempty" name:"SpotOptions"`

	// Market option type. Currently `spot` is the only supported value.
	MarketType *string `json:"MarketType,omitempty" name:"MarketType"`
}

type InstanceStatus struct {

	// Instance `ID`.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The instance state. Valid values: <br><li>PENDING: creating<br></li><li>LAUNCH_FAILED: creation failed<br></li><li>RUNNING: running<br></li><li>STOPPED: shut down<br></li><li>STARTING: starting<br></li><li>STOPPING: shutting down<br></li><li>REBOOTING: rebooting<br></li><li>SHUTDOWN: shut down and to be terminated<br></li><li>TERMINATING: terminating.<br></li>
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
}

type InstanceTypeConfig struct {

	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Number of GPU cores.
	GPU *int64 `json:"GPU,omitempty" name:"GPU"`

	// Number of CPU cores.
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Memory capacity; unit: `GB`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Number of FPGA cores; unit: core.
	FPGA *int64 `json:"FPGA,omitempty" name:"FPGA"`
}

type InstanceTypeQuotaItem struct {

	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance billing plan. Valid values: <br><li>POSTPAID_BY_HOUR: pay after use. You are billed for your traffic by the hour.<br><li>`CDHPAID`: [`CDH`](https://cloud.tencent.com/document/product/416) billing plan. Applicable to `CDH` only, not the instances on the host.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// ENI type. For example, 25 represents an ENI of 25 GB.
	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// Additional data.
	// Note: This field may return null, indicating that no valid value is found.
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// Number of CPU cores of an instance model.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance memory capacity; unit: `GB`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Model name.
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// List of local disk specifications. If the parameter returns null, it means that local disks cannot be created.
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList" list`

	// Whether an instance model is available. Valid values: <br><li>SELL: available <br><li>SOLD_OUT: sold out
	Status *string `json:"Status,omitempty" name:"Status"`

	// Price of an instance model.
	Price *ItemPrice `json:"Price,omitempty" name:"Price"`

	// Details of out-of-stock items
	// Note: this field may return null, indicating that no valid value is obtained.
	SoldOutReason *string `json:"SoldOutReason,omitempty" name:"SoldOutReason"`
}

type InternetAccessible struct {

	// Network connection billing plan. Valid value: <br><li>TRAFFIC_POSTPAID_BY_HOUR: pay after use. You are billed for your traffic, by the hour.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// The maximum outbound bandwidth of the public network, in Mbps. The default value is 0 Mbps. The upper limit of bandwidth varies for different models. For more information, see [Purchase Network Bandwidth](https://cloud.tencent.com/document/product/213/12523).
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP. Valid values: <br><li>TRUE: Assign a public IP <br><li>FALSE: Do not assign a public IP <br><br>If the public network bandwidth is greater than 0 Mbps, you can choose whether to assign a public IP; by default a public IP will be assigned. If the public network bandwidth is 0 Mbps, you will not be able to assign a public IP.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`

	// Bandwidth package ID. To obatin the IDs, you can call [`DescribeBandwidthPackages`](https://cloud.tencent.com/document/api/215/19209) and look for the `BandwidthPackageId` fields in the response.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

type InternetChargeTypeConfig struct {

	// Network billing method.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// Description of the network billing method.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ItemPrice struct {

	// The original unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to the unit price for the usage between 0 and ∞ hours.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Billing unit for pay-as-you-go mode. Valid values: <br><li>HOUR: billed on an hourly basis. It's used for hourly postpaid instances (`POSTPAID_BY_HOUR`). <br><li>GB: bill by traffic in GB. It's used for postpaid products that are billed by the hourly traffic (`TRAFFIC_POSTPAID_BY_HOUR`).
	// Note: this field may return null, indicating that no valid value is obtained.
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// The original price of a pay-in-advance instance, in USD.
	// Note: this field may return null, indicating that no valid value is obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discount price of a prepaid instance, in USD.
	// Note: this field may return null, indicating that no valid value is obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Percentage of the original price. For example, if you enter "20", the discounted price will be 20% of the original price.
	// Note: this field may return null, indicating that no valid value is obtained.
	Discount *uint64 `json:"Discount,omitempty" name:"Discount"`

	// The discounted unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to the unit price for the usage between 0 and ∞ hours.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// Original unit price for the usage between 96 to 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceSecondStep *float64 `json:"UnitPriceSecondStep,omitempty" name:"UnitPriceSecondStep"`

	// Discounted unit price for the usage between 96 to 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscountSecondStep *float64 `json:"UnitPriceDiscountSecondStep,omitempty" name:"UnitPriceDiscountSecondStep"`

	// Original unit price for the usage after 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceThirdStep *float64 `json:"UnitPriceThirdStep,omitempty" name:"UnitPriceThirdStep"`

	// Discounted unit price for the usage after 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscountThirdStep *float64 `json:"UnitPriceDiscountThirdStep,omitempty" name:"UnitPriceDiscountThirdStep"`
}

type KeyPair struct {

	// Key pair `ID`, the unique identifier of a key pair.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Key pair name.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// `ID` of the project to which a key pair belongs.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Key pair description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Content of public key in a key pair.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Content of private key in a key pair. Tencent Cloud do not keep private keys. Please keep it properly.
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// `ID` list of instances associated with a key.
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds" list`

	// Creation time, which follows the `ISO8601` standard and uses `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type LocalDiskType struct {

	// Type of a local disk.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Attributes of a local disk.
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// Minimum size of a local disk.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum size of a local disk.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 
	Required *string `json:"Required,omitempty" name:"Required"`
}

type LoginSettings struct {

	// Login password of the instance. The password requirements vary among different operating systems: <br><li>For Linux instances, the password must be 8-30 characters long and contain at least two of the following types: [a-z], [A-Z], [0-9] and [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]. <br><li>For Windows instances, the password must be 12-30 characters long and contain at least three of the following categories: [a-z], [A-Z], [0-9] and [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]. <br><br>If this parameter is not specified, a random password will be generated and sent to you via the Message Center.
	// Note: this field may return null, indicating that no valid value is obtained.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call [`DescribeKeyPairs`](https://cloud.tencent.com/document/api/213/15699) to obtain `KeyId`. A key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	// Note: this field may return null, indicating that no valid value is obtained.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// Whether to keep the original settings of an image. You cannot specify this parameter and `Password` or `KeyIds.N` at the same time. You can specify this parameter as `TRUE` only when you create an instance using a custom image, a shared image, or an imported image. Valid values: <br><li>TRUE: keep the login settings of the image <br><li>FALSE: do not keep the login settings of the image <br><br>Default value: FALSE.
	// Note: This field may return null, indicating that no valid value is found.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ModifyDisasterRecoverGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// Spread placement group ID, which can be obtained by calling the [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/api/213/17810) API.
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// Name of a spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDisasterRecoverGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDisasterRecoverGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDisasterRecoverGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDisasterRecoverGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDisasterRecoverGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeRequest struct {
	*tchttp.BaseRequest

	// CDH instance ID(s).
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`

	// CDH instance name to be displayed. You can specify any name you like, but its length cannot exceed 60 characters.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Auto renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

func (r *ModifyHostsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest

	// Image ID such as `img-gvbnzy6f`. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image).
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// New image name, which must meet the following requirements: <br> <li>No more than 20 characters. <br> <li>Must be unique.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// New image description, which must meet the following requirement: <br> <li> No more than 60 characters.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest

	// Image ID such as `img-gvbnzy6f`. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image). <br>You can only specify an image in the `NORMAL` state. For more information on image states, see [here](/document/api/213/9452#image_state).
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// List of account IDs with which an image is shared. For the format of array-type parameters, see [API Introduction](/document/api/213/568). The account ID is different from the QQ number. You can find the account ID in [Account Information](https://console.cloud.tencent.com/developer). 
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds" list`

	// Operations. Valid values: `SHARE`, sharing an image; `CANCEL`, cancelling an image sharing. 
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Instance name. You can specify any name you like, but its length cannot exceed 60 characters.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// ID list of security groups of the instance. The instance will be associated with the specified security groups and will be disassociated from the original security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups" list`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
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

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Project ID. You can create a project by calling [AddProject](https://cloud.tencent.com/doc/api/403/4398). When calling [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) to query instances, the project IDs can be used to filter the results.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesVpcAttributeRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// VPC configurations. You can use this parameter to specify the VPC ID, subnet ID, VPC IP, etc. If the specified VPC ID and subnet ID (the subnet must be in the same availability zone as the instance) are different from the VPC where the specified instance resides, the instance will be migrated to a subnet of the specified VPC. You can use `PrivateIpAddresses` to specify the VPC subnet IP. If you want to specify the subnet IP, you will need to specify a subnet IP for each of the specified instances, and each `InstanceIds` will match a `PrivateIpAddresses`. If `PrivateIpAddresses` is not specified, the VPC subnet IP will be assigned randomly.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Whether to force shut down a running instances. Default value: TRUE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Whether to keep the host name. Default value: FALSE.
	ReserveHostName *bool `json:"ReserveHostName,omitempty" name:"ReserveHostName"`
}

func (r *ModifyInstancesVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesVpcAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesVpcAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeRequest struct {
	*tchttp.BaseRequest

	// Key pair ID in the format of `skey-xxxxxxxx`. <br><br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) and look for `KeyId` in the response.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// New key pair name, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// New key pair description. You can specify any name you like, but its length cannot exceed 60 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyKeyPairAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyKeyPairAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperationCountLimit struct {

	// Instance operation. Valid values: <br><li>`INSTANCE_DEGRADE`: downgrade an instance<br><li>`INTERNET_CHARGE_TYPE_CHANGE`: modify the billing plan of the network connection
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Number of operations already performed. If it returns `-1`, it means there is no limit on the times of the operation.
	CurrentCount *int64 `json:"CurrentCount,omitempty" name:"CurrentCount"`

	// Maximum number of times you can perform an operation. If it returns `-1`, it means there is no limit on the times of the operation. If it returns `0`, it means that configuration modification is not supported.
	LimitCount *int64 `json:"LimitCount,omitempty" name:"LimitCount"`
}

type OsVersion struct {

	// Operating system type
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// Supported operating system versions
	OsVersions []*string `json:"OsVersions,omitempty" name:"OsVersions" list`

	// Supported operating system architecture
	Architecture []*string `json:"Architecture,omitempty" name:"Architecture" list`
}

type Placement struct {

	// The ID of [availability zone](https://cloud.tencent.com/document/product/213/15753#ZoneInfo) where the instance locates. It can obtained in the `Zone` field returned by [DescribeZones](https://cloud.tencent.com/document/213/15707) API.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ID of the project to which the instance belongs. To obtain the project IDs, you can call [DescribeProject](/document/api/378/4400) and look for the `projectId` fields in the response. If this parameter is not specified, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID list of CDHs from which the instance can be created. If you have purchased CDHs and specify this parameter, the instances you purchase will be randomly deployed on the CDHs.
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds" list`

	// Master host IP used to create the CVM
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps" list`

	// 
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type Price struct {

	// Instance price.
	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// Network price.
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Whether to force restart an instance after a normal restart fails. Valid values: <br><li>TRUE: force restart an instance after a normal restart fails <br><li>FALSE: do not force restart an instance after a normal restart fails <br><br>Default value: FALSE.
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// Shutdown type. Valid values: <br><li>SOFT: soft shutdown<br><li>HARD: hard shutdown<br><li>SOFT_FIRST: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails<br><br>Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
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

func (r *RebootInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// Region name, such as `ap-guangzhou`
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region description, such as South China (Guangzhou)
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Whether the region is available
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest

	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// [Image](/document/product/213/4940) ID in the format of `img-xxx`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For IDs of `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the information; for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response.</li>
	// <br>The current image will be used by default.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Configuration of the system disk of the instance. For instances with a cloud disk as the system disk, you can expand the system disk by using this parameter to specify the new capacity after reinstallation. If the parameter is not specified, the system disk capacity remains unchanged by default. You can only expand the capacity of the system disk; reducing its capacity is not supported. When reinstalling the system, you can only modify the capacity of the system disk, not the type.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Monitor and Cloud Security. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// You can use this parameter to specify a new HostName for the instance when reinstalling the system.
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstanceRequest) FromJsonString(s string) error {
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

func (r *ResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100. When changing the bandwidth of instances with `BANDWIDTH_PREPAID` or `BANDWIDTH_POSTPAID_BY_HOUR` as the network billing method, you can only specify one instance at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Configuration of public network egress bandwidth. The maximum bandwidth varies among different models. For more information, see the documentation on bandwidth limits. Currently only the `InternetMaxBandwidthOut` parameter is supported.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Date from which the new bandwidth takes effect. Format: `YYYY-MM-DD`, such as `2016-10-30`. The starting date cannot be earlier than the current date. If the starting date is the current date, the new bandwidth takes effect immediately. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Date until which the new bandwidth is effective. Format: `YYYY-MM-DD`, such as `2016-10-30`. The validity period of the new bandwidth covers the end date. The end date cannot be later than the expiration date of a prepaid instance. You can query the expiration time of an instance by calling [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728) and looking for `ExpiredTime` in the response. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Login password of the instance(s). The password requirements vary among different operating systems:
	// For a Linux instance, the password must be 8 to 30 characters in length; password with more than 12 characters is recommended. It cannot begin with "/", and must contain at least one character from three of the following categories: <br><li>Lowercase letters: [a-z]<br><li>Uppercase letters: [A-Z]<br><li>Numbers: 0-9<br><li>Special characters: ()\`~!@#$%^&\*-+=\_|{}[]:;'<>,.?/:
	// For a Windows CVM, the password must be 12 to 30 characters in length. It cannot begin with "/" or contain your username. It must contain at least one character from three of the following categories: <br><li>Lowercase letters: [a-z]<br><li>Uppercase letters: [A-Z]<br><li>Numbers: 0-9<br><li>Special characters: ()\`~!@#$%^&\*-+=\_|{}[]:;' <>,.?/:<br><li>If the specified instances include both `Linux` and `Windows` instances, you will need to follow the password requirements for `Windows` instances.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Operating system username of the instance for which you want to reset the password. The length of the username cannot exceed 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
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

func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 1.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Instance model. Different resource specifications are specified for different models. For specific values, call [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) to get the latest specification list or refer to [Instance Types](https://cloud.tencent.com/document/product/213/11518).
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest

	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Configuration of data disks to be expanded. Currently you can only use the API to expand non-elastic data disks whose [disk type](/document/api/213/9452#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://cloud.tencent.com/document/api/362/16315) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is not elastic. Data disk capacity unit: GB; minimum increment: 10 GB. For more information on selecting the data disk type, see the [product overview on cloud disks](https://cloud.tencent.com/document/product/362/2353). Available data disk types are subject to the instance type (`InstanceType`). In addition, the maximum capacity allowed for expansion varies by data disk type.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone, project, and CDH. You can specify a CDH for a CVM by creating the CVM on the CDH.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The [image](/document/product/213/4940) ID in the format of `img-xxx`. There are four types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><li>Marketplace images</li><br/>You can retrieve available image IDs in the following ways:<br/><li>For the IDs of `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the information. For the IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715), pass in `InstanceType` to retrieve the list of images supported by the current model, and then find the `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// The instance [billing method](https://cloud.tencent.com/document/product/213/2180). Valid values: <br><li>`POSTPAID_BY_HOUR`: hourly, pay-as-you-go<br><li>`CDHPAID`: you are only billed for CDH instances, not the CVMs running on the CDH instances.<br>Default value: POSTPAID_BY_HOUR.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Configuration of prepaid instances. You can use the parameter to specify the attributes of prepaid instances, such as the subscription period and the auto-renewal plan. This parameter is required for prepaid instances.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// The instance model. Different resource specifications are specified for different instance models.
	// <br><li>To view specific values for `POSTPAID_BY_HOUR` instances, you can call [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) or refer to [Instance Types](https://cloud.tencent.com/document/product/213/11518). If this parameter is not specified, `S1.SMALL1` will be used by default.<br><li>For `CDHPAID` instances, the value of this parameter is in the format of `CDH_XCXG` based on the number of CPU cores and memory capacity. For example, if you want to create a CDH instance with a single-core CPU and 1 GB memory, specify this parameter as `CDH_1C1G`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default. When purchasing, you can specify 21 data disks, which can contain at most 1 LOCAL_BASIC data disk or LOCAL_SSD data disk, and at most 20 CLOUD_BASIC data disks, CLOUD_PREMIUM data disks, or CLOUD_SSD data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// VPC configurations. You can use this parameter to specify the VPC ID, subnet ID, etc. If this parameter is not specified, the basic network will be used by default. If a VPC IP is specified in this parameter, it will represent the primary ENI IP of each instance. The value of `InstanceCount` must be the same as the number of VPC IPs.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// The number of instances to be purchased. Value range: [1, 100]; default value: 1. The specified number of instances to be purchased cannot exceed the remaining quota allowed for the user. For more information on the quota, see [CVM instance purchase limit](https://intl.cloud.tencent.com/document/product/213/2664).
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instance name to be displayed. <br><li>If this parameter is not specified, "Unnamed" will be displayed by default. </li><li>If you purchase multiple instances at the same time and specify a pattern string `{R:x}`, numbers `[x, x+n-1]` will be generated, where `n` represents the number of instances purchased. For example, you specify a pattern string, `server_{R:3}`. If you only purchase 1 instance, the instance will be named `server_3`; if you purchase 2, they will be named `server_3` and `server_4`. You can specify multiple pattern strings in the format of `{R:x}`. </li><li>If you purchase multiple instances at the same time and do not specify a pattern string, the instance names will be suffixed by `1, 2...n`, where `n` represents the number of instances purchased. For example, if you purchase 2 instances and the instance name body is `server_`, the instance names will be `server_1` and `server_2`.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response. If this parameter is not specified, the instance will be associated with default security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// Specifies whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default. However, for custom images and images from the marketplace, Anti-DDoS and Cloud Monitor are not enabled by default. The original services in the image will be retained.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. <br>For more information, see “How to ensure idempotency”.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Host name of the CVM. <br><li>Periods (.) or hyphens (-) cannot be the start or end of a host name or appear consecutively in a host name.<br><li>For Windows instances, the host name must be 2-15 characters long and can contain uppercase and lowercase letters, numbers, and hyphens (-). It cannot contain periods (.) or contain only numbers. <br><li>For other instances, such as Linux instances, the host name must be 2-60 characters long. It supports multiple periods (.) and allows uppercase and lowercase letters, numbers, and hyphens (-) between any two periods (.).
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Scheduled tasks. You can use this parameter to specify scheduled tasks for the instance. Only scheduled termination is supported.
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// Placement group ID. You can only specify one.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds" list`

	// The tag description list. This parameter is used to bind a tag to a resource instance. A tag can only be bound to CVM instances.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// The market options of the instance.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// User data provided to the instance, which needs to be encoded in base64 format with the maximum size of 16KB. For more information on how to get the value of this parameter, see the commands you need to execute on startup for [Windows](https://intl.cloud.tencent.com/document/product/213/17526) or [Linux](https://intl.cloud.tencent.com/document/product/213/17525).
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Whether the request is a dry run only.
	// true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// false (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// If you use this API to create instance(s), this parameter will be returned, representing one or more instance `ID`s. Retuning the instance `ID` list does not necessarily mean that the instance(s) were created successfully. To check whether the instance(s) were created successfully, you can call [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) and check the states of the instances in `InstancesSet` in the response. If the state of an instance changes from "pending" to "running", it means that the instance has been created successfully.
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// Whether to enable [Cloud Monitor](/document/product/248). Valid values: <br><li>TRUE: enable Cloud Monitor <br><li>FALSE: do not enable Cloud Monitor <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// Whether to enable [Cloud Security](/document/product/296). Valid values: <br><li>TRUE: enable Cloud Security <br><li>FALSE: do not enable Cloud Security <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type SharePermission struct {

	// Time when an image was shared.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// ID of the account with which the image is shared.
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`
}

type Snapshot struct {

	// Snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// Type of the cloud disk used to create the snapshot. Valid values:
	// SYSTEM_DISK: system disk
	// DATA_DISK: data disk
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// Size of the cloud disk used to create the snapshot; unit: GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type SpotMarketOptions struct {

	// Bidding price
	MaxPrice *string `json:"MaxPrice,omitempty" name:"MaxPrice"`

	// Bidding request type. Currently only "one-time" is supported.
	SpotInstanceType *string `json:"SpotInstanceType,omitempty" name:"SpotInstanceType"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
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

func (r *StartInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// Whether to force shut down an instance after a normal shutdown fails. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails <br><li>FALSE: do not force shut down an instance after a normal shutdown fails <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Instance shutdown mode. Valid values: <br><li>SOFT_FIRST: perform a soft shutdown first, and force shut down the instance if the soft shutdown fails <br><li>HARD: force shut down the instance directly <br><li>SOFT: soft shutdown only <br>Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`

	// Billing method of a pay-as-you-go instance after shutdown.
	// Valid values: <br><li>KEEP_CHARGING: billing continues after shutdown <br><li>STOP_CHARGING: billing stops after shutdown <br>Default value: KEEP_CHARGING.
	// This parameter is only valid for some pay-as-you-go instances using cloud disks. For more information, see [No charges when shut down for pay-as-you-go instances](https://intl.cloud.tencent.com/document/product/213/19918).
	StoppedMode *string `json:"StoppedMode,omitempty" name:"StoppedMode"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
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

func (r *StopInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StorageBlock struct {

	// Local HDD storage type. Value: LOCAL_PRO.
	// Note: This field may return null, indicating that no valid value is found.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Minimum capacity of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum capacity of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest

	// List of image IDs. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://cloud.tencent.com/document/api/213/15715) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image). <br>The specified images must meet the following requirements: <br><li>The images must be in the `NORMAL` state. <br><li>The image size must be smaller than 50 GB. <br>For more information on image states, see [here](/document/api/213/9452#image_state).
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds" list`

	// List of destination regions for synchronization. A destination region must meet the following requirements: <br><li>It cannot be the source region. <br><li>It must be valid. <br><li>Currently some regions do not support image synchronization. <br>For specific regions, see [Region](https://cloud.tencent.com/document/product/213/6091).
	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions" list`
}

func (r *SyncImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {

	// System disk type. For more information on system disk types and their limits, refer to [Storage Overview](https://cloud.tencent.com/document/product/213/4952). Valid values: <br><li>LOCAL_BASIC: Local disk <br><li>LOCAL_SSD: Local SSD disk <br><li>CLOUD_BASIC: HDD cloud disk <br><li>CLOUD_PREMIUM: Premium cloud disk <br><li>CLOUD_SSD: SSD cloud disk <br><br>Default value: LOCAL_BASIC.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk ID. System disks whose type is `LOCAL_BASIC` or `LOCAL_SSD` do not have an ID and do not support this parameter currently.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// System disk size; unit: GB; default value: 50 GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Tag struct {

	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {

	// Type of the resources associated with the tags. Currently only "instance" and "host" are supported.
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// List of tags
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
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

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {

	// 
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID in the format `subnet-xxx`. To obtain valid subnet IDs, you can log in to the [console](https://console.cloud.tencent.com/vpc/subnet?rid=1) or call [DescribeSubnets](/document/api/215/15784) and look for the `unSubnetId` fields in the response. If you specify `DEFAULT` for both `SubnetId` and `VpcId` when creating an instance, the default VPC will be used.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether to use an instance as a public gateway. An instance can be used as a public gateway only when it has a public IP and resides in a VPC. Valid values: <br><li>TRUE: use the instance as a public gateway <br><li>FALSE: do not use the instance as a public gateway <br><br>Default value: FALSE.
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// Array of VPC subnet IPs. You can use this parameter when creating instances or modifying VPC attributes of instances. Currently you can specify multiple IPs in one subnet only when creating multiple instances at the same time.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type ZoneInfo struct {

	// Availability zone name, such as ap-guangzhou-3.
	// The following is a list of all availability zones:
	// <li> ap-chongqing-1 </li>
	// <li> ap-seoul-1 </li>
	// <li> ap-chengdu-1 </li>
	// <li> ap-chengdu-2 </li>
	// <li> ap-hongkong-1 </li>
	// <li> ap-hongkong-2 </li>
	// <li> ap-shenzhen-fsi-1 </li>
	// <li> ap-shenzhen-fsi-2 </li>
	// <li> ap-shenzhen-fsi-3 </li>
	// <li> ap-guangzhou-1 (sold out)</li>
	// <li> ap-guangzhou-2 (sold out)</li>
	// <li> ap-guangzhou-3 </li>
	// <li> ap-guangzhou-4 </li>
	// <li> ap-tokyo-1 </li>
	// <li> ap-singapore-1 </li>
	// <li> ap-shanghai-fsi-1 </li>
	// <li> ap-shanghai-fsi-2 </li>
	// <li> ap-shanghai-fsi-3 </li>
	// <li> ap-bangkok-1 </li>
	// <li> ap-shanghai-1 (sold out) </li>
	// <li> ap-shanghai-2 </li>
	// <li> ap-shanghai-3 </li>
	// <li> ap-shanghai-4 </li>
	// <li> ap-mumbai-1 </li>
	// <li> ap-mumbai-2 </li>
	// <li> eu-moscow-1 </li>
	// <li> ap-beijing-1 </li>
	// <li> ap-beijing-2 </li>
	// <li> ap-beijing-3 </li>
	// <li> ap-beijing-4 </li>
	// <li> na-siliconvalley-1 </li>
	// <li> na-siliconvalley-2 </li>
	// <li> eu-frankfurt-1 </li>
	// <li> na-toronto-1 </li>
	// <li> na-ashburn-1 </li>
	// <li> na-ashburn-2 </li>
	// <li> ap-nanjing-1 </li>
	// <li> ap-nanjing-2 </li>
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Availability zone description, such as Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Availability zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Availability zone status. Valid values: `AVAILABLE`: available; `UNAVAILABLE`: unavailable.
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}
