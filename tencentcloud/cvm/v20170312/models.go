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

type ActionTimer struct {
	// Additional data
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// Timer name. Currently `TerminateInstances` is the only supported value.
	TimerAction *string `json:"TimerAction,omitempty" name:"TimerAction"`

	// Execution time, which must be at least 5 minutes later than the current time. For example, 2018-5-29 11:26:40.
	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
}

// Predefined struct for user
type AllocateHostsRequestParams struct {
	// Instance location. This parameter is used to specify the attributes of an instance, such as its availability zone and project.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// A string used to ensure the idempotency of the request.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Configuration of prepaid instances. You can use the parameter to specify the attributes of prepaid instances, such as the subscription period and the auto-renewal plan. This parameter is required for prepaid instances.
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// Instance billing model, only monthly or yearly subscription supported. Default value: `PREPAID'.
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// CDH instance model. Default value: `HS1`.
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// Quantity of CDH instances purchased. Default value: 1.
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// Tag description. You can specify the parameter to associate a tag with an instance.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type AllocateHostsRequest struct {
	*tchttp.BaseRequest
	
	// Instance location. This parameter is used to specify the attributes of an instance, such as its availability zone and project.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// A string used to ensure the idempotency of the request.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Configuration of prepaid instances. You can use the parameter to specify the attributes of prepaid instances, such as the subscription period and the auto-renewal plan. This parameter is required for prepaid instances.
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid,omitempty" name:"HostChargePrepaid"`

	// Instance billing model, only monthly or yearly subscription supported. Default value: `PREPAID'.
	HostChargeType *string `json:"HostChargeType,omitempty" name:"HostChargeType"`

	// CDH instance model. Default value: `HS1`.
	HostType *string `json:"HostType,omitempty" name:"HostType"`

	// Quantity of CDH instances purchased. Default value: 1.
	HostCount *uint64 `json:"HostCount,omitempty" name:"HostCount"`

	// Tag description. You can specify the parameter to associate a tag with an instance.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

func (r *AllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ClientToken")
	delete(f, "HostChargePrepaid")
	delete(f, "HostChargeType")
	delete(f, "HostType")
	delete(f, "HostCount")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateHostsResponseParams struct {
	// The ID list of the CVM instances newly created on the CDH.
	HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *AllocateHostsResponseParams `json:"Response"`
}

func (r *AllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateInstancesKeyPairsRequestParams struct {
	// Instance ID(s). The maximum number of instances in each request is 100. <br>You can obtain the available instance IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the instance IDs. <br><li>Call [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Key ID(s). The maximum number of key pairs in each request is 100. The key pair ID is in the format of `skey-3glfot13`. <br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before associating a key pair with it. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). The maximum number of instances in each request is 100. <br>You can obtain the available instance IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the instance IDs. <br><li>Call [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Key ID(s). The maximum number of key pairs in each request is 100. The key pair ID is in the format of `skey-3glfot13`. <br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before associating a key pair with it. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
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
	delete(f, "InstanceIds")
	delete(f, "KeyIds")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateInstancesKeyPairsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type AssociateSecurityGroupsRequestParams struct {
	// ID of the security group to be associated, such as `sg-efil73jd`. Only one security group can be associated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID of the instance bound in the format of ins-lesecurk. You can specify up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the security group to be associated, such as `sg-efil73jd`. Only one security group can be associated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID of the instance bound in the format of ins-lesecurk. You can specify up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {
	// Purchased usage period, in month. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// Auto renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>Default value: NOTIFY_AND_AUTO_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type ChcDeployExtraConfig struct {

}

type ChcHost struct {
	// CHC host ID
	ChcId *string `json:"ChcId,omitempty" name:"ChcId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Server serial number
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

	// CHC host status<br/>
	// <ul>
	// <li>REGISTERED: The CHC host is registered, but the out-of-band network and deployment network are not configured.</li>
	// <li>VPC_READY: The out-of-band network and deployment network are configured.</li>
	// <li>PREPARED: It’s ready and can be associated with a CVM.</li>
	// <li>ONLINE: It’s already associated with a CVM.</li>
	// </ul>
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// Device type
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Availability zone
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Out-of-band network
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BmcVirtualPrivateCloud *VirtualPrivateCloud `json:"BmcVirtualPrivateCloud,omitempty" name:"BmcVirtualPrivateCloud"`

	// Out-of-band network IP
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BmcIp *string `json:"BmcIp,omitempty" name:"BmcIp"`

	// Out-of-band network security group ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitempty" name:"BmcSecurityGroupIds"`

	// Deployment network
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitempty" name:"DeployVirtualPrivateCloud"`

	// Deployment network IP
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeployIp *string `json:"DeployIp,omitempty" name:"DeployIp"`

	// Deployment network security group ID
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitempty" name:"DeploySecurityGroupIds"`

	// ID of the associated CVM
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`

	// Server creation time
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// Instance hardware description, including CPU cores, memory capacity and disk capacity.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HardwareDescription *string `json:"HardwareDescription,omitempty" name:"HardwareDescription"`

	// CPU cores of the CHC host
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// Memory capacity of the CHC host (unit: GB)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Disk capacity of the CHC host
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Disk *string `json:"Disk,omitempty" name:"Disk"`

	// MAC address assigned under the out-of-band network
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	BmcMAC *string `json:"BmcMAC,omitempty" name:"BmcMAC"`

	// MAC address assigned under the deployment network
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeployMAC *string `json:"DeployMAC,omitempty" name:"DeployMAC"`

	// Management type
	// HOSTING: Hosting
	// TENANT: Leasing
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TenantType *string `json:"TenantType,omitempty" name:"TenantType"`

	// CHC DHCP option, which is used for MiniOS debugging.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DeployExtraConfig *ChcDeployExtraConfig `json:"DeployExtraConfig,omitempty" name:"DeployExtraConfig"`
}

type ChcHostDeniedActions struct {
	// CHC instance ID
	ChcId *string `json:"ChcId,omitempty" name:"ChcId"`

	// CHC instance status
	State *string `json:"State,omitempty" name:"State"`

	// Actions not allowed for the current CHC instance
	DenyActions []*string `json:"DenyActions,omitempty" name:"DenyActions"`
}

// Predefined struct for user
type ConfigureChcAssistVpcRequestParams struct {
	// CHC host IDs
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// Out-of-band network information
	BmcVirtualPrivateCloud *VirtualPrivateCloud `json:"BmcVirtualPrivateCloud,omitempty" name:"BmcVirtualPrivateCloud"`

	// Out-of-band network security group list
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitempty" name:"BmcSecurityGroupIds"`

	// Deployment network information
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitempty" name:"DeployVirtualPrivateCloud"`

	// Deployment network security group list
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitempty" name:"DeploySecurityGroupIds"`
}

type ConfigureChcAssistVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC host IDs
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// Out-of-band network information
	BmcVirtualPrivateCloud *VirtualPrivateCloud `json:"BmcVirtualPrivateCloud,omitempty" name:"BmcVirtualPrivateCloud"`

	// Out-of-band network security group list
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitempty" name:"BmcSecurityGroupIds"`

	// Deployment network information
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitempty" name:"DeployVirtualPrivateCloud"`

	// Deployment network security group list
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitempty" name:"DeploySecurityGroupIds"`
}

func (r *ConfigureChcAssistVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcAssistVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "BmcVirtualPrivateCloud")
	delete(f, "BmcSecurityGroupIds")
	delete(f, "DeployVirtualPrivateCloud")
	delete(f, "DeploySecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureChcAssistVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureChcAssistVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ConfigureChcAssistVpcResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureChcAssistVpcResponseParams `json:"Response"`
}

func (r *ConfigureChcAssistVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcAssistVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureChcDeployVpcRequestParams struct {
	// CHC instance ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// Deployment network information
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitempty" name:"DeployVirtualPrivateCloud"`

	// Deployment network security group list
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitempty" name:"DeploySecurityGroupIds"`
}

type ConfigureChcDeployVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC instance ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// Deployment network information
	DeployVirtualPrivateCloud *VirtualPrivateCloud `json:"DeployVirtualPrivateCloud,omitempty" name:"DeployVirtualPrivateCloud"`

	// Deployment network security group list
	DeploySecurityGroupIds []*string `json:"DeploySecurityGroupIds,omitempty" name:"DeploySecurityGroupIds"`
}

func (r *ConfigureChcDeployVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcDeployVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "DeployVirtualPrivateCloud")
	delete(f, "DeploySecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureChcDeployVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureChcDeployVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ConfigureChcDeployVpcResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureChcDeployVpcResponseParams `json:"Response"`
}

func (r *ConfigureChcDeployVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureChcDeployVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoverGroupRequestParams struct {
	// Name of the spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of the spread placement group. Valid values: <br><li>HOST: physical machine <br><li>SW: switch <br><li>RACK: rack
	Type *string `json:"Type,omitempty" name:"Type"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. <br>For more information, see 'How to ensure idempotency'.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

type CreateDisasterRecoverGroupRequest struct {
	*tchttp.BaseRequest
	
	// Name of the spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type of the spread placement group. Valid values: <br><li>HOST: physical machine <br><li>SW: switch <br><li>RACK: rack
	Type *string `json:"Type,omitempty" name:"Type"`

	// A string used to ensure the idempotency of the request, which is generated by the user and must be unique to each request. The maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed. <br>For more information, see 'How to ensure idempotency'.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *CreateDisasterRecoverGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoverGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisasterRecoverGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisasterRecoverGroupResponseParams struct {
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
}

type CreateDisasterRecoverGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisasterRecoverGroupResponseParams `json:"Response"`
}

func (r *CreateDisasterRecoverGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisasterRecoverGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageRequestParams struct {
	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// ID of the instance from which an image will be created. This parameter is required when using instance to create an image.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Whether to force shut down an instance to create an image when a soft shutdown fails
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`

	// Whether to enable Sysprep when creating a Windows image.
	// Valid values: `TRUE` and `FALSE`; default value: `FALSE`.
	// 
	// Click [here](https://intl.cloud.tencent.com/document/product/213/43498?from_cn_redirect=1) to learn more about Sysprep.
	Sysprep *string `json:"Sysprep,omitempty" name:"Sysprep"`

	// Specified data disk ID included in the full image created from the instance.
	DataDiskIds []*string `json:"DataDiskIds,omitempty" name:"DataDiskIds"`

	// Specified snapshot ID used to create an image. A system disk snapshot must be included. It cannot be passed together with `InstanceId`.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Success status of this request, without affecting the resources involved
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	
	// Image name
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// ID of the instance from which an image will be created. This parameter is required when using instance to create an image.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Image description
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// Whether to force shut down an instance to create an image when a soft shutdown fails
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`

	// Whether to enable Sysprep when creating a Windows image.
	// Valid values: `TRUE` and `FALSE`; default value: `FALSE`.
	// 
	// Click [here](https://intl.cloud.tencent.com/document/product/213/43498?from_cn_redirect=1) to learn more about Sysprep.
	Sysprep *string `json:"Sysprep,omitempty" name:"Sysprep"`

	// Specified data disk ID included in the full image created from the instance.
	DataDiskIds []*string `json:"DataDiskIds,omitempty" name:"DataDiskIds"`

	// Specified snapshot ID used to create an image. A system disk snapshot must be included. It cannot be passed together with `InstanceId`.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// Success status of this request, without affecting the resources involved
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageName")
	delete(f, "InstanceId")
	delete(f, "ImageDescription")
	delete(f, "ForcePoweroff")
	delete(f, "Sysprep")
	delete(f, "DataDiskIds")
	delete(f, "SnapshotIds")
	delete(f, "DryRun")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageResponseParams struct {
	// Image ID.
	// Note: This field may return null, indicating that no valid value was found.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageResponseParams `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairRequestParams struct {
	// Name of the key pair, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// The ID of the project to which the new key pair belongs.
	// You can query the project IDs in two ways:
	// <li>Query the project IDs in the project list.
	// <li>Call `DescribeProject` and look for `projectId` in the response.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Tag description list. This parameter is used to bind a tag to a key pair.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
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

	// Tag description list. This parameter is used to bind a tag to a key pair.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
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
	delete(f, "ProjectId")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairResponseParams struct {
	// Key pair information.
	KeyPair *KeyPair `json:"KeyPair,omitempty" name:"KeyPair"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateLaunchTemplateVersionRequestParams struct {
	// Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone, project, and CDH (for dedicated CVMs)
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance launch template ID. This parameter is used as a basis for creating new template versions.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// This parameter, when specified, is used to create instance launch templates. If this parameter is not specified, the default version will be used.
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitempty" name:"LaunchTemplateVersion"`

	// Description of instance launch template versions. This parameter can contain 2-256 characters.
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`

	// The instance model. Different resource specifications are specified for different instance models.
	// <br><li>To view specific values for `PREPAID` or `POSTPAID_BY_HOUR` instances, you can call [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) or refer to [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1). If this parameter is not specified, the system will specify the default model according to the dynamic resource sales in the current region. <br><li>For `CDHPAID` instances, the value of this parameter is in the format of `CDH_XCXG` based on the number of CPU cores and memory capacity. For example, if you want to create a CDH instance with a single-core CPU and 1 GB memory, you need to specify this parameter as `CDH_1C1G`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The [image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are four types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><li>Marketplace images (for Chinese mainland only)</li><br/>To check the image ID:<br/><li>For public images, custom images, and shared images, go to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE). For marketplace images, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1), pass in `InstanceType` to retrieve the list of images supported by the current model, and then find the `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default. When purchasing, you can specify 21 data disks, which can contain at most 1 LOCAL_BASIC or LOCAL_SSD data disk, and at most 20 CLOUD_BASIC, CLOUD_PREMIUM, or CLOUD_SSD data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Configuration information of VPC. This parameter is used to specify VPC ID and subnet ID, etc. If this parameter is not specified, the classic network is used by default. If a VPC IP is specified in this parameter, it indicates the primary ENI IP of each instance. The value of parameter InstanceCount must be same as the number of VPC IPs, which cannot be greater than 20.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Number of instances to be purchased. Value range for monthly-subscribed instances: [1, 300]. Value range for pay-as-you-go instances: [1, 100]. Default value: 1. The specified number of instances to be purchased cannot exceed the remaining quota allowed for the user. For more information on quota, see CVM instance [Purchase Limits](https://intl.cloud.tencent.com/document/product/213/2664).
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instance name to be displayed. <br><li>If this parameter is not specified, "Unnamed" will be displayed by default. </li><li>If you purchase multiple instances at the same time and specify a pattern string `{R:x}`, numbers `[x, x+n-1]` will be generated, where `n` represents the number of instances purchased. For example, you specify a pattern string, `server_{R:3}`. If you only purchase 1 instance, the instance will be named `server_3`; if you purchase 2, they will be named `server_3` and `server_4`. You can specify multiple pattern strings in the format of `{R:x}`. </li><li>If you purchase multiple instances at the same time and do not specify a pattern string, the instance names will be suffixed by `1, 2...n`, where `n` represents the number of instances purchased. For example, if you purchase 2 instances and the instance name body is `server_`, the instance names will be `server_1` and `server_2`. </li><li>This parameter can contain up to 60 characters, including the pattern string.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response. If this parameter is not specified, the instance will be associated with default security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default. However, for custom images and images from the marketplace, Anti-DDoS and Cloud Monitor are not enabled by default. The original services in the image will be retained.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Host name of the CVM. <br><li>Dots (.) or hyphens (-) cannot be the start or end of a host name or appear consecutively in a host name. <br><li>For Windows instances, the host name must be 2-15 characters long and can contain uppercase and lowercase letters, numbers, and hyphens (-). It cannot contain dots (.) or contain only numbers. <br><li>For other instances, such as Linux instances, the host name must be 2-60 characters long. It supports multiple dots (.) and allows uppercase and lowercase letters, numbers, and hyphens (-) between any two dots (.).
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Scheduled tasks. You can use this parameter to specify scheduled tasks for the instance. Only scheduled termination is supported.
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// Placement group ID. You can only specify one.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// The tag description list. This parameter is used to bind a tag to a resource instance. A tag can only be bound to CVM instances.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16KB. For more information on how to get the value of this parameter, see the commands you need to execute on startup for [Windows](https://intl.cloud.tencent.com/document/product/213/17526) or [Linux](https://intl.cloud.tencent.com/document/product/213/17525).
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Whether the request is a dry run only.
	// `true`: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// `false` (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// CAM role name, which can be obtained from the `roleName` field in the response of the [`DescribeRoleList`](https://intl.cloud.tencent.com/document/product/598/36223?from_cn_redirect=1) API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`

	// Instance [billing mode](https://intl.cloud.tencent.com/document/product/213/2180?from_cn_redirect=1). Valid values: <br><li>`PREPAID`: monthly subscription <br><li>`POSTPAID_BY_HOUR`: pay-as-you-go on an hourly basis <br><li>`CDHPAID`: billed on the associated CDH instance. <br><li>`SPOTPAID`: spot instances. <br>Default value: `POSTPAID_BY_HOUR`.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// Whether the termination protection is enabled. Values: <br><li>`TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.<br><li>`FALSE`: Do not enable the instance protection.<br><br>Default value: `FALSE`.
	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

type CreateLaunchTemplateVersionRequest struct {
	*tchttp.BaseRequest
	
	// Location of the instance. You can use this parameter to specify the attributes of the instance, such as its availability zone, project, and CDH (for dedicated CVMs)
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance launch template ID. This parameter is used as a basis for creating new template versions.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// This parameter, when specified, is used to create instance launch templates. If this parameter is not specified, the default version will be used.
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitempty" name:"LaunchTemplateVersion"`

	// Description of instance launch template versions. This parameter can contain 2-256 characters.
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`

	// The instance model. Different resource specifications are specified for different instance models.
	// <br><li>To view specific values for `PREPAID` or `POSTPAID_BY_HOUR` instances, you can call [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) or refer to [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1). If this parameter is not specified, the system will specify the default model according to the dynamic resource sales in the current region. <br><li>For `CDHPAID` instances, the value of this parameter is in the format of `CDH_XCXG` based on the number of CPU cores and memory capacity. For example, if you want to create a CDH instance with a single-core CPU and 1 GB memory, you need to specify this parameter as `CDH_1C1G`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The [image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are four types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><li>Marketplace images (for Chinese mainland only)</li><br/>To check the image ID:<br/><li>For public images, custom images, and shared images, go to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE). For marketplace images, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1), pass in `InstanceType` to retrieve the list of images supported by the current model, and then find the `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default. When purchasing, you can specify 21 data disks, which can contain at most 1 LOCAL_BASIC or LOCAL_SSD data disk, and at most 20 CLOUD_BASIC, CLOUD_PREMIUM, or CLOUD_SSD data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Configuration information of VPC. This parameter is used to specify VPC ID and subnet ID, etc. If this parameter is not specified, the classic network is used by default. If a VPC IP is specified in this parameter, it indicates the primary ENI IP of each instance. The value of parameter InstanceCount must be same as the number of VPC IPs, which cannot be greater than 20.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Number of instances to be purchased. Value range for monthly-subscribed instances: [1, 300]. Value range for pay-as-you-go instances: [1, 100]. Default value: 1. The specified number of instances to be purchased cannot exceed the remaining quota allowed for the user. For more information on quota, see CVM instance [Purchase Limits](https://intl.cloud.tencent.com/document/product/213/2664).
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instance name to be displayed. <br><li>If this parameter is not specified, "Unnamed" will be displayed by default. </li><li>If you purchase multiple instances at the same time and specify a pattern string `{R:x}`, numbers `[x, x+n-1]` will be generated, where `n` represents the number of instances purchased. For example, you specify a pattern string, `server_{R:3}`. If you only purchase 1 instance, the instance will be named `server_3`; if you purchase 2, they will be named `server_3` and `server_4`. You can specify multiple pattern strings in the format of `{R:x}`. </li><li>If you purchase multiple instances at the same time and do not specify a pattern string, the instance names will be suffixed by `1, 2...n`, where `n` represents the number of instances purchased. For example, if you purchase 2 instances and the instance name body is `server_`, the instance names will be `server_1` and `server_2`. </li><li>This parameter can contain up to 60 characters, including the pattern string.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response. If this parameter is not specified, the instance will be associated with default security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default. However, for custom images and images from the marketplace, Anti-DDoS and Cloud Monitor are not enabled by default. The original services in the image will be retained.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Host name of the CVM. <br><li>Dots (.) or hyphens (-) cannot be the start or end of a host name or appear consecutively in a host name. <br><li>For Windows instances, the host name must be 2-15 characters long and can contain uppercase and lowercase letters, numbers, and hyphens (-). It cannot contain dots (.) or contain only numbers. <br><li>For other instances, such as Linux instances, the host name must be 2-60 characters long. It supports multiple dots (.) and allows uppercase and lowercase letters, numbers, and hyphens (-) between any two dots (.).
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Scheduled tasks. You can use this parameter to specify scheduled tasks for the instance. Only scheduled termination is supported.
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// Placement group ID. You can only specify one.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// The tag description list. This parameter is used to bind a tag to a resource instance. A tag can only be bound to CVM instances.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16KB. For more information on how to get the value of this parameter, see the commands you need to execute on startup for [Windows](https://intl.cloud.tencent.com/document/product/213/17526) or [Linux](https://intl.cloud.tencent.com/document/product/213/17525).
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Whether the request is a dry run only.
	// `true`: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// `false` (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// CAM role name, which can be obtained from the `roleName` field in the response of the [`DescribeRoleList`](https://intl.cloud.tencent.com/document/product/598/36223?from_cn_redirect=1) API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`

	// Instance [billing mode](https://intl.cloud.tencent.com/document/product/213/2180?from_cn_redirect=1). Valid values: <br><li>`PREPAID`: monthly subscription <br><li>`POSTPAID_BY_HOUR`: pay-as-you-go on an hourly basis <br><li>`CDHPAID`: billed on the associated CDH instance. <br><li>`SPOTPAID`: spot instances. <br>Default value: `POSTPAID_BY_HOUR`.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// Whether the termination protection is enabled. Values: <br><li>`TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.<br><li>`FALSE`: Do not enable the instance protection.<br><br>Default value: `FALSE`.
	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

func (r *CreateLaunchTemplateVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchTemplateVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "LaunchTemplateId")
	delete(f, "LaunchTemplateVersion")
	delete(f, "LaunchTemplateVersionDescription")
	delete(f, "InstanceType")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InternetAccessible")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "ClientToken")
	delete(f, "HostName")
	delete(f, "ActionTimer")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "TagSpecification")
	delete(f, "InstanceMarketOptions")
	delete(f, "UserData")
	delete(f, "DryRun")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DisableApiTermination")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaunchTemplateVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchTemplateVersionResponseParams struct {
	// Version number of the new instance launch template.
	LaunchTemplateVersionNumber *int64 `json:"LaunchTemplateVersionNumber,omitempty" name:"LaunchTemplateVersionNumber"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLaunchTemplateVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaunchTemplateVersionResponseParams `json:"Response"`
}

func (r *CreateLaunchTemplateVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchTemplateVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// Data disk size (in GB). The minimum adjustment increment is 10 GB. The value range varies by data disk type. For more information on limits, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952?from_cn_redirect=1). The default value is 0, indicating that no data disk is purchased. For more information, see the product documentation.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952?from_cn_redirect=1). Valid values: <br><li>LOCAL_BASIC: local disk<br><li>LOCAL_SSD: local SSD disk<br><li>LOCAL_NVME: local NVME disk, specified in the `InstanceType`<br><li>LOCAL_PRO: local HDD disk, specified in the `InstanceType`<br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_PREMIUM: Premium Cloud Storage<br><li>CLOUD_SSD: SSD<br><li>CLOUD_HSSD: Enhanced SSD<br><li>CLOUD_TSSD: Tremendous SSD<br><li>CLOUD_BSSD: Balanced SSD<br><br>Default value: LOCAL_BASIC.<br><br>This parameter is invalid for the `ResizeInstanceDisk` API.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Data disk ID. Note that it’s not available for `LOCAL_BASIC` and `LOCAL_SSD` disks.
	// It is only used as a response parameter for APIs such as `DescribeInstances`, and cannot be used as a request parameter for APIs such as `RunInstances`.
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

	// ID of the custom CMK in the format of UUID or “kms-abcd1234”. This parameter is used to encrypt cloud disks.
	// 
	// Currently, this parameter is only used in the `RunInstances` API.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// Cloud disk performance, in MB/s
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ThroughputPerformance *int64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// ID of the dedicated cluster to which the instance belongs.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

// Predefined struct for user
type DeleteDisasterRecoverGroupsRequestParams struct {
	// ID list of spread placement groups, obtainable via the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/api/213/17810?from_cn_redirect=1) API. You can operate up to 100 spread placement groups in each request.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
}

type DeleteDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID list of spread placement groups, obtainable via the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/api/213/17810?from_cn_redirect=1) API. You can operate up to 100 spread placement groups in each request.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`
}

func (r *DeleteDisasterRecoverGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoverGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDisasterRecoverGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDisasterRecoverGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDisasterRecoverGroupsResponseParams `json:"Response"`
}

func (r *DeleteDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesRequestParams struct {
	// List of the IDs of the instances to be deleted.
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// Whether to delete the snapshot associated with the image
	DeleteBindedSnap *bool `json:"DeleteBindedSnap,omitempty" name:"DeleteBindedSnap"`

	// Check whether deleting an image is supported
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest
	
	// List of the IDs of the instances to be deleted.
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// Whether to delete the snapshot associated with the image
	DeleteBindedSnap *bool `json:"DeleteBindedSnap,omitempty" name:"DeleteBindedSnap"`

	// Check whether deleting an image is supported
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "DeleteBindedSnap")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImagesResponseParams `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeyPairsRequestParams struct {
	// Key ID(s). The maximum number of key pairs in each request is 100. <br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Key ID(s). The maximum number of key pairs in each request is 100. <br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
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

// Predefined struct for user
type DeleteKeyPairsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteLaunchTemplateRequestParams struct {
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
}

type DeleteLaunchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`
}

func (r *DeleteLaunchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaunchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLaunchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaunchTemplateResponseParams `json:"Response"`
}

func (r *DeleteLaunchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateVersionsRequestParams struct {
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// List of instance launch template versions.
	LaunchTemplateVersions []*int64 `json:"LaunchTemplateVersions,omitempty" name:"LaunchTemplateVersions"`
}

type DeleteLaunchTemplateVersionsRequest struct {
	*tchttp.BaseRequest
	
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// List of instance launch template versions.
	LaunchTemplateVersions []*int64 `json:"LaunchTemplateVersions,omitempty" name:"LaunchTemplateVersions"`
}

func (r *DeleteLaunchTemplateVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	delete(f, "LaunchTemplateVersions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaunchTemplateVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchTemplateVersionsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLaunchTemplateVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaunchTemplateVersionsResponseParams `json:"Response"`
}

func (r *DeleteLaunchTemplateVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchTemplateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcDeniedActionsRequestParams struct {
	// CHC instance ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`
}

type DescribeChcDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// CHC instance ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`
}

func (r *DescribeChcDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChcDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcDeniedActionsResponseParams struct {
	// Actions not allowed for the CHC instance
	ChcHostDeniedActionSet []*ChcHostDeniedActions `json:"ChcHostDeniedActionSet,omitempty" name:"ChcHostDeniedActionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChcDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChcDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeChcDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcHostsRequestParams struct {
	// CHC host ID. Up to 100 instances per request is allowed. `ChcIds` and `Filters` cannot be specified at the same time.
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>availability zone</strong>, such as `ap-guangzhou-1`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p><p style="padding-left: 30px;">Valid values: See <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Regions and Availability Zones</a></p>
	// <li><strong>instance-name</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>instance name</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>instance-state</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>instance status</strong>. For status details, see [InstanceStatus](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#InstanceStatus).</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>device-type</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>device type</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>vpc-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>unique VPC ID</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>subnet-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>unique VPC subnet ID</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The offset. Default value: `0`. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: `20`. Maximum value: `100`. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeChcHostsRequest struct {
	*tchttp.BaseRequest
	
	// CHC host ID. Up to 100 instances per request is allowed. `ChcIds` and `Filters` cannot be specified at the same time.
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>availability zone</strong>, such as `ap-guangzhou-1`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p><p style="padding-left: 30px;">Valid values: See <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Regions and Availability Zones</a></p>
	// <li><strong>instance-name</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>instance name</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>instance-state</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>instance status</strong>. For status details, see [InstanceStatus](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#InstanceStatus).</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>device-type</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>device type</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>vpc-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>unique VPC ID</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>subnet-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>unique VPC subnet ID</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The offset. Default value: `0`. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: `20`. Maximum value: `100`. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeChcHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChcHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChcHostsResponseParams struct {
	// Number of eligible instances
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of returned instances
	ChcHostSet []*ChcHost `json:"ChcHostSet,omitempty" name:"ChcHostSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChcHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChcHostsResponseParams `json:"Response"`
}

func (r *DescribeChcHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChcHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupQuotaRequestParams struct {

}

type DescribeDisasterRecoverGroupQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDisasterRecoverGroupQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverGroupQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupQuotaResponseParams struct {
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
}

type DescribeDisasterRecoverGroupQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverGroupQuotaResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupsRequestParams struct {
	// ID list of spread placement groups. You can operate up to 100 spread placement groups in each request.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// Name of a spread placement group. Fuzzy match is supported.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDisasterRecoverGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID list of spread placement groups. You can operate up to 100 spread placement groups in each request.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisasterRecoverGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisasterRecoverGroupsResponseParams struct {
	// Information on spread placement groups.
	DisasterRecoverGroupSet []*DisasterRecoverGroup `json:"DisasterRecoverGroupSet,omitempty" name:"DisasterRecoverGroupSet"`

	// Total number of placement groups of the user.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDisasterRecoverGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisasterRecoverGroupsResponseParams `json:"Response"`
}

func (r *DescribeDisasterRecoverGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisasterRecoverGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsRequestParams struct {
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">Filter by the availability zone, such as `ap-guangzhou-1`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p><p style="padding-left: 30px;">Valid values: See <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Regions and Availability Zones</a></p>
	// <li><strong>project-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the project ID. You can query the list of created projects by calling `DescribeProject` or logging in to the [CVM console](https://console.cloud.tencent.com/cvm/index). You can also call `AddProject` to create projects. The project ID is like 1002189. </p><p style="padding-left: 30px;">Type: Integer</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>host-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the CDH instance ID. Format: host-xxxxxxxx. </p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>host-name</strong></li>
	// <p style="padding-left: 30px;">Filter by the host name.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>host-state</strong></li>
	// <p style="padding-left: 30px;">Filter by the CDH instance status. (`PENDING`: Creating | `LAUNCH_FAILURE`: Failed to create | `RUNNING`: Running | `EXPIRED`: Expired)</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest
	
	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">Filter by the availability zone, such as `ap-guangzhou-1`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p><p style="padding-left: 30px;">Valid values: See <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Regions and Availability Zones</a></p>
	// <li><strong>project-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the project ID. You can query the list of created projects by calling `DescribeProject` or logging in to the [CVM console](https://console.cloud.tencent.com/cvm/index). You can also call `AddProject` to create projects. The project ID is like 1002189. </p><p style="padding-left: 30px;">Type: Integer</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>host-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the CDH instance ID. Format: host-xxxxxxxx. </p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>host-name</strong></li>
	// <p style="padding-left: 30px;">Filter by the host name.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>host-state</strong></li>
	// <p style="padding-left: 30px;">Filter by the CDH instance status. (`PENDING`: Creating | `LAUNCH_FAILURE`: Failed to create | `RUNNING`: Running | `EXPIRED`: Expired)</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsResponseParams struct {
	// Total number of CDH instances meeting the query conditions
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Information on CDH instances
	HostSet []*HostItem `json:"HostSet,omitempty" name:"HostSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsResponseParams `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageQuotaRequestParams struct {

}

type DescribeImageQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageQuotaResponseParams struct {
	// The image quota of an account
	ImageNumQuota *int64 `json:"ImageNumQuota,omitempty" name:"ImageNumQuota"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageQuotaResponseParams `json:"Response"`
}

func (r *DescribeImageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSharePermissionRequestParams struct {
	// The ID of the image to be shared
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSharePermissionResponseParams struct {
	// Information on image sharing.
	SharePermissionSet []*SharePermission `json:"SharePermissionSet,omitempty" name:"SharePermissionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageSharePermissionResponseParams `json:"Response"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// List of image IDs, such as `img-gvbnzy6f`. For the format of array-type parameters, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response. <br><li>View the image IDs in the [Image Console](https://console.cloud.tencent.com/cvm/image).
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// Filters. Each request can have up to 10 `Filters`, and 5 `Filters.Values` for each filter. `ImageIds` and `Filters` cannot be specified at the same time. See details:
	// 
	// <li><strong>image-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image ID</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>image-type</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image type</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p><p style="padding-left: 30px;">Options:</p><p style="padding-left: 30px;">`PRIVATE_IMAGE`: Private images (images created by this account)</p><p style="padding-left: 30px;">`PUBLIC_IMAGE`: Public images (Tencent Cloud official images)</p><p style="padding-left: 30px;">`SHARED_IMAGE`: Shared images (images shared by other accounts to this account)</p>
	// <li><strong>image-name</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image name</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>platform</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image operating system</strong>, such as `CentOS`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>tag key</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>tag-value</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>tag value</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>tag:tag-key</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>tag key-value pair</strong>. Replace “tag-key” with the actual value. </p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0. For more information on `Offset`, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Instance type, e.g. `S1.SMALL1`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
	// List of image IDs, such as `img-gvbnzy6f`. For the format of array-type parameters, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response. <br><li>View the image IDs in the [Image Console](https://console.cloud.tencent.com/cvm/image).
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// Filters. Each request can have up to 10 `Filters`, and 5 `Filters.Values` for each filter. `ImageIds` and `Filters` cannot be specified at the same time. See details:
	// 
	// <li><strong>image-id</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image ID</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>image-type</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image type</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p><p style="padding-left: 30px;">Options:</p><p style="padding-left: 30px;">`PRIVATE_IMAGE`: Private images (images created by this account)</p><p style="padding-left: 30px;">`PUBLIC_IMAGE`: Public images (Tencent Cloud official images)</p><p style="padding-left: 30px;">`SHARED_IMAGE`: Shared images (images shared by other accounts to this account)</p>
	// <li><strong>image-name</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image name</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>platform</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>image operating system</strong>, such as `CentOS`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>tag-key</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>tag key</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>tag-value</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>tag value</strong>.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// <li><strong>tag:tag-key</strong></li>
	// <p style="padding-left: 30px;">Filter by the <strong>tag key-value pair</strong>. Replace “tag-key” with the actual value. </p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0. For more information on `Offset`, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Instance type, e.g. `S1.SMALL1`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// Information on an image, including its state and attributes.
	ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet"`

	// Number of images meeting the filtering conditions.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportImageOsRequestParams struct {

}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImportImageOsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportImageOsResponseParams struct {
	// Supported operating system types of imported images.
	ImportImageOsListSupported *ImageOsList `json:"ImportImageOsListSupported,omitempty" name:"ImportImageOsListSupported"`

	// Supported operating system versions of imported images. 
	ImportImageOsVersionSet []*OsVersion `json:"ImportImageOsVersionSet,omitempty" name:"ImportImageOsVersionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImportImageOsResponseParams `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceFamilyConfigsRequestParams struct {

}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceFamilyConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceFamilyConfigsResponseParams struct {
	// List of instance model families
	InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet,omitempty" name:"InstanceFamilyConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceFamilyConfigsResponseParams `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesOperationLimitRequestParams struct {
	// Query by instance ID(s). You can obtain the instance IDs from the value of `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API. For example, instance ID: ins-xxxxxxxx. (For the specific format, refer to section `ids.N` of the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).) You can query up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Operation on the instance(s).
	// <li> INSTANCE_DEGRADE: downgrade the instance configurations</li>
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type DescribeInstancesOperationLimitRequest struct {
	*tchttp.BaseRequest
	
	// Query by instance ID(s). You can obtain the instance IDs from the value of `InstanceId` returned by the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API. For example, instance ID: ins-xxxxxxxx. (For the specific format, refer to section `ids.N` of the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).) You can query up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Operation on the instance(s).
	// <li> INSTANCE_DEGRADE: downgrade the instance configurations</li>
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeInstancesOperationLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesOperationLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesOperationLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesOperationLimitResponseParams struct {
	// The maximum number of times you can modify the instance configurations (degrading the configurations)
	InstanceOperationLimitSet []*OperationCountLimit `json:"InstanceOperationLimitSet,omitempty" name:"InstanceOperationLimitSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesOperationLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesOperationLimitResponseParams `json:"Response"`
}

func (r *DescribeInstancesOperationLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesOperationLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// Query by instance ID(s). For example, instance ID: `ins-xxxxxxxx`. For the specific format, refer to section `Ids.N` of the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You can query up to 100 instances in each request. However, `InstanceIds` and `Filters` cannot be specified at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Filters
	// <li> `zone` - String - Optional - Filter by the availability zone.</li>
	// <li> `project-id` - Integer - Optional - Filter by the project ID. You can query the list of created projects by calling `DescribeProject` or logging in to the [CVM console](https://console.cloud.tencent.com/cvm/index). You can also call `AddProject` to create projects. </li>
	// <li> `host-id` - String - Optional - Filter by the CDH instance ID. Format: `host-xxxxxxxx`.</li>
	// </li>`vpc-id` - String - Optional - Filter by the VPC ID. Format: `vpc-xxxxxxxx`.</li>
	// <li> `subnet-id` - String - Optional - Filter by the subnet ID. Format: `subnet-xxxxxxxx`.</li>
	// </li>`instance-id` - String - Optional - Filter by the instance ID. Format: `ins-xxxxxxxx`.</li>
	// </li>`security-group-id` - String - Optional - Filter by the security group ID. Format: `sg-8jlk3f3r`.</li>
	// </li>`instance-name` - String - Optional - Filter by the instance name.</li>
	// </li>`instance-charge-type` - String - Optional - Filter by the instance billing method. `POSTPAID_BY_HOUR`: pay-as-you-go | `CDHPAID`: You are only billed for [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances, not the CVMs running on the [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.</li>
	// </li>`private-ip-address` - String - Optional - Filter by the private IP address of the instance's primary ENI.</li>
	// </li>`public-ip-address` - String - Optional - Filter by the public IP address of the instance's primary ENI, including the IP addresses automatically assigned during the instance creation and the EIPs manually associated after the instance creation.</li>
	// <li> `tag-key` - String - Optional - Filter by the tag key.</li>
	// </li>`tag-value` - String - Optional - Filter by the tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Filter by the tag key-value pair. Replace `tag-key` with the actual tag keys. See example 2.</li>
	// Each request can have up to 10 `Filters` and 5 `Filters.Values`. You cannot specify `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Query by instance ID(s). For example, instance ID: `ins-xxxxxxxx`. For the specific format, refer to section `Ids.N` of the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You can query up to 100 instances in each request. However, `InstanceIds` and `Filters` cannot be specified at the same time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Filters
	// <li> `zone` - String - Optional - Filter by the availability zone.</li>
	// <li> `project-id` - Integer - Optional - Filter by the project ID. You can query the list of created projects by calling `DescribeProject` or logging in to the [CVM console](https://console.cloud.tencent.com/cvm/index). You can also call `AddProject` to create projects. </li>
	// <li> `host-id` - String - Optional - Filter by the CDH instance ID. Format: `host-xxxxxxxx`.</li>
	// </li>`vpc-id` - String - Optional - Filter by the VPC ID. Format: `vpc-xxxxxxxx`.</li>
	// <li> `subnet-id` - String - Optional - Filter by the subnet ID. Format: `subnet-xxxxxxxx`.</li>
	// </li>`instance-id` - String - Optional - Filter by the instance ID. Format: `ins-xxxxxxxx`.</li>
	// </li>`security-group-id` - String - Optional - Filter by the security group ID. Format: `sg-8jlk3f3r`.</li>
	// </li>`instance-name` - String - Optional - Filter by the instance name.</li>
	// </li>`instance-charge-type` - String - Optional - Filter by the instance billing method. `POSTPAID_BY_HOUR`: pay-as-you-go | `CDHPAID`: You are only billed for [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances, not the CVMs running on the [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.</li>
	// </li>`private-ip-address` - String - Optional - Filter by the private IP address of the instance's primary ENI.</li>
	// </li>`public-ip-address` - String - Optional - Filter by the public IP address of the instance's primary ENI, including the IP addresses automatically assigned during the instance creation and the EIPs manually associated after the instance creation.</li>
	// <li> `tag-key` - String - Optional - Filter by the tag key.</li>
	// </li>`tag-value` - String - Optional - Filter by the tag value.</li>
	// <li> `tag:tag-key` - String - Optional - Filter by the tag key-value pair. Replace `tag-key` with the actual tag keys. See example 2.</li>
	// Each request can have up to 10 `Filters` and 5 `Filters.Values`. You cannot specify `InstanceIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
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

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Number of instances meeting the filtering conditions.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Detailed instance information.
	InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstancesStatusRequestParams struct {
	// Query by instance ID(s). For example, instance ID: `ins-xxxxxxxx`. For the specific format, refer to section `Ids.N` of the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You can query up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest
	
	// Query by instance ID(s). For example, instance ID: `ins-xxxxxxxx`. For the specific format, refer to section `Ids.N` of the API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You can query up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesStatusResponseParams struct {
	// Number of instance states meeting the filtering conditions.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// [Instance status](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) list.
	InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet,omitempty" name:"InstanceStatusSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesStatusResponseParams `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetChargeTypeConfigsRequestParams struct {

}

type DescribeInternetChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInternetChargeTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetChargeTypeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetChargeTypeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetChargeTypeConfigsResponseParams struct {
	// List of network billing methods.
	InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet,omitempty" name:"InternetChargeTypeConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInternetChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetChargeTypeConfigsResponseParams `json:"Response"`
}

func (r *DescribeInternetChargeTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetChargeTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyPairsRequestParams struct {
	// Key pair ID(s) in the format of `skey-11112222`. This API supports using multiple IDs as filters at the same time. For more information on the format of this parameter, see the `id.N` section in [API Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You cannot specify `KeyIds` and `Filters` at the same time. You can log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the key pair IDs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Filters
	// <li> `project-id` - Integer - Optional - Filter by project ID. To view the list of project IDs, you can go to [Project Management](https://console.cloud.tencent.com/project), or call the `DescribeProject` API and look for `projectId` in the response.</li>
	// <li> `key-name` - String - Optional - Filter by key pair name.</li>You cannot specify `KeyIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding sections in API [Introduction](https://intl.cloud.tencent.com/document/product/377). Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Key pair ID(s) in the format of `skey-11112222`. This API supports using multiple IDs as filters at the same time. For more information on the format of this parameter, see the `id.N` section in [API Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1). You cannot specify `KeyIds` and `Filters` at the same time. You can log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the key pair IDs.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Filters
	// <li> `project-id` - Integer - Optional - Filter by project ID. To view the list of project IDs, you can go to [Project Management](https://console.cloud.tencent.com/project), or call the `DescribeProject` API and look for `projectId` in the response.</li>
	// <li> `key-name` - String - Optional - Filter by key pair name.</li>You cannot specify `KeyIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset; default value: 0. For more information on `Offset`, see the corresponding sections in API [Introduction](https://intl.cloud.tencent.com/document/product/377). Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results returned; default value: 20; maximum: 100. For more information on `Limit`, see the corresponding section in API [Introduction](https://intl.cloud.tencent.com/document/product/377). 
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyPairsResponseParams struct {
	// Number of key pairs meeting the filtering conditions.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Detailed information on key pairs.
	KeyPairSet []*KeyPair `json:"KeyPairSet,omitempty" name:"KeyPairSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeLaunchTemplateVersionsRequestParams struct {
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// List of instance launch templates.
	LaunchTemplateVersions []*uint64 `json:"LaunchTemplateVersions,omitempty" name:"LaunchTemplateVersions"`

	// The minimum version number specified, which defaults to 0.
	MinVersion *uint64 `json:"MinVersion,omitempty" name:"MinVersion"`

	// The maximum version number specified, which defaults to 30.
	MaxVersion *uint64 `json:"MaxVersion,omitempty" name:"MaxVersion"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Specify whether to query the default version. This parameter and `LaunchTemplateVersions` cannot be specified at the same time.
	DefaultVersion *bool `json:"DefaultVersion,omitempty" name:"DefaultVersion"`
}

type DescribeLaunchTemplateVersionsRequest struct {
	*tchttp.BaseRequest
	
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// List of instance launch templates.
	LaunchTemplateVersions []*uint64 `json:"LaunchTemplateVersions,omitempty" name:"LaunchTemplateVersions"`

	// The minimum version number specified, which defaults to 0.
	MinVersion *uint64 `json:"MinVersion,omitempty" name:"MinVersion"`

	// The maximum version number specified, which defaults to 30.
	MaxVersion *uint64 `json:"MaxVersion,omitempty" name:"MaxVersion"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Specify whether to query the default version. This parameter and `LaunchTemplateVersions` cannot be specified at the same time.
	DefaultVersion *bool `json:"DefaultVersion,omitempty" name:"DefaultVersion"`
}

func (r *DescribeLaunchTemplateVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplateVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	delete(f, "LaunchTemplateVersions")
	delete(f, "MinVersion")
	delete(f, "MaxVersion")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DefaultVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaunchTemplateVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplateVersionsResponseParams struct {
	// Total number of instance launch templates.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Set of instance launch template versions.
	LaunchTemplateVersionSet []*LaunchTemplateVersionInfo `json:"LaunchTemplateVersionSet,omitempty" name:"LaunchTemplateVersionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLaunchTemplateVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaunchTemplateVersionsResponseParams `json:"Response"`
}

func (r *DescribeLaunchTemplateVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplatesRequestParams struct {
	// Instance launch template ID. ID of one or more instance launch templates. If not specified, all templates of the user will be displayed.
	LaunchTemplateIds []*string `json:"LaunchTemplateIds,omitempty" name:"LaunchTemplateIds"`

	// <p style="padding-left: 30px;">Filter by [<strong>LaunchTemplateName</strong>].</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// The maximum number of `Filters` in each request is 10. The upper limit for `Filter.Values` is 5. This parameter cannot specify `LaunchTemplateIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeLaunchTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Instance launch template ID. ID of one or more instance launch templates. If not specified, all templates of the user will be displayed.
	LaunchTemplateIds []*string `json:"LaunchTemplateIds,omitempty" name:"LaunchTemplateIds"`

	// <p style="padding-left: 30px;">Filter by [<strong>LaunchTemplateName</strong>].</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Optional</p>
	// The maximum number of `Filters` in each request is 10. The upper limit for `Filter.Values` is 5. This parameter cannot specify `LaunchTemplateIds` and `Filters` at the same time.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLaunchTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaunchTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchTemplatesResponseParams struct {
	// Number of eligible instance templates.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of instance details
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LaunchTemplateSet []*LaunchTemplateInfo `json:"LaunchTemplateSet,omitempty" name:"LaunchTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLaunchTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaunchTemplatesResponseParams `json:"Response"`
}

func (r *DescribeLaunchTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchTemplatesResponse) FromJsonString(s string) error {
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
	// Number of regions
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of regions
	RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeReservedInstancesConfigInfosRequestParams struct {
	// zone
	// Filters by the availability zones in which the reserved instance can be purchased, such as `ap-guangzhou-1`.
	// Type: String
	// Required: no
	// Valid values: list of regions/availability zones
	// 
	// product-description
	// Filters by the platform description (operating system) of the reserved instance, such as `linux`.
	// Type: String
	// Required: no
	// Valid value: linux
	// 
	// duration
	// Filters by the **validity** of the reserved instance, which is the purchased usage period. For example, `31536000`.
	// Type: Integer
	// Unit: second
	// Required: no
	// Valid value: 31536000 (1 year)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeReservedInstancesConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// zone
	// Filters by the availability zones in which the reserved instance can be purchased, such as `ap-guangzhou-1`.
	// Type: String
	// Required: no
	// Valid values: list of regions/availability zones
	// 
	// product-description
	// Filters by the platform description (operating system) of the reserved instance, such as `linux`.
	// Type: String
	// Required: no
	// Valid value: linux
	// 
	// duration
	// Filters by the **validity** of the reserved instance, which is the purchased usage period. For example, `31536000`.
	// Type: Integer
	// Unit: second
	// Required: no
	// Valid value: 31536000 (1 year)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeReservedInstancesConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedInstancesConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesConfigInfosResponseParams struct {
	// Static configurations of the reserved instance.
	ReservedInstanceConfigInfos []*ReservedInstanceConfigInfoItem `json:"ReservedInstanceConfigInfos,omitempty" name:"ReservedInstanceConfigInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReservedInstancesConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedInstancesConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeReservedInstancesConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesOfferingsRequestParams struct {
	// Dry run. Default value: false.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The maximum duration as a filter, 
	// in seconds.
	// Default value: 94608000.
	MaxDuration *int64 `json:"MaxDuration,omitempty" name:"MaxDuration"`

	// The minimum duration as a filter, 
	// in seconds.
	// Default value: 2592000.
	MinDuration *int64 `json:"MinDuration,omitempty" name:"MinDuration"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">Filters by the <strong>availability zones</strong> in which the Reserved Instances can be purchased, such as ap-guangzhou-1.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid values: please see <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Availability Zones</a></p>
	// <li><strong>duration</strong></li>
	// <p style="padding-left: 30px;">Filters by the <strong>duration</strong> of the Reserved Instance, in seconds. For example, 31536000.</p><p style="padding-left: 30px;">Type: Integer</p><p style="padding-left: 30px;">Unit: second</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid values: 31536000 (1 year) | 94608000 (3 years)</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">Filters by <strong>type of the Reserved Instance</strong>, such as `S3.MEDIUM4`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid values: please see <a href="https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1">Instance Types</a></p>
	// <li><strong>offering-type</strong></li>
	// <p style="padding-left: 30px;">Filters by **<strong>payment term</strong>**, such as `All Upfront`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid value: All Upfront</p>
	// <li><strong>product-description</strong></li>
	// <p style="padding-left: 30px;">Filters by the <strong>platform description</strong> (operating system) of the Reserved Instance, such as `linux`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid value: linux</p>
	// <li><strong>reserved-instances-offering-id</strong></li>
	// <p style="padding-left: 30px;">Filters by <strong>Reserved Instance ID</strong>, in the form of 650c138f-ae7e-4750-952a-96841d6e9fc1.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p>
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeReservedInstancesOfferingsRequest struct {
	*tchttp.BaseRequest
	
	// Dry run. Default value: false.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// The offset. Default value: 0. For more information on `Offset`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100. For more information on `Limit`, see the relevant sections in API [Introduction](https://intl.cloud.tencent.com/document/api/213/15688?from_cn_redirect=1).
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The maximum duration as a filter, 
	// in seconds.
	// Default value: 94608000.
	MaxDuration *int64 `json:"MaxDuration,omitempty" name:"MaxDuration"`

	// The minimum duration as a filter, 
	// in seconds.
	// Default value: 2592000.
	MinDuration *int64 `json:"MinDuration,omitempty" name:"MinDuration"`

	// <li><strong>zone</strong></li>
	// <p style="padding-left: 30px;">Filters by the <strong>availability zones</strong> in which the Reserved Instances can be purchased, such as ap-guangzhou-1.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid values: please see <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Availability Zones</a></p>
	// <li><strong>duration</strong></li>
	// <p style="padding-left: 30px;">Filters by the <strong>duration</strong> of the Reserved Instance, in seconds. For example, 31536000.</p><p style="padding-left: 30px;">Type: Integer</p><p style="padding-left: 30px;">Unit: second</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid values: 31536000 (1 year) | 94608000 (3 years)</p>
	// <li><strong>instance-type</strong></li>
	// <p style="padding-left: 30px;">Filters by <strong>type of the Reserved Instance</strong>, such as `S3.MEDIUM4`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid values: please see <a href="https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1">Instance Types</a></p>
	// <li><strong>offering-type</strong></li>
	// <p style="padding-left: 30px;">Filters by **<strong>payment term</strong>**, such as `All Upfront`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid value: All Upfront</p>
	// <li><strong>product-description</strong></li>
	// <p style="padding-left: 30px;">Filters by the <strong>platform description</strong> (operating system) of the Reserved Instance, such as `linux`.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p><p style="padding-left: 30px;">Valid value: linux</p>
	// <li><strong>reserved-instances-offering-id</strong></li>
	// <p style="padding-left: 30px;">Filters by <strong>Reserved Instance ID</strong>, in the form of 650c138f-ae7e-4750-952a-96841d6e9fc1.</p><p style="padding-left: 30px;">Type: String</p><p style="padding-left: 30px;">Required: no</p>
	// Each request can have up to 10 `Filters` and 5 `Filter.Values`.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeReservedInstancesOfferingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesOfferingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DryRun")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MaxDuration")
	delete(f, "MinDuration")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedInstancesOfferingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesOfferingsResponseParams struct {
	// The number of Reserved Instances that meet the condition.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of Reserved Instances that meet the condition.
	ReservedInstancesOfferingsSet []*ReservedInstancesOffering `json:"ReservedInstancesOfferingsSet,omitempty" name:"ReservedInstancesOfferingsSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReservedInstancesOfferingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedInstancesOfferingsResponseParams `json:"Response"`
}

func (r *DescribeReservedInstancesOfferingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesOfferingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneInstanceConfigInfosRequestParams struct {
	// <li> instance-charge-type-String-required: no-(filter) billing mode of instances. (POSTPAID_BY_HOUR: pay-as-you-go billing by hour | SPOTPAID: spot billing, which is suitable for a [spot instance] (https://intl.cloud.Tencent.com/document/product/213/17817) | CDHPAID: CDH billing, that is, billing only for CDH, but not for the instances on CDH. )  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// <li> instance-charge-type-String-required: no-(filter) billing mode of instances. (POSTPAID_BY_HOUR: pay-as-you-go billing by hour | SPOTPAID: spot billing, which is suitable for a [spot instance] (https://intl.cloud.Tencent.com/document/product/213/17817) | CDHPAID: CDH billing, that is, billing only for CDH, but not for the instances on CDH. )  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneInstanceConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneInstanceConfigInfosResponseParams struct {
	// List of model configurations for the availability zone.
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneInstanceConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {

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

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of availability zones.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of availability zones.
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DisassociateInstancesKeyPairsRequestParams struct {
	// Instance ID(s). The maximum number of instances in each request is 100. <br><br>You can obtain the available instance IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the instance IDs. <br><li>Call [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of key pair IDs. The maximum number of key pairs in each request is 100. The key pair ID is in the format of `skey-11112222`. <br><br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before disassociating a key pair from it. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). The maximum number of instances in each request is 100. <br><br>You can obtain the available instance IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/index) to query the instance IDs. <br><li>Call [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// List of key pair IDs. The maximum number of key pairs in each request is 100. The key pair ID is in the format of `skey-11112222`. <br><br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before disassociating a key pair from it. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
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
	delete(f, "InstanceIds")
	delete(f, "KeyIds")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// ID of the security group to be disassociated, such as `sg-efil73jd`. Only one security group can be disassociated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID(s) of the instance(s) to be disassociated,such as `ins-lesecurk`. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// ID of the security group to be disassociated, such as `sg-efil73jd`. Only one security group can be disassociated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// ID(s) of the instance(s) to be disassociated,such as `ins-lesecurk`. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Creation time of a spread placement group.
	// Note: This field may return null, indicating that no valid value is found.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type EnhancedService struct {
	// Enables cloud security service. If this parameter is not specified, the cloud security service will be enabled by default.
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// Enables cloud monitor service. If this parameter is not specified, the cloud monitor service will be enabled by default.
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`

	// Enables the TAT service. If this parameter is not specified, the TAT service will not be enabled.
	AutomationService *RunAutomationServiceEnabled `json:"AutomationService,omitempty" name:"AutomationService"`
}

// Predefined struct for user
type ExportImagesRequestParams struct {
	// COS bucket name
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// List of image IDs
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// Format of the exported image file. Valid values: `RAW`, `QCOW2`, `VHD` and `VMDK`. Default value: `RAW`.
	ExportFormat *string `json:"ExportFormat,omitempty" name:"ExportFormat"`

	// Prefix list of the name of exported files
	FileNamePrefixList []*string `json:"FileNamePrefixList,omitempty" name:"FileNamePrefixList"`

	// Whether to export only the system disk
	OnlyExportRootDisk *bool `json:"OnlyExportRootDisk,omitempty" name:"OnlyExportRootDisk"`

	// Check whether the image can be exported
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Role name (Default: `CVM_QcsRole`). Before exporting the images, make sure the role exists, and it has write permission to COS.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type ExportImagesRequest struct {
	*tchttp.BaseRequest
	
	// COS bucket name
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// List of image IDs
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// Format of the exported image file. Valid values: `RAW`, `QCOW2`, `VHD` and `VMDK`. Default value: `RAW`.
	ExportFormat *string `json:"ExportFormat,omitempty" name:"ExportFormat"`

	// Prefix list of the name of exported files
	FileNamePrefixList []*string `json:"FileNamePrefixList,omitempty" name:"FileNamePrefixList"`

	// Whether to export only the system disk
	OnlyExportRootDisk *bool `json:"OnlyExportRootDisk,omitempty" name:"OnlyExportRootDisk"`

	// Check whether the image can be exported
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// Role name (Default: `CVM_QcsRole`). Before exporting the images, make sure the role exists, and it has write permission to COS.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *ExportImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketName")
	delete(f, "ImageIds")
	delete(f, "ExportFormat")
	delete(f, "FileNamePrefixList")
	delete(f, "OnlyExportRootDisk")
	delete(f, "DryRun")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportImagesResponseParams struct {
	// ID of the image export task
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportImagesResponse struct {
	*tchttp.BaseResponse
	Response *ExportImagesResponseParams `json:"Response"`
}

func (r *ExportImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Externals struct {
	// Release address
	// Note: This field may return null, indicating that no valid value is found.
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// Not supported network. Value: <br><li>BASIC: classic network<br><li>VPC1.0: VPC1.0
	// Note: This field may return null, indicating that no valid value was found.
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty" name:"UnsupportNetworks"`

	// Attributes of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty" name:"StorageBlockAttr"`
}

type Filter struct {
	// Filters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GPUInfo struct {
	// Number of GPUs. 
	// Note: this field may return `null`, indicating that no valid value can be found.
	GPUCount *float64 `json:"GPUCount,omitempty" name:"GPUCount"`

	// GPU address
	// Note: this field may return `null`, indicating that no valid value can be found.
	GPUId []*string `json:"GPUId,omitempty" name:"GPUId"`

	// GPU type of the instance.
	// Note: this field may return `null`, indicating that no valid value can be found.
	GPUType *string `json:"GPUType,omitempty" name:"GPUType"`
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
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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

	// Total memory size of the CDH instance (unit: GiB)
	MemTotal *float64 `json:"MemTotal,omitempty" name:"MemTotal"`

	// Available memory size of the CDH instance (unit: GiB)
	MemAvailable *float64 `json:"MemAvailable,omitempty" name:"MemAvailable"`

	// Total disk size of the CDH instance (unit: GiB)
	DiskTotal *uint64 `json:"DiskTotal,omitempty" name:"DiskTotal"`

	// Available disk size of the CDH instance (unit: GiB)
	DiskAvailable *uint64 `json:"DiskAvailable,omitempty" name:"DiskAvailable"`

	// Disk type of the CDH instance
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// Total number of GPU cards in the CDH instance
	GpuTotal *uint64 `json:"GpuTotal,omitempty" name:"GpuTotal"`

	// Number of available GPU cards in the CDH instance
	GpuAvailable *uint64 `json:"GpuAvailable,omitempty" name:"GpuAvailable"`
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
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`

	// The list of tags bound to the image.
	// Note: This field may return `null`, indicating that no valid value was found.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Image license type
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`
}

type ImageOsList struct {
	// Supported Windows OS
	// Note: This field may return null, indicating that no valid value is found.
	Windows []*string `json:"Windows,omitempty" name:"Windows"`

	// Supported Linux OS
	// Note: This field may return null, indicating that no valid value is found.
	Linux []*string `json:"Linux,omitempty" name:"Linux"`
}

// Predefined struct for user
type ImportImageRequestParams struct {
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

	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// The license type used to activate the OS after importing an image.
	// Valid values:
	// `TencentCloud`: Tencent Cloud official license
	// `BYOL`: Bring Your Own License
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`
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

	// Tag description list. This parameter is used to bind a tag to a custom image.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// The license type used to activate the OS after importing an image.
	// Valid values:
	// `TencentCloud`: Tencent Cloud official license
	// `BYOL`: Bring Your Own License
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Architecture")
	delete(f, "OsType")
	delete(f, "OsVersion")
	delete(f, "ImageUrl")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	delete(f, "DryRun")
	delete(f, "Force")
	delete(f, "TagSpecification")
	delete(f, "LicenseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *ImportImageResponseParams `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyPairRequestParams struct {
	// Key pair name, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// The project ID to which the key pair belongs after it is created. <br><br>You can obtain the project ID in the following ways: <br><li>Check the project list in the [Project management](https://console.cloud.tencent.com/project) page.<br><li>Call the `DescribeProject` API and view the `projectId` in the response.
	// 
	// If you want to use the default project, specify 0 for the parameter.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Content of the public key in the key pair in the `OpenSSH RSA` format.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Tag description list. This parameter is used to bind a tag to a key pair.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// Key pair name, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// The project ID to which the key pair belongs after it is created. <br><br>You can obtain the project ID in the following ways: <br><li>Check the project list in the [Project management](https://console.cloud.tencent.com/project) page.<br><li>Call the `DescribeProject` API and view the `projectId` in the response.
	// 
	// If you want to use the default project, specify 0 for the parameter.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Content of the public key in the key pair in the `OpenSSH RSA` format.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Tag description list. This parameter is used to bind a tag to a key pair.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
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
	delete(f, "ProjectId")
	delete(f, "PublicKey")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyPairResponseParams struct {
	// Key pair ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type InquirePricePurchaseReservedInstancesOfferingRequestParams struct {
	// The number of the reserved instances you are purchasing.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// The ID of the reserved instance offering.
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitempty" name:"ReservedInstancesOfferingId"`

	// Dry run.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.<br>For more information, see Ensuring Idempotency.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Reserved instance name.<br><li>The RI name defaults to “unnamed” if this parameter is left empty.</li><li>You can enter any name within 60 characters (including the pattern string).</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" name:"ReservedInstanceName"`
}

type InquirePricePurchaseReservedInstancesOfferingRequest struct {
	*tchttp.BaseRequest
	
	// The number of the reserved instances you are purchasing.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// The ID of the reserved instance offering.
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitempty" name:"ReservedInstancesOfferingId"`

	// Dry run.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.<br>For more information, see Ensuring Idempotency.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Reserved instance name.<br><li>The RI name defaults to “unnamed” if this parameter is left empty.</li><li>You can enter any name within 60 characters (including the pattern string).</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" name:"ReservedInstanceName"`
}

func (r *InquirePricePurchaseReservedInstancesOfferingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePricePurchaseReservedInstancesOfferingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceCount")
	delete(f, "ReservedInstancesOfferingId")
	delete(f, "DryRun")
	delete(f, "ClientToken")
	delete(f, "ReservedInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePricePurchaseReservedInstancesOfferingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePricePurchaseReservedInstancesOfferingResponseParams struct {
	// Price of the reserved instance with specified configuration.
	Price *ReservedInstancePrice `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquirePricePurchaseReservedInstancesOfferingResponse struct {
	*tchttp.BaseResponse
	Response *InquirePricePurchaseReservedInstancesOfferingResponseParams `json:"Response"`
}

func (r *InquirePricePurchaseReservedInstancesOfferingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePricePurchaseReservedInstancesOfferingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstanceRequestParams struct {
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For IDs of `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the information; for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Configuration of the system disk of the instance. For instances with a cloud disk as the system disk, you can expand the system disk by using this parameter to specify the new capacity after reinstallation. If the parameter is not specified, the system disk capacity remains unchanged by default. You can only expand the capacity of the system disk; reducing its capacity is not supported. When reinstalling the system, you can only modify the capacity of the system disk, not the type.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Monitor and Cloud Security. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

type InquiryPriceResetInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// [Image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are four types of images: <br/><li>Public images </li><li>Custom images </li><li>Shared images </li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways: <br/><li>For IDs of `public images`, `custom images`, and `shared images`, log in to the [console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE) to query the information; for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list). </li><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "LoginSettings")
	delete(f, "EnhancedService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstanceResponseParams struct {
	// Price of reinstalling the instance with the specified configuration.
	Price *Price `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesInternetMaxBandwidthRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100. When changing the bandwidth of instances with `BANDWIDTH_PREPAID` or `BANDWIDTH_POSTPAID_BY_HOUR` as the network billing method, you can only specify one instance at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Configuration of public network egress bandwidth. The maximum bandwidth varies among different models. For more information, see the documentation on bandwidth limits. Currently only the `InternetMaxBandwidthOut` parameter is supported.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Date from which the new bandwidth takes effect. Format: `YYYY-MM-DD`, such as `2016-10-30`. The starting date cannot be earlier than the current date. If the starting date is the current date, the new bandwidth takes effect immediately. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Date until which the new bandwidth is effective. Format: `YYYY-MM-DD`, such as `2016-10-30`. The validity period of the new bandwidth covers the end date. The end date cannot be later than the expiration date of a prepaid instance. You can query the expiration time of an instance by calling [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728) and looking for `ExpiredTime` in the response. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100. When changing the bandwidth of instances with `BANDWIDTH_PREPAID` or `BANDWIDTH_POSTPAID_BY_HOUR` as the network billing method, you can only specify one instance at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InternetAccessible")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetInstancesInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesInternetMaxBandwidthResponseParams struct {
	// Price of the new bandwidth
	Price *Price `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetInstancesInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesTypeRequestParams struct {
	// Instance ID(s). You can obtain the instance IDs from the value of `InstanceId` returned by the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API. The maximum number of instances in each request is 1.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance model. Resources vary with the instance model. Specific values can be found in the tables of [Instance Types] (https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1) or in the latest specifications via the [DescribeInstanceTypeConfigs] (https://intl.cloud.tencent.com/document/product/213/15749?from_cn_redirect=1) API.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type InquiryPriceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). You can obtain the instance IDs from the value of `InstanceId` returned by the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API. The maximum number of instances in each request is 1.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance model. Resources vary with the instance model. Specific values can be found in the tables of [Instance Types] (https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1) or in the latest specifications via the [DescribeInstanceTypeConfigs] (https://intl.cloud.tencent.com/document/product/213/15749?from_cn_redirect=1) API.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *InquiryPriceResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetInstancesTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetInstancesTypeResponseParams struct {
	// Price of the instance using the specified model
	Price *Price `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetInstancesTypeResponseParams `json:"Response"`
}

func (r *InquiryPriceResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeInstanceDisksRequestParams struct {
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The configuration of data disks to be expanded. Currently, you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#DataDisk) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic. Data disk capacity unit: GB; minimum increment: 10 GB. For more information about selecting a data disk type, see the product overview on cloud disks. Available data disk types are subject to the instance type (`InstanceType`). In addition, the maximum capacity allowed for expansion varies by data disk type.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type InquiryPriceResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// The configuration of data disks to be expanded. Currently, you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#DataDisk) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic. Data disk capacity unit: GB; minimum increment: 10 GB. For more information about selecting a data disk type, see the product overview on cloud disks. Available data disk types are subject to the instance type (`InstanceType`). In addition, the maximum capacity allowed for expansion varies by data disk type.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *InquiryPriceResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeInstanceDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DataDisks")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResizeInstanceDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResizeInstanceDisksResponseParams struct {
	// Price of the disks after being expanded to the specified configurations
	Price *Price `json:"Price,omitempty" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResizeInstanceDisksResponseParams `json:"Response"`
}

func (r *InquiryPriceResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResizeInstanceDisksResponse) FromJsonString(s string) error {
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

	// Information of the instance data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// List of private IPs of the instance's primary ENI.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// List of public IPs of the instance's primary ENI.
	// Note: This field may return null, indicating that no valid value is found.
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

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
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Login settings of the instance. Currently only the key associated with the instance is returned.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Instance state. Valid values: <br><li>PENDING: creating <br></li><li>LAUNCH_FAILED: creation failed <br></li><li>RUNNING: running <br></li><li>STOPPED: shut down <br></li><li>STARTING: starting <br></li><li>STOPPING: shutting down <br></li><li>REBOOTING: rebooting <br></li><li>SHUTDOWN: shut down and to be terminated <br></li><li>TERMINATING: terminating. <br></li>
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// List of tags associated with the instance.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Instance billing method after shutdown.
	// Valid values: <br><li>KEEP_CHARGING: billing continues after shutdown <br><li>STOP_CHARGING: billing stops after shutdown <li>NOT_APPLICABLE: the instance is not shut down or stopping billing after shutdown is not applicable to the instance. <br>
	StopChargingMode *string `json:"StopChargingMode,omitempty" name:"StopChargingMode"`

	// Globally unique ID of the instance.
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// Last operation of the instance, such as StopInstances or ResetInstance.
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// The latest operation status of the instance. Valid values:<br><li>SUCCESS: operation succeeded<br><li>OPERATING: operation in progress<br><li>FAILED: operation failed
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// Unique request ID for the last operation of the instance.
	LatestOperationRequestId *string `json:"LatestOperationRequestId,omitempty" name:"LatestOperationRequestId"`

	// ID of a spread placement group.
	// Note: this field may return null, indicating that no valid value is obtained.
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// IPv6 address of the instance.
	// Note: this field may return null, indicating that no valid value is obtained.
	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`

	// CAM role name.
	// Note: this field may return null, indicating that no valid value is obtained.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// HPC cluster ID.
	// Note: this field may return null, indicating that no valid value was found.
	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`

	// IP list of HPC cluster.
	// Note: this field may return null, indicating that no valid value was found.
	RdmaIpAddresses []*string `json:"RdmaIpAddresses,omitempty" name:"RdmaIpAddresses"`

	// The isolation status of the instance. Valid values:<br><li>`ARREAR`: isolated due to overdue payment;<br></li><li>`EXPIRE`: isolated upon expiration;<br></li><li>`MANMADE`: isolated after manual returning;<br></li><li>`NOTISOLATED`: not isolated<br></li>
	// Note: this field may return null, indicating that no valid value was found.
	IsolatedSource *string `json:"IsolatedSource,omitempty" name:"IsolatedSource"`

	// GPU information. This field is only returned for GPU instances.
	// Note: this field may return null, indicating that no valid value was found.
	GPUInfo *GPUInfo `json:"GPUInfo,omitempty" name:"GPUInfo"`

	// Instance OS license type. Default value: `TencentCloud`
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`
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

type InstanceTypeQuotaItem struct {
	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance billing plan. Valid values: <br><li>POSTPAID_BY_HOUR: pay after use. You are billed for your traffic by the hour.<br><li>`CDHPAID`: [`CDH`](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) billing plan. Applicable to `CDH` only, not the instances on the host.
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
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`

	// Whether an instance model is available. Valid values: <br><li>SELL: available <br><li>SOLD_OUT: sold out
	Status *string `json:"Status,omitempty" name:"Status"`

	// Price of an instance model.
	Price *ItemPrice `json:"Price,omitempty" name:"Price"`

	// Details of out-of-stock items
	// Note: this field may return null, indicating that no valid value is obtained.
	SoldOutReason *string `json:"SoldOutReason,omitempty" name:"SoldOutReason"`

	// Private network bandwidth, in Gbps.
	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitempty" name:"InstanceBandwidth"`

	// The max packet sending and receiving capability (in 10k PPS).
	InstancePps *int64 `json:"InstancePps,omitempty" name:"InstancePps"`

	// Number of local storage blocks.
	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitempty" name:"StorageBlockAmount"`

	// CPU type.
	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`

	// Number of GPUs of the instance.
	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`

	// Number of FPGAs of the instance.
	Fpga *int64 `json:"Fpga,omitempty" name:"Fpga"`

	// Descriptive information of the instance.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type InternetAccessible struct {
	// Network connection billing plan. Valid value: <br><li>TRAFFIC_POSTPAID_BY_HOUR: pay after use. You are billed for your traffic, by the hour.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// The maximum outbound bandwidth of the public network, in Mbps. The default value is 0 Mbps. The upper limit of bandwidth varies for different models. For more information, see [Purchase Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/12523?from_cn_redirect=1).
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// Whether to assign a public IP. Valid values: <br><li>TRUE: Assign a public IP <br><li>FALSE: Do not assign a public IP <br><br>If the public network bandwidth is greater than 0 Mbps, you can choose whether to assign a public IP; by default a public IP will be assigned. If the public network bandwidth is 0 Mbps, you will not be able to assign a public IP.
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`

	// Bandwidth package ID. To obatin the IDs, you can call [`DescribeBandwidthPackages`](https://intl.cloud.tencent.com/document/api/215/19209?from_cn_redirect=1) and look for the `BandwidthPackageId` fields in the response.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
}

type InternetChargeTypeConfig struct {
	// Network billing method.
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// Description of the network billing method.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ItemPrice struct {
	// The original unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to that the unit price for unlimited usage.
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

	// Percentage of the original price. For example, if you enter "20.0", the discounted price will be 20% of the original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// The discounted unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to that the unit price for unlimited usage.
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

	// Original 3-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceThreeYear *float64 `json:"OriginalPriceThreeYear,omitempty" name:"OriginalPriceThreeYear"`

	// Discounted 3-year upfront payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceThreeYear *float64 `json:"DiscountPriceThreeYear,omitempty" name:"DiscountPriceThreeYear"`

	// Discount for 3-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountThreeYear *float64 `json:"DiscountThreeYear,omitempty" name:"DiscountThreeYear"`

	// Original 5-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceFiveYear *float64 `json:"OriginalPriceFiveYear,omitempty" name:"OriginalPriceFiveYear"`

	// Discounted 5-year upfront payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceFiveYear *float64 `json:"DiscountPriceFiveYear,omitempty" name:"DiscountPriceFiveYear"`

	// Discount for 5-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountFiveYear *float64 `json:"DiscountFiveYear,omitempty" name:"DiscountFiveYear"`

	// Original 1-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceOneYear *float64 `json:"OriginalPriceOneYear,omitempty" name:"OriginalPriceOneYear"`

	// Discounted 1-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceOneYear *float64 `json:"DiscountPriceOneYear,omitempty" name:"DiscountPriceOneYear"`

	// Discount for 1-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountOneYear *float64 `json:"DiscountOneYear,omitempty" name:"DiscountOneYear"`
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
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`

	// Creation time, which follows the `ISO8601` standard and uses `UTC` time in the format of `YYYY-MM-DDThh:mm:ssZ`.
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// The list of tags bound to the key.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type LaunchTemplate struct {
	// Instance launch template ID. This parameter enables you to create an instance using the preset parameters in the template.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// Instance launch template version number. If specified, this parameter will be used to create a new instance launch template.
	LaunchTemplateVersion *uint64 `json:"LaunchTemplateVersion,omitempty" name:"LaunchTemplateVersion"`
}

type LaunchTemplateInfo struct {
	// Instance launch template version number.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LatestVersionNumber *uint64 `json:"LatestVersionNumber,omitempty" name:"LatestVersionNumber"`

	// Instance launch template ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// Instance launch template name.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" name:"LaunchTemplateName"`

	// Default instance launch template version number.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DefaultVersionNumber *uint64 `json:"DefaultVersionNumber,omitempty" name:"DefaultVersionNumber"`

	// Total number of versions that the instance launch template contains.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LaunchTemplateVersionCount *uint64 `json:"LaunchTemplateVersionCount,omitempty" name:"LaunchTemplateVersionCount"`

	// UIN of the user who created the template.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`

	// Creation time of the template.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type LaunchTemplateVersionData struct {
	// Location of the instance.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// Instance model.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance name.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance billing mode. Valid values: <br><li>`POSTPAID_BY_HOUR`: postpaid for pay-as-you-go instances <br><li>`CDHPAID`: billed for CDH instances, not the CVMs running on the CDHs.<br>
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Instance system disk information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Instance data disk information. This parameter only covers the data disks purchased together with the instance.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Instance bandwidth information.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Information of the VPC where the instance resides.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// `ID` of the image used to create the instance.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Login settings of the instance. Currently, only the key associated with the instance is returned.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// CAM role name.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// HPC cluster `ID`.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`

	// Number of instances purchased.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Enhanced service.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16KB.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Placement group ID. You can only specify one.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// Scheduled tasks. You can use this parameter to specify scheduled tasks for the instance. Only scheduled termination is supported.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// Market options of the instance, such as parameters related to spot instances. This parameter is required for spot instances.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// Hostname of a CVM.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// A string used to ensure the idempotency of the request.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Prepaid mode. This parameter indicates relevant parameter settings for monthly-subscribed instances.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// List of tag description. By specifying this parameter, the tag can be bound to the corresponding CVM and CBS instances at the same time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type LaunchTemplateVersionInfo struct {
	// Instance launch template version number.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LaunchTemplateVersion *uint64 `json:"LaunchTemplateVersion,omitempty" name:"LaunchTemplateVersion"`

	// Details of instance launch template versions.
	LaunchTemplateVersionData *LaunchTemplateVersionData `json:"LaunchTemplateVersionData,omitempty" name:"LaunchTemplateVersionData"`

	// Creation time of the instance launch template version.
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Instance launch template ID.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// Specifies whether it’s the default launch template version.
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" name:"IsDefaultVersion"`

	// Information of instance launch template version description.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	LaunchTemplateVersionDescription *string `json:"LaunchTemplateVersionDescription,omitempty" name:"LaunchTemplateVersionDescription"`

	// Creator account
	CreatedBy *string `json:"CreatedBy,omitempty" name:"CreatedBy"`
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

	// Whether a local disk is required during purchase. Valid values:<br><li>REQUIRED: required<br><li>OPTIONAL: optional
	Required *string `json:"Required,omitempty" name:"Required"`
}

type LoginSettings struct {
	// Login password of the instance. <br><li>For Linux instances, the password must include 8-30 characters, and contain at least two of the following character sets: [a-z], [A-Z], [0-9] and [()\`~!@#$%^&*-+=|{}[]:;',.?/]. <br><li>For Windows instances, the password must include 12-30 characters, and contain at least three of the following character sets: [a-z], [A-Z], [0-9] and [()\`~!@#$%^&*-+=|{}[]:;',.?/]. <br><br>If it's not specified, the user needs to set the login password using the **Reset password** option in the CVM console or calling the API `ResetInstancesPassword` to complete the creation of the CVM instance(s).
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitempty" name:"Password"`

	// List of key IDs. After an instance is associated with a key, you can access the instance with the private key in the key pair. You can call [`DescribeKeyPairs`](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) to obtain `KeyId`. A key and password cannot be specified at the same time. Windows instances do not support keys. Currently, you can only specify one key when purchasing an instance.
	// Note: this field may return null, indicating that no valid value is obtained.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// Whether to keep the original settings of an image. You cannot specify this parameter and `Password` or `KeyIds.N` at the same time. You can specify this parameter as `TRUE` only when you create an instance using a custom image, a shared image, or an imported image. Valid values: <br><li>TRUE: keep the login settings of the image <br><li>FALSE: do not keep the login settings of the image <br><br>Default value: FALSE.
	// Note: This field may return null, indicating that no valid value is found.
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

// Predefined struct for user
type ModifyChcAttributeRequestParams struct {
	// CHC host IDs
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// CHC host name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Server type
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Valid characters: Letters, numbers, hyphens and underscores
	BmcUser *string `json:"BmcUser,omitempty" name:"BmcUser"`

	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&*-+=_|{}).
	Password *string `json:"Password,omitempty" name:"Password"`

	// BMC network security group list
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitempty" name:"BmcSecurityGroupIds"`
}

type ModifyChcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// CHC host IDs
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// CHC host name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Server type
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// Valid characters: Letters, numbers, hyphens and underscores
	BmcUser *string `json:"BmcUser,omitempty" name:"BmcUser"`

	// The password can contain 8 to 16 characters, including letters, numbers and special symbols (()`~!@#$%^&*-+=_|{}).
	Password *string `json:"Password,omitempty" name:"Password"`

	// BMC network security group list
	BmcSecurityGroupIds []*string `json:"BmcSecurityGroupIds,omitempty" name:"BmcSecurityGroupIds"`
}

func (r *ModifyChcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	delete(f, "InstanceName")
	delete(f, "DeviceType")
	delete(f, "BmcUser")
	delete(f, "Password")
	delete(f, "BmcSecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyChcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChcAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyChcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyChcAttributeResponseParams `json:"Response"`
}

func (r *ModifyChcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisasterRecoverGroupAttributeRequestParams struct {
	// Spread placement group ID, which can be obtained by calling the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/api/213/17810?from_cn_redirect=1) API.
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// Name of a spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyDisasterRecoverGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Spread placement group ID, which can be obtained by calling the [DescribeDisasterRecoverGroups](https://intl.cloud.tencent.com/document/api/213/17810?from_cn_redirect=1) API.
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitempty" name:"DisasterRecoverGroupId"`

	// Name of a spread placement group. The name must be 1-60 characters long and can contain both Chinese characters and English letters.
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDisasterRecoverGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisasterRecoverGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisasterRecoverGroupId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDisasterRecoverGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDisasterRecoverGroupAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDisasterRecoverGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDisasterRecoverGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyDisasterRecoverGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDisasterRecoverGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsAttributeRequestParams struct {
	// CDH instance ID(s).
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`

	// CDH instance name to be displayed. You can specify any name you like, but its length cannot exceed 60 characters.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Auto renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Project ID. You can use the `AddProject` API to create projects, and obtain the `projectId` field in the response of the `DescribeProject` API. When using the [DescribeHosts](https://intl.cloud.tencent.com/document/api/213/16474?from_cn_redirect=1) API to query instances later, you can filter the results by the project ID.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyHostsAttributeRequest struct {
	*tchttp.BaseRequest
	
	// CDH instance ID(s).
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`

	// CDH instance name to be displayed. You can specify any name you like, but its length cannot exceed 60 characters.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Auto renewal flag. Valid values: <br><li>NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically <br><li>NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically <br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically <br><br>If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// Project ID. You can use the `AddProject` API to create projects, and obtain the `projectId` field in the response of the `DescribeProject` API. When using the [DescribeHosts](https://intl.cloud.tencent.com/document/api/213/16474?from_cn_redirect=1) API to query instances later, you can filter the results by the project ID.
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyHostsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostIds")
	delete(f, "HostName")
	delete(f, "RenewFlag")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHostsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostsAttributeResponseParams `json:"Response"`
}

func (r *ModifyHostsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeRequestParams struct {
	// Image ID such as `img-gvbnzy6f`. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image).
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// New image name, which must meet the following requirements: <br> <li>No more than 20 characters. <br> <li>Must be unique.
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// New image description, which must meet the following requirement: <br> <li> No more than 60 characters.
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Image ID such as `img-gvbnzy6f`. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image).
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageAttributeResponseParams `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSharePermissionRequestParams struct {
	// Image ID such as `img-gvbnzy6f`. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image). <br>You can only specify an image in the `NORMAL` state. For more information on image states, see [here](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#image_state).
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// List of account IDs with which an image is shared. For the format of array-type parameters, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1). The account ID is different from the QQ number. You can find the account ID in [Account Information](https://console.cloud.tencent.com/developer). 
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`

	// Operations. Valid values: `SHARE`, sharing an image; `CANCEL`, cancelling an image sharing. 
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest
	
	// Image ID such as `img-gvbnzy6f`. You can obtain the image IDs in two ways: <br><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response. <br><li>Look for the information in the [Image Console](https://console.cloud.tencent.com/cvm/image). <br>You can only specify an image in the `NORMAL` state. For more information on image states, see [here](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#image_state).
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// List of account IDs with which an image is shared. For the format of array-type parameters, see [API Introduction](https://intl.cloud.tencent.com/document/api/213/568?from_cn_redirect=1). The account ID is different from the QQ number. You can find the account ID in [Account Information](https://console.cloud.tencent.com/developer). 
	AccountIds []*string `json:"AccountIds,omitempty" name:"AccountIds"`

	// Operations. Valid values: `SHARE`, sharing an image; `CANCEL`, cancelling an image sharing. 
	Permission *string `json:"Permission,omitempty" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "AccountIds")
	delete(f, "Permission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageSharePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageSharePermissionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageSharePermissionResponseParams `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The instance name, which can not exceed 60 characters
	// <dx-alert infotype="explain" title="">Either `InstanceName` or `SecurityGroups` must be specified, but they can not be both specified.</dx-alert>
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// IDs of security groups associated with the specified instance. You can associate with a security group by adding its ID, or cancel the association with a security group by removing its ID. <dx-alert infotype="explain" title="">Either `InstanceName` or `SecurityGroups` must be specified, but they cannot be both set.</dx-alert>
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// The role bound with the instance. If it is not specified, it indicates to unbind the current role of the CVM.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// Whether the termination protection is enabled. Values: <br><li>`TRUE`: enable instance protection, which means that this instance can not be deleted by an API action.<br><li>`FALSE`: do not enable the instance protection.<br><br>Default Value: `FALSE`.
	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`

	// The role type, which is used in conjunction with `CamRoleName`. The value is obtained in `RoleType` field, returning by `CAM DescribeRoleList` and `GetRole` APIs. Valid value: `user`, `system` and `service_linked`.
	// For example, when `LinkedRoleIn` is contained in `CamRoleName` (such as `TKE_QCSLinkedRoleInPrometheusService`), the returned `RoleType` of `DescribeRoleList` and `GetRoleis` is `service_linked`, and the `CamRoleType` `service_linked`.
	// When the value obtained in `RoleType` is `user` (default) or `system`, `CamRoleType` can be left empty.
	CamRoleType *string `json:"CamRoleType,omitempty" name:"CamRoleType"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// The instance name, which can not exceed 60 characters
	// <dx-alert infotype="explain" title="">Either `InstanceName` or `SecurityGroups` must be specified, but they can not be both specified.</dx-alert>
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// IDs of security groups associated with the specified instance. You can associate with a security group by adding its ID, or cancel the association with a security group by removing its ID. <dx-alert infotype="explain" title="">Either `InstanceName` or `SecurityGroups` must be specified, but they cannot be both set.</dx-alert>
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// The role bound with the instance. If it is not specified, it indicates to unbind the current role of the CVM.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// Whether the termination protection is enabled. Values: <br><li>`TRUE`: enable instance protection, which means that this instance can not be deleted by an API action.<br><li>`FALSE`: do not enable the instance protection.<br><br>Default Value: `FALSE`.
	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`

	// The role type, which is used in conjunction with `CamRoleName`. The value is obtained in `RoleType` field, returning by `CAM DescribeRoleList` and `GetRole` APIs. Valid value: `user`, `system` and `service_linked`.
	// For example, when `LinkedRoleIn` is contained in `CamRoleName` (such as `TKE_QCSLinkedRoleInPrometheusService`), the returned `RoleType` of `DescribeRoleList` and `GetRoleis` is `service_linked`, and the `CamRoleType` `service_linked`.
	// When the value obtained in `RoleType` is `user` (default) or `system`, `CamRoleType` can be left empty.
	CamRoleType *string `json:"CamRoleType,omitempty" name:"CamRoleType"`
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
	delete(f, "SecurityGroups")
	delete(f, "CamRoleName")
	delete(f, "DisableApiTermination")
	delete(f, "CamRoleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyInstancesProjectRequestParams struct {
	// Instance IDs. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. You can operate up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Project ID. You can use the API `AddProject` to create projects, and obtain the `projectId` field in the response of the `DescribeProject` API. When using the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API to query instances later, you can filter the results by the project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// Instance IDs. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. You can operate up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Project ID. You can use the API `AddProject` to create projects, and obtain the `projectId` field in the response of the `DescribeProject` API. When using the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API to query instances later, you can filter the results by the project ID.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesProjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesProjectResponseParams `json:"Response"`
}

func (r *ModifyInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesVpcAttributeRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// VPC configurations. You can use this parameter to specify the VPC ID, subnet ID, VPC IP, etc. If the specified VPC ID and subnet ID (the subnet must be in the same availability zone as the instance) are different from the VPC where the specified instance resides, the instance will be migrated to a subnet of the specified VPC. You can use `PrivateIpAddresses` to specify the VPC subnet IP. If you want to specify the subnet IP, you will need to specify a subnet IP for each of the specified instances, and each `InstanceIds` will match a `PrivateIpAddresses`. If `PrivateIpAddresses` is not specified, the VPC subnet IP will be assigned randomly.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Whether to force shut down a running instances. Default value: TRUE.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Whether to keep the host name. Default value: FALSE.
	ReserveHostName *bool `json:"ReserveHostName,omitempty" name:"ReserveHostName"`
}

type ModifyInstancesVpcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesVpcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "VirtualPrivateCloud")
	delete(f, "ForceStop")
	delete(f, "ReserveHostName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesVpcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesVpcAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesVpcAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstancesVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeyPairAttributeRequestParams struct {
	// Key pair ID in the format of `skey-xxxxxxxx`. <br><br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// New key pair name, which can contain numbers, letters, and underscores, with a maximum length of 25 characters.
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// New key pair description. You can specify any name you like, but its length cannot exceed 60 characters.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Key pair ID in the format of `skey-xxxxxxxx`. <br><br>You can obtain the available key pair IDs in two ways: <br><li>Log in to the [console](https://console.cloud.tencent.com/cvm/sshkey) to query the key pair IDs. <br><li>Call [DescribeKeyPairs](https://intl.cloud.tencent.com/document/api/213/15699?from_cn_redirect=1) and look for `KeyId` in the response.
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeyPairAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "KeyName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKeyPairAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeyPairAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKeyPairAttributeResponseParams `json:"Response"`
}

func (r *ModifyKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeyPairAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchTemplateDefaultVersionRequestParams struct {
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// The number of the version that you want to set as the default version
	DefaultVersion *int64 `json:"DefaultVersion,omitempty" name:"DefaultVersion"`
}

type ModifyLaunchTemplateDefaultVersionRequest struct {
	*tchttp.BaseRequest
	
	// The launch template ID
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" name:"LaunchTemplateId"`

	// The number of the version that you want to set as the default version
	DefaultVersion *int64 `json:"DefaultVersion,omitempty" name:"DefaultVersion"`
}

func (r *ModifyLaunchTemplateDefaultVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchTemplateDefaultVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchTemplateId")
	delete(f, "DefaultVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaunchTemplateDefaultVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchTemplateDefaultVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLaunchTemplateDefaultVersionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLaunchTemplateDefaultVersionResponseParams `json:"Response"`
}

func (r *ModifyLaunchTemplateDefaultVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchTemplateDefaultVersionResponse) FromJsonString(s string) error {
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
	OsVersions []*string `json:"OsVersions,omitempty" name:"OsVersions"`

	// Supported operating system architecture
	Architecture []*string `json:"Architecture,omitempty" name:"Architecture"`
}

type Placement struct {
	// ID of the availability zone where the instance resides. You can call the [DescribeZones](https://intl.cloud.tencent.com/document/product/213/35071) API and obtain the ID in the returned `Zone` field.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ID of the project to which the instance belongs. To obtain the project IDs, you can call [DescribeProject](https://intl.cloud.tencent.com/document/api/378/4400?from_cn_redirect=1) and look for the `projectId` fields in the response. If this parameter is not specified, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// ID list of CDHs from which the instance can be created. If you have purchased CDHs and specify this parameter, the instances you purchase will be randomly deployed on the CDHs.
	HostIds []*string `json:"HostIds,omitempty" name:"HostIds"`

	// Master host IP used to create the CVM
	HostIps []*string `json:"HostIps,omitempty" name:"HostIps"`

	// The ID of the CDH to which the instance belongs, only used as an output parameter.
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type Price struct {
	// Instance price.
	InstancePrice *ItemPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// Network price.
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitempty" name:"BandwidthPrice"`
}

// Predefined struct for user
type PurchaseReservedInstancesOfferingRequestParams struct {
	// The number of the Reserved Instance you are purchasing.
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// The ID of the Reserved Instance.
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitempty" name:"ReservedInstancesOfferingId"`

	// Dry run
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.<br>For more information, see Ensuring Idempotency.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Reserved instance name.<br><li>The RI name defaults to “unnamed” if this parameter is left empty.</li><li>You can enter any name within 60 characters (including the pattern string).</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" name:"ReservedInstanceName"`
}

type PurchaseReservedInstancesOfferingRequest struct {
	*tchttp.BaseRequest
	
	// The number of the Reserved Instance you are purchasing.
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// The ID of the Reserved Instance.
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitempty" name:"ReservedInstancesOfferingId"`

	// Dry run
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idempotency of the request cannot be guaranteed.<br>For more information, see Ensuring Idempotency.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Reserved instance name.<br><li>The RI name defaults to “unnamed” if this parameter is left empty.</li><li>You can enter any name within 60 characters (including the pattern string).</li>
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" name:"ReservedInstanceName"`
}

func (r *PurchaseReservedInstancesOfferingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurchaseReservedInstancesOfferingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceCount")
	delete(f, "ReservedInstancesOfferingId")
	delete(f, "DryRun")
	delete(f, "ClientToken")
	delete(f, "ReservedInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurchaseReservedInstancesOfferingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurchaseReservedInstancesOfferingResponseParams struct {
	// The ID of the Reserved Instance purchased.
	ReservedInstanceId *string `json:"ReservedInstanceId,omitempty" name:"ReservedInstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PurchaseReservedInstancesOfferingResponse struct {
	*tchttp.BaseResponse
	Response *PurchaseReservedInstancesOfferingResponseParams `json:"Response"`
}

func (r *PurchaseReservedInstancesOfferingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurchaseReservedInstancesOfferingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// Instance IDs. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. You can operate up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Whether to forcibly restart an instance after a normal restart fails. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default value: `FALSE`. This parameter has been disused. We recommend using `StopType` instead. Note that `ForceReboot` and `StopType` parameters cannot be specified at the same time.
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// Shutdown type. Valid values: <br><li>SOFT: soft shutdown<br><li>HARD: hard shutdown<br><li>SOFT_FIRST: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails<br><br>Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance IDs. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. You can operate up to 100 instances in each request.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Whether to forcibly restart an instance after a normal restart fails. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default value: `FALSE`. This parameter has been disused. We recommend using `StopType` instead. Note that `ForceReboot` and `StopType` parameters cannot be specified at the same time.
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// Shutdown type. Valid values: <br><li>SOFT: soft shutdown<br><li>HARD: hard shutdown<br><li>SOFT_FIRST: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails<br><br>Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`
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
	delete(f, "ForceReboot")
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Region name, such as `ap-guangzhou`
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region description, such as South China (Guangzhou)
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Whether the region is available
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

// Predefined struct for user
type RemoveChcAssistVpcRequestParams struct {
	// CHC ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`
}

type RemoveChcAssistVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`
}

func (r *RemoveChcAssistVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcAssistVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveChcAssistVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChcAssistVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveChcAssistVpcResponse struct {
	*tchttp.BaseResponse
	Response *RemoveChcAssistVpcResponseParams `json:"Response"`
}

func (r *RemoveChcAssistVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcAssistVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChcDeployVpcRequestParams struct {
	// CHC ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`
}

type RemoveChcDeployVpcRequest struct {
	*tchttp.BaseRequest
	
	// CHC ID
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`
}

func (r *RemoveChcDeployVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcDeployVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveChcDeployVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveChcDeployVpcResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveChcDeployVpcResponse struct {
	*tchttp.BaseResponse
	Response *RemoveChcDeployVpcResponseParams `json:"Response"`
}

func (r *RemoveChcDeployVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveChcDeployVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReservedInstanceConfigInfoItem struct {
	// Abbreviation name of the instance type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Full name of the instance type.
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// Priority.
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// List of instance families.
	InstanceFamilies []*ReservedInstanceFamilyItem `json:"InstanceFamilies,omitempty" name:"InstanceFamilies"`
}

type ReservedInstanceFamilyItem struct {
	// Instance family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Priority.
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// List of instance types.
	InstanceTypes []*ReservedInstanceTypeItem `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
}

type ReservedInstancePrice struct {
	// Original upfront payment, in USD.
	OriginalFixedPrice *float64 `json:"OriginalFixedPrice,omitempty" name:"OriginalFixedPrice"`

	// Discounted upfront payment, in USD.
	DiscountFixedPrice *float64 `json:"DiscountFixedPrice,omitempty" name:"DiscountFixedPrice"`

	// Original subsequent unit price, in USD/hr.
	OriginalUsagePrice *float64 `json:"OriginalUsagePrice,omitempty" name:"OriginalUsagePrice"`

	// Discounted subsequent unit price, in USD/hr.
	DiscountUsagePrice *float64 `json:"DiscountUsagePrice,omitempty" name:"DiscountUsagePrice"`
}

type ReservedInstancePriceItem struct {
	// Payment method. Valid values: All Upfront, Partial Upfront, and No Upfront.
	OfferingType *string `json:"OfferingType,omitempty" name:"OfferingType"`

	// Upfront payment, in USD.
	FixedPrice *float64 `json:"FixedPrice,omitempty" name:"FixedPrice"`

	// Subsequent unit price, in USD/hr.
	UsagePrice *float64 `json:"UsagePrice,omitempty" name:"UsagePrice"`

	// The ID of the reserved instance offering.
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitempty" name:"ReservedInstancesOfferingId"`

	// The availability zone in which the reserved instance can be purchased.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The **validity** of the reserved instance in seconds, which is the purchased usage period. For example, `31536000`.
	// Unit: second
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// The operating system of the reserved instance, such as `linux`.
	// Valid value: linux.
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`
}

type ReservedInstanceTypeItem struct {
	// Instance type.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Number of CPU cores.
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// Memory size.
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Number of GPUs.
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// Number of FPGAs.
	Fpga *uint64 `json:"Fpga,omitempty" name:"Fpga"`

	// Number of local storage blocks.
	StorageBlock *uint64 `json:"StorageBlock,omitempty" name:"StorageBlock"`

	// Number of NICs.
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// Maximum bandwidth.
	MaxBandwidth *float64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// CPU frequency.
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// CPU type.
	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`

	// Packet forwarding rate.
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// Other information.
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Price information about the reserved instance.
	Prices []*ReservedInstancePriceItem `json:"Prices,omitempty" name:"Prices"`
}

type ReservedInstancesOffering struct {
	// The availability zones in which the Reserved Instance can be purchased, such as ap-guangzhou-1.
	// Valid value: <a href="https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1">Availability Zones</a>
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// The billing currency of the Reserved Instance you are purchasing. It's specified using ISO 4217 standard currency.
	// Value: USD.
	CurrencyCode *string `json:"CurrencyCode,omitempty" name:"CurrencyCode"`

	// The **validity** of the Reserved Instance in seconds, which is the purchased usage period. For example, 31536000.
	// Unit: second
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// The purchase price of the Reserved Instance, such as 4000.0.
	// Unit: this field uses the currency code specified in `currencyCode`, and only supports “USD” at this time.
	FixedPrice *float64 `json:"FixedPrice,omitempty" name:"FixedPrice"`

	// The instance model of the Reserved Instance, such as S3.MEDIUM4.
	// Valid values: please see <a href="https://intl.cloud.tencent.com/document/product/213/11518">Reserved Instance Types</a>
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The payment term of the Reserved Instance, such as **All Upfront**.
	// Valid value: All Upfront.
	OfferingType *string `json:"OfferingType,omitempty" name:"OfferingType"`

	// The ID of the Reserved Instance offering, such as 650c138f-ae7e-4750-952a-96841d6e9fc1.
	ReservedInstancesOfferingId *string `json:"ReservedInstancesOfferingId,omitempty" name:"ReservedInstancesOfferingId"`

	// The operating system of the Reserved Instance, such as **linux**.
	// Valid value: linux.
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// The hourly usage price of the Reserved Instance, such as 0.0.
	// Currently, the only supported payment mode is **All Upfront**, so the default value of `UsagePrice` is 0 USD/hr.
	// Unit: USD/hr. This field uses the currency code specified in `currencyCode`, and only supports “USD” at this time.
	UsagePrice *float64 `json:"UsagePrice,omitempty" name:"UsagePrice"`
}

// Predefined struct for user
type ResetInstanceRequestParams struct {
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Specified effective [image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are four types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways:<br/><li>for IDs of `public images`, `custom images`, and `shared images`, log in to the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE); for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list).</li><li>Call the API [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	// <br>Default value: current image.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Configurations of the system disk. For an instance whose system disk is a cloud disk, this parameter can be used to expand the system disk by specifying a new capacity after reinstallation. The system disk capacity can only be expanded but not shrunk. Reinstalling the system can only resize rather than changing the type of the system disk.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Monitor and Cloud Security. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Host name of the CVM, editable during the system reinstallation. <br><li>Periods (.) or hyphens (-) cannot be the start or end of a host name or appear consecutively in a host name.<br><li>For Windows instances, the host name must consist of 2-15 characters , including uppercase and lowercase letters, numbers, or hyphens (-). It cannot contain periods (.) or contain only numbers.<br><li>For other instances, such as Linux instances, the host name must consist of 2-60 characters, including multiple periods (.), and allows uppercase and lowercase letters, numbers, or hyphens (-) between any two periods (.).
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Specified effective [image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are four types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><li>Marketplace images </li><br/>You can obtain the available image IDs in the following ways:<br/><li>for IDs of `public images`, `custom images`, and `shared images`, log in to the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE); for IDs of `marketplace images`, go to [Cloud Marketplace](https://market.cloud.tencent.com/list).</li><li>Call the API [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) and look for `ImageId` in the response.</li>
	// <br>Default value: current image.
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// Configurations of the system disk. For an instance whose system disk is a cloud disk, this parameter can be used to expand the system disk by specifying a new capacity after reinstallation. The system disk capacity can only be expanded but not shrunk. Reinstalling the system can only resize rather than changing the type of the system disk.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Login settings of the instance. You can use this parameter to set the login method, password, and key of the instance or keep the login settings of the original image. By default, a random password will be generated and sent to you via the Message Center.
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Enhanced services. You can use this parameter to specify whether to enable services such as Cloud Monitor and Cloud Security. If this parameter is not specified, Cloud Monitor and Cloud Security will be enabled by default.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// Host name of the CVM, editable during the system reinstallation. <br><li>Periods (.) or hyphens (-) cannot be the start or end of a host name or appear consecutively in a host name.<br><li>For Windows instances, the host name must consist of 2-15 characters , including uppercase and lowercase letters, numbers, or hyphens (-). It cannot contain periods (.) or contain only numbers.<br><li>For other instances, such as Linux instances, the host name must consist of 2-60 characters, including multiple periods (.), and allows uppercase and lowercase letters, numbers, or hyphens (-) between any two periods (.).
	HostName *string `json:"HostName,omitempty" name:"HostName"`
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
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "LoginSettings")
	delete(f, "EnhancedService")
	delete(f, "HostName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ResetInstancesInternetMaxBandwidthRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100. When changing the bandwidth of instances with `BANDWIDTH_PREPAID` or `BANDWIDTH_POSTPAID_BY_HOUR` as the network billing method, you can only specify one instance at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Configuration of public network egress bandwidth. The maximum bandwidth varies among different models. For more information, see the documentation on bandwidth limits. Currently only the `InternetMaxBandwidthOut` parameter is supported.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// Date from which the new bandwidth takes effect. Format: `YYYY-MM-DD`, such as `2016-10-30`. The starting date cannot be earlier than the current date. If the starting date is the current date, the new bandwidth takes effect immediately. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// Date until which the new bandwidth is effective. Format: `YYYY-MM-DD`, such as `2016-10-30`. The validity period of the new bandwidth covers the end date. The end date cannot be later than the expiration date of a prepaid instance. You can query the expiration time of an instance by calling [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728) and looking for `ExpiredTime` in the response. This parameter is only valid for prepaid bandwidth. If you specify the parameter for bandwidth with other network billing methods, an error code will be returned.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100. When changing the bandwidth of instances with `BANDWIDTH_PREPAID` or `BANDWIDTH_POSTPAID_BY_HOUR` as the network billing method, you can only specify one instance at a time.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InternetAccessible")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesInternetMaxBandwidthResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *ResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Login password of the instance. The password requirements vary among different operating systems:
	// For a Linux instance, the password must be 8 to 30 characters in length; password with more than 12 characters is recommended. It cannot begin with "/", and must contain at least one character from three of the following categories: <br><li>Lowercase letters: [a-z]<br><li>Uppercase letters: [A-Z]<br><li>Numbers: 0-9<br><li>Special characters: ()\`\~!@#$%^&\*-+=\_|{}[]:;'<>,.?/
	// For a Windows CVM, the password must be 12 to 30 characters in length. It cannot begin with "/" or contain a username. It must contain at least one character from three of the following categories: <br><li>Lowercase letters: [a-z]<br><li>Uppercase letters: [A-Z]<br><li>Numbers: 0-9<br><li>Special characters: ()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>If the specified instances include both `Linux` and `Windows` instances, you will need to follow the password requirements for `Windows` instances.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Username of the instance operating system for which the password needs to be reset. This parameter is limited to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Login password of the instance. The password requirements vary among different operating systems:
	// For a Linux instance, the password must be 8 to 30 characters in length; password with more than 12 characters is recommended. It cannot begin with "/", and must contain at least one character from three of the following categories: <br><li>Lowercase letters: [a-z]<br><li>Uppercase letters: [A-Z]<br><li>Numbers: 0-9<br><li>Special characters: ()\`\~!@#$%^&\*-+=\_|{}[]:;'<>,.?/
	// For a Windows CVM, the password must be 12 to 30 characters in length. It cannot begin with "/" or contain a username. It must contain at least one character from three of the following categories: <br><li>Lowercase letters: [a-z]<br><li>Uppercase letters: [A-Z]<br><li>Numbers: 0-9<br><li>Special characters: ()\`\~!@#$%^&\*-+=\_|{}[]:;' <>,.?/<br><li>If the specified instances include both `Linux` and `Windows` instances, you will need to follow the password requirements for `Windows` instances.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Username of the instance operating system for which the password needs to be reset. This parameter is limited to 64 characters.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
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
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ResetInstancesTypeRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API and find the value `InstanceId` in the response. The maximum number of instances in each request is 1.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance model. Different resource specifications are specified for different models. For specific values, call [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) to get the latest specification list or refer to [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1).
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Forced shutdown of a running instances. We recommend you firstly try to shut down a running instance manually. Valid values: <br><li>TRUE: forced shutdown of an instance after a normal shutdown fails.<br><li>FALSE: no forced shutdown of an instance after a normal shutdown fails.<br><br>Default value: FALSE.<br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force a CVM to shut off if the normal shutdown fails.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API and find the value `InstanceId` in the response. The maximum number of instances in each request is 1.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Instance model. Different resource specifications are specified for different models. For specific values, call [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) to get the latest specification list or refer to [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1).
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Forced shutdown of a running instances. We recommend you firstly try to shut down a running instance manually. Valid values: <br><li>TRUE: forced shutdown of an instance after a normal shutdown fails.<br><li>FALSE: no forced shutdown of an instance after a normal shutdown fails.<br><br>Default value: FALSE.<br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force a CVM to shut off if the normal shutdown fails.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesTypeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesTypeResponseParams `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeInstanceDisksRequestParams struct {
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Configuration of data disks to be expanded. Currently you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is not elastic. Data disk capacity unit: GB; minimum increment: 10 GB. For more information on selecting the data disk type, see the [product overview on cloud disks](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1). Available data disk types are subject to the instance type (`InstanceType`). In addition, the maximum capacity allowed for expansion varies by data disk type.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Configuration of the system disk to be expanded. Only cloud disks are supported.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Whether the cloud disk is expanded online.
	ResizeOnline *bool `json:"ResizeOnline,omitempty" name:"ResizeOnline"`
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response.
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Configuration of data disks to be expanded. Currently you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is not elastic. Data disk capacity unit: GB; minimum increment: 10 GB. For more information on selecting the data disk type, see the [product overview on cloud disks](https://intl.cloud.tencent.com/document/product/362/2353?from_cn_redirect=1). Available data disk types are subject to the instance type (`InstanceType`). In addition, the maximum capacity allowed for expansion varies by data disk type.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Whether to force shut down a running instances. It is recommended to manually shut down a running instance before resetting the user password. Valid values: <br><li>TRUE: force shut down an instance after a normal shutdown fails. <br><li>FALSE: do not force shut down an instance after a normal shutdown fails. <br><br>Default value: FALSE. <br><br>A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be shut down normally.
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Configuration of the system disk to be expanded. Only cloud disks are supported.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// Whether the cloud disk is expanded online.
	ResizeOnline *bool `json:"ResizeOnline,omitempty" name:"ResizeOnline"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DataDisks")
	delete(f, "ForceStop")
	delete(f, "SystemDisk")
	delete(f, "ResizeOnline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeInstanceDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeInstanceDisksResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *ResizeInstanceDisksResponseParams `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunAutomationServiceEnabled struct {
	// Whether to enable the TAT service. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default: `FALSE`.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

// Predefined struct for user
type RunInstancesRequestParams struct {
	// Instance [billing type](https://intl.cloud.tencent.com/document/product/213/2180?from_cn_redirect=1). <br><li>`POSTPAID_BY_HOUR`: Hourly-based pay-as-you-go <br><li>`CDHPAID`: Dedicated CVM (associated with a dedicated host. Resource usage of the dedicated host is free of charge.) <br><li>`SPOTPAID`: [Spot instance](https://intl.cloud.tencent.com/document/product/213/17817)<br>Default value: `POSTPAID_BY_HOUR`.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// Location of the instance. You can use this parameter to specify the availability zone, project, and CDH (for dedicated CVMs).
	//  <b>Note: `Placement` is required when `LaunchTemplate` is not specified. If both the parameters are passed in, `Placement` prevails.</b>
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The instance model. 
	// <br><li>To view specific values for `POSTPAID_BY_HOUR` instances, you can call [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) or refer to [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1). If this parameter is not specified, `S1.SMALL1` will be used by default.<br><li>For `CDHPAID` instances, the value of this parameter is in the format of `CDH_XCXG` based on the number of CPU cores and memory capacity. For example, if you want to create a CDH instance with a single-core CPU and 1 GB memory, specify this parameter as `CDH_1C1G`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The [image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><br/>To check the image ID:<br/><li>For public images, custom images, and shared images, go to the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE). </li><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1), pass in `InstanceType` to retrieve the list of images supported by the current model, and then find the `ImageId` in the response.</li>
	//  <b>Note: `ImageId` is required when `LaunchTemplate` is not specified. If both the parameters are passed in, `ImageId` prevails.</b>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default. When purchasing, you can specify 21 data disks, which can contain at most 1 LOCAL_BASIC or LOCAL_SSD data disk, and at most 20 CLOUD_BASIC, CLOUD_PREMIUM, or CLOUD_SSD data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Configuration information of VPC. This parameter is used to specify VPC ID and subnet ID, etc. If a VPC IP is specified in this parameter, it indicates the primary ENI IP of each instance. The value of parameter InstanceCount must be the same as the number of VPC IPs, which cannot be greater than 20.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// The number of instances to be purchased. Value range for pay-as-you-go instances: [1, 100]. Default value: `1`. The specified number of instances to be purchased cannot exceed the remaining quota allowed for the user. For more information on the quota, see [Quota for CVM Instances](https://intl.cloud.tencent.com/document/product/213/2664).
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instance name to be displayed. <br><li>If this parameter is not specified, "Unnamed" will be displayed by default. </li><li>If you purchase multiple instances at the same time and specify a pattern string `{R:x}`, numbers `[x, x+n-1]` will be generated, where `n` represents the number of instances purchased. For example, you specify a pattern string, `server_{R:3}`. If you only purchase 1 instance, the instance will be named `server_3`; if you purchase 2, they will be named `server_3` and `server_4`. You can specify multiple pattern strings in the format of `{R:x}`. </li><li>If you purchase multiple instances at the same time and do not specify a pattern string, the instance names will be suffixed by `1, 2...n`, where `n` represents the number of instances purchased. For example, if you purchase 2 instances and the instance name body is `server_`, the instance names will be `server_1` and `server_2`. </li><li>This parameter can contain up to 60 characters, including the pattern string.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance login settings. You can use this parameter to set the login method, password and key of the instance, or keep the original login settings of the image. If it's not specified, the user needs to set the login password using the "Reset password" option in the CVM console or calling the API `ResetInstancesPassword` to complete the creation of the CVM instance(s).
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response. If this parameter is not specified, the instance will be associated with default security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default. However, for custom images and images from the marketplace, Anti-DDoS and Cloud Monitor are not enabled by default. The original services in the image will be retained.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Instance hostname. <br><li>Dots (.) and dashes (-) can not be used as the first or last character of HostName nor used consecutively. <br<li>Windows instances: 2 to 15 characters, including English letters (case-insensitive), numbers and dashes (-). Dots and numeric-only names are not allowed. <br><li>Other instances (Linux, etc.): 2 to 60 characters, including English letters (case-insensitive), numbers, dashes (-) and dots. Note that consecutive dots are not allowed.<br><li>Batch naming: use `{R:x}` for batch naming if multiple instances are purchased. `x` is the serial number of the instance. It’s generated by [x, x+n-1], where `n` refers to the number of instances purchased. For example, if `server{R:3}` is input, if you purchase one instance, the hostname is `server3`. If you purchase two instances, the hostnames are `server3` and `server4` respectively. You can specify multiple pattern strings `{R:x}`. </li><br><li>Purchasing multiple instances: If no pattern string is specified, you shall add suffixes `1, 2...n` to the instance hostname. `n` represents the number of purchased instances. For example, if the instance hostname is `server`,  when two instances are purchased, the hostnames of instances purchased are respectively `server1` and `server2`.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Scheduled tasks. You can use this parameter to specify scheduled tasks for the instance. Only scheduled termination is supported.
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// Placement group ID. You can only specify one.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// List of tag description. By specifying this parameter, the tag can be bound to the corresponding CVM and CBS instances at the same time.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// The market options of the instance.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16 KB. For more information on how to get the value of this parameter, see the commands you need to execute on startup for [Windows](https://intl.cloud.tencent.com/document/product/213/17526) or [Linux](https://intl.cloud.tencent.com/document/product/213/17525).
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Whether the request is a dry run only.
	// `true`: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// `false` (default value): Send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// CAM role name, which can be obtained from the `roleName` field in the response of the [`DescribeRoleList`](https://intl.cloud.tencent.com/document/product/598/36223?from_cn_redirect=1) API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`

	// Instance launch template.
	LaunchTemplate *LaunchTemplate `json:"LaunchTemplate,omitempty" name:"LaunchTemplate"`

	// Specify the ID of the dedicated cluster where the CVM is created.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`

	// Specify the CHC physical server that used to create the CHC CVM.
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// Whether the termination protection is enabled. Values: <br><li>`TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.<br><li>`FALSE`: Do not enable the instance protection.<br><br>Default value: `FALSE`.
	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance [billing type](https://intl.cloud.tencent.com/document/product/213/2180?from_cn_redirect=1). <br><li>`POSTPAID_BY_HOUR`: Hourly-based pay-as-you-go <br><li>`CDHPAID`: Dedicated CVM (associated with a dedicated host. Resource usage of the dedicated host is free of charge.) <br><li>`SPOTPAID`: [Spot instance](https://intl.cloud.tencent.com/document/product/213/17817)<br>Default value: `POSTPAID_BY_HOUR`.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// Details of the monthly subscription, including the purchase period, auto-renewal. It is required if the `InstanceChargeType` is `PREPAID`.
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// Location of the instance. You can use this parameter to specify the availability zone, project, and CDH (for dedicated CVMs).
	//  <b>Note: `Placement` is required when `LaunchTemplate` is not specified. If both the parameters are passed in, `Placement` prevails.</b>
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// The instance model. 
	// <br><li>To view specific values for `POSTPAID_BY_HOUR` instances, you can call [DescribeInstanceTypeConfigs](https://intl.cloud.tencent.com/document/api/213/15749?from_cn_redirect=1) or refer to [Instance Types](https://intl.cloud.tencent.com/document/product/213/11518?from_cn_redirect=1). If this parameter is not specified, `S1.SMALL1` will be used by default.<br><li>For `CDHPAID` instances, the value of this parameter is in the format of `CDH_XCXG` based on the number of CPU cores and memory capacity. For example, if you want to create a CDH instance with a single-core CPU and 1 GB memory, specify this parameter as `CDH_1C1G`.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The [image](https://intl.cloud.tencent.com/document/product/213/4940?from_cn_redirect=1) ID in the format of `img-xxx`. There are three types of images:<br/><li>Public images</li><li>Custom images</li><li>Shared images</li><br/>To check the image ID:<br/><li>For public images, custom images, and shared images, go to the [CVM console](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE). </li><li>Call [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1), pass in `InstanceType` to retrieve the list of images supported by the current model, and then find the `ImageId` in the response.</li>
	//  <b>Note: `ImageId` is required when `LaunchTemplate` is not specified. If both the parameters are passed in, `ImageId` prevails.</b>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// System disk configuration of the instance. If this parameter is not specified, the default value will be used.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// The configuration information of instance data disks. If this parameter is not specified, no data disk will be purchased by default. When purchasing, you can specify 21 data disks, which can contain at most 1 LOCAL_BASIC or LOCAL_SSD data disk, and at most 20 CLOUD_BASIC, CLOUD_PREMIUM, or CLOUD_SSD data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// Configuration information of VPC. This parameter is used to specify VPC ID and subnet ID, etc. If a VPC IP is specified in this parameter, it indicates the primary ENI IP of each instance. The value of parameter InstanceCount must be the same as the number of VPC IPs, which cannot be greater than 20.
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// Configuration of public network bandwidth. If this parameter is not specified, 0 Mbps will be used by default.
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// The number of instances to be purchased. Value range for pay-as-you-go instances: [1, 100]. Default value: `1`. The specified number of instances to be purchased cannot exceed the remaining quota allowed for the user. For more information on the quota, see [Quota for CVM Instances](https://intl.cloud.tencent.com/document/product/213/2664).
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// Instance name to be displayed. <br><li>If this parameter is not specified, "Unnamed" will be displayed by default. </li><li>If you purchase multiple instances at the same time and specify a pattern string `{R:x}`, numbers `[x, x+n-1]` will be generated, where `n` represents the number of instances purchased. For example, you specify a pattern string, `server_{R:3}`. If you only purchase 1 instance, the instance will be named `server_3`; if you purchase 2, they will be named `server_3` and `server_4`. You can specify multiple pattern strings in the format of `{R:x}`. </li><li>If you purchase multiple instances at the same time and do not specify a pattern string, the instance names will be suffixed by `1, 2...n`, where `n` represents the number of instances purchased. For example, if you purchase 2 instances and the instance name body is `server_`, the instance names will be `server_1` and `server_2`. </li><li>This parameter can contain up to 60 characters, including the pattern string.
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// Instance login settings. You can use this parameter to set the login method, password and key of the instance, or keep the original login settings of the image. If it's not specified, the user needs to set the login password using the "Reset password" option in the CVM console or calling the API `ResetInstancesPassword` to complete the creation of the CVM instance(s).
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// Security groups to which the instance belongs. To obtain the security group IDs, you can call [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808) and look for the `sgld` fields in the response. If this parameter is not specified, the instance will be associated with default security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// Enhanced service. You can use this parameter to specify whether to enable services such as Anti-DDoS and Cloud Monitor. If this parameter is not specified, Cloud Monitor and Anti-DDoS are enabled for public images by default. However, for custom images and images from the marketplace, Anti-DDoS and Cloud Monitor are not enabled by default. The original services in the image will be retained.
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// Instance hostname. <br><li>Dots (.) and dashes (-) can not be used as the first or last character of HostName nor used consecutively. <br<li>Windows instances: 2 to 15 characters, including English letters (case-insensitive), numbers and dashes (-). Dots and numeric-only names are not allowed. <br><li>Other instances (Linux, etc.): 2 to 60 characters, including English letters (case-insensitive), numbers, dashes (-) and dots. Note that consecutive dots are not allowed.<br><li>Batch naming: use `{R:x}` for batch naming if multiple instances are purchased. `x` is the serial number of the instance. It’s generated by [x, x+n-1], where `n` refers to the number of instances purchased. For example, if `server{R:3}` is input, if you purchase one instance, the hostname is `server3`. If you purchase two instances, the hostnames are `server3` and `server4` respectively. You can specify multiple pattern strings `{R:x}`. </li><br><li>Purchasing multiple instances: If no pattern string is specified, you shall add suffixes `1, 2...n` to the instance hostname. `n` represents the number of purchased instances. For example, if the instance hostname is `server`,  when two instances are purchased, the hostnames of instances purchased are respectively `server1` and `server2`.
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// Scheduled tasks. You can use this parameter to specify scheduled tasks for the instance. Only scheduled termination is supported.
	ActionTimer *ActionTimer `json:"ActionTimer,omitempty" name:"ActionTimer"`

	// Placement group ID. You can only specify one.
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// List of tag description. By specifying this parameter, the tag can be bound to the corresponding CVM and CBS instances at the same time.
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// The market options of the instance.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// User data provided to the instance. This parameter needs to be encoded in base64 format with the maximum size of 16 KB. For more information on how to get the value of this parameter, see the commands you need to execute on startup for [Windows](https://intl.cloud.tencent.com/document/product/213/17526) or [Linux](https://intl.cloud.tencent.com/document/product/213/17525).
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// Whether the request is a dry run only.
	// `true`: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available.
	// If the dry run fails, the corresponding error code will be returned.
	// If the dry run succeeds, the RequestId will be returned.
	// `false` (default value): Send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// CAM role name, which can be obtained from the `roleName` field in the response of the [`DescribeRoleList`](https://intl.cloud.tencent.com/document/product/598/36223?from_cn_redirect=1) API.
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// HPC cluster ID. The HPC cluster must and can only be specified for a high-performance computing instance.
	HpcClusterId *string `json:"HpcClusterId,omitempty" name:"HpcClusterId"`

	// Instance launch template.
	LaunchTemplate *LaunchTemplate `json:"LaunchTemplate,omitempty" name:"LaunchTemplate"`

	// Specify the ID of the dedicated cluster where the CVM is created.
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`

	// Specify the CHC physical server that used to create the CHC CVM.
	ChcIds []*string `json:"ChcIds,omitempty" name:"ChcIds"`

	// Whether the termination protection is enabled. Values: <br><li>`TRUE`: Enable instance protection, which means that this instance can not be deleted by an API action.<br><li>`FALSE`: Do not enable the instance protection.<br><br>Default value: `FALSE`.
	DisableApiTermination *bool `json:"DisableApiTermination,omitempty" name:"DisableApiTermination"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "Placement")
	delete(f, "InstanceType")
	delete(f, "ImageId")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "VirtualPrivateCloud")
	delete(f, "InternetAccessible")
	delete(f, "InstanceCount")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "ClientToken")
	delete(f, "HostName")
	delete(f, "ActionTimer")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "TagSpecification")
	delete(f, "InstanceMarketOptions")
	delete(f, "UserData")
	delete(f, "DryRun")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "LaunchTemplate")
	delete(f, "DedicatedClusterId")
	delete(f, "ChcIds")
	delete(f, "DisableApiTermination")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesResponseParams struct {
	// If you use this API to create instance(s), this parameter will be returned, representing one or more instance IDs. Retuning the instance ID list does not necessarily mean that the instance(s) were created successfully. To check whether the instance(s) were created successfully, you can call [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and check the status of the instances in `InstancesSet` in the response. If the status of an instance changes from "PENDING" to "RUNNING", it means that the instance has been created successfully.
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RunInstancesResponseParams `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {
	// Whether to enable [Cloud Monitor](https://intl.cloud.tencent.com/document/product/248?from_cn_redirect=1). Valid values: <br><li>TRUE: enable Cloud Monitor <br><li>FALSE: do not enable Cloud Monitor <br><br>Default value: TRUE.
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// Whether to enable [Cloud Security](https://intl.cloud.tencent.com/document/product/296?from_cn_redirect=1). Valid values: <br><li>TRUE: enable Cloud Security <br><li>FALSE: do not enable Cloud Security <br><br>Default value: TRUE.
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

// Predefined struct for user
type StartInstancesRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
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

// Predefined struct for user
type StartInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// (Disused. Please use `StopType` instead.) Whether to forcibly shut down an instance after a normal shutdown fails. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default value: `FALSE`. 
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// Instance shutdown mode. Valid values: <br><li>SOFT_FIRST: perform a soft shutdown first, and force shut down the instance if the soft shutdown fails <br><li>HARD: force shut down the instance directly <br><li>SOFT: soft shutdown only <br>Default value: SOFT.
	StopType *string `json:"StopType,omitempty" name:"StopType"`

	// Billing method of a pay-as-you-go instance after shutdown.
	// Valid values: <br><li>KEEP_CHARGING: billing continues after shutdown <br><li>STOP_CHARGING: billing stops after shutdown <br>Default value: KEEP_CHARGING.
	// This parameter is only valid for some pay-as-you-go instances using cloud disks. For more information, see [No charges when shut down for pay-as-you-go instances](https://intl.cloud.tencent.com/document/product/213/19918).
	StoppedMode *string `json:"StoppedMode,omitempty" name:"StoppedMode"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// (Disused. Please use `StopType` instead.) Whether to forcibly shut down an instance after a normal shutdown fails. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default value: `FALSE`. 
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ForceStop")
	delete(f, "StopType")
	delete(f, "StoppedMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type SyncImagesRequestParams struct {
	// List of image IDs. You can obtain the image IDs in two ways: <br><li>Call the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) API and find the value of `ImageId` in the response. <br><li>Find the image IDs in the [Image Console](https://console.cloud.tencent.com/cvm/image). <br>The specified images must meet the following requirement: <br><li>the images must be in the `NORMAL` state. <br>For more information on image status, see [Image Data Table](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#Image).
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// List of destination regions for synchronization. A destination region must meet the following requirements: <br><li>It cannot be the source region. <br><li>It must be valid. <br><li>Currently some regions do not support image synchronization. <br>For specific regions, see [Region](https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1).
	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`

	// Checks whether image synchronization can be initiated 
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest
	
	// List of image IDs. You can obtain the image IDs in two ways: <br><li>Call the [DescribeImages](https://intl.cloud.tencent.com/document/api/213/15715?from_cn_redirect=1) API and find the value of `ImageId` in the response. <br><li>Find the image IDs in the [Image Console](https://console.cloud.tencent.com/cvm/image). <br>The specified images must meet the following requirement: <br><li>the images must be in the `NORMAL` state. <br>For more information on image status, see [Image Data Table](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#Image).
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// List of destination regions for synchronization. A destination region must meet the following requirements: <br><li>It cannot be the source region. <br><li>It must be valid. <br><li>Currently some regions do not support image synchronization. <br>For specific regions, see [Region](https://intl.cloud.tencent.com/document/product/213/6091?from_cn_redirect=1).
	DestinationRegions []*string `json:"DestinationRegions,omitempty" name:"DestinationRegions"`

	// Checks whether image synchronization can be initiated 
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SyncImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "DestinationRegions")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncImagesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse
	Response *SyncImagesResponseParams `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// System disk type. For more information about the limits of system disk types, please see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952?from_cn_redirect=1). Valid values:<br><li>LOCAL_BASIC: local disk<br><li>LOCAL_SSD: local SSD disk<br><li>CLOUD_BASIC: HDD cloud disk<br><li>CLOUD_SSD: SSD cloud disk<br><li>CLOUD_PREMIUM: Premium cloud storage<br><li>CLOUD_BSSD: Balanced SSD<br><br>The disk currently in stock will be used by default.
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// System disk ID. System disks whose type is `LOCAL_BASIC` or `LOCAL_SSD` do not have an ID and do not support this parameter.
	// It is only used as a response parameter for APIs such as `DescribeInstances`, and cannot be used as a request parameter for APIs such as `RunInstances`.
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// System disk size; unit: GB; default value: 50 GB.
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ID of the dedicated cluster to which the instance belongs.
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type Tag struct {
	// Tag key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {
	// The type of resource that the tag is bound to. Valid values: `instance` (for CVM), `host` (for CDH), `image` (for image), and `keypair` (for key).
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// List of tags
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID(s). To obtain the instance IDs, you can call [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) and look for `InstanceId` in the response. The maximum number of instances in each request is 100.
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

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type VirtualPrivateCloud struct {
	// VPC ID in the format of `vpc-xxx`. To obtain valid VPC IDs, you can log in to the [console](https://console.cloud.tencent.com/vpc/vpc?rid=1) or call the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API and look for the `unVpcId` fields in the response. If you specify `DEFAULT` for both `VpcId` and `SubnetId` when creating an instance, the default VPC will be used.
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC subnet ID in the format `subnet-xxx`. To obtain valid subnet IDs, you can log in to the [console](https://console.cloud.tencent.com/vpc/subnet?rid=1) or call [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) and look for the `unSubnetId` fields in the response. If you specify `DEFAULT` for both `SubnetId` and `VpcId` when creating an instance, the default VPC will be used.
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Whether to use a CVM instance as a public gateway. The public gateway is only available when the instance has a public IP and resides in a VPC. Valid values: <br><li>`TRUE`: yes;<br><li>`FALSE`: no<br><br>Default: `FALSE`.
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// Array of VPC subnet IPs. You can use this parameter when creating instances or modifying VPC attributes of instances. Currently you can specify multiple IPs in one subnet only when creating multiple instances at the same time.
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// Number of IPv6 addresses randomly generated for the ENI.
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type ZoneInfo struct {
	// Availability zone name, such as `ap-guangzhou-3`.
	// The following is a list of all availability zones:
	// <li> ap-chongqing-1 </li>
	// <li> ap-seoul-1 </li>
	// <li> ap-seoul-2 </li>
	// <li> ap-chengdu-1 </li>
	// <li> ap-chengdu-2 </li>
	// <li> ap-hongkong-1 (sold out)</li>
	// <li> ap-hongkong-2 </li>
	// <li> ap-hongkong-3 </li>
	// <li> ap-shenzhen-fsi-1 </li>
	// <li> ap-shenzhen-fsi-2 </li>
	// <li> ap-shenzhen-fsi-3 </li>
	// <li> ap-guangzhou-1 (sold out)</li>
	// <li> ap-guangzhou-2 (sold out)</li>
	// <li> ap-guangzhou-3 </li>
	// <li> ap-guangzhou-4 </li>
	// <li> ap-guangzhou-6 </li>
	// <li> ap-guangzhou-7 </li>
	// <li> ap-tokyo-1 </li>
	// <li> ap-tokyo-2 </li>
	// <li> ap-singapore-1 </li>
	// <li> ap-singapore-2 </li>
	// <li> ap-singapore-3 </li>
	// <li> ap-shanghai-fsi-1 </li>
	// <li> ap-shanghai-fsi-2 </li>
	// <li> ap-shanghai-fsi-3 </li>
	// <li> ap-bangkok-1 </li>
	// <li> ap-bangkok-2 </li>
	// <li> ap-shanghai-1 (sold out) </li>
	// <li> ap-shanghai-2 </li>
	// <li> ap-shanghai-3 </li>
	// <li> ap-shanghai-4 </li>
	// <li> ap-shanghai-5 </li>
	// <li> ap-mumbai-1 </li>
	// <li> ap-mumbai-2 </li>
	// <li> eu-moscow-1 </li>
	// <li> ap-beijing-1 (sold out) </li>
	// <li> ap-beijing-2 </li>
	// <li> ap-beijing-3 </li>
	// <li> ap-beijing-4 </li>
	// <li> ap-beijing-5 </li>
	// <li> ap-beijing-6 </li>
	// <li> ap-beijing-7 </li>
	// <li> na-siliconvalley-1 </li>
	// <li> na-siliconvalley-2 </li>
	// <li> eu-frankfurt-1 </li>
	// <li> eu-frankfurt-2 </li>
	// <li> na-toronto-1 </li>
	// <li> na-ashburn-1 </li>
	// <li> na-ashburn-2 </li>
	// <li> ap-nanjing-1 </li>
	// <li> ap-nanjing-2 </li>
	// <li> sa-saopaulo-1</li>
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Availability zone description, such as Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Availability zone ID.
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Availability zone status. Valid values: `AVAILABLE`: available; `UNAVAILABLE`: unavailable.
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`
}